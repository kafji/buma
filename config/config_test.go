package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFile(t *testing.T) {
	path := "../example.buma.toml"
	cfg, err := FromFile(path)

	if assert.Nil(t, err) {
		assert.Equal(t, Config{
			Database: Database{
				Name:     "postgres",
				User:     "postgres",
				Password: "password",
				Host:     "127.0.0.1",
				Port:     5432,
			},
			HTTPServer: HTTPServer{
				Port: 3000,
			},
			GRPCServer: GRPCServer{
				Port: 4000,
			},
		}, cfg)
	}
}
