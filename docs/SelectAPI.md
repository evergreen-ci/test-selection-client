# \SelectAPI

All URIs are relative to *https://evergreen.mongodb.com/rest/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SelectTestsPost**](SelectAPI.md#SelectTestsPost) | **Post** /select/tests | Select tests



## SelectTestsPost

> RouteSelectTestsRequest SelectTestsPost(ctx).Object(object).Execute()

Select tests



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
	object := *openapiclient.NewRouteSelectTestsRequest() // RouteSelectTestsRequest | Select tests request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SelectAPI.SelectTestsPost(context.Background()).Object(object).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SelectAPI.SelectTestsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectTestsPost`: RouteSelectTestsRequest
	fmt.Fprintf(os.Stdout, "Response from `SelectAPI.SelectTestsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectTestsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **object** | [**RouteSelectTestsRequest**](RouteSelectTestsRequest.md) | Select tests request | 

### Return type

[**RouteSelectTestsRequest**](RouteSelectTestsRequest.md)

### Authorization

[Api-User](../README.md#Api-User), [Api-Key](../README.md#Api-Key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

