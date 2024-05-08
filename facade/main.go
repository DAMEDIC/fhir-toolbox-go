package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"fhir-toolbox/backend/fake"
	"fhir-toolbox/facade/config"
	dispatchR4 "fhir-toolbox/server/dispatch/gen/r4"
)

var Version = "dev"

func main() {
	configFile := flag.String("f", "", "config file path")
	flag.Parse()

	c, err := config.Load(*configFile)
	if err != nil {
		slog.Error("error loading config", "error", err)
		os.Exit(1)
	}

	slog.Info("starting FHIR facade",
		"version", Version,
		"config", c)

	client := fake.Fake{}

	ctx := context.TODO()

	fmt.Println(dispatchR4.Read(ctx, &client, "Patient", "1234"))
	fmt.Println(dispatchR4.Read(ctx, &client, "Observation", "1234"))
}
