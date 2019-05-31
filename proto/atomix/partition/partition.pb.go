// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/partition/partition.proto

package partition

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Storage level
type StorageLevel int32

const (
	StorageLevel_DISK   StorageLevel = 0
	StorageLevel_MAPPED StorageLevel = 1
)

var StorageLevel_name = map[int32]string{
	0: "DISK",
	1: "MAPPED",
}
var StorageLevel_value = map[string]int32{
	"DISK":   0,
	"MAPPED": 1,
}

func (x StorageLevel) String() string {
	return proto.EnumName(StorageLevel_name, int32(x))
}
func (StorageLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{0}
}

// Member group strategy
type MemberGroupStrategy int32

const (
	MemberGroupStrategy_HOST_AWARE MemberGroupStrategy = 0
	MemberGroupStrategy_RACK_AWARE MemberGroupStrategy = 1
	MemberGroupStrategy_ZONE_AWARE MemberGroupStrategy = 2
)

var MemberGroupStrategy_name = map[int32]string{
	0: "HOST_AWARE",
	1: "RACK_AWARE",
	2: "ZONE_AWARE",
}
var MemberGroupStrategy_value = map[string]int32{
	"HOST_AWARE": 0,
	"RACK_AWARE": 1,
	"ZONE_AWARE": 2,
}

func (x MemberGroupStrategy) String() string {
	return proto.EnumName(MemberGroupStrategy_name, int32(x))
}
func (MemberGroupStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{1}
}

// Partition identifier
type PartitionId struct {
	Partition            int32             `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	Group                *PartitionGroupId `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PartitionId) Reset()         { *m = PartitionId{} }
func (m *PartitionId) String() string { return proto.CompactTextString(m) }
func (*PartitionId) ProtoMessage()    {}
func (*PartitionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{0}
}
func (m *PartitionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionId.Unmarshal(m, b)
}
func (m *PartitionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionId.Marshal(b, m, deterministic)
}
func (dst *PartitionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionId.Merge(dst, src)
}
func (m *PartitionId) XXX_Size() int {
	return xxx_messageInfo_PartitionId.Size(m)
}
func (m *PartitionId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionId proto.InternalMessageInfo

func (m *PartitionId) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *PartitionId) GetGroup() *PartitionGroupId {
	if m != nil {
		return m.Group
	}
	return nil
}

// Partition group name
type PartitionGroupId struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace            string   `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionGroupId) Reset()         { *m = PartitionGroupId{} }
func (m *PartitionGroupId) String() string { return proto.CompactTextString(m) }
func (*PartitionGroupId) ProtoMessage()    {}
func (*PartitionGroupId) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{1}
}
func (m *PartitionGroupId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroupId.Unmarshal(m, b)
}
func (m *PartitionGroupId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroupId.Marshal(b, m, deterministic)
}
func (dst *PartitionGroupId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroupId.Merge(dst, src)
}
func (m *PartitionGroupId) XXX_Size() int {
	return xxx_messageInfo_PartitionGroupId.Size(m)
}
func (m *PartitionGroupId) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroupId.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroupId proto.InternalMessageInfo

func (m *PartitionGroupId) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PartitionGroupId) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// Partition group info
type PartitionGroup struct {
	Id                   *PartitionGroupId   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                 *PartitionGroupSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Partitions           []*Partition        `protobuf:"bytes,3,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PartitionGroup) Reset()         { *m = PartitionGroup{} }
func (m *PartitionGroup) String() string { return proto.CompactTextString(m) }
func (*PartitionGroup) ProtoMessage()    {}
func (*PartitionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{2}
}
func (m *PartitionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroup.Unmarshal(m, b)
}
func (m *PartitionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroup.Marshal(b, m, deterministic)
}
func (dst *PartitionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroup.Merge(dst, src)
}
func (m *PartitionGroup) XXX_Size() int {
	return xxx_messageInfo_PartitionGroup.Size(m)
}
func (m *PartitionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroup proto.InternalMessageInfo

func (m *PartitionGroup) GetId() *PartitionGroupId {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PartitionGroup) GetSpec() *PartitionGroupSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *PartitionGroup) GetPartitions() []*Partition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

