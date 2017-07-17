package httpserver

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type handlerStruct struct {
	gorilla http.Handler
}

var Handler handlerStruct = handlerStruct{
	gorilla: mux.NewRouter(),
}

func (h handlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("We got request. Route it to Gorilla")
	h.gorilla.ServeHTTP(w, r)
}
