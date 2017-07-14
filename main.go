package main

import (
	"./config"
	"./httpserver"
)

func main() {
	config.Read()
	httpserver.Run()
}
