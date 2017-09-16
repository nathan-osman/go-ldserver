package manager

import (
	"encoding/json"
)

// Config is used to load configuration for lights from a JSON file.
type Config struct {
	Lights map[string]*lightConfig `json:"lights"`
}

type lightConfig struct {
	Type   string          `json:"type"`
	Config json.RawMessage `json:"config"`
}

type debugConfig struct {
	Name string `json:"name"`
}

type gpioConfig struct {
	Number int `json:"number"`
}
