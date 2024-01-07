package main

import (
	"errors"
	"log"

	"github.com/kluff-com/kluff-go"
	"github.com/kluff-com/kluff-go/db"
)

func main() {
	r := kluff.NewRouter()
	r.GET("/logs", func(c *kluff.Context) {
		data, err := c.SDK.Object("tab_fields").GetRecords(&db.RecordQuery{
			Limit: 10,
		})
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, data.Result())
	})

	r.RegisterAction(kluff.Action{
		ID:     "convert lead",
		Name:   "Convert Lead",
		Object: "lead",
		Handler: func(i *kluff.Interactor, r *kluff.Record) error {
			return errors.New("new error")
		},
	})

	r.RegisterAction(kluff.Action{
		ID:     "Convert Lead 2",
		Name:   "Convert Lead to Something",
		Object: "lead",
		Handler: func(i *kluff.Interactor, r *kluff.Record) error {
			return nil
		},
	})
	log.Fatal(r.Start())
}
