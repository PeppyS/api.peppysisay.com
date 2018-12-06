package main

import (
	"os"

	"github.com/PeppyS/api.peppysisay.com/api"
)

func main() {
	port := ":8080"
	apiVersion := "dev"

	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	if v := os.Getenv("SOURCE_VERSION"); v != "" {
		apiVersion = v
	}

	app := api.New(apiVersion)
	app.Run(port)
}
