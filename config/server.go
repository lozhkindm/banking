package config

import (
	"fmt"
	"os"
)

type Server struct {
	Address string
	Port    string
}

func NewServerConfig() Server {
	conf := Server{
		Address: os.Getenv("SERVER_ADDRESS"),
		Port:    os.Getenv("SERVER_PORT"),
	}

	if conf.Address == "" {
		conf.Address = "localhost"
	}

	if conf.Port == "" {
		conf.Port = "8000"
	}

	return conf
}

func (server Server) AsString() string {
	return fmt.Sprintf("%s:%s", server.Address, server.Port)
}
