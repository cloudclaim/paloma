// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/evm_reference_block_attestation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type ReferenceBlockAttestation struct {
	FromBlockTime time.Time `protobuf:"bytes,1,opt,name=fromBlockTime,proto3,stdtime" json:"fromBlockTime"`
}

func (m *ReferenceBlockAttestation) Reset()         { *m = ReferenceBlockAttestation{} }
func (m *ReferenceBlockAttestation) String() string { return proto.CompactTextString(m) }
func (*ReferenceBlockAttestation) ProtoMessage()    {}
func (*ReferenceBlockAttestation) Descriptor() ([]byte, []int) {
	return fileDescriptor_181c5ad50f766137, []int{0}
}
func (m *ReferenceBlockAttestation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReferenceBlockAttestation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReferenceBlockAttestation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReferenceBlockAttestation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferenceBlockAttestation.Merge(m, src)
}
func (m *ReferenceBlockAttestation) XXX_Size() int {
	return m.Size()
}
func (m *ReferenceBlockAttestation) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferenceBlockAttestation.DiscardUnknown(m)
}

var xxx_messageInfo_ReferenceBlockAttestation proto.InternalMessageInfo

func (m *ReferenceBlockAttestation) GetFromBlockTime() time.Time {
	if m != nil {
		return m.FromBlockTime
	}
	return time.Time{}
}

type ReferenceBlockAttestationRes struct {
	BlockHeight uint64 `protobuf:"varint,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	BlockHash   string `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
}

func (m *ReferenceBlockAttestationRes) Reset()         { *m = ReferenceBlockAttestationRes{} }
func (m *ReferenceBlockAttestationRes) String() string { return proto.CompactTextString(m) }
func (*ReferenceBlockAttestationRes) ProtoMessage()    {}
func (*ReferenceBlockAttestationRes) Descriptor() ([]byte, []int) {
	return fileDescriptor_181c5ad50f766137, []int{1}
}
func (m *ReferenceBlockAttestationRes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReferenceBlockAttestationRes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ReferenceBlockAttestationRes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ReferenceBlockAttestationRes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReferenceBlockAttestationRes.Merge(m, src)
}
func (m *ReferenceBlockAttestationRes) XXX_Size() int {
	return m.Size()
}
func (m *ReferenceBlockAttestationRes) XXX_DiscardUnknown() {
	xxx_messageInfo_ReferenceBlockAttestationRes.DiscardUnknown(m)
}

var xxx_messageInfo_ReferenceBlockAttestationRes proto.InternalMessageInfo

func (m *ReferenceBlockAttestationRes) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ReferenceBlockAttestationRes) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func init() {
	proto.RegisterType((*ReferenceBlockAttestation)(nil), "palomachain.paloma.evm.ReferenceBlockAttestation")
	proto.RegisterType((*ReferenceBlockAttestationRes)(nil), "palomachain.paloma.evm.ReferenceBlockAttestationRes")
}

func init() {
	proto.RegisterFile("palomachain/paloma/evm/evm_reference_block_attestation.proto", fileDescriptor_181c5ad50f766137)
}

