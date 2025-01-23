# \TasksAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildsBuildIdTasksGet**](TasksAPI.md#BuildsBuildIdTasksGet) | **Get** /builds/{build_id}/tasks | List tasks by build
[**ProjectsProjectIdTaskExecutionsGet**](TasksAPI.md#ProjectsProjectIdTaskExecutionsGet) | **Get** /projects/{project_id}/task_executions | Get execution info for a task
[**ProjectsProjectIdTasksTaskNameGet**](TasksAPI.md#ProjectsProjectIdTasksTaskNameGet) | **Get** /projects/{project_id}/tasks/{task_name} | Get tasks for a project
[**ProjectsProjectNameRevisionsCommitHashTasksGet**](TasksAPI.md#ProjectsProjectNameRevisionsCommitHashTasksGet) | **Get** /projects/{project_name}/revisions/{commit_hash}/tasks | List tasks by project and commit
[**TasksTaskIdAbortPost**](TasksAPI.md#TasksTaskIdAbortPost) | **Post** /tasks/{task_id}/abort | Abort a task
[**TasksTaskIdBuildTaskLogsGet**](TasksAPI.md#TasksTaskIdBuildTaskLogsGet) | **Get** /tasks/{task_id}/build/TaskLogs | Get task logs for a task.
[**TasksTaskIdBuildTestLogsPathGet**](TasksAPI.md#TasksTaskIdBuildTestLogsPathGet) | **Get** /tasks/{task_id}/build/TestLogs/{path} | Get test logs for a task.
[**TasksTaskIdGeneratedTasksGet**](TasksAPI.md#TasksTaskIdGeneratedTasksGet) | **Get** /tasks/{task_id}/generated_tasks | Get info about generated tasks
[**TasksTaskIdGet**](TasksAPI.md#TasksTaskIdGet) | **Get** /tasks/{task_id} | Get a single task
[**TasksTaskIdPatch**](TasksAPI.md#TasksTaskIdPatch) | **Patch** /tasks/{task_id} | Change a task&#39;s execution status
[**TasksTaskIdRestartPost**](TasksAPI.md#TasksTaskIdRestartPost) | **Post** /tasks/{task_id}/restart | Restart a task



## BuildsBuildIdTasksGet

> []ModelAPITask BuildsBuildIdTasksGet(ctx, buildId).StartAt(startAt).Limit(limit).FetchAllExecutions(fetchAllExecutions).FetchParentIds(fetchParentIds).Execute()

List tasks by build



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
	buildId := "buildId_example" // string | the build ID
	startAt := "startAt_example" // string | The identifier of the task to start at in the pagination (optional)
	limit := int32(56) // int32 | The number of tasks to be returned per page of pagination. Defaults to 100 (optional)
	fetchAllExecutions := true // bool | Fetches previous executions of tasks if they are available (optional)
	fetchParentIds := true // bool | Fetches the parent display task ID for each returned execution task (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.BuildsBuildIdTasksGet(context.Background(), buildId).StartAt(startAt).Limit(limit).FetchAllExecutions(fetchAllExecutions).FetchParentIds(fetchParentIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.BuildsBuildIdTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBuildIdTasksGet`: []ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.BuildsBuildIdTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | the build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildIdTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **string** | The identifier of the task to start at in the pagination | 
 **limit** | **int32** | The number of tasks to be returned per page of pagination. Defaults to 100 | 
 **fetchAllExecutions** | **bool** | Fetches previous executions of tasks if they are available | 
 **fetchParentIds** | **bool** | Fetches the parent display task ID for each returned execution task | 

### Return type

[**[]ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdTaskExecutionsGet

> ModelProjectTaskExecutionResp ProjectsProjectIdTaskExecutionsGet(ctx, projectId).TaskName(taskName).BuildVariant(buildVariant).StartTime(startTime).EndTime(endTime).Requesters(requesters).Execute()

Get execution info for a task



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
	projectId := "projectId_example" // string | the project ID
	taskName := "taskName_example" // string | The task to return execution info for.
	buildVariant := "buildVariant_example" // string | The build variant to return task execution info for.
	startTime := "startTime_example" // string | Will only return execution info after this time. Format should be 2022-12-01T12:30:00.000Z
	endTime := "endTime_example" // string | If not provided, will default to the current time. (optional)
	requesters := []string{"Inner_example"} // []string | If not provided, will default to gitter_request (versions created by git commit). Can also be github_pull_request, trigger_request (Project Trigger versions) , github_merge_request (GitHub merge queue), or ad_hoc (periodic builds) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ProjectsProjectIdTaskExecutionsGet(context.Background(), projectId).TaskName(taskName).BuildVariant(buildVariant).StartTime(startTime).EndTime(endTime).Requesters(requesters).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ProjectsProjectIdTaskExecutionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdTaskExecutionsGet`: ModelProjectTaskExecutionResp
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ProjectsProjectIdTaskExecutionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdTaskExecutionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **taskName** | **string** | The task to return execution info for. | 
 **buildVariant** | **string** | The build variant to return task execution info for. | 
 **startTime** | **string** | Will only return execution info after this time. Format should be 2022-12-01T12:30:00.000Z | 
 **endTime** | **string** | If not provided, will default to the current time. | 
 **requesters** | **[]string** | If not provided, will default to gitter_request (versions created by git commit). Can also be github_pull_request, trigger_request (Project Trigger versions) , github_merge_request (GitHub merge queue), or ad_hoc (periodic builds) | 

### Return type

[**ModelProjectTaskExecutionResp**](ModelProjectTaskExecutionResp.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdTasksTaskNameGet

> []ModelAPITask ProjectsProjectIdTasksTaskNameGet(ctx, projectId, taskName).BuildVariant(buildVariant).Execute()

Get tasks for a project



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
	projectId := "projectId_example" // string | the project ID
	taskName := "taskName_example" // string | the task name
	buildVariant := "buildVariant_example" // string | If set, will only include tasks that have run on this build variant. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ProjectsProjectIdTasksTaskNameGet(context.Background(), projectId, taskName).BuildVariant(buildVariant).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ProjectsProjectIdTasksTaskNameGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdTasksTaskNameGet`: []ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ProjectsProjectIdTasksTaskNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 
**taskName** | **string** | the task name | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdTasksTaskNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **buildVariant** | **string** | If set, will only include tasks that have run on this build variant. | 

### Return type

[**[]ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectNameRevisionsCommitHashTasksGet

> []ModelAPITask ProjectsProjectNameRevisionsCommitHashTasksGet(ctx, projectName, commitHash).StartAt(startAt).Limit(limit).Variant(variant).VariantRegex(variantRegex).TaskName(taskName).Status(status).Execute()

List tasks by project and commit



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
	projectName := "projectName_example" // string | project name
	commitHash := "commitHash_example" // string | commit hash
	startAt := "startAt_example" // string | The identifier of the task to start at in the pagination (optional)
	limit := int32(56) // int32 | The number of tasks to be returned per page of pagination. Defaults to 100 (optional)
	variant := "variant_example" // string | Only return tasks within this variant (optional)
	variantRegex := "variantRegex_example" // string | Only return tasks within variants that match this regex (optional)
	taskName := "taskName_example" // string | Only return tasks with this display name (optional)
	status := "status_example" // string | Only return tasks with this status (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.ProjectsProjectNameRevisionsCommitHashTasksGet(context.Background(), projectName, commitHash).StartAt(startAt).Limit(limit).Variant(variant).VariantRegex(variantRegex).TaskName(taskName).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.ProjectsProjectNameRevisionsCommitHashTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectNameRevisionsCommitHashTasksGet`: []ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.ProjectsProjectNameRevisionsCommitHashTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectName** | **string** | project name | 
**commitHash** | **string** | commit hash | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectNameRevisionsCommitHashTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **string** | The identifier of the task to start at in the pagination | 
 **limit** | **int32** | The number of tasks to be returned per page of pagination. Defaults to 100 | 
 **variant** | **string** | Only return tasks within this variant | 
 **variantRegex** | **string** | Only return tasks within variants that match this regex | 
 **taskName** | **string** | Only return tasks with this display name | 
 **status** | **string** | Only return tasks with this status | 

### Return type

[**[]ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdAbortPost

> ModelAPITask TasksTaskIdAbortPost(ctx, taskId).Execute()

Abort a task



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksTaskIdAbortPost(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdAbortPost`: ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdBuildTaskLogsGet

> string TasksTaskIdBuildTaskLogsGet(ctx, taskId).Execution(execution).Type_(type_).Start(start).End(end).LineLimit(lineLimit).TailLimit(tailLimit).PrintTime(printTime).PrintPriority(printPriority).Paginate(paginate).Execute()

Get task logs for a task.



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
	taskId := "taskId_example" // string | Task ID.
	execution := int32(56) // int32 | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution. (optional)
	type_ := "type__example" // string | Task log type. Must be one of: `agent_log`, `system_log`, `task_log`, `all_logs`. Defaults to `all_logs`. (optional)
	start := "start_example" // string | Start of targeted time interval (inclusive) in RFC3339 format. Defaults to the first timestamp of the requested logs. (optional)
	end := "end_example" // string | End of targeted time interval (inclusive) in RFC3339 format. Defaults to the last timestamp of the requested logs. (optional)
	lineLimit := int32(56) // int32 | If set greater than 0, limits the number of log lines returned. (optional)
	tailLimit := int32(56) // int32 | If set greater than 0, returns the last N log lines. (optional)
	printTime := true // bool | If set to true, returns log lines prefixed with their timestamp. (optional)
	printPriority := true // bool | If set to true, returns log lines prefixed with their priority. (optional)
	paginate := true // bool | If set to true, paginates the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksTaskIdBuildTaskLogsGet(context.Background(), taskId).Execution(execution).Type_(type_).Start(start).End(end).LineLimit(lineLimit).TailLimit(tailLimit).PrintTime(printTime).PrintPriority(printPriority).Paginate(paginate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdBuildTaskLogsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdBuildTaskLogsGet`: string
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdBuildTaskLogsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Task ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdBuildTaskLogsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execution** | **int32** | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution. | 
 **type_** | **string** | Task log type. Must be one of: &#x60;agent_log&#x60;, &#x60;system_log&#x60;, &#x60;task_log&#x60;, &#x60;all_logs&#x60;. Defaults to &#x60;all_logs&#x60;. | 
 **start** | **string** | Start of targeted time interval (inclusive) in RFC3339 format. Defaults to the first timestamp of the requested logs. | 
 **end** | **string** | End of targeted time interval (inclusive) in RFC3339 format. Defaults to the last timestamp of the requested logs. | 
 **lineLimit** | **int32** | If set greater than 0, limits the number of log lines returned. | 
 **tailLimit** | **int32** | If set greater than 0, returns the last N log lines. | 
 **printTime** | **bool** | If set to true, returns log lines prefixed with their timestamp. | 
 **printPriority** | **bool** | If set to true, returns log lines prefixed with their priority. | 
 **paginate** | **bool** | If set to true, paginates the response. | 

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


## TasksTaskIdBuildTestLogsPathGet

> string TasksTaskIdBuildTestLogsPathGet(ctx, taskId, path).Execution(execution).LogsToMerge(logsToMerge).Start(start).End(end).LineLimit(lineLimit).TailLimit(tailLimit).PrintTime(printTime).PrintPriority(printPriority).Paginate(paginate).Execute()

Get test logs for a task.



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
	taskId := "taskId_example" // string | Task ID.
	path := "path_example" // string | Test log path relative to the task's test logs directory.
	execution := int32(56) // int32 | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution. (optional)
	logsToMerge := "logsToMerge_example" // string | Test log path, relative to the task's test log directory, to merge with test log specified in the URL path. Can be a prefix. Merging is stable and timestamp-based. Repeat the parameter key if more than one value. (optional)
	start := "start_example" // string | Start of targeted time interval (inclusive) in RFC3339 format. Defaults to the first timestamp of the test log specified in the URL path. (optional)
	end := "end_example" // string | End of targeted time interval (inclusive) in RFC3339 format. Defaults to the last timestamp of the test log specified in the URL path. (optional)
	lineLimit := int32(56) // int32 | If set greater than 0, limits the number of log lines returned. (optional)
	tailLimit := int32(56) // int32 | If set greater than 0, returns the last N log lines. (optional)
	printTime := true // bool | If set to true, returns log lines prefixed with their timestamp. (optional)
	printPriority := true // bool | If set to true, returns log lines prefixed with their priority. (optional)
	paginate := true // bool | If set to true, paginates the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksTaskIdBuildTestLogsPathGet(context.Background(), taskId, path).Execution(execution).LogsToMerge(logsToMerge).Start(start).End(end).LineLimit(lineLimit).TailLimit(tailLimit).PrintTime(printTime).PrintPriority(printPriority).Paginate(paginate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdBuildTestLogsPathGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdBuildTestLogsPathGet`: string
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdBuildTestLogsPathGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | Task ID. | 
**path** | **string** | Test log path relative to the task&#39;s test logs directory. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdBuildTestLogsPathGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **execution** | **int32** | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution. | 
 **logsToMerge** | **string** | Test log path, relative to the task&#39;s test log directory, to merge with test log specified in the URL path. Can be a prefix. Merging is stable and timestamp-based. Repeat the parameter key if more than one value. | 
 **start** | **string** | Start of targeted time interval (inclusive) in RFC3339 format. Defaults to the first timestamp of the test log specified in the URL path. | 
 **end** | **string** | End of targeted time interval (inclusive) in RFC3339 format. Defaults to the last timestamp of the test log specified in the URL path. | 
 **lineLimit** | **int32** | If set greater than 0, limits the number of log lines returned. | 
 **tailLimit** | **int32** | If set greater than 0, returns the last N log lines. | 
 **printTime** | **bool** | If set to true, returns log lines prefixed with their timestamp. | 
 **printPriority** | **bool** | If set to true, returns log lines prefixed with their priority. | 
 **paginate** | **bool** | If set to true, paginates the response. | 

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


## TasksTaskIdGeneratedTasksGet

> []ModelAPIGeneratedTaskInfo TasksTaskIdGeneratedTasksGet(ctx, taskId).Execute()

Get info about generated tasks



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksTaskIdGeneratedTasksGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdGeneratedTasksGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdGeneratedTasksGet`: []ModelAPIGeneratedTaskInfo
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdGeneratedTasksGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdGeneratedTasksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelAPIGeneratedTaskInfo**](ModelAPIGeneratedTaskInfo.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdGet

> ModelAPITask TasksTaskIdGet(ctx, taskId).Execution(execution).FetchAllExecutions(fetchAllExecutions).Execute()

Get a single task



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
	resp, r, err := apiClient.TasksAPI.TasksTaskIdGet(context.Background(), taskId).Execution(execution).FetchAllExecutions(fetchAllExecutions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdGet`: ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **execution** | **int32** | The 0-based number corresponding to the execution of the task ID. Defaults to the latest execution | 
 **fetchAllExecutions** | **string** | Fetches previous executions of the task if they are available | 

### Return type

[**ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdPatch

> ModelAPITask TasksTaskIdPatch(ctx, taskId).Object(object).Execute()

Change a task's execution status



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
	object := *openapiclient.NewRouteTaskExecutionPatchHandler() // RouteTaskExecutionPatchHandler | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksTaskIdPatch(context.Background(), taskId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdPatch`: ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteTaskExecutionPatchHandler**](RouteTaskExecutionPatchHandler.md) | parameters | 

### Return type

[**ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TasksTaskIdRestartPost

> ModelAPITask TasksTaskIdRestartPost(ctx, taskId).Object(object).Execute()

Restart a task



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
	object := *openapiclient.NewRouteTaskRestartHandler() // RouteTaskRestartHandler | parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TasksAPI.TasksTaskIdRestartPost(context.Background(), taskId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TasksAPI.TasksTaskIdRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdRestartPost`: ModelAPITask
	fmt.Fprintf(os.Stdout, "Response from `TasksAPI.TasksTaskIdRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteTaskRestartHandler**](RouteTaskRestartHandler.md) | parameters | 

### Return type

[**ModelAPITask**](ModelAPITask.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

