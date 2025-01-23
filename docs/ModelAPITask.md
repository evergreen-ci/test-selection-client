# ModelAPITask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbortInfo** | Pointer to [**ModelAPIAbortInfo**](ModelAPIAbortInfo.md) |  | [optional] 
**Aborted** | Pointer to **bool** |  | [optional] 
**Activated** | Pointer to **bool** | Whether the task is currently active | [optional] 
**ActivatedBy** | Pointer to **string** | Identifier of the process or user that activated this task | [optional] 
**ActivatedTime** | Pointer to **string** |  | [optional] 
**Ami** | Pointer to **string** |  | [optional] 
**Archived** | Pointer to **bool** |  | [optional] 
**Artifacts** | Pointer to [**[]ModelAPIFile**](ModelAPIFile.md) | The list of artifacts associated with the task. | [optional] 
**BaseTask** | Pointer to [**ModelAPIBaseTaskInfo**](ModelAPIBaseTaskInfo.md) |  | [optional] 
**Blocked** | Pointer to **bool** |  | [optional] 
**BuildId** | Pointer to **string** | Identifier of the build that this task is part of | [optional] 
**BuildVariant** | Pointer to **string** | Name of the buildvariant that this task runs on | [optional] 
**BuildVariantDisplayName** | Pointer to **string** |  | [optional] 
**Container** | Pointer to **string** |  | [optional] 
**ContainerAllocated** | Pointer to **bool** |  | [optional] 
**ContainerAllocatedTime** | Pointer to **string** |  | [optional] 
**ContainerAllocationAttempts** | Pointer to **int32** |  | [optional] 
**ContainerOptions** | Pointer to [**ModelAPIContainerOptions**](ModelAPIContainerOptions.md) |  | [optional] 
**CreateTime** | Pointer to **string** | Time that this task was first created | [optional] 
**DependsOn** | Pointer to [**[]ModelAPIDependency**](ModelAPIDependency.md) | List of task_ids of task that this task depends on before beginning | [optional] 
**DispatchTime** | Pointer to **string** | Time that this time was dispatched | [optional] 
**DisplayName** | Pointer to **string** | Name of this task displayed in the UI | [optional] 
**DisplayOnly** | Pointer to **bool** |  | [optional] 
**DisplayStatus** | Pointer to **string** | The status of this task that is displayed in the UI (possible values are \&quot;will-run\&quot;, \&quot;unscheduled\&quot;, \&quot;blocked\&quot;, \&quot;dispatched\&quot;, \&quot;started\&quot;, \&quot;success\&quot;, \&quot;failed\&quot;, \&quot;aborted\&quot;, \&quot;system-failed\&quot;, \&quot;system-unresponsive\&quot;, \&quot;system-timed-out\&quot;, \&quot;task-timed-out\&quot;, \&quot;known-issue\&quot;) | [optional] 
**DistroId** | Pointer to **string** | Identifier of the distro that this task runs on | [optional] 
**EstWaitToStartMs** | Pointer to **int32** |  | [optional] 
**Execution** | Pointer to **int32** | The number of the execution of this particular task | [optional] 
**ExecutionTasks** | Pointer to **[]string** |  | [optional] 
**ExpectedDurationMs** | Pointer to **int32** | Number of milliseconds expected for this task to execute | [optional] 
**FinishTime** | Pointer to **string** | Time that this task finished execution | [optional] 
**GenerateTask** | Pointer to **bool** |  | [optional] 
**GeneratedBy** | Pointer to **string** |  | [optional] 
**HostId** | Pointer to **string** | The ID of the host this task ran or is running on | [optional] 
**IngestTime** | Pointer to **string** |  | [optional] 
**Logs** | Pointer to [**ModelLogLinks**](ModelLogLinks.md) | Object containing raw and event logs for this task | [optional] 
**Mainline** | Pointer to **bool** |  | [optional] 
**MustHaveTestResults** | Pointer to **bool** |  | [optional] 
**Order** | Pointer to **int32** | For mainline commits, represents the position in the commit history of commit this task is associated with. For patches, this represents the number of total patches submitted by the user. | [optional] 
**ParentTaskId** | Pointer to **string** | The ID of the task&#39;s parent display task, if requested and available | [optional] 
**ParsleyLogs** | Pointer to [**ModelLogLinks**](ModelLogLinks.md) | Object containing parsley logs for this task | [optional] 
**PodId** | Pointer to **string** |  | [optional] 
**PreviousExecutions** | Pointer to [**[]ModelAPITask**](ModelAPITask.md) | Contains previous executions of the task if they were requested, and available. May be empty | [optional] 
**Priority** | Pointer to **int32** | The priority of this task to be run | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ProjectIdentifier** | Pointer to **string** |  | [optional] 
**Requester** | Pointer to **string** | Version created by one of patch_request\&quot;, \&quot;github_pull_request\&quot;, \&quot;gitter_request\&quot; (caused by git commit, aka the repotracker requester), \&quot;trigger_request\&quot; (Project Trigger versions) , \&quot;merge_test\&quot; (commit queue patches), \&quot;ad_hoc\&quot; (periodic builds) | [optional] 
**ResetWhenFinished** | Pointer to **bool** |  | [optional] 
**Revision** | Pointer to **string** | The version control identifier associated with this task | [optional] 
**ScheduledTime** | Pointer to **string** | Time that this task is scheduled to begin | [optional] 
**StartTime** | Pointer to **string** | Time that this task began execution | [optional] 
**Status** | Pointer to **string** | The current status of this task (possible values are \&quot;undispatched\&quot;, \&quot;dispatched\&quot;, \&quot;started\&quot;, \&quot;success\&quot;, and \&quot;failed\&quot;) | [optional] 
**StatusDetails** | Pointer to [**ModelApiTaskEndDetail**](ModelApiTaskEndDetail.md) | Object containing additional information about the status | [optional] 
**StepbackInfo** | Pointer to [**ModelAPIStepbackInfo**](ModelAPIStepbackInfo.md) | The information, if any, about stepback | [optional] 
**Tags** | Pointer to **[]string** | List of tags defined for the task, if any | [optional] 
**TaskGroup** | Pointer to **string** |  | [optional] 
**TaskGroupMaxHosts** | Pointer to **int32** |  | [optional] 
**TaskId** | Pointer to **string** | Unique identifier of this task | [optional] 
**TestResults** | Pointer to [**[]ModelAPITest**](ModelAPITest.md) |  | [optional] 
**TimeTakenMs** | Pointer to **int32** | Number of milliseconds this task took during execution | [optional] 
**VersionId** | Pointer to **string** | An identifier of this task by its project and commit hash | [optional] 

