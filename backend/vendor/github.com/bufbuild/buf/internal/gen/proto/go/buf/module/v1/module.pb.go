// Copyright 2020 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: buf/module/v1/module.proto

package modulev1

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

// Module is a module.
type Module struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// files are the files that make up the set.
	//
	// Sorted by path.
	// Path must be unique.
	// Only the target files. No imports.
	//
	// Maximum total size of all content: 32MB.
	Files []*ModuleFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// dependencies are the dependencies.
	//
	// These must be resolved.
	Dependencies []*ModuleName `protobuf:"bytes,2,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *Module) Reset() {
	*x = Module{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_module_v1_module_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Module) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Module) ProtoMessage() {}

func (x *Module) ProtoReflect() protoreflect.Message {
	mi := &file_buf_module_v1_module_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Module.ProtoReflect.Descriptor instead.
func (*Module) Descriptor() ([]byte, []int) {
	return file_buf_module_v1_module_proto_rawDescGZIP(), []int{0}
}

func (x *Module) GetFiles() []*ModuleFile {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *Module) GetDependencies() []*ModuleName {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

// ModuleFile is a file within a FileSet.
type ModuleFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// path is the relative path of the file.
	// Path can only use '/' as the separator character, and includes no ".." components.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// content is the content of the file.
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ModuleFile) Reset() {
	*x = ModuleFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_module_v1_module_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleFile) ProtoMessage() {}

func (x *ModuleFile) ProtoReflect() protoreflect.Message {
	mi := &file_buf_module_v1_module_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleFile.ProtoReflect.Descriptor instead.
func (*ModuleFile) Descriptor() ([]byte, []int) {
	return file_buf_module_v1_module_proto_rawDescGZIP(), []int{1}
}

func (x *ModuleFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ModuleFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type ModuleName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DNS name or IP address of the server that hosts the module.
	Remote string `protobuf:"bytes,1,opt,name=remote,proto3" json:"remote,omitempty"`
	// The users username or organization name
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	// The repository name
	Repository string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
	// The repository track
	Track string `protobuf:"bytes,4,opt,name=track,proto3" json:"track,omitempty"`
	// The repository commit digest
	Digest string `protobuf:"bytes,5,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *ModuleName) Reset() {
	*x = ModuleName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_module_v1_module_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuleName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuleName) ProtoMessage() {}

func (x *ModuleName) ProtoReflect() protoreflect.Message {
	mi := &file_buf_module_v1_module_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuleName.ProtoReflect.Descriptor instead.
func (*ModuleName) Descriptor() ([]byte, []int) {
	return file_buf_module_v1_module_proto_rawDescGZIP(), []int{2}
}

func (x *ModuleName) GetRemote() string {
	if x != nil {
		return x.Remote
	}
	return ""
}

func (x *ModuleName) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *ModuleName) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *ModuleName) GetTrack() string {
	if x != nil {
		return x.Track
	}
	return ""
}

func (x *ModuleName) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

var File_buf_module_v1_module_proto protoreflect.FileDescriptor

var file_buf_module_v1_module_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x62, 0x75, 0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x75,
	0x66, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x78, 0x0a, 0x06, 0x4d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64,
	0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x88, 0x01, 0x0a, 0x0a, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x42, 0x46, 0x5a, 0x44,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62, 0x75,
	0x66, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_module_v1_module_proto_rawDescOnce sync.Once
	file_buf_module_v1_module_proto_rawDescData = file_buf_module_v1_module_proto_rawDesc
)

func file_buf_module_v1_module_proto_rawDescGZIP() []byte {
	file_buf_module_v1_module_proto_rawDescOnce.Do(func() {
		file_buf_module_v1_module_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_module_v1_module_proto_rawDescData)
	})
	return file_buf_module_v1_module_proto_rawDescData
}

var file_buf_module_v1_module_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_module_v1_module_proto_goTypes = []interface{}{
	(*Module)(nil),     // 0: buf.module.v1.Module
	(*ModuleFile)(nil), // 1: buf.module.v1.ModuleFile
	(*ModuleName)(nil), // 2: buf.module.v1.ModuleName
}
var file_buf_module_v1_module_proto_depIdxs = []int32{
	1, // 0: buf.module.v1.Module.files:type_name -> buf.module.v1.ModuleFile
	2, // 1: buf.module.v1.Module.dependencies:type_name -> buf.module.v1.ModuleName
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_module_v1_module_proto_init() }
func file_buf_module_v1_module_proto_init() {
	if File_buf_module_v1_module_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_module_v1_module_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Module); i {
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
		file_buf_module_v1_module_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleFile); i {
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
		file_buf_module_v1_module_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModuleName); i {
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
			RawDescriptor: file_buf_module_v1_module_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_module_v1_module_proto_goTypes,
		DependencyIndexes: file_buf_module_v1_module_proto_depIdxs,
		MessageInfos:      file_buf_module_v1_module_proto_msgTypes,
	}.Build()
	File_buf_module_v1_module_proto = out.File
	file_buf_module_v1_module_proto_rawDesc = nil
	file_buf_module_v1_module_proto_goTypes = nil
	file_buf_module_v1_module_proto_depIdxs = nil
}
