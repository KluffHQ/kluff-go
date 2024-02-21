package kluff

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/kluffHQ/kluff-go/db"
	"google.golang.org/grpc"
)

type Interactor struct {
	cl db.DbClient
}

type ObjectSchema struct {
	Name          string      `json:"name"`
	CreatedByID   float64     `json:"created_by_id"`
	Fields        []*db.Field `json:"fields"`
	Owner         string      `json:"owner"`
	PluralLabel   string      `json:"plural_label"`
	SingularLabel string      `json:"singular_label"`
	TotalCount    float64     `json:"total_count"`
}

func NewDBInteractor(conn grpc.ClientConnInterface) Interactor {
	client := db.NewDbClient(conn)
	return Interactor{
		cl: client,
	}
}

func (i *Interactor) Object(name string) *Object {
	return &Object{
		cl: i.cl,
		Object: &db.Object{
			Base: &db.BaseObject{
				Name: name,
			},
		},
	}
}

func (i *Interactor) NewObject(obj *db.Object) *Object {
	return &Object{
		cl:     i.cl,
		Object: obj,
	}
}

func (i *Interactor) NewRecord(objectName string, data map[string]any) *Record {
	return &Record{
		cl:     i.cl,
		Object: objectName,
		Data:   data,
	}
}

func (i *Interactor) SendPing(context context.Context, ping *db.Ping) error {
	_, err := i.cl.SendPing(context, ping)
	return err
}

func (i *Interactor) GetObjects(context context.Context, q *db.ObjectQuery) ([]map[string]any, error) {
	data, err := i.cl.GetObjects(context, q)
	if err != nil {
		return nil, err
	}
	return data.ParseMapSlice()
}

func (i *Interactor) GetSession(context context.Context) (*db.Session, error) {
	return i.cl.GetSession(context, &db.Empty{})
}

/*
Object Methods
*/
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

func (o *Object) GetSingleRecord(q *db.RecordQuery) (*Record, error) {
	data, err := o.cl.GetSingleRecord(context.Background(), q)
	if err != nil {
		return nil, err
	}
	m, err := data.ParseMap()
	if err != nil {
		return nil, err
	}
	return newRecord(o.cl, o.Base.Name, m), nil
}

func (o *Object) CreateSingleRecord(q *db.SingleObject) (*db.SingleObject, error) {
	data, err := o.cl.CreateSingleObjectRecord(context.Background(), q)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (i *Interactor) ExecuteRawSql(q string, variables ...any) ([]map[string]any, error) {
	b, err := json.Marshal(&variables)
	if err != nil {
		return nil, err
	}
	data, err := i.cl.ExecuteRaw(context.Background(), &db.SqlQuery{
		Query:     q,
		Variables: b,
	})
	if err != nil {
		return nil, err
	}
	return data.ParseMapSlice()

}

func (o *Object) ExecuteSoql(q *db.String) (*Record, error) {
	data, err := o.cl.ExecuteSoql(context.Background(), q)
	if err != nil {
		return nil, err
	}
	m, err := data.ParseMap()
	if err != nil {
		return nil, err
	}
	return newRecord(o.cl, o.Base.Name, m), nil
}

func (o *Object) GetFields() ([]*db.Field, error) {
	data, err := o.cl.GetFields(context.Background(), &db.String{Value: o.Base.Name})
	return data.Fields, err
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

/*
Record Methods
*/
type Record struct {
	cl       db.DbClient
	Object   string `json:"-"`
	Data     map[string]any
	_updates map[string]any // this is used to store updates that can be saved by calling the save method
}

type Response struct {
	Data     []*Record `json:"data"`
	Count    float64   `json:"count"`
	PerPage  float64   `json:"per_page"`
	NextPage float64   `json:"next_page"`
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

func (r *Record) parseAllData() ([]byte, error) {
	update := r.Data
	for k, v := range r._updates {
		update[k] = v
	}
	return json.Marshal(update)
}

/*
Appends child record to the parent

Note: the fieldName has to have a parent-to-child relation with the main record
*/
func (r *Record) AppendChild(fieldName string, data []map[string]any) error {
	existingData := make([]map[string]any, 0)
	existing, ok := r.Data[fieldName]
	if !ok {
		r.Data[fieldName] = make([]map[string]any, 0)
	} else {
		single, ok := existing.(map[string]any)
		if ok {
			existingData = append(existingData, single)
		} else {
			mult, ok := existing.([]map[string]any)
			if ok {
				existingData = append(existingData, mult...)
			}
		}
	}

	if existingData != nil {
		data = append(data, existingData...)
	}
	r.Data[fieldName] = data
	return nil
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
