package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Mysql struct {
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
	DBname   string `yaml:"dbname,omitempty"`
	Host     string `yaml:"host,omitempty"`
	Port     string `yaml:"port,omitempty"`
}

type Config struct {
	Port  string `yaml:"port,omitempty"`
	Mysql *Mysql `yaml:"mysql"`
}

func NewConfig(path string) (*Config, error) {
	var config Config

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
