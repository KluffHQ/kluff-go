package test

import (
	"testing"

	"github.com/kluff-com/kluff-go"
	"github.com/kluff-com/kluff-go/db"
)

const testToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbnYiOnsic2NoZW1hIjoicHVibGljIn0sInR5cGUiOiJhdXRoZW50aWNhdGlvbiIsImlkIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwidXNlcl9pZCI6MSwib3JnYW5pemF0aW9uX2lkIjoxLCJhcHBfdG9rZW4iOnRydWUsImlzc3VlZF9hdCI6IjIwMjQtMDEtMDZUMDA6MjQ6MzIuNzY4NjE5WiIsImV4cGlyZWRfYXQiOiIyMDI0LTAxLTA2VDA4OjI0OjMyLjc2ODYxOVoifQ.D3NTpBHGm80w5OEmN3oHkD5VO8aXdsQOMUE27t6nJpM"

func TestMethods(t *testing.T) {
	inter, err := kluff.Get(testToken)
	if err != nil {
		t.Log(err)
		t.Error(err)
	}

	// getting a record
	obj := inter.Object("page_layout")
	rec, err := obj.GetRecord(&db.RecordQuery{
		Filters: []*db.Filter{
			{
				Field:    "object_name",
				Operator: "=",
				Value:    "contract",
			},
		},
	})

	if err != nil {
		t.Error(err)
	}

	_, err = rec.Get("id")
	if err != nil {
		t.Error(err)
	}

	// getting records

	recs, err := obj.GetRecords(&db.RecordQuery{
		Limit: 10,
	})

	if err != nil {
		t.Error(err)
	}

	if len(recs.Data) != 10 {
		t.Error("limit not working")
	}

	// create object

	ob := inter.NewObject(&db.Object{
		Base: &db.BaseObject{
			Name:       "new_object",
			NamingRule: "random",
		},
		Fields: []*db.Field{
			{
				Name:      "age",
				FieldType: "Text",
				Label:     "Enter age",
			},
		},
	})

	e, err := ob.Exists()
	if err != nil {
		t.Error(err)
	}
	if e {
		err = ob.Delete()
		if err != nil {
			t.Error(err)
		}
	}

	err = ob.Create()
	if err != nil {
		t.Error(err)
	}

	err = ob.AddFields([]*db.Field{
		{
			Name:      "label",
			FieldType: "Int",
			Label:     "Enter some label",
		},
	})

	if err != nil {
		t.Error(err)
	}

	ex, err := ob.FieldExists("label")
	if err != nil {
		t.Error(err)
	}

	if !ex {
		t.Error("field not added")
	}
	rec, err = ob.Insert(map[string]any{
		"age":   "12",
		"label": 12,
	})

	if err != nil {
		t.Error(err)
	}

	rec.Set("age", "14")
	err = rec.Save()
	if err != nil {
		t.Error(err)
	}

	recs, err = ob.GetRecords(&db.RecordQuery{})
	if err != nil {
		return
	}

	if recs.Count != 1 {
		t.Error("there must be only one record")
	}

	age, err := recs.Data[0].Get("age")
	if err != nil {
		t.Error(err)
	}

	if age.(string) != "14" {
		t.Error("age must be updated to 14")
	}

	err = rec.Delete()
	if err != nil {
		t.Error(err)
	}

	err = ob.Delete()
	if err != nil {
		t.Error(err)
	}
}
