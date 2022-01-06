// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/core/types.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Define a set of simple types.
type SimpleType int32

const (
	SimpleType_NONE     SimpleType = 0
	SimpleType_INTEGER  SimpleType = 1
	SimpleType_FLOAT    SimpleType = 2
	SimpleType_STRING   SimpleType = 3
	SimpleType_BOOLEAN  SimpleType = 4
	SimpleType_DATETIME SimpleType = 5
	SimpleType_DURATION SimpleType = 6
	SimpleType_BINARY   SimpleType = 7
	SimpleType_ERROR    SimpleType = 8
	SimpleType_STRUCT   SimpleType = 9
)

var SimpleType_name = map[int32]string{
	0: "NONE",
	1: "INTEGER",
	2: "FLOAT",
	3: "STRING",
	4: "BOOLEAN",
	5: "DATETIME",
	6: "DURATION",
	7: "BINARY",
	8: "ERROR",
	9: "STRUCT",
}

var SimpleType_value = map[string]int32{
	"NONE":     0,
	"INTEGER":  1,
	"FLOAT":    2,
	"STRING":   3,
	"BOOLEAN":  4,
	"DATETIME": 5,
	"DURATION": 6,
	"BINARY":   7,
	"ERROR":    8,
	"STRUCT":   9,
}

func (x SimpleType) String() string {
	return proto.EnumName(SimpleType_name, int32(x))
}

func (SimpleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{0}
}

type SchemaType_SchemaColumn_SchemaColumnType int32

const (
	SchemaType_SchemaColumn_INTEGER  SchemaType_SchemaColumn_SchemaColumnType = 0
	SchemaType_SchemaColumn_FLOAT    SchemaType_SchemaColumn_SchemaColumnType = 1
	SchemaType_SchemaColumn_STRING   SchemaType_SchemaColumn_SchemaColumnType = 2
	SchemaType_SchemaColumn_BOOLEAN  SchemaType_SchemaColumn_SchemaColumnType = 3
	SchemaType_SchemaColumn_DATETIME SchemaType_SchemaColumn_SchemaColumnType = 4
	SchemaType_SchemaColumn_DURATION SchemaType_SchemaColumn_SchemaColumnType = 5
)

var SchemaType_SchemaColumn_SchemaColumnType_name = map[int32]string{
	0: "INTEGER",
	1: "FLOAT",
	2: "STRING",
	3: "BOOLEAN",
	4: "DATETIME",
	5: "DURATION",
}

var SchemaType_SchemaColumn_SchemaColumnType_value = map[string]int32{
	"INTEGER":  0,
	"FLOAT":    1,
	"STRING":   2,
	"BOOLEAN":  3,
	"DATETIME": 4,
	"DURATION": 5,
}

func (x SchemaType_SchemaColumn_SchemaColumnType) String() string {
	return proto.EnumName(SchemaType_SchemaColumn_SchemaColumnType_name, int32(x))
}

func (SchemaType_SchemaColumn_SchemaColumnType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{0, 0, 0}
}

type BlobType_BlobDimensionality int32

const (
	BlobType_SINGLE    BlobType_BlobDimensionality = 0
	BlobType_MULTIPART BlobType_BlobDimensionality = 1
)

var BlobType_BlobDimensionality_name = map[int32]string{
	0: "SINGLE",
	1: "MULTIPART",
}

var BlobType_BlobDimensionality_value = map[string]int32{
	"SINGLE":    0,
	"MULTIPART": 1,
}

func (x BlobType_BlobDimensionality) String() string {
	return proto.EnumName(BlobType_BlobDimensionality_name, int32(x))
}

func (BlobType_BlobDimensionality) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{2, 0}
}

