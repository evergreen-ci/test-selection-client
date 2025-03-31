# Explanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TestName** | **string** |  | 
**Selected** | **bool** |  | 
**ExpectedOutcome** | [**ExpectedOutcome**](ExpectedOutcome.md) |  | 
**Explanation** | **string** |  | 

## Methods

### NewExplanation

`func NewExplanation(testName string, selected bool, expectedOutcome ExpectedOutcome, explanation string, ) *Explanation`

NewExplanation instantiates a new Explanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplanationWithDefaults

`func NewExplanationWithDefaults() *Explanation`

NewExplanationWithDefaults instantiates a new Explanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTestName

`func (o *Explanation) GetTestName() string`

GetTestName returns the TestName field if non-nil, zero value otherwise.

### GetTestNameOk

`func (o *Explanation) GetTestNameOk() (*string, bool)`

GetTestNameOk returns a tuple with the TestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestName

`func (o *Explanation) SetTestName(v string)`

SetTestName sets TestName field to given value.


### GetSelected

`func (o *Explanation) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *Explanation) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *Explanation) SetSelected(v bool)`

SetSelected sets Selected field to given value.


### GetExpectedOutcome

`func (o *Explanation) GetExpectedOutcome() ExpectedOutcome`

GetExpectedOutcome returns the ExpectedOutcome field if non-nil, zero value otherwise.

### GetExpectedOutcomeOk

`func (o *Explanation) GetExpectedOutcomeOk() (*ExpectedOutcome, bool)`

GetExpectedOutcomeOk returns a tuple with the ExpectedOutcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutcome

`func (o *Explanation) SetExpectedOutcome(v ExpectedOutcome)`

SetExpectedOutcome sets ExpectedOutcome field to given value.


### GetExplanation

`func (o *Explanation) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *Explanation) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *Explanation) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


