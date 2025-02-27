package server

import (
	"fmt"
)

type AppConfig struct {
	Port         int
	ReadTimeout  int
	WriteTimeout int
}

func (c *AppConfig) getAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}
