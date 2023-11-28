package test

import (
	"context"
	"testing"

	"github.com/kluff-com/kluff-go"
	"github.com/kluff-com/kluff-go/data/db"
)

const testToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0eXBlIjoiYXV0aGVudGljYXRpb24iLCJpZCI6IjAwMDAwMDAwLTAwMDAtMDAwMC0wMDAwLTAwMDAwMDAwMDAwMCIsInVzZXJfaWQiOjIsIm9yZ2FuaXphdGlvbl9pZCI6MiwiYXBwX3Rva2VuIjp0cnVlLCJpc3N1ZWRfYXQiOiIyMDIzLTA5LTI4VDIzOjA2OjMxLjkwMzIzNVoiLCJleHBpcmVkX2F0IjoiMjAyMy0xMC0wMVQyMzowNjozMS45MDMyMzVaIn0.4p1CdOYKI3CPUMP2kAviT3MjVi8-iHbWsSiq1DHk_Ec"

func TestDB(t *testing.T) {
	sdk, err := kluff.Get(testToken)
	if err != nil {
		t.Fatal(err)
	}
	objectName := "st_test_users"

	ok, _ := sdk.ObjectExists(context.Background(), objectName)
	if ok {
		t.Errorf("%s must be exist", objectName)
	}
	// Create object
	err = sdk.CreateObject(context.Background(), &db.Object{
		Meta: &db.ObjectMeta{
			Name:        objectName,
			Description: "Some description",
		},
		Fields: []*db.Field{
			{
				Name:      "fullname",
				FieldType: "Text",
				Required:  true,
			},
			{
				Name:      "age",
				FieldType: "Int",
				Default:   "18",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	ok, _ = sdk.ObjectExists(context.Background(), objectName)
	if !ok {
		t.Errorf("there must an %s object", objectName)
	}
	// get the fields
	fields, err := sdk.GetFields(context.Background(), objectName)
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 2 {
		t.Error("invalid field length")
	}

	// create some records
	data := []map[string]any{
		{
			"fullname": "abel",
			"age":      21,
		},
		{
			"fullname": "liagiba",
			"age":      16,
		},
		{
			"fullname": "another name",
			"age":      34,
		},
	}

	for _, d := range data {
		_, err := sdk.CreateRecord(context.Background(), objectName, d)
		if err != nil {
			t.Error(err)
		}
	}

	// get objects
	records, err := sdk.GetRecords(context.Background(), &db.RecordQuery{
		APIName: objectName,
		Fields:  []string{"fullname", "age", "id"},
	})

	if err != nil {
		t.Error(err)
	}

	if len(records) != 3 {
		t.Error("records not up to")
	}

	// get a specific record
	record, err := sdk.GetARecord(context.Background(), &db.RecordQuery{
		APIName: objectName,
		Fields:  []string{"fullname"},
		Filters: []*db.Filter{
			{
				Field:    "fullname",
				Operator: "=",
				Value:    "abel",
			},
		},
	})

	if err != nil {
		t.Error(err)
	}

	if record["fullname"] != "abel" {
		t.Error("invalid record")
	}

	// Delete A record
	for _, v := range records {
		err := sdk.DeleteARecord(context.Background(), objectName, v["id"].(float64))
		if err != nil {
			t.Error(err)
		}
	}

	// check if all records are deleted
	records, err = sdk.GetRecords(context.Background(), &db.RecordQuery{
		APIName: objectName,
	})

	if err != nil {
		t.Error(err)
	}

	if len(records) != 0 {
		t.Error("records are not deleted")
	}

	if len(records) != 0 {
		t.Error("records not deleted")
	}

	err = sdk.DeleteObject(context.Background(), objectName)
	if err != nil {
		t.Error(err)
	}

}
