package config

const defaultFile = "./config.toml"

type httpServerType struct {
	Port                 int
	Address              string
	ReadTimeoutSec       int
	WriteTimeoutSec      int
	ReadHeaderTimeoutSec int
	IdleTimeout          int
	MaxHeaderBytes       int
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
		"0.0.0.0",
		10,
		10,
		5,
		30,
		1 << 20, // 1MB
	},
}

type FlagsStr struct {
	help  string
	value interface{}
}

var FlagsMap = map[string]FlagsStr{
	"port":           FlagsStr{"The internal port of the Proxy management", &Config.Httpserver.Port},
	"bind":           FlagsStr{"The ip address where the proxy management will listen", &Config.Httpserver.Address},
	"rtimeout":       FlagsStr{"The HTTP read timeout in seconds", &Config.Httpserver.ReadTimeoutSec},
	"wtimeout":       FlagsStr{"The HTTP write timeout in seconds", &Config.Httpserver.WriteTimeoutSec},
	"maxheaderbytes": FlagsStr{"The maximum header size", &Config.Httpserver.MaxHeaderBytes},
	//	"config": FlagsStr{help: "The configuration file name", value: &Config.Filename}, // Not necessary
}
