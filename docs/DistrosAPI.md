# \DistrosAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DistrosDistroIdDelete**](DistrosAPI.md#DistrosDistroIdDelete) | **Delete** /distros/{distro_id} | Delete a distro
[**DistrosDistroIdGet**](DistrosAPI.md#DistrosDistroIdGet) | **Get** /distros/{distro_id} | Get a single distro
[**DistrosDistroIdPatch**](DistrosAPI.md#DistrosDistroIdPatch) | **Patch** /distros/{distro_id} | Update an existing distro
[**DistrosDistroIdPut**](DistrosAPI.md#DistrosDistroIdPut) | **Put** /distros/{distro_id} | Replace or add a distro
[**DistrosDistroIdSetupGet**](DistrosAPI.md#DistrosDistroIdSetupGet) | **Get** /distros/{distro_id}/setup | Get the setup script for a distro
[**DistrosDistroIdSetupPatch**](DistrosAPI.md#DistrosDistroIdSetupPatch) | **Patch** /distros/{distro_id}/setup | Update the setup script for a distro
[**DistrosGet**](DistrosAPI.md#DistrosGet) | **Get** /distros | Get all distros



## DistrosDistroIdDelete

> DistrosDistroIdDelete(ctx, distroId).Execute()

Delete a distro



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
	distroId := "distroId_example" // string | distro ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DistrosAPI.DistrosDistroIdDelete(context.Background(), distroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosDistroIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroId** | **string** | distro ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosDistroIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DistrosDistroIdGet

> ModelAPIDistro DistrosDistroIdGet(ctx, distroId).Execute()

Get a single distro



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
	distroId := "distroId_example" // string | distro ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistrosAPI.DistrosDistroIdGet(context.Background(), distroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosDistroIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistrosDistroIdGet`: ModelAPIDistro
	fmt.Fprintf(os.Stdout, "Response from `DistrosAPI.DistrosDistroIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroId** | **string** | distro ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosDistroIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIDistro**](ModelAPIDistro.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistrosDistroIdPatch

> ModelAPIDistro DistrosDistroIdPatch(ctx, distroId).Object(object).Execute()

Update an existing distro



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
	distroId := "distroId_example" // string | distro ID
	object := *openapiclient.NewModelAPIDistro() // ModelAPIDistro | the updated distro

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistrosAPI.DistrosDistroIdPatch(context.Background(), distroId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosDistroIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistrosDistroIdPatch`: ModelAPIDistro
	fmt.Fprintf(os.Stdout, "Response from `DistrosAPI.DistrosDistroIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroId** | **string** | distro ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosDistroIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPIDistro**](ModelAPIDistro.md) | the updated distro | 

### Return type

[**ModelAPIDistro**](ModelAPIDistro.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistrosDistroIdPut

> DistrosDistroIdPut(ctx, distroId).Object(object).Execute()

Replace or add a distro



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
	distroId := "distroId_example" // string | distro ID
	object := *openapiclient.NewModelAPIDistro() // ModelAPIDistro | the new distro

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DistrosAPI.DistrosDistroIdPut(context.Background(), distroId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosDistroIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroId** | **string** | distro ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosDistroIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPIDistro**](ModelAPIDistro.md) | the new distro | 

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


## DistrosDistroIdSetupGet

> string DistrosDistroIdSetupGet(ctx, distroId).Execute()

Get the setup script for a distro



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
	distroId := "distroId_example" // string | distro ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistrosAPI.DistrosDistroIdSetupGet(context.Background(), distroId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosDistroIdSetupGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistrosDistroIdSetupGet`: string
	fmt.Fprintf(os.Stdout, "Response from `DistrosAPI.DistrosDistroIdSetupGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroId** | **string** | distro ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosDistroIdSetupGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DistrosDistroIdSetupPatch

> ModelAPIDistro DistrosDistroIdSetupPatch(ctx, distroId).Object(object).Execute()

Update the setup script for a distro



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
	distroId := "distroId_example" // string | distro ID
	object := *openapiclient.NewRouteDistroIDChangeSetupHandler() // RouteDistroIDChangeSetupHandler | the updated setup script

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistrosAPI.DistrosDistroIdSetupPatch(context.Background(), distroId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosDistroIdSetupPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistrosDistroIdSetupPatch`: ModelAPIDistro
	fmt.Fprintf(os.Stdout, "Response from `DistrosAPI.DistrosDistroIdSetupPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**distroId** | **string** | distro ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosDistroIdSetupPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteDistroIDChangeSetupHandler**](RouteDistroIDChangeSetupHandler.md) | the updated setup script | 

### Return type

[**ModelAPIDistro**](ModelAPIDistro.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DistrosGet

> []ModelAPIDistro DistrosGet(ctx).Execute()

Get all distros



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DistrosAPI.DistrosGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DistrosAPI.DistrosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DistrosGet`: []ModelAPIDistro
	fmt.Fprintf(os.Stdout, "Response from `DistrosAPI.DistrosGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDistrosGetRequest struct via the builder pattern


### Return type

[**[]ModelAPIDistro**](ModelAPIDistro.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

