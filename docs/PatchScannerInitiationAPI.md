# \PatchScannerInitiationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ScanPatchesApiTestSelectionPatchScannerProjectIdPost**](PatchScannerInitiationAPI.md#ScanPatchesApiTestSelectionPatchScannerProjectIdPost) | **Post** /api/test_selection/patch_scanner/{project_id} | Scan Patches
[**ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost**](PatchScannerInitiationAPI.md#ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost) | **Post** /api/test_selection/single_patch_scanner/{patch_id} | Scan Single Patch



## ScanPatchesApiTestSelectionPatchScannerProjectIdPost

> interface{} ScanPatchesApiTestSelectionPatchScannerProjectIdPost(ctx, projectId).StrategyEnum(strategyEnum).Execute()

Scan Patches



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
	strategyEnum := []openapiclient.StrategyEnum{openapiclient.StrategyEnum("Existential")} // []StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchScannerInitiationAPI.ScanPatchesApiTestSelectionPatchScannerProjectIdPost(context.Background(), projectId).StrategyEnum(strategyEnum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchScannerInitiationAPI.ScanPatchesApiTestSelectionPatchScannerProjectIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanPatchesApiTestSelectionPatchScannerProjectIdPost`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PatchScannerInitiationAPI.ScanPatchesApiTestSelectionPatchScannerProjectIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanPatchesApiTestSelectionPatchScannerProjectIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **strategyEnum** | [**[]StrategyEnum**](StrategyEnum.md) |  | 

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


## ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost

> PatchExplanation ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost(ctx, patchId).StrategyEnum(strategyEnum).Execute()

Scan Single Patch



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
	patchId := "patchId_example" // string | 
	strategyEnum := []openapiclient.StrategyEnum{openapiclient.StrategyEnum("Existential")} // []StrategyEnum |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchScannerInitiationAPI.ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost(context.Background(), patchId).StrategyEnum(strategyEnum).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchScannerInitiationAPI.ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost`: PatchExplanation
	fmt.Fprintf(os.Stdout, "Response from `PatchScannerInitiationAPI.ScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanSinglePatchApiTestSelectionSinglePatchScannerPatchIdPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **strategyEnum** | [**[]StrategyEnum**](StrategyEnum.md) |  | 

### Return type

[**PatchExplanation**](PatchExplanation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

