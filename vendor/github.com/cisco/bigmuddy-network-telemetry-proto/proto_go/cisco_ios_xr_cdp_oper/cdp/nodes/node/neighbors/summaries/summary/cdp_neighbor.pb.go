// Code generated by protoc-gen-go.
// source: cdp_neighbor.proto
// DO NOT EDIT!

/*
Package cisco_ios_xr_cdp_oper_cdp_nodes_node_neighbors_summaries_summary is a generated protocol buffer package.

It is generated from these files:
	cdp_neighbor.proto

It has these top-level messages:
	CdpNeighbor_KEYS
	CdpNeighbor
	CdpNeighborItem
	In6AddrTd
	CdpL3Addr
	CdpAddrEntry
	CdpAddrEntryItem
	CdpProtHelloEntry
	CdpProtHelloEntryItem
	CdpNeighborDetail
*/
package cisco_ios_xr_cdp_oper_cdp_nodes_node_neighbors_summaries_summary

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// CDP neighbor info
type CdpNeighbor_KEYS struct {
	NodeName      string `protobuf:"bytes,1,opt,name=node_name,json=nodeName" json:"node_name,omitempty"`
	InterfaceName string `protobuf:"bytes,2,opt,name=interface_name,json=interfaceName" json:"interface_name,omitempty"`
	DeviceId      string `protobuf:"bytes,3,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
}

func (m *CdpNeighbor_KEYS) Reset()                    { *m = CdpNeighbor_KEYS{} }
func (m *CdpNeighbor_KEYS) String() string            { return proto.CompactTextString(m) }
func (*CdpNeighbor_KEYS) ProtoMessage()               {}
func (*CdpNeighbor_KEYS) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CdpNeighbor_KEYS) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *CdpNeighbor_KEYS) GetInterfaceName() string {
	if m != nil {
		return m.InterfaceName
	}
	return ""
}

func (m *CdpNeighbor_KEYS) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

type CdpNeighbor struct {
	// Next neighbor in the list
	CdpNeighbor []*CdpNeighborItem `protobuf:"bytes,50,rep,name=cdp_neighbor,json=cdpNeighbor" json:"cdp_neighbor,omitempty"`
}

func (m *CdpNeighbor) Reset()                    { *m = CdpNeighbor{} }
func (m *CdpNeighbor) String() string            { return proto.CompactTextString(m) }
func (*CdpNeighbor) ProtoMessage()               {}
func (*CdpNeighbor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CdpNeighbor) GetCdpNeighbor() []*CdpNeighborItem {
	if m != nil {
		return m.CdpNeighbor
	}
	return nil
}

type CdpNeighborItem struct {
	// Interface the neighbor entry was received on
	ReceivingInterfaceName string `protobuf:"bytes,1,opt,name=receiving_interface_name,json=receivingInterfaceName" json:"receiving_interface_name,omitempty"`
	// Device identifier
	DeviceId string `protobuf:"bytes,2,opt,name=device_id,json=deviceId" json:"device_id,omitempty"`
	// Outgoing port identifier
	PortId string `protobuf:"bytes,3,opt,name=port_id,json=portId" json:"port_id,omitempty"`
	// Version number
	HeaderVersion uint32 `protobuf:"varint,4,opt,name=header_version,json=headerVersion" json:"header_version,omitempty"`
	// Remaining hold time
	HoldTime uint32 `protobuf:"varint,5,opt,name=hold_time,json=holdTime" json:"hold_time,omitempty"`
	// Capabilities
	Capabilities string `protobuf:"bytes,6,opt,name=capabilities" json:"capabilities,omitempty"`
	// Platform type
	Platform string `protobuf:"bytes,7,opt,name=platform" json:"platform,omitempty"`
	// Detailed neighbor info
	Detail *CdpNeighborDetail `protobuf:"bytes,8,opt,name=detail" json:"detail,omitempty"`
}

func (m *CdpNeighborItem) Reset()                    { *m = CdpNeighborItem{} }
func (m *CdpNeighborItem) String() string            { return proto.CompactTextString(m) }
func (*CdpNeighborItem) ProtoMessage()               {}
func (*CdpNeighborItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CdpNeighborItem) GetReceivingInterfaceName() string {
	if m != nil {
		return m.ReceivingInterfaceName
	}
	return ""
}

func (m *CdpNeighborItem) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *CdpNeighborItem) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *CdpNeighborItem) GetHeaderVersion() uint32 {
	if m != nil {
		return m.HeaderVersion
	}
	return 0
}

func (m *CdpNeighborItem) GetHoldTime() uint32 {
	if m != nil {
		return m.HoldTime
	}
	return 0
}

func (m *CdpNeighborItem) GetCapabilities() string {
	if m != nil {
		return m.Capabilities
	}
	return ""
}

func (m *CdpNeighborItem) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *CdpNeighborItem) GetDetail() *CdpNeighborDetail {
	if m != nil {
		return m.Detail
	}
	return nil
}

type In6AddrTd struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *In6AddrTd) Reset()                    { *m = In6AddrTd{} }
func (m *In6AddrTd) String() string            { return proto.CompactTextString(m) }
func (*In6AddrTd) ProtoMessage()               {}
func (*In6AddrTd) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *In6AddrTd) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type CdpL3Addr struct {
	AddressType string `protobuf:"bytes,1,opt,name=address_type,json=addressType" json:"address_type,omitempty"`
	// IPv4 address
	Ipv4Address string `protobuf:"bytes,2,opt,name=ipv4_address,json=ipv4Address" json:"ipv4_address,omitempty"`
	// IPv6 address
	Ipv6Address *In6AddrTd `protobuf:"bytes,3,opt,name=ipv6_address,json=ipv6Address" json:"ipv6_address,omitempty"`
}

func (m *CdpL3Addr) Reset()                    { *m = CdpL3Addr{} }
func (m *CdpL3Addr) String() string            { return proto.CompactTextString(m) }
func (*CdpL3Addr) ProtoMessage()               {}
func (*CdpL3Addr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CdpL3Addr) GetAddressType() string {
	if m != nil {
		return m.AddressType
	}
	return ""
}

func (m *CdpL3Addr) GetIpv4Address() string {
	if m != nil {
		return m.Ipv4Address
	}
	return ""
}

func (m *CdpL3Addr) GetIpv6Address() *In6AddrTd {
	if m != nil {
		return m.Ipv6Address
	}
	return nil
}

type CdpAddrEntry struct {
	// Next address entry in list
	CdpAddrEntry []*CdpAddrEntryItem `protobuf:"bytes,1,rep,name=cdp_addr_entry,json=cdpAddrEntry" json:"cdp_addr_entry,omitempty"`
}

func (m *CdpAddrEntry) Reset()                    { *m = CdpAddrEntry{} }
func (m *CdpAddrEntry) String() string            { return proto.CompactTextString(m) }
func (*CdpAddrEntry) ProtoMessage()               {}
func (*CdpAddrEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CdpAddrEntry) GetCdpAddrEntry() []*CdpAddrEntryItem {
	if m != nil {
		return m.CdpAddrEntry
	}
	return nil
}

type CdpAddrEntryItem struct {
	// Network layer address
	Address *CdpL3Addr `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
}

