// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.4
// source: syslog.proto

// proto 包名

package syslogrpc

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
		mi := &file_syslog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[0]
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
	return file_syslog_proto_rawDescGZIP(), []int{0}
}

type EmptyResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResp) Reset() {
	*x = EmptyResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResp) ProtoMessage() {}

func (x *EmptyResp) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[1]
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
	return file_syslog_proto_rawDescGZIP(), []int{1}
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
		mi := &file_syslog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdReq) ProtoMessage() {}

func (x *IdReq) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[2]
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
	return file_syslog_proto_rawDescGZIP(), []int{2}
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
		mi := &file_syslog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdsReq) ProtoMessage() {}

func (x *IdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[3]
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
	return file_syslog_proto_rawDescGZIP(), []int{3}
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
		mi := &file_syslog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIdReq) ProtoMessage() {}

func (x *UserIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[4]
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
	return file_syslog_proto_rawDescGZIP(), []int{4}
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
		mi := &file_syslog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchResp) ProtoMessage() {}

func (x *BatchResp) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[5]
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
	return file_syslog_proto_rawDescGZIP(), []int{5}
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
		mi := &file_syslog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResp) ProtoMessage() {}

func (x *CountResp) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[6]
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
	return file_syslog_proto_rawDescGZIP(), []int{6}
}

func (x *CountResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// ************* blog 日志管理 *************
type OperationLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                // 主键id
	UserId         int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                          // 用户id
	Nickname       string `protobuf:"bytes,3,opt,name=nickname,proto3" json:"nickname,omitempty"`                                     // 用户昵称
	IpAddress      string `protobuf:"bytes,4,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`                  // 操作ip
	IpSource       string `protobuf:"bytes,5,opt,name=ip_source,json=ipSource,proto3" json:"ip_source,omitempty"`                     // 操作地址
	OptModule      string `protobuf:"bytes,6,opt,name=opt_module,json=optModule,proto3" json:"opt_module,omitempty"`                  // 操作模块
	OptDesc        string `protobuf:"bytes,7,opt,name=opt_desc,json=optDesc,proto3" json:"opt_desc,omitempty"`                        // 操作描述
	RequestUrl     string `protobuf:"bytes,8,opt,name=request_url,json=requestUrl,proto3" json:"request_url,omitempty"`               // 请求地址
	RequestMethod  string `protobuf:"bytes,9,opt,name=request_method,json=requestMethod,proto3" json:"request_method,omitempty"`      // 请求方式
	RequestHeader  string `protobuf:"bytes,10,opt,name=request_header,json=requestHeader,proto3" json:"request_header,omitempty"`     // 请求头参数
	RequestData    string `protobuf:"bytes,11,opt,name=request_data,json=requestData,proto3" json:"request_data,omitempty"`           // 请求参数
	ResponseData   string `protobuf:"bytes,12,opt,name=response_data,json=responseData,proto3" json:"response_data,omitempty"`        // 返回数据
	ResponseStatus int64  `protobuf:"varint,13,opt,name=response_status,json=responseStatus,proto3" json:"response_status,omitempty"` // 响应状态码
	Cost           string `protobuf:"bytes,14,opt,name=cost,proto3" json:"cost,omitempty"`                                            // 耗时（ms）
	CreatedAt      int64  `protobuf:"varint,15,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                // 创建时间
	UpdatedAt      int64  `protobuf:"varint,16,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                // 更新时间
}

func (x *OperationLog) Reset() {
	*x = OperationLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperationLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperationLog) ProtoMessage() {}

func (x *OperationLog) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperationLog.ProtoReflect.Descriptor instead.
func (*OperationLog) Descriptor() ([]byte, []int) {
	return file_syslog_proto_rawDescGZIP(), []int{7}
}

func (x *OperationLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OperationLog) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *OperationLog) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *OperationLog) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *OperationLog) GetIpSource() string {
	if x != nil {
		return x.IpSource
	}
	return ""
}

func (x *OperationLog) GetOptModule() string {
	if x != nil {
		return x.OptModule
	}
	return ""
}

func (x *OperationLog) GetOptDesc() string {
	if x != nil {
		return x.OptDesc
	}
	return ""
}

func (x *OperationLog) GetRequestUrl() string {
	if x != nil {
		return x.RequestUrl
	}
	return ""
}

func (x *OperationLog) GetRequestMethod() string {
	if x != nil {
		return x.RequestMethod
	}
	return ""
}

func (x *OperationLog) GetRequestHeader() string {
	if x != nil {
		return x.RequestHeader
	}
	return ""
}

func (x *OperationLog) GetRequestData() string {
	if x != nil {
		return x.RequestData
	}
	return ""
}

