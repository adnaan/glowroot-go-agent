// Code generated by protoc-gen-go.
// source: Aggregate.proto
// DO NOT EDIT!

package org_glowroot_wire_api_model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// deprecated in 0.9.5
type OldAggregatesByType struct {
	TransactionType      string                     `protobuf:"bytes,1,opt,name=transaction_type,json=transactionType" json:"transaction_type,omitempty"`
	OverallAggregate     *Aggregate                 `protobuf:"bytes,2,opt,name=overall_aggregate,json=overallAggregate" json:"overall_aggregate,omitempty"`
	TransactionAggregate []*OldTransactionAggregate `protobuf:"bytes,3,rep,name=transaction_aggregate,json=transactionAggregate" json:"transaction_aggregate,omitempty"`
}

func (m *OldAggregatesByType) Reset()                    { *m = OldAggregatesByType{} }
func (m *OldAggregatesByType) String() string            { return proto.CompactTextString(m) }
func (*OldAggregatesByType) ProtoMessage()               {}
func (*OldAggregatesByType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *OldAggregatesByType) GetTransactionType() string {
	if m != nil {
		return m.TransactionType
	}
	return ""
}

func (m *OldAggregatesByType) GetOverallAggregate() *Aggregate {
	if m != nil {
		return m.OverallAggregate
	}
	return nil
}

func (m *OldAggregatesByType) GetTransactionAggregate() []*OldTransactionAggregate {
	if m != nil {
		return m.TransactionAggregate
	}
	return nil
}

// deprecated in 0.9.5
type OldTransactionAggregate struct {
	TransactionName string     `protobuf:"bytes,2,opt,name=transaction_name,json=transactionName" json:"transaction_name,omitempty"`
	Aggregate       *Aggregate `protobuf:"bytes,3,opt,name=aggregate" json:"aggregate,omitempty"`
}

func (m *OldTransactionAggregate) Reset()                    { *m = OldTransactionAggregate{} }
func (m *OldTransactionAggregate) String() string            { return proto.CompactTextString(m) }
func (*OldTransactionAggregate) ProtoMessage()               {}
func (*OldTransactionAggregate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *OldTransactionAggregate) GetTransactionName() string {
	if m != nil {
		return m.TransactionName
	}
	return ""
}

func (m *OldTransactionAggregate) GetAggregate() *Aggregate {
	if m != nil {
		return m.Aggregate
	}
	return nil
}

// aggregate uses double instead of int64 to avoid (unlikely) 292 year nanosecond rollover
type Aggregate struct {
	TotalDurationNanos     float64                `protobuf:"fixed64,1,opt,name=total_duration_nanos,json=totalDurationNanos" json:"total_duration_nanos,omitempty"`
	TransactionCount       int64                  `protobuf:"varint,2,opt,name=transaction_count,json=transactionCount" json:"transaction_count,omitempty"`
	ErrorCount             int64                  `protobuf:"varint,3,opt,name=error_count,json=errorCount" json:"error_count,omitempty"`
	AsyncTransactions      bool                   `protobuf:"varint,4,opt,name=async_transactions,json=asyncTransactions" json:"async_transactions,omitempty"`
	MainThreadRootTimer    []*Aggregate_Timer     `protobuf:"bytes,5,rep,name=main_thread_root_timer,json=mainThreadRootTimer" json:"main_thread_root_timer,omitempty"`
	AuxThreadRootTimer     []*Aggregate_Timer     `protobuf:"bytes,6,rep,name=aux_thread_root_timer,json=auxThreadRootTimer" json:"aux_thread_root_timer,omitempty"`
	AsyncTimer             []*Aggregate_Timer     `protobuf:"bytes,7,rep,name=async_timer,json=asyncTimer" json:"async_timer,omitempty"`
	MainThreadStats        *Aggregate_ThreadStats `protobuf:"bytes,8,opt,name=main_thread_stats,json=mainThreadStats" json:"main_thread_stats,omitempty"`
	AuxThreadStats         *Aggregate_ThreadStats `protobuf:"bytes,9,opt,name=aux_thread_stats,json=auxThreadStats" json:"aux_thread_stats,omitempty"`
	DurationNanosHistogram *Aggregate_Histogram   `protobuf:"bytes,10,opt,name=duration_nanos_histogram,json=durationNanosHistogram" json:"duration_nanos_histogram,omitempty"`
	// precision
	QueriesByType      []*Aggregate_QueriesByType      `protobuf:"bytes,11,rep,name=queries_by_type,json=queriesByType" json:"queries_by_type,omitempty"`
	ServiceCallsByType []*Aggregate_ServiceCallsByType `protobuf:"bytes,12,rep,name=service_calls_by_type,json=serviceCallsByType" json:"service_calls_by_type,omitempty"`
	MainThreadProfile  *Profile                        `protobuf:"bytes,13,opt,name=main_thread_profile,json=mainThreadProfile" json:"main_thread_profile,omitempty"`
	AuxThreadProfile   *Profile                        `protobuf:"bytes,14,opt,name=aux_thread_profile,json=auxThreadProfile" json:"aux_thread_profile,omitempty"`
}

