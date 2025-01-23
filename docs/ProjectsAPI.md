# \ProjectsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AliasProjectIdGet**](ProjectsAPI.md#AliasProjectIdGet) | **Get** /alias/{project_id} | Get a project&#39;s aliases
[**ProjectsGet**](ProjectsAPI.md#ProjectsGet) | **Get** /projects | Fetch all projects
[**ProjectsProjectIdCopyPost**](ProjectsAPI.md#ProjectsProjectIdCopyPost) | **Post** /projects/{project_id}/copy | Copy a project
[**ProjectsProjectIdCopyVariablesPost**](ProjectsAPI.md#ProjectsProjectIdCopyVariablesPost) | **Post** /projects/{project_id}/copy/variables | Copy variables to an existing project
[**ProjectsProjectIdGet**](ProjectsAPI.md#ProjectsProjectIdGet) | **Get** /projects/{project_id} | Get a project
[**ProjectsProjectIdParametersGet**](ProjectsAPI.md#ProjectsProjectIdParametersGet) | **Get** /projects/{project_id}/parameters | Get current parameters for a project
[**ProjectsProjectIdPatch**](ProjectsAPI.md#ProjectsProjectIdPatch) | **Patch** /projects/{project_id} | Modify a project
[**ProjectsProjectIdPut**](ProjectsAPI.md#ProjectsProjectIdPut) | **Put** /projects/{project_id} | Put a project
[**ProjectsTestAliasGet**](ProjectsAPI.md#ProjectsTestAliasGet) | **Get** /projects/test_alias | Check project alias results



## AliasProjectIdGet

> []ModelAPIProjectAlias AliasProjectIdGet(ctx, projectId).IncludeProjectConfig(includeProjectConfig).Execute()

Get a project's aliases



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
	includeProjectConfig := true // bool | Setting to true will return the merged result of the project and the config properties set in the project YAML. Defaults to false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.AliasProjectIdGet(context.Background(), projectId).IncludeProjectConfig(includeProjectConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.AliasProjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AliasProjectIdGet`: []ModelAPIProjectAlias
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.AliasProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAliasProjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeProjectConfig** | **bool** | Setting to true will return the merged result of the project and the config properties set in the project YAML. Defaults to false | 

### Return type

[**[]ModelAPIProjectAlias**](ModelAPIProjectAlias.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsGet

> []ModelAPIProjectRef ProjectsGet(ctx).StartAt(startAt).Limit(limit).Execute()

Fetch all projects



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
	startAt := "startAt_example" // string | The identifier of the host to start at in the pagination (optional)
	limit := int32(56) // int32 | The number of hosts to be returned per page of pagination. Defaults to 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsGet(context.Background()).StartAt(startAt).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsGet`: []ModelAPIProjectRef
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The identifier of the host to start at in the pagination | 
 **limit** | **int32** | The number of hosts to be returned per page of pagination. Defaults to 100 | 

### Return type

[**[]ModelAPIProjectRef**](ModelAPIProjectRef.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdCopyPost

> ModelAPIProjectRef ProjectsProjectIdCopyPost(ctx, projectId).NewProject(newProject).Execute()

Copy a project



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
	newProject := "newProject_example" // string | the new project ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsProjectIdCopyPost(context.Background(), projectId).NewProject(newProject).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsProjectIdCopyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdCopyPost`: ModelAPIProjectRef
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsProjectIdCopyPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdCopyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **newProject** | **string** | the new project ID | 

### Return type

[**ModelAPIProjectRef**](ModelAPIProjectRef.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdCopyVariablesPost

> ProjectsProjectIdCopyVariablesPost(ctx, projectId).Object(object).Execute()

Copy variables to an existing project



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
	object := *openapiclient.NewRouteCopyVariablesOptions() // RouteCopyVariablesOptions | parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectsAPI.ProjectsProjectIdCopyVariablesPost(context.Background(), projectId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsProjectIdCopyVariablesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiProjectsProjectIdCopyVariablesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteCopyVariablesOptions**](RouteCopyVariablesOptions.md) | parameters | 

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


## ProjectsProjectIdGet

> ModelAPIProjectRef ProjectsProjectIdGet(ctx, projectId).IncludeRepo(includeRepo).IncludeProjectConfig(includeProjectConfig).Execute()

Get a project



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
	includeRepo := true // bool | Setting to true will return the merged result of project and repo level settings. Defaults to false (optional)
	includeProjectConfig := true // bool | Setting to true will return the merged result of the project and the config properties set in the project YAML. Defaults to false (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsProjectIdGet(context.Background(), projectId).IncludeRepo(includeRepo).IncludeProjectConfig(includeProjectConfig).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsProjectIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdGet`: ModelAPIProjectRef
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsProjectIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRepo** | **bool** | Setting to true will return the merged result of project and repo level settings. Defaults to false | 
 **includeProjectConfig** | **bool** | Setting to true will return the merged result of the project and the config properties set in the project YAML. Defaults to false | 

### Return type

[**ModelAPIProjectRef**](ModelAPIProjectRef.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdParametersGet

> []ModelAPIParameterInfo ProjectsProjectIdParametersGet(ctx, projectId).Execute()

Get current parameters for a project



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsProjectIdParametersGet(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsProjectIdParametersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdParametersGet`: []ModelAPIParameterInfo
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsProjectIdParametersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdParametersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ModelAPIParameterInfo**](ModelAPIParameterInfo.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdPatch

> ModelAPIProjectRef ProjectsProjectIdPatch(ctx, projectId).Object(object).Execute()

Modify a project



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
	object := *openapiclient.NewModelAPIProjectRef() // ModelAPIProjectRef | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsProjectIdPatch(context.Background(), projectId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsProjectIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdPatch`: ModelAPIProjectRef
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsProjectIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPIProjectRef**](ModelAPIProjectRef.md) | parameters | 

### Return type

[**ModelAPIProjectRef**](ModelAPIProjectRef.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsProjectIdPut

> ModelAPIProjectRef ProjectsProjectIdPut(ctx, projectId).Object(object).Execute()

Put a project



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
	object := *openapiclient.NewModelAPIProjectRef() // ModelAPIProjectRef | parameters (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsProjectIdPut(context.Background(), projectId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsProjectIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsProjectIdPut`: ModelAPIProjectRef
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsProjectIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** | the project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsProjectIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPIProjectRef**](ModelAPIProjectRef.md) | parameters | 

### Return type

[**ModelAPIProjectRef**](ModelAPIProjectRef.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsTestAliasGet

> ModelAPIVariantTasks ProjectsTestAliasGet(ctx).Version(version).Alias(alias).IncludeDeps(includeDeps).Execute()

Check project alias results



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
	version := "version_example" // string | version
	alias := "alias_example" // string | alias
	includeDeps := true // bool | include dependencies (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectsAPI.ProjectsTestAliasGet(context.Background()).Version(version).Alias(alias).IncludeDeps(includeDeps).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectsAPI.ProjectsTestAliasGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProjectsTestAliasGet`: ModelAPIVariantTasks
	fmt.Fprintf(os.Stdout, "Response from `ProjectsAPI.ProjectsTestAliasGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectsTestAliasGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | **string** | version | 
 **alias** | **string** | alias | 
 **includeDeps** | **bool** | include dependencies | 

### Return type

[**ModelAPIVariantTasks**](ModelAPIVariantTasks.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

