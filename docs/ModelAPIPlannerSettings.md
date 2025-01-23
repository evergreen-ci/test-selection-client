# ModelAPIPlannerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommitQueueFactor** | Pointer to **int32** |  | [optional] 
**ExpectedRuntimeFactor** | Pointer to **int32** |  | [optional] 
**GenerateTaskFactor** | Pointer to **int32** |  | [optional] 
**GroupVersions** | Pointer to **bool** |  | [optional] 
**MainlineTimeInQueueFactor** | Pointer to **int32** |  | [optional] 
**NumDependentsFactor** | Pointer to **float32** |  | [optional] 
**PatchFactor** | Pointer to **int32** |  | [optional] 
**PatchTimeInQueueFactor** | Pointer to **int32** |  | [optional] 
**TargetTime** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIPlannerSettings

`func NewModelAPIPlannerSettings() *ModelAPIPlannerSettings`

NewModelAPIPlannerSettings instantiates a new ModelAPIPlannerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIPlannerSettingsWithDefaults

`func NewModelAPIPlannerSettingsWithDefaults() *ModelAPIPlannerSettings`

NewModelAPIPlannerSettingsWithDefaults instantiates a new ModelAPIPlannerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommitQueueFactor

`func (o *ModelAPIPlannerSettings) GetCommitQueueFactor() int32`

GetCommitQueueFactor returns the CommitQueueFactor field if non-nil, zero value otherwise.

### GetCommitQueueFactorOk

`func (o *ModelAPIPlannerSettings) GetCommitQueueFactorOk() (*int32, bool)`

GetCommitQueueFactorOk returns a tuple with the CommitQueueFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitQueueFactor

`func (o *ModelAPIPlannerSettings) SetCommitQueueFactor(v int32)`

SetCommitQueueFactor sets CommitQueueFactor field to given value.

### HasCommitQueueFactor

`func (o *ModelAPIPlannerSettings) HasCommitQueueFactor() bool`

HasCommitQueueFactor returns a boolean if a field has been set.

### GetExpectedRuntimeFactor

`func (o *ModelAPIPlannerSettings) GetExpectedRuntimeFactor() int32`

GetExpectedRuntimeFactor returns the ExpectedRuntimeFactor field if non-nil, zero value otherwise.

### GetExpectedRuntimeFactorOk

`func (o *ModelAPIPlannerSettings) GetExpectedRuntimeFactorOk() (*int32, bool)`

GetExpectedRuntimeFactorOk returns a tuple with the ExpectedRuntimeFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedRuntimeFactor

`func (o *ModelAPIPlannerSettings) SetExpectedRuntimeFactor(v int32)`

SetExpectedRuntimeFactor sets ExpectedRuntimeFactor field to given value.

### HasExpectedRuntimeFactor

`func (o *ModelAPIPlannerSettings) HasExpectedRuntimeFactor() bool`

HasExpectedRuntimeFactor returns a boolean if a field has been set.

### GetGenerateTaskFactor

`func (o *ModelAPIPlannerSettings) GetGenerateTaskFactor() int32`

GetGenerateTaskFactor returns the GenerateTaskFactor field if non-nil, zero value otherwise.

### GetGenerateTaskFactorOk

`func (o *ModelAPIPlannerSettings) GetGenerateTaskFactorOk() (*int32, bool)`

GetGenerateTaskFactorOk returns a tuple with the GenerateTaskFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateTaskFactor

`func (o *ModelAPIPlannerSettings) SetGenerateTaskFactor(v int32)`

SetGenerateTaskFactor sets GenerateTaskFactor field to given value.

### HasGenerateTaskFactor

`func (o *ModelAPIPlannerSettings) HasGenerateTaskFactor() bool`

HasGenerateTaskFactor returns a boolean if a field has been set.

### GetGroupVersions

`func (o *ModelAPIPlannerSettings) GetGroupVersions() bool`

GetGroupVersions returns the GroupVersions field if non-nil, zero value otherwise.

### GetGroupVersionsOk

`func (o *ModelAPIPlannerSettings) GetGroupVersionsOk() (*bool, bool)`

GetGroupVersionsOk returns a tuple with the GroupVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupVersions

`func (o *ModelAPIPlannerSettings) SetGroupVersions(v bool)`

SetGroupVersions sets GroupVersions field to given value.

### HasGroupVersions

