# ModelAPITaskCache

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**FailedTestNames** | Pointer to **[]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**TaskEndDetails** | Pointer to [**ApimodelsTaskEndDetail**](ApimodelsTaskEndDetail.md) |  | [optional] 
**TimeTaken** | Pointer to **int32** |  | [optional] 
**TimeTakenMs** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelAPITaskCache

`func NewModelAPITaskCache() *ModelAPITaskCache`

NewModelAPITaskCache instantiates a new ModelAPITaskCache object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITaskCacheWithDefaults

`func NewModelAPITaskCacheWithDefaults() *ModelAPITaskCache`

NewModelAPITaskCacheWithDefaults instantiates a new ModelAPITaskCache object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *ModelAPITaskCache) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ModelAPITaskCache) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ModelAPITaskCache) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *ModelAPITaskCache) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelAPITaskCache) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPITaskCache) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPITaskCache) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPITaskCache) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFailedTestNames

`func (o *ModelAPITaskCache) GetFailedTestNames() []string`

GetFailedTestNames returns the FailedTestNames field if non-nil, zero value otherwise.

### GetFailedTestNamesOk

`func (o *ModelAPITaskCache) GetFailedTestNamesOk() (*[]string, bool)`

GetFailedTestNamesOk returns a tuple with the FailedTestNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedTestNames

`func (o *ModelAPITaskCache) SetFailedTestNames(v []string)`

SetFailedTestNames sets FailedTestNames field to given value.

### HasFailedTestNames

`func (o *ModelAPITaskCache) HasFailedTestNames() bool`

HasFailedTestNames returns a boolean if a field has been set.

### GetId

`func (o *ModelAPITaskCache) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPITaskCache) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPITaskCache) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPITaskCache) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStartTime

`func (o *ModelAPITaskCache) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ModelAPITaskCache) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ModelAPITaskCache) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *ModelAPITaskCache) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPITaskCache) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPITaskCache) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPITaskCache) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPITaskCache) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTaskEndDetails

`func (o *ModelAPITaskCache) GetTaskEndDetails() ApimodelsTaskEndDetail`

GetTaskEndDetails returns the TaskEndDetails field if non-nil, zero value otherwise.

### GetTaskEndDetailsOk

`func (o *ModelAPITaskCache) GetTaskEndDetailsOk() (*ApimodelsTaskEndDetail, bool)`

GetTaskEndDetailsOk returns a tuple with the TaskEndDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskEndDetails

`func (o *ModelAPITaskCache) SetTaskEndDetails(v ApimodelsTaskEndDetail)`

SetTaskEndDetails sets TaskEndDetails field to given value.

### HasTaskEndDetails

`func (o *ModelAPITaskCache) HasTaskEndDetails() bool`

HasTaskEndDetails returns a boolean if a field has been set.

### GetTimeTaken

`func (o *ModelAPITaskCache) GetTimeTaken() int32`

GetTimeTaken returns the TimeTaken field if non-nil, zero value otherwise.

### GetTimeTakenOk

`func (o *ModelAPITaskCache) GetTimeTakenOk() (*int32, bool)`

GetTimeTakenOk returns a tuple with the TimeTaken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTaken

`func (o *ModelAPITaskCache) SetTimeTaken(v int32)`

SetTimeTaken sets TimeTaken field to given value.

### HasTimeTaken

`func (o *ModelAPITaskCache) HasTimeTaken() bool`

HasTimeTaken returns a boolean if a field has been set.

### GetTimeTakenMs

`func (o *ModelAPITaskCache) GetTimeTakenMs() int32`

GetTimeTakenMs returns the TimeTakenMs field if non-nil, zero value otherwise.

### GetTimeTakenMsOk

`func (o *ModelAPITaskCache) GetTimeTakenMsOk() (*int32, bool)`

GetTimeTakenMsOk returns a tuple with the TimeTakenMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeTakenMs

`func (o *ModelAPITaskCache) SetTimeTakenMs(v int32)`

SetTimeTakenMs sets TimeTakenMs field to given value.

### HasTimeTakenMs

`func (o *ModelAPITaskCache) HasTimeTakenMs() bool`

HasTimeTakenMs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


