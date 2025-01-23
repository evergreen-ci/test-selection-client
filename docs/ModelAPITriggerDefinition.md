# ModelAPITriggerDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias to run for the trigger. | [optional] 
**ConfigFile** | Pointer to **string** | Project configuration file for the trigger. | [optional] 
**DateCutoff** | Pointer to **int32** | Number of days after commit when the trigger cannot run. | [optional] 
**DefinitionId** | Pointer to **string** | Identifier for the definition. | [optional] 
**Level** | Pointer to **string** | Trigger on build, task, or push. | [optional] 
**Project** | Pointer to **string** | Identifier of project to watch. | [optional] 
**Status** | Pointer to **string** | Task status to trigger for (or \&quot;*\&quot; for all). | [optional] 
**TaskRegex** | Pointer to **string** | Task regex to match. | [optional] 
**UnscheduleDownstreamVersions** | Pointer to **bool** | Deactivate downstream versions created by this trigger. | [optional] 
**VariantRegex** | Pointer to **string** | Build variant regex to match. | [optional] 

## Methods

### NewModelAPITriggerDefinition

`func NewModelAPITriggerDefinition() *ModelAPITriggerDefinition`

NewModelAPITriggerDefinition instantiates a new ModelAPITriggerDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITriggerDefinitionWithDefaults

`func NewModelAPITriggerDefinitionWithDefaults() *ModelAPITriggerDefinition`

NewModelAPITriggerDefinitionWithDefaults instantiates a new ModelAPITriggerDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ModelAPITriggerDefinition) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelAPITriggerDefinition) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelAPITriggerDefinition) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelAPITriggerDefinition) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetConfigFile

`func (o *ModelAPITriggerDefinition) GetConfigFile() string`

GetConfigFile returns the ConfigFile field if non-nil, zero value otherwise.

### GetConfigFileOk

`func (o *ModelAPITriggerDefinition) GetConfigFileOk() (*string, bool)`

GetConfigFileOk returns a tuple with the ConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFile

`func (o *ModelAPITriggerDefinition) SetConfigFile(v string)`

SetConfigFile sets ConfigFile field to given value.

### HasConfigFile

`func (o *ModelAPITriggerDefinition) HasConfigFile() bool`

HasConfigFile returns a boolean if a field has been set.

### GetDateCutoff

`func (o *ModelAPITriggerDefinition) GetDateCutoff() int32`

GetDateCutoff returns the DateCutoff field if non-nil, zero value otherwise.

### GetDateCutoffOk

`func (o *ModelAPITriggerDefinition) GetDateCutoffOk() (*int32, bool)`

GetDateCutoffOk returns a tuple with the DateCutoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCutoff

`func (o *ModelAPITriggerDefinition) SetDateCutoff(v int32)`

SetDateCutoff sets DateCutoff field to given value.

### HasDateCutoff

`func (o *ModelAPITriggerDefinition) HasDateCutoff() bool`

HasDateCutoff returns a boolean if a field has been set.

### GetDefinitionId

`func (o *ModelAPITriggerDefinition) GetDefinitionId() string`

GetDefinitionId returns the DefinitionId field if non-nil, zero value otherwise.

### GetDefinitionIdOk

`func (o *ModelAPITriggerDefinition) GetDefinitionIdOk() (*string, bool)`

GetDefinitionIdOk returns a tuple with the DefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionId

`func (o *ModelAPITriggerDefinition) SetDefinitionId(v string)`

SetDefinitionId sets DefinitionId field to given value.

### HasDefinitionId

`func (o *ModelAPITriggerDefinition) HasDefinitionId() bool`

HasDefinitionId returns a boolean if a field has been set.

### GetLevel

`func (o *ModelAPITriggerDefinition) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ModelAPITriggerDefinition) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ModelAPITriggerDefinition) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ModelAPITriggerDefinition) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetProject

`func (o *ModelAPITriggerDefinition) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ModelAPITriggerDefinition) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ModelAPITriggerDefinition) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ModelAPITriggerDefinition) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPITriggerDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPITriggerDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPITriggerDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPITriggerDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskRegex

`func (o *ModelAPITriggerDefinition) GetTaskRegex() string`

GetTaskRegex returns the TaskRegex field if non-nil, zero value otherwise.

### GetTaskRegexOk

`func (o *ModelAPITriggerDefinition) GetTaskRegexOk() (*string, bool)`

GetTaskRegexOk returns a tuple with the TaskRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskRegex

`func (o *ModelAPITriggerDefinition) SetTaskRegex(v string)`

SetTaskRegex sets TaskRegex field to given value.

### HasTaskRegex

`func (o *ModelAPITriggerDefinition) HasTaskRegex() bool`

HasTaskRegex returns a boolean if a field has been set.

### GetUnscheduleDownstreamVersions

`func (o *ModelAPITriggerDefinition) GetUnscheduleDownstreamVersions() bool`

GetUnscheduleDownstreamVersions returns the UnscheduleDownstreamVersions field if non-nil, zero value otherwise.

### GetUnscheduleDownstreamVersionsOk

`func (o *ModelAPITriggerDefinition) GetUnscheduleDownstreamVersionsOk() (*bool, bool)`

GetUnscheduleDownstreamVersionsOk returns a tuple with the UnscheduleDownstreamVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnscheduleDownstreamVersions

`func (o *ModelAPITriggerDefinition) SetUnscheduleDownstreamVersions(v bool)`

SetUnscheduleDownstreamVersions sets UnscheduleDownstreamVersions field to given value.

### HasUnscheduleDownstreamVersions

`func (o *ModelAPITriggerDefinition) HasUnscheduleDownstreamVersions() bool`

HasUnscheduleDownstreamVersions returns a boolean if a field has been set.

### GetVariantRegex

`func (o *ModelAPITriggerDefinition) GetVariantRegex() string`

GetVariantRegex returns the VariantRegex field if non-nil, zero value otherwise.

### GetVariantRegexOk

`func (o *ModelAPITriggerDefinition) GetVariantRegexOk() (*string, bool)`

GetVariantRegexOk returns a tuple with the VariantRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantRegex

`func (o *ModelAPITriggerDefinition) SetVariantRegex(v string)`

SetVariantRegex sets VariantRegex field to given value.

### HasVariantRegex

`func (o *ModelAPITriggerDefinition) HasVariantRegex() bool`

HasVariantRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


