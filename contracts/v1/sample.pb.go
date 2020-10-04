// Code generated by protoc-gen-go. DO NOT EDIT.
// source: jarvis/contracts/v1/sample.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A sample represents a single measured value with with a measurement; a measurement can have multiple samples.
type Sample struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DisplayName          string           `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	AggregationLevel     AggregationLevel `protobuf:"varint,3,opt,name=aggregation_level,json=aggregationLevel,proto3,enum=jarvis.contracts.v1.AggregationLevel" json:"aggregation_level,omitempty"`
	MetricType           MetricType       `protobuf:"varint,4,opt,name=metric_type,json=metricType,proto3,enum=jarvis.contracts.v1.MetricType" json:"metric_type,omitempty"`
	Value                float64          `protobuf:"fixed64,5,opt,name=value,proto3" json:"value,omitempty"`
	SampleType           SampleType       `protobuf:"varint,6,opt,name=sample_type,json=sampleType,proto3,enum=jarvis.contracts.v1.SampleType" json:"sample_type,omitempty"`
	SampleUnit           SampleUnit       `protobuf:"varint,7,opt,name=sample_unit,json=sampleUnit,proto3,enum=jarvis.contracts.v1.SampleUnit" json:"sample_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_aeafe98fcaf65e33, []int{0}
}

func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sample) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Sample) GetAggregationLevel() AggregationLevel {
	if m != nil {
		return m.AggregationLevel
	}
	return AggregationLevel_AGGREGATION_LEVEL_INVALID
}

func (m *Sample) GetMetricType() MetricType {
	if m != nil {
		return m.MetricType
	}
	return MetricType_METRIC_TYPE_INVALID
}

func (m *Sample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sample) GetSampleType() SampleType {
	if m != nil {
		return m.SampleType
	}
	return SampleType_SAMPLE_TYPE_INVALID
}

func (m *Sample) GetSampleUnit() SampleUnit {
	if m != nil {
		return m.SampleUnit
	}
	return SampleUnit_SAMPLE_UNIT_INVALID
}

func init() {
	proto.RegisterType((*Sample)(nil), "jarvis.contracts.v1.Sample")
}

func init() { proto.RegisterFile("jarvis/contracts/v1/sample.proto", fileDescriptor_aeafe98fcaf65e33) }

var fileDescriptor_aeafe98fcaf65e33 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xc1, 0x4b, 0xfb, 0x30,
	0x14, 0xc7, 0xe9, 0x7e, 0xdb, 0x7e, 0x98, 0x89, 0x68, 0xf4, 0x10, 0xbc, 0x58, 0x85, 0xc1, 0x40,
	0x96, 0x32, 0xbd, 0xcb, 0xf4, 0x28, 0xea, 0xa1, 0xd3, 0x8b, 0x97, 0xf1, 0xd6, 0x85, 0x18, 0x49,
	0x93, 0x92, 0x66, 0x81, 0xfd, 0xd5, 0xfe, 0x0b, 0xd2, 0x64, 0xd6, 0xaa, 0xb3, 0xde, 0xbe, 0x7d,
	0xef, 0xf3, 0x3e, 0x25, 0xef, 0xa1, 0xf8, 0x15, 0x8c, 0x13, 0x65, 0x92, 0x69, 0x65, 0x0d, 0x64,
	0xb6, 0x4c, 0xdc, 0x24, 0x29, 0x21, 0x2f, 0x24, 0xa3, 0x85, 0xd1, 0x56, 0xe3, 0xc3, 0x40, 0xd0,
	0x9a, 0xa0, 0x6e, 0x72, 0x3c, 0xfc, 0x7d, 0x6c, 0x6e, 0xd7, 0xc5, 0x66, 0xb6, 0x15, 0x5b, 0x29,
	0x61, 0xdb, 0xb0, 0x9c, 0x59, 0x23, 0xb2, 0xa6, 0xed, 0x7c, 0x1b, 0x06, 0x9c, 0x1b, 0xc6, 0xc1,
	0x0a, 0xad, 0xe6, 0x92, 0x39, 0x26, 0x03, 0x7c, 0xf6, 0xd6, 0x41, 0xfd, 0x99, 0xff, 0x13, 0xc6,
	0xa8, 0xab, 0x20, 0x67, 0x24, 0x8a, 0xa3, 0xd1, 0x4e, 0xea, 0x33, 0x3e, 0x45, 0xbb, 0x4b, 0x51,
	0x16, 0x12, 0xd6, 0x73, 0xdf, 0xeb, 0xf8, 0xde, 0x60, 0x53, 0x7b, 0xa8, 0x90, 0x14, 0x1d, 0xfc,
	0x90, 0x93, 0x7f, 0x71, 0x34, 0xda, 0xbb, 0x18, 0xd2, 0x2d, 0x4b, 0xa1, 0xd7, 0x9f, 0xf4, 0x5d,
	0x05, 0xa7, 0xfb, 0xf0, 0xad, 0x82, 0xa7, 0x68, 0xd0, 0x78, 0x17, 0xe9, 0x7a, 0xdb, 0xc9, 0x56,
	0xdb, 0xbd, 0xe7, 0x1e, 0xd7, 0x05, 0x4b, 0x51, 0x5e, 0x67, 0x7c, 0x84, 0x7a, 0x0e, 0xe4, 0x8a,
	0x91, 0x5e, 0x1c, 0x8d, 0xa2, 0x34, 0x7c, 0x54, 0xde, 0xc6, 0xf6, 0x49, 0xbf, 0xc5, 0x1b, 0x96,
	0x12, 0xbc, 0x65, 0x9d, 0x1b, 0x86, 0xea, 0x30, 0xe4, 0xff, 0x9f, 0x86, 0x27, 0x25, 0xec, 0x87,
	0xa1, 0xca, 0x37, 0xd3, 0xe7, 0x2b, 0x2e, 0xec, 0xcb, 0x6a, 0x41, 0x33, 0x9d, 0x27, 0xb7, 0xda,
	0x18, 0x61, 0x67, 0x20, 0x1d, 0x33, 0x4b, 0x48, 0x82, 0x67, 0x5c, 0x7b, 0xc6, 0x5c, 0x4b, 0x50,
	0xfc, 0xcb, 0x2d, 0x17, 0x7d, 0x7f, 0xba, 0xcb, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x61, 0x66,
	0xb9, 0x80, 0x95, 0x02, 0x00, 0x00,
}
