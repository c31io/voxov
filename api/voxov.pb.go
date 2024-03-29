// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: voxov.proto

package api

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

type CreateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion int32 `protobuf:"varint,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Ttl        int64 `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *CreateSessionRequest) Reset() {
	*x = CreateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionRequest) ProtoMessage() {}

func (x *CreateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionRequest.ProtoReflect.Descriptor instead.
func (*CreateSessionRequest) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSessionRequest) GetApiVersion() int32 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *CreateSessionRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type CreateSessionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion int32  `protobuf:"varint,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Token      []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CreateSessionReply) Reset() {
	*x = CreateSessionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSessionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSessionReply) ProtoMessage() {}

func (x *CreateSessionReply) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSessionReply.ProtoReflect.Descriptor instead.
func (*CreateSessionReply) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSessionReply) GetApiVersion() int32 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *CreateSessionReply) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type UpdateSessionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ttl   int64  `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Token []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *UpdateSessionRequest) Reset() {
	*x = UpdateSessionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSessionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSessionRequest) ProtoMessage() {}

func (x *UpdateSessionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSessionRequest.ProtoReflect.Descriptor instead.
func (*UpdateSessionRequest) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateSessionRequest) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *UpdateSessionRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type UpdateSessionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *UpdateSessionReply) Reset() {
	*x = UpdateSessionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSessionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSessionReply) ProtoMessage() {}

func (x *UpdateSessionReply) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSessionReply.ProtoReflect.Descriptor instead.
func (*UpdateSessionReply) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateSessionReply) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type AuthenticateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthenticateRequest) Reset() {
	*x = AuthenticateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateRequest) ProtoMessage() {}

func (x *AuthenticateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateRequest.ProtoReflect.Descriptor instead.
func (*AuthenticateRequest) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{4}
}

func (x *AuthenticateRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type AuthenticateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tel string `protobuf:"bytes,1,opt,name=tel,proto3" json:"tel,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AuthenticateReply) Reset() {
	*x = AuthenticateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticateReply) ProtoMessage() {}

func (x *AuthenticateReply) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticateReply.ProtoReflect.Descriptor instead.
func (*AuthenticateReply) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{5}
}

func (x *AuthenticateReply) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *AuthenticateReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type WhoAmIRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tel   string `protobuf:"bytes,1,opt,name=tel,proto3" json:"tel,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Token []byte `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *WhoAmIRequest) Reset() {
	*x = WhoAmIRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIRequest) ProtoMessage() {}

func (x *WhoAmIRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIRequest.ProtoReflect.Descriptor instead.
func (*WhoAmIRequest) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{6}
}

func (x *WhoAmIRequest) GetTel() string {
	if x != nil {
		return x.Tel
	}
	return ""
}

func (x *WhoAmIRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *WhoAmIRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type WhoAmIReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person int64 `protobuf:"varint,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *WhoAmIReply) Reset() {
	*x = WhoAmIReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WhoAmIReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WhoAmIReply) ProtoMessage() {}

func (x *WhoAmIReply) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WhoAmIReply.ProtoReflect.Descriptor instead.
func (*WhoAmIReply) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{7}
}

func (x *WhoAmIReply) GetPerson() int64 {
	if x != nil {
		return x.Person
	}
	return 0
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did     int64  `protobuf:"varint,1,opt,name=did,proto3" json:"did,omitempty"`
	Dtoken  []byte `protobuf:"bytes,2,opt,name=dtoken,proto3" json:"dtoken,omitempty"`
	Dname   string `protobuf:"bytes,3,opt,name=dname,proto3" json:"dname,omitempty"`
	Dinfo   string `protobuf:"bytes,4,opt,name=dinfo,proto3" json:"dinfo,omitempty"`
	Pid     int64  `protobuf:"varint,5,opt,name=pid,proto3" json:"pid,omitempty"`
	Created int64  `protobuf:"varint,6,opt,name=created,proto3" json:"created,omitempty"` // UTC UnixMicro
	LastIn  int64  `protobuf:"varint,7,opt,name=last_in,json=lastIn,proto3" json:"last_in,omitempty"`
	Token   []byte `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{8}
}

func (x *Device) GetDid() int64 {
	if x != nil {
		return x.Did
	}
	return 0
}

func (x *Device) GetDtoken() []byte {
	if x != nil {
		return x.Dtoken
	}
	return nil
}

func (x *Device) GetDname() string {
	if x != nil {
		return x.Dname
	}
	return ""
}

func (x *Device) GetDinfo() string {
	if x != nil {
		return x.Dinfo
	}
	return ""
}

func (x *Device) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Device) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Device) GetLastIn() int64 {
	if x != nil {
		return x.LastIn
	}
	return 0
}

