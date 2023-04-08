# \PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi

All URIs are relative to *https://api.pingone.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitalWalletApp**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.md#CreateDigitalWalletApp) | **Post** /environments/{environmentID}/digitalWalletApplications | Create Digital Wallet App
[**DeleteDigitalWalletApp**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.md#DeleteDigitalWalletApp) | **Delete** /environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID} | Delete Digital Wallet App
[**ReadAllDigitalWalletApps**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.md#ReadAllDigitalWalletApps) | **Get** /environments/{environmentID}/digitalWalletApplications | Read All Digital Wallet Apps
[**ReadOneDigitalWalletApp**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.md#ReadOneDigitalWalletApp) | **Get** /environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID} | Read One Digital Wallet App
[**UpdateDigitalWalletApp**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.md#UpdateDigitalWalletApp) | **Put** /environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID} | Update Digital Wallet App



## CreateDigitalWalletApp

> CredentialDigitalWalletApplication CreateDigitalWalletApp(ctx, environmentID).CredentialDigitalWalletApplication(credentialDigitalWalletApplication).Execute()

Create Digital Wallet App

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
    environmentID := "environmentID_example" // string | 
    credentialDigitalWalletApplication := *openapiclient.NewCredentialDigitalWalletApplication("AppOpenUrl_example", "Name_example") // CredentialDigitalWalletApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.CreateDigitalWalletApp(context.Background(), environmentID).CredentialDigitalWalletApplication(credentialDigitalWalletApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.CreateDigitalWalletApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDigitalWalletApp`: CredentialDigitalWalletApplication
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.CreateDigitalWalletApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **credentialDigitalWalletApplication** | [**CredentialDigitalWalletApplication**](CredentialDigitalWalletApplication.md) |  | 

### Return type

[**CredentialDigitalWalletApplication**](CredentialDigitalWalletApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDigitalWalletApp

> DeleteDigitalWalletApp(ctx, environmentID, digitalWalletApplicationID).Execute()

Delete Digital Wallet App

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
    environmentID := "environmentID_example" // string | 
    digitalWalletApplicationID := "digitalWalletApplicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.DeleteDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.DeleteDigitalWalletApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**digitalWalletApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ReadAllDigitalWalletApps

> EntityArray ReadAllDigitalWalletApps(ctx, environmentID).Execute()

Read All Digital Wallet Apps

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
    environmentID := "environmentID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadAllDigitalWalletApps(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadAllDigitalWalletApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllDigitalWalletApps`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadAllDigitalWalletApps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllDigitalWalletAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneDigitalWalletApp

> CredentialDigitalWalletApplication ReadOneDigitalWalletApp(ctx, environmentID, digitalWalletApplicationID).Execute()

Read One Digital Wallet App

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
    environmentID := "environmentID_example" // string | 
    digitalWalletApplicationID := "digitalWalletApplicationID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadOneDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadOneDigitalWalletApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneDigitalWalletApp`: CredentialDigitalWalletApplication
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.ReadOneDigitalWalletApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**digitalWalletApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CredentialDigitalWalletApplication**](CredentialDigitalWalletApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDigitalWalletApp

> CredentialDigitalWalletApplication UpdateDigitalWalletApp(ctx, environmentID, digitalWalletApplicationID).CredentialDigitalWalletApplication(credentialDigitalWalletApplication).Execute()

Update Digital Wallet App

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
    environmentID := "environmentID_example" // string | 
    digitalWalletApplicationID := "digitalWalletApplicationID_example" // string | 
    credentialDigitalWalletApplication := *openapiclient.NewCredentialDigitalWalletApplication("AppOpenUrl_example", "Name_example") // CredentialDigitalWalletApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.UpdateDigitalWalletApp(context.Background(), environmentID, digitalWalletApplicationID).CredentialDigitalWalletApplication(credentialDigitalWalletApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.UpdateDigitalWalletApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDigitalWalletApp`: CredentialDigitalWalletApplication
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi.UpdateDigitalWalletApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**digitalWalletApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **credentialDigitalWalletApplication** | [**CredentialDigitalWalletApplication**](CredentialDigitalWalletApplication.md) |  | 

### Return type

[**CredentialDigitalWalletApplication**](CredentialDigitalWalletApplication.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

