package manager

import (
	"fmt"
)

type debugLight struct {
	Name string `json:"name"`
}

func (d *debugLight) SetState(state bool) {
	fmt.Printf("[%s] state: %v\n", d.Name, state)
}

func (d *debugLight) Close() {}
