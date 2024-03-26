// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evmos/recovery/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// GenesisState defines the recovery module's genesis state.
type GenesisState struct {
	// params defines all the paramaters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a3e70cb61e26f25, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// Params holds parameters for the recovery module
type Params struct {
	// enable_recovery IBC middleware
	EnableRecovery bool `protobuf:"varint,1,opt,name=enable_recovery,json=enableRecovery,proto3" json:"enable_recovery,omitempty"`
	// packet_timeout_duration is the duration added to timeout timestamp for balances recovered via IBC packets
	PacketTimeoutDuration time.Duration `protobuf:"bytes,2,opt,name=packet_timeout_duration,json=packetTimeoutDuration,proto3,stdduration" json:"packet_timeout_duration"`
	// authorized_channels is the list of authorized channel identifiers that can perform address
	// attestations via IBC.
	AuthorizedChannels []string `protobuf:"bytes,6,rep,name=authorized_channels,json=authorizedChannels,proto3" json:"authorized_channels,omitempty"`
	// evm_channels is the list of channel identifiers from EVM compatible chains
	EVMChannels []string `protobuf:"bytes,7,rep,name=evm_channels,json=evmChannels,proto3" json:"evm_channels,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a3e70cb61e26f25, []int{1}
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

func (m *Params) GetEnableRecovery() bool {
	if m != nil {
		return m.EnableRecovery
	}
	return false
}

func (m *Params) GetPacketTimeoutDuration() time.Duration {
	if m != nil {
		return m.PacketTimeoutDuration
	}
	return 0
}

func (m *Params) GetAuthorizedChannels() []string {
	if m != nil {
		return m.AuthorizedChannels
	}
	return nil
}

func (m *Params) GetEVMChannels() []string {
	if m != nil {
		return m.EVMChannels
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "evmos.recovery.v1.GenesisState")
	proto.RegisterType((*Params)(nil), "evmos.recovery.v1.Params")
}

func init() { proto.RegisterFile("evmos/recovery/v1/genesis.proto", fileDescriptor_8a3e70cb61e26f25) }

var fileDescriptor_8a3e70cb61e26f25 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x13, 0x95, 0x58, 0xb7, 0xc5, 0x62, 0x54, 0xac, 0x3d, 0x24, 0xa5, 0x17, 0x0b, 0xc2,
	0x2e, 0xad, 0x82, 0xf7, 0x68, 0xe9, 0x49, 0x90, 0x28, 0x1e, 0xf4, 0x10, 0x36, 0xe9, 0x98, 0x06,
	0x9b, 0x6c, 0x48, 0x36, 0xc1, 0xfa, 0x14, 0x1e, 0x7d, 0xa4, 0x1e, 0x7b, 0xf4, 0x54, 0x25, 0x7d,
	0x04, 0x5f, 0x40, 0x9a, 0xcd, 0xd2, 0x83, 0x97, 0x65, 0x67, 0xfe, 0x6f, 0xff, 0xf9, 0x97, 0x41,
	0x26, 0xe4, 0x21, 0x4b, 0x49, 0x02, 0x1e, 0xcb, 0x21, 0x99, 0x91, 0xbc, 0x4f, 0x7c, 0x88, 0x20,
	0x0d, 0x52, 0x1c, 0x27, 0x8c, 0x33, 0xfd, 0xa0, 0x04, 0xb0, 0x04, 0x70, 0xde, 0x6f, 0x1f, 0xf9,
	0xcc, 0x67, 0xa5, 0x4a, 0xd6, 0x37, 0x01, 0xb6, 0x0d, 0x9f, 0x31, 0x7f, 0x0a, 0xa4, 0xac, 0xdc,
	0xec, 0x85, 0x8c, 0xb3, 0x84, 0xf2, 0x80, 0x45, 0x42, 0xef, 0x8e, 0x50, 0x63, 0x24, 0x9c, 0xef,
	0x39, 0xe5, 0xa0, 0x5f, 0x21, 0x2d, 0xa6, 0x09, 0x0d, 0xd3, 0x96, 0xda, 0x51, 0x7b, 0xf5, 0xc1,
	0x29, 0xfe, 0x37, 0x09, 0xdf, 0x95, 0x80, 0xb5, 0x33, 0x5f, 0x9a, 0x8a, 0x5d, 0xe1, 0xdd, 0x5f,
	0x15, 0x69, 0x42, 0xd0, 0xcf, 0x50, 0x13, 0x22, 0xea, 0x4e, 0xc1, 0x91, 0xaf, 0x4a, 0xb3, 0x9a,
	0xbd, 0x2f, 0xda, 0x76, 0xd5, 0xd5, 0x9f, 0xd1, 0x49, 0x4c, 0xbd, 0x57, 0xe0, 0x0e, 0x0f, 0x42,
	0x60, 0x19, 0x77, 0x64, 0xba, 0xd6, 0x56, 0x35, 0x5d, 0xc4, 0xc7, 0x32, 0x3e, 0xbe, 0xa9, 0x00,
	0xab, 0xb6, 0x9e, 0xfe, 0xf9, 0x6d, 0xaa, 0xf6, 0xb1, 0xf0, 0x78, 0x10, 0x16, 0x12, 0xd0, 0x09,
	0x3a, 0xa4, 0x19, 0x9f, 0xb0, 0x24, 0x78, 0x87, 0xb1, 0xe3, 0x4d, 0x68, 0x14, 0xc1, 0x34, 0x6d,
	0x69, 0x9d, 0xed, 0xde, 0x9e, 0xad, 0x6f, 0xa4, 0xeb, 0x4a, 0xd1, 0x07, 0xa8, 0x01, 0x79, 0xb8,
	0x21, 0x77, 0xd7, 0xa4, 0xd5, 0x2c, 0x96, 0x66, 0x7d, 0xf8, 0x78, 0x2b, 0x31, 0xbb, 0x0e, 0x79,
	0x28, 0x0b, 0x6b, 0x38, 0x2f, 0x0c, 0x75, 0x51, 0x18, 0xea, 0x4f, 0x61, 0xa8, 0x1f, 0x2b, 0x43,
	0x59, 0xac, 0x0c, 0xe5, 0x6b, 0x65, 0x28, 0x4f, 0xe7, 0x7e, 0xc0, 0x27, 0x99, 0x8b, 0x3d, 0x16,
	0x12, 0xb1, 0x4d, 0x71, 0xe6, 0xfd, 0x4b, 0xf2, 0xb6, 0xd9, 0x2c, 0x9f, 0xc5, 0x90, 0xba, 0x5a,
	0xf9, 0xbf, 0x8b, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x32, 0xa6, 0xde, 0xf8, 0x01, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if len(m.EVMChannels) > 0 {
		for iNdEx := len(m.EVMChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.EVMChannels[iNdEx])
			copy(dAtA[i:], m.EVMChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.EVMChannels[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.AuthorizedChannels) > 0 {
		for iNdEx := len(m.AuthorizedChannels) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AuthorizedChannels[iNdEx])
			copy(dAtA[i:], m.AuthorizedChannels[iNdEx])
			i = encodeVarintGenesis(dAtA, i, uint64(len(m.AuthorizedChannels[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.PacketTimeoutDuration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.PacketTimeoutDuration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintGenesis(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x12
	if m.EnableRecovery {
		i--
		if m.EnableRecovery {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.EnableRecovery {
		n += 2
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.PacketTimeoutDuration)
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.AuthorizedChannels) > 0 {
		for _, s := range m.AuthorizedChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.EVMChannels) > 0 {
		for _, s := range m.EVMChannels {
			l = len(s)
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnableRecovery", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
			m.EnableRecovery = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PacketTimeoutDuration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.PacketTimeoutDuration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizedChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizedChannels = append(m.AuthorizedChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EVMChannels", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EVMChannels = append(m.EVMChannels, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
