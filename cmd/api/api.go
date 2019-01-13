package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/PeppyS/api.peppysisay.com/api"
	"github.com/PeppyS/api.peppysisay.com/api/blog"
	"github.com/PeppyS/api.peppysisay.com/api/blog/comments"
	"github.com/PeppyS/api.peppysisay.com/api/blog/posts"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	port := ":8080"
	apiVersion := "dev"

	godotenv.Load()
	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	config := api.SetupConfig()

	if p := os.Getenv("PORT"); p != "" {
		port = ":" + p
	}

	if v := os.Getenv("SOURCE_VERSION"); v != "" {
		apiVersion = v
	}

	creds, err := json.Marshal(map[string]interface{}{
		"type":                        config.GoogleCredentialsType,
		"project_id":                  config.GoogleProjectID,
		"private_key_id":              config.GooglePrivateKeyID,
		"private_key":                 config.GooglePrivateKey,
		"client_email":                config.GoogleClientEmail,
		"client_id":                   config.GoogleClientID,
		"auth_uri":                    config.GoogleAuthURI,
		"token_uri":                   config.GoogleTokenURI,
		"auth_provider_x509_cert_url": config.GoogleAuthProviderX509CertURL,
		"client_x509_cert_url":        config.GoogleClientX509CertURL,
	})
	if err != nil {
		log.Fatalf("JSON encode credentials: %v", err)
	}

	dbClient, err := firestore.NewClient(context.Background(), "personal-site-staging-a449f", option.WithCredentialsJSON(creds))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// TODO use lib for DI
	commentsService := comments.NewService(dbClient)
	postsService := posts.NewService(dbClient, commentsService)

	postsAPI := posts.NewAPI(postsService)

	blogAPI := blog.NewAPI(postsAPI)

	router := gin.Default()

	app := api.New(router, blogAPI, api.Opts{Version: apiVersion})
	app.Run(port)
}
