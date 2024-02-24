package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Addr string `yaml:"addr"`
}

func (c *Config) LogAll() {
	s := reflect.ValueOf(c).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		log.Printf("%s %s = %v\n", typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}

func New(configPath string) (*Config, error) {
	var err error

	once.Do(func() {
		var configRaw []byte
		configRaw, err = os.ReadFile(configPath)
		if err != nil {
			return
		}

		if err = yaml.Unmarshal(configRaw, &cfg); err != nil {
			return
		}
	})
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
