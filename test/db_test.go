package test

import (
	"context"
	"testing"

	"github.com/kluff-com/kluff-go"
	"github.com/kluff-com/kluff-go/data/db"
)

func TestDB(t *testing.T) {
	sdk, err := kluff.New(kluff.Config{
		APIKey: "some api Key",
	})
	if err != nil {
		t.Error(err)
		_ = sdk
	}
	apiName := "st_test_users"

	ok, _ := sdk.ObjectExists(context.Background(), apiName)
	if ok {
		t.Errorf("%s must be exist", apiName)
	}
	// Create object
	err = sdk.CreateObject(context.Background(), &db.Object{
		Meta: &db.ObjectMeta{
			APIName:     apiName,
			Description: "Some description",
			Owner:       "abel",
		},
		Fields: []*db.Field{
			{
				FieldName: "fullname",
				FieldType: "Text",
				Required:  true,
			},
			{
				FieldName: "age",
				FieldType: "Int",
				Default:   "18",
			},
		},
	})
	if err != nil {
		t.Error(err)
	}

	defer sdk.DeleteObject(context.Background(), apiName)

	ok, _ = sdk.ObjectExists(context.Background(), apiName)
	if !ok {
		t.Errorf("there must an %s object", apiName)
	}
	// get the fields
	fields, err := sdk.GetFields(context.Background(), apiName)
	if err != nil {
		t.Error(err)
	}

	if len(fields) != 2 {
		t.Error("unable to find fields")
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
		_, err := sdk.CreateRecord(context.Background(), apiName, d)
		if err != nil {
			t.Error(err)
		}
	}

	// get objects
	records, err := sdk.GetRecords(context.Background(), &db.RecordQuery{
		APIName: apiName,
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
		APIName: apiName,
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
		err := sdk.DeleteARecord(context.Background(), apiName, v["id"].(float64))
		if err != nil {
			t.Error(err)
		}
	}

	// check if all records are deleted
	records, err = sdk.GetRecords(context.Background(), &db.RecordQuery{
		APIName: apiName,
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

}
