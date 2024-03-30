// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as/external/api/device.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Device struct {
	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=UUID,proto3" json:"dev_eui,omitempty"`
	// Name of the device (if left blank, it will be set to the DevEUI).
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the device.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Device is onboarded to respective server.
	IsOnboarded bool `protobuf:"varint,4,opt,name=is_onboarded,json=isOnboarded,proto3" json:"is_onboarded,omitempty"`
	// Device is agateway.
	IsGateway bool `protobuf:"varint,5,opt,name=is_gateway,json=isGateway,proto3" json:"is_gateway,omitempty"`
	// Device is onboarding server internal or external.
	IsOnboardingServerExternal bool `protobuf:"varint,6,opt,name=is_onboarding_server_external,json=isOnboardingServerExternal,proto3" json:"is_onboarding_server_external,omitempty"`
	// Device is mode of communication like LoRaWAN, MQTT, GSM, TCP, CoAP.
	Communication string `protobuf:"bytes,7,opt,name=communication,proto3" json:"communication,omitempty"`
	// Device type like WaterMeter, TempSensor, Energy Meter.
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty"`
	// Variables (user defined).
	// These variables can be used together with integrations to store tokens /
	// secrets that must be configured per device. These variables are not
	// exposed in the event payloads.
	Variables map[string]string `protobuf:"bytes,9,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Tags (user defined).
	// These tags are exposed in the event payloads or to integration. Tags are
	// intended for aggregation and filtering.
	Tags map[string]string `protobuf:"bytes,10,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	//Device Profile
	Profile string `protobuf:"bytes,11,opt,name=profile,proto3" json:"profile,omitempty"`
	// Created at timestamp.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last update timestamp.
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{0}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Device) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Device) GetIsOnboarded() bool {
	if m != nil {
		return m.IsOnboarded
	}
	return false
}

func (m *Device) GetIsGateway() bool {
	if m != nil {
		return m.IsGateway
	}
	return false
}

func (m *Device) GetIsOnboardingServerExternal() bool {
	if m != nil {
		return m.IsOnboardingServerExternal
	}
	return false
}

func (m *Device) GetCommunication() string {
	if m != nil {
		return m.Communication
	}
	return ""
}

func (m *Device) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Device) GetVariables() map[string]string {
	if m != nil {
		return m.Variables
	}
	return nil
}

