// Code generated by protoc-gen-go. DO NOT EDIT.
// source: im.proto

package protobuf

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

// 群发
type SendMsgAllReq struct {
	AppId                string   `protobuf:"bytes,1,opt,name=appId,proto3" json:"appId,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	TargetIds            []string `protobuf:"bytes,3,rep,name=targetIds,proto3" json:"targetIds,omitempty"`
	Seq                  string   `protobuf:"bytes,4,opt,name=seq,proto3" json:"seq,omitempty"`
	Message              string   `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgAllReq) Reset()         { *m = SendMsgAllReq{} }
func (m *SendMsgAllReq) String() string { return proto.CompactTextString(m) }
func (*SendMsgAllReq) ProtoMessage()    {}
func (*SendMsgAllReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f2114a3e4ddb9e, []int{0}
}

func (m *SendMsgAllReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgAllReq.Unmarshal(m, b)
}
func (m *SendMsgAllReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgAllReq.Marshal(b, m, deterministic)
}
func (m *SendMsgAllReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgAllReq.Merge(m, src)
}
func (m *SendMsgAllReq) XXX_Size() int {
	return xxx_messageInfo_SendMsgAllReq.Size(m)
}
func (m *SendMsgAllReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgAllReq.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgAllReq proto.InternalMessageInfo

func (m *SendMsgAllReq) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *SendMsgAllReq) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *SendMsgAllReq) GetTargetIds() []string {
	if m != nil {
		return m.TargetIds
	}
	return nil
}

func (m *SendMsgAllReq) GetSeq() string {
	if m != nil {
		return m.Seq
	}
	return ""
}

func (m *SendMsgAllReq) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type SendMsgAllRsp struct {
	PushIds              []string `protobuf:"bytes,1,rep,name=pushIds,proto3" json:"pushIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgAllRsp) Reset()         { *m = SendMsgAllRsp{} }
func (m *SendMsgAllRsp) String() string { return proto.CompactTextString(m) }
func (*SendMsgAllRsp) ProtoMessage()    {}
func (*SendMsgAllRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_36f2114a3e4ddb9e, []int{1}
}

func (m *SendMsgAllRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgAllRsp.Unmarshal(m, b)
}
func (m *SendMsgAllRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgAllRsp.Marshal(b, m, deterministic)
}
func (m *SendMsgAllRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgAllRsp.Merge(m, src)
}
func (m *SendMsgAllRsp) XXX_Size() int {
	return xxx_messageInfo_SendMsgAllRsp.Size(m)
}
func (m *SendMsgAllRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgAllRsp.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgAllRsp proto.InternalMessageInfo

func (m *SendMsgAllRsp) GetPushIds() []string {
	if m != nil {
		return m.PushIds
	}
	return nil
}

func init() {
	proto.RegisterType((*SendMsgAllReq)(nil), "protobuf.SendMsgAllReq")
	proto.RegisterType((*SendMsgAllRsp)(nil), "protobuf.SendMsgAllRsp")
}

func init() { proto.RegisterFile("im.proto", fileDescriptor_36f2114a3e4ddb9e) }

var fileDescriptor_36f2114a3e4ddb9e = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0xcc, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0x8d, 0x8c, 0x5c, 0xbc,
	0xc1, 0xa9, 0x79, 0x29, 0xbe, 0xc5, 0xe9, 0x8e, 0x39, 0x39, 0x41, 0xa9, 0x85, 0x42, 0x22, 0x5c,
	0xac, 0x89, 0x05, 0x05, 0x9e, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90,
	0x00, 0x17, 0x73, 0x69, 0x66, 0x8a, 0x04, 0x13, 0x58, 0x0c, 0xc4, 0x14, 0x92, 0xe1, 0xe2, 0x2c,
	0x49, 0x2c, 0x4a, 0x4f, 0x2d, 0xf1, 0x4c, 0x29, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x0c, 0x42,
	0x08, 0x80, 0xd4, 0x17, 0xa7, 0x16, 0x4a, 0xb0, 0x40, 0xd4, 0x17, 0xa7, 0x16, 0x0a, 0x49, 0x70,
	0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0xb0, 0x81, 0x45, 0x61, 0x5c, 0x25, 0x4d,
	0x14, 0x27, 0x14, 0x17, 0x80, 0x94, 0x16, 0x94, 0x16, 0x67, 0x80, 0x0c, 0x66, 0x04, 0x1b, 0x0c,
	0xe3, 0x1a, 0xf9, 0x70, 0x71, 0x78, 0xfa, 0x06, 0xa7, 0x16, 0x95, 0xa5, 0x16, 0x09, 0x39, 0x70,
	0x71, 0x21, 0xb4, 0x09, 0x89, 0xeb, 0xc1, 0xfc, 0xa4, 0x87, 0xe2, 0x1f, 0x29, 0xec, 0x12, 0xc5,
	0x05, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x19, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x5f,
	0xb2, 0x84, 0x19, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IMServerClient is the client API for IMServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IMServerClient interface {
	SendMsgAll(ctx context.Context, in *SendMsgAllReq, opts ...grpc.CallOption) (*SendMsgAllRsp, error)
}

type iMServerClient struct {
	cc *grpc.ClientConn
}

func NewIMServerClient(cc *grpc.ClientConn) IMServerClient {
	return &iMServerClient{cc}
}

func (c *iMServerClient) SendMsgAll(ctx context.Context, in *SendMsgAllReq, opts ...grpc.CallOption) (*SendMsgAllRsp, error) {
	out := new(SendMsgAllRsp)
	err := c.cc.Invoke(ctx, "/protobuf.IMServer/SendMsgAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IMServerServer is the server API for IMServer service.
type IMServerServer interface {
	SendMsgAll(context.Context, *SendMsgAllReq) (*SendMsgAllRsp, error)
}

// UnimplementedIMServerServer can be embedded to have forward compatible implementations.
type UnimplementedIMServerServer struct {
}

func (*UnimplementedIMServerServer) SendMsgAll(ctx context.Context, req *SendMsgAllReq) (*SendMsgAllRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMsgAll not implemented")
}

func RegisterIMServerServer(s *grpc.Server, srv IMServerServer) {
	s.RegisterService(&_IMServer_serviceDesc, srv)
}

func _IMServer_SendMsgAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMsgAllReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IMServerServer).SendMsgAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.IMServer/SendMsgAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IMServerServer).SendMsgAll(ctx, req.(*SendMsgAllReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _IMServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.IMServer",
	HandlerType: (*IMServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendMsgAll",
			Handler:    _IMServer_SendMsgAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "im.proto",
}
