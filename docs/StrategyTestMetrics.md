# StrategyTestMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestName** | **string** |  | 
**WasFiltered** | **bool** |  | 
**StrategyName** | [**StrategyEnum**](StrategyEnum.md) |  | 
**ExpectedOutcome** | **string** |  | 
**Explanation** | **string** |  | 
**BaselineResult** | Pointer to [**NullableTestResult**](TestResult.md) |  | [optional] 
**StrategyResult** | Pointer to [**NullableTestResult**](TestResult.md) |  | [optional] 
**SubsequentResult** | Pointer to [**NullableTestResult**](TestResult.md) |  | [optional] 
**Classification** | [**TestClassification**](TestClassification.md) |  | 
**KnownIssues** | Pointer to [**[]KnownIssue**](KnownIssue.md) |  | [optional] 
**IsKnownIssue** | **bool** |  | 

## Methods

### NewStrategyTestMetrics

`func NewStrategyTestMetrics(testName string, wasFiltered bool, strategyName StrategyEnum, expectedOutcome string, explanation string, classification TestClassification, isKnownIssue bool, ) *StrategyTestMetrics`

NewStrategyTestMetrics instantiates a new StrategyTestMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyTestMetricsWithDefaults

`func NewStrategyTestMetricsWithDefaults() *StrategyTestMetrics`

NewStrategyTestMetricsWithDefaults instantiates a new StrategyTestMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestName

`func (o *StrategyTestMetrics) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *StrategyTestMetrics) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *StrategyTestMetrics) SetTestName(v string)`

SetTestName sets TestName field to given value.


### GetWasFiltered

`func (o *StrategyTestMetrics) GetWasFiltered() bool`

GetWasFiltered returns the WasFiltered field if non-nil, zero value otherwise.

### GetWasFilteredOk

`func (o *StrategyTestMetrics) GetWasFilteredOk() (*bool, bool)`

GetWasFilteredOk returns a tuple with the WasFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWasFiltered

`func (o *StrategyTestMetrics) SetWasFiltered(v bool)`

SetWasFiltered sets WasFiltered field to given value.


### GetStrategyName

`func (o *StrategyTestMetrics) GetStrategyName() StrategyEnum`

GetStrategyName returns the StrategyName field if non-nil, zero value otherwise.

### GetStrategyNameOk

`func (o *StrategyTestMetrics) GetStrategyNameOk() (*StrategyEnum, bool)`

GetStrategyNameOk returns a tuple with the StrategyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyName

`func (o *StrategyTestMetrics) SetStrategyName(v StrategyEnum)`

SetStrategyName sets StrategyName field to given value.


### GetExpectedOutcome

`func (o *StrategyTestMetrics) GetExpectedOutcome() string`

GetExpectedOutcome returns the ExpectedOutcome field if non-nil, zero value otherwise.

### GetExpectedOutcomeOk

`func (o *StrategyTestMetrics) GetExpectedOutcomeOk() (*string, bool)`

GetExpectedOutcomeOk returns a tuple with the ExpectedOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutcome

`func (o *StrategyTestMetrics) SetExpectedOutcome(v string)`

SetExpectedOutcome sets ExpectedOutcome field to given value.


### GetExplanation

