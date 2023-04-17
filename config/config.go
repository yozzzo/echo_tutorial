package config

import (
	"gopkg.in/ini.v1"
	"log"
)

type ConfigList struct {
	SqlDriver string
	DBName    string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("Fail to read file: %v", err)
	}
	Config = ConfigList{
		SqlDriver: cfg.Section("database").Key("driver").String(),
		DBName:    cfg.Section("database").Key("name").String(),
		LogFile:   cfg.Section("log").Key("logFile").String(),
	}
}
