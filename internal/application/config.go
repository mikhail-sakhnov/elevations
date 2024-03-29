package application

// Config general config
type Config struct {
	ListenOn    string `envconfig:"LISTEN_ON" default:":8080"`
	MapboxToken string `envconfig:"MAPBOX_TOKEN"`
}
