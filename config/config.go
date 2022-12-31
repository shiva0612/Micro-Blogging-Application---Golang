package config

import (
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

var (
	Cfg *Config
)

type Config struct {
	Mongo struct {
		Port     string `json:"port"`
		Username string `json:"username"`
		Password string `json:"password"`
		Authdb   string `json:"authdb"`
		Dbname   string `json:"dbname"`
	}
	App struct {
		Port       string        `json:"port"`
		Jwt_secret string        `json:"jwt_secret"`
		Token_time time.Duration `json:"token_time"`
	}
}

func (c *Config) LoadConfig(path string) {
	viper.AddConfigPath(filepath.Dir(path))     //dir path
	viper.SetConfigName(filepath.Base(path))    //filename
	viper.SetConfigType(filepath.Ext(path)[1:]) //.json	err := viper.ReadInConfig()

	err := viper.ReadInConfig()
	if err != nil {
		panic("error reading config: " + err.Error())
	}

	err = viper.Unmarshal(c)
	if err != nil {
		panic("error while unmarshalling config: " + err.Error())
	}

}

func NewConfig(path string) {
	Cfg = &Config{}
	Cfg.LoadConfig(path)
}
