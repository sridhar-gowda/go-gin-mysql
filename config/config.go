package config

import (
	"fmt"
	"github.com/sridhar-gowda/go-gin-mysql/helper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	//GetConf is variable which will contain all configurations
	GetConf *Config
	//PropFilePath is variable which will contain path of properties file
	defaultPropFilePath = "./config.yaml"
	EnvVar = "CONFIG_FILE"

)

func init() {
	GetConf = LoadConfig()
}

type Config struct {
	ServicePort int    `yaml:"service_port"`
	DbPort      int    `yaml:"db_port"`
	DbHost      string `yaml:"db_host"`
	DbUsername  string `yaml:"db_username"`
	DbPassword  string `yaml:"db_password"`
	DbDbname    string `yaml:"db_dbname"`
}

func LoadConfig() *Config {
	confObject := &Config{}
	filePath := helper.GetEnv(EnvVar, defaultPropFilePath)
	configYaml, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error while loading file")
	}
	err = yaml.Unmarshal(configYaml, confObject)
	if err != nil {
		fmt.Println("Error while UnMarshaling file")
	}
	return confObject
}
