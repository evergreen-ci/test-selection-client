# RouteTaskRestartHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedOnly** | Pointer to **bool** | If set for a display task, restarts only failed execution tasks. When used with a non-display task, this parameter has no effect. | [optional] 

## Methods

### NewRouteTaskRestartHandler

`func NewRouteTaskRestartHandler() *RouteTaskRestartHandler`

NewRouteTaskRestartHandler instantiates a new RouteTaskRestartHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteTaskRestartHandlerWithDefaults

`func NewRouteTaskRestartHandlerWithDefaults() *RouteTaskRestartHandler`

NewRouteTaskRestartHandlerWithDefaults instantiates a new RouteTaskRestartHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedOnly

`func (o *RouteTaskRestartHandler) GetFailedOnly() bool`

GetFailedOnly returns the FailedOnly field if non-nil, zero value otherwise.

### GetFailedOnlyOk

`func (o *RouteTaskRestartHandler) GetFailedOnlyOk() (*bool, bool)`

GetFailedOnlyOk returns a tuple with the FailedOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedOnly

`func (o *RouteTaskRestartHandler) SetFailedOnly(v bool)`

SetFailedOnly sets FailedOnly field to given value.

### HasFailedOnly

`func (o *RouteTaskRestartHandler) HasFailedOnly() bool`

HasFailedOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


