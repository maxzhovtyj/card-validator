// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/card.proto

package proto

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

type Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number          string `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	ExpirationYear  int64  `protobuf:"varint,2,opt,name=expirationYear,proto3" json:"expirationYear,omitempty"`
	ExpirationMonth string `protobuf:"bytes,3,opt,name=expirationMonth,proto3" json:"expirationMonth,omitempty"`
}

func (x *Card) Reset() {
	*x = Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Card) ProtoMessage() {}

func (x *Card) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Card.ProtoReflect.Descriptor instead.
func (*Card) Descriptor() ([]byte, []int) {
	return file_proto_card_proto_rawDescGZIP(), []int{0}
}

func (x *Card) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Card) GetExpirationYear() int64 {
	if x != nil {
		return x.ExpirationYear
	}
	return 0
}

func (x *Card) GetExpirationMonth() string {
	if x != nil {
		return x.ExpirationMonth
	}
	return ""
}

type ValidateCardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Card *Card `protobuf:"bytes,1,opt,name=card,proto3" json:"card,omitempty"`
}

func (x *ValidateCardRequest) Reset() {
	*x = ValidateCardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCardRequest) ProtoMessage() {}

func (x *ValidateCardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCardRequest.ProtoReflect.Descriptor instead.
func (*ValidateCardRequest) Descriptor() ([]byte, []int) {
	return file_proto_card_proto_rawDescGZIP(), []int{1}
}

func (x *ValidateCardRequest) GetCard() *Card {
	if x != nil {
		return x.Card
	}
	return nil
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_proto_card_proto_rawDescGZIP(), []int{2}
}

func (x *Error) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ValidateCardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool   `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Error *Error `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *ValidateCardResponse) Reset() {
	*x = ValidateCardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_card_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateCardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateCardResponse) ProtoMessage() {}

func (x *ValidateCardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_card_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateCardResponse.ProtoReflect.Descriptor instead.
func (*ValidateCardResponse) Descriptor() ([]byte, []int) {
	return file_proto_card_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateCardResponse) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *ValidateCardResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_proto_card_proto protoreflect.FileDescriptor

var file_proto_card_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x70, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x59, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x59, 0x65, 0x61, 0x72, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x22, 0x30, 0x0a, 0x13, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x04, 0x63,
	0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x64,
	0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x22, 0x35, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x59, 0x0a,
	0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08,
	0x0a, 0x06, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x32, 0x46, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_card_proto_rawDescOnce sync.Once
	file_proto_card_proto_rawDescData = file_proto_card_proto_rawDesc
)

func file_proto_card_proto_rawDescGZIP() []byte {
	file_proto_card_proto_rawDescOnce.Do(func() {
		file_proto_card_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_card_proto_rawDescData)
	})
	return file_proto_card_proto_rawDescData
}

var file_proto_card_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_card_proto_goTypes = []interface{}{
	(*Card)(nil),                 // 0: Card
	(*ValidateCardRequest)(nil),  // 1: ValidateCardRequest
	(*Error)(nil),                // 2: Error
	(*ValidateCardResponse)(nil), // 3: ValidateCardResponse
}
var file_proto_card_proto_depIdxs = []int32{
	0, // 0: ValidateCardRequest.card:type_name -> Card
	2, // 1: ValidateCardResponse.error:type_name -> Error
	1, // 2: CardService.Validate:input_type -> ValidateCardRequest
	3, // 3: CardService.Validate:output_type -> ValidateCardResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_card_proto_init() }
func file_proto_card_proto_init() {
	if File_proto_card_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_card_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Card); i {
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
		file_proto_card_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCardRequest); i {
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
		file_proto_card_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_proto_card_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateCardResponse); i {
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
	file_proto_card_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_card_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_card_proto_goTypes,
		DependencyIndexes: file_proto_card_proto_depIdxs,
		MessageInfos:      file_proto_card_proto_msgTypes,
	}.Build()
	File_proto_card_proto = out.File
	file_proto_card_proto_rawDesc = nil
	file_proto_card_proto_goTypes = nil
	file_proto_card_proto_depIdxs = nil
}
