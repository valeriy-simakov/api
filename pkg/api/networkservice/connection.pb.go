// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connection.proto

package networkservice

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type State int32

const (
	State_UP   State = 0
	State_DOWN State = 1
)

var State_name = map[int32]string{
	0: "UP",
	1: "DOWN",
}

var State_value = map[string]int32{
	"UP":   0,
	"DOWN": 1,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

type ConnectionEventType int32

const (
	ConnectionEventType_INITIAL_STATE_TRANSFER ConnectionEventType = 0
	ConnectionEventType_UPDATE                 ConnectionEventType = 1
	ConnectionEventType_DELETE                 ConnectionEventType = 2
)

var ConnectionEventType_name = map[int32]string{
	0: "INITIAL_STATE_TRANSFER",
	1: "UPDATE",
	2: "DELETE",
}

var ConnectionEventType_value = map[string]int32{
	"INITIAL_STATE_TRANSFER": 0,
	"UPDATE":                 1,
	"DELETE":                 2,
}

func (x ConnectionEventType) String() string {
	return proto.EnumName(ConnectionEventType_name, int32(x))
}

func (ConnectionEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

type Mechanism struct {
	Cls                  string            `protobuf:"bytes,1,opt,name=cls,proto3" json:"cls,omitempty"`
	Type                 string            `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Parameters           map[string]string `protobuf:"bytes,3,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Mechanism) Reset()         { *m = Mechanism{} }
func (m *Mechanism) String() string { return proto.CompactTextString(m) }
func (*Mechanism) ProtoMessage()    {}
func (*Mechanism) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{0}
}

func (m *Mechanism) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Mechanism.Unmarshal(m, b)
}
func (m *Mechanism) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Mechanism.Marshal(b, m, deterministic)
}
func (m *Mechanism) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mechanism.Merge(m, src)
}
func (m *Mechanism) XXX_Size() int {
	return xxx_messageInfo_Mechanism.Size(m)
}
func (m *Mechanism) XXX_DiscardUnknown() {
	xxx_messageInfo_Mechanism.DiscardUnknown(m)
}

var xxx_messageInfo_Mechanism proto.InternalMessageInfo

func (m *Mechanism) GetCls() string {
	if m != nil {
		return m.Cls
	}
	return ""
}

func (m *Mechanism) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Mechanism) GetParameters() map[string]string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type PathSegment struct {
	Name                 string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string               `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Token                string               `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Expires              *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires,proto3" json:"expires,omitempty"`
	Metrics              map[string]string    `protobuf:"bytes,5,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PathSegment) Reset()         { *m = PathSegment{} }
func (m *PathSegment) String() string { return proto.CompactTextString(m) }
func (*PathSegment) ProtoMessage()    {}
func (*PathSegment) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{1}
}

func (m *PathSegment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathSegment.Unmarshal(m, b)
}
func (m *PathSegment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathSegment.Marshal(b, m, deterministic)
}
func (m *PathSegment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathSegment.Merge(m, src)
}
func (m *PathSegment) XXX_Size() int {
	return xxx_messageInfo_PathSegment.Size(m)
}
func (m *PathSegment) XXX_DiscardUnknown() {
	xxx_messageInfo_PathSegment.DiscardUnknown(m)
}

var xxx_messageInfo_PathSegment proto.InternalMessageInfo

func (m *PathSegment) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PathSegment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PathSegment) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *PathSegment) GetExpires() *timestamp.Timestamp {
	if m != nil {
		return m.Expires
	}
	return nil
}

func (m *PathSegment) GetMetrics() map[string]string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

type Path struct {
	Index                uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PathSegments         []*PathSegment `protobuf:"bytes,2,rep,name=path_segments,json=pathSegments,proto3" json:"path_segments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Path) Reset()         { *m = Path{} }
func (m *Path) String() string { return proto.CompactTextString(m) }
func (*Path) ProtoMessage()    {}
func (*Path) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{2}
}

