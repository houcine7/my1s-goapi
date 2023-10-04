package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/houcine7/my1s-goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) // to get the line number of the log
	// create a new router
	var router *chi.Mux = chi.NewRouter()
	handler.Handler(router)
	
	fmt.Println("Starting the API...")

	err :=http.ListenAndServe("localhost:8080",router)
	if err != nil{
		log.Fatal(err)
	}

}