func (m *CdpAddrEntryItem) Reset()                    { *m = CdpAddrEntryItem{} }
func (m *CdpAddrEntryItem) String() string            { return proto.CompactTextString(m) }
func (*CdpAddrEntryItem) ProtoMessage()               {}
func (*CdpAddrEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CdpAddrEntryItem) GetAddress() *CdpL3Addr {
	if m != nil {
		return m.Address
	}
	return nil
}

type CdpProtHelloEntry struct {
	// Next protocol hello entry in list
	CdpProtHelloEntry []*CdpProtHelloEntryItem `protobuf:"bytes,1,rep,name=cdp_prot_hello_entry,json=cdpProtHelloEntry" json:"cdp_prot_hello_entry,omitempty"`
}

func (m *CdpProtHelloEntry) Reset()                    { *m = CdpProtHelloEntry{} }
func (m *CdpProtHelloEntry) String() string            { return proto.CompactTextString(m) }
func (*CdpProtHelloEntry) ProtoMessage()               {}
func (*CdpProtHelloEntry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CdpProtHelloEntry) GetCdpProtHelloEntry() []*CdpProtHelloEntryItem {
	if m != nil {
		return m.CdpProtHelloEntry
	}
	return nil
}

type CdpProtHelloEntryItem struct {
	// Protocol Hello msg
	HelloMessage []byte `protobuf:"bytes,1,opt,name=hello_message,json=helloMessage,proto3" json:"hello_message,omitempty"`
}

func (m *CdpProtHelloEntryItem) Reset()                    { *m = CdpProtHelloEntryItem{} }
func (m *CdpProtHelloEntryItem) String() string            { return proto.CompactTextString(m) }
func (*CdpProtHelloEntryItem) ProtoMessage()               {}
func (*CdpProtHelloEntryItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CdpProtHelloEntryItem) GetHelloMessage() []byte {
	if m != nil {
		return m.HelloMessage
	}
	return nil
}

type CdpNeighborDetail struct {
	// List of network addresses
	NetworkAddresses *CdpAddrEntry `protobuf:"bytes,1,opt,name=network_addresses,json=networkAddresses" json:"network_addresses,omitempty"`
	// Version TLV
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	// List of protocol hello entries
	ProtocolHelloList *CdpProtHelloEntry `protobuf:"bytes,3,opt,name=protocol_hello_list,json=protocolHelloList" json:"protocol_hello_list,omitempty"`
	// VTP domain
	VtpDomain string `protobuf:"bytes,4,opt,name=vtp_domain,json=vtpDomain" json:"vtp_domain,omitempty"`
	// Native VLAN
	NativeVlan uint32 `protobuf:"varint,5,opt,name=native_vlan,json=nativeVlan" json:"native_vlan,omitempty"`
	// Duplex setting
	Duplex string `protobuf:"bytes,6,opt,name=duplex" json:"duplex,omitempty"`
	// SysName
	SystemName string `protobuf:"bytes,7,opt,name=system_name,json=systemName" json:"system_name,omitempty"`
}

func (m *CdpNeighborDetail) Reset()                    { *m = CdpNeighborDetail{} }
func (m *CdpNeighborDetail) String() string            { return proto.CompactTextString(m) }
func (*CdpNeighborDetail) ProtoMessage()               {}
func (*CdpNeighborDetail) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CdpNeighborDetail) GetNetworkAddresses() *CdpAddrEntry {
	if m != nil {
		return m.NetworkAddresses
	}
	return nil
}

func (m *CdpNeighborDetail) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *CdpNeighborDetail) GetProtocolHelloList() *CdpProtHelloEntry {
	if m != nil {
		return m.ProtocolHelloList
	}
	return nil
}

func (m *CdpNeighborDetail) GetVtpDomain() string {
	if m != nil {
		return m.VtpDomain
	}
	return ""
}

func (m *CdpNeighborDetail) GetNativeVlan() uint32 {
	if m != nil {
		return m.NativeVlan
	}
	return 0
}

func (m *CdpNeighborDetail) GetDuplex() string {
	if m != nil {
		return m.Duplex
	}
	return ""
}

func (m *CdpNeighborDetail) GetSystemName() string {
	if m != nil {
		return m.SystemName
	}
	return ""
}

