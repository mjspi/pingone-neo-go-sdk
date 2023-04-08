# \PingOneCredentialsUserCredentialsApi

All URIs are relative to *https://api.pingone.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAUserCredential**](PingOneCredentialsUserCredentialsApi.md#CreateAUserCredential) | **Post** /environments/{environmentID}/users/{userID}/credentials | Create a User Credential
[**ReadAllUserCredentials**](PingOneCredentialsUserCredentialsApi.md#ReadAllUserCredentials) | **Get** /environments/{environmentID}/users/{userID}/credentials | Read All User Credentials
[**ReadOneUserCredential**](PingOneCredentialsUserCredentialsApi.md#ReadOneUserCredential) | **Get** /environments/{environmentID}/users/{userID}/credentials/{credentialID} | Read One User Credential
[**ReadOneUserCredentialWallets**](PingOneCredentialsUserCredentialsApi.md#ReadOneUserCredentialWallets) | **Get** /environments/{environmentID}/users/{userID}/credentials/{credentialID}/provisionedCredentials | Read One User Credential Wallets
[**UpdateAUserCredential**](PingOneCredentialsUserCredentialsApi.md#UpdateAUserCredential) | **Put** /environments/{environmentID}/users/{userID}/credentials/{credentialID} | Update a User Credential



## CreateAUserCredential

> UserCredential CreateAUserCredential(ctx, environmentID, userID).UserCredential(userCredential).Execute()

Create a User Credential

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
    userID := "userID_example" // string | 
    userCredential := *openapiclient.NewUserCredential() // UserCredential |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsUserCredentialsApi.CreateAUserCredential(context.Background(), environmentID, userID).UserCredential(userCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsUserCredentialsApi.CreateAUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAUserCredential`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsUserCredentialsApi.CreateAUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userCredential** | [**UserCredential**](UserCredential.md) |  | 

### Return type

[**UserCredential**](UserCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllUserCredentials

> EntityArray ReadAllUserCredentials(ctx, environmentID, userID).Execute()

Read All User Credentials

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
    userID := "userID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsUserCredentialsApi.ReadAllUserCredentials(context.Background(), environmentID, userID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsUserCredentialsApi.ReadAllUserCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllUserCredentials`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsUserCredentialsApi.ReadAllUserCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllUserCredentialsRequest struct via the builder pattern


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


## ReadOneUserCredential

> UserCredential ReadOneUserCredential(ctx, environmentID, userID, credentialID).Execute()

Read One User Credential

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
    userID := "userID_example" // string | 
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsUserCredentialsApi.ReadOneUserCredential(context.Background(), environmentID, userID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsUserCredentialsApi.ReadOneUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneUserCredential`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsUserCredentialsApi.ReadOneUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**UserCredential**](UserCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneUserCredentialWallets

> EntityArray ReadOneUserCredentialWallets(ctx, environmentID, userID, credentialID).Execute()

Read One User Credential Wallets

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
    userID := "userID_example" // string | 
    credentialID := "credentialID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsUserCredentialsApi.ReadOneUserCredentialWallets(context.Background(), environmentID, userID, credentialID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsUserCredentialsApi.ReadOneUserCredentialWallets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneUserCredentialWallets`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsUserCredentialsApi.ReadOneUserCredentialWallets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneUserCredentialWalletsRequest struct via the builder pattern


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


## UpdateAUserCredential

> UserCredential UpdateAUserCredential(ctx, environmentID, userID, credentialID).UserCredential(userCredential).Execute()

Update a User Credential

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
    userID := "userID_example" // string | 
    credentialID := "credentialID_example" // string | 
    userCredential := *openapiclient.NewUserCredential() // UserCredential |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PingOneCredentialsUserCredentialsApi.UpdateAUserCredential(context.Background(), environmentID, userID, credentialID).UserCredential(userCredential).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PingOneCredentialsUserCredentialsApi.UpdateAUserCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAUserCredential`: UserCredential
    fmt.Fprintf(os.Stdout, "Response from `PingOneCredentialsUserCredentialsApi.UpdateAUserCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**credentialID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAUserCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userCredential** | [**UserCredential**](UserCredential.md) |  | 

### Return type

[**UserCredential**](UserCredential.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

