/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingONe Credentials service

API version: 2023-03-30
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// PingOneCredentialsUserCredentialsApiService PingOneCredentialsUserCredentialsApi service
type PingOneCredentialsUserCredentialsApiService service

type ApiCreateAUserCredentialRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsUserCredentialsApiService
	envID string
	userID string
	authorization *string
	body *map[string]interface{}
}

func (r ApiCreateAUserCredentialRequest) Authorization(authorization string) ApiCreateAUserCredentialRequest {
	r.authorization = &authorization
	return r
}

func (r ApiCreateAUserCredentialRequest) Body(body map[string]interface{}) ApiCreateAUserCredentialRequest {
	r.body = &body
	return r
}

func (r ApiCreateAUserCredentialRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateAUserCredentialExecute(r)
}

/*
CreateAUserCredential Create a User Credential

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiCreateAUserCredentialRequest
*/
func (a *PingOneCredentialsUserCredentialsApiService) CreateAUserCredential(ctx context.Context, envID string, userID string) ApiCreateAUserCredentialRequest {
	return ApiCreateAUserCredentialRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsUserCredentialsApiService) CreateAUserCredentialExecute(r ApiCreateAUserCredentialRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsUserCredentialsApiService.CreateAUserCredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{envID}/users/{userID}/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterValueToString(r.envID, "envID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiReadAllUserCredentialsRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsUserCredentialsApiService
	envID string
	userID string
	authorization *string
}

func (r ApiReadAllUserCredentialsRequest) Authorization(authorization string) ApiReadAllUserCredentialsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiReadAllUserCredentialsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadAllUserCredentialsExecute(r)
}

/*
ReadAllUserCredentials Read All User Credentials

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @return ApiReadAllUserCredentialsRequest
*/
func (a *PingOneCredentialsUserCredentialsApiService) ReadAllUserCredentials(ctx context.Context, envID string, userID string) ApiReadAllUserCredentialsRequest {
	return ApiReadAllUserCredentialsRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsUserCredentialsApiService) ReadAllUserCredentialsExecute(r ApiReadAllUserCredentialsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsUserCredentialsApiService.ReadAllUserCredentials")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{envID}/users/{userID}/credentials"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterValueToString(r.envID, "envID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiReadOneUserCredentialRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsUserCredentialsApiService
	envID string
	userID string
	credentialID string
	authorization *string
}

func (r ApiReadOneUserCredentialRequest) Authorization(authorization string) ApiReadOneUserCredentialRequest {
	r.authorization = &authorization
	return r
}

func (r ApiReadOneUserCredentialRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadOneUserCredentialExecute(r)
}

/*
ReadOneUserCredential Read One User Credential

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @param credentialID
 @return ApiReadOneUserCredentialRequest
*/
func (a *PingOneCredentialsUserCredentialsApiService) ReadOneUserCredential(ctx context.Context, envID string, userID string, credentialID string) ApiReadOneUserCredentialRequest {
	return ApiReadOneUserCredentialRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		credentialID: credentialID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsUserCredentialsApiService) ReadOneUserCredentialExecute(r ApiReadOneUserCredentialRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsUserCredentialsApiService.ReadOneUserCredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{envID}/users/{userID}/credentials/{credentialID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterValueToString(r.envID, "envID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentialID"+"}", url.PathEscape(parameterValueToString(r.credentialID, "credentialID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiReadOneUserCredentialWalletsRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsUserCredentialsApiService
	envID string
	userID string
	credentialID string
	authorization *string
}

func (r ApiReadOneUserCredentialWalletsRequest) Authorization(authorization string) ApiReadOneUserCredentialWalletsRequest {
	r.authorization = &authorization
	return r
}

func (r ApiReadOneUserCredentialWalletsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadOneUserCredentialWalletsExecute(r)
}

/*
ReadOneUserCredentialWallets Read One User Credential Wallets

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @param credentialID
 @return ApiReadOneUserCredentialWalletsRequest
*/
func (a *PingOneCredentialsUserCredentialsApiService) ReadOneUserCredentialWallets(ctx context.Context, envID string, userID string, credentialID string) ApiReadOneUserCredentialWalletsRequest {
	return ApiReadOneUserCredentialWalletsRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		credentialID: credentialID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsUserCredentialsApiService) ReadOneUserCredentialWalletsExecute(r ApiReadOneUserCredentialWalletsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsUserCredentialsApiService.ReadOneUserCredentialWallets")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{envID}/users/{userID}/credentials/{credentialID}/provisionedCredentials"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterValueToString(r.envID, "envID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentialID"+"}", url.PathEscape(parameterValueToString(r.credentialID, "credentialID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiUpdateAUserCredentialRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsUserCredentialsApiService
	envID string
	userID string
	credentialID string
	authorization *string
	body *map[string]interface{}
}

func (r ApiUpdateAUserCredentialRequest) Authorization(authorization string) ApiUpdateAUserCredentialRequest {
	r.authorization = &authorization
	return r
}

func (r ApiUpdateAUserCredentialRequest) Body(body map[string]interface{}) ApiUpdateAUserCredentialRequest {
	r.body = &body
	return r
}

func (r ApiUpdateAUserCredentialRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateAUserCredentialExecute(r)
}

/*
UpdateAUserCredential Update a User Credential

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param envID
 @param userID
 @param credentialID
 @return ApiUpdateAUserCredentialRequest
*/
func (a *PingOneCredentialsUserCredentialsApiService) UpdateAUserCredential(ctx context.Context, envID string, userID string, credentialID string) ApiUpdateAUserCredentialRequest {
	return ApiUpdateAUserCredentialRequest{
		ApiService: a,
		ctx: ctx,
		envID: envID,
		userID: userID,
		credentialID: credentialID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsUserCredentialsApiService) UpdateAUserCredentialExecute(r ApiUpdateAUserCredentialRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsUserCredentialsApiService.UpdateAUserCredential")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{envID}/users/{userID}/credentials/{credentialID}"
	localVarPath = strings.Replace(localVarPath, "{"+"envID"+"}", url.PathEscape(parameterValueToString(r.envID, "envID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userID"+"}", url.PathEscape(parameterValueToString(r.userID, "userID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"credentialID"+"}", url.PathEscape(parameterValueToString(r.credentialID, "credentialID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.authorization != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Authorization", r.authorization, "")
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
