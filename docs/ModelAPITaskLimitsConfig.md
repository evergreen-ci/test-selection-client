# ModelAPITaskLimitsConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxConcurrentLargeParserProjectTasks** | Pointer to **int32** | MaxConcurrentLargeParserProjectTasks is the maximum number of tasks with parser projects stored in S3 that can be running at once. | [optional] 
**MaxDailyAutomaticRestarts** | Pointer to **int32** | MaxDailyAutomaticRestarts is the maximum number of times a project can automatically restart a task within a 24-hour period. | [optional] 
**MaxDegradedModeConcurrentLargeParserProjectTasks** | Pointer to **int32** | MaxDegradedModeConcurrentLargeParserProjectTasks is the maximum number of tasks with parser projects stored in S3 that can be running at once during CPU degraded mode. | [optional] 
**MaxDegradedModeParserProjectSize** | Pointer to **int32** | MaxDegradedModeParserProjectSize is the maximum parser project size in MB during CPU degraded mode. | [optional] 
**MaxExecTimeoutSecs** | Pointer to **int32** | MaxExecTimeoutSecs is the maximum number of seconds a task can run and set their timeout to. | [optional] 
**MaxGenerateTaskJsonSize** | Pointer to **int32** | MaxGenerateTaskJSONSize is the maximum size of a JSON file in MB that can be specified in the GenerateTasks command. | [optional] 
**MaxHourlyPatchTasks** | Pointer to **int32** | MaxHourlyPatchTasks is the maximum number of patch tasks a single user can schedule per hour. | [optional] 
**MaxIncludesPerVersion** | Pointer to **int32** | MaxIncludesPerVersion is the maximum number of includes that a single version can have. | [optional] 
**MaxParserProjectSize** | Pointer to **int32** | MaxParserProjectSize is the maximum allowed size in MB for parser projects that are stored in S3. | [optional] 
**MaxPendingGeneratedTasks** | Pointer to **int32** | MaxPendingGeneratedTasks is the maximum number of tasks that can be created by all generated task at once. | [optional] 
**MaxTaskExecution** | Pointer to **int32** | MaxTaskExecution is the maximum task (zero based) execution number. | [optional] 
**MaxTasksPerVersion** | Pointer to **int32** | MaxTasksPerVersion is the maximum number of tasks that a single version can have. | [optional] 

## Methods

### NewModelAPITaskLimitsConfig

`func NewModelAPITaskLimitsConfig() *ModelAPITaskLimitsConfig`

NewModelAPITaskLimitsConfig instantiates a new ModelAPITaskLimitsConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITaskLimitsConfigWithDefaults

`func NewModelAPITaskLimitsConfigWithDefaults() *ModelAPITaskLimitsConfig`

NewModelAPITaskLimitsConfigWithDefaults instantiates a new ModelAPITaskLimitsConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxConcurrentLargeParserProjectTasks

`func (o *ModelAPITaskLimitsConfig) GetMaxConcurrentLargeParserProjectTasks() int32`

GetMaxConcurrentLargeParserProjectTasks returns the MaxConcurrentLargeParserProjectTasks field if non-nil, zero value otherwise.

### GetMaxConcurrentLargeParserProjectTasksOk

`func (o *ModelAPITaskLimitsConfig) GetMaxConcurrentLargeParserProjectTasksOk() (*int32, bool)`

GetMaxConcurrentLargeParserProjectTasksOk returns a tuple with the MaxConcurrentLargeParserProjectTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentLargeParserProjectTasks

`func (o *ModelAPITaskLimitsConfig) SetMaxConcurrentLargeParserProjectTasks(v int32)`

SetMaxConcurrentLargeParserProjectTasks sets MaxConcurrentLargeParserProjectTasks field to given value.

### HasMaxConcurrentLargeParserProjectTasks

`func (o *ModelAPITaskLimitsConfig) HasMaxConcurrentLargeParserProjectTasks() bool`

HasMaxConcurrentLargeParserProjectTasks returns a boolean if a field has been set.

### GetMaxDailyAutomaticRestarts

`func (o *ModelAPITaskLimitsConfig) GetMaxDailyAutomaticRestarts() int32`

GetMaxDailyAutomaticRestarts returns the MaxDailyAutomaticRestarts field if non-nil, zero value otherwise.

### GetMaxDailyAutomaticRestartsOk

`func (o *ModelAPITaskLimitsConfig) GetMaxDailyAutomaticRestartsOk() (*int32, bool)`

GetMaxDailyAutomaticRestartsOk returns a tuple with the MaxDailyAutomaticRestarts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDailyAutomaticRestarts

`func (o *ModelAPITaskLimitsConfig) SetMaxDailyAutomaticRestarts(v int32)`

SetMaxDailyAutomaticRestarts sets MaxDailyAutomaticRestarts field to given value.

### HasMaxDailyAutomaticRestarts

`func (o *ModelAPITaskLimitsConfig) HasMaxDailyAutomaticRestarts() bool`

HasMaxDailyAutomaticRestarts returns a boolean if a field has been set.