// Defines schema columns and types to strongly type-validate schemas interoperability.
type SchemaType struct {
	// A list of ordered columns this schema comprises of.
	Columns              []*SchemaType_SchemaColumn `protobuf:"bytes,3,rep,name=columns,proto3" json:"columns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SchemaType) Reset()         { *m = SchemaType{} }
func (m *SchemaType) String() string { return proto.CompactTextString(m) }
func (*SchemaType) ProtoMessage()    {}
func (*SchemaType) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{0}
}

func (m *SchemaType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaType.Unmarshal(m, b)
}
func (m *SchemaType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaType.Marshal(b, m, deterministic)
}
func (m *SchemaType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaType.Merge(m, src)
}
func (m *SchemaType) XXX_Size() int {
	return xxx_messageInfo_SchemaType.Size(m)
}
func (m *SchemaType) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaType.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaType proto.InternalMessageInfo

func (m *SchemaType) GetColumns() []*SchemaType_SchemaColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

type SchemaType_SchemaColumn struct {
	// A unique name -within the schema type- for the column
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The column type. This allows a limited set of types currently.
	Type                 SchemaType_SchemaColumn_SchemaColumnType `protobuf:"varint,2,opt,name=type,proto3,enum=flyteidl.core.SchemaType_SchemaColumn_SchemaColumnType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *SchemaType_SchemaColumn) Reset()         { *m = SchemaType_SchemaColumn{} }
func (m *SchemaType_SchemaColumn) String() string { return proto.CompactTextString(m) }
func (*SchemaType_SchemaColumn) ProtoMessage()    {}
func (*SchemaType_SchemaColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{0, 0}
}

func (m *SchemaType_SchemaColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SchemaType_SchemaColumn.Unmarshal(m, b)
}
func (m *SchemaType_SchemaColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SchemaType_SchemaColumn.Marshal(b, m, deterministic)
}
func (m *SchemaType_SchemaColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SchemaType_SchemaColumn.Merge(m, src)
}
func (m *SchemaType_SchemaColumn) XXX_Size() int {
	return xxx_messageInfo_SchemaType_SchemaColumn.Size(m)
}
func (m *SchemaType_SchemaColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_SchemaType_SchemaColumn.DiscardUnknown(m)
}

var xxx_messageInfo_SchemaType_SchemaColumn proto.InternalMessageInfo

func (m *SchemaType_SchemaColumn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SchemaType_SchemaColumn) GetType() SchemaType_SchemaColumn_SchemaColumnType {
	if m != nil {
		return m.Type
	}
	return SchemaType_SchemaColumn_INTEGER
}

type StructuredDatasetType struct {
	// A list of ordered columns this schema comprises of.
	Columns []*StructuredDatasetType_DatasetColumn `protobuf:"bytes,1,rep,name=columns,proto3" json:"columns,omitempty"`
	// This is the storage format, the format of the bits at rest
	// parquet, feather, csv, etc.
	// For two types to be compatible, the format will need to be an exact match.
	Format string `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	// This is a string representing the type that the bytes in external_schema_bytes are formatted in.
	// This is an optional field that will not be used for type checking.
	ExternalSchemaType string `protobuf:"bytes,3,opt,name=external_schema_type,json=externalSchemaType,proto3" json:"external_schema_type,omitempty"`
	// The serialized bytes of a third-party schema library like Arrow.
	// This is an optional field that will not be used for type checking.
	ExternalSchemaBytes  []byte   `protobuf:"bytes,4,opt,name=external_schema_bytes,json=externalSchemaBytes,proto3" json:"external_schema_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StructuredDatasetType) Reset()         { *m = StructuredDatasetType{} }
func (m *StructuredDatasetType) String() string { return proto.CompactTextString(m) }
func (*StructuredDatasetType) ProtoMessage()    {}
func (*StructuredDatasetType) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{1}
}

func (m *StructuredDatasetType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredDatasetType.Unmarshal(m, b)
}
func (m *StructuredDatasetType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredDatasetType.Marshal(b, m, deterministic)
}
func (m *StructuredDatasetType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredDatasetType.Merge(m, src)
}
func (m *StructuredDatasetType) XXX_Size() int {
	return xxx_messageInfo_StructuredDatasetType.Size(m)
}
func (m *StructuredDatasetType) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredDatasetType.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredDatasetType proto.InternalMessageInfo

func (m *StructuredDatasetType) GetColumns() []*StructuredDatasetType_DatasetColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *StructuredDatasetType) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *StructuredDatasetType) GetExternalSchemaType() string {
	if m != nil {
		return m.ExternalSchemaType
	}
	return ""
}

