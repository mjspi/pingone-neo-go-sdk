/*
PingOne Platform API - Credentials

Testing PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService CreateDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string

		httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.CreateDigitalWalletApp(context.Background(), envID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService DeleteDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string
		var digitalWalletApplicationID string

		httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.DeleteDigitalWalletApp(context.Background(), envID, digitalWalletApplicationID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService ReadAllDigitalWalletApps", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string

		httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadAllDigitalWalletApps(context.Background(), envID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService ReadOneDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string
		var digitalWalletApplicationID string

		httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadOneDigitalWalletApp(context.Background(), envID, digitalWalletApplicationID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService UpdateDigitalWalletApp", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string
		var digitalWalletApplicationID string

		httpRes, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.UpdateDigitalWalletApp(context.Background(), envID, digitalWalletApplicationID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
