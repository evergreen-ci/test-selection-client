# \StateTransitionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/get_tests_state/{project_id}/{build_variant_name}/{task_name}/ | Get Tests State
[**GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost**](StateTransitionAPI.md#GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost) | **Post** /api/test_selection/get_variant_state/{project_id}/{build_variant_name}/ | Get Variant State
[**MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_task/{project_id}/{build_variant_name}/{task_name}/ | Mark Task As Manually Quarantined
[**MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_tests/{project_id}/{build_variant_name}/{task_name}/ | Mark Tests As Manually Quarantined
[**MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost**](StateTransitionAPI.md#MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost) | **Post** /api/test_selection/transition_variant/{project_id}/{build_variant_name}/ | Mark Variant As Manually Quarantined



## GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost

> map[string]TestStateInfo GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).RequestBody(requestBody).Execute()

Get Tests State



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
	requestBody := []string{"Property_example"} // []string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost`: map[string]TestStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **requestBody** | **[]string** |  | 

### Return type

[**map[string]TestStateInfo**](TestStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost

> map[string]TaskStateInfo GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost(ctx, projectId, buildVariantName).Execute()

Get Variant State



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost(context.Background(), projectId, buildVariantName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost`: map[string]TaskStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**buildVariantName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**map[string]TaskStateInfo**](TaskStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost

> interface{} MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).IsManuallyQuarantined(isManuallyQuarantined).Execute()

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
	isManuallyQuarantined := true // bool | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).IsManuallyQuarantined(isManuallyQuarantined).Execute()
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
	requestBody := []*string{"Property_example"} // []*string | 

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


## MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost

> interface{} MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost(ctx, projectId, buildVariantName).IsManuallyQuarantined(isManuallyQuarantined).Execute()

Mark Variant As Manually Quarantined



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
	isManuallyQuarantined := true // bool | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost(context.Background(), projectId, buildVariantName).IsManuallyQuarantined(isManuallyQuarantined).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**buildVariantName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

