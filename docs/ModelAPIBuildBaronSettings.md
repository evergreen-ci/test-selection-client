# ModelAPIBuildBaronSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BfSuggestionFeaturesUrl** | Pointer to **string** |  | [optional] 
**BfSuggestionPassword** | Pointer to **string** |  | [optional] 
**BfSuggestionServer** | Pointer to **string** |  | [optional] 
**BfSuggestionTimeoutSecs** | Pointer to **int32** |  | [optional] 
**BfSuggestionUsername** | Pointer to **string** |  | [optional] 
**TicketCreateIssueType** | Pointer to **string** | Type of ticket to create. | [optional] 
**TicketCreateProject** | Pointer to **string** | Jira project where tickets should be created. | [optional] 
**TicketSearchProjects** | Pointer to **[]string** | Jira project to search for tickets. | [optional] 

## Methods

### NewModelAPIBuildBaronSettings

`func NewModelAPIBuildBaronSettings() *ModelAPIBuildBaronSettings`

NewModelAPIBuildBaronSettings instantiates a new ModelAPIBuildBaronSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIBuildBaronSettingsWithDefaults

`func NewModelAPIBuildBaronSettingsWithDefaults() *ModelAPIBuildBaronSettings`

NewModelAPIBuildBaronSettingsWithDefaults instantiates a new ModelAPIBuildBaronSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBfSuggestionFeaturesUrl

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionFeaturesUrl() string`

GetBfSuggestionFeaturesUrl returns the BfSuggestionFeaturesUrl field if non-nil, zero value otherwise.

### GetBfSuggestionFeaturesUrlOk

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionFeaturesUrlOk() (*string, bool)`

GetBfSuggestionFeaturesUrlOk returns a tuple with the BfSuggestionFeaturesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfSuggestionFeaturesUrl

`func (o *ModelAPIBuildBaronSettings) SetBfSuggestionFeaturesUrl(v string)`

SetBfSuggestionFeaturesUrl sets BfSuggestionFeaturesUrl field to given value.

### HasBfSuggestionFeaturesUrl

`func (o *ModelAPIBuildBaronSettings) HasBfSuggestionFeaturesUrl() bool`

HasBfSuggestionFeaturesUrl returns a boolean if a field has been set.

### GetBfSuggestionPassword

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionPassword() string`

GetBfSuggestionPassword returns the BfSuggestionPassword field if non-nil, zero value otherwise.

### GetBfSuggestionPasswordOk

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionPasswordOk() (*string, bool)`

GetBfSuggestionPasswordOk returns a tuple with the BfSuggestionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfSuggestionPassword

`func (o *ModelAPIBuildBaronSettings) SetBfSuggestionPassword(v string)`

SetBfSuggestionPassword sets BfSuggestionPassword field to given value.

### HasBfSuggestionPassword

`func (o *ModelAPIBuildBaronSettings) HasBfSuggestionPassword() bool`

HasBfSuggestionPassword returns a boolean if a field has been set.

### GetBfSuggestionServer

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionServer() string`

GetBfSuggestionServer returns the BfSuggestionServer field if non-nil, zero value otherwise.

### GetBfSuggestionServerOk

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionServerOk() (*string, bool)`

GetBfSuggestionServerOk returns a tuple with the BfSuggestionServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfSuggestionServer

`func (o *ModelAPIBuildBaronSettings) SetBfSuggestionServer(v string)`

SetBfSuggestionServer sets BfSuggestionServer field to given value.

### HasBfSuggestionServer

`func (o *ModelAPIBuildBaronSettings) HasBfSuggestionServer() bool`

HasBfSuggestionServer returns a boolean if a field has been set.

### GetBfSuggestionTimeoutSecs

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionTimeoutSecs() int32`

GetBfSuggestionTimeoutSecs returns the BfSuggestionTimeoutSecs field if non-nil, zero value otherwise.

### GetBfSuggestionTimeoutSecsOk

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionTimeoutSecsOk() (*int32, bool)`

