package main

import (
	"os"

	"github.com/PeppyS/api.peppysisay.com/api"
)

func main() {
	port := ":8080"

	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	app := api.New()
	app.Run(port)
}
