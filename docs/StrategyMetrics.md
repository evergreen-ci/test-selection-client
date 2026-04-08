# StrategyMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StrategyName** | [**StrategyEnum**](StrategyEnum.md) |  | 
**TruePositives** | Pointer to **int32** |  | [optional] [default to 0]
**FalsePositives** | Pointer to **int32** |  | [optional] [default to 0]
**TrueNegatives** | Pointer to **int32** |  | [optional] [default to 0]
**FalseNegatives** | Pointer to **int32** |  | [optional] [default to 0]
**Unknown** | Pointer to **int32** |  | [optional] [default to 0]
**TotalTests** | **int32** |  | 
**TestsSelected** | **int32** |  | 
**TestsFiltered** | **int32** |  | 

## Methods

### NewStrategyMetrics

`func NewStrategyMetrics(strategyName StrategyEnum, totalTests int32, testsSelected int32, testsFiltered int32, ) *StrategyMetrics`

NewStrategyMetrics instantiates a new StrategyMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStrategyMetricsWithDefaults

`func NewStrategyMetricsWithDefaults() *StrategyMetrics`

NewStrategyMetricsWithDefaults instantiates a new StrategyMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrategyName

`func (o *StrategyMetrics) GetStrategyName() StrategyEnum`

GetStrategyName returns the StrategyName field if non-nil, zero value otherwise.

### GetStrategyNameOk

`func (o *StrategyMetrics) GetStrategyNameOk() (*StrategyEnum, bool)`

GetStrategyNameOk returns a tuple with the StrategyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategyName

`func (o *StrategyMetrics) SetStrategyName(v StrategyEnum)`

SetStrategyName sets StrategyName field to given value.


### GetTruePositives

`func (o *StrategyMetrics) GetTruePositives() int32`

GetTruePositives returns the TruePositives field if non-nil, zero value otherwise.

### GetTruePositivesOk

`func (o *StrategyMetrics) GetTruePositivesOk() (*int32, bool)`

GetTruePositivesOk returns a tuple with the TruePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruePositives

`func (o *StrategyMetrics) SetTruePositives(v int32)`

SetTruePositives sets TruePositives field to given value.

### HasTruePositives

`func (o *StrategyMetrics) HasTruePositives() bool`

HasTruePositives returns a boolean if a field has been set.

### GetFalsePositives

`func (o *StrategyMetrics) GetFalsePositives() int32`

GetFalsePositives returns the FalsePositives field if non-nil, zero value otherwise.

### GetFalsePositivesOk

`func (o *StrategyMetrics) GetFalsePositivesOk() (*int32, bool)`

GetFalsePositivesOk returns a tuple with the FalsePositives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalsePositives

`func (o *StrategyMetrics) SetFalsePositives(v int32)`

SetFalsePositives sets FalsePositives field to given value.

### HasFalsePositives

`func (o *StrategyMetrics) HasFalsePositives() bool`

HasFalsePositives returns a boolean if a field has been set.

### GetTrueNegatives

`func (o *StrategyMetrics) GetTrueNegatives() int32`

GetTrueNegatives returns the TrueNegatives field if non-nil, zero value otherwise.

### GetTrueNegativesOk

`func (o *StrategyMetrics) GetTrueNegativesOk() (*int32, bool)`

GetTrueNegativesOk returns a tuple with the TrueNegatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueNegatives

`func (o *StrategyMetrics) SetTrueNegatives(v int32)`

SetTrueNegatives sets TrueNegatives field to given value.

### HasTrueNegatives

`func (o *StrategyMetrics) HasTrueNegatives() bool`

HasTrueNegatives returns a boolean if a field has been set.

### GetFalseNegatives

`func (o *StrategyMetrics) GetFalseNegatives() int32`

GetFalseNegatives returns the FalseNegatives field if non-nil, zero value otherwise.

### GetFalseNegativesOk

`func (o *StrategyMetrics) GetFalseNegativesOk() (*int32, bool)`

GetFalseNegativesOk returns a tuple with the FalseNegatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFalseNegatives

`func (o *StrategyMetrics) SetFalseNegatives(v int32)`

SetFalseNegatives sets FalseNegatives field to given value.

### HasFalseNegatives

`func (o *StrategyMetrics) HasFalseNegatives() bool`

HasFalseNegatives returns a boolean if a field has been set.

### GetUnknown

`func (o *StrategyMetrics) GetUnknown() int32`

GetUnknown returns the Unknown field if non-nil, zero value otherwise.

### GetUnknownOk

`func (o *StrategyMetrics) GetUnknownOk() (*int32, bool)`

GetUnknownOk returns a tuple with the Unknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknown

`func (o *StrategyMetrics) SetUnknown(v int32)`

SetUnknown sets Unknown field to given value.

### HasUnknown

`func (o *StrategyMetrics) HasUnknown() bool`

HasUnknown returns a boolean if a field has been set.

### GetTotalTests

`func (o *StrategyMetrics) GetTotalTests() int32`

GetTotalTests returns the TotalTests field if non-nil, zero value otherwise.

### GetTotalTestsOk

`func (o *StrategyMetrics) GetTotalTestsOk() (*int32, bool)`

GetTotalTestsOk returns a tuple with the TotalTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTests

`func (o *StrategyMetrics) SetTotalTests(v int32)`

SetTotalTests sets TotalTests field to given value.


### GetTestsSelected

`func (o *StrategyMetrics) GetTestsSelected() int32`

GetTestsSelected returns the TestsSelected field if non-nil, zero value otherwise.

### GetTestsSelectedOk

`func (o *StrategyMetrics) GetTestsSelectedOk() (*int32, bool)`

GetTestsSelectedOk returns a tuple with the TestsSelected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsSelected

`func (o *StrategyMetrics) SetTestsSelected(v int32)`

SetTestsSelected sets TestsSelected field to given value.


### GetTestsFiltered

`func (o *StrategyMetrics) GetTestsFiltered() int32`

GetTestsFiltered returns the TestsFiltered field if non-nil, zero value otherwise.

### GetTestsFilteredOk

`func (o *StrategyMetrics) GetTestsFilteredOk() (*int32, bool)`

GetTestsFilteredOk returns a tuple with the TestsFiltered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestsFiltered

`func (o *StrategyMetrics) SetTestsFiltered(v int32)`

SetTestsFiltered sets TestsFiltered field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


