// Package used to read and export the default configuration of the application
package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"reflect"
)

const defaultFile = "./config.toml"

type httpServerType struct {
	Port  int
	Sport int
}

type ConfigType struct {
	Filename   string
	Httpserver httpServerType
}

// Config - Set the default values of the configuration
var Config ConfigType = ConfigType{
	Filename: defaultFile,
	Httpserver: httpServerType{
		1000,
		1000,
	},
}

type FlagsStr struct {
	help  string
	value interface{}
}

var FlagsMap = map[string]FlagsStr{
	"port":   FlagsStr{help: "The internal port of the Proxy", value: &Config.Httpserver.Port},
	"config": FlagsStr{help: "The configuration file name", value: &Config.Filename},
}

func runFlags() {
	for key, s := range FlagsMap {
		v := s.value
		fmt.Println(v)
		switch reflect.TypeOf(s.value).String() {
		default:
			log.Fatal("Unknown type in flagsMap", key, s.value)
		case "*bool":
			s.value.(*bool) = flag.Bool(key, *s.value.(*bool), s.help)
			fmt.Println("bool", key, *s.value.(*bool))
		case "*int":
			s.value.(*int) = flag.Int(key, *s.value.(*int), s.help)
			fmt.Println("int", key, s.value)
		case "*string":
			s.value.(*string) = flag.String(key, *s.value.(*string), s.help)
			fmt.Println("string", key, s.value)
		}
	}
	flag.Parse()
}

// Read method reads the configuration file specified
// or if it is empty string it will read the file stored in defaultFile
// or by -config command line property
func Read() {

	runFlags()

	if _, err := toml.DecodeFile(Config.Filename, &Config); err != nil {
		log.Fatal("Error reading the config file", err)
	}

	Config.Httpserver.Port = Config.Httpserver.Port
	fmt.Println("Port", Config.Httpserver.Port, Config.Httpserver.Sport)
}