var fileDescriptor_181c5ad50f766137 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x1b, 0x11, 0x71, 0x19, 0x5e, 0x8a, 0xc8, 0x1c, 0x23, 0x1b, 0x3b, 0xed, 0x20, 0x09,
	0xcc, 0xab, 0x17, 0x77, 0x10, 0xf1, 0x58, 0x76, 0xf2, 0x60, 0x49, 0xcb, 0x6b, 0x1a, 0x6c, 0x9a,
	0xd2, 0x66, 0x45, 0xff, 0x8b, 0xfd, 0x59, 0x3b, 0xee, 0xe8, 0x49, 0xa5, 0xfd, 0x47, 0xa4, 0x09,
	0x65, 0x13, 0xf4, 0x10, 0xf8, 0xf2, 0xf1, 0xbd, 0xef, 0xc7, 0x7b, 0xf8, 0xae, 0xe0, 0x99, 0x56,
	0x3c, 0x4e, 0xb9, 0xcc, 0x99, 0xd3, 0x0c, 0x6a, 0xd5, 0xbd, 0xb0, 0x84, 0x04, 0x4a, 0xc8, 0x63,
	0x08, 0xa3, 0x4c, 0xc7, 0xaf, 0x21, 0x37, 0x06, 0x2a, 0xc3, 0x8d, 0xd4, 0x39, 0x2d, 0x4a, 0x6d,
	0xb4, 0x7f, 0x75, 0x34, 0x4d, 0x9d, 0xa6, 0x50, 0xab, 0xf1, 0xa5, 0xd0, 0x42, 0xdb, 0x08, 0xeb,
	0x94, 0x4b, 0x8f, 0xa7, 0x42, 0x6b, 0x91, 0x01, 0xb3, 0xbf, 0x68, 0x93, 0x30, 0x23, 0x55, 0x57,
	0xa8, 0x0a, 0x17, 0x98, 0x0b, 0x7c, 0x1d, 0xf4, 0xcc, 0x55, 0x87, 0xbc, 0x3f, 0x10, 0xfd, 0x27,
	0x7c, 0x91, 0x94, 0x5a, 0x59, 0x7f, 0x2d, 0x15, 0x8c, 0xd0, 0x0c, 0x2d, 0x86, 0xcb, 0x31, 0x75,
	0xad, 0xb4, 0x6f, 0xa5, 0xeb, 0xbe, 0x75, 0x75, 0xbe, 0xfb, 0x9c, 0x7a, 0xdb, 0xaf, 0x29, 0x0a,
	0x7e, 0x8f, 0xce, 0x5f, 0xf0, 0xe4, 0x5f, 0x50, 0x00, 0x95, 0x3f, 0xc3, 0x43, 0xbb, 0xf2, 0x23,
	0x48, 0x91, 0x1a, 0x4b, 0x3a, 0x0d, 0x8e, 0x2d, 0x7f, 0x82, 0x07, 0xee, 0xcb, 0xab, 0x74, 0x74,
	0x32, 0x43, 0x8b, 0x41, 0x70, 0x30, 0x56, 0x0f, 0xbb, 0x86, 0xa0, 0x7d, 0x43, 0xd0, 0x77, 0x43,
	0xd0, 0xb6, 0x25, 0xde, 0xbe, 0x25, 0xde, 0x47, 0x4b, 0xbc, 0xe7, 0x1b, 0x21, 0x4d, 0xba, 0x89,
	0x68, 0xac, 0x15, 0xfb, 0xe3, 0xf4, 0xf5, 0x92, 0xbd, 0xd9, 0xfb, 0x9b, 0xf7, 0x02, 0xaa, 0xe8,
	0xcc, 0x2e, 0x75, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x16, 0xfb, 0x72, 0xf1, 0xa6, 0x01, 0x00,
	0x00,
}

func (m *ReferenceBlockAttestation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReferenceBlockAttestation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReferenceBlockAttestation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.FromBlockTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FromBlockTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEvmReferenceBlockAttestation(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *ReferenceBlockAttestationRes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReferenceBlockAttestationRes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReferenceBlockAttestationRes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintEvmReferenceBlockAttestation(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintEvmReferenceBlockAttestation(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvmReferenceBlockAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvmReferenceBlockAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReferenceBlockAttestation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.FromBlockTime)
	n += 1 + l + sovEvmReferenceBlockAttestation(uint64(l))
	return n
}

func (m *ReferenceBlockAttestationRes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovEvmReferenceBlockAttestation(uint64(m.BlockHeight))
	}
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovEvmReferenceBlockAttestation(uint64(l))
	}
	return n
}

func sovEvmReferenceBlockAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvmReferenceBlockAttestation(x uint64) (n int) {
	return sovEvmReferenceBlockAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReferenceBlockAttestation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvmReferenceBlockAttestation
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
			return fmt.Errorf("proto: ReferenceBlockAttestation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReferenceBlockAttestation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromBlockTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmReferenceBlockAttestation
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
				return ErrInvalidLengthEvmReferenceBlockAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvmReferenceBlockAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.FromBlockTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvmReferenceBlockAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvmReferenceBlockAttestation
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
func (m *ReferenceBlockAttestationRes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvmReferenceBlockAttestation
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
			return fmt.Errorf("proto: ReferenceBlockAttestationRes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReferenceBlockAttestationRes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmReferenceBlockAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvmReferenceBlockAttestation
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
				return ErrInvalidLengthEvmReferenceBlockAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvmReferenceBlockAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvmReferenceBlockAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvmReferenceBlockAttestation
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
func skipEvmReferenceBlockAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvmReferenceBlockAttestation
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
					return 0, ErrIntOverflowEvmReferenceBlockAttestation
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
					return 0, ErrIntOverflowEvmReferenceBlockAttestation
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
				return 0, ErrInvalidLengthEvmReferenceBlockAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvmReferenceBlockAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvmReferenceBlockAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvmReferenceBlockAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvmReferenceBlockAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvmReferenceBlockAttestation = fmt.Errorf("proto: unexpected end of group")
)
