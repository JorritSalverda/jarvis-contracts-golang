// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: jarvis/contracts/v1/measurement.proto

package v1

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// A measurements is a range of samples read at the same time from the same source.
type Measurement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source         string               `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Location       string               `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	Samples        []*Sample            `protobuf:"bytes,3,rep,name=samples,proto3" json:"samples,omitempty"`
	MeasuredAtTime *timestamp.Timestamp `protobuf:"bytes,4,opt,name=measured_at_time,json=measuredAtTime,proto3" json:"measured_at_time,omitempty"`
}

func (x *Measurement) Reset() {
	*x = Measurement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jarvis_contracts_v1_measurement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Measurement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Measurement) ProtoMessage() {}

func (x *Measurement) ProtoReflect() protoreflect.Message {
	mi := &file_jarvis_contracts_v1_measurement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Measurement.ProtoReflect.Descriptor instead.
func (*Measurement) Descriptor() ([]byte, []int) {
	return file_jarvis_contracts_v1_measurement_proto_rawDescGZIP(), []int{0}
}

func (x *Measurement) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Measurement) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Measurement) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

func (x *Measurement) GetMeasuredAtTime() *timestamp.Timestamp {
	if x != nil {
		return x.MeasuredAtTime
	}
	return nil
}

var File_jarvis_contracts_v1_measurement_proto protoreflect.FileDescriptor

var file_jarvis_contracts_v1_measurement_proto_rawDesc = []byte{
	0x0a, 0x25, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbe, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6a, 0x61, 0x72, 0x76, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x10, 0x6d, 0x65,
	0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0e, 0x6d, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x64, 0x41, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a,
	0x6f, 0x72, 0x72, 0x69, 0x74, 0x53, 0x61, 0x6c, 0x76, 0x65, 0x72, 0x64, 0x61, 0x2f, 0x6a, 0x61,
	0x72, 0x76, 0x69, 0x73, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2d, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jarvis_contracts_v1_measurement_proto_rawDescOnce sync.Once
	file_jarvis_contracts_v1_measurement_proto_rawDescData = file_jarvis_contracts_v1_measurement_proto_rawDesc
)

func file_jarvis_contracts_v1_measurement_proto_rawDescGZIP() []byte {
	file_jarvis_contracts_v1_measurement_proto_rawDescOnce.Do(func() {
		file_jarvis_contracts_v1_measurement_proto_rawDescData = protoimpl.X.CompressGZIP(file_jarvis_contracts_v1_measurement_proto_rawDescData)
	})
	return file_jarvis_contracts_v1_measurement_proto_rawDescData
}

var file_jarvis_contracts_v1_measurement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_jarvis_contracts_v1_measurement_proto_goTypes = []interface{}{
	(*Measurement)(nil),         // 0: jarvis.contracts.v1.Measurement
	(*Sample)(nil),              // 1: jarvis.contracts.v1.Sample
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_jarvis_contracts_v1_measurement_proto_depIdxs = []int32{
	1, // 0: jarvis.contracts.v1.Measurement.samples:type_name -> jarvis.contracts.v1.Sample
	2, // 1: jarvis.contracts.v1.Measurement.measured_at_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_jarvis_contracts_v1_measurement_proto_init() }
func file_jarvis_contracts_v1_measurement_proto_init() {
	if File_jarvis_contracts_v1_measurement_proto != nil {
		return
	}
	file_jarvis_contracts_v1_sample_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_jarvis_contracts_v1_measurement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Measurement); i {
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
			RawDescriptor: file_jarvis_contracts_v1_measurement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jarvis_contracts_v1_measurement_proto_goTypes,
		DependencyIndexes: file_jarvis_contracts_v1_measurement_proto_depIdxs,
		MessageInfos:      file_jarvis_contracts_v1_measurement_proto_msgTypes,
	}.Build()
	File_jarvis_contracts_v1_measurement_proto = out.File
	file_jarvis_contracts_v1_measurement_proto_rawDesc = nil
	file_jarvis_contracts_v1_measurement_proto_goTypes = nil
	file_jarvis_contracts_v1_measurement_proto_depIdxs = nil
}
