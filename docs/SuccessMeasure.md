# SuccessMeasure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalNumberOfTests** | Pointer to **int32** |  | [optional] [default to 0]
**NumberOfSelectedTests** | Pointer to **int32** |  | [optional] [default to 0]
**NumberOfUnselectedTests** | Pointer to **int32** |  | [optional] [default to 0]
**NumberOfCorrectlySelectedExpectedOutcomePassingTests** | Pointer to **int32** |  | [optional] [default to 0]
**NumberOfCorrectlyUnselectedExpectedOutcomePassingTests** | Pointer to **int32** |  | [optional] [default to 0]
**SelectedWithoutFilter** | Pointer to **int32** |  | [optional] [default to 0]
**NumberOfSelectedUnknownExpectedOutcomeTests** | Pointer to **int32** |  | [optional] [default to 0]
**NumberOfUnselectedUnknownExpectedOutcomeTests** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewSuccessMeasure

`func NewSuccessMeasure() *SuccessMeasure`

NewSuccessMeasure instantiates a new SuccessMeasure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSuccessMeasureWithDefaults

`func NewSuccessMeasureWithDefaults() *SuccessMeasure`

NewSuccessMeasureWithDefaults instantiates a new SuccessMeasure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalNumberOfTests

`func (o *SuccessMeasure) GetTotalNumberOfTests() int32`

GetTotalNumberOfTests returns the TotalNumberOfTests field if non-nil, zero value otherwise.

### GetTotalNumberOfTestsOk

`func (o *SuccessMeasure) GetTotalNumberOfTestsOk() (*int32, bool)`

GetTotalNumberOfTestsOk returns a tuple with the TotalNumberOfTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNumberOfTests

`func (o *SuccessMeasure) SetTotalNumberOfTests(v int32)`

SetTotalNumberOfTests sets TotalNumberOfTests field to given value.

### HasTotalNumberOfTests

`func (o *SuccessMeasure) HasTotalNumberOfTests() bool`

HasTotalNumberOfTests returns a boolean if a field has been set.

### GetNumberOfSelectedTests

`func (o *SuccessMeasure) GetNumberOfSelectedTests() int32`

GetNumberOfSelectedTests returns the NumberOfSelectedTests field if non-nil, zero value otherwise.

### GetNumberOfSelectedTestsOk

`func (o *SuccessMeasure) GetNumberOfSelectedTestsOk() (*int32, bool)`

GetNumberOfSelectedTestsOk returns a tuple with the NumberOfSelectedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSelectedTests

`func (o *SuccessMeasure) SetNumberOfSelectedTests(v int32)`

SetNumberOfSelectedTests sets NumberOfSelectedTests field to given value.

### HasNumberOfSelectedTests

`func (o *SuccessMeasure) HasNumberOfSelectedTests() bool`

HasNumberOfSelectedTests returns a boolean if a field has been set.

### GetNumberOfUnselectedTests

`func (o *SuccessMeasure) GetNumberOfUnselectedTests() int32`

GetNumberOfUnselectedTests returns the NumberOfUnselectedTests field if non-nil, zero value otherwise.

### GetNumberOfUnselectedTestsOk

`func (o *SuccessMeasure) GetNumberOfUnselectedTestsOk() (*int32, bool)`

GetNumberOfUnselectedTestsOk returns a tuple with the NumberOfUnselectedTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUnselectedTests

`func (o *SuccessMeasure) SetNumberOfUnselectedTests(v int32)`

SetNumberOfUnselectedTests sets NumberOfUnselectedTests field to given value.

### HasNumberOfUnselectedTests

`func (o *SuccessMeasure) HasNumberOfUnselectedTests() bool`

HasNumberOfUnselectedTests returns a boolean if a field has been set.

### GetNumberOfCorrectlySelectedExpectedOutcomePassingTests

`func (o *SuccessMeasure) GetNumberOfCorrectlySelectedExpectedOutcomePassingTests() int32`

GetNumberOfCorrectlySelectedExpectedOutcomePassingTests returns the NumberOfCorrectlySelectedExpectedOutcomePassingTests field if non-nil, zero value otherwise.

### GetNumberOfCorrectlySelectedExpectedOutcomePassingTestsOk

`func (o *SuccessMeasure) GetNumberOfCorrectlySelectedExpectedOutcomePassingTestsOk() (*int32, bool)`

GetNumberOfCorrectlySelectedExpectedOutcomePassingTestsOk returns a tuple with the NumberOfCorrectlySelectedExpectedOutcomePassingTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCorrectlySelectedExpectedOutcomePassingTests

`func (o *SuccessMeasure) SetNumberOfCorrectlySelectedExpectedOutcomePassingTests(v int32)`

SetNumberOfCorrectlySelectedExpectedOutcomePassingTests sets NumberOfCorrectlySelectedExpectedOutcomePassingTests field to given value.

### HasNumberOfCorrectlySelectedExpectedOutcomePassingTests

`func (o *SuccessMeasure) HasNumberOfCorrectlySelectedExpectedOutcomePassingTests() bool`

