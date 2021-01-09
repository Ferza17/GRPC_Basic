// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: calculator/calculatorpb/calculator.proto

package calculatorpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum1 int32 `protobuf:"varint,1,opt,name=sum1,proto3" json:"sum1,omitempty"`
	Sum2 int32 `protobuf:"varint,2,opt,name=sum2,proto3" json:"sum2,omitempty"`
}

func (x *Sum) Reset() {
	*x = Sum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sum) ProtoMessage() {}

func (x *Sum) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sum.ProtoReflect.Descriptor instead.
func (*Sum) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *Sum) GetSum1() int32 {
	if x != nil {
		return x.Sum1
	}
	return 0
}

func (x *Sum) GetSum2() int32 {
	if x != nil {
		return x.Sum2
	}
	return 0
}

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sum *Sum `protobuf:"bytes,1,opt,name=sum,proto3" json:"sum,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *SumRequest) GetSum() *Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

type SumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumResponse) Reset() {
	*x = SumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumResponse) ProtoMessage() {}

func (x *SumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumResponse.ProtoReflect.Descriptor instead.
func (*SumResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *SumResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

type SumManyTimesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *SumManyTimesRequest) Reset() {
	*x = SumManyTimesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumManyTimesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumManyTimesRequest) ProtoMessage() {}

func (x *SumManyTimesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumManyTimesRequest.ProtoReflect.Descriptor instead.
func (*SumManyTimesRequest) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *SumManyTimesRequest) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

type SumManyTimesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int32 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SumManyTimesResponse) Reset() {
	*x = SumManyTimesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumManyTimesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumManyTimesResponse) ProtoMessage() {}

func (x *SumManyTimesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calculator_calculatorpb_calculator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumManyTimesResponse.ProtoReflect.Descriptor instead.
func (*SumManyTimesResponse) Descriptor() ([]byte, []int) {
	return file_calculator_calculatorpb_calculator_proto_rawDescGZIP(), []int{4}
}

func (x *SumManyTimesResponse) GetResult() int32 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calculator_calculatorpb_calculator_proto protoreflect.FileDescriptor

var file_calculator_calculatorpb_calculator_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x2d, 0x0a, 0x03, 0x53, 0x75, 0x6d, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x75, 0x6d, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x75, 0x6d,
	0x31, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x75, 0x6d, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x73, 0x75, 0x6d, 0x32, 0x22, 0x2f, 0x0a, 0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75,
	0x6d, 0x52, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x25, 0x0a, 0x0b, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2b, 0x0a,
	0x13, 0x53, 0x75, 0x6d, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x2e, 0x0a, 0x14, 0x53, 0x75,
	0x6d, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xa1, 0x01, 0x0a, 0x0a, 0x53,
	0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x53, 0x75, 0x6d,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x75, 0x6d, 0x4d, 0x61,
	0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75,
	0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x53, 0x75, 0x6d, 0x4d, 0x61, 0x6e, 0x79, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x19,
	0x5a, 0x17, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c,
	0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_calculator_calculatorpb_calculator_proto_rawDescOnce sync.Once
	file_calculator_calculatorpb_calculator_proto_rawDescData = file_calculator_calculatorpb_calculator_proto_rawDesc
)

func file_calculator_calculatorpb_calculator_proto_rawDescGZIP() []byte {
	file_calculator_calculatorpb_calculator_proto_rawDescOnce.Do(func() {
		file_calculator_calculatorpb_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calculator_calculatorpb_calculator_proto_rawDescData)
	})
	return file_calculator_calculatorpb_calculator_proto_rawDescData
}

var file_calculator_calculatorpb_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_calculator_calculatorpb_calculator_proto_goTypes = []interface{}{
	(*Sum)(nil),                  // 0: calculator.Sum
	(*SumRequest)(nil),           // 1: calculator.SumRequest
	(*SumResponse)(nil),          // 2: calculator.SumResponse
	(*SumManyTimesRequest)(nil),  // 3: calculator.SumManyTimesRequest
	(*SumManyTimesResponse)(nil), // 4: calculator.SumManyTimesResponse
}
var file_calculator_calculatorpb_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.SumRequest.sum:type_name -> calculator.Sum
	1, // 1: calculator.SumService.SumData:input_type -> calculator.SumRequest
	3, // 2: calculator.SumService.SumManyTimes:input_type -> calculator.SumManyTimesRequest
	2, // 3: calculator.SumService.SumData:output_type -> calculator.SumResponse
	4, // 4: calculator.SumService.SumManyTimes:output_type -> calculator.SumManyTimesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_calculator_calculatorpb_calculator_proto_init() }
func file_calculator_calculatorpb_calculator_proto_init() {
	if File_calculator_calculatorpb_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calculator_calculatorpb_calculator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumManyTimesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_calculator_calculatorpb_calculator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumManyTimesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_calculator_calculatorpb_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calculator_calculatorpb_calculator_proto_goTypes,
		DependencyIndexes: file_calculator_calculatorpb_calculator_proto_depIdxs,
		MessageInfos:      file_calculator_calculatorpb_calculator_proto_msgTypes,
	}.Build()
	File_calculator_calculatorpb_calculator_proto = out.File
	file_calculator_calculatorpb_calculator_proto_rawDesc = nil
	file_calculator_calculatorpb_calculator_proto_goTypes = nil
	file_calculator_calculatorpb_calculator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SumServiceClient is the client API for SumService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SumServiceClient interface {
	// UNARY API
	SumData(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	//Server Streaming API
	SumManyTimes(ctx context.Context, in *SumManyTimesRequest, opts ...grpc.CallOption) (SumService_SumManyTimesClient, error)
}

type sumServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSumServiceClient(cc grpc.ClientConnInterface) SumServiceClient {
	return &sumServiceClient{cc}
}

func (c *sumServiceClient) SumData(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.SumService/SumData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sumServiceClient) SumManyTimes(ctx context.Context, in *SumManyTimesRequest, opts ...grpc.CallOption) (SumService_SumManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SumService_serviceDesc.Streams[0], "/calculator.SumService/SumManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumServiceSumManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SumService_SumManyTimesClient interface {
	Recv() (*SumManyTimesResponse, error)
	grpc.ClientStream
}

type sumServiceSumManyTimesClient struct {
	grpc.ClientStream
}

func (x *sumServiceSumManyTimesClient) Recv() (*SumManyTimesResponse, error) {
	m := new(SumManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SumServiceServer is the server API for SumService service.
type SumServiceServer interface {
	// UNARY API
	SumData(context.Context, *SumRequest) (*SumResponse, error)
	//Server Streaming API
	SumManyTimes(*SumManyTimesRequest, SumService_SumManyTimesServer) error
}

// UnimplementedSumServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSumServiceServer struct {
}

func (*UnimplementedSumServiceServer) SumData(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SumData not implemented")
}
func (*UnimplementedSumServiceServer) SumManyTimes(*SumManyTimesRequest, SumService_SumManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method SumManyTimes not implemented")
}

func RegisterSumServiceServer(s *grpc.Server, srv SumServiceServer) {
	s.RegisterService(&_SumService_serviceDesc, srv)
}

func _SumService_SumData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SumServiceServer).SumData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.SumService/SumData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SumServiceServer).SumData(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SumService_SumManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SumManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SumServiceServer).SumManyTimes(m, &sumServiceSumManyTimesServer{stream})
}

type SumService_SumManyTimesServer interface {
	Send(*SumManyTimesResponse) error
	grpc.ServerStream
}

type sumServiceSumManyTimesServer struct {
	grpc.ServerStream
}

func (x *sumServiceSumManyTimesServer) Send(m *SumManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _SumService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.SumService",
	HandlerType: (*SumServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SumData",
			Handler:    _SumService_SumData_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SumManyTimes",
			Handler:       _SumService_SumManyTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
