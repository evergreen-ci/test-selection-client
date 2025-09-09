# TestExplanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestName** | **string** |  | 
**IsPassing** | **bool** |  | 
**Selected** | **bool** |  | 
**Explanation** | **string** |  | 
**Explanations** | Pointer to [**map[string]Explanation**](Explanation.md) |  | [optional] [default to {}]
**ExpectedOutcome** | [**ExpectedOutcome**](ExpectedOutcome.md) |  | 

## Methods

### NewTestExplanation

`func NewTestExplanation(testName string, isPassing bool, selected bool, explanation string, expectedOutcome ExpectedOutcome, ) *TestExplanation`

NewTestExplanation instantiates a new TestExplanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestExplanationWithDefaults

`func NewTestExplanationWithDefaults() *TestExplanation`

NewTestExplanationWithDefaults instantiates a new TestExplanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestName

`func (o *TestExplanation) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *TestExplanation) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *TestExplanation) SetTestName(v string)`

SetTestName sets TestName field to given value.


### GetIsPassing

`func (o *TestExplanation) GetIsPassing() bool`

GetIsPassing returns the IsPassing field if non-nil, zero value otherwise.

### GetIsPassingOk

`func (o *TestExplanation) GetIsPassingOk() (*bool, bool)`

GetIsPassingOk returns a tuple with the IsPassing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassing

`func (o *TestExplanation) SetIsPassing(v bool)`

SetIsPassing sets IsPassing field to given value.


### GetSelected

`func (o *TestExplanation) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *TestExplanation) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *TestExplanation) SetSelected(v bool)`

SetSelected sets Selected field to given value.


### GetExplanation

`func (o *TestExplanation) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *TestExplanation) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *TestExplanation) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.


### GetExplanations

`func (o *TestExplanation) GetExplanations() map[string]Explanation`

GetExplanations returns the Explanations field if non-nil, zero value otherwise.

### GetExplanationsOk

`func (o *TestExplanation) GetExplanationsOk() (*map[string]Explanation, bool)`

GetExplanationsOk returns a tuple with the Explanations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanations

`func (o *TestExplanation) SetExplanations(v map[string]Explanation)`

SetExplanations sets Explanations field to given value.

### HasExplanations

`func (o *TestExplanation) HasExplanations() bool`

HasExplanations returns a boolean if a field has been set.

### GetExpectedOutcome

`func (o *TestExplanation) GetExpectedOutcome() ExpectedOutcome`

GetExpectedOutcome returns the ExpectedOutcome field if non-nil, zero value otherwise.

### GetExpectedOutcomeOk

`func (o *TestExplanation) GetExpectedOutcomeOk() (*ExpectedOutcome, bool)`

GetExpectedOutcomeOk returns a tuple with the ExpectedOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutcome

`func (o *TestExplanation) SetExpectedOutcome(v ExpectedOutcome)`

SetExpectedOutcome sets ExpectedOutcome field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


