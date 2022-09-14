// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pylons/pylons/accounts.proto

package types

import (
	fmt "fmt"
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

type UserMap struct {
	AccountAddr string `protobuf:"bytes,1,opt,name=account_addr,json=accountAddr,proto3" json:"account_addr,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (m *UserMap) Reset()         { *m = UserMap{} }
func (m *UserMap) String() string { return proto.CompactTextString(m) }
func (*UserMap) ProtoMessage()    {}
func (*UserMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25da355fd1de5b5, []int{0}
}
func (m *UserMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UserMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UserMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UserMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMap.Merge(m, src)
}
func (m *UserMap) XXX_Size() int {
	return m.Size()
}
func (m *UserMap) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMap.DiscardUnknown(m)
}

var xxx_messageInfo_UserMap proto.InternalMessageInfo

func (m *UserMap) GetAccountAddr() string {
	if m != nil {
		return m.AccountAddr
	}
	return ""
}

func (m *UserMap) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type Username struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *Username) Reset()         { *m = Username{} }
func (m *Username) String() string { return proto.CompactTextString(m) }
func (*Username) ProtoMessage()    {}
func (*Username) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25da355fd1de5b5, []int{1}
}
func (m *Username) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Username) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Username.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Username) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Username.Merge(m, src)
}
func (m *Username) XXX_Size() int {
	return m.Size()
}
func (m *Username) XXX_DiscardUnknown() {
	xxx_messageInfo_Username.DiscardUnknown(m)
}

var xxx_messageInfo_Username proto.InternalMessageInfo

func (m *Username) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AccountAddr struct {
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *AccountAddr) Reset()         { *m = AccountAddr{} }
func (m *AccountAddr) String() string { return proto.CompactTextString(m) }
func (*AccountAddr) ProtoMessage()    {}
func (*AccountAddr) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25da355fd1de5b5, []int{2}
}
func (m *AccountAddr) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AccountAddr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AccountAddr.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AccountAddr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAddr.Merge(m, src)
}
func (m *AccountAddr) XXX_Size() int {
	return m.Size()
}
func (m *AccountAddr) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAddr.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAddr proto.InternalMessageInfo

func (m *AccountAddr) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ReferralKV struct {
	Address string           `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Users   []*RefereeSignup `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (m *ReferralKV) Reset()         { *m = ReferralKV{} }
func (m *ReferralKV) String() string { return proto.CompactTextString(m) }
func (*ReferralKV) ProtoMessage()    {}
func (*ReferralKV) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25da355fd1de5b5, []int{3}
}
func (m *ReferralKV) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReferralKV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReferralKV.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReferralKV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferralKV.Merge(m, src)
}
func (m *ReferralKV) XXX_Size() int {
	return m.Size()
}
func (m *ReferralKV) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferralKV.DiscardUnknown(m)
}

var xxx_messageInfo_ReferralKV proto.InternalMessageInfo

func (m *ReferralKV) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ReferralKV) GetUsers() []*RefereeSignup {
	if m != nil {
		return m.Users
	}
	return nil
}

type RefereeSignup struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *RefereeSignup) Reset()         { *m = RefereeSignup{} }
func (m *RefereeSignup) String() string { return proto.CompactTextString(m) }
func (*RefereeSignup) ProtoMessage()    {}
func (*RefereeSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25da355fd1de5b5, []int{4}
}
func (m *RefereeSignup) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RefereeSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RefereeSignup.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RefereeSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefereeSignup.Merge(m, src)
}
func (m *RefereeSignup) XXX_Size() int {
	return m.Size()
}
func (m *RefereeSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_RefereeSignup.DiscardUnknown(m)
}

var xxx_messageInfo_RefereeSignup proto.InternalMessageInfo

func (m *RefereeSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RefereeSignup) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type KYCAccount struct {
	AccountAddr string `protobuf:"bytes,1,opt,name=account_addr,json=accountAddr,proto3" json:"account_addr,omitempty"`
	Username    string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (m *KYCAccount) Reset()         { *m = KYCAccount{} }
func (m *KYCAccount) String() string { return proto.CompactTextString(m) }
func (*KYCAccount) ProtoMessage()    {}
func (*KYCAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d25da355fd1de5b5, []int{5}
}
func (m *KYCAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KYCAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KYCAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KYCAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KYCAccount.Merge(m, src)
}
func (m *KYCAccount) XXX_Size() int {
	return m.Size()
}
func (m *KYCAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_KYCAccount.DiscardUnknown(m)
}

var xxx_messageInfo_KYCAccount proto.InternalMessageInfo

func (m *KYCAccount) GetAccountAddr() string {
	if m != nil {
		return m.AccountAddr
	}
	return ""
}

func (m *KYCAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*UserMap)(nil), "pylons.pylons.UserMap")
	proto.RegisterType((*Username)(nil), "pylons.pylons.Username")
	proto.RegisterType((*AccountAddr)(nil), "pylons.pylons.AccountAddr")
	proto.RegisterType((*ReferralKV)(nil), "pylons.pylons.ReferralKV")
	proto.RegisterType((*RefereeSignup)(nil), "pylons.pylons.RefereeSignup")
	proto.RegisterType((*KYCAccount)(nil), "pylons.pylons.KYCAccount")
}

func init() { proto.RegisterFile("pylons/pylons/accounts.proto", fileDescriptor_d25da355fd1de5b5) }

var fileDescriptor_d25da355fd1de5b5 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0xa8, 0xcc, 0xc9,
	0xcf, 0x2b, 0xd6, 0x87, 0x52, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0xc5, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0xbc, 0x10, 0x61, 0x3d, 0x08, 0x25, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f,
	0x96, 0xd1, 0x07, 0xb1, 0x20, 0x8a, 0x94, 0x3c, 0xb8, 0xd8, 0x43, 0x8b, 0x53, 0x8b, 0x7c, 0x13,
	0x0b, 0x84, 0x14, 0xb9, 0x78, 0xa0, 0x26, 0xc4, 0x27, 0xa6, 0xa4, 0x14, 0x49, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x71, 0x43, 0xc5, 0x1c, 0x53, 0x52, 0x8a, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0x8b,
	0x53, 0x8b, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x98, 0xc0, 0xd2, 0x70, 0xbe, 0x92, 0x02, 0x17, 0x47,
	0x28, 0x94, 0x2d, 0x24, 0xc2, 0xc5, 0x5a, 0x96, 0x98, 0x53, 0x9a, 0x0a, 0x35, 0x03, 0xc2, 0x51,
	0x52, 0xe6, 0xe2, 0x76, 0x44, 0x32, 0x0c, 0xae, 0x88, 0x09, 0x59, 0x51, 0x14, 0x17, 0x57, 0x50,
	0x6a, 0x5a, 0x6a, 0x51, 0x51, 0x62, 0x8e, 0x77, 0x98, 0x90, 0x04, 0x17, 0x3b, 0xc8, 0x2d, 0xa9,
	0xc5, 0xc5, 0x50, 0xa3, 0x60, 0x5c, 0x21, 0x23, 0x2e, 0x56, 0x90, 0xd5, 0xc5, 0x12, 0x4c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x32, 0x7a, 0x28, 0xbe, 0xd5, 0x03, 0x9b, 0x91, 0x9a, 0x1a, 0x9c, 0x99,
	0x9e, 0x57, 0x5a, 0x10, 0x04, 0x51, 0xaa, 0xe4, 0xca, 0xc5, 0x8b, 0x22, 0x8e, 0xe2, 0x1f, 0x46,
	0x54, 0xff, 0x20, 0x5b, 0xcd, 0x84, 0x62, 0xb5, 0x92, 0x37, 0x17, 0x97, 0x77, 0xa4, 0x33, 0xd4,
	0x2b, 0x14, 0x06, 0x9b, 0x93, 0xdb, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78,
	0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44,
	0xe9, 0xa4, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x07, 0x80, 0x7d, 0xa5,
	0x5b, 0x92, 0x9a, 0x9c, 0x01, 0x8b, 0xed, 0x0a, 0x18, 0xa3, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89,
	0x0d, 0x1c, 0x9f, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x89, 0xa0, 0x05, 0x14, 0x02,
	0x00, 0x00,
}

func (m *UserMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UserMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountAddr) > 0 {
		i -= len(m.AccountAddr)
		copy(dAtA[i:], m.AccountAddr)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.AccountAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Username) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Username) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Username) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AccountAddr) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AccountAddr) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AccountAddr) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *ReferralKV) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReferralKV) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReferralKV) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Users) > 0 {
		for iNdEx := len(m.Users) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Users[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAccounts(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RefereeSignup) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RefereeSignup) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RefereeSignup) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *KYCAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KYCAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KYCAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Username) > 0 {
		i -= len(m.Username)
		copy(dAtA[i:], m.Username)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.Username)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AccountAddr) > 0 {
		i -= len(m.AccountAddr)
		copy(dAtA[i:], m.AccountAddr)
		i = encodeVarintAccounts(dAtA, i, uint64(len(m.AccountAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAccounts(dAtA []byte, offset int, v uint64) int {
	offset -= sovAccounts(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UserMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountAddr)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	return n
}

func (m *Username) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	return n
}

func (m *AccountAddr) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	return n
}

func (m *ReferralKV) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	if len(m.Users) > 0 {
		for _, e := range m.Users {
			l = e.Size()
			n += 1 + l + sovAccounts(uint64(l))
		}
	}
	return n
}

func (m *RefereeSignup) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	return n
}

func (m *KYCAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AccountAddr)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovAccounts(uint64(l))
	}
	return n
}

func sovAccounts(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAccounts(x uint64) (n int) {
	return sovAccounts(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: UserMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func (m *Username) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: Username: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Username: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func (m *AccountAddr) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: AccountAddr: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AccountAddr: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func (m *ReferralKV) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: ReferralKV: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReferralKV: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Users", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Users = append(m.Users, &RefereeSignup{})
			if err := m.Users[len(m.Users)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func (m *RefereeSignup) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: RefereeSignup: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RefereeSignup: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func (m *KYCAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAccounts
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
			return fmt.Errorf("proto: KYCAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KYCAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAccounts
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
				return ErrInvalidLengthAccounts
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAccounts
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAccounts(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAccounts
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
func skipAccounts(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAccounts
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
					return 0, ErrIntOverflowAccounts
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
					return 0, ErrIntOverflowAccounts
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
				return 0, ErrInvalidLengthAccounts
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAccounts
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAccounts
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAccounts        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAccounts          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAccounts = fmt.Errorf("proto: unexpected end of group")
)
