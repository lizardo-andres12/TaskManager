// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: proto/taskservice.proto

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

// defined to allow all task details to be returned with responses
type Task struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Title         string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status        uint32                 `protobuf:"varint,4,opt,name=status,proto3" json:"status,omitempty"`
	Deadline      string                 `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Priority      bool                   `protobuf:"varint,6,opt,name=priority,proto3" json:"priority,omitempty"`
	CreatorId     uint64                 `protobuf:"varint,7,opt,name=creatorId,proto3" json:"creatorId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Task) Reset() {
	*x = Task{}
	mi := &file_proto_taskservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{0}
}

func (x *Task) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *Task) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Task) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Task) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Task) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *Task) GetPriority() bool {
	if x != nil {
		return x.Priority
	}
	return false
}

func (x *Task) GetCreatorId() uint64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Task          *Task                  `protobuf:"bytes,1,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_proto_taskservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRequest) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type AssignRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	AssigneeId    uint64                 `protobuf:"varint,2,opt,name=assigneeId,proto3" json:"assigneeId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssignRequest) Reset() {
	*x = AssignRequest{}
	mi := &file_proto_taskservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssignRequest) ProtoMessage() {}

func (x *AssignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssignRequest.ProtoReflect.Descriptor instead.
func (*AssignRequest) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{2}
}

func (x *AssignRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *AssignRequest) GetAssigneeId() uint64 {
	if x != nil {
		return x.AssigneeId
	}
	return 0
}

type SuccessResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SuccessResponse) Reset() {
	*x = SuccessResponse{}
	mi := &file_proto_taskservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SuccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessResponse) ProtoMessage() {}

func (x *SuccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessResponse.ProtoReflect.Descriptor instead.
func (*SuccessResponse) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{3}
}

func (x *SuccessResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// ambiguous id so message can be used for multiple get types
type IDOnlyRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // bool priorityOnly = 2;
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IDOnlyRequest) Reset() {
	*x = IDOnlyRequest{}
	mi := &file_proto_taskservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IDOnlyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDOnlyRequest) ProtoMessage() {}

func (x *IDOnlyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDOnlyRequest.ProtoReflect.Descriptor instead.
func (*IDOnlyRequest) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{4}
}

func (x *IDOnlyRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Task          *Task                  `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_proto_taskservice_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetResponse) GetTask() *Task {
	if x != nil {
		return x.Task
	}
	return nil
}

type GetAllResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Tasks         []*Task                `protobuf:"bytes,2,rep,name=tasks,proto3" json:"tasks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	mi := &file_proto_taskservice_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetAllResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type UpdateStringRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStringRequest) Reset() {
	*x = UpdateStringRequest{}
	mi := &file_proto_taskservice_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStringRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStringRequest) ProtoMessage() {}

func (x *UpdateStringRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStringRequest.ProtoReflect.Descriptor instead.
func (*UpdateStringRequest) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateStringRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateStringRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type UpdateStatusRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Status        uint32                 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStatusRequest) Reset() {
	*x = UpdateStatusRequest{}
	mi := &file_proto_taskservice_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStatusRequest) ProtoMessage() {}

func (x *UpdateStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStatusRequest.ProtoReflect.Descriptor instead.
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateStatusRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdateStatusRequest) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

type UpdatePriorityRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TaskId        uint64                 `protobuf:"varint,1,opt,name=taskId,proto3" json:"taskId,omitempty"`
	Priority      bool                   `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdatePriorityRequest) Reset() {
	*x = UpdatePriorityRequest{}
	mi := &file_proto_taskservice_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdatePriorityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriorityRequest) ProtoMessage() {}

func (x *UpdatePriorityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_taskservice_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriorityRequest.ProtoReflect.Descriptor instead.
func (*UpdatePriorityRequest) Descriptor() ([]byte, []int) {
	return file_proto_taskservice_proto_rawDescGZIP(), []int{9}
}

func (x *UpdatePriorityRequest) GetTaskId() uint64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *UpdatePriorityRequest) GetPriority() bool {
	if x != nil {
		return x.Priority
	}
	return false
}

var File_proto_taskservice_proto protoreflect.FileDescriptor

var file_proto_taskservice_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x04, 0x54, 0x61,
	0x73, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x22, 0x2a, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x47, 0x0a, 0x0d,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69, 0x67,
	0x6e, 0x65, 0x65, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x0f, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x1f, 0x0a, 0x0d, 0x49, 0x44, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x04,
	0x74, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x47, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73,
	0x22, 0x41, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61,
	0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x15, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x70,
	0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x32, 0x85, 0x05, 0x0a, 0x0f, 0x54, 0x61, 0x73, 0x6b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x52, 0x50, 0x43, 0x12, 0x2e, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0c, 0x41,
	0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x6f, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x49, 0x44, 0x4f, 0x6e, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x49, 0x44, 0x4f, 0x6e, 0x6c,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x2e, 0x49, 0x44, 0x4f,
	0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x36, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6f, 0x72,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e, 0x49, 0x44,
	0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a,
	0x0c, 0x55, 0x6e, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x0e, 0x2e,
	0x49, 0x44, 0x4f, 0x6e, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x1e, 0x5a, 0x1c, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_taskservice_proto_rawDescOnce sync.Once
	file_proto_taskservice_proto_rawDescData = file_proto_taskservice_proto_rawDesc
)

func file_proto_taskservice_proto_rawDescGZIP() []byte {
	file_proto_taskservice_proto_rawDescOnce.Do(func() {
		file_proto_taskservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_taskservice_proto_rawDescData)
	})
	return file_proto_taskservice_proto_rawDescData
}

var file_proto_taskservice_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_taskservice_proto_goTypes = []any{
	(*Task)(nil),                  // 0: Task
	(*CreateRequest)(nil),         // 1: CreateRequest
	(*AssignRequest)(nil),         // 2: AssignRequest
	(*SuccessResponse)(nil),       // 3: SuccessResponse
	(*IDOnlyRequest)(nil),         // 4: IDOnlyRequest
	(*GetResponse)(nil),           // 5: GetResponse
	(*GetAllResponse)(nil),        // 6: GetAllResponse
	(*UpdateStringRequest)(nil),   // 7: UpdateStringRequest
	(*UpdateStatusRequest)(nil),   // 8: UpdateStatusRequest
	(*UpdatePriorityRequest)(nil), // 9: UpdatePriorityRequest
}
var file_proto_taskservice_proto_depIdxs = []int32{
	0,  // 0: CreateRequest.task:type_name -> Task
	0,  // 1: GetResponse.task:type_name -> Task
	0,  // 2: GetAllResponse.tasks:type_name -> Task
	1,  // 3: TaskServiceGRPC.CreateTask:input_type -> CreateRequest
	2,  // 4: TaskServiceGRPC.AssignToTask:input_type -> AssignRequest
	4,  // 5: TaskServiceGRPC.GetTask:input_type -> IDOnlyRequest
	4,  // 6: TaskServiceGRPC.GetAllAssigned:input_type -> IDOnlyRequest
	4,  // 7: TaskServiceGRPC.GetAllCreated:input_type -> IDOnlyRequest
	7,  // 8: TaskServiceGRPC.UpdateTitle:input_type -> UpdateStringRequest
	7,  // 9: TaskServiceGRPC.UpdateDescription:input_type -> UpdateStringRequest
	8,  // 10: TaskServiceGRPC.UpdateStatus:input_type -> UpdateStatusRequest
	7,  // 11: TaskServiceGRPC.UpdateDeadline:input_type -> UpdateStringRequest
	9,  // 12: TaskServiceGRPC.UpdatePriority:input_type -> UpdatePriorityRequest
	4,  // 13: TaskServiceGRPC.DeleteTask:input_type -> IDOnlyRequest
	4,  // 14: TaskServiceGRPC.UnassignTask:input_type -> IDOnlyRequest
	3,  // 15: TaskServiceGRPC.CreateTask:output_type -> SuccessResponse
	3,  // 16: TaskServiceGRPC.AssignToTask:output_type -> SuccessResponse
	5,  // 17: TaskServiceGRPC.GetTask:output_type -> GetResponse
	6,  // 18: TaskServiceGRPC.GetAllAssigned:output_type -> GetAllResponse
	6,  // 19: TaskServiceGRPC.GetAllCreated:output_type -> GetAllResponse
	3,  // 20: TaskServiceGRPC.UpdateTitle:output_type -> SuccessResponse
	3,  // 21: TaskServiceGRPC.UpdateDescription:output_type -> SuccessResponse
	3,  // 22: TaskServiceGRPC.UpdateStatus:output_type -> SuccessResponse
	3,  // 23: TaskServiceGRPC.UpdateDeadline:output_type -> SuccessResponse
	3,  // 24: TaskServiceGRPC.UpdatePriority:output_type -> SuccessResponse
	3,  // 25: TaskServiceGRPC.DeleteTask:output_type -> SuccessResponse
	3,  // 26: TaskServiceGRPC.UnassignTask:output_type -> SuccessResponse
	15, // [15:27] is the sub-list for method output_type
	3,  // [3:15] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_taskservice_proto_init() }
func file_proto_taskservice_proto_init() {
	if File_proto_taskservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_taskservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_taskservice_proto_goTypes,
		DependencyIndexes: file_proto_taskservice_proto_depIdxs,
		MessageInfos:      file_proto_taskservice_proto_msgTypes,
	}.Build()
	File_proto_taskservice_proto = out.File
	file_proto_taskservice_proto_rawDesc = nil
	file_proto_taskservice_proto_goTypes = nil
	file_proto_taskservice_proto_depIdxs = nil
}
