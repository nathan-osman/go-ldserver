package manager

import (
	"errors"
)

// Light is the interface that each type of light must implement.
type Light interface {
	SetState(bool)
	Close()
}

func newLight(lCfg *lightConfig) (Light, error) {
	switch {
	case lCfg.GPIO != nil:
		return newGPIOLight(lCfg.GPIO.Number)
	default:
		return nil, errors.New("invalid light configuration")
	}
}
