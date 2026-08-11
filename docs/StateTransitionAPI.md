# \StateTransitionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost**](StateTransitionAPI.md#GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost) | **Post** /api/test_selection/get_project_quarantined/{project_id}/ | Get Project Quarantined
[**GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost**](StateTransitionAPI.md#GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost) | **Post** /api/test_selection/get_project_quarantined/ | Get Project Quarantined Data In Body
[**GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/get_task_quarantined/{project_id}/{build_variant_name}/{task_name}/ | Get Task Quarantined
[**GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost**](StateTransitionAPI.md#GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost) | **Post** /api/test_selection/get_task_quarantined/ | Get Task Quarantined Data In Body
[**GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#GetTestsStateApiTestSelectionGetTestsStateProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/get_tests_state/{project_id}/{build_variant_name}/{task_name}/ | Get Tests State
[**GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost**](StateTransitionAPI.md#GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost) | **Post** /api/test_selection/get_tests_state/ | Get Tests State Data In Body
[**GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost**](StateTransitionAPI.md#GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost) | **Post** /api/test_selection/get_variant_quarantined/{project_id}/{build_variant_name}/ | Get Variant Quarantined
[**GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost**](StateTransitionAPI.md#GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost) | **Post** /api/test_selection/get_variant_quarantined/ | Get Variant Quarantined Data In Body
[**GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost**](StateTransitionAPI.md#GetVariantStateApiTestSelectionGetVariantStateProjectIdBuildVariantNamePost) | **Post** /api/test_selection/get_variant_state/{project_id}/{build_variant_name}/ | Get Variant State
[**GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost**](StateTransitionAPI.md#GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost) | **Post** /api/test_selection/get_variant_state/ | Get Variant State Data In Body
[**MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTaskAsManuallyQuarantinedApiTestSelectionTransitionTaskProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_task/{project_id}/{build_variant_name}/{task_name}/ | Mark Task As Manually Quarantined
[**MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost**](StateTransitionAPI.md#MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost) | **Post** /api/test_selection/transition_task/ | Mark Task As Manually Quarantined Data In Body
[**MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost**](StateTransitionAPI.md#MarkTestsAsManuallyQuarantinedApiTestSelectionTransitionTestsProjectIdBuildVariantNameTaskNamePost) | **Post** /api/test_selection/transition_tests/{project_id}/{build_variant_name}/{task_name}/ | Mark Tests As Manually Quarantined
[**MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost**](StateTransitionAPI.md#MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost) | **Post** /api/test_selection/transition_tests/ | Mark Tests As Manually Quarantined Data In Body
[**MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost**](StateTransitionAPI.md#MarkVariantAsManuallyQuarantinedApiTestSelectionTransitionVariantProjectIdBuildVariantNamePost) | **Post** /api/test_selection/transition_variant/{project_id}/{build_variant_name}/ | Mark Variant As Manually Quarantined
[**MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost**](StateTransitionAPI.md#MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost) | **Post** /api/test_selection/transition_variant/ | Mark Variant As Manually Quarantined Data In Body



## GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost

> map[string]VariantStateInfo GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost(ctx, projectId).Execute()

Get Project Quarantined



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost`: map[string]VariantStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectQuarantinedApiTestSelectionGetProjectQuarantinedProjectIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]VariantStateInfo**](VariantStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost

> map[string]VariantStateInfo GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost(ctx).BodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost(bodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost).Execute()

Get Project Quarantined Data In Body



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
	bodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost := *openapiclient.NewBodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost("ProjectId_example") // BodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost(context.Background()).BodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost(bodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost`: map[string]VariantStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost** | [**BodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost**](BodyGetProjectQuarantinedDataInBodyApiTestSelectionGetProjectQuarantinedPost.md) |  | 

### Return type

[**map[string]VariantStateInfo**](VariantStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost

> TaskStateInfo GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost(ctx, projectId, buildVariantName, taskName).Execute()

Get Task Quarantined



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost(context.Background(), projectId, buildVariantName, taskName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost`: TaskStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePost`: %v\n", resp)
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

Other parameters are passed through a pointer to a apiGetTaskQuarantinedApiTestSelectionGetTaskQuarantinedProjectIdBuildVariantNameTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**TaskStateInfo**](TaskStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost

> TaskStateInfo GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost(ctx).BodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost(bodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost).Execute()

Get Task Quarantined Data In Body



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
	bodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost := *openapiclient.NewBodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost("ProjectId_example", "BuildVariantName_example", "TaskName_example") // BodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost(context.Background()).BodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost(bodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost`: TaskStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost** | [**BodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost**](BodyGetTaskQuarantinedDataInBodyApiTestSelectionGetTaskQuarantinedPost.md) |  | 

### Return type

[**TaskStateInfo**](TaskStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost

> map[string]TestStateInfo GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost(ctx).BodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost(bodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost).Execute()

Get Tests State Data In Body



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
	bodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost := *openapiclient.NewBodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost("ProjectId_example", "BuildVariantName_example", "TaskName_example", []string{"TestNames_example"}) // BodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost(context.Background()).BodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost(bodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost`: map[string]TestStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetTestsStateDataInBodyApiTestSelectionGetTestsStatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTestsStateDataInBodyApiTestSelectionGetTestsStatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost** | [**BodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost**](BodyGetTestsStateDataInBodyApiTestSelectionGetTestsStatePost.md) |  | 

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


## GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost

> map[string]TaskStateInfo GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost(ctx, projectId, buildVariantName).Execute()

Get Variant Quarantined



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
	resp, r, err := apiClient.StateTransitionAPI.GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost(context.Background(), projectId, buildVariantName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost`: map[string]TaskStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**buildVariantName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantQuarantinedApiTestSelectionGetVariantQuarantinedProjectIdBuildVariantNamePostRequest struct via the builder pattern


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


## GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost

> map[string]TaskStateInfo GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost(ctx).BodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost(bodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost).Execute()

Get Variant Quarantined Data In Body



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
	bodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost := *openapiclient.NewBodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost("ProjectId_example", "BuildVariantName_example") // BodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost(context.Background()).BodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost(bodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost`: map[string]TaskStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost** | [**BodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost**](BodyGetVariantQuarantinedDataInBodyApiTestSelectionGetVariantQuarantinedPost.md) |  | 

### Return type

[**map[string]TaskStateInfo**](TaskStateInfo.md)

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


## GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost

> map[string]TaskStateInfo GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost(ctx).BodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost(bodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost).Execute()

Get Variant State Data In Body



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
	bodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost := *openapiclient.NewBodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost("ProjectId_example", "BuildVariantName_example") // BodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost(context.Background()).BodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost(bodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost`: map[string]TaskStateInfo
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.GetVariantStateDataInBodyApiTestSelectionGetVariantStatePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetVariantStateDataInBodyApiTestSelectionGetVariantStatePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost** | [**BodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost**](BodyGetVariantStateDataInBodyApiTestSelectionGetVariantStatePost.md) |  | 

### Return type

[**map[string]TaskStateInfo**](TaskStateInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
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


## MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost

> interface{} MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost(ctx).BodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost(bodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost).Execute()

Mark Task As Manually Quarantined Data In Body



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
	bodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost := *openapiclient.NewBodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost("ProjectId_example", "BuildVariantName_example", "TaskName_example", false) // BodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost(context.Background()).BodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost(bodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost** | [**BodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost**](BodyMarkTaskAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTaskPost.md) |  | 

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


## MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost

> interface{} MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost(ctx).BodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost(bodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost).Execute()

Mark Tests As Manually Quarantined Data In Body



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
	bodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost := *openapiclient.NewBodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost("ProjectId_example", "BuildVariantName_example", "TaskName_example", []string{"TestNames_example"}, false) // BodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost(context.Background()).BodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost(bodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost** | [**BodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost**](BodyMarkTestsAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionTestsPost.md) |  | 

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


## MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost

> interface{} MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost(ctx).BodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost(bodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost).Execute()

Mark Variant As Manually Quarantined Data In Body



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
	bodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost := *openapiclient.NewBodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost("ProjectId_example", "BuildVariantName_example", false) // BodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StateTransitionAPI.MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost(context.Background()).BodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost(bodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StateTransitionAPI.MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StateTransitionAPI.MarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost** | [**BodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost**](BodyMarkVariantAsManuallyQuarantinedDataInBodyApiTestSelectionTransitionVariantPost.md) |  | 

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

