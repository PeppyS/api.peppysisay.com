package main

import (
	"github.com/PeppyS/api.peppysisay.com/api"
)

func main() {
	app := api.New()
	app.Run(":8080")
}
