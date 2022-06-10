package config

import (
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/jinzhu/gorm/dialects/postgres" //driver for postgres
	log "github.com/sirupsen/logrus"
)

//Config struct
type Config struct {
	Address  string `envconfig:"SERVER_HOST"`
	Port     string `envconfig:"SERVER_PORT"`
	Env      string `envconfig:"ENV"`
	Database DbConfig
}

//DefaultConfig returns default config
//used if no config file is found
func DefaultConfig() *Config {
	return &Config{
		Address:  "0.0.0.0",
		Port:     os.Getenv("PORT"),
		Env:      "dev",
		Database: HerokuDbConfig(),
	}
}

//DbConfig represents data useful to db connection
type DbConfig struct {
	Host     string `envconfig:"PSQL_HOST"`
	Port     string `envconfig:"PSQL_PORT"`
	User     string `envconfig:"PSQL_USER"`
	Password string `envconfig:"PSQL_PSW"`
	Name     string `envconfig:"PSQL_DATABASE"`
}

//Dialect returns mocked postgres dialect
func (c DbConfig) Dialect() string {
	return "postgres"
}

//ConnectionInfo returns string to connect to postgres db
func (c DbConfig) ConnectionInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require", c.Host, c.Port, c.User, c.Password, c.Name)
	//sslmode=require for heroku postgres instance
}

//DefaultDbConfig returns default config for connecting to postgres
//func DefaultDbConfig() DbConfig { //local db
//	return DbConfig{
//		Host:     "localhost",0
//		Port:     "5432",
//		User:     "postgres",
//		Password: "postgres",
//		Name:     "postgres",
//	}
//}

func HerokuDbConfig() DbConfig { // "production"
	var dbConfig DbConfig
	dbConfigFile, err := os.Open("dbConfigFile.json")
	defer func(configFile *os.File) {
		err := dbConfigFile.Close()
		if err != nil {
			log.Fatalf("could not decode json db config file %s\n", err.Error())
		}
	}(dbConfigFile)
	jsonParser := json.NewDecoder(dbConfigFile)
	err = jsonParser.Decode(&dbConfig)
	if err != nil {
		log.Fatalf("could not parse json db config file %s\n", err.Error())
	}
	return dbConfig
}

//Load returns config based on .config file if exists, else use default config
func Load() *Config {
	return DefaultConfig()
}

//SetLogConfig initialization of logrus
func SetLogConfig() {

	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:            true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05",
		PadLevelText:           true,
		DisableLevelTruncation: false,
	})
	log.SetReportCaller(true)

	log.Info("Setup Logging")
}