func (m *Aggregate) Reset()                    { *m = Aggregate{} }
func (m *Aggregate) String() string            { return proto.CompactTextString(m) }
func (*Aggregate) ProtoMessage()               {}
func (*Aggregate) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Aggregate) GetTotalDurationNanos() float64 {
	if m != nil {
		return m.TotalDurationNanos
	}
	return 0
}

func (m *Aggregate) GetTransactionCount() int64 {
	if m != nil {
		return m.TransactionCount
	}
	return 0
}

func (m *Aggregate) GetErrorCount() int64 {
	if m != nil {
		return m.ErrorCount
	}
	return 0
}

func (m *Aggregate) GetAsyncTransactions() bool {
	if m != nil {
		return m.AsyncTransactions
	}
	return false
}

func (m *Aggregate) GetMainThreadRootTimer() []*Aggregate_Timer {
	if m != nil {
		return m.MainThreadRootTimer
	}
	return nil
}

func (m *Aggregate) GetAuxThreadRootTimer() []*Aggregate_Timer {
	if m != nil {
		return m.AuxThreadRootTimer
	}
	return nil
}

func (m *Aggregate) GetAsyncTimer() []*Aggregate_Timer {
	if m != nil {
		return m.AsyncTimer
	}
	return nil
}

func (m *Aggregate) GetMainThreadStats() *Aggregate_ThreadStats {
	if m != nil {
		return m.MainThreadStats
	}
	return nil
}

func (m *Aggregate) GetAuxThreadStats() *Aggregate_ThreadStats {
	if m != nil {
		return m.AuxThreadStats
	}
	return nil
}

func (m *Aggregate) GetDurationNanosHistogram() *Aggregate_Histogram {
	if m != nil {
		return m.DurationNanosHistogram
	}
	return nil
}

func (m *Aggregate) GetQueriesByType() []*Aggregate_QueriesByType {
	if m != nil {
		return m.QueriesByType
	}
	return nil
}

func (m *Aggregate) GetServiceCallsByType() []*Aggregate_ServiceCallsByType {
	if m != nil {
		return m.ServiceCallsByType
	}
	return nil
}

func (m *Aggregate) GetMainThreadProfile() *Profile {
	if m != nil {
		return m.MainThreadProfile
	}
	return nil
}

func (m *Aggregate) GetAuxThreadProfile() *Profile {
	if m != nil {
		return m.AuxThreadProfile
	}
	return nil
}

type Aggregate_Histogram struct {
	// for smaller numbers of transactions, the individual raw values are stored directly
	// these values are ordered for fast percentile calculation on retrieval
	OrderedRawValue []int64 `protobuf:"varint,1,rep,packed,name=ordered_raw_value,json=orderedRawValue" json:"ordered_raw_value,omitempty"`
	// for larger numbers of transactions, the histogram is tracked and encoded using HdrHistogram
	EncodedBytes []byte `protobuf:"bytes,2,opt,name=encoded_bytes,json=encodedBytes,proto3" json:"encoded_bytes,omitempty"`
}

func (m *Aggregate_Histogram) Reset()                    { *m = Aggregate_Histogram{} }
func (m *Aggregate_Histogram) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_Histogram) ProtoMessage()               {}
func (*Aggregate_Histogram) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

