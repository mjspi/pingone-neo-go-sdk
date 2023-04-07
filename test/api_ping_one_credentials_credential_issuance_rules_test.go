/*
PingOne Platform API - Credentials

Testing PingOneCredentialsCredentialIssuanceRulesApiService

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

func Test_openapi_PingOneCredentialsCredentialIssuanceRulesApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService ApplyCredentialIssuanceRuleStagedChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService CreateCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.CreateCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService DeleteCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.DeleteCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService ReadAllCredentialIssuanceRules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules(context.Background(), environmentID, credentialTypeID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService ReadCredentialIssuanceRuleStagedChanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService ReadCredentialIssuanceRuleUsageCounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService ReadCredentialIssuanceRuleUsageDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService ReadOneCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PingOneCredentialsCredentialIssuanceRulesApiService UpdateCredentialIssuanceRule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var environmentID string
		var credentialTypeID string
		var credentialIssuanceRuleID string

		resp, httpRes, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.UpdateCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
