# RoutePatchTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Optional, if sent will update the patch&#39;s description | [optional] 
**Variants** | Pointer to [**[]RouteVariant**](RouteVariant.md) | Required, these are the variants and tasks that the patch should run. For an already-scheduled patch, any new tasks in this array will be created and any existing tasks not in this array will be unscheduled. | [optional] 

## Methods

### NewRoutePatchTasks

`func NewRoutePatchTasks() *RoutePatchTasks`

NewRoutePatchTasks instantiates a new RoutePatchTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutePatchTasksWithDefaults

`func NewRoutePatchTasksWithDefaults() *RoutePatchTasks`

NewRoutePatchTasksWithDefaults instantiates a new RoutePatchTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *RoutePatchTasks) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoutePatchTasks) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoutePatchTasks) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoutePatchTasks) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariants

`func (o *RoutePatchTasks) GetVariants() []RouteVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *RoutePatchTasks) GetVariantsOk() (*[]RouteVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *RoutePatchTasks) SetVariants(v []RouteVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *RoutePatchTasks) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


