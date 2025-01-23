# \AnnotationsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksBuildIdAnnotationsGet**](AnnotationsAPI.md#TasksBuildIdAnnotationsGet) | **Get** /tasks/{build_id}/annotations | List task annotations by build
[**TasksTaskIdAnnotationPatch**](AnnotationsAPI.md#TasksTaskIdAnnotationPatch) | **Patch** /tasks/{task_id}/annotation | Create or update a new task annotation by appending
[**TasksTaskIdAnnotationPut**](AnnotationsAPI.md#TasksTaskIdAnnotationPut) | **Put** /tasks/{task_id}/annotation | Create or update a new task annotation
[**TasksTaskIdAnnotationsGet**](AnnotationsAPI.md#TasksTaskIdAnnotationsGet) | **Get** /tasks/{task_id}/annotations | Fetch task annotations
[**TasksTaskIdCreatedTicketPut**](AnnotationsAPI.md#TasksTaskIdCreatedTicketPut) | **Put** /tasks/{task_id}/created_ticket | Send a newly created ticket for a task.
[**TasksVersionIdAnnotationsGet**](AnnotationsAPI.md#TasksVersionIdAnnotationsGet) | **Get** /tasks/{version_id}/annotations | List task annotations by version



## TasksBuildIdAnnotationsGet

> []ModelAPITaskAnnotation TasksBuildIdAnnotationsGet(ctx, buildId).FetchAllExecutions(fetchAllExecutions).Execute()

List task annotations by build



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
	buildId := "buildId_example" // string | build_id
	fetchAllExecutions := "fetchAllExecutions_example" // string | Fetches previous executions of the task if they are available (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationsAPI.TasksBuildIdAnnotationsGet(context.Background(), buildId).FetchAllExecutions(fetchAllExecutions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsAPI.TasksBuildIdAnnotationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksBuildIdAnnotationsGet`: []ModelAPITaskAnnotation
	fmt.Fprintf(os.Stdout, "Response from `AnnotationsAPI.TasksBuildIdAnnotationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | build_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksBuildIdAnnotationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchAllExecutions** | **string** | Fetches previous executions of the task if they are available | 

### Return type

[**[]ModelAPITaskAnnotation**](ModelAPITaskAnnotation.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdAnnotationPatch

> TasksTaskIdAnnotationPatch(ctx, taskId).Object(object).Execution(execution).Upsert(upsert).Execute()

Create or update a new task annotation by appending



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
	object := *openapiclient.NewModelAPITaskAnnotation() // ModelAPITaskAnnotation | parameters
	execution := int32(56) // int32 | Can be set in lieu of specifying task_execution in the request body. (optional)
	upsert := true // bool | Will create a new annotation if task annotation isn't found and upsert is true. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnotationsAPI.TasksTaskIdAnnotationPatch(context.Background(), taskId).Object(object).Execution(execution).Upsert(upsert).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsAPI.TasksTaskIdAnnotationPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdAnnotationPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPITaskAnnotation**](ModelAPITaskAnnotation.md) | parameters | 
 **execution** | **int32** | Can be set in lieu of specifying task_execution in the request body. | 
 **upsert** | **bool** | Will create a new annotation if task annotation isn&#39;t found and upsert is true. | 

### Return type

 (empty response body)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdAnnotationPut

> TasksTaskIdAnnotationPut(ctx, taskId).Object(object).Execution(execution).Execute()

Create or update a new task annotation



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
	object := *openapiclient.NewModelAPITaskAnnotation() // ModelAPITaskAnnotation | parameters
	execution := int32(56) // int32 | Can be set in lieu of specifying task_execution in the request body. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnotationsAPI.TasksTaskIdAnnotationPut(context.Background(), taskId).Object(object).Execution(execution).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsAPI.TasksTaskIdAnnotationPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdAnnotationPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPITaskAnnotation**](ModelAPITaskAnnotation.md) | parameters | 
 **execution** | **int32** | Can be set in lieu of specifying task_execution in the request body. | 

### Return type

 (empty response body)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdAnnotationsGet

> []ModelAPITaskAnnotation TasksTaskIdAnnotationsGet(ctx, taskId).Execution(execution).FetchAllExecutions(fetchAllExecutions).Execute()

Fetch task annotations



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
	execution := int32(56) // int32 | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution (optional)
	fetchAllExecutions := "fetchAllExecutions_example" // string | Fetches previous executions of the task if they are available (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationsAPI.TasksTaskIdAnnotationsGet(context.Background(), taskId).Execution(execution).FetchAllExecutions(fetchAllExecutions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsAPI.TasksTaskIdAnnotationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdAnnotationsGet`: []ModelAPITaskAnnotation
	fmt.Fprintf(os.Stdout, "Response from `AnnotationsAPI.TasksTaskIdAnnotationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdAnnotationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execution** | **int32** | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution | 
 **fetchAllExecutions** | **string** | Fetches previous executions of the task if they are available | 

### Return type

[**[]ModelAPITaskAnnotation**](ModelAPITaskAnnotation.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdCreatedTicketPut

> TasksTaskIdCreatedTicketPut(ctx, taskId).Object(object).Execute()

Send a newly created ticket for a task.



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
	object := *openapiclient.NewModelAPIIssueLink() // ModelAPIIssueLink | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AnnotationsAPI.TasksTaskIdCreatedTicketPut(context.Background(), taskId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsAPI.TasksTaskIdCreatedTicketPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdCreatedTicketPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPIIssueLink**](ModelAPIIssueLink.md) | parameters | 

### Return type

 (empty response body)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksVersionIdAnnotationsGet

> []ModelAPITaskAnnotation TasksVersionIdAnnotationsGet(ctx, versionId).FetchAllExecutions(fetchAllExecutions).Execute()

List task annotations by version



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
	versionId := "versionId_example" // string | version_id
	fetchAllExecutions := "fetchAllExecutions_example" // string | Fetches previous executions of the task if they are available (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnotationsAPI.TasksVersionIdAnnotationsGet(context.Background(), versionId).FetchAllExecutions(fetchAllExecutions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnotationsAPI.TasksVersionIdAnnotationsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksVersionIdAnnotationsGet`: []ModelAPITaskAnnotation
	fmt.Fprintf(os.Stdout, "Response from `AnnotationsAPI.TasksVersionIdAnnotationsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | version_id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksVersionIdAnnotationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fetchAllExecutions** | **string** | Fetches previous executions of the task if they are available | 

### Return type

[**[]ModelAPITaskAnnotation**](ModelAPITaskAnnotation.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

