// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kava/committee/v1beta1/permissions.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GodPermission allows any governance proposal. It is used mainly for testing.
type GodPermission struct {
}

func (m *GodPermission) Reset()         { *m = GodPermission{} }
func (m *GodPermission) String() string { return proto.CompactTextString(m) }
func (*GodPermission) ProtoMessage()    {}
func (*GodPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdfaf7be16465ae4, []int{0}
}
func (m *GodPermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GodPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GodPermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GodPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GodPermission.Merge(m, src)
}
func (m *GodPermission) XXX_Size() int {
	return m.Size()
}
func (m *GodPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_GodPermission.DiscardUnknown(m)
}

var xxx_messageInfo_GodPermission proto.InternalMessageInfo

// SoftwareUpgradePermission permission type for software upgrade proposals
type SoftwareUpgradePermission struct {
}

func (m *SoftwareUpgradePermission) Reset()         { *m = SoftwareUpgradePermission{} }
func (m *SoftwareUpgradePermission) String() string { return proto.CompactTextString(m) }
func (*SoftwareUpgradePermission) ProtoMessage()    {}
func (*SoftwareUpgradePermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdfaf7be16465ae4, []int{1}
}
func (m *SoftwareUpgradePermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SoftwareUpgradePermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SoftwareUpgradePermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SoftwareUpgradePermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoftwareUpgradePermission.Merge(m, src)
}
func (m *SoftwareUpgradePermission) XXX_Size() int {
	return m.Size()
}
func (m *SoftwareUpgradePermission) XXX_DiscardUnknown() {
	xxx_messageInfo_SoftwareUpgradePermission.DiscardUnknown(m)
}

var xxx_messageInfo_SoftwareUpgradePermission proto.InternalMessageInfo

// TextPermission allows any text governance proposal.
type TextPermission struct {
}

func (m *TextPermission) Reset()         { *m = TextPermission{} }
func (m *TextPermission) String() string { return proto.CompactTextString(m) }
func (*TextPermission) ProtoMessage()    {}
func (*TextPermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdfaf7be16465ae4, []int{2}
}
func (m *TextPermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TextPermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TextPermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TextPermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextPermission.Merge(m, src)
}
func (m *TextPermission) XXX_Size() int {
	return m.Size()
}
func (m *TextPermission) XXX_DiscardUnknown() {
	xxx_messageInfo_TextPermission.DiscardUnknown(m)
}

var xxx_messageInfo_TextPermission proto.InternalMessageInfo

// ParamsChangePermission allows any parameter or sub parameter change proposal.
type ParamsChangePermission struct {
	AllowedParamsChanges AllowedParamsChanges `protobuf:"bytes,1,rep,name=allowed_params_changes,json=allowedParamsChanges,proto3,castrepeated=AllowedParamsChanges" json:"allowed_params_changes"`
}

func (m *ParamsChangePermission) Reset()         { *m = ParamsChangePermission{} }
func (m *ParamsChangePermission) String() string { return proto.CompactTextString(m) }
func (*ParamsChangePermission) ProtoMessage()    {}
func (*ParamsChangePermission) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdfaf7be16465ae4, []int{3}
}
func (m *ParamsChangePermission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamsChangePermission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamsChangePermission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamsChangePermission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamsChangePermission.Merge(m, src)
}
func (m *ParamsChangePermission) XXX_Size() int {
	return m.Size()
}
func (m *ParamsChangePermission) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamsChangePermission.DiscardUnknown(m)
}

var xxx_messageInfo_ParamsChangePermission proto.InternalMessageInfo

func (m *ParamsChangePermission) GetAllowedParamsChanges() AllowedParamsChanges {
	if m != nil {
		return m.AllowedParamsChanges
	}
	return nil
}

