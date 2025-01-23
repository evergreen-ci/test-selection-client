# ModelAPIVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aborted** | Pointer to **bool** |  | [optional] 
**Activated** | Pointer to **bool** | Will be null for versions created before this field was added. | [optional] 
**Author** | Pointer to **string** | Author of the version | [optional] 
**AuthorEmail** | Pointer to **string** | Email of the author of the version | [optional] 
**Branch** | Pointer to **string** | The version control branch where the commit was made | [optional] 
**BuildVariantsStatus** | Pointer to [**[]ModelBuildDetail**](ModelBuildDetail.md) | List of documents of the associated build variant and the build id | [optional] 
**Builds** | Pointer to [**[]ModelAPIBuild**](ModelAPIBuild.md) |  | [optional] 
**CreateTime** | Pointer to **string** | Time that the version was first created | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 
**FinishTime** | Pointer to **string** | Time at which tasks associated with this version finished running | [optional] 
**GitTags** | Pointer to [**[]ModelAPIGitTag**](ModelAPIGitTag.md) | Git tags that were pushed to this version. | [optional] 
**Ignored** | Pointer to **bool** | Indicates if the version was ignored due to only making changes to ignored files. | [optional] 
**Message** | Pointer to **string** | Message left with the commit | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to [**[]ModelAPIParameter**](ModelAPIParameter.md) |  | [optional] 
**Project** | Pointer to **string** |  | [optional] 
**ProjectIdentifier** | Pointer to **string** |  | [optional] 
**Repo** | Pointer to **string** | The github repository where the commit was made | [optional] 
**Requester** | Pointer to **string** | Version created by one of \&quot;patch_request\&quot;, \&quot;github_pull_request\&quot;, \&quot;gitter_request\&quot; (caused by git commit, aka the repotracker requester), \&quot;trigger_request\&quot; (Project Trigger versions) , \&quot;github_merge_request\&quot; (GitHub merge queue), \&quot;ad_hoc\&quot; (periodic builds) | [optional] 
**Revision** | Pointer to **string** | The version control identifier | [optional] 
**StartTime** | Pointer to **string** | Time at which tasks associated with this version started running | [optional] 
**Status** | Pointer to **string** | The status of the version (possible values are \&quot;created\&quot;, \&quot;started\&quot;, \&quot;success\&quot;, or \&quot;failed\&quot;) | [optional] 
**TriggeredByGitTag** | Pointer to [**ModelAPIGitTag**](ModelAPIGitTag.md) | The git tag that triggered this version, if any. | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIVersion

`func NewModelAPIVersion() *ModelAPIVersion`

NewModelAPIVersion instantiates a new ModelAPIVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIVersionWithDefaults

`func NewModelAPIVersionWithDefaults() *ModelAPIVersion`

NewModelAPIVersionWithDefaults instantiates a new ModelAPIVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAborted

`func (o *ModelAPIVersion) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *ModelAPIVersion) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *ModelAPIVersion) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *ModelAPIVersion) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetActivated

`func (o *ModelAPIVersion) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ModelAPIVersion) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ModelAPIVersion) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ModelAPIVersion) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetAuthor

`func (o *ModelAPIVersion) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelAPIVersion) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelAPIVersion) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelAPIVersion) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorEmail

`func (o *ModelAPIVersion) GetAuthorEmail() string`

GetAuthorEmail returns the AuthorEmail field if non-nil, zero value otherwise.

### GetAuthorEmailOk

`func (o *ModelAPIVersion) GetAuthorEmailOk() (*string, bool)`

GetAuthorEmailOk returns a tuple with the AuthorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorEmail

`func (o *ModelAPIVersion) SetAuthorEmail(v string)`

SetAuthorEmail sets AuthorEmail field to given value.

### HasAuthorEmail

`func (o *ModelAPIVersion) HasAuthorEmail() bool`

HasAuthorEmail returns a boolean if a field has been set.

### GetBranch

`func (o *ModelAPIVersion) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ModelAPIVersion) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ModelAPIVersion) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ModelAPIVersion) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuildVariantsStatus

`func (o *ModelAPIVersion) GetBuildVariantsStatus() []ModelBuildDetail`

GetBuildVariantsStatus returns the BuildVariantsStatus field if non-nil, zero value otherwise.

### GetBuildVariantsStatusOk

`func (o *ModelAPIVersion) GetBuildVariantsStatusOk() (*[]ModelBuildDetail, bool)`

GetBuildVariantsStatusOk returns a tuple with the BuildVariantsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVariantsStatus

`func (o *ModelAPIVersion) SetBuildVariantsStatus(v []ModelBuildDetail)`

SetBuildVariantsStatus sets BuildVariantsStatus field to given value.

### HasBuildVariantsStatus

`func (o *ModelAPIVersion) HasBuildVariantsStatus() bool`

HasBuildVariantsStatus returns a boolean if a field has been set.

### GetBuilds

`func (o *ModelAPIVersion) GetBuilds() []ModelAPIBuild`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *ModelAPIVersion) GetBuildsOk() (*[]ModelAPIBuild, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *ModelAPIVersion) SetBuilds(v []ModelAPIBuild)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *ModelAPIVersion) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetCreateTime

`func (o *ModelAPIVersion) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ModelAPIVersion) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ModelAPIVersion) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ModelAPIVersion) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetErrors

`func (o *ModelAPIVersion) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ModelAPIVersion) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ModelAPIVersion) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ModelAPIVersion) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFinishTime

`func (o *ModelAPIVersion) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ModelAPIVersion) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ModelAPIVersion) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ModelAPIVersion) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetGitTags

`func (o *ModelAPIVersion) GetGitTags() []ModelAPIGitTag`

GetGitTags returns the GitTags field if non-nil, zero value otherwise.

### GetGitTagsOk

`func (o *ModelAPIVersion) GetGitTagsOk() (*[]ModelAPIGitTag, bool)`

GetGitTagsOk returns a tuple with the GitTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTags

`func (o *ModelAPIVersion) SetGitTags(v []ModelAPIGitTag)`

SetGitTags sets GitTags field to given value.

### HasGitTags

`func (o *ModelAPIVersion) HasGitTags() bool`

HasGitTags returns a boolean if a field has been set.

### GetIgnored

`func (o *ModelAPIVersion) GetIgnored() bool`

GetIgnored returns the Ignored field if non-nil, zero value otherwise.

### GetIgnoredOk

`func (o *ModelAPIVersion) GetIgnoredOk() (*bool, bool)`

GetIgnoredOk returns a tuple with the Ignored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnored

`func (o *ModelAPIVersion) SetIgnored(v bool)`

SetIgnored sets Ignored field to given value.

### HasIgnored

`func (o *ModelAPIVersion) HasIgnored() bool`

HasIgnored returns a boolean if a field has been set.

### GetMessage

`func (o *ModelAPIVersion) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelAPIVersion) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelAPIVersion) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelAPIVersion) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOrder

`func (o *ModelAPIVersion) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelAPIVersion) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelAPIVersion) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelAPIVersion) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParameters

`func (o *ModelAPIVersion) GetParameters() []ModelAPIParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ModelAPIVersion) GetParametersOk() (*[]ModelAPIParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ModelAPIVersion) SetParameters(v []ModelAPIParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ModelAPIVersion) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetProject

`func (o *ModelAPIVersion) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ModelAPIVersion) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ModelAPIVersion) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ModelAPIVersion) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetProjectIdentifier

`func (o *ModelAPIVersion) GetProjectIdentifier() string`

GetProjectIdentifier returns the ProjectIdentifier field if non-nil, zero value otherwise.

### GetProjectIdentifierOk

`func (o *ModelAPIVersion) GetProjectIdentifierOk() (*string, bool)`

GetProjectIdentifierOk returns a tuple with the ProjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdentifier

`func (o *ModelAPIVersion) SetProjectIdentifier(v string)`

SetProjectIdentifier sets ProjectIdentifier field to given value.

### HasProjectIdentifier

`func (o *ModelAPIVersion) HasProjectIdentifier() bool`

HasProjectIdentifier returns a boolean if a field has been set.

### GetRepo

`func (o *ModelAPIVersion) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ModelAPIVersion) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ModelAPIVersion) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ModelAPIVersion) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRequester

`func (o *ModelAPIVersion) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ModelAPIVersion) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ModelAPIVersion) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ModelAPIVersion) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetRevision

`func (o *ModelAPIVersion) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ModelAPIVersion) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ModelAPIVersion) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ModelAPIVersion) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelAPIVersion) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelAPIVersion) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelAPIVersion) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelAPIVersion) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPIVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPIVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPIVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPIVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTriggeredByGitTag

`func (o *ModelAPIVersion) GetTriggeredByGitTag() ModelAPIGitTag`

GetTriggeredByGitTag returns the TriggeredByGitTag field if non-nil, zero value otherwise.

### GetTriggeredByGitTagOk

`func (o *ModelAPIVersion) GetTriggeredByGitTagOk() (*ModelAPIGitTag, bool)`

GetTriggeredByGitTagOk returns a tuple with the TriggeredByGitTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredByGitTag

`func (o *ModelAPIVersion) SetTriggeredByGitTag(v ModelAPIGitTag)`

SetTriggeredByGitTag sets TriggeredByGitTag field to given value.

### HasTriggeredByGitTag

`func (o *ModelAPIVersion) HasTriggeredByGitTag() bool`

HasTriggeredByGitTag returns a boolean if a field has been set.

### GetVersionId

`func (o *ModelAPIVersion) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ModelAPIVersion) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ModelAPIVersion) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ModelAPIVersion) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


