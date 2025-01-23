# ModelAPIProjectAlias

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier for the project alias. | [optional] 
**Alias** | Pointer to **string** | Name of the alias. | [optional] 
**Delete** | Pointer to **bool** | If set, deletes the project alias by name. | [optional] 
**Description** | Pointer to **string** | Human-friendly description for the alias. | [optional] 
**GitTag** | Pointer to **string** | Regex for matching git tags to run git tag versions. | [optional] 
**Parameters** | Pointer to [**[]ModelAPIParameter**](ModelAPIParameter.md) | List of allowed parameters to the alias. | [optional] 
**RemotePath** | Pointer to **string** | Path to project config file to use. | [optional] 
**Tags** | Pointer to **[]string** | Task tag selectors to match. | [optional] 
**Task** | Pointer to **string** | Regex for tasks to match. | [optional] 
**Variant** | Pointer to **string** | Regex for build variants to match. | [optional] 
**VariantTags** | Pointer to **[]string** | Build variant tags selectors to match. | [optional] 

## Methods

### NewModelAPIProjectAlias

`func NewModelAPIProjectAlias() *ModelAPIProjectAlias`

NewModelAPIProjectAlias instantiates a new ModelAPIProjectAlias object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIProjectAliasWithDefaults

`func NewModelAPIProjectAliasWithDefaults() *ModelAPIProjectAlias`

NewModelAPIProjectAliasWithDefaults instantiates a new ModelAPIProjectAlias object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelAPIProjectAlias) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPIProjectAlias) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPIProjectAlias) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPIProjectAlias) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAlias

`func (o *ModelAPIProjectAlias) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelAPIProjectAlias) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelAPIProjectAlias) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelAPIProjectAlias) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetDelete

`func (o *ModelAPIProjectAlias) GetDelete() bool`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ModelAPIProjectAlias) GetDeleteOk() (*bool, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ModelAPIProjectAlias) SetDelete(v bool)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ModelAPIProjectAlias) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *ModelAPIProjectAlias) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelAPIProjectAlias) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelAPIProjectAlias) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelAPIProjectAlias) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitTag

`func (o *ModelAPIProjectAlias) GetGitTag() string`

GetGitTag returns the GitTag field if non-nil, zero value otherwise.

### GetGitTagOk

`func (o *ModelAPIProjectAlias) GetGitTagOk() (*string, bool)`

GetGitTagOk returns a tuple with the GitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTag

`func (o *ModelAPIProjectAlias) SetGitTag(v string)`

SetGitTag sets GitTag field to given value.

### HasGitTag

`func (o *ModelAPIProjectAlias) HasGitTag() bool`

HasGitTag returns a boolean if a field has been set.

### GetParameters

`func (o *ModelAPIProjectAlias) GetParameters() []ModelAPIParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ModelAPIProjectAlias) GetParametersOk() (*[]ModelAPIParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ModelAPIProjectAlias) SetParameters(v []ModelAPIParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ModelAPIProjectAlias) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRemotePath

`func (o *ModelAPIProjectAlias) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *ModelAPIProjectAlias) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *ModelAPIProjectAlias) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *ModelAPIProjectAlias) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetTags

`func (o *ModelAPIProjectAlias) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelAPIProjectAlias) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelAPIProjectAlias) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelAPIProjectAlias) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTask

`func (o *ModelAPIProjectAlias) GetTask() string`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *ModelAPIProjectAlias) GetTaskOk() (*string, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *ModelAPIProjectAlias) SetTask(v string)`

SetTask sets Task field to given value.

### HasTask

`func (o *ModelAPIProjectAlias) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetVariant

`func (o *ModelAPIProjectAlias) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ModelAPIProjectAlias) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ModelAPIProjectAlias) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ModelAPIProjectAlias) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetVariantTags

`func (o *ModelAPIProjectAlias) GetVariantTags() []string`

GetVariantTags returns the VariantTags field if non-nil, zero value otherwise.

### GetVariantTagsOk

`func (o *ModelAPIProjectAlias) GetVariantTagsOk() (*[]string, bool)`

GetVariantTagsOk returns a tuple with the VariantTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantTags

`func (o *ModelAPIProjectAlias) SetVariantTags(v []string)`

SetVariantTags sets VariantTags field to given value.

### HasVariantTags

`func (o *ModelAPIProjectAlias) HasVariantTags() bool`

HasVariantTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


