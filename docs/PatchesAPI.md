# \PatchesAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PatchesPatchIdAbortPost**](PatchesAPI.md#PatchesPatchIdAbortPost) | **Post** /patches/{patch_id}/abort | Abort a patch
[**PatchesPatchIdConfigurePost**](PatchesAPI.md#PatchesPatchIdConfigurePost) | **Post** /patches/{patch_id}/configure | Configure/schedule a patch
[**PatchesPatchIdGet**](PatchesAPI.md#PatchesPatchIdGet) | **Get** /patches/{patch_id} | Fetch patch by ID
[**PatchesPatchIdPatch**](PatchesAPI.md#PatchesPatchIdPatch) | **Patch** /patches/{patch_id} | Change patch status
[**PatchesPatchIdRawGet**](PatchesAPI.md#PatchesPatchIdRawGet) | **Get** /patches/{patch_id}/raw | Get patch diff
[**PatchesPatchIdRestartPost**](PatchesAPI.md#PatchesPatchIdRestartPost) | **Post** /patches/{patch_id}/restart | restart a patch
[**ProjectsPatchIdRawModulesGet**](PatchesAPI.md#ProjectsPatchIdRawModulesGet) | **Get** /projects/{patch_id}/raw_modules | Get patch diff with module diffs
[**ProjectsProjectIdPatchesGet**](PatchesAPI.md#ProjectsProjectIdPatchesGet) | **Get** /projects/{project_id}/patches | Fetch patches by project
[**UsersUserIdPatchesGet**](PatchesAPI.md#UsersUserIdPatchesGet) | **Get** /users/{user_id}/patches | Fetch patches by user



## PatchesPatchIdAbortPost

> ModelAPIPatch PatchesPatchIdAbortPost(ctx, patchId).Execute()

Abort a patch



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
	patchId := "patchId_example" // string | the patch ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.PatchesPatchIdAbortPost(context.Background(), patchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.PatchesPatchIdAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchesPatchIdAbortPost`: ModelAPIPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.PatchesPatchIdAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | the patch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchesPatchIdAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIPatch**](ModelAPIPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchesPatchIdConfigurePost

> ModelAPIVersion PatchesPatchIdConfigurePost(ctx, patchId).Object(object).Execute()

Configure/schedule a patch



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
	patchId := "patchId_example" // string | the patch ID
	object := *openapiclient.NewRoutePatchTasks() // RoutePatchTasks | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.PatchesPatchIdConfigurePost(context.Background(), patchId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.PatchesPatchIdConfigurePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchesPatchIdConfigurePost`: ModelAPIVersion
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.PatchesPatchIdConfigurePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | the patch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchesPatchIdConfigurePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RoutePatchTasks**](RoutePatchTasks.md) | parameters | 

### Return type

[**ModelAPIVersion**](ModelAPIVersion.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchesPatchIdGet

> ModelAPIPatch PatchesPatchIdGet(ctx, patchId).Module(module).Execute()

Fetch patch by ID



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
	patchId := "patchId_example" // string | patch ID
	module := "module_example" // string | A module to get the diff for. Returns the empty string when no patch exists for the module. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.PatchesPatchIdGet(context.Background(), patchId).Module(module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.PatchesPatchIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchesPatchIdGet`: ModelAPIPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.PatchesPatchIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | patch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchesPatchIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **module** | **string** | A module to get the diff for. Returns the empty string when no patch exists for the module. | 

### Return type

[**ModelAPIPatch**](ModelAPIPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchesPatchIdPatch

> ModelAPIPatch PatchesPatchIdPatch(ctx, patchId).Object(object).Execute()

Change patch status



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
	patchId := "patchId_example" // string | patch ID
	object := *openapiclient.NewRoutePatchChangeStatusHandler() // RoutePatchChangeStatusHandler | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.PatchesPatchIdPatch(context.Background(), patchId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.PatchesPatchIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchesPatchIdPatch`: ModelAPIPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.PatchesPatchIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | patch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchesPatchIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RoutePatchChangeStatusHandler**](RoutePatchChangeStatusHandler.md) | parameters | 

### Return type

[**ModelAPIPatch**](ModelAPIPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchesPatchIdRawGet

> string PatchesPatchIdRawGet(ctx, patchId).Module(module).Execute()

Get patch diff



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
	patchId := "patchId_example" // string | patch ID
	module := "module_example" // string | A module to get the diff for. Returns the empty string when no patch exists for the module. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.PatchesPatchIdRawGet(context.Background(), patchId).Module(module).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.PatchesPatchIdRawGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchesPatchIdRawGet`: string
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.PatchesPatchIdRawGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | patch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchesPatchIdRawGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **module** | **string** | A module to get the diff for. Returns the empty string when no patch exists for the module. | 

### Return type

**string**

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchesPatchIdRestartPost

> ModelAPIPatch PatchesPatchIdRestartPost(ctx, patchId).Execute()

restart a patch



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
	patchId := "patchId_example" // string | the patch ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.PatchesPatchIdRestartPost(context.Background(), patchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.PatchesPatchIdRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PatchesPatchIdRestartPost`: ModelAPIPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.PatchesPatchIdRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | the patch ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchesPatchIdRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIPatch**](ModelAPIPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsPatchIdRawModulesGet

> ModelAPIRawPatch ProjectsPatchIdRawModulesGet(ctx, patchId).Execute()

Get patch diff with module diffs



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
	patchId := "patchId_example" // string | the project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.ProjectsPatchIdRawModulesGet(context.Background(), patchId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.ProjectsPatchIdRawModulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsPatchIdRawModulesGet`: ModelAPIRawPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.ProjectsPatchIdRawModulesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**patchId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsPatchIdRawModulesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIRawPatch**](ModelAPIRawPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdPatchesGet

> []ModelAPIPatch ProjectsProjectIdPatchesGet(ctx, projectId).StartAt(startAt).Limit(limit).Execute()

Fetch patches by project



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
	projectId := "projectId_example" // string | the project ID
	startAt := "startAt_example" // string | The create_time of the patch to start at in the pagination. Defaults to now (optional)
	limit := int32(56) // int32 | The number of patches to be returned per page of pagination. Defaults to 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.ProjectsProjectIdPatchesGet(context.Background(), projectId).StartAt(startAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.ProjectsProjectIdPatchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdPatchesGet`: []ModelAPIPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.ProjectsProjectIdPatchesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdPatchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **string** | The create_time of the patch to start at in the pagination. Defaults to now | 
 **limit** | **int32** | The number of patches to be returned per page of pagination. Defaults to 100 | 

### Return type

[**[]ModelAPIPatch**](ModelAPIPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdPatchesGet

> []ModelAPIPatch UsersUserIdPatchesGet(ctx, userId, projectId).StartAt(startAt).Limit(limit).Execute()

Fetch patches by user



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
	userId := "userId_example" // string | the user's ID
	projectId := "projectId_example" // string | the project ID
	startAt := "startAt_example" // string | The create_time of the patch to start at in the pagination. Defaults to now (optional)
	limit := int32(56) // int32 | The number of patches to be returned per page of pagination. Defaults to 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PatchesAPI.UsersUserIdPatchesGet(context.Background(), userId, projectId).StartAt(startAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PatchesAPI.UsersUserIdPatchesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdPatchesGet`: []ModelAPIPatch
	fmt.Fprintf(os.Stdout, "Response from `PatchesAPI.UsersUserIdPatchesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID | 
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdPatchesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **string** | The create_time of the patch to start at in the pagination. Defaults to now | 
 **limit** | **int32** | The number of patches to be returned per page of pagination. Defaults to 100 | 

### Return type

[**[]ModelAPIPatch**](ModelAPIPatch.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

