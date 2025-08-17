package main

import (
	"github.com/vinizer4/go-expert-fullcycle/challenges/rate-limiter/config"
	"github.com/vinizer4/go-expert-fullcycle/challenges/rate-limiter/internal/pkg/dependencyinjector"
)

func main() {
	configs, err := config.Load(".")
	if err != nil {
		panic(err)
	}

	di := dependencyinjector.NewDependencyInjector(configs)

	deps, err := di.Inject()
	if err != nil {
		panic(err)
	}

	deps.WebServer.Start()
}
