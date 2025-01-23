# \ManifestsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TasksTaskIdManifestGet**](ManifestsAPI.md#TasksTaskIdManifestGet) | **Get** /tasks/{task_id}/manifest | Get manifest for task



## TasksTaskIdManifestGet

> ManifestManifest TasksTaskIdManifestGet(ctx, taskId).Execute()

Get manifest for task



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
	resp, r, err := apiClient.ManifestsAPI.TasksTaskIdManifestGet(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestsAPI.TasksTaskIdManifestGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TasksTaskIdManifestGet`: ManifestManifest
	fmt.Fprintf(os.Stdout, "Response from `ManifestsAPI.TasksTaskIdManifestGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTasksTaskIdManifestGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ManifestManifest**](ManifestManifest.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

