// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb/client_tls.proto

package tlspb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TLSClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnString string `protobuf:"bytes,1,opt,name=conn_string,json=connString,proto3" json:"conn_string,omitempty"`
	SqlString  string `protobuf:"bytes,2,opt,name=sql_string,json=sqlString,proto3" json:"sql_string,omitempty"`
}

func (x *TLSClientRequest) Reset() {
	*x = TLSClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSClientRequest) ProtoMessage() {}

func (x *TLSClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSClientRequest.ProtoReflect.Descriptor instead.
func (*TLSClientRequest) Descriptor() ([]byte, []int) {
	return file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescGZIP(), []int{0}
}

func (x *TLSClientRequest) GetConnString() string {
	if x != nil {
		return x.ConnString
	}
	return ""
}

func (x *TLSClientRequest) GetSqlString() string {
	if x != nil {
		return x.SqlString
	}
	return ""
}

type TLSClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool   `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TLSClientResponse) Reset() {
	*x = TLSClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TLSClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSClientResponse) ProtoMessage() {}

func (x *TLSClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSClientResponse.ProtoReflect.Descriptor instead.
func (*TLSClientResponse) Descriptor() ([]byte, []int) {
	return file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescGZIP(), []int{1}
}

func (x *TLSClientResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *TLSClientResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *TLSClientResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto protoreflect.FileDescriptor

var file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDesc = []byte{
	0x0a, 0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72,
	0x73, 0x61, 0x61, 0x6b, 0x62, 0x61, 0x72, 0x69, 0x31, 0x32, 0x30, 0x39, 0x2f, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x2d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2d, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73, 0x70, 0x62, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x6c, 0x73, 0x70, 0x62, 0x22, 0x52, 0x0a, 0x10, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6e, 0x6e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x71,
	0x6c, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x71, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x5b, 0x0a, 0x11, 0x54, 0x4c, 0x53,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x5a, 0x0a, 0x10, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x54, 0x4c,
	0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x17, 0x2e, 0x74, 0x6c,
	0x73, 0x70, 0x62, 0x2e, 0x54, 0x4c, 0x53, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6c, 0x73, 0x70, 0x62, 0x2e, 0x54, 0x4c, 0x53,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x61, 0x72, 0x73, 0x61, 0x61, 0x6b, 0x62, 0x61, 0x72, 0x69, 0x31, 0x32, 0x30, 0x39,
	0x2f, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2d, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x6c, 0x73, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescOnce sync.Once
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescData = file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDesc
)

func file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescGZIP() []byte {
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescOnce.Do(func() {
		file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescData)
	})
	return file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDescData
}

var file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_goTypes = []interface{}{
	(*TLSClientRequest)(nil),  // 0: tlspb.TLSClientRequest
	(*TLSClientResponse)(nil), // 1: tlspb.TLSClientResponse
}
var file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_depIdxs = []int32{
	0, // 0: tlspb.TLSClientService.TLSClientSend:input_type -> tlspb.TLSClientRequest
	1, // 1: tlspb.TLSClientService.TLSClientSend:output_type -> tlspb.TLSClientResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_init()
}
func file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_init() {
	if File_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSClientRequest); i {
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
		file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TLSClientResponse); i {
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
			RawDescriptor: file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_goTypes,
		DependencyIndexes: file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_depIdxs,
		MessageInfos:      file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_msgTypes,
	}.Build()
	File_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto = out.File
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_rawDesc = nil
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_goTypes = nil
	file_github_com_parsaakbari1209_Database_Security_Proxy_client_application_tlspb_client_tls_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TLSClientServiceClient is the client API for TLSClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TLSClientServiceClient interface {
	// This rpc is for client application to call the tls service running on the client behalf.
	// The client side tls service will send the request to the database-secuirty-service securly.
	TLSClientSend(ctx context.Context, in *TLSClientRequest, opts ...grpc.CallOption) (TLSClientService_TLSClientSendClient, error)
}

type tLSClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTLSClientServiceClient(cc grpc.ClientConnInterface) TLSClientServiceClient {
	return &tLSClientServiceClient{cc}
}

func (c *tLSClientServiceClient) TLSClientSend(ctx context.Context, in *TLSClientRequest, opts ...grpc.CallOption) (TLSClientService_TLSClientSendClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TLSClientService_serviceDesc.Streams[0], "/tlspb.TLSClientService/TLSClientSend", opts...)
	if err != nil {
		return nil, err
	}
	x := &tLSClientServiceTLSClientSendClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TLSClientService_TLSClientSendClient interface {
	Recv() (*TLSClientResponse, error)
	grpc.ClientStream
}

type tLSClientServiceTLSClientSendClient struct {
	grpc.ClientStream
}

func (x *tLSClientServiceTLSClientSendClient) Recv() (*TLSClientResponse, error) {
	m := new(TLSClientResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TLSClientServiceServer is the server API for TLSClientService service.
type TLSClientServiceServer interface {
	// This rpc is for client application to call the tls service running on the client behalf.
	// The client side tls service will send the request to the database-secuirty-service securly.
	TLSClientSend(*TLSClientRequest, TLSClientService_TLSClientSendServer) error
}

// UnimplementedTLSClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTLSClientServiceServer struct {
}

func (*UnimplementedTLSClientServiceServer) TLSClientSend(*TLSClientRequest, TLSClientService_TLSClientSendServer) error {
	return status.Errorf(codes.Unimplemented, "method TLSClientSend not implemented")
}

func RegisterTLSClientServiceServer(s *grpc.Server, srv TLSClientServiceServer) {
	s.RegisterService(&_TLSClientService_serviceDesc, srv)
}

func _TLSClientService_TLSClientSend_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TLSClientRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TLSClientServiceServer).TLSClientSend(m, &tLSClientServiceTLSClientSendServer{stream})
}

type TLSClientService_TLSClientSendServer interface {
	Send(*TLSClientResponse) error
	grpc.ServerStream
}

type tLSClientServiceTLSClientSendServer struct {
	grpc.ServerStream
}

func (x *tLSClientServiceTLSClientSendServer) Send(m *TLSClientResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TLSClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tlspb.TLSClientService",
	HandlerType: (*TLSClientServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TLSClientSend",
			Handler:       _TLSClientService_TLSClientSend_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/parsaakbari1209/Database-Security-Proxy/client/application/tlspb/client_tls.proto",
}
