package server

import "github.com/zc2638/swag"

type Controller interface {
	GetEndpoints() []*swag.Endpoint
}

type ServerConfig struct {
	Version string
	Name    string
	Port    int
}

type ValueObject struct {
	Value string `json:"value"`
}