func (m *Aggregate_Histogram) GetOrderedRawValue() []int64 {
	if m != nil {
		return m.OrderedRawValue
	}
	return nil
}

func (m *Aggregate_Histogram) GetEncodedBytes() []byte {
	if m != nil {
		return m.EncodedBytes
	}
	return nil
}

type Aggregate_Timer struct {
	// name is null for synthetic root
	Name       string             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Extended   bool               `protobuf:"varint,2,opt,name=extended" json:"extended,omitempty"`
	TotalNanos float64            `protobuf:"fixed64,3,opt,name=total_nanos,json=totalNanos" json:"total_nanos,omitempty"`
	Count      int64              `protobuf:"varint,4,opt,name=count" json:"count,omitempty"`
	ChildTimer []*Aggregate_Timer `protobuf:"bytes,5,rep,name=child_timer,json=childTimer" json:"child_timer,omitempty"`
}

func (m *Aggregate_Timer) Reset()                    { *m = Aggregate_Timer{} }
func (m *Aggregate_Timer) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_Timer) ProtoMessage()               {}
func (*Aggregate_Timer) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 1} }

func (m *Aggregate_Timer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Aggregate_Timer) GetExtended() bool {
	if m != nil {
		return m.Extended
	}
	return false
}

func (m *Aggregate_Timer) GetTotalNanos() float64 {
	if m != nil {
		return m.TotalNanos
	}
	return 0
}

func (m *Aggregate_Timer) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Aggregate_Timer) GetChildTimer() []*Aggregate_Timer {
	if m != nil {
		return m.ChildTimer
	}
	return nil
}

type Aggregate_ThreadStats struct {
	TotalCpuNanos       *OptionalDouble `protobuf:"bytes,1,opt,name=total_cpu_nanos,json=totalCpuNanos" json:"total_cpu_nanos,omitempty"`
	TotalBlockedNanos   *OptionalDouble `protobuf:"bytes,2,opt,name=total_blocked_nanos,json=totalBlockedNanos" json:"total_blocked_nanos,omitempty"`
	TotalWaitedNanos    *OptionalDouble `protobuf:"bytes,3,opt,name=total_waited_nanos,json=totalWaitedNanos" json:"total_waited_nanos,omitempty"`
	TotalAllocatedBytes *OptionalDouble `protobuf:"bytes,4,opt,name=total_allocated_bytes,json=totalAllocatedBytes" json:"total_allocated_bytes,omitempty"`
}

func (m *Aggregate_ThreadStats) Reset()                    { *m = Aggregate_ThreadStats{} }
func (m *Aggregate_ThreadStats) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_ThreadStats) ProtoMessage()               {}
func (*Aggregate_ThreadStats) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 2} }

func (m *Aggregate_ThreadStats) GetTotalCpuNanos() *OptionalDouble {
	if m != nil {
		return m.TotalCpuNanos
	}
	return nil
}

func (m *Aggregate_ThreadStats) GetTotalBlockedNanos() *OptionalDouble {
	if m != nil {
		return m.TotalBlockedNanos
	}
	return nil
}

func (m *Aggregate_ThreadStats) GetTotalWaitedNanos() *OptionalDouble {
	if m != nil {
		return m.TotalWaitedNanos
	}
	return nil
}

func (m *Aggregate_ThreadStats) GetTotalAllocatedBytes() *OptionalDouble {
	if m != nil {
		return m.TotalAllocatedBytes
	}
	return nil
}