GetBfSuggestionTimeoutSecsOk returns a tuple with the BfSuggestionTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfSuggestionTimeoutSecs

`func (o *ModelAPIBuildBaronSettings) SetBfSuggestionTimeoutSecs(v int32)`

SetBfSuggestionTimeoutSecs sets BfSuggestionTimeoutSecs field to given value.

### HasBfSuggestionTimeoutSecs

`func (o *ModelAPIBuildBaronSettings) HasBfSuggestionTimeoutSecs() bool`

HasBfSuggestionTimeoutSecs returns a boolean if a field has been set.

### GetBfSuggestionUsername

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionUsername() string`

GetBfSuggestionUsername returns the BfSuggestionUsername field if non-nil, zero value otherwise.

### GetBfSuggestionUsernameOk

`func (o *ModelAPIBuildBaronSettings) GetBfSuggestionUsernameOk() (*string, bool)`

GetBfSuggestionUsernameOk returns a tuple with the BfSuggestionUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBfSuggestionUsername

`func (o *ModelAPIBuildBaronSettings) SetBfSuggestionUsername(v string)`

SetBfSuggestionUsername sets BfSuggestionUsername field to given value.

### HasBfSuggestionUsername

`func (o *ModelAPIBuildBaronSettings) HasBfSuggestionUsername() bool`

HasBfSuggestionUsername returns a boolean if a field has been set.

### GetTicketCreateIssueType

`func (o *ModelAPIBuildBaronSettings) GetTicketCreateIssueType() string`

GetTicketCreateIssueType returns the TicketCreateIssueType field if non-nil, zero value otherwise.

### GetTicketCreateIssueTypeOk

`func (o *ModelAPIBuildBaronSettings) GetTicketCreateIssueTypeOk() (*string, bool)`

GetTicketCreateIssueTypeOk returns a tuple with the TicketCreateIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketCreateIssueType

`func (o *ModelAPIBuildBaronSettings) SetTicketCreateIssueType(v string)`

SetTicketCreateIssueType sets TicketCreateIssueType field to given value.

### HasTicketCreateIssueType

`func (o *ModelAPIBuildBaronSettings) HasTicketCreateIssueType() bool`

HasTicketCreateIssueType returns a boolean if a field has been set.

### GetTicketCreateProject

`func (o *ModelAPIBuildBaronSettings) GetTicketCreateProject() string`

GetTicketCreateProject returns the TicketCreateProject field if non-nil, zero value otherwise.

### GetTicketCreateProjectOk

`func (o *ModelAPIBuildBaronSettings) GetTicketCreateProjectOk() (*string, bool)`

GetTicketCreateProjectOk returns a tuple with the TicketCreateProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketCreateProject

`func (o *ModelAPIBuildBaronSettings) SetTicketCreateProject(v string)`

SetTicketCreateProject sets TicketCreateProject field to given value.

### HasTicketCreateProject

`func (o *ModelAPIBuildBaronSettings) HasTicketCreateProject() bool`

HasTicketCreateProject returns a boolean if a field has been set.

### GetTicketSearchProjects

`func (o *ModelAPIBuildBaronSettings) GetTicketSearchProjects() []string`

GetTicketSearchProjects returns the TicketSearchProjects field if non-nil, zero value otherwise.

### GetTicketSearchProjectsOk

`func (o *ModelAPIBuildBaronSettings) GetTicketSearchProjectsOk() (*[]string, bool)`

GetTicketSearchProjectsOk returns a tuple with the TicketSearchProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTicketSearchProjects

`func (o *ModelAPIBuildBaronSettings) SetTicketSearchProjects(v []string)`

SetTicketSearchProjects sets TicketSearchProjects field to given value.

### HasTicketSearchProjects

`func (o *ModelAPIBuildBaronSettings) HasTicketSearchProjects() bool`

HasTicketSearchProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