## Methods

### NewModelAPITask

`func NewModelAPITask() *ModelAPITask`

NewModelAPITask instantiates a new ModelAPITask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITaskWithDefaults

`func NewModelAPITaskWithDefaults() *ModelAPITask`

NewModelAPITaskWithDefaults instantiates a new ModelAPITask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbortInfo

`func (o *ModelAPITask) GetAbortInfo() ModelAPIAbortInfo`

GetAbortInfo returns the AbortInfo field if non-nil, zero value otherwise.

### GetAbortInfoOk

`func (o *ModelAPITask) GetAbortInfoOk() (*ModelAPIAbortInfo, bool)`

GetAbortInfoOk returns a tuple with the AbortInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbortInfo

`func (o *ModelAPITask) SetAbortInfo(v ModelAPIAbortInfo)`

SetAbortInfo sets AbortInfo field to given value.

### HasAbortInfo

`func (o *ModelAPITask) HasAbortInfo() bool`

HasAbortInfo returns a boolean if a field has been set.

### GetAborted

`func (o *ModelAPITask) GetAborted() bool`

GetAborted returns the Aborted field if non-nil, zero value otherwise.

### GetAbortedOk

`func (o *ModelAPITask) GetAbortedOk() (*bool, bool)`

GetAbortedOk returns a tuple with the Aborted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAborted

`func (o *ModelAPITask) SetAborted(v bool)`

SetAborted sets Aborted field to given value.

### HasAborted

`func (o *ModelAPITask) HasAborted() bool`

HasAborted returns a boolean if a field has been set.

### GetActivated

