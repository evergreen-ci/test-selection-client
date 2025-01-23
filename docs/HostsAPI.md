# \HostsAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HostsGet**](HostsAPI.md#HostsGet) | **Get** /hosts | Fetch all hosts
[**HostsHostIdChangePasswordPost**](HostsAPI.md#HostsHostIdChangePasswordPost) | **Post** /hosts/{host_id}/change_password | Change RDP password of a host
[**HostsHostIdExtendExpirationPost**](HostsAPI.md#HostsHostIdExtendExpirationPost) | **Post** /hosts/{host_id}/extend_expiration | Extend the expiration of a host
[**HostsHostIdGet**](HostsAPI.md#HostsHostIdGet) | **Get** /hosts/{host_id} | Fetch host by ID
[**HostsHostIdStartPost**](HostsAPI.md#HostsHostIdStartPost) | **Post** /hosts/{host_id}/start | Start host
[**HostsHostIdStopPost**](HostsAPI.md#HostsHostIdStopPost) | **Post** /hosts/{host_id}/stop | Stop host
[**HostsHostIdTerminatePost**](HostsAPI.md#HostsHostIdTerminatePost) | **Post** /hosts/{host_id}/terminate | Terminate a host
[**HostsPost**](HostsAPI.md#HostsPost) | **Post** /hosts | Spawn a host
[**UsersUserIdHostsGet**](HostsAPI.md#UsersUserIdHostsGet) | **Get** /users/{user_id}/hosts | Fetch hosts spawned by user



## HostsGet

> ModelAPIHost HostsGet(ctx).StartAt(startAt).Limit(limit).Status(status).Execute()

Fetch all hosts



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
	status := "status_example" // string | A status of host to limit the results to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.HostsGet(context.Background()).StartAt(startAt).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostsGet`: ModelAPIHost
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.HostsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The identifier of the host to start at in the pagination | 
 **limit** | **int32** | The number of hosts to be returned per page of pagination. Defaults to 100 | 
 **status** | **string** | A status of host to limit the results to | 

### Return type

[**ModelAPIHost**](ModelAPIHost.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsHostIdChangePasswordPost

> HostsHostIdChangePasswordPost(ctx, hostId).Object(object).Execute()

Change RDP password of a host



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
	hostId := "hostId_example" // string | the host ID
	object := *openapiclient.NewModelAPISpawnHostModify() // ModelAPISpawnHostModify | New RDP password; must meet RDP password criteria as provided by Microsoft at: https://technet.microsoft.com/en-us/library/cc786468(v=ws.10).aspx and be between 6 and 255 characters long

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.HostsHostIdChangePasswordPost(context.Background(), hostId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsHostIdChangePasswordPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | the host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsHostIdChangePasswordPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPISpawnHostModify**](ModelAPISpawnHostModify.md) | New RDP password; must meet RDP password criteria as provided by Microsoft at: https://technet.microsoft.com/en-us/library/cc786468(v&#x3D;ws.10).aspx and be between 6 and 255 characters long | 

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


## HostsHostIdExtendExpirationPost

> HostsHostIdExtendExpirationPost(ctx, hostId).Object(object).Execute()

Extend the expiration of a host



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
	hostId := "hostId_example" // string | the host ID
	object := *openapiclient.NewModelAPISpawnHostModify() // ModelAPISpawnHostModify | Number of hours to extend expiration; not to exceed 168

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.HostsHostIdExtendExpirationPost(context.Background(), hostId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsHostIdExtendExpirationPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | the host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsHostIdExtendExpirationPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**ModelAPISpawnHostModify**](ModelAPISpawnHostModify.md) | Number of hours to extend expiration; not to exceed 168 | 

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


## HostsHostIdGet

> ModelAPIHost HostsHostIdGet(ctx, hostId).Execute()

Fetch host by ID



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
	hostId := "hostId_example" // string | the host ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.HostsHostIdGet(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsHostIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostsHostIdGet`: ModelAPIHost
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.HostsHostIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | the host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsHostIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIHost**](ModelAPIHost.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HostsHostIdStartPost

> HostsHostIdStartPost(ctx, hostId).Object(object).Execute()

Start host



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
	hostId := "hostId_example" // string | the host ID
	object := *openapiclient.NewRouteHostSubscriptionInfo() // RouteHostSubscriptionInfo | subscription_type (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.HostsHostIdStartPost(context.Background(), hostId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsHostIdStartPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | the host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsHostIdStartPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteHostSubscriptionInfo**](RouteHostSubscriptionInfo.md) | subscription_type | 

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


## HostsHostIdStopPost

> HostsHostIdStopPost(ctx, hostId).Object(object).Execute()

Stop host



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
	hostId := "hostId_example" // string | the host ID
	object := *openapiclient.NewRouteHostStopOptions() // RouteHostStopOptions | subscription type and whether an unexpirable host should be kept off (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.HostsHostIdStopPost(context.Background(), hostId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsHostIdStopPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | the host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsHostIdStopPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteHostStopOptions**](RouteHostStopOptions.md) | subscription type and whether an unexpirable host should be kept off | 

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


## HostsHostIdTerminatePost

> HostsHostIdTerminatePost(ctx, hostId).Execute()

Terminate a host



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
	hostId := "hostId_example" // string | the host ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HostsAPI.HostsHostIdTerminatePost(context.Background(), hostId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsHostIdTerminatePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**hostId** | **string** | the host ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiHostsHostIdTerminatePostRequest struct via the builder pattern


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


## HostsPost

> ModelAPIHost HostsPost(ctx).Object(object).Execute()

Spawn a host



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
	object := map[string]interface{}{ ... } // map[string]interface{} | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.HostsPost(context.Background()).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.HostsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HostsPost`: ModelAPIHost
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.HostsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHostsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **object** | **map[string]interface{}** | parameters | 

### Return type

[**ModelAPIHost**](ModelAPIHost.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdHostsGet

> ModelAPIHost UsersUserIdHostsGet(ctx, userId).StartAt(startAt).Limit(limit).Status(status).Execute()

Fetch hosts spawned by user



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
	userId := "userId_example" // string | the user ID
	startAt := "startAt_example" // string | The identifier of the host to start at in the pagination (optional)
	limit := int32(56) // int32 | The number of hosts to be returned per page of pagination. Defaults to 100 (optional)
	status := "status_example" // string | A status of host to limit the results to (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HostsAPI.UsersUserIdHostsGet(context.Background(), userId).StartAt(startAt).Limit(limit).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HostsAPI.UsersUserIdHostsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdHostsGet`: ModelAPIHost
	fmt.Fprintf(os.Stdout, "Response from `HostsAPI.UsersUserIdHostsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdHostsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **string** | The identifier of the host to start at in the pagination | 
 **limit** | **int32** | The number of hosts to be returned per page of pagination. Defaults to 100 | 
 **status** | **string** | A status of host to limit the results to | 

### Return type

[**ModelAPIHost**](ModelAPIHost.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

