package httpserver

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strings"
)

type handlerStruct struct {
	gorilla http.Handler
	files   http.Handler
}

var Handler handlerStruct = handlerStruct{
	gorilla: mux.NewRouter(),
	files:   http.FileServer(http.Dir("/Users/delian/src/grogi")),
}

func (h handlerStruct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("We got request. Route it to Gorilla", r.RequestURI)

	if strings.HasPrefix(r.RequestURI, "http://") ||
		strings.HasPrefix(r.RequestURI, "https://") {
		log.Println("Must get from ", r.RequestURI)
		client := &http.Client{}
		if req, err := http.NewRequest(r.Method, r.RequestURI, nil); err == nil {
			if resp, err := client.Do(req); err == nil {
				log.Println("Got", resp)
				io.Copy(w, resp.Body)

			} else {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}

	} else {
		h.files.ServeHTTP(w, r)
	}

	//h.gorilla.ServeHTTP(w, r)
}
