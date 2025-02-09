// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: v1/restic.proto

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

// ResticSnapshot represents a restic snapshot.
type ResticSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UnixTimeMs int64    `protobuf:"varint,2,opt,name=unix_time_ms,json=unixTimeMs,proto3" json:"unix_time_ms,omitempty"`
	Hostname   string   `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Username   string   `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Tree       string   `protobuf:"bytes,5,opt,name=tree,proto3" json:"tree,omitempty"`     // tree hash
	Parent     string   `protobuf:"bytes,6,opt,name=parent,proto3" json:"parent,omitempty"` // parent snapshot's id
	Paths      []string `protobuf:"bytes,7,rep,name=paths,proto3" json:"paths,omitempty"`
	Tags       []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ResticSnapshot) Reset() {
	*x = ResticSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_restic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResticSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResticSnapshot) ProtoMessage() {}

func (x *ResticSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_v1_restic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResticSnapshot.ProtoReflect.Descriptor instead.
func (*ResticSnapshot) Descriptor() ([]byte, []int) {
	return file_v1_restic_proto_rawDescGZIP(), []int{0}
}

func (x *ResticSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResticSnapshot) GetUnixTimeMs() int64 {
	if x != nil {
		return x.UnixTimeMs
	}
	return 0
}

func (x *ResticSnapshot) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ResticSnapshot) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ResticSnapshot) GetTree() string {
	if x != nil {
		return x.Tree
	}
	return ""
}

func (x *ResticSnapshot) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ResticSnapshot) GetPaths() []string {
	if x != nil {
		return x.Paths
	}
	return nil
}

func (x *ResticSnapshot) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

// ResticSnapshotList represents a list of restic snapshots.
type ResticSnapshotList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snapshots []*ResticSnapshot `protobuf:"bytes,1,rep,name=snapshots,proto3" json:"snapshots,omitempty"`
}

func (x *ResticSnapshotList) Reset() {
	*x = ResticSnapshotList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_restic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResticSnapshotList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResticSnapshotList) ProtoMessage() {}

func (x *ResticSnapshotList) ProtoReflect() protoreflect.Message {
	mi := &file_v1_restic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResticSnapshotList.ProtoReflect.Descriptor instead.
func (*ResticSnapshotList) Descriptor() ([]byte, []int) {
	return file_v1_restic_proto_rawDescGZIP(), []int{1}
}

func (x *ResticSnapshotList) GetSnapshots() []*ResticSnapshot {
	if x != nil {
		return x.Snapshots
	}
	return nil
}

// BackupProgressEntriy represents a single entry in the backup progress stream.
type BackupProgressEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Entry:
	//
	//	*BackupProgressEntry_Status
	//	*BackupProgressEntry_Summary
	Entry isBackupProgressEntry_Entry `protobuf_oneof:"entry"`
}

func (x *BackupProgressEntry) Reset() {
	*x = BackupProgressEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_restic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupProgressEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupProgressEntry) ProtoMessage() {}

