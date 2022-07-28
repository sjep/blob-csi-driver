// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: cache_volume_service.proto

package cache_volume_service

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	blob_cache_volume "sigs.k8s.io/blob-csi-driver/pkg/edgecache/blob_cache_volume"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volume *blob_cache_volume.BlobCacheVolume `protobuf:"bytes,1,opt,name=volume,proto3,oneof" json:"volume,omitempty"`
}

func (x *CreateBlobRequest) Reset() {
	*x = CreateBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlobRequest) ProtoMessage() {}

func (x *CreateBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlobRequest.ProtoReflect.Descriptor instead.
func (*CreateBlobRequest) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlobRequest) GetVolume() *blob_cache_volume.BlobCacheVolume {
	if x != nil {
		return x.Volume
	}
	return nil
}

type CreateBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBlobResponse) Reset() {
	*x = CreateBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlobResponse) ProtoMessage() {}

func (x *CreateBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlobResponse.ProtoReflect.Descriptor instead.
func (*CreateBlobResponse) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{1}
}

type DeleteBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       *blob_cache_volume.Name `protobuf:"bytes,1,opt,name=name,proto3,oneof" json:"name,omitempty"`
	DoNotFlush *bool                   `protobuf:"varint,2,opt,name=do_not_flush,json=doNotFlush,proto3,oneof" json:"do_not_flush,omitempty"`
}

func (x *DeleteBlobRequest) Reset() {
	*x = DeleteBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlobRequest) ProtoMessage() {}

func (x *DeleteBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlobRequest.ProtoReflect.Descriptor instead.
func (*DeleteBlobRequest) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteBlobRequest) GetName() *blob_cache_volume.Name {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *DeleteBlobRequest) GetDoNotFlush() bool {
	if x != nil && x.DoNotFlush != nil {
		return *x.DoNotFlush
	}
	return false
}

type DeleteBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteBlobResponse) Reset() {
	*x = DeleteBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBlobResponse) ProtoMessage() {}

func (x *DeleteBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBlobResponse.ProtoReflect.Descriptor instead.
func (*DeleteBlobResponse) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{3}
}

type GetBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Names []*blob_cache_volume.Name `protobuf:"bytes,1,rep,name=names,proto3" json:"names,omitempty"`
}

func (x *GetBlobRequest) Reset() {
	*x = GetBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobRequest) ProtoMessage() {}

func (x *GetBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobRequest.ProtoReflect.Descriptor instead.
func (*GetBlobRequest) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetBlobRequest) GetNames() []*blob_cache_volume.Name {
	if x != nil {
		return x.Names
	}
	return nil
}

type GetBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volumes []*blob_cache_volume.BlobCacheVolume `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *GetBlobResponse) Reset() {
	*x = GetBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlobResponse) ProtoMessage() {}

func (x *GetBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlobResponse.ProtoReflect.Descriptor instead.
func (*GetBlobResponse) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetBlobResponse) GetVolumes() []*blob_cache_volume.BlobCacheVolume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

type ListBlobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListBlobRequest) Reset() {
	*x = ListBlobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlobRequest) ProtoMessage() {}

func (x *ListBlobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlobRequest.ProtoReflect.Descriptor instead.
func (*ListBlobRequest) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{6}
}

type ListBlobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Volumes []*blob_cache_volume.BlobCacheVolume `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
}

func (x *ListBlobResponse) Reset() {
	*x = ListBlobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cache_volume_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBlobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBlobResponse) ProtoMessage() {}

func (x *ListBlobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cache_volume_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBlobResponse.ProtoReflect.Descriptor instead.
func (*ListBlobResponse) Descriptor() ([]byte, []int) {
	return file_cache_volume_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListBlobResponse) GetVolumes() []*blob_cache_volume.BlobCacheVolume {
	if x != nil {
		return x.Volumes
	}
	return nil
}

var File_cache_volume_service_proto protoreflect.FileDescriptor

var file_cache_volume_service_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x1a, 0x17, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3f, 0x0a, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x43, 0x61, 0x63, 0x68, 0x65, 0x56, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x48, 0x00, 0x52, 0x06, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x48,
	0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0c, 0x64, 0x6f,
	0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x01, 0x52, 0x0a, 0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x46, 0x6c, 0x75, 0x73, 0x68, 0x88, 0x01,
	0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64,
	0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x66, 0x6c, 0x75, 0x73, 0x68, 0x22, 0x14, 0x0a, 0x12, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x3f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x62, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52, 0x07, 0x76, 0x6f, 0x6c, 0x75,
	0x6d, 0x65, 0x73, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x50, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x62, 0x6c,
	0x6f, 0x62, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x2e,
	0x42, 0x6c, 0x6f, 0x62, 0x43, 0x61, 0x63, 0x68, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x52,
	0x07, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x73, 0x32, 0x82, 0x03, 0x0a, 0x0b, 0x43, 0x61, 0x63,
	0x68, 0x65, 0x56, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x27, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76,
	0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x28, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f,
	0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5f, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x27, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x28, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x24, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x59, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x12, 0x25,
	0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x63, 0x61, 0x63, 0x68, 0x65, 0x5f, 0x76, 0x6f,
	0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x6c, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a,
	0x3e, 0x73, 0x69, 0x67, 0x73, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x6f,
	0x62, 0x2d, 0x63, 0x73, 0x69, 0x2d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x63, 0x61, 0x63, 0x68, 0x65,
	0x5f, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cache_volume_service_proto_rawDescOnce sync.Once
	file_cache_volume_service_proto_rawDescData = file_cache_volume_service_proto_rawDesc
)

