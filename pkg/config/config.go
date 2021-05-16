package config

import (
	"flag"
	"io/ioutil"
	"sync/atomic"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var confPath string
var gConf atomic.Value

func init() {
	flag.StringVar(&confPath, "conf", "configs/config.yaml", "default config path")
}

type Config struct {
	Addr  string      `yaml:"addr"`
	MySQL MySQLConfig `yaml:"mysql"`
}

type MySQLConfig struct {
	DSN         string        `yaml:"dsn"`
	MaxOpen     int           `yaml:"max_open_conn"`
	MaxIdle     int           `yaml:"max_idle_conn"`
	MaxLifetime time.Duration `yaml:"max_life_time"`
}

func NewConfig() (*Config, error) {
	configContent, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(configContent, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func Set(c *Config) {
	gConf.Store(c)
}

func Get() *Config {
	v, _ := gConf.Load().(*Config)
	return v
}
