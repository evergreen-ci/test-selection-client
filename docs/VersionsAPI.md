# \VersionsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsProjectIdVersionsGet**](VersionsAPI.md#ProjectsProjectIdVersionsGet) | **Get** /projects/{project_id}/versions | Get versions for a project
[**ProjectsProjectIdVersionsPatch**](VersionsAPI.md#ProjectsProjectIdVersionsPatch) | **Patch** /projects/{project_id}/versions | Modify versions for a project
[**VersionsVersionIdAbortPost**](VersionsAPI.md#VersionsVersionIdAbortPost) | **Post** /versions/{version_id}/abort | Abort a version
[**VersionsVersionIdGet**](VersionsAPI.md#VersionsVersionIdGet) | **Get** /versions/{version_id} | Fetch version by ID
[**VersionsVersionIdPatch**](VersionsAPI.md#VersionsVersionIdPatch) | **Patch** /versions/{version_id} | Activate or deactivate a version
[**VersionsVersionIdRestartPost**](VersionsAPI.md#VersionsVersionIdRestartPost) | **Post** /versions/{version_id}/restart | Restart a version



## ProjectsProjectIdVersionsGet

> []ModelAPIVersion ProjectsProjectIdVersionsGet(ctx, projectId).Skip(skip).Limit(limit).Start(start).RevisionEnd(revisionEnd).Requester(requester).IncludeBuilds(includeBuilds).ByBuildVariant(byBuildVariant).IncludeTasks(includeTasks).ByTask(byTask).Execute()

Get versions for a project



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
	skip := int32(56) // int32 | Number of versions to skip. (optional)
	limit := int32(56) // int32 | The number of versions to be returned per page of pagination. Defaults to 20. (optional)
	start := int32(56) // int32 | The version order number to start at, for pagination. Will return the versions that are less than (and therefore older) the revision number specified. (optional)
	revisionEnd := int32(56) // int32 | Will return the versions that are greater than (and therefore more recent) or equal to revision number specified. (optional)
	requester := "requester_example" // string | Returns versions for this requester only. Defaults to gitter_request (caused by git commit, aka the repotracker requester). Can also be set to patch_request, github_pull_request, trigger_request (Project Trigger versions) , github_merge_request (GitHub merge queue),, and ad_hoc (periodic builds). (optional)
	includeBuilds := true // bool | If set, will return some information for each build in the version. (optional)
	byBuildVariant := "byBuildVariant_example" // string | If set, will only include information for this build, and only return versions with this build activated. Must have include_builds set. (optional)
	includeTasks := true // bool | If set, will return some information for each task in the included builds. This is only allowed if include_builds is set. (optional)
	byTask := "byTask_example" // string | If set, will only include information for this task, and will only return versions with this task activated. Must have include_tasks set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.ProjectsProjectIdVersionsGet(context.Background(), projectId).Skip(skip).Limit(limit).Start(start).RevisionEnd(revisionEnd).Requester(requester).IncludeBuilds(includeBuilds).ByBuildVariant(byBuildVariant).IncludeTasks(includeTasks).ByTask(byTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ProjectsProjectIdVersionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdVersionsGet`: []ModelAPIVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.ProjectsProjectIdVersionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdVersionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skip** | **int32** | Number of versions to skip. | 
 **limit** | **int32** | The number of versions to be returned per page of pagination. Defaults to 20. | 
 **start** | **int32** | The version order number to start at, for pagination. Will return the versions that are less than (and therefore older) the revision number specified. | 
 **revisionEnd** | **int32** | Will return the versions that are greater than (and therefore more recent) or equal to revision number specified. | 
 **requester** | **string** | Returns versions for this requester only. Defaults to gitter_request (caused by git commit, aka the repotracker requester). Can also be set to patch_request, github_pull_request, trigger_request (Project Trigger versions) , github_merge_request (GitHub merge queue),, and ad_hoc (periodic builds). | 
 **includeBuilds** | **bool** | If set, will return some information for each build in the version. | 
 **byBuildVariant** | **string** | If set, will only include information for this build, and only return versions with this build activated. Must have include_builds set. | 
 **includeTasks** | **bool** | If set, will return some information for each task in the included builds. This is only allowed if include_builds is set. | 
 **byTask** | **string** | If set, will only include information for this task, and will only return versions with this task activated. Must have include_tasks set. | 

### Return type

[**[]ModelAPIVersion**](ModelAPIVersion.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdVersionsPatch

> ProjectsProjectIdVersionsPatch(ctx, projectId).StartTimeStr(startTimeStr).Priority(priority).EndTimeStr(endTimeStr).RevisionStart(revisionStart).RevisionEnd(revisionEnd).Requester(requester).ByBuildVariant(byBuildVariant).ByTask(byTask).Execute()

Modify versions for a project



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
	startTimeStr := "startTimeStr_example" // string | Timestamp to start looking for applicable versions.
	priority := int32(56) // int32 | Priority to set for all tasks within applicable versions.
	endTimeStr := "endTimeStr_example" // string | Timestamp to stop looking for applicable versions. (optional)
	revisionStart := int32(56) // int32 | The version order number to start at. (optional)
	revisionEnd := int32(56) // int32 | The version order number to end at. (optional)
	requester := "requester_example" // string | Returns versions for this requester only. Defaults to gitter_request (caused by git commit, aka the repotracker requester). Can also be set to patch_request, github_pull_request, trigger_request (Project Trigger versions) , github_merge_request (GitHub merge queue), and ad_hoc (periodic builds). (optional)
	byBuildVariant := "byBuildVariant_example" // string | If set, will only include information for this build, and only return versions with this build activated. Must have include_builds set. (optional)
	byTask := "byTask_example" // string | If set, will only include information for this task, and will only return versions with this task activated. Must have include_tasks set. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VersionsAPI.ProjectsProjectIdVersionsPatch(context.Background(), projectId).StartTimeStr(startTimeStr).Priority(priority).EndTimeStr(endTimeStr).RevisionStart(revisionStart).RevisionEnd(revisionEnd).Requester(requester).ByBuildVariant(byBuildVariant).ByTask(byTask).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.ProjectsProjectIdVersionsPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdVersionsPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startTimeStr** | **string** | Timestamp to start looking for applicable versions. | 
 **priority** | **int32** | Priority to set for all tasks within applicable versions. | 
 **endTimeStr** | **string** | Timestamp to stop looking for applicable versions. | 
 **revisionStart** | **int32** | The version order number to start at. | 
 **revisionEnd** | **int32** | The version order number to end at. | 
 **requester** | **string** | Returns versions for this requester only. Defaults to gitter_request (caused by git commit, aka the repotracker requester). Can also be set to patch_request, github_pull_request, trigger_request (Project Trigger versions) , github_merge_request (GitHub merge queue), and ad_hoc (periodic builds). | 
 **byBuildVariant** | **string** | If set, will only include information for this build, and only return versions with this build activated. Must have include_builds set. | 
 **byTask** | **string** | If set, will only include information for this task, and will only return versions with this task activated. Must have include_tasks set. | 

### Return type

 (empty response body)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsVersionIdAbortPost

> ModelAPIVersion VersionsVersionIdAbortPost(ctx, versionId).Execute()

Abort a version



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
	versionId := "versionId_example" // string | version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.VersionsVersionIdAbortPost(context.Background(), versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.VersionsVersionIdAbortPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionsVersionIdAbortPost`: ModelAPIVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.VersionsVersionIdAbortPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVersionsVersionIdAbortPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIVersion**](ModelAPIVersion.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsVersionIdGet

> ModelAPIVersion VersionsVersionIdGet(ctx, versionId).Execute()

Fetch version by ID



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
	versionId := "versionId_example" // string | version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.VersionsVersionIdGet(context.Background(), versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.VersionsVersionIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionsVersionIdGet`: ModelAPIVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.VersionsVersionIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVersionsVersionIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIVersion**](ModelAPIVersion.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsVersionIdPatch

> VersionsVersionIdPatch(ctx, versionId).Object(object).Execute()

Activate or deactivate a version



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
	object := *openapiclient.NewRouteVersionPatchHandler(false) // RouteVersionPatchHandler | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VersionsAPI.VersionsVersionIdPatch(context.Background(), versionId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.VersionsVersionIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | the version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVersionsVersionIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteVersionPatchHandler**](RouteVersionPatchHandler.md) | parameters | 

### Return type

 (empty response body)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionsVersionIdRestartPost

> ModelAPIVersion VersionsVersionIdRestartPost(ctx, versionId).Execute()

Restart a version



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
	versionId := "versionId_example" // string | version ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VersionsAPI.VersionsVersionIdRestartPost(context.Background(), versionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VersionsAPI.VersionsVersionIdRestartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VersionsVersionIdRestartPost`: ModelAPIVersion
	fmt.Fprintf(os.Stdout, "Response from `VersionsAPI.VersionsVersionIdRestartPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | version ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiVersionsVersionIdRestartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIVersion**](ModelAPIVersion.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

