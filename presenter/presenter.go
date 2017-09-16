package presenter

import (
	"os"

	"github.com/nathan-osman/go-ldserver/manager"
)

// Presenter runs a performance, executing each of the events in realtime.
type Presenter struct {
	perf performance
}

func (p *Presenter) run() {
	//...
}

// NewPresenter creates a presenter for the specified file.
func NewPresenter(m manager.Manager, filename string) (*Presenter, error) {
	r, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	perf, err := compile(m, r)
	if err != nil {
		return nil, err
	}
	p := &Presenter{
		perf: perf,
	}
	return p, nil
}