func (x *Device) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type ListDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token []byte `protobuf:"bytes,8,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ListDeviceRequest) Reset() {
	*x = ListDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceRequest) ProtoMessage() {}

func (x *ListDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceRequest.ProtoReflect.Descriptor instead.
func (*ListDeviceRequest) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{9}
}

func (x *ListDeviceRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type ListDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Devices []*Device `protobuf:"bytes,1,rep,name=devices,proto3" json:"devices,omitempty"`
}

func (x *ListDeviceReply) Reset() {
	*x = ListDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDeviceReply) ProtoMessage() {}

func (x *ListDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDeviceReply.ProtoReflect.Descriptor instead.
func (*ListDeviceReply) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{10}
}

func (x *ListDeviceReply) GetDevices() []*Device {
	if x != nil {
		return x.Devices
	}
	return nil
}

type AuthDeviceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dtoken []byte `protobuf:"bytes,1,opt,name=dtoken,proto3" json:"dtoken,omitempty"`
	Token  []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthDeviceRequest) Reset() {
	*x = AuthDeviceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthDeviceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthDeviceRequest) ProtoMessage() {}

func (x *AuthDeviceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthDeviceRequest.ProtoReflect.Descriptor instead.
func (*AuthDeviceRequest) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{11}
}

func (x *AuthDeviceRequest) GetDtoken() []byte {
	if x != nil {
		return x.Dtoken
	}
	return nil
}

func (x *AuthDeviceRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type AuthDeviceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Did int64 `protobuf:"varint,1,opt,name=did,proto3" json:"did,omitempty"`
	Pid int64 `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *AuthDeviceReply) Reset() {
	*x = AuthDeviceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthDeviceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthDeviceReply) ProtoMessage() {}

func (x *AuthDeviceReply) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthDeviceReply.ProtoReflect.Descriptor instead.
func (*AuthDeviceReply) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{12}
}

func (x *AuthDeviceReply) GetDid() int64 {
	if x != nil {
		return x.Did
	}
	return 0
}

