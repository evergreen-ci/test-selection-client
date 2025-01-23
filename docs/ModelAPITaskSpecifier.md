# ModelAPITaskSpecifier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PatchAlias** | Pointer to **string** | Patch alias to run. | [optional] 
**TaskRegex** | Pointer to **string** | Regex matching tasks to run. | [optional] 
**VariantRegex** | Pointer to **string** | Regex matching build variants to run. | [optional] 

## Methods

### NewModelAPITaskSpecifier

`func NewModelAPITaskSpecifier() *ModelAPITaskSpecifier`

NewModelAPITaskSpecifier instantiates a new ModelAPITaskSpecifier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITaskSpecifierWithDefaults

`func NewModelAPITaskSpecifierWithDefaults() *ModelAPITaskSpecifier`

NewModelAPITaskSpecifierWithDefaults instantiates a new ModelAPITaskSpecifier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatchAlias

`func (o *ModelAPITaskSpecifier) GetPatchAlias() string`

GetPatchAlias returns the PatchAlias field if non-nil, zero value otherwise.

### GetPatchAliasOk

`func (o *ModelAPITaskSpecifier) GetPatchAliasOk() (*string, bool)`

GetPatchAliasOk returns a tuple with the PatchAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchAlias

`func (o *ModelAPITaskSpecifier) SetPatchAlias(v string)`

SetPatchAlias sets PatchAlias field to given value.

### HasPatchAlias

`func (o *ModelAPITaskSpecifier) HasPatchAlias() bool`

HasPatchAlias returns a boolean if a field has been set.

### GetTaskRegex

`func (o *ModelAPITaskSpecifier) GetTaskRegex() string`

GetTaskRegex returns the TaskRegex field if non-nil, zero value otherwise.

### GetTaskRegexOk

`func (o *ModelAPITaskSpecifier) GetTaskRegexOk() (*string, bool)`

GetTaskRegexOk returns a tuple with the TaskRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRegex

`func (o *ModelAPITaskSpecifier) SetTaskRegex(v string)`

SetTaskRegex sets TaskRegex field to given value.

### HasTaskRegex

`func (o *ModelAPITaskSpecifier) HasTaskRegex() bool`

HasTaskRegex returns a boolean if a field has been set.

### GetVariantRegex

`func (o *ModelAPITaskSpecifier) GetVariantRegex() string`

GetVariantRegex returns the VariantRegex field if non-nil, zero value otherwise.

### GetVariantRegexOk

`func (o *ModelAPITaskSpecifier) GetVariantRegexOk() (*string, bool)`

GetVariantRegexOk returns a tuple with the VariantRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantRegex

`func (o *ModelAPITaskSpecifier) SetVariantRegex(v string)`

SetVariantRegex sets VariantRegex field to given value.

### HasVariantRegex

`func (o *ModelAPITaskSpecifier) HasVariantRegex() bool`

HasVariantRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