`func (o *ModelAPITask) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ModelAPITask) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ModelAPITask) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ModelAPITask) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivatedBy

`func (o *ModelAPITask) GetActivatedBy() string`

GetActivatedBy returns the ActivatedBy field if non-nil, zero value otherwise.

### GetActivatedByOk

`func (o *ModelAPITask) GetActivatedByOk() (*string, bool)`

GetActivatedByOk returns a tuple with the ActivatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedBy

`func (o *ModelAPITask) SetActivatedBy(v string)`

SetActivatedBy sets ActivatedBy field to given value.

### HasActivatedBy

`func (o *ModelAPITask) HasActivatedBy() bool`

HasActivatedBy returns a boolean if a field has been set.

### GetActivatedTime

`func (o *ModelAPITask) GetActivatedTime() string`

GetActivatedTime returns the ActivatedTime field if non-nil, zero value otherwise.

### GetActivatedTimeOk

`func (o *ModelAPITask) GetActivatedTimeOk() (*string, bool)`

GetActivatedTimeOk returns a tuple with the ActivatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedTime

`func (o *ModelAPITask) SetActivatedTime(v string)`

SetActivatedTime sets ActivatedTime field to given value.

### HasActivatedTime

`func (o *ModelAPITask) HasActivatedTime() bool`

HasActivatedTime returns a boolean if a field has been set.

### GetAmi

`func (o *ModelAPITask) GetAmi() string`

GetAmi returns the Ami field if non-nil, zero value otherwise.

### GetAmiOk

`func (o *ModelAPITask) GetAmiOk() (*string, bool)`

GetAmiOk returns a tuple with the Ami field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmi

`func (o *ModelAPITask) SetAmi(v string)`

SetAmi sets Ami field to given value.

### HasAmi

`func (o *ModelAPITask) HasAmi() bool`

HasAmi returns a boolean if a field has been set.

### GetArchived

`func (o *ModelAPITask) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ModelAPITask) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ModelAPITask) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *ModelAPITask) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetArtifacts

`func (o *ModelAPITask) GetArtifacts() []ModelAPIFile`

GetArtifacts returns the Artifacts field if non-nil, zero value otherwise.

### GetArtifactsOk

`func (o *ModelAPITask) GetArtifactsOk() (*[]ModelAPIFile, bool)`

GetArtifactsOk returns a tuple with the Artifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifacts

`func (o *ModelAPITask) SetArtifacts(v []ModelAPIFile)`

SetArtifacts sets Artifacts field to given value.

### HasArtifacts

`func (o *ModelAPITask) HasArtifacts() bool`

HasArtifacts returns a boolean if a field has been set.

### GetBaseTask

`func (o *ModelAPITask) GetBaseTask() ModelAPIBaseTaskInfo`

GetBaseTask returns the BaseTask field if non-nil, zero value otherwise.

### GetBaseTaskOk

`func (o *ModelAPITask) GetBaseTaskOk() (*ModelAPIBaseTaskInfo, bool)`

GetBaseTaskOk returns a tuple with the BaseTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseTask

`func (o *ModelAPITask) SetBaseTask(v ModelAPIBaseTaskInfo)`

SetBaseTask sets BaseTask field to given value.

### HasBaseTask

`func (o *ModelAPITask) HasBaseTask() bool`

HasBaseTask returns a boolean if a field has been set.

### GetBlocked

`func (o *ModelAPITask) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ModelAPITask) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ModelAPITask) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *ModelAPITask) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBuildId

`func (o *ModelAPITask) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *ModelAPITask) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *ModelAPITask) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *ModelAPITask) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetBuildVariant

`func (o *ModelAPITask) GetBuildVariant() string`

GetBuildVariant returns the BuildVariant field if non-nil, zero value otherwise.

### GetBuildVariantOk

`func (o *ModelAPITask) GetBuildVariantOk() (*string, bool)`

GetBuildVariantOk returns a tuple with the BuildVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVariant

`func (o *ModelAPITask) SetBuildVariant(v string)`

SetBuildVariant sets BuildVariant field to given value.

### HasBuildVariant

`func (o *ModelAPITask) HasBuildVariant() bool`

HasBuildVariant returns a boolean if a field has been set.

### GetBuildVariantDisplayName

`func (o *ModelAPITask) GetBuildVariantDisplayName() string`

GetBuildVariantDisplayName returns the BuildVariantDisplayName field if non-nil, zero value otherwise.

### GetBuildVariantDisplayNameOk

`func (o *ModelAPITask) GetBuildVariantDisplayNameOk() (*string, bool)`

GetBuildVariantDisplayNameOk returns a tuple with the BuildVariantDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVariantDisplayName

`func (o *ModelAPITask) SetBuildVariantDisplayName(v string)`

SetBuildVariantDisplayName sets BuildVariantDisplayName field to given value.

### HasBuildVariantDisplayName

`func (o *ModelAPITask) HasBuildVariantDisplayName() bool`

HasBuildVariantDisplayName returns a boolean if a field has been set.

### GetContainer

`func (o *ModelAPITask) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelAPITask) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelAPITask) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelAPITask) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetContainerAllocated

`func (o *ModelAPITask) GetContainerAllocated() bool`

GetContainerAllocated returns the ContainerAllocated field if non-nil, zero value otherwise.

### GetContainerAllocatedOk

`func (o *ModelAPITask) GetContainerAllocatedOk() (*bool, bool)`

GetContainerAllocatedOk returns a tuple with the ContainerAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAllocated

`func (o *ModelAPITask) SetContainerAllocated(v bool)`

SetContainerAllocated sets ContainerAllocated field to given value.

### HasContainerAllocated

`func (o *ModelAPITask) HasContainerAllocated() bool`

HasContainerAllocated returns a boolean if a field has been set.

### GetContainerAllocatedTime

`func (o *ModelAPITask) GetContainerAllocatedTime() string`

GetContainerAllocatedTime returns the ContainerAllocatedTime field if non-nil, zero value otherwise.

### GetContainerAllocatedTimeOk

`func (o *ModelAPITask) GetContainerAllocatedTimeOk() (*string, bool)`

GetContainerAllocatedTimeOk returns a tuple with the ContainerAllocatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAllocatedTime

`func (o *ModelAPITask) SetContainerAllocatedTime(v string)`

SetContainerAllocatedTime sets ContainerAllocatedTime field to given value.

### HasContainerAllocatedTime

`func (o *ModelAPITask) HasContainerAllocatedTime() bool`

HasContainerAllocatedTime returns a boolean if a field has been set.

### GetContainerAllocationAttempts

`func (o *ModelAPITask) GetContainerAllocationAttempts() int32`

GetContainerAllocationAttempts returns the ContainerAllocationAttempts field if non-nil, zero value otherwise.

### GetContainerAllocationAttemptsOk

`func (o *ModelAPITask) GetContainerAllocationAttemptsOk() (*int32, bool)`

GetContainerAllocationAttemptsOk returns a tuple with the ContainerAllocationAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerAllocationAttempts

`func (o *ModelAPITask) SetContainerAllocationAttempts(v int32)`

SetContainerAllocationAttempts sets ContainerAllocationAttempts field to given value.

### HasContainerAllocationAttempts

`func (o *ModelAPITask) HasContainerAllocationAttempts() bool`

HasContainerAllocationAttempts returns a boolean if a field has been set.

### GetContainerOptions

`func (o *ModelAPITask) GetContainerOptions() ModelAPIContainerOptions`

GetContainerOptions returns the ContainerOptions field if non-nil, zero value otherwise.

### GetContainerOptionsOk

`func (o *ModelAPITask) GetContainerOptionsOk() (*ModelAPIContainerOptions, bool)`

GetContainerOptionsOk returns a tuple with the ContainerOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerOptions

`func (o *ModelAPITask) SetContainerOptions(v ModelAPIContainerOptions)`

SetContainerOptions sets ContainerOptions field to given value.

### HasContainerOptions

`func (o *ModelAPITask) HasContainerOptions() bool`

HasContainerOptions returns a boolean if a field has been set.

### GetCreateTime

`func (o *ModelAPITask) GetCreateTime() string`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ModelAPITask) GetCreateTimeOk() (*string, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ModelAPITask) SetCreateTime(v string)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ModelAPITask) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDependsOn

`func (o *ModelAPITask) GetDependsOn() []ModelAPIDependency`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *ModelAPITask) GetDependsOnOk() (*[]ModelAPIDependency, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *ModelAPITask) SetDependsOn(v []ModelAPIDependency)`

SetDependsOn sets DependsOn field to given value.

### HasDependsOn

`func (o *ModelAPITask) HasDependsOn() bool`

HasDependsOn returns a boolean if a field has been set.

### GetDispatchTime

`func (o *ModelAPITask) GetDispatchTime() string`

GetDispatchTime returns the DispatchTime field if non-nil, zero value otherwise.

### GetDispatchTimeOk

`func (o *ModelAPITask) GetDispatchTimeOk() (*string, bool)`

GetDispatchTimeOk returns a tuple with the DispatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchTime

`func (o *ModelAPITask) SetDispatchTime(v string)`

SetDispatchTime sets DispatchTime field to given value.

### HasDispatchTime

`func (o *ModelAPITask) HasDispatchTime() bool`

HasDispatchTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelAPITask) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPITask) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPITask) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPITask) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayOnly

`func (o *ModelAPITask) GetDisplayOnly() bool`

GetDisplayOnly returns the DisplayOnly field if non-nil, zero value otherwise.

### GetDisplayOnlyOk

`func (o *ModelAPITask) GetDisplayOnlyOk() (*bool, bool)`

GetDisplayOnlyOk returns a tuple with the DisplayOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOnly

`func (o *ModelAPITask) SetDisplayOnly(v bool)`

SetDisplayOnly sets DisplayOnly field to given value.

### HasDisplayOnly

`func (o *ModelAPITask) HasDisplayOnly() bool`

HasDisplayOnly returns a boolean if a field has been set.

### GetDisplayStatus

`func (o *ModelAPITask) GetDisplayStatus() string`

GetDisplayStatus returns the DisplayStatus field if non-nil, zero value otherwise.

### GetDisplayStatusOk

`func (o *ModelAPITask) GetDisplayStatusOk() (*string, bool)`

GetDisplayStatusOk returns a tuple with the DisplayStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayStatus

`func (o *ModelAPITask) SetDisplayStatus(v string)`

SetDisplayStatus sets DisplayStatus field to given value.

### HasDisplayStatus

`func (o *ModelAPITask) HasDisplayStatus() bool`

HasDisplayStatus returns a boolean if a field has been set.

### GetDistroId

`func (o *ModelAPITask) GetDistroId() string`

GetDistroId returns the DistroId field if non-nil, zero value otherwise.

### GetDistroIdOk

`func (o *ModelAPITask) GetDistroIdOk() (*string, bool)`

GetDistroIdOk returns a tuple with the DistroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroId

`func (o *ModelAPITask) SetDistroId(v string)`

SetDistroId sets DistroId field to given value.

### HasDistroId

`func (o *ModelAPITask) HasDistroId() bool`

HasDistroId returns a boolean if a field has been set.

### GetEstWaitToStartMs

`func (o *ModelAPITask) GetEstWaitToStartMs() int32`

GetEstWaitToStartMs returns the EstWaitToStartMs field if non-nil, zero value otherwise.

### GetEstWaitToStartMsOk

`func (o *ModelAPITask) GetEstWaitToStartMsOk() (*int32, bool)`

GetEstWaitToStartMsOk returns a tuple with the EstWaitToStartMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstWaitToStartMs

`func (o *ModelAPITask) SetEstWaitToStartMs(v int32)`

SetEstWaitToStartMs sets EstWaitToStartMs field to given value.

### HasEstWaitToStartMs

`func (o *ModelAPITask) HasEstWaitToStartMs() bool`

HasEstWaitToStartMs returns a boolean if a field has been set.

### GetExecution

`func (o *ModelAPITask) GetExecution() int32`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ModelAPITask) GetExecutionOk() (*int32, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ModelAPITask) SetExecution(v int32)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *ModelAPITask) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetExecutionTasks

`func (o *ModelAPITask) GetExecutionTasks() []string`

GetExecutionTasks returns the ExecutionTasks field if non-nil, zero value otherwise.

### GetExecutionTasksOk

`func (o *ModelAPITask) GetExecutionTasksOk() (*[]string, bool)`

GetExecutionTasksOk returns a tuple with the ExecutionTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionTasks

`func (o *ModelAPITask) SetExecutionTasks(v []string)`

SetExecutionTasks sets ExecutionTasks field to given value.

### HasExecutionTasks

`func (o *ModelAPITask) HasExecutionTasks() bool`

HasExecutionTasks returns a boolean if a field has been set.

### GetExpectedDurationMs

`func (o *ModelAPITask) GetExpectedDurationMs() int32`

GetExpectedDurationMs returns the ExpectedDurationMs field if non-nil, zero value otherwise.

### GetExpectedDurationMsOk

`func (o *ModelAPITask) GetExpectedDurationMsOk() (*int32, bool)`

GetExpectedDurationMsOk returns a tuple with the ExpectedDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDurationMs

`func (o *ModelAPITask) SetExpectedDurationMs(v int32)`

SetExpectedDurationMs sets ExpectedDurationMs field to given value.

### HasExpectedDurationMs

`func (o *ModelAPITask) HasExpectedDurationMs() bool`

HasExpectedDurationMs returns a boolean if a field has been set.

### GetFinishTime

`func (o *ModelAPITask) GetFinishTime() string`

GetFinishTime returns the FinishTime field if non-nil, zero value otherwise.

