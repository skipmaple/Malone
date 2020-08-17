// Copyright © 2020. Drew Lee. All rights reserved.

// Package config provides config data to project.
package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

var (
	Database database
	Redis    redis
	Logger   logger
	Server   server
)

var cfgViper = func() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("./config")
	v.AddConfigPath(".")
	v.SetConfigName("config")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("Fatal error config file: %v\n", err)
	}
	return v
}()

type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string `yaml:"table_prefix, omitempty"`
}

type redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type logger struct {
	Dir string
}

type server struct {
	Port      string
	JwtSecret string
}

func init() {
	env := os.Getenv("malone_env")
	switch env {
	case "dev", "prod", "test":
		initConf(env)
	default:
		initConf("dev")
	}
}

func initConf(env string) {
	cfg := cfgViper.Sub(env)
	bindingConfig(cfg.Sub("database"), &Database)
	bindingConfig(cfg.Sub("redis"), &Redis)
	bindingConfig(cfg.Sub("logger"), &Logger)
	bindingConfig(cfg.Sub("server"), &Server)
}

func bindingConfig(cfg *viper.Viper, rawVal interface{}) {
	err := cfg.Unmarshal(&rawVal)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