### GetMaxDegradedModeConcurrentLargeParserProjectTasks

`func (o *ModelAPITaskLimitsConfig) GetMaxDegradedModeConcurrentLargeParserProjectTasks() int32`

GetMaxDegradedModeConcurrentLargeParserProjectTasks returns the MaxDegradedModeConcurrentLargeParserProjectTasks field if non-nil, zero value otherwise.

### GetMaxDegradedModeConcurrentLargeParserProjectTasksOk

`func (o *ModelAPITaskLimitsConfig) GetMaxDegradedModeConcurrentLargeParserProjectTasksOk() (*int32, bool)`

GetMaxDegradedModeConcurrentLargeParserProjectTasksOk returns a tuple with the MaxDegradedModeConcurrentLargeParserProjectTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDegradedModeConcurrentLargeParserProjectTasks

`func (o *ModelAPITaskLimitsConfig) SetMaxDegradedModeConcurrentLargeParserProjectTasks(v int32)`

SetMaxDegradedModeConcurrentLargeParserProjectTasks sets MaxDegradedModeConcurrentLargeParserProjectTasks field to given value.

### HasMaxDegradedModeConcurrentLargeParserProjectTasks

`func (o *ModelAPITaskLimitsConfig) HasMaxDegradedModeConcurrentLargeParserProjectTasks() bool`

HasMaxDegradedModeConcurrentLargeParserProjectTasks returns a boolean if a field has been set.

### GetMaxDegradedModeParserProjectSize

`func (o *ModelAPITaskLimitsConfig) GetMaxDegradedModeParserProjectSize() int32`

GetMaxDegradedModeParserProjectSize returns the MaxDegradedModeParserProjectSize field if non-nil, zero value otherwise.

### GetMaxDegradedModeParserProjectSizeOk

`func (o *ModelAPITaskLimitsConfig) GetMaxDegradedModeParserProjectSizeOk() (*int32, bool)`

GetMaxDegradedModeParserProjectSizeOk returns a tuple with the MaxDegradedModeParserProjectSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDegradedModeParserProjectSize

`func (o *ModelAPITaskLimitsConfig) SetMaxDegradedModeParserProjectSize(v int32)`

SetMaxDegradedModeParserProjectSize sets MaxDegradedModeParserProjectSize field to given value.

### HasMaxDegradedModeParserProjectSize

`func (o *ModelAPITaskLimitsConfig) HasMaxDegradedModeParserProjectSize() bool`

HasMaxDegradedModeParserProjectSize returns a boolean if a field has been set.

### GetMaxExecTimeoutSecs

`func (o *ModelAPITaskLimitsConfig) GetMaxExecTimeoutSecs() int32`

GetMaxExecTimeoutSecs returns the MaxExecTimeoutSecs field if non-nil, zero value otherwise.

### GetMaxExecTimeoutSecsOk

`func (o *ModelAPITaskLimitsConfig) GetMaxExecTimeoutSecsOk() (*int32, bool)`

GetMaxExecTimeoutSecsOk returns a tuple with the MaxExecTimeoutSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxExecTimeoutSecs

`func (o *ModelAPITaskLimitsConfig) SetMaxExecTimeoutSecs(v int32)`

SetMaxExecTimeoutSecs sets MaxExecTimeoutSecs field to given value.

### HasMaxExecTimeoutSecs

`func (o *ModelAPITaskLimitsConfig) HasMaxExecTimeoutSecs() bool`

HasMaxExecTimeoutSecs returns a boolean if a field has been set.

### GetMaxGenerateTaskJsonSize

`func (o *ModelAPITaskLimitsConfig) GetMaxGenerateTaskJsonSize() int32`

GetMaxGenerateTaskJsonSize returns the MaxGenerateTaskJsonSize field if non-nil, zero value otherwise.

### GetMaxGenerateTaskJsonSizeOk

`func (o *ModelAPITaskLimitsConfig) GetMaxGenerateTaskJsonSizeOk() (*int32, bool)`

GetMaxGenerateTaskJsonSizeOk returns a tuple with the MaxGenerateTaskJsonSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGenerateTaskJsonSize

`func (o *ModelAPITaskLimitsConfig) SetMaxGenerateTaskJsonSize(v int32)`

SetMaxGenerateTaskJsonSize sets MaxGenerateTaskJsonSize field to given value.

### HasMaxGenerateTaskJsonSize

`func (o *ModelAPITaskLimitsConfig) HasMaxGenerateTaskJsonSize() bool`

HasMaxGenerateTaskJsonSize returns a boolean if a field has been set.

### GetMaxHourlyPatchTasks

`func (o *ModelAPITaskLimitsConfig) GetMaxHourlyPatchTasks() int32`

GetMaxHourlyPatchTasks returns the MaxHourlyPatchTasks field if non-nil, zero value otherwise.

### GetMaxHourlyPatchTasksOk

`func (o *ModelAPITaskLimitsConfig) GetMaxHourlyPatchTasksOk() (*int32, bool)`

GetMaxHourlyPatchTasksOk returns a tuple with the MaxHourlyPatchTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHourlyPatchTasks

`func (o *ModelAPITaskLimitsConfig) SetMaxHourlyPatchTasks(v int32)`

SetMaxHourlyPatchTasks sets MaxHourlyPatchTasks field to given value.

### HasMaxHourlyPatchTasks

`func (o *ModelAPITaskLimitsConfig) HasMaxHourlyPatchTasks() bool`

HasMaxHourlyPatchTasks returns a boolean if a field has been set.

### GetMaxIncludesPerVersion

`func (o *ModelAPITaskLimitsConfig) GetMaxIncludesPerVersion() int32`

GetMaxIncludesPerVersion returns the MaxIncludesPerVersion field if non-nil, zero value otherwise.

### GetMaxIncludesPerVersionOk

`func (o *ModelAPITaskLimitsConfig) GetMaxIncludesPerVersionOk() (*int32, bool)`

GetMaxIncludesPerVersionOk returns a tuple with the MaxIncludesPerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIncludesPerVersion

`func (o *ModelAPITaskLimitsConfig) SetMaxIncludesPerVersion(v int32)`

SetMaxIncludesPerVersion sets MaxIncludesPerVersion field to given value.

### HasMaxIncludesPerVersion

`func (o *ModelAPITaskLimitsConfig) HasMaxIncludesPerVersion() bool`

HasMaxIncludesPerVersion returns a boolean if a field has been set.

### GetMaxParserProjectSize

`func (o *ModelAPITaskLimitsConfig) GetMaxParserProjectSize() int32`

GetMaxParserProjectSize returns the MaxParserProjectSize field if non-nil, zero value otherwise.

### GetMaxParserProjectSizeOk

`func (o *ModelAPITaskLimitsConfig) GetMaxParserProjectSizeOk() (*int32, bool)`

GetMaxParserProjectSizeOk returns a tuple with the MaxParserProjectSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxParserProjectSize

`func (o *ModelAPITaskLimitsConfig) SetMaxParserProjectSize(v int32)`

SetMaxParserProjectSize sets MaxParserProjectSize field to given value.

### HasMaxParserProjectSize

`func (o *ModelAPITaskLimitsConfig) HasMaxParserProjectSize() bool`

HasMaxParserProjectSize returns a boolean if a field has been set.

### GetMaxPendingGeneratedTasks

`func (o *ModelAPITaskLimitsConfig) GetMaxPendingGeneratedTasks() int32`

GetMaxPendingGeneratedTasks returns the MaxPendingGeneratedTasks field if non-nil, zero value otherwise.

### GetMaxPendingGeneratedTasksOk

`func (o *ModelAPITaskLimitsConfig) GetMaxPendingGeneratedTasksOk() (*int32, bool)`

GetMaxPendingGeneratedTasksOk returns a tuple with the MaxPendingGeneratedTasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPendingGeneratedTasks

`func (o *ModelAPITaskLimitsConfig) SetMaxPendingGeneratedTasks(v int32)`

SetMaxPendingGeneratedTasks sets MaxPendingGeneratedTasks field to given value.

### HasMaxPendingGeneratedTasks

`func (o *ModelAPITaskLimitsConfig) HasMaxPendingGeneratedTasks() bool`

HasMaxPendingGeneratedTasks returns a boolean if a field has been set.

### GetMaxTaskExecution

`func (o *ModelAPITaskLimitsConfig) GetMaxTaskExecution() int32`

GetMaxTaskExecution returns the MaxTaskExecution field if non-nil, zero value otherwise.

### GetMaxTaskExecutionOk

`func (o *ModelAPITaskLimitsConfig) GetMaxTaskExecutionOk() (*int32, bool)`

GetMaxTaskExecutionOk returns a tuple with the MaxTaskExecution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTaskExecution

`func (o *ModelAPITaskLimitsConfig) SetMaxTaskExecution(v int32)`

SetMaxTaskExecution sets MaxTaskExecution field to given value.

### HasMaxTaskExecution

`func (o *ModelAPITaskLimitsConfig) HasMaxTaskExecution() bool`

HasMaxTaskExecution returns a boolean if a field has been set.

### GetMaxTasksPerVersion

`func (o *ModelAPITaskLimitsConfig) GetMaxTasksPerVersion() int32`

GetMaxTasksPerVersion returns the MaxTasksPerVersion field if non-nil, zero value otherwise.

### GetMaxTasksPerVersionOk

`func (o *ModelAPITaskLimitsConfig) GetMaxTasksPerVersionOk() (*int32, bool)`

GetMaxTasksPerVersionOk returns a tuple with the MaxTasksPerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTasksPerVersion

`func (o *ModelAPITaskLimitsConfig) SetMaxTasksPerVersion(v int32)`

SetMaxTasksPerVersion sets MaxTasksPerVersion field to given value.

### HasMaxTasksPerVersion

`func (o *ModelAPITaskLimitsConfig) HasMaxTasksPerVersion() bool`

HasMaxTasksPerVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


