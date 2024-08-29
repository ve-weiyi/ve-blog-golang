// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: remark.proto

// proto 包名

package remarkrpc

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

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{0}
}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResp.ProtoReflect.Descriptor instead.
func (*EmptyResp) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{1}
}

type IdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *IdReq) Reset() {
	*x = IdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdReq.ProtoReflect.Descriptor instead.
func (*IdReq) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{2}
}

func (x *IdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int64 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IdsReq) Reset() {
	*x = IdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsReq) ProtoMessage() {}

func (x *IdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdsReq.ProtoReflect.Descriptor instead.
func (*IdsReq) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{3}
}

func (x *IdsReq) GetIds() []int64 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type UserIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UserIdReq) Reset() {
	*x = UserIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIdReq.ProtoReflect.Descriptor instead.
func (*UserIdReq) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{4}
}

func (x *UserIdReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type BatchResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCount int64 `protobuf:"varint,1,opt,name=success_count,json=successCount,proto3" json:"success_count,omitempty"`
}

func (x *BatchResp) Reset() {
	*x = BatchResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResp) ProtoMessage() {}

func (x *BatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchResp.ProtoReflect.Descriptor instead.
func (*BatchResp) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{5}
}

func (x *BatchResp) GetSuccessCount() int64 {
	if x != nil {
		return x.SuccessCount
	}
	return 0
}

type CountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountResp) Reset() {
	*x = CountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResp) ProtoMessage() {}

func (x *CountResp) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResp.ProtoReflect.Descriptor instead.
func (*CountResp) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{6}
}

func (x *CountResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type RemarkNew struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                              // 主键id
	Nickname       string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                                   // 昵称
	Avatar         string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                                       // 头像
	MessageContent string `protobuf:"bytes,4,opt,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"` // 留言内容
	IpAddress      string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`                // 用户ip
	IpSource       string `protobuf:"bytes,6,opt,name=ip_source,json=ipSource,proto3" json:"ip_source,omitempty"`                   // 用户地址
	Time           int64  `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`                                          // 弹幕速度
	IsReview       int64  `protobuf:"varint,8,opt,name=is_review,json=isReview,proto3" json:"is_review,omitempty"`                  // 是否审核
}

func (x *RemarkNew) Reset() {
	*x = RemarkNew{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemarkNew) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemarkNew) ProtoMessage() {}

func (x *RemarkNew) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemarkNew.ProtoReflect.Descriptor instead.
func (*RemarkNew) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{7}
}

func (x *RemarkNew) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RemarkNew) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RemarkNew) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RemarkNew) GetMessageContent() string {
	if x != nil {
		return x.MessageContent
	}
	return ""
}

func (x *RemarkNew) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *RemarkNew) GetIpSource() string {
	if x != nil {
		return x.IpSource
	}
	return ""
}

func (x *RemarkNew) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RemarkNew) GetIsReview() int64 {
	if x != nil {
		return x.IsReview
	}
	return 0
}

type RemarkDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                              // 主键id
	Nickname       string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`                                   // 昵称
	Avatar         string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`                                       // 头像
	MessageContent string `protobuf:"bytes,4,opt,name=message_content,json=messageContent,proto3" json:"message_content,omitempty"` // 留言内容
	IpAddress      string `protobuf:"bytes,5,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`                // 用户ip
	IpSource       string `protobuf:"bytes,6,opt,name=ip_source,json=ipSource,proto3" json:"ip_source,omitempty"`                   // 用户地址
	Time           int64  `protobuf:"varint,7,opt,name=time,proto3" json:"time,omitempty"`                                          // 弹幕速度
	IsReview       int64  `protobuf:"varint,8,opt,name=is_review,json=isReview,proto3" json:"is_review,omitempty"`                  // 是否审核
	CreatedAt      int64  `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`               // 发布时间
	UpdatedAt      int64  `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`              // 更新时间
}

func (x *RemarkDetails) Reset() {
	*x = RemarkDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemarkDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemarkDetails) ProtoMessage() {}

func (x *RemarkDetails) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemarkDetails.ProtoReflect.Descriptor instead.
func (*RemarkDetails) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{8}
}

func (x *RemarkDetails) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RemarkDetails) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RemarkDetails) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *RemarkDetails) GetMessageContent() string {
	if x != nil {
		return x.MessageContent
	}
	return ""
}

func (x *RemarkDetails) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *RemarkDetails) GetIpSource() string {
	if x != nil {
		return x.IpSource
	}
	return ""
}

func (x *RemarkDetails) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *RemarkDetails) GetIsReview() int64 {
	if x != nil {
		return x.IsReview
	}
	return 0
}

func (x *RemarkDetails) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *RemarkDetails) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type FindRemarkListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorts    string `protobuf:"bytes,3,opt,name=sorts,proto3" json:"sorts,omitempty"` // 排序
}

func (x *FindRemarkListReq) Reset() {
	*x = FindRemarkListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRemarkListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRemarkListReq) ProtoMessage() {}

func (x *FindRemarkListReq) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRemarkListReq.ProtoReflect.Descriptor instead.
func (*FindRemarkListReq) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{9}
}

func (x *FindRemarkListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindRemarkListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FindRemarkListReq) GetSorts() string {
	if x != nil {
		return x.Sorts
	}
	return ""
}

type FindRemarkListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*RemarkDetails `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FindRemarkListResp) Reset() {
	*x = FindRemarkListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_remark_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindRemarkListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindRemarkListResp) ProtoMessage() {}

func (x *FindRemarkListResp) ProtoReflect() protoreflect.Message {
	mi := &file_remark_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindRemarkListResp.ProtoReflect.Descriptor instead.
func (*FindRemarkListResp) Descriptor() ([]byte, []int) {
	return file_remark_proto_rawDescGZIP(), []int{10}
}

func (x *FindRemarkListResp) GetList() []*RemarkDetails {
	if x != nil {
		return x.List
	}
	return nil
}

var File_remark_proto protoreflect.FileDescriptor

var file_remark_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x22, 0x0b, 0x0a, 0x09, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x06, 0x49,
	0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x03, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x24, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x30, 0x0a,
	0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x21, 0x0a, 0x09, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xe5, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4e, 0x65, 0x77,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x70, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x70, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x22, 0xa7, 0x02, 0x0a, 0x0d, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x70, 0x53,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x73,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x5a, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f,
	0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72, 0x74, 0x73,
	0x22, 0x42, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x32, 0x86, 0x03, 0x0a, 0x09, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x52,
	0x70, 0x63, 0x12, 0x3b, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x14, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x4e, 0x65, 0x77, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x3e, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x14, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x4e, 0x65, 0x77, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x36, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12,
	0x10, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x14, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x72, 0x65,
	0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x38, 0x0a, 0x0a, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x61,
	0x72, 0x6b, 0x12, 0x10, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x4d,
	0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x1c, 0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e,
	0x64, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x52,
	0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a,
	0x0b, 0x2e, 0x2f, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_remark_proto_rawDescOnce sync.Once
	file_remark_proto_rawDescData = file_remark_proto_rawDesc
)

func file_remark_proto_rawDescGZIP() []byte {
	file_remark_proto_rawDescOnce.Do(func() {
		file_remark_proto_rawDescData = protoimpl.X.CompressGZIP(file_remark_proto_rawDescData)
	})
	return file_remark_proto_rawDescData
}

var file_remark_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_remark_proto_goTypes = []any{
	(*EmptyReq)(nil),           // 0: remarkrpc.EmptyReq
	(*EmptyResp)(nil),          // 1: remarkrpc.EmptyResp
	(*IdReq)(nil),              // 2: remarkrpc.IdReq
	(*IdsReq)(nil),             // 3: remarkrpc.IdsReq
	(*UserIdReq)(nil),          // 4: remarkrpc.UserIdReq
	(*BatchResp)(nil),          // 5: remarkrpc.BatchResp
	(*CountResp)(nil),          // 6: remarkrpc.CountResp
	(*RemarkNew)(nil),          // 7: remarkrpc.RemarkNew
	(*RemarkDetails)(nil),      // 8: remarkrpc.RemarkDetails
	(*FindRemarkListReq)(nil),  // 9: remarkrpc.FindRemarkListReq
	(*FindRemarkListResp)(nil), // 10: remarkrpc.FindRemarkListResp
}
var file_remark_proto_depIdxs = []int32{
	8,  // 0: remarkrpc.FindRemarkListResp.list:type_name -> remarkrpc.RemarkDetails
	7,  // 1: remarkrpc.RemarkRpc.AddRemark:input_type -> remarkrpc.RemarkNew
	7,  // 2: remarkrpc.RemarkRpc.UpdateRemark:input_type -> remarkrpc.RemarkNew
	2,  // 3: remarkrpc.RemarkRpc.DeleteRemark:input_type -> remarkrpc.IdReq
	3,  // 4: remarkrpc.RemarkRpc.DeleteRemarkList:input_type -> remarkrpc.IdsReq
	2,  // 5: remarkrpc.RemarkRpc.FindRemark:input_type -> remarkrpc.IdReq
	9,  // 6: remarkrpc.RemarkRpc.FindRemarkList:input_type -> remarkrpc.FindRemarkListReq
	8,  // 7: remarkrpc.RemarkRpc.AddRemark:output_type -> remarkrpc.RemarkDetails
	8,  // 8: remarkrpc.RemarkRpc.UpdateRemark:output_type -> remarkrpc.RemarkDetails
	5,  // 9: remarkrpc.RemarkRpc.DeleteRemark:output_type -> remarkrpc.BatchResp
	5,  // 10: remarkrpc.RemarkRpc.DeleteRemarkList:output_type -> remarkrpc.BatchResp
	8,  // 11: remarkrpc.RemarkRpc.FindRemark:output_type -> remarkrpc.RemarkDetails
	10, // 12: remarkrpc.RemarkRpc.FindRemarkList:output_type -> remarkrpc.FindRemarkListResp
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_remark_proto_init() }
func file_remark_proto_init() {
	if File_remark_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_remark_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyReq); i {
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
		file_remark_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EmptyResp); i {
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
		file_remark_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*IdReq); i {
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
		file_remark_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*IdsReq); i {
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
		file_remark_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UserIdReq); i {
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
		file_remark_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BatchResp); i {
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
		file_remark_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CountResp); i {
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
		file_remark_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RemarkNew); i {
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
		file_remark_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RemarkDetails); i {
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
		file_remark_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*FindRemarkListReq); i {
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
		file_remark_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*FindRemarkListResp); i {
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
			RawDescriptor: file_remark_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_remark_proto_goTypes,
		DependencyIndexes: file_remark_proto_depIdxs,
		MessageInfos:      file_remark_proto_msgTypes,
	}.Build()
	File_remark_proto = out.File
	file_remark_proto_rawDesc = nil
	file_remark_proto_goTypes = nil
	file_remark_proto_depIdxs = nil
}