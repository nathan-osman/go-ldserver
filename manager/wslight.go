package manager

type wsLight struct {
	name   string
	client *wsClient
}

func newWSLights(cfg *wsConfig) ([]Light, error) {
	c, err := newWSClient(cfg.Host)
	if err != nil {
		return nil, err
	}
	lights := []Light{}
	for _, n := range cfg.Names {
		lights = append(lights, &wsLight{
			name:   n,
			client: c,
		})
	}
	if err != nil {
		c.Close()
		return nil, err
	}
	return lights, nil
}

func (w *wsLight) GetName() string {
	return w.name
}

func (w *wsLight) SetState(state bool) {
	w.client.Send(&wsState{
		Name:  w.name,
		State: state,
	})
}

func (w *wsLight) Close() {
	w.client.Release()
}
