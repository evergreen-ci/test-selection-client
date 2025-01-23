# ModelAPICommitQueueParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enable/disable the commit queue. | [optional] 
**MergeMethod** | Pointer to **string** | Method of merging (squash, merge, or rebase). | [optional] 
**Message** | Pointer to **string** | Message to display when users interact with the commit queue. | [optional] 

## Methods

### NewModelAPICommitQueueParams

`func NewModelAPICommitQueueParams() *ModelAPICommitQueueParams`

NewModelAPICommitQueueParams instantiates a new ModelAPICommitQueueParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPICommitQueueParamsWithDefaults

`func NewModelAPICommitQueueParamsWithDefaults() *ModelAPICommitQueueParams`

NewModelAPICommitQueueParamsWithDefaults instantiates a new ModelAPICommitQueueParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ModelAPICommitQueueParams) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelAPICommitQueueParams) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelAPICommitQueueParams) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelAPICommitQueueParams) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMergeMethod

`func (o *ModelAPICommitQueueParams) GetMergeMethod() string`

GetMergeMethod returns the MergeMethod field if non-nil, zero value otherwise.

### GetMergeMethodOk

`func (o *ModelAPICommitQueueParams) GetMergeMethodOk() (*string, bool)`

GetMergeMethodOk returns a tuple with the MergeMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeMethod

`func (o *ModelAPICommitQueueParams) SetMergeMethod(v string)`

SetMergeMethod sets MergeMethod field to given value.

### HasMergeMethod

`func (o *ModelAPICommitQueueParams) HasMergeMethod() bool`

HasMergeMethod returns a boolean if a field has been set.

### GetMessage

`func (o *ModelAPICommitQueueParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelAPICommitQueueParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelAPICommitQueueParams) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelAPICommitQueueParams) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


