// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: db.proto

package db

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Db_SendPing_FullMethodName                 = "/db.db/SendPing"
	Db_CreateObject_FullMethodName             = "/db.db/CreateObject"
	Db_AddObjectFields_FullMethodName          = "/db.db/AddObjectFields"
	Db_GetObjects_FullMethodName               = "/db.db/GetObjects"
	Db_UpdateObjectMeta_FullMethodName         = "/db.db/UpdateObjectMeta"
	Db_DeleteObject_FullMethodName             = "/db.db/DeleteObject"
	Db_DeleteFields_FullMethodName             = "/db.db/DeleteFields"
	Db_UpdateObjectField_FullMethodName        = "/db.db/UpdateObjectField"
	Db_GetObjectSchema_FullMethodName          = "/db.db/GetObjectSchema"
	Db_ObjectExists_FullMethodName             = "/db.db/ObjectExists"
	Db_GetRelatedFields_FullMethodName         = "/db.db/GetRelatedFields"
	Db_GetLatestRecord_FullMethodName          = "/db.db/GetLatestRecord"
	Db_ObjectFieldExists_FullMethodName        = "/db.db/ObjectFieldExists"
	Db_GetFields_FullMethodName                = "/db.db/GetFields"
	Db_CreateRecord_FullMethodName             = "/db.db/CreateRecord"
	Db_GetRecord_FullMethodName                = "/db.db/GetRecord"
	Db_UpdateARecord_FullMethodName            = "/db.db/UpdateARecord"
	Db_DeleteARecord_FullMethodName            = "/db.db/DeleteARecord"
	Db_GetAllRecords_FullMethodName            = "/db.db/GetAllRecords"
	Db_GetSingleRecord_FullMethodName          = "/db.db/GetSingleRecord"
	Db_CreateSingleObjectRecord_FullMethodName = "/db.db/CreateSingleObjectRecord"
	Db_ExcuteRaw_FullMethodName                = "/db.db/ExcuteRaw"
	Db_ExcuteSoql_FullMethodName               = "/db.db/ExcuteSoql"
)

// DbClient is the client API for Db service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DbClient interface {
	SendPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
	// Object Methods
	CreateObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*Object, error)
	AddObjectFields(ctx context.Context, in *ObjectFieldParam, opts ...grpc.CallOption) (*Data, error)
	GetObjects(ctx context.Context, in *ObjectQuery, opts ...grpc.CallOption) (*Data, error)
	UpdateObjectMeta(ctx context.Context, in *BaseObject, opts ...grpc.CallOption) (*Empty, error)
	DeleteObject(ctx context.Context, in *String, opts ...grpc.CallOption) (*Empty, error)
	DeleteFields(ctx context.Context, in *DeleteField, opts ...grpc.CallOption) (*Empty, error)
	UpdateObjectField(ctx context.Context, in *UpdateField, opts ...grpc.CallOption) (*Empty, error)
	GetObjectSchema(ctx context.Context, in *String, opts ...grpc.CallOption) (*BaseObject, error)
	ObjectExists(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error)
	GetRelatedFields(ctx context.Context, in *String, opts ...grpc.CallOption) (*Fields, error)
	GetLatestRecord(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error)
	ObjectFieldExists(ctx context.Context, in *FieldData, opts ...grpc.CallOption) (*Bool, error)
	GetFields(ctx context.Context, in *String, opts ...grpc.CallOption) (*Fields, error)
	// Record Methods
	CreateRecord(ctx context.Context, in *CreateRecordParam, opts ...grpc.CallOption) (*Data, error)
	GetRecord(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error)
	UpdateARecord(ctx context.Context, in *UpdateRecord, opts ...grpc.CallOption) (*Empty, error)
	DeleteARecord(ctx context.Context, in *RecordData, opts ...grpc.CallOption) (*Empty, error)
	GetAllRecords(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error)
	GetSingleRecord(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error)
	CreateSingleObjectRecord(ctx context.Context, in *SingleObject, opts ...grpc.CallOption) (*SingleObject, error)
	ExcuteRaw(ctx context.Context, in *String, opts ...grpc.CallOption) (*Data, error)
	ExcuteSoql(ctx context.Context, in *String, opts ...grpc.CallOption) (*Data, error)
}

type dbClient struct {
	cc grpc.ClientConnInterface
}

func NewDbClient(cc grpc.ClientConnInterface) DbClient {
	return &dbClient{cc}
}