func (x *OperationLog) GetResponseData() string {
	if x != nil {
		return x.ResponseData
	}
	return ""
}

func (x *OperationLog) GetResponseStatus() int64 {
	if x != nil {
		return x.ResponseStatus
	}
	return 0
}

func (x *OperationLog) GetCost() string {
	if x != nil {
		return x.Cost
	}
	return ""
}

func (x *OperationLog) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *OperationLog) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type FindOperationLogListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int64  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PageSize int64  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	Sorts    string `protobuf:"bytes,3,opt,name=sorts,proto3" json:"sorts,omitempty"` // 排序
}

func (x *FindOperationLogListReq) Reset() {
	*x = FindOperationLogListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOperationLogListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOperationLogListReq) ProtoMessage() {}

func (x *FindOperationLogListReq) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOperationLogListReq.ProtoReflect.Descriptor instead.
func (*FindOperationLogListReq) Descriptor() ([]byte, []int) {
	return file_syslog_proto_rawDescGZIP(), []int{8}
}

func (x *FindOperationLogListReq) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *FindOperationLogListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *FindOperationLogListReq) GetSorts() string {
	if x != nil {
		return x.Sorts
	}
	return ""
}

type FindOperationLogListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*OperationLog `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *FindOperationLogListResp) Reset() {
	*x = FindOperationLogListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOperationLogListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOperationLogListResp) ProtoMessage() {}

func (x *FindOperationLogListResp) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOperationLogListResp.ProtoReflect.Descriptor instead.
func (*FindOperationLogListResp) Descriptor() ([]byte, []int) {
	return file_syslog_proto_rawDescGZIP(), []int{9}
}

func (x *FindOperationLogListResp) GetList() []*OperationLog {
	if x != nil {
		return x.List
	}
	return nil
}

type UploadLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                             // id
	UserId   int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`       // 用户id
	Label    string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`                        // 标签
	FileName string `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`  // 文件名称
	FileSize int64  `protobuf:"varint,5,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"` // 文件大小
	FileMd5  string `protobuf:"bytes,6,opt,name=file_md5,json=fileMd5,proto3" json:"file_md5,omitempty"`     // 文件md5值
	FileUrl  string `protobuf:"bytes,7,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`     // 上传路径
}

func (x *UploadLogReq) Reset() {
	*x = UploadLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslog_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLogReq) ProtoMessage() {}

func (x *UploadLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadLogReq.ProtoReflect.Descriptor instead.
func (*UploadLogReq) Descriptor() ([]byte, []int) {
	return file_syslog_proto_rawDescGZIP(), []int{10}
}

func (x *UploadLogReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadLogReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UploadLogReq) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UploadLogReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadLogReq) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *UploadLogReq) GetFileMd5() string {
	if x != nil {
		return x.FileMd5
	}
	return ""
}

func (x *UploadLogReq) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type UploadLogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                // id
	UserId    int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`          // 用户id
	Label     string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`                           // 标签
	FileName  string `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`     // 文件名称
	FileSize  int64  `protobuf:"varint,5,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`    // 文件大小
	FileMd5   string `protobuf:"bytes,6,opt,name=file_md5,json=fileMd5,proto3" json:"file_md5,omitempty"`        // 文件md5值
	FileUrl   string `protobuf:"bytes,7,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`        // 上传路径
	CreatedAt int64  `protobuf:"varint,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"` // 创建时间
	UpdatedAt int64  `protobuf:"varint,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"` // 更新时间
}

func (x *UploadLogResp) Reset() {
	*x = UploadLogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_syslog_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadLogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadLogResp) ProtoMessage() {}

func (x *UploadLogResp) ProtoReflect() protoreflect.Message {
	mi := &file_syslog_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadLogResp.ProtoReflect.Descriptor instead.
func (*UploadLogResp) Descriptor() ([]byte, []int) {
	return file_syslog_proto_rawDescGZIP(), []int{11}
}

func (x *UploadLogResp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UploadLogResp) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UploadLogResp) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UploadLogResp) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadLogResp) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *UploadLogResp) GetFileMd5() string {
	if x != nil {
		return x.FileMd5
	}
	return ""
}

func (x *UploadLogResp) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

func (x *UploadLogResp) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UploadLogResp) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_syslog_proto protoreflect.FileDescriptor

