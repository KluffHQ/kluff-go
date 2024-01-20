package main

import (
	"errors"
	"log"

	"github.com/kluff-com/kluff-go"
)

func main() {
	r := kluff.NewRouter()
	r.GET("/logs", func(c *kluff.Context) {
		data, err := c.Inter.Object("page_layout").GetFields()
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, data)
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

	r.RegisterTrigger(kluff.Trigger{
		ID:     "new_id",
		Action: kluff.ON_CREATE,
		Handler: func(i *kluff.Context, m map[string]any) error {
			return nil
		},
	})
	log.Fatal(r.Start())
}
