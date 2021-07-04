// Copyright 2021 TsumiNa
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//*
//All files should be ordered in the following manner:
//
//1. License header (if applicable)
//2. File overview
//3. Syntax
//4. Package
//5. Imports (sorted)
//6. File options
//7. Everything else
//
//see also: https://developers.google.com/protocol-buffers/docs/style

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.2
// source: protobuf/scheduler/scheduler.protobuf

package scheduler

import (
	proto "github.com/golang/protobuf/proto"
	task "github.com/yoshida-lab/avalon/protobuf/task"
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

type SchedulerForwardFlag int32

const (
	SchedulerForwardFlag_FORWARD   SchedulerForwardFlag = 0
	SchedulerForwardFlag_FINISH    SchedulerForwardFlag = 1
	SchedulerForwardFlag_TERMINATE SchedulerForwardFlag = 2
)

// Enum value maps for SchedulerForwardFlag.
var (
	SchedulerForwardFlag_name = map[int32]string{
		0: "FORWARD",
		1: "FINISH",
		2: "TERMINATE",
	}
	SchedulerForwardFlag_value = map[string]int32{
		"FORWARD":   0,
		"FINISH":    1,
		"TERMINATE": 2,
	}
)

func (x SchedulerForwardFlag) Enum() *SchedulerForwardFlag {
	p := new(SchedulerForwardFlag)
	*p = x
	return p
}

func (x SchedulerForwardFlag) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SchedulerForwardFlag) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_scheduler_scheduler_proto_enumTypes[0].Descriptor()
}

func (SchedulerForwardFlag) Type() protoreflect.EnumType {
	return &file_proto_scheduler_scheduler_proto_enumTypes[0]
}

func (x SchedulerForwardFlag) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SchedulerForwardFlag.Descriptor instead.
func (SchedulerForwardFlag) EnumDescriptor() ([]byte, []int) {
	return file_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

type SchedulerSubmission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JobDir    string            `protobuf:"bytes,1,opt,hello=job_dir,json=jobDir,proto3" json:"job_dir,omitempty"`
	Hostname  string            `protobuf:"bytes,2,opt,hello=hostname,proto3" json:"hostname,omitempty"`
	Overwrite map[string]string `protobuf:"bytes,3,rep,hello=overwrite,proto3" json:"overwrite,omitempty" protobuf_key:"bytes,1,opt,hello=key,proto3" protobuf_val:"bytes,2,opt,hello=value,proto3"`
}

func (x *SchedulerSubmission) Reset() {
	*x = SchedulerSubmission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerSubmission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerSubmission) ProtoMessage() {}

func (x *SchedulerSubmission) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerSubmission.ProtoReflect.Descriptor instead.
func (*SchedulerSubmission) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *SchedulerSubmission) GetJobDir() string {
	if x != nil {
		return x.JobDir
	}
	return ""
}

func (x *SchedulerSubmission) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *SchedulerSubmission) GetOverwrite() map[string]string {
	if x != nil {
		return x.Overwrite
	}
	return nil
}

type SchedulerForwardReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId      uint32                 `protobuf:"varint,1,opt,hello=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Flag        SchedulerForwardFlag   `protobuf:"varint,2,opt,hello=flag,proto3,enum=scheduler.SchedulerForwardFlag" json:"flag,omitempty"`
	Variables   map[string]string      `protobuf:"bytes,3,rep,hello=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,hello=key,proto3" protobuf_val:"bytes,2,opt,hello=value,proto3"`
	Submissions []*SchedulerSubmission `protobuf:"bytes,4,rep,hello=submissions,proto3" json:"submissions,omitempty"`
}

func (x *SchedulerForwardReq) Reset() {
	*x = SchedulerForwardReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerForwardReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerForwardReq) ProtoMessage() {}

func (x *SchedulerForwardReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerForwardReq.ProtoReflect.Descriptor instead.
func (*SchedulerForwardReq) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *SchedulerForwardReq) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *SchedulerForwardReq) GetFlag() SchedulerForwardFlag {
	if x != nil {
		return x.Flag
	}
	return SchedulerForwardFlag_FORWARD
}

func (x *SchedulerForwardReq) GetVariables() map[string]string {
	if x != nil {
		return x.Variables
	}
	return nil
}

func (x *SchedulerForwardReq) GetSubmissions() []*SchedulerSubmission {
	if x != nil {
		return x.Submissions
	}
	return nil
}

type SchedulerForwardRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId uint32               `protobuf:"varint,1,opt,hello=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Flag   SchedulerForwardFlag `protobuf:"varint,2,opt,hello=flag,proto3,enum=scheduler.SchedulerForwardFlag" json:"flag,omitempty"`
	JobIds []uint32             `protobuf:"varint,3,rep,packed,hello=job_ids,json=jobIds,proto3" json:"job_ids,omitempty"`
}

func (x *SchedulerForwardRes) Reset() {
	*x = SchedulerForwardRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerForwardRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerForwardRes) ProtoMessage() {}

func (x *SchedulerForwardRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerForwardRes.ProtoReflect.Descriptor instead.
func (*SchedulerForwardRes) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *SchedulerForwardRes) GetTaskId() uint32 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

