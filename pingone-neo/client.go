package pingone

import (
	"context"
	"fmt"
	"log"

	"github.com/mjspi/pingone-neo-go-sdk/credentials"
	"github.com/mjspi/pingone-neo-go-sdk/pingone-neo/model"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *Config) APIClient(ctx context.Context) (*Client, error) {

	token, err := getToken(ctx, c)
	if err != nil {
		return nil, err
	}

	credentialsClient, err := CredentialsAPIClient(token)
	if err != nil {
		return nil, err
	}

	apiClient := &Client{
		CredentialsAPIClient: credentialsClient,
		Region:               model.FindRegionByName(c.Region),
	}

	log.Printf("[INFO] PingOne Client configured")
	return apiClient, nil
}

func CredentialsAPIClient(token *oauth2.Token) (*credentials.APIClient, error) {

	var client *credentials.APIClient

	clientcfg := credentials.NewConfiguration()
	clientcfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	client = credentials.NewAPIClient(clientcfg)

	if client == nil {
		return nil, fmt.Errorf("Cannot initialise PingOne Agreement Management client")
	}

	log.Printf("[INFO] PingOne Agreement Management Client initialised")

	return client, nil

}

func getToken(ctx context.Context, c *Config) (*oauth2.Token, error) {

	if c.AccessToken == "" {

		if c.ClientID == "" || c.ClientSecret == "" || c.EnvironmentID == "" || c.Region == "" {
			return nil, fmt.Errorf("Required parameter missing.  Must provide ClientID, ClientSecret, EnvironmentID and Region.")
		}

		regionSuffix := model.FindRegionByName(c.Region).URLSuffix

		//Get URL from SDK
		authURL := fmt.Sprintf("https://auth.pingone.%s", regionSuffix)
		log.Printf("[INFO] Getting token from %s", authURL)

		//OAuth 2.0 config for client creds
		config := clientcredentials.Config{
			ClientID:     c.ClientID,
			ClientSecret: c.ClientSecret,
			TokenURL:     fmt.Sprintf("%s/%s/as/token", authURL, c.EnvironmentID),
			AuthStyle:    oauth2.AuthStyleAutoDetect,
		}

		token, err := config.Token(ctx)
		if err != nil {
			return nil, err
		}
		log.Printf("[INFO] Token retrieved")

		return token, nil

	} else {
		token := oauth2.Token{
			AccessToken: c.AccessToken,
			TokenType:   "Bearer",
		}
		return &token, nil
	}
}
