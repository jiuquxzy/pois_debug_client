// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pois_api.proto

package pois_rpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ResponseMinerRegister struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMinerRegister) Reset()         { *m = ResponseMinerRegister{} }
func (m *ResponseMinerRegister) String() string { return proto.CompactTextString(m) }
func (*ResponseMinerRegister) ProtoMessage()    {}
func (*ResponseMinerRegister) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{0}
}

func (m *ResponseMinerRegister) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMinerRegister.Unmarshal(m, b)
}
func (m *ResponseMinerRegister) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMinerRegister.Marshal(b, m, deterministic)
}
func (m *ResponseMinerRegister) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMinerRegister.Merge(m, src)
}
func (m *ResponseMinerRegister) XXX_Size() int {
	return xxx_messageInfo_ResponseMinerRegister.Size(m)
}
func (m *ResponseMinerRegister) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMinerRegister.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMinerRegister proto.InternalMessageInfo

func (m *ResponseMinerRegister) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type RequestMinerInitParam struct {
	MinerId              []byte   `protobuf:"bytes,1,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMinerInitParam) Reset()         { *m = RequestMinerInitParam{} }
func (m *RequestMinerInitParam) String() string { return proto.CompactTextString(m) }
func (*RequestMinerInitParam) ProtoMessage()    {}
func (*RequestMinerInitParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{1}
}

func (m *RequestMinerInitParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMinerInitParam.Unmarshal(m, b)
}
func (m *RequestMinerInitParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMinerInitParam.Marshal(b, m, deterministic)
}
func (m *RequestMinerInitParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMinerInitParam.Merge(m, src)
}
func (m *RequestMinerInitParam) XXX_Size() int {
	return xxx_messageInfo_RequestMinerInitParam.Size(m)
}
func (m *RequestMinerInitParam) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMinerInitParam.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMinerInitParam proto.InternalMessageInfo

func (m *RequestMinerInitParam) GetMinerId() []byte {
	if m != nil {
		return m.MinerId
	}
	return nil
}

type ResponseMinerInitParam struct {
	Acc                  []byte   `protobuf:"bytes,1,opt,name=acc,proto3" json:"acc,omitempty"`
	KeyN                 []byte   `protobuf:"bytes,2,opt,name=key_n,json=keyN,proto3" json:"key_n,omitempty"`
	KeyG                 []byte   `protobuf:"bytes,3,opt,name=key_g,json=keyG,proto3" json:"key_g,omitempty"`
	Front                int64    `protobuf:"varint,4,opt,name=front,proto3" json:"front,omitempty"`
	Rear                 int64    `protobuf:"varint,5,opt,name=rear,proto3" json:"rear,omitempty"`
	MinerId              []byte   `protobuf:"bytes,6,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
	Signature            []byte   `protobuf:"bytes,7,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseMinerInitParam) Reset()         { *m = ResponseMinerInitParam{} }
func (m *ResponseMinerInitParam) String() string { return proto.CompactTextString(m) }
func (*ResponseMinerInitParam) ProtoMessage()    {}
func (*ResponseMinerInitParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{2}
}

func (m *ResponseMinerInitParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseMinerInitParam.Unmarshal(m, b)
}
func (m *ResponseMinerInitParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseMinerInitParam.Marshal(b, m, deterministic)
}
func (m *ResponseMinerInitParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseMinerInitParam.Merge(m, src)
}
func (m *ResponseMinerInitParam) XXX_Size() int {
	return xxx_messageInfo_ResponseMinerInitParam.Size(m)
}
func (m *ResponseMinerInitParam) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseMinerInitParam.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseMinerInitParam proto.InternalMessageInfo

func (m *ResponseMinerInitParam) GetAcc() []byte {
	if m != nil {
		return m.Acc
	}
	return nil
}

func (m *ResponseMinerInitParam) GetKeyN() []byte {
	if m != nil {
		return m.KeyN
	}
	return nil
}

func (m *ResponseMinerInitParam) GetKeyG() []byte {
	if m != nil {
		return m.KeyG
	}
	return nil
}

func (m *ResponseMinerInitParam) GetFront() int64 {
	if m != nil {
		return m.Front
	}
	return 0
}

func (m *ResponseMinerInitParam) GetRear() int64 {
	if m != nil {
		return m.Rear
	}
	return 0
}

func (m *ResponseMinerInitParam) GetMinerId() []byte {
	if m != nil {
		return m.MinerId
	}
	return nil
}

func (m *ResponseMinerInitParam) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type RequestMinerCommitGenChall struct {
	MinerId []byte    `protobuf:"bytes,1,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
	Commit  []*Commit `protobuf:"bytes,2,rep,name=commit,proto3" json:"commit,omitempty"`
	// This is acc data to set the ProverNode
	ProverAcc            *ProverAcc `protobuf:"bytes,3,opt,name=prover_acc,json=proverAcc,proto3" json:"prover_acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RequestMinerCommitGenChall) Reset()         { *m = RequestMinerCommitGenChall{} }
func (m *RequestMinerCommitGenChall) String() string { return proto.CompactTextString(m) }
func (*RequestMinerCommitGenChall) ProtoMessage()    {}
func (*RequestMinerCommitGenChall) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{3}
}

func (m *RequestMinerCommitGenChall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMinerCommitGenChall.Unmarshal(m, b)
}
func (m *RequestMinerCommitGenChall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMinerCommitGenChall.Marshal(b, m, deterministic)
}
func (m *RequestMinerCommitGenChall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMinerCommitGenChall.Merge(m, src)
}
func (m *RequestMinerCommitGenChall) XXX_Size() int {
	return xxx_messageInfo_RequestMinerCommitGenChall.Size(m)
}
func (m *RequestMinerCommitGenChall) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMinerCommitGenChall.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMinerCommitGenChall proto.InternalMessageInfo

func (m *RequestMinerCommitGenChall) GetMinerId() []byte {
	if m != nil {
		return m.MinerId
	}
	return nil
}

func (m *RequestMinerCommitGenChall) GetCommit() []*Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *RequestMinerCommitGenChall) GetProverAcc() *ProverAcc {
	if m != nil {
		return m.ProverAcc
	}
	return nil
}

type ProverAcc struct {
	Acc                  []byte   `protobuf:"bytes,1,opt,name=acc,proto3" json:"acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProverAcc) Reset()         { *m = ProverAcc{} }
func (m *ProverAcc) String() string { return proto.CompactTextString(m) }
func (*ProverAcc) ProtoMessage()    {}
func (*ProverAcc) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{4}
}

func (m *ProverAcc) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProverAcc.Unmarshal(m, b)
}
func (m *ProverAcc) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProverAcc.Marshal(b, m, deterministic)
}
func (m *ProverAcc) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProverAcc.Merge(m, src)
}
func (m *ProverAcc) XXX_Size() int {
	return xxx_messageInfo_ProverAcc.Size(m)
}
func (m *ProverAcc) XXX_DiscardUnknown() {
	xxx_messageInfo_ProverAcc.DiscardUnknown(m)
}

var xxx_messageInfo_ProverAcc proto.InternalMessageInfo

func (m *ProverAcc) GetAcc() []byte {
	if m != nil {
		return m.Acc
	}
	return nil
}

type Commit struct {
	FileIndex            int64    `protobuf:"varint,1,opt,name=file_index,json=fileIndex,proto3" json:"file_index,omitempty"`
	Roots                [][]byte `protobuf:"bytes,2,rep,name=roots,proto3" json:"roots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Commit) Reset()         { *m = Commit{} }
func (m *Commit) String() string { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()    {}
func (*Commit) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{5}
}

func (m *Commit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Commit.Unmarshal(m, b)
}
func (m *Commit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Commit.Marshal(b, m, deterministic)
}
func (m *Commit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Commit.Merge(m, src)
}
func (m *Commit) XXX_Size() int {
	return xxx_messageInfo_Commit.Size(m)
}
func (m *Commit) XXX_DiscardUnknown() {
	xxx_messageInfo_Commit.DiscardUnknown(m)
}

var xxx_messageInfo_Commit proto.InternalMessageInfo

func (m *Commit) GetFileIndex() int64 {
	if m != nil {
		return m.FileIndex
	}
	return 0
}

func (m *Commit) GetRoots() [][]byte {
	if m != nil {
		return m.Roots
	}
	return nil
}

type Challenge struct {
	Rows                 []*Int64Slice `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Challenge) Reset()         { *m = Challenge{} }
func (m *Challenge) String() string { return proto.CompactTextString(m) }
func (*Challenge) ProtoMessage()    {}
func (*Challenge) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{6}
}

func (m *Challenge) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Challenge.Unmarshal(m, b)
}
func (m *Challenge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Challenge.Marshal(b, m, deterministic)
}
func (m *Challenge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Challenge.Merge(m, src)
}
func (m *Challenge) XXX_Size() int {
	return xxx_messageInfo_Challenge.Size(m)
}
func (m *Challenge) XXX_DiscardUnknown() {
	xxx_messageInfo_Challenge.DiscardUnknown(m)
}

var xxx_messageInfo_Challenge proto.InternalMessageInfo

func (m *Challenge) GetRows() []*Int64Slice {
	if m != nil {
		return m.Rows
	}
	return nil
}

