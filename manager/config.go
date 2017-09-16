package manager

// Config is used to load configuration for lights from a JSON file.
type Config struct {
	Lights map[string]*lightConfig `json:"lights"`
}

type lightConfig struct {
	GPIO *gpioConfig `json:"gpio"`
}

type gpioConfig struct {
	Number int `json:"number"`
}
