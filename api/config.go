package api

import (
	"os"
	"strings"
)

type Config struct {
	Port                          string
	GoogleCredentialsType         string
	GoogleProjectID               string
	GooglePrivateKeyID            string
	GooglePrivateKey              string
	GoogleClientEmail             string
	GoogleClientID                string
	GoogleAuthURI                 string
	GoogleTokenURI                string
	GoogleAuthProviderX509CertURL string
	GoogleClientX509CertURL       string
}

func SetupConfig() Config {
	return Config{
		Port:                          os.Getenv("PORT"),
		GoogleCredentialsType:         os.Getenv("GOOGLE_CREDENTIALS_TYPE"),
		GoogleProjectID:               os.Getenv("GOOGLE_PROJECT_ID"),
		GooglePrivateKeyID:            os.Getenv("GOOGLE_PRIVATE_KEY_ID"),
		GooglePrivateKey:              strings.Replace(os.Getenv("GOOGLE_PRIVATE_KEY"), `\n`, "\n", -1),
		GoogleClientEmail:             os.Getenv("GOOGLE_CLIENT_EMAIL"),
		GoogleClientID:                os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleAuthURI:                 os.Getenv("GOOGLE_AUTH_URI"),
		GoogleTokenURI:                os.Getenv("GOOGLE_TOKEN_URI"),
		GoogleAuthProviderX509CertURL: os.Getenv("GOOGLE_AUTH_PROVIDER_X509_CERT_URL"),
		GoogleClientX509CertURL:       os.Getenv("GOOGLE_CLIENT_X509_CERT_URL"),
	}
}
