// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: todo/todo.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message for a to do item
type ToDo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *ToDo) Reset() {
	*x = ToDo{}
	mi := &file_todo_todo_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToDo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToDo) ProtoMessage() {}

func (x *ToDo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToDo.ProtoReflect.Descriptor instead.
func (*ToDo) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{0}
}

func (x *ToDo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ToDo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ToDo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ToDo) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

// Message for creating to do items
type CreateToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *ToDo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *CreateToDoRequest) Reset() {
	*x = CreateToDoRequest{}
	mi := &file_todo_todo_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateToDoRequest) ProtoMessage() {}

func (x *CreateToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateToDoRequest.ProtoReflect.Descriptor instead.
func (*CreateToDoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateToDoRequest) GetTodo() *ToDo {
	if x != nil {
		return x.Todo
	}
	return nil
}

// Message for updating to do items
type UpdateToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Done        bool   `protobuf:"varint,4,opt,name=done,proto3" json:"done,omitempty"`
}

func (x *UpdateToDoRequest) Reset() {
	*x = UpdateToDoRequest{}
	mi := &file_todo_todo_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateToDoRequest) ProtoMessage() {}

func (x *UpdateToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateToDoRequest.ProtoReflect.Descriptor instead.
func (*UpdateToDoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateToDoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateToDoRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateToDoRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateToDoRequest) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

// Message to fetch a to do item
type GetToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetToDoRequest) Reset() {
	*x = GetToDoRequest{}
	mi := &file_todo_todo_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetToDoRequest) ProtoMessage() {}

func (x *GetToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetToDoRequest.ProtoReflect.Descriptor instead.
func (*GetToDoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{3}
}

func (x *GetToDoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Message to fetch all to do items
type ListToDoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListToDoRequest) Reset() {
	*x = ListToDoRequest{}
	mi := &file_todo_todo_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListToDoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListToDoRequest) ProtoMessage() {}

func (x *ListToDoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListToDoRequest.ProtoReflect.Descriptor instead.
func (*ListToDoRequest) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{4}
}

// Message for a list of to do items
type ListToDoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todos []*ToDo `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
}

func (x *ListToDoResponse) Reset() {
	*x = ListToDoResponse{}
	mi := &file_todo_todo_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListToDoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListToDoResponse) ProtoMessage() {}

func (x *ListToDoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_todo_todo_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListToDoResponse.ProtoReflect.Descriptor instead.
func (*ListToDoResponse) Descriptor() ([]byte, []int) {
	return file_todo_todo_proto_rawDescGZIP(), []int{5}
}

func (x *ListToDoResponse) GetTodos() []*ToDo {
	if x != nil {
		return x.Todos
	}
	return nil
}

var File_todo_todo_proto protoreflect.FileDescriptor

var file_todo_todo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x62, 0x0a, 0x04, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x33, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x04, 0x74,
	0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f,
	0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x6f, 0x0a, 0x11, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65, 0x22, 0x20, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x11,
	0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x34, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f,
	0x52, 0x05, 0x74, 0x6f, 0x64, 0x6f, 0x73, 0x32, 0x87, 0x03, 0x0a, 0x0b, 0x54, 0x6f, 0x44, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0d, 0x3a, 0x01, 0x2a, 0x22, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12,
	0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x17, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f,
	0x44, 0x6f, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x1a, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x42, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x54, 0x6f, 0x44, 0x6f, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x4c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x51,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x44, 0x6f, 0x12, 0x14, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x44, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x0c, 0x5a, 0x0a, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_todo_proto_rawDescOnce sync.Once
	file_todo_todo_proto_rawDescData = file_todo_todo_proto_rawDesc
)

func file_todo_todo_proto_rawDescGZIP() []byte {
	file_todo_todo_proto_rawDescOnce.Do(func() {
		file_todo_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_todo_proto_rawDescData)
	})
	return file_todo_todo_proto_rawDescData
}

var file_todo_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_todo_todo_proto_goTypes = []any{
	(*ToDo)(nil),              // 0: todo.ToDo
	(*CreateToDoRequest)(nil), // 1: todo.CreateToDoRequest
	(*UpdateToDoRequest)(nil), // 2: todo.UpdateToDoRequest
	(*GetToDoRequest)(nil),    // 3: todo.GetToDoRequest
	(*ListToDoRequest)(nil),   // 4: todo.ListToDoRequest
	(*ListToDoResponse)(nil),  // 5: todo.ListToDoResponse
	(*emptypb.Empty)(nil),     // 6: google.protobuf.Empty
}
var file_todo_todo_proto_depIdxs = []int32{
	0, // 0: todo.CreateToDoRequest.todo:type_name -> todo.ToDo
	0, // 1: todo.ListToDoResponse.todos:type_name -> todo.ToDo
	1, // 2: todo.ToDoService.CreateToDo:input_type -> todo.CreateToDoRequest
	2, // 3: todo.ToDoService.UpdateToDo:input_type -> todo.UpdateToDoRequest
	3, // 4: todo.ToDoService.GetToDo:input_type -> todo.GetToDoRequest
	6, // 5: todo.ToDoService.ListToDo:input_type -> google.protobuf.Empty
	3, // 6: todo.ToDoService.DeleteToDo:input_type -> todo.GetToDoRequest
	0, // 7: todo.ToDoService.CreateToDo:output_type -> todo.ToDo
	0, // 8: todo.ToDoService.UpdateToDo:output_type -> todo.ToDo
	0, // 9: todo.ToDoService.GetToDo:output_type -> todo.ToDo
	5, // 10: todo.ToDoService.ListToDo:output_type -> todo.ListToDoResponse
	6, // 11: todo.ToDoService.DeleteToDo:output_type -> google.protobuf.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_todo_todo_proto_init() }
func file_todo_todo_proto_init() {
	if File_todo_todo_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_todo_todo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_todo_todo_proto_goTypes,
		DependencyIndexes: file_todo_todo_proto_depIdxs,
		MessageInfos:      file_todo_todo_proto_msgTypes,
	}.Build()
	File_todo_todo_proto = out.File
	file_todo_todo_proto_rawDesc = nil
	file_todo_todo_proto_goTypes = nil
	file_todo_todo_proto_depIdxs = nil
}