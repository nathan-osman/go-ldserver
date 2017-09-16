package presenter

import (
	"io"
	"time"

	"github.com/nathan-osman/go-ldserver/manager"
)

// Presenter runs a performance, executing each of the events in realtime.
type Presenter struct {
	perf      performance
	stopCh    chan bool
	stoppedCh chan bool
}

func (p *Presenter) run() {
	defer close(p.stoppedCh)
	start := time.Now()
	for _, e := range p.perf {
		select {
		case <-time.After(start.Add(e.offset).Sub(time.Now())):
			e.light.SetState(e.state)
		case <-p.stopCh:
			return
		}
	}
}

// NewPresenter creates a presenter for the specified reader.
func NewPresenter(m manager.Manager, r io.Reader) (*Presenter, error) {
	p, err := compile(m, r)
	if err != nil {
		return nil, err
	}
	return &Presenter{
		perf:      p,
		stopCh:    make(chan bool),
		stoppedCh: make(chan bool),
	}, nil
}

// Start begins the performance.
func (p *Presenter) Start() {
	go p.run()
}

// Stop ends the performance.
func (p *Presenter) Stop() {
	close(p.stopCh)
	<-p.stoppedCh
}
