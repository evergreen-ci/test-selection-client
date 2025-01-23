# ModelAPIHostAllocatorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptableHostIdleTime** | Pointer to **int32** |  | [optional] 
**FeedbackRule** | Pointer to **string** |  | [optional] 
**FutureHostFraction** | Pointer to **float32** |  | [optional] 
**HostsOverallocatedRule** | Pointer to **string** |  | [optional] 
**MaximumHosts** | Pointer to **int32** |  | [optional] 
**MinimumHosts** | Pointer to **int32** |  | [optional] 
**RoundingRule** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIHostAllocatorSettings

`func NewModelAPIHostAllocatorSettings() *ModelAPIHostAllocatorSettings`

NewModelAPIHostAllocatorSettings instantiates a new ModelAPIHostAllocatorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIHostAllocatorSettingsWithDefaults

`func NewModelAPIHostAllocatorSettingsWithDefaults() *ModelAPIHostAllocatorSettings`

NewModelAPIHostAllocatorSettingsWithDefaults instantiates a new ModelAPIHostAllocatorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptableHostIdleTime

`func (o *ModelAPIHostAllocatorSettings) GetAcceptableHostIdleTime() int32`

GetAcceptableHostIdleTime returns the AcceptableHostIdleTime field if non-nil, zero value otherwise.

### GetAcceptableHostIdleTimeOk

`func (o *ModelAPIHostAllocatorSettings) GetAcceptableHostIdleTimeOk() (*int32, bool)`

GetAcceptableHostIdleTimeOk returns a tuple with the AcceptableHostIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableHostIdleTime

`func (o *ModelAPIHostAllocatorSettings) SetAcceptableHostIdleTime(v int32)`

SetAcceptableHostIdleTime sets AcceptableHostIdleTime field to given value.

### HasAcceptableHostIdleTime

`func (o *ModelAPIHostAllocatorSettings) HasAcceptableHostIdleTime() bool`

HasAcceptableHostIdleTime returns a boolean if a field has been set.

### GetFeedbackRule

`func (o *ModelAPIHostAllocatorSettings) GetFeedbackRule() string`

GetFeedbackRule returns the FeedbackRule field if non-nil, zero value otherwise.

### GetFeedbackRuleOk

`func (o *ModelAPIHostAllocatorSettings) GetFeedbackRuleOk() (*string, bool)`

GetFeedbackRuleOk returns a tuple with the FeedbackRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackRule

`func (o *ModelAPIHostAllocatorSettings) SetFeedbackRule(v string)`

SetFeedbackRule sets FeedbackRule field to given value.

### HasFeedbackRule

`func (o *ModelAPIHostAllocatorSettings) HasFeedbackRule() bool`

HasFeedbackRule returns a boolean if a field has been set.

### GetFutureHostFraction

`func (o *ModelAPIHostAllocatorSettings) GetFutureHostFraction() float32`

GetFutureHostFraction returns the FutureHostFraction field if non-nil, zero value otherwise.

### GetFutureHostFractionOk

`func (o *ModelAPIHostAllocatorSettings) GetFutureHostFractionOk() (*float32, bool)`

GetFutureHostFractionOk returns a tuple with the FutureHostFraction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFutureHostFraction

`func (o *ModelAPIHostAllocatorSettings) SetFutureHostFraction(v float32)`

SetFutureHostFraction sets FutureHostFraction field to given value.

### HasFutureHostFraction

`func (o *ModelAPIHostAllocatorSettings) HasFutureHostFraction() bool`

HasFutureHostFraction returns a boolean if a field has been set.

### GetHostsOverallocatedRule

`func (o *ModelAPIHostAllocatorSettings) GetHostsOverallocatedRule() string`

GetHostsOverallocatedRule returns the HostsOverallocatedRule field if non-nil, zero value otherwise.

### GetHostsOverallocatedRuleOk

`func (o *ModelAPIHostAllocatorSettings) GetHostsOverallocatedRuleOk() (*string, bool)`

GetHostsOverallocatedRuleOk returns a tuple with the HostsOverallocatedRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostsOverallocatedRule

`func (o *ModelAPIHostAllocatorSettings) SetHostsOverallocatedRule(v string)`

SetHostsOverallocatedRule sets HostsOverallocatedRule field to given value.

### HasHostsOverallocatedRule

`func (o *ModelAPIHostAllocatorSettings) HasHostsOverallocatedRule() bool`

HasHostsOverallocatedRule returns a boolean if a field has been set.

### GetMaximumHosts

`func (o *ModelAPIHostAllocatorSettings) GetMaximumHosts() int32`

GetMaximumHosts returns the MaximumHosts field if non-nil, zero value otherwise.

### GetMaximumHostsOk

`func (o *ModelAPIHostAllocatorSettings) GetMaximumHostsOk() (*int32, bool)`

GetMaximumHostsOk returns a tuple with the MaximumHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumHosts

`func (o *ModelAPIHostAllocatorSettings) SetMaximumHosts(v int32)`

SetMaximumHosts sets MaximumHosts field to given value.

### HasMaximumHosts

`func (o *ModelAPIHostAllocatorSettings) HasMaximumHosts() bool`

HasMaximumHosts returns a boolean if a field has been set.

### GetMinimumHosts

`func (o *ModelAPIHostAllocatorSettings) GetMinimumHosts() int32`

GetMinimumHosts returns the MinimumHosts field if non-nil, zero value otherwise.

### GetMinimumHostsOk

`func (o *ModelAPIHostAllocatorSettings) GetMinimumHostsOk() (*int32, bool)`

GetMinimumHostsOk returns a tuple with the MinimumHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinimumHosts

`func (o *ModelAPIHostAllocatorSettings) SetMinimumHosts(v int32)`

SetMinimumHosts sets MinimumHosts field to given value.

### HasMinimumHosts

`func (o *ModelAPIHostAllocatorSettings) HasMinimumHosts() bool`

HasMinimumHosts returns a boolean if a field has been set.

### GetRoundingRule

`func (o *ModelAPIHostAllocatorSettings) GetRoundingRule() string`

GetRoundingRule returns the RoundingRule field if non-nil, zero value otherwise.

### GetRoundingRuleOk

`func (o *ModelAPIHostAllocatorSettings) GetRoundingRuleOk() (*string, bool)`

GetRoundingRuleOk returns a tuple with the RoundingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundingRule

`func (o *ModelAPIHostAllocatorSettings) SetRoundingRule(v string)`

SetRoundingRule sets RoundingRule field to given value.

### HasRoundingRule

`func (o *ModelAPIHostAllocatorSettings) HasRoundingRule() bool`

HasRoundingRule returns a boolean if a field has been set.

### GetVersion

`func (o *ModelAPIHostAllocatorSettings) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ModelAPIHostAllocatorSettings) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ModelAPIHostAllocatorSettings) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ModelAPIHostAllocatorSettings) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


