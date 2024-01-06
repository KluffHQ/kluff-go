package sdk

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kluff-com/kluff-go/db"
)

type Record struct {
	cl       db.DbClient
	Object   string
	Data     map[string]any
	_updates map[string]any // this is used to store updates that can be saved by calling the save method
}

type Response struct {
	Data     []*Record
	Count    float64
	PerPage  float64
	NextPage float64
}

func (r *Response) Result() []map[string]any {
	res := []map[string]any{}
	for _, v := range r.Data {
		res = append(res, v.Data)
	}
	return res
}

func ParseResponse(object *Object, resp map[string]any) (*Response, error) {
	data := []*Record{}

	d, ok := resp["data"].([]any)
	if ok {
		for _, v := range d {
			rec, ok := v.(map[string]any)
			if ok {
				data = append(data, newRecord(object.cl, object.Base.Name, rec))
			}
		}
	}
	count, ok := resp["count"].(float64)
	if !ok {
		return nil, errors.New("invalid response")
	}
	perPage, ok := resp["per_page"].(float64)
	if !ok {
		return nil, errors.New("invalid response")
	}
	nextPage, ok := resp["next_page"].(float64)
	if !ok {
		return nil, errors.New("invalid response")
	}
	return &Response{
		Data:     data,
		Count:    count,
		PerPage:  perPage,
		NextPage: nextPage,
	}, nil

}

func newRecord(cl db.DbClient, object string, data map[string]any) *Record {
	return &Record{
		cl:       cl,
		Data:     data,
		Object:   object,
		_updates: map[string]any{},
	}
}

func (r *Record) hasID() bool {
	_, ok := r.Data["id"]
	return ok
}

func (r *Record) parseAllData() ([]byte, error) {
	update := r.Data
	for k, v := range r._updates {
		update[k] = v
	}
	return json.Marshal(update)
}

func (r *Record) Get(field string) (any, error) {
	// if the value is not found in the updates get the read data
	d, ok := r._updates[field]
	if ok {
		return d, nil
	}
	d, ok = r.Data[field]
	if !ok {
		return nil, errors.New("field not found")
	}
	return d, nil
}

func (r *Record) Set(field string, data any) *Record {
	if field != "" && field != "id" {
		r._updates[field] = data
	}
	return r
}

func (r *Record) SetAll(data map[string]any) *Record {
	if len(data) > 0 {
		for k, v := range data {
			r.Set(k, v)
		}
	}
	return r
}

func (r *Record) Create() error {
	fields, err := r.parseAllData()
	if err != nil {
		return err
	}
	data, err := r.cl.CreateRecord(context.Background(), &db.CreateRecordParam{
		ObjectName: r.Object,
		Fields:     fields,
	})
	if err != nil {
		return err
	}
	dt, err := data.ParseMap()
	if err != nil {
		return err
	}
	r.Data = dt
	r._updates = map[string]any{}
	return nil
}

func (r *Record) getID() (int64, error) {
	res, err := r.Get("id")
	if err != nil {
		return 0, err
	}
	id, ok := res.(float64)
	if !ok {
		return 0, errors.New("record id must be `int64`")
	}
	return int64(id), nil
}

// save the record data. ND: the values in set using the `Set` or `SetAll` methods is not saved, call the `Save` method to save those changes
func (r *Record) Update(opt map[string]any) error {
	id, err := r.getID()
	if err != nil {
		return err
	}
	data, err := json.Marshal(opt)
	if err != nil {
		return err
	}
	_, err = r.cl.UpdateARecord(context.Background(), &db.UpdateRecord{
		Name:     r.Object,
		RecordId: id,
		Options:  data,
	})

	if err != nil {
		return err
	}
	// merge the update and data
	//TODO: This might break for formula field, return the updated data after update
	for k, v := range opt {
		r.Data[k] = v
	}
	return nil
}

// save all updates which are made using the `Set` and `SetAll` method
func (r *Record) Save() error {
	err := r.Update(r._updates)
	if err != nil {
		return err
	}
	r._updates = map[string]any{}
	return nil
}

func (r *Record) Delete() error {
	id, err := r.getID()
	if err != nil {
		return err
	}
	_, err = r.cl.DeleteARecord(context.Background(), &db.RecordData{
		Name:     r.Object,
		RecordID: id,
	})
	return err
}
