# \PingOneCredentialsCredentialVerificationsApi

All URIs are relative to *https://api.pingone.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialVerificationPresentationSessionNative**](PingOneCredentialsCredentialVerificationsApi.md#CreateCredentialVerificationPresentationSessionNative) | **Post** /environments/{envID}/presentationSessions | Create Credential Verification Presentation Session (NATIVE)
[**ReadCredentialVerificationCredentialData**](PingOneCredentialsCredentialVerificationsApi.md#ReadCredentialVerificationCredentialData) | **Get** /environments/{envID}/presentationSessions/{credentialsVerificationID}/credentialData | Read Credential Verification Credential Data
[**ReadOneCredentialVerificationStatus**](PingOneCredentialsCredentialVerificationsApi.md#ReadOneCredentialVerificationStatus) | **Get** /environments/{envID}/presentationSessions/{credentialsVerificationID} | Read One Credential Verification Status



## CreateCredentialVerificationPresentationSessionNative

> CreateCredentialVerificationPresentationSessionNative(ctx, envID).Authorization(authorization).Body(body).Execute()

Create Credential Verification Presentation Session (NATIVE)



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
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialVerificationsApi.CreateCredentialVerificationPresentationSessionNative(context.Background(), envID).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialVerificationsApi.CreateCredentialVerificationPresentationSessionNative``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialVerificationPresentationSessionNativeRequest struct via the builder pattern


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


## ReadCredentialVerificationCredentialData

> ReadCredentialVerificationCredentialData(ctx, envID, credentialsVerificationID).Authorization(authorization).Execute()

Read Credential Verification Credential Data



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
    credentialsVerificationID := "credentialsVerificationID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialVerificationsApi.ReadCredentialVerificationCredentialData(context.Background(), envID, credentialsVerificationID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialVerificationsApi.ReadCredentialVerificationCredentialData``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialsVerificationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialVerificationCredentialDataRequest struct via the builder pattern


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


## ReadOneCredentialVerificationStatus

> ReadOneCredentialVerificationStatus(ctx, envID, credentialsVerificationID).Authorization(authorization).Execute()

Read One Credential Verification Status



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
    credentialsVerificationID := "credentialsVerificationID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialVerificationsApi.ReadOneCredentialVerificationStatus(context.Background(), envID, credentialsVerificationID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialVerificationsApi.ReadOneCredentialVerificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**credentialsVerificationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneCredentialVerificationStatusRequest struct via the builder pattern


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

