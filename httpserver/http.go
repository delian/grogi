package httpserver

import (
	"../config"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Run() {
	port := config.Config.Httpserver.Port
	addr := config.Config.Httpserver.Address + ":" + strconv.Itoa(port)
	log.Println("HTTP Proxy started. Management is listening on", addr)
	s := &http.Server{
		Addr:              addr,
		Handler:           Handler,
		ReadTimeout:       time.Duration(config.Config.Httpserver.ReadTimeoutSec) * time.Second,
		WriteTimeout:      time.Duration(config.Config.Httpserver.WriteTimeoutSec) * time.Second,
		ReadHeaderTimeout: time.Duration(config.Config.Httpserver.ReadHeaderTimeoutSec) * time.Second,
		IdleTimeout:       time.Duration(config.Config.Httpserver.IdleTimeout) * time.Second,
		MaxHeaderBytes:    config.Config.Httpserver.MaxHeaderBytes,
	}

	http.Handle("/pesho", http.FileServer(http.Dir("/Users/delian/src/grogi")))
	log.Fatal(s.ListenAndServe())
}
