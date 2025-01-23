# ModelAPIBuild

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Activated** | Pointer to **bool** | Whether this build was manually initiated | [optional] 
**ActivatedBy** | Pointer to **string** | Who initiated the build | [optional] 
**ActivatedTime** | Pointer to **string** | When the build was initiated | [optional] 
**ActualMakespanMs** | Pointer to **int32** | Actual makespan measured during execution | [optional] 
**BuildVariant** | Pointer to **string** | Build distro and architecture information | [optional] 
**CreateTime** | Pointer to **string** | Time at which build was created | [optional] 
**DefinitionInfo** | Pointer to [**ModelDefinitionInfo**](ModelDefinitionInfo.md) | Some routes will return information about the variant as defined in the project. Does not expand expansions; they will be returned as written in the project yaml (i.e. ${syntax}) | [optional] 
**DisplayName** | Pointer to **string** | Displayed title of the build showing version and variant running | [optional] 
**FinishTime** | Pointer to **string** | Time at which build finished running all tasks | [optional] 
**GitHash** | Pointer to **string** | Hash of the revision on which this build is running | [optional] 
**Order** | Pointer to **int32** | Incrementing counter of project&#39;s builds | [optional] 
**Origin** | Pointer to **string** | The source of the patch, a commit or a patch | [optional] 
**PredictedMakespanMs** | Pointer to **int32** | Predicted makespan by the scheduler prior to execution | [optional] 
**ProjectId** | Pointer to **string** | The identifier of the project this build represents | [optional] 
**ProjectIdentifier** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** | Time at which build started running tasks | [optional] 
**Status** | Pointer to **string** | The status of the build (possible values are \&quot;created\&quot;, \&quot;started\&quot;, \&quot;success\&quot;, or \&quot;failed\&quot;) | [optional] 
**StatusCounts** | Pointer to [**TaskTaskStatusCount**](TaskTaskStatusCount.md) | Contains aggregated data about the statuses of tasks in this build. The keys of this object are statuses and the values are the number of tasks within this build in that status. Note that this field provides data that you can get yourself by querying tasks for this build. | [optional] 
**Tags** | Pointer to **[]string** | List of tags defined for the build variant, if any | [optional] 
**TaskCache** | Pointer to [**[]ModelAPITaskCache**](ModelAPITaskCache.md) | Contains a subset of information about tasks for the build; this is not provided/accurate for most routes (get versions for project is an exception). | [optional] 
**Tasks** | Pointer to **[]string** | Tasks is the build&#39;s task cache with just the names | [optional] 
**TimeTakenMs** | Pointer to **int32** | How long the build took to complete all tasks | [optional] 
**Version** | Pointer to **string** | The version this build is running tasks for | [optional] 

## Methods

### NewModelAPIBuild

`func NewModelAPIBuild() *ModelAPIBuild`

NewModelAPIBuild instantiates a new ModelAPIBuild object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIBuildWithDefaults

`func NewModelAPIBuildWithDefaults() *ModelAPIBuild`

NewModelAPIBuildWithDefaults instantiates a new ModelAPIBuild object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelAPIBuild) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPIBuild) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPIBuild) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPIBuild) HasId() bool`

HasId returns a boolean if a field has been set.

### GetActivated

`func (o *ModelAPIBuild) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ModelAPIBuild) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ModelAPIBuild) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ModelAPIBuild) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivatedBy

`func (o *ModelAPIBuild) GetActivatedBy() string`

GetActivatedBy returns the ActivatedBy field if non-nil, zero value otherwise.

### GetActivatedByOk

`func (o *ModelAPIBuild) GetActivatedByOk() (*string, bool)`

GetActivatedByOk returns a tuple with the ActivatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedBy

`func (o *ModelAPIBuild) SetActivatedBy(v string)`

SetActivatedBy sets ActivatedBy field to given value.

### HasActivatedBy

`func (o *ModelAPIBuild) HasActivatedBy() bool`

HasActivatedBy returns a boolean if a field has been set.

### GetActivatedTime

`func (o *ModelAPIBuild) GetActivatedTime() string`

GetActivatedTime returns the ActivatedTime field if non-nil, zero value otherwise.

### GetActivatedTimeOk

`func (o *ModelAPIBuild) GetActivatedTimeOk() (*string, bool)`

GetActivatedTimeOk returns a tuple with the ActivatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedTime

`func (o *ModelAPIBuild) SetActivatedTime(v string)`

SetActivatedTime sets ActivatedTime field to given value.

### HasActivatedTime

`func (o *ModelAPIBuild) HasActivatedTime() bool`

HasActivatedTime returns a boolean if a field has been set.

### GetActualMakespanMs

`func (o *ModelAPIBuild) GetActualMakespanMs() int32`

GetActualMakespanMs returns the ActualMakespanMs field if non-nil, zero value otherwise.

### GetActualMakespanMsOk

`func (o *ModelAPIBuild) GetActualMakespanMsOk() (*int32, bool)`

GetActualMakespanMsOk returns a tuple with the ActualMakespanMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMakespanMs

`func (o *ModelAPIBuild) SetActualMakespanMs(v int32)`

SetActualMakespanMs sets ActualMakespanMs field to given value.

### HasActualMakespanMs

`func (o *ModelAPIBuild) HasActualMakespanMs() bool`

HasActualMakespanMs returns a boolean if a field has been set.

### GetBuildVariant

`func (o *ModelAPIBuild) GetBuildVariant() string`

GetBuildVariant returns the BuildVariant field if non-nil, zero value otherwise.

### GetBuildVariantOk

`func (o *ModelAPIBuild) GetBuildVariantOk() (*string, bool)`

GetBuildVariantOk returns a tuple with the BuildVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVariant

`func (o *ModelAPIBuild) SetBuildVariant(v string)`

SetBuildVariant sets BuildVariant field to given value.

### HasBuildVariant

`func (o *ModelAPIBuild) HasBuildVariant() bool`

HasBuildVariant returns a boolean if a field has been set.

### GetCreateTime

`func (o *ModelAPIBuild) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ModelAPIBuild) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ModelAPIBuild) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ModelAPIBuild) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDefinitionInfo

`func (o *ModelAPIBuild) GetDefinitionInfo() ModelDefinitionInfo`

GetDefinitionInfo returns the DefinitionInfo field if non-nil, zero value otherwise.

### GetDefinitionInfoOk

`func (o *ModelAPIBuild) GetDefinitionInfoOk() (*ModelDefinitionInfo, bool)`

GetDefinitionInfoOk returns a tuple with the DefinitionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinitionInfo

`func (o *ModelAPIBuild) SetDefinitionInfo(v ModelDefinitionInfo)`

SetDefinitionInfo sets DefinitionInfo field to given value.

### HasDefinitionInfo

`func (o *ModelAPIBuild) HasDefinitionInfo() bool`

HasDefinitionInfo returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelAPIBuild) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPIBuild) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPIBuild) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPIBuild) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFinishTime

`func (o *ModelAPIBuild) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ModelAPIBuild) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ModelAPIBuild) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ModelAPIBuild) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetGitHash

`func (o *ModelAPIBuild) GetGitHash() string`

GetGitHash returns the GitHash field if non-nil, zero value otherwise.

### GetGitHashOk

`func (o *ModelAPIBuild) GetGitHashOk() (*string, bool)`

GetGitHashOk returns a tuple with the GitHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitHash

`func (o *ModelAPIBuild) SetGitHash(v string)`

SetGitHash sets GitHash field to given value.

### HasGitHash

`func (o *ModelAPIBuild) HasGitHash() bool`

HasGitHash returns a boolean if a field has been set.

### GetOrder

`func (o *ModelAPIBuild) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelAPIBuild) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelAPIBuild) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelAPIBuild) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrigin

`func (o *ModelAPIBuild) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *ModelAPIBuild) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *ModelAPIBuild) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *ModelAPIBuild) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetPredictedMakespanMs

`func (o *ModelAPIBuild) GetPredictedMakespanMs() int32`

GetPredictedMakespanMs returns the PredictedMakespanMs field if non-nil, zero value otherwise.

### GetPredictedMakespanMsOk

