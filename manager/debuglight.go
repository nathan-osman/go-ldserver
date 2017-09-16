package manager

import (
	"github.com/sirupsen/logrus"
)

type debugLight struct {
	log *logrus.Entry
}

func newDebugLight(name string) *debugLight {
	return &debugLight{
		log: logrus.WithField("context", name),
	}
}

func (d *debugLight) SetState(state bool) {
	d.log.Infof("state: %v", state)
}

func (d *debugLight) Close() {}
