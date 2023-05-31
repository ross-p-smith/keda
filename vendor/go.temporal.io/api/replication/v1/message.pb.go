// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/replication/v1/message.proto

package replication

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
	time "time"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	v1 "go.temporal.io/api/enums/v1"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ClusterReplicationConfig struct {
	ClusterName string `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
}

func (m *ClusterReplicationConfig) Reset()      { *m = ClusterReplicationConfig{} }
func (*ClusterReplicationConfig) ProtoMessage() {}
func (*ClusterReplicationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4566ff8a441e2e, []int{0}
}
func (m *ClusterReplicationConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClusterReplicationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClusterReplicationConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClusterReplicationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterReplicationConfig.Merge(m, src)
}
func (m *ClusterReplicationConfig) XXX_Size() int {
	return m.Size()
}
func (m *ClusterReplicationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterReplicationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterReplicationConfig proto.InternalMessageInfo

func (m *ClusterReplicationConfig) GetClusterName() string {
	if m != nil {
		return m.ClusterName
	}
	return ""
}

type NamespaceReplicationConfig struct {
	ActiveClusterName string                      `protobuf:"bytes,1,opt,name=active_cluster_name,json=activeClusterName,proto3" json:"active_cluster_name,omitempty"`
	Clusters          []*ClusterReplicationConfig `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	State             v1.ReplicationState         `protobuf:"varint,3,opt,name=state,proto3,enum=temporal.api.enums.v1.ReplicationState" json:"state,omitempty"`
}

