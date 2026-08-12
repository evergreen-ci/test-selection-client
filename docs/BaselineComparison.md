# BaselineComparison

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComparisonId** | **string** |  | 
**Project** | **string** |  | 
**StrategyPatchId** | Pointer to **NullableString** |  | [optional] 
**TaskId** | **string** |  | 
**ExecutionNumber** | Pointer to **int32** |  | [optional] [default to 0]
**BaseCommitSha** | Pointer to **NullableString** |  | [optional] 
**BaselineVersionId** | Pointer to **NullableString** |  | [optional] 
**SubsequentCommitSha** | Pointer to **NullableString** |  | [optional] 
**SubsequentTaskId** | Pointer to **NullableString** |  | [optional] 
**BaselineTaskId** | Pointer to **NullableString** |  | [optional] 
**BaselineOrder** | Pointer to **NullableInt32** |  | [optional] 
**SubsequentOrder** | Pointer to **NullableInt32** |  | [optional] 
**Status** | [**ComparisonStatus**](ComparisonStatus.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**Strategies** | **[]string** |  | 
**StrategyExecutionOrder** | [**[]StrategyEnum**](StrategyEnum.md) |  | 
**TaskMetrics** | Pointer to [**[]TaskMetrics**](TaskMetrics.md) |  | [optional] 
**StrategyMetrics** | Pointer to [**map[string]StrategyMetrics**](StrategyMetrics.md) |  | [optional] 
**RetroactiveCorrections** | Pointer to **[]string** |  | [optional] 
**TotalTestsAcrossAllTasks** | Pointer to **int32** |  | [optional] [default to 0]
**TestNamesHashes** | Pointer to **[]string** |  | [optional] 

## Methods

### NewBaselineComparison

`func NewBaselineComparison(comparisonId string, project string, taskId string, status ComparisonStatus, createdAt time.Time, updatedAt time.Time, strategies []string, strategyExecutionOrder []StrategyEnum, ) *BaselineComparison`

NewBaselineComparison instantiates a new BaselineComparison object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaselineComparisonWithDefaults

`func NewBaselineComparisonWithDefaults() *BaselineComparison`

NewBaselineComparisonWithDefaults instantiates a new BaselineComparison object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComparisonId

`func (o *BaselineComparison) GetComparisonId() string`

GetComparisonId returns the ComparisonId field if non-nil, zero value otherwise.

### GetComparisonIdOk

`func (o *BaselineComparison) GetComparisonIdOk() (*string, bool)`

GetComparisonIdOk returns a tuple with the ComparisonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparisonId

`func (o *BaselineComparison) SetComparisonId(v string)`

SetComparisonId sets ComparisonId field to given value.


### GetProject

`func (o *BaselineComparison) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *BaselineComparison) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *BaselineComparison) SetProject(v string)`

SetProject sets Project field to given value.


### GetStrategyPatchId

`func (o *BaselineComparison) GetStrategyPatchId() string`

GetStrategyPatchId returns the StrategyPatchId field if non-nil, zero value otherwise.

### GetStrategyPatchIdOk

`func (o *BaselineComparison) GetStrategyPatchIdOk() (*string, bool)`

GetStrategyPatchIdOk returns a tuple with the StrategyPatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyPatchId

`func (o *BaselineComparison) SetStrategyPatchId(v string)`

SetStrategyPatchId sets StrategyPatchId field to given value.

### HasStrategyPatchId

`func (o *BaselineComparison) HasStrategyPatchId() bool`

HasStrategyPatchId returns a boolean if a field has been set.

### SetStrategyPatchIdNil

`func (o *BaselineComparison) SetStrategyPatchIdNil(b bool)`

 SetStrategyPatchIdNil sets the value for StrategyPatchId to be an explicit nil

### UnsetStrategyPatchId
`func (o *BaselineComparison) UnsetStrategyPatchId()`

UnsetStrategyPatchId ensures that no value is present for StrategyPatchId, not even an explicit nil
### GetTaskId

`func (o *BaselineComparison) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BaselineComparison) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BaselineComparison) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetExecutionNumber

`func (o *BaselineComparison) GetExecutionNumber() int32`

GetExecutionNumber returns the ExecutionNumber field if non-nil, zero value otherwise.

### GetExecutionNumberOk

`func (o *BaselineComparison) GetExecutionNumberOk() (*int32, bool)`

GetExecutionNumberOk returns a tuple with the ExecutionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionNumber

`func (o *BaselineComparison) SetExecutionNumber(v int32)`

SetExecutionNumber sets ExecutionNumber field to given value.

### HasExecutionNumber

`func (o *BaselineComparison) HasExecutionNumber() bool`

HasExecutionNumber returns a boolean if a field has been set.

### GetBaseCommitSha

`func (o *BaselineComparison) GetBaseCommitSha() string`

GetBaseCommitSha returns the BaseCommitSha field if non-nil, zero value otherwise.

### GetBaseCommitShaOk

`func (o *BaselineComparison) GetBaseCommitShaOk() (*string, bool)`

GetBaseCommitShaOk returns a tuple with the BaseCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCommitSha

`func (o *BaselineComparison) SetBaseCommitSha(v string)`

SetBaseCommitSha sets BaseCommitSha field to given value.

### HasBaseCommitSha

`func (o *BaselineComparison) HasBaseCommitSha() bool`

HasBaseCommitSha returns a boolean if a field has been set.

### SetBaseCommitShaNil

`func (o *BaselineComparison) SetBaseCommitShaNil(b bool)`

 SetBaseCommitShaNil sets the value for BaseCommitSha to be an explicit nil

### UnsetBaseCommitSha
`func (o *BaselineComparison) UnsetBaseCommitSha()`

UnsetBaseCommitSha ensures that no value is present for BaseCommitSha, not even an explicit nil
### GetBaselineVersionId

`func (o *BaselineComparison) GetBaselineVersionId() string`

GetBaselineVersionId returns the BaselineVersionId field if non-nil, zero value otherwise.

### GetBaselineVersionIdOk

`func (o *BaselineComparison) GetBaselineVersionIdOk() (*string, bool)`

GetBaselineVersionIdOk returns a tuple with the BaselineVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineVersionId

`func (o *BaselineComparison) SetBaselineVersionId(v string)`

SetBaselineVersionId sets BaselineVersionId field to given value.

### HasBaselineVersionId

`func (o *BaselineComparison) HasBaselineVersionId() bool`

HasBaselineVersionId returns a boolean if a field has been set.

### SetBaselineVersionIdNil

`func (o *BaselineComparison) SetBaselineVersionIdNil(b bool)`

 SetBaselineVersionIdNil sets the value for BaselineVersionId to be an explicit nil

### UnsetBaselineVersionId
`func (o *BaselineComparison) UnsetBaselineVersionId()`

UnsetBaselineVersionId ensures that no value is present for BaselineVersionId, not even an explicit nil
### GetSubsequentCommitSha

`func (o *BaselineComparison) GetSubsequentCommitSha() string`

GetSubsequentCommitSha returns the SubsequentCommitSha field if non-nil, zero value otherwise.

### GetSubsequentCommitShaOk

`func (o *BaselineComparison) GetSubsequentCommitShaOk() (*string, bool)`

GetSubsequentCommitShaOk returns a tuple with the SubsequentCommitSha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsequentCommitSha

`func (o *BaselineComparison) SetSubsequentCommitSha(v string)`

SetSubsequentCommitSha sets SubsequentCommitSha field to given value.

### HasSubsequentCommitSha

`func (o *BaselineComparison) HasSubsequentCommitSha() bool`

HasSubsequentCommitSha returns a boolean if a field has been set.

### SetSubsequentCommitShaNil

`func (o *BaselineComparison) SetSubsequentCommitShaNil(b bool)`

 SetSubsequentCommitShaNil sets the value for SubsequentCommitSha to be an explicit nil

### UnsetSubsequentCommitSha
`func (o *BaselineComparison) UnsetSubsequentCommitSha()`

UnsetSubsequentCommitSha ensures that no value is present for SubsequentCommitSha, not even an explicit nil
### GetSubsequentTaskId

`func (o *BaselineComparison) GetSubsequentTaskId() string`

GetSubsequentTaskId returns the SubsequentTaskId field if non-nil, zero value otherwise.

### GetSubsequentTaskIdOk

`func (o *BaselineComparison) GetSubsequentTaskIdOk() (*string, bool)`

GetSubsequentTaskIdOk returns a tuple with the SubsequentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsequentTaskId

`func (o *BaselineComparison) SetSubsequentTaskId(v string)`

SetSubsequentTaskId sets SubsequentTaskId field to given value.

### HasSubsequentTaskId

`func (o *BaselineComparison) HasSubsequentTaskId() bool`

HasSubsequentTaskId returns a boolean if a field has been set.

### SetSubsequentTaskIdNil

`func (o *BaselineComparison) SetSubsequentTaskIdNil(b bool)`

 SetSubsequentTaskIdNil sets the value for SubsequentTaskId to be an explicit nil

### UnsetSubsequentTaskId
`func (o *BaselineComparison) UnsetSubsequentTaskId()`

UnsetSubsequentTaskId ensures that no value is present for SubsequentTaskId, not even an explicit nil
### GetBaselineTaskId

`func (o *BaselineComparison) GetBaselineTaskId() string`

GetBaselineTaskId returns the BaselineTaskId field if non-nil, zero value otherwise.

### GetBaselineTaskIdOk

`func (o *BaselineComparison) GetBaselineTaskIdOk() (*string, bool)`

GetBaselineTaskIdOk returns a tuple with the BaselineTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTaskId

`func (o *BaselineComparison) SetBaselineTaskId(v string)`

SetBaselineTaskId sets BaselineTaskId field to given value.

### HasBaselineTaskId

`func (o *BaselineComparison) HasBaselineTaskId() bool`

HasBaselineTaskId returns a boolean if a field has been set.

### SetBaselineTaskIdNil

`func (o *BaselineComparison) SetBaselineTaskIdNil(b bool)`

 SetBaselineTaskIdNil sets the value for BaselineTaskId to be an explicit nil

### UnsetBaselineTaskId
`func (o *BaselineComparison) UnsetBaselineTaskId()`

UnsetBaselineTaskId ensures that no value is present for BaselineTaskId, not even an explicit nil
### GetBaselineOrder

`func (o *BaselineComparison) GetBaselineOrder() int32`

GetBaselineOrder returns the BaselineOrder field if non-nil, zero value otherwise.

### GetBaselineOrderOk

`func (o *BaselineComparison) GetBaselineOrderOk() (*int32, bool)`

GetBaselineOrderOk returns a tuple with the BaselineOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineOrder

`func (o *BaselineComparison) SetBaselineOrder(v int32)`

SetBaselineOrder sets BaselineOrder field to given value.

### HasBaselineOrder

`func (o *BaselineComparison) HasBaselineOrder() bool`

HasBaselineOrder returns a boolean if a field has been set.

### SetBaselineOrderNil

`func (o *BaselineComparison) SetBaselineOrderNil(b bool)`

 SetBaselineOrderNil sets the value for BaselineOrder to be an explicit nil

### UnsetBaselineOrder
`func (o *BaselineComparison) UnsetBaselineOrder()`

UnsetBaselineOrder ensures that no value is present for BaselineOrder, not even an explicit nil
### GetSubsequentOrder

`func (o *BaselineComparison) GetSubsequentOrder() int32`

GetSubsequentOrder returns the SubsequentOrder field if non-nil, zero value otherwise.

### GetSubsequentOrderOk

`func (o *BaselineComparison) GetSubsequentOrderOk() (*int32, bool)`

GetSubsequentOrderOk returns a tuple with the SubsequentOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsequentOrder

`func (o *BaselineComparison) SetSubsequentOrder(v int32)`

SetSubsequentOrder sets SubsequentOrder field to given value.

### HasSubsequentOrder

`func (o *BaselineComparison) HasSubsequentOrder() bool`

HasSubsequentOrder returns a boolean if a field has been set.

### SetSubsequentOrderNil

`func (o *BaselineComparison) SetSubsequentOrderNil(b bool)`

 SetSubsequentOrderNil sets the value for SubsequentOrder to be an explicit nil

### UnsetSubsequentOrder
`func (o *BaselineComparison) UnsetSubsequentOrder()`

UnsetSubsequentOrder ensures that no value is present for SubsequentOrder, not even an explicit nil
### GetStatus

`func (o *BaselineComparison) GetStatus() ComparisonStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaselineComparison) GetStatusOk() (*ComparisonStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaselineComparison) SetStatus(v ComparisonStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BaselineComparison) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BaselineComparison) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BaselineComparison) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BaselineComparison) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BaselineComparison) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BaselineComparison) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCompletedAt

