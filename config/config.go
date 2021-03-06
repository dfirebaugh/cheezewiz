package config

import (
	"embed"
	_ "embed"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Title    string  `yaml:"title"`
	TileSize float64 `yaml:"tile-size"`
	Window   struct {
		Height int `yaml:"height"`
		Width  int `yaml:"width"`
	} `yaml:"window"`
	Server       string `yaml:"server"`
	Port         int    `yaml:"port"`
	ScaleFactor  int    `yaml:"scale-factor"`
	DebugEnabled bool   `yaml:"debug"`
}

//go:embed levels/*
var LevelFS embed.FS

//go:embed *
var EntityFS embed.FS

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
