package main

import (
	"context"
	"fmt"
	"log/slog"

	"fhir-toolbox/backend/fake"
	"fhir-toolbox/facade/config"
	dispatchR4 "fhir-toolbox/server/dispatch/gen/r4"
)

var Version = "dev"

func main() {
	c := config.Load()

	slog.Info("starting FHIR facade",
		"version", Version,
		"config", c)

	client := fake.Fake{}

	ctx := context.TODO()

	fmt.Println(dispatchR4.Read(ctx, &client, "Patient", "1234"))
	fmt.Println(dispatchR4.Read(ctx, &client, "Observation", "1234"))
}
