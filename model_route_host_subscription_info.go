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

// checks if the RouteHostSubscriptionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteHostSubscriptionInfo{}

// RouteHostSubscriptionInfo struct for RouteHostSubscriptionInfo
type RouteHostSubscriptionInfo struct {
	// The type of subscription to send when the host is running (\"slack\" or \"email\")
	SubscriptionType *string `json:"subscription_type,omitempty"`
}

// NewRouteHostSubscriptionInfo instantiates a new RouteHostSubscriptionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteHostSubscriptionInfo() *RouteHostSubscriptionInfo {
	this := RouteHostSubscriptionInfo{}
	return &this
}

// NewRouteHostSubscriptionInfoWithDefaults instantiates a new RouteHostSubscriptionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteHostSubscriptionInfoWithDefaults() *RouteHostSubscriptionInfo {
	this := RouteHostSubscriptionInfo{}
	return &this
}

// GetSubscriptionType returns the SubscriptionType field value if set, zero value otherwise.
func (o *RouteHostSubscriptionInfo) GetSubscriptionType() string {
	if o == nil || IsNil(o.SubscriptionType) {
		var ret string
		return ret
	}
	return *o.SubscriptionType
}

// GetSubscriptionTypeOk returns a tuple with the SubscriptionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteHostSubscriptionInfo) GetSubscriptionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubscriptionType) {
		return nil, false
	}
	return o.SubscriptionType, true
}

// HasSubscriptionType returns a boolean if a field has been set.
func (o *RouteHostSubscriptionInfo) HasSubscriptionType() bool {
	if o != nil && !IsNil(o.SubscriptionType) {
		return true
	}

	return false
}

// SetSubscriptionType gets a reference to the given string and assigns it to the SubscriptionType field.
func (o *RouteHostSubscriptionInfo) SetSubscriptionType(v string) {
	o.SubscriptionType = &v
}

func (o RouteHostSubscriptionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteHostSubscriptionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscriptionType) {
		toSerialize["subscription_type"] = o.SubscriptionType
	}
	return toSerialize, nil
}

type NullableRouteHostSubscriptionInfo struct {
	value *RouteHostSubscriptionInfo
	isSet bool
}

func (v NullableRouteHostSubscriptionInfo) Get() *RouteHostSubscriptionInfo {
	return v.value
}

func (v *NullableRouteHostSubscriptionInfo) Set(val *RouteHostSubscriptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteHostSubscriptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteHostSubscriptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteHostSubscriptionInfo(val *RouteHostSubscriptionInfo) *NullableRouteHostSubscriptionInfo {
	return &NullableRouteHostSubscriptionInfo{value: val, isSet: true}
}

func (v NullableRouteHostSubscriptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteHostSubscriptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