`func (o *ModelAPIPlannerSettings) HasGroupVersions() bool`

HasGroupVersions returns a boolean if a field has been set.

### GetMainlineTimeInQueueFactor

`func (o *ModelAPIPlannerSettings) GetMainlineTimeInQueueFactor() int32`

GetMainlineTimeInQueueFactor returns the MainlineTimeInQueueFactor field if non-nil, zero value otherwise.

### GetMainlineTimeInQueueFactorOk

`func (o *ModelAPIPlannerSettings) GetMainlineTimeInQueueFactorOk() (*int32, bool)`

GetMainlineTimeInQueueFactorOk returns a tuple with the MainlineTimeInQueueFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainlineTimeInQueueFactor

`func (o *ModelAPIPlannerSettings) SetMainlineTimeInQueueFactor(v int32)`

SetMainlineTimeInQueueFactor sets MainlineTimeInQueueFactor field to given value.

### HasMainlineTimeInQueueFactor

`func (o *ModelAPIPlannerSettings) HasMainlineTimeInQueueFactor() bool`

HasMainlineTimeInQueueFactor returns a boolean if a field has been set.

### GetNumDependentsFactor

`func (o *ModelAPIPlannerSettings) GetNumDependentsFactor() float32`

GetNumDependentsFactor returns the NumDependentsFactor field if non-nil, zero value otherwise.

### GetNumDependentsFactorOk

`func (o *ModelAPIPlannerSettings) GetNumDependentsFactorOk() (*float32, bool)`

GetNumDependentsFactorOk returns a tuple with the NumDependentsFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumDependentsFactor

`func (o *ModelAPIPlannerSettings) SetNumDependentsFactor(v float32)`

SetNumDependentsFactor sets NumDependentsFactor field to given value.

### HasNumDependentsFactor

`func (o *ModelAPIPlannerSettings) HasNumDependentsFactor() bool`

HasNumDependentsFactor returns a boolean if a field has been set.

### GetPatchFactor

`func (o *ModelAPIPlannerSettings) GetPatchFactor() int32`

GetPatchFactor returns the PatchFactor field if non-nil, zero value otherwise.

### GetPatchFactorOk

`func (o *ModelAPIPlannerSettings) GetPatchFactorOk() (*int32, bool)`

GetPatchFactorOk returns a tuple with the PatchFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchFactor

`func (o *ModelAPIPlannerSettings) SetPatchFactor(v int32)`

SetPatchFactor sets PatchFactor field to given value.

### HasPatchFactor

`func (o *ModelAPIPlannerSettings) HasPatchFactor() bool`

HasPatchFactor returns a boolean if a field has been set.

### GetPatchTimeInQueueFactor

`func (o *ModelAPIPlannerSettings) GetPatchTimeInQueueFactor() int32`

GetPatchTimeInQueueFactor returns the PatchTimeInQueueFactor field if non-nil, zero value otherwise.

### GetPatchTimeInQueueFactorOk

`func (o *ModelAPIPlannerSettings) GetPatchTimeInQueueFactorOk() (*int32, bool)`

GetPatchTimeInQueueFactorOk returns a tuple with the PatchTimeInQueueFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchTimeInQueueFactor

`func (o *ModelAPIPlannerSettings) SetPatchTimeInQueueFactor(v int32)`

SetPatchTimeInQueueFactor sets PatchTimeInQueueFactor field to given value.

### HasPatchTimeInQueueFactor

`func (o *ModelAPIPlannerSettings) HasPatchTimeInQueueFactor() bool`

HasPatchTimeInQueueFactor returns a boolean if a field has been set.

### GetTargetTime

`func (o *ModelAPIPlannerSettings) GetTargetTime() int32`

GetTargetTime returns the TargetTime field if non-nil, zero value otherwise.

### GetTargetTimeOk

`func (o *ModelAPIPlannerSettings) GetTargetTimeOk() (*int32, bool)`

GetTargetTimeOk returns a tuple with the TargetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetTime

`func (o *ModelAPIPlannerSettings) SetTargetTime(v int32)`

SetTargetTime sets TargetTime field to given value.

### HasTargetTime

`func (o *ModelAPIPlannerSettings) HasTargetTime() bool`

HasTargetTime returns a boolean if a field has been set.

### GetVersion

`func (o *ModelAPIPlannerSettings) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelAPIPlannerSettings) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelAPIPlannerSettings) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelAPIPlannerSettings) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


