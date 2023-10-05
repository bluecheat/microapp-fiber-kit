package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	confOnce  sync.Once
	appConfig *Config
)

type Config struct {
	Host     string         `json:"host"`
	Port     string         `json:"port"`
	Develop  bool           `json:"develop"`
	Cors     CorsConfig     `json:"cors"`
	Log      LoggerConfig   `json:"log"`
	Database DatabaseConfig `json:"database"`
}

// DatabaseConfig struct
type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
	Password string `json:"password"`
	Type     string `json:"type"`
	MaxOpen  int    `json:"maxOpen"`
	MaxIdle  int    `json:"maxIdle"`
}

// RedisConfig struct
type RedisConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Db   int    `json:"db"`
}

// CorsConfig struct
type CorsConfig struct {
	Origins     []string `json:"origins"`
	Methods     []string `json:"methods"`
	Headers     []string `json:"headers"`
	Credentials bool     `json:"credentials"`
}

// LoggerConfig struct
type LoggerConfig struct {
	Level    string `json:"level"` // debug, info, error...
	Type     string `json:"type"`  // options: file, stdout
	Filename string `json:"filename"`
}

// LoadConfigFile config 파일(yaml)을 읽고 global struct 에 저장한다.
func LoadConfigFile(filename string) *Config {
	confOnce.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(filename)
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Fatal(err)
		}
		err = viper.Unmarshal(&appConfig)
		if err != nil {
			panic(err)
		}
		data, _ := json.MarshalIndent(appConfig, "", "  ")
		log.Printf("Config Load %s", data)
	})
	return appConfig
}