### GetFinishTimeOk

`func (o *ModelAPITask) GetFinishTimeOk() (*string, bool)`

GetFinishTimeOk returns a tuple with the FinishTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishTime

`func (o *ModelAPITask) SetFinishTime(v string)`

SetFinishTime sets FinishTime field to given value.

### HasFinishTime

`func (o *ModelAPITask) HasFinishTime() bool`

HasFinishTime returns a boolean if a field has been set.

### GetGenerateTask

`func (o *ModelAPITask) GetGenerateTask() bool`

GetGenerateTask returns the GenerateTask field if non-nil, zero value otherwise.

### GetGenerateTaskOk

`func (o *ModelAPITask) GetGenerateTaskOk() (*bool, bool)`

GetGenerateTaskOk returns a tuple with the GenerateTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateTask

`func (o *ModelAPITask) SetGenerateTask(v bool)`

SetGenerateTask sets GenerateTask field to given value.

### HasGenerateTask

`func (o *ModelAPITask) HasGenerateTask() bool`

HasGenerateTask returns a boolean if a field has been set.

### GetGeneratedBy

`func (o *ModelAPITask) GetGeneratedBy() string`

GetGeneratedBy returns the GeneratedBy field if non-nil, zero value otherwise.

### GetGeneratedByOk

`func (o *ModelAPITask) GetGeneratedByOk() (*string, bool)`

GetGeneratedByOk returns a tuple with the GeneratedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedBy

`func (o *ModelAPITask) SetGeneratedBy(v string)`

SetGeneratedBy sets GeneratedBy field to given value.

### HasGeneratedBy

`func (o *ModelAPITask) HasGeneratedBy() bool`

HasGeneratedBy returns a boolean if a field has been set.

### GetHostId

`func (o *ModelAPITask) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ModelAPITask) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ModelAPITask) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ModelAPITask) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetIngestTime

`func (o *ModelAPITask) GetIngestTime() string`

GetIngestTime returns the IngestTime field if non-nil, zero value otherwise.

### GetIngestTimeOk

`func (o *ModelAPITask) GetIngestTimeOk() (*string, bool)`

GetIngestTimeOk returns a tuple with the IngestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestTime

`func (o *ModelAPITask) SetIngestTime(v string)`

SetIngestTime sets IngestTime field to given value.

### HasIngestTime

`func (o *ModelAPITask) HasIngestTime() bool`

HasIngestTime returns a boolean if a field has been set.

### GetLogs

`func (o *ModelAPITask) GetLogs() ModelLogLinks`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ModelAPITask) GetLogsOk() (*ModelLogLinks, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ModelAPITask) SetLogs(v ModelLogLinks)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ModelAPITask) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetMainline

`func (o *ModelAPITask) GetMainline() bool`

GetMainline returns the Mainline field if non-nil, zero value otherwise.

### GetMainlineOk

`func (o *ModelAPITask) GetMainlineOk() (*bool, bool)`

GetMainlineOk returns a tuple with the Mainline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainline

`func (o *ModelAPITask) SetMainline(v bool)`

SetMainline sets Mainline field to given value.

### HasMainline

`func (o *ModelAPITask) HasMainline() bool`

HasMainline returns a boolean if a field has been set.

### GetMustHaveTestResults

`func (o *ModelAPITask) GetMustHaveTestResults() bool`

GetMustHaveTestResults returns the MustHaveTestResults field if non-nil, zero value otherwise.

### GetMustHaveTestResultsOk

`func (o *ModelAPITask) GetMustHaveTestResultsOk() (*bool, bool)`

GetMustHaveTestResultsOk returns a tuple with the MustHaveTestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustHaveTestResults

`func (o *ModelAPITask) SetMustHaveTestResults(v bool)`

SetMustHaveTestResults sets MustHaveTestResults field to given value.

### HasMustHaveTestResults

`func (o *ModelAPITask) HasMustHaveTestResults() bool`

HasMustHaveTestResults returns a boolean if a field has been set.

### GetOrder

`func (o *ModelAPITask) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ModelAPITask) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ModelAPITask) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ModelAPITask) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetParentTaskId

`func (o *ModelAPITask) GetParentTaskId() string`

GetParentTaskId returns the ParentTaskId field if non-nil, zero value otherwise.

### GetParentTaskIdOk

`func (o *ModelAPITask) GetParentTaskIdOk() (*string, bool)`

GetParentTaskIdOk returns a tuple with the ParentTaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTaskId

`func (o *ModelAPITask) SetParentTaskId(v string)`

SetParentTaskId sets ParentTaskId field to given value.

### HasParentTaskId

`func (o *ModelAPITask) HasParentTaskId() bool`

HasParentTaskId returns a boolean if a field has been set.

### GetParsleyLogs

`func (o *ModelAPITask) GetParsleyLogs() ModelLogLinks`

GetParsleyLogs returns the ParsleyLogs field if non-nil, zero value otherwise.

### GetParsleyLogsOk

`func (o *ModelAPITask) GetParsleyLogsOk() (*ModelLogLinks, bool)`

GetParsleyLogsOk returns a tuple with the ParsleyLogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsleyLogs

`func (o *ModelAPITask) SetParsleyLogs(v ModelLogLinks)`

SetParsleyLogs sets ParsleyLogs field to given value.

### HasParsleyLogs

`func (o *ModelAPITask) HasParsleyLogs() bool`

HasParsleyLogs returns a boolean if a field has been set.

### GetPodId

`func (o *ModelAPITask) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *ModelAPITask) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *ModelAPITask) SetPodId(v string)`

SetPodId sets PodId field to given value.

### HasPodId

`func (o *ModelAPITask) HasPodId() bool`

HasPodId returns a boolean if a field has been set.

### GetPreviousExecutions

`func (o *ModelAPITask) GetPreviousExecutions() []ModelAPITask`

GetPreviousExecutions returns the PreviousExecutions field if non-nil, zero value otherwise.

### GetPreviousExecutionsOk

`func (o *ModelAPITask) GetPreviousExecutionsOk() (*[]ModelAPITask, bool)`

GetPreviousExecutionsOk returns a tuple with the PreviousExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousExecutions

`func (o *ModelAPITask) SetPreviousExecutions(v []ModelAPITask)`

SetPreviousExecutions sets PreviousExecutions field to given value.

### HasPreviousExecutions

`func (o *ModelAPITask) HasPreviousExecutions() bool`

HasPreviousExecutions returns a boolean if a field has been set.

### GetPriority

`func (o *ModelAPITask) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ModelAPITask) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ModelAPITask) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ModelAPITask) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetProjectId

`func (o *ModelAPITask) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ModelAPITask) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ModelAPITask) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ModelAPITask) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetProjectIdentifier

`func (o *ModelAPITask) GetProjectIdentifier() string`

GetProjectIdentifier returns the ProjectIdentifier field if non-nil, zero value otherwise.

### GetProjectIdentifierOk

`func (o *ModelAPITask) GetProjectIdentifierOk() (*string, bool)`

GetProjectIdentifierOk returns a tuple with the ProjectIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIdentifier

`func (o *ModelAPITask) SetProjectIdentifier(v string)`

SetProjectIdentifier sets ProjectIdentifier field to given value.

### HasProjectIdentifier

`func (o *ModelAPITask) HasProjectIdentifier() bool`

HasProjectIdentifier returns a boolean if a field has been set.

### GetRequester

`func (o *ModelAPITask) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ModelAPITask) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ModelAPITask) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ModelAPITask) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetResetWhenFinished

`func (o *ModelAPITask) GetResetWhenFinished() bool`

GetResetWhenFinished returns the ResetWhenFinished field if non-nil, zero value otherwise.

### GetResetWhenFinishedOk

`func (o *ModelAPITask) GetResetWhenFinishedOk() (*bool, bool)`

GetResetWhenFinishedOk returns a tuple with the ResetWhenFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetWhenFinished

`func (o *ModelAPITask) SetResetWhenFinished(v bool)`

SetResetWhenFinished sets ResetWhenFinished field to given value.

### HasResetWhenFinished

`func (o *ModelAPITask) HasResetWhenFinished() bool`

HasResetWhenFinished returns a boolean if a field has been set.

### GetRevision

`func (o *ModelAPITask) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ModelAPITask) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ModelAPITask) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ModelAPITask) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetScheduledTime

`func (o *ModelAPITask) GetScheduledTime() string`

GetScheduledTime returns the ScheduledTime field if non-nil, zero value otherwise.

### GetScheduledTimeOk

`func (o *ModelAPITask) GetScheduledTimeOk() (*string, bool)`

GetScheduledTimeOk returns a tuple with the ScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledTime

`func (o *ModelAPITask) SetScheduledTime(v string)`

SetScheduledTime sets ScheduledTime field to given value.

