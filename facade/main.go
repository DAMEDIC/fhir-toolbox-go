package main

import (
	"fmt"
	"log/slog"

	"fhir-toolbox/facade/config"
	"fhir-toolbox/fake"
)

var Version = "dev"

func main() {
	c := config.Load()

	slog.Info("starting FHIR facade",
		"version", Version,
		"config", c)

	f := fake.Fake{}

	fmt.Println(f.GetPatient())
	fmt.Println(f.GetObservtion())
}
