# ModelAPITaskAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedIssues** | Pointer to [**[]ModelAPIIssueLink**](ModelAPIIssueLink.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Issues** | Pointer to [**[]ModelAPIIssueLink**](ModelAPIIssueLink.md) | Links to tickets definitely related | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Structured data about the task. Since this is user-given json data, the structure can differ between annotations | [optional] 
**MetadataLinks** | Pointer to [**[]ModelAPIMetadataLink**](ModelAPIMetadataLink.md) | List of links associated with a task, to be displayed in the task metadata sidebar, currently limited to 1 | [optional] 
**Note** | Pointer to [**ModelAPINote**](ModelAPINote.md) | Comment about the task failure | [optional] 
**SuspectedIssues** | Pointer to [**[]ModelAPIIssueLink**](ModelAPIIssueLink.md) | Links to tickets possibly related | [optional] 
**TaskExecution** | Pointer to **int32** | The number of the execution of the task that the annotation is for | [optional] 
**TaskId** | Pointer to **string** | Identifier of the task that this annotation is for | [optional] 

## Methods

### NewModelAPITaskAnnotation

`func NewModelAPITaskAnnotation() *ModelAPITaskAnnotation`

NewModelAPITaskAnnotation instantiates a new ModelAPITaskAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITaskAnnotationWithDefaults

`func NewModelAPITaskAnnotationWithDefaults() *ModelAPITaskAnnotation`

NewModelAPITaskAnnotationWithDefaults instantiates a new ModelAPITaskAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedIssues

`func (o *ModelAPITaskAnnotation) GetCreatedIssues() []ModelAPIIssueLink`

GetCreatedIssues returns the CreatedIssues field if non-nil, zero value otherwise.

### GetCreatedIssuesOk

`func (o *ModelAPITaskAnnotation) GetCreatedIssuesOk() (*[]ModelAPIIssueLink, bool)`

GetCreatedIssuesOk returns a tuple with the CreatedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedIssues

`func (o *ModelAPITaskAnnotation) SetCreatedIssues(v []ModelAPIIssueLink)`

SetCreatedIssues sets CreatedIssues field to given value.

### HasCreatedIssues

`func (o *ModelAPITaskAnnotation) HasCreatedIssues() bool`

HasCreatedIssues returns a boolean if a field has been set.

### GetId

`func (o *ModelAPITaskAnnotation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPITaskAnnotation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPITaskAnnotation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPITaskAnnotation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssues

`func (o *ModelAPITaskAnnotation) GetIssues() []ModelAPIIssueLink`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *ModelAPITaskAnnotation) GetIssuesOk() (*[]ModelAPIIssueLink, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *ModelAPITaskAnnotation) SetIssues(v []ModelAPIIssueLink)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *ModelAPITaskAnnotation) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelAPITaskAnnotation) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelAPITaskAnnotation) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelAPITaskAnnotation) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelAPITaskAnnotation) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetadataLinks

`func (o *ModelAPITaskAnnotation) GetMetadataLinks() []ModelAPIMetadataLink`

GetMetadataLinks returns the MetadataLinks field if non-nil, zero value otherwise.

### GetMetadataLinksOk

`func (o *ModelAPITaskAnnotation) GetMetadataLinksOk() (*[]ModelAPIMetadataLink, bool)`

GetMetadataLinksOk returns a tuple with the MetadataLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataLinks

`func (o *ModelAPITaskAnnotation) SetMetadataLinks(v []ModelAPIMetadataLink)`

SetMetadataLinks sets MetadataLinks field to given value.

### HasMetadataLinks

`func (o *ModelAPITaskAnnotation) HasMetadataLinks() bool`

HasMetadataLinks returns a boolean if a field has been set.

### GetNote

`func (o *ModelAPITaskAnnotation) GetNote() ModelAPINote`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ModelAPITaskAnnotation) GetNoteOk() (*ModelAPINote, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ModelAPITaskAnnotation) SetNote(v ModelAPINote)`

SetNote sets Note field to given value.

### HasNote

`func (o *ModelAPITaskAnnotation) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetSuspectedIssues

`func (o *ModelAPITaskAnnotation) GetSuspectedIssues() []ModelAPIIssueLink`

GetSuspectedIssues returns the SuspectedIssues field if non-nil, zero value otherwise.

### GetSuspectedIssuesOk

`func (o *ModelAPITaskAnnotation) GetSuspectedIssuesOk() (*[]ModelAPIIssueLink, bool)`

GetSuspectedIssuesOk returns a tuple with the SuspectedIssues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspectedIssues

`func (o *ModelAPITaskAnnotation) SetSuspectedIssues(v []ModelAPIIssueLink)`

SetSuspectedIssues sets SuspectedIssues field to given value.

### HasSuspectedIssues

`func (o *ModelAPITaskAnnotation) HasSuspectedIssues() bool`

HasSuspectedIssues returns a boolean if a field has been set.

### GetTaskExecution

`func (o *ModelAPITaskAnnotation) GetTaskExecution() int32`

GetTaskExecution returns the TaskExecution field if non-nil, zero value otherwise.

### GetTaskExecutionOk

`func (o *ModelAPITaskAnnotation) GetTaskExecutionOk() (*int32, bool)`

GetTaskExecutionOk returns a tuple with the TaskExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskExecution

`func (o *ModelAPITaskAnnotation) SetTaskExecution(v int32)`

SetTaskExecution sets TaskExecution field to given value.

### HasTaskExecution

`func (o *ModelAPITaskAnnotation) HasTaskExecution() bool`

HasTaskExecution returns a boolean if a field has been set.

### GetTaskId

`func (o *ModelAPITaskAnnotation) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelAPITaskAnnotation) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelAPITaskAnnotation) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelAPITaskAnnotation) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