`func (o *ModelAPIBuild) GetPredictedMakespanMsOk() (*int32, bool)`

GetPredictedMakespanMsOk returns a tuple with the PredictedMakespanMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredictedMakespanMs

`func (o *ModelAPIBuild) SetPredictedMakespanMs(v int32)`

SetPredictedMakespanMs sets PredictedMakespanMs field to given value.

### HasPredictedMakespanMs

`func (o *ModelAPIBuild) HasPredictedMakespanMs() bool`

HasPredictedMakespanMs returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelAPIBuild) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelAPIBuild) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelAPIBuild) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelAPIBuild) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectIdentifier

`func (o *ModelAPIBuild) GetProjectIdentifier() string`

GetProjectIdentifier returns the ProjectIdentifier field if non-nil, zero value otherwise.

### GetProjectIdentifierOk

`func (o *ModelAPIBuild) GetProjectIdentifierOk() (*string, bool)`

GetProjectIdentifierOk returns a tuple with the ProjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdentifier

`func (o *ModelAPIBuild) SetProjectIdentifier(v string)`

SetProjectIdentifier sets ProjectIdentifier field to given value.

### HasProjectIdentifier

`func (o *ModelAPIBuild) HasProjectIdentifier() bool`

HasProjectIdentifier returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelAPIBuild) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelAPIBuild) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelAPIBuild) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelAPIBuild) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPIBuild) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPIBuild) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPIBuild) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPIBuild) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCounts

`func (o *ModelAPIBuild) GetStatusCounts() TaskTaskStatusCount`

GetStatusCounts returns the StatusCounts field if non-nil, zero value otherwise.

### GetStatusCountsOk

`func (o *ModelAPIBuild) GetStatusCountsOk() (*TaskTaskStatusCount, bool)`

GetStatusCountsOk returns a tuple with the StatusCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCounts

`func (o *ModelAPIBuild) SetStatusCounts(v TaskTaskStatusCount)`

SetStatusCounts sets StatusCounts field to given value.

### HasStatusCounts

`func (o *ModelAPIBuild) HasStatusCounts() bool`

HasStatusCounts returns a boolean if a field has been set.

### GetTags

`func (o *ModelAPIBuild) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelAPIBuild) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelAPIBuild) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelAPIBuild) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskCache

`func (o *ModelAPIBuild) GetTaskCache() []ModelAPITaskCache`

GetTaskCache returns the TaskCache field if non-nil, zero value otherwise.

### GetTaskCacheOk

`func (o *ModelAPIBuild) GetTaskCacheOk() (*[]ModelAPITaskCache, bool)`

GetTaskCacheOk returns a tuple with the TaskCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskCache

`func (o *ModelAPIBuild) SetTaskCache(v []ModelAPITaskCache)`

SetTaskCache sets TaskCache field to given value.

### HasTaskCache

`func (o *ModelAPIBuild) HasTaskCache() bool`

HasTaskCache returns a boolean if a field has been set.

### GetTasks

`func (o *ModelAPIBuild) GetTasks() []string`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *ModelAPIBuild) GetTasksOk() (*[]string, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *ModelAPIBuild) SetTasks(v []string)`

SetTasks sets Tasks field to given value.

### HasTasks

`func (o *ModelAPIBuild) HasTasks() bool`

HasTasks returns a boolean if a field has been set.

### GetTimeTakenMs

`func (o *ModelAPIBuild) GetTimeTakenMs() int32`

GetTimeTakenMs returns the TimeTakenMs field if non-nil, zero value otherwise.

### GetTimeTakenMsOk

`func (o *ModelAPIBuild) GetTimeTakenMsOk() (*int32, bool)`

GetTimeTakenMsOk returns a tuple with the TimeTakenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenMs

`func (o *ModelAPIBuild) SetTimeTakenMs(v int32)`

SetTimeTakenMs sets TimeTakenMs field to given value.

### HasTimeTakenMs

`func (o *ModelAPIBuild) HasTimeTakenMs() bool`

HasTimeTakenMs returns a boolean if a field has been set.

### GetVersion

`func (o *ModelAPIBuild) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelAPIBuild) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelAPIBuild) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelAPIBuild) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


