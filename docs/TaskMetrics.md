# TaskMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskName** | **string** |  | 
**Variant** | **string** |  | 
**TotalTests** | **int32** |  | 
**TestsSelected** | **int32** |  | 
**TestsFiltered** | **int32** |  | 
**TestMetrics** | [**[]StrategyTestMetrics**](StrategyTestMetrics.md) |  | 

## Methods

### NewTaskMetrics

`func NewTaskMetrics(taskName string, variant string, totalTests int32, testsSelected int32, testsFiltered int32, testMetrics []StrategyTestMetrics, ) *TaskMetrics`

NewTaskMetrics instantiates a new TaskMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskMetricsWithDefaults

`func NewTaskMetricsWithDefaults() *TaskMetrics`

NewTaskMetricsWithDefaults instantiates a new TaskMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskName

`func (o *TaskMetrics) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *TaskMetrics) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *TaskMetrics) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetVariant

`func (o *TaskMetrics) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *TaskMetrics) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *TaskMetrics) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetTotalTests

`func (o *TaskMetrics) GetTotalTests() int32`

GetTotalTests returns the TotalTests field if non-nil, zero value otherwise.

### GetTotalTestsOk

`func (o *TaskMetrics) GetTotalTestsOk() (*int32, bool)`

GetTotalTestsOk returns a tuple with the TotalTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTests

`func (o *TaskMetrics) SetTotalTests(v int32)`

SetTotalTests sets TotalTests field to given value.


### GetTestsSelected

`func (o *TaskMetrics) GetTestsSelected() int32`

GetTestsSelected returns the TestsSelected field if non-nil, zero value otherwise.

### GetTestsSelectedOk

`func (o *TaskMetrics) GetTestsSelectedOk() (*int32, bool)`

GetTestsSelectedOk returns a tuple with the TestsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsSelected

`func (o *TaskMetrics) SetTestsSelected(v int32)`

SetTestsSelected sets TestsSelected field to given value.


### GetTestsFiltered

`func (o *TaskMetrics) GetTestsFiltered() int32`

GetTestsFiltered returns the TestsFiltered field if non-nil, zero value otherwise.

### GetTestsFilteredOk

`func (o *TaskMetrics) GetTestsFilteredOk() (*int32, bool)`

GetTestsFilteredOk returns a tuple with the TestsFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsFiltered

`func (o *TaskMetrics) SetTestsFiltered(v int32)`

SetTestsFiltered sets TestsFiltered field to given value.


### GetTestMetrics

`func (o *TaskMetrics) GetTestMetrics() []StrategyTestMetrics`

GetTestMetrics returns the TestMetrics field if non-nil, zero value otherwise.

### GetTestMetricsOk

`func (o *TaskMetrics) GetTestMetricsOk() (*[]StrategyTestMetrics, bool)`

GetTestMetricsOk returns a tuple with the TestMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestMetrics

`func (o *TaskMetrics) SetTestMetrics(v []StrategyTestMetrics)`

SetTestMetrics sets TestMetrics field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


