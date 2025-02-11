/*
Test Selection Services

Test Selection services, owner: DevProd Services & Integrations team

API version: 1.0.0
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


// TestSelectionAPIService TestSelectionAPI service
type TestSelectionAPIService service

type ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest struct {
	ctx context.Context
	ApiService *TestSelectionAPIService
	projectId string
	requester string
	buildVariantName string
	taskId string
	taskName string
	bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost *BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost
}

func (r ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest) BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest {
	r.bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost = &bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost
	return r
}

func (r ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest) Execute() ([]Explanation, *http.Response, error) {
	return r.ApiService.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostExecute(r)
}

/*
ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost Explain Select Tests

Get the results of the default strategy as well as the explanation.
:param project_id: The project id.
:param requester: The requester.
:param build_variant_name: The build variant name.
:param task_id: The task id.
:param task_name: The task name.
:param test_names: The test names.
:param strategies: The set of strategies to use
:return: A mapping from {Build variant x task x test} to a
         true/false value indicating if the test should run

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId
 @param requester
 @param buildVariantName
 @param taskId
 @param taskName
 @return ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest
*/
func (a *TestSelectionAPIService) ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx context.Context, projectId string, requester string, buildVariantName string, taskId string, taskName string) ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest {
	return ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		requester: requester,
		buildVariantName: buildVariantName,
		taskId: taskId,
		taskName: taskName,
	}
}

// Execute executes the request
//  @return []Explanation
func (a *TestSelectionAPIService) ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostExecute(r ApiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest) ([]Explanation, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Explanation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSelectionAPIService.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/test_selection/explain_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requester"+"}", url.PathEscape(parameterValueToString(r.requester, "requester")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build_variant_name"+"}", url.PathEscape(parameterValueToString(r.buildVariantName, "buildVariantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_name"+"}", url.PathEscape(parameterValueToString(r.taskName, "taskName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost == nil {
		return localVarReturnValue, nil, reportError("bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost is required and must be specified")
	}

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
	localVarPostBody = r.bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest struct {
	ctx context.Context
	ApiService *TestSelectionAPIService
	projectId string
	requester string
	buildVariantName string
	taskId string
	taskName string
	bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost *BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost
}

func (r ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest) BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest {
	r.bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost = &bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost
	return r
}

func (r ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostExecute(r)
}

/*
SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost Select Tests

Get the results of the default strategy.
:param project_id: The project id.
:param requester: The requester.
:param build_variant_name: The build variant name.
:param task_id: The task id.
:param task_name: The task name.
:param test_names: The test names.
:param strategies: The set of strategies to use
:return: A mapping from {Build variant x task x test} to a
         true/false value indicating if the test should run

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param projectId
 @param requester
 @param buildVariantName
 @param taskId
 @param taskName
 @return ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest
*/
func (a *TestSelectionAPIService) SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx context.Context, projectId string, requester string, buildVariantName string, taskId string, taskName string) ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest {
	return ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest{
		ApiService: a,
		ctx: ctx,
		projectId: projectId,
		requester: requester,
		buildVariantName: buildVariantName,
		taskId: taskId,
		taskName: taskName,
	}
}

// Execute executes the request
//  @return []string
func (a *TestSelectionAPIService) SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostExecute(r ApiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSelectionAPIService.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/test_selection/select_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"project_id"+"}", url.PathEscape(parameterValueToString(r.projectId, "projectId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"requester"+"}", url.PathEscape(parameterValueToString(r.requester, "requester")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"build_variant_name"+"}", url.PathEscape(parameterValueToString(r.buildVariantName, "buildVariantName")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"task_name"+"}", url.PathEscape(parameterValueToString(r.taskName, "taskName")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost == nil {
		return localVarReturnValue, nil, reportError("bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost is required and must be specified")
	}

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
	localVarPostBody = r.bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
