/*
PingOne Platform API - Credentials

Testing PingOneCredentialsUserCredentialsApiService

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

func Test_openapi_PingOneCredentialsUserCredentialsApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingOneCredentialsUserCredentialsApiService CreateAUserCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string

		resp, httpRes, err := apiClient.PingOneCredentialsUserCredentialsApi.CreateAUserCredential(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsUserCredentialsApiService ReadAllUserCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string

		resp, httpRes, err := apiClient.PingOneCredentialsUserCredentialsApi.ReadAllUserCredentials(context.Background(), environmentID, userID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsUserCredentialsApiService ReadOneUserCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var credentialID string

		resp, httpRes, err := apiClient.PingOneCredentialsUserCredentialsApi.ReadOneUserCredential(context.Background(), environmentID, userID, credentialID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsUserCredentialsApiService ReadOneUserCredentialWallets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var credentialID string

		resp, httpRes, err := apiClient.PingOneCredentialsUserCredentialsApi.ReadOneUserCredentialWallets(context.Background(), environmentID, userID, credentialID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsUserCredentialsApiService UpdateAUserCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var userID string
		var credentialID string

		resp, httpRes, err := apiClient.PingOneCredentialsUserCredentialsApi.UpdateAUserCredential(context.Background(), environmentID, userID, credentialID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