type Int64Slice struct {
	Values               []int64  `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Slice) Reset()         { *m = Int64Slice{} }
func (m *Int64Slice) String() string { return proto.CompactTextString(m) }
func (*Int64Slice) ProtoMessage()    {}
func (*Int64Slice) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{7}
}

func (m *Int64Slice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Slice.Unmarshal(m, b)
}
func (m *Int64Slice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Slice.Marshal(b, m, deterministic)
}
func (m *Int64Slice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Slice.Merge(m, src)
}
func (m *Int64Slice) XXX_Size() int {
	return xxx_messageInfo_Int64Slice.Size(m)
}
func (m *Int64Slice) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Slice.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Slice proto.InternalMessageInfo

func (m *Int64Slice) GetValues() []int64 {
	if m != nil {
		return m.Values
	}
	return nil
}

type RequestVerifyCommitAndAccProof struct {
	CommitProofGroup     *CommitProofGroup `protobuf:"bytes,1,opt,name=commit_proof_group,json=commitProofGroup,proto3" json:"commit_proof_group,omitempty"`
	AccProof             *AccProof         `protobuf:"bytes,2,opt,name=acc_proof,json=accProof,proto3" json:"acc_proof,omitempty"`
	MinerId              []byte            `protobuf:"bytes,3,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RequestVerifyCommitAndAccProof) Reset()         { *m = RequestVerifyCommitAndAccProof{} }
func (m *RequestVerifyCommitAndAccProof) String() string { return proto.CompactTextString(m) }
func (*RequestVerifyCommitAndAccProof) ProtoMessage()    {}
func (*RequestVerifyCommitAndAccProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{8}
}

func (m *RequestVerifyCommitAndAccProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVerifyCommitAndAccProof.Unmarshal(m, b)
}
func (m *RequestVerifyCommitAndAccProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVerifyCommitAndAccProof.Marshal(b, m, deterministic)
}
func (m *RequestVerifyCommitAndAccProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVerifyCommitAndAccProof.Merge(m, src)
}
func (m *RequestVerifyCommitAndAccProof) XXX_Size() int {
	return xxx_messageInfo_RequestVerifyCommitAndAccProof.Size(m)
}
func (m *RequestVerifyCommitAndAccProof) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVerifyCommitAndAccProof.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVerifyCommitAndAccProof proto.InternalMessageInfo

func (m *RequestVerifyCommitAndAccProof) GetCommitProofGroup() *CommitProofGroup {
	if m != nil {
		return m.CommitProofGroup
	}
	return nil
}

func (m *RequestVerifyCommitAndAccProof) GetAccProof() *AccProof {
	if m != nil {
		return m.AccProof
	}
	return nil
}

func (m *RequestVerifyCommitAndAccProof) GetMinerId() []byte {
	if m != nil {
		return m.MinerId
	}
	return nil
}

type CommitProofGroup struct {
	CommitProofGroupInner []*CommitProofGroupInner `protobuf:"bytes,1,rep,name=commit_proof_group_inner,json=commitProofGroupInner,proto3" json:"commit_proof_group_inner,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                 `json:"-"`
	XXX_unrecognized      []byte                   `json:"-"`
	XXX_sizecache         int32                    `json:"-"`
}

func (m *CommitProofGroup) Reset()         { *m = CommitProofGroup{} }
func (m *CommitProofGroup) String() string { return proto.CompactTextString(m) }
func (*CommitProofGroup) ProtoMessage()    {}
func (*CommitProofGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{9}
}

func (m *CommitProofGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitProofGroup.Unmarshal(m, b)
}
func (m *CommitProofGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitProofGroup.Marshal(b, m, deterministic)
}
func (m *CommitProofGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitProofGroup.Merge(m, src)
}
func (m *CommitProofGroup) XXX_Size() int {
	return xxx_messageInfo_CommitProofGroup.Size(m)
}
func (m *CommitProofGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitProofGroup.DiscardUnknown(m)
}

var xxx_messageInfo_CommitProofGroup proto.InternalMessageInfo

func (m *CommitProofGroup) GetCommitProofGroupInner() []*CommitProofGroupInner {
	if m != nil {
		return m.CommitProofGroupInner
	}
	return nil
}

type CommitProofGroupInner struct {
	CommitProof          []*CommitProof `protobuf:"bytes,1,rep,name=commit_proof,json=commitProof,proto3" json:"commit_proof,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CommitProofGroupInner) Reset()         { *m = CommitProofGroupInner{} }
func (m *CommitProofGroupInner) String() string { return proto.CompactTextString(m) }
func (*CommitProofGroupInner) ProtoMessage()    {}
func (*CommitProofGroupInner) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{10}
}

func (m *CommitProofGroupInner) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitProofGroupInner.Unmarshal(m, b)
}
func (m *CommitProofGroupInner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitProofGroupInner.Marshal(b, m, deterministic)
}
func (m *CommitProofGroupInner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitProofGroupInner.Merge(m, src)
}
func (m *CommitProofGroupInner) XXX_Size() int {
	return xxx_messageInfo_CommitProofGroupInner.Size(m)
}
func (m *CommitProofGroupInner) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitProofGroupInner.DiscardUnknown(m)
}

var xxx_messageInfo_CommitProofGroupInner proto.InternalMessageInfo

func (m *CommitProofGroupInner) GetCommitProof() []*CommitProof {
	if m != nil {
		return m.CommitProof
	}
	return nil
}

