package config

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	DbUrl string
	Port  string
	Db    *gorm.DB
}

var Conf Config

func Load() (Config, error) {

	path, _ := os.Getwd()

	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln(err)
	}

	db, err := gorm.Open(mysql.Open(viper.GetString("DB_URL")), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	result := &Config{
		DbUrl: viper.GetString("DB_URL"),
		Port:  viper.GetString("PORT"),
		Db:    db,
	}

	Conf = *result

	return *result, nil
}
