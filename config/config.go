package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Title      string  `yaml:"title"`
	SpriteSize float64 `yaml:"sprite-size"`
	Window     struct {
		Height int `yaml:"height"`
		Width  int `yaml:"width"`
	} `yaml:"window"`
	Server      string `yaml:"server"`
	Port        int    `yaml:"port"`
	ScaleFactor int    `yaml:"scale-factor"`
}

//go:embed config.yml
var ConfigRaw []byte

func (c *Config) Unmarshal(raw []byte) {
	yaml.Unmarshal(raw, c)
}

func Get() *Config {
	c := &Config{}
	c.Unmarshal(ConfigRaw)
	return c
}
