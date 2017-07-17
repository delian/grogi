all: build

godeps:
	go get github.com/gorilla/mux
	go get github.com/BurntSushi/toml

build: godeps
	go build

clean:
	rm -rf grogi

run:
	./grogi
