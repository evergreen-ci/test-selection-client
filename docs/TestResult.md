# TestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestName** | **string** |  | 
**Status** | **string** |  | 
**TaskId** | **string** |  | 
**Execution** | **int32** |  | 

## Methods

### NewTestResult

`func NewTestResult(testName string, status string, taskId string, execution int32, ) *TestResult`

NewTestResult instantiates a new TestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestResultWithDefaults

`func NewTestResultWithDefaults() *TestResult`

NewTestResultWithDefaults instantiates a new TestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestName

`func (o *TestResult) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestResult) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestResult) SetTestName(v string)`

SetTestName sets TestName field to given value.


### GetStatus

`func (o *TestResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TestResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TestResult) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTaskId

`func (o *TestResult) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TestResult) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TestResult) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetExecution

`func (o *TestResult) GetExecution() int32`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *TestResult) GetExecutionOk() (*int32, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *TestResult) SetExecution(v int32)`

SetExecution sets Execution field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


