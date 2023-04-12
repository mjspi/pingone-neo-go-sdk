package pingone

import (
	"github.com/mjspi/pingone-neo-go-sdk-v2/credentials"
	"github.com/mjspi/pingone-neo-go-sdk/pingone-neo/model"
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
