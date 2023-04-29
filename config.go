package main

import (
	"fmt"
	pg_conf "github.com/joegasewicz/pg-conf"
	"os"
	"strconv"
)

const (
	DEFAULT_PROTOCOL   = "http"
	DEFAULT_HOST       = "localhost"
	DEFAULT_PORT       = 4444
	ALLOWED_FILE_TYPES = "jpg,jpeg,png"
)

type Config struct {
	*pg_conf.PostgresConfig
	PROTOCOL           string
	HOST               string
	PORT               int
	ALLOWED_FILE_TYPES string
}

func NewConfig() *Config {
	envProtocol := os.Getenv("DEFAULT_PROTOCOL")
	if envProtocol == "" {
		envProtocol = DEFAULT_PROTOCOL
	}
	envHost := os.Getenv("DEFAULT_HOST")
	if envHost == "" {
		envHost = DEFAULT_HOST
	}
	envPort, _ := strconv.Atoi(os.Getenv("DEFAULT_PORT"))
	if envPort == 0 {
		envPort = DEFAULT_PORT
	}
	return &Config{
		PROTOCOL:           envProtocol,
		HOST:               envHost,
		PORT:               envPort,
		ALLOWED_FILE_TYPES: ALLOWED_FILE_TYPES, // TODO
	}
}

func (c *Config) GetUrl() string {
	return fmt.Sprintf("%s://%s:%d", c.PROTOCOL, c.HOST, c.PORT)
}
