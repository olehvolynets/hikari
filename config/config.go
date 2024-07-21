package config

import (
	"io"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Events []Event
	Types  []Type
}

func LoadConfig(r io.Reader) (*Config, error) {
	var cfg Config
	dec := yaml.NewDecoder(r)

	if err := dec.Decode(&cfg); err != nil && err != io.EOF {
		return nil, err
	}

	return &cfg, nil
}
