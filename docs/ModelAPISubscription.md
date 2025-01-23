# ModelAPISubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier for the subscription. | [optional] 
**Owner** | Pointer to **string** | The subscription owner. | [optional] 
**OwnerType** | Pointer to **string** | Type of subscription owner. | [optional] 
**RegexSelectors** | Pointer to [**[]ModelAPISelector**](ModelAPISelector.md) | List of resource regex selectors. | [optional] 
**ResourceType** | Pointer to **string** | Type of resource to subscribe to. | [optional] 
**Selectors** | Pointer to [**[]ModelAPISelector**](ModelAPISelector.md) | List of resource selectors. | [optional] 
**Subscriber** | Pointer to [**ModelAPISubscriber**](ModelAPISubscriber.md) | Options for the subscriber. | [optional] 
**Trigger** | Pointer to **string** | Type of trigger for the subscription. | [optional] 
**TriggerData** | Pointer to **map[string]string** | Data for the particular condition that triggers the subscription. | [optional] 

## Methods

### NewModelAPISubscription

`func NewModelAPISubscription() *ModelAPISubscription`

NewModelAPISubscription instantiates a new ModelAPISubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPISubscriptionWithDefaults

`func NewModelAPISubscriptionWithDefaults() *ModelAPISubscription`

NewModelAPISubscriptionWithDefaults instantiates a new ModelAPISubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelAPISubscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPISubscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPISubscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPISubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOwner

`func (o *ModelAPISubscription) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ModelAPISubscription) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ModelAPISubscription) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ModelAPISubscription) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetOwnerType

`func (o *ModelAPISubscription) GetOwnerType() string`

GetOwnerType returns the OwnerType field if non-nil, zero value otherwise.

### GetOwnerTypeOk

`func (o *ModelAPISubscription) GetOwnerTypeOk() (*string, bool)`

GetOwnerTypeOk returns a tuple with the OwnerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerType

`func (o *ModelAPISubscription) SetOwnerType(v string)`

SetOwnerType sets OwnerType field to given value.

### HasOwnerType

`func (o *ModelAPISubscription) HasOwnerType() bool`

HasOwnerType returns a boolean if a field has been set.

### GetRegexSelectors

`func (o *ModelAPISubscription) GetRegexSelectors() []ModelAPISelector`

GetRegexSelectors returns the RegexSelectors field if non-nil, zero value otherwise.

### GetRegexSelectorsOk

`func (o *ModelAPISubscription) GetRegexSelectorsOk() (*[]ModelAPISelector, bool)`

GetRegexSelectorsOk returns a tuple with the RegexSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexSelectors

`func (o *ModelAPISubscription) SetRegexSelectors(v []ModelAPISelector)`

SetRegexSelectors sets RegexSelectors field to given value.

### HasRegexSelectors

`func (o *ModelAPISubscription) HasRegexSelectors() bool`

HasRegexSelectors returns a boolean if a field has been set.

### GetResourceType

`func (o *ModelAPISubscription) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ModelAPISubscription) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ModelAPISubscription) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *ModelAPISubscription) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetSelectors

`func (o *ModelAPISubscription) GetSelectors() []ModelAPISelector`

GetSelectors returns the Selectors field if non-nil, zero value otherwise.

### GetSelectorsOk

`func (o *ModelAPISubscription) GetSelectorsOk() (*[]ModelAPISelector, bool)`

GetSelectorsOk returns a tuple with the Selectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectors

`func (o *ModelAPISubscription) SetSelectors(v []ModelAPISelector)`

SetSelectors sets Selectors field to given value.

### HasSelectors

`func (o *ModelAPISubscription) HasSelectors() bool`

HasSelectors returns a boolean if a field has been set.

### GetSubscriber

`func (o *ModelAPISubscription) GetSubscriber() ModelAPISubscriber`

GetSubscriber returns the Subscriber field if non-nil, zero value otherwise.

### GetSubscriberOk

`func (o *ModelAPISubscription) GetSubscriberOk() (*ModelAPISubscriber, bool)`

GetSubscriberOk returns a tuple with the Subscriber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriber

`func (o *ModelAPISubscription) SetSubscriber(v ModelAPISubscriber)`

SetSubscriber sets Subscriber field to given value.

### HasSubscriber

`func (o *ModelAPISubscription) HasSubscriber() bool`

HasSubscriber returns a boolean if a field has been set.

### GetTrigger

`func (o *ModelAPISubscription) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ModelAPISubscription) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ModelAPISubscription) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *ModelAPISubscription) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetTriggerData

`func (o *ModelAPISubscription) GetTriggerData() map[string]string`

GetTriggerData returns the TriggerData field if non-nil, zero value otherwise.

### GetTriggerDataOk

`func (o *ModelAPISubscription) GetTriggerDataOk() (*map[string]string, bool)`

GetTriggerDataOk returns a tuple with the TriggerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerData

`func (o *ModelAPISubscription) SetTriggerData(v map[string]string)`

SetTriggerData sets TriggerData field to given value.

### HasTriggerData

`func (o *ModelAPISubscription) HasTriggerData() bool`

HasTriggerData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


