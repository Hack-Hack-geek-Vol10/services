// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: v1/token-service.proto

package v1

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

type CreateTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Authority string `protobuf:"bytes,2,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (x *CreateTokenRequest) Reset() {
	*x = CreateTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_token_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenRequest) ProtoMessage() {}

func (x *CreateTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_token_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_token_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTokenRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CreateTokenRequest) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

type CreateTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateTokenResponse) Reset() {
	*x = CreateTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_token_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTokenResponse) ProtoMessage() {}

func (x *CreateTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_token_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_token_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_token_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_token_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_token_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId   string `protobuf:"bytes,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Authority string `protobuf:"bytes,3,opt,name=authority,proto3" json:"authority,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_token_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_token_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_token_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTokenResponse) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *GetTokenResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GetTokenResponse) GetAuthority() string {
	if x != nil {
		return x.Authority
	}
	return ""
}

type DeleteTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeleteTokenRequest) Reset() {
	*x = DeleteTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_token_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenRequest) ProtoMessage() {}

func (x *DeleteTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_token_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) {
	return file_v1_token_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteTokenRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type DeleteTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *DeleteTokenResponse) Reset() {
	*x = DeleteTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_token_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTokenResponse) ProtoMessage() {}

func (x *DeleteTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_token_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTokenResponse.ProtoReflect.Descriptor instead.
func (*DeleteTokenResponse) Descriptor() ([]byte, []int) {
	return file_v1_token_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTokenResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

var File_v1_token_service_proto protoreflect.FileDescriptor

var file_v1_token_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x76, 0x31, 0x22, 0x51, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x2b, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x6a, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x22, 0x33, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x34, 0x0a, 0x13,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x32, 0xe9, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x2e, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16,
	0x5a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_token_service_proto_rawDescOnce sync.Once
	file_v1_token_service_proto_rawDescData = file_v1_token_service_proto_rawDesc
)

func file_v1_token_service_proto_rawDescGZIP() []byte {
	file_v1_token_service_proto_rawDescOnce.Do(func() {
		file_v1_token_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_token_service_proto_rawDescData)
	})
	return file_v1_token_service_proto_rawDescData
}

var file_v1_token_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_token_service_proto_goTypes = []interface{}{
	(*CreateTokenRequest)(nil),  // 0: token.v1.CreateTokenRequest
	(*CreateTokenResponse)(nil), // 1: token.v1.CreateTokenResponse
	(*GetTokenRequest)(nil),     // 2: token.v1.GetTokenRequest
	(*GetTokenResponse)(nil),    // 3: token.v1.GetTokenResponse
	(*DeleteTokenRequest)(nil),  // 4: token.v1.DeleteTokenRequest
	(*DeleteTokenResponse)(nil), // 5: token.v1.DeleteTokenResponse
}
var file_v1_token_service_proto_depIdxs = []int32{
	0, // 0: token.v1.TokenService.CreateToken:input_type -> token.v1.CreateTokenRequest
	2, // 1: token.v1.TokenService.GetToken:input_type -> token.v1.GetTokenRequest
	4, // 2: token.v1.TokenService.DeleteToken:input_type -> token.v1.DeleteTokenRequest
	1, // 3: token.v1.TokenService.CreateToken:output_type -> token.v1.CreateTokenResponse
	3, // 4: token.v1.TokenService.GetToken:output_type -> token.v1.GetTokenResponse
	5, // 5: token.v1.TokenService.DeleteToken:output_type -> token.v1.DeleteTokenResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_token_service_proto_init() }
func file_v1_token_service_proto_init() {
	if File_v1_token_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_token_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenRequest); i {
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
		file_v1_token_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTokenResponse); i {
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
		file_v1_token_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_v1_token_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenResponse); i {
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
		file_v1_token_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTokenRequest); i {
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
		file_v1_token_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTokenResponse); i {
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
			RawDescriptor: file_v1_token_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_token_service_proto_goTypes,
		DependencyIndexes: file_v1_token_service_proto_depIdxs,
		MessageInfos:      file_v1_token_service_proto_msgTypes,
	}.Build()
	File_v1_token_service_proto = out.File
	file_v1_token_service_proto_rawDesc = nil
	file_v1_token_service_proto_goTypes = nil
	file_v1_token_service_proto_depIdxs = nil
}