type Aggregate_QueriesByType struct {
	// e.g. "SQL", "CQL"
	Type  string             `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Query []*Aggregate_Query `protobuf:"bytes,2,rep,name=query" json:"query,omitempty"`
}

func (m *Aggregate_QueriesByType) Reset()                    { *m = Aggregate_QueriesByType{} }
func (m *Aggregate_QueriesByType) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_QueriesByType) ProtoMessage()               {}
func (*Aggregate_QueriesByType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 3} }

func (m *Aggregate_QueriesByType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Aggregate_QueriesByType) GetQuery() []*Aggregate_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type Aggregate_Query struct {
	FullText             string         `protobuf:"bytes,1,opt,name=full_text,json=fullText" json:"full_text,omitempty"`
	SharedQueryTextIndex int32          `protobuf:"varint,5,opt,name=shared_query_text_index,json=sharedQueryTextIndex" json:"shared_query_text_index,omitempty"`
	TotalDurationNanos   float64        `protobuf:"fixed64,2,opt,name=total_duration_nanos,json=totalDurationNanos" json:"total_duration_nanos,omitempty"`
	ExecutionCount       int64          `protobuf:"varint,3,opt,name=execution_count,json=executionCount" json:"execution_count,omitempty"`
	TotalRows            *OptionalInt64 `protobuf:"bytes,4,opt,name=total_rows,json=totalRows" json:"total_rows,omitempty"`
}

func (m *Aggregate_Query) Reset()                    { *m = Aggregate_Query{} }
func (m *Aggregate_Query) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_Query) ProtoMessage()               {}
func (*Aggregate_Query) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 4} }

func (m *Aggregate_Query) GetFullText() string {
	if m != nil {
		return m.FullText
	}
	return ""
}

func (m *Aggregate_Query) GetSharedQueryTextIndex() int32 {
	if m != nil {
		return m.SharedQueryTextIndex
	}
	return 0
}

func (m *Aggregate_Query) GetTotalDurationNanos() float64 {
	if m != nil {
		return m.TotalDurationNanos
	}
	return 0
}

func (m *Aggregate_Query) GetExecutionCount() int64 {
	if m != nil {
		return m.ExecutionCount
	}
	return 0
}

func (m *Aggregate_Query) GetTotalRows() *OptionalInt64 {
	if m != nil {
		return m.TotalRows
	}
	return nil
}

type Aggregate_ServiceCallsByType struct {
	// e.g. "HTTP"
	Type        string                   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	ServiceCall []*Aggregate_ServiceCall `protobuf:"bytes,2,rep,name=service_call,json=serviceCall" json:"service_call,omitempty"`
}

func (m *Aggregate_ServiceCallsByType) Reset()                    { *m = Aggregate_ServiceCallsByType{} }
func (m *Aggregate_ServiceCallsByType) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_ServiceCallsByType) ProtoMessage()               {}
func (*Aggregate_ServiceCallsByType) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 5} }

func (m *Aggregate_ServiceCallsByType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Aggregate_ServiceCallsByType) GetServiceCall() []*Aggregate_ServiceCall {
	if m != nil {
		return m.ServiceCall
	}
	return nil
}

type Aggregate_ServiceCall struct {
	Text               string  `protobuf:"bytes,1,opt,name=text" json:"text,omitempty"`
	TotalDurationNanos float64 `protobuf:"fixed64,2,opt,name=total_duration_nanos,json=totalDurationNanos" json:"total_duration_nanos,omitempty"`
	ExecutionCount     int64   `protobuf:"varint,3,opt,name=execution_count,json=executionCount" json:"execution_count,omitempty"`
}

func (m *Aggregate_ServiceCall) Reset()                    { *m = Aggregate_ServiceCall{} }
func (m *Aggregate_ServiceCall) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_ServiceCall) ProtoMessage()               {}
func (*Aggregate_ServiceCall) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 6} }

func (m *Aggregate_ServiceCall) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Aggregate_ServiceCall) GetTotalDurationNanos() float64 {
	if m != nil {
		return m.TotalDurationNanos
	}
	return 0
}

func (m *Aggregate_ServiceCall) GetExecutionCount() int64 {
	if m != nil {
		return m.ExecutionCount
	}
	return 0
}

type Aggregate_SharedQueryText struct {
	FullText      string `protobuf:"bytes,1,opt,name=full_text,json=fullText" json:"full_text,omitempty"`
	TruncatedText string `protobuf:"bytes,2,opt,name=truncated_text,json=truncatedText" json:"truncated_text,omitempty"`
	FullTextSha1  string `protobuf:"bytes,3,opt,name=full_text_sha1,json=fullTextSha1" json:"full_text_sha1,omitempty"`
}

func (m *Aggregate_SharedQueryText) Reset()                    { *m = Aggregate_SharedQueryText{} }
func (m *Aggregate_SharedQueryText) String() string            { return proto.CompactTextString(m) }
func (*Aggregate_SharedQueryText) ProtoMessage()               {}
func (*Aggregate_SharedQueryText) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 7} }

func (m *Aggregate_SharedQueryText) GetFullText() string {
	if m != nil {
		return m.FullText
	}
	return ""
}

func (m *Aggregate_SharedQueryText) GetTruncatedText() string {
	if m != nil {
		return m.TruncatedText
	}
	return ""
}

func (m *Aggregate_SharedQueryText) GetFullTextSha1() string {
	if m != nil {
		return m.FullTextSha1
	}
	return ""
}

func init() {
	proto.RegisterType((*OldAggregatesByType)(nil), "org_glowroot_wire_api_model.OldAggregatesByType")
	proto.RegisterType((*OldTransactionAggregate)(nil), "org_glowroot_wire_api_model.OldTransactionAggregate")
	proto.RegisterType((*Aggregate)(nil), "org_glowroot_wire_api_model.Aggregate")
	proto.RegisterType((*Aggregate_Histogram)(nil), "org_glowroot_wire_api_model.Aggregate.Histogram")
	proto.RegisterType((*Aggregate_Timer)(nil), "org_glowroot_wire_api_model.Aggregate.Timer")
	proto.RegisterType((*Aggregate_ThreadStats)(nil), "org_glowroot_wire_api_model.Aggregate.ThreadStats")
	proto.RegisterType((*Aggregate_QueriesByType)(nil), "org_glowroot_wire_api_model.Aggregate.QueriesByType")
	proto.RegisterType((*Aggregate_Query)(nil), "org_glowroot_wire_api_model.Aggregate.Query")
	proto.RegisterType((*Aggregate_ServiceCallsByType)(nil), "org_glowroot_wire_api_model.Aggregate.ServiceCallsByType")
	proto.RegisterType((*Aggregate_ServiceCall)(nil), "org_glowroot_wire_api_model.Aggregate.ServiceCall")
	proto.RegisterType((*Aggregate_SharedQueryText)(nil), "org_glowroot_wire_api_model.Aggregate.SharedQueryText")
}

func init() { proto.RegisterFile("Aggregate.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1029 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x96, 0x41, 0x6f, 0x23, 0x35,
	0x14, 0xc7, 0x95, 0xa4, 0x59, 0x9a, 0x97, 0xa4, 0x69, 0xdc, 0x76, 0x37, 0xca, 0x0a, 0x51, 0x95,
	0x05, 0xc2, 0x2e, 0x44, 0xbb, 0x65, 0x41, 0xe2, 0xb8, 0x6d, 0x0f, 0xf4, 0xc0, 0x2e, 0x38, 0x01,
	0x84, 0xa8, 0xb0, 0xdc, 0x19, 0x6f, 0x32, 0xe0, 0x8c, 0xb3, 0x1e, 0x4f, 0x93, 0x08, 0x24, 0xe0,
	0xca, 0x77, 0xe2, 0x4b, 0x71, 0xe5, 0x82, 0xfc, 0x3c, 0x99, 0x71, 0x77, 0xdb, 0x2a, 0x29, 0xe2,
	0x36, 0xf3, 0xec, 0xf7, 0x7b, 0xe3, 0xff, 0xb3, 0xff, 0x1e, 0x68, 0x3d, 0x1b, 0x8d, 0xb4, 0x18,
	0x71, 0x23, 0xfa, 0x53, 0xad, 0x8c, 0x22, 0xf7, 0x95, 0x1e, 0xb1, 0x91, 0x54, 0x33, 0xad, 0x94,
	0x61, 0xb3, 0x48, 0x0b, 0xc6, 0xa7, 0x11, 0x9b, 0xa8, 0x50, 0xc8, 0x6e, 0xf3, 0x2b, 0xad, 0x5e,
	0x46, 0x32, 0x9b, 0xdb, 0x6d, 0x1c, 0xab, 0xc9, 0x44, 0xc5, 0xee, 0xed, 0xe0, 0xf7, 0x32, 0xec,
	0xbc, 0x90, 0x61, 0x0e, 0x4c, 0x8e, 0x16, 0xc3, 0xc5, 0x54, 0x90, 0x0f, 0x61, 0xdb, 0x68, 0x1e,
	0x27, 0x3c, 0x30, 0x91, 0x8a, 0x99, 0x59, 0x4c, 0x45, 0xa7, 0xb4, 0x5f, 0xea, 0xd5, 0x68, 0xcb,
	0x8b, 0xe3, 0xd4, 0x01, 0xb4, 0xd5, 0x85, 0xd0, 0x5c, 0x4a, 0xc6, 0x97, 0x98, 0x4e, 0x79, 0xbf,
	0xd4, 0xab, 0x1f, 0xbe, 0xdf, 0xbf, 0xe1, 0xc3, 0xfa, 0x79, 0x51, 0xba, 0x9d, 0x01, 0xf2, 0x08,
	0x89, 0x60, 0xcf, 0xaf, 0x5f, 0x80, 0x2b, 0xfb, 0x95, 0x5e, 0xfd, 0xf0, 0xe9, 0x8d, 0xe0, 0x17,
	0x32, 0x1c, 0x16, 0xc9, 0x45, 0x99, 0x5d, 0x73, 0x45, 0xf4, 0xe0, 0xcf, 0x12, 0xdc, 0xbb, 0x26,
	0xe3, 0x75, 0x19, 0x62, 0x3e, 0x71, 0x4b, 0xbb, 0x2c, 0xc3, 0x73, 0x3e, 0x11, 0xe4, 0x04, 0x6a,
	0xfe, 0x57, 0xae, 0xb3, 0xfc, 0x22, 0xf1, 0xe0, 0xef, 0x1d, 0xa8, 0x15, 0xe5, 0x1f, 0xc3, 0xae,
	0x51, 0x86, 0x4b, 0x16, 0xa6, 0x9a, 0x67, 0x5f, 0x10, 0xab, 0x04, 0x3b, 0x51, 0xa2, 0x04, 0xc7,
	0x4e, 0xb2, 0xa1, 0xe7, 0x76, 0x84, 0x3c, 0x82, 0xb6, 0xff, 0xc1, 0x81, 0x4a, 0x63, 0x83, 0x5f,
	0x5c, 0xa1, 0xfe, 0x4a, 0x8e, 0x6d, 0x9c, 0xbc, 0x03, 0x75, 0xa1, 0xb5, 0xd2, 0xd9, 0xb4, 0x0a,
	0x4e, 0x03, 0x0c, 0xb9, 0x09, 0x1f, 0x03, 0xe1, 0xc9, 0x22, 0x0e, 0x98, 0x97, 0x9a, 0x74, 0x36,
	0xf6, 0x4b, 0xbd, 0x4d, 0xda, 0xc6, 0x11, 0x4f, 0xb5, 0x84, 0x70, 0xb8, 0x3b, 0xe1, 0x51, 0xcc,
	0xcc, 0x58, 0x0b, 0x1e, 0x32, 0x5c, 0xb4, 0x89, 0x26, 0x42, 0x77, 0xaa, 0xd8, 0xb5, 0x8f, 0x56,
	0xd3, 0xa3, 0x3f, 0xb4, 0x39, 0x74, 0xc7, 0xb2, 0x86, 0x88, 0xa2, 0x4a, 0x19, 0x0c, 0x12, 0x06,
	0x7b, 0x3c, 0x9d, 0x5f, 0x51, 0xe1, 0xce, 0x2d, 0x2a, 0x10, 0x9e, 0xce, 0x5f, 0x2f, 0xf0, 0x25,
	0xd4, 0xb3, 0x25, 0x23, 0xf6, 0xad, 0x5b, 0x60, 0xc1, 0x29, 0x83, 0xb8, 0x1f, 0xa1, 0xed, 0x4b,
	0x92, 0x18, 0x6e, 0x92, 0xce, 0x26, 0xee, 0x8e, 0xc3, 0x55, 0xa1, 0x98, 0x3a, 0xb0, 0x99, 0xb4,
	0x55, 0x68, 0x82, 0x01, 0x72, 0x06, 0xdb, 0x9e, 0x1e, 0x0e, 0x5f, 0xbb, 0x35, 0x7e, 0x2b, 0x17,
	0xc4, 0xd1, 0x7f, 0x82, 0xce, 0xe5, 0x9d, 0xc7, 0xc6, 0x51, 0x62, 0xd4, 0x48, 0xf3, 0x49, 0x07,
	0xb0, 0xca, 0xe3, 0x15, 0xab, 0x7c, 0xb1, 0xcc, 0xa3, 0x77, 0x43, 0x7f, 0xc3, 0xe6, 0x71, 0x72,
	0x06, 0xad, 0x57, 0xa9, 0xd0, 0x91, 0x48, 0xd8, 0xf9, 0xc2, 0x19, 0x4e, 0x7d, 0x85, 0xb3, 0x5e,
	0x94, 0xf8, 0xda, 0x65, 0x3b, 0x03, 0xa3, 0xcd, 0x57, 0xfe, 0x2b, 0x91, 0xb0, 0x97, 0x08, 0x7d,
	0x11, 0x05, 0x82, 0x05, 0x5c, 0xca, 0xa2, 0x46, 0x03, 0x6b, 0x7c, 0xbe, 0x62, 0x8d, 0x81, 0x63,
	0x1c, 0x5b, 0x44, 0x56, 0x88, 0x24, 0x6f, 0xc4, 0xc8, 0x10, 0x76, 0xfc, 0xae, 0x4f, 0x9d, 0x01,
	0x77, 0x9a, 0x28, 0xd9, 0x83, 0x1b, 0x6b, 0x65, 0x66, 0x4d, 0xdb, 0x45, 0xa7, 0xb3, 0x10, 0xa1,
	0x40, 0xbc, 0x5e, 0x2f, 0xa1, 0x5b, 0x6b, 0x40, 0xb7, 0xf3, 0xfe, 0x66, 0x91, 0xee, 0x19, 0xd4,
	0x8a, 0x16, 0x3c, 0x84, 0xb6, 0xd2, 0xa1, 0xd0, 0x22, 0x64, 0x9a, 0xcf, 0xd8, 0x05, 0x97, 0xa9,
	0x75, 0xfd, 0x4a, 0xaf, 0x42, 0x5b, 0xd9, 0x00, 0xe5, 0xb3, 0x6f, 0x6d, 0x98, 0xbc, 0x0b, 0x4d,
	0x11, 0x07, 0x2a, 0x14, 0x21, 0x3b, 0x5f, 0x18, 0x91, 0xa0, 0xc9, 0x34, 0x68, 0x23, 0x0b, 0x1e,
	0xd9, 0x58, 0xf7, 0xaf, 0x12, 0x54, 0xdd, 0x39, 0x20, 0xb0, 0x81, 0xe6, 0xe9, 0xee, 0x10, 0x7c,
	0x26, 0x5d, 0xd8, 0x14, 0x73, 0x23, 0xe2, 0x50, 0x84, 0x98, 0xbd, 0x49, 0xf3, 0x77, 0x6b, 0x4d,
	0xce, 0xf9, 0x9c, 0xe1, 0x55, 0xd0, 0xf0, 0x00, 0x43, 0xce, 0xe8, 0x76, 0xa1, 0xea, 0x5c, 0x6b,
	0x03, 0x5d, 0xcb, 0xbd, 0xd8, 0xd3, 0x1b, 0x8c, 0x23, 0x19, 0xfe, 0x07, 0xdb, 0x01, 0x04, 0xe0,
	0x73, 0xf7, 0x9f, 0x32, 0xd4, 0xfd, 0xf3, 0x30, 0x80, 0x96, 0xfb, 0xaa, 0x60, 0x9a, 0x7a, 0x56,
	0x5c, 0x3f, 0x7c, 0x74, 0xf3, 0x7d, 0x34, 0xb5, 0xfb, 0x9d, 0xcb, 0x13, 0x95, 0x9e, 0x4b, 0x41,
	0x9b, 0xc8, 0x38, 0x9e, 0xa6, 0x6e, 0x25, 0x3f, 0xc0, 0x8e, 0x83, 0x9e, 0x4b, 0x15, 0xfc, 0x2c,
	0xc2, 0x0c, 0x5c, 0x5e, 0x1f, 0xdc, 0x46, 0xce, 0x91, 0xc3, 0x38, 0xf8, 0xf7, 0xe0, 0x6e, 0x09,
	0x36, 0xe3, 0x91, 0xc9, 0xd9, 0x95, 0xf5, 0xd9, 0xdb, 0x88, 0xf9, 0x0e, 0x29, 0x0e, 0xcd, 0x60,
	0xcf, 0xa1, 0xb9, 0x94, 0x2a, 0xe0, 0x26, 0xdf, 0x09, 0x1b, 0xeb, 0xd3, 0x9d, 0x02, 0xcf, 0x96,
	0x20, 0xb7, 0x7b, 0x46, 0xd0, 0xbc, 0x74, 0xa6, 0xed, 0x26, 0xf2, 0x7e, 0x44, 0xf0, 0x99, 0x1c,
	0x41, 0xd5, 0x9e, 0xf4, 0x45, 0xa7, 0xbc, 0x56, 0xaf, 0x2d, 0x78, 0x41, 0x5d, 0x6a, 0xf7, 0x8f,
	0x32, 0x54, 0x31, 0x40, 0xee, 0x43, 0xed, 0x65, 0x2a, 0x25, 0x33, 0x62, 0x6e, 0xb2, 0x32, 0x9b,
	0x36, 0x30, 0x14, 0x73, 0x43, 0x3e, 0x85, 0x7b, 0xc9, 0x98, 0xdb, 0xd3, 0x81, 0x69, 0x38, 0x89,
	0x45, 0x71, 0x28, 0xe6, 0x9d, 0xea, 0x7e, 0xa9, 0x57, 0xa5, 0xbb, 0x6e, 0x18, 0x51, 0x36, 0xe3,
	0xd4, 0x8e, 0x5d, 0x7b, 0x89, 0x97, 0xaf, 0xbd, 0xc4, 0x3f, 0x80, 0x96, 0x98, 0x8b, 0x20, 0xf5,
	0xae, 0x70, 0x77, 0x37, 0x6f, 0xe5, 0x61, 0x77, 0x3f, 0x9f, 0x82, 0x3b, 0x12, 0x4c, 0xab, 0xd9,
	0x52, 0xf7, 0x87, 0x2b, 0xe9, 0x7e, 0x1a, 0x9b, 0xcf, 0x9e, 0xd2, 0x1a, 0x66, 0x53, 0x35, 0x4b,
	0xba, 0xbf, 0x01, 0x79, 0xd3, 0xdc, 0xae, 0x54, 0xfc, 0x1b, 0x68, 0xf8, 0x56, 0x9a, 0x09, 0x7f,
	0xb8, 0xbe, 0x83, 0xd2, 0xba, 0x67, 0x9d, 0xdd, 0x5f, 0xa1, 0xee, 0x8d, 0x61, 0xe5, 0xa2, 0x09,
	0xf8, 0xfc, 0x3f, 0x2a, 0xd9, 0xfd, 0x05, 0x5a, 0x83, 0xcb, 0xcd, 0xbb, 0x79, 0x2f, 0xbc, 0x07,
	0x5b, 0x46, 0xa7, 0xb1, 0xdb, 0xf6, 0x38, 0xc3, 0xfd, 0x16, 0x36, 0xf3, 0x28, 0x4e, 0x7b, 0x00,
	0x5b, 0x39, 0x83, 0x25, 0x63, 0xfe, 0x04, 0xcb, 0xd7, 0x68, 0x63, 0x09, 0x1a, 0x8c, 0xf9, 0x93,
	0xa3, 0xb7, 0xc1, 0xfe, 0xc0, 0xf7, 0x97, 0xe2, 0xf5, 0xad, 0x78, 0x7d, 0x3e, 0x8d, 0xfa, 0x28,
	0xde, 0xf9, 0x1d, 0xfc, 0x55, 0xff, 0xe4, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x18, 0x1e,
	0x4d, 0xf7, 0x0b, 0x00, 0x00,
}
