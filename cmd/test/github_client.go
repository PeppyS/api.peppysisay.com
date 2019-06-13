package main

import (
	"fmt"
	"net/http"

	"github.com/PeppyS/api.peppysisay.com/pkg/github"
)

func main() {
	// TODO - Move to config
	github.NewClient(http.DefaultClient)

	fmt.Println("Sup")
}
