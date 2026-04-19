package core_http_server

import (
	"time"
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr 			string 		  `envconfig:"ADDR" 			 required:"true"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT"  default:"30s"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("HTTP", &config); err != nil {
		return Config{}, fmt.Errorf("process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get HTTP server: %w", err)
		panic(err)
	}

	return config
}