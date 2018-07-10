package config

var (
	Loaded bool
	Config *config
)

type config struct {
	Host    string
	Port    string
	Static  string
	Uploads string
	EP      *demoEndpointsConfig
}

type demoEndpointsConfig struct {
	Inception string
	Other     string
}

func SetDefaults() {
	Config = &config{
		Port:    "8888",
		Host:    "localhost",
		Static:  "./static",
		Uploads: "./static/demo",
		EP: &demoEndpointsConfig{
			Inception: "localhost:9000",
		},
	}
}

func URL() string {
	return "http://" + Config.Host + ":" + Config.Port + "/"
}
