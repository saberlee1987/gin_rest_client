package dto

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Port int `yaml:"port" json:"port"`
}

type Config struct {
	Server Server `yaml:"server" json:"server"`
	Gin    Gin    `yaml:"gin" json:"gin"`
}

type Application struct {
	Name string `yaml:"name" json:"name"`
}
type Gin struct {
	Application Application `yaml:"application" json:"application"`
}

func (c Config) String() string {
	marshal, err := json.Marshal(c)

	if err != nil {
		return fmt.Sprintf("{\"server\":%s,\"gin\":%s}", c.Server, c.Gin)
	}
	return string(marshal)
}

func (server Server) String() string {
	marshal, err := json.Marshal(server)

	if err != nil {
		return fmt.Sprintf("{\"port\":%d}", server.Port)
	}
	return string(marshal)
}

func (application Application) String() string {
	marshal, err := json.Marshal(application)

	if err != nil {
		return fmt.Sprintf("{\"name\":\"%s\"}", application.Name)
	}
	return string(marshal)
}

func (gin Gin) String() string {
	marshal, err := json.Marshal(gin)

	if err != nil {
		return fmt.Sprintf("{\"application\":%s}", gin.Application)
	}
	return string(marshal)
}
