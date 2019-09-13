// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cat/cat.proto

package cat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetMyCatMessage struct {
	TargetCat            string   `protobuf:"bytes,1,opt,name=target_cat,json=targetCat,proto3" json:"target_cat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMyCatMessage) Reset()         { *m = GetMyCatMessage{} }
func (m *GetMyCatMessage) String() string { return proto.CompactTextString(m) }
func (*GetMyCatMessage) ProtoMessage()    {}
func (*GetMyCatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a17bbdc5b4d1d91, []int{0}
}

func (m *GetMyCatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMyCatMessage.Unmarshal(m, b)
}
func (m *GetMyCatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMyCatMessage.Marshal(b, m, deterministic)
}
func (m *GetMyCatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMyCatMessage.Merge(m, src)
}
func (m *GetMyCatMessage) XXX_Size() int {
	return xxx_messageInfo_GetMyCatMessage.Size(m)
}
func (m *GetMyCatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMyCatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetMyCatMessage proto.InternalMessageInfo

func (m *GetMyCatMessage) GetTargetCat() string {
	if m != nil {
		return m.TargetCat
	}
	return ""
}

type MyCatResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyCatResponse) Reset()         { *m = MyCatResponse{} }
func (m *MyCatResponse) String() string { return proto.CompactTextString(m) }
func (*MyCatResponse) ProtoMessage()    {}
func (*MyCatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a17bbdc5b4d1d91, []int{1}
}

func (m *MyCatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyCatResponse.Unmarshal(m, b)
}
func (m *MyCatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyCatResponse.Marshal(b, m, deterministic)
}
func (m *MyCatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyCatResponse.Merge(m, src)
}
func (m *MyCatResponse) XXX_Size() int {
	return xxx_messageInfo_MyCatResponse.Size(m)
}
func (m *MyCatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MyCatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MyCatResponse proto.InternalMessageInfo

func (m *MyCatResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MyCatResponse) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMyCatMessage)(nil), "GetMyCatMessage")
	proto.RegisterType((*MyCatResponse)(nil), "MyCatResponse")
}

func init() { proto.RegisterFile("cat/cat.proto", fileDescriptor_4a17bbdc5b4d1d91) }

var fileDescriptor_4a17bbdc5b4d1d91 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4e, 0x2c, 0xd1,
	0x4f, 0x4e, 0x2c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe0, 0xe2, 0x77, 0x4f, 0x2d,
	0xf1, 0xad, 0x74, 0x4e, 0x2c, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x15, 0x92, 0xe5, 0xe2,
	0x2a, 0x49, 0x2c, 0x4a, 0x4f, 0x2d, 0x89, 0x4f, 0x4e, 0x2c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0xe2, 0x84, 0x88, 0x38, 0x27, 0x96, 0x28, 0x99, 0x73, 0xf1, 0x82, 0x95, 0x07, 0xa5, 0x16,
	0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x42, 0x55, 0x82,
	0xd9, 0x20, 0xb1, 0xec, 0xcc, 0xbc, 0x14, 0x09, 0x26, 0x88, 0x18, 0x88, 0x6d, 0x64, 0xca, 0xc5,
	0xec, 0x9c, 0x58, 0x22, 0xa4, 0xc7, 0xc5, 0x01, 0xb3, 0x51, 0x48, 0x40, 0x0f, 0xcd, 0x72, 0x29,
	0x3e, 0x3d, 0x14, 0xc3, 0x95, 0x18, 0x92, 0xd8, 0xc0, 0x0e, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x5c, 0xc1, 0x4a, 0x84, 0xb9, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatClient is the client API for Cat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatClient interface {
	GetMyCat(ctx context.Context, in *GetMyCatMessage, opts ...grpc.CallOption) (*MyCatResponse, error)
}

type catClient struct {
	cc *grpc.ClientConn
}

func NewCatClient(cc *grpc.ClientConn) CatClient {
	return &catClient{cc}
}

func (c *catClient) GetMyCat(ctx context.Context, in *GetMyCatMessage, opts ...grpc.CallOption) (*MyCatResponse, error) {
	out := new(MyCatResponse)
	err := c.cc.Invoke(ctx, "/Cat/GetMyCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServer is the server API for Cat service.
type CatServer interface {
	GetMyCat(context.Context, *GetMyCatMessage) (*MyCatResponse, error)
}

// UnimplementedCatServer can be embedded to have forward compatible implementations.
type UnimplementedCatServer struct {
}

func (*UnimplementedCatServer) GetMyCat(ctx context.Context, req *GetMyCatMessage) (*MyCatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyCat not implemented")
}

func RegisterCatServer(s *grpc.Server, srv CatServer) {
	s.RegisterService(&_Cat_serviceDesc, srv)
}

func _Cat_GetMyCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).GetMyCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cat/GetMyCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).GetMyCat(ctx, req.(*GetMyCatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cat",
	HandlerType: (*CatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyCat",
			Handler:    _Cat_GetMyCat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cat/cat.proto",
}
