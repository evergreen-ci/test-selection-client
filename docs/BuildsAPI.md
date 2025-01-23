# \BuildsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BuildsBuildIdAbortPost**](BuildsAPI.md#BuildsBuildIdAbortPost) | **Post** /builds/{build_id}/abort | Abort a build
[**BuildsBuildIdGet**](BuildsAPI.md#BuildsBuildIdGet) | **Get** /builds/{build_id} | Fetch build by ID
[**BuildsBuildIdPatch**](BuildsAPI.md#BuildsBuildIdPatch) | **Patch** /builds/{build_id} | Change a build&#39;s execution status
[**BuildsBuildIdRestartPost**](BuildsAPI.md#BuildsBuildIdRestartPost) | **Post** /builds/{build_id}/restart | Restart a build
[**VersionsVersionIdBuildsGet**](BuildsAPI.md#VersionsVersionIdBuildsGet) | **Get** /versions/{version_id}/builds | Get builds from a version



## BuildsBuildIdAbortPost

> ModelAPIBuild BuildsBuildIdAbortPost(ctx, buildId).Execute()

Abort a build



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
	buildId := "buildId_example" // string | build ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBuildIdAbortPost(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBuildIdAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBuildIdAbortPost`: ModelAPIBuild
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBuildIdAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildIdAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIBuild**](ModelAPIBuild.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBuildIdGet

> ModelAPIBuild BuildsBuildIdGet(ctx, buildId).Execute()

Fetch build by ID



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
	buildId := "buildId_example" // string | the build ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBuildIdGet(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBuildIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBuildIdGet`: ModelAPIBuild
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBuildIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | the build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIBuild**](ModelAPIBuild.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBuildIdPatch

> ModelAPIBuild BuildsBuildIdPatch(ctx, buildId).Object(object).Execute()

Change a build's execution status



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
	buildId := "buildId_example" // string | the build ID
	object := *openapiclient.NewRouteBuildChangeStatusHandler() // RouteBuildChangeStatusHandler | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBuildIdPatch(context.Background(), buildId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBuildIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBuildIdPatch`: ModelAPIBuild
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBuildIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | the build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteBuildChangeStatusHandler**](RouteBuildChangeStatusHandler.md) | parameters | 

### Return type

[**ModelAPIBuild**](ModelAPIBuild.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BuildsBuildIdRestartPost

> ModelAPIBuild BuildsBuildIdRestartPost(ctx, buildId).Execute()

Restart a build



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
	buildId := "buildId_example" // string | build ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.BuildsBuildIdRestartPost(context.Background(), buildId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.BuildsBuildIdRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BuildsBuildIdRestartPost`: ModelAPIBuild
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.BuildsBuildIdRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buildId** | **string** | build ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiBuildsBuildIdRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIBuild**](ModelAPIBuild.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsVersionIdBuildsGet

> []ModelAPIBuild VersionsVersionIdBuildsGet(ctx, versionId).Variant(variant).IncludeTaskInfo(includeTaskInfo).Execute()

Get builds from a version



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
	versionId := "versionId_example" // string | the version ID
	variant := "variant_example" // string | Only return the build with this variant (using Distro identifier). (optional)
	includeTaskInfo := "includeTaskInfo_example" // string | if set, include additional information about tasks in each build (this is expensive) (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BuildsAPI.VersionsVersionIdBuildsGet(context.Background(), versionId).Variant(variant).IncludeTaskInfo(includeTaskInfo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BuildsAPI.VersionsVersionIdBuildsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionsVersionIdBuildsGet`: []ModelAPIBuild
	fmt.Fprintf(os.Stdout, "Response from `BuildsAPI.VersionsVersionIdBuildsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | the version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVersionsVersionIdBuildsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variant** | **string** | Only return the build with this variant (using Distro identifier). | 
 **includeTaskInfo** | **string** | if set, include additional information about tasks in each build (this is expensive) | 

### Return type

[**[]ModelAPIBuild**](ModelAPIBuild.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

