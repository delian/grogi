package httpserver

import (
	"../config"
	"log"
	"net/http"
	"strconv"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {

}

func Run() {
	port := config.Config.Httpserver.Port
	log.Println("HTTP Proxy started. Management is listening on", "0.0.0.0:"+strconv.Itoa(port))
	s := &http.Server{
		Addr: "0.0.0.0:" + strconv.Itoa(port),
		//Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
