// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf.proto

package protobuf

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Event struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StartTime            int64    `protobuf:"varint,2,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime              int64    `protobuf:"varint,3,opt,name=endTime,proto3" json:"endTime,omitempty"`
	Title                string   `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Event) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *Event) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *Event) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Event) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type Id struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Id) Reset()         { *m = Id{} }
func (m *Id) String() string { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()    {}
func (*Id) Descriptor() ([]byte, []int) {
	return fileDescriptor_c77a803fcbc0c059, []int{1}
}

func (m *Id) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Id.Unmarshal(m, b)
}
func (m *Id) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Id.Marshal(b, m, deterministic)
}
func (m *Id) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Id.Merge(m, src)
}
func (m *Id) XXX_Size() int {
	return xxx_messageInfo_Id.Size(m)
}
func (m *Id) XXX_DiscardUnknown() {
	xxx_messageInfo_Id.DiscardUnknown(m)
}

var xxx_messageInfo_Id proto.InternalMessageInfo

func (m *Id) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*Event)(nil), "Event")
	proto.RegisterType((*Id)(nil), "Id")
}

func init() { proto.RegisterFile("protobuf.proto", fileDescriptor_c77a803fcbc0c059) }

var fileDescriptor_c77a803fcbc0c059 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0x87, 0xbb, 0x49, 0x37, 0xff, 0xed, 0xfc, 0xa1, 0xc8, 0x52, 0x64, 0x8d, 0x1e, 0x42, 0x4e,
	0x39, 0xc8, 0x16, 0xf4, 0x09, 0x44, 0x8a, 0xd4, 0x63, 0xf4, 0x05, 0x52, 0x67, 0x2c, 0x0b, 0x69,
	0x36, 0x6c, 0xa7, 0x82, 0x4f, 0xe2, 0xeb, 0x4a, 0x77, 0xad, 0x8a, 0xb7, 0xf9, 0xbe, 0xef, 0x32,
	0xfc, 0x60, 0x3e, 0x06, 0xcf, 0x7e, 0x73, 0x78, 0xb5, 0xf1, 0x28, 0x2f, 0xb7, 0xde, 0x6f, 0x7b,
	0x5a, 0x9e, 0xf4, 0x92, 0x76, 0x23, 0xbf, 0xa7, 0x58, 0x7f, 0x08, 0x90, 0xab, 0x37, 0x1a, 0x58,
	0xcf, 0x21, 0x73, 0x68, 0x44, 0x25, 0x1a, 0xd9, 0x66, 0x0e, 0xf5, 0x15, 0xcc, 0xf6, 0xdc, 0x05,
	0x7e, 0x76, 0x3b, 0x32, 0x59, 0x25, 0x9a, 0xbc, 0xfd, 0x11, 0xda, 0xc0, 0x3f, 0x1a, 0x30, 0xb6,
	0x3c, 0xb6, 0x13, 0xea, 0x05, 0x48, 0x76, 0xdc, 0x93, 0x91, 0x95, 0x68, 0x66, 0x6d, 0x02, 0x5d,
	0xc1, 0x7f, 0xa4, 0xfd, 0x4b, 0x70, 0x23, 0x3b, 0x3f, 0x98, 0x22, 0xb6, 0xdf, 0xea, 0x71, 0xaa,
	0xa6, 0x67, 0xb2, 0x55, 0x78, 0x08, 0xdd, 0x91, 0xeb, 0x05, 0x64, 0x6b, 0xfc, 0xfb, 0xd5, 0xcd,
	0x13, 0xa8, 0xfb, 0xae, 0xa7, 0x01, 0xbb, 0xa0, 0x2f, 0x40, 0x3d, 0x10, 0xa7, 0xef, 0x73, 0xbb,
	0xc6, 0xb2, 0xb0, 0x11, 0xea, 0x89, 0xbe, 0x06, 0x75, 0x87, 0x98, 0xd2, 0x97, 0x2d, 0xcf, 0x6d,
	0x1a, 0xc2, 0x7e, 0xef, 0xb3, 0x3a, 0x0e, 0x51, 0x4f, 0x36, 0x45, 0x34, 0xb7, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xbb, 0x81, 0x60, 0xa0, 0x3a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalendarClient is the client API for Calendar service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalendarClient interface {
	GetEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Event, error)
	AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error)
}

type calendarClient struct {
	cc grpc.ClientConnInterface
}

func NewCalendarClient(cc grpc.ClientConnInterface) CalendarClient {
	return &calendarClient{cc}
}

func (c *calendarClient) GetEvent(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Event, error) {
	out := new(Event)
	err := c.cc.Invoke(ctx, "/Calendar/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calendarClient) AddEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Calendar/AddEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalendarServer is the server API for Calendar service.
type CalendarServer interface {
	GetEvent(context.Context, *Id) (*Event, error)
	AddEvent(context.Context, *Event) (*empty.Empty, error)
}

// UnimplementedCalendarServer can be embedded to have forward compatible implementations.
type UnimplementedCalendarServer struct {
}

func (*UnimplementedCalendarServer) GetEvent(ctx context.Context, req *Id) (*Event, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (*UnimplementedCalendarServer) AddEvent(ctx context.Context, req *Event) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEvent not implemented")
}

func RegisterCalendarServer(s *grpc.Server, srv CalendarServer) {
	s.RegisterService(&_Calendar_serviceDesc, srv)
}

func _Calendar_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).GetEvent(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calendar_AddEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalendarServer).AddEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Calendar/AddEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalendarServer).AddEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _Calendar_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Calendar",
	HandlerType: (*CalendarServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _Calendar_GetEvent_Handler,
		},
		{
			MethodName: "AddEvent",
			Handler:    _Calendar_AddEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf.proto",
}
