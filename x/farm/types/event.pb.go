// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/farm/v1beta1/event.proto

package types

import (
	fmt "fmt"
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

type EventCreatePrivatePlan struct {
	Creator            string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	PlanId             uint64 `protobuf:"varint,2,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
	FarmingPoolAddress string `protobuf:"bytes,3,opt,name=farming_pool_address,json=farmingPoolAddress,proto3" json:"farming_pool_address,omitempty"`
}

func (m *EventCreatePrivatePlan) Reset()         { *m = EventCreatePrivatePlan{} }
func (m *EventCreatePrivatePlan) String() string { return proto.CompactTextString(m) }
func (*EventCreatePrivatePlan) ProtoMessage()    {}
func (*EventCreatePrivatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a95bdaaa134d6ec, []int{0}
}
func (m *EventCreatePrivatePlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventCreatePrivatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventCreatePrivatePlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventCreatePrivatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCreatePrivatePlan.Merge(m, src)
}
func (m *EventCreatePrivatePlan) XXX_Size() int {
	return m.Size()
}
func (m *EventCreatePrivatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCreatePrivatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventCreatePrivatePlan proto.InternalMessageInfo

type EventFarm struct {
	Farmer         string                                   `protobuf:"bytes,1,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Coin           types.Coin                               `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	WithdrawnCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=withdrawn_coins,json=withdrawnCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_coins"`
}

func (m *EventFarm) Reset()         { *m = EventFarm{} }
func (m *EventFarm) String() string { return proto.CompactTextString(m) }
func (*EventFarm) ProtoMessage()    {}
func (*EventFarm) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a95bdaaa134d6ec, []int{1}
}
func (m *EventFarm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFarm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFarm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFarm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFarm.Merge(m, src)
}
func (m *EventFarm) XXX_Size() int {
	return m.Size()
}
func (m *EventFarm) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFarm.DiscardUnknown(m)
}

var xxx_messageInfo_EventFarm proto.InternalMessageInfo

type EventUnfarm struct {
	Farmer         string                                   `protobuf:"bytes,1,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Coin           types.Coin                               `protobuf:"bytes,2,opt,name=coin,proto3" json:"coin"`
	WithdrawnCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=withdrawn_coins,json=withdrawnCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_coins"`
}

func (m *EventUnfarm) Reset()         { *m = EventUnfarm{} }
func (m *EventUnfarm) String() string { return proto.CompactTextString(m) }
func (*EventUnfarm) ProtoMessage()    {}
func (*EventUnfarm) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a95bdaaa134d6ec, []int{2}
}
func (m *EventUnfarm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventUnfarm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventUnfarm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventUnfarm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventUnfarm.Merge(m, src)
}
func (m *EventUnfarm) XXX_Size() int {
	return m.Size()
}
func (m *EventUnfarm) XXX_DiscardUnknown() {
	xxx_messageInfo_EventUnfarm.DiscardUnknown(m)
}

var xxx_messageInfo_EventUnfarm proto.InternalMessageInfo

type EventHarvest struct {
	Farmer         string                                   `protobuf:"bytes,1,opt,name=farmer,proto3" json:"farmer,omitempty"`
	Denom          string                                   `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	WithdrawnCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=withdrawn_coins,json=withdrawnCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"withdrawn_coins"`
}

func (m *EventHarvest) Reset()         { *m = EventHarvest{} }
func (m *EventHarvest) String() string { return proto.CompactTextString(m) }
func (*EventHarvest) ProtoMessage()    {}
func (*EventHarvest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a95bdaaa134d6ec, []int{3}
}
func (m *EventHarvest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventHarvest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventHarvest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventHarvest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHarvest.Merge(m, src)
}
func (m *EventHarvest) XXX_Size() int {
	return m.Size()
}
func (m *EventHarvest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHarvest.DiscardUnknown(m)
}

var xxx_messageInfo_EventHarvest proto.InternalMessageInfo

type EventTerminatePlan struct {
	PlanId uint64 `protobuf:"varint,1,opt,name=plan_id,json=planId,proto3" json:"plan_id,omitempty"`
}

