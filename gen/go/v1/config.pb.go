// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/config.proto

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

// Config is the top level config object for restic UI.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// modification number, used for read-modify-write consistency in the UI. Incremented on every write.
	Modno int32 `protobuf:"varint,1,opt,name=modno,proto3" json:"modno,omitempty"`
	// override the hostname tagged on backups. If provided it will be used in addition to tags to group backups.
	Host  string  `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	Repos []*Repo `protobuf:"bytes,3,rep,name=repos,proto3" json:"repos,omitempty"`
	Plans []*Plan `protobuf:"bytes,4,rep,name=plans,proto3" json:"plans,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetModno() int32 {
	if x != nil {
		return x.Modno
	}
	return 0
}

func (x *Config) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Config) GetRepos() []*Repo {
	if x != nil {
		return x.Repos
	}
	return nil
}

func (x *Config) GetPlans() []*Plan {
	if x != nil {
		return x.Plans
	}
	return nil
}

type Repo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                                      // unique but human readable ID for this repo.
	Uri         string       `protobuf:"bytes,2,opt,name=uri,proto3" json:"uri,omitempty"`                                    // restic repo URI
	Password    string       `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                          // plaintext password
	Env         []string     `protobuf:"bytes,4,rep,name=env,proto3" json:"env,omitempty"`                                    // extra environment variables to set for restic.
	Flags       []string     `protobuf:"bytes,5,rep,name=flags,proto3" json:"flags,omitempty"`                                // extra flags set on the restic command.
	PrunePolicy *PrunePolicy `protobuf:"bytes,6,opt,name=prune_policy,json=prunePolicy,proto3" json:"prune_policy,omitempty"` // policy for when to run prune.
}

func (x *Repo) Reset() {
	*x = Repo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repo) ProtoMessage() {}

func (x *Repo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repo.ProtoReflect.Descriptor instead.
func (*Repo) Descriptor() ([]byte, []int) {
	return file_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *Repo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Repo) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

func (x *Repo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Repo) GetEnv() []string {
	if x != nil {
		return x.Env
	}
	return nil
}

func (x *Repo) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

func (x *Repo) GetPrunePolicy() *PrunePolicy {
	if x != nil {
		return x.PrunePolicy
	}
	return nil
}

type Plan struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`               // unique but human readable ID for this plan.
	Repo      string           `protobuf:"bytes,2,opt,name=repo,proto3" json:"repo,omitempty"`           // ID of the repo to use.
	Paths     []string         `protobuf:"bytes,4,rep,name=paths,proto3" json:"paths,omitempty"`         // paths to include in the backup.
	Excludes  []string         `protobuf:"bytes,5,rep,name=excludes,proto3" json:"excludes,omitempty"`   // glob patterns to exclude.
	Cron      string           `protobuf:"bytes,6,opt,name=cron,proto3" json:"cron,omitempty"`           // cron expression describing the backup schedule.
	Retention *RetentionPolicy `protobuf:"bytes,7,opt,name=retention,proto3" json:"retention,omitempty"` // retention policy for snapshots.
}

func (x *Plan) Reset() {
	*x = Plan{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plan) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plan) ProtoMessage() {}

func (x *Plan) ProtoReflect() protoreflect.Message {
	mi := &file_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plan.ProtoReflect.Descriptor instead.
func (*Plan) Descriptor() ([]byte, []int) {
	return file_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *Plan) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Plan) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *Plan) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *Plan) GetExcludes() []string {
	if x != nil {
		return x.Excludes
	}
	return nil
}

func (x *Plan) GetCron() string {
	if x != nil {
		return x.Cron
	}
	return ""
}

func (x *Plan) GetRetention() *RetentionPolicy {
	if x != nil {
		return x.Retention
	}
	return nil
}

type RetentionPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// max_unused_limit is used to decide when forget should be run.
	MaxUnusedLimit     string `protobuf:"bytes,1,opt,name=max_unused_limit,json=maxUnusedLimit,proto3" json:"max_unused_limit,omitempty"`             // e.g. a percentage i.e. 25% or a number of megabytes.
	KeepLastN          int32  `protobuf:"varint,2,opt,name=keep_last_n,json=keepLastN,proto3" json:"keep_last_n,omitempty"`                           // keep the last n snapshots.
	KeepHourly         int32  `protobuf:"varint,3,opt,name=keep_hourly,json=keepHourly,proto3" json:"keep_hourly,omitempty"`                          // keep the last n hourly snapshots.
	KeepDaily          int32  `protobuf:"varint,4,opt,name=keep_daily,json=keepDaily,proto3" json:"keep_daily,omitempty"`                             // keep the last n daily snapshots.
	KeepWeekly         int32  `protobuf:"varint,5,opt,name=keep_weekly,json=keepWeekly,proto3" json:"keep_weekly,omitempty"`                          // keep the last n weekly snapshots.
	KeepMonthly        int32  `protobuf:"varint,6,opt,name=keep_monthly,json=keepMonthly,proto3" json:"keep_monthly,omitempty"`                       // keep the last n monthly snapshots.
	KeepYearly         int32  `protobuf:"varint,7,opt,name=keep_yearly,json=keepYearly,proto3" json:"keep_yearly,omitempty"`                          // keep the last n yearly snapshots.
	KeepWithinDuration string `protobuf:"bytes,8,opt,name=keep_within_duration,json=keepWithinDuration,proto3" json:"keep_within_duration,omitempty"` // keep snapshots within a duration e.g. 1y2m3d4h5m6s
}

func (x *RetentionPolicy) Reset() {
	*x = RetentionPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetentionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetentionPolicy) ProtoMessage() {}

func (x *RetentionPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetentionPolicy.ProtoReflect.Descriptor instead.
func (*RetentionPolicy) Descriptor() ([]byte, []int) {
	return file_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *RetentionPolicy) GetMaxUnusedLimit() string {
	if x != nil {
		return x.MaxUnusedLimit
	}
	return ""
}

func (x *RetentionPolicy) GetKeepLastN() int32 {
	if x != nil {
		return x.KeepLastN
	}
	return 0
}

func (x *RetentionPolicy) GetKeepHourly() int32 {
	if x != nil {
		return x.KeepHourly
	}
	return 0
}

func (x *RetentionPolicy) GetKeepDaily() int32 {
	if x != nil {
		return x.KeepDaily
	}
	return 0
}

func (x *RetentionPolicy) GetKeepWeekly() int32 {
	if x != nil {
		return x.KeepWeekly
	}
	return 0
}

func (x *RetentionPolicy) GetKeepMonthly() int32 {
	if x != nil {
		return x.KeepMonthly
	}
	return 0
}

func (x *RetentionPolicy) GetKeepYearly() int32 {
	if x != nil {
		return x.KeepYearly
	}
	return 0
}

func (x *RetentionPolicy) GetKeepWithinDuration() string {
	if x != nil {
		return x.KeepWithinDuration
	}
	return ""
}

type PrunePolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxFrequencyDays int32 `protobuf:"varint,1,opt,name=max_frequency_days,json=maxFrequencyDays,proto3" json:"max_frequency_days,omitempty"`   // max frequency of prune runs in days. If 0, prune will be run on every backup.
	MaxUnusedPercent int32 `protobuf:"varint,100,opt,name=max_unused_percent,json=maxUnusedPercent,proto3" json:"max_unused_percent,omitempty"` // max percentage of repo size that can be unused before prune is run.
	MaxUnusedBytes   int32 `protobuf:"varint,101,opt,name=max_unused_bytes,json=maxUnusedBytes,proto3" json:"max_unused_bytes,omitempty"`       // max number of bytes that can be unused before prune is run.
}

func (x *PrunePolicy) Reset() {
	*x = PrunePolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrunePolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrunePolicy) ProtoMessage() {}

