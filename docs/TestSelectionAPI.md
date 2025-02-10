# \TestSelectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](TestSelectionAPI.md#ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) | **Post** /api/test_selection/explain_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/ | Explain Select Tests
[**SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost**](TestSelectionAPI.md#SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost) | **Post** /api/test_selection/select_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/ | Select Tests



## ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost

> []Explanation ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx, projectId, requester, buildVariantName, taskId, taskName).RequestBody(requestBody).Strategies(strategies).Execute()

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
	requestBody := []string{"Property_example"} // []string | 
	strategies := openapiclient.StrategyEnum("Optimistic") // StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(context.Background(), projectId, requester, buildVariantName, taskId, taskName).RequestBody(requestBody).Strategies(strategies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: []Explanation
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





 **requestBody** | **[]string** |  | 
 **strategies** | [**StrategyEnum**](StrategyEnum.md) |  | 

### Return type

[**[]Explanation**](Explanation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost

> []string SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(ctx, projectId, requester, buildVariantName, taskId, taskName).RequestBody(requestBody).Strategies(strategies).Execute()

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
	requestBody := []string{"Property_example"} // []string | 
	strategies := openapiclient.StrategyEnum("Optimistic") // StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost(context.Background(), projectId, requester, buildVariantName, taskId, taskName).RequestBody(requestBody).Strategies(strategies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNamePost`: []string
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





 **requestBody** | **[]string** |  | 
 **strategies** | [**StrategyEnum**](StrategyEnum.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

