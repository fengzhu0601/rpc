// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_service.proto

package pb

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

type CreateLaptopRequest struct {
	Laptop               *Laptop  `protobuf:"bytes,1,opt,name=laptop,proto3" json:"laptop,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopRequest) Reset()         { *m = CreateLaptopRequest{} }
func (m *CreateLaptopRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopRequest) ProtoMessage()    {}
func (*CreateLaptopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{0}
}

func (m *CreateLaptopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopRequest.Unmarshal(m, b)
}
func (m *CreateLaptopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopRequest.Marshal(b, m, deterministic)
}
func (m *CreateLaptopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopRequest.Merge(m, src)
}
func (m *CreateLaptopRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopRequest.Size(m)
}
func (m *CreateLaptopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopRequest proto.InternalMessageInfo

func (m *CreateLaptopRequest) GetLaptop() *Laptop {
	if m != nil {
		return m.Laptop
	}
	return nil
}

type CreateLaptopResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLaptopResponse) Reset()         { *m = CreateLaptopResponse{} }
func (m *CreateLaptopResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLaptopResponse) ProtoMessage()    {}
func (*CreateLaptopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_240c60d9fb227e71, []int{1}
}

func (m *CreateLaptopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLaptopResponse.Unmarshal(m, b)
}
func (m *CreateLaptopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLaptopResponse.Marshal(b, m, deterministic)
}
func (m *CreateLaptopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLaptopResponse.Merge(m, src)
}
func (m *CreateLaptopResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLaptopResponse.Size(m)
}
func (m *CreateLaptopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLaptopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLaptopResponse proto.InternalMessageInfo

func (m *CreateLaptopResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateLaptopRequest)(nil), "rpc.CreateLaptopRequest")
	proto.RegisterType((*CreateLaptopResponse)(nil), "rpc.CreateLaptopResponse")
}

func init() { proto.RegisterFile("laptop_service.proto", fileDescriptor_240c60d9fb227e71) }

var fileDescriptor_240c60d9fb227e71 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0x49, 0x2c, 0x28,
	0xc9, 0x2f, 0x88, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0x82, 0x49, 0xe5, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0x43, 0xa5,
	0x94, 0xac, 0xb8, 0x84, 0x9d, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x7d, 0xc0, 0xb2, 0x41, 0xa9, 0x85,
	0xa5, 0xa9, 0xc5, 0x25, 0x42, 0xca, 0x5c, 0x6c, 0x10, 0xe5, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc,
	0x46, 0xdc, 0x7a, 0x45, 0x05, 0xc9, 0x7a, 0x50, 0x35, 0x50, 0x29, 0x25, 0x35, 0x2e, 0x11, 0x54,
	0xbd, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0x8d, 0x9c,
	0x41, 0x4c, 0x99, 0x29, 0x46, 0x61, 0x5c, 0xbc, 0x10, 0x15, 0xc1, 0x10, 0x57, 0x09, 0xb9, 0x72,
	0xf1, 0x20, 0x6b, 0x14, 0x92, 0x00, 0x9b, 0x8e, 0xc5, 0x1d, 0x52, 0x92, 0x58, 0x64, 0x20, 0xb6,
	0x28, 0x31, 0x38, 0xb1, 0x44, 0x31, 0x15, 0x24, 0x25, 0xb1, 0x81, 0x3d, 0x62, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0xef, 0xf8, 0x8c, 0xa1, 0xfb, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LaptopServiceClient is the client API for LaptopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LaptopServiceClient interface {
	CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error)
}

type laptopServiceClient struct {
	cc *grpc.ClientConn
}

func NewLaptopServiceClient(cc *grpc.ClientConn) LaptopServiceClient {
	return &laptopServiceClient{cc}
}

func (c *laptopServiceClient) CreateLaptop(ctx context.Context, in *CreateLaptopRequest, opts ...grpc.CallOption) (*CreateLaptopResponse, error) {
	out := new(CreateLaptopResponse)
	err := c.cc.Invoke(ctx, "/rpc.LaptopService/CreateLaptop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaptopServiceServer is the server API for LaptopService service.
type LaptopServiceServer interface {
	CreateLaptop(context.Context, *CreateLaptopRequest) (*CreateLaptopResponse, error)
}

// UnimplementedLaptopServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLaptopServiceServer struct {
}

func (*UnimplementedLaptopServiceServer) CreateLaptop(ctx context.Context, req *CreateLaptopRequest) (*CreateLaptopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaptop not implemented")
}

func RegisterLaptopServiceServer(s *grpc.Server, srv LaptopServiceServer) {
	s.RegisterService(&_LaptopService_serviceDesc, srv)
}

func _LaptopService_CreateLaptop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLaptopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.LaptopService/CreateLaptop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaptopServiceServer).CreateLaptop(ctx, req.(*CreateLaptopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LaptopService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.LaptopService",
	HandlerType: (*LaptopServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaptop",
			Handler:    _LaptopService_CreateLaptop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "laptop_service.proto",
}
