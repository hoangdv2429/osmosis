// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/concentrated-liquidity/tickInfo.proto

// this is a legacy package that requires additional migration logic
// in order to use the correct packge. Decision made to use legacy package path
// until clear steps for migration logic and the unknowns for state breaking are
// investigated for changing proto package.

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

type TickInfo struct {
	LiquidityGross github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=liquidity_gross,json=liquidityGross,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidity_gross" yaml:"liquidity_gross"`
	LiquidityNet   github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=liquidity_net,json=liquidityNet,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidity_net" yaml:"liquidity_net"`
	// Total spread rewards accumulated in the opposite direction that the tick
	// was last crossed. i.e. if the current tick is to the right of this tick
	// (meaning its currently a greater price), then this is the total spread
	// rewards accumulated below the tick. If the current tick is to the left of
	// this tick (meaning its currently at a lower price), then this is the total
	// spread rewards accumulated above the tick.
	//
	// Note: the way this value is used depends on the direction of spread rewards
	// we are calculating for. If we are calculating spread rewards below the
	// lower tick and the lower tick is the active tick, then this is the
	// spreadRewardGrowthGlobal - the lower tick's
	// spreadRewardGrowthOppositeDirectionOfLastTraversal. If we are calculating
	// spread rewards above the upper tick and the upper tick is the active tick,
	// then this is just the tick's
	// spreadRewardGrowthOppositeDirectionOfLastTraversal value.
	SpreadRewardGrowthOppositeDirectionOfLastTraversal github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,3,rep,name=spread_reward_growth_opposite_direction_of_last_traversal,json=spreadRewardGrowthOppositeDirectionOfLastTraversal,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"spread_reward_growth_opposite_direction_of_last_traversal"`
	// uptime_trackers is a container encapsulating the uptime trackers.
	// We use a container instead of a "repeated UptimeTracker" directly
	// because we need the ability to serialize and deserialize the
	// container easily for events when crossing a tick.
	UptimeTrackers UptimeTrackers `protobuf:"bytes,4,opt,name=uptime_trackers,json=uptimeTrackers,proto3" json:"uptime_trackers" yaml:"uptime_trackers"`
}

func (m *TickInfo) Reset()         { *m = TickInfo{} }
func (m *TickInfo) String() string { return proto.CompactTextString(m) }
func (*TickInfo) ProtoMessage()    {}
func (*TickInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ccb7e45032b943a, []int{0}
}
func (m *TickInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TickInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickInfo.Merge(m, src)
}
func (m *TickInfo) XXX_Size() int {
	return m.Size()
}
func (m *TickInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TickInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TickInfo proto.InternalMessageInfo

func (m *TickInfo) GetSpreadRewardGrowthOppositeDirectionOfLastTraversal() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.SpreadRewardGrowthOppositeDirectionOfLastTraversal
	}
	return nil
}

func (m *TickInfo) GetUptimeTrackers() UptimeTrackers {
	if m != nil {
		return m.UptimeTrackers
	}
	return UptimeTrackers{}
}

type UptimeTrackers struct {
	List []UptimeTracker `protobuf:"bytes,1,rep,name=list,proto3" json:"list" yaml:"list"`
}

func (m *UptimeTrackers) Reset()         { *m = UptimeTrackers{} }
func (m *UptimeTrackers) String() string { return proto.CompactTextString(m) }
func (*UptimeTrackers) ProtoMessage()    {}
func (*UptimeTrackers) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ccb7e45032b943a, []int{1}
}
func (m *UptimeTrackers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UptimeTrackers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UptimeTrackers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UptimeTrackers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeTrackers.Merge(m, src)
}
func (m *UptimeTrackers) XXX_Size() int {
	return m.Size()
}
func (m *UptimeTrackers) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeTrackers.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeTrackers proto.InternalMessageInfo

func (m *UptimeTrackers) GetList() []UptimeTracker {
	if m != nil {
		return m.List
	}
	return nil
}

type UptimeTracker struct {
	UptimeGrowthOutside github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,1,rep,name=uptime_growth_outside,json=uptimeGrowthOutside,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"uptime_growth_outside"`
}

