// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: empowerchain/plasticcredit/genesis.proto

package plasticcredit

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// GenesisState defines the plasticcredit module's genesis state.
type GenesisState struct {
	Params            Params             `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	IdCounters        IDCounters         `protobuf:"bytes,2,opt,name=id_counters,json=idCounters,proto3" json:"id_counters"`
	Issuers           []Issuer           `protobuf:"bytes,3,rep,name=issuers,proto3" json:"issuers"`
	Applicants        []Applicant        `protobuf:"bytes,4,rep,name=applicants,proto3" json:"applicants"`
	CreditTypes       []CreditType       `protobuf:"bytes,5,rep,name=credit_types,json=creditTypes,proto3" json:"credit_types"`
	Projects          []Project          `protobuf:"bytes,6,rep,name=projects,proto3" json:"projects"`
	CreditCollections []CreditCollection `protobuf:"bytes,7,rep,name=credit_collections,json=creditCollections,proto3" json:"credit_collections"`
	CreditBalances    []CreditBalance    `protobuf:"bytes,8,rep,name=credit_balances,json=creditBalances,proto3" json:"credit_balances"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3d9a84508647bd6f, []int{0}
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

func (m *GenesisState) GetIdCounters() IDCounters {
	if m != nil {
		return m.IdCounters
	}
	return IDCounters{}
}

func (m *GenesisState) GetIssuers() []Issuer {
	if m != nil {
		return m.Issuers
	}
	return nil
}

func (m *GenesisState) GetApplicants() []Applicant {
	if m != nil {
		return m.Applicants
	}
	return nil
}

func (m *GenesisState) GetCreditTypes() []CreditType {
	if m != nil {
		return m.CreditTypes
	}
	return nil
}

func (m *GenesisState) GetProjects() []Project {
	if m != nil {
		return m.Projects
	}
	return nil
}

func (m *GenesisState) GetCreditCollections() []CreditCollection {
	if m != nil {
		return m.CreditCollections
	}
	return nil
}

func (m *GenesisState) GetCreditBalances() []CreditBalance {
	if m != nil {
		return m.CreditBalances
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "empowerchain.plasticcredit.GenesisState")
}

func init() {
	proto.RegisterFile("empowerchain/plasticcredit/genesis.proto", fileDescriptor_3d9a84508647bd6f)
}

var fileDescriptor_3d9a84508647bd6f = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0x93, 0xdb, 0xbf, 0x4c, 0xcb, 0xbd, 0xdc, 0xe1, 0x2e, 0x42, 0x17, 0xb9, 0xa5, 0x62,
	0xa9, 0x20, 0x09, 0x28, 0xb8, 0xd6, 0xd4, 0x22, 0x45, 0xc4, 0xa2, 0x2e, 0xc4, 0x4d, 0x99, 0x4e,
	0x87, 0x74, 0x24, 0xcd, 0x0c, 0x99, 0x29, 0xda, 0x07, 0x70, 0xef, 0x63, 0x75, 0xd9, 0xa5, 0x2b,
	0x91, 0xf6, 0x45, 0xa4, 0x33, 0x93, 0xda, 0x0a, 0xa6, 0xab, 0x0c, 0x87, 0x73, 0x7e, 0xdf, 0xc9,
	0xc7, 0x07, 0x5a, 0x64, 0xcc, 0xd9, 0x13, 0x49, 0xf0, 0x08, 0xd1, 0xd8, 0xe7, 0x11, 0x12, 0x92,
	0x62, 0x9c, 0x90, 0x21, 0x95, 0x7e, 0x48, 0x62, 0x22, 0xa8, 0xf0, 0x78, 0xc2, 0x24, 0x83, 0xb5,
	0x4d, 0xa7, 0xb7, 0xe5, 0xac, 0x35, 0x33, 0x28, 0x72, 0xca, 0x89, 0x61, 0xd4, 0xfe, 0x85, 0x2c,
	0x64, 0xea, 0xe9, 0xaf, 0x5e, 0x5a, 0x6d, 0xbc, 0x14, 0x40, 0xf5, 0x42, 0xcf, 0xba, 0x95, 0x48,
	0x12, 0x78, 0x0a, 0x8a, 0x1c, 0x25, 0x68, 0x2c, 0x1c, 0xbb, 0x6e, 0xb7, 0x2a, 0x47, 0x0d, 0xef,
	0xe7, 0xd9, 0x5e, 0x4f, 0x39, 0x83, 0xfc, 0xec, 0xfd, 0xbf, 0x75, 0x63, 0x72, 0xf0, 0x0a, 0x54,
	0xe8, 0xb0, 0x8f, 0xd9, 0x24, 0x96, 0x24, 0x11, 0xce, 0x2f, 0x85, 0x69, 0x66, 0x61, 0xba, 0xe7,
	0x6d, 0xe3, 0x36, 0x28, 0x40, 0x87, 0xa9, 0x02, 0x03, 0x50, 0xa2, 0x42, 0x4c, 0x56, 0xa8, 0x5c,
	0x3d, 0xb7, 0xab, 0x51, 0x57, 0x59, 0x0d, 0x26, 0x0d, 0xc2, 0x4b, 0x00, 0x10, 0xe7, 0x11, 0xc5,
	0x28, 0x96, 0xc2, 0xc9, 0x2b, 0xcc, 0x7e, 0x16, 0xe6, 0x2c, 0x75, 0xa7, 0x85, 0xbe, 0xe2, 0xf0,
	0x1a, 0x54, 0xb5, 0xab, 0xaf, 0xd6, 0xeb, 0x14, 0x14, 0x2e, 0xf3, 0x07, 0xdb, 0xea, 0x73, 0x37,
	0xe5, 0xc4, 0xf0, 0x2a, 0x78, 0xad, 0x08, 0xd8, 0x01, 0x65, 0x9e, 0xb0, 0x47, 0x82, 0xa5, 0x70,
	0x8a, 0x0a, 0xb6, 0x97, 0xb9, 0x74, 0xed, 0x35, 0xa4, 0x75, 0x14, 0x22, 0x00, 0x4d, 0x2f, 0xcc,
	0xa2, 0x88, 0x60, 0x49, 0x59, 0x2c, 0x9c, 0x92, 0x02, 0x1e, 0xee, 0x6e, 0xd7, 0x5e, 0x87, 0x0c,
	0xf9, 0x2f, 0xfe, 0xa6, 0x0b, 0x78, 0x0f, 0xfe, 0x98, 0x11, 0x03, 0x14, 0xa1, 0x18, 0x13, 0xe1,
	0x94, 0x15, 0xff, 0x60, 0x37, 0x3f, 0xd0, 0x09, 0x03, 0xff, 0x8d, 0x37, 0x45, 0x11, 0xf4, 0x66,
	0x0b, 0xd7, 0x9e, 0x2f, 0x5c, 0xfb, 0x63, 0xe1, 0xda, 0xaf, 0x4b, 0xd7, 0x9a, 0x2f, 0x5d, 0xeb,
	0x6d, 0xe9, 0x5a, 0x0f, 0x27, 0x21, 0x95, 0xa3, 0xc9, 0xc0, 0xc3, 0x6c, 0xec, 0x77, 0xf4, 0x90,
	0x9e, 0xc6, 0xfb, 0x5b, 0x97, 0xff, 0xbc, 0x7d, 0xfb, 0x83, 0xa2, 0x3a, 0xf0, 0xe3, 0xcf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x5b, 0x09, 0x60, 0xe4, 0x66, 0x03, 0x00, 0x00,
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
	if len(m.CreditBalances) > 0 {
		for iNdEx := len(m.CreditBalances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditBalances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if len(m.CreditCollections) > 0 {
		for iNdEx := len(m.CreditCollections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditCollections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.Projects) > 0 {
		for iNdEx := len(m.Projects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Projects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.CreditTypes) > 0 {
		for iNdEx := len(m.CreditTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Applicants) > 0 {
		for iNdEx := len(m.Applicants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Applicants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Issuers) > 0 {
		for iNdEx := len(m.Issuers) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issuers[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.IdCounters.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
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
	l = m.IdCounters.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Issuers) > 0 {
		for _, e := range m.Issuers {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Applicants) > 0 {
		for _, e := range m.Applicants {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CreditTypes) > 0 {
		for _, e := range m.CreditTypes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Projects) > 0 {
		for _, e := range m.Projects {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CreditCollections) > 0 {
		for _, e := range m.CreditCollections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CreditBalances) > 0 {
		for _, e := range m.CreditBalances {
			l = e.Size()
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdCounters", wireType)
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
			if err := m.IdCounters.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuers", wireType)
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
			m.Issuers = append(m.Issuers, Issuer{})
			if err := m.Issuers[len(m.Issuers)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Applicants", wireType)
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
			m.Applicants = append(m.Applicants, Applicant{})
			if err := m.Applicants[len(m.Applicants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditTypes", wireType)
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
			m.CreditTypes = append(m.CreditTypes, CreditType{})
			if err := m.CreditTypes[len(m.CreditTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Projects", wireType)
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
			m.Projects = append(m.Projects, Project{})
			if err := m.Projects[len(m.Projects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditCollections", wireType)
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
			m.CreditCollections = append(m.CreditCollections, CreditCollection{})
			if err := m.CreditCollections[len(m.CreditCollections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditBalances", wireType)
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
			m.CreditBalances = append(m.CreditBalances, CreditBalance{})
			if err := m.CreditBalances[len(m.CreditBalances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
