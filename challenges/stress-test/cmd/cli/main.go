package main

import (
	"log"
	"os"

	"github.com/vinizer4/go-expert-fullcycle/stress_test/internal/pkg/dependencyinjector"
)

func main() {
	di := dependencyinjector.NewDependencyInjector()

	deps, err := di.Inject()
	if err != nil {
		log.Fatalf("There was an error while injecting dependencies: %s", err)
	}

	if err := deps.CLI.Start(); err != nil {
		os.Exit(1)
	}
}
