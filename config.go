package main

import (
	"fmt"
	pg_conf "github.com/joegasewicz/pg-conf"
	"log"
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
	PROTOCOL           string
	HOST               string
	PORT               int
	ALLOWED_FILE_TYPES string
	*pg_conf.PostgresConfig
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
	c := &Config{
		PROTOCOL:           envProtocol,
		HOST:               envHost,
		PORT:               envPort,
		ALLOWED_FILE_TYPES: ALLOWED_FILE_TYPES, // TODO
		PostgresConfig: &pg_conf.PostgresConfig{
			PGPort:     "", // TODO https://github.com/joegasewicz/pg-conf/issues/1
			PGDatabase: "",
			PGUser:     "",
			PGPassword: "",
		},
	}
	err := pg_conf.Update(c.PostgresConfig)
	if err != nil {
		log.Panic(err)
	}
	return c
}

func (c *Config) GetUrl() string {
	return fmt.Sprintf("%s://%s:%d", c.PROTOCOL, c.HOST, c.PORT)
}
