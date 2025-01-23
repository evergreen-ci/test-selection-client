# ModelApiTaskEndDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Desc** | Pointer to **string** | Description of the final status of this task. | [optional] 
**DiskDevices** | Pointer to **[]string** |  | [optional] 
**FailingCommand** | Pointer to **string** | Command which indiciates the task failure. | [optional] 
**FailureMetadataTags** | Pointer to **[]string** | FailureMetadataTags contains the metadata tags associated with the command that caused the task to fail. These are not set if the task succeeded. | [optional] 
**OomTrackerInfo** | Pointer to [**ModelAPIOomTrackerInfo**](ModelAPIOomTrackerInfo.md) |  | [optional] 
**OtherFailingCommands** | Pointer to [**[]ModelAPIFailingCommand**](ModelAPIFailingCommand.md) | OtherFailingCommands contain information about commands that failed but did not cause the task to fail. | [optional] 
**PostErrored** | Pointer to **bool** | PostErrored is true when the post command errored. | [optional] 
**Status** | Pointer to **string** | The status of the completed task. | [optional] 
**TimedOut** | Pointer to **bool** | Whether this task ended in a timeout. | [optional] 
**TimeoutType** | Pointer to **string** |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | The method by which the task failed. | [optional] 

## Methods

### NewModelApiTaskEndDetail

`func NewModelApiTaskEndDetail() *ModelApiTaskEndDetail`

NewModelApiTaskEndDetail instantiates a new ModelApiTaskEndDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelApiTaskEndDetailWithDefaults

`func NewModelApiTaskEndDetailWithDefaults() *ModelApiTaskEndDetail`

NewModelApiTaskEndDetailWithDefaults instantiates a new ModelApiTaskEndDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesc

