// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

/*
Package api is a generated protocol buffer package.

API Service

It is generated from these files:
	api/api.proto

It has these top-level messages:
	Record
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Message represents a simple message sent to the Echo service.
type Record struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Record Type see https://github.com/miekg/dns/blob/master/types.go#L27
	Type string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	// Record Name
	Domain string `protobuf:"bytes,3,opt,name=domain" json:"domain,omitempty"`
	// TTL time to live of the record
	Expires int32 `protobuf:"varint,4,opt,name=expires" json:"expires,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Record) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Record) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Record) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Record) GetExpires() int32 {
	if m != nil {
		return m.Expires
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "api.Record")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DDNSService service

type DDNSServiceClient interface {
	// Echo method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	SaveRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error)
	DeleteRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error)
}

type dDNSServiceClient struct {
	cc *grpc.ClientConn
}

func NewDDNSServiceClient(cc *grpc.ClientConn) DDNSServiceClient {
	return &dDNSServiceClient{cc}
}

func (c *dDNSServiceClient) SaveRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := grpc.Invoke(ctx, "/api.DDNSService/SaveRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dDNSServiceClient) DeleteRecord(ctx context.Context, in *Record, opts ...grpc.CallOption) (*Record, error) {
	out := new(Record)
	err := grpc.Invoke(ctx, "/api.DDNSService/DeleteRecord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DDNSService service

type DDNSServiceServer interface {
	// Echo method receives a simple message and returns it.
	// The message posted as the id parameter will also be returned.
	SaveRecord(context.Context, *Record) (*Record, error)
	DeleteRecord(context.Context, *Record) (*Record, error)
}

func RegisterDDNSServiceServer(s *grpc.Server, srv DDNSServiceServer) {
	s.RegisterService(&_DDNSService_serviceDesc, srv)
}

func _DDNSService_SaveRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDNSServiceServer).SaveRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DDNSService/SaveRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDNSServiceServer).SaveRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

func _DDNSService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Record)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DDNSServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DDNSService/DeleteRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DDNSServiceServer).DeleteRecord(ctx, req.(*Record))
	}
	return interceptor(ctx, in, info, handler)
}

var _DDNSService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DDNSService",
	HandlerType: (*DDNSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveRecord",
			Handler:    _DDNSService_SaveRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _DDNSService_DeleteRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0xe5, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64,
	0xe6, 0xe7, 0x15, 0x43, 0x94, 0x28, 0xc5, 0x71, 0xb1, 0x05, 0xa5, 0x26, 0xe7, 0x17, 0xa5, 0x08,
	0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08,
	0x09, 0x71, 0xb1, 0x94, 0x54, 0x16, 0xa4, 0x4a, 0x30, 0x81, 0x45, 0xc0, 0x6c, 0x21, 0x31, 0x2e,
	0xb6, 0x94, 0xfc, 0xdc, 0xc4, 0xcc, 0x3c, 0x09, 0x66, 0xb0, 0x28, 0x94, 0x27, 0x24, 0xc1, 0xc5,
	0x9e, 0x5a, 0x51, 0x90, 0x59, 0x94, 0x5a, 0x2c, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0xe3,
	0x1a, 0xb5, 0x33, 0x72, 0x71, 0xbb, 0xb8, 0xf8, 0x05, 0x07, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7,
	0x0a, 0x59, 0x71, 0x71, 0x05, 0x27, 0x96, 0xa5, 0x42, 0xed, 0xe4, 0xd6, 0x03, 0x39, 0x16, 0xc2,
	0x91, 0x42, 0xe6, 0x28, 0x09, 0x35, 0x5d, 0x7e, 0x32, 0x99, 0x89, 0x47, 0x89, 0x4b, 0xbf, 0xcc,
	0x50, 0xbf, 0x08, 0xa2, 0xda, 0x86, 0x8b, 0xc7, 0x25, 0x35, 0x27, 0xb5, 0x84, 0x68, 0xdd, 0x5a,
	0x48, 0xba, 0x9d, 0x58, 0xa3, 0x40, 0xc1, 0x91, 0xc4, 0x06, 0xf6, 0xb7, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x62, 0xd6, 0x20, 0x41, 0x2b, 0x01, 0x00, 0x00,
}
