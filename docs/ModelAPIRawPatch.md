# ModelAPIRawPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Patch** | Pointer to [**ModelAPIRawModule**](ModelAPIRawModule.md) | The main patch | [optional] 
**RawModules** | Pointer to [**[]ModelAPIRawModule**](ModelAPIRawModule.md) | The list of module diffs | [optional] 

## Methods

### NewModelAPIRawPatch

`func NewModelAPIRawPatch() *ModelAPIRawPatch`

NewModelAPIRawPatch instantiates a new ModelAPIRawPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIRawPatchWithDefaults

`func NewModelAPIRawPatchWithDefaults() *ModelAPIRawPatch`

NewModelAPIRawPatchWithDefaults instantiates a new ModelAPIRawPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPatch

`func (o *ModelAPIRawPatch) GetPatch() ModelAPIRawModule`

GetPatch returns the Patch field if non-nil, zero value otherwise.

### GetPatchOk

`func (o *ModelAPIRawPatch) GetPatchOk() (*ModelAPIRawModule, bool)`

GetPatchOk returns a tuple with the Patch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatch

`func (o *ModelAPIRawPatch) SetPatch(v ModelAPIRawModule)`

SetPatch sets Patch field to given value.

### HasPatch

`func (o *ModelAPIRawPatch) HasPatch() bool`

HasPatch returns a boolean if a field has been set.

### GetRawModules

`func (o *ModelAPIRawPatch) GetRawModules() []ModelAPIRawModule`

GetRawModules returns the RawModules field if non-nil, zero value otherwise.

### GetRawModulesOk

`func (o *ModelAPIRawPatch) GetRawModulesOk() (*[]ModelAPIRawModule, bool)`

GetRawModulesOk returns a tuple with the RawModules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawModules

`func (o *ModelAPIRawPatch) SetRawModules(v []ModelAPIRawModule)`

SetRawModules sets RawModules field to given value.

### HasRawModules

`func (o *ModelAPIRawPatch) HasRawModules() bool`

HasRawModules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


