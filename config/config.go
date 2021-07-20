package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Port string `required:"true" split_words:"true"`
}

var once sync.Once
var C Config

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &C); err != nil {
			log.Errorf("Error parsing environment vars %#v", err)
		}
	})

	return C
}