// Partition info
type Partition struct {
	PartitionId          int32                `protobuf:"varint,1,opt,name=partition_id,json=partitionId,proto3" json:"partition_id,omitempty"`
	Endpoints            []*PartitionEndpoint `protobuf:"bytes,2,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Partition) Reset()         { *m = Partition{} }
func (m *Partition) String() string { return proto.CompactTextString(m) }
func (*Partition) ProtoMessage()    {}
func (*Partition) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{3}
}
func (m *Partition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Partition.Unmarshal(m, b)
}
func (m *Partition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Partition.Marshal(b, m, deterministic)
}
func (dst *Partition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Partition.Merge(dst, src)
}
func (m *Partition) XXX_Size() int {
	return xxx_messageInfo_Partition.Size(m)
}
func (m *Partition) XXX_DiscardUnknown() {
	xxx_messageInfo_Partition.DiscardUnknown(m)
}

var xxx_messageInfo_Partition proto.InternalMessageInfo

func (m *Partition) GetPartitionId() int32 {
	if m != nil {
		return m.PartitionId
	}
	return 0
}

func (m *Partition) GetEndpoints() []*PartitionEndpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

// Partition endpoint
type PartitionEndpoint struct {
	Host                 string   `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PartitionEndpoint) Reset()         { *m = PartitionEndpoint{} }
func (m *PartitionEndpoint) String() string { return proto.CompactTextString(m) }
func (*PartitionEndpoint) ProtoMessage()    {}
func (*PartitionEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{4}
}
func (m *PartitionEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionEndpoint.Unmarshal(m, b)
}
func (m *PartitionEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionEndpoint.Marshal(b, m, deterministic)
}
func (dst *PartitionEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionEndpoint.Merge(dst, src)
}
func (m *PartitionEndpoint) XXX_Size() int {
	return xxx_messageInfo_PartitionEndpoint.Size(m)
}
func (m *PartitionEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionEndpoint proto.InternalMessageInfo

func (m *PartitionEndpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *PartitionEndpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

// Partition group specification
type PartitionGroupSpec struct {
	Replicas      uint32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	Partitions    uint32 `protobuf:"varint,2,opt,name=partitions,proto3" json:"partitions,omitempty"`
	PartitionSize uint32 `protobuf:"varint,3,opt,name=partition_size,json=partitionSize,proto3" json:"partition_size,omitempty"`
	// Types that are valid to be assigned to Group:
	//	*PartitionGroupSpec_Raft
	//	*PartitionGroupSpec_PrimaryBackup
	//	*PartitionGroupSpec_Log
	Group                isPartitionGroupSpec_Group `protobuf_oneof:"group"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PartitionGroupSpec) Reset()         { *m = PartitionGroupSpec{} }
func (m *PartitionGroupSpec) String() string { return proto.CompactTextString(m) }
func (*PartitionGroupSpec) ProtoMessage()    {}
func (*PartitionGroupSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{5}
}
func (m *PartitionGroupSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartitionGroupSpec.Unmarshal(m, b)
}
func (m *PartitionGroupSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartitionGroupSpec.Marshal(b, m, deterministic)
}
func (dst *PartitionGroupSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartitionGroupSpec.Merge(dst, src)
}
func (m *PartitionGroupSpec) XXX_Size() int {
	return xxx_messageInfo_PartitionGroupSpec.Size(m)
}
func (m *PartitionGroupSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PartitionGroupSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PartitionGroupSpec proto.InternalMessageInfo

func (m *PartitionGroupSpec) GetReplicas() uint32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *PartitionGroupSpec) GetPartitions() uint32 {
	if m != nil {
		return m.Partitions
	}
	return 0
}

func (m *PartitionGroupSpec) GetPartitionSize() uint32 {
	if m != nil {
		return m.PartitionSize
	}
	return 0
}

type isPartitionGroupSpec_Group interface {
	isPartitionGroupSpec_Group()
}

type PartitionGroupSpec_Raft struct {
	Raft *RaftPartitionGroup `protobuf:"bytes,4,opt,name=raft,proto3,oneof"`
}

type PartitionGroupSpec_PrimaryBackup struct {
	PrimaryBackup *PrimaryBackupPartitionGroup `protobuf:"bytes,5,opt,name=primary_backup,json=primaryBackup,proto3,oneof"`
}

type PartitionGroupSpec_Log struct {
	Log *DistributedLogPartitionGroup `protobuf:"bytes,6,opt,name=log,proto3,oneof"`
}

func (*PartitionGroupSpec_Raft) isPartitionGroupSpec_Group() {}

func (*PartitionGroupSpec_PrimaryBackup) isPartitionGroupSpec_Group() {}

func (*PartitionGroupSpec_Log) isPartitionGroupSpec_Group() {}

func (m *PartitionGroupSpec) GetGroup() isPartitionGroupSpec_Group {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *PartitionGroupSpec) GetRaft() *RaftPartitionGroup {
	if x, ok := m.GetGroup().(*PartitionGroupSpec_Raft); ok {
		return x.Raft
	}
	return nil
}

func (m *PartitionGroupSpec) GetPrimaryBackup() *PrimaryBackupPartitionGroup {
	if x, ok := m.GetGroup().(*PartitionGroupSpec_PrimaryBackup); ok {
		return x.PrimaryBackup
	}
	return nil
}

func (m *PartitionGroupSpec) GetLog() *DistributedLogPartitionGroup {
	if x, ok := m.GetGroup().(*PartitionGroupSpec_Log); ok {
		return x.Log
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PartitionGroupSpec) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _PartitionGroupSpec_OneofMarshaler, _PartitionGroupSpec_OneofUnmarshaler, _PartitionGroupSpec_OneofSizer, []interface{}{
		(*PartitionGroupSpec_Raft)(nil),
		(*PartitionGroupSpec_PrimaryBackup)(nil),
		(*PartitionGroupSpec_Log)(nil),
	}
}

func _PartitionGroupSpec_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PartitionGroupSpec)
	// group
	switch x := m.Group.(type) {
	case *PartitionGroupSpec_Raft:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Raft); err != nil {
			return err
		}
	case *PartitionGroupSpec_PrimaryBackup:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrimaryBackup); err != nil {
			return err
		}
	case *PartitionGroupSpec_Log:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PartitionGroupSpec.Group has unexpected type %T", x)
	}
	return nil
}

func _PartitionGroupSpec_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PartitionGroupSpec)
	switch tag {
	case 4: // group.raft
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RaftPartitionGroup)
		err := b.DecodeMessage(msg)
		m.Group = &PartitionGroupSpec_Raft{msg}
		return true, err
	case 5: // group.primary_backup
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrimaryBackupPartitionGroup)
		err := b.DecodeMessage(msg)
		m.Group = &PartitionGroupSpec_PrimaryBackup{msg}
		return true, err
	case 6: // group.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DistributedLogPartitionGroup)
		err := b.DecodeMessage(msg)
		m.Group = &PartitionGroupSpec_Log{msg}
		return true, err
	default:
		return false, nil
	}
}

func _PartitionGroupSpec_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*PartitionGroupSpec)
	// group
	switch x := m.Group.(type) {
	case *PartitionGroupSpec_Raft:
		s := proto.Size(x.Raft)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PartitionGroupSpec_PrimaryBackup:
		s := proto.Size(x.PrimaryBackup)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *PartitionGroupSpec_Log:
		s := proto.Size(x.Log)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Raft partition group spec
type RaftPartitionGroup struct {
	ElectionTimeout      *duration.Duration `protobuf:"bytes,1,opt,name=election_timeout,json=electionTimeout,proto3" json:"election_timeout,omitempty"`
	HeartbeatInterval    *duration.Duration `protobuf:"bytes,2,opt,name=heartbeat_interval,json=heartbeatInterval,proto3" json:"heartbeat_interval,omitempty"`
	Storage              *StorageSpec       `protobuf:"bytes,3,opt,name=storage,proto3" json:"storage,omitempty"`
	Compaction           *CompactionSpec    `protobuf:"bytes,4,opt,name=compaction,proto3" json:"compaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RaftPartitionGroup) Reset()         { *m = RaftPartitionGroup{} }
func (m *RaftPartitionGroup) String() string { return proto.CompactTextString(m) }
func (*RaftPartitionGroup) ProtoMessage()    {}
func (*RaftPartitionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{6}
}
func (m *RaftPartitionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RaftPartitionGroup.Unmarshal(m, b)
}
func (m *RaftPartitionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RaftPartitionGroup.Marshal(b, m, deterministic)
}
func (dst *RaftPartitionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaftPartitionGroup.Merge(dst, src)
}
func (m *RaftPartitionGroup) XXX_Size() int {
	return xxx_messageInfo_RaftPartitionGroup.Size(m)
}
func (m *RaftPartitionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_RaftPartitionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_RaftPartitionGroup proto.InternalMessageInfo

func (m *RaftPartitionGroup) GetElectionTimeout() *duration.Duration {
	if m != nil {
		return m.ElectionTimeout
	}
	return nil
}

func (m *RaftPartitionGroup) GetHeartbeatInterval() *duration.Duration {
	if m != nil {
		return m.HeartbeatInterval
	}
	return nil
}

func (m *RaftPartitionGroup) GetStorage() *StorageSpec {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *RaftPartitionGroup) GetCompaction() *CompactionSpec {
	if m != nil {
		return m.Compaction
	}
	return nil
}

// Primary-backup partition group spec
type PrimaryBackupPartitionGroup struct {
	MemberGroupStrategy  MemberGroupStrategy `protobuf:"varint,1,opt,name=member_group_strategy,json=memberGroupStrategy,proto3,enum=atomix.partition.MemberGroupStrategy" json:"member_group_strategy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *PrimaryBackupPartitionGroup) Reset()         { *m = PrimaryBackupPartitionGroup{} }
func (m *PrimaryBackupPartitionGroup) String() string { return proto.CompactTextString(m) }
func (*PrimaryBackupPartitionGroup) ProtoMessage()    {}
func (*PrimaryBackupPartitionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{7}
}
func (m *PrimaryBackupPartitionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimaryBackupPartitionGroup.Unmarshal(m, b)
}
func (m *PrimaryBackupPartitionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimaryBackupPartitionGroup.Marshal(b, m, deterministic)
}
func (dst *PrimaryBackupPartitionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimaryBackupPartitionGroup.Merge(dst, src)
}
func (m *PrimaryBackupPartitionGroup) XXX_Size() int {
	return xxx_messageInfo_PrimaryBackupPartitionGroup.Size(m)
}
func (m *PrimaryBackupPartitionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimaryBackupPartitionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_PrimaryBackupPartitionGroup proto.InternalMessageInfo

func (m *PrimaryBackupPartitionGroup) GetMemberGroupStrategy() MemberGroupStrategy {
	if m != nil {
		return m.MemberGroupStrategy
	}
	return MemberGroupStrategy_HOST_AWARE
}

// Log partition group spec
type DistributedLogPartitionGroup struct {
	MemberGroupStrategy  MemberGroupStrategy `protobuf:"varint,1,opt,name=member_group_strategy,json=memberGroupStrategy,proto3,enum=atomix.partition.MemberGroupStrategy" json:"member_group_strategy,omitempty"`
	Storage              *StorageSpec        `protobuf:"bytes,2,opt,name=storage,proto3" json:"storage,omitempty"`
	Compaction           *CompactionSpec     `protobuf:"bytes,3,opt,name=compaction,proto3" json:"compaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DistributedLogPartitionGroup) Reset()         { *m = DistributedLogPartitionGroup{} }
func (m *DistributedLogPartitionGroup) String() string { return proto.CompactTextString(m) }
func (*DistributedLogPartitionGroup) ProtoMessage()    {}
func (*DistributedLogPartitionGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{8}
}
func (m *DistributedLogPartitionGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedLogPartitionGroup.Unmarshal(m, b)
}
func (m *DistributedLogPartitionGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedLogPartitionGroup.Marshal(b, m, deterministic)
}
func (dst *DistributedLogPartitionGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedLogPartitionGroup.Merge(dst, src)
}
func (m *DistributedLogPartitionGroup) XXX_Size() int {
	return xxx_messageInfo_DistributedLogPartitionGroup.Size(m)
}
func (m *DistributedLogPartitionGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedLogPartitionGroup.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedLogPartitionGroup proto.InternalMessageInfo

func (m *DistributedLogPartitionGroup) GetMemberGroupStrategy() MemberGroupStrategy {
	if m != nil {
		return m.MemberGroupStrategy
	}
	return MemberGroupStrategy_HOST_AWARE
}

func (m *DistributedLogPartitionGroup) GetStorage() *StorageSpec {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *DistributedLogPartitionGroup) GetCompaction() *CompactionSpec {
	if m != nil {
		return m.Compaction
	}
	return nil
}

// Partition group storage configuration
type StorageSpec struct {
	MaxEntrySize         uint32       `protobuf:"varint,1,opt,name=max_entry_size,json=maxEntrySize,proto3" json:"max_entry_size,omitempty"`
	SegmentSize          uint32       `protobuf:"varint,2,opt,name=segment_size,json=segmentSize,proto3" json:"segment_size,omitempty"`
	Level                StorageLevel `protobuf:"varint,3,opt,name=level,proto3,enum=atomix.partition.StorageLevel" json:"level,omitempty"`
	FlushOnCommit        bool         `protobuf:"varint,4,opt,name=flush_on_commit,json=flushOnCommit,proto3" json:"flush_on_commit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StorageSpec) Reset()         { *m = StorageSpec{} }
func (m *StorageSpec) String() string { return proto.CompactTextString(m) }
func (*StorageSpec) ProtoMessage()    {}
func (*StorageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{9}
}
func (m *StorageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageSpec.Unmarshal(m, b)
}
func (m *StorageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageSpec.Marshal(b, m, deterministic)
}
func (dst *StorageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageSpec.Merge(dst, src)
}
func (m *StorageSpec) XXX_Size() int {
	return xxx_messageInfo_StorageSpec.Size(m)
}
func (m *StorageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_StorageSpec proto.InternalMessageInfo

func (m *StorageSpec) GetMaxEntrySize() uint32 {
	if m != nil {
		return m.MaxEntrySize
	}
	return 0
}

func (m *StorageSpec) GetSegmentSize() uint32 {
	if m != nil {
		return m.SegmentSize
	}
	return 0
}

func (m *StorageSpec) GetLevel() StorageLevel {
	if m != nil {
		return m.Level
	}
	return StorageLevel_DISK
}

func (m *StorageSpec) GetFlushOnCommit() bool {
	if m != nil {
		return m.FlushOnCommit
	}
	return false
}

// Partition group compaction configuration
type CompactionSpec struct {
	Dynamic              bool     `protobuf:"varint,1,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	FreeDiskBuffer       float64  `protobuf:"fixed64,2,opt,name=free_disk_buffer,json=freeDiskBuffer,proto3" json:"free_disk_buffer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompactionSpec) Reset()         { *m = CompactionSpec{} }
func (m *CompactionSpec) String() string { return proto.CompactTextString(m) }
func (*CompactionSpec) ProtoMessage()    {}
func (*CompactionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_partition_e6b455c56c47317c, []int{10}
}
func (m *CompactionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompactionSpec.Unmarshal(m, b)
}
func (m *CompactionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompactionSpec.Marshal(b, m, deterministic)
}
func (dst *CompactionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompactionSpec.Merge(dst, src)
}
func (m *CompactionSpec) XXX_Size() int {
	return xxx_messageInfo_CompactionSpec.Size(m)
}
func (m *CompactionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_CompactionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_CompactionSpec proto.InternalMessageInfo

func (m *CompactionSpec) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *CompactionSpec) GetFreeDiskBuffer() float64 {
	if m != nil {
		return m.FreeDiskBuffer
	}
	return 0
}

func init() {
	proto.RegisterType((*PartitionId)(nil), "atomix.partition.PartitionId")
	proto.RegisterType((*PartitionGroupId)(nil), "atomix.partition.PartitionGroupId")
	proto.RegisterType((*PartitionGroup)(nil), "atomix.partition.PartitionGroup")
	proto.RegisterType((*Partition)(nil), "atomix.partition.Partition")
	proto.RegisterType((*PartitionEndpoint)(nil), "atomix.partition.PartitionEndpoint")
	proto.RegisterType((*PartitionGroupSpec)(nil), "atomix.partition.PartitionGroupSpec")
	proto.RegisterType((*RaftPartitionGroup)(nil), "atomix.partition.RaftPartitionGroup")
	proto.RegisterType((*PrimaryBackupPartitionGroup)(nil), "atomix.partition.PrimaryBackupPartitionGroup")
	proto.RegisterType((*DistributedLogPartitionGroup)(nil), "atomix.partition.DistributedLogPartitionGroup")
	proto.RegisterType((*StorageSpec)(nil), "atomix.partition.StorageSpec")
	proto.RegisterType((*CompactionSpec)(nil), "atomix.partition.CompactionSpec")
	proto.RegisterEnum("atomix.partition.StorageLevel", StorageLevel_name, StorageLevel_value)
	proto.RegisterEnum("atomix.partition.MemberGroupStrategy", MemberGroupStrategy_name, MemberGroupStrategy_value)
}

func init() {
	proto.RegisterFile("atomix/partition/partition.proto", fileDescriptor_partition_e6b455c56c47317c)
}

var fileDescriptor_partition_e6b455c56c47317c = []byte{
	// 832 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x51, 0x6e, 0xdb, 0x46,
	0x10, 0x35, 0x29, 0xc9, 0xb6, 0x46, 0x36, 0xa3, 0x6c, 0x50, 0x94, 0xb1, 0x53, 0x43, 0x61, 0x9d,
	0xc2, 0x08, 0x50, 0x19, 0x50, 0x0b, 0x34, 0x68, 0x7e, 0x2a, 0x59, 0x42, 0x6d, 0x24, 0xa9, 0x8d,
	0x95, 0xd1, 0xa2, 0xfd, 0x21, 0x56, 0xe4, 0x4a, 0x5e, 0x58, 0xe4, 0xb2, 0xcb, 0x65, 0x60, 0xe5,
	0x32, 0xbd, 0x44, 0xff, 0x7a, 0x91, 0x1e, 0xa4, 0x07, 0x28, 0x38, 0xa4, 0x28, 0xca, 0x4c, 0xd4,
	0xa2, 0x45, 0xbe, 0xb4, 0x3b, 0xf3, 0xde, 0xdb, 0x99, 0x37, 0x1c, 0x41, 0x87, 0x69, 0x19, 0x88,
	0xbb, 0xd3, 0x88, 0x29, 0x2d, 0xb4, 0x90, 0xe1, 0xea, 0xd4, 0x8d, 0x94, 0xd4, 0x92, 0xb4, 0x33,
	0x44, 0xb7, 0x88, 0x1f, 0x3c, 0x9e, 0x49, 0x39, 0x9b, 0xf3, 0x53, 0xcc, 0x4f, 0x92, 0xe9, 0x29,
	0x0b, 0x17, 0x19, 0xf8, 0xe0, 0xe8, 0x7e, 0xca, 0x4f, 0x14, 0x5b, 0x89, 0x39, 0x1c, 0x5a, 0x57,
	0x4b, 0x9d, 0x0b, 0x9f, 0x3c, 0x81, 0x66, 0x21, 0x6b, 0x1b, 0x1d, 0xe3, 0xa4, 0x41, 0x57, 0x01,
	0xf2, 0x02, 0x1a, 0x33, 0x25, 0x93, 0xc8, 0x36, 0x3b, 0xc6, 0x49, 0xab, 0xe7, 0x74, 0xef, 0x57,
	0xd2, 0x2d, 0xb4, 0xbe, 0x4f, 0x71, 0x17, 0x3e, 0xcd, 0x08, 0xce, 0x10, 0xda, 0xf7, 0x53, 0x84,
	0x40, 0x3d, 0x64, 0x01, 0xc7, 0x67, 0x9a, 0x14, 0xcf, 0xe9, 0xfb, 0xe9, 0x6f, 0x1c, 0x31, 0x8f,
	0xe3, 0x2b, 0x4d, 0xba, 0x0a, 0x38, 0x7f, 0x18, 0x60, 0xad, 0xcb, 0x90, 0x1e, 0x98, 0xc2, 0x47,
	0x89, 0x7f, 0x57, 0x8f, 0x29, 0x7c, 0xf2, 0x02, 0xea, 0x71, 0xc4, 0xbd, 0xbc, 0x8b, 0xe3, 0x7f,
	0x62, 0x8d, 0x23, 0xee, 0x51, 0x64, 0x90, 0x97, 0x00, 0x05, 0x2a, 0xb6, 0x6b, 0x9d, 0xda, 0x49,
	0xab, 0x77, 0xb8, 0x81, 0x4f, 0x4b, 0x70, 0xe7, 0x57, 0x68, 0x16, 0x09, 0xf2, 0x14, 0xf6, 0x8a,
	0x94, 0x9b, 0x77, 0xd0, 0xa0, 0xad, 0xa8, 0x34, 0x8b, 0x3e, 0x34, 0x79, 0xe8, 0x47, 0x52, 0x84,
	0x3a, 0xb6, 0x4d, 0x7c, 0xeb, 0xf3, 0x0d, 0x6f, 0x8d, 0x72, 0x2c, 0x5d, 0xb1, 0x9c, 0x97, 0xf0,
	0xb0, 0x92, 0x4f, 0x7d, 0xbf, 0x91, 0xb1, 0x5e, 0xfa, 0x9e, 0x9e, 0xd3, 0x58, 0x24, 0x95, 0x46,
	0x4b, 0x1a, 0x14, 0xcf, 0xce, 0x9f, 0x26, 0x90, 0xaa, 0x13, 0xe4, 0x00, 0x76, 0x15, 0x8f, 0xe6,
	0xc2, 0x63, 0x31, 0x4a, 0xec, 0xd3, 0xe2, 0x4e, 0x8e, 0xd6, 0xfc, 0x31, 0x31, 0x5b, 0x8a, 0x90,
	0x67, 0x60, 0xad, 0xba, 0x8e, 0xc5, 0x3b, 0x6e, 0xd7, 0x10, 0xb3, 0x5f, 0x44, 0xc7, 0xe2, 0x1d,
	0x27, 0xdf, 0x42, 0x5d, 0xb1, 0xa9, 0xb6, 0xeb, 0x1f, 0x1a, 0x10, 0x65, 0x53, 0xbd, 0x5e, 0xda,
	0xf9, 0x16, 0x45, 0x0e, 0xf9, 0x11, 0xac, 0x48, 0x89, 0x80, 0xa9, 0x85, 0x3b, 0x61, 0xde, 0x6d,
	0x12, 0xd9, 0x0d, 0x54, 0xf9, 0xf2, 0x3d, 0xd6, 0x65, 0xb8, 0x01, 0xc2, 0x2a, 0x72, 0xfb, 0x51,
	0x39, 0x4d, 0x06, 0x50, 0x9b, 0xcb, 0x99, 0xbd, 0x8d, 0x62, 0xdd, 0xaa, 0xd8, 0x50, 0xc4, 0x5a,
	0x89, 0x49, 0xa2, 0xb9, 0xff, 0x5a, 0xce, 0x2a, 0x6a, 0x29, 0x79, 0xb0, 0x93, 0xef, 0x8f, 0xf3,
	0x9b, 0x09, 0xa4, 0xda, 0x03, 0x19, 0x42, 0x9b, 0xcf, 0xb9, 0x87, 0xee, 0x68, 0x11, 0x70, 0x99,
	0xe8, 0xfc, 0xd3, 0x7e, 0xdc, 0xcd, 0xf6, 0xb8, 0xbb, 0xdc, 0xe3, 0xee, 0x30, 0xdf, 0x63, 0xfa,
	0x60, 0x49, 0xb9, 0xce, 0x18, 0xe4, 0x1c, 0xc8, 0x0d, 0x67, 0x4a, 0x4f, 0x38, 0xd3, 0xae, 0x08,
	0x35, 0x57, 0x6f, 0xd9, 0x3c, 0xff, 0xd8, 0x37, 0xe8, 0x3c, 0x2c, 0x48, 0x17, 0x39, 0x87, 0x7c,
	0x03, 0x3b, 0xb1, 0x96, 0x8a, 0xcd, 0xb2, 0x39, 0xb5, 0x7a, 0x9f, 0x55, 0xfb, 0x1e, 0x67, 0x00,
	0x5c, 0x92, 0x25, 0x9a, 0x7c, 0x07, 0xe0, 0xc9, 0x20, 0x62, 0x58, 0x57, 0x3e, 0xc6, 0x4e, 0x95,
	0x7b, 0x56, 0x60, 0x90, 0x5e, 0xe2, 0x38, 0x77, 0x70, 0xb8, 0x61, 0x3c, 0xe4, 0x67, 0xf8, 0x24,
	0xe0, 0xc1, 0x84, 0x2b, 0x17, 0x0d, 0x75, 0x63, 0xad, 0x98, 0xe6, 0xb3, 0x05, 0xda, 0x65, 0xf5,
	0x9e, 0x55, 0xdf, 0x7a, 0x83, 0xf0, 0xec, 0x33, 0xce, 0xc1, 0xf4, 0x51, 0x50, 0x0d, 0x3a, 0x7f,
	0x19, 0xf0, 0x64, 0xd3, 0x30, 0x3f, 0xe2, 0xdb, 0x65, 0xc3, 0xcd, 0xff, 0x61, 0x78, 0xed, 0x3f,
	0x18, 0xfe, 0xbb, 0x01, 0xad, 0x92, 0x34, 0x39, 0x06, 0x2b, 0x60, 0x77, 0x2e, 0x0f, 0xb5, 0x5a,
	0x64, 0xab, 0x9a, 0x2d, 0xfb, 0x5e, 0xc0, 0xee, 0x46, 0x69, 0x10, 0x37, 0xf5, 0x29, 0xec, 0xc5,
	0x7c, 0x16, 0xf0, 0x50, 0x67, 0x98, 0x6c, 0xe5, 0x5b, 0x79, 0x0c, 0x21, 0x5f, 0x43, 0x63, 0xce,
	0xdf, 0xf2, 0x39, 0x56, 0x65, 0xf5, 0x8e, 0x3e, 0xd8, 0xd1, 0xeb, 0x14, 0x45, 0x33, 0x30, 0xf9,
	0x02, 0x1e, 0x4c, 0xe7, 0x49, 0x7c, 0xe3, 0xca, 0xd0, 0xf5, 0x64, 0x10, 0x88, 0xec, 0xdf, 0x60,
	0x97, 0xee, 0x63, 0xf8, 0x32, 0x3c, 0xc3, 0xa0, 0x73, 0x0d, 0xd6, 0x7a, 0x53, 0xc4, 0x86, 0x1d,
	0x7f, 0x11, 0xb2, 0x40, 0x78, 0x58, 0xf1, 0x2e, 0x5d, 0x5e, 0xc9, 0x09, 0xb4, 0xa7, 0x8a, 0x73,
	0xd7, 0x17, 0xf1, 0xad, 0x3b, 0x49, 0xa6, 0x53, 0xae, 0xb0, 0x60, 0x83, 0x5a, 0x69, 0x7c, 0x28,
	0xe2, 0xdb, 0x01, 0x46, 0x9f, 0x1f, 0xc3, 0x5e, 0xb9, 0x28, 0xb2, 0x0b, 0xf5, 0xe1, 0xc5, 0xf8,
	0x55, 0x7b, 0x8b, 0x00, 0x6c, 0xbf, 0xe9, 0x5f, 0x5d, 0x8d, 0x86, 0x6d, 0xe3, 0xf9, 0x08, 0x1e,
	0xbd, 0x67, 0xb2, 0xc4, 0x02, 0x38, 0xbf, 0x1c, 0x5f, 0xbb, 0xfd, 0x9f, 0xfa, 0x74, 0xd4, 0xde,
	0x4a, 0xef, 0xb4, 0x7f, 0xf6, 0x2a, 0xbf, 0x1b, 0xe9, 0xfd, 0x97, 0xcb, 0x1f, 0x46, 0xf9, 0xdd,
	0x1c, 0x1c, 0xc2, 0xa7, 0x42, 0x2e, 0x5d, 0x61, 0x91, 0x58, 0x39, 0x73, 0x65, 0x4c, 0xb6, 0x71,
	0x51, 0xbf, 0xfa, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x34, 0xaf, 0x79, 0x8d, 0x17, 0x08, 0x00, 0x00,
}