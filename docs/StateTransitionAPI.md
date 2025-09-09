# \StateTransitionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_task/{project_id}/{build_variant_name}/{task_name}/ | Mark Task Random By Design
[**MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_tests/{project_id}/{build_variant_name}/{task_name}/ | Mark Tests Random By Design



## MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost

> interface{} MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).Requester(requester).IsRandom(isRandom).Execute()

Mark Task Random By Design



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
	projectId := "projectId_example" // string | 
	buildVariantName := "buildVariantName_example" // string | 
	taskName := "taskName_example" // string | 
	requester := "requester_example" // string | 
	isRandom := true // bool | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).Requester(requester).IsRandom(isRandom).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**buildVariantName** | **string** |  | 
**taskName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkTaskRandomByDesignApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requester** | **string** |  | 
 **isRandom** | **bool** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost

> interface{} MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).IsRandom(isRandom).RequestBody(requestBody).Execute()

Mark Tests Random By Design



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
	projectId := "projectId_example" // string | 
	buildVariantName := "buildVariantName_example" // string | 
	taskName := "taskName_example" // string | 
	isRandom := true // bool | 
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).IsRandom(isRandom).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**buildVariantName** | **string** |  | 
**taskName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkTestsRandomByDesignApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isRandom** | **bool** |  | 
 **requestBody** | **[]string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

