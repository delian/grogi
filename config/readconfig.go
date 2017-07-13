// Package used to read and export the default configuration of the application
package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"reflect"
)

func runFlags() {
	for key, s := range FlagsMap {
		switch reflect.TypeOf(s.value).String() {
		case "*bool":
			flag.BoolVar(s.value.(*bool), key, *s.value.(*bool), s.help)
		case "*int":
			flag.IntVar(s.value.(*int), key, *s.value.(*int), s.help)
		case "*string":
			flag.StringVar(s.value.(*string), key, *s.value.(*string), s.help)
		default:
			log.Fatal("Unknown type in flagsMap", key, s.value)
		}
	}
	flag.Parse()
}

// Read method reads the configuration file specified
// or if it is empty string it will read the file stored in defaultFile
// or by -config command line property
func Read() {

	if _, err := toml.DecodeFile(Config.Filename, &Config); err != nil {
		log.Fatal("Error reading the config file", err)
	}

	runFlags()

	fmt.Println("Port", Config.Httpserver.Port, Config.Httpserver.Sport)
}
