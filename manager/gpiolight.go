package manager

import (
	"github.com/nathan-osman/go-rpigpio"
)

type gpioLight struct {
	name string
	pin  *rpi.Pin
}

func newGPIOLights(cfg *gpioConfig) ([]Light, error) {
	var (
		lights = []Light{}
		err    = func() error {
			for _, c := range cfg.Pins {
				p, err := rpi.OpenPin(c.Number, rpi.OUT)
				if err != nil {
					return err
				}
				lights = append(lights, &gpioLight{
					name: c.Name,
					pin:  p,
				})
			}
			return nil
		}()
	)
	if err != nil {
		for _, l := range lights {
			l.Close()
		}
		return nil, err
	}
	return lights, nil
}

func (g *gpioLight) GetName() string {
	return g.name
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