func (x *SchedulerForwardRes) GetFlag() SchedulerForwardFlag {
	if x != nil {
		return x.Flag
	}
	return SchedulerForwardFlag_FORWARD
}

func (x *SchedulerForwardRes) GetJobIds() []uint32 {
	if x != nil {
		return x.JobIds
	}
	return nil
}

var File_proto_scheduler_scheduler_proto protoreflect.FileDescriptor

var file_proto_scheduler_scheduler_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x15, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6a,
	0x6f, 0x62, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6a, 0x6f,
	0x62, 0x44, 0x69, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x4b, 0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x1a, 0x3c, 0x0a,
	0x0e, 0x4f, 0x76, 0x65, 0x72, 0x77, 0x72, 0x69, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb0, 0x02, 0x0a, 0x13,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x12, 0x4b, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x12, 0x40,
	0x0a, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x1a, 0x3c, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x80,
	0x01, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x33, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x52, 0x04,
	0x66, 0x6c, 0x61, 0x67, 0x12, 0x1b, 0x0a, 0x07, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0d, 0x42, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6a, 0x6f, 0x62, 0x49, 0x64,
	0x73, 0x2a, 0x3e, 0x0a, 0x14, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x0b, 0x0a, 0x07, 0x46, 0x4f, 0x52,
	0x57, 0x41, 0x52, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x10,
	0x02, 0x32, 0x82, 0x01, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12,
	0x26, 0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x6f, 0x6f, 0x74, 0x1a, 0x0a, 0x2e, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x4d, 0x0a, 0x0b, 0x53, 0x74, 0x65, 0x70, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x12, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x73, 0x68, 0x69, 0x64, 0x61, 0x2d, 0x6c, 0x61, 0x62,
	0x2f, 0x61, 0x76, 0x61, 0x6c, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scheduler_scheduler_proto_rawDescOnce sync.Once
	file_proto_scheduler_scheduler_proto_rawDescData = file_proto_scheduler_scheduler_proto_rawDesc
)

func file_proto_scheduler_scheduler_proto_rawDescGZIP() []byte {
	file_proto_scheduler_scheduler_proto_rawDescOnce.Do(func() {
		file_proto_scheduler_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scheduler_scheduler_proto_rawDescData)
	})
	return file_proto_scheduler_scheduler_proto_rawDescData
}

var file_proto_scheduler_scheduler_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_scheduler_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_scheduler_scheduler_proto_goTypes = []interface{}{
	(SchedulerForwardFlag)(0),   // 0: scheduler.SchedulerForwardFlag
	(*SchedulerSubmission)(nil), // 1: scheduler.SchedulerSubmission
	(*SchedulerForwardReq)(nil), // 2: scheduler.SchedulerForwardReq
	(*SchedulerForwardRes)(nil), // 3: scheduler.SchedulerForwardRes
	nil,                         // 4: scheduler.SchedulerSubmission.OverwriteEntry
	nil,                         // 5: scheduler.SchedulerForwardReq.VariablesEntry
	(*task.TaskRoot)(nil),       // 6: task.TaskRoot
	(*task.Task)(nil),           // 7: task.Task
}
var file_proto_scheduler_scheduler_proto_depIdxs = []int32{
	4, // 0: scheduler.SchedulerSubmission.overwrite:type_name -> scheduler.SchedulerSubmission.OverwriteEntry
	0, // 1: scheduler.SchedulerForwardReq.flag:type_name -> scheduler.SchedulerForwardFlag
	5, // 2: scheduler.SchedulerForwardReq.variables:type_name -> scheduler.SchedulerForwardReq.VariablesEntry
	1, // 3: scheduler.SchedulerForwardReq.submissions:type_name -> scheduler.SchedulerSubmission
	0, // 4: scheduler.SchedulerForwardRes.flag:type_name -> scheduler.SchedulerForwardFlag
	6, // 5: scheduler.Scheduler.InitFrom:input_type -> task.TaskRoot
	2, // 6: scheduler.Scheduler.StepForward:input_type -> scheduler.SchedulerForwardReq
	7, // 7: scheduler.Scheduler.InitFrom:output_type -> task.Task
	3, // 8: scheduler.Scheduler.StepForward:output_type -> scheduler.SchedulerForwardRes
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_scheduler_scheduler_proto_init() }
func file_proto_scheduler_scheduler_proto_init() {
	if File_proto_scheduler_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scheduler_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerSubmission); i {
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
		file_proto_scheduler_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerForwardReq); i {
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
		file_proto_scheduler_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerForwardRes); i {
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
			RawDescriptor: file_proto_scheduler_scheduler_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scheduler_scheduler_proto_goTypes,
		DependencyIndexes: file_proto_scheduler_scheduler_proto_depIdxs,
		EnumInfos:         file_proto_scheduler_scheduler_proto_enumTypes,
		MessageInfos:      file_proto_scheduler_scheduler_proto_msgTypes,
	}.Build()
	File_proto_scheduler_scheduler_proto = out.File
	file_proto_scheduler_scheduler_proto_rawDesc = nil
	file_proto_scheduler_scheduler_proto_goTypes = nil
	file_proto_scheduler_scheduler_proto_depIdxs = nil
}