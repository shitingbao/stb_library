// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: api/slog/v1/slog.proto

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

// The request message containing the user's name.
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *HelloReply) Reset() {
	*x = HelloReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReply) ProtoMessage() {}

func (x *HelloReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReply.ProtoReflect.Descriptor instead.
func (*HelloReply) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{1}
}

func (x *HelloReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RequestLogMessages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys     string      `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`
	Msg     *LogMessage `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Level   int64       `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Version string      `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	LogTime string      `protobuf:"bytes,5,opt,name=log_time,json=logTime,proto3" json:"log_time,omitempty"`
}

func (x *RequestLogMessages) Reset() {
	*x = RequestLogMessages{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogMessages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogMessages) ProtoMessage() {}

func (x *RequestLogMessages) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogMessages.ProtoReflect.Descriptor instead.
func (*RequestLogMessages) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{2}
}

func (x *RequestLogMessages) GetSys() string {
	if x != nil {
		return x.Sys
	}
	return ""
}

func (x *RequestLogMessages) GetMsg() *LogMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *RequestLogMessages) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *RequestLogMessages) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RequestLogMessages) GetLogTime() string {
	if x != nil {
		return x.LogTime
	}
	return ""
}

type RequestLogIdenticalMessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sys     string        `protobuf:"bytes,1,opt,name=sys,proto3" json:"sys,omitempty"`
	Msg     []*LogMessage `protobuf:"bytes,2,rep,name=msg,proto3" json:"msg,omitempty"`
	Level   int64         `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Version string        `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	LogTime string        `protobuf:"bytes,5,opt,name=log_time,json=logTime,proto3" json:"log_time,omitempty"`
}

func (x *RequestLogIdenticalMessageList) Reset() {
	*x = RequestLogIdenticalMessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogIdenticalMessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogIdenticalMessageList) ProtoMessage() {}

func (x *RequestLogIdenticalMessageList) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogIdenticalMessageList.ProtoReflect.Descriptor instead.
func (*RequestLogIdenticalMessageList) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{3}
}

func (x *RequestLogIdenticalMessageList) GetSys() string {
	if x != nil {
		return x.Sys
	}
	return ""
}

func (x *RequestLogIdenticalMessageList) GetMsg() []*LogMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

func (x *RequestLogIdenticalMessageList) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *RequestLogIdenticalMessageList) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RequestLogIdenticalMessageList) GetLogTime() string {
	if x != nil {
		return x.LogTime
	}
	return ""
}

type RequestLogDifferentMessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg []*RequestLogMessages `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RequestLogDifferentMessageList) Reset() {
	*x = RequestLogDifferentMessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogDifferentMessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogDifferentMessageList) ProtoMessage() {}

func (x *RequestLogDifferentMessageList) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogDifferentMessageList.ProtoReflect.Descriptor instead.
func (*RequestLogDifferentMessageList) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{4}
}

func (x *RequestLogDifferentMessageList) GetMsg() []*RequestLogMessages {
	if x != nil {
		return x.Msg
	}
	return nil
}

type LogMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic   string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *LogMessage) Reset() {
	*x = LogMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogMessage) ProtoMessage() {}

func (x *LogMessage) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogMessage.ProtoReflect.Descriptor instead.
func (*LogMessage) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{5}
}

func (x *LogMessage) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *LogMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type RespondLogRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RespondLogRes) Reset() {
	*x = RespondLogRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondLogRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondLogRes) ProtoMessage() {}

func (x *RespondLogRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondLogRes.ProtoReflect.Descriptor instead.
func (*RespondLogRes) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{6}
}

func (x *RespondLogRes) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespondLogRes) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RequestLogFindParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogSys       string    `protobuf:"bytes,1,opt,name=log_sys,json=logSys,proto3" json:"log_sys,omitempty"`
	LogStartTime string    `protobuf:"bytes,2,opt,name=log_start_time,json=logStartTime,proto3" json:"log_start_time,omitempty"`
	LogEndTime   string    `protobuf:"bytes,3,opt,name=log_end_time,json=logEndTime,proto3" json:"log_end_time,omitempty"`
	LogLevel     string    `protobuf:"bytes,4,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"` //逗号隔开
	Topic        string    `protobuf:"bytes,5,opt,name=topic,proto3" json:"topic,omitempty"`
	Content      string    `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	Page         int64     `protobuf:"varint,7,opt,name=page,proto3" json:"page,omitempty"`
	PageSize     int64     `protobuf:"varint,8,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Order        *ArgOrder `protobuf:"bytes,9,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *RequestLogFindParam) Reset() {
	*x = RequestLogFindParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestLogFindParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestLogFindParam) ProtoMessage() {}

func (x *RequestLogFindParam) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestLogFindParam.ProtoReflect.Descriptor instead.
func (*RequestLogFindParam) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{7}
}

func (x *RequestLogFindParam) GetLogSys() string {
	if x != nil {
		return x.LogSys
	}
	return ""
}