func (m *NamespaceReplicationConfig) Reset()      { *m = NamespaceReplicationConfig{} }
func (*NamespaceReplicationConfig) ProtoMessage() {}
func (*NamespaceReplicationConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4566ff8a441e2e, []int{1}
}
func (m *NamespaceReplicationConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NamespaceReplicationConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NamespaceReplicationConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NamespaceReplicationConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NamespaceReplicationConfig.Merge(m, src)
}
func (m *NamespaceReplicationConfig) XXX_Size() int {
	return m.Size()
}
func (m *NamespaceReplicationConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NamespaceReplicationConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NamespaceReplicationConfig proto.InternalMessageInfo

func (m *NamespaceReplicationConfig) GetActiveClusterName() string {
	if m != nil {
		return m.ActiveClusterName
	}
	return ""
}

func (m *NamespaceReplicationConfig) GetClusters() []*ClusterReplicationConfig {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *NamespaceReplicationConfig) GetState() v1.ReplicationState {
	if m != nil {
		return m.State
	}
	return v1.REPLICATION_STATE_UNSPECIFIED
}

// Represents a historical replication status of a Namespace
type FailoverStatus struct {
	// Timestamp when the Cluster switched to the following failover_version
	FailoverTime    *time.Time `protobuf:"bytes,1,opt,name=failover_time,json=failoverTime,proto3,stdtime" json:"failover_time,omitempty"`
	FailoverVersion int64      `protobuf:"varint,2,opt,name=failover_version,json=failoverVersion,proto3" json:"failover_version,omitempty"`
}

func (m *FailoverStatus) Reset()      { *m = FailoverStatus{} }
func (*FailoverStatus) ProtoMessage() {}
func (*FailoverStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd4566ff8a441e2e, []int{2}
}
func (m *FailoverStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FailoverStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FailoverStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FailoverStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FailoverStatus.Merge(m, src)
}
func (m *FailoverStatus) XXX_Size() int {
	return m.Size()
}
func (m *FailoverStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FailoverStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FailoverStatus proto.InternalMessageInfo

func (m *FailoverStatus) GetFailoverTime() *time.Time {
	if m != nil {
		return m.FailoverTime
	}
	return nil
}

func (m *FailoverStatus) GetFailoverVersion() int64 {
	if m != nil {
		return m.FailoverVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*ClusterReplicationConfig)(nil), "temporal.api.replication.v1.ClusterReplicationConfig")
	proto.RegisterType((*NamespaceReplicationConfig)(nil), "temporal.api.replication.v1.NamespaceReplicationConfig")
	proto.RegisterType((*FailoverStatus)(nil), "temporal.api.replication.v1.FailoverStatus")
}

func init() {
	proto.RegisterFile("temporal/api/replication/v1/message.proto", fileDescriptor_dd4566ff8a441e2e)
}

var fileDescriptor_dd4566ff8a441e2e = []byte{
	// 472 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xf4, 0x26, 0x80, 0x60, 0x13, 0x0a, 0x98, 0x8b, 0x15, 0xc4, 0xd6, 0x8d, 0x84, 0x48, 0x0f,
	0xac, 0x95, 0x20, 0x2e, 0x46, 0x3d, 0x34, 0x11, 0xdc, 0xa8, 0x8a, 0xa9, 0x72, 0xe0, 0x12, 0x6d,
	0xdd, 0x2f, 0xd6, 0x4a, 0xb6, 0x77, 0xe5, 0xdd, 0xf8, 0x8c, 0x78, 0x82, 0x3e, 0x06, 0xe2, 0x49,
	0x38, 0xe6, 0x84, 0x7a, 0x02, 0xe2, 0x5c, 0x10, 0xa7, 0x3e, 0x02, 0xf2, 0x5f, 0xe3, 0x8a, 0x92,
	0x9b, 0x3d, 0x33, 0xdf, 0x68, 0x76, 0xbe, 0x0f, 0xef, 0x6b, 0x88, 0xa4, 0x48, 0x58, 0xe8, 0x30,
	0xc9, 0x9d, 0x04, 0x64, 0xc8, 0x7d, 0xa6, 0xb9, 0x88, 0x9d, 0x74, 0xe8, 0x44, 0xa0, 0x14, 0x0b,
	0x80, 0xca, 0x44, 0x68, 0x61, 0x3e, 0xa9, 0xa5, 0x94, 0x49, 0x4e, 0x1b, 0x52, 0x9a, 0x0e, 0x7b,
	0xbb, 0x81, 0x10, 0x41, 0x08, 0x4e, 0x21, 0x3d, 0x5d, 0xcc, 0x1d, 0xcd, 0x23, 0x50, 0x9a, 0x45,
	0xb2, 0x9c, 0xee, 0xed, 0x9d, 0x81, 0x84, 0xf8, 0x0c, 0x62, 0x9f, 0x83, 0x72, 0x02, 0x11, 0x88,
	0x02, 0x2f, 0xbe, 0x2a, 0xc9, 0xb3, 0x6b, 0x59, 0x20, 0x5e, 0x44, 0x2a, 0x4f, 0x11, 0xb3, 0x08,
	0x94, 0x64, 0x7e, 0x95, 0xa3, 0x7f, 0x80, 0xad, 0x49, 0xb8, 0x50, 0x1a, 0x12, 0x6f, 0x93, 0x61,
	0x22, 0xe2, 0x39, 0x0f, 0xcc, 0x3d, 0xdc, 0xf5, 0x4b, 0x6e, 0x96, 0x8f, 0x59, 0xc8, 0x46, 0x83,
	0x7b, 0x5e, 0xa7, 0xc2, 0x8e, 0x58, 0x04, 0xfd, 0x1f, 0x08, 0xf7, 0x8e, 0x6a, 0xcb, 0x7f, 0x1d,
	0x28, 0x7e, 0xcc, 0x7c, 0xcd, 0x53, 0x98, 0xdd, 0x60, 0xf4, 0xa8, 0xa4, 0x26, 0x1b, 0x3b, 0xf3,
	0x3d, 0xbe, 0x5b, 0x09, 0x95, 0xd5, 0xb2, 0xdb, 0x83, 0xce, 0xe8, 0x15, 0xdd, 0x52, 0x14, 0xfd,
	0x5f, 0x74, 0xef, 0xca, 0xc6, 0x3c, 0xc0, 0xb7, 0x95, 0x66, 0x1a, 0xac, 0xb6, 0x8d, 0x06, 0x3b,
	0xa3, 0xe7, 0xd7, 0xfd, 0x8a, 0x5e, 0x72, 0xa7, 0x86, 0xc5, 0x87, 0x5c, 0xee, 0x95, 0x53, 0xfd,
	0xcf, 0x08, 0xef, 0xbc, 0x65, 0x3c, 0x14, 0x29, 0x24, 0x39, 0xb1, 0x50, 0xe6, 0x1b, 0x7c, 0x7f,
	0x5e, 0x21, 0xb3, 0x7c, 0x31, 0xc5, 0x73, 0x3a, 0xa3, 0x1e, 0x2d, 0xb7, 0x46, 0xeb, 0xad, 0xd1,
	0x93, 0x7a, 0x6b, 0xe3, 0x5b, 0xe7, 0x3f, 0x77, 0x91, 0xd7, 0xad, 0xc7, 0x72, 0xc2, 0xdc, 0xc7,
	0x0f, 0xaf, 0x6c, 0x52, 0x48, 0x14, 0x17, 0xb1, 0xd5, 0xb2, 0xd1, 0xa0, 0xed, 0x3d, 0xa8, 0xf1,
	0x69, 0x09, 0x8f, 0xbf, 0xa3, 0xe5, 0x8a, 0x18, 0x17, 0x2b, 0x62, 0x5c, 0xae, 0x08, 0xfa, 0x94,
	0x11, 0xf4, 0x25, 0x23, 0xe8, 0x5b, 0x46, 0xd0, 0x32, 0x23, 0xe8, 0x57, 0x46, 0xd0, 0xef, 0x8c,
	0x18, 0x97, 0x19, 0x41, 0xe7, 0x6b, 0x62, 0x2c, 0xd7, 0xc4, 0xb8, 0x58, 0x13, 0x03, 0x13, 0x2e,
	0xb6, 0xb5, 0x37, 0xee, 0xbe, 0x2b, 0x4f, 0xf2, 0x38, 0x0f, 0x7c, 0x8c, 0x3e, 0xbe, 0x08, 0x1a,
	0x7a, 0x2e, 0x6e, 0x38, 0xe2, 0xd7, 0x8d, 0xdf, 0xaf, 0xad, 0xa7, 0x27, 0x95, 0x98, 0x0b, 0x7a,
	0x28, 0x79, 0xb3, 0x43, 0x3a, 0x1d, 0xfe, 0x69, 0xd9, 0x1b, 0xde, 0x75, 0x0f, 0x25, 0x77, 0xdd,
	0x86, 0xc2, 0x75, 0xa7, 0xc3, 0xd3, 0x3b, 0x45, 0x57, 0x2f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xbc, 0xc9, 0x0f, 0x40, 0x39, 0x03, 0x00, 0x00,
}

func (this *ClusterReplicationConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClusterReplicationConfig)
	if !ok {
		that2, ok := that.(ClusterReplicationConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ClusterName != that1.ClusterName {
		return false
	}
	return true
}
func (this *NamespaceReplicationConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NamespaceReplicationConfig)
	if !ok {
		that2, ok := that.(NamespaceReplicationConfig)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ActiveClusterName != that1.ActiveClusterName {
		return false
	}
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if !this.Clusters[i].Equal(that1.Clusters[i]) {
			return false
		}
	}
	if this.State != that1.State {
		return false
	}
	return true
}
func (this *FailoverStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*FailoverStatus)
	if !ok {
		that2, ok := that.(FailoverStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.FailoverTime == nil {
		if this.FailoverTime != nil {
			return false
		}
	} else if !this.FailoverTime.Equal(*that1.FailoverTime) {
		return false
	}
	if this.FailoverVersion != that1.FailoverVersion {
		return false
	}
	return true
}
func (this *ClusterReplicationConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&replication.ClusterReplicationConfig{")
	s = append(s, "ClusterName: "+fmt.Sprintf("%#v", this.ClusterName)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *NamespaceReplicationConfig) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&replication.NamespaceReplicationConfig{")
	s = append(s, "ActiveClusterName: "+fmt.Sprintf("%#v", this.ActiveClusterName)+",\n")
	if this.Clusters != nil {
		s = append(s, "Clusters: "+fmt.Sprintf("%#v", this.Clusters)+",\n")
	}
	s = append(s, "State: "+fmt.Sprintf("%#v", this.State)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *FailoverStatus) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&replication.FailoverStatus{")
	s = append(s, "FailoverTime: "+fmt.Sprintf("%#v", this.FailoverTime)+",\n")
	s = append(s, "FailoverVersion: "+fmt.Sprintf("%#v", this.FailoverVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMessage(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *ClusterReplicationConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterReplicationConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClusterReplicationConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClusterName) > 0 {
		i -= len(m.ClusterName)
		copy(dAtA[i:], m.ClusterName)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *NamespaceReplicationConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NamespaceReplicationConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NamespaceReplicationConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.State != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Clusters) > 0 {
		for iNdEx := len(m.Clusters) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Clusters[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMessage(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ActiveClusterName) > 0 {
		i -= len(m.ActiveClusterName)
		copy(dAtA[i:], m.ActiveClusterName)
		i = encodeVarintMessage(dAtA, i, uint64(len(m.ActiveClusterName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FailoverStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FailoverStatus) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FailoverStatus) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FailoverVersion != 0 {
		i = encodeVarintMessage(dAtA, i, uint64(m.FailoverVersion))
		i--
		dAtA[i] = 0x10
	}
	if m.FailoverTime != nil {
		n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.FailoverTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.FailoverTime):])
		if err1 != nil {
			return 0, err1
		}
		i -= n1
		i = encodeVarintMessage(dAtA, i, uint64(n1))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClusterReplicationConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClusterName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *NamespaceReplicationConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ActiveClusterName)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if len(m.Clusters) > 0 {
		for _, e := range m.Clusters {
			l = e.Size()
			n += 1 + l + sovMessage(uint64(l))
		}
	}
	if m.State != 0 {
		n += 1 + sovMessage(uint64(m.State))
	}
	return n
}

func (m *FailoverStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FailoverTime != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.FailoverTime)
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.FailoverVersion != 0 {
		n += 1 + sovMessage(uint64(m.FailoverVersion))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ClusterReplicationConfig) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterReplicationConfig{`,
		`ClusterName:` + fmt.Sprintf("%v", this.ClusterName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *NamespaceReplicationConfig) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForClusters := "[]*ClusterReplicationConfig{"
	for _, f := range this.Clusters {
		repeatedStringForClusters += strings.Replace(f.String(), "ClusterReplicationConfig", "ClusterReplicationConfig", 1) + ","
	}
	repeatedStringForClusters += "}"
	s := strings.Join([]string{`&NamespaceReplicationConfig{`,
		`ActiveClusterName:` + fmt.Sprintf("%v", this.ActiveClusterName) + `,`,
		`Clusters:` + repeatedStringForClusters + `,`,
		`State:` + fmt.Sprintf("%v", this.State) + `,`,
		`}`,
	}, "")
	return s
}
func (this *FailoverStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&FailoverStatus{`,
		`FailoverTime:` + strings.Replace(fmt.Sprintf("%v", this.FailoverTime), "Timestamp", "types.Timestamp", 1) + `,`,
		`FailoverVersion:` + fmt.Sprintf("%v", this.FailoverVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMessage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ClusterReplicationConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClusterReplicationConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterReplicationConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NamespaceReplicationConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NamespaceReplicationConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NamespaceReplicationConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActiveClusterName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActiveClusterName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, &ClusterReplicationConfig{})
			if err := m.Clusters[len(m.Clusters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= v1.ReplicationState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *FailoverStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: FailoverStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FailoverStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailoverTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMessage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FailoverTime == nil {
				m.FailoverTime = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.FailoverTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FailoverVersion", wireType)
			}
			m.FailoverVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FailoverVersion |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessage = fmt.Errorf("proto: unexpected end of group")
)