### HasScheduledTime

`func (o *ModelAPITask) HasScheduledTime() bool`

HasScheduledTime returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelAPITask) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelAPITask) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelAPITask) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelAPITask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPITask) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPITask) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPITask) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPITask) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDetails

`func (o *ModelAPITask) GetStatusDetails() ModelApiTaskEndDetail`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *ModelAPITask) GetStatusDetailsOk() (*ModelApiTaskEndDetail, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *ModelAPITask) SetStatusDetails(v ModelApiTaskEndDetail)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *ModelAPITask) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

### GetStepbackInfo

`func (o *ModelAPITask) GetStepbackInfo() ModelAPIStepbackInfo`

GetStepbackInfo returns the StepbackInfo field if non-nil, zero value otherwise.

### GetStepbackInfoOk

`func (o *ModelAPITask) GetStepbackInfoOk() (*ModelAPIStepbackInfo, bool)`

GetStepbackInfoOk returns a tuple with the StepbackInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepbackInfo

`func (o *ModelAPITask) SetStepbackInfo(v ModelAPIStepbackInfo)`

SetStepbackInfo sets StepbackInfo field to given value.

### HasStepbackInfo

`func (o *ModelAPITask) HasStepbackInfo() bool`

HasStepbackInfo returns a boolean if a field has been set.

### GetTags

`func (o *ModelAPITask) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ModelAPITask) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ModelAPITask) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ModelAPITask) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTaskGroup

`func (o *ModelAPITask) GetTaskGroup() string`

GetTaskGroup returns the TaskGroup field if non-nil, zero value otherwise.

### GetTaskGroupOk

`func (o *ModelAPITask) GetTaskGroupOk() (*string, bool)`

GetTaskGroupOk returns a tuple with the TaskGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroup

`func (o *ModelAPITask) SetTaskGroup(v string)`

SetTaskGroup sets TaskGroup field to given value.

### HasTaskGroup

`func (o *ModelAPITask) HasTaskGroup() bool`

HasTaskGroup returns a boolean if a field has been set.

### GetTaskGroupMaxHosts

`func (o *ModelAPITask) GetTaskGroupMaxHosts() int32`

GetTaskGroupMaxHosts returns the TaskGroupMaxHosts field if non-nil, zero value otherwise.

### GetTaskGroupMaxHostsOk

`func (o *ModelAPITask) GetTaskGroupMaxHostsOk() (*int32, bool)`

GetTaskGroupMaxHostsOk returns a tuple with the TaskGroupMaxHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskGroupMaxHosts

`func (o *ModelAPITask) SetTaskGroupMaxHosts(v int32)`

SetTaskGroupMaxHosts sets TaskGroupMaxHosts field to given value.

### HasTaskGroupMaxHosts

`func (o *ModelAPITask) HasTaskGroupMaxHosts() bool`

HasTaskGroupMaxHosts returns a boolean if a field has been set.

### GetTaskId

`func (o *ModelAPITask) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelAPITask) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelAPITask) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelAPITask) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTestResults

`func (o *ModelAPITask) GetTestResults() []ModelAPITest`

GetTestResults returns the TestResults field if non-nil, zero value otherwise.

### GetTestResultsOk

`func (o *ModelAPITask) GetTestResultsOk() (*[]ModelAPITest, bool)`

GetTestResultsOk returns a tuple with the TestResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResults

`func (o *ModelAPITask) SetTestResults(v []ModelAPITest)`

SetTestResults sets TestResults field to given value.

### HasTestResults

`func (o *ModelAPITask) HasTestResults() bool`

HasTestResults returns a boolean if a field has been set.

### GetTimeTakenMs

`func (o *ModelAPITask) GetTimeTakenMs() int32`

GetTimeTakenMs returns the TimeTakenMs field if non-nil, zero value otherwise.

### GetTimeTakenMsOk

`func (o *ModelAPITask) GetTimeTakenMsOk() (*int32, bool)`

GetTimeTakenMsOk returns a tuple with the TimeTakenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenMs

`func (o *ModelAPITask) SetTimeTakenMs(v int32)`

SetTimeTakenMs sets TimeTakenMs field to given value.

### HasTimeTakenMs

`func (o *ModelAPITask) HasTimeTakenMs() bool`

HasTimeTakenMs returns a boolean if a field has been set.

### GetVersionId

`func (o *ModelAPITask) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ModelAPITask) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ModelAPITask) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ModelAPITask) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


