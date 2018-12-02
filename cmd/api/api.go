package main

import (
	"github.com/PeppyS/api.peppysisay.com/app"
)

func main() {
	api := app.NewAPI()
	api.Run(":8080")
}
