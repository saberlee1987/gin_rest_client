package consul

import (
	"fmt"
	consul "github.com/hashicorp/consul/api"
)

type Client interface {
	// Get a Service from consul
	Service(string, string) ([]string, error)
	// Register a service with local agent
	Register(string, int) error
	// Deregister a service with local agent
	DeRegister(string) error
}

type client struct {
	consul *consul.Client
}

//NewConsul returns a Client interface for given consul address
func NewConsulClient(host string, port int) (*client, error) {
	config := consul.DefaultConfig()
	address := fmt.Sprintf("%s:%d", host, port)
	config.Address = address
	c, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &client{consul: c}, nil
}

// Register a service with consul local agent
func (c *client) Register(name string, port int) error {
	reg := &consul.AgentServiceRegistration{
		ID:   name,
		Name: name,
		Port: port,
	}
	return c.consul.Agent().ServiceRegister(reg)
}

// DeRegister a service with consul local agent
func (c *client) DeRegister(id string) error {
	return c.consul.Agent().ServiceDeregister(id)
}

// Service return a service
func (c *client) Service(service, tag string) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	passingOnly := true
	address, meta, err := c.consul.Health().Service(service, tag, passingOnly, nil)
	if len(address) == 0 && err == nil {
		return nil, nil, fmt.Errorf("service ( %s ) was not found", service)
	}
	if err != nil {
		return nil, nil, err
	}
	return address, meta, nil
}
