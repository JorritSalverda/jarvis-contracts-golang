// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: jarvis/contracts/v1/aggregation_level.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
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

// At what level of aggregation is the sample recorded.
type AggregationLevel int32

const (
	AggregationLevel_AGGREGATION_LEVEL_INVALID    AggregationLevel = 0
	AggregationLevel_AGGREGATION_LEVEL_METER      AggregationLevel = 1
	AggregationLevel_AGGREGATION_LEVEL_GROUP      AggregationLevel = 2
	AggregationLevel_AGGREGATION_LEVEL_DEVICE     AggregationLevel = 3
	AggregationLevel_AGGREGATION_LEVEL_DEVICE_SUB AggregationLevel = 4
)

// Enum value maps for AggregationLevel.
var (
	AggregationLevel_name = map[int32]string{
		0: "AGGREGATION_LEVEL_INVALID",
		1: "AGGREGATION_LEVEL_METER",
		2: "AGGREGATION_LEVEL_GROUP",
		3: "AGGREGATION_LEVEL_DEVICE",
		4: "AGGREGATION_LEVEL_DEVICE_SUB",
	}
	AggregationLevel_value = map[string]int32{
		"AGGREGATION_LEVEL_INVALID":    0,
		"AGGREGATION_LEVEL_METER":      1,
		"AGGREGATION_LEVEL_GROUP":      2,
		"AGGREGATION_LEVEL_DEVICE":     3,
		"AGGREGATION_LEVEL_DEVICE_SUB": 4,
	}
)

func (x AggregationLevel) Enum() *AggregationLevel {
	p := new(AggregationLevel)
	*p = x
	return p
}

func (x AggregationLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AggregationLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_jarvis_contracts_v1_aggregation_level_proto_enumTypes[0].Descriptor()
}

func (AggregationLevel) Type() protoreflect.EnumType {
	return &file_jarvis_contracts_v1_aggregation_level_proto_enumTypes[0]
}

func (x AggregationLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AggregationLevel.Descriptor instead.
func (AggregationLevel) EnumDescriptor() ([]byte, []int) {
	return file_jarvis_contracts_v1_aggregation_level_proto_rawDescGZIP(), []int{0}
}

var File_jarvis_contracts_v1_aggregation_level_proto protoreflect.FileDescriptor

var file_jarvis_contracts_v1_aggregation_level_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6a,
	0x61, 0x72, 0x76, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2a, 0xab, 0x01, 0x0a, 0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1d, 0x0a, 0x19, 0x41, 0x47, 0x47, 0x52, 0x45,
	0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47,
	0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x4d, 0x45, 0x54, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x47, 0x52, 0x4f, 0x55, 0x50, 0x10, 0x02,
	0x12, 0x1c, 0x0a, 0x18, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x10, 0x03, 0x12, 0x20,
	0x0a, 0x1c, 0x41, 0x47, 0x47, 0x52, 0x45, 0x47, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x55, 0x42, 0x10, 0x04,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a,
	0x6f, 0x72, 0x72, 0x69, 0x74, 0x53, 0x61, 0x6c, 0x76, 0x65, 0x72, 0x64, 0x61, 0x2f, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2d, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jarvis_contracts_v1_aggregation_level_proto_rawDescOnce sync.Once
	file_jarvis_contracts_v1_aggregation_level_proto_rawDescData = file_jarvis_contracts_v1_aggregation_level_proto_rawDesc
)

func file_jarvis_contracts_v1_aggregation_level_proto_rawDescGZIP() []byte {
	file_jarvis_contracts_v1_aggregation_level_proto_rawDescOnce.Do(func() {
		file_jarvis_contracts_v1_aggregation_level_proto_rawDescData = protoimpl.X.CompressGZIP(file_jarvis_contracts_v1_aggregation_level_proto_rawDescData)
	})
	return file_jarvis_contracts_v1_aggregation_level_proto_rawDescData
}

var file_jarvis_contracts_v1_aggregation_level_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_jarvis_contracts_v1_aggregation_level_proto_goTypes = []interface{}{
	(AggregationLevel)(0), // 0: jarvis.contracts.v1.AggregationLevel
}
var file_jarvis_contracts_v1_aggregation_level_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_jarvis_contracts_v1_aggregation_level_proto_init() }
func file_jarvis_contracts_v1_aggregation_level_proto_init() {
	if File_jarvis_contracts_v1_aggregation_level_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_jarvis_contracts_v1_aggregation_level_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jarvis_contracts_v1_aggregation_level_proto_goTypes,
		DependencyIndexes: file_jarvis_contracts_v1_aggregation_level_proto_depIdxs,
		EnumInfos:         file_jarvis_contracts_v1_aggregation_level_proto_enumTypes,
	}.Build()
	File_jarvis_contracts_v1_aggregation_level_proto = out.File
	file_jarvis_contracts_v1_aggregation_level_proto_rawDesc = nil
	file_jarvis_contracts_v1_aggregation_level_proto_goTypes = nil
	file_jarvis_contracts_v1_aggregation_level_proto_depIdxs = nil
}