func (x *RequestLogFindParam) GetLogStartTime() string {
	if x != nil {
		return x.LogStartTime
	}
	return ""
}

func (x *RequestLogFindParam) GetLogEndTime() string {
	if x != nil {
		return x.LogEndTime
	}
	return ""
}

func (x *RequestLogFindParam) GetLogLevel() string {
	if x != nil {
		return x.LogLevel
	}
	return ""
}

func (x *RequestLogFindParam) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *RequestLogFindParam) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *RequestLogFindParam) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *RequestLogFindParam) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *RequestLogFindParam) GetOrder() *ArgOrder {
	if x != nil {
		return x.Order
	}
	return nil
}

type ArgOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderField string `protobuf:"bytes,1,opt,name=order_field,json=orderField,proto3" json:"order_field,omitempty"`
	OrderVal   int64  `protobuf:"varint,2,opt,name=order_val,json=orderVal,proto3" json:"order_val,omitempty"`
}

func (x *ArgOrder) Reset() {
	*x = ArgOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArgOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgOrder) ProtoMessage() {}

func (x *ArgOrder) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgOrder.ProtoReflect.Descriptor instead.
func (*ArgOrder) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{8}
}

func (x *ArgOrder) GetOrderField() string {
	if x != nil {
		return x.OrderField
	}
	return ""
}

func (x *ArgOrder) GetOrderVal() int64 {
	if x != nil {
		return x.OrderVal
	}
	return 0
}

type RespondLogFindList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *RespondLogFindList) Reset() {
	*x = RespondLogFindList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondLogFindList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondLogFindList) ProtoMessage() {}

func (x *RespondLogFindList) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondLogFindList.ProtoReflect.Descriptor instead.
func (*RespondLogFindList) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{9}
}

func (x *RespondLogFindList) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type RespondLogFind struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	LogSys   string `protobuf:"bytes,2,opt,name=log_sys,json=logSys,proto3" json:"log_sys,omitempty"`
	LogTime  string `protobuf:"bytes,3,opt,name=log_time,json=logTime,proto3" json:"log_time,omitempty"`
	LogLevel int64  `protobuf:"varint,4,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
	Content  string `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	Host     string `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	Topic    string `protobuf:"bytes,7,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *RespondLogFind) Reset() {
	*x = RespondLogFind{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_slog_v1_slog_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespondLogFind) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespondLogFind) ProtoMessage() {}

func (x *RespondLogFind) ProtoReflect() protoreflect.Message {
	mi := &file_api_slog_v1_slog_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespondLogFind.ProtoReflect.Descriptor instead.
func (*RespondLogFind) Descriptor() ([]byte, []int) {
	return file_api_slog_v1_slog_proto_rawDescGZIP(), []int{10}
}

func (x *RespondLogFind) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RespondLogFind) GetLogSys() string {
	if x != nil {
		return x.LogSys
	}
	return ""
}

func (x *RespondLogFind) GetLogTime() string {
	if x != nil {
		return x.LogTime
	}
	return ""
}

func (x *RespondLogFind) GetLogLevel() int64 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

func (x *RespondLogFind) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *RespondLogFind) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RespondLogFind) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

var File_api_slog_v1_slog_proto protoreflect.FileDescriptor

var file_api_slog_v1_slog_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6c,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c,
	0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22,
	0x22, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x0a, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa6, 0x01, 0x0a, 0x12,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x79, 0x73, 0x12, 0x33, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0xb2, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x79, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x79, 0x73, 0x12, 0x33, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x19,
	0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x1e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f,
	0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3c, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xab, 0x02,
	0x0a, 0x13, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6e, 0x64,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x79, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67, 0x53, 0x79, 0x73, 0x12, 0x24,
	0x0a, 0x0e, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x35, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x67, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x48, 0x0a, 0x08, 0x41,
	0x72, 0x67, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x56, 0x61, 0x6c, 0x22, 0x26, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0xbf, 0x01,
	0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6e, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x6f,
	0x67, 0x5f, 0x73, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x67,
	0x53, 0x79, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x6c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x32,
	0xb6, 0x03, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x5d, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x4c, 0x6f, 0x67, 0x12, 0x29, 0x2e, 0x77, 0x6f,
	0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x24, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f,
	0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x73, 0x0a, 0x14,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x4c, 0x6f, 0x67, 0x12, 0x35, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x6f,
	0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x73, 0x12, 0x73, 0x0a, 0x14, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x6e, 0x79, 0x44, 0x69, 0x66,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x35, 0x2e, 0x77, 0x6f, 0x6e, 0x65,
	0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x66, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64,
	0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x12, 0x60, 0x0a, 0x07, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6e,
	0x64, 0x12, 0x2a, 0x2e, 0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x6f, 0x67, 0x46, 0x69, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x29, 0x2e,
	0x77, 0x6f, 0x6e, 0x65, 0x5f, 0x6c, 0x6f, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x4c, 0x6f, 0x67,
	0x46, 0x69, 0x6e, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x12, 0x50, 0x01, 0x5a, 0x0e, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_slog_v1_slog_proto_rawDescOnce sync.Once
	file_api_slog_v1_slog_proto_rawDescData = file_api_slog_v1_slog_proto_rawDesc
)

