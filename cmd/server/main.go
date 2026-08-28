package main

import (
	"log"
	"net/http"
	"os"
	"repairdesk/api"
	"repairdesk/service"
	"repairdesk/storage"
)

func main() {
	path := os.Getenv("REPAIRDESK_DB")
	if path == "" {
		path = "repairdesk.db"
	}
	s, e := storage.Open(path)
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	d := service.New(s)
	log.Println(http.ListenAndServe(":8080", (&api.Server{Desk: d}).Handler()))
}
