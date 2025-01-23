# ModelAPITest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseStatus** | Pointer to **string** |  | [optional] 
**Duration** | Pointer to **float32** |  | [optional] 
**EndTime** | Pointer to **string** | Time that this test stopped execution | [optional] 
**Execution** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Logs** | Pointer to [**ModelTestLogs**](ModelTestLogs.md) | Object containing information about the logs for this test | [optional] 
**StartTime** | Pointer to **string** | Time that this test began execution | [optional] 
**Status** | Pointer to **string** | Execution status of the test | [optional] 
**TaskId** | Pointer to **string** | Identifier of the task this test is a part of | [optional] 
**TestFile** | Pointer to **string** | Name of the test file that this test was run in | [optional] 
**TestId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPITest

`func NewModelAPITest() *ModelAPITest`

NewModelAPITest instantiates a new ModelAPITest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITestWithDefaults

`func NewModelAPITestWithDefaults() *ModelAPITest`

NewModelAPITestWithDefaults instantiates a new ModelAPITest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseStatus

`func (o *ModelAPITest) GetBaseStatus() string`

GetBaseStatus returns the BaseStatus field if non-nil, zero value otherwise.

### GetBaseStatusOk

`func (o *ModelAPITest) GetBaseStatusOk() (*string, bool)`

GetBaseStatusOk returns a tuple with the BaseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseStatus

`func (o *ModelAPITest) SetBaseStatus(v string)`

SetBaseStatus sets BaseStatus field to given value.

### HasBaseStatus

`func (o *ModelAPITest) HasBaseStatus() bool`

HasBaseStatus returns a boolean if a field has been set.

### GetDuration

`func (o *ModelAPITest) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ModelAPITest) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ModelAPITest) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ModelAPITest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndTime

`func (o *ModelAPITest) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ModelAPITest) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ModelAPITest) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *ModelAPITest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetExecution

`func (o *ModelAPITest) GetExecution() int32`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ModelAPITest) GetExecutionOk() (*int32, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ModelAPITest) SetExecution(v int32)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *ModelAPITest) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetGroupId

`func (o *ModelAPITest) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ModelAPITest) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ModelAPITest) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ModelAPITest) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetLogs

`func (o *ModelAPITest) GetLogs() ModelTestLogs`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ModelAPITest) GetLogsOk() (*ModelTestLogs, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ModelAPITest) SetLogs(v ModelTestLogs)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ModelAPITest) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelAPITest) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelAPITest) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelAPITest) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelAPITest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPITest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPITest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPITest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPITest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskId

`func (o *ModelAPITest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelAPITest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelAPITest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelAPITest) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTestFile

`func (o *ModelAPITest) GetTestFile() string`

GetTestFile returns the TestFile field if non-nil, zero value otherwise.

### GetTestFileOk

`func (o *ModelAPITest) GetTestFileOk() (*string, bool)`

GetTestFileOk returns a tuple with the TestFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFile

`func (o *ModelAPITest) SetTestFile(v string)`

SetTestFile sets TestFile field to given value.

### HasTestFile

`func (o *ModelAPITest) HasTestFile() bool`

HasTestFile returns a boolean if a field has been set.

### GetTestId

`func (o *ModelAPITest) GetTestId() string`

GetTestId returns the TestId field if non-nil, zero value otherwise.

### GetTestIdOk

`func (o *ModelAPITest) GetTestIdOk() (*string, bool)`

GetTestIdOk returns a tuple with the TestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestId

`func (o *ModelAPITest) SetTestId(v string)`

SetTestId sets TestId field to given value.

### HasTestId

`func (o *ModelAPITest) HasTestId() bool`

HasTestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


