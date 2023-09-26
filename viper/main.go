package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName  string
	LogLevel string

	MySQL MySQLConfig
	Redis RedisConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type RedisConfig struct {
	IP   string
	Port int
}

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./viper/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}

	var c Config
	viper.Unmarshal(&c)

	fmt.Println(c.MySQL)
	fmt.Println(c.Redis)
}
