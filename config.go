package main

const (
	DEFAULT_PROTOCOL = "http"
	DEFAULT_HOST     = "localhost"
	DEFAULT_PORT     = 4444
)

type Config struct {
	PROTOCOL string
	HOST     string
	PORT     uint8
}

func (c *Config) New(protocol, host string, port uint8) *Config {
	return &Config{
		PROTOCOL: protocol,
		HOST:     host,
		PORT:     port,
	}
}
