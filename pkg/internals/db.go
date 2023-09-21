package internals

import (
	"context"
	"encoding/json"

	"github.com/kluff-com/kluff-go/data/db"
	"google.golang.org/grpc"
)

type session struct {
	*Interactor
}

type Interactor struct {
	cl db.DbClient
}

func NewDBInteractor(conn grpc.ClientConnInterface) Interactor {
	client := db.NewDbClient(conn)
	return Interactor{
		cl: client,
	}
}

func unMarshalMap(data *db.Data) (map[string]any, error) {
	v := map[string]any{}
	err := json.Unmarshal(data.Result, &v)
	return v, err
}
func unMarshalMapList(data *db.Data) ([]map[string]any, error) {
	v := []map[string]any{}
	err := json.Unmarshal(data.Result, &v)
	return v, err
}

func (i *Interactor) CreateObject(context context.Context, obj *db.Object) error {
	_, err := i.cl.CreateObject(context, obj)
	return err
}

func (i *Interactor) ObjectExists(context context.Context, apiName string) (bool, error) {
	v, err := i.cl.ObjectExists(context, &db.String{Value: apiName})
	if err != nil {
		return false, err
	}
	return v.Value, nil
}

func (i *Interactor) SendPing(context context.Context, ping *db.Ping) error {
	_, err := i.cl.SendPing(context, ping)
	return err
}

func (i *Interactor) AddObjectFields(context context.Context, param *db.ObjectFieldParam) (map[string]any, error) {
	data, err := i.cl.AddObjectFields(context, param)
	if err != nil {
		return nil, err
	}
	return unMarshalMap(data)
}

func (i *Interactor) UpdateObjectMeta(context context.Context, param *db.ObjectMeta) error {
	_, err := i.cl.UpdateObjectMeta(context, param)
	return err
}

// FIXME: data returned by json.Marshal is `float64` regardless, passing record record["id"] as `recID` directly panics
// unless you parse it as float64 and converting it to int64 to be parsed by the grpc
func (i *Interactor) DeleteARecord(context context.Context, apiName string, recID float64) error {
	_, err := i.cl.DeleteARecord(context, &db.DeleteRecord{
		APIName:  apiName,
		RecordID: int64(recID),
	})
	return err
}

func (i *Interactor) DeleteObject(context context.Context, apiName string) error {
	_, err := i.cl.DeleteObject(context, &db.String{
		Value: apiName,
	})
	return err
}

func (i *Interactor) DeleteFields(context context.Context, apiName string, fields ...string) error {
	_, err := i.cl.DeleteFields(context, &db.DeleteField{
		APIName: apiName,
		Fields:  fields,
	})
	return err
}

func (i *Interactor) UpdateObjectField(context context.Context, apiName string, fields []*db.FieldUpdate) error {
	_, err := i.cl.UpdateObjectField(context, &db.UpdateField{
		APIName: apiName,
		Fields:  fields,
	})
	return err
}

func (i *Interactor) GetObjectSchema(context context.Context, apiName string) (map[string]any, error) {
	data, err := i.cl.GetObjectSchema(context, &db.String{
		Value: apiName,
	})
	if err != nil {
		return nil, err
	}
	return unMarshalMap(data)
}

func (i *Interactor) CreateRecord(context context.Context, apiName string, data map[string]any) (map[string]any, error) {
	v, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	d, err := i.cl.CreateRecord(context, &db.CreateRecordParam{
		APIName: apiName,
		Owner:   "some owner",
		Fields:  v,
	})

	if err != nil {
		return nil, err
	}
	return unMarshalMap(d)
}

func (i *Interactor) GetARecord(context context.Context, q *db.RecordQuery) (map[string]any, error) {
	data, err := i.cl.GetRecord(context, q)
	if err != nil {
		return nil, err
	}
	return unMarshalMap(data)
}

func (i *Interactor) GetRecords(context context.Context, q *db.RecordQuery) ([]map[string]any, error) {
	data, err := i.cl.GetAllRecords(context, q)
	if err != nil {
		return nil, err
	}
	return unMarshalMapList(data)
}

func (i *Interactor) UpdateRecord(context context.Context, apiName string, recordID int64, data map[string]any) error {
	v, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = i.cl.UpdateARecord(context, &db.UpdateRecord{
		APIName:  apiName,
		RecordId: recordID,
		Options:  v,
	})
	return err
}

func (i *Interactor) DeleteRecord(context context.Context, apiName string, recordID int64) error {
	_, err := i.cl.DeleteARecord(context, &db.DeleteRecord{
		APIName:  apiName,
		RecordID: recordID,
	})
	return err
}

func (i *Interactor) GetFields(context context.Context, apiName string) ([]*db.Field, error) {
	data, err := i.cl.GetFields(context, &db.String{
		Value: apiName,
	})
	if err != nil {
		return nil, err
	}
	res := []*db.Field{}
	err = json.Unmarshal(data.Result, &res)
	return res, err
}
