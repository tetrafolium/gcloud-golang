// Code generated by protoc-gen-go.
// source: github.com/tetrafolium/gcloud-golang/bigtable/internal/service_proto/bigtable_service.proto
// DO NOT EDIT!

package google_bigtable_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_bigtable_v11 "github.com/tetrafolium/gcloud-golang/bigtable/internal/data_proto"
import google_protobuf "github.com/tetrafolium/gcloud-golang/bigtable/internal/empty"

import (
	context "golang.org/x/net/context"
	grpc "github.com/tetrafolium/grpc-go"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for BigtableService service

type BigtableServiceClient interface {
	// Streams back the contents of all requested rows, optionally applying
	// the same Reader filter to each. Depending on their size, rows may be
	// broken up across multiple responses, but atomicity of each row will still
	// be preserved.
	ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (BigtableService_ReadRowsClient, error)
	// Returns a sample of row keys in the table. The returned row keys will
	// delimit contiguous sections of the table of approximately equal size,
	// which can be used to break up the data for distributed tasks like
	// mapreduces.
	SampleRowKeys(ctx context.Context, in *SampleRowKeysRequest, opts ...grpc.CallOption) (BigtableService_SampleRowKeysClient, error)
	// Mutates a row atomically. Cells already present in the row are left
	// unchanged unless explicitly changed by 'mutation'.
	MutateRow(ctx context.Context, in *MutateRowRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	// Mutates a row atomically based on the output of a predicate Reader filter.
	CheckAndMutateRow(ctx context.Context, in *CheckAndMutateRowRequest, opts ...grpc.CallOption) (*CheckAndMutateRowResponse, error)
	// Modifies a row atomically, reading the latest existing timestamp/value from
	// the specified columns and writing a new value at
	// max(existing timestamp, current server time) based on pre-defined
	// read/modify/write rules. Returns the new contents of all modified cells.
	ReadModifyWriteRow(ctx context.Context, in *ReadModifyWriteRowRequest, opts ...grpc.CallOption) (*google_bigtable_v11.Row, error)
}

type bigtableServiceClient struct {
	cc *grpc.ClientConn
}

func NewBigtableServiceClient(cc *grpc.ClientConn) BigtableServiceClient {
	return &bigtableServiceClient{cc}
}

func (c *bigtableServiceClient) ReadRows(ctx context.Context, in *ReadRowsRequest, opts ...grpc.CallOption) (BigtableService_ReadRowsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BigtableService_serviceDesc.Streams[0], c.cc, "/google.bigtable.v1.BigtableService/ReadRows", opts...)
	if err != nil {
		return nil, err
	}
	x := &bigtableServiceReadRowsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BigtableService_ReadRowsClient interface {
	Recv() (*ReadRowsResponse, error)
	grpc.ClientStream
}

type bigtableServiceReadRowsClient struct {
	grpc.ClientStream
}

func (x *bigtableServiceReadRowsClient) Recv() (*ReadRowsResponse, error) {
	m := new(ReadRowsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bigtableServiceClient) SampleRowKeys(ctx context.Context, in *SampleRowKeysRequest, opts ...grpc.CallOption) (BigtableService_SampleRowKeysClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_BigtableService_serviceDesc.Streams[1], c.cc, "/google.bigtable.v1.BigtableService/SampleRowKeys", opts...)
	if err != nil {
		return nil, err
	}
	x := &bigtableServiceSampleRowKeysClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BigtableService_SampleRowKeysClient interface {
	Recv() (*SampleRowKeysResponse, error)
	grpc.ClientStream
}

type bigtableServiceSampleRowKeysClient struct {
	grpc.ClientStream
}

func (x *bigtableServiceSampleRowKeysClient) Recv() (*SampleRowKeysResponse, error) {
	m := new(SampleRowKeysResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bigtableServiceClient) MutateRow(ctx context.Context, in *MutateRowRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/google.bigtable.v1.BigtableService/MutateRow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableServiceClient) CheckAndMutateRow(ctx context.Context, in *CheckAndMutateRowRequest, opts ...grpc.CallOption) (*CheckAndMutateRowResponse, error) {
	out := new(CheckAndMutateRowResponse)
	err := grpc.Invoke(ctx, "/google.bigtable.v1.BigtableService/CheckAndMutateRow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bigtableServiceClient) ReadModifyWriteRow(ctx context.Context, in *ReadModifyWriteRowRequest, opts ...grpc.CallOption) (*google_bigtable_v11.Row, error) {
	out := new(google_bigtable_v11.Row)
	err := grpc.Invoke(ctx, "/google.bigtable.v1.BigtableService/ReadModifyWriteRow", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BigtableService service

type BigtableServiceServer interface {
	// Streams back the contents of all requested rows, optionally applying
	// the same Reader filter to each. Depending on their size, rows may be
	// broken up across multiple responses, but atomicity of each row will still
	// be preserved.
	ReadRows(*ReadRowsRequest, BigtableService_ReadRowsServer) error
	// Returns a sample of row keys in the table. The returned row keys will
	// delimit contiguous sections of the table of approximately equal size,
	// which can be used to break up the data for distributed tasks like
	// mapreduces.
	SampleRowKeys(*SampleRowKeysRequest, BigtableService_SampleRowKeysServer) error
	// Mutates a row atomically. Cells already present in the row are left
	// unchanged unless explicitly changed by 'mutation'.
	MutateRow(context.Context, *MutateRowRequest) (*google_protobuf.Empty, error)
	// Mutates a row atomically based on the output of a predicate Reader filter.
	CheckAndMutateRow(context.Context, *CheckAndMutateRowRequest) (*CheckAndMutateRowResponse, error)
	// Modifies a row atomically, reading the latest existing timestamp/value from
	// the specified columns and writing a new value at
	// max(existing timestamp, current server time) based on pre-defined
	// read/modify/write rules. Returns the new contents of all modified cells.
	ReadModifyWriteRow(context.Context, *ReadModifyWriteRowRequest) (*google_bigtable_v11.Row, error)
}

func RegisterBigtableServiceServer(s *grpc.Server, srv BigtableServiceServer) {
	s.RegisterService(&_BigtableService_serviceDesc, srv)
}

func _BigtableService_ReadRows_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReadRowsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BigtableServiceServer).ReadRows(m, &bigtableServiceReadRowsServer{stream})
}

type BigtableService_ReadRowsServer interface {
	Send(*ReadRowsResponse) error
	grpc.ServerStream
}

type bigtableServiceReadRowsServer struct {
	grpc.ServerStream
}

func (x *bigtableServiceReadRowsServer) Send(m *ReadRowsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BigtableService_SampleRowKeys_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SampleRowKeysRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BigtableServiceServer).SampleRowKeys(m, &bigtableServiceSampleRowKeysServer{stream})
}

type BigtableService_SampleRowKeysServer interface {
	Send(*SampleRowKeysResponse) error
	grpc.ServerStream
}

type bigtableServiceSampleRowKeysServer struct {
	grpc.ServerStream
}

func (x *bigtableServiceSampleRowKeysServer) Send(m *SampleRowKeysResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BigtableService_MutateRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(MutateRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableServiceServer).MutateRow(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableService_CheckAndMutateRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CheckAndMutateRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableServiceServer).CheckAndMutateRow(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _BigtableService_ReadModifyWriteRow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ReadModifyWriteRowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(BigtableServiceServer).ReadModifyWriteRow(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _BigtableService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.bigtable.v1.BigtableService",
	HandlerType: (*BigtableServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MutateRow",
			Handler:    _BigtableService_MutateRow_Handler,
		},
		{
			MethodName: "CheckAndMutateRow",
			Handler:    _BigtableService_CheckAndMutateRow_Handler,
		},
		{
			MethodName: "ReadModifyWriteRow",
			Handler:    _BigtableService_ReadModifyWriteRow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ReadRows",
			Handler:       _BigtableService_ReadRows_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SampleRowKeys",
			Handler:       _BigtableService_SampleRowKeys_Handler,
			ServerStreams: true,
		},
	},
}

var fileDescriptor1 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x1b, 0x04, 0xd1, 0x05, 0x11, 0x17, 0xac, 0xd0, 0x63, 0xf5, 0xa0, 0x60, 0x37, 0xfe,
	0xbb, 0x79, 0xb2, 0x22, 0x08, 0xa5, 0x52, 0xda, 0x43, 0xf1, 0x62, 0xd9, 0x24, 0xd3, 0x6d, 0x70,
	0x93, 0x89, 0xbb, 0x9b, 0x96, 0x7e, 0x03, 0xbf, 0xb1, 0x57, 0x4d, 0x93, 0xad, 0xb6, 0x8d, 0x92,
	0x83, 0x97, 0x1c, 0xe6, 0xfd, 0xe6, 0xbd, 0xc9, 0x63, 0xc9, 0x93, 0x40, 0x14, 0x12, 0x98, 0x40,
	0xc9, 0x63, 0xc1, 0x50, 0x09, 0xd7, 0x97, 0x98, 0x06, 0xae, 0x17, 0x0a, 0xc3, 0x3d, 0x09, 0x6e,
	0x18, 0x1b, 0x50, 0x31, 0x97, 0xae, 0x06, 0x35, 0x0d, 0x7d, 0x18, 0x25, 0x0a, 0x0d, 0x2e, 0xf5,
	0x51, 0x31, 0x66, 0x8b, 0x31, 0xa5, 0x85, 0x9f, 0x95, 0xd9, 0xf4, 0xb2, 0xf1, 0x58, 0x3d, 0x23,
	0xe0, 0x86, 0xaf, 0x07, 0x64, 0xb3, 0xdc, 0xbd, 0x31, 0xfc, 0xaf, 0x6b, 0x47, 0x11, 0x68, 0xcd,
	0x05, 0xe8, 0xc2, 0xf8, 0xb6, 0xba, 0x31, 0x44, 0x89, 0x99, 0xe7, 0xdf, 0x7c, 0xf9, 0xea, 0x63,
	0x8b, 0xec, 0xb7, 0x0b, 0x6e, 0x90, 0xfb, 0xd3, 0x67, 0xb2, 0xd3, 0x07, 0x1e, 0xf4, 0x71, 0xa6,
	0xe9, 0x31, 0xdb, 0x2c, 0x85, 0x59, 0xb5, 0x0f, 0x6f, 0x29, 0x68, 0xd3, 0x38, 0xf9, 0x1b, 0xd2,
	0x09, 0xc6, 0x1a, 0x9a, 0xb5, 0x0b, 0x87, 0x4e, 0xc8, 0xde, 0x80, 0x47, 0x89, 0x84, 0x2f, 0xa5,
	0x03, 0x73, 0x4d, 0x4f, 0xcb, 0x56, 0x57, 0x10, 0x1b, 0x72, 0x56, 0x81, 0xfc, 0x91, 0xd4, 0x21,
	0xbb, 0xdd, 0xd4, 0x70, 0x93, 0x89, 0xb4, 0xf4, 0xc0, 0xa5, 0x6c, 0x13, 0xea, 0x96, 0x5a, 0x54,
	0xe3, 0xa5, 0x63, 0xf6, 0x90, 0x35, 0xd5, 0xac, 0x51, 0x45, 0x0e, 0xee, 0x27, 0xe0, 0xbf, 0xde,
	0xc5, 0xc1, 0xb7, 0xe9, 0x79, 0x99, 0xe9, 0x06, 0x66, 0xcd, 0x5b, 0x15, 0x69, 0xfb, 0x0b, 0xf4,
	0x85, 0xd0, 0xac, 0xc2, 0x2e, 0x06, 0xe1, 0x78, 0x3e, 0x54, 0x61, 0x1e, 0xda, 0xfa, 0xad, 0xea,
	0x55, 0xce, 0xa6, 0x1e, 0x95, 0xe2, 0x38, 0x6b, 0xd6, 0xda, 0x37, 0xa4, 0xee, 0x63, 0x54, 0xa2,
	0xb7, 0x0f, 0xd7, 0x1e, 0x84, 0xee, 0x65, 0x7d, 0xf4, 0x9c, 0x77, 0xc7, 0xf1, 0xb6, 0x17, 0xdd,
	0x5c, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x63, 0x50, 0x8b, 0x8e, 0x7c, 0x03, 0x00, 0x00,
}
