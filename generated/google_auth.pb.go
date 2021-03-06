// Borrowed from https://raw.githubusercontent.com/beemdevelopment/Aegis/39f18b79cb4d5b09f9c98692bca88532f93edb31/app/src/main/proto/google_auth.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: google_auth.proto

package generated

import (
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

type MigrationPayload_Algorithm int32

const (
	MigrationPayload_ALGORITHM_UNSPECIFIED MigrationPayload_Algorithm = 0
	MigrationPayload_ALGORITHM_SHA1        MigrationPayload_Algorithm = 1
	MigrationPayload_ALGORITHM_SHA256      MigrationPayload_Algorithm = 2
	MigrationPayload_ALGORITHM_SHA512      MigrationPayload_Algorithm = 3
	MigrationPayload_ALGORITHM_MD5         MigrationPayload_Algorithm = 4
)

// Enum value maps for MigrationPayload_Algorithm.
var (
	MigrationPayload_Algorithm_name = map[int32]string{
		0: "ALGORITHM_UNSPECIFIED",
		1: "ALGORITHM_SHA1",
		2: "ALGORITHM_SHA256",
		3: "ALGORITHM_SHA512",
		4: "ALGORITHM_MD5",
	}
	MigrationPayload_Algorithm_value = map[string]int32{
		"ALGORITHM_UNSPECIFIED": 0,
		"ALGORITHM_SHA1":        1,
		"ALGORITHM_SHA256":      2,
		"ALGORITHM_SHA512":      3,
		"ALGORITHM_MD5":         4,
	}
)

func (x MigrationPayload_Algorithm) Enum() *MigrationPayload_Algorithm {
	p := new(MigrationPayload_Algorithm)
	*p = x
	return p
}

func (x MigrationPayload_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationPayload_Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_google_auth_proto_enumTypes[0].Descriptor()
}

func (MigrationPayload_Algorithm) Type() protoreflect.EnumType {
	return &file_google_auth_proto_enumTypes[0]
}

func (x MigrationPayload_Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationPayload_Algorithm.Descriptor instead.
func (MigrationPayload_Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_google_auth_proto_rawDescGZIP(), []int{0, 0}
}

type MigrationPayload_DigitCount int32

const (
	MigrationPayload_DIGIT_COUNT_UNSPECIFIED MigrationPayload_DigitCount = 0
	MigrationPayload_DIGIT_COUNT_SIX         MigrationPayload_DigitCount = 1
	MigrationPayload_DIGIT_COUNT_EIGHT       MigrationPayload_DigitCount = 2
)

// Enum value maps for MigrationPayload_DigitCount.
var (
	MigrationPayload_DigitCount_name = map[int32]string{
		0: "DIGIT_COUNT_UNSPECIFIED",
		1: "DIGIT_COUNT_SIX",
		2: "DIGIT_COUNT_EIGHT",
	}
	MigrationPayload_DigitCount_value = map[string]int32{
		"DIGIT_COUNT_UNSPECIFIED": 0,
		"DIGIT_COUNT_SIX":         1,
		"DIGIT_COUNT_EIGHT":       2,
	}
)

func (x MigrationPayload_DigitCount) Enum() *MigrationPayload_DigitCount {
	p := new(MigrationPayload_DigitCount)
	*p = x
	return p
}

func (x MigrationPayload_DigitCount) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationPayload_DigitCount) Descriptor() protoreflect.EnumDescriptor {
	return file_google_auth_proto_enumTypes[1].Descriptor()
}

func (MigrationPayload_DigitCount) Type() protoreflect.EnumType {
	return &file_google_auth_proto_enumTypes[1]
}

func (x MigrationPayload_DigitCount) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationPayload_DigitCount.Descriptor instead.
func (MigrationPayload_DigitCount) EnumDescriptor() ([]byte, []int) {
	return file_google_auth_proto_rawDescGZIP(), []int{0, 1}
}

type MigrationPayload_OtpType int32

const (
	MigrationPayload_OTP_TYPE_UNSPECIFIED MigrationPayload_OtpType = 0
	MigrationPayload_OTP_TYPE_HOTP        MigrationPayload_OtpType = 1
	MigrationPayload_OTP_TYPE_TOTP        MigrationPayload_OtpType = 2
)

// Enum value maps for MigrationPayload_OtpType.
var (
	MigrationPayload_OtpType_name = map[int32]string{
		0: "OTP_TYPE_UNSPECIFIED",
		1: "OTP_TYPE_HOTP",
		2: "OTP_TYPE_TOTP",
	}
	MigrationPayload_OtpType_value = map[string]int32{
		"OTP_TYPE_UNSPECIFIED": 0,
		"OTP_TYPE_HOTP":        1,
		"OTP_TYPE_TOTP":        2,
	}
)

func (x MigrationPayload_OtpType) Enum() *MigrationPayload_OtpType {
	p := new(MigrationPayload_OtpType)
	*p = x
	return p
}

func (x MigrationPayload_OtpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MigrationPayload_OtpType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_auth_proto_enumTypes[2].Descriptor()
}

func (MigrationPayload_OtpType) Type() protoreflect.EnumType {
	return &file_google_auth_proto_enumTypes[2]
}

func (x MigrationPayload_OtpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MigrationPayload_OtpType.Descriptor instead.
func (MigrationPayload_OtpType) EnumDescriptor() ([]byte, []int) {
	return file_google_auth_proto_rawDescGZIP(), []int{0, 2}
}

type MigrationPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtpParameters []*MigrationPayload_OtpParameters `protobuf:"bytes,1,rep,name=otp_parameters,json=otpParameters,proto3" json:"otp_parameters,omitempty"`
	Version       int32                             `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	BatchSize     int32                             `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	BatchIndex    int32                             `protobuf:"varint,4,opt,name=batch_index,json=batchIndex,proto3" json:"batch_index,omitempty"`
	BatchId       int32                             `protobuf:"varint,5,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
}

func (x *MigrationPayload) Reset() {
	*x = MigrationPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationPayload) ProtoMessage() {}

func (x *MigrationPayload) ProtoReflect() protoreflect.Message {
	mi := &file_google_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationPayload.ProtoReflect.Descriptor instead.
func (*MigrationPayload) Descriptor() ([]byte, []int) {
	return file_google_auth_proto_rawDescGZIP(), []int{0}
}

func (x *MigrationPayload) GetOtpParameters() []*MigrationPayload_OtpParameters {
	if x != nil {
		return x.OtpParameters
	}
	return nil
}

func (x *MigrationPayload) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *MigrationPayload) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *MigrationPayload) GetBatchIndex() int32 {
	if x != nil {
		return x.BatchIndex
	}
	return 0
}

func (x *MigrationPayload) GetBatchId() int32 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

type MigrationPayload_OtpParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret    []byte                      `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Name      string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Issuer    string                      `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Algorithm MigrationPayload_Algorithm  `protobuf:"varint,4,opt,name=algorithm,proto3,enum=generated.MigrationPayload_Algorithm" json:"algorithm,omitempty"`
	Digits    MigrationPayload_DigitCount `protobuf:"varint,5,opt,name=digits,proto3,enum=generated.MigrationPayload_DigitCount" json:"digits,omitempty"`
	Type      MigrationPayload_OtpType    `protobuf:"varint,6,opt,name=type,proto3,enum=generated.MigrationPayload_OtpType" json:"type,omitempty"`
	Counter   int64                       `protobuf:"varint,7,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *MigrationPayload_OtpParameters) Reset() {
	*x = MigrationPayload_OtpParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MigrationPayload_OtpParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MigrationPayload_OtpParameters) ProtoMessage() {}

func (x *MigrationPayload_OtpParameters) ProtoReflect() protoreflect.Message {
	mi := &file_google_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MigrationPayload_OtpParameters.ProtoReflect.Descriptor instead.
func (*MigrationPayload_OtpParameters) Descriptor() ([]byte, []int) {
	return file_google_auth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *MigrationPayload_OtpParameters) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *MigrationPayload_OtpParameters) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MigrationPayload_OtpParameters) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *MigrationPayload_OtpParameters) GetAlgorithm() MigrationPayload_Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return MigrationPayload_ALGORITHM_UNSPECIFIED
}

