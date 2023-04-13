# \PingOneCredentialsCredentialIssuersApi

All URIs are relative to *https://api.pingone.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCredentialIssuerProfile**](PingOneCredentialsCredentialIssuersApi.md#CreateCredentialIssuerProfile) | **Post** /environments/{environmentID}/credentialIssuerProfile | Create Credential Issuer Profile
[**ReadCredentialIssuerProfile**](PingOneCredentialsCredentialIssuersApi.md#ReadCredentialIssuerProfile) | **Get** /environments/{environmentID}/credentialIssuerProfile | Read Credential Issuer Profile
[**UpdateCredentialIssuerProfile**](PingOneCredentialsCredentialIssuersApi.md#UpdateCredentialIssuerProfile) | **Put** /environments/{environmentID}/credentialIssuerProfile | Update Credential Issuer Profile



## CreateCredentialIssuerProfile

> CredentialIssuerProfile CreateCredentialIssuerProfile(ctx, environmentID).CredentialIssuerProfile(credentialIssuerProfile).Execute()

Create Credential Issuer Profile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mjspi/pingone-neo-go-sdk/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialIssuerProfile := *openapiclient.NewCredentialIssuerProfile("Name_example") // CredentialIssuerProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialIssuersApi.CreateCredentialIssuerProfile(context.Background(), environmentID).CredentialIssuerProfile(credentialIssuerProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuersApi.CreateCredentialIssuerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentialIssuerProfile`: CredentialIssuerProfile
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialIssuersApi.CreateCredentialIssuerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialIssuerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialIssuerProfile** | [**CredentialIssuerProfile**](CredentialIssuerProfile.md) |  | 

### Return type

[**CredentialIssuerProfile**](CredentialIssuerProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuerProfile

> CredentialIssuerProfile ReadCredentialIssuerProfile(ctx, environmentID).Execute()

Read Credential Issuer Profile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mjspi/pingone-neo-go-sdk/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialIssuersApi.ReadCredentialIssuerProfile(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuersApi.ReadCredentialIssuerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCredentialIssuerProfile`: CredentialIssuerProfile
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialIssuersApi.ReadCredentialIssuerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CredentialIssuerProfile**](CredentialIssuerProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialIssuerProfile

> CredentialIssuerProfile UpdateCredentialIssuerProfile(ctx, environmentID).CredentialIssuerProfile(credentialIssuerProfile).Execute()

Update Credential Issuer Profile

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/mjspi/pingone-neo-go-sdk/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialIssuerProfile := *openapiclient.NewCredentialIssuerProfile("Name_example") // CredentialIssuerProfile |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialIssuersApi.UpdateCredentialIssuerProfile(context.Background(), environmentID).CredentialIssuerProfile(credentialIssuerProfile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialIssuersApi.UpdateCredentialIssuerProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredentialIssuerProfile`: CredentialIssuerProfile
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialIssuersApi.UpdateCredentialIssuerProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialIssuerProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialIssuerProfile** | [**CredentialIssuerProfile**](CredentialIssuerProfile.md) |  | 

### Return type

[**CredentialIssuerProfile**](CredentialIssuerProfile.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

