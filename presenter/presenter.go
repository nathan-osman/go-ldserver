package presenter

import (
	"io"
	"sync"
	"time"

	"github.com/nathan-osman/go-ldserver/manager"
)

// Presenter runs a performance, executing each of the events in realtime.
type Presenter struct {
	mutex     sync.Mutex
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
		perf: p,
	}, nil
}

// Start begins the performance.
func (p *Presenter) Start() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.stoppedCh == nil {
		p.stopCh = make(chan bool)
		p.stoppedCh = make(chan bool)
		go p.run()
	}
}

// Stop ends the performance.
func (p *Presenter) Stop() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if p.stopCh != nil {
		close(p.stopCh)
		<-p.stoppedCh
		p.stopCh = nil
		p.stoppedCh = nil
	}
}
