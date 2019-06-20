/*
 * Copyright 2019 Open Networking Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package partitionset

import (
	"context"
	"github.com/atomix/atomix-k8s-controller/pkg/apis/k8s/v1alpha1"
	"github.com/atomix/atomix-k8s-controller/pkg/controller/protocol"
	k8sutil "github.com/atomix/atomix-k8s-controller/pkg/controller/util/k8s"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_partitiongroup")

// AddController creates a new PartitionSet controller and adds it to the Manager. The Manager will set fields on the
// controller and Start it when the Manager is Started.
func Add(mgr manager.Manager, protocols *protocol.ProtocolManager) error {
	r := &PartitionSetReconciler{
		client: mgr.GetClient(),
		scheme: mgr.GetScheme(),
		config: mgr.GetConfig(),
		protocols: protocols,
	}

	// Create a new controller
	c, err := controller.New("partitionset-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource PartitionSet
	err = c.Watch(&source.Kind{Type: &v1alpha1.PartitionSet{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Partitions and requeue the owner PartitionSet
	err = c.Watch(&source.Kind{Type: &v1alpha1.Partition{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &v1alpha1.PartitionSet{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &PartitionSetReconciler{}

// PartitionSetReconciler reconciles a PartitionSet object
type PartitionSetReconciler struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
	config *rest.Config
	protocols *protocol.ProtocolManager
}

// Reconcile reads that state of the partition for a PartitionSet object and makes changes based on the state read
// and what is in the PartitionSet.Spec
func (r *PartitionSetReconciler) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	logger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	logger.Info("Reconciling PartitionSet")

	// Fetch the PartitionSet instance
	set := &v1alpha1.PartitionSet{}
	err := r.client.Get(context.TODO(), request.NamespacedName, set)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	v1alpha1.SetDefaults_PartitionGroup(set)

	if err = r.reconcileService(set); err != nil {
		return reconcile.Result{}, err
	}

	if err = r.reconcileEndpoints(set); err != nil {
		return reconcile.Result{}, err
	}

	for i := 1; i <= set.Spec.Partitions; i++ {
		if err = r.reconcilePartition(set, i); err != nil {
			return reconcile.Result{}, err
		}
		if err = r.reconcileStatus(set, i); err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func (r *PartitionSetReconciler) reconcileService(set *v1alpha1.PartitionSet) error {
	service := &corev1.Service{}
	err := r.client.Get(context.TODO(), k8sutil.GetPartitionSetServiceNamespacedName(set), service)
	if err != nil && errors.IsNotFound(err) {
		err = r.addService(set)
	}
	return err
}

func (r *PartitionSetReconciler) addService(set *v1alpha1.PartitionSet) error {
	log.Info("Creating service", "Name", set.Name, "Namespace", set.Namespace)
	service := k8sutil.NewPartitionSetService(set)
	if err := controllerutil.SetControllerReference(set, service, r.scheme); err != nil {
		return err
	}
	return r.client.Create(context.TODO(), service)
}

func (r *PartitionSetReconciler) reconcileEndpoints(set *v1alpha1.PartitionSet) error {
	endpoints := &corev1.Endpoints{}
	err := r.client.Get(context.TODO(), k8sutil.GetPartitionSetServiceNamespacedName(set), endpoints)
	if err != nil && errors.IsNotFound(err) {
		err = r.addEndpoints(set)
	}
	return err
}

func (r *PartitionSetReconciler) addEndpoints(set *v1alpha1.PartitionSet) error {
	log.Info("Creating endpoints", "Name", set.Name, "Namespace", set.Namespace)
	endpoints := k8sutil.NewPartitionSetEndpoints(set)
	if err := controllerutil.SetControllerReference(set, endpoints, r.scheme); err != nil {
		return err
	}
	return r.client.Create(context.TODO(), endpoints)
}

func (r *PartitionSetReconciler) reconcileStatus(set *v1alpha1.PartitionSet, id int) error {
	partition := &v1alpha1.Partition{}
	err := r.client.Get(context.TODO(), k8sutil.GetPartitionNamespacedName(set, id), partition)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}

	if set.Status.ReadyPartitions != partition.Status.ReadyReplicas {
		log.Info("Updating partition set status", "Name", set.Name, "Namespace", set.Namespace)
		set.Status.ReadyPartitions = partition.Status.ReadyReplicas
		err = r.client.Status().Update(context.TODO(), partition)
	}
	return err
}

func (r *PartitionSetReconciler) reconcilePartition(set *v1alpha1.PartitionSet, id int) error {
	partition := &v1alpha1.Partition{}
	err := r.client.Get(context.TODO(), k8sutil.GetPartitionNamespacedName(set, id), partition)
	if err != nil && errors.IsNotFound(err) {
		err = r.addPartition(set, id)
	}
	return err
}

func (r *PartitionSetReconciler) addPartition(set *v1alpha1.PartitionSet, id int) error {
	log.Info("Creating partition", "Name", k8sutil.GetPartitionName(set, id), "Namespace", set.Namespace)
	partition := k8sutil.NewPartition(set, id)
	if err := controllerutil.SetControllerReference(set, partition, r.scheme); err != nil {
		return err
	}
	return r.client.Create(context.TODO(), partition)
}