/*
PingOne Platform API - Credentials

Testing PingOneCredentialsCredentialTypesApiService

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

func Test_openapi_PingOneCredentialsCredentialTypesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingOneCredentialsCredentialTypesApiService CreateCredentialType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string

		httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.CreateCredentialType(context.Background(), envID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService DeleteACredentialType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string
		var credentialTypeID string

		httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.DeleteACredentialType(context.Background(), envID, credentialTypeID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService ReadAllCredentialTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string

		httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.ReadAllCredentialTypes(context.Background(), envID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService ReadOneCredentialType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string
		var credentialTypeID string

		httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.ReadOneCredentialType(context.Background(), envID, credentialTypeID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialTypesApiService UpdateACredentialType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var envID string
		var credentialTypeID string

		httpRes, err := apiClient.PingOneCredentialsCredentialTypesApi.UpdateACredentialType(context.Background(), envID, credentialTypeID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
