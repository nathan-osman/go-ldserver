package manager

// Manager maintains a map of light names to lights.
type Manager map[string]Light

// TODO: check for duplicate names

// NewManager creates and initializes lights from the specified configuration.
func NewManager(cfg *Config) (Manager, error) {
	var (
		m   = Manager{}
		err = func() error {
			for lType, lCfg := range cfg.Lights {
				lights, err := newLights(lType, lCfg)
				if err != nil {
					return err
				}
				for _, l := range lights {
					m[l.GetName()] = l
				}
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

// Reset sets all lights to the unlit state.
func (m Manager) Reset() {
	for _, l := range m {
		l.SetState(false)
	}
}

// Close frees all of the resources used by the lights.
func (m Manager) Close() {
	for _, l := range m {
		l.Close()
	}
}
