package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/soider/elevations/internal/application"
	"log"
)

func main() {
	cfg := application.Config{}
	if err := envconfig.Process("elevations", &cfg); err != nil {
		log.Fatalf("can't parse configuration %s", err)
	}
	r := application.Build(cfg)

	if err := r.Run(); err != nil {
		log.Fatalf("can't start http server %s", err)
	}

}
