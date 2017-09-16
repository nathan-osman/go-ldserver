package main

import (
	"encoding/json"
	"os"

	"github.com/nathan-osman/go-ldserver/manager"
)

type Config struct {
	Manager *manager.Config `json:"manager"`
}

func LoadConfig(name string) (*Config, error) {
	r, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	c := &Config{}
	if err := json.NewDecoder(r).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}
