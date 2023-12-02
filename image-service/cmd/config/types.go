package config

var Config = &config{}

type config struct {
	Server   Server
	Firebase Firebase
}

type Server struct {
	ServerAddr string `env:"SERVER_ADDR" envDefault:"8084"`
}

type Firebase struct {
	Bucket string `env:"FIREBASE_BUCKET" envDefault:"geek-vol10.appspot.com"`
}
