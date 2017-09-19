package manager

import (
	"github.com/sirupsen/logrus"
)

type debugLight struct {
	name string
	log  *logrus.Entry
}

func newDebugLights(cfg *debugConfig) []Light {
	lights := []Light{}
	for _, n := range cfg.Names {
		lights = append(lights, &debugLight{
			name: n,
			log:  logrus.WithField("context", n),
		})
	}
	return lights
}

func (d *debugLight) GetName() string {
	return d.name
}

func (d *debugLight) SetState(state bool) {
	d.log.Infof("state: %v", state)
}

func (d *debugLight) Close() {}
