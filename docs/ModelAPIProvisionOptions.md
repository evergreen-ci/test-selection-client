# ModelAPIProvisionOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** | ID of the task that the host was spawned from. | [optional] 

## Methods

### NewModelAPIProvisionOptions

`func NewModelAPIProvisionOptions() *ModelAPIProvisionOptions`

NewModelAPIProvisionOptions instantiates a new ModelAPIProvisionOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIProvisionOptionsWithDefaults

`func NewModelAPIProvisionOptionsWithDefaults() *ModelAPIProvisionOptions`

NewModelAPIProvisionOptionsWithDefaults instantiates a new ModelAPIProvisionOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ModelAPIProvisionOptions) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelAPIProvisionOptions) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelAPIProvisionOptions) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelAPIProvisionOptions) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


