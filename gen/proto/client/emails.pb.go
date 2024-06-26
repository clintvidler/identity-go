// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: proto/client/emails.proto

package __

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

type ResponseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ResponseRequest) Reset() {
	*x = ResponseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_emails_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseRequest) ProtoMessage() {}

func (x *ResponseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_emails_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseRequest.ProtoReflect.Descriptor instead.
func (*ResponseRequest) Descriptor() ([]byte, []int) {
	return file_proto_client_emails_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type EmailSendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To      string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Body    string `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Subject string `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
}

func (x *EmailSendRequest) Reset() {
	*x = EmailSendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_emails_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailSendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailSendRequest) ProtoMessage() {}

func (x *EmailSendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_emails_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailSendRequest.ProtoReflect.Descriptor instead.
func (*EmailSendRequest) Descriptor() ([]byte, []int) {
	return file_proto_client_emails_proto_rawDescGZIP(), []int{1}
}

func (x *EmailSendRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *EmailSendRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *EmailSendRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *EmailSendRequest) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

type EmailSendResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *EmailSendResponse) Reset() {
	*x = EmailSendResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_client_emails_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailSendResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailSendResponse) ProtoMessage() {}

func (x *EmailSendResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_client_emails_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailSendResponse.ProtoReflect.Descriptor instead.
func (*EmailSendResponse) Descriptor() ([]byte, []int) {
	return file_proto_client_emails_proto_rawDescGZIP(), []int{2}
}

func (x *EmailSendResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_client_emails_proto protoreflect.FileDescriptor

var file_proto_client_emails_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x22, 0x23, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x64, 0x0a, 0x10, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x2d, 0x0a, 0x11,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x7a, 0x0a, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x36, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x15, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x04,
	0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_client_emails_proto_rawDescOnce sync.Once
	file_proto_client_emails_proto_rawDescData = file_proto_client_emails_proto_rawDesc
)

func file_proto_client_emails_proto_rawDescGZIP() []byte {
	file_proto_client_emails_proto_rawDescOnce.Do(func() {
		file_proto_client_emails_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_client_emails_proto_rawDescData)
	})
	return file_proto_client_emails_proto_rawDescData
}

var file_proto_client_emails_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_client_emails_proto_goTypes = []interface{}{
	(*ResponseRequest)(nil),   // 0: grpc.ResponseRequest
	(*EmailSendRequest)(nil),  // 1: grpc.EmailSendRequest
	(*EmailSendResponse)(nil), // 2: grpc.EmailSendResponse
}
var file_proto_client_emails_proto_depIdxs = []int32{
	0, // 0: grpc.Email.Echo:input_type -> grpc.ResponseRequest
	1, // 1: grpc.Email.Send:input_type -> grpc.EmailSendRequest
	0, // 2: grpc.Email.Echo:output_type -> grpc.ResponseRequest
	2, // 3: grpc.Email.Send:output_type -> grpc.EmailSendResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_client_emails_proto_init() }
func file_proto_client_emails_proto_init() {
	if File_proto_client_emails_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_client_emails_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseRequest); i {
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
		file_proto_client_emails_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailSendRequest); i {
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
		file_proto_client_emails_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailSendResponse); i {
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
			RawDescriptor: file_proto_client_emails_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_client_emails_proto_goTypes,
		DependencyIndexes: file_proto_client_emails_proto_depIdxs,
		MessageInfos:      file_proto_client_emails_proto_msgTypes,
	}.Build()
	File_proto_client_emails_proto = out.File
	file_proto_client_emails_proto_rawDesc = nil
	file_proto_client_emails_proto_goTypes = nil
	file_proto_client_emails_proto_depIdxs = nil
}
