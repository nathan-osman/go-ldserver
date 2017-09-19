package manager

import (
	"encoding/json"
)

// Config is used to load configuration for lights from a JSON file.
type Config struct {
	Lights map[string]json.RawMessage `json:"lights"`
}

type debugConfig struct {
	Names []string `json:"names"`
}

type gpioConfig struct {
	Pins []*gpioPinConfig `json:"pins"`
}

type gpioPinConfig struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

type wsConfig struct {
	Host  string   `json:"host"`
	Names []string `json:"names"`
}
