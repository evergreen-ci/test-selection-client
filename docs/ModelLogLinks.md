# ModelLogLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentLog** | Pointer to **string** | Link to logs created by the agent process | [optional] 
**AllLog** | Pointer to **string** | Link to logs containing merged copy of all other logs | [optional] 
**EventLog** | Pointer to **string** |  | [optional] 
**SystemLog** | Pointer to **string** | Link to logs created by the machine running the task | [optional] 
**TaskLog** | Pointer to **string** | Link to logs created by the task execution | [optional] 

## Methods

### NewModelLogLinks

`func NewModelLogLinks() *ModelLogLinks`

NewModelLogLinks instantiates a new ModelLogLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLogLinksWithDefaults

`func NewModelLogLinksWithDefaults() *ModelLogLinks`

NewModelLogLinksWithDefaults instantiates a new ModelLogLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentLog

`func (o *ModelLogLinks) GetAgentLog() string`

GetAgentLog returns the AgentLog field if non-nil, zero value otherwise.

### GetAgentLogOk

`func (o *ModelLogLinks) GetAgentLogOk() (*string, bool)`

GetAgentLogOk returns a tuple with the AgentLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentLog

`func (o *ModelLogLinks) SetAgentLog(v string)`

SetAgentLog sets AgentLog field to given value.

### HasAgentLog

`func (o *ModelLogLinks) HasAgentLog() bool`

HasAgentLog returns a boolean if a field has been set.

### GetAllLog

`func (o *ModelLogLinks) GetAllLog() string`

GetAllLog returns the AllLog field if non-nil, zero value otherwise.

### GetAllLogOk

`func (o *ModelLogLinks) GetAllLogOk() (*string, bool)`

GetAllLogOk returns a tuple with the AllLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLog

`func (o *ModelLogLinks) SetAllLog(v string)`

SetAllLog sets AllLog field to given value.

### HasAllLog

`func (o *ModelLogLinks) HasAllLog() bool`

HasAllLog returns a boolean if a field has been set.

### GetEventLog

`func (o *ModelLogLinks) GetEventLog() string`

GetEventLog returns the EventLog field if non-nil, zero value otherwise.

### GetEventLogOk

`func (o *ModelLogLinks) GetEventLogOk() (*string, bool)`

GetEventLogOk returns a tuple with the EventLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLog

`func (o *ModelLogLinks) SetEventLog(v string)`

SetEventLog sets EventLog field to given value.

### HasEventLog

`func (o *ModelLogLinks) HasEventLog() bool`

HasEventLog returns a boolean if a field has been set.

### GetSystemLog

`func (o *ModelLogLinks) GetSystemLog() string`

GetSystemLog returns the SystemLog field if non-nil, zero value otherwise.

### GetSystemLogOk

`func (o *ModelLogLinks) GetSystemLogOk() (*string, bool)`

GetSystemLogOk returns a tuple with the SystemLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemLog

`func (o *ModelLogLinks) SetSystemLog(v string)`

SetSystemLog sets SystemLog field to given value.

### HasSystemLog

`func (o *ModelLogLinks) HasSystemLog() bool`

HasSystemLog returns a boolean if a field has been set.

### GetTaskLog

`func (o *ModelLogLinks) GetTaskLog() string`

GetTaskLog returns the TaskLog field if non-nil, zero value otherwise.

### GetTaskLogOk

`func (o *ModelLogLinks) GetTaskLogOk() (*string, bool)`

GetTaskLogOk returns a tuple with the TaskLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskLog

`func (o *ModelLogLinks) SetTaskLog(v string)`

SetTaskLog sets TaskLog field to given value.

### HasTaskLog

`func (o *ModelLogLinks) HasTaskLog() bool`

HasTaskLog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


