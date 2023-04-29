package main

import (
	"os"
	"strconv"
)

const (
	DEFAULT_PROTOCOL = "http"
	DEFAULT_HOST     = "localhost"
	DEFAULT_PORT     = 4444
)

type Config struct {
	PROTOCOL string
	HOST     string
	PORT     int
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
		PROTOCOL: envProtocol,
		HOST:     envHost,
		PORT:     envPort,
	}
}
