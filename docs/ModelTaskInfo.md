# ModelTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildId** | Pointer to **string** | Unique identifier for the build of the project that this task is run as part of | [optional] 
**DispatchTime** | Pointer to **string** | Time that this task was dispatched to this host | [optional] 
**Name** | Pointer to **string** | The name of this task | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**TaskId** | Pointer to **string** | Unique Identifier of this task. Can be used to fetch more informaiton about this task | [optional] 
**VersionId** | Pointer to **string** | Unique identifier for the version of the project that this task is run as part of | [optional] 

## Methods

### NewModelTaskInfo

`func NewModelTaskInfo() *ModelTaskInfo`

NewModelTaskInfo instantiates a new ModelTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelTaskInfoWithDefaults

`func NewModelTaskInfoWithDefaults() *ModelTaskInfo`

NewModelTaskInfoWithDefaults instantiates a new ModelTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildId

`func (o *ModelTaskInfo) GetBuildId() string`

GetBuildId returns the BuildId field if non-nil, zero value otherwise.

### GetBuildIdOk

`func (o *ModelTaskInfo) GetBuildIdOk() (*string, bool)`

GetBuildIdOk returns a tuple with the BuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildId

`func (o *ModelTaskInfo) SetBuildId(v string)`

SetBuildId sets BuildId field to given value.

### HasBuildId

`func (o *ModelTaskInfo) HasBuildId() bool`

HasBuildId returns a boolean if a field has been set.

### GetDispatchTime

`func (o *ModelTaskInfo) GetDispatchTime() string`

GetDispatchTime returns the DispatchTime field if non-nil, zero value otherwise.

### GetDispatchTimeOk

`func (o *ModelTaskInfo) GetDispatchTimeOk() (*string, bool)`

GetDispatchTimeOk returns a tuple with the DispatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchTime

`func (o *ModelTaskInfo) SetDispatchTime(v string)`

SetDispatchTime sets DispatchTime field to given value.

### HasDispatchTime

`func (o *ModelTaskInfo) HasDispatchTime() bool`

HasDispatchTime returns a boolean if a field has been set.

### GetName

`func (o *ModelTaskInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelTaskInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelTaskInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelTaskInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelTaskInfo) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelTaskInfo) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelTaskInfo) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelTaskInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetTaskId

`func (o *ModelTaskInfo) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ModelTaskInfo) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ModelTaskInfo) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ModelTaskInfo) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetVersionId

`func (o *ModelTaskInfo) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ModelTaskInfo) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ModelTaskInfo) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *ModelTaskInfo) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


