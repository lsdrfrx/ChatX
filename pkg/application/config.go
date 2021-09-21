package application

import "os"

type Config struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{
		Addr: os.Getenv("ADDR"),
	}
}