# \InfoAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminBannerGet**](InfoAPI.md#AdminBannerGet) | **Get** /admin/banner | Get the banner
[**PermissionsGet**](InfoAPI.md#PermissionsGet) | **Get** /permissions | Get permissions
[**StatusCliVersionGet**](InfoAPI.md#StatusCliVersionGet) | **Get** /status/cli_version | Fetch CLI client version



## AdminBannerGet

> ModelAPIBanner AdminBannerGet(ctx).Execute()

Get the banner



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
	resp, r, err := apiClient.InfoAPI.AdminBannerGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.AdminBannerGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AdminBannerGet`: ModelAPIBanner
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.AdminBannerGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminBannerGetRequest struct via the builder pattern


### Return type

[**ModelAPIBanner**](ModelAPIBanner.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PermissionsGet

> ModelAPIPermissions PermissionsGet(ctx).Execute()

Get permissions



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
	resp, r, err := apiClient.InfoAPI.PermissionsGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.PermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsGet`: ModelAPIPermissions
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.PermissionsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsGetRequest struct via the builder pattern


### Return type

[**ModelAPIPermissions**](ModelAPIPermissions.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatusCliVersionGet

> ModelAPICLIUpdate StatusCliVersionGet(ctx).Execute()

Fetch CLI client version



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
	resp, r, err := apiClient.InfoAPI.StatusCliVersionGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InfoAPI.StatusCliVersionGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatusCliVersionGet`: ModelAPICLIUpdate
	fmt.Fprintf(os.Stdout, "Response from `InfoAPI.StatusCliVersionGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatusCliVersionGetRequest struct via the builder pattern


### Return type

[**ModelAPICLIUpdate**](ModelAPICLIUpdate.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

