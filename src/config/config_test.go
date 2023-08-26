package config

import (
	"log"
	"testing"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("config.toml")
	if cfg != nil {
		log.Printf("value :%+v\n", cfg)
	}

}
