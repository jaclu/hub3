// Copyright 2020 Delving B.V.
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
// 	protoc-gen-go v1.20.1
// 	protoc        v3.11.4
// source: ikuzo/domain/domainpb/domain.proto

package domainpb

import (
	proto "github.com/golang/protobuf/proto"
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

type RecordType int32

const (
	RecordType_NARTHEX    RecordType = 0
	RecordType_SCHEMA     RecordType = 1
	RecordType_VOCABULARY RecordType = 2
	RecordType_SOURCE     RecordType = 3
	RecordType_CACHE      RecordType = 4
)

// Enum value maps for RecordType.
var (
	RecordType_name = map[int32]string{
		0: "NARTHEX",
		1: "SCHEMA",
		2: "VOCABULARY",
		3: "SOURCE",
		4: "CACHE",
	}
	RecordType_value = map[string]int32{
		"NARTHEX":    0,
		"SCHEMA":     1,
		"VOCABULARY": 2,
		"SOURCE":     3,
		"CACHE":      4,
	}
)

func (x RecordType) Enum() *RecordType {
	p := new(RecordType)
	*p = x
	return p
}

func (x RecordType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RecordType) Descriptor() protoreflect.EnumDescriptor {
	return file_ikuzo_domain_domainpb_domain_proto_enumTypes[0].Descriptor()
}

func (RecordType) Type() protoreflect.EnumType {
	return &file_ikuzo_domain_domainpb_domain_proto_enumTypes[0]
}

func (x RecordType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RecordType.Descriptor instead.
func (RecordType) EnumDescriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{0}
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The tenant identifier for this RecordType.
	OrgID string `protobuf:"bytes,1,opt,name=orgID,proto3" json:"orgID,omitempty"`
	// The spec is the unique dataset string to identify which dataset  this
	// Fragment belongs to
	Spec string `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// The revision is used to determine which version is an orphan and should be removed
	Revision int32 `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	// The hubId is the unique identifier for any document record in hub3
	HubID string `protobuf:"bytes,4,opt,name=hubID,proto3" json:"hubID,omitempty"`
	// Each fragment can be tagged with additional metadata. This can be queried for.
	// Some examples are 'resource', 'literal', 'bnode', 'rdfType', etc.
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	// The document type for ElasticSearch. This is a constant value
	DocType string `protobuf:"bytes,6,opt,name=docType,proto3" json:"docType,omitempty"`
	// The subject of the graph stored
	EntryURI string `protobuf:"bytes,7,opt,name=entryURI,proto3" json:"entryURI,omitempty"`
	// the graph name of the graph stored
	NamedGraphURI string `protobuf:"bytes,8,opt,name=namedGraphURI,proto3" json:"namedGraphURI,omitempty"`
	// miliseconds since epoch
	Modified int64 `protobuf:"varint,9,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetOrgID() string {
	if x != nil {
		return x.OrgID
	}
	return ""
}

func (x *Header) GetSpec() string {
	if x != nil {
		return x.Spec
	}
	return ""
}

func (x *Header) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

func (x *Header) GetHubID() string {
	if x != nil {
		return x.HubID
	}
	return ""
}

func (x *Header) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Header) GetDocType() string {
	if x != nil {
		return x.DocType
	}
	return ""
}

func (x *Header) GetEntryURI() string {
	if x != nil {
		return x.EntryURI
	}
	return ""
}

func (x *Header) GetNamedGraphURI() string {
	if x != nil {
		return x.NamedGraphURI
	}
	return ""
}

func (x *Header) GetModified() int64 {
	if x != nil {
		return x.Modified
	}
	return 0
}

type RecordGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta          *Header     `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	EntryURI      string      `protobuf:"bytes,2,opt,name=entryURI,proto3" json:"entryURI,omitempty"`
	NamedGraphURI string      `protobuf:"bytes,3,opt,name=namedGraphURI,proto3" json:"namedGraphURI,omitempty"`
	RecordType    RecordType  `protobuf:"varint,4,opt,name=recordType,proto3,enum=domainpb.RecordType" json:"recordType,omitempty"`
	Resources     []*Resource `protobuf:"bytes,5,rep,name=resources,proto3" json:"resources,omitempty"`
}

func (x *RecordGraph) Reset() {
	*x = RecordGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordGraph) ProtoMessage() {}

func (x *RecordGraph) ProtoReflect() protoreflect.Message {
	mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordGraph.ProtoReflect.Descriptor instead.
func (*RecordGraph) Descriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{1}
}

func (x *RecordGraph) GetMeta() *Header {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *RecordGraph) GetEntryURI() string {
	if x != nil {
		return x.EntryURI
	}
	return ""
}

func (x *RecordGraph) GetNamedGraphURI() string {
	if x != nil {
		return x.NamedGraphURI
	}
	return ""
}

func (x *RecordGraph) GetRecordType() RecordType {
	if x != nil {
		return x.RecordType
	}
	return RecordType_NARTHEX
}

func (x *RecordGraph) GetResources() []*Resource {
	if x != nil {
		return x.Resources
	}
	return nil
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                   string             `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Types                []string           `protobuf:"bytes,2,rep,name=Types,proto3" json:"Types,omitempty"`
	GraphExternalContext []*ReferrerContext `protobuf:"bytes,3,rep,name=GraphExternalContext,proto3" json:"GraphExternalContext,omitempty"`
	Context              []*ReferrerContext `protobuf:"bytes,4,rep,name=Context,proto3" json:"Context,omitempty"`
	Predicates           map[string]*Entry  `protobuf:"bytes,6,rep,name=Predicates,proto3" json:"Predicates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ObjectIDs            []*ReferrerContext `protobuf:"bytes,7,rep,name=ObjectIDs,proto3" json:"ObjectIDs,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{2}
}

func (x *Resource) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Resource) GetTypes() []string {
	if x != nil {
		return x.Types
	}
	return nil
}

