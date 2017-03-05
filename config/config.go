package config

import "github.com/jinzhu/configor"

var Config = struct {
	Port uint `default:"9000" env:"PORT"`
	DB   struct {
		Name string `default:"data/fw.db"`
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml"); err != nil {
		panic(err)
	}
}
