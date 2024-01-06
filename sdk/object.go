package sdk

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kluff-com/kluff-go/db"
)

type Object struct {
	*db.Object
	cl db.DbClient
}

func (o *Object) validate() bool {
	if o.cl != nil && o.Base.Name != "" {
		return true
	}
	return false
}

// Load Object metadata if the object exists.
// if fields are to the object it will update the object with the `saved` fields
func (o *Object) LoadMeta() error {
	// TODO: get all meta and populate the object
	return nil
}

// Create object and update it's meta according to the created objects
func (o *Object) Create() error {
	if !o.validate() {
		return errors.New("object name is required")
	}
	obj, err := o.cl.CreateObject(context.Background(), o.Object)
	if err != nil {
		return err
	}
	o.Object = obj
	return nil
}

// Get a record
func (o *Object) GetRecord(q *db.RecordQuery) (*Record, error) {
	q.Name = o.Base.Name
	data, err := o.cl.GetRecord(context.Background(), q)
	if err != nil {
		return nil, err
	}
	m, err := data.ParseMap()
	if err != nil {
		return nil, err
	}
	return newRecord(o.cl, o.Base.Name, m), nil
}

// Query for records
func (o *Object) GetRecords(q *db.RecordQuery) (*Response, error) {
	q.Name = o.Base.Name
	d, err := o.cl.GetAllRecords(context.Background(), q)
	if err != nil {
		return nil, err
	}
	dt, err := d.ParseMap()
	if err != nil {
		return nil, err
	}
	return ParseResponse(o, dt)
}

func (o *Object) GetLatestRecord(q *db.RecordQuery) (*Record, error) {
	data, err := o.cl.GetLatestRecord(context.Background(), q)
	if err != nil {
		return nil, err
	}
	m, err := data.ParseMap()
	if err != nil {
		return nil, err
	}
	return newRecord(o.cl, o.Base.Name, m), nil
}

func (o *Object) FieldExists(fieldName string) (bool, error) {
	v, err := o.cl.ObjectFieldExists(context.Background(), &db.FieldData{
		Object:    o.Base.Name,
		FieldName: fieldName,
	})
	if err != nil {
		return false, err
	}
	return v.Value, nil
}

func (o *Object) AddFields(f []*db.Field) error {
	_, err := o.cl.AddObjectFields(context.Background(), &db.ObjectFieldParam{
		Fields: f,
		Name:   o.Base.Name,
	})
	return err
}

func (o *Object) DeleteFields(fields []string) error {
	_, err := o.cl.DeleteFields(context.Background(), &db.DeleteField{
		Fields: fields,
		Name:   o.Base.Name,
	})
	return err
}

func (o *Object) UpdateFields(update []*db.FieldUpdate) error {
	_, err := o.cl.UpdateObjectField(context.Background(), &db.UpdateField{
		Name:   o.Base.Name,
		Fields: update,
	})
	return err
}

func (o *Object) UpdateMeta(obj *db.BaseObject) error {
	_, err := o.cl.UpdateObjectMeta(context.Background(), obj)
	return err
}

func (o *Object) GetRelatedFields() ([]*db.Field, error) {
	f, err := o.cl.GetRelatedFields(context.Background(), &db.String{Value: o.Base.Name})
	return f.Fields, err
}

func (o *Object) Exists() (bool, error) {
	v, err := o.cl.ObjectExists(context.Background(), &db.String{Value: o.Base.Name})
	if err != nil {
		return false, err
	}
	return v.Value, err
}

func (o *Object) Delete() error {
	_, err := o.cl.DeleteObject(context.Background(), &db.String{Value: o.Base.Name})
	return err
}

func (o *Object) Insert(data map[string]any) (*Record, error) {
	v, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	res, err := o.cl.CreateRecord(context.Background(), &db.CreateRecordParam{
		ObjectName: o.Base.Name,
		Fields:     v,
	})
	if err != nil {
		return nil, err
	}
	m, err := res.ParseMap()
	if err != nil {
		return nil, err
	}
	return newRecord(o.cl, o.Base.Name, m), nil
}
