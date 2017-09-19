package manager

import (
	"encoding/json"
	"errors"
)

// Light is the interface that each type of light must implement.
type Light interface {
	GetName() string
	SetState(bool)
	Close()
}

func newLights(lType string, lCfg json.RawMessage) ([]Light, error) {
	switch lType {
	case "debug":
		c := &debugConfig{}
		if err := json.Unmarshal(lCfg, c); err != nil {
			return nil, err
		}
		return newDebugLights(c), nil
	case "gpio":
		c := &gpioConfig{}
		if err := json.Unmarshal(lCfg, c); err != nil {
			return nil, err
		}
		return newGPIOLights(c)
	case "ws":
		c := &wsConfig{}
		if err := json.Unmarshal(lCfg, c); err != nil {
			return nil, err
		}
		return newWSLights(c)
	default:
		return nil, errors.New("invalid light type")
	}
}
