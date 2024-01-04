package sdk

import (
	"context"
	"encoding/json"

	"github.com/kluff-com/kluff-go/dt"
	"google.golang.org/grpc"
)

type Interactor struct {
	cl dt.DbClient
}

type ObjectSchema struct {
	Name          string      `json:"name"`
	CreatedByID   float64     `json:"created_by_id"`
	Fields        []*dt.Field `json:"fields"`
	Owner         string      `json:"owner"`
	PluralLabel   string      `json:"plural_label"`
	SingularLabel string      `json:"singular_label"`
	TotalCount    float64     `json:"total_count"`
}

func NewDBInteractor(conn grpc.ClientConnInterface) Interactor {
	client := dt.NewDbClient(conn)
	return Interactor{
		cl: client,
	}
}

func unMarshalMap(data *dt.Data) (map[string]any, error) {
	v := map[string]any{}
	err := json.Unmarshal(data.Result, &v)
	return v, err
}

func unMarshalMapList(data *dt.Data) ([]map[string]any, error) {
	v := []map[string]any{}
	err := json.Unmarshal(data.Result, &v)
	return v, err
}

// New Methods
func (i *Interactor) Object(name string) *Object {
	return &Object{
		cl: &i.cl,
		Base: dt.ObjectBase{
			Name: name,
		},
	}
}

// Old Methods
func (i *Interactor) CreateObject(context context.Context, obj *dt.Object) error {
	_, err := i.cl.CreateObject(context, obj)
	return err
}

func (i *Interactor) ObjectExists(context context.Context, apiName string) (bool, error) {
	v, err := i.cl.ObjectExists(context, &dt.String{Value: apiName})
	if err != nil {
		return false, err
	}
	return v.Value, nil
}

func (i *Interactor) SendPing(context context.Context, ping *dt.Ping) error {
	_, err := i.cl.SendPing(context, ping)
	return err
}

func (i *Interactor) GetObjects(context context.Context, q *dt.ObjectQuery) ([]map[string]any, error) {
	data, err := i.cl.GetObjects(context, q)
	if err != nil {
		return nil, err
	}
	return unMarshalMapList(data)
}

func (i *Interactor) AddObjectFields(context context.Context, param *dt.ObjectFieldParam) (map[string]any, error) {
	data, err := i.cl.AddObjectFields(context, param)
	if err != nil {
		return nil, err
	}
	return unMarshalMap(data)
}

func (i *Interactor) UpdateObjectMeta(context context.Context, param *dt.ObjectBase) error {
	_, err := i.cl.UpdateObjectMeta(context, param)
	return err
}

// FIXME: data returned by json.Marshal is `float64` regardless, passing record record["id"] as `recID` directly panics
// unless you parse it as float64 and converting it to int64 to be parsed by the grpc
func (i *Interactor) DeleteARecord(context context.Context, apiName string, recID float64) error {
	_, err := i.cl.DeleteARecord(context, &dt.DeleteRecord{
		APIName:  apiName,
		RecordID: int64(recID),
	})
	return err
}

func (i *Interactor) DeleteObject(context context.Context, apiName string) error {
	_, err := i.cl.DeleteObject(context, &dt.String{
		Value: apiName,
	})
	return err
}

func (i *Interactor) DeleteFields(context context.Context, name string, fields ...string) error {
	_, err := i.cl.DeleteFields(context, &dt.DeleteField{
		Name:   name,
		Fields: fields,
	})
	return err
}

func (i *Interactor) UpdateObjectField(context context.Context, apiName string, fields []*dt.FieldUpdate) error {
	_, err := i.cl.UpdateObjectField(context, &dt.UpdateField{
		APIName: apiName,
		Fields:  fields,
	})
	return err
}

func (i *Interactor) GetObjectSchema(context context.Context, apiName string) (map[string]any, error) {
	data, err := i.cl.GetObjectSchema(context, &dt.String{
		Value: apiName,
	})
	if err != nil {
		return nil, err
	}
	return unMarshalMap(data)
}

func (i *Interactor) CreateRecord(context context.Context, Name string, data map[string]any) (map[string]any, error) {
	v, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	d, err := i.cl.CreateRecord(context, &dt.CreateRecordParam{
		ObjectName: Name,
		Fields:     v,
	})

	if err != nil {
		return nil, err
	}
	return unMarshalMap(d)
}

func (i *Interactor) GetARecord(context context.Context, q *dt.RecordQuery) (map[string]any, error) {
	data, err := i.cl.GetRecord(context, q)
	if err != nil {
		return nil, err
	}
	return unMarshalMap(data)
}

func (i *Interactor) GetRecords(context context.Context, q *dt.RecordQuery) ([]map[string]any, error) {
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
	_, err = i.cl.UpdateARecord(context, &dt.UpdateRecord{
		APIName:  apiName,
		RecordId: recordID,
		Options:  v,
	})
	return err
}

func (i *Interactor) DeleteRecord(context context.Context, apiName string, recordID int64) error {
	_, err := i.cl.DeleteARecord(context, &dt.DeleteRecord{
		APIName:  apiName,
		RecordID: recordID,
	})
	return err
}

func (i *Interactor) GetFields(context context.Context, apiName string) ([]*dt.Field, error) {
	data, err := i.cl.GetFields(context, &dt.String{
		Value: apiName,
	})
	if err != nil {
		return nil, err
	}

	// fmt.Println(string(data.Result))
	res := ObjectSchema{}
	err = json.Unmarshal(data.Result, &res)
	return res.Fields, err
}
