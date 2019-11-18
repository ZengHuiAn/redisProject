// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/Chat.proto

package proto

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

// 创建聊天室的请求
type CreateChatRequest struct {
	ChatMAX              int32    `protobuf:"varint,1,opt,name=ChatMAX,proto3" json:"ChatMAX,omitempty"`
	ChatName             string   `protobuf:"bytes,2,opt,name=ChatName,proto3" json:"ChatName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatRequest) Reset()         { *m = CreateChatRequest{} }
func (m *CreateChatRequest) String() string { return proto.CompactTextString(m) }
func (*CreateChatRequest) ProtoMessage()    {}
func (*CreateChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bff7ac5e1649b6, []int{0}
}

func (m *CreateChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatRequest.Unmarshal(m, b)
}
func (m *CreateChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatRequest.Marshal(b, m, deterministic)
}
func (m *CreateChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatRequest.Merge(m, src)
}
func (m *CreateChatRequest) XXX_Size() int {
	return xxx_messageInfo_CreateChatRequest.Size(m)
}
func (m *CreateChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatRequest proto.InternalMessageInfo

func (m *CreateChatRequest) GetChatMAX() int32 {
	if m != nil {
		return m.ChatMAX
	}
	return 0
}

func (m *CreateChatRequest) GetChatName() string {
	if m != nil {
		return m.ChatName
	}
	return ""
}

// 创建聊天室的返回
type CreateChatResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ChatID               int32    `protobuf:"varint,2,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateChatResponse) Reset()         { *m = CreateChatResponse{} }
func (m *CreateChatResponse) String() string { return proto.CompactTextString(m) }
func (*CreateChatResponse) ProtoMessage()    {}
func (*CreateChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bff7ac5e1649b6, []int{1}
}

func (m *CreateChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateChatResponse.Unmarshal(m, b)
}
func (m *CreateChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateChatResponse.Marshal(b, m, deterministic)
}
func (m *CreateChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateChatResponse.Merge(m, src)
}
func (m *CreateChatResponse) XXX_Size() int {
	return xxx_messageInfo_CreateChatResponse.Size(m)
}
func (m *CreateChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateChatResponse proto.InternalMessageInfo

func (m *CreateChatResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *CreateChatResponse) GetChatID() int32 {
	if m != nil {
		return m.ChatID
	}
	return 0
}

//删除聊天室请求
type DropChatRequest struct {
	ChatID               int32    `protobuf:"varint,1,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DropChatRequest) Reset()         { *m = DropChatRequest{} }
func (m *DropChatRequest) String() string { return proto.CompactTextString(m) }
func (*DropChatRequest) ProtoMessage()    {}
func (*DropChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bff7ac5e1649b6, []int{2}
}

func (m *DropChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropChatRequest.Unmarshal(m, b)
}
func (m *DropChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropChatRequest.Marshal(b, m, deterministic)
}
func (m *DropChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropChatRequest.Merge(m, src)
}
func (m *DropChatRequest) XXX_Size() int {
	return xxx_messageInfo_DropChatRequest.Size(m)
}
func (m *DropChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DropChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DropChatRequest proto.InternalMessageInfo

func (m *DropChatRequest) GetChatID() int32 {
	if m != nil {
		return m.ChatID
	}
	return 0
}

// 删除聊天室返回
type DropChatResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DropChatResponse) Reset()         { *m = DropChatResponse{} }
func (m *DropChatResponse) String() string { return proto.CompactTextString(m) }
func (*DropChatResponse) ProtoMessage()    {}
func (*DropChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bff7ac5e1649b6, []int{3}
}

func (m *DropChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropChatResponse.Unmarshal(m, b)
}
func (m *DropChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropChatResponse.Marshal(b, m, deterministic)
}
func (m *DropChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropChatResponse.Merge(m, src)
}
func (m *DropChatResponse) XXX_Size() int {
	return xxx_messageInfo_DropChatResponse.Size(m)
}
func (m *DropChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DropChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DropChatResponse proto.InternalMessageInfo

func (m *DropChatResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

// 主动离开聊天室的请求
type LeaveChatRequest struct {
	ChatID               int32    `protobuf:"varint,1,opt,name=ChatID,proto3" json:"ChatID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveChatRequest) Reset()         { *m = LeaveChatRequest{} }
func (m *LeaveChatRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveChatRequest) ProtoMessage()    {}
func (*LeaveChatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bff7ac5e1649b6, []int{4}
}

func (m *LeaveChatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveChatRequest.Unmarshal(m, b)
}
func (m *LeaveChatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveChatRequest.Marshal(b, m, deterministic)
}
func (m *LeaveChatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveChatRequest.Merge(m, src)
}
func (m *LeaveChatRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveChatRequest.Size(m)
}
func (m *LeaveChatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveChatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveChatRequest proto.InternalMessageInfo

func (m *LeaveChatRequest) GetChatID() int32 {
	if m != nil {
		return m.ChatID
	}
	return 0
}

// 主动离开聊天室的返回
type LeaveChatResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveChatResponse) Reset()         { *m = LeaveChatResponse{} }
func (m *LeaveChatResponse) String() string { return proto.CompactTextString(m) }
func (*LeaveChatResponse) ProtoMessage()    {}
func (*LeaveChatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b7bff7ac5e1649b6, []int{5}
}

func (m *LeaveChatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveChatResponse.Unmarshal(m, b)
}
func (m *LeaveChatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveChatResponse.Marshal(b, m, deterministic)
}
func (m *LeaveChatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveChatResponse.Merge(m, src)
}
func (m *LeaveChatResponse) XXX_Size() int {
	return xxx_messageInfo_LeaveChatResponse.Size(m)
}
func (m *LeaveChatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveChatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveChatResponse proto.InternalMessageInfo

func (m *LeaveChatResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateChatRequest)(nil), "proto.CreateChatRequest")
	proto.RegisterType((*CreateChatResponse)(nil), "proto.CreateChatResponse")
	proto.RegisterType((*DropChatRequest)(nil), "proto.DropChatRequest")
	proto.RegisterType((*DropChatResponse)(nil), "proto.DropChatResponse")
	proto.RegisterType((*LeaveChatRequest)(nil), "proto.LeaveChatRequest")
	proto.RegisterType((*LeaveChatResponse)(nil), "proto.LeaveChatResponse")
}

func init() { proto.RegisterFile("proto/Chat.proto", fileDescriptor_b7bff7ac5e1649b6) }

var fileDescriptor_b7bff7ac5e1649b6 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x77, 0xce, 0x48, 0x2c, 0xd1, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x27, 0x97,
	0xa0, 0x73, 0x51, 0x6a, 0x62, 0x49, 0x2a, 0x48, 0x2a, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44,
	0x48, 0x82, 0x8b, 0x1d, 0xc4, 0xf5, 0x75, 0x8c, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82,
	0x71, 0x85, 0xa4, 0xb8, 0x38, 0x40, 0x4c, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x38, 0x5f, 0xc9, 0x8b, 0x4b, 0x08, 0xd9, 0xa8, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x19, 0x2e, 0xce, 0xd4, 0xa2, 0xa2, 0xfc, 0x22, 0xe7, 0xfc, 0x94, 0x54, 0xa8, 0x69, 0x08,
	0x01, 0x21, 0x31, 0x2e, 0x36, 0x90, 0x6a, 0x4f, 0x17, 0xb0, 0x69, 0xac, 0x41, 0x50, 0x9e, 0x92,
	0x26, 0x17, 0xbf, 0x4b, 0x51, 0x7e, 0x01, 0xb2, 0xa3, 0x10, 0x4a, 0x19, 0x51, 0x94, 0x1a, 0x70,
	0x09, 0x20, 0x94, 0x12, 0x63, 0xa9, 0x92, 0x16, 0x97, 0x80, 0x4f, 0x6a, 0x62, 0x59, 0x2a, 0x31,
	0xa6, 0x1b, 0x72, 0x09, 0x22, 0xa9, 0x25, 0xc6, 0x78, 0xa3, 0x09, 0x8c, 0x5c, 0xdc, 0x20, 0xe5,
	0xc1, 0xa9, 0x45, 0x65, 0x99, 0xc9, 0xa9, 0x42, 0xce, 0x5c, 0x5c, 0x88, 0x70, 0x11, 0x92, 0x80,
	0x84, 0xbf, 0x1e, 0x46, 0xa8, 0x4b, 0x49, 0x62, 0x91, 0x81, 0x58, 0xa8, 0xc4, 0x20, 0x64, 0xc3,
	0xc5, 0x01, 0xf3, 0xa5, 0x90, 0x18, 0x54, 0x21, 0x5a, 0x08, 0x49, 0xe1, 0x10, 0x57, 0x62, 0x48,
	0x62, 0x03, 0x4b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb8, 0xf6, 0xc9, 0x5d, 0x07, 0x02,
	0x00, 0x00,
}