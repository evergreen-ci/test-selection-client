# RouteHostStopOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShouldKeepOff** | Pointer to **bool** | If this host is an unexpirable host with a sleep schedule, setting this to true will keep the host off (and ignore the sleep schedule) until it&#39;s started back up manually. | [optional] 
**SubscriptionType** | Pointer to **string** | The type of subscription to send when the host is stopped (\&quot;slack\&quot; or \&quot;email\&quot;) | [optional] 

## Methods

### NewRouteHostStopOptions

`func NewRouteHostStopOptions() *RouteHostStopOptions`

NewRouteHostStopOptions instantiates a new RouteHostStopOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteHostStopOptionsWithDefaults

`func NewRouteHostStopOptionsWithDefaults() *RouteHostStopOptions`

NewRouteHostStopOptionsWithDefaults instantiates a new RouteHostStopOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShouldKeepOff

`func (o *RouteHostStopOptions) GetShouldKeepOff() bool`

GetShouldKeepOff returns the ShouldKeepOff field if non-nil, zero value otherwise.

### GetShouldKeepOffOk

`func (o *RouteHostStopOptions) GetShouldKeepOffOk() (*bool, bool)`

GetShouldKeepOffOk returns a tuple with the ShouldKeepOff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldKeepOff

`func (o *RouteHostStopOptions) SetShouldKeepOff(v bool)`

SetShouldKeepOff sets ShouldKeepOff field to given value.

### HasShouldKeepOff

`func (o *RouteHostStopOptions) HasShouldKeepOff() bool`

HasShouldKeepOff returns a boolean if a field has been set.

### GetSubscriptionType

`func (o *RouteHostStopOptions) GetSubscriptionType() string`

GetSubscriptionType returns the SubscriptionType field if non-nil, zero value otherwise.

### GetSubscriptionTypeOk

`func (o *RouteHostStopOptions) GetSubscriptionTypeOk() (*string, bool)`

GetSubscriptionTypeOk returns a tuple with the SubscriptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionType

`func (o *RouteHostStopOptions) SetSubscriptionType(v string)`

SetSubscriptionType sets SubscriptionType field to given value.

### HasSubscriptionType

`func (o *RouteHostStopOptions) HasSubscriptionType() bool`

HasSubscriptionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


