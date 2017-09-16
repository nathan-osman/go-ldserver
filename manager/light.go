package manager

import (
	"encoding/json"
	"errors"
)

// Light is the interface that each type of light must implement.
type Light interface {
	SetState(bool)
	Close()
}

func newLight(lCfg *lightConfig) (Light, error) {
	switch lCfg.Type {
	case "debug":
		c := &debugConfig{}
		if err := json.Unmarshal(lCfg.Config, c); err != nil {
			return nil, err
		}
		return newDebugLight(c.Name), nil
	case "gpio":
		c := &gpioConfig{}
		if err := json.Unmarshal(lCfg.Config, c); err != nil {
			return nil, err
		}
		return newGPIOLight(c.Number)
	default:
		return nil, errors.New("invalid light type")
	}
}
