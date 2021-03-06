// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/findMaxNumber.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FindMaxNumberRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Sig                  []byte   `protobuf:"bytes,2,opt,name=sig,proto3" json:"sig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxNumberRequest) Reset()         { *m = FindMaxNumberRequest{} }
func (m *FindMaxNumberRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaxNumberRequest) ProtoMessage()    {}
func (*FindMaxNumberRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_076cc4f30fd5f537, []int{0}
}

func (m *FindMaxNumberRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxNumberRequest.Unmarshal(m, b)
}
func (m *FindMaxNumberRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxNumberRequest.Marshal(b, m, deterministic)
}
func (m *FindMaxNumberRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxNumberRequest.Merge(m, src)
}
func (m *FindMaxNumberRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaxNumberRequest.Size(m)
}
func (m *FindMaxNumberRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxNumberRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxNumberRequest proto.InternalMessageInfo

func (m *FindMaxNumberRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *FindMaxNumberRequest) GetSig() []byte {
	if m != nil {
		return m.Sig
	}
	return nil
}

type FindMaxNumberResponse struct {
	MaxNumber            int64    `protobuf:"varint,1,opt,name=maxNumber,proto3" json:"maxNumber,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxNumberResponse) Reset()         { *m = FindMaxNumberResponse{} }
func (m *FindMaxNumberResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxNumberResponse) ProtoMessage()    {}
func (*FindMaxNumberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_076cc4f30fd5f537, []int{1}
}

func (m *FindMaxNumberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxNumberResponse.Unmarshal(m, b)
}
func (m *FindMaxNumberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxNumberResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxNumberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxNumberResponse.Merge(m, src)
}
func (m *FindMaxNumberResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxNumberResponse.Size(m)
}
func (m *FindMaxNumberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxNumberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxNumberResponse proto.InternalMessageInfo

func (m *FindMaxNumberResponse) GetMaxNumber() int64 {
	if m != nil {
		return m.MaxNumber
	}
	return 0
}

func init() {
	proto.RegisterType((*FindMaxNumberRequest)(nil), "findMaxNumber.grpc.proto.FindMaxNumberRequest")
	proto.RegisterType((*FindMaxNumberResponse)(nil), "findMaxNumber.grpc.proto.FindMaxNumberResponse")
}

func init() { proto.RegisterFile("proto/findMaxNumber.proto", fileDescriptor_076cc4f30fd5f537) }

var fileDescriptor_076cc4f30fd5f537 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xcb, 0xcc, 0x4b, 0xf1, 0x4d, 0xac, 0xf0, 0x2b, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2,
	0x03, 0x8b, 0x09, 0x49, 0xa0, 0x0a, 0xa6, 0x17, 0x15, 0x24, 0x43, 0x64, 0x94, 0x1c, 0xb8, 0x44,
	0xdc, 0x90, 0xe5, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xc4, 0xb8, 0xd8, 0xf2, 0xc0,
	0x02, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x50, 0x9e, 0x90, 0x00, 0x17, 0x73, 0x71, 0x66,
	0xba, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0x88, 0xa9, 0x64, 0xca, 0x25, 0x8a, 0x66, 0x42,
	0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x0c, 0x17, 0x67, 0x2e, 0x4c, 0x10, 0x6a, 0x0a, 0x42,
	0xc0, 0xa8, 0x91, 0x91, 0x8b, 0x1d, 0xc2, 0x2c, 0x16, 0x2a, 0xe3, 0xe2, 0x45, 0x31, 0x42, 0x48,
	0x4f, 0x0f, 0x97, 0x83, 0xf5, 0xb0, 0xb9, 0x56, 0x4a, 0x9f, 0x68, 0xf5, 0x10, 0xb7, 0x29, 0x31,
	0x68, 0x30, 0x1a, 0x30, 0x3a, 0xb1, 0x47, 0xb1, 0x82, 0x15, 0x25, 0xb1, 0x81, 0x29, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xae, 0x1d, 0x5b, 0x43, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NumbersClient is the client API for Numbers service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NumbersClient interface {
	FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (Numbers_FindMaxNumberClient, error)
}

type numbersClient struct {
	cc *grpc.ClientConn
}

func NewNumbersClient(cc *grpc.ClientConn) NumbersClient {
	return &numbersClient{cc}
}

func (c *numbersClient) FindMaxNumber(ctx context.Context, opts ...grpc.CallOption) (Numbers_FindMaxNumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Numbers_serviceDesc.Streams[0], "/findMaxNumber.grpc.proto.Numbers/FindMaxNumber", opts...)
	if err != nil {
		return nil, err
	}
	x := &numbersFindMaxNumberClient{stream}
	return x, nil
}

type Numbers_FindMaxNumberClient interface {
	Send(*FindMaxNumberRequest) error
	Recv() (*FindMaxNumberResponse, error)
	grpc.ClientStream
}

type numbersFindMaxNumberClient struct {
	grpc.ClientStream
}

func (x *numbersFindMaxNumberClient) Send(m *FindMaxNumberRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *numbersFindMaxNumberClient) Recv() (*FindMaxNumberResponse, error) {
	m := new(FindMaxNumberResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NumbersServer is the server API for Numbers service.
type NumbersServer interface {
	FindMaxNumber(Numbers_FindMaxNumberServer) error
}

func RegisterNumbersServer(s *grpc.Server, srv NumbersServer) {
	s.RegisterService(&_Numbers_serviceDesc, srv)
}

func _Numbers_FindMaxNumber_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NumbersServer).FindMaxNumber(&numbersFindMaxNumberServer{stream})
}

type Numbers_FindMaxNumberServer interface {
	Send(*FindMaxNumberResponse) error
	Recv() (*FindMaxNumberRequest, error)
	grpc.ServerStream
}

type numbersFindMaxNumberServer struct {
	grpc.ServerStream
}

func (x *numbersFindMaxNumberServer) Send(m *FindMaxNumberResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *numbersFindMaxNumberServer) Recv() (*FindMaxNumberRequest, error) {
	m := new(FindMaxNumberRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Numbers_serviceDesc = grpc.ServiceDesc{
	ServiceName: "findMaxNumber.grpc.proto.Numbers",
	HandlerType: (*NumbersServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FindMaxNumber",
			Handler:       _Numbers_FindMaxNumber_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/findMaxNumber.proto",
}
