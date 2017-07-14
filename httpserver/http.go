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
	addr := config.Config.Httpserver.Address + ":" + strconv.Itoa(port)
	log.Println("HTTP Proxy started. Management is listening on", addr)
	s := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadTimeout:       time.Duration(config.Config.Httpserver.ReadTimeoutSec) * time.Second,
		WriteTimeout:      time.Duration(config.Config.Httpserver.WriteTimeoutSec) * time.Second,
		ReadHeaderTimeout: time.Duration(config.Config.Httpserver.ReadHeaderTimeoutSec) * time.Second,
		IdleTimeout:       time.Duration(config.Config.Httpserver.IdleTimeout) * time.Second,
		MaxHeaderBytes:    config.Config.Httpserver.MaxHeaderBytes,
	}
	log.Fatal(s.ListenAndServe())
}
