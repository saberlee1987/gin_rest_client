package config

import (
	"time"
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
	BasePath string `yaml:"base-path"`
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
	ReadTimeout        time.Duration `yaml:"readTimeout"`
	ConnectionDuration time.Duration `yaml:"connectionDuration"`
	ConnectTimeout     time.Duration `yaml:"connectTimeout"`
	ConnectionPerHost  int           `yaml:"ConnectionPerHost"`
}

type Consul struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
