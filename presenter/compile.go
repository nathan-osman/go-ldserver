package presenter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/nathan-osman/go-ldserver/manager"
)

func compile(m manager.Manager, r io.Reader) (performance, error) {
	var (
		s = bufio.NewScanner(r)
		p = performance{}
	)
	for s.Scan() {
		var (
			start float64
			end   float64
			name  string
		)
		if _, err := fmt.Sscanf(s.Text(), "%f %f %s", &start, &end, &name); err != nil {
			return nil, err
		}
		l, ok := m[name]
		if !ok {
			return nil, errors.New("invalid light specified")
		}
		p = append(p, &event{
			light:  l,
			state:  true,
			offset: time.Duration(start * float64(time.Second)),
		})
		p = append(p, &event{
			light:  l,
			state:  false,
			offset: time.Duration(end * float64(time.Second)),
		})
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	sort.Sort(p)
	return p, nil
}
