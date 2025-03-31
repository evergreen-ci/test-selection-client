# BuildExplanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildVariant** | **string** | Build variant name (not display name) | 
**Tasks** | Pointer to [**map[string]TaskExplanation**](TaskExplanation.md) |  | [optional] [default to {}]

## Methods

### NewBuildExplanation

`func NewBuildExplanation(buildVariant string, ) *BuildExplanation`

NewBuildExplanation instantiates a new BuildExplanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildExplanationWithDefaults

`func NewBuildExplanationWithDefaults() *BuildExplanation`

NewBuildExplanationWithDefaults instantiates a new BuildExplanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildVariant

`func (o *BuildExplanation) GetBuildVariant() string`

GetBuildVariant returns the BuildVariant field if non-nil, zero value otherwise.

### GetBuildVariantOk

`func (o *BuildExplanation) GetBuildVariantOk() (*string, bool)`

GetBuildVariantOk returns a tuple with the BuildVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVariant

`func (o *BuildExplanation) SetBuildVariant(v string)`

SetBuildVariant sets BuildVariant field to given value.


### GetTasks

`func (o *BuildExplanation) GetTasks() map[string]TaskExplanation`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *BuildExplanation) GetTasksOk() (*map[string]TaskExplanation, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *BuildExplanation) SetTasks(v map[string]TaskExplanation)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *BuildExplanation) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


