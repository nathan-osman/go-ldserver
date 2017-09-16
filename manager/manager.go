package manager

import (
	"errors"
)

// Manager maintains a map of light names to lights.
type Manager map[string]Light

// NewManager creates and initializes lights from the specified configuration.
func NewManager(cfg *Config) (Manager, error) {
	var (
		m   = Manager{}
		err = func() error {
			for name, lCfg := range cfg.Lights {
				if _, ok := m[name]; ok {
					return errors.New("duplicate light name")
				}
				l, err := newLight(lCfg)
				if err != nil {
					return err
				}
				m[name] = l
			}
			return nil
		}()
	)
	if err != nil {
		m.Close()
		return nil, err
	}
	return m, nil
}

// Close frees all of the resources used by the lights.
func (m Manager) Close() {
	for _, l := range m {
		l.Close()
	}
}
