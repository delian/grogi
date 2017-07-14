package httpserver

import (
	"../config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Run() {
	router := mux.NewRouter()
	port := config.Config.Httpserver.Port
	log.Println("HTTP Proxy started. Management is listening on", "0.0.0.0:"+strconv.Itoa(port))
	s := &http.Server{
		Addr:           "0.0.0.0:" + strconv.Itoa(port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
