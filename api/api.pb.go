/*
 * Copyright 2019 American Express Travel Related Services Company, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 */

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	InputRequest
	OutputResponse
	HealthCheckRequest
	HealthCheckResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN     HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING     HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING HealthCheckResponse_ServingStatus = 2
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
}
var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":     0,
	"SERVING":     1,
	"NOT_SERVING": 2,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{3, 0}
}

// the serialized message that client sends
// each field in the message definition has a unique number used to identify fields in the message binary format, and should not be changed once your message type is in use. numbers in the range 1 through 15 take one byte to encode.
type InputRequest struct {
	Text       string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	ClientName string `protobuf:"bytes,2,opt,name=clientName" json:"clientName,omitempty"`
}

func (m *InputRequest) Reset()                    { *m = InputRequest{} }
func (m *InputRequest) String() string            { return proto.CompactTextString(m) }
func (*InputRequest) ProtoMessage()               {}
func (*InputRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *InputRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *InputRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

// what server responds to as a result of getting InputRequest
type OutputResponse struct {
	Text       string `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	ServerName string `protobuf:"bytes,2,opt,name=serverName" json:"serverName,omitempty"`
}

func (m *OutputResponse) Reset()                    { *m = OutputResponse{} }
func (m *OutputResponse) String() string            { return proto.CompactTextString(m) }
func (*OutputResponse) ProtoMessage()               {}
func (*OutputResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *OutputResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *OutputResponse) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

// health check request to find out the readiness
type HealthCheckRequest struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
}

func (m *HealthCheckRequest) Reset()                    { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()               {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

// health check response
type HealthCheckResponse struct {
	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,enum=api.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (m *HealthCheckResponse) Reset()                    { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string            { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()               {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterType((*InputRequest)(nil), "api.InputRequest")
	proto.RegisterType((*OutputResponse)(nil), "api.OutputResponse")
	proto.RegisterType((*HealthCheckRequest)(nil), "api.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "api.HealthCheckResponse")
	proto.RegisterEnum("api.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ProcessText service

type ProcessTextClient interface {
	Upper(ctx context.Context, in *InputRequest, opts ...grpc.CallOption) (*OutputResponse, error)
}

type processTextClient struct {
	cc *grpc.ClientConn
}

func NewProcessTextClient(cc *grpc.ClientConn) ProcessTextClient {
	return &processTextClient{cc}
}

func (c *processTextClient) Upper(ctx context.Context, in *InputRequest, opts ...grpc.CallOption) (*OutputResponse, error) {
	out := new(OutputResponse)
	err := grpc.Invoke(ctx, "/api.ProcessText/upper", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProcessText service

type ProcessTextServer interface {
	Upper(context.Context, *InputRequest) (*OutputResponse, error)
}

func RegisterProcessTextServer(s *grpc.Server, srv ProcessTextServer) {
	s.RegisterService(&_ProcessText_serviceDesc, srv)
}

func _ProcessText_Upper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessTextServer).Upper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProcessText/Upper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessTextServer).Upper(ctx, req.(*InputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessText_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProcessText",
	HandlerType: (*ProcessTextServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "upper",
			Handler:    _ProcessText_Upper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for Health service

type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc *grpc.ClientConn
}

func NewHealthClient(cc *grpc.ClientConn) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/api.Health/Check", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := grpc.Invoke(ctx, "/api.Health/Watch", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Health service

type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Watch(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Watch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Watch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Health/Watch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Watch(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "Watch",
			Handler:    _Health_Watch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 301 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x6c, 0xaa, 0x4d, 0xe9, 0x8b, 0xd6, 0xba, 0x3d, 0x18, 0x3c, 0x88, 0xec, 0x41, 0x3c, 0x05,
	0x8c, 0xb7, 0x1e, 0x44, 0xfc, 0x40, 0x8b, 0xb0, 0x91, 0xa4, 0xda, 0xa3, 0xac, 0xe1, 0x61, 0x82,
	0x35, 0x59, 0xb3, 0x1b, 0xf1, 0xe8, 0x5f, 0xf0, 0x1f, 0xcb, 0x6e, 0x13, 0x4c, 0xd0, 0x5e, 0x7a,
	0xcb, 0x9b, 0x37, 0x33, 0x99, 0x79, 0x2c, 0x0c, 0xb8, 0x48, 0x3d, 0x51, 0xe4, 0x2a, 0x27, 0x1b,
	0x5c, 0xa4, 0xf4, 0x02, 0xb6, 0xa6, 0x99, 0x28, 0x55, 0x88, 0xef, 0x25, 0x4a, 0x45, 0x08, 0x6c,
	0x2a, 0xfc, 0x54, 0xae, 0x75, 0x68, 0x1d, 0x0f, 0x42, 0xf3, 0x4d, 0x0e, 0x00, 0xe2, 0x45, 0x8a,
	0x99, 0x62, 0xfc, 0x0d, 0xdd, 0xae, 0xd9, 0x34, 0x10, 0x7a, 0x05, 0xc3, 0xa0, 0x54, 0xc6, 0x44,
	0x8a, 0x3c, 0x93, 0xb8, 0xca, 0x45, 0x62, 0xf1, 0x81, 0x45, 0xd3, 0xe5, 0x17, 0xa1, 0x1e, 0x90,
	0x5b, 0xe4, 0x0b, 0x95, 0x5c, 0x26, 0x18, 0xbf, 0xd6, 0x79, 0x5c, 0xe8, 0x6b, 0x4e, 0x1a, 0x63,
	0x65, 0x56, 0x8f, 0xf4, 0xdb, 0x82, 0x71, 0x4b, 0x50, 0xfd, 0xfb, 0x0c, 0x6c, 0xa9, 0xb8, 0x2a,
	0xa5, 0x11, 0x0c, 0xfd, 0x23, 0x4f, 0x57, 0xfe, 0x87, 0xe9, 0x45, 0xda, 0x29, 0x7b, 0x89, 0x0c,
	0x3b, 0xac, 0x54, 0x74, 0x02, 0xdb, 0xad, 0x05, 0x71, 0xa0, 0xff, 0xc0, 0xee, 0x58, 0x30, 0x67,
	0xa3, 0x8e, 0x1e, 0xa2, 0xeb, 0xf0, 0x71, 0xca, 0x6e, 0x46, 0x16, 0xd9, 0x01, 0x87, 0x05, 0xb3,
	0xa7, 0x1a, 0xe8, 0xfa, 0xe7, 0xe0, 0xdc, 0x17, 0x79, 0x8c, 0x52, 0xce, 0x74, 0xe5, 0x13, 0xe8,
	0x95, 0x42, 0x60, 0x41, 0x76, 0x4d, 0x86, 0xe6, 0xa1, 0xf7, 0xc7, 0x06, 0x6a, 0xdf, 0x8d, 0x76,
	0xfc, 0x2f, 0x0b, 0xec, 0x65, 0x56, 0x32, 0x81, 0x9e, 0xc9, 0x4b, 0xf6, 0xfe, 0x36, 0x58, 0x7a,
	0xb8, 0xab, 0xaa, 0x69, 0xed, 0x9c, 0xab, 0x38, 0x59, 0x43, 0xfb, 0x6c, 0x9b, 0xe7, 0x71, 0xfa,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x41, 0xc8, 0x82, 0x2b, 0x02, 0x00, 0x00,
}
