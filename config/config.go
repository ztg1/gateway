package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

const (
	LogPath = "/data/goapp/log/"
)

type Config struct {
	Debug      bool             `json:"debug"`
	Redis      RedisConfig      `json:"redis"`
	HttpServer HttpServerConfig `json:"httpServer"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

//web端请求地址
type HttpServerConfig struct {
	Addr string `json:"addr"`
}

var config Config
var configOnce sync.Once

func GetConfig() *Config {
	configOnce.Do(func() {
		bytes, err := ioutil.ReadFile("etc/config.json")
		if err != nil {
			log.Panic("config.json-err", err.Error())
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			log.Panic("Unmarshal-err", err.Error())
		}
	})
	return &config
}
