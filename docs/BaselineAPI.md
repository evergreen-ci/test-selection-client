# \BaselineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ComputeMetricsApiBaselineComputePost**](BaselineAPI.md#ComputeMetricsApiBaselineComputePost) | **Post** /api/baseline/compute | Compute Metrics
[**GetBaselineStatusApiBaselineStatusGet**](BaselineAPI.md#GetBaselineStatusApiBaselineStatusGet) | **Get** /api/baseline/status | Get Baseline Status



## ComputeMetricsApiBaselineComputePost

> map[string]interface{} ComputeMetricsApiBaselineComputePost(ctx).Project(project).LookbackDays(lookbackDays).Execute()

Compute Metrics



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
	project := "project_example" // string | Optional project filter (optional)
	lookbackDays := int32(56) // int32 | Days of pending comparisons to process (optional) (default to 14)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaselineAPI.ComputeMetricsApiBaselineComputePost(context.Background()).Project(project).LookbackDays(lookbackDays).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaselineAPI.ComputeMetricsApiBaselineComputePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ComputeMetricsApiBaselineComputePost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `BaselineAPI.ComputeMetricsApiBaselineComputePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiComputeMetricsApiBaselineComputePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | **string** | Optional project filter | 
 **lookbackDays** | **int32** | Days of pending comparisons to process | [default to 14]

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBaselineStatusApiBaselineStatusGet

> BaselineComparison GetBaselineStatusApiBaselineStatusGet(ctx).TaskId(taskId).ComparisonId(comparisonId).Execute()

Get Baseline Status



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
	taskId := "taskId_example" // string | Evergreen task ID (optional)
	comparisonId := "comparisonId_example" // string | Baseline comparison UUID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BaselineAPI.GetBaselineStatusApiBaselineStatusGet(context.Background()).TaskId(taskId).ComparisonId(comparisonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BaselineAPI.GetBaselineStatusApiBaselineStatusGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBaselineStatusApiBaselineStatusGet`: BaselineComparison
	fmt.Fprintf(os.Stdout, "Response from `BaselineAPI.GetBaselineStatusApiBaselineStatusGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBaselineStatusApiBaselineStatusGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **taskId** | **string** | Evergreen task ID | 
 **comparisonId** | **string** | Baseline comparison UUID | 

### Return type

[**BaselineComparison**](BaselineComparison.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

