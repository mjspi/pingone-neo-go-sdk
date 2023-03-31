# \PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi

All URIs are relative to *https://api.pingone.*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDigitalWallet**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.md#CreateDigitalWallet) | **Post** /environments/{envID}/users/{userID}/digitalWallets | Create Digital Wallet
[**DeleteDigitalWallet**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.md#DeleteDigitalWallet) | **Delete** /environments/{envID}/users/{userID}/digitalWallets/{digitalWalletID} | Delete Digital Wallet
[**ReadAllDigitalWallets**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.md#ReadAllDigitalWallets) | **Get** /environments/{envID}/users/{userID}/digitalWallets | Read All Digital Wallets
[**ReadOneDigitalWallet**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.md#ReadOneDigitalWallet) | **Get** /environments/{envID}/users/{userID}/digitalWallets/{digitalWalletID} | Read One Digital Wallet
[**ReadOneDigitalWalletCredentials**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.md#ReadOneDigitalWalletCredentials) | **Get** /environments/{envID}/users/{userID}/digitalWallets/{digitalWalletID}/provisionedCredentials | Read One Digital Wallet Credentials
[**UpdateDigitalWallet**](PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.md#UpdateDigitalWallet) | **Put** /environments/{envID}/users/{userID}/digitalWallets/{digitalWalletID} | Update Digital Wallet



## CreateDigitalWallet

> CreateDigitalWallet(ctx, envID, userID).Authorization(authorization).Body(body).Execute()

Create Digital Wallet



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
    userID := "userID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.CreateDigitalWallet(context.Background(), envID, userID).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.CreateDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDigitalWalletRequest struct via the builder pattern


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


## DeleteDigitalWallet

> DeleteDigitalWallet(ctx, envID, userID, digitalWalletID).Authorization(authorization).Execute()

Delete Digital Wallet



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
    userID := "userID_example" // string | 
    digitalWalletID := "digitalWalletID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.DeleteDigitalWallet(context.Background(), envID, userID, digitalWalletID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.DeleteDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDigitalWalletRequest struct via the builder pattern


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


## ReadAllDigitalWallets

> ReadAllDigitalWallets(ctx, envID, userID).Authorization(authorization).Execute()

Read All Digital Wallets



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
    userID := "userID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadAllDigitalWallets(context.Background(), envID, userID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadAllDigitalWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllDigitalWalletsRequest struct via the builder pattern


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


## ReadOneDigitalWallet

> ReadOneDigitalWallet(ctx, envID, userID, digitalWalletID).Authorization(authorization).Execute()

Read One Digital Wallet



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
    userID := "userID_example" // string | 
    digitalWalletID := "digitalWalletID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadOneDigitalWallet(context.Background(), envID, userID, digitalWalletID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadOneDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDigitalWalletRequest struct via the builder pattern


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


## ReadOneDigitalWalletCredentials

> ReadOneDigitalWalletCredentials(ctx, envID, userID, digitalWalletID).Authorization(authorization).Execute()

Read One Digital Wallet Credentials



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
    userID := "userID_example" // string | 
    digitalWalletID := "digitalWalletID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadOneDigitalWalletCredentials(context.Background(), envID, userID, digitalWalletID).Authorization(authorization).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.ReadOneDigitalWalletCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneDigitalWalletCredentialsRequest struct via the builder pattern


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


## UpdateDigitalWallet

> UpdateDigitalWallet(ctx, envID, userID, digitalWalletID).Authorization(authorization).Body(body).Execute()

Update Digital Wallet



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
    userID := "userID_example" // string | 
    digitalWalletID := "digitalWalletID_example" // string | 
    authorization := "Bearer {{jwtToken}}" // string |  (optional)
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.UpdateDigitalWallet(context.Background(), envID, userID, digitalWalletID).Authorization(authorization).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsCredentialDigitalWalletsDigitalWalletsApi.UpdateDigitalWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**envID** | **string** |  | 
**userID** | **string** |  | 
**digitalWalletID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDigitalWalletRequest struct via the builder pattern


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

