// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: service_stocks.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_service_stocks_proto protoreflect.FileDescriptor

var file_service_stocks_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x13, 0x72, 0x70, 0x63, 0x5f,
	0x67, 0x65, 0x74, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb6, 0x01, 0x0a, 0x06, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x37, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x75, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6c, 0x69, 0x6e, 0x2d, 0x6d, 0x63, 0x6c, 0x2f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_service_stocks_proto_goTypes = []interface{}{
	(*GetQuoteRequest)(nil),    // 0: pb.GetQuoteRequest
	(*CreateUserRequest)(nil),  // 1: pb.CreateUserRequest
	(*GetUserRequest)(nil),     // 2: pb.GetUserRequest
	(*GetQuoteResponse)(nil),   // 3: pb.GetQuoteResponse
	(*CreateUserResponse)(nil), // 4: pb.CreateUserResponse
	(*GetUserResponse)(nil),    // 5: pb.GetUserResponse
}
var file_service_stocks_proto_depIdxs = []int32{
	0, // 0: pb.Stocks.GetQuote:input_type -> pb.GetQuoteRequest
	1, // 1: pb.Stocks.CreateUser:input_type -> pb.CreateUserRequest
	2, // 2: pb.Stocks.GetUser:input_type -> pb.GetUserRequest
	3, // 3: pb.Stocks.GetQuote:output_type -> pb.GetQuoteResponse
	4, // 4: pb.Stocks.CreateUser:output_type -> pb.CreateUserResponse
	5, // 5: pb.Stocks.GetUser:output_type -> pb.GetUserResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_service_stocks_proto_init() }
func file_service_stocks_proto_init() {
	if File_service_stocks_proto != nil {
		return
	}
	file_rpc_get_quote_proto_init()
	file_rpc_create_user_proto_init()
	file_rpc_get_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_stocks_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_stocks_proto_goTypes,
		DependencyIndexes: file_service_stocks_proto_depIdxs,
	}.Build()
	File_service_stocks_proto = out.File
	file_service_stocks_proto_rawDesc = nil
	file_service_stocks_proto_goTypes = nil
	file_service_stocks_proto_depIdxs = nil
}
