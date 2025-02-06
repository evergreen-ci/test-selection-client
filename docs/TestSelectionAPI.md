# \TestSelectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet**](TestSelectionAPI.md#ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet) | **Get** /api/test_selection/explain_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/{test_names}/ | Explain Select Tests
[**SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet**](TestSelectionAPI.md#SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet) | **Get** /api/test_selection/select_tests/{project_id}/{requester}/{build_variant_name}/{task_id}/{task_name}/{test_names}/ | Select Tests



## ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet

> []Explanation ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet(ctx, projectId, requester, buildVariantName, taskId, taskName, testNames).Strategies(strategies).Execute()

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
	testNames := "testNames_example" // string | 
	strategies := openapiclient.StrategyEnum("Optimistic") // StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet(context.Background(), projectId, requester, buildVariantName, taskId, taskName, testNames).Strategies(strategies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet`: []Explanation
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.ExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet`: %v\n", resp)
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
**testNames** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExplainSelectTestsApiTestSelectionExplainTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **strategies** | [**StrategyEnum**](StrategyEnum.md) |  | 

### Return type

[**[]Explanation**](Explanation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet

> []string SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet(ctx, projectId, requester, buildVariantName, taskId, taskName, testNames).Strategies(strategies).Execute()

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
	testNames := "testNames_example" // string | 
	strategies := openapiclient.StrategyEnum("Optimistic") // StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet(context.Background(), projectId, requester, buildVariantName, taskId, taskName, testNames).Strategies(strategies).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet`: []string
	fmt.Fprintf(os.Stdout, "Response from `TestSelectionAPI.SelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGet`: %v\n", resp)
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
**testNames** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSelectTestsApiTestSelectionSelectTestsProjectIdRequesterBuildVariantNameTaskIdTaskNameTestNamesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **strategies** | [**StrategyEnum**](StrategyEnum.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

