package httpserver

import "fmt"
import "../config"

func Test() {
	port := config.Config.Httpserver.Port
	fmt.Println("Test", port)
}