func (m *EventTerminatePlan) Reset()         { *m = EventTerminatePlan{} }
func (m *EventTerminatePlan) String() string { return proto.CompactTextString(m) }
func (*EventTerminatePlan) ProtoMessage()    {}
func (*EventTerminatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a95bdaaa134d6ec, []int{4}
}
func (m *EventTerminatePlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventTerminatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventTerminatePlan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventTerminatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventTerminatePlan.Merge(m, src)
}
func (m *EventTerminatePlan) XXX_Size() int {
	return m.Size()
}
func (m *EventTerminatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_EventTerminatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_EventTerminatePlan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EventCreatePrivatePlan)(nil), "squad.farm.v1beta1.EventCreatePrivatePlan")
	proto.RegisterType((*EventFarm)(nil), "squad.farm.v1beta1.EventFarm")
	proto.RegisterType((*EventUnfarm)(nil), "squad.farm.v1beta1.EventUnfarm")
	proto.RegisterType((*EventHarvest)(nil), "squad.farm.v1beta1.EventHarvest")
	proto.RegisterType((*EventTerminatePlan)(nil), "squad.farm.v1beta1.EventTerminatePlan")
}

func init() { proto.RegisterFile("squad/farm/v1beta1/event.proto", fileDescriptor_0a95bdaaa134d6ec) }

var fileDescriptor_0a95bdaaa134d6ec = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x53, 0x4d, 0x8e, 0xd3, 0x30,
	0x18, 0x8d, 0x99, 0xd2, 0x51, 0x3d, 0x08, 0x24, 0xab, 0x1a, 0xc2, 0x2c, 0x3c, 0xa3, 0xae, 0xba,
	0x49, 0x3c, 0xc3, 0x9c, 0x80, 0x19, 0x81, 0x40, 0x6c, 0x46, 0x11, 0x6c, 0xd8, 0x44, 0x4e, 0xec,
	0xa6, 0x11, 0x89, 0x1d, 0x6c, 0x37, 0x85, 0x05, 0x77, 0xe0, 0x1c, 0x2c, 0x39, 0x45, 0x57, 0xa8,
	0x4b, 0x56, 0xfc, 0xb4, 0x17, 0x41, 0xb6, 0x93, 0xaa, 0x12, 0x82, 0x2d, 0xd2, 0xac, 0x9c, 0xe7,
	0xf7, 0x7d, 0xfe, 0xde, 0x8b, 0xbe, 0x07, 0xb1, 0x7e, 0xb7, 0xa0, 0x8c, 0xcc, 0xa8, 0xaa, 0x49,
	0x7b, 0x91, 0x71, 0x43, 0x2f, 0x08, 0x6f, 0xb9, 0x30, 0x71, 0xa3, 0xa4, 0x91, 0x08, 0x39, 0x3e,
	0xb6, 0x7c, 0xdc, 0xf1, 0x27, 0xe3, 0x42, 0x16, 0xd2, 0xd1, 0xc4, 0x7e, 0xf9, 0xca, 0x13, 0x9c,
	0x4b, 0x5d, 0x4b, 0x4d, 0x32, 0xaa, 0xf9, 0xee, 0xa9, 0x5c, 0x96, 0xc2, 0xf3, 0x93, 0x8f, 0xf0,
	0xf8, 0xa9, 0x7d, 0xf8, 0x5a, 0x71, 0x6a, 0xf8, 0x8d, 0x2a, 0x5b, 0x7b, 0x54, 0x54, 0xa0, 0x10,
	0x1e, 0xe6, 0xf6, 0x52, 0xaa, 0x10, 0x9c, 0x81, 0xe9, 0x28, 0xe9, 0x21, 0x7a, 0x08, 0x0f, 0x9b,
	0x8a, 0x8a, 0xb4, 0x64, 0xe1, 0x9d, 0x33, 0x30, 0x1d, 0x24, 0x43, 0x0b, 0x5f, 0x30, 0x74, 0x0e,
	0xc7, 0x56, 0x52, 0x29, 0x8a, 0xb4, 0x91, 0xb2, 0x4a, 0x29, 0x63, 0x8a, 0x6b, 0x1d, 0x1e, 0xb8,
	0x7e, 0xd4, 0x71, 0x37, 0x52, 0x56, 0x4f, 0x3c, 0x33, 0xf9, 0x0a, 0xe0, 0xc8, 0xcd, 0x7f, 0x46,
	0x55, 0x8d, 0x8e, 0xe1, 0xd0, 0xd6, 0xf0, 0x7e, 0x62, 0x87, 0xd0, 0x25, 0x1c, 0x58, 0xc9, 0x6e,
	0xda, 0xd1, 0xe3, 0x47, 0xb1, 0xf7, 0x14, 0x5b, 0x4f, 0xbd, 0xfd, 0xf8, 0x5a, 0x96, 0xe2, 0x6a,
	0xb0, 0xfa, 0x7e, 0x1a, 0x24, 0xae, 0x18, 0x19, 0xf8, 0x60, 0x59, 0x9a, 0x39, 0x53, 0x74, 0x29,
	0x52, 0x7b, 0x63, 0x75, 0x1c, 0xfc, 0xbb, 0xff, 0xdc, 0xf6, 0x7f, 0xfe, 0x71, 0x3a, 0x2d, 0x4a,
	0x33, 0x5f, 0x64, 0x71, 0x2e, 0x6b, 0xd2, 0xfd, 0x40, 0x7f, 0x44, 0x9a, 0xbd, 0x25, 0xe6, 0x43,
	0xc3, 0xb5, 0x6b, 0xd0, 0xc9, 0xfd, 0xdd, 0x0c, 0x87, 0x27, 0x6b, 0x00, 0x8f, 0x9c, 0xa1, 0xd7,
	0x62, 0x76, 0x4b, 0x2c, 0x7d, 0x01, 0xf0, 0x9e, 0xb3, 0xf4, 0x9c, 0xaa, 0x96, 0x6b, 0xf3, 0x57,
	0x4f, 0x63, 0x78, 0x97, 0x71, 0x21, 0x6b, 0x67, 0x6a, 0x94, 0x78, 0xf0, 0x9f, 0x44, 0x47, 0x10,
	0x39, 0xcd, 0xaf, 0xb8, 0xdd, 0xb9, 0x7e, 0xa7, 0xf7, 0x36, 0x17, 0xec, 0x6f, 0xee, 0xd5, 0xcb,
	0xd5, 0x2f, 0x1c, 0xac, 0x36, 0x18, 0xac, 0x37, 0x18, 0xfc, 0xdc, 0x60, 0xf0, 0x69, 0x8b, 0x83,
	0xf5, 0x16, 0x07, 0xdf, 0xb6, 0x38, 0x78, 0x13, 0xfd, 0x21, 0xc3, 0xc6, 0x2f, 0xaa, 0x68, 0xa6,
	0x89, 0x4f, 0xea, 0x7b, 0x9f, 0x55, 0xa7, 0x28, 0x1b, 0xba, 0x68, 0x5d, 0xfe, 0x0e, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x03, 0x62, 0xad, 0xc6, 0x03, 0x00, 0x00,
}

