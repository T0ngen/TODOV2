package config

import (
	
	"log"

	"github.com/spf13/viper"
)


type Config struct {
	Port      string
	DbPath string
	Password string
}

func (c *Config) InitConfig() (err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./pkg/common/env")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
		return err
	}
	
	
	c.Port = viper.GetString("TODO_PORT")
	c.DbPath = viper.GetString("TODO_DBFILE")
	c.Password = viper.GetString("TODO_PASSWORD")
	return nil

}