var file_syslog_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70,
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
	0x6e, 0x74, 0x22, 0xfb, 0x03, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x70, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x70, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6f, 0x70, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x70, 0x74, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x70, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x44, 0x65, 0x73, 0x63, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x55, 0x72, 0x6c, 0x12,
	0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x60, 0x0a, 0x17, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x6f, 0x72,
	0x74, 0x73, 0x22, 0x47, 0x0a, 0x18, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2b,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73,
	0x79, 0x73, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xbd, 0x01, 0x0a, 0x0c,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x64,
	0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x64, 0x35,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0xfc, 0x01, 0x0a, 0x0d,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d,
	0x64, 0x35, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x64,
	0x35, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xb7, 0x02, 0x0a, 0x09, 0x53,
	0x79, 0x73, 0x6c, 0x6f, 0x67, 0x52, 0x70, 0x63, 0x12, 0x43, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x17, 0x2e, 0x73, 0x79,
	0x73, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x1a, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x12, 0x41, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x73, 0x79, 0x73,
	0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x5f, 0x0a, 0x14, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x73,
	0x79, 0x73, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x41, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f,
	0x67, 0x12, 0x17, 0x2e, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x73, 0x79, 0x73,
	0x6c, 0x6f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x6f, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x73, 0x79, 0x73, 0x6c, 0x6f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_syslog_proto_rawDescOnce sync.Once
	file_syslog_proto_rawDescData = file_syslog_proto_rawDesc
)

func file_syslog_proto_rawDescGZIP() []byte {
	file_syslog_proto_rawDescOnce.Do(func() {
		file_syslog_proto_rawDescData = protoimpl.X.CompressGZIP(file_syslog_proto_rawDescData)
	})
	return file_syslog_proto_rawDescData
}

var file_syslog_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_syslog_proto_goTypes = []any{
	(*EmptyReq)(nil),                 // 0: syslogrpc.EmptyReq
	(*EmptyResp)(nil),                // 1: syslogrpc.EmptyResp
	(*IdReq)(nil),                    // 2: syslogrpc.IdReq
	(*IdsReq)(nil),                   // 3: syslogrpc.IdsReq
	(*UserIdReq)(nil),                // 4: syslogrpc.UserIdReq
	(*BatchResp)(nil),                // 5: syslogrpc.BatchResp
	(*CountResp)(nil),                // 6: syslogrpc.CountResp
	(*OperationLog)(nil),             // 7: syslogrpc.OperationLog
	(*FindOperationLogListReq)(nil),  // 8: syslogrpc.FindOperationLogListReq
	(*FindOperationLogListResp)(nil), // 9: syslogrpc.FindOperationLogListResp
	(*UploadLogReq)(nil),             // 10: syslogrpc.UploadLogReq
	(*UploadLogResp)(nil),            // 11: syslogrpc.UploadLogResp
}
var file_syslog_proto_depIdxs = []int32{
	7,  // 0: syslogrpc.FindOperationLogListResp.list:type_name -> syslogrpc.OperationLog
	7,  // 1: syslogrpc.SyslogRpc.AddOperationLog:input_type -> syslogrpc.OperationLog
	3,  // 2: syslogrpc.SyslogRpc.DeleteOperationLogList:input_type -> syslogrpc.IdsReq
	8,  // 3: syslogrpc.SyslogRpc.FindOperationLogList:input_type -> syslogrpc.FindOperationLogListReq
	10, // 4: syslogrpc.SyslogRpc.AddUploadLog:input_type -> syslogrpc.UploadLogReq
	7,  // 5: syslogrpc.SyslogRpc.AddOperationLog:output_type -> syslogrpc.OperationLog
	5,  // 6: syslogrpc.SyslogRpc.DeleteOperationLogList:output_type -> syslogrpc.BatchResp
	9,  // 7: syslogrpc.SyslogRpc.FindOperationLogList:output_type -> syslogrpc.FindOperationLogListResp
	11, // 8: syslogrpc.SyslogRpc.AddUploadLog:output_type -> syslogrpc.UploadLogResp
	5,  // [5:9] is the sub-list for method output_type
	1,  // [1:5] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_syslog_proto_init() }
func file_syslog_proto_init() {
	if File_syslog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_syslog_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[3].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[5].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[6].Exporter = func(v any, i int) any {
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
		file_syslog_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*OperationLog); i {
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
		file_syslog_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*FindOperationLogListReq); i {
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
		file_syslog_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*FindOperationLogListResp); i {
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
		file_syslog_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*UploadLogReq); i {
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
		file_syslog_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*UploadLogResp); i {
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
			RawDescriptor: file_syslog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_syslog_proto_goTypes,
		DependencyIndexes: file_syslog_proto_depIdxs,
		MessageInfos:      file_syslog_proto_msgTypes,
	}.Build()
	File_syslog_proto = out.File
	file_syslog_proto_rawDesc = nil
	file_syslog_proto_goTypes = nil
	file_syslog_proto_depIdxs = nil
}