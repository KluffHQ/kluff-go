package main

import (
	"log"

	"github.com/kluff-com/kluff-go/db"
	"github.com/kluff-com/kluff-go/router"
)

func main() {
	r := router.NewRouter()
	r.GET("/logs", func(c *router.Context) {
		data, err := c.SDK.Object("tab_fields").GetRecords(&db.RecordQuery{
			Limit: 10,
		})
		if err != nil {
			c.JSON(500, err.Error())
			return
		}
		c.JSON(200, data.Result())
	})
	log.Fatal(r.Start())
}
