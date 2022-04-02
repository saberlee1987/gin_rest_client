package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type Server struct {
	Port int `yaml:"port" json:"port"`
}

type Config struct {
	Server  Server  `yaml:"server" json:"server"`
	Gin     Gin     `yaml:"gin" json:"gin"`
	Service Service `yaml:"service" json:"service"`
}

type Application struct {
	Name string `yaml:"name" json:"name"`
}
type Gin struct {
	Application Application `yaml:"application" json:"application"`
	Consul      Consul      `yaml:"consul"`
}

type Service struct {
	Api          Api          `yaml:"api"`
	PersonClient PersonClient `yaml:"person-client"`
}

type Api struct {
	BasePath       string `yaml:"base-path"`
	SwaggerPath    string `yaml:"swagger-path"`
	SwaggerTitle   string `yaml:"swagger-title"`
	SwaggerVersion string `yaml:"swagger-version"`
}
type PersonClient struct {
	URL                string           `yaml:"url"`
	Port               int              `yaml:"port"`
	Baseurl            string           `yaml:"baseUrl"`
	FindAll            string           `yaml:"findAll"`
	FindByNationalCode string           `yaml:"findByNationalCode"`
	AddPerson          string           `yaml:"addPerson"`
	UpdatePerson       string           `yaml:"updatePerson"`
	DeletePerson       string           `yaml:"deletePerson"`
	Authorization      Authorization    `yaml:"authorization"`
	ConnectionConfig   ConnectionConfig `yaml:"connection"`
}

type Authorization struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type ConnectionConfig struct {
	ReadTimeout        int `yaml:"readTimeout"`
	ConnectionDuration int `yaml:"connectionDuration"`
	ConnectTimeout     int `yaml:"connectTimeout"`
	ConnectionPerHost  int `yaml:"ConnectionPerHost"`
}

type Consul struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func ReadConfigFromYamlFile() Config {
	file, err := ioutil.ReadFile("application.yml")
	if err != nil {
		log.Fatal(err)
	}
	config := Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatal(err)
	}

	serverPort := config.Server.Port
	applicationName := config.Gin.Application.Name
	fmt.Printf("Server port %d with application name is %s \n", serverPort, applicationName)

	return config
}
