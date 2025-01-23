# ApimodelsTaskEndDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** |  | [optional] 
**DiskDevices** | Pointer to **[]string** |  | [optional] 
**FailingCommand** | Pointer to **string** |  | [optional] 
**FailureMetadataTags** | Pointer to **[]string** | FailureMetadataTags are user metadata tags associated with the command that caused the task to fail. | [optional] 
**Modules** | Pointer to [**ApimodelsModuleCloneInfo**](ApimodelsModuleCloneInfo.md) |  | [optional] 
**OomKiller** | Pointer to [**ApimodelsOOMTrackerInfo**](ApimodelsOOMTrackerInfo.md) |  | [optional] 
**OtherFailingCommands** | Pointer to [**[]ApimodelsFailingCommand**](ApimodelsFailingCommand.md) | OtherFailingCommands contains information about commands that failed while the task was running but did not cause the task to fail. | [optional] 
**PostErrored** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TimedOut** | Pointer to **bool** |  | [optional] 
**TimeoutDuration** | Pointer to **int32** |  | [optional] 
**TimeoutType** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewApimodelsTaskEndDetail

`func NewApimodelsTaskEndDetail() *ApimodelsTaskEndDetail`

NewApimodelsTaskEndDetail instantiates a new ApimodelsTaskEndDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApimodelsTaskEndDetailWithDefaults

`func NewApimodelsTaskEndDetailWithDefaults() *ApimodelsTaskEndDetail`

NewApimodelsTaskEndDetailWithDefaults instantiates a new ApimodelsTaskEndDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *ApimodelsTaskEndDetail) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ApimodelsTaskEndDetail) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ApimodelsTaskEndDetail) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ApimodelsTaskEndDetail) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetDiskDevices

`func (o *ApimodelsTaskEndDetail) GetDiskDevices() []string`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *ApimodelsTaskEndDetail) GetDiskDevicesOk() (*[]string, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *ApimodelsTaskEndDetail) SetDiskDevices(v []string)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *ApimodelsTaskEndDetail) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetFailingCommand

`func (o *ApimodelsTaskEndDetail) GetFailingCommand() string`

GetFailingCommand returns the FailingCommand field if non-nil, zero value otherwise.

### GetFailingCommandOk

`func (o *ApimodelsTaskEndDetail) GetFailingCommandOk() (*string, bool)`

GetFailingCommandOk returns a tuple with the FailingCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingCommand

`func (o *ApimodelsTaskEndDetail) SetFailingCommand(v string)`

SetFailingCommand sets FailingCommand field to given value.

### HasFailingCommand

`func (o *ApimodelsTaskEndDetail) HasFailingCommand() bool`

HasFailingCommand returns a boolean if a field has been set.

### GetFailureMetadataTags

`func (o *ApimodelsTaskEndDetail) GetFailureMetadataTags() []string`

GetFailureMetadataTags returns the FailureMetadataTags field if non-nil, zero value otherwise.

### GetFailureMetadataTagsOk

`func (o *ApimodelsTaskEndDetail) GetFailureMetadataTagsOk() (*[]string, bool)`

GetFailureMetadataTagsOk returns a tuple with the FailureMetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMetadataTags

`func (o *ApimodelsTaskEndDetail) SetFailureMetadataTags(v []string)`

SetFailureMetadataTags sets FailureMetadataTags field to given value.

### HasFailureMetadataTags

`func (o *ApimodelsTaskEndDetail) HasFailureMetadataTags() bool`

HasFailureMetadataTags returns a boolean if a field has been set.

### GetModules

`func (o *ApimodelsTaskEndDetail) GetModules() ApimodelsModuleCloneInfo`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ApimodelsTaskEndDetail) GetModulesOk() (*ApimodelsModuleCloneInfo, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ApimodelsTaskEndDetail) SetModules(v ApimodelsModuleCloneInfo)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ApimodelsTaskEndDetail) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetOomKiller

`func (o *ApimodelsTaskEndDetail) GetOomKiller() ApimodelsOOMTrackerInfo`

GetOomKiller returns the OomKiller field if non-nil, zero value otherwise.

### GetOomKillerOk