func (m *UptimeTracker) Reset()         { *m = UptimeTracker{} }
func (m *UptimeTracker) String() string { return proto.CompactTextString(m) }
func (*UptimeTracker) ProtoMessage()    {}
func (*UptimeTracker) Descriptor() ([]byte, []int) {
	return fileDescriptor_1ccb7e45032b943a, []int{2}
}
func (m *UptimeTracker) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UptimeTracker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UptimeTracker.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UptimeTracker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UptimeTracker.Merge(m, src)
}
func (m *UptimeTracker) XXX_Size() int {
	return m.Size()
}
func (m *UptimeTracker) XXX_DiscardUnknown() {
	xxx_messageInfo_UptimeTracker.DiscardUnknown(m)
}

var xxx_messageInfo_UptimeTracker proto.InternalMessageInfo

func (m *UptimeTracker) GetUptimeGrowthOutside() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.UptimeGrowthOutside
	}
	return nil
}

func init() {
	proto.RegisterType((*TickInfo)(nil), "osmosis.concentratedliquidity.v1beta1.TickInfo")
	proto.RegisterType((*UptimeTrackers)(nil), "osmosis.concentratedliquidity.v1beta1.UptimeTrackers")
	proto.RegisterType((*UptimeTracker)(nil), "osmosis.concentratedliquidity.v1beta1.UptimeTracker")
}

func init() {
	proto.RegisterFile("osmosis/concentrated-liquidity/tickInfo.proto", fileDescriptor_1ccb7e45032b943a)
}

