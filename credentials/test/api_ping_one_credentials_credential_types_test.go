/*
PingOne Platform API - Credentials

Testing PingOneCredentialsCredentialTypesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package credentials

import (
	"context"
	"testing"

	openapiclient "github.com/mjspi/pingone-neo-go-sdk/credentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/* talk to patrick
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
*/

func Test_openapi_PingOneCredentialsCredentialTypesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingOneCredentialsCredentialTypesApiService CreateCredentialType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.CreateCredentialType(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService DeleteACredentialType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string

		httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.DeleteACredentialType(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService ReadAllCredentialTypes", func(t *testing.T) {

		// t.Skip("skip test") // remove to run test

		/* var ctx = context.Background()
		config := &Config{
			ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
			ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
			EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
			Region:        os.Getenv("PINGONE_REGION"),
		}

		client, err := config.APIClient(ctx) */

		var environmentID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.ReadAllCredentialTypes(context.Background(), environmentID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService ReadOneCredentialType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.ReadOneCredentialType(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService UpdateACredentialType", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var credentialTypeID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.UpdateACredentialType(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
