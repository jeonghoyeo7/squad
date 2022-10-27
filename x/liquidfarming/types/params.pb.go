// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: squad/liquidfarming/v1beta1/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
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

// Params defines the parameters for the module.
type Params struct {
	FeeCollector           string        `protobuf:"bytes,1,opt,name=fee_collector,json=feeCollector,proto3" json:"fee_collector,omitempty"`
	RewardsAuctionDuration time.Duration `protobuf:"bytes,2,opt,name=rewards_auction_duration,json=rewardsAuctionDuration,proto3,stdduration" json:"rewards_auction_duration"`
	LiquidFarms            []LiquidFarm  `protobuf:"bytes,3,rep,name=liquid_farms,json=liquidFarms,proto3" json:"liquid_farms"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_6012e16b27fcc811, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

// LiquidFarm defines liquid farm object that provides auto compounding functionality
// for the liquidity pool and undergoes farming rewards auction process.
// See the technical spec for more detailed information.
type LiquidFarm struct {
	PoolId        uint64                                 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	MinFarmAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=min_farm_amount,json=minFarmAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_farm_amount"`
	MinBidAmount  github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=min_bid_amount,json=minBidAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"min_bid_amount"`
	FeeRate       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate"`
}

func (m *LiquidFarm) Reset()      { *m = LiquidFarm{} }
func (*LiquidFarm) ProtoMessage() {}
func (*LiquidFarm) Descriptor() ([]byte, []int) {
	return fileDescriptor_6012e16b27fcc811, []int{1}
}
func (m *LiquidFarm) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LiquidFarm) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LiquidFarm.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LiquidFarm) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiquidFarm.Merge(m, src)
}
func (m *LiquidFarm) XXX_Size() int {
	return m.Size()
}
func (m *LiquidFarm) XXX_DiscardUnknown() {
	xxx_messageInfo_LiquidFarm.DiscardUnknown(m)
}

var xxx_messageInfo_LiquidFarm proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "squad.liquidfarming.v1beta1.Params")
	proto.RegisterType((*LiquidFarm)(nil), "squad.liquidfarming.v1beta1.LiquidFarm")
}

func init() {
	proto.RegisterFile("squad/liquidfarming/v1beta1/params.proto", fileDescriptor_6012e16b27fcc811)
}

