package pingoneneo

import (
	"github.com/mjspi/pingone-neo-go-sdk/credentials"
	"github.com/mjspi/pingone-neo-go-sdk/pingoneneo/model"
)

type Config struct {
	ClientID      string
	ClientSecret  string
	EnvironmentID string
	AccessToken   string
	Region        string
}

type Client struct {
	CredentialsAPIClient *credentials.APIClient
	Region               model.RegionMapping
}
