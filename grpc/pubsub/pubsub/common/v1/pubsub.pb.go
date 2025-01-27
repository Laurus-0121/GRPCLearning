// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.24.3
// source: pubsub.proto

package pubsub

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

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pubsub_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_pubsub_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_pubsub_proto_rawDescGZIP(), []int{0}
}

func (x *String) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pubsub_proto protoreflect.FileDescriptor

var file_pubsub_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x69, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x73, 0x75, 0x62,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x29, 0x0a, 0x07, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x12, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x0e, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x2d, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12,
	0x0e, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a,
	0x0e, 0x2e, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x30,
	0x01, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31,
	0x3b, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pubsub_proto_rawDescOnce sync.Once
	file_pubsub_proto_rawDescData = file_pubsub_proto_rawDesc
)

func file_pubsub_proto_rawDescGZIP() []byte {
	file_pubsub_proto_rawDescOnce.Do(func() {
		file_pubsub_proto_rawDescData = protoimpl.X.CompressGZIP(file_pubsub_proto_rawDescData)
	})
	return file_pubsub_proto_rawDescData
}

var file_pubsub_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pubsub_proto_goTypes = []interface{}{
	(*String)(nil), // 0: pubsub.String
}
var file_pubsub_proto_depIdxs = []int32{
	0, // 0: pubsub.PubsubService.Publish:input_type -> pubsub.String
	0, // 1: pubsub.PubsubService.Subscribe:input_type -> pubsub.String
	0, // 2: pubsub.PubsubService.Publish:output_type -> pubsub.String
	0, // 3: pubsub.PubsubService.Subscribe:output_type -> pubsub.String
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pubsub_proto_init() }
func file_pubsub_proto_init() {
	if File_pubsub_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pubsub_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
			RawDescriptor: file_pubsub_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pubsub_proto_goTypes,
		DependencyIndexes: file_pubsub_proto_depIdxs,
		MessageInfos:      file_pubsub_proto_msgTypes,
	}.Build()
	File_pubsub_proto = out.File
	file_pubsub_proto_rawDesc = nil
	file_pubsub_proto_goTypes = nil
	file_pubsub_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PubsubServiceClient is the clientSub API for PubsubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubsubServiceClient interface {
	Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error)
	Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error)
}

type pubsubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPubsubServiceClient(cc grpc.ClientConnInterface) PubsubServiceClient {
	return &pubsubServiceClient{cc}
}

func (c *pubsubServiceClient) Publish(ctx context.Context, in *String, opts ...grpc.CallOption) (*String, error) {
	out := new(String)
	err := c.cc.Invoke(ctx, "/pubsub.PubsubService/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubsubServiceClient) Subscribe(ctx context.Context, in *String, opts ...grpc.CallOption) (PubsubService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubsubService_serviceDesc.Streams[0], "/pubsub.PubsubService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubsubServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubsubService_SubscribeClient interface {
	Recv() (*String, error)
	grpc.ClientStream
}

type pubsubServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubsubServiceSubscribeClient) Recv() (*String, error) {
	m := new(String)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubsubServiceServer is the server API for PubsubService service.
type PubsubServiceServer interface {
	Publish(context.Context, *String) (*String, error)
	Subscribe(*String, PubsubService_SubscribeServer) error
}

// UnimplementedPubsubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPubsubServiceServer struct {
}

func (*UnimplementedPubsubServiceServer) Publish(context.Context, *String) (*String, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedPubsubServiceServer) Subscribe(*String, PubsubService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterPubsubServiceServer(s *grpc.Server, srv PubsubServiceServer) {
	s.RegisterService(&_PubsubService_serviceDesc, srv)
}

func _PubsubService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubsubServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pubsub.PubsubService/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubsubServiceServer).Publish(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubsubService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(String)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubsubServiceServer).Subscribe(m, &pubsubServiceSubscribeServer{stream})
}

type PubsubService_SubscribeServer interface {
	Send(*String) error
	grpc.ServerStream
}

type pubsubServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubsubServiceSubscribeServer) Send(m *String) error {
	return x.ServerStream.SendMsg(m)
}

var _PubsubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pubsub.PubsubService",
	HandlerType: (*PubsubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubsubService_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubsubService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pubsub.proto",
}
