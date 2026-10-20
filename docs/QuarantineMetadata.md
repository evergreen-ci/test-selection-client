# QuarantineMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Justification** | Pointer to **NullableString** |  | [optional] 
**JiraTicketsToLinkTo** | Pointer to **[]string** |  | [optional] 
**TaskInitiatedFrom** | Pointer to **NullableString** |  | [optional] 
**JiraAssignTo** | Pointer to **NullableString** |  | [optional] 
**InitiatedBy** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewQuarantineMetadata

`func NewQuarantineMetadata() *QuarantineMetadata`

NewQuarantineMetadata instantiates a new QuarantineMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuarantineMetadataWithDefaults

`func NewQuarantineMetadataWithDefaults() *QuarantineMetadata`

NewQuarantineMetadataWithDefaults instantiates a new QuarantineMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJustification

`func (o *QuarantineMetadata) GetJustification() string`

GetJustification returns the Justification field if non-nil, zero value otherwise.

### GetJustificationOk

`func (o *QuarantineMetadata) GetJustificationOk() (*string, bool)`

GetJustificationOk returns a tuple with the Justification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJustification

`func (o *QuarantineMetadata) SetJustification(v string)`

SetJustification sets Justification field to given value.

### HasJustification

`func (o *QuarantineMetadata) HasJustification() bool`

HasJustification returns a boolean if a field has been set.

### SetJustificationNil

`func (o *QuarantineMetadata) SetJustificationNil(b bool)`

 SetJustificationNil sets the value for Justification to be an explicit nil

### UnsetJustification
`func (o *QuarantineMetadata) UnsetJustification()`

UnsetJustification ensures that no value is present for Justification, not even an explicit nil
### GetJiraTicketsToLinkTo

`func (o *QuarantineMetadata) GetJiraTicketsToLinkTo() []string`

GetJiraTicketsToLinkTo returns the JiraTicketsToLinkTo field if non-nil, zero value otherwise.

### GetJiraTicketsToLinkToOk

`func (o *QuarantineMetadata) GetJiraTicketsToLinkToOk() (*[]string, bool)`

GetJiraTicketsToLinkToOk returns a tuple with the JiraTicketsToLinkTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraTicketsToLinkTo

`func (o *QuarantineMetadata) SetJiraTicketsToLinkTo(v []string)`

SetJiraTicketsToLinkTo sets JiraTicketsToLinkTo field to given value.

### HasJiraTicketsToLinkTo

`func (o *QuarantineMetadata) HasJiraTicketsToLinkTo() bool`

HasJiraTicketsToLinkTo returns a boolean if a field has been set.

### SetJiraTicketsToLinkToNil

`func (o *QuarantineMetadata) SetJiraTicketsToLinkToNil(b bool)`

 SetJiraTicketsToLinkToNil sets the value for JiraTicketsToLinkTo to be an explicit nil

### UnsetJiraTicketsToLinkTo
`func (o *QuarantineMetadata) UnsetJiraTicketsToLinkTo()`

UnsetJiraTicketsToLinkTo ensures that no value is present for JiraTicketsToLinkTo, not even an explicit nil
### GetTaskInitiatedFrom

`func (o *QuarantineMetadata) GetTaskInitiatedFrom() string`

GetTaskInitiatedFrom returns the TaskInitiatedFrom field if non-nil, zero value otherwise.

### GetTaskInitiatedFromOk

`func (o *QuarantineMetadata) GetTaskInitiatedFromOk() (*string, bool)`

GetTaskInitiatedFromOk returns a tuple with the TaskInitiatedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskInitiatedFrom

`func (o *QuarantineMetadata) SetTaskInitiatedFrom(v string)`

SetTaskInitiatedFrom sets TaskInitiatedFrom field to given value.

### HasTaskInitiatedFrom

`func (o *QuarantineMetadata) HasTaskInitiatedFrom() bool`

HasTaskInitiatedFrom returns a boolean if a field has been set.

### SetTaskInitiatedFromNil

`func (o *QuarantineMetadata) SetTaskInitiatedFromNil(b bool)`

 SetTaskInitiatedFromNil sets the value for TaskInitiatedFrom to be an explicit nil

### UnsetTaskInitiatedFrom
`func (o *QuarantineMetadata) UnsetTaskInitiatedFrom()`

UnsetTaskInitiatedFrom ensures that no value is present for TaskInitiatedFrom, not even an explicit nil
### GetJiraAssignTo

`func (o *QuarantineMetadata) GetJiraAssignTo() string`

GetJiraAssignTo returns the JiraAssignTo field if non-nil, zero value otherwise.

### GetJiraAssignToOk

`func (o *QuarantineMetadata) GetJiraAssignToOk() (*string, bool)`

GetJiraAssignToOk returns a tuple with the JiraAssignTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJiraAssignTo

`func (o *QuarantineMetadata) SetJiraAssignTo(v string)`

SetJiraAssignTo sets JiraAssignTo field to given value.

### HasJiraAssignTo

`func (o *QuarantineMetadata) HasJiraAssignTo() bool`

HasJiraAssignTo returns a boolean if a field has been set.

### SetJiraAssignToNil

`func (o *QuarantineMetadata) SetJiraAssignToNil(b bool)`

 SetJiraAssignToNil sets the value for JiraAssignTo to be an explicit nil

### UnsetJiraAssignTo
`func (o *QuarantineMetadata) UnsetJiraAssignTo()`

UnsetJiraAssignTo ensures that no value is present for JiraAssignTo, not even an explicit nil
### GetInitiatedBy

`func (o *QuarantineMetadata) GetInitiatedBy() string`

GetInitiatedBy returns the InitiatedBy field if non-nil, zero value otherwise.

### GetInitiatedByOk

`func (o *QuarantineMetadata) GetInitiatedByOk() (*string, bool)`

GetInitiatedByOk returns a tuple with the InitiatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatedBy

`func (o *QuarantineMetadata) SetInitiatedBy(v string)`

SetInitiatedBy sets InitiatedBy field to given value.

### HasInitiatedBy

`func (o *QuarantineMetadata) HasInitiatedBy() bool`

HasInitiatedBy returns a boolean if a field has been set.

### SetInitiatedByNil

`func (o *QuarantineMetadata) SetInitiatedByNil(b bool)`

 SetInitiatedByNil sets the value for InitiatedBy to be an explicit nil

### UnsetInitiatedBy
`func (o *QuarantineMetadata) UnsetInitiatedBy()`

UnsetInitiatedBy ensures that no value is present for InitiatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


