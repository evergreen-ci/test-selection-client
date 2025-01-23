# RoutePatchChangeStatusHandler

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** | The priority to set the patch to | [optional] 
**Priority** | Pointer to **int32** | The activation status to set the patch to | [optional] 

## Methods

### NewRoutePatchChangeStatusHandler

`func NewRoutePatchChangeStatusHandler() *RoutePatchChangeStatusHandler`

NewRoutePatchChangeStatusHandler instantiates a new RoutePatchChangeStatusHandler object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutePatchChangeStatusHandlerWithDefaults

`func NewRoutePatchChangeStatusHandlerWithDefaults() *RoutePatchChangeStatusHandler`

NewRoutePatchChangeStatusHandlerWithDefaults instantiates a new RoutePatchChangeStatusHandler object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *RoutePatchChangeStatusHandler) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *RoutePatchChangeStatusHandler) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *RoutePatchChangeStatusHandler) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *RoutePatchChangeStatusHandler) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetPriority

`func (o *RoutePatchChangeStatusHandler) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *RoutePatchChangeStatusHandler) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *RoutePatchChangeStatusHandler) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *RoutePatchChangeStatusHandler) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