func (x *MigrationPayload_OtpParameters) GetDigits() MigrationPayload_DigitCount {
	if x != nil {
		return x.Digits
	}
	return MigrationPayload_DIGIT_COUNT_UNSPECIFIED
}

func (x *MigrationPayload_OtpParameters) GetType() MigrationPayload_OtpType {
	if x != nil {
		return x.Type
	}
	return MigrationPayload_OTP_TYPE_UNSPECIFIED
}

func (x *MigrationPayload_OtpParameters) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

var File_google_auth_proto protoreflect.FileDescriptor

var file_google_auth_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x22, 0xa4,
	0x06, 0x0a, 0x10, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x50, 0x0a, 0x0e, 0x6f, 0x74, 0x70, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4f, 0x74, 0x70, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0d, 0x6f, 0x74, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x1a, 0xab, 0x02, 0x0a, 0x0d, 0x4f,
	0x74, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x65,
	0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x12, 0x43, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x3e, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x44, 0x69, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x64,
	0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x4f, 0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x79, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54,
	0x48, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x53, 0x48,
	0x41, 0x31, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48,
	0x4d, 0x5f, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x4c,
	0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x53, 0x48, 0x41, 0x35, 0x31, 0x32, 0x10, 0x03,
	0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54, 0x48, 0x4d, 0x5f, 0x4d, 0x44,
	0x35, 0x10, 0x04, 0x22, 0x55, 0x0a, 0x0a, 0x44, 0x69, 0x67, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x49, 0x47, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x44, 0x49, 0x47, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x53, 0x49,
	0x58, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x47, 0x49, 0x54, 0x5f, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x45, 0x49, 0x47, 0x48, 0x54, 0x10, 0x02, 0x22, 0x49, 0x0a, 0x07, 0x4f, 0x74,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x4f, 0x54, 0x50, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x11, 0x0a, 0x0d, 0x4f, 0x54, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x4f, 0x54, 0x50,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x54, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54,
	0x4f, 0x54, 0x50, 0x10, 0x02, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_auth_proto_rawDescOnce sync.Once
	file_google_auth_proto_rawDescData = file_google_auth_proto_rawDesc
)

func file_google_auth_proto_rawDescGZIP() []byte {
	file_google_auth_proto_rawDescOnce.Do(func() {
		file_google_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_auth_proto_rawDescData)
	})
	return file_google_auth_proto_rawDescData
}

var file_google_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_google_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_auth_proto_goTypes = []interface{}{
	(MigrationPayload_Algorithm)(0),        // 0: generated.MigrationPayload.Algorithm
	(MigrationPayload_DigitCount)(0),       // 1: generated.MigrationPayload.DigitCount
	(MigrationPayload_OtpType)(0),          // 2: generated.MigrationPayload.OtpType
	(*MigrationPayload)(nil),               // 3: generated.MigrationPayload
	(*MigrationPayload_OtpParameters)(nil), // 4: generated.MigrationPayload.OtpParameters
}
var file_google_auth_proto_depIdxs = []int32{
	4, // 0: generated.MigrationPayload.otp_parameters:type_name -> generated.MigrationPayload.OtpParameters
	0, // 1: generated.MigrationPayload.OtpParameters.algorithm:type_name -> generated.MigrationPayload.Algorithm
	1, // 2: generated.MigrationPayload.OtpParameters.digits:type_name -> generated.MigrationPayload.DigitCount
	2, // 3: generated.MigrationPayload.OtpParameters.type:type_name -> generated.MigrationPayload.OtpType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_auth_proto_init() }
func file_google_auth_proto_init() {
	if File_google_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationPayload); i {
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
		file_google_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MigrationPayload_OtpParameters); i {
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
			RawDescriptor: file_google_auth_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_auth_proto_goTypes,
		DependencyIndexes: file_google_auth_proto_depIdxs,
		EnumInfos:         file_google_auth_proto_enumTypes,
		MessageInfos:      file_google_auth_proto_msgTypes,
	}.Build()
	File_google_auth_proto = out.File
	file_google_auth_proto_rawDesc = nil
	file_google_auth_proto_goTypes = nil
	file_google_auth_proto_depIdxs = nil
}