func (x *AuthDeviceReply) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid     int64  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Hid     int64  `protobuf:"varint,2,opt,name=hid,proto3" json:"hid,omitempty"`
	Balance int64  `protobuf:"varint,3,opt,name=balance,proto3" json:"balance,omitempty"`
	Phone   string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Pname   string `protobuf:"bytes,5,opt,name=pname,proto3" json:"pname,omitempty"`
	IdDoc   string `protobuf:"bytes,6,opt,name=id_doc,json=idDoc,proto3" json:"id_doc,omitempty"`
	Dlimit  int32  `protobuf:"varint,7,opt,name=dlimit,proto3" json:"dlimit,omitempty"`
	Created int64  `protobuf:"varint,8,opt,name=created,proto3" json:"created,omitempty"`
	LastIn  int64  `protobuf:"varint,9,opt,name=last_in,json=lastIn,proto3" json:"last_in,omitempty"`
	Token   []byte `protobuf:"bytes,10,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voxov_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_voxov_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_voxov_proto_rawDescGZIP(), []int{13}
}

func (x *Person) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Person) GetHid() int64 {
	if x != nil {
		return x.Hid
	}
	return 0
}

func (x *Person) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *Person) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Person) GetPname() string {
	if x != nil {
		return x.Pname
	}
	return ""
}

func (x *Person) GetIdDoc() string {
	if x != nil {
		return x.IdDoc
	}
	return ""
}

func (x *Person) GetDlimit() int32 {
	if x != nil {
		return x.Dlimit
	}
	return 0
}

func (x *Person) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Person) GetLastIn() int64 {
	if x != nil {
		return x.LastIn
	}
	return 0
}

func (x *Person) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

var File_voxov_proto protoreflect.FileDescriptor

var file_voxov_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x76,
	0x6f, 0x78, 0x6f, 0x76, 0x22, 0x49, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x22,
	0x4b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x22, 0x2b, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x37, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x49, 0x0a, 0x0d, 0x57, 0x68, 0x6f, 0x41,
	0x6d, 0x49, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x25, 0x0a, 0x0b, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0xb9, 0x01, 0x0a, 0x06, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x3a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x27, 0x0a, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x41, 0x0a,
	0x11, 0x41, 0x75, 0x74, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x64, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x35, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x64, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0xea, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x68, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06,
	0x69, 0x64, 0x5f, 0x64, 0x6f, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x64,
	0x44, 0x6f, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x64, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69, 0x6e,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x6e, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xf5, 0x04, 0x0a, 0x05, 0x56, 0x4f, 0x78, 0x4f, 0x56, 0x12, 0x47,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x1b, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76,
	0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x47, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x44, 0x0a, 0x0c, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x1a, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x76,
	0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49,
	0x12, 0x14, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x57,
	0x68, 0x6f, 0x41, 0x6d, 0x49, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x0c, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x76, 0x6f, 0x78,
	0x6f, 0x76, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f,
	0x76, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x0d, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18,
	0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x3e, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18,
	0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2a, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0d,
	0x2e, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e,
	0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x1c, 0x5a, 0x1a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x33, 0x31, 0x69, 0x6f,
	0x2f, 0x76, 0x6f, 0x78, 0x6f, 0x76, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_voxov_proto_rawDescOnce sync.Once
	file_voxov_proto_rawDescData = file_voxov_proto_rawDesc
)

func file_voxov_proto_rawDescGZIP() []byte {
	file_voxov_proto_rawDescOnce.Do(func() {
		file_voxov_proto_rawDescData = protoimpl.X.CompressGZIP(file_voxov_proto_rawDescData)
	})
	return file_voxov_proto_rawDescData
}

var file_voxov_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_voxov_proto_goTypes = []interface{}{
	(*CreateSessionRequest)(nil), // 0: voxov.CreateSessionRequest
	(*CreateSessionReply)(nil),   // 1: voxov.CreateSessionReply
	(*UpdateSessionRequest)(nil), // 2: voxov.UpdateSessionRequest
	(*UpdateSessionReply)(nil),   // 3: voxov.UpdateSessionReply
	(*AuthenticateRequest)(nil),  // 4: voxov.AuthenticateRequest
	(*AuthenticateReply)(nil),    // 5: voxov.AuthenticateReply
	(*WhoAmIRequest)(nil),        // 6: voxov.WhoAmIRequest
	(*WhoAmIReply)(nil),          // 7: voxov.WhoAmIReply
	(*Device)(nil),               // 8: voxov.Device
	(*ListDeviceRequest)(nil),    // 9: voxov.ListDeviceRequest
	(*ListDeviceReply)(nil),      // 10: voxov.ListDeviceReply
	(*AuthDeviceRequest)(nil),    // 11: voxov.AuthDeviceRequest
	(*AuthDeviceReply)(nil),      // 12: voxov.AuthDeviceReply
	(*Person)(nil),               // 13: voxov.Person
}
var file_voxov_proto_depIdxs = []int32{
	8,  // 0: voxov.ListDeviceReply.devices:type_name -> voxov.Device
	0,  // 1: voxov.VOxOV.CreateSession:input_type -> voxov.CreateSessionRequest
	2,  // 2: voxov.VOxOV.UpdateSession:input_type -> voxov.UpdateSessionRequest
	4,  // 3: voxov.VOxOV.Authenticate:input_type -> voxov.AuthenticateRequest
	6,  // 4: voxov.VOxOV.WhoAmI:input_type -> voxov.WhoAmIRequest
	8,  // 5: voxov.VOxOV.CreateDevice:input_type -> voxov.Device
	8,  // 6: voxov.VOxOV.ReadDevice:input_type -> voxov.Device
	8,  // 7: voxov.VOxOV.UpdateDevice:input_type -> voxov.Device
	8,  // 8: voxov.VOxOV.DeleteDevice:input_type -> voxov.Device
	9,  // 9: voxov.VOxOV.ListDevice:input_type -> voxov.ListDeviceRequest
	11, // 10: voxov.VOxOV.AuthDevice:input_type -> voxov.AuthDeviceRequest
	13, // 11: voxov.VOxOV.ReadPerson:input_type -> voxov.Person
	1,  // 12: voxov.VOxOV.CreateSession:output_type -> voxov.CreateSessionReply
	3,  // 13: voxov.VOxOV.UpdateSession:output_type -> voxov.UpdateSessionReply
	5,  // 14: voxov.VOxOV.Authenticate:output_type -> voxov.AuthenticateReply
	7,  // 15: voxov.VOxOV.WhoAmI:output_type -> voxov.WhoAmIReply
	8,  // 16: voxov.VOxOV.CreateDevice:output_type -> voxov.Device
	8,  // 17: voxov.VOxOV.ReadDevice:output_type -> voxov.Device
	8,  // 18: voxov.VOxOV.UpdateDevice:output_type -> voxov.Device
	8,  // 19: voxov.VOxOV.DeleteDevice:output_type -> voxov.Device
	10, // 20: voxov.VOxOV.ListDevice:output_type -> voxov.ListDeviceReply
	12, // 21: voxov.VOxOV.AuthDevice:output_type -> voxov.AuthDeviceReply
	13, // 22: voxov.VOxOV.ReadPerson:output_type -> voxov.Person
	12, // [12:23] is the sub-list for method output_type
	1,  // [1:12] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_voxov_proto_init() }
func file_voxov_proto_init() {
	if File_voxov_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_voxov_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionRequest); i {
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
		file_voxov_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSessionReply); i {
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
		file_voxov_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSessionRequest); i {
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
		file_voxov_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSessionReply); i {
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
		file_voxov_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateRequest); i {
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
		file_voxov_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticateReply); i {
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
		file_voxov_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIRequest); i {
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
		file_voxov_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WhoAmIReply); i {
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
		file_voxov_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_voxov_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceRequest); i {
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
		file_voxov_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListDeviceReply); i {
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
		file_voxov_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthDeviceRequest); i {
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
		file_voxov_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthDeviceReply); i {
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
		file_voxov_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
			RawDescriptor: file_voxov_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voxov_proto_goTypes,
		DependencyIndexes: file_voxov_proto_depIdxs,
		MessageInfos:      file_voxov_proto_msgTypes,
	}.Build()
	File_voxov_proto = out.File
	file_voxov_proto_rawDesc = nil
	file_voxov_proto_goTypes = nil
	file_voxov_proto_depIdxs = nil
}