func (x *PrunePolicy) ProtoReflect() protoreflect.Message {
	mi := &file_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrunePolicy.ProtoReflect.Descriptor instead.
func (*PrunePolicy) Descriptor() ([]byte, []int) {
	return file_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *PrunePolicy) GetMaxFrequencyDays() int32 {
	if x != nil {
		return x.MaxFrequencyDays
	}
	return 0
}

func (x *PrunePolicy) GetMaxUnusedPercent() int32 {
	if x != nil {
		return x.MaxUnusedPercent
	}
	return 0
}

func (x *PrunePolicy) GetMaxUnusedBytes() int32 {
	if x != nil {
		return x.MaxUnusedBytes
	}
	return 0
}

var File_v1_config_proto protoreflect.FileDescriptor

var file_v1_config_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0x72, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6d, 0x6f, 0x64, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x70, 0x6f, 0x52, 0x05, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x1e, 0x0a, 0x05, 0x70, 0x6c, 0x61,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c,
	0x61, 0x6e, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x04, 0x52, 0x65,
	0x70, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x69, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x65,
	0x6e, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x32, 0x0a, 0x0c, 0x70, 0x72, 0x75, 0x6e,
	0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x0b, 0x70, 0x72, 0x75, 0x6e, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0xa3, 0x01, 0x0a,
	0x04, 0x50, 0x6c, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x68, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x72, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x72, 0x6f, 0x6e, 0x12,
	0x31, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x09, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xb2, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e,
	0x75, 0x73, 0x65, 0x64, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x1e, 0x0a, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x4c, 0x61, 0x73, 0x74, 0x4e,
	0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x6c, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x48, 0x6f, 0x75, 0x72, 0x6c,
	0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x44, 0x61, 0x69, 0x6c, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x77, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x57, 0x65, 0x65, 0x6b, 0x6c,
	0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x79, 0x65, 0x61,
	0x72, 0x6c, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x59,
	0x65, 0x61, 0x72, 0x6c, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x6b, 0x65, 0x65, 0x70, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x50, 0x72, 0x75, 0x6e,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x66,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x46, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x79, 0x44, 0x61, 0x79, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x75,
	0x73, 0x65, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x50, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x75, 0x73, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x65, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d,
	0x61, 0x78, 0x55, 0x6e, 0x75, 0x73, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73, 0x42, 0x2e, 0x5a,
	0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x65,
	0x74, 0x68, 0x67, 0x65, 0x6f, 0x72, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x69, 0x63, 0x75,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_config_proto_rawDescOnce sync.Once
	file_v1_config_proto_rawDescData = file_v1_config_proto_rawDesc
)

func file_v1_config_proto_rawDescGZIP() []byte {
	file_v1_config_proto_rawDescOnce.Do(func() {
		file_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_config_proto_rawDescData)
	})
	return file_v1_config_proto_rawDescData
}

var file_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_v1_config_proto_goTypes = []interface{}{
	(*Config)(nil),          // 0: v1.Config
	(*Repo)(nil),            // 1: v1.Repo
	(*Plan)(nil),            // 2: v1.Plan
	(*RetentionPolicy)(nil), // 3: v1.RetentionPolicy
	(*PrunePolicy)(nil),     // 4: v1.PrunePolicy
}
var file_v1_config_proto_depIdxs = []int32{
	1, // 0: v1.Config.repos:type_name -> v1.Repo
	2, // 1: v1.Config.plans:type_name -> v1.Plan
	4, // 2: v1.Repo.prune_policy:type_name -> v1.PrunePolicy
	3, // 3: v1.Plan.retention:type_name -> v1.RetentionPolicy
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_config_proto_init() }
func file_v1_config_proto_init() {
	if File_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
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
		file_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repo); i {
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
		file_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plan); i {
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
		file_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetentionPolicy); i {
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
		file_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrunePolicy); i {
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
			RawDescriptor: file_v1_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_config_proto_goTypes,
		DependencyIndexes: file_v1_config_proto_depIdxs,
		MessageInfos:      file_v1_config_proto_msgTypes,
	}.Build()
	File_v1_config_proto = out.File
	file_v1_config_proto_rawDesc = nil
	file_v1_config_proto_goTypes = nil
	file_v1_config_proto_depIdxs = nil
}