func (c *dbClient) SendPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, Db_SendPing_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) CreateObject(ctx context.Context, in *Object, opts ...grpc.CallOption) (*Object, error) {
	out := new(Object)
	err := c.cc.Invoke(ctx, Db_CreateObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) AddObjectFields(ctx context.Context, in *ObjectFieldParam, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_AddObjectFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetObjects(ctx context.Context, in *ObjectQuery, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_GetObjects_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) UpdateObjectMeta(ctx context.Context, in *BaseObject, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Db_UpdateObjectMeta_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) DeleteObject(ctx context.Context, in *String, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Db_DeleteObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) DeleteFields(ctx context.Context, in *DeleteField, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Db_DeleteFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) UpdateObjectField(ctx context.Context, in *UpdateField, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Db_UpdateObjectField_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetObjectSchema(ctx context.Context, in *String, opts ...grpc.CallOption) (*BaseObject, error) {
	out := new(BaseObject)
	err := c.cc.Invoke(ctx, Db_GetObjectSchema_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) ObjectExists(ctx context.Context, in *String, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, Db_ObjectExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetRelatedFields(ctx context.Context, in *String, opts ...grpc.CallOption) (*Fields, error) {
	out := new(Fields)
	err := c.cc.Invoke(ctx, Db_GetRelatedFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetLatestRecord(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_GetLatestRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) ObjectFieldExists(ctx context.Context, in *FieldData, opts ...grpc.CallOption) (*Bool, error) {
	out := new(Bool)
	err := c.cc.Invoke(ctx, Db_ObjectFieldExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetFields(ctx context.Context, in *String, opts ...grpc.CallOption) (*Fields, error) {
	out := new(Fields)
	err := c.cc.Invoke(ctx, Db_GetFields_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) CreateRecord(ctx context.Context, in *CreateRecordParam, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_CreateRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetRecord(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_GetRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) UpdateARecord(ctx context.Context, in *UpdateRecord, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Db_UpdateARecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) DeleteARecord(ctx context.Context, in *RecordData, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Db_DeleteARecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetAllRecords(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_GetAllRecords_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) GetSingleRecord(ctx context.Context, in *RecordQuery, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_GetSingleRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) CreateSingleObjectRecord(ctx context.Context, in *SingleObject, opts ...grpc.CallOption) (*SingleObject, error) {
	out := new(SingleObject)
	err := c.cc.Invoke(ctx, Db_CreateSingleObjectRecord_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) ExcuteRaw(ctx context.Context, in *String, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_ExcuteRaw_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dbClient) ExcuteSoql(ctx context.Context, in *String, opts ...grpc.CallOption) (*Data, error) {
	out := new(Data)
	err := c.cc.Invoke(ctx, Db_ExcuteSoql_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DbServer is the server API for Db service.
// All implementations must embed UnimplementedDbServer
// for forward compatibility
type DbServer interface {
	SendPing(context.Context, *Ping) (*Pong, error)
	// Object Methods
	CreateObject(context.Context, *Object) (*Object, error)
	AddObjectFields(context.Context, *ObjectFieldParam) (*Data, error)
	GetObjects(context.Context, *ObjectQuery) (*Data, error)
	UpdateObjectMeta(context.Context, *BaseObject) (*Empty, error)
	DeleteObject(context.Context, *String) (*Empty, error)
	DeleteFields(context.Context, *DeleteField) (*Empty, error)
	UpdateObjectField(context.Context, *UpdateField) (*Empty, error)
	GetObjectSchema(context.Context, *String) (*BaseObject, error)
	ObjectExists(context.Context, *String) (*Bool, error)
	GetRelatedFields(context.Context, *String) (*Fields, error)
	GetLatestRecord(context.Context, *RecordQuery) (*Data, error)
	ObjectFieldExists(context.Context, *FieldData) (*Bool, error)
	GetFields(context.Context, *String) (*Fields, error)
	// Record Methods
	CreateRecord(context.Context, *CreateRecordParam) (*Data, error)
	GetRecord(context.Context, *RecordQuery) (*Data, error)
	UpdateARecord(context.Context, *UpdateRecord) (*Empty, error)
	DeleteARecord(context.Context, *RecordData) (*Empty, error)
	GetAllRecords(context.Context, *RecordQuery) (*Data, error)
	GetSingleRecord(context.Context, *RecordQuery) (*Data, error)
	CreateSingleObjectRecord(context.Context, *SingleObject) (*SingleObject, error)
	ExcuteRaw(context.Context, *String) (*Data, error)
	ExcuteSoql(context.Context, *String) (*Data, error)
	mustEmbedUnimplementedDbServer()
}

// UnimplementedDbServer must be embedded to have forward compatible implementations.
type UnimplementedDbServer struct {
}

func (UnimplementedDbServer) SendPing(context.Context, *Ping) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPing not implemented")
}
func (UnimplementedDbServer) CreateObject(context.Context, *Object) (*Object, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateObject not implemented")
}
func (UnimplementedDbServer) AddObjectFields(context.Context, *ObjectFieldParam) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddObjectFields not implemented")
}
func (UnimplementedDbServer) GetObjects(context.Context, *ObjectQuery) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjects not implemented")
}
func (UnimplementedDbServer) UpdateObjectMeta(context.Context, *BaseObject) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjectMeta not implemented")
}
func (UnimplementedDbServer) DeleteObject(context.Context, *String) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteObject not implemented")
}
func (UnimplementedDbServer) DeleteFields(context.Context, *DeleteField) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFields not implemented")
}
func (UnimplementedDbServer) UpdateObjectField(context.Context, *UpdateField) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObjectField not implemented")
}
func (UnimplementedDbServer) GetObjectSchema(context.Context, *String) (*BaseObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetObjectSchema not implemented")
}
func (UnimplementedDbServer) ObjectExists(context.Context, *String) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectExists not implemented")
}
func (UnimplementedDbServer) GetRelatedFields(context.Context, *String) (*Fields, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRelatedFields not implemented")
}
func (UnimplementedDbServer) GetLatestRecord(context.Context, *RecordQuery) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestRecord not implemented")
}
func (UnimplementedDbServer) ObjectFieldExists(context.Context, *FieldData) (*Bool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObjectFieldExists not implemented")
}
func (UnimplementedDbServer) GetFields(context.Context, *String) (*Fields, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFields not implemented")
}
func (UnimplementedDbServer) CreateRecord(context.Context, *CreateRecordParam) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRecord not implemented")
}
func (UnimplementedDbServer) GetRecord(context.Context, *RecordQuery) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecord not implemented")
}
func (UnimplementedDbServer) UpdateARecord(context.Context, *UpdateRecord) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateARecord not implemented")
}
func (UnimplementedDbServer) DeleteARecord(context.Context, *RecordData) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteARecord not implemented")
}
func (UnimplementedDbServer) GetAllRecords(context.Context, *RecordQuery) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRecords not implemented")
}
func (UnimplementedDbServer) GetSingleRecord(context.Context, *RecordQuery) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSingleRecord not implemented")
}
func (UnimplementedDbServer) CreateSingleObjectRecord(context.Context, *SingleObject) (*SingleObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSingleObjectRecord not implemented")
}
func (UnimplementedDbServer) ExcuteRaw(context.Context, *String) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExcuteRaw not implemented")
}
func (UnimplementedDbServer) ExcuteSoql(context.Context, *String) (*Data, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExcuteSoql not implemented")
}
func (UnimplementedDbServer) mustEmbedUnimplementedDbServer() {}

// UnsafeDbServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DbServer will
// result in compilation errors.
type UnsafeDbServer interface {
	mustEmbedUnimplementedDbServer()
}

func RegisterDbServer(s grpc.ServiceRegistrar, srv DbServer) {
	s.RegisterService(&Db_ServiceDesc, srv)
}

func _Db_SendPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).SendPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_SendPing_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).SendPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_CreateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Object)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).CreateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_CreateObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).CreateObject(ctx, req.(*Object))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_AddObjectFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectFieldParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).AddObjectFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_AddObjectFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).AddObjectFields(ctx, req.(*ObjectFieldParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetObjects_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetObjects(ctx, req.(*ObjectQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_UpdateObjectMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).UpdateObjectMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_UpdateObjectMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).UpdateObjectMeta(ctx, req.(*BaseObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_DeleteObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).DeleteObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_DeleteObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).DeleteObject(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_DeleteFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteField)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).DeleteFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_DeleteFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).DeleteFields(ctx, req.(*DeleteField))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_UpdateObjectField_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateField)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).UpdateObjectField(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_UpdateObjectField_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).UpdateObjectField(ctx, req.(*UpdateField))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetObjectSchema_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetObjectSchema(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetObjectSchema_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetObjectSchema(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_ObjectExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).ObjectExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_ObjectExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).ObjectExists(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetRelatedFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetRelatedFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetRelatedFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetRelatedFields(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetLatestRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetLatestRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetLatestRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetLatestRecord(ctx, req.(*RecordQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_ObjectFieldExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FieldData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).ObjectFieldExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_ObjectFieldExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).ObjectFieldExists(ctx, req.(*FieldData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetFields_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetFields(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetFields_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetFields(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_CreateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRecordParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).CreateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_CreateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).CreateRecord(ctx, req.(*CreateRecordParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetRecord(ctx, req.(*RecordQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_UpdateARecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).UpdateARecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_UpdateARecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).UpdateARecord(ctx, req.(*UpdateRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_DeleteARecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).DeleteARecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_DeleteARecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).DeleteARecord(ctx, req.(*RecordData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetAllRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetAllRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetAllRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetAllRecords(ctx, req.(*RecordQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_GetSingleRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecordQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).GetSingleRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_GetSingleRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).GetSingleRecord(ctx, req.(*RecordQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_CreateSingleObjectRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SingleObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).CreateSingleObjectRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_CreateSingleObjectRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).CreateSingleObjectRecord(ctx, req.(*SingleObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_ExcuteRaw_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).ExcuteRaw(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_ExcuteRaw_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).ExcuteRaw(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _Db_ExcuteSoql_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DbServer).ExcuteSoql(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Db_ExcuteSoql_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DbServer).ExcuteSoql(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

// Db_ServiceDesc is the grpc.ServiceDesc for Db service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Db_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "db.db",
	HandlerType: (*DbServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPing",
			Handler:    _Db_SendPing_Handler,
		},
		{
			MethodName: "CreateObject",
			Handler:    _Db_CreateObject_Handler,
		},
		{
			MethodName: "AddObjectFields",
			Handler:    _Db_AddObjectFields_Handler,
		},
		{
			MethodName: "GetObjects",
			Handler:    _Db_GetObjects_Handler,
		},
		{
			MethodName: "UpdateObjectMeta",
			Handler:    _Db_UpdateObjectMeta_Handler,
		},
		{
			MethodName: "DeleteObject",
			Handler:    _Db_DeleteObject_Handler,
		},
		{
			MethodName: "DeleteFields",
			Handler:    _Db_DeleteFields_Handler,
		},
		{
			MethodName: "UpdateObjectField",
			Handler:    _Db_UpdateObjectField_Handler,
		},
		{
			MethodName: "GetObjectSchema",
			Handler:    _Db_GetObjectSchema_Handler,
		},
		{
			MethodName: "ObjectExists",
			Handler:    _Db_ObjectExists_Handler,
		},
		{
			MethodName: "GetRelatedFields",
			Handler:    _Db_GetRelatedFields_Handler,
		},
		{
			MethodName: "GetLatestRecord",
			Handler:    _Db_GetLatestRecord_Handler,
		},
		{
			MethodName: "ObjectFieldExists",
			Handler:    _Db_ObjectFieldExists_Handler,
		},
		{
			MethodName: "GetFields",
			Handler:    _Db_GetFields_Handler,
		},
		{
			MethodName: "CreateRecord",
			Handler:    _Db_CreateRecord_Handler,
		},
		{
			MethodName: "GetRecord",
			Handler:    _Db_GetRecord_Handler,
		},
		{
			MethodName: "UpdateARecord",
			Handler:    _Db_UpdateARecord_Handler,
		},
		{
			MethodName: "DeleteARecord",
			Handler:    _Db_DeleteARecord_Handler,
		},
		{
			MethodName: "GetAllRecords",
			Handler:    _Db_GetAllRecords_Handler,
		},
		{
			MethodName: "GetSingleRecord",
			Handler:    _Db_GetSingleRecord_Handler,
		},
		{
			MethodName: "CreateSingleObjectRecord",
			Handler:    _Db_CreateSingleObjectRecord_Handler,
		},
		{
			MethodName: "ExcuteRaw",
			Handler:    _Db_ExcuteRaw_Handler,
		},
		{
			MethodName: "ExcuteSoql",
			Handler:    _Db_ExcuteSoql_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "db.proto",
}