type CommitProof struct {
	Node                 *MhtProof   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Parents              []*MhtProof `protobuf:"bytes,2,rep,name=parents,proto3" json:"parents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CommitProof) Reset()         { *m = CommitProof{} }
func (m *CommitProof) String() string { return proto.CompactTextString(m) }
func (*CommitProof) ProtoMessage()    {}
func (*CommitProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{11}
}

func (m *CommitProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommitProof.Unmarshal(m, b)
}
func (m *CommitProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommitProof.Marshal(b, m, deterministic)
}
func (m *CommitProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommitProof.Merge(m, src)
}
func (m *CommitProof) XXX_Size() int {
	return xxx_messageInfo_CommitProof.Size(m)
}
func (m *CommitProof) XXX_DiscardUnknown() {
	xxx_messageInfo_CommitProof.DiscardUnknown(m)
}

var xxx_messageInfo_CommitProof proto.InternalMessageInfo

func (m *CommitProof) GetNode() *MhtProof {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *CommitProof) GetParents() []*MhtProof {
	if m != nil {
		return m.Parents
	}
	return nil
}

type MhtProof struct {
	Index                int32    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Label                []byte   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Paths                [][]byte `protobuf:"bytes,3,rep,name=paths,proto3" json:"paths,omitempty"`
	Locs                 []byte   `protobuf:"bytes,4,opt,name=locs,proto3" json:"locs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MhtProof) Reset()         { *m = MhtProof{} }
func (m *MhtProof) String() string { return proto.CompactTextString(m) }
func (*MhtProof) ProtoMessage()    {}
func (*MhtProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{12}
}

func (m *MhtProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MhtProof.Unmarshal(m, b)
}
func (m *MhtProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MhtProof.Marshal(b, m, deterministic)
}
func (m *MhtProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MhtProof.Merge(m, src)
}
func (m *MhtProof) XXX_Size() int {
	return xxx_messageInfo_MhtProof.Size(m)
}
func (m *MhtProof) XXX_DiscardUnknown() {
	xxx_messageInfo_MhtProof.DiscardUnknown(m)
}

var xxx_messageInfo_MhtProof proto.InternalMessageInfo

func (m *MhtProof) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *MhtProof) GetLabel() []byte {
	if m != nil {
		return m.Label
	}
	return nil
}

func (m *MhtProof) GetPaths() [][]byte {
	if m != nil {
		return m.Paths
	}
	return nil
}

func (m *MhtProof) GetLocs() []byte {
	if m != nil {
		return m.Locs
	}
	return nil
}

type AccProof struct {
	Indexs               []int64         `protobuf:"varint,1,rep,packed,name=indexs,proto3" json:"indexs,omitempty"`
	Labels               [][]byte        `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty"`
	WitChains            *AccWitnessNode `protobuf:"bytes,3,opt,name=wit_chains,json=witChains,proto3" json:"wit_chains,omitempty"`
	AccPath              [][]byte        `protobuf:"bytes,4,rep,name=acc_path,json=accPath,proto3" json:"acc_path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccProof) Reset()         { *m = AccProof{} }
func (m *AccProof) String() string { return proto.CompactTextString(m) }
func (*AccProof) ProtoMessage()    {}
func (*AccProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{13}
}

func (m *AccProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccProof.Unmarshal(m, b)
}
func (m *AccProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccProof.Marshal(b, m, deterministic)
}
func (m *AccProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccProof.Merge(m, src)
}
func (m *AccProof) XXX_Size() int {
	return xxx_messageInfo_AccProof.Size(m)
}
func (m *AccProof) XXX_DiscardUnknown() {
	xxx_messageInfo_AccProof.DiscardUnknown(m)
}

var xxx_messageInfo_AccProof proto.InternalMessageInfo

func (m *AccProof) GetIndexs() []int64 {
	if m != nil {
		return m.Indexs
	}
	return nil
}

func (m *AccProof) GetLabels() [][]byte {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *AccProof) GetWitChains() *AccWitnessNode {
	if m != nil {
		return m.WitChains
	}
	return nil
}

func (m *AccProof) GetAccPath() [][]byte {
	if m != nil {
		return m.AccPath
	}
	return nil
}

type AccWitnessNode struct {
	Elem                 []byte          `protobuf:"bytes,1,opt,name=elem,proto3" json:"elem,omitempty"`
	Wit                  []byte          `protobuf:"bytes,2,opt,name=wit,proto3" json:"wit,omitempty"`
	Acc                  *AccWitnessNode `protobuf:"bytes,3,opt,name=acc,proto3" json:"acc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccWitnessNode) Reset()         { *m = AccWitnessNode{} }
func (m *AccWitnessNode) String() string { return proto.CompactTextString(m) }
func (*AccWitnessNode) ProtoMessage()    {}
func (*AccWitnessNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{14}
}

func (m *AccWitnessNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccWitnessNode.Unmarshal(m, b)
}
func (m *AccWitnessNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccWitnessNode.Marshal(b, m, deterministic)
}
func (m *AccWitnessNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccWitnessNode.Merge(m, src)
}
func (m *AccWitnessNode) XXX_Size() int {
	return xxx_messageInfo_AccWitnessNode.Size(m)
}
func (m *AccWitnessNode) XXX_DiscardUnknown() {
	xxx_messageInfo_AccWitnessNode.DiscardUnknown(m)
}

var xxx_messageInfo_AccWitnessNode proto.InternalMessageInfo

func (m *AccWitnessNode) GetElem() []byte {
	if m != nil {
		return m.Elem
	}
	return nil
}

func (m *AccWitnessNode) GetWit() []byte {
	if m != nil {
		return m.Wit
	}
	return nil
}

func (m *AccWitnessNode) GetAcc() *AccWitnessNode {
	if m != nil {
		return m.Acc
	}
	return nil
}

type ResponseVerifyCommitAndAccProof struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponseVerifyCommitAndAccProof) Reset()         { *m = ResponseVerifyCommitAndAccProof{} }
func (m *ResponseVerifyCommitAndAccProof) String() string { return proto.CompactTextString(m) }
func (*ResponseVerifyCommitAndAccProof) ProtoMessage()    {}
func (*ResponseVerifyCommitAndAccProof) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{15}
}

