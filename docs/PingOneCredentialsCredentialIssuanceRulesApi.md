# \PingOneCredentialsCredentialIssuanceRulesApi

All URIs are relative to *https://api.pingone.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCredentialIssuanceRuleStagedChanges**](PingOneCredentialsCredentialIssuanceRulesApi.md#ApplyCredentialIssuanceRuleStagedChanges) | **Post** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/stagedChanges | Apply Credential Issuance Rule Staged Changes
[**CreateCredentialIssuanceRule**](PingOneCredentialsCredentialIssuanceRulesApi.md#CreateCredentialIssuanceRule) | **Post** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules | Create Credential Issuance Rule
[**DeleteCredentialIssuanceRule**](PingOneCredentialsCredentialIssuanceRulesApi.md#DeleteCredentialIssuanceRule) | **Delete** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID} | Delete Credential Issuance Rule
[**ReadAllCredentialIssuanceRules**](PingOneCredentialsCredentialIssuanceRulesApi.md#ReadAllCredentialIssuanceRules) | **Get** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules | Read All Credential Issuance Rules
[**ReadCredentialIssuanceRuleStagedChanges**](PingOneCredentialsCredentialIssuanceRulesApi.md#ReadCredentialIssuanceRuleStagedChanges) | **Get** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/stagedChanges | Read Credential Issuance Rule Staged Changes
[**ReadCredentialIssuanceRuleUsageCounts**](PingOneCredentialsCredentialIssuanceRulesApi.md#ReadCredentialIssuanceRuleUsageCounts) | **Get** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/usageCounts | Read Credential Issuance Rule Usage Counts
[**ReadCredentialIssuanceRuleUsageDetails**](PingOneCredentialsCredentialIssuanceRulesApi.md#ReadCredentialIssuanceRuleUsageDetails) | **Get** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/usageDetails | Read Credential Issuance Rule Usage Details
[**ReadOneCredentialIssuanceRule**](PingOneCredentialsCredentialIssuanceRulesApi.md#ReadOneCredentialIssuanceRule) | **Get** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID} | Read One Credential Issuance Rule
[**UpdateCredentialIssuanceRule**](PingOneCredentialsCredentialIssuanceRulesApi.md#UpdateCredentialIssuanceRule) | **Put** /environments/{envID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID} | Update Credential Issuance Rule



## ApplyCredentialIssuanceRuleStagedChanges

> ApplyCredentialIssuanceRuleStagedChanges(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).ContentType(contentType).Body(body).Execute()

Apply Credential Issuance Rule Staged Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    contentType := "application/vnd.pingidentity.credentials.applyStagedChanges+json" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).ContentType(contentType).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyCredentialIssuanceRuleStagedChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 
 **contentType** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredentialIssuanceRule

> CreateCredentialIssuanceRule(ctx, envID, credentialTypeID).Authorization(authorization).Body(body).Execute()

Create Credential Issuance Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.CreateCredentialIssuanceRule(context.Background(), envID, credentialTypeID).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.CreateCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialIssuanceRule

> DeleteCredentialIssuanceRule(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()

Delete Credential Issuance Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.DeleteCredentialIssuanceRule(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.DeleteCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllCredentialIssuanceRules

> ReadAllCredentialIssuanceRules(ctx, envID, credentialTypeID).Authorization(authorization).Execute()

Read All Credential Issuance Rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules(context.Background(), envID, credentialTypeID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllCredentialIssuanceRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuanceRuleStagedChanges

> ReadCredentialIssuanceRuleStagedChanges(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()

Read Credential Issuance Rule Staged Changes



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuanceRuleStagedChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuanceRuleUsageCounts

> ReadCredentialIssuanceRuleUsageCounts(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()

Read Credential Issuance Rule Usage Counts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuanceRuleUsageCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuanceRuleUsageDetails

> ReadCredentialIssuanceRuleUsageDetails(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()

Read Credential Issuance Rule Usage Details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuanceRuleUsageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneCredentialIssuanceRule

> ReadOneCredentialIssuanceRule(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()

Read One Credential Issuance Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialIssuanceRule

> UpdateCredentialIssuanceRule(ctx, envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Body(body).Execute()

Update Credential Issuance Rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    envID := "envID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialIssuanceRulesApi.UpdateCredentialIssuanceRule(context.Background(), envID, credentialTypeID, credentialIssuanceRuleID).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuanceRulesApi.UpdateCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** |  | 
 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

