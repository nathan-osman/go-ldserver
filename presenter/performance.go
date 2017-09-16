package presenter

import (
	"time"

	"github.com/nathan-osman/go-ldserver/manager"
)

type performance []*event

func (p performance) Len() int {
	return len(p)
}

func (p performance) Less(i, j int) bool {
	return p[i].offset < p[j].offset
}

func (p performance) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type event struct {
	light  manager.Light
	state  bool
	offset time.Duration
}
