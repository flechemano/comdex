// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/asset/v1beta1/extendedAsset.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type ExtendedAsset struct {
	Id                   uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AssetId              uint64                                 `protobuf:"varint,2,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty" yaml:"asset_id"`
	CollateralWeight     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=collateral_weight,json=collateralWeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"collateral_weight" yaml:"collateral_weight"`
	LiquidationThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=liquidation_threshold,json=liquidationThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidation_threshold" yaml:"liquidation_threshold"`
	// The rate indicates the price of the asset in USD
	Rate           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=rate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"rate" yaml:"rate"`
	IsBridgedAsset bool                                   `protobuf:"varint,6,opt,name=is_bridged_asset,json=isBridgedAsset,proto3" json:"is_bridged_asset,omitempty" yaml:"is_bridged_asset"`
}

func (m *ExtendedAsset) Reset()         { *m = ExtendedAsset{} }
func (m *ExtendedAsset) String() string { return proto.CompactTextString(m) }
func (*ExtendedAsset) ProtoMessage()    {}
func (*ExtendedAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_f9f1ce091d29adf2, []int{0}
}
func (m *ExtendedAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtendedAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtendedAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtendedAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtendedAsset.Merge(m, src)
}
func (m *ExtendedAsset) XXX_Size() int {
	return m.Size()
}
func (m *ExtendedAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtendedAsset.DiscardUnknown(m)
}

var xxx_messageInfo_ExtendedAsset proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtendedAsset)(nil), "comdex.asset.v1beta1.ExtendedAsset")
}

func init() {
	proto.RegisterFile("comdex/asset/v1beta1/extendedAsset.proto", fileDescriptor_f9f1ce091d29adf2)
}

var fileDescriptor_f9f1ce091d29adf2 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0xae, 0xd3, 0x30,
	0x14, 0xc6, 0xe3, 0xde, 0x72, 0xb9, 0x18, 0x51, 0x2e, 0xa1, 0x88, 0x08, 0x21, 0xa7, 0xca, 0x80,
	0xb2, 0x34, 0x56, 0xc5, 0x86, 0xc4, 0x40, 0x04, 0x43, 0x19, 0x90, 0x1a, 0x21, 0x55, 0x62, 0x89,
	0x9c, 0xd8, 0x4d, 0x2c, 0x92, 0xba, 0xc4, 0x2e, 0x6d, 0x67, 0x5e, 0x80, 0xc7, 0xe0, 0x51, 0x3a,
	0x76, 0x44, 0x0c, 0x11, 0xa4, 0x6f, 0xd0, 0x91, 0x09, 0xc5, 0x69, 0x45, 0x0b, 0x77, 0xe9, 0x94,
	0x93, 0xef, 0x7c, 0xfe, 0x7e, 0xfe, 0x73, 0xa0, 0x1b, 0x8b, 0x9c, 0xb2, 0x25, 0x26, 0x52, 0x32,
	0x85, 0x3f, 0x0f, 0x22, 0xa6, 0xc8, 0x00, 0xb3, 0xa5, 0x62, 0x53, 0xca, 0xe8, 0xab, 0x5a, 0xf5,
	0x66, 0x85, 0x50, 0xc2, 0xec, 0x36, 0x4e, 0x4f, 0x3b, 0xbd, 0xbd, 0xf3, 0x49, 0x37, 0x11, 0x89,
	0xd0, 0x06, 0x5c, 0x57, 0x8d, 0xd7, 0xf9, 0x7d, 0x01, 0xef, 0xbd, 0x39, 0xce, 0x30, 0x3b, 0xb0,
	0xc5, 0xa9, 0x05, 0x7a, 0xc0, 0x6d, 0x07, 0x2d, 0x4e, 0x4d, 0x0f, 0x5e, 0xe9, 0xa0, 0x90, 0x53,
	0xab, 0x55, 0xab, 0xfe, 0xc3, 0x5d, 0x69, 0xdf, 0x5f, 0x91, 0x3c, 0x7b, 0xe1, 0x1c, 0x3a, 0x4e,
	0x70, 0x5b, 0x97, 0x43, 0x6a, 0x2e, 0xe0, 0x83, 0x58, 0x64, 0x19, 0x51, 0xac, 0x20, 0x59, 0xb8,
	0x60, 0x3c, 0x49, 0x95, 0x75, 0xd1, 0x03, 0xee, 0x1d, 0xff, 0xed, 0xba, 0xb4, 0x8d, 0x1f, 0xa5,
	0xfd, 0x2c, 0xe1, 0x2a, 0x9d, 0x47, 0x5e, 0x2c, 0x72, 0x1c, 0x0b, 0x99, 0x0b, 0xb9, 0xff, 0xf4,
	0x25, 0xfd, 0x88, 0xd5, 0x6a, 0xc6, 0xa4, 0xf7, 0x9a, 0xc5, 0xbb, 0xd2, 0xb6, 0x1a, 0xcc, 0x7f,
	0x81, 0x4e, 0x70, 0xfd, 0x57, 0x1b, 0x6b, 0xc9, 0xfc, 0x02, 0xe0, 0xa3, 0x8c, 0x7f, 0x9a, 0x73,
	0x4a, 0x14, 0x17, 0xd3, 0x50, 0xa5, 0x05, 0x93, 0xa9, 0xc8, 0xa8, 0xd5, 0xd6, 0xf4, 0x77, 0x67,
	0xd3, 0x9f, 0x36, 0xf4, 0x1b, 0x43, 0x9d, 0xa0, 0x7b, 0xa4, 0xbf, 0x3f, 0xc8, 0xe6, 0x08, 0xb6,
	0x0b, 0xa2, 0x98, 0x75, 0x4b, 0x33, 0x5f, 0x9e, 0xcd, 0xbc, 0xdb, 0x30, 0xeb, 0x0c, 0x27, 0xd0,
	0x51, 0xe6, 0x18, 0x5e, 0x73, 0x19, 0x46, 0x05, 0xa7, 0x09, 0xa3, 0xa1, 0xbe, 0x67, 0xeb, 0xb2,
	0x07, 0xdc, 0x2b, 0xbf, 0x5f, 0x95, 0x76, 0x67, 0x28, 0xfd, 0xa6, 0xa5, 0xdf, 0x6f, 0x57, 0xda,
	0x8f, 0x9b, 0x88, 0x7f, 0xd7, 0x38, 0x41, 0x87, 0x9f, 0x58, 0xfd, 0xd1, 0xfa, 0x17, 0x32, 0xbe,
	0x55, 0xc8, 0x58, 0x57, 0x08, 0x6c, 0x2a, 0x04, 0x7e, 0x56, 0x08, 0x7c, 0xdd, 0x22, 0x63, 0xb3,
	0x45, 0xc6, 0xf7, 0x2d, 0x32, 0x3e, 0xe0, 0x93, 0x7d, 0xd7, 0x53, 0xd5, 0x17, 0x93, 0x09, 0x8f,
	0x39, 0xc9, 0xf6, 0xff, 0xf8, 0x30, 0x91, 0xfa, 0x10, 0xd1, 0xa5, 0x1e, 0xab, 0xe7, 0x7f, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x19, 0xb2, 0xb2, 0x46, 0xae, 0x02, 0x00, 0x00,
}

func (m *ExtendedAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtendedAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtendedAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsBridgedAsset {
		i--
		if m.IsBridgedAsset {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.Rate.Size()
		i -= size
		if _, err := m.Rate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.LiquidationThreshold.Size()
		i -= size
		if _, err := m.LiquidationThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.CollateralWeight.Size()
		i -= size
		if _, err := m.CollateralWeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintExtendedAsset(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AssetId != 0 {
		i = encodeVarintExtendedAsset(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintExtendedAsset(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintExtendedAsset(dAtA []byte, offset int, v uint64) int {
	offset -= sovExtendedAsset(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtendedAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovExtendedAsset(uint64(m.Id))
	}
	if m.AssetId != 0 {
		n += 1 + sovExtendedAsset(uint64(m.AssetId))
	}
	l = m.CollateralWeight.Size()
	n += 1 + l + sovExtendedAsset(uint64(l))
	l = m.LiquidationThreshold.Size()
	n += 1 + l + sovExtendedAsset(uint64(l))
	l = m.Rate.Size()
	n += 1 + l + sovExtendedAsset(uint64(l))
	if m.IsBridgedAsset {
		n += 2
	}
	return n
}

func sovExtendedAsset(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExtendedAsset(x uint64) (n int) {
	return sovExtendedAsset(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtendedAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExtendedAsset
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
			return fmt.Errorf("proto: ExtendedAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtendedAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralWeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedAsset
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
				return ErrInvalidLengthExtendedAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CollateralWeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidationThreshold", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedAsset
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
				return ErrInvalidLengthExtendedAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LiquidationThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedAsset
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
				return ErrInvalidLengthExtendedAsset
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthExtendedAsset
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Rate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsBridgedAsset", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExtendedAsset
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsBridgedAsset = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipExtendedAsset(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthExtendedAsset
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
func skipExtendedAsset(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExtendedAsset
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
					return 0, ErrIntOverflowExtendedAsset
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
					return 0, ErrIntOverflowExtendedAsset
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
				return 0, ErrInvalidLengthExtendedAsset
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupExtendedAsset
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthExtendedAsset
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthExtendedAsset        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExtendedAsset          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupExtendedAsset = fmt.Errorf("proto: unexpected end of group")
)
