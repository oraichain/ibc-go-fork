// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/packetserver/v1/counterparty.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	v2 "github.com/cosmos/ibc-go/v9/modules/core/23-commitment/types/v2"
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

// Counterparty defines the counterparty for a light client to implement IBC eureka protocol
type Counterparty struct {
	// the client identifier of the light client representing the counterparty chain
	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	// the counterparty identifier that must be used by the packet
	CounterpartyChannelId string `protobuf:"bytes,2,opt,name=counterparty_channel_id,json=counterpartyChannelId,proto3" json:"counterparty_channel_id,omitempty"`
	// the key path used to store packet flow messages that the counterparty
	// will use to send to us. In backwards compatible cases, we will append the channelID and sequence in order to create
	// the final path.
	MerklePathPrefix v2.MerklePath `protobuf:"bytes,3,opt,name=merkle_path_prefix,json=merklePathPrefix,proto3" json:"merkle_path_prefix"`
}

func (m *Counterparty) Reset()         { *m = Counterparty{} }
func (m *Counterparty) String() string { return proto.CompactTextString(m) }
func (*Counterparty) ProtoMessage()    {}
func (*Counterparty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0c60a0709a0040c, []int{0}
}
func (m *Counterparty) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Counterparty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Counterparty.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Counterparty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Counterparty.Merge(m, src)
}
func (m *Counterparty) XXX_Size() int {
	return m.Size()
}
func (m *Counterparty) XXX_DiscardUnknown() {
	xxx_messageInfo_Counterparty.DiscardUnknown(m)
}

var xxx_messageInfo_Counterparty proto.InternalMessageInfo

func (m *Counterparty) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *Counterparty) GetCounterpartyChannelId() string {
	if m != nil {
		return m.CounterpartyChannelId
	}
	return ""
}

func (m *Counterparty) GetMerklePathPrefix() v2.MerklePath {
	if m != nil {
		return m.MerklePathPrefix
	}
	return v2.MerklePath{}
}

func init() {
	proto.RegisterType((*Counterparty)(nil), "ibc.core.packetserver.v1.Counterparty")
}

func init() {
	proto.RegisterFile("ibc/core/packetserver/v1/counterparty.proto", fileDescriptor_e0c60a0709a0040c)
}

