package main

import (
	"database/sql"
	"fmt"
	"os"
	"tinyapps/api-base/handlers"
	"tinyapps/api-base/routing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

func main() {
	config := loadConfig()
	env := dbConnect(config)
	defer env.DB.Close()

	r := gin.Default()
	routing.Setup(env, r)
	r.Run(config.Server.Address)
}

type Config struct {
	Server struct {
		Address		string	`yaml:"address"`
	} `yaml:"server"`
	Database struct {
		Server		string	`yaml:"server"`
		Port		int		`yaml:"port"`
		Database	string	`yaml:"db"`
		Username	string	`yaml:"user"`
		Password	string	`yaml:"pass"`
	} `yaml:"database"`
}

func loadConfig() Config {
	f, ioErr := os.Open("config/main.yml")
	if ioErr != nil {
		panic(ioErr.Error())
	}
	defer f.Close()

	var c Config
	decoder := yaml.NewDecoder(f)
	parseErr := decoder.Decode(&c)
	if parseErr != nil {
		panic(parseErr.Error())
	}

	return c
}

func dbConnect(c Config) *handlers.Env {
	var db, dbErr = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.Database.Username, c.Database.Password, c.Database.Server, c.Database.Port, c.Database.Database))
	if dbErr != nil {
		panic(dbErr.Error())
	}

	return &handlers.Env{DB: db}
}
