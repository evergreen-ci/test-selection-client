# ModelAPIPatchTriggerDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias to run in the downstream project. | [optional] 
**ChildProjectId** | Pointer to **string** | ID of the downstream project. | [optional] 
**ChildProjectIdentifier** | Pointer to **string** | Identifier of the downstream project. | [optional] 
**DownstreamRevision** | Pointer to **string** | An optional field representing the revision at which to create the downstream patch. By default, this field is empty and the downstream patch will be based off of its most recent commit. | [optional] 
**ParentAsModule** | Pointer to **string** | Name of the module corresponding to the upstream project in the downstream project&#39;s YAML. | [optional] 
**Status** | Pointer to **string** | Status for the parent patch to conditionally kick off the child patch. | [optional] 
**TaskSpecifiers** | Pointer to [**[]ModelAPITaskSpecifier**](ModelAPITaskSpecifier.md) | List of task specifiers. | [optional] 
**VariantsTasks** | Pointer to [**[]ModelVariantTask**](ModelVariantTask.md) | The list of variants/tasks from the alias that will run in the downstream project. | [optional] 

## Methods

### NewModelAPIPatchTriggerDefinition

`func NewModelAPIPatchTriggerDefinition() *ModelAPIPatchTriggerDefinition`

NewModelAPIPatchTriggerDefinition instantiates a new ModelAPIPatchTriggerDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIPatchTriggerDefinitionWithDefaults

`func NewModelAPIPatchTriggerDefinitionWithDefaults() *ModelAPIPatchTriggerDefinition`

NewModelAPIPatchTriggerDefinitionWithDefaults instantiates a new ModelAPIPatchTriggerDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ModelAPIPatchTriggerDefinition) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelAPIPatchTriggerDefinition) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelAPIPatchTriggerDefinition) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelAPIPatchTriggerDefinition) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetChildProjectId

`func (o *ModelAPIPatchTriggerDefinition) GetChildProjectId() string`

GetChildProjectId returns the ChildProjectId field if non-nil, zero value otherwise.

### GetChildProjectIdOk

`func (o *ModelAPIPatchTriggerDefinition) GetChildProjectIdOk() (*string, bool)`

GetChildProjectIdOk returns a tuple with the ChildProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildProjectId

`func (o *ModelAPIPatchTriggerDefinition) SetChildProjectId(v string)`

SetChildProjectId sets ChildProjectId field to given value.

### HasChildProjectId

`func (o *ModelAPIPatchTriggerDefinition) HasChildProjectId() bool`

HasChildProjectId returns a boolean if a field has been set.

### GetChildProjectIdentifier

`func (o *ModelAPIPatchTriggerDefinition) GetChildProjectIdentifier() string`

GetChildProjectIdentifier returns the ChildProjectIdentifier field if non-nil, zero value otherwise.

### GetChildProjectIdentifierOk

`func (o *ModelAPIPatchTriggerDefinition) GetChildProjectIdentifierOk() (*string, bool)`

GetChildProjectIdentifierOk returns a tuple with the ChildProjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildProjectIdentifier

`func (o *ModelAPIPatchTriggerDefinition) SetChildProjectIdentifier(v string)`

SetChildProjectIdentifier sets ChildProjectIdentifier field to given value.

### HasChildProjectIdentifier

`func (o *ModelAPIPatchTriggerDefinition) HasChildProjectIdentifier() bool`

HasChildProjectIdentifier returns a boolean if a field has been set.

### GetDownstreamRevision

`func (o *ModelAPIPatchTriggerDefinition) GetDownstreamRevision() string`

GetDownstreamRevision returns the DownstreamRevision field if non-nil, zero value otherwise.

### GetDownstreamRevisionOk

`func (o *ModelAPIPatchTriggerDefinition) GetDownstreamRevisionOk() (*string, bool)`

GetDownstreamRevisionOk returns a tuple with the DownstreamRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamRevision

`func (o *ModelAPIPatchTriggerDefinition) SetDownstreamRevision(v string)`

SetDownstreamRevision sets DownstreamRevision field to given value.

### HasDownstreamRevision

`func (o *ModelAPIPatchTriggerDefinition) HasDownstreamRevision() bool`

HasDownstreamRevision returns a boolean if a field has been set.

### GetParentAsModule

`func (o *ModelAPIPatchTriggerDefinition) GetParentAsModule() string`

GetParentAsModule returns the ParentAsModule field if non-nil, zero value otherwise.

### GetParentAsModuleOk

`func (o *ModelAPIPatchTriggerDefinition) GetParentAsModuleOk() (*string, bool)`

GetParentAsModuleOk returns a tuple with the ParentAsModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAsModule

`func (o *ModelAPIPatchTriggerDefinition) SetParentAsModule(v string)`

SetParentAsModule sets ParentAsModule field to given value.

### HasParentAsModule

`func (o *ModelAPIPatchTriggerDefinition) HasParentAsModule() bool`

HasParentAsModule returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPIPatchTriggerDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPIPatchTriggerDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPIPatchTriggerDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPIPatchTriggerDefinition) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskSpecifiers

`func (o *ModelAPIPatchTriggerDefinition) GetTaskSpecifiers() []ModelAPITaskSpecifier`

GetTaskSpecifiers returns the TaskSpecifiers field if non-nil, zero value otherwise.

### GetTaskSpecifiersOk

`func (o *ModelAPIPatchTriggerDefinition) GetTaskSpecifiersOk() (*[]ModelAPITaskSpecifier, bool)`

GetTaskSpecifiersOk returns a tuple with the TaskSpecifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSpecifiers

`func (o *ModelAPIPatchTriggerDefinition) SetTaskSpecifiers(v []ModelAPITaskSpecifier)`

SetTaskSpecifiers sets TaskSpecifiers field to given value.

### HasTaskSpecifiers

`func (o *ModelAPIPatchTriggerDefinition) HasTaskSpecifiers() bool`

HasTaskSpecifiers returns a boolean if a field has been set.

### GetVariantsTasks

`func (o *ModelAPIPatchTriggerDefinition) GetVariantsTasks() []ModelVariantTask`

GetVariantsTasks returns the VariantsTasks field if non-nil, zero value otherwise.

### GetVariantsTasksOk

`func (o *ModelAPIPatchTriggerDefinition) GetVariantsTasksOk() (*[]ModelVariantTask, bool)`

GetVariantsTasksOk returns a tuple with the VariantsTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantsTasks

`func (o *ModelAPIPatchTriggerDefinition) SetVariantsTasks(v []ModelVariantTask)`

SetVariantsTasks sets VariantsTasks field to given value.

### HasVariantsTasks

`func (o *ModelAPIPatchTriggerDefinition) HasVariantsTasks() bool`

HasVariantsTasks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