func (m *StructuredDatasetType) GetExternalSchemaBytes() []byte {
	if m != nil {
		return m.ExternalSchemaBytes
	}
	return nil
}

type StructuredDatasetType_DatasetColumn struct {
	// A unique name within the schema type for the column.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The column type.
	LiteralType          *LiteralType `protobuf:"bytes,2,opt,name=literal_type,json=literalType,proto3" json:"literal_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *StructuredDatasetType_DatasetColumn) Reset()         { *m = StructuredDatasetType_DatasetColumn{} }
func (m *StructuredDatasetType_DatasetColumn) String() string { return proto.CompactTextString(m) }
func (*StructuredDatasetType_DatasetColumn) ProtoMessage()    {}
func (*StructuredDatasetType_DatasetColumn) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{1, 0}
}

func (m *StructuredDatasetType_DatasetColumn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StructuredDatasetType_DatasetColumn.Unmarshal(m, b)
}
func (m *StructuredDatasetType_DatasetColumn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StructuredDatasetType_DatasetColumn.Marshal(b, m, deterministic)
}
func (m *StructuredDatasetType_DatasetColumn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StructuredDatasetType_DatasetColumn.Merge(m, src)
}
func (m *StructuredDatasetType_DatasetColumn) XXX_Size() int {
	return xxx_messageInfo_StructuredDatasetType_DatasetColumn.Size(m)
}
func (m *StructuredDatasetType_DatasetColumn) XXX_DiscardUnknown() {
	xxx_messageInfo_StructuredDatasetType_DatasetColumn.DiscardUnknown(m)
}

var xxx_messageInfo_StructuredDatasetType_DatasetColumn proto.InternalMessageInfo

func (m *StructuredDatasetType_DatasetColumn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StructuredDatasetType_DatasetColumn) GetLiteralType() *LiteralType {
	if m != nil {
		return m.LiteralType
	}
	return nil
}

// Defines type behavior for blob objects
type BlobType struct {
	// Format can be a free form string understood by SDK/UI etc like
	// csv, parquet etc
	Format               string                      `protobuf:"bytes,1,opt,name=format,proto3" json:"format,omitempty"`
	Dimensionality       BlobType_BlobDimensionality `protobuf:"varint,2,opt,name=dimensionality,proto3,enum=flyteidl.core.BlobType_BlobDimensionality" json:"dimensionality,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *BlobType) Reset()         { *m = BlobType{} }
func (m *BlobType) String() string { return proto.CompactTextString(m) }
func (*BlobType) ProtoMessage()    {}
func (*BlobType) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{2}
}

func (m *BlobType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlobType.Unmarshal(m, b)
}
func (m *BlobType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlobType.Marshal(b, m, deterministic)
}
func (m *BlobType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlobType.Merge(m, src)
}
func (m *BlobType) XXX_Size() int {
	return xxx_messageInfo_BlobType.Size(m)
}
func (m *BlobType) XXX_DiscardUnknown() {
	xxx_messageInfo_BlobType.DiscardUnknown(m)
}

var xxx_messageInfo_BlobType proto.InternalMessageInfo

func (m *BlobType) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *BlobType) GetDimensionality() BlobType_BlobDimensionality {
	if m != nil {
		return m.Dimensionality
	}
	return BlobType_SINGLE
}

// Enables declaring enum types, with predefined string values
// For len(values) > 0, the first value in the ordered list is regarded as the default value. If you wish
// To provide no defaults, make the first value as undefined.
type EnumType struct {
	// Predefined set of enum values.
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumType) Reset()         { *m = EnumType{} }
func (m *EnumType) String() string { return proto.CompactTextString(m) }
func (*EnumType) ProtoMessage()    {}
func (*EnumType) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{3}
}

