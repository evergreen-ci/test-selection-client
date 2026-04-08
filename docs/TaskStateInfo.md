# TaskStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** |  | 
**TestStats** | [**map[string]TestStateInfo**](TestStateInfo.md) |  | 

## Methods

### NewTaskStateInfo

`func NewTaskStateInfo(taskName string, testStats map[string]TestStateInfo, ) *TaskStateInfo`

NewTaskStateInfo instantiates a new TaskStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStateInfoWithDefaults

`func NewTaskStateInfoWithDefaults() *TaskStateInfo`

NewTaskStateInfoWithDefaults instantiates a new TaskStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *TaskStateInfo) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *TaskStateInfo) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *TaskStateInfo) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetTestStats

`func (o *TaskStateInfo) GetTestStats() map[string]TestStateInfo`

GetTestStats returns the TestStats field if non-nil, zero value otherwise.

### GetTestStatsOk

`func (o *TaskStateInfo) GetTestStatsOk() (*map[string]TestStateInfo, bool)`

GetTestStatsOk returns a tuple with the TestStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStats

`func (o *TaskStateInfo) SetTestStats(v map[string]TestStateInfo)`

SetTestStats sets TestStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