func (m *Path) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Path.Unmarshal(m, b)
}
func (m *Path) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Path.Marshal(b, m, deterministic)
}
func (m *Path) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Path.Merge(m, src)
}
func (m *Path) XXX_Size() int {
	return xxx_messageInfo_Path.Size(m)
}
func (m *Path) XXX_DiscardUnknown() {
	xxx_messageInfo_Path.DiscardUnknown(m)
}

var xxx_messageInfo_Path proto.InternalMessageInfo

func (m *Path) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Path) GetPathSegments() []*PathSegment {
	if m != nil {
		return m.PathSegments
	}
	return nil
}

type Connection struct {
	Id                         string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	NetworkService             string             `protobuf:"bytes,2,opt,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	Mechanism                  *Mechanism         `protobuf:"bytes,3,opt,name=mechanism,proto3" json:"mechanism,omitempty"`
	Context                    *ConnectionContext `protobuf:"bytes,4,opt,name=context,proto3" json:"context,omitempty"`
	Labels                     map[string]string  `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Path                       *Path              `protobuf:"bytes,6,opt,name=path,proto3" json:"path,omitempty"`
	NetworkServiceEndpointName string             `protobuf:"bytes,7,opt,name=network_service_endpoint_name,json=networkServiceEndpointName,proto3" json:"network_service_endpoint_name,omitempty"`
	Payload                    string             `protobuf:"bytes,8,opt,name=payload,proto3" json:"payload,omitempty"`
	State                      State              `protobuf:"varint,9,opt,name=state,proto3,enum=connection.State" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}           `json:"-"`
	XXX_unrecognized           []byte             `json:"-"`
	XXX_sizecache              int32              `json:"-"`
}

func (m *Connection) Reset()         { *m = Connection{} }
func (m *Connection) String() string { return proto.CompactTextString(m) }
func (*Connection) ProtoMessage()    {}
func (*Connection) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{3}
}

func (m *Connection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connection.Unmarshal(m, b)
}
func (m *Connection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connection.Marshal(b, m, deterministic)
}
func (m *Connection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connection.Merge(m, src)
}
func (m *Connection) XXX_Size() int {
	return xxx_messageInfo_Connection.Size(m)
}
func (m *Connection) XXX_DiscardUnknown() {
	xxx_messageInfo_Connection.DiscardUnknown(m)
}

var xxx_messageInfo_Connection proto.InternalMessageInfo

func (m *Connection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Connection) GetNetworkService() string {
	if m != nil {
		return m.NetworkService
	}
	return ""
}

func (m *Connection) GetMechanism() *Mechanism {
	if m != nil {
		return m.Mechanism
	}
	return nil
}

func (m *Connection) GetContext() *ConnectionContext {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Connection) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Connection) GetPath() *Path {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Connection) GetNetworkServiceEndpointName() string {
	if m != nil {
		return m.NetworkServiceEndpointName
	}
	return ""
}

func (m *Connection) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *Connection) GetState() State {
	if m != nil {
		return m.State
	}
	return State_UP
}

type ConnectionEvent struct {
	Type                 ConnectionEventType    `protobuf:"varint,1,opt,name=type,proto3,enum=connection.ConnectionEventType" json:"type,omitempty"`
	Connections          map[string]*Connection `protobuf:"bytes,2,rep,name=connections,proto3" json:"connections,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConnectionEvent) Reset()         { *m = ConnectionEvent{} }
func (m *ConnectionEvent) String() string { return proto.CompactTextString(m) }
func (*ConnectionEvent) ProtoMessage()    {}
func (*ConnectionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{4}
}

func (m *ConnectionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionEvent.Unmarshal(m, b)
}
func (m *ConnectionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionEvent.Marshal(b, m, deterministic)
}
func (m *ConnectionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionEvent.Merge(m, src)
}
func (m *ConnectionEvent) XXX_Size() int {
	return xxx_messageInfo_ConnectionEvent.Size(m)
}
func (m *ConnectionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionEvent proto.InternalMessageInfo

func (m *ConnectionEvent) GetType() ConnectionEventType {
	if m != nil {
		return m.Type
	}
	return ConnectionEventType_INITIAL_STATE_TRANSFER
}

func (m *ConnectionEvent) GetConnections() map[string]*Connection {
	if m != nil {
		return m.Connections
	}
	return nil
}

type MonitorScopeSelector struct {
	PathSegments         []*PathSegment `protobuf:"bytes,1,rep,name=path_segments,json=pathSegments,proto3" json:"path_segments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MonitorScopeSelector) Reset()         { *m = MonitorScopeSelector{} }
func (m *MonitorScopeSelector) String() string { return proto.CompactTextString(m) }
func (*MonitorScopeSelector) ProtoMessage()    {}
func (*MonitorScopeSelector) Descriptor() ([]byte, []int) {
	return fileDescriptor_51baa40a1cc6b48b, []int{5}
}

func (m *MonitorScopeSelector) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MonitorScopeSelector.Unmarshal(m, b)
}
func (m *MonitorScopeSelector) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MonitorScopeSelector.Marshal(b, m, deterministic)
}
func (m *MonitorScopeSelector) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitorScopeSelector.Merge(m, src)
}
func (m *MonitorScopeSelector) XXX_Size() int {
	return xxx_messageInfo_MonitorScopeSelector.Size(m)
}
func (m *MonitorScopeSelector) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitorScopeSelector.DiscardUnknown(m)
}

var xxx_messageInfo_MonitorScopeSelector proto.InternalMessageInfo

func (m *MonitorScopeSelector) GetPathSegments() []*PathSegment {
	if m != nil {
		return m.PathSegments
	}
	return nil
}

func init() {
	proto.RegisterEnum("connection.State", State_name, State_value)
	proto.RegisterEnum("connection.ConnectionEventType", ConnectionEventType_name, ConnectionEventType_value)
	proto.RegisterType((*Mechanism)(nil), "connection.Mechanism")
	proto.RegisterMapType((map[string]string)(nil), "connection.Mechanism.ParametersEntry")
	proto.RegisterType((*PathSegment)(nil), "connection.PathSegment")
	proto.RegisterMapType((map[string]string)(nil), "connection.PathSegment.MetricsEntry")
	proto.RegisterType((*Path)(nil), "connection.Path")
	proto.RegisterType((*Connection)(nil), "connection.Connection")
	proto.RegisterMapType((map[string]string)(nil), "connection.Connection.LabelsEntry")
	proto.RegisterType((*ConnectionEvent)(nil), "connection.ConnectionEvent")
	proto.RegisterMapType((map[string]*Connection)(nil), "connection.ConnectionEvent.ConnectionsEntry")
	proto.RegisterType((*MonitorScopeSelector)(nil), "connection.MonitorScopeSelector")
}

func init() { proto.RegisterFile("connection.proto", fileDescriptor_51baa40a1cc6b48b) }

var fileDescriptor_51baa40a1cc6b48b = []byte{
	// 743 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x4e, 0xdb, 0x4c,
	0x10, 0xc5, 0xf9, 0x25, 0x13, 0x7e, 0xcc, 0x7e, 0x7c, 0xe0, 0xba, 0xaa, 0x1a, 0x45, 0x54, 0x44,
	0x08, 0x39, 0x55, 0xe8, 0x05, 0x8d, 0xda, 0x4a, 0x29, 0xb8, 0x52, 0x24, 0x48, 0x23, 0xc7, 0xb4,
	0x12, 0x37, 0x91, 0xe3, 0x4c, 0x13, 0x37, 0xb6, 0xd7, 0xb2, 0x17, 0x4a, 0x9e, 0xac, 0x4f, 0xd3,
	0x57, 0xe8, 0x7d, 0xef, 0x2a, 0xaf, 0xed, 0xd8, 0x84, 0x14, 0x89, 0x2b, 0xef, 0xcc, 0x9c, 0xd9,
	0x99, 0x39, 0x7b, 0xc6, 0x20, 0x9a, 0xd4, 0x75, 0xd1, 0x64, 0x16, 0x75, 0x15, 0xcf, 0xa7, 0x8c,
	0x12, 0x48, 0x3d, 0x72, 0xcd, 0x63, 0x73, 0x0f, 0x83, 0x26, 0xb3, 0x1c, 0x0c, 0x98, 0xe1, 0x78,
	0xe9, 0x29, 0x42, 0xcb, 0xfb, 0x29, 0xda, 0xa4, 0x2e, 0xc3, 0x3b, 0x16, 0x05, 0xea, 0x3f, 0x05,
	0xa8, 0x5c, 0xa2, 0x39, 0x35, 0x5c, 0x2b, 0x70, 0x88, 0x08, 0x79, 0xd3, 0x0e, 0x24, 0xa1, 0x26,
	0x34, 0x2a, 0x5a, 0x78, 0x24, 0x04, 0x0a, 0xe1, 0xdd, 0x52, 0x8e, 0xbb, 0xf8, 0x99, 0xa8, 0x00,
	0x9e, 0xe1, 0x1b, 0x0e, 0x32, 0xf4, 0x03, 0x29, 0x5f, 0xcb, 0x37, 0xaa, 0xad, 0x57, 0x4a, 0xa6,
	0xc3, 0xc5, 0x85, 0x4a, 0x7f, 0x81, 0x53, 0x5d, 0xe6, 0xcf, 0xb5, 0x4c, 0xa2, 0xfc, 0x1e, 0xb6,
	0x97, 0xc2, 0x61, 0xfd, 0x19, 0xce, 0x93, 0xfa, 0x33, 0x9c, 0x93, 0x5d, 0x28, 0xde, 0x1a, 0xf6,
	0x4d, 0xd2, 0x40, 0x64, 0xb4, 0x73, 0xa7, 0x42, 0xfd, 0x8f, 0x00, 0xd5, 0xbe, 0xc1, 0xa6, 0x03,
	0x9c, 0x38, 0xe8, 0xb2, 0xb0, 0x53, 0xd7, 0x70, 0x30, 0x4e, 0xe6, 0x67, 0xb2, 0x05, 0x39, 0x6b,
	0x1c, 0xa7, 0xe6, 0xac, 0x71, 0x78, 0x1b, 0xa3, 0x33, 0x74, 0xa5, 0x7c, 0x74, 0x1b, 0x37, 0xc8,
	0x1b, 0x28, 0xe3, 0x9d, 0x67, 0xf9, 0x18, 0x48, 0x85, 0x9a, 0xd0, 0xa8, 0xb6, 0x64, 0x65, 0x42,
	0xe9, 0xc4, 0xc6, 0x88, 0xa3, 0xd1, 0xcd, 0x37, 0x45, 0x4f, 0xf8, 0xd4, 0x12, 0x28, 0xf9, 0x00,
	0x65, 0x07, 0x99, 0x6f, 0x99, 0x81, 0x54, 0xe4, 0x14, 0x1c, 0x64, 0x29, 0xc8, 0x74, 0xa6, 0x5c,
	0x46, 0xb0, 0x88, 0x81, 0x24, 0x49, 0x6e, 0xc3, 0x46, 0x36, 0xf0, 0xa4, 0xd9, 0xaf, 0xa1, 0x10,
	0x16, 0x08, 0x11, 0x96, 0x3b, 0xc6, 0x3b, 0x9e, 0xb5, 0xa9, 0x45, 0x06, 0x79, 0x07, 0x9b, 0x9e,
	0xc1, 0xa6, 0xc3, 0x20, 0xaa, 0x1f, 0x48, 0x39, 0xde, 0xdf, 0xfe, 0x3f, 0xfa, 0xd3, 0x36, 0xbc,
	0xd4, 0x08, 0xea, 0xbf, 0xf2, 0x00, 0x67, 0x0b, 0x60, 0x4c, 0xa1, 0xb0, 0xa0, 0xf0, 0x10, 0xb6,
	0x5d, 0x64, 0x3f, 0xa8, 0x3f, 0x1b, 0x06, 0xe8, 0xdf, 0x5a, 0x66, 0xd2, 0xde, 0x56, 0xec, 0x1e,
	0x44, 0x5e, 0x72, 0x02, 0x15, 0x27, 0xd1, 0x01, 0xe7, 0xbb, 0xda, 0xfa, 0x7f, 0xa5, 0x48, 0xb4,
	0x14, 0x17, 0x92, 0x1a, 0xeb, 0x33, 0x7e, 0x8a, 0x2c, 0xa9, 0x89, 0x72, 0xd3, 0xee, 0xce, 0x22,
	0x8f, 0x96, 0x24, 0x91, 0x36, 0x94, 0x6c, 0x63, 0x84, 0x76, 0xf2, 0x26, 0xf5, 0x6c, 0xc5, 0x34,
	0x4f, 0xb9, 0xe0, 0xa0, 0xe8, 0x45, 0xe2, 0x0c, 0x72, 0x00, 0x85, 0x90, 0x08, 0xa9, 0xc4, 0x0b,
	0x8b, 0xcb, 0x6c, 0x69, 0x3c, 0x4a, 0x3a, 0xf0, 0x62, 0x69, 0xfe, 0x21, 0xba, 0x63, 0x8f, 0x5a,
	0x2e, 0x1b, 0x72, 0xfd, 0x95, 0x39, 0x1b, 0xf2, 0x7d, 0x36, 0xd4, 0x18, 0xd2, 0x0b, 0x55, 0x29,
	0x41, 0xd9, 0x33, 0xe6, 0x36, 0x35, 0xc6, 0xd2, 0x3a, 0x07, 0x27, 0x26, 0x39, 0x84, 0x62, 0xc0,
	0x0c, 0x86, 0x52, 0xa5, 0x26, 0x34, 0xb6, 0x5a, 0x3b, 0xd9, 0x1e, 0x06, 0x61, 0x40, 0x8b, 0xe2,
	0xf2, 0x5b, 0xa8, 0x66, 0x46, 0x78, 0x92, 0x76, 0x7e, 0x0b, 0xb0, 0x9d, 0x32, 0xa1, 0xde, 0x86,
	0xbb, 0x73, 0x12, 0x6f, 0xb9, 0xc0, 0xcb, 0xbe, 0x5c, 0x4d, 0x1a, 0x87, 0xea, 0x73, 0x0f, 0xe3,
	0xdf, 0x40, 0x0f, 0xaa, 0x29, 0x2e, 0x11, 0xd9, 0xf1, 0x23, 0xb9, 0x19, 0x3b, 0xa6, 0x3e, 0x7b,
	0x81, 0xfc, 0x05, 0xc4, 0x65, 0xc0, 0x8a, 0xc1, 0x8e, 0xb3, 0x83, 0x55, 0x5b, 0x7b, 0xab, 0xeb,
	0x65, 0x07, 0xd6, 0x61, 0xf7, 0x92, 0xba, 0x16, 0xa3, 0xfe, 0xc0, 0xa4, 0x1e, 0x0e, 0xd0, 0x46,
	0x93, 0x51, 0xff, 0xe1, 0x9a, 0x08, 0x4f, 0x58, 0x93, 0xa3, 0x67, 0x50, 0xe4, 0x2f, 0x42, 0x4a,
	0x90, 0xbb, 0xea, 0x8b, 0x6b, 0x64, 0x1d, 0x0a, 0xe7, 0x9f, 0xbf, 0xf6, 0x44, 0xe1, 0xa8, 0x0b,
	0xff, 0xad, 0x60, 0x8d, 0xc8, 0xb0, 0xd7, 0xed, 0x75, 0xf5, 0x6e, 0xe7, 0x62, 0x38, 0xd0, 0x3b,
	0xba, 0x3a, 0xd4, 0xb5, 0x4e, 0x6f, 0xf0, 0x49, 0xd5, 0xc4, 0x35, 0x02, 0x50, 0xba, 0xea, 0x9f,
	0x77, 0x74, 0x55, 0x14, 0xc2, 0xf3, 0xb9, 0x7a, 0xa1, 0xea, 0xaa, 0x98, 0x6b, 0x7d, 0x87, 0x9d,
	0xb8, 0xf7, 0xcc, 0x4a, 0x5e, 0x01, 0x79, 0xe0, 0x0c, 0x48, 0xed, 0xde, 0x72, 0xad, 0x18, 0x58,
	0x7e, 0xfe, 0xc8, 0xdb, 0xbc, 0x16, 0x3e, 0xb6, 0xaf, 0x4f, 0x27, 0x16, 0x9b, 0xde, 0x8c, 0x14,
	0x93, 0x3a, 0xcd, 0x58, 0xbf, 0xb1, 0xc6, 0x1d, 0x0c, 0xa6, 0x4d, 0xc3, 0xb3, 0x9a, 0xde, 0x6c,
	0xc2, 0xbf, 0xf7, 0xc3, 0xa3, 0x12, 0xff, 0x53, 0x9e, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x70,
	0x37, 0xbd, 0xa2, 0xa8, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MonitorConnectionClient is the client API for MonitorConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MonitorConnectionClient interface {
	MonitorConnections(ctx context.Context, in *MonitorScopeSelector, opts ...grpc.CallOption) (MonitorConnection_MonitorConnectionsClient, error)
}

type monitorConnectionClient struct {
	cc grpc.ClientConnInterface
}

func NewMonitorConnectionClient(cc grpc.ClientConnInterface) MonitorConnectionClient {
	return &monitorConnectionClient{cc}
}

func (c *monitorConnectionClient) MonitorConnections(ctx context.Context, in *MonitorScopeSelector, opts ...grpc.CallOption) (MonitorConnection_MonitorConnectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MonitorConnection_serviceDesc.Streams[0], "/connection.MonitorConnection/MonitorConnections", opts...)
	if err != nil {
		return nil, err
	}
	x := &monitorConnectionMonitorConnectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MonitorConnection_MonitorConnectionsClient interface {
	Recv() (*ConnectionEvent, error)
	grpc.ClientStream
}

type monitorConnectionMonitorConnectionsClient struct {
	grpc.ClientStream
}

func (x *monitorConnectionMonitorConnectionsClient) Recv() (*ConnectionEvent, error) {
	m := new(ConnectionEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MonitorConnectionServer is the server API for MonitorConnection service.
type MonitorConnectionServer interface {
	MonitorConnections(*MonitorScopeSelector, MonitorConnection_MonitorConnectionsServer) error
}

// UnimplementedMonitorConnectionServer can be embedded to have forward compatible implementations.
type UnimplementedMonitorConnectionServer struct {
}

func (*UnimplementedMonitorConnectionServer) MonitorConnections(req *MonitorScopeSelector, srv MonitorConnection_MonitorConnectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method MonitorConnections not implemented")
}

func RegisterMonitorConnectionServer(s *grpc.Server, srv MonitorConnectionServer) {
	s.RegisterService(&_MonitorConnection_serviceDesc, srv)
}

func _MonitorConnection_MonitorConnections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MonitorScopeSelector)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MonitorConnectionServer).MonitorConnections(m, &monitorConnectionMonitorConnectionsServer{stream})
}

type MonitorConnection_MonitorConnectionsServer interface {
	Send(*ConnectionEvent) error
	grpc.ServerStream
}

type monitorConnectionMonitorConnectionsServer struct {
	grpc.ServerStream
}

func (x *monitorConnectionMonitorConnectionsServer) Send(m *ConnectionEvent) error {
	return x.ServerStream.SendMsg(m)
}

var _MonitorConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "connection.MonitorConnection",
	HandlerType: (*MonitorConnectionServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MonitorConnections",
			Handler:       _MonitorConnection_MonitorConnections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "connection.proto",
}
