/*
Evergreen REST v2 API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelAPISubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPISubscription{}

// ModelAPISubscription struct for ModelAPISubscription
type ModelAPISubscription struct {
	// Identifier for the subscription.
	Id *string `json:"id,omitempty"`
	// The subscription owner.
	Owner *string `json:"owner,omitempty"`
	// Type of subscription owner.
	OwnerType *string `json:"owner_type,omitempty"`
	// List of resource regex selectors.
	RegexSelectors []ModelAPISelector `json:"regex_selectors,omitempty"`
	// Type of resource to subscribe to.
	ResourceType *string `json:"resource_type,omitempty"`
	// List of resource selectors.
	Selectors []ModelAPISelector `json:"selectors,omitempty"`
	// Options for the subscriber.
	Subscriber *ModelAPISubscriber `json:"subscriber,omitempty"`
	// Type of trigger for the subscription.
	Trigger *string `json:"trigger,omitempty"`
	// Data for the particular condition that triggers the subscription.
	TriggerData *map[string]string `json:"trigger_data,omitempty"`
}

// NewModelAPISubscription instantiates a new ModelAPISubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPISubscription() *ModelAPISubscription {
	this := ModelAPISubscription{}
	return &this
}

// NewModelAPISubscriptionWithDefaults instantiates a new ModelAPISubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPISubscriptionWithDefaults() *ModelAPISubscription {
	this := ModelAPISubscription{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelAPISubscription) SetId(v string) {
	o.Id = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *ModelAPISubscription) SetOwner(v string) {
	o.Owner = &v
}

// GetOwnerType returns the OwnerType field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetOwnerType() string {
	if o == nil || IsNil(o.OwnerType) {
		var ret string
		return ret
	}
	return *o.OwnerType
}

// GetOwnerTypeOk returns a tuple with the OwnerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetOwnerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerType) {
		return nil, false
	}
	return o.OwnerType, true
}

// HasOwnerType returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasOwnerType() bool {
	if o != nil && !IsNil(o.OwnerType) {
		return true
	}

	return false
}

// SetOwnerType gets a reference to the given string and assigns it to the OwnerType field.
func (o *ModelAPISubscription) SetOwnerType(v string) {
	o.OwnerType = &v
}

// GetRegexSelectors returns the RegexSelectors field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetRegexSelectors() []ModelAPISelector {
	if o == nil || IsNil(o.RegexSelectors) {
		var ret []ModelAPISelector
		return ret
	}
	return o.RegexSelectors
}

// GetRegexSelectorsOk returns a tuple with the RegexSelectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetRegexSelectorsOk() ([]ModelAPISelector, bool) {
	if o == nil || IsNil(o.RegexSelectors) {
		return nil, false
	}
	return o.RegexSelectors, true
}

// HasRegexSelectors returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasRegexSelectors() bool {
	if o != nil && !IsNil(o.RegexSelectors) {
		return true
	}

	return false
}

// SetRegexSelectors gets a reference to the given []ModelAPISelector and assigns it to the RegexSelectors field.
func (o *ModelAPISubscription) SetRegexSelectors(v []ModelAPISelector) {
	o.RegexSelectors = v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ModelAPISubscription) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSelectors returns the Selectors field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetSelectors() []ModelAPISelector {
	if o == nil || IsNil(o.Selectors) {
		var ret []ModelAPISelector
		return ret
	}
	return o.Selectors
}

// GetSelectorsOk returns a tuple with the Selectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetSelectorsOk() ([]ModelAPISelector, bool) {
	if o == nil || IsNil(o.Selectors) {
		return nil, false
	}
	return o.Selectors, true
}

// HasSelectors returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasSelectors() bool {
	if o != nil && !IsNil(o.Selectors) {
		return true
	}

	return false
}

// SetSelectors gets a reference to the given []ModelAPISelector and assigns it to the Selectors field.
func (o *ModelAPISubscription) SetSelectors(v []ModelAPISelector) {
	o.Selectors = v
}

// GetSubscriber returns the Subscriber field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetSubscriber() ModelAPISubscriber {
	if o == nil || IsNil(o.Subscriber) {
		var ret ModelAPISubscriber
		return ret
	}
	return *o.Subscriber
}

// GetSubscriberOk returns a tuple with the Subscriber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetSubscriberOk() (*ModelAPISubscriber, bool) {
	if o == nil || IsNil(o.Subscriber) {
		return nil, false
	}
	return o.Subscriber, true
}

// HasSubscriber returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasSubscriber() bool {
	if o != nil && !IsNil(o.Subscriber) {
		return true
	}

	return false
}

// SetSubscriber gets a reference to the given ModelAPISubscriber and assigns it to the Subscriber field.
func (o *ModelAPISubscription) SetSubscriber(v ModelAPISubscriber) {
	o.Subscriber = &v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetTrigger() string {
	if o == nil || IsNil(o.Trigger) {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *ModelAPISubscription) SetTrigger(v string) {
	o.Trigger = &v
}

// GetTriggerData returns the TriggerData field value if set, zero value otherwise.
func (o *ModelAPISubscription) GetTriggerData() map[string]string {
	if o == nil || IsNil(o.TriggerData) {
		var ret map[string]string
		return ret
	}
	return *o.TriggerData
}

// GetTriggerDataOk returns a tuple with the TriggerData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPISubscription) GetTriggerDataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.TriggerData) {
		return nil, false
	}
	return o.TriggerData, true
}

// HasTriggerData returns a boolean if a field has been set.
func (o *ModelAPISubscription) HasTriggerData() bool {
	if o != nil && !IsNil(o.TriggerData) {
		return true
	}

	return false
}

// SetTriggerData gets a reference to the given map[string]string and assigns it to the TriggerData field.
func (o *ModelAPISubscription) SetTriggerData(v map[string]string) {
	o.TriggerData = &v
}

func (o ModelAPISubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPISubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.OwnerType) {
		toSerialize["owner_type"] = o.OwnerType
	}
	if !IsNil(o.RegexSelectors) {
		toSerialize["regex_selectors"] = o.RegexSelectors
	}
	if !IsNil(o.ResourceType) {
		toSerialize["resource_type"] = o.ResourceType
	}
	if !IsNil(o.Selectors) {
		toSerialize["selectors"] = o.Selectors
	}
	if !IsNil(o.Subscriber) {
		toSerialize["subscriber"] = o.Subscriber
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !IsNil(o.TriggerData) {
		toSerialize["trigger_data"] = o.TriggerData
	}
	return toSerialize, nil
}

type NullableModelAPISubscription struct {
	value *ModelAPISubscription
	isSet bool
}

func (v NullableModelAPISubscription) Get() *ModelAPISubscription {
	return v.value
}

func (v *NullableModelAPISubscription) Set(val *ModelAPISubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPISubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPISubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPISubscription(val *ModelAPISubscription) *NullableModelAPISubscription {
	return &NullableModelAPISubscription{value: val, isSet: true}
}

func (v NullableModelAPISubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPISubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