HasNumberOfCorrectlySelectedExpectedOutcomePassingTests returns a boolean if a field has been set.

### GetNumberOfCorrectlyUnselectedExpectedOutcomePassingTests

`func (o *SuccessMeasure) GetNumberOfCorrectlyUnselectedExpectedOutcomePassingTests() int32`

GetNumberOfCorrectlyUnselectedExpectedOutcomePassingTests returns the NumberOfCorrectlyUnselectedExpectedOutcomePassingTests field if non-nil, zero value otherwise.

### GetNumberOfCorrectlyUnselectedExpectedOutcomePassingTestsOk

`func (o *SuccessMeasure) GetNumberOfCorrectlyUnselectedExpectedOutcomePassingTestsOk() (*int32, bool)`

GetNumberOfCorrectlyUnselectedExpectedOutcomePassingTestsOk returns a tuple with the NumberOfCorrectlyUnselectedExpectedOutcomePassingTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCorrectlyUnselectedExpectedOutcomePassingTests

`func (o *SuccessMeasure) SetNumberOfCorrectlyUnselectedExpectedOutcomePassingTests(v int32)`

SetNumberOfCorrectlyUnselectedExpectedOutcomePassingTests sets NumberOfCorrectlyUnselectedExpectedOutcomePassingTests field to given value.

### HasNumberOfCorrectlyUnselectedExpectedOutcomePassingTests

`func (o *SuccessMeasure) HasNumberOfCorrectlyUnselectedExpectedOutcomePassingTests() bool`

HasNumberOfCorrectlyUnselectedExpectedOutcomePassingTests returns a boolean if a field has been set.

### GetSelectedWithoutFilter

`func (o *SuccessMeasure) GetSelectedWithoutFilter() int32`

GetSelectedWithoutFilter returns the SelectedWithoutFilter field if non-nil, zero value otherwise.

### GetSelectedWithoutFilterOk

`func (o *SuccessMeasure) GetSelectedWithoutFilterOk() (*int32, bool)`

GetSelectedWithoutFilterOk returns a tuple with the SelectedWithoutFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedWithoutFilter

`func (o *SuccessMeasure) SetSelectedWithoutFilter(v int32)`

SetSelectedWithoutFilter sets SelectedWithoutFilter field to given value.

### HasSelectedWithoutFilter

`func (o *SuccessMeasure) HasSelectedWithoutFilter() bool`

HasSelectedWithoutFilter returns a boolean if a field has been set.

### GetNumberOfSelectedUnknownExpectedOutcomeTests

`func (o *SuccessMeasure) GetNumberOfSelectedUnknownExpectedOutcomeTests() int32`

GetNumberOfSelectedUnknownExpectedOutcomeTests returns the NumberOfSelectedUnknownExpectedOutcomeTests field if non-nil, zero value otherwise.

### GetNumberOfSelectedUnknownExpectedOutcomeTestsOk

`func (o *SuccessMeasure) GetNumberOfSelectedUnknownExpectedOutcomeTestsOk() (*int32, bool)`

GetNumberOfSelectedUnknownExpectedOutcomeTestsOk returns a tuple with the NumberOfSelectedUnknownExpectedOutcomeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfSelectedUnknownExpectedOutcomeTests

`func (o *SuccessMeasure) SetNumberOfSelectedUnknownExpectedOutcomeTests(v int32)`

SetNumberOfSelectedUnknownExpectedOutcomeTests sets NumberOfSelectedUnknownExpectedOutcomeTests field to given value.

### HasNumberOfSelectedUnknownExpectedOutcomeTests

`func (o *SuccessMeasure) HasNumberOfSelectedUnknownExpectedOutcomeTests() bool`

HasNumberOfSelectedUnknownExpectedOutcomeTests returns a boolean if a field has been set.

### GetNumberOfUnselectedUnknownExpectedOutcomeTests

`func (o *SuccessMeasure) GetNumberOfUnselectedUnknownExpectedOutcomeTests() int32`

GetNumberOfUnselectedUnknownExpectedOutcomeTests returns the NumberOfUnselectedUnknownExpectedOutcomeTests field if non-nil, zero value otherwise.

### GetNumberOfUnselectedUnknownExpectedOutcomeTestsOk

`func (o *SuccessMeasure) GetNumberOfUnselectedUnknownExpectedOutcomeTestsOk() (*int32, bool)`

GetNumberOfUnselectedUnknownExpectedOutcomeTestsOk returns a tuple with the NumberOfUnselectedUnknownExpectedOutcomeTests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfUnselectedUnknownExpectedOutcomeTests

`func (o *SuccessMeasure) SetNumberOfUnselectedUnknownExpectedOutcomeTests(v int32)`

SetNumberOfUnselectedUnknownExpectedOutcomeTests sets NumberOfUnselectedUnknownExpectedOutcomeTests field to given value.

### HasNumberOfUnselectedUnknownExpectedOutcomeTests

`func (o *SuccessMeasure) HasNumberOfUnselectedUnknownExpectedOutcomeTests() bool`

HasNumberOfUnselectedUnknownExpectedOutcomeTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