func (m *Device) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Device) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *Device) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Device) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type DeviceListItem struct {
	// Device EUI (HEX encoded).
	DevEui string `protobuf:"bytes,1,opt,name=dev_eui,json=UUID,proto3" json:"dev_eui,omitempty"`
	// Name of the device.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the device.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceListItem) Reset()         { *m = DeviceListItem{} }
func (m *DeviceListItem) String() string { return proto.CompactTextString(m) }
func (*DeviceListItem) ProtoMessage()    {}
func (*DeviceListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{1}
}

func (m *DeviceListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceListItem.Unmarshal(m, b)
}
func (m *DeviceListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceListItem.Marshal(b, m, deterministic)
}
func (m *DeviceListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceListItem.Merge(m, src)
}
func (m *DeviceListItem) XXX_Size() int {
	return xxx_messageInfo_DeviceListItem.Size(m)
}
func (m *DeviceListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceListItem proto.InternalMessageInfo

func (m *DeviceListItem) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

func (m *DeviceListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeviceListItem) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type CreateDeviceRequest struct {
	// Device object to create.
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateDeviceRequest) Reset()         { *m = CreateDeviceRequest{} }
func (m *CreateDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*CreateDeviceRequest) ProtoMessage()    {}
func (*CreateDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{2}
}

func (m *CreateDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateDeviceRequest.Unmarshal(m, b)
}
func (m *CreateDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateDeviceRequest.Marshal(b, m, deterministic)
}
func (m *CreateDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateDeviceRequest.Merge(m, src)
}
func (m *CreateDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_CreateDeviceRequest.Size(m)
}
func (m *CreateDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateDeviceRequest proto.InternalMessageInfo

func (m *CreateDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type GetDeviceRequest struct {
	// Device EUI (HEX encoded).
	DevEui               string   `protobuf:"bytes,1,opt,name=dev_eui,json=devEUI,proto3" json:"dev_eui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceRequest) Reset()         { *m = GetDeviceRequest{} }
func (m *GetDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeviceRequest) ProtoMessage()    {}
func (*GetDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{3}
}

func (m *GetDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceRequest.Unmarshal(m, b)
}
func (m *GetDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceRequest.Marshal(b, m, deterministic)
}
func (m *GetDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceRequest.Merge(m, src)
}
func (m *GetDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeviceRequest.Size(m)
}
func (m *GetDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceRequest proto.InternalMessageInfo

func (m *GetDeviceRequest) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

type GetDeviceResponse struct {
	// Device object.
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeviceResponse) Reset()         { *m = GetDeviceResponse{} }
func (m *GetDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*GetDeviceResponse) ProtoMessage()    {}
func (*GetDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{4}
}

func (m *GetDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeviceResponse.Unmarshal(m, b)
}
func (m *GetDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeviceResponse.Marshal(b, m, deterministic)
}
func (m *GetDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeviceResponse.Merge(m, src)
}
func (m *GetDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_GetDeviceResponse.Size(m)
}
func (m *GetDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeviceResponse proto.InternalMessageInfo

func (m *GetDeviceResponse) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

type ListDeviceRequest struct {
	// Max number of devices to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	// Offset in the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	// Search on name or DevEUI.
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
	// Tags to filter on.
	Tags                 map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListDeviceRequest) Reset()         { *m = ListDeviceRequest{} }
func (m *ListDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*ListDeviceRequest) ProtoMessage()    {}
func (*ListDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{5}
}

func (m *ListDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceRequest.Unmarshal(m, b)
}
func (m *ListDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceRequest.Marshal(b, m, deterministic)
}
func (m *ListDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceRequest.Merge(m, src)
}
func (m *ListDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_ListDeviceRequest.Size(m)
}
func (m *ListDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceRequest proto.InternalMessageInfo

func (m *ListDeviceRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListDeviceRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *ListDeviceRequest) GetSearch() string {
	if m != nil {
		return m.Search
	}
	return ""
}

func (m *ListDeviceRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ListDeviceResponse struct {
	// Total number of devices available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Devices within this result-set.
	Result               []*DeviceListItem `protobuf:"bytes,2,rep,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ListDeviceResponse) Reset()         { *m = ListDeviceResponse{} }
func (m *ListDeviceResponse) String() string { return proto.CompactTextString(m) }
func (*ListDeviceResponse) ProtoMessage()    {}
func (*ListDeviceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{6}
}

func (m *ListDeviceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDeviceResponse.Unmarshal(m, b)
}
func (m *ListDeviceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDeviceResponse.Marshal(b, m, deterministic)
}
func (m *ListDeviceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDeviceResponse.Merge(m, src)
}
func (m *ListDeviceResponse) XXX_Size() int {
	return xxx_messageInfo_ListDeviceResponse.Size(m)
}
func (m *ListDeviceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDeviceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDeviceResponse proto.InternalMessageInfo

func (m *ListDeviceResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListDeviceResponse) GetResult() []*DeviceListItem {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteDeviceRequest struct {
	// Device EUI (HEX encoded).
	DevEui               string   `protobuf:"bytes,1,opt,name=dev_eui,json=UUID,proto3" json:"dev_eui,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteDeviceRequest) Reset()         { *m = DeleteDeviceRequest{} }
func (m *DeleteDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteDeviceRequest) ProtoMessage()    {}
func (*DeleteDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{7}
}

func (m *DeleteDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteDeviceRequest.Unmarshal(m, b)
}
func (m *DeleteDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteDeviceRequest.Marshal(b, m, deterministic)
}
func (m *DeleteDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteDeviceRequest.Merge(m, src)
}
func (m *DeleteDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteDeviceRequest.Size(m)
}
func (m *DeleteDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteDeviceRequest proto.InternalMessageInfo

func (m *DeleteDeviceRequest) GetDevEui() string {
	if m != nil {
		return m.DevEui
	}
	return ""
}

type UpdateDeviceRequest struct {
	// Device object to update.
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateDeviceRequest) Reset()         { *m = UpdateDeviceRequest{} }
func (m *UpdateDeviceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDeviceRequest) ProtoMessage()    {}
func (*UpdateDeviceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_57ec3c2ed36f7cf9, []int{8}
}

func (m *UpdateDeviceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDeviceRequest.Unmarshal(m, b)
}
func (m *UpdateDeviceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDeviceRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDeviceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDeviceRequest.Merge(m, src)
}
func (m *UpdateDeviceRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDeviceRequest.Size(m)
}
func (m *UpdateDeviceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDeviceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDeviceRequest proto.InternalMessageInfo

func (m *UpdateDeviceRequest) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func init() {
	proto.RegisterType((*Device)(nil), "api.Device")
	proto.RegisterMapType((map[string]string)(nil), "api.Device.TagsEntry")
	proto.RegisterMapType((map[string]string)(nil), "api.Device.VariablesEntry")
	proto.RegisterType((*DeviceListItem)(nil), "api.DeviceListItem")
	proto.RegisterType((*CreateDeviceRequest)(nil), "api.CreateDeviceRequest")
	proto.RegisterType((*GetDeviceRequest)(nil), "api.GetDeviceRequest")
	proto.RegisterType((*GetDeviceResponse)(nil), "api.GetDeviceResponse")
	proto.RegisterType((*ListDeviceRequest)(nil), "api.ListDeviceRequest")
	proto.RegisterMapType((map[string]string)(nil), "api.ListDeviceRequest.TagsEntry")
	proto.RegisterType((*ListDeviceResponse)(nil), "api.ListDeviceResponse")
	proto.RegisterType((*DeleteDeviceRequest)(nil), "api.DeleteDeviceRequest")
	proto.RegisterType((*UpdateDeviceRequest)(nil), "api.UpdateDeviceRequest")
}

func init() {
	proto.RegisterFile("as/external/api/device.proto", fileDescriptor_57ec3c2ed36f7cf9)
}

var fileDescriptor_57ec3c2ed36f7cf9 = []byte{
	// 817 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x56, 0xea, 0xd4, 0xdd, 0x9c, 0xb4, 0xd5, 0x76, 0xda, 0x4d, 0x2d, 0xb3, 0x65, 0x83, 0xe1,
	0xa2, 0x74, 0x21, 0x96, 0xba, 0x2b, 0x28, 0x15, 0x37, 0x65, 0x5b, 0x55, 0x95, 0x40, 0x48, 0xde,
	0x96, 0x8b, 0x95, 0x50, 0x34, 0xb1, 0x4f, 0xbc, 0x23, 0x6c, 0x8f, 0xf1, 0x8c, 0x03, 0x11, 0x42,
	0x42, 0xbc, 0x02, 0x2f, 0x85, 0xc4, 0x25, 0xaf, 0xc0, 0x83, 0xa0, 0xf9, 0xf1, 0xae, 0x93, 0xa6,
	0x82, 0x22, 0xee, 0x66, 0xce, 0x7c, 0xe7, 0xef, 0x3b, 0x9f, 0x8f, 0xe1, 0x31, 0x15, 0x21, 0xfe,
	0x28, 0xb1, 0x2a, 0x68, 0x16, 0xd2, 0x92, 0x85, 0x09, 0xce, 0x58, 0x8c, 0xa3, 0xb2, 0xe2, 0x92,
	0x13, 0x87, 0x96, 0xcc, 0x7f, 0x9c, 0x72, 0x9e, 0x66, 0xa8, 0x5f, 0x69, 0x51, 0x70, 0x49, 0x25,
	0xe3, 0x85, 0x30, 0x10, 0xff, 0x89, 0x7d, 0xd5, 0xb7, 0x49, 0x3d, 0x0d, 0x25, 0xcb, 0x51, 0x48,
	0x9a, 0x97, 0x16, 0xf0, 0xce, 0x32, 0x00, 0xf3, 0x52, 0xce, 0xcd, 0x63, 0xf0, 0xcb, 0x3a, 0xb8,
	0xe7, 0x3a, 0x23, 0x79, 0x04, 0x1b, 0x09, 0xce, 0xc6, 0x58, 0x33, 0xaf, 0x33, 0xec, 0x1c, 0xf6,
	0xa2, 0xee, 0xcd, 0xcd, 0xd5, 0x39, 0x21, 0xd0, 0x2d, 0x68, 0x8e, 0xde, 0x9a, 0xb1, 0xa9, 0x33,
	0x19, 0x42, 0x3f, 0x41, 0x11, 0x57, 0xac, 0x54, 0x95, 0x78, 0x8e, 0x7e, 0x6a, 0x9b, 0xc8, 0x7b,
	0xb0, 0xc9, 0xc4, 0x98, 0x17, 0x13, 0x4e, 0xab, 0x04, 0x13, 0xaf, 0x3b, 0xec, 0x1c, 0x3e, 0x88,
	0xfa, 0x4c, 0x7c, 0xdd, 0x98, 0xc8, 0x01, 0x00, 0x13, 0xe3, 0x94, 0x4a, 0xfc, 0x81, 0xce, 0xbd,
	0x75, 0x0d, 0xe8, 0x31, 0x71, 0x69, 0x0c, 0xe4, 0x0c, 0x0e, 0xde, 0x46, 0x60, 0x45, 0x3a, 0x16,
	0x58, 0xcd, 0xb0, 0x1a, 0x37, 0x64, 0x79, 0xae, 0xf6, 0xf0, 0xdf, 0x84, 0x64, 0x45, 0xfa, 0x52,
	0x43, 0x2e, 0x2c, 0x82, 0x7c, 0x00, 0x5b, 0x31, 0xcf, 0xf3, 0xba, 0x60, 0xb1, 0xa6, 0xcc, 0xdb,
	0xd0, 0x85, 0x2e, 0x1a, 0x55, 0x83, 0x72, 0x5e, 0xa2, 0xf7, 0xc0, 0x34, 0xa8, 0xce, 0xe4, 0x04,
	0x7a, 0x33, 0x5a, 0x31, 0x3a, 0xc9, 0x50, 0x78, 0xbd, 0xa1, 0x73, 0xd8, 0x3f, 0xf6, 0x47, 0xb4,
	0x64, 0x23, 0xc3, 0xd5, 0xe8, 0x9b, 0xe6, 0xf1, 0xa2, 0x90, 0xd5, 0x3c, 0x7a, 0x0b, 0x26, 0x1f,
	0x42, 0x57, 0xd2, 0x54, 0x78, 0xa0, 0x9d, 0x1e, 0xb5, 0x9d, 0xae, 0x69, 0x6a, 0xf1, 0x1a, 0x42,
	0x3c, 0xd8, 0x28, 0x2b, 0x3e, 0x65, 0x19, 0x7a, 0x7d, 0x9d, 0xbb, 0xb9, 0x92, 0xcf, 0x00, 0xe2,
	0x0a, 0xa9, 0xc4, 0x64, 0x4c, 0xa5, 0xb7, 0x39, 0xec, 0xe8, 0xfc, 0x66, 0x8e, 0xa3, 0x66, 0x8e,
	0xa3, 0xeb, 0x66, 0xd0, 0x51, 0xcf, 0xa2, 0xcf, 0xa4, 0x72, 0xad, 0xcb, 0xa4, 0x71, 0xdd, 0xfa,
	0x67, 0x57, 0x8b, 0x3e, 0x93, 0xfe, 0xe7, 0xb0, 0xbd, 0xd8, 0x17, 0x79, 0x08, 0xce, 0x77, 0x38,
	0xb7, 0x72, 0x50, 0x47, 0xb2, 0x07, 0xeb, 0x33, 0x9a, 0xd5, 0x8d, 0x1c, 0xcc, 0xe5, 0x74, 0xed,
	0xa4, 0xe3, 0x7f, 0x0a, 0xbd, 0x37, 0x0d, 0xde, 0xc7, 0x31, 0xf8, 0x16, 0xb6, 0x0d, 0x41, 0x5f,
	0x32, 0x21, 0xaf, 0x24, 0xe6, 0xff, 0xab, 0x12, 0x83, 0x53, 0xd8, 0x7d, 0xa1, 0xd9, 0x31, 0x49,
	0x22, 0xfc, 0xbe, 0x46, 0x21, 0xc9, 0xfb, 0xe0, 0x9a, 0x2f, 0x4d, 0xa7, 0xe8, 0x1f, 0xf7, 0x5b,
	0x93, 0x8a, 0xec, 0x53, 0xf0, 0x14, 0x1e, 0x5e, 0xa2, 0x5c, 0x74, 0xdc, 0x5f, 0x2e, 0x4e, 0x81,
	0x2f, 0x6e, 0xae, 0x82, 0x13, 0xd8, 0x69, 0x81, 0x45, 0xc9, 0x0b, 0x81, 0xff, 0x2e, 0xcd, 0xef,
	0x1d, 0xd8, 0x51, 0xcd, 0x2f, 0x26, 0xda, 0x83, 0xf5, 0x8c, 0xe5, 0x4c, 0x6a, 0x4f, 0x27, 0x32,
	0x17, 0x32, 0x00, 0x97, 0x4f, 0xa7, 0x02, 0xa5, 0xa6, 0xc1, 0x89, 0xec, 0x4d, 0xd9, 0x05, 0xd2,
	0x2a, 0x7e, 0x6d, 0x39, 0xb0, 0x37, 0xf2, 0xdc, 0xea, 0xb1, 0xab, 0xf5, 0x38, 0xd4, 0xe9, 0x6f,
	0xe5, 0x5a, 0x96, 0xe6, 0x7f, 0x1f, 0xe6, 0x04, 0x48, 0x3b, 0xba, 0x65, 0xe1, 0x09, 0xf4, 0x25,
	0x97, 0x34, 0x1b, 0xc7, 0xbc, 0x2e, 0x9a, 0x86, 0x40, 0x9b, 0x5e, 0x28, 0x0b, 0x79, 0x0a, 0x6e,
	0x85, 0xa2, 0xce, 0x54, 0x57, 0xaa, 0xce, 0xdd, 0x16, 0x4d, 0x8d, 0x2c, 0x22, 0x0b, 0x09, 0x3e,
	0x82, 0xdd, 0x73, 0xcc, 0x70, 0x79, 0xa2, 0xab, 0x55, 0xa3, 0xe6, 0x7f, 0xa3, 0x25, 0x7e, 0xff,
	0xf9, 0x1f, 0xff, 0xe1, 0xc0, 0x96, 0x31, 0xa9, 0xcd, 0xa2, 0x96, 0xe4, 0x4b, 0x70, 0x8d, 0x9a,
	0x88, 0xa7, 0x1d, 0x56, 0x48, 0xcb, 0x1f, 0xdc, 0xfa, 0xdc, 0x2e, 0xd4, 0xc6, 0x0d, 0xf6, 0x7f,
	0xfd, 0xf3, 0xaf, 0xdf, 0xd6, 0x76, 0x82, 0xcd, 0xd6, 0x9e, 0x17, 0xa7, 0x9d, 0x23, 0x72, 0x0d,
	0xce, 0x25, 0x4a, 0x62, 0x96, 0xc5, 0xb2, 0xe0, 0xfc, 0xc1, 0xb2, 0xd9, 0x90, 0x1a, 0xbc, 0xab,
	0xc3, 0x79, 0x64, 0xd0, 0x0e, 0x17, 0xfe, 0x64, 0x39, 0xf8, 0x99, 0x7c, 0x05, 0x5d, 0x45, 0x1d,
	0x19, 0xac, 0x9e, 0xb9, 0xbf, 0x7f, 0xcb, 0x6e, 0x03, 0xef, 0xe9, 0xc0, 0xdb, 0x64, 0xa1, 0x4e,
	0xf2, 0x4a, 0xfd, 0x28, 0x14, 0xeb, 0xb6, 0xf3, 0x15, 0x23, 0xb8, 0xb3, 0x73, 0x5b, 0xea, 0xd1,
	0x5d, 0xa5, 0x26, 0xe0, 0x9a, 0x19, 0xd9, 0xd8, 0x2b, 0x06, 0x76, 0x67, 0xec, 0x43, 0x1d, 0x3b,
	0xf0, 0x0f, 0x6e, 0xc5, 0x56, 0x3b, 0xb7, 0x49, 0x71, 0xda, 0x39, 0xfa, 0xe2, 0x93, 0x57, 0xcf,
	0x53, 0x26, 0x5f, 0xd7, 0x93, 0x51, 0xcc, 0xf3, 0x30, 0xce, 0x78, 0x52, 0xb2, 0xb0, 0xcc, 0xa8,
	0x9c, 0xf2, 0x2a, 0xff, 0x58, 0x79, 0xa7, 0x3c, 0x9c, 0x3d, 0x0b, 0x97, 0xfe, 0xc8, 0x13, 0x57,
	0x67, 0x7c, 0xf6, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xbf, 0x5e, 0xef, 0xab, 0x07, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DeviceServiceClient interface {
	// Create creates the given device.
	Create(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Get returns the device matching the given DevEUI.
	Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error)
	// List returns the available devices.
	List(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceResponse, error)
	// Delete deletes the device matching the given DevEUI.
	Delete(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Update updates the device matching the given DevEUI.
	Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) Create(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*GetDeviceResponse, error) {
	out := new(GetDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) List(ctx context.Context, in *ListDeviceRequest, opts ...grpc.CallOption) (*ListDeviceResponse, error) {
	out := new(ListDeviceResponse)
	err := c.cc.Invoke(ctx, "/api.DeviceService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Delete(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.DeviceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
type DeviceServiceServer interface {
	// Create creates the given device.
	Create(context.Context, *CreateDeviceRequest) (*empty.Empty, error)
	// Get returns the device matching the given DevEUI.
	Get(context.Context, *GetDeviceRequest) (*GetDeviceResponse, error)
	// List returns the available devices.
	List(context.Context, *ListDeviceRequest) (*ListDeviceResponse, error)
	// Delete deletes the device matching the given DevEUI.
	Delete(context.Context, *DeleteDeviceRequest) (*empty.Empty, error)
	// Update updates the device matching the given DevEUI.
	Update(context.Context, *UpdateDeviceRequest) (*empty.Empty, error)
}

// UnimplementedDeviceServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDeviceServiceServer struct {
}

func (*UnimplementedDeviceServiceServer) Create(ctx context.Context, req *CreateDeviceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedDeviceServiceServer) Get(ctx context.Context, req *GetDeviceRequest) (*GetDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedDeviceServiceServer) List(ctx context.Context, req *ListDeviceRequest) (*ListDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedDeviceServiceServer) Delete(ctx context.Context, req *DeleteDeviceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedDeviceServiceServer) Update(ctx context.Context, req *UpdateDeviceRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterDeviceServiceServer(s *grpc.Server, srv DeviceServiceServer) {
	s.RegisterService(&_DeviceService_serviceDesc, srv)
}

func _DeviceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Create(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Get(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).List(ctx, req.(*ListDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Delete(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DeviceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Update(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DeviceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DeviceService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DeviceService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as/external/api/device.proto",
}