var fileDescriptor_6012e16b27fcc811 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x6b, 0xd4, 0x40,
	0x18, 0xc6, 0x93, 0xee, 0xb2, 0xdd, 0xce, 0x6e, 0x15, 0x82, 0x68, 0xac, 0x90, 0x2c, 0x15, 0x74,
	0x2f, 0x9d, 0xa1, 0x15, 0x3c, 0xf4, 0xd6, 0x58, 0x84, 0x05, 0x0f, 0x4b, 0x10, 0x0f, 0x82, 0x84,
	0x49, 0x66, 0x12, 0x07, 0x33, 0x99, 0x74, 0x66, 0xe2, 0x9f, 0x6f, 0xe0, 0xd1, 0x63, 0x8f, 0x3d,
	0xfa, 0x51, 0x7a, 0xec, 0x51, 0x3c, 0x54, 0xd9, 0xf5, 0x83, 0xc8, 0x4c, 0x92, 0x8a, 0x0a, 0x82,
	0x3d, 0x65, 0xf2, 0xf2, 0xcc, 0xef, 0x79, 0x9f, 0x77, 0x5e, 0x30, 0x57, 0x27, 0x0d, 0x26, 0xa8,
	0x64, 0x27, 0x0d, 0x23, 0x39, 0x96, 0x9c, 0x55, 0x05, 0x7a, 0xbb, 0x9f, 0x52, 0x8d, 0xf7, 0x51,
	0x8d, 0x25, 0xe6, 0x0a, 0xd6, 0x52, 0x68, 0xe1, 0xdd, 0xb3, 0x4a, 0xf8, 0x9b, 0x12, 0x76, 0xca,
	0x9d, 0xa0, 0x10, 0xa2, 0x28, 0x29, 0xb2, 0xd2, 0xb4, 0xc9, 0x11, 0x69, 0x24, 0xd6, 0x4c, 0x54,
	0xed, 0xe5, 0x9d, 0x5b, 0x85, 0x28, 0x84, 0x3d, 0x22, 0x73, 0xea, 0xaa, 0x41, 0x26, 0x14, 0x17,
	0x0a, 0xa5, 0x58, 0xd1, 0x2b, 0xd3, 0x4c, 0xb0, 0xee, 0xd6, 0xee, 0x0f, 0x17, 0x8c, 0x96, 0xb6,
	0x07, 0xef, 0x3e, 0xd8, 0xce, 0x29, 0x4d, 0x32, 0x51, 0x96, 0x34, 0xd3, 0x42, 0xfa, 0xee, 0xcc,
	0x9d, 0x6f, 0xc5, 0xd3, 0x9c, 0xd2, 0x27, 0x7d, 0xcd, 0x7b, 0x05, 0x7c, 0x49, 0xdf, 0x61, 0x49,
	0x54, 0x82, 0x9b, 0xcc, 0xd8, 0x27, 0x7d, 0x1f, 0xfe, 0xc6, 0xcc, 0x9d, 0x4f, 0x0e, 0xee, 0xc2,
	0xb6, 0x51, 0xd8, 0x37, 0x0a, 0x8f, 0x3b, 0x41, 0x34, 0x3e, 0xbf, 0x0c, 0x9d, 0xd3, 0x6f, 0xa1,
	0x1b, 0xdf, 0xee, 0x20, 0x47, 0x2d, 0xa3, 0x57, 0x78, 0x4b, 0x30, 0x6d, 0xd3, 0x27, 0x26, 0xbe,
	0xf2, 0x07, 0xb3, 0xc1, 0x7c, 0x72, 0xf0, 0x10, 0xfe, 0x63, 0x30, 0xf0, 0x99, 0xad, 0x3e, 0xc5,
	0x92, 0x47, 0x43, 0x63, 0x10, 0x4f, 0xca, 0xab, 0x8a, 0x3a, 0x1c, 0x7e, 0x3c, 0x0b, 0x9d, 0xdd,
	0xcf, 0x1b, 0x00, 0xfc, 0xd2, 0x79, 0x77, 0xc0, 0x66, 0x2d, 0x44, 0x99, 0x30, 0x62, 0x43, 0x0e,
	0xe3, 0x91, 0xf9, 0x5d, 0x10, 0xef, 0x05, 0xb8, 0xc9, 0x59, 0x65, 0xcd, 0x13, 0xcc, 0x45, 0x53,
	0x69, 0x9b, 0x6a, 0x2b, 0x82, 0x86, 0xfc, 0xf5, 0x32, 0x7c, 0x50, 0x30, 0xfd, 0xba, 0x49, 0x61,
	0x26, 0x38, 0xea, 0x46, 0xdb, 0x7e, 0xf6, 0x14, 0x79, 0x83, 0xf4, 0x87, 0x9a, 0x2a, 0xb8, 0xa8,
	0x74, 0xbc, 0xcd, 0x59, 0x65, 0xac, 0x8e, 0x2c, 0xc4, 0x7b, 0x0e, 0x6e, 0x18, 0x6e, 0xca, 0x48,
	0x8f, 0x1d, 0x5c, 0x0b, 0x3b, 0xe5, 0xac, 0x8a, 0x18, 0xe9, 0xa8, 0x0b, 0x30, 0x36, 0x2f, 0x26,
	0xb1, 0xa6, 0xfe, 0xf0, 0xbf, 0x79, 0xc7, 0x34, 0x8b, 0x37, 0x73, 0x4a, 0x63, 0xac, 0xe9, 0xe1,
	0xd8, 0x8c, 0xe9, 0xf4, 0x2c, 0x74, 0xa2, 0xe5, 0xf9, 0x2a, 0x70, 0x2f, 0x56, 0x81, 0xfb, 0x7d,
	0x15, 0xb8, 0x9f, 0xd6, 0x81, 0x73, 0xb1, 0x0e, 0x9c, 0x2f, 0xeb, 0xc0, 0x79, 0xf9, 0xf8, 0x2f,
	0xa8, 0x79, 0x95, 0xbd, 0x12, 0xa7, 0x0a, 0xb5, 0x3b, 0xfe, 0xfe, 0x8f, 0x2d, 0xb7, 0x46, 0xe9,
	0xc8, 0x6e, 0xc2, 0xa3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x03, 0x0a, 0x90, 0x09, 0x03,
	0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.LiquidFarms) > 0 {
		for iNdEx := len(m.LiquidFarms) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LiquidFarms[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.RewardsAuctionDuration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.RewardsAuctionDuration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintParams(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.FeeCollector) > 0 {
		i -= len(m.FeeCollector)
		copy(dAtA[i:], m.FeeCollector)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FeeCollector)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *LiquidFarm) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LiquidFarm) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LiquidFarm) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.MinBidAmount.Size()
		i -= size
		if _, err := m.MinBidAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.MinFarmAmount.Size()
		i -= size
		if _, err := m.MinFarmAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.PoolId != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FeeCollector)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.RewardsAuctionDuration)
	n += 1 + l + sovParams(uint64(l))
	if len(m.LiquidFarms) > 0 {
		for _, e := range m.LiquidFarms {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	return n
}

func (m *LiquidFarm) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovParams(uint64(m.PoolId))
	}
	l = m.MinFarmAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.MinBidAmount.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.FeeRate.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeCollector", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeeCollector = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardsAuctionDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.RewardsAuctionDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidFarms", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LiquidFarms = append(m.LiquidFarms, LiquidFarm{})
			if err := m.LiquidFarms[len(m.LiquidFarms)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *LiquidFarm) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: LiquidFarm: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LiquidFarm: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinFarmAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinFarmAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinBidAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinBidAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
