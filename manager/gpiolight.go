package manager

import (
	"github.com/nathan-osman/go-rpigpio"
)

type gpioLight struct {
	pin *rpi.Pin
}

func newGPIOLight(number int) (*gpioLight, error) {
	p, err := rpi.OpenPin(number, rpi.OUT)
	if err != nil {
		return nil, err
	}
	return &gpioLight{
		pin: p,
	}, nil
}

func (g *gpioLight) SetState(state bool) {
	if state {
		g.pin.Write(rpi.HIGH)
	} else {
		g.pin.Write(rpi.LOW)
	}
}

func (g *gpioLight) Close() {
	g.pin.Close()
}