var fileDescriptor_1ccb7e45032b943a = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x3b, 0xee, 0x22, 0x3a, 0x75, 0xbb, 0x90, 0x5d, 0x25, 0x2e, 0x92, 0x96, 0x80, 0x52,
	0x90, 0x26, 0xb4, 0xab, 0x82, 0x1e, 0xe3, 0xe2, 0x2a, 0x88, 0x0b, 0xa1, 0x5e, 0x44, 0x09, 0x93,
	0x64, 0xda, 0x1d, 0x9a, 0x64, 0xb2, 0x33, 0x2f, 0x5d, 0x7b, 0xf1, 0xe6, 0xdd, 0x9b, 0xdf, 0xc1,
	0xbb, 0xdf, 0xa1, 0xc7, 0x3d, 0x8a, 0x87, 0x2a, 0xed, 0x37, 0xf0, 0x13, 0x48, 0x92, 0x49, 0x77,
	0x5b, 0x10, 0x56, 0xc1, 0x53, 0x32, 0x2f, 0xff, 0xf7, 0xff, 0xff, 0x78, 0x2f, 0x09, 0xee, 0x70,
	0x19, 0x73, 0xc9, 0xa4, 0x1d, 0xf0, 0x24, 0xa0, 0x09, 0x08, 0x02, 0x34, 0xec, 0x44, 0xec, 0x24,
	0x63, 0x21, 0x83, 0x89, 0x0d, 0x2c, 0x18, 0xbd, 0x48, 0x06, 0xdc, 0x4a, 0x05, 0x07, 0xae, 0xdd,
	0x55, 0x72, 0xeb, 0xa2, 0x7c, 0xa9, 0xb6, 0xc6, 0x5d, 0x9f, 0x02, 0xe9, 0xee, 0xdd, 0x0e, 0x0a,
	0x9d, 0x57, 0x34, 0xd9, 0xe5, 0xa1, 0x74, 0xd8, 0xdb, 0x1d, 0xf2, 0x21, 0x2f, 0xeb, 0xf9, 0x9d,
	0xaa, 0x1a, 0xa5, 0xc6, 0xf6, 0x89, 0xa4, 0xb6, 0x72, 0xb1, 0x03, 0xce, 0x92, 0xf2, 0xb9, 0xf9,
	0x75, 0x13, 0x5f, 0xeb, 0x2b, 0x14, 0xed, 0x04, 0x6f, 0x2f, 0x23, 0xbd, 0xa1, 0xe0, 0x52, 0xea,
	0xa8, 0x85, 0xda, 0xd7, 0x9d, 0xe7, 0xd3, 0x59, 0xb3, 0xf6, 0x7d, 0xd6, 0xbc, 0x37, 0x64, 0x70,
	0x9c, 0xf9, 0x56, 0xc0, 0x63, 0x15, 0xae, 0x2e, 0x1d, 0x19, 0x8e, 0x6c, 0x98, 0xa4, 0x54, 0x5a,
	0x07, 0x34, 0xf8, 0x35, 0x6b, 0xde, 0x9a, 0x90, 0x38, 0x7a, 0x62, 0xae, 0xd9, 0x99, 0x6e, 0x63,
	0x59, 0x39, 0xcc, 0x0b, 0xda, 0x08, 0x6f, 0x9d, 0x6b, 0x12, 0x0a, 0xfa, 0x95, 0x22, 0xf0, 0xd9,
	0x5f, 0x07, 0xee, 0xae, 0x07, 0x26, 0x14, 0x4c, 0xf7, 0xc6, 0xf2, 0xfc, 0x8a, 0x82, 0x36, 0x45,
	0xf8, 0xb1, 0x4c, 0x05, 0x25, 0xa1, 0x27, 0xe8, 0x29, 0x11, 0x61, 0x4e, 0x75, 0x0a, 0xc7, 0x1e,
	0x4f, 0x53, 0x2e, 0x19, 0x50, 0x2f, 0x64, 0x82, 0x06, 0xc0, 0x78, 0xe2, 0xf1, 0x81, 0x17, 0x11,
	0x09, 0x1e, 0x08, 0x32, 0xa6, 0x42, 0x92, 0x48, 0xdf, 0x68, 0x6d, 0xb4, 0xeb, 0xbd, 0x3b, 0x96,
	0x9a, 0x7a, 0x3e, 0xd1, 0x6a, 0x2f, 0x39, 0xc3, 0x53, 0xce, 0x12, 0x67, 0x3f, 0xe7, 0xfe, 0xf2,
	0xa3, 0x79, 0xff, 0x72, 0xdc, 0x79, 0x8f, 0x74, 0x7b, 0x25, 0x93, 0x5b, 0x20, 0x1d, 0x16, 0x44,
	0x47, 0x0a, 0xe8, 0xa0, 0xe2, 0x39, 0x1a, 0xbc, 0x24, 0x12, 0xfa, 0x15, 0x8c, 0xf6, 0x01, 0x6f,
	0x67, 0x29, 0xb0, 0x98, 0xe6, 0x80, 0xc1, 0x88, 0x0a, 0xa9, 0x6f, 0xb6, 0x50, 0xbb, 0xde, 0x7b,
	0x68, 0x5d, 0xea, 0x4d, 0xb2, 0x5e, 0x17, 0xdd, 0x7d, 0xd5, 0xec, 0x18, 0x39, 0xf8, 0xf9, 0xde,
	0xd6, 0xbc, 0x4d, 0xb7, 0x91, 0xad, 0xe8, 0x4d, 0x8e, 0x1b, 0xab, 0x0e, 0xda, 0x3b, 0xbc, 0x19,
	0x31, 0x09, 0x3a, 0x2a, 0xc6, 0xf4, 0xe0, 0x5f, 0x30, 0x9c, 0x1d, 0x45, 0x51, 0xaf, 0x96, 0x29,
	0xc1, 0x74, 0x0b, 0x5b, 0xf3, 0x33, 0xc2, 0x5b, 0x2b, 0x62, 0xed, 0x23, 0xc2, 0x37, 0x15, 0x67,
	0xb5, 0xc6, 0x0c, 0x24, 0x0b, 0xa9, 0x42, 0xf8, 0x0f, 0x9b, 0xda, 0x29, 0xf3, 0xd4, 0x8e, 0xca,
	0x34, 0xe7, 0xed, 0x74, 0x6e, 0xa0, 0xb3, 0xb9, 0x81, 0x7e, 0xce, 0x0d, 0xf4, 0x69, 0x61, 0xd4,
	0xce, 0x16, 0x46, 0xed, 0xdb, 0xc2, 0xa8, 0xbd, 0x71, 0x2e, 0x78, 0xab, 0x71, 0x74, 0x22, 0xe2,
	0xcb, 0xea, 0x60, 0x8f, 0xbb, 0x8f, 0xec, 0xf7, 0x7f, 0xfa, 0x43, 0xc4, 0x3c, 0xa4, 0x91, 0x7f,
	0xb5, 0xf8, 0x4e, 0xf7, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb5, 0x8c, 0xe7, 0x50, 0x04,
	0x00, 0x00,
}

func (m *TickInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TickInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.UptimeTrackers.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTickInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.SpreadRewardGrowthOppositeDirectionOfLastTraversal) > 0 {
		for iNdEx := len(m.SpreadRewardGrowthOppositeDirectionOfLastTraversal) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpreadRewardGrowthOppositeDirectionOfLastTraversal[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTickInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.LiquidityNet.Size()
		i -= size
		if _, err := m.LiquidityNet.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTickInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.LiquidityGross.Size()
		i -= size
		if _, err := m.LiquidityGross.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTickInfo(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *UptimeTrackers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UptimeTrackers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UptimeTrackers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.List) > 0 {
		for iNdEx := len(m.List) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.List[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTickInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *UptimeTracker) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UptimeTracker) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UptimeTracker) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.UptimeGrowthOutside) > 0 {
		for iNdEx := len(m.UptimeGrowthOutside) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UptimeGrowthOutside[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTickInfo(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTickInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovTickInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TickInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LiquidityGross.Size()
	n += 1 + l + sovTickInfo(uint64(l))
	l = m.LiquidityNet.Size()
	n += 1 + l + sovTickInfo(uint64(l))
	if len(m.SpreadRewardGrowthOppositeDirectionOfLastTraversal) > 0 {
		for _, e := range m.SpreadRewardGrowthOppositeDirectionOfLastTraversal {
			l = e.Size()
			n += 1 + l + sovTickInfo(uint64(l))
		}
	}
	l = m.UptimeTrackers.Size()
	n += 1 + l + sovTickInfo(uint64(l))
	return n
}

func (m *UptimeTrackers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.List) > 0 {
		for _, e := range m.List {
			l = e.Size()
			n += 1 + l + sovTickInfo(uint64(l))
		}
	}
	return n
}

func (m *UptimeTracker) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.UptimeGrowthOutside) > 0 {
		for _, e := range m.UptimeGrowthOutside {
			l = e.Size()
			n += 1 + l + sovTickInfo(uint64(l))
		}
	}
	return n
}

func sovTickInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTickInfo(x uint64) (n int) {
	return sovTickInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TickInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTickInfo
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
			return fmt.Errorf("proto: TickInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityGross", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickInfo
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
				return ErrInvalidLengthTickInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTickInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityGross.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityNet", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickInfo
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
				return ErrInvalidLengthTickInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTickInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidityNet.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpreadRewardGrowthOppositeDirectionOfLastTraversal", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickInfo
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
				return ErrInvalidLengthTickInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpreadRewardGrowthOppositeDirectionOfLastTraversal = append(m.SpreadRewardGrowthOppositeDirectionOfLastTraversal, types.DecCoin{})
			if err := m.SpreadRewardGrowthOppositeDirectionOfLastTraversal[len(m.SpreadRewardGrowthOppositeDirectionOfLastTraversal)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UptimeTrackers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickInfo
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
				return ErrInvalidLengthTickInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.UptimeTrackers.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTickInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTickInfo
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
func (m *UptimeTrackers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTickInfo
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
			return fmt.Errorf("proto: UptimeTrackers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UptimeTrackers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field List", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickInfo
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
				return ErrInvalidLengthTickInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.List = append(m.List, UptimeTracker{})
			if err := m.List[len(m.List)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTickInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTickInfo
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
func (m *UptimeTracker) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTickInfo
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
			return fmt.Errorf("proto: UptimeTracker: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UptimeTracker: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UptimeGrowthOutside", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTickInfo
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
				return ErrInvalidLengthTickInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTickInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UptimeGrowthOutside = append(m.UptimeGrowthOutside, types.DecCoin{})
			if err := m.UptimeGrowthOutside[len(m.UptimeGrowthOutside)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTickInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTickInfo
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
func skipTickInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTickInfo
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
					return 0, ErrIntOverflowTickInfo
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
					return 0, ErrIntOverflowTickInfo
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
				return 0, ErrInvalidLengthTickInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTickInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTickInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTickInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTickInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTickInfo = fmt.Errorf("proto: unexpected end of group")
)