var fileDescriptor_e0c60a0709a0040c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x30,
	0x1c, 0xc6, 0x1b, 0x15, 0x71, 0xd5, 0x83, 0x14, 0xc5, 0x31, 0x21, 0x8e, 0x5d, 0x1c, 0xc8, 0x12,
	0x36, 0x41, 0x10, 0x3c, 0x6d, 0xa7, 0x1d, 0x84, 0xb1, 0xc3, 0x0e, 0x5e, 0x4a, 0x9b, 0xfe, 0x6d,
	0xc3, 0x9a, 0x26, 0xa4, 0x59, 0x70, 0x6f, 0xe1, 0xfb, 0xf8, 0x02, 0x3b, 0xee, 0xe8, 0x49, 0x64,
	0x7b, 0x11, 0x69, 0x2b, 0x23, 0xb7, 0xe4, 0xff, 0xfd, 0xf2, 0xe5, 0xfb, 0x7f, 0xfe, 0x03, 0x8f,
	0x19, 0x65, 0x52, 0x03, 0x55, 0x11, 0x5b, 0x82, 0x29, 0x41, 0x5b, 0xd0, 0xd4, 0x0e, 0x29, 0x93,
	0xab, 0xc2, 0x80, 0x56, 0x91, 0x36, 0x6b, 0xa2, 0xb4, 0x34, 0x32, 0x68, 0xf3, 0x98, 0x91, 0x0a,
	0x26, 0x2e, 0x4c, 0xec, 0xb0, 0x73, 0x95, 0xca, 0x54, 0xd6, 0x10, 0xad, 0x4e, 0x0d, 0xdf, 0xb9,
	0x3f, 0x98, 0x33, 0x29, 0x04, 0x37, 0x02, 0x0a, 0x43, 0xed, 0xc8, 0xb9, 0x35, 0x60, 0xef, 0x0b,
	0xf9, 0x17, 0x13, 0xe7, 0xbf, 0xe0, 0xd6, 0x6f, 0xb1, 0x9c, 0x43, 0x61, 0x42, 0x9e, 0xb4, 0x51,
	0x17, 0xf5, 0x5b, 0xf3, 0xb3, 0x66, 0x30, 0x4d, 0x82, 0x27, 0xff, 0xc6, 0x0d, 0x17, 0xb2, 0x2c,
	0x2a, 0x0a, 0xc8, 0x2b, 0xf4, 0xa8, 0x46, 0xaf, 0x5d, 0x79, 0xd2, 0xa8, 0xd3, 0x24, 0x58, 0xf8,
	0x81, 0x00, 0xbd, 0xcc, 0x21, 0x54, 0x91, 0xc9, 0x42, 0xa5, 0xe1, 0x9d, 0x7f, 0xb4, 0x8f, 0xbb,
	0xa8, 0x7f, 0x3e, 0xea, 0x91, 0xc3, 0x6e, 0x4e, 0x3a, 0x3b, 0x22, 0xaf, 0xf5, 0x8b, 0x59, 0x64,
	0xb2, 0xf1, 0xc9, 0xe6, 0xe7, 0xce, 0x9b, 0x5f, 0x8a, 0xc3, 0x64, 0x56, 0x3b, 0x8c, 0x17, 0x9b,
	0x1d, 0x46, 0xdb, 0x1d, 0x46, 0xbf, 0x3b, 0x8c, 0x3e, 0xf7, 0xd8, 0xdb, 0xee, 0xb1, 0xf7, 0xbd,
	0xc7, 0xde, 0xdb, 0x4b, 0xca, 0x4d, 0xb6, 0x8a, 0x2b, 0x4b, 0xca, 0x64, 0x29, 0x64, 0x49, 0x79,
	0xcc, 0x06, 0xa9, 0xa4, 0xf6, 0x99, 0x0a, 0x99, 0xac, 0x72, 0x28, 0xdd, 0xf6, 0x07, 0xff, 0xf5,
	0x9b, 0xb5, 0x82, 0x32, 0x3e, 0xad, 0xcb, 0x79, 0xfc, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x34, 0xa3,
	0x29, 0x7a, 0xa4, 0x01, 0x00, 0x00,
}

func (m *Counterparty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Counterparty) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Counterparty) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.MerklePathPrefix.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCounterparty(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.CounterpartyChannelId) > 0 {
		i -= len(m.CounterpartyChannelId)
		copy(dAtA[i:], m.CounterpartyChannelId)
		i = encodeVarintCounterparty(dAtA, i, uint64(len(m.CounterpartyChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClientId) > 0 {
		i -= len(m.ClientId)
		copy(dAtA[i:], m.ClientId)
		i = encodeVarintCounterparty(dAtA, i, uint64(len(m.ClientId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCounterparty(dAtA []byte, offset int, v uint64) int {
	offset -= sovCounterparty(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Counterparty) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientId)
	if l > 0 {
		n += 1 + l + sovCounterparty(uint64(l))
	}
	l = len(m.CounterpartyChannelId)
	if l > 0 {
		n += 1 + l + sovCounterparty(uint64(l))
	}
	l = m.MerklePathPrefix.Size()
	n += 1 + l + sovCounterparty(uint64(l))
	return n
}

func sovCounterparty(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCounterparty(x uint64) (n int) {
	return sovCounterparty(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Counterparty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCounterparty
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
			return fmt.Errorf("proto: Counterparty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Counterparty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounterparty
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
				return ErrInvalidLengthCounterparty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCounterparty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CounterpartyChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounterparty
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
				return ErrInvalidLengthCounterparty
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCounterparty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CounterpartyChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerklePathPrefix", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCounterparty
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
				return ErrInvalidLengthCounterparty
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCounterparty
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MerklePathPrefix.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCounterparty(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCounterparty
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
func skipCounterparty(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCounterparty
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
					return 0, ErrIntOverflowCounterparty
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
					return 0, ErrIntOverflowCounterparty
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
				return 0, ErrInvalidLengthCounterparty
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCounterparty
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCounterparty
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCounterparty        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCounterparty          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCounterparty = fmt.Errorf("proto: unexpected end of group")
)