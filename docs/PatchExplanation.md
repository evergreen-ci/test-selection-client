# PatchExplanation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchId** | **string** |  | 
**Builds** | Pointer to [**map[string]BuildExplanation**](BuildExplanation.md) |  | [optional] [default to {}]
**SuccessMeasure** | Pointer to [**map[string]SuccessMeasure**](SuccessMeasure.md) |  | [optional] [default to {}]

## Methods

### NewPatchExplanation

`func NewPatchExplanation(patchId string, ) *PatchExplanation`

NewPatchExplanation instantiates a new PatchExplanation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatchExplanationWithDefaults

`func NewPatchExplanationWithDefaults() *PatchExplanation`

NewPatchExplanationWithDefaults instantiates a new PatchExplanation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchId

`func (o *PatchExplanation) GetPatchId() string`

GetPatchId returns the PatchId field if non-nil, zero value otherwise.

### GetPatchIdOk

`func (o *PatchExplanation) GetPatchIdOk() (*string, bool)`

GetPatchIdOk returns a tuple with the PatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchId

`func (o *PatchExplanation) SetPatchId(v string)`

SetPatchId sets PatchId field to given value.


### GetBuilds

`func (o *PatchExplanation) GetBuilds() map[string]BuildExplanation`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *PatchExplanation) GetBuildsOk() (*map[string]BuildExplanation, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *PatchExplanation) SetBuilds(v map[string]BuildExplanation)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *PatchExplanation) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetSuccessMeasure

`func (o *PatchExplanation) GetSuccessMeasure() map[string]SuccessMeasure`

GetSuccessMeasure returns the SuccessMeasure field if non-nil, zero value otherwise.

### GetSuccessMeasureOk

`func (o *PatchExplanation) GetSuccessMeasureOk() (*map[string]SuccessMeasure, bool)`

GetSuccessMeasureOk returns a tuple with the SuccessMeasure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMeasure

`func (o *PatchExplanation) SetSuccessMeasure(v map[string]SuccessMeasure)`

SetSuccessMeasure sets SuccessMeasure field to given value.

### HasSuccessMeasure

`func (o *PatchExplanation) HasSuccessMeasure() bool`

HasSuccessMeasure returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


