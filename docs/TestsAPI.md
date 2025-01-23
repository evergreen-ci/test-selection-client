# \TestsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksTaskIdTestsCountGet**](TestsAPI.md#TasksTaskIdTestsCountGet) | **Get** /tasks/{task_id}/tests/count | Get the test count from a task
[**TasksTaskIdTestsGet**](TestsAPI.md#TasksTaskIdTestsGet) | **Get** /tasks/{task_id}/tests | Get tests from a task



## TasksTaskIdTestsCountGet

> string TasksTaskIdTestsCountGet(ctx, taskId).Execution(execution).Execute()

Get the test count from a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evergreen-ci/test-selection-client"
)

func main() {
	taskId := "taskId_example" // string | task ID
	execution := int32(56) // int32 | The 0-based number corresponding to the execution of the task. Defaults to 0, meaning the first time the task was run. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TasksTaskIdTestsCountGet(context.Background(), taskId).Execution(execution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TasksTaskIdTestsCountGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdTestsCountGet`: string
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TasksTaskIdTestsCountGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdTestsCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execution** | **int32** | The 0-based number corresponding to the execution of the task. Defaults to 0, meaning the first time the task was run. | 

### Return type

**string**

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdTestsGet

> []ModelAPITest TasksTaskIdTestsGet(ctx, taskId).StartAt(startAt).Limit(limit).Status(status).Execution(execution).TestName(testName).Latest(latest).Execute()

Get tests from a task



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/evergreen-ci/test-selection-client"
)

func main() {
	taskId := "taskId_example" // string | task ID
	startAt := "startAt_example" // string | The identifier of the test to start at in the pagination (optional)
	limit := int32(56) // int32 | The number of tests to be returned per page of pagination. Defaults to 100 (optional)
	status := "status_example" // string | A status of test to limit the results to. (optional)
	execution := int32(56) // int32 | The 0-based number corresponding to the execution of the task. Defaults to 0, meaning the first time the task was run. (optional)
	testName := "testName_example" // string | Only return the test matching the name. (optional)
	latest := true // bool | Return tests from the latest execution. Cannot be used with execution. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestsAPI.TasksTaskIdTestsGet(context.Background(), taskId).StartAt(startAt).Limit(limit).Status(status).Execution(execution).TestName(testName).Latest(latest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestsAPI.TasksTaskIdTestsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdTestsGet`: []ModelAPITest
	fmt.Fprintf(os.Stdout, "Response from `TestsAPI.TasksTaskIdTestsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdTestsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **string** | The identifier of the test to start at in the pagination | 
 **limit** | **int32** | The number of tests to be returned per page of pagination. Defaults to 100 | 
 **status** | **string** | A status of test to limit the results to. | 
 **execution** | **int32** | The 0-based number corresponding to the execution of the task. Defaults to 0, meaning the first time the task was run. | 
 **testName** | **string** | Only return the test matching the name. | 
 **latest** | **bool** | Return tests from the latest execution. Cannot be used with execution. | 

### Return type

[**[]ModelAPITest**](ModelAPITest.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