`func (o *BaselineComparison) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *BaselineComparison) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *BaselineComparison) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *BaselineComparison) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *BaselineComparison) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *BaselineComparison) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetStrategies

`func (o *BaselineComparison) GetStrategies() []string`

GetStrategies returns the Strategies field if non-nil, zero value otherwise.

### GetStrategiesOk

`func (o *BaselineComparison) GetStrategiesOk() (*[]string, bool)`

GetStrategiesOk returns a tuple with the Strategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategies

`func (o *BaselineComparison) SetStrategies(v []string)`

SetStrategies sets Strategies field to given value.


### GetStrategyExecutionOrder

`func (o *BaselineComparison) GetStrategyExecutionOrder() []StrategyEnum`

GetStrategyExecutionOrder returns the StrategyExecutionOrder field if non-nil, zero value otherwise.

### GetStrategyExecutionOrderOk

`func (o *BaselineComparison) GetStrategyExecutionOrderOk() (*[]StrategyEnum, bool)`

GetStrategyExecutionOrderOk returns a tuple with the StrategyExecutionOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyExecutionOrder

`func (o *BaselineComparison) SetStrategyExecutionOrder(v []StrategyEnum)`

SetStrategyExecutionOrder sets StrategyExecutionOrder field to given value.


### GetTaskMetrics

`func (o *BaselineComparison) GetTaskMetrics() []TaskMetrics`

GetTaskMetrics returns the TaskMetrics field if non-nil, zero value otherwise.

### GetTaskMetricsOk

`func (o *BaselineComparison) GetTaskMetricsOk() (*[]TaskMetrics, bool)`

GetTaskMetricsOk returns a tuple with the TaskMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskMetrics

`func (o *BaselineComparison) SetTaskMetrics(v []TaskMetrics)`

SetTaskMetrics sets TaskMetrics field to given value.

### HasTaskMetrics

`func (o *BaselineComparison) HasTaskMetrics() bool`

HasTaskMetrics returns a boolean if a field has been set.

### GetStrategyMetrics

`func (o *BaselineComparison) GetStrategyMetrics() map[string]StrategyMetrics`

GetStrategyMetrics returns the StrategyMetrics field if non-nil, zero value otherwise.

### GetStrategyMetricsOk

`func (o *BaselineComparison) GetStrategyMetricsOk() (*map[string]StrategyMetrics, bool)`

GetStrategyMetricsOk returns a tuple with the StrategyMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyMetrics

`func (o *BaselineComparison) SetStrategyMetrics(v map[string]StrategyMetrics)`

SetStrategyMetrics sets StrategyMetrics field to given value.

### HasStrategyMetrics

`func (o *BaselineComparison) HasStrategyMetrics() bool`

HasStrategyMetrics returns a boolean if a field has been set.

### GetRetroactiveCorrections

`func (o *BaselineComparison) GetRetroactiveCorrections() []string`

GetRetroactiveCorrections returns the RetroactiveCorrections field if non-nil, zero value otherwise.

### GetRetroactiveCorrectionsOk

`func (o *BaselineComparison) GetRetroactiveCorrectionsOk() (*[]string, bool)`

GetRetroactiveCorrectionsOk returns a tuple with the RetroactiveCorrections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetroactiveCorrections

`func (o *BaselineComparison) SetRetroactiveCorrections(v []string)`

SetRetroactiveCorrections sets RetroactiveCorrections field to given value.

### HasRetroactiveCorrections

`func (o *BaselineComparison) HasRetroactiveCorrections() bool`

HasRetroactiveCorrections returns a boolean if a field has been set.

### GetTotalTestsAcrossAllTasks

`func (o *BaselineComparison) GetTotalTestsAcrossAllTasks() int32`

GetTotalTestsAcrossAllTasks returns the TotalTestsAcrossAllTasks field if non-nil, zero value otherwise.

### GetTotalTestsAcrossAllTasksOk

`func (o *BaselineComparison) GetTotalTestsAcrossAllTasksOk() (*int32, bool)`

GetTotalTestsAcrossAllTasksOk returns a tuple with the TotalTestsAcrossAllTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTestsAcrossAllTasks

`func (o *BaselineComparison) SetTotalTestsAcrossAllTasks(v int32)`

SetTotalTestsAcrossAllTasks sets TotalTestsAcrossAllTasks field to given value.

### HasTotalTestsAcrossAllTasks

`func (o *BaselineComparison) HasTotalTestsAcrossAllTasks() bool`

HasTotalTestsAcrossAllTasks returns a boolean if a field has been set.

### GetTestNamesHashes

`func (o *BaselineComparison) GetTestNamesHashes() []string`

GetTestNamesHashes returns the TestNamesHashes field if non-nil, zero value otherwise.

### GetTestNamesHashesOk

`func (o *BaselineComparison) GetTestNamesHashesOk() (*[]string, bool)`

GetTestNamesHashesOk returns a tuple with the TestNamesHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestNamesHashes

`func (o *BaselineComparison) SetTestNamesHashes(v []string)`

SetTestNamesHashes sets TestNamesHashes field to given value.

### HasTestNamesHashes

`func (o *BaselineComparison) HasTestNamesHashes() bool`

HasTestNamesHashes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