func init() {
	proto.RegisterType((*CdpNeighbor_KEYS)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_neighbor_KEYS")
	proto.RegisterType((*CdpNeighbor)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_neighbor")
	proto.RegisterType((*CdpNeighborItem)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_neighbor_item")
	proto.RegisterType((*In6AddrTd)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.in6_addr_td")
	proto.RegisterType((*CdpL3Addr)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_l3_addr")
	proto.RegisterType((*CdpAddrEntry)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_addr_entry")
	proto.RegisterType((*CdpAddrEntryItem)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_addr_entry_item")
	proto.RegisterType((*CdpProtHelloEntry)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_prot_hello_entry")
	proto.RegisterType((*CdpProtHelloEntryItem)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_prot_hello_entry_item")
	proto.RegisterType((*CdpNeighborDetail)(nil), "cisco_ios_xr_cdp_oper.cdp.nodes.node.neighbors.summaries.summary.cdp_neighbor_detail")
}

func init() { proto.RegisterFile("cdp_neighbor.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 698 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcf, 0x4f, 0x1b, 0x39,
	0x14, 0xd6, 0x90, 0x25, 0x3f, 0x5e, 0x12, 0xb4, 0x18, 0xc4, 0xce, 0x2e, 0x5a, 0x6d, 0x76, 0xd0,
	0x4a, 0x39, 0xe5, 0x00, 0x2b, 0xb4, 0x47, 0x90, 0x16, 0xa9, 0xa8, 0x05, 0xa1, 0x81, 0x22, 0x55,
	0x3d, 0x58, 0x66, 0x6c, 0x12, 0xab, 0x33, 0x63, 0xcb, 0x36, 0x53, 0xa2, 0xaa, 0x3d, 0x72, 0x6a,
	0xff, 0x8c, 0xde, 0xfb, 0x4f, 0xf4, 0x5f, 0xea, 0xb9, 0xf2, 0x8f, 0x81, 0x24, 0x2a, 0x27, 0xd2,
	0x4b, 0x92, 0xf7, 0xf9, 0xf3, 0x7b, 0xfe, 0x3e, 0x3f, 0xbf, 0x00, 0xca, 0xa8, 0xc4, 0x25, 0xe3,
	0xe3, 0xc9, 0x95, 0x50, 0x23, 0xa9, 0x84, 0x11, 0xe8, 0x20, 0xe3, 0x3a, 0x13, 0x98, 0x0b, 0x8d,
	0x6f, 0x15, 0xb6, 0x04, 0x21, 0x99, 0x1a, 0x65, 0x54, 0x8e, 0x4a, 0x41, 0x99, 0x76, 0x9f, 0xa3,
	0x7a, 0x93, 0x1e, 0xe9, 0x9b, 0xa2, 0x20, 0x8a, 0xb3, 0xfa, 0xd7, 0x34, 0x31, 0xb0, 0x3e, 0x9b,
	0x17, 0x3f, 0x3f, 0x7a, 0x75, 0x8e, 0xb6, 0xa1, 0x63, 0x37, 0xe2, 0x92, 0x14, 0x2c, 0x8e, 0x06,
	0xd1, 0xb0, 0x93, 0xb6, 0x2d, 0x70, 0x4a, 0x0a, 0x86, 0xfe, 0x81, 0x35, 0x5e, 0x1a, 0xa6, 0xae,
	0x49, 0x16, 0x18, 0x2b, 0x8e, 0xd1, 0xbf, 0x47, 0x1d, 0x6d, 0x1b, 0x3a, 0x94, 0x55, 0x3c, 0x63,
	0x98, 0xd3, 0xb8, 0xe1, 0x73, 0x78, 0xe0, 0x98, 0x26, 0x77, 0x11, 0xf4, 0x66, 0xcb, 0xa2, 0x6a,
	0x3e, 0x8e, 0x77, 0x07, 0x8d, 0x61, 0x77, 0xf7, 0x7c, 0xf4, 0x54, 0x7d, 0xa3, 0x39, 0x71, 0xdc,
	0xb0, 0x22, 0xed, 0x66, 0x54, 0x9e, 0x06, 0x24, 0xf9, 0xb6, 0xb2, 0xa0, 0xdf, 0x52, 0xd0, 0x7f,
	0x10, 0x2b, 0x96, 0x31, 0x5e, 0xf1, 0x72, 0x8c, 0x17, 0xc4, 0x7a, 0x3b, 0xb6, 0xee, 0xd7, 0x8f,
	0x1f, 0x57, 0xbd, 0x32, 0xaf, 0x1a, 0xfd, 0x06, 0x2d, 0x29, 0x94, 0x79, 0x30, 0xa4, 0x69, 0xc3,
	0x63, 0x6a, 0x2d, 0x9d, 0x30, 0x42, 0x99, 0xc2, 0x15, 0x53, 0x9a, 0x8b, 0x32, 0xfe, 0x65, 0x10,
	0x0d, 0xfb, 0x69, 0xdf, 0xa3, 0x97, 0x1e, 0xb4, 0xc9, 0x27, 0x22, 0xa7, 0xd8, 0xf0, 0x82, 0xc5,
	0xab, 0x8e, 0xd1, 0xb6, 0xc0, 0x05, 0x2f, 0x18, 0x4a, 0xa0, 0x97, 0x11, 0x49, 0xae, 0x78, 0xce,
	0x0d, 0x67, 0x3a, 0x6e, 0xba, 0x0a, 0x73, 0x18, 0xfa, 0x03, 0xda, 0x32, 0x27, 0xe6, 0x5a, 0xa8,
	0x22, 0x6e, 0xf9, 0xc3, 0xd5, 0x31, 0x2a, 0xa0, 0x49, 0x99, 0x21, 0x3c, 0x8f, 0xdb, 0x83, 0x68,
	0xd8, 0xdd, 0x7d, 0xb9, 0x64, 0xef, 0x7d, 0xf2, 0x34, 0x14, 0x49, 0x76, 0xa0, 0xcb, 0xcb, 0x7d,
	0x4c, 0x28, 0x55, 0xd8, 0x50, 0xb4, 0x09, 0xab, 0x15, 0xc9, 0x6f, 0x6a, 0x7b, 0x7d, 0x90, 0x7c,
	0x8d, 0xc0, 0xde, 0x16, 0xce, 0xf7, 0x1c, 0x11, 0xfd, 0x0d, 0x3d, 0xfb, 0xcd, 0xb4, 0xc6, 0x66,
	0x2a, 0x6b, 0x72, 0x37, 0x60, 0x17, 0x53, 0xc9, 0x2c, 0x85, 0xcb, 0xea, 0x5f, 0x1c, 0xb0, 0x70,
	0x07, 0x5d, 0x8b, 0x1d, 0x7a, 0x08, 0x49, 0x47, 0xd9, 0xbf, 0xa7, 0x34, 0x9c, 0xde, 0x93, 0xa7,
	0xeb, 0x9d, 0x11, 0xe4, 0x2a, 0xee, 0x87, 0x8a, 0xc9, 0xa7, 0x08, 0xd6, 0x6c, 0x42, 0xb7, 0xc8,
	0x4a, 0xa3, 0xa6, 0xe8, 0xdd, 0x22, 0x12, 0x47, 0xae, 0xe5, 0x97, 0x64, 0xfb, 0x43, 0x5e, 0xdf,
	0xf4, 0xf6, 0x75, 0xd9, 0xd3, 0x1c, 0x59, 0x28, 0xf9, 0x00, 0x1b, 0x3f, 0x20, 0xa1, 0x31, 0xb4,
	0x6a, 0x4f, 0xa2, 0x65, 0x79, 0x32, 0x73, 0x7d, 0x69, 0x9d, 0x3d, 0xf9, 0x1c, 0xc1, 0xa6, 0x5d,
	0xb0, 0x43, 0x0c, 0x4f, 0x58, 0x9e, 0x8b, 0xe0, 0xca, 0xc7, 0x47, 0x16, 0x82, 0x39, 0xaf, 0x97,
	0x73, 0x9e, 0xc5, 0xec, 0xde, 0x22, 0x3b, 0x07, 0xce, 0x94, 0x30, 0xcf, 0xec, 0x82, 0xf7, 0xe9,
	0x00, 0x7e, 0x7f, 0x94, 0x8f, 0x76, 0xa0, 0xef, 0xb1, 0x82, 0x69, 0x4d, 0xc6, 0xbe, 0x1b, 0x7b,
	0x69, 0xcf, 0x81, 0x27, 0x1e, 0x4b, 0xbe, 0x34, 0xbc, 0xd5, 0x0b, 0xcf, 0x00, 0xbd, 0x87, 0xf5,
	0x92, 0x99, 0xb7, 0x42, 0xbd, 0xa9, 0xdb, 0x90, 0xd5, 0xa6, 0x9f, 0x2d, 0xbb, 0x03, 0xd2, 0x5f,
	0x43, 0xa9, 0xc3, 0xba, 0x12, 0x8a, 0xa1, 0x55, 0x4f, 0x1a, 0xff, 0x40, 0xea, 0x10, 0xdd, 0x45,
	0xb0, 0xe1, 0xfe, 0x5b, 0x32, 0x91, 0x07, 0xcd, 0x39, 0xd7, 0x26, 0x3c, 0x92, 0xcb, 0x9f, 0x73,
	0x01, 0xe9, 0x7a, 0x5d, 0xd2, 0x99, 0xff, 0x82, 0x6b, 0x83, 0xfe, 0x04, 0xa8, 0x8c, 0xc4, 0x54,
	0x14, 0x84, 0xfb, 0x79, 0xd8, 0x49, 0x3b, 0x95, 0x91, 0xff, 0x3b, 0x00, 0xfd, 0x05, 0xdd, 0x92,
	0x18, 0x5e, 0x31, 0x5c, 0xe5, 0xa4, 0x0c, 0xd3, 0x10, 0x3c, 0x74, 0x99, 0x93, 0x12, 0x6d, 0x41,
	0x93, 0xde, 0xc8, 0x9c, 0xdd, 0x86, 0x49, 0x18, 0x22, 0xbb, 0x51, 0x4f, 0xb5, 0x61, 0x85, 0x1f,
	0xe7, 0x7e, 0x0c, 0x82, 0x87, 0xec, 0x08, 0xbf, 0x6a, 0xba, 0xb3, 0xec, 0x7d, 0x0f, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x61, 0x6a, 0xad, 0x70, 0x07, 0x00, 0x00,
}