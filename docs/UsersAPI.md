# \UsersAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PermissionsUsersGet**](UsersAPI.md#PermissionsUsersGet) | **Get** /permissions/users | Get all user permissions for resource
[**RolesRoleIdUsersGet**](UsersAPI.md#RolesRoleIdUsersGet) | **Get** /roles/{role_id}/users | Get users for role
[**UsersOffboardUserPost**](UsersAPI.md#UsersOffboardUserPost) | **Post** /users/offboard_user | Offboard user
[**UsersRenameUserPost**](UsersAPI.md#UsersRenameUserPost) | **Post** /users/rename_user | Rename user
[**UsersUserIdGet**](UsersAPI.md#UsersUserIdGet) | **Get** /users/{user_id} | Get user
[**UsersUserIdPermissionsDelete**](UsersAPI.md#UsersUserIdPermissionsDelete) | **Delete** /users/{user_id}/permissions | Delete user permissions
[**UsersUserIdPermissionsGet**](UsersAPI.md#UsersUserIdPermissionsGet) | **Get** /users/{user_id}/permissions | Get user permissions
[**UsersUserIdPermissionsPost**](UsersAPI.md#UsersUserIdPermissionsPost) | **Post** /users/{user_id}/permissions | Give permissions to user
[**UsersUserIdRolesPost**](UsersAPI.md#UsersUserIdRolesPost) | **Post** /users/{user_id}/roles | Give roles to user



## PermissionsUsersGet

> map[string]map[string]int32 PermissionsUsersGet(ctx).Object(object).Execute()

Get all user permissions for resource



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
	object := *openapiclient.NewRouteUsersPermissionsInput() // RouteUsersPermissionsInput | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.PermissionsUsersGet(context.Background()).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.PermissionsUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PermissionsUsersGet`: map[string]map[string]int32
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.PermissionsUsersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPermissionsUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **object** | [**RouteUsersPermissionsInput**](RouteUsersPermissionsInput.md) | parameters | 

### Return type

[**map[string]map[string]int32**](map.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RolesRoleIdUsersGet

> RouteUsersWithRoleResponse RolesRoleIdUsersGet(ctx, roleId).Execute()

Get users for role



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
	roleId := "roleId_example" // string | role ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.RolesRoleIdUsersGet(context.Background(), roleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RolesRoleIdUsersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RolesRoleIdUsersGet`: RouteUsersWithRoleResponse
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.RolesRoleIdUsersGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roleId** | **string** | role ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRolesRoleIdUsersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RouteUsersWithRoleResponse**](RouteUsersWithRoleResponse.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersOffboardUserPost

> ModelAPIOffboardUserResults UsersOffboardUserPost(ctx).Object(object).DryRun(dryRun).Execute()

Offboard user



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
	object := *openapiclient.NewRouteOffboardUserEmail("Email_example") // RouteOffboardUserEmail | parameters
	dryRun := true // bool | If set to true, route returns the IDs of the hosts/volumes that *would* be modified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersOffboardUserPost(context.Background()).Object(object).DryRun(dryRun).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersOffboardUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersOffboardUserPost`: ModelAPIOffboardUserResults
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersOffboardUserPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersOffboardUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **object** | [**RouteOffboardUserEmail**](RouteOffboardUserEmail.md) | parameters | 
 **dryRun** | **bool** | If set to true, route returns the IDs of the hosts/volumes that *would* be modified. | 

### Return type

[**ModelAPIOffboardUserResults**](ModelAPIOffboardUserResults.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRenameUserPost

> UsersRenameUserPost(ctx).Object(object).Execute()

Rename user



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
	object := *openapiclient.NewRouteRenameUserInfo("Email_example", "NewEmail_example") // RouteRenameUserInfo | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersRenameUserPost(context.Background()).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersRenameUserPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersRenameUserPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **object** | [**RouteRenameUserInfo**](RouteRenameUserInfo.md) | parameters | 

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


## UsersUserIdGet

> ModelAPIDBUser UsersUserIdGet(ctx, userId).Execute()

Get user



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
	userId := "userId_example" // string | User ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdGet(context.Background(), userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdGet`: ModelAPIDBUser
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | User ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelAPIDBUser**](ModelAPIDBUser.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdPermissionsDelete

> UsersUserIdPermissionsDelete(ctx, userId).Object(object).Execute()

Delete user permissions



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
	object := *openapiclient.NewRouteDeletePermissionsRequest() // RouteDeletePermissionsRequest | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUserIdPermissionsDelete(context.Background(), userId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdPermissionsDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdPermissionsDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteDeletePermissionsRequest**](RouteDeletePermissionsRequest.md) | parameters | 

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


## UsersUserIdPermissionsGet

> []RouteSwaggerPermissionSummary UsersUserIdPermissionsGet(ctx, userId).All(all).Execute()

Get user permissions



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
	all := true // bool | If included, we will not filter out basic permissions (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UsersAPI.UsersUserIdPermissionsGet(context.Background(), userId).All(all).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdPermissionsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UsersUserIdPermissionsGet`: []RouteSwaggerPermissionSummary
	fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UsersUserIdPermissionsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdPermissionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **all** | **bool** | If included, we will not filter out basic permissions | 

### Return type

[**[]RouteSwaggerPermissionSummary**](RouteSwaggerPermissionSummary.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersUserIdPermissionsPost

> UsersUserIdPermissionsPost(ctx, userId).Object(object).Execute()

Give permissions to user



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
	object := map[string]interface{}{ ... } // map[string]interface{} | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUserIdPermissionsPost(context.Background(), userId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdPermissionsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdPermissionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | **map[string]interface{}** | parameters | 

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


## UsersUserIdRolesPost

> UsersUserIdRolesPost(ctx, userId).Object(object).Execute()

Give roles to user



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
	userId := "userId_example" // string | user ID
	object := *openapiclient.NewRouteRolesPostRequest() // RouteRolesPostRequest | parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UsersAPI.UsersUserIdRolesPost(context.Background(), userId).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UsersUserIdRolesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | user ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersUserIdRolesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **object** | [**RouteRolesPostRequest**](RouteRolesPostRequest.md) | parameters | 

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

