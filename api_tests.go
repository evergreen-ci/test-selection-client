/*
Evergreen REST v2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3
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


// TestsAPIService TestsAPI service
type TestsAPIService service

type ApiTasksTaskIdTestsCountGetRequest struct {
	ctx context.Context
	ApiService *TestsAPIService
	taskId string
	execution *int32
}

// The 0-based number corresponding to the execution of the task. Defaults to 0, meaning the first time the task was run.
func (r ApiTasksTaskIdTestsCountGetRequest) Execution(execution int32) ApiTasksTaskIdTestsCountGetRequest {
	r.execution = &execution
	return r
}

func (r ApiTasksTaskIdTestsCountGetRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.TasksTaskIdTestsCountGetExecute(r)
}

/*
TasksTaskIdTestsCountGet Get the test count from a task

Returns an integer representing the number of tests that ran as part of the given task.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId task ID
 @return ApiTasksTaskIdTestsCountGetRequest
*/
func (a *TestsAPIService) TasksTaskIdTestsCountGet(ctx context.Context, taskId string) ApiTasksTaskIdTestsCountGetRequest {
	return ApiTasksTaskIdTestsCountGetRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return string
func (a *TestsAPIService) TasksTaskIdTestsCountGetExecute(r ApiTasksTaskIdTestsCountGetRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestsAPIService.TasksTaskIdTestsCountGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks/{task_id}/tests/count"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.execution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "execution", r.execution, "", "")
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-User"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-User"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-Key"] = key
			}
		}
	}
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

type ApiTasksTaskIdTestsGetRequest struct {
	ctx context.Context
	ApiService *TestsAPIService
	taskId string
	startAt *string
	limit *int32
	status *string
	execution *int32
	testName *string
	latest *bool
}

// The identifier of the test to start at in the pagination
func (r ApiTasksTaskIdTestsGetRequest) StartAt(startAt string) ApiTasksTaskIdTestsGetRequest {
	r.startAt = &startAt
	return r
}

// The number of tests to be returned per page of pagination. Defaults to 100
func (r ApiTasksTaskIdTestsGetRequest) Limit(limit int32) ApiTasksTaskIdTestsGetRequest {
	r.limit = &limit
	return r
}

// A status of test to limit the results to.
func (r ApiTasksTaskIdTestsGetRequest) Status(status string) ApiTasksTaskIdTestsGetRequest {
	r.status = &status
	return r
}

// The 0-based number corresponding to the execution of the task. Defaults to 0, meaning the first time the task was run.
func (r ApiTasksTaskIdTestsGetRequest) Execution(execution int32) ApiTasksTaskIdTestsGetRequest {
	r.execution = &execution
	return r
}

// Only return the test matching the name.
func (r ApiTasksTaskIdTestsGetRequest) TestName(testName string) ApiTasksTaskIdTestsGetRequest {
	r.testName = &testName
	return r
}

// Return tests from the latest execution. Cannot be used with execution.
func (r ApiTasksTaskIdTestsGetRequest) Latest(latest bool) ApiTasksTaskIdTestsGetRequest {
	r.latest = &latest
	return r
}

func (r ApiTasksTaskIdTestsGetRequest) Execute() ([]ModelAPITest, *http.Response, error) {
	return r.ApiService.TasksTaskIdTestsGetExecute(r)
}

/*
TasksTaskIdTestsGet Get tests from a task

Fetches a paginated list of tests that ran as part of the given task. To filter the tasks, add the following parameters into the query string.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param taskId task ID
 @return ApiTasksTaskIdTestsGetRequest
*/
func (a *TestsAPIService) TasksTaskIdTestsGet(ctx context.Context, taskId string) ApiTasksTaskIdTestsGetRequest {
	return ApiTasksTaskIdTestsGetRequest{
		ApiService: a,
		ctx: ctx,
		taskId: taskId,
	}
}

// Execute executes the request
//  @return []ModelAPITest
func (a *TestsAPIService) TasksTaskIdTestsGetExecute(r ApiTasksTaskIdTestsGetRequest) ([]ModelAPITest, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ModelAPITest
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestsAPIService.TasksTaskIdTestsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks/{task_id}/tests"
	localVarPath = strings.Replace(localVarPath, "{"+"task_id"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startAt != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start_at", r.startAt, "", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "", "")
	}
	if r.execution != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "execution", r.execution, "", "")
	}
	if r.testName != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "test_name", r.testName, "", "")
	}
	if r.latest != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "latest", r.latest, "", "")
	}
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-User"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-User"] = key
			}
		}
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Api-Key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Api-Key"] = key
			}
		}
	}
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
