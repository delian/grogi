package config

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