func file_api_slog_v1_slog_proto_rawDescGZIP() []byte {
	file_api_slog_v1_slog_proto_rawDescOnce.Do(func() {
		file_api_slog_v1_slog_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_slog_v1_slog_proto_rawDescData)
	})
	return file_api_slog_v1_slog_proto_rawDescData
}

var file_api_slog_v1_slog_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_slog_v1_slog_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),                   // 0: wone_logger.server.v1.HelloRequest
	(*HelloReply)(nil),                     // 1: wone_logger.server.v1.HelloReply
	(*RequestLogMessages)(nil),             // 2: wone_logger.server.v1.RequestLogMessages
	(*RequestLogIdenticalMessageList)(nil), // 3: wone_logger.server.v1.RequestLogIdenticalMessageList
	(*RequestLogDifferentMessageList)(nil), // 4: wone_logger.server.v1.RequestLogDifferentMessageList
	(*LogMessage)(nil),                     // 5: wone_logger.server.v1.LogMessage
	(*RespondLogRes)(nil),                  // 6: wone_logger.server.v1.RespondLogRes
	(*RequestLogFindParam)(nil),            // 7: wone_logger.server.v1.RequestLogFindParam
	(*ArgOrder)(nil),                       // 8: wone_logger.server.v1.ArgOrder
	(*RespondLogFindList)(nil),             // 9: wone_logger.server.v1.RespondLogFindList
	(*RespondLogFind)(nil),                 // 10: wone_logger.server.v1.RespondLogFind
}
var file_api_slog_v1_slog_proto_depIdxs = []int32{
	5, // 0: wone_logger.server.v1.RequestLogMessages.msg:type_name -> wone_logger.server.v1.LogMessage
	5, // 1: wone_logger.server.v1.RequestLogIdenticalMessageList.msg:type_name -> wone_logger.server.v1.LogMessage
	2, // 2: wone_logger.server.v1.RequestLogDifferentMessageList.msg:type_name -> wone_logger.server.v1.RequestLogMessages
	8, // 3: wone_logger.server.v1.RequestLogFindParam.order:type_name -> wone_logger.server.v1.ArgOrder
	2, // 4: wone_logger.server.v1.LogServer.SendOneLog:input_type -> wone_logger.server.v1.RequestLogMessages
	3, // 5: wone_logger.server.v1.LogServer.SendManyIdenticalLog:input_type -> wone_logger.server.v1.RequestLogIdenticalMessageList
	4, // 6: wone_logger.server.v1.LogServer.SendManyDifferentLog:input_type -> wone_logger.server.v1.RequestLogDifferentMessageList
	7, // 7: wone_logger.server.v1.LogServer.LogFind:input_type -> wone_logger.server.v1.RequestLogFindParam
	6, // 8: wone_logger.server.v1.LogServer.SendOneLog:output_type -> wone_logger.server.v1.RespondLogRes
	6, // 9: wone_logger.server.v1.LogServer.SendManyIdenticalLog:output_type -> wone_logger.server.v1.RespondLogRes
	6, // 10: wone_logger.server.v1.LogServer.SendManyDifferentLog:output_type -> wone_logger.server.v1.RespondLogRes
	9, // 11: wone_logger.server.v1.LogServer.LogFind:output_type -> wone_logger.server.v1.RespondLogFindList
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_slog_v1_slog_proto_init() }
func file_api_slog_v1_slog_proto_init() {
	if File_api_slog_v1_slog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_slog_v1_slog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_api_slog_v1_slog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReply); i {
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
		file_api_slog_v1_slog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogMessages); i {
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
		file_api_slog_v1_slog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogIdenticalMessageList); i {
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
		file_api_slog_v1_slog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogDifferentMessageList); i {
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
		file_api_slog_v1_slog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogMessage); i {
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
		file_api_slog_v1_slog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondLogRes); i {
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
		file_api_slog_v1_slog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestLogFindParam); i {
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
		file_api_slog_v1_slog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArgOrder); i {
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
		file_api_slog_v1_slog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondLogFindList); i {
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
		file_api_slog_v1_slog_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespondLogFind); i {
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
			RawDescriptor: file_api_slog_v1_slog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_slog_v1_slog_proto_goTypes,
		DependencyIndexes: file_api_slog_v1_slog_proto_depIdxs,
		MessageInfos:      file_api_slog_v1_slog_proto_msgTypes,
	}.Build()
	File_api_slog_v1_slog_proto = out.File
	file_api_slog_v1_slog_proto_rawDesc = nil
	file_api_slog_v1_slog_proto_goTypes = nil
	file_api_slog_v1_slog_proto_depIdxs = nil
}