`func (o *ModelApiTaskEndDetail) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *ModelApiTaskEndDetail) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *ModelApiTaskEndDetail) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *ModelApiTaskEndDetail) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetDiskDevices

`func (o *ModelApiTaskEndDetail) GetDiskDevices() []string`

GetDiskDevices returns the DiskDevices field if non-nil, zero value otherwise.

### GetDiskDevicesOk

`func (o *ModelApiTaskEndDetail) GetDiskDevicesOk() (*[]string, bool)`

GetDiskDevicesOk returns a tuple with the DiskDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskDevices

`func (o *ModelApiTaskEndDetail) SetDiskDevices(v []string)`

SetDiskDevices sets DiskDevices field to given value.

### HasDiskDevices

`func (o *ModelApiTaskEndDetail) HasDiskDevices() bool`

HasDiskDevices returns a boolean if a field has been set.

### GetFailingCommand

`func (o *ModelApiTaskEndDetail) GetFailingCommand() string`

GetFailingCommand returns the FailingCommand field if non-nil, zero value otherwise.

### GetFailingCommandOk

`func (o *ModelApiTaskEndDetail) GetFailingCommandOk() (*string, bool)`

GetFailingCommandOk returns a tuple with the FailingCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailingCommand

`func (o *ModelApiTaskEndDetail) SetFailingCommand(v string)`

SetFailingCommand sets FailingCommand field to given value.

### HasFailingCommand

`func (o *ModelApiTaskEndDetail) HasFailingCommand() bool`

HasFailingCommand returns a boolean if a field has been set.

### GetFailureMetadataTags

`func (o *ModelApiTaskEndDetail) GetFailureMetadataTags() []string`

GetFailureMetadataTags returns the FailureMetadataTags field if non-nil, zero value otherwise.

### GetFailureMetadataTagsOk

`func (o *ModelApiTaskEndDetail) GetFailureMetadataTagsOk() (*[]string, bool)`

GetFailureMetadataTagsOk returns a tuple with the FailureMetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureMetadataTags

`func (o *ModelApiTaskEndDetail) SetFailureMetadataTags(v []string)`

SetFailureMetadataTags sets FailureMetadataTags field to given value.

### HasFailureMetadataTags

`func (o *ModelApiTaskEndDetail) HasFailureMetadataTags() bool`

HasFailureMetadataTags returns a boolean if a field has been set.

### GetOomTrackerInfo

`func (o *ModelApiTaskEndDetail) GetOomTrackerInfo() ModelAPIOomTrackerInfo`

GetOomTrackerInfo returns the OomTrackerInfo field if non-nil, zero value otherwise.

### GetOomTrackerInfoOk

`func (o *ModelApiTaskEndDetail) GetOomTrackerInfoOk() (*ModelAPIOomTrackerInfo, bool)`

GetOomTrackerInfoOk returns a tuple with the OomTrackerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOomTrackerInfo

`func (o *ModelApiTaskEndDetail) SetOomTrackerInfo(v ModelAPIOomTrackerInfo)`

SetOomTrackerInfo sets OomTrackerInfo field to given value.

### HasOomTrackerInfo

`func (o *ModelApiTaskEndDetail) HasOomTrackerInfo() bool`

HasOomTrackerInfo returns a boolean if a field has been set.

### GetOtherFailingCommands

`func (o *ModelApiTaskEndDetail) GetOtherFailingCommands() []ModelAPIFailingCommand`

GetOtherFailingCommands returns the OtherFailingCommands field if non-nil, zero value otherwise.

### GetOtherFailingCommandsOk

`func (o *ModelApiTaskEndDetail) GetOtherFailingCommandsOk() (*[]ModelAPIFailingCommand, bool)`

GetOtherFailingCommandsOk returns a tuple with the OtherFailingCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFailingCommands

`func (o *ModelApiTaskEndDetail) SetOtherFailingCommands(v []ModelAPIFailingCommand)`

SetOtherFailingCommands sets OtherFailingCommands field to given value.

### HasOtherFailingCommands

`func (o *ModelApiTaskEndDetail) HasOtherFailingCommands() bool`

HasOtherFailingCommands returns a boolean if a field has been set.

### GetPostErrored

`func (o *ModelApiTaskEndDetail) GetPostErrored() bool`

GetPostErrored returns the PostErrored field if non-nil, zero value otherwise.

### GetPostErroredOk

`func (o *ModelApiTaskEndDetail) GetPostErroredOk() (*bool, bool)`

GetPostErroredOk returns a tuple with the PostErrored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostErrored

`func (o *ModelApiTaskEndDetail) SetPostErrored(v bool)`

SetPostErrored sets PostErrored field to given value.

### HasPostErrored

`func (o *ModelApiTaskEndDetail) HasPostErrored() bool`

HasPostErrored returns a boolean if a field has been set.

### GetStatus

`func (o *ModelApiTaskEndDetail) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelApiTaskEndDetail) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelApiTaskEndDetail) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelApiTaskEndDetail) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimedOut

`func (o *ModelApiTaskEndDetail) GetTimedOut() bool`

GetTimedOut returns the TimedOut field if non-nil, zero value otherwise.

### GetTimedOutOk

`func (o *ModelApiTaskEndDetail) GetTimedOutOk() (*bool, bool)`

GetTimedOutOk returns a tuple with the TimedOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimedOut

`func (o *ModelApiTaskEndDetail) SetTimedOut(v bool)`

SetTimedOut sets TimedOut field to given value.

### HasTimedOut

`func (o *ModelApiTaskEndDetail) HasTimedOut() bool`

HasTimedOut returns a boolean if a field has been set.

### GetTimeoutType

`func (o *ModelApiTaskEndDetail) GetTimeoutType() string`

GetTimeoutType returns the TimeoutType field if non-nil, zero value otherwise.

### GetTimeoutTypeOk

`func (o *ModelApiTaskEndDetail) GetTimeoutTypeOk() (*string, bool)`

GetTimeoutTypeOk returns a tuple with the TimeoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutType

`func (o *ModelApiTaskEndDetail) SetTimeoutType(v string)`

SetTimeoutType sets TimeoutType field to given value.

### HasTimeoutType

`func (o *ModelApiTaskEndDetail) HasTimeoutType() bool`

HasTimeoutType returns a boolean if a field has been set.

### GetTraceId

`func (o *ModelApiTaskEndDetail) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ModelApiTaskEndDetail) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ModelApiTaskEndDetail) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ModelApiTaskEndDetail) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *ModelApiTaskEndDetail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelApiTaskEndDetail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelApiTaskEndDetail) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelApiTaskEndDetail) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