func (x *BackupProgressEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_restic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupProgressEntry.ProtoReflect.Descriptor instead.
func (*BackupProgressEntry) Descriptor() ([]byte, []int) {
	return file_v1_restic_proto_rawDescGZIP(), []int{2}
}

func (m *BackupProgressEntry) GetEntry() isBackupProgressEntry_Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (x *BackupProgressEntry) GetStatus() *BackupProgressStatusEntry {
	if x, ok := x.GetEntry().(*BackupProgressEntry_Status); ok {
		return x.Status
	}
	return nil
}

func (x *BackupProgressEntry) GetSummary() *BackupProgressSummary {
	if x, ok := x.GetEntry().(*BackupProgressEntry_Summary); ok {
		return x.Summary
	}
	return nil
}

type isBackupProgressEntry_Entry interface {
	isBackupProgressEntry_Entry()
}

type BackupProgressEntry_Status struct {
	Status *BackupProgressStatusEntry `protobuf:"bytes,1,opt,name=status,proto3,oneof"`
}

type BackupProgressEntry_Summary struct {
	Summary *BackupProgressSummary `protobuf:"bytes,2,opt,name=summary,proto3,oneof"`
}

func (*BackupProgressEntry_Status) isBackupProgressEntry_Entry() {}

func (*BackupProgressEntry_Summary) isBackupProgressEntry_Entry() {}

// BackupProgressStatusEntry represents a single status entry in the backup progress stream.
type BackupProgressStatusEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PercentDone float64  `protobuf:"fixed64,1,opt,name=percent_done,json=percentDone,proto3" json:"percent_done,omitempty"` // 0.0 - 1.0
	TotalFiles  int64    `protobuf:"varint,2,opt,name=total_files,json=totalFiles,proto3" json:"total_files,omitempty"`
	TotalBytes  int64    `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	FilesDone   int64    `protobuf:"varint,4,opt,name=files_done,json=filesDone,proto3" json:"files_done,omitempty"`
	BytesDone   int64    `protobuf:"varint,5,opt,name=bytes_done,json=bytesDone,proto3" json:"bytes_done,omitempty"`
	CurrentFile []string `protobuf:"bytes,6,rep,name=current_file,json=currentFile,proto3" json:"current_file,omitempty"`
}

func (x *BackupProgressStatusEntry) Reset() {
	*x = BackupProgressStatusEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_restic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupProgressStatusEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupProgressStatusEntry) ProtoMessage() {}

func (x *BackupProgressStatusEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_restic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupProgressStatusEntry.ProtoReflect.Descriptor instead.
func (*BackupProgressStatusEntry) Descriptor() ([]byte, []int) {
	return file_v1_restic_proto_rawDescGZIP(), []int{3}
}

func (x *BackupProgressStatusEntry) GetPercentDone() float64 {
	if x != nil {
		return x.PercentDone
	}
	return 0
}

func (x *BackupProgressStatusEntry) GetTotalFiles() int64 {
	if x != nil {
		return x.TotalFiles
	}
	return 0
}

func (x *BackupProgressStatusEntry) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *BackupProgressStatusEntry) GetFilesDone() int64 {
	if x != nil {
		return x.FilesDone
	}
	return 0
}

func (x *BackupProgressStatusEntry) GetBytesDone() int64 {
	if x != nil {
		return x.BytesDone
	}
	return 0
}

func (x *BackupProgressStatusEntry) GetCurrentFile() []string {
	if x != nil {
		return x.CurrentFile
	}
	return nil
}

// BackupProgressSummary represents a the summary event emitted at the end of a backup stream.
type BackupProgressSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FilesNew            int64   `protobuf:"varint,1,opt,name=files_new,json=filesNew,proto3" json:"files_new,omitempty"`
	FilesChanged        int64   `protobuf:"varint,2,opt,name=files_changed,json=filesChanged,proto3" json:"files_changed,omitempty"`
	FilesUnmodified     int64   `protobuf:"varint,3,opt,name=files_unmodified,json=filesUnmodified,proto3" json:"files_unmodified,omitempty"`
	DirsNew             int64   `protobuf:"varint,4,opt,name=dirs_new,json=dirsNew,proto3" json:"dirs_new,omitempty"`
	DirsChanged         int64   `protobuf:"varint,5,opt,name=dirs_changed,json=dirsChanged,proto3" json:"dirs_changed,omitempty"`
	DirsUnmodified      int64   `protobuf:"varint,6,opt,name=dirs_unmodified,json=dirsUnmodified,proto3" json:"dirs_unmodified,omitempty"`
	DataBlobs           int64   `protobuf:"varint,7,opt,name=data_blobs,json=dataBlobs,proto3" json:"data_blobs,omitempty"`
	TreeBlobs           int64   `protobuf:"varint,8,opt,name=tree_blobs,json=treeBlobs,proto3" json:"tree_blobs,omitempty"`
	DataAdded           int64   `protobuf:"varint,9,opt,name=data_added,json=dataAdded,proto3" json:"data_added,omitempty"`
	TotalFilesProcessed int64   `protobuf:"varint,10,opt,name=total_files_processed,json=totalFilesProcessed,proto3" json:"total_files_processed,omitempty"`
	TotalBytesProcessed int64   `protobuf:"varint,11,opt,name=total_bytes_processed,json=totalBytesProcessed,proto3" json:"total_bytes_processed,omitempty"`
	TotalDuration       float64 `protobuf:"fixed64,12,opt,name=total_duration,json=totalDuration,proto3" json:"total_duration,omitempty"`
	SnapshotId          string  `protobuf:"bytes,13,opt,name=snapshot_id,json=snapshotId,proto3" json:"snapshot_id,omitempty"`
}

func (x *BackupProgressSummary) Reset() {
	*x = BackupProgressSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_restic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BackupProgressSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupProgressSummary) ProtoMessage() {}

func (x *BackupProgressSummary) ProtoReflect() protoreflect.Message {
	mi := &file_v1_restic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupProgressSummary.ProtoReflect.Descriptor instead.
func (*BackupProgressSummary) Descriptor() ([]byte, []int) {
	return file_v1_restic_proto_rawDescGZIP(), []int{4}
}

func (x *BackupProgressSummary) GetFilesNew() int64 {
	if x != nil {
		return x.FilesNew
	}
	return 0
}

func (x *BackupProgressSummary) GetFilesChanged() int64 {
	if x != nil {
		return x.FilesChanged
	}
	return 0
}

func (x *BackupProgressSummary) GetFilesUnmodified() int64 {
	if x != nil {
		return x.FilesUnmodified
	}
	return 0
}

func (x *BackupProgressSummary) GetDirsNew() int64 {
	if x != nil {
		return x.DirsNew
	}
	return 0
}

func (x *BackupProgressSummary) GetDirsChanged() int64 {
	if x != nil {
		return x.DirsChanged
	}
	return 0
}

func (x *BackupProgressSummary) GetDirsUnmodified() int64 {
	if x != nil {
		return x.DirsUnmodified
	}
	return 0
}

func (x *BackupProgressSummary) GetDataBlobs() int64 {
	if x != nil {
		return x.DataBlobs
	}
	return 0
}

func (x *BackupProgressSummary) GetTreeBlobs() int64 {
	if x != nil {
		return x.TreeBlobs
	}
	return 0
}

func (x *BackupProgressSummary) GetDataAdded() int64 {
	if x != nil {
		return x.DataAdded
	}
	return 0
}

func (x *BackupProgressSummary) GetTotalFilesProcessed() int64 {
	if x != nil {
		return x.TotalFilesProcessed
	}
	return 0
}

func (x *BackupProgressSummary) GetTotalBytesProcessed() int64 {
	if x != nil {
		return x.TotalBytesProcessed
	}
	return 0
}

func (x *BackupProgressSummary) GetTotalDuration() float64 {
	if x != nil {
		return x.TotalDuration
	}
	return 0
}

func (x *BackupProgressSummary) GetSnapshotId() string {
	if x != nil {
		return x.SnapshotId
	}
	return ""
}

// RestoreProgressEvent represents a single entry in the restore progress stream.
type RestoreProgressEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType    string  `protobuf:"bytes,1,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"` // "summary" or "status"
	SecondsElapsed float64 `protobuf:"fixed64,2,opt,name=seconds_elapsed,json=secondsElapsed,proto3" json:"seconds_elapsed,omitempty"`
	TotalBytes     int64   `protobuf:"varint,3,opt,name=total_bytes,json=totalBytes,proto3" json:"total_bytes,omitempty"`
	BytesRestored  int64   `protobuf:"varint,4,opt,name=bytes_restored,json=bytesRestored,proto3" json:"bytes_restored,omitempty"`
	TotalFiles     int64   `protobuf:"varint,5,opt,name=total_files,json=totalFiles,proto3" json:"total_files,omitempty"`
	FilesRestored  int64   `protobuf:"varint,6,opt,name=files_restored,json=filesRestored,proto3" json:"files_restored,omitempty"`
	PercentDone    float64 `protobuf:"fixed64,7,opt,name=percent_done,json=percentDone,proto3" json:"percent_done,omitempty"` // 0.0 - 1.0
}

func (x *RestoreProgressEntry) Reset() {
	*x = RestoreProgressEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_restic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestoreProgressEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestoreProgressEntry) ProtoMessage() {}

func (x *RestoreProgressEntry) ProtoReflect() protoreflect.Message {
	mi := &file_v1_restic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestoreProgressEntry.ProtoReflect.Descriptor instead.
func (*RestoreProgressEntry) Descriptor() ([]byte, []int) {
	return file_v1_restic_proto_rawDescGZIP(), []int{5}
}

func (x *RestoreProgressEntry) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *RestoreProgressEntry) GetSecondsElapsed() float64 {
	if x != nil {
		return x.SecondsElapsed
	}
	return 0
}

func (x *RestoreProgressEntry) GetTotalBytes() int64 {
	if x != nil {
		return x.TotalBytes
	}
	return 0
}

func (x *RestoreProgressEntry) GetBytesRestored() int64 {
	if x != nil {
		return x.BytesRestored
	}
	return 0
}

func (x *RestoreProgressEntry) GetTotalFiles() int64 {
	if x != nil {
		return x.TotalFiles
	}
	return 0
}

func (x *RestoreProgressEntry) GetFilesRestored() int64 {
	if x != nil {
		return x.FilesRestored
	}
	return 0
}

func (x *RestoreProgressEntry) GetPercentDone() float64 {
	if x != nil {
		return x.PercentDone
	}
	return 0
}

var File_v1_restic_proto protoreflect.FileDescriptor

var file_v1_restic_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xd0, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74, 0x69, 0x63,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x75, 0x6e, 0x69, 0x78,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x75, 0x6e, 0x69, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x72, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x72, 0x65, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x68, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x46, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x74,
	0x69, 0x63, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30,
	0x0a, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x69, 0x63, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73,
	0x22, 0x8e, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x35, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x72,
	0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x48, 0x00, 0x52,
	0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0xe1, 0x01, 0x0a, 0x19, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x44, 0x6f,
	0x6e, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x64, 0x6f,
	0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x44,
	0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x64, 0x6f, 0x6e,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x44, 0x6f,
	0x6e, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69,
	0x6c, 0x65, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x22, 0xf8, 0x03, 0x0a, 0x15, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70,
	0x50, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x23, 0x0a, 0x0d,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x75, 0x6e, 0x6d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x55, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x64, 0x69, 0x72, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x64, 0x69, 0x72, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x73, 0x5f,
	0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64,
	0x69, 0x72, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69,
	0x72, 0x73, 0x5f, 0x75, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x64, 0x69, 0x72, 0x73, 0x55, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x62, 0x6c, 0x6f, 0x62,
	0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x42, 0x6c, 0x6f,
	0x62, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x72, 0x65, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x62, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x72, 0x65, 0x65, 0x42, 0x6c, 0x6f, 0x62,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x61, 0x74, 0x61, 0x41, 0x64, 0x64, 0x65, 0x64,
	0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x15, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x49, 0x64,
	0x22, 0x95, 0x02, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x5f, 0x65, 0x6c, 0x61, 0x70, 0x73, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0e, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x45, 0x6c,
	0x61, 0x70, 0x73, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62,
	0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x25,
	0x0a, 0x0e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f, 0x72, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x5f, 0x64, 0x6f, 0x6e, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x44, 0x6f, 0x6e, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x72, 0x65, 0x74, 0x68, 0x67, 0x65, 0x6f,
	0x72, 0x67, 0x65, 0x2f, 0x72, 0x65, 0x73, 0x74, 0x69, 0x63, 0x75, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_restic_proto_rawDescOnce sync.Once
	file_v1_restic_proto_rawDescData = file_v1_restic_proto_rawDesc
)

func file_v1_restic_proto_rawDescGZIP() []byte {
	file_v1_restic_proto_rawDescOnce.Do(func() {
		file_v1_restic_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_restic_proto_rawDescData)
	})
	return file_v1_restic_proto_rawDescData
}

var file_v1_restic_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_v1_restic_proto_goTypes = []interface{}{
	(*ResticSnapshot)(nil),            // 0: v1.ResticSnapshot
	(*ResticSnapshotList)(nil),        // 1: v1.ResticSnapshotList
	(*BackupProgressEntry)(nil),       // 2: v1.BackupProgressEntry
	(*BackupProgressStatusEntry)(nil), // 3: v1.BackupProgressStatusEntry
	(*BackupProgressSummary)(nil),     // 4: v1.BackupProgressSummary
	(*RestoreProgressEntry)(nil),      // 5: v1.RestoreProgressEntry
}
var file_v1_restic_proto_depIdxs = []int32{
	0, // 0: v1.ResticSnapshotList.snapshots:type_name -> v1.ResticSnapshot
	3, // 1: v1.BackupProgressEntry.status:type_name -> v1.BackupProgressStatusEntry
	4, // 2: v1.BackupProgressEntry.summary:type_name -> v1.BackupProgressSummary
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_v1_restic_proto_init() }
func file_v1_restic_proto_init() {
	if File_v1_restic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_restic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResticSnapshot); i {
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
		file_v1_restic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResticSnapshotList); i {
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
		file_v1_restic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupProgressEntry); i {
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
		file_v1_restic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupProgressStatusEntry); i {
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
		file_v1_restic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BackupProgressSummary); i {
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
		file_v1_restic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestoreProgressEntry); i {
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
	file_v1_restic_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*BackupProgressEntry_Status)(nil),
		(*BackupProgressEntry_Summary)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_restic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_restic_proto_goTypes,
		DependencyIndexes: file_v1_restic_proto_depIdxs,
		MessageInfos:      file_v1_restic_proto_msgTypes,
	}.Build()
	File_v1_restic_proto = out.File
	file_v1_restic_proto_rawDesc = nil
	file_v1_restic_proto_goTypes = nil
	file_v1_restic_proto_depIdxs = nil
}