`func (o *ApimodelsTaskEndDetail) GetOomKillerOk() (*ApimodelsOOMTrackerInfo, bool)`

GetOomKillerOk returns a tuple with the OomKiller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOomKiller

`func (o *ApimodelsTaskEndDetail) SetOomKiller(v ApimodelsOOMTrackerInfo)`

SetOomKiller sets OomKiller field to given value.

### HasOomKiller

`func (o *ApimodelsTaskEndDetail) HasOomKiller() bool`

HasOomKiller returns a boolean if a field has been set.

### GetOtherFailingCommands

`func (o *ApimodelsTaskEndDetail) GetOtherFailingCommands() []ApimodelsFailingCommand`

GetOtherFailingCommands returns the OtherFailingCommands field if non-nil, zero value otherwise.

### GetOtherFailingCommandsOk

`func (o *ApimodelsTaskEndDetail) GetOtherFailingCommandsOk() (*[]ApimodelsFailingCommand, bool)`

GetOtherFailingCommandsOk returns a tuple with the OtherFailingCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFailingCommands

`func (o *ApimodelsTaskEndDetail) SetOtherFailingCommands(v []ApimodelsFailingCommand)`

SetOtherFailingCommands sets OtherFailingCommands field to given value.

### HasOtherFailingCommands

`func (o *ApimodelsTaskEndDetail) HasOtherFailingCommands() bool`

HasOtherFailingCommands returns a boolean if a field has been set.

### GetPostErrored

`func (o *ApimodelsTaskEndDetail) GetPostErrored() bool`

GetPostErrored returns the PostErrored field if non-nil, zero value otherwise.

### GetPostErroredOk

`func (o *ApimodelsTaskEndDetail) GetPostErroredOk() (*bool, bool)`

GetPostErroredOk returns a tuple with the PostErrored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostErrored

`func (o *ApimodelsTaskEndDetail) SetPostErrored(v bool)`

SetPostErrored sets PostErrored field to given value.

### HasPostErrored

`func (o *ApimodelsTaskEndDetail) HasPostErrored() bool`

HasPostErrored returns a boolean if a field has been set.

### GetStatus

`func (o *ApimodelsTaskEndDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApimodelsTaskEndDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApimodelsTaskEndDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApimodelsTaskEndDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimedOut

`func (o *ApimodelsTaskEndDetail) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *ApimodelsTaskEndDetail) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *ApimodelsTaskEndDetail) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *ApimodelsTaskEndDetail) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetTimeoutDuration

`func (o *ApimodelsTaskEndDetail) GetTimeoutDuration() int32`

GetTimeoutDuration returns the TimeoutDuration field if non-nil, zero value otherwise.

### GetTimeoutDurationOk

`func (o *ApimodelsTaskEndDetail) GetTimeoutDurationOk() (*int32, bool)`

GetTimeoutDurationOk returns a tuple with the TimeoutDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutDuration

`func (o *ApimodelsTaskEndDetail) SetTimeoutDuration(v int32)`

SetTimeoutDuration sets TimeoutDuration field to given value.

### HasTimeoutDuration

`func (o *ApimodelsTaskEndDetail) HasTimeoutDuration() bool`

HasTimeoutDuration returns a boolean if a field has been set.

### GetTimeoutType

`func (o *ApimodelsTaskEndDetail) GetTimeoutType() string`

GetTimeoutType returns the TimeoutType field if non-nil, zero value otherwise.

### GetTimeoutTypeOk

`func (o *ApimodelsTaskEndDetail) GetTimeoutTypeOk() (*string, bool)`

GetTimeoutTypeOk returns a tuple with the TimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutType

`func (o *ApimodelsTaskEndDetail) SetTimeoutType(v string)`

SetTimeoutType sets TimeoutType field to given value.

### HasTimeoutType

`func (o *ApimodelsTaskEndDetail) HasTimeoutType() bool`

HasTimeoutType returns a boolean if a field has been set.

### GetTraceId

`func (o *ApimodelsTaskEndDetail) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ApimodelsTaskEndDetail) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ApimodelsTaskEndDetail) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ApimodelsTaskEndDetail) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *ApimodelsTaskEndDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApimodelsTaskEndDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApimodelsTaskEndDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApimodelsTaskEndDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