// AllowedParamsChange contains data on the allowed parameter changes for subspace, key, and sub params requirements.
type AllowedParamsChange struct {
	Subspace string `protobuf:"bytes,1,opt,name=subspace,proto3" json:"subspace,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Requirements for when the subparam value is a single record. This contains list of allowed attribute keys that can
	// be changed on the subparam record.
	SingleSubparamAllowedAttrs []string `protobuf:"bytes,3,rep,name=single_subparam_allowed_attrs,json=singleSubparamAllowedAttrs,proto3" json:"single_subparam_allowed_attrs,omitempty"`
	// Requirements for when the subparam value is a list of records. The requirements contains requirements for each
	// record in the list.
	MultiSubparamsRequirements []SubparamRequirement `protobuf:"bytes,4,rep,name=multi_subparams_requirements,json=multiSubparamsRequirements,proto3" json:"multi_subparams_requirements"`
}

func (m *AllowedParamsChange) Reset()         { *m = AllowedParamsChange{} }
func (m *AllowedParamsChange) String() string { return proto.CompactTextString(m) }
func (*AllowedParamsChange) ProtoMessage()    {}
func (*AllowedParamsChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdfaf7be16465ae4, []int{4}
}
func (m *AllowedParamsChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedParamsChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedParamsChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedParamsChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedParamsChange.Merge(m, src)
}
func (m *AllowedParamsChange) XXX_Size() int {
	return m.Size()
}
func (m *AllowedParamsChange) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedParamsChange.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedParamsChange proto.InternalMessageInfo

func (m *AllowedParamsChange) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func (m *AllowedParamsChange) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AllowedParamsChange) GetSingleSubparamAllowedAttrs() []string {
	if m != nil {
		return m.SingleSubparamAllowedAttrs
	}
	return nil
}

func (m *AllowedParamsChange) GetMultiSubparamsRequirements() []SubparamRequirement {
	if m != nil {
		return m.MultiSubparamsRequirements
	}
	return nil
}

// SubparamRequirement contains requirements for a single record in a subparam value list
type SubparamRequirement struct {
	// The required attr key of the param record.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The required param value for the param record key. The key and value is used to match to the target param record.
	Val string `protobuf:"bytes,2,opt,name=val,proto3" json:"val,omitempty"`
	// The sub param attrs that are allowed to be changed.
	AllowedSubparamAttrChanges []string `protobuf:"bytes,3,rep,name=allowed_subparam_attr_changes,json=allowedSubparamAttrChanges,proto3" json:"allowed_subparam_attr_changes,omitempty"`
}

func (m *SubparamRequirement) Reset()         { *m = SubparamRequirement{} }
func (m *SubparamRequirement) String() string { return proto.CompactTextString(m) }
func (*SubparamRequirement) ProtoMessage()    {}
func (*SubparamRequirement) Descriptor() ([]byte, []int) {
	return fileDescriptor_bdfaf7be16465ae4, []int{5}
}
func (m *SubparamRequirement) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SubparamRequirement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SubparamRequirement.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SubparamRequirement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubparamRequirement.Merge(m, src)
}
func (m *SubparamRequirement) XXX_Size() int {
	return m.Size()
}
func (m *SubparamRequirement) XXX_DiscardUnknown() {
	xxx_messageInfo_SubparamRequirement.DiscardUnknown(m)
}

var xxx_messageInfo_SubparamRequirement proto.InternalMessageInfo

func (m *SubparamRequirement) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SubparamRequirement) GetVal() string {
	if m != nil {
		return m.Val
	}
	return ""
}

func (m *SubparamRequirement) GetAllowedSubparamAttrChanges() []string {
	if m != nil {
		return m.AllowedSubparamAttrChanges
	}
	return nil
}

func init() {
	proto.RegisterType((*GodPermission)(nil), "kava.committee.v1beta1.GodPermission")
	proto.RegisterType((*SoftwareUpgradePermission)(nil), "kava.committee.v1beta1.SoftwareUpgradePermission")
	proto.RegisterType((*TextPermission)(nil), "kava.committee.v1beta1.TextPermission")
	proto.RegisterType((*ParamsChangePermission)(nil), "kava.committee.v1beta1.ParamsChangePermission")
	proto.RegisterType((*AllowedParamsChange)(nil), "kava.committee.v1beta1.AllowedParamsChange")
	proto.RegisterType((*SubparamRequirement)(nil), "kava.committee.v1beta1.SubparamRequirement")
}

func init() {
	proto.RegisterFile("kava/committee/v1beta1/permissions.proto", fileDescriptor_bdfaf7be16465ae4)
}

var fileDescriptor_bdfaf7be16465ae4 = []byte{
	// 451 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0x80, 0xb3, 0xb8, 0x42, 0x74, 0x11, 0x55, 0xe5, 0x46, 0x91, 0x6b, 0x15, 0x27, 0xca, 0x29,
	0x52, 0x54, 0x5b, 0x85, 0x1b, 0xb7, 0x04, 0x21, 0xae, 0x95, 0x0b, 0x17, 0x2e, 0xd6, 0x3a, 0x59,
	0x5c, 0xab, 0x76, 0xd6, 0xec, 0x8c, 0xd3, 0x56, 0x42, 0xe2, 0x15, 0x78, 0x0d, 0x38, 0xf3, 0x10,
	0x15, 0xa7, 0x1e, 0x39, 0x01, 0x4a, 0x1e, 0x83, 0x0b, 0xda, 0xf5, 0xaf, 0x54, 0x2b, 0xb7, 0x99,
	0xd9, 0x6f, 0x66, 0xfd, 0xed, 0x7a, 0xe9, 0xe4, 0x8a, 0xad, 0x99, 0xb7, 0x10, 0x69, 0x1a, 0x23,
	0x72, 0xee, 0xad, 0xcf, 0x42, 0x8e, 0xec, 0xcc, 0xcb, 0xb8, 0x4c, 0x63, 0x80, 0x58, 0xac, 0xc0,
	0xcd, 0xa4, 0x40, 0x61, 0x0e, 0x14, 0xe9, 0xd6, 0xa4, 0x5b, 0x92, 0xf6, 0xf1, 0x42, 0x40, 0x2a,
	0x20, 0xd0, 0x94, 0x57, 0x24, 0x45, 0x8b, 0xdd, 0x8f, 0x44, 0x24, 0x8a, 0xba, 0x8a, 0x8a, 0xea,
	0x78, 0x48, 0x9f, 0xbd, 0x15, 0xcb, 0xf3, 0x7a, 0x83, 0x57, 0x07, 0x3f, 0x7f, 0x9c, 0xd2, 0x26,
	0x1f, 0x4f, 0xe9, 0xf1, 0x85, 0xf8, 0x88, 0xd7, 0x4c, 0xf2, 0xf7, 0x59, 0x24, 0xd9, 0x92, 0xef,
	0x80, 0x47, 0xf4, 0xe0, 0x1d, 0xbf, 0xc1, 0x1d, 0xc4, 0x37, 0x42, 0x07, 0xe7, 0x4c, 0xb2, 0x14,
	0x5e, 0x5f, 0xb2, 0x55, 0xd4, 0x1a, 0x66, 0x7e, 0xa1, 0x03, 0x96, 0x24, 0xe2, 0x9a, 0x2f, 0x83,
	0x4c, 0x13, 0xc1, 0x42, 0x23, 0x60, 0x91, 0x91, 0x31, 0x79, 0xfa, 0x62, 0xea, 0x76, 0x4b, 0xbb,
	0xb3, 0xa2, 0xab, 0x3d, 0x76, 0x7e, 0x72, 0xf7, 0x7b, 0xd8, 0xfb, 0xfe, 0x67, 0xd8, 0xef, 0x58,
	0x04, 0xbf, 0xcf, 0x3a, 0xaa, 0x0f, 0xbe, 0xf5, 0x1f, 0xa1, 0x47, 0x1d, 0xed, 0xa6, 0x4d, 0x9f,
	0x40, 0x1e, 0x42, 0xc6, 0x16, 0xdc, 0x22, 0x23, 0x32, 0xd9, 0xf7, 0xeb, 0xdc, 0x3c, 0xa4, 0xc6,
	0x15, 0xbf, 0xb5, 0x1e, 0xe9, 0xb2, 0x0a, 0xcd, 0x19, 0x7d, 0x0e, 0xf1, 0x2a, 0x4a, 0x78, 0x00,
	0x79, 0xa8, 0xc5, 0x82, 0x4a, 0x93, 0x21, 0x4a, 0xb0, 0x8c, 0x91, 0x31, 0xd9, 0xf7, 0xed, 0x02,
	0xba, 0x28, 0x99, 0x72, 0xdf, 0x99, 0x22, 0x4c, 0xa0, 0x27, 0x69, 0x9e, 0x60, 0x5c, 0x4f, 0x80,
	0x40, 0xf2, 0x4f, 0x79, 0x2c, 0x79, 0xca, 0x57, 0x08, 0xd6, 0xde, 0xee, 0xf3, 0xa9, 0x66, 0xfa,
	0x4d, 0xcf, 0x7c, 0x4f, 0x9d, 0x8f, 0x6f, 0xeb, 0xb1, 0xd5, 0x3a, 0xb4, 0x00, 0x18, 0x7f, 0xa6,
	0x47, 0x1d, 0x8d, 0x95, 0x20, 0x69, 0x04, 0x0f, 0xa9, 0xb1, 0x66, 0x49, 0xa5, 0xbc, 0x66, 0x89,
	0x52, 0xae, 0x14, 0x1b, 0x67, 0x44, 0x59, 0x5f, 0x68, 0xa9, 0x5c, 0x42, 0xb5, 0x33, 0xa2, 0x2c,
	0xef, 0x62, 0xfe, 0xe6, 0x6e, 0xe3, 0x90, 0xfb, 0x8d, 0x43, 0xfe, 0x6e, 0x1c, 0xf2, 0x75, 0xeb,
	0xf4, 0xee, 0xb7, 0x4e, 0xef, 0xd7, 0xd6, 0xe9, 0x7d, 0x98, 0x46, 0x31, 0x5e, 0xe6, 0xa1, 0xf2,
	0xf4, 0x94, 0xf0, 0x69, 0xc2, 0x42, 0xd0, 0x91, 0x77, 0xd3, 0x7a, 0x3b, 0x78, 0x9b, 0x71, 0x08,
	0x1f, 0xeb, 0xbf, 0xfc, 0xe5, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0x09, 0x36, 0x56, 0x5a,
	0x03, 0x00, 0x00,
}

func (m *GodPermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GodPermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GodPermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *SoftwareUpgradePermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SoftwareUpgradePermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SoftwareUpgradePermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *TextPermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TextPermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TextPermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ParamsChangePermission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamsChangePermission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamsChangePermission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedParamsChanges) > 0 {
		for iNdEx := len(m.AllowedParamsChanges) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AllowedParamsChanges[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPermissions(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *AllowedParamsChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedParamsChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedParamsChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MultiSubparamsRequirements) > 0 {
		for iNdEx := len(m.MultiSubparamsRequirements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MultiSubparamsRequirements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPermissions(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.SingleSubparamAllowedAttrs) > 0 {
		for iNdEx := len(m.SingleSubparamAllowedAttrs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SingleSubparamAllowedAttrs[iNdEx])
			copy(dAtA[i:], m.SingleSubparamAllowedAttrs[iNdEx])
			i = encodeVarintPermissions(dAtA, i, uint64(len(m.SingleSubparamAllowedAttrs[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintPermissions(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintPermissions(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SubparamRequirement) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SubparamRequirement) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SubparamRequirement) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowedSubparamAttrChanges) > 0 {
		for iNdEx := len(m.AllowedSubparamAttrChanges) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedSubparamAttrChanges[iNdEx])
			copy(dAtA[i:], m.AllowedSubparamAttrChanges[iNdEx])
			i = encodeVarintPermissions(dAtA, i, uint64(len(m.AllowedSubparamAttrChanges[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Val) > 0 {
		i -= len(m.Val)
		copy(dAtA[i:], m.Val)
		i = encodeVarintPermissions(dAtA, i, uint64(len(m.Val)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintPermissions(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPermissions(dAtA []byte, offset int, v uint64) int {
	offset -= sovPermissions(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GodPermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SoftwareUpgradePermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *TextPermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ParamsChangePermission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AllowedParamsChanges) > 0 {
		for _, e := range m.AllowedParamsChanges {
			l = e.Size()
			n += 1 + l + sovPermissions(uint64(l))
		}
	}
	return n
}

func (m *AllowedParamsChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovPermissions(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovPermissions(uint64(l))
	}
	if len(m.SingleSubparamAllowedAttrs) > 0 {
		for _, s := range m.SingleSubparamAllowedAttrs {
			l = len(s)
			n += 1 + l + sovPermissions(uint64(l))
		}
	}
	if len(m.MultiSubparamsRequirements) > 0 {
		for _, e := range m.MultiSubparamsRequirements {
			l = e.Size()
			n += 1 + l + sovPermissions(uint64(l))
		}
	}
	return n
}

func (m *SubparamRequirement) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovPermissions(uint64(l))
	}
	l = len(m.Val)
	if l > 0 {
		n += 1 + l + sovPermissions(uint64(l))
	}
	if len(m.AllowedSubparamAttrChanges) > 0 {
		for _, s := range m.AllowedSubparamAttrChanges {
			l = len(s)
			n += 1 + l + sovPermissions(uint64(l))
		}
	}
	return n
}

func sovPermissions(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPermissions(x uint64) (n int) {
	return sovPermissions(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GodPermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: GodPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GodPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *SoftwareUpgradePermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: SoftwareUpgradePermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SoftwareUpgradePermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *TextPermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: TextPermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TextPermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *ParamsChangePermission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: ParamsChangePermission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamsChangePermission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedParamsChanges", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedParamsChanges = append(m.AllowedParamsChanges, AllowedParamsChange{})
			if err := m.AllowedParamsChanges[len(m.AllowedParamsChanges)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *AllowedParamsChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: AllowedParamsChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedParamsChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SingleSubparamAllowedAttrs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SingleSubparamAllowedAttrs = append(m.SingleSubparamAllowedAttrs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MultiSubparamsRequirements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MultiSubparamsRequirements = append(m.MultiSubparamsRequirements, SubparamRequirement{})
			if err := m.MultiSubparamsRequirements[len(m.MultiSubparamsRequirements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func (m *SubparamRequirement) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPermissions
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
			return fmt.Errorf("proto: SubparamRequirement: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SubparamRequirement: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Val", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Val = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedSubparamAttrChanges", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPermissions
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
				return ErrInvalidLengthPermissions
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPermissions
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedSubparamAttrChanges = append(m.AllowedSubparamAttrChanges, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPermissions(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPermissions
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
func skipPermissions(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPermissions
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
					return 0, ErrIntOverflowPermissions
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
					return 0, ErrIntOverflowPermissions
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
				return 0, ErrInvalidLengthPermissions
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPermissions
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPermissions
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPermissions        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPermissions          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPermissions = fmt.Errorf("proto: unexpected end of group")
)
