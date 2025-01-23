# ModelDownstreamTasks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** |  | [optional] 
**Tasks** | Pointer to **[]string** |  | [optional] 
**VariantTasks** | Pointer to [**[]ModelVariantTask**](ModelVariantTask.md) |  | [optional] 

## Methods

### NewModelDownstreamTasks

`func NewModelDownstreamTasks() *ModelDownstreamTasks`

NewModelDownstreamTasks instantiates a new ModelDownstreamTasks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDownstreamTasksWithDefaults

`func NewModelDownstreamTasksWithDefaults() *ModelDownstreamTasks`

NewModelDownstreamTasksWithDefaults instantiates a new ModelDownstreamTasks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ModelDownstreamTasks) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ModelDownstreamTasks) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ModelDownstreamTasks) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ModelDownstreamTasks) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetTasks

`func (o *ModelDownstreamTasks) GetTasks() []string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ModelDownstreamTasks) GetTasksOk() (*[]string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ModelDownstreamTasks) SetTasks(v []string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ModelDownstreamTasks) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetVariantTasks

`func (o *ModelDownstreamTasks) GetVariantTasks() []ModelVariantTask`

GetVariantTasks returns the VariantTasks field if non-nil, zero value otherwise.

### GetVariantTasksOk

`func (o *ModelDownstreamTasks) GetVariantTasksOk() (*[]ModelVariantTask, bool)`

GetVariantTasksOk returns a tuple with the VariantTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantTasks

`func (o *ModelDownstreamTasks) SetVariantTasks(v []ModelVariantTask)`

SetVariantTasks sets VariantTasks field to given value.

### HasVariantTasks

`func (o *ModelDownstreamTasks) HasVariantTasks() bool`

HasVariantTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


