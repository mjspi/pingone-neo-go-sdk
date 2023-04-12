/*
PingOne Platform API - Credentials

Testing PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService

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

func Test_openapi_PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService CreateDigitalWallet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var userID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.CreateDigitalWallet(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService DeleteDigitalWallet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var userID string
		var digitalWalletID string

		httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.DeleteDigitalWallet(context.Background(), environmentID, userID, digitalWalletID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService ReadAllDigitalWallets", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var userID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadAllDigitalWallets(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService ReadOneDigitalWallet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var userID string
		var digitalWalletID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadOneDigitalWallet(context.Background(), environmentID, userID, digitalWalletID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService ReadOneDigitalWalletCredentials", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var userID string
		var digitalWalletID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadOneDigitalWalletCredentials(context.Background(), environmentID, userID, digitalWalletID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApiService UpdateDigitalWallet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var environmentID string
		var userID string
		var digitalWalletID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.UpdateDigitalWallet(context.Background(), environmentID, userID, digitalWalletID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
