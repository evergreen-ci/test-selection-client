# ModelVariantTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of build variant | [optional] 
**Tasks** | Pointer to **[]string** | All tasks available to run on this build variant | [optional] 

## Methods

### NewModelVariantTask

`func NewModelVariantTask() *ModelVariantTask`

NewModelVariantTask instantiates a new ModelVariantTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVariantTaskWithDefaults

`func NewModelVariantTaskWithDefaults() *ModelVariantTask`

NewModelVariantTaskWithDefaults instantiates a new ModelVariantTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelVariantTask) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelVariantTask) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelVariantTask) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelVariantTask) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTasks

`func (o *ModelVariantTask) GetTasks() []string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ModelVariantTask) GetTasksOk() (*[]string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ModelVariantTask) SetTasks(v []string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ModelVariantTask) HasTasks() bool`

HasTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


