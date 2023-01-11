package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

func Index(w http.ResponseWriter, p *http.Request, ps httprouter.Params) {
	logger, _ := zap.NewProduction()
	logger.Info("successfuly performed http request")
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