func (m *EnumType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumType.Unmarshal(m, b)
}
func (m *EnumType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumType.Marshal(b, m, deterministic)
}
func (m *EnumType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumType.Merge(m, src)
}
func (m *EnumType) XXX_Size() int {
	return xxx_messageInfo_EnumType.Size(m)
}
func (m *EnumType) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumType.DiscardUnknown(m)
}

var xxx_messageInfo_EnumType proto.InternalMessageInfo

func (m *EnumType) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

// Defines a strong type to allow type checking between interfaces.
type LiteralType struct {
	// Types that are valid to be assigned to Type:
	//	*LiteralType_Simple
	//	*LiteralType_Schema
	//	*LiteralType_CollectionType
	//	*LiteralType_MapValueType
	//	*LiteralType_Blob
	//	*LiteralType_EnumType
	//	*LiteralType_StructuredDatasetType
	Type isLiteralType_Type `protobuf_oneof:"type"`
	// This field contains type metadata that is descriptive of the type, but is NOT considered in type-checking.  This might be used by
	// consumers to identify special behavior or display extended information for the type.
	Metadata             *_struct.Struct `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *LiteralType) Reset()         { *m = LiteralType{} }
func (m *LiteralType) String() string { return proto.CompactTextString(m) }
func (*LiteralType) ProtoMessage()    {}
func (*LiteralType) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{4}
}

func (m *LiteralType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiteralType.Unmarshal(m, b)
}
func (m *LiteralType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiteralType.Marshal(b, m, deterministic)
}
func (m *LiteralType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiteralType.Merge(m, src)
}
func (m *LiteralType) XXX_Size() int {
	return xxx_messageInfo_LiteralType.Size(m)
}
func (m *LiteralType) XXX_DiscardUnknown() {
	xxx_messageInfo_LiteralType.DiscardUnknown(m)
}

var xxx_messageInfo_LiteralType proto.InternalMessageInfo

type isLiteralType_Type interface {
	isLiteralType_Type()
}

type LiteralType_Simple struct {
	Simple SimpleType `protobuf:"varint,1,opt,name=simple,proto3,enum=flyteidl.core.SimpleType,oneof"`
}

type LiteralType_Schema struct {
	Schema *SchemaType `protobuf:"bytes,2,opt,name=schema,proto3,oneof"`
}

type LiteralType_CollectionType struct {
	CollectionType *LiteralType `protobuf:"bytes,3,opt,name=collection_type,json=collectionType,proto3,oneof"`
}

type LiteralType_MapValueType struct {
	MapValueType *LiteralType `protobuf:"bytes,4,opt,name=map_value_type,json=mapValueType,proto3,oneof"`
}

type LiteralType_Blob struct {
	Blob *BlobType `protobuf:"bytes,5,opt,name=blob,proto3,oneof"`
}

type LiteralType_EnumType struct {
	EnumType *EnumType `protobuf:"bytes,7,opt,name=enum_type,json=enumType,proto3,oneof"`
}

type LiteralType_StructuredDatasetType struct {
	StructuredDatasetType *StructuredDatasetType `protobuf:"bytes,8,opt,name=structured_dataset_type,json=structuredDatasetType,proto3,oneof"`
}

func (*LiteralType_Simple) isLiteralType_Type() {}

func (*LiteralType_Schema) isLiteralType_Type() {}

func (*LiteralType_CollectionType) isLiteralType_Type() {}

func (*LiteralType_MapValueType) isLiteralType_Type() {}

func (*LiteralType_Blob) isLiteralType_Type() {}

func (*LiteralType_EnumType) isLiteralType_Type() {}

func (*LiteralType_StructuredDatasetType) isLiteralType_Type() {}

func (m *LiteralType) GetType() isLiteralType_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *LiteralType) GetSimple() SimpleType {
	if x, ok := m.GetType().(*LiteralType_Simple); ok {
		return x.Simple
	}
	return SimpleType_NONE
}

func (m *LiteralType) GetSchema() *SchemaType {
	if x, ok := m.GetType().(*LiteralType_Schema); ok {
		return x.Schema
	}
	return nil
}

func (m *LiteralType) GetCollectionType() *LiteralType {
	if x, ok := m.GetType().(*LiteralType_CollectionType); ok {
		return x.CollectionType
	}
	return nil
}

func (m *LiteralType) GetMapValueType() *LiteralType {
	if x, ok := m.GetType().(*LiteralType_MapValueType); ok {
		return x.MapValueType
	}
	return nil
}

func (m *LiteralType) GetBlob() *BlobType {
	if x, ok := m.GetType().(*LiteralType_Blob); ok {
		return x.Blob
	}
	return nil
}

func (m *LiteralType) GetEnumType() *EnumType {
	if x, ok := m.GetType().(*LiteralType_EnumType); ok {
		return x.EnumType
	}
	return nil
}

func (m *LiteralType) GetStructuredDatasetType() *StructuredDatasetType {
	if x, ok := m.GetType().(*LiteralType_StructuredDatasetType); ok {
		return x.StructuredDatasetType
	}
	return nil
}

func (m *LiteralType) GetMetadata() *_struct.Struct {
	if m != nil {
		return m.Metadata
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*LiteralType) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*LiteralType_Simple)(nil),
		(*LiteralType_Schema)(nil),
		(*LiteralType_CollectionType)(nil),
		(*LiteralType_MapValueType)(nil),
		(*LiteralType_Blob)(nil),
		(*LiteralType_EnumType)(nil),
		(*LiteralType_StructuredDatasetType)(nil),
	}
}

// A reference to an output produced by a node. The type can be retrieved -and validated- from
// the underlying interface of the node.
type OutputReference struct {
	// Node id must exist at the graph layer.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Variable name must refer to an output variable for the node.
	Var                  string   `protobuf:"bytes,2,opt,name=var,proto3" json:"var,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputReference) Reset()         { *m = OutputReference{} }
func (m *OutputReference) String() string { return proto.CompactTextString(m) }
func (*OutputReference) ProtoMessage()    {}
func (*OutputReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{5}
}

func (m *OutputReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputReference.Unmarshal(m, b)
}
func (m *OutputReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputReference.Marshal(b, m, deterministic)
}
func (m *OutputReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputReference.Merge(m, src)
}
func (m *OutputReference) XXX_Size() int {
	return xxx_messageInfo_OutputReference.Size(m)
}
func (m *OutputReference) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputReference.DiscardUnknown(m)
}

