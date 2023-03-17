package config

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Database Database `toml:"database"`
	Server   Server   `toml:"server"`
}

type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

type Server struct {
	Port int
}

func FromFile(path string) (cfg Config, err error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return
	}

	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		return
	}

	return
}

func FromEnv(key string) (cfg Config, err error) {
	b64 := os.Getenv(key)
	if b64 == "" {
		err = errors.New("config: variable is empty")
		return
	}

	buf, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return
	}

	err = toml.Unmarshal(buf, &cfg)
	if err != nil {
		return
	}

	return
}

func ReadConfig() Config {
	cfg, err := FromEnv("BUMA_CONFIG")
	if err == nil {
		return cfg
	}

	log.Info().Err(err).Msg("config: read config from env var fail, reading from local file")

	cfg, err = FromFile("./buma.toml")
	if err == nil {
		return cfg
	}

	log.Panic().Err(err).Msg("config: read config from local file is also fail")

	panic("unreachable")
}
