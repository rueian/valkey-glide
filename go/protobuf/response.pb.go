// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: response.proto

package protobuf

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

type RequestErrorType int32

const (
	RequestErrorType_Unspecified RequestErrorType = 0
	RequestErrorType_ExecAbort   RequestErrorType = 1
	RequestErrorType_Timeout     RequestErrorType = 2
	RequestErrorType_Disconnect  RequestErrorType = 3
)

// Enum value maps for RequestErrorType.
var (
	RequestErrorType_name = map[int32]string{
		0: "Unspecified",
		1: "ExecAbort",
		2: "Timeout",
		3: "Disconnect",
	}
	RequestErrorType_value = map[string]int32{
		"Unspecified": 0,
		"ExecAbort":   1,
		"Timeout":     2,
		"Disconnect":  3,
	}
)

func (x RequestErrorType) Enum() *RequestErrorType {
	p := new(RequestErrorType)
	*p = x
	return p
}

func (x RequestErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_response_proto_enumTypes[0].Descriptor()
}

func (RequestErrorType) Type() protoreflect.EnumType {
	return &file_response_proto_enumTypes[0]
}

func (x RequestErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestErrorType.Descriptor instead.
func (RequestErrorType) EnumDescriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{0}
}

type ConstantResponse int32

const (
	ConstantResponse_OK ConstantResponse = 0
)

// Enum value maps for ConstantResponse.
var (
	ConstantResponse_name = map[int32]string{
		0: "OK",
	}
	ConstantResponse_value = map[string]int32{
		"OK": 0,
	}
)

func (x ConstantResponse) Enum() *ConstantResponse {
	p := new(ConstantResponse)
	*p = x
	return p
}

func (x ConstantResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConstantResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_response_proto_enumTypes[1].Descriptor()
}

func (ConstantResponse) Type() protoreflect.EnumType {
	return &file_response_proto_enumTypes[1]
}

func (x ConstantResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConstantResponse.Descriptor instead.
func (ConstantResponse) EnumDescriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{1}
}

type RequestError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    RequestErrorType `protobuf:"varint,1,opt,name=type,proto3,enum=response.RequestErrorType" json:"type,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestError) Reset() {
	*x = RequestError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestError) ProtoMessage() {}

func (x *RequestError) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestError.ProtoReflect.Descriptor instead.
func (*RequestError) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{0}
}

func (x *RequestError) GetType() RequestErrorType {
	if x != nil {
		return x.Type
	}
	return RequestErrorType_Unspecified
}

func (x *RequestError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CallbackIdx uint32 `protobuf:"varint,1,opt,name=callback_idx,json=callbackIdx,proto3" json:"callback_idx,omitempty"`
	// Types that are assignable to Value:
	//
	//	*Response_RespPointer
	//	*Response_ConstantResponse
	//	*Response_RequestError
	//	*Response_ClosingError
	Value  isResponse_Value `protobuf_oneof:"value"`
	IsPush bool             `protobuf:"varint,6,opt,name=is_push,json=isPush,proto3" json:"is_push,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_response_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCallbackIdx() uint32 {
	if x != nil {
		return x.CallbackIdx
	}
	return 0
}

func (m *Response) GetValue() isResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Response) GetRespPointer() uint64 {
	if x, ok := x.GetValue().(*Response_RespPointer); ok {
		return x.RespPointer
	}
	return 0
}

func (x *Response) GetConstantResponse() ConstantResponse {
	if x, ok := x.GetValue().(*Response_ConstantResponse); ok {
		return x.ConstantResponse
	}
	return ConstantResponse_OK
}

func (x *Response) GetRequestError() *RequestError {
	if x, ok := x.GetValue().(*Response_RequestError); ok {
		return x.RequestError
	}
	return nil
}

func (x *Response) GetClosingError() string {
	if x, ok := x.GetValue().(*Response_ClosingError); ok {
		return x.ClosingError
	}
	return ""
}

func (x *Response) GetIsPush() bool {
	if x != nil {
		return x.IsPush
	}
	return false
}

type isResponse_Value interface {
	isResponse_Value()
}

type Response_RespPointer struct {
	RespPointer uint64 `protobuf:"varint,2,opt,name=resp_pointer,json=respPointer,proto3,oneof"`
}

type Response_ConstantResponse struct {
	ConstantResponse ConstantResponse `protobuf:"varint,3,opt,name=constant_response,json=constantResponse,proto3,enum=response.ConstantResponse,oneof"`
}

type Response_RequestError struct {
	RequestError *RequestError `protobuf:"bytes,4,opt,name=request_error,json=requestError,proto3,oneof"`
}

type Response_ClosingError struct {
	ClosingError string `protobuf:"bytes,5,opt,name=closing_error,json=closingError,proto3,oneof"`
}

func (*Response_RespPointer) isResponse_Value() {}

func (*Response_ConstantResponse) isResponse_Value() {}

func (*Response_RequestError) isResponse_Value() {}

func (*Response_ClosingError) isResponse_Value() {}

var File_response_proto protoreflect.FileDescriptor

var file_response_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x64,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63,
	0x6b, 0x49, 0x64, 0x78, 0x12, 0x23, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x70, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x11, 0x63, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x48, 0x00, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0d, 0x63, 0x6c, 0x6f, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x63, 0x6c,
	0x6f, 0x73, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73,
	0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50,
	0x75, 0x73, 0x68, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x4f, 0x0a, 0x10,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x78, 0x65, 0x63, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x10, 0x03, 0x2a, 0x1a, 0x0a,
	0x10, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_response_proto_rawDescOnce sync.Once
	file_response_proto_rawDescData = file_response_proto_rawDesc
)

func file_response_proto_rawDescGZIP() []byte {
	file_response_proto_rawDescOnce.Do(func() {
		file_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_response_proto_rawDescData)
	})
	return file_response_proto_rawDescData
}

var file_response_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_response_proto_goTypes = []interface{}{
	(RequestErrorType)(0), // 0: response.RequestErrorType
	(ConstantResponse)(0), // 1: response.ConstantResponse
	(*RequestError)(nil),  // 2: response.RequestError
	(*Response)(nil),      // 3: response.Response
}
var file_response_proto_depIdxs = []int32{
	0, // 0: response.RequestError.type:type_name -> response.RequestErrorType
	1, // 1: response.Response.constant_response:type_name -> response.ConstantResponse
	2, // 2: response.Response.request_error:type_name -> response.RequestError
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_response_proto_init() }
func file_response_proto_init() {
	if File_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestError); i {
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
		file_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
	file_response_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_RespPointer)(nil),
		(*Response_ConstantResponse)(nil),
		(*Response_RequestError)(nil),
		(*Response_ClosingError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_response_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_response_proto_goTypes,
		DependencyIndexes: file_response_proto_depIdxs,
		EnumInfos:         file_response_proto_enumTypes,
		MessageInfos:      file_response_proto_msgTypes,
	}.Build()
	File_response_proto = out.File
	file_response_proto_rawDesc = nil
	file_response_proto_goTypes = nil
	file_response_proto_depIdxs = nil
}