func file_cache_volume_service_proto_rawDescGZIP() []byte {
	file_cache_volume_service_proto_rawDescOnce.Do(func() {
		file_cache_volume_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_cache_volume_service_proto_rawDescData)
	})
	return file_cache_volume_service_proto_rawDescData
}

var file_cache_volume_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cache_volume_service_proto_goTypes = []interface{}{
	(*CreateBlobRequest)(nil),                 // 0: cache_volume_service.CreateBlobRequest
	(*CreateBlobResponse)(nil),                // 1: cache_volume_service.CreateBlobResponse
	(*DeleteBlobRequest)(nil),                 // 2: cache_volume_service.DeleteBlobRequest
	(*DeleteBlobResponse)(nil),                // 3: cache_volume_service.DeleteBlobResponse
	(*GetBlobRequest)(nil),                    // 4: cache_volume_service.GetBlobRequest
	(*GetBlobResponse)(nil),                   // 5: cache_volume_service.GetBlobResponse
	(*ListBlobRequest)(nil),                   // 6: cache_volume_service.ListBlobRequest
	(*ListBlobResponse)(nil),                  // 7: cache_volume_service.ListBlobResponse
	(*blob_cache_volume.BlobCacheVolume)(nil), // 8: blob_cache_volume.BlobCacheVolume
	(*blob_cache_volume.Name)(nil),            // 9: blob_cache_volume.Name
}
var file_cache_volume_service_proto_depIdxs = []int32{
	8, // 0: cache_volume_service.CreateBlobRequest.volume:type_name -> blob_cache_volume.BlobCacheVolume
	9, // 1: cache_volume_service.DeleteBlobRequest.name:type_name -> blob_cache_volume.Name
	9, // 2: cache_volume_service.GetBlobRequest.names:type_name -> blob_cache_volume.Name
	8, // 3: cache_volume_service.GetBlobResponse.volumes:type_name -> blob_cache_volume.BlobCacheVolume
	8, // 4: cache_volume_service.ListBlobResponse.volumes:type_name -> blob_cache_volume.BlobCacheVolume
	0, // 5: cache_volume_service.CacheVolume.CreateBlob:input_type -> cache_volume_service.CreateBlobRequest
	2, // 6: cache_volume_service.CacheVolume.DeleteBlob:input_type -> cache_volume_service.DeleteBlobRequest
	4, // 7: cache_volume_service.CacheVolume.GetBlob:input_type -> cache_volume_service.GetBlobRequest
	6, // 8: cache_volume_service.CacheVolume.ListBlob:input_type -> cache_volume_service.ListBlobRequest
	1, // 9: cache_volume_service.CacheVolume.CreateBlob:output_type -> cache_volume_service.CreateBlobResponse
	3, // 10: cache_volume_service.CacheVolume.DeleteBlob:output_type -> cache_volume_service.DeleteBlobResponse
	5, // 11: cache_volume_service.CacheVolume.GetBlob:output_type -> cache_volume_service.GetBlobResponse
	7, // 12: cache_volume_service.CacheVolume.ListBlob:output_type -> cache_volume_service.ListBlobResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cache_volume_service_proto_init() }
func file_cache_volume_service_proto_init() {
	if File_cache_volume_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cache_volume_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlobRequest); i {
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
		file_cache_volume_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlobResponse); i {
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
		file_cache_volume_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlobRequest); i {
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
		file_cache_volume_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBlobResponse); i {
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
		file_cache_volume_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobRequest); i {
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
		file_cache_volume_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlobResponse); i {
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
		file_cache_volume_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlobRequest); i {
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
		file_cache_volume_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBlobResponse); i {
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
	file_cache_volume_service_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_cache_volume_service_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cache_volume_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cache_volume_service_proto_goTypes,
		DependencyIndexes: file_cache_volume_service_proto_depIdxs,
		MessageInfos:      file_cache_volume_service_proto_msgTypes,
	}.Build()
	File_cache_volume_service_proto = out.File
	file_cache_volume_service_proto_rawDesc = nil
	file_cache_volume_service_proto_goTypes = nil
	file_cache_volume_service_proto_depIdxs = nil
}