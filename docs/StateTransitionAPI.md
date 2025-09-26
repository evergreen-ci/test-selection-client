# \StateTransitionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_task/{project_id}/{build_variant_name}/{task_name}/ | Mark Task As Manually Quarantined
[**MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_tests/{project_id}/{build_variant_name}/{task_name}/ | Mark Tests As Manually Quarantined



## MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost

> interface{} MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).Requester(requester).IsManuallyQuarantined(isManuallyQuarantined).Execute()

Mark Task As Manually Quarantined



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
	isManuallyQuarantined := true // bool | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).Requester(requester).IsManuallyQuarantined(isManuallyQuarantined).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiMarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requester** | **string** |  | 
 **isManuallyQuarantined** | **bool** |  | 

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


## MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost

> interface{} MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).IsManuallyQuarantined(isManuallyQuarantined).RequestBody(requestBody).Execute()

Mark Tests As Manually Quarantined



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
	isManuallyQuarantined := true // bool | 
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).IsManuallyQuarantined(isManuallyQuarantined).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiMarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **isManuallyQuarantined** | **bool** |  | 
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

