package main

import (
	"log"
	"net/http"

	"github.com/kluff-com/kluff-go"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		event, err := kluff.ParseEvent(w, r)
		if err != nil {
			w.WriteHeader(500)
			log.Println(err)
			w.Write([]byte("cannot parse event"))
			return
		}
		data, err := event.ParseTriggerData()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(data.Data)
		log.Println(event)
	})

	http.ListenAndServe(":5000", nil)

}