`func (o *StrategyTestMetrics) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *StrategyTestMetrics) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *StrategyTestMetrics) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### GetBaselineResult

`func (o *StrategyTestMetrics) GetBaselineResult() TestResult`

GetBaselineResult returns the BaselineResult field if non-nil, zero value otherwise.

### GetBaselineResultOk

`func (o *StrategyTestMetrics) GetBaselineResultOk() (*TestResult, bool)`

GetBaselineResultOk returns a tuple with the BaselineResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineResult

`func (o *StrategyTestMetrics) SetBaselineResult(v TestResult)`

SetBaselineResult sets BaselineResult field to given value.

### HasBaselineResult

`func (o *StrategyTestMetrics) HasBaselineResult() bool`

HasBaselineResult returns a boolean if a field has been set.

### SetBaselineResultNil

`func (o *StrategyTestMetrics) SetBaselineResultNil(b bool)`

 SetBaselineResultNil sets the value for BaselineResult to be an explicit nil

### UnsetBaselineResult
`func (o *StrategyTestMetrics) UnsetBaselineResult()`

UnsetBaselineResult ensures that no value is present for BaselineResult, not even an explicit nil
### GetStrategyResult

`func (o *StrategyTestMetrics) GetStrategyResult() TestResult`

GetStrategyResult returns the StrategyResult field if non-nil, zero value otherwise.

### GetStrategyResultOk

`func (o *StrategyTestMetrics) GetStrategyResultOk() (*TestResult, bool)`

GetStrategyResultOk returns a tuple with the StrategyResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyResult

`func (o *StrategyTestMetrics) SetStrategyResult(v TestResult)`

SetStrategyResult sets StrategyResult field to given value.

### HasStrategyResult

`func (o *StrategyTestMetrics) HasStrategyResult() bool`

HasStrategyResult returns a boolean if a field has been set.

### SetStrategyResultNil

`func (o *StrategyTestMetrics) SetStrategyResultNil(b bool)`

 SetStrategyResultNil sets the value for StrategyResult to be an explicit nil

### UnsetStrategyResult
`func (o *StrategyTestMetrics) UnsetStrategyResult()`

UnsetStrategyResult ensures that no value is present for StrategyResult, not even an explicit nil
### GetSubsequentResult

`func (o *StrategyTestMetrics) GetSubsequentResult() TestResult`

GetSubsequentResult returns the SubsequentResult field if non-nil, zero value otherwise.

### GetSubsequentResultOk

`func (o *StrategyTestMetrics) GetSubsequentResultOk() (*TestResult, bool)`

GetSubsequentResultOk returns a tuple with the SubsequentResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsequentResult

`func (o *StrategyTestMetrics) SetSubsequentResult(v TestResult)`

SetSubsequentResult sets SubsequentResult field to given value.

### HasSubsequentResult

`func (o *StrategyTestMetrics) HasSubsequentResult() bool`

HasSubsequentResult returns a boolean if a field has been set.

### SetSubsequentResultNil

`func (o *StrategyTestMetrics) SetSubsequentResultNil(b bool)`

 SetSubsequentResultNil sets the value for SubsequentResult to be an explicit nil

### UnsetSubsequentResult
`func (o *StrategyTestMetrics) UnsetSubsequentResult()`

UnsetSubsequentResult ensures that no value is present for SubsequentResult, not even an explicit nil
### GetClassification

`func (o *StrategyTestMetrics) GetClassification() TestClassification`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *StrategyTestMetrics) GetClassificationOk() (*TestClassification, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *StrategyTestMetrics) SetClassification(v TestClassification)`

SetClassification sets Classification field to given value.


### GetKnownIssues

`func (o *StrategyTestMetrics) GetKnownIssues() []KnownIssue`

GetKnownIssues returns the KnownIssues field if non-nil, zero value otherwise.

### GetKnownIssuesOk

`func (o *StrategyTestMetrics) GetKnownIssuesOk() (*[]KnownIssue, bool)`

GetKnownIssuesOk returns a tuple with the KnownIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnownIssues

`func (o *StrategyTestMetrics) SetKnownIssues(v []KnownIssue)`

SetKnownIssues sets KnownIssues field to given value.

### HasKnownIssues

`func (o *StrategyTestMetrics) HasKnownIssues() bool`

HasKnownIssues returns a boolean if a field has been set.

### GetIsKnownIssue

`func (o *StrategyTestMetrics) GetIsKnownIssue() bool`

GetIsKnownIssue returns the IsKnownIssue field if non-nil, zero value otherwise.

### GetIsKnownIssueOk

`func (o *StrategyTestMetrics) GetIsKnownIssueOk() (*bool, bool)`

GetIsKnownIssueOk returns a tuple with the IsKnownIssue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKnownIssue

`func (o *StrategyTestMetrics) SetIsKnownIssue(v bool)`

SetIsKnownIssue sets IsKnownIssue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