var xxx_messageInfo_OutputReference proto.InternalMessageInfo

func (m *OutputReference) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *OutputReference) GetVar() string {
	if m != nil {
		return m.Var
	}
	return ""
}

// Represents an error thrown from a node.
type Error struct {
	// The node id that threw the error.
	FailedNodeId string `protobuf:"bytes,1,opt,name=failed_node_id,json=failedNodeId,proto3" json:"failed_node_id,omitempty"`
	// Error message thrown.
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_51e8add38f4caaed, []int{6}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetFailedNodeId() string {
	if m != nil {
		return m.FailedNodeId
	}
	return ""
}

func (m *Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("flyteidl.core.SimpleType", SimpleType_name, SimpleType_value)
	proto.RegisterEnum("flyteidl.core.SchemaType_SchemaColumn_SchemaColumnType", SchemaType_SchemaColumn_SchemaColumnType_name, SchemaType_SchemaColumn_SchemaColumnType_value)
	proto.RegisterEnum("flyteidl.core.BlobType_BlobDimensionality", BlobType_BlobDimensionality_name, BlobType_BlobDimensionality_value)
	proto.RegisterType((*SchemaType)(nil), "flyteidl.core.SchemaType")
	proto.RegisterType((*SchemaType_SchemaColumn)(nil), "flyteidl.core.SchemaType.SchemaColumn")
	proto.RegisterType((*StructuredDatasetType)(nil), "flyteidl.core.StructuredDatasetType")
	proto.RegisterType((*StructuredDatasetType_DatasetColumn)(nil), "flyteidl.core.StructuredDatasetType.DatasetColumn")
	proto.RegisterType((*BlobType)(nil), "flyteidl.core.BlobType")
	proto.RegisterType((*EnumType)(nil), "flyteidl.core.EnumType")
	proto.RegisterType((*LiteralType)(nil), "flyteidl.core.LiteralType")
	proto.RegisterType((*OutputReference)(nil), "flyteidl.core.OutputReference")
	proto.RegisterType((*Error)(nil), "flyteidl.core.Error")
}

func init() { proto.RegisterFile("flyteidl/core/types.proto", fileDescriptor_51e8add38f4caaed) }

var fileDescriptor_51e8add38f4caaed = []byte{
	// 812 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x13, 0xe7, 0xef, 0x24, 0xcd, 0x5a, 0x03, 0xa5, 0x69, 0xc5, 0x45, 0x65, 0xad, 0x50,
	0xb5, 0xd2, 0xc6, 0x28, 0x45, 0xcb, 0x0d, 0x48, 0x24, 0x5b, 0xd3, 0x44, 0x64, 0x1d, 0x34, 0x75,
	0x91, 0xe0, 0x02, 0xcb, 0x3f, 0x93, 0xac, 0xa5, 0xb1, 0xc7, 0xb2, 0xc7, 0x2b, 0xf2, 0x02, 0xfb,
	0x1e, 0x5c, 0xf1, 0x28, 0xdc, 0xf2, 0x48, 0xc8, 0x63, 0xbb, 0x8e, 0x4d, 0x16, 0xf5, 0x6e, 0x8e,
	0xcf, 0xf7, 0x7d, 0x73, 0xfe, 0xe6, 0x18, 0x2e, 0x77, 0xf4, 0xc0, 0x89, 0xef, 0x51, 0xcd, 0x65,
	0x31, 0xd1, 0xf8, 0x21, 0x22, 0xc9, 0x2c, 0x8a, 0x19, 0x67, 0xe8, 0xac, 0x74, 0xcd, 0x32, 0xd7,
	0xd5, 0x97, 0x7b, 0xc6, 0xf6, 0x94, 0x68, 0xc2, 0xe9, 0xa4, 0x3b, 0x2d, 0xe1, 0x71, 0xea, 0xf2,
	0x1c, 0xac, 0xfe, 0xd9, 0x06, 0x78, 0x70, 0xdf, 0x93, 0xc0, 0x36, 0x0f, 0x11, 0x41, 0x3f, 0x40,
	0xdf, 0x65, 0x34, 0x0d, 0xc2, 0x64, 0xda, 0xb9, 0xee, 0xdc, 0x8c, 0xe6, 0x5f, 0xcd, 0x6a, 0x6a,
	0xb3, 0x0a, 0x5b, 0x1c, 0xdf, 0x0a, 0x38, 0x2e, 0x69, 0x57, 0xff, 0x48, 0x30, 0x3e, 0xf6, 0x20,
	0x04, 0x72, 0x68, 0x07, 0x64, 0x2a, 0x5d, 0x4b, 0x37, 0x43, 0x2c, 0xce, 0xe8, 0x27, 0x90, 0xb3,
	0x88, 0xa7, 0xed, 0x6b, 0xe9, 0x66, 0x32, 0xff, 0xf6, 0x79, 0x77, 0xd4, 0x8c, 0xcc, 0x8b, 0x85,
	0x88, 0x6a, 0x81, 0xd2, 0xf4, 0xa0, 0x11, 0xf4, 0xd7, 0x86, 0xa9, 0xdf, 0xeb, 0x58, 0x69, 0xa1,
	0x21, 0x74, 0x7f, 0xdc, 0x6c, 0x17, 0xa6, 0x22, 0x21, 0x80, 0xde, 0x83, 0x89, 0xd7, 0xc6, 0xbd,
	0xd2, 0xce, 0x30, 0xcb, 0xed, 0x76, 0xa3, 0x2f, 0x0c, 0xa5, 0x83, 0xc6, 0x30, 0xb8, 0x5b, 0x98,
	0xba, 0xb9, 0x7e, 0xa7, 0x2b, 0xb2, 0xb0, 0x1e, 0xf1, 0xc2, 0x5c, 0x6f, 0x0d, 0xa5, 0xab, 0xfe,
	0xdd, 0x86, 0xf3, 0x07, 0x51, 0xb4, 0x34, 0x26, 0xde, 0x9d, 0xcd, 0xed, 0x84, 0x70, 0x71, 0xcd,
	0xa6, 0x2a, 0x97, 0x24, 0xca, 0x35, 0x6f, 0xa6, 0x72, 0x8a, 0x36, 0x2b, 0xce, 0x8d, 0xd2, 0xa1,
	0x2f, 0xa0, 0xb7, 0x63, 0x71, 0x60, 0x73, 0x51, 0x97, 0x21, 0x2e, 0x2c, 0xf4, 0x35, 0x7c, 0x4e,
	0xfe, 0xe0, 0x24, 0x0e, 0x6d, 0x6a, 0x25, 0x22, 0x53, 0x4b, 0x54, 0xaf, 0x23, 0x50, 0xa8, 0xf4,
	0x1d, 0xb5, 0x71, 0x0e, 0xe7, 0x4d, 0x86, 0x73, 0xe0, 0x24, 0x99, 0xca, 0xd7, 0xd2, 0xcd, 0x18,
	0x7f, 0x56, 0xa7, 0x2c, 0x33, 0xd7, 0x95, 0x03, 0x67, 0xb5, 0xb8, 0x4e, 0x36, 0xee, 0x7b, 0x18,
	0x53, 0x9f, 0x93, 0xd8, 0xa6, 0xd6, 0x53, 0x03, 0x47, 0xf3, 0xab, 0x46, 0xd6, 0x9b, 0x1c, 0x22,
	0x7a, 0x34, 0xa2, 0x95, 0xa1, 0xfe, 0x25, 0xc1, 0x60, 0x49, 0x99, 0x23, 0x82, 0xac, 0xd2, 0x95,
	0x6a, 0xe9, 0x62, 0x98, 0x78, 0x7e, 0x40, 0xc2, 0xc4, 0x67, 0xa1, 0x4d, 0x7d, 0x7e, 0x28, 0xc6,
	0xe4, 0x55, 0xe3, 0x96, 0x52, 0x48, 0x1c, 0xee, 0x6a, 0x0c, 0xdc, 0x50, 0x50, 0x35, 0x40, 0xff,
	0x45, 0x89, 0x69, 0x58, 0x1b, 0xf7, 0x1b, 0x5d, 0x69, 0xa1, 0x33, 0x18, 0xbe, 0x7b, 0xdc, 0x98,
	0xeb, 0x9f, 0x17, 0xd8, 0x54, 0x24, 0x55, 0x85, 0x81, 0x1e, 0xa6, 0x41, 0x19, 0xe8, 0x07, 0x9b,
	0xa6, 0x24, 0x6f, 0xf2, 0x10, 0x17, 0x96, 0xfa, 0x51, 0x86, 0xd1, 0x51, 0xaa, 0xe8, 0x16, 0x7a,
	0x89, 0x1f, 0x44, 0x34, 0x2f, 0xd9, 0x64, 0x7e, 0xd9, 0x1c, 0x06, 0xe1, 0xcc, 0xa0, 0xab, 0x16,
	0x2e, 0xa0, 0x82, 0x24, 0xba, 0x50, 0xd4, 0xf2, 0xf2, 0x93, 0x8f, 0x41, 0x90, 0x84, 0x85, 0x74,
	0x78, 0xe1, 0x32, 0x4a, 0x89, 0xcb, 0x7d, 0x16, 0x56, 0xc3, 0xf0, 0xbf, 0x9d, 0x58, 0xb5, 0xf0,
	0xa4, 0x22, 0x89, 0x80, 0x97, 0x30, 0x09, 0xec, 0xc8, 0x12, 0xe9, 0xe4, 0x2a, 0xf2, 0x33, 0x54,
	0xc6, 0x81, 0x1d, 0xfd, 0x92, 0x51, 0x84, 0xc6, 0x6b, 0x90, 0x1d, 0xca, 0x9c, 0x69, 0x57, 0x30,
	0x2f, 0x3e, 0xd1, 0xa3, 0x55, 0x0b, 0x0b, 0x18, 0x7a, 0x03, 0x43, 0x12, 0xa6, 0x41, 0x7e, 0x5b,
	0xff, 0x24, 0xa7, 0xac, 0xfb, 0xaa, 0x85, 0x07, 0xa4, 0xec, 0xc1, 0xef, 0x70, 0x91, 0x3c, 0xbd,
	0x25, 0xcb, 0xcb, 0x07, 0x35, 0x57, 0x19, 0x08, 0x95, 0x97, 0xcf, 0x79, 0x79, 0xab, 0x16, 0x3e,
	0x4f, 0x4e, 0xbe, 0xe4, 0x5b, 0x18, 0x04, 0x84, 0xdb, 0x99, 0xf0, 0xb4, 0x57, 0x84, 0x95, 0x2f,
	0xce, 0x59, 0xb9, 0x38, 0x0b, 0x49, 0xfc, 0x04, 0x5c, 0xf6, 0xf2, 0x35, 0xa6, 0x7e, 0x07, 0x2f,
	0xb6, 0x29, 0x8f, 0x52, 0x8e, 0xc9, 0x8e, 0xc4, 0x24, 0x74, 0x09, 0xba, 0x80, 0x7e, 0xc8, 0x3c,
	0x62, 0xf9, 0x5e, 0x39, 0xdd, 0x99, 0xb9, 0xf6, 0x90, 0x02, 0x9d, 0x0f, 0x76, 0x5c, 0xbc, 0xf0,
	0xec, 0xa8, 0xde, 0x43, 0x57, 0x8f, 0x63, 0x16, 0xa3, 0x97, 0x30, 0xd9, 0xd9, 0x3e, 0x25, 0x9e,
	0x55, 0xa7, 0x8e, 0xf3, 0xaf, 0x46, 0x2e, 0x30, 0x85, 0x7e, 0x40, 0x92, 0xc4, 0xde, 0x93, 0x42,
	0xa4, 0x34, 0x5f, 0x7d, 0x94, 0x00, 0xaa, 0x19, 0x43, 0x03, 0x90, 0x8d, 0xad, 0x91, 0xcd, 0xf6,
	0xd1, 0x36, 0x94, 0xaa, 0x6d, 0xd8, 0x3e, 0xda, 0x86, 0x9d, 0xe3, 0x6d, 0x28, 0xd7, 0xb6, 0x61,
	0xb7, 0xb6, 0x0d, 0x7b, 0x19, 0x69, 0xb9, 0x36, 0x16, 0xf8, 0x57, 0xa5, 0x9f, 0x69, 0xe9, 0x18,
	0x6f, 0xb1, 0x32, 0x28, 0xb4, 0x1e, 0xdf, 0x9a, 0xca, 0x70, 0xf9, 0xe6, 0xb7, 0x6f, 0xf6, 0x3e,
	0x7f, 0x9f, 0x3a, 0x33, 0x97, 0x05, 0x9a, 0xe8, 0x0b, 0x8b, 0xf7, 0xda, 0xd3, 0x2f, 0x6b, 0x4f,
	0x42, 0x2d, 0x72, 0x5e, 0xef, 0x99, 0x56, 0xfb, 0x8b, 0x39, 0x3d, 0x51, 0xea, 0xdb, 0x7f, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xfc, 0x32, 0x51, 0x68, 0xdd, 0x06, 0x00, 0x00,
}
