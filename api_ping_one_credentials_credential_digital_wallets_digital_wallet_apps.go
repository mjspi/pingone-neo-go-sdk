/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

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


// PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApi service
type PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService service

type ApiCreateDigitalWalletAppRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService
	environmentID string
	body *map[string]interface{}
}

func (r ApiCreateDigitalWalletAppRequest) Body(body map[string]interface{}) ApiCreateDigitalWalletAppRequest {
	r.body = &body
	return r
}

func (r ApiCreateDigitalWalletAppRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateDigitalWalletAppExecute(r)
}

/*
CreateDigitalWalletApp Create Digital Wallet App

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiCreateDigitalWalletAppRequest
*/
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) CreateDigitalWalletApp(ctx context.Context, environmentID string) ApiCreateDigitalWalletAppRequest {
	return ApiCreateDigitalWalletAppRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) CreateDigitalWalletAppExecute(r ApiCreateDigitalWalletAppRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService.CreateDigitalWalletApp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/digitalWalletApplications"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

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

type ApiDeleteDigitalWalletAppRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService
	environmentID string
	digitalWalletApplicationID string
}

func (r ApiDeleteDigitalWalletAppRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteDigitalWalletAppExecute(r)
}

/*
DeleteDigitalWalletApp Delete Digital Wallet App

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param digitalWalletApplicationID
 @return ApiDeleteDigitalWalletAppRequest
*/
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) DeleteDigitalWalletApp(ctx context.Context, environmentID string, digitalWalletApplicationID string) ApiDeleteDigitalWalletAppRequest {
	return ApiDeleteDigitalWalletAppRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		digitalWalletApplicationID: digitalWalletApplicationID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) DeleteDigitalWalletAppExecute(r ApiDeleteDigitalWalletAppRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService.DeleteDigitalWalletApp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"digitalWalletApplicationID"+"}", url.PathEscape(parameterValueToString(r.digitalWalletApplicationID, "digitalWalletApplicationID")), -1)

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

type ApiReadAllDigitalWalletAppsRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService
	environmentID string
}

func (r ApiReadAllDigitalWalletAppsRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadAllDigitalWalletAppsExecute(r)
}

/*
ReadAllDigitalWalletApps Read All Digital Wallet Apps

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiReadAllDigitalWalletAppsRequest
*/
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) ReadAllDigitalWalletApps(ctx context.Context, environmentID string) ApiReadAllDigitalWalletAppsRequest {
	return ApiReadAllDigitalWalletAppsRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) ReadAllDigitalWalletAppsExecute(r ApiReadAllDigitalWalletAppsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService.ReadAllDigitalWalletApps")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/digitalWalletApplications"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

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

type ApiReadOneDigitalWalletAppRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService
	environmentID string
	digitalWalletApplicationID string
}

func (r ApiReadOneDigitalWalletAppRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReadOneDigitalWalletAppExecute(r)
}

/*
ReadOneDigitalWalletApp Read One Digital Wallet App

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param digitalWalletApplicationID
 @return ApiReadOneDigitalWalletAppRequest
*/
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) ReadOneDigitalWalletApp(ctx context.Context, environmentID string, digitalWalletApplicationID string) ApiReadOneDigitalWalletAppRequest {
	return ApiReadOneDigitalWalletAppRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		digitalWalletApplicationID: digitalWalletApplicationID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) ReadOneDigitalWalletAppExecute(r ApiReadOneDigitalWalletAppRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService.ReadOneDigitalWalletApp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"digitalWalletApplicationID"+"}", url.PathEscape(parameterValueToString(r.digitalWalletApplicationID, "digitalWalletApplicationID")), -1)

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

type ApiUpdateDigitalWalletAppRequest struct {
	ctx context.Context
	ApiService *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService
	environmentID string
	digitalWalletApplicationID string
	body *map[string]interface{}
}

func (r ApiUpdateDigitalWalletAppRequest) Body(body map[string]interface{}) ApiUpdateDigitalWalletAppRequest {
	r.body = &body
	return r
}

func (r ApiUpdateDigitalWalletAppRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateDigitalWalletAppExecute(r)
}

/*
UpdateDigitalWalletApp Update Digital Wallet App

This PingOne collection contains only the REST API request examples without documentation. For complete documentation, go to <a href="https://apidocs.pingidentity.com/pingone/platform/v1/api/">apidocs.pingidentity.com</a>.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @param digitalWalletApplicationID
 @return ApiUpdateDigitalWalletAppRequest
*/
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) UpdateDigitalWalletApp(ctx context.Context, environmentID string, digitalWalletApplicationID string) ApiUpdateDigitalWalletAppRequest {
	return ApiUpdateDigitalWalletAppRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
		digitalWalletApplicationID: digitalWalletApplicationID,
	}
}

// Execute executes the request
func (a *PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService) UpdateDigitalWalletAppExecute(r ApiUpdateDigitalWalletAppRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PingOneCredentialsCredentialDigitalWalletsDigitalWalletAppsApiService.UpdateDigitalWalletApp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/digitalWalletApplications/{digitalWalletApplicationID}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"digitalWalletApplicationID"+"}", url.PathEscape(parameterValueToString(r.digitalWalletApplicationID, "digitalWalletApplicationID")), -1)

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
