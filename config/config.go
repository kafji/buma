package config

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/BurntSushi/toml"
	"golang.org/x/exp/slog"
)

type Config struct {
	Database   Database   `toml:"database"`
	HTTPServer HTTPServer `toml:"http_server"`
	GRPCServer GRPCServer `toml:"grpc_server"`
}

type Database struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     int
}

type HTTPServer struct {
	Port int
}

type GRPCServer struct {
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
	errs := []error{}

	cfg, err := FromEnv("BUMA_CONFIG")
	if err == nil {
		return cfg
	}
	errs = append(errs, err)

	cfg, err = FromFile("./buma.toml")
	if err == nil {
		return cfg
	}
	errs = append(errs, err)

	err = errors.Join(errs...)
	slog.Error("failed to read config", err)
	panic(err)
}