func (m *EventCreatePrivatePlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventCreatePrivatePlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventCreatePrivatePlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FarmingPoolAddress) > 0 {
		i -= len(m.FarmingPoolAddress)
		copy(dAtA[i:], m.FarmingPoolAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.FarmingPoolAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PlanId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PlanId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventFarm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFarm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFarm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnCoins) > 0 {
		for iNdEx := len(m.WithdrawnCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventUnfarm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventUnfarm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventUnfarm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnCoins) > 0 {
		for iNdEx := len(m.WithdrawnCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.Coin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventHarvest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventHarvest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventHarvest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.WithdrawnCoins) > 0 {
		for iNdEx := len(m.WithdrawnCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.WithdrawnCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Farmer) > 0 {
		i -= len(m.Farmer)
		copy(dAtA[i:], m.Farmer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Farmer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventTerminatePlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventTerminatePlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventTerminatePlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PlanId != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PlanId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventCreatePrivatePlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.PlanId != 0 {
		n += 1 + sovEvent(uint64(m.PlanId))
	}
	l = len(m.FarmingPoolAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EventFarm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovEvent(uint64(l))
	if len(m.WithdrawnCoins) > 0 {
		for _, e := range m.WithdrawnCoins {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *EventUnfarm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.Coin.Size()
	n += 1 + l + sovEvent(uint64(l))
	if len(m.WithdrawnCoins) > 0 {
		for _, e := range m.WithdrawnCoins {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *EventHarvest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Farmer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if len(m.WithdrawnCoins) > 0 {
		for _, e := range m.WithdrawnCoins {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *EventTerminatePlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PlanId != 0 {
		n += 1 + sovEvent(uint64(m.PlanId))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventCreatePrivatePlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventCreatePrivatePlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventCreatePrivatePlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanId", wireType)
			}
			m.PlanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FarmingPoolAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FarmingPoolAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventFarm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventFarm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFarm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnCoins = append(m.WithdrawnCoins, types.Coin{})
			if err := m.WithdrawnCoins[len(m.WithdrawnCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventUnfarm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventUnfarm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventUnfarm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Coin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnCoins = append(m.WithdrawnCoins, types.Coin{})
			if err := m.WithdrawnCoins[len(m.WithdrawnCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventHarvest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventHarvest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventHarvest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Farmer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Farmer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithdrawnCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WithdrawnCoins = append(m.WithdrawnCoins, types.Coin{})
			if err := m.WithdrawnCoins[len(m.WithdrawnCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventTerminatePlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventTerminatePlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventTerminatePlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlanId", wireType)
			}
			m.PlanId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlanId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
