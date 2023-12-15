package config

var Config *config

type config struct {
	Server struct {
		Addr string `env:"SERVER_ADDR" envDefault:"localhost:8080"`
	}

	Mongo struct {
		Host     string `env:"MONGO_HOST" envDefault:"localhost"`
		Port     string `env:"MONGO_PORT" envDefault:"27017"`
		Username string `env:"MONGO_USERNAME" envDefault:"mongo"`
		Password string `env:"MONGO_PASSWORD" envDefault:"mongo"`
		Database string `env:"MONGO_DATABASE" envDefault:"test"`

		ConnectWaitTime   int `env:"MONGO_CONNECT_WAIT_TIME" envDefault:"5"`
		ConnectionTimeout int `env:"MONGO_CONNECTION_TIMEOUT" envDefault:"10"`
		ConnectAttempts   int `env:"MONGO_CONNECT_ATTEMPTS" envDefault:"3"`
	}

	NewRelic struct {
		NewRelicLicense string `env:"NEWRELIC_LICENSE_KEY" envDefault:""`
		NewRelicAppName string `env:"NEWRELIC_APP_NAME" envDefault:""`
	}
}