func (x *Resource) GetGraphExternalContext() []*ReferrerContext {
	if x != nil {
		return x.GraphExternalContext
	}
	return nil
}

func (x *Resource) GetContext() []*ReferrerContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *Resource) GetPredicates() map[string]*Entry {
	if x != nil {
		return x.Predicates
	}
	return nil
}

func (x *Resource) GetObjectIDs() []*ReferrerContext {
	if x != nil {
		return x.ObjectIDs
	}
	return nil
}

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Value       string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	Language    string `protobuf:"bytes,3,opt,name=Language,proto3" json:"Language,omitempty"`
	Datatype    string `protobuf:"bytes,4,opt,name=Datatype,proto3" json:"Datatype,omitempty"`
	Entrytype   string `protobuf:"bytes,5,opt,name=Entrytype,proto3" json:"Entrytype,omitempty"`
	Predicate   string `protobuf:"bytes,6,opt,name=Predicate,proto3" json:"Predicate,omitempty"`
	SearchLabel string `protobuf:"bytes,7,opt,name=SearchLabel,proto3" json:"SearchLabel,omitempty"`
	// the rest are index specific fields that take the Value and add it to typed fields
	Level     int32       `protobuf:"varint,8,opt,name=level,proto3" json:"level,omitempty"`
	Tags      []string    `protobuf:"bytes,9,rep,name=tags,proto3" json:"tags,omitempty"`
	Date      []string    `protobuf:"bytes,10,rep,name=date,proto3" json:"date,omitempty"`
	Integer   int32       `protobuf:"varint,11,opt,name=integer,proto3" json:"integer,omitempty"`
	Float     float32     `protobuf:"fixed32,12,opt,name=Float,proto3" json:"Float,omitempty"`
	IntRange  *IndexRange `protobuf:"bytes,13,opt,name=intRange,proto3" json:"intRange,omitempty"`
	DateRange *IndexRange `protobuf:"bytes,14,opt,name=dateRange,proto3" json:"dateRange,omitempty"`
	LatLong   string      `protobuf:"bytes,15,opt,name=latLong,proto3" json:"latLong,omitempty"`
	Order     int32       `protobuf:"varint,16,opt,name=order,proto3" json:"order,omitempty"`
	// Inline is only used for presentation after the graph is resolved
	Inline *Resource `protobuf:"bytes,17,opt,name=Inline,proto3" json:"Inline,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{3}
}

func (x *Entry) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Entry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Entry) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Entry) GetDatatype() string {
	if x != nil {
		return x.Datatype
	}
	return ""
}

func (x *Entry) GetEntrytype() string {
	if x != nil {
		return x.Entrytype
	}
	return ""
}

func (x *Entry) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

func (x *Entry) GetSearchLabel() string {
	if x != nil {
		return x.SearchLabel
	}
	return ""
}

func (x *Entry) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *Entry) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Entry) GetDate() []string {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Entry) GetInteger() int32 {
	if x != nil {
		return x.Integer
	}
	return 0
}

func (x *Entry) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Entry) GetIntRange() *IndexRange {
	if x != nil {
		return x.IntRange
	}
	return nil
}

func (x *Entry) GetDateRange() *IndexRange {
	if x != nil {
		return x.DateRange
	}
	return nil
}

func (x *Entry) GetLatLong() string {
	if x != nil {
		return x.LatLong
	}
	return ""
}

func (x *Entry) GetOrder() int32 {
	if x != nil {
		return x.Order
	}
	return 0
}

func (x *Entry) GetInline() *Resource {
	if x != nil {
		return x.Inline
	}
	return nil
}

type IndexRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gte string `protobuf:"bytes,1,opt,name=gte,proto3" json:"gte,omitempty"`
	Lte string `protobuf:"bytes,2,opt,name=lte,proto3" json:"lte,omitempty"`
}

func (x *IndexRange) Reset() {
	*x = IndexRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRange) ProtoMessage() {}

func (x *IndexRange) ProtoReflect() protoreflect.Message {
	mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRange.ProtoReflect.Descriptor instead.
func (*IndexRange) Descriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{4}
}

func (x *IndexRange) GetGte() string {
	if x != nil {
		return x.Gte
	}
	return ""
}

func (x *IndexRange) GetLte() string {
	if x != nil {
		return x.Lte
	}
	return ""
}

// ReferrerContext holds the referrer in formation for creating new fragments
type ReferrerContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subject      string   `protobuf:"bytes,1,opt,name=Subject,proto3" json:"Subject,omitempty"`
	SubjectClass []string `protobuf:"bytes,2,rep,name=SubjectClass,proto3" json:"SubjectClass,omitempty"`
	Predicate    string   `protobuf:"bytes,3,opt,name=Predicate,proto3" json:"Predicate,omitempty"`
	SearchLabel  string   `protobuf:"bytes,4,opt,name=SearchLabel,proto3" json:"SearchLabel,omitempty"`
	Level        int32    `protobuf:"varint,5,opt,name=Level,proto3" json:"Level,omitempty"`
	ObjectID     string   `protobuf:"bytes,6,opt,name=ObjectID,proto3" json:"ObjectID,omitempty"`
	SortKey      int32    `protobuf:"varint,7,opt,name=SortKey,proto3" json:"SortKey,omitempty"`
	Label        string   `protobuf:"bytes,8,opt,name=Label,proto3" json:"Label,omitempty"`
}

func (x *ReferrerContext) Reset() {
	*x = ReferrerContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReferrerContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReferrerContext) ProtoMessage() {}

func (x *ReferrerContext) ProtoReflect() protoreflect.Message {
	mi := &file_ikuzo_domain_domainpb_domain_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReferrerContext.ProtoReflect.Descriptor instead.
func (*ReferrerContext) Descriptor() ([]byte, []int) {
	return file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP(), []int{5}
}

func (x *ReferrerContext) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *ReferrerContext) GetSubjectClass() []string {
	if x != nil {
		return x.SubjectClass
	}
	return nil
}

func (x *ReferrerContext) GetPredicate() string {
	if x != nil {
		return x.Predicate
	}
	return ""
}

func (x *ReferrerContext) GetSearchLabel() string {
	if x != nil {
		return x.SearchLabel
	}
	return ""
}

func (x *ReferrerContext) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ReferrerContext) GetObjectID() string {
	if x != nil {
		return x.ObjectID
	}
	return ""
}

func (x *ReferrerContext) GetSortKey() int32 {
	if x != nil {
		return x.SortKey
	}
	return 0
}

func (x *ReferrerContext) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_ikuzo_domain_domainpb_domain_proto protoreflect.FileDescriptor

var file_ikuzo_domain_domainpb_domain_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x6b, 0x75, 0x7a, 0x6f, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x22, 0xf0,
	0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x44, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x68, 0x75, 0x62, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x68, 0x75, 0x62, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6f, 0x63, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x55, 0x52, 0x49, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x55, 0x52, 0x49, 0x12,
	0x24, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x52, 0x49,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x55, 0x52, 0x49, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x22, 0xdd, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x12, 0x24, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x55, 0x52, 0x49, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x72, 0x79,
	0x55, 0x52, 0x49, 0x12, 0x24, 0x0a, 0x0d, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x55, 0x52, 0x49, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x61, 0x6d, 0x65,
	0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x55, 0x52, 0x49, 0x12, 0x34, 0x0a, 0x0a, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x22, 0x81, 0x03, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14,
	0x0a, 0x05, 0x54, 0x79, 0x70, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x14, 0x47, 0x72, 0x61, 0x70, 0x68, 0x45, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x14, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x33, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e,
	0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x0a, 0x50, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0a, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x37, 0x0a, 0x09,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x09, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x73, 0x1a, 0x4e, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf3, 0x03, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x61, 0x74, 0x61, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x50,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6c, 0x61, 0x74, 0x4c, 0x6f, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x2a, 0x0a, 0x06, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x52, 0x06, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x30, 0x0a, 0x0a, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x67, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6c, 0x74, 0x65, 0x22, 0xf1, 0x01,
	0x0a, 0x0f, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x53, 0x6f, 0x72, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x2a, 0x4c, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x4e, 0x41, 0x52, 0x54, 0x48, 0x45, 0x58, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x43, 0x48, 0x45, 0x4d, 0x41, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x4f, 0x43, 0x41,
	0x42, 0x55, 0x4c, 0x41, 0x52, 0x59, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x41, 0x43, 0x48, 0x45, 0x10, 0x04, 0x42,
	0x17, 0x5a, 0x15, 0x69, 0x6b, 0x75, 0x7a, 0x6f, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ikuzo_domain_domainpb_domain_proto_rawDescOnce sync.Once
	file_ikuzo_domain_domainpb_domain_proto_rawDescData = file_ikuzo_domain_domainpb_domain_proto_rawDesc
)

func file_ikuzo_domain_domainpb_domain_proto_rawDescGZIP() []byte {
	file_ikuzo_domain_domainpb_domain_proto_rawDescOnce.Do(func() {
		file_ikuzo_domain_domainpb_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_ikuzo_domain_domainpb_domain_proto_rawDescData)
	})
	return file_ikuzo_domain_domainpb_domain_proto_rawDescData
}

var file_ikuzo_domain_domainpb_domain_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ikuzo_domain_domainpb_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ikuzo_domain_domainpb_domain_proto_goTypes = []interface{}{
	(RecordType)(0),         // 0: domainpb.RecordType
	(*Header)(nil),          // 1: domainpb.Header
	(*RecordGraph)(nil),     // 2: domainpb.RecordGraph
	(*Resource)(nil),        // 3: domainpb.Resource
	(*Entry)(nil),           // 4: domainpb.Entry
	(*IndexRange)(nil),      // 5: domainpb.IndexRange
	(*ReferrerContext)(nil), // 6: domainpb.ReferrerContext
	nil,                     // 7: domainpb.Resource.PredicatesEntry
}
var file_ikuzo_domain_domainpb_domain_proto_depIdxs = []int32{
	1,  // 0: domainpb.RecordGraph.meta:type_name -> domainpb.Header
	0,  // 1: domainpb.RecordGraph.recordType:type_name -> domainpb.RecordType
	3,  // 2: domainpb.RecordGraph.resources:type_name -> domainpb.Resource
	6,  // 3: domainpb.Resource.GraphExternalContext:type_name -> domainpb.ReferrerContext
	6,  // 4: domainpb.Resource.Context:type_name -> domainpb.ReferrerContext
	7,  // 5: domainpb.Resource.Predicates:type_name -> domainpb.Resource.PredicatesEntry
	6,  // 6: domainpb.Resource.ObjectIDs:type_name -> domainpb.ReferrerContext
	5,  // 7: domainpb.Entry.intRange:type_name -> domainpb.IndexRange
	5,  // 8: domainpb.Entry.dateRange:type_name -> domainpb.IndexRange
	3,  // 9: domainpb.Entry.Inline:type_name -> domainpb.Resource
	4,  // 10: domainpb.Resource.PredicatesEntry.value:type_name -> domainpb.Entry
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_ikuzo_domain_domainpb_domain_proto_init() }
func file_ikuzo_domain_domainpb_domain_proto_init() {
	if File_ikuzo_domain_domainpb_domain_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ikuzo_domain_domainpb_domain_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_ikuzo_domain_domainpb_domain_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordGraph); i {
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
		file_ikuzo_domain_domainpb_domain_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_ikuzo_domain_domainpb_domain_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_ikuzo_domain_domainpb_domain_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRange); i {
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
		file_ikuzo_domain_domainpb_domain_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReferrerContext); i {
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
			RawDescriptor: file_ikuzo_domain_domainpb_domain_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ikuzo_domain_domainpb_domain_proto_goTypes,
		DependencyIndexes: file_ikuzo_domain_domainpb_domain_proto_depIdxs,
		EnumInfos:         file_ikuzo_domain_domainpb_domain_proto_enumTypes,
		MessageInfos:      file_ikuzo_domain_domainpb_domain_proto_msgTypes,
	}.Build()
	File_ikuzo_domain_domainpb_domain_proto = out.File
	file_ikuzo_domain_domainpb_domain_proto_rawDesc = nil
	file_ikuzo_domain_domainpb_domain_proto_goTypes = nil
	file_ikuzo_domain_domainpb_domain_proto_depIdxs = nil
}
