package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"fhir-toolbox/backend/fake"
	"fhir-toolbox/model"
	"fhir-toolbox/rest"
)

func main() {
	var backendUrl = os.Args[1]

	slog.SetLogLoggerLevel(slog.LevelDebug)
	log.Printf("FHIR Server: %s", backendUrl)

	backend := fake.Fake{}
	log.Fatal(http.ListenAndServe(":80", rest.NewServer[model.R4](&backend, rest.DefaultConfig)))
}