func (m *ResponseVerifyCommitAndAccProof) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponseVerifyCommitAndAccProof.Unmarshal(m, b)
}
func (m *ResponseVerifyCommitAndAccProof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponseVerifyCommitAndAccProof.Marshal(b, m, deterministic)
}
func (m *ResponseVerifyCommitAndAccProof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseVerifyCommitAndAccProof.Merge(m, src)
}
func (m *ResponseVerifyCommitAndAccProof) XXX_Size() int {
	return xxx_messageInfo_ResponseVerifyCommitAndAccProof.Size(m)
}
func (m *ResponseVerifyCommitAndAccProof) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseVerifyCommitAndAccProof.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseVerifyCommitAndAccProof proto.InternalMessageInfo

func (m *ResponseVerifyCommitAndAccProof) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

type SpaceVerify struct {
	MinerId              []byte        `protobuf:"bytes,1,opt,name=miner_id,json=minerId,proto3" json:"miner_id,omitempty"`
	SpaceChals           []*Int64Slice `protobuf:"bytes,2,rep,name=space_chals,json=spaceChals,proto3" json:"space_chals,omitempty"`
	SpaceProofLeft       int64         `protobuf:"varint,3,opt,name=space_proof_left,json=spaceProofLeft,proto3" json:"space_proof_left,omitempty"`
	SpaceProofRight      int64         `protobuf:"varint,4,opt,name=space_proof_right,json=spaceProofRight,proto3" json:"space_proof_right,omitempty"`
	MinerProofHash       []byte        `protobuf:"bytes,5,opt,name=miner_proof_hash,json=minerProofHash,proto3" json:"miner_proof_hash,omitempty"`
	VerifyResult         bool          `protobuf:"varint,6,opt,name=verify_result,json=verifyResult,proto3" json:"verify_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SpaceVerify) Reset()         { *m = SpaceVerify{} }
func (m *SpaceVerify) String() string { return proto.CompactTextString(m) }
func (*SpaceVerify) ProtoMessage()    {}
func (*SpaceVerify) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea31f2b7d8186a5c, []int{16}
}

func (m *SpaceVerify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpaceVerify.Unmarshal(m, b)
}
func (m *SpaceVerify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpaceVerify.Marshal(b, m, deterministic)
}
func (m *SpaceVerify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpaceVerify.Merge(m, src)
}
func (m *SpaceVerify) XXX_Size() int {
	return xxx_messageInfo_SpaceVerify.Size(m)
}
func (m *SpaceVerify) XXX_DiscardUnknown() {
	xxx_messageInfo_SpaceVerify.DiscardUnknown(m)
}

var xxx_messageInfo_SpaceVerify proto.InternalMessageInfo

func (m *SpaceVerify) GetMinerId() []byte {
	if m != nil {
		return m.MinerId
	}
	return nil
}

func (m *SpaceVerify) GetSpaceChals() []*Int64Slice {
	if m != nil {
		return m.SpaceChals
	}
	return nil
}

func (m *SpaceVerify) GetSpaceProofLeft() int64 {
	if m != nil {
		return m.SpaceProofLeft
	}
	return 0
}

func (m *SpaceVerify) GetSpaceProofRight() int64 {
	if m != nil {
		return m.SpaceProofRight
	}
	return 0
}

func (m *SpaceVerify) GetMinerProofHash() []byte {
	if m != nil {
		return m.MinerProofHash
	}
	return nil
}

func (m *SpaceVerify) GetVerifyResult() bool {
	if m != nil {
		return m.VerifyResult
	}
	return false
}

func init() {
	proto.RegisterType((*ResponseMinerRegister)(nil), "pois.api.ResponseMinerRegister")
	proto.RegisterType((*RequestMinerInitParam)(nil), "pois.api.RequestMinerInitParam")
	proto.RegisterType((*ResponseMinerInitParam)(nil), "pois.api.ResponseMinerInitParam")
	proto.RegisterType((*RequestMinerCommitGenChall)(nil), "pois.api.RequestMinerCommitGenChall")
	proto.RegisterType((*ProverAcc)(nil), "pois.api.ProverAcc")
	proto.RegisterType((*Commit)(nil), "pois.api.Commit")
	proto.RegisterType((*Challenge)(nil), "pois.api.Challenge")
	proto.RegisterType((*Int64Slice)(nil), "pois.api.Int64Slice")
	proto.RegisterType((*RequestVerifyCommitAndAccProof)(nil), "pois.api.RequestVerifyCommitAndAccProof")
	proto.RegisterType((*CommitProofGroup)(nil), "pois.api.CommitProofGroup")
	proto.RegisterType((*CommitProofGroupInner)(nil), "pois.api.CommitProofGroupInner")
	proto.RegisterType((*CommitProof)(nil), "pois.api.CommitProof")
	proto.RegisterType((*MhtProof)(nil), "pois.api.MhtProof")
	proto.RegisterType((*AccProof)(nil), "pois.api.AccProof")
	proto.RegisterType((*AccWitnessNode)(nil), "pois.api.AccWitnessNode")
	proto.RegisterType((*ResponseVerifyCommitAndAccProof)(nil), "pois.api.ResponseVerifyCommitAndAccProof")
	proto.RegisterType((*SpaceVerify)(nil), "pois.api.SpaceVerify")
}

func init() { proto.RegisterFile("pois_api.proto", fileDescriptor_ea31f2b7d8186a5c) }

var fileDescriptor_ea31f2b7d8186a5c = []byte{
	// 938 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0x8f, 0xb3, 0x8e, 0x63, 0x3f, 0xbb, 0xc1, 0x4c, 0x9b, 0x68, 0x6b, 0x68, 0x13, 0x0d, 0x15,
	0x32, 0x08, 0x39, 0x92, 0xa1, 0xfc, 0x39, 0x70, 0x08, 0x39, 0xa4, 0x96, 0x68, 0x15, 0xa6, 0x12,
	0x20, 0x04, 0xda, 0x4e, 0xd6, 0x63, 0xef, 0x28, 0xeb, 0xd9, 0x65, 0x66, 0x1c, 0x93, 0x3b, 0x67,
	0x6e, 0x5c, 0xf9, 0x18, 0xf0, 0xf5, 0xd0, 0xbc, 0xd9, 0xb5, 0xd7, 0x89, 0xdd, 0xde, 0xe6, 0xfd,
	0xe6, 0xf7, 0xfe, 0xee, 0x7b, 0x6f, 0x16, 0x0e, 0xf2, 0x4c, 0x9a, 0x88, 0xe7, 0x72, 0x90, 0xeb,
	0xcc, 0x66, 0xa4, 0xe9, 0xe4, 0x01, 0xcf, 0x25, 0x3d, 0x85, 0x43, 0x26, 0x4c, 0x9e, 0x29, 0x23,
	0x5e, 0x4a, 0x25, 0x34, 0x13, 0x53, 0x69, 0xac, 0xd0, 0xe4, 0x08, 0x1a, 0xc6, 0x72, 0x3b, 0x37,
	0x61, 0xed, 0xa4, 0xd6, 0x6f, 0xb2, 0x42, 0xa2, 0x43, 0xa7, 0xf0, 0xfb, 0x5c, 0x18, 0x8b, 0xfc,
	0x91, 0x92, 0xf6, 0x92, 0x6b, 0x3e, 0x23, 0x8f, 0xa1, 0x39, 0x73, 0x48, 0x24, 0xc7, 0xa8, 0xd2,
	0x61, 0xfb, 0x28, 0x8f, 0xc6, 0xf4, 0xbf, 0x1a, 0x1c, 0xad, 0x79, 0x59, 0x69, 0x75, 0x21, 0xe0,
	0x71, 0x5c, 0x28, 0xb8, 0x23, 0x79, 0x08, 0x7b, 0xd7, 0xe2, 0x36, 0x52, 0xe1, 0x2e, 0x62, 0xf5,
	0x6b, 0x71, 0xfb, 0xaa, 0x04, 0xa7, 0x61, 0xb0, 0x04, 0x2f, 0xc8, 0x23, 0xd8, 0x9b, 0xe8, 0x4c,
	0xd9, 0xb0, 0x7e, 0x52, 0xeb, 0x07, 0xcc, 0x0b, 0x84, 0x40, 0x5d, 0x0b, 0xae, 0xc3, 0x3d, 0x04,
	0xf1, 0xbc, 0x16, 0x5b, 0x63, 0x2d, 0x36, 0xf2, 0x21, 0xb4, 0x8c, 0x9c, 0x2a, 0x6e, 0xe7, 0x5a,
	0x84, 0xfb, 0x78, 0xb7, 0x02, 0xe8, 0xdf, 0x35, 0xe8, 0x55, 0xd3, 0x3d, 0xcf, 0x66, 0x33, 0x69,
	0x2f, 0x84, 0x3a, 0x4f, 0x78, 0x9a, 0xbe, 0x25, 0x67, 0xd2, 0x87, 0x46, 0x8c, 0xe4, 0x70, 0xf7,
	0x24, 0xe8, 0xb7, 0x87, 0xdd, 0x41, 0x59, 0xf3, 0x81, 0x37, 0xc2, 0x8a, 0x7b, 0x32, 0x04, 0xc8,
	0x75, 0x76, 0x23, 0x74, 0xe4, 0x2a, 0xe1, 0x12, 0x6c, 0x0f, 0x1f, 0xae, 0xd8, 0x97, 0x78, 0x77,
	0x16, 0xc7, 0xac, 0x95, 0x97, 0x47, 0xfa, 0x04, 0x5a, 0x4b, 0xfc, 0x7e, 0x0d, 0xe9, 0xb7, 0xd0,
	0xf0, 0x4e, 0xc8, 0x13, 0x80, 0x89, 0x4c, 0x45, 0x24, 0xd5, 0x58, 0xfc, 0x81, 0x94, 0x80, 0xb5,
	0x1c, 0x32, 0x72, 0x80, 0x2b, 0xa1, 0xce, 0x32, 0x6b, 0x30, 0xc8, 0x0e, 0xf3, 0x02, 0x7d, 0x0e,
	0x2d, 0xcc, 0x4f, 0xa8, 0xa9, 0x20, 0x7d, 0xa8, 0xeb, 0x6c, 0xe1, 0xda, 0xc0, 0xa5, 0xf1, 0x68,
	0x15, 0xd8, 0x48, 0xd9, 0x2f, 0xbf, 0x78, 0x9d, 0xca, 0x58, 0x30, 0x64, 0xd0, 0x67, 0x00, 0x2b,
	0xcc, 0x35, 0xd0, 0x0d, 0x4f, 0xe7, 0xc2, 0x6b, 0x06, 0xac, 0x90, 0xe8, 0xbf, 0x35, 0x78, 0x5a,
	0x94, 0xf4, 0x47, 0xa1, 0xe5, 0xe4, 0xd6, 0x47, 0x7a, 0xa6, 0xc6, 0x67, 0x71, 0x7c, 0xa9, 0xb3,
	0x6c, 0x42, 0x5e, 0x00, 0xf1, 0xb5, 0x89, 0x72, 0x27, 0x47, 0x53, 0x9d, 0xcd, 0x73, 0x0c, 0xbe,
	0x3d, 0xec, 0xdd, 0xad, 0x23, 0xaa, 0x5c, 0x38, 0x06, 0xeb, 0xc6, 0x77, 0x10, 0x72, 0x0a, 0x2d,
	0x1e, 0xc7, 0xde, 0x0c, 0x36, 0x54, 0x7b, 0x48, 0x56, 0x06, 0x4a, 0x87, 0xac, 0xc9, 0x4b, 0xd7,
	0xd5, 0x2f, 0x1a, 0xac, 0x77, 0x71, 0x0a, 0xdd, 0xbb, 0x1e, 0xc9, 0xcf, 0x10, 0xde, 0x8f, 0x34,
	0x92, 0x4a, 0x09, 0x5d, 0x14, 0xec, 0x78, 0x7b, 0xbc, 0x23, 0x47, 0x63, 0x87, 0xf1, 0x26, 0x98,
	0xfe, 0x00, 0x87, 0x1b, 0xf9, 0xe4, 0x6b, 0xe8, 0x54, 0x5d, 0x16, 0x6e, 0x0e, 0x37, 0xba, 0x61,
	0xed, 0x8a, 0x71, 0x1a, 0x43, 0xbb, 0x72, 0x47, 0x3e, 0x86, 0xba, 0xca, 0xc6, 0xa2, 0xa8, 0x6b,
	0xa5, 0x2c, 0x2f, 0x93, 0x42, 0x1b, 0xef, 0xc9, 0x67, 0xb0, 0x9f, 0x73, 0x2d, 0x54, 0xd1, 0x25,
	0x9b, 0xa9, 0x25, 0x85, 0xbe, 0x81, 0x66, 0x09, 0xba, 0xee, 0x5a, 0xf5, 0xdd, 0x1e, 0xf3, 0x82,
	0x43, 0x53, 0x7e, 0x25, 0xd2, 0x62, 0xc0, 0xbd, 0xe0, 0xd0, 0x9c, 0xdb, 0xc4, 0x84, 0x81, 0xef,
	0x44, 0x14, 0xdc, 0x30, 0xa7, 0x59, 0x6c, 0x70, 0xc2, 0x3b, 0x0c, 0xcf, 0xf4, 0xaf, 0x1a, 0x34,
	0x97, 0xad, 0x72, 0x04, 0x0d, 0xb4, 0xba, 0xec, 0x32, 0x2f, 0x39, 0x1c, 0xed, 0x96, 0x9d, 0x5d,
	0x48, 0xe4, 0x2b, 0x80, 0x85, 0xb4, 0x51, 0x9c, 0x70, 0xa9, 0x4c, 0x31, 0x6c, 0xe1, 0x5a, 0x47,
	0xfc, 0x24, 0xad, 0x12, 0xc6, 0xbc, 0xca, 0xc6, 0x82, 0xb5, 0x16, 0xd2, 0x9e, 0x23, 0xd5, 0x35,
	0x06, 0x76, 0x12, 0xb7, 0x49, 0x58, 0x47, 0x93, 0xfb, 0xae, 0x69, 0xb8, 0x4d, 0xe8, 0x15, 0x1c,
	0xac, 0xeb, 0xb9, 0xb0, 0x45, 0x2a, 0x66, 0xc5, 0x48, 0xe2, 0xd9, 0x4d, 0xe9, 0x02, 0xb7, 0x01,
	0x4e, 0xe9, 0x42, 0x5a, 0xf2, 0xa9, 0x9f, 0xdb, 0x77, 0x05, 0x81, 0x13, 0xfd, 0x0d, 0x1c, 0x97,
	0x1b, 0x74, 0xdb, 0xd4, 0x6c, 0xdb, 0xd8, 0x7f, 0xee, 0x42, 0xfb, 0x75, 0xce, 0xe3, 0x42, 0xf1,
	0x6d, 0x4b, 0xeb, 0x39, 0xb4, 0x8d, 0x63, 0xba, 0xfa, 0xa4, 0xe5, 0xe7, 0xde, 0x3c, 0xf2, 0x80,
	0x44, 0xb7, 0x26, 0x0c, 0xe9, 0x43, 0xd7, 0xab, 0xf9, 0x21, 0x48, 0xc5, 0xc4, 0x62, 0x56, 0x01,
	0x3b, 0x40, 0x1c, 0xe3, 0xfb, 0x5e, 0x4c, 0x5c, 0xca, 0xef, 0x57, 0x99, 0x5a, 0x4e, 0x93, 0x72,
	0x7d, 0xbf, 0xb7, 0xa2, 0x32, 0x07, 0x3b, 0xab, 0x3e, 0x4e, 0xcf, 0x4d, 0xb8, 0x49, 0x70, 0xa9,
	0x77, 0xd8, 0x01, 0xe2, 0x48, 0x7d, 0xc1, 0x4d, 0x42, 0x3e, 0x82, 0x07, 0x37, 0x98, 0x5b, 0xa4,
	0x85, 0x99, 0xa7, 0x16, 0x77, 0x7c, 0x93, 0x75, 0x3c, 0xc8, 0x10, 0x1b, 0xfe, 0x13, 0xc0, 0xfe,
	0x65, 0x26, 0xcd, 0x59, 0x2e, 0xc9, 0x1b, 0x78, 0xac, 0xfd, 0x0a, 0x8a, 0xbc, 0x8b, 0xa9, 0xb0,
	0x91, 0x12, 0x8b, 0xe8, 0x5a, 0xdc, 0x92, 0xca, 0xc4, 0x6e, 0x7c, 0xe9, 0x7a, 0x27, 0x55, 0xc2,
	0xa6, 0x57, 0x8d, 0xee, 0x90, 0x5f, 0xe1, 0x68, 0xdd, 0x83, 0x2e, 0x1f, 0xd6, 0x77, 0x9a, 0x3f,
	0xde, 0x62, 0xbe, 0x7c, 0x9a, 0xe9, 0x0e, 0xf9, 0x0d, 0x9e, 0xae, 0x5b, 0x2f, 0x36, 0xc2, 0x54,
	0x28, 0xfc, 0x74, 0x29, 0x79, 0xb6, 0xd9, 0xcb, 0xfa, 0xfb, 0xd5, 0xab, 0x3c, 0x33, 0xcb, 0x85,
	0x4f, 0x77, 0x88, 0x85, 0x0f, 0x4a, 0xf3, 0x45, 0x5d, 0xab, 0x1b, 0x87, 0xf4, 0xef, 0xd9, 0xde,
	0xd2, 0x92, 0xbd, 0x4f, 0xee, 0xa7, 0xb2, 0x85, 0x4a, 0x77, 0xbe, 0x7b, 0xf0, 0x4b, 0x7b, 0x30,
	0x38, 0xc5, 0x3f, 0x15, 0x9d, 0xc7, 0x57, 0x0d, 0xfc, 0x55, 0xf9, 0xfc, 0xff, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x14, 0x11, 0x35, 0x4c, 0xbc, 0x08, 0x00, 0x00,
}