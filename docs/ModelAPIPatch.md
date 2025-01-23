# ModelAPIPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** | Whether the patch has been finalized and activated | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** | Author of the patch | [optional] 
**Branch** | Pointer to **string** | The branch on which the patch was initiated | [optional] 
**Builds** | Pointer to **[]string** | List of identifiers of builds to run for this patch | [optional] 
**ChildPatchAliases** | Pointer to [**[]ModelAPIChildPatchAlias**](ModelAPIChildPatchAlias.md) |  | [optional] 
**ChildPatches** | Pointer to [**[]ModelAPIPatch**](ModelAPIPatch.md) |  | [optional] 
**CreateTime** | Pointer to **string** | Time patch was created | [optional] 
**Description** | Pointer to **string** | Description of the patch | [optional] 
**DownstreamTasks** | Pointer to [**[]ModelDownstreamTasks**](ModelDownstreamTasks.md) |  | [optional] 
**FinishTime** | Pointer to **string** | Time at patch completion | [optional] 
**GitHash** | Pointer to **string** | Hash of commit off which the patch was initiated | [optional] 
**GithubPatchData** | Pointer to [**ModelGithubPatch**](ModelGithubPatch.md) |  | [optional] 
**Hidden** | Pointer to **bool** |  | [optional] 
**MergedFrom** | Pointer to **string** |  | [optional] 
**ModuleCodeChanges** | Pointer to [**[]ModelAPIModulePatch**](ModelAPIModulePatch.md) |  | [optional] 
**Parameters** | Pointer to [**[]ModelAPIParameter**](ModelAPIParameter.md) |  | [optional] 
**PatchId** | Pointer to **string** | Unique identifier of a specific patch | [optional] 
**PatchNumber** | Pointer to **int32** | Incrementing counter of user&#39;s patches | [optional] 
**ProjectId** | Pointer to **string** | Name of the project | [optional] 
**ProjectIdentifier** | Pointer to **string** |  | [optional] 
**ProjectStorageMethod** | Pointer to **string** |  | [optional] 
**Requester** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** | Time patch started to run | [optional] 
**Status** | Pointer to **string** | Status of patch (possible values are \&quot;created\&quot;, \&quot;started\&quot;, \&quot;success\&quot;, or \&quot;failed\&quot;) | [optional] 
**Tasks** | Pointer to **[]string** | List of identifiers of tasks used in this patch | [optional] 
**VariantsTasks** | Pointer to [**[]ModelVariantTask**](ModelVariantTask.md) | List of documents of available tasks and associated build variant | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIPatch

`func NewModelAPIPatch() *ModelAPIPatch`

NewModelAPIPatch instantiates a new ModelAPIPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIPatchWithDefaults

`func NewModelAPIPatchWithDefaults() *ModelAPIPatch`

NewModelAPIPatchWithDefaults instantiates a new ModelAPIPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *ModelAPIPatch) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ModelAPIPatch) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ModelAPIPatch) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ModelAPIPatch) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetAlias

`func (o *ModelAPIPatch) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelAPIPatch) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelAPIPatch) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelAPIPatch) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetAuthor

`func (o *ModelAPIPatch) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelAPIPatch) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelAPIPatch) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelAPIPatch) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetBranch

`func (o *ModelAPIPatch) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ModelAPIPatch) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ModelAPIPatch) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ModelAPIPatch) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetBuilds

`func (o *ModelAPIPatch) GetBuilds() []string`

GetBuilds returns the Builds field if non-nil, zero value otherwise.

### GetBuildsOk

`func (o *ModelAPIPatch) GetBuildsOk() (*[]string, bool)`

GetBuildsOk returns a tuple with the Builds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilds

`func (o *ModelAPIPatch) SetBuilds(v []string)`

SetBuilds sets Builds field to given value.

### HasBuilds

`func (o *ModelAPIPatch) HasBuilds() bool`

HasBuilds returns a boolean if a field has been set.

### GetChildPatchAliases

`func (o *ModelAPIPatch) GetChildPatchAliases() []ModelAPIChildPatchAlias`

GetChildPatchAliases returns the ChildPatchAliases field if non-nil, zero value otherwise.

### GetChildPatchAliasesOk

`func (o *ModelAPIPatch) GetChildPatchAliasesOk() (*[]ModelAPIChildPatchAlias, bool)`

GetChildPatchAliasesOk returns a tuple with the ChildPatchAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildPatchAliases

`func (o *ModelAPIPatch) SetChildPatchAliases(v []ModelAPIChildPatchAlias)`

SetChildPatchAliases sets ChildPatchAliases field to given value.

### HasChildPatchAliases

`func (o *ModelAPIPatch) HasChildPatchAliases() bool`

HasChildPatchAliases returns a boolean if a field has been set.

### GetChildPatches

`func (o *ModelAPIPatch) GetChildPatches() []ModelAPIPatch`

GetChildPatches returns the ChildPatches field if non-nil, zero value otherwise.

### GetChildPatchesOk

`func (o *ModelAPIPatch) GetChildPatchesOk() (*[]ModelAPIPatch, bool)`

GetChildPatchesOk returns a tuple with the ChildPatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildPatches

`func (o *ModelAPIPatch) SetChildPatches(v []ModelAPIPatch)`

SetChildPatches sets ChildPatches field to given value.

### HasChildPatches

`func (o *ModelAPIPatch) HasChildPatches() bool`

HasChildPatches returns a boolean if a field has been set.

### GetCreateTime

`func (o *ModelAPIPatch) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ModelAPIPatch) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ModelAPIPatch) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ModelAPIPatch) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDescription

`func (o *ModelAPIPatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelAPIPatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelAPIPatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelAPIPatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDownstreamTasks

`func (o *ModelAPIPatch) GetDownstreamTasks() []ModelDownstreamTasks`

GetDownstreamTasks returns the DownstreamTasks field if non-nil, zero value otherwise.

### GetDownstreamTasksOk

`func (o *ModelAPIPatch) GetDownstreamTasksOk() (*[]ModelDownstreamTasks, bool)`

GetDownstreamTasksOk returns a tuple with the DownstreamTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamTasks

`func (o *ModelAPIPatch) SetDownstreamTasks(v []ModelDownstreamTasks)`

SetDownstreamTasks sets DownstreamTasks field to given value.

### HasDownstreamTasks

`func (o *ModelAPIPatch) HasDownstreamTasks() bool`

HasDownstreamTasks returns a boolean if a field has been set.

### GetFinishTime

`func (o *ModelAPIPatch) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ModelAPIPatch) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ModelAPIPatch) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ModelAPIPatch) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetGitHash

`func (o *ModelAPIPatch) GetGitHash() string`

GetGitHash returns the GitHash field if non-nil, zero value otherwise.

### GetGitHashOk

`func (o *ModelAPIPatch) GetGitHashOk() (*string, bool)`

GetGitHashOk returns a tuple with the GitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHash

`func (o *ModelAPIPatch) SetGitHash(v string)`

SetGitHash sets GitHash field to given value.

### HasGitHash

`func (o *ModelAPIPatch) HasGitHash() bool`

HasGitHash returns a boolean if a field has been set.

### GetGithubPatchData

`func (o *ModelAPIPatch) GetGithubPatchData() ModelGithubPatch`

GetGithubPatchData returns the GithubPatchData field if non-nil, zero value otherwise.

### GetGithubPatchDataOk

`func (o *ModelAPIPatch) GetGithubPatchDataOk() (*ModelGithubPatch, bool)`

GetGithubPatchDataOk returns a tuple with the GithubPatchData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubPatchData

`func (o *ModelAPIPatch) SetGithubPatchData(v ModelGithubPatch)`

SetGithubPatchData sets GithubPatchData field to given value.

### HasGithubPatchData

`func (o *ModelAPIPatch) HasGithubPatchData() bool`

HasGithubPatchData returns a boolean if a field has been set.

### GetHidden

`func (o *ModelAPIPatch) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ModelAPIPatch) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ModelAPIPatch) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ModelAPIPatch) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetMergedFrom

`func (o *ModelAPIPatch) GetMergedFrom() string`

GetMergedFrom returns the MergedFrom field if non-nil, zero value otherwise.

### GetMergedFromOk

`func (o *ModelAPIPatch) GetMergedFromOk() (*string, bool)`

GetMergedFromOk returns a tuple with the MergedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedFrom

`func (o *ModelAPIPatch) SetMergedFrom(v string)`

SetMergedFrom sets MergedFrom field to given value.

### HasMergedFrom

`func (o *ModelAPIPatch) HasMergedFrom() bool`

HasMergedFrom returns a boolean if a field has been set.

### GetModuleCodeChanges

`func (o *ModelAPIPatch) GetModuleCodeChanges() []ModelAPIModulePatch`

GetModuleCodeChanges returns the ModuleCodeChanges field if non-nil, zero value otherwise.

### GetModuleCodeChangesOk

`func (o *ModelAPIPatch) GetModuleCodeChangesOk() (*[]ModelAPIModulePatch, bool)`

GetModuleCodeChangesOk returns a tuple with the ModuleCodeChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleCodeChanges

`func (o *ModelAPIPatch) SetModuleCodeChanges(v []ModelAPIModulePatch)`

SetModuleCodeChanges sets ModuleCodeChanges field to given value.

### HasModuleCodeChanges

`func (o *ModelAPIPatch) HasModuleCodeChanges() bool`

HasModuleCodeChanges returns a boolean if a field has been set.

### GetParameters

`func (o *ModelAPIPatch) GetParameters() []ModelAPIParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ModelAPIPatch) GetParametersOk() (*[]ModelAPIParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ModelAPIPatch) SetParameters(v []ModelAPIParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ModelAPIPatch) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPatchId

`func (o *ModelAPIPatch) GetPatchId() string`

GetPatchId returns the PatchId field if non-nil, zero value otherwise.

### GetPatchIdOk

`func (o *ModelAPIPatch) GetPatchIdOk() (*string, bool)`

GetPatchIdOk returns a tuple with the PatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchId

`func (o *ModelAPIPatch) SetPatchId(v string)`

SetPatchId sets PatchId field to given value.

### HasPatchId

`func (o *ModelAPIPatch) HasPatchId() bool`

HasPatchId returns a boolean if a field has been set.

### GetPatchNumber

`func (o *ModelAPIPatch) GetPatchNumber() int32`

GetPatchNumber returns the PatchNumber field if non-nil, zero value otherwise.

### GetPatchNumberOk

`func (o *ModelAPIPatch) GetPatchNumberOk() (*int32, bool)`

GetPatchNumberOk returns a tuple with the PatchNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchNumber

`func (o *ModelAPIPatch) SetPatchNumber(v int32)`

SetPatchNumber sets PatchNumber field to given value.

### HasPatchNumber

`func (o *ModelAPIPatch) HasPatchNumber() bool`

HasPatchNumber returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelAPIPatch) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelAPIPatch) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelAPIPatch) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelAPIPatch) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectIdentifier

`func (o *ModelAPIPatch) GetProjectIdentifier() string`

GetProjectIdentifier returns the ProjectIdentifier field if non-nil, zero value otherwise.

### GetProjectIdentifierOk

`func (o *ModelAPIPatch) GetProjectIdentifierOk() (*string, bool)`

GetProjectIdentifierOk returns a tuple with the ProjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdentifier

`func (o *ModelAPIPatch) SetProjectIdentifier(v string)`

SetProjectIdentifier sets ProjectIdentifier field to given value.

### HasProjectIdentifier

`func (o *ModelAPIPatch) HasProjectIdentifier() bool`

HasProjectIdentifier returns a boolean if a field has been set.

### GetProjectStorageMethod

`func (o *ModelAPIPatch) GetProjectStorageMethod() string`

GetProjectStorageMethod returns the ProjectStorageMethod field if non-nil, zero value otherwise.

### GetProjectStorageMethodOk

`func (o *ModelAPIPatch) GetProjectStorageMethodOk() (*string, bool)`

GetProjectStorageMethodOk returns a tuple with the ProjectStorageMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectStorageMethod

`func (o *ModelAPIPatch) SetProjectStorageMethod(v string)`

SetProjectStorageMethod sets ProjectStorageMethod field to given value.

### HasProjectStorageMethod

`func (o *ModelAPIPatch) HasProjectStorageMethod() bool`

HasProjectStorageMethod returns a boolean if a field has been set.

### GetRequester

`func (o *ModelAPIPatch) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ModelAPIPatch) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ModelAPIPatch) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ModelAPIPatch) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelAPIPatch) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelAPIPatch) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelAPIPatch) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelAPIPatch) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPIPatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPIPatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPIPatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPIPatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTasks

`func (o *ModelAPIPatch) GetTasks() []string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ModelAPIPatch) GetTasksOk() (*[]string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ModelAPIPatch) SetTasks(v []string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ModelAPIPatch) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetVariantsTasks

`func (o *ModelAPIPatch) GetVariantsTasks() []ModelVariantTask`

GetVariantsTasks returns the VariantsTasks field if non-nil, zero value otherwise.

### GetVariantsTasksOk

`func (o *ModelAPIPatch) GetVariantsTasksOk() (*[]ModelVariantTask, bool)`

GetVariantsTasksOk returns a tuple with the VariantsTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantsTasks

`func (o *ModelAPIPatch) SetVariantsTasks(v []ModelVariantTask)`

SetVariantsTasks sets VariantsTasks field to given value.

### HasVariantsTasks

`func (o *ModelAPIPatch) HasVariantsTasks() bool`

HasVariantsTasks returns a boolean if a field has been set.

### GetVersion

`func (o *ModelAPIPatch) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelAPIPatch) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelAPIPatch) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelAPIPatch) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


