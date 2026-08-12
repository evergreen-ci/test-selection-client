# \TestSelectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](TestSelectionAPI.md#ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) | **Post** /api/test_selection/explain_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/ | Explain Select Tests
[**ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost**](TestSelectionAPI.md#ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost) | **Post** /api/test_selection/explain_tests/ | Explain Select Tests With Data In Body
[**SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](TestSelectionAPI.md#SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) | **Post** /api/test_selection/select_known_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/ | Select All Known Tests Of A Task
[**SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost**](TestSelectionAPI.md#SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost) | **Post** /api/test_selection/select_known_tests/ | Select All Known Tests Of A Task With Data In Body
[**SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](TestSelectionAPI.md#SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) | **Post** /api/test_selection/select_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/ | Select Tests
[**SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost**](TestSelectionAPI.md#SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost) | **Post** /api/test_selection/select_tests/ | Select Tests With Data In Body



## ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost

> map[string]map[string]Explanation ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx, projectId, requester, buildVariantName, taskId, taskName).BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost).Execute()

Explain Select Tests



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
	requester := "requester_example" // string | 
	buildVariantName := "buildVariantName_example" // string | 
	taskId := "taskId_example" // string | 
	taskName := "taskName_example" // string | 
	bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost := *openapiclient.NewBodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost([]string{"TestNames_example"}) // BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(context.Background(), projectId, requester, buildVariantName, taskId, taskName).BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: map[string]map[string]Explanation
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**requester** | **string** |  | 
**buildVariantName** | **string** |  | 
**taskId** | **string** |  | 
**taskName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost** | [**BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](BodyExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost.md) |  | 

### Return type

[**map[string]map[string]Explanation**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost

> map[string]map[string]Explanation ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost(ctx).BodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost(bodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost).Execute()

Explain Select Tests With Data In Body



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
	bodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost := *openapiclient.NewBodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost("ProjectId_example", "Requester_example", "BuildVariantName_example", "TaskId_example", "TaskName_example", []string{"TestNames_example"}) // BodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost(context.Background()).BodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost(bodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost`: map[string]map[string]Explanation
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.ExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost** | [**BodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost**](BodyExplainSelectTestsWithDataInBodyApiTestSelectionExplainTestsPost.md) |  | 

### Return type

[**map[string]map[string]Explanation**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost

> []*string SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx, projectId, requester, buildVariantName, taskId, taskName).StrategyEnum(strategyEnum).Execute()

Select All Known Tests Of A Task



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
	requester := "requester_example" // string | 
	buildVariantName := "buildVariantName_example" // string | 
	taskId := "taskId_example" // string | 
	taskName := "taskName_example" // string | 
	strategyEnum := []openapiclient.StrategyEnum{openapiclient.StrategyEnum("ExcludeManuallyQuarantined")} // []StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(context.Background(), projectId, requester, buildVariantName, taskId, taskName).StrategyEnum(strategyEnum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: []*string
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.SelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**requester** | **string** |  | 
**buildVariantName** | **string** |  | 
**taskId** | **string** |  | 
**taskName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectAllKnownTestsOfATaskApiTestSelectionSelectKnownTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **strategyEnum** | [**[]StrategyEnum**](StrategyEnum.md) |  | 

### Return type

**[]*string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost

> []*string SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost(ctx).BodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost(bodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost).Execute()

Select All Known Tests Of A Task With Data In Body



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
	bodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost := *openapiclient.NewBodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost("ProjectId_example", "Requester_example", "BuildVariantName_example", "TaskId_example", "TaskName_example") // BodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost(context.Background()).BodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost(bodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost`: []*string
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.SelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost** | [**BodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost**](BodySelectAllKnownTestsOfATaskWithDataInBodyApiTestSelectionSelectKnownTestsPost.md) |  | 

### Return type

**[]*string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost

> []*string SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx, projectId, requester, buildVariantName, taskId, taskName).BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost).Execute()

Select Tests



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
	requester := "requester_example" // string | 
	buildVariantName := "buildVariantName_example" // string | 
	taskId := "taskId_example" // string | 
	taskName := "taskName_example" // string | 
	bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost := *openapiclient.NewBodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost([]string{"TestNames_example"}) // BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(context.Background(), projectId, requester, buildVariantName, taskId, taskName).BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: []*string
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 
**requester** | **string** |  | 
**buildVariantName** | **string** |  | 
**taskId** | **string** |  | 
**taskName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **bodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost** | [**BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](BodySelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost.md) |  | 

### Return type

**[]*string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost

> []*string SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost(ctx).BodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost(bodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost).Execute()

Select Tests With Data In Body



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
	bodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost := *openapiclient.NewBodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost("ProjectId_example", "Requester_example", "BuildVariantName_example", "TaskId_example", "TaskName_example", []string{"TestNames_example"}) // BodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost(context.Background()).BodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost(bodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost`: []*string
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.SelectTestsWithDataInBodyApiTestSelectionSelectTestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectTestsWithDataInBodyApiTestSelectionSelectTestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost** | [**BodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost**](BodySelectTestsWithDataInBodyApiTestSelectionSelectTestsPost.md) |  | 

### Return type

**[]*string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

