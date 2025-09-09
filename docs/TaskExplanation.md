# TaskExplanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** |  | 
**Tests** | Pointer to [**map[string]TestExplanation**](TestExplanation.md) |  | [optional] [default to {}]

## Methods

### NewTaskExplanation

`func NewTaskExplanation(taskName string, ) *TaskExplanation`

NewTaskExplanation instantiates a new TaskExplanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskExplanationWithDefaults

`func NewTaskExplanationWithDefaults() *TaskExplanation`

NewTaskExplanationWithDefaults instantiates a new TaskExplanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *TaskExplanation) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *TaskExplanation) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *TaskExplanation) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetTests

`func (o *TaskExplanation) GetTests() map[string]TestExplanation`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *TaskExplanation) GetTestsOk() (*map[string]TestExplanation, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *TaskExplanation) SetTests(v map[string]TestExplanation)`

SetTests sets Tests field to given value.

### HasTests

`func (o *TaskExplanation) HasTests() bool`

HasTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


