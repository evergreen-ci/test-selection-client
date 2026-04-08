# \BaselineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBaselineStatusApiBaselineStatusGet**](BaselineAPI.md#GetBaselineStatusApiBaselineStatusGet) | **Get** /api/baseline/status | Get Baseline Status



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

