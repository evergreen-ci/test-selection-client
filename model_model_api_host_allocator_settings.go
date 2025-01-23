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

// checks if the ModelAPIHostAllocatorSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIHostAllocatorSettings{}

// ModelAPIHostAllocatorSettings struct for ModelAPIHostAllocatorSettings
type ModelAPIHostAllocatorSettings struct {
	AcceptableHostIdleTime *int32 `json:"acceptable_host_idle_time,omitempty"`
	FeedbackRule *string `json:"feedback_rule,omitempty"`
	FutureHostFraction *float32 `json:"future_host_fraction,omitempty"`
	HostsOverallocatedRule *string `json:"hosts_overallocated_rule,omitempty"`
	MaximumHosts *int32 `json:"maximum_hosts,omitempty"`
	MinimumHosts *int32 `json:"minimum_hosts,omitempty"`
	RoundingRule *string `json:"rounding_rule,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewModelAPIHostAllocatorSettings instantiates a new ModelAPIHostAllocatorSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIHostAllocatorSettings() *ModelAPIHostAllocatorSettings {
	this := ModelAPIHostAllocatorSettings{}
	return &this
}

// NewModelAPIHostAllocatorSettingsWithDefaults instantiates a new ModelAPIHostAllocatorSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIHostAllocatorSettingsWithDefaults() *ModelAPIHostAllocatorSettings {
	this := ModelAPIHostAllocatorSettings{}
	return &this
}

// GetAcceptableHostIdleTime returns the AcceptableHostIdleTime field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetAcceptableHostIdleTime() int32 {
	if o == nil || IsNil(o.AcceptableHostIdleTime) {
		var ret int32
		return ret
	}
	return *o.AcceptableHostIdleTime
}

// GetAcceptableHostIdleTimeOk returns a tuple with the AcceptableHostIdleTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetAcceptableHostIdleTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AcceptableHostIdleTime) {
		return nil, false
	}
	return o.AcceptableHostIdleTime, true
}

// HasAcceptableHostIdleTime returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasAcceptableHostIdleTime() bool {
	if o != nil && !IsNil(o.AcceptableHostIdleTime) {
		return true
	}

	return false
}

// SetAcceptableHostIdleTime gets a reference to the given int32 and assigns it to the AcceptableHostIdleTime field.
func (o *ModelAPIHostAllocatorSettings) SetAcceptableHostIdleTime(v int32) {
	o.AcceptableHostIdleTime = &v
}

// GetFeedbackRule returns the FeedbackRule field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetFeedbackRule() string {
	if o == nil || IsNil(o.FeedbackRule) {
		var ret string
		return ret
	}
	return *o.FeedbackRule
}

// GetFeedbackRuleOk returns a tuple with the FeedbackRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetFeedbackRuleOk() (*string, bool) {
	if o == nil || IsNil(o.FeedbackRule) {
		return nil, false
	}
	return o.FeedbackRule, true
}

// HasFeedbackRule returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasFeedbackRule() bool {
	if o != nil && !IsNil(o.FeedbackRule) {
		return true
	}

	return false
}

// SetFeedbackRule gets a reference to the given string and assigns it to the FeedbackRule field.
func (o *ModelAPIHostAllocatorSettings) SetFeedbackRule(v string) {
	o.FeedbackRule = &v
}

// GetFutureHostFraction returns the FutureHostFraction field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetFutureHostFraction() float32 {
	if o == nil || IsNil(o.FutureHostFraction) {
		var ret float32
		return ret
	}
	return *o.FutureHostFraction
}

// GetFutureHostFractionOk returns a tuple with the FutureHostFraction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetFutureHostFractionOk() (*float32, bool) {
	if o == nil || IsNil(o.FutureHostFraction) {
		return nil, false
	}
	return o.FutureHostFraction, true
}

// HasFutureHostFraction returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasFutureHostFraction() bool {
	if o != nil && !IsNil(o.FutureHostFraction) {
		return true
	}

	return false
}

// SetFutureHostFraction gets a reference to the given float32 and assigns it to the FutureHostFraction field.
func (o *ModelAPIHostAllocatorSettings) SetFutureHostFraction(v float32) {
	o.FutureHostFraction = &v
}

// GetHostsOverallocatedRule returns the HostsOverallocatedRule field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetHostsOverallocatedRule() string {
	if o == nil || IsNil(o.HostsOverallocatedRule) {
		var ret string
		return ret
	}
	return *o.HostsOverallocatedRule
}

// GetHostsOverallocatedRuleOk returns a tuple with the HostsOverallocatedRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetHostsOverallocatedRuleOk() (*string, bool) {
	if o == nil || IsNil(o.HostsOverallocatedRule) {
		return nil, false
	}
	return o.HostsOverallocatedRule, true
}

// HasHostsOverallocatedRule returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasHostsOverallocatedRule() bool {
	if o != nil && !IsNil(o.HostsOverallocatedRule) {
		return true
	}

	return false
}

// SetHostsOverallocatedRule gets a reference to the given string and assigns it to the HostsOverallocatedRule field.
func (o *ModelAPIHostAllocatorSettings) SetHostsOverallocatedRule(v string) {
	o.HostsOverallocatedRule = &v
}

// GetMaximumHosts returns the MaximumHosts field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetMaximumHosts() int32 {
	if o == nil || IsNil(o.MaximumHosts) {
		var ret int32
		return ret
	}
	return *o.MaximumHosts
}

// GetMaximumHostsOk returns a tuple with the MaximumHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetMaximumHostsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumHosts) {
		return nil, false
	}
	return o.MaximumHosts, true
}

// HasMaximumHosts returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasMaximumHosts() bool {
	if o != nil && !IsNil(o.MaximumHosts) {
		return true
	}

	return false
}

// SetMaximumHosts gets a reference to the given int32 and assigns it to the MaximumHosts field.
func (o *ModelAPIHostAllocatorSettings) SetMaximumHosts(v int32) {
	o.MaximumHosts = &v
}

// GetMinimumHosts returns the MinimumHosts field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetMinimumHosts() int32 {
	if o == nil || IsNil(o.MinimumHosts) {
		var ret int32
		return ret
	}
	return *o.MinimumHosts
}

// GetMinimumHostsOk returns a tuple with the MinimumHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetMinimumHostsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumHosts) {
		return nil, false
	}
	return o.MinimumHosts, true
}

// HasMinimumHosts returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasMinimumHosts() bool {
	if o != nil && !IsNil(o.MinimumHosts) {
		return true
	}

	return false
}

// SetMinimumHosts gets a reference to the given int32 and assigns it to the MinimumHosts field.
func (o *ModelAPIHostAllocatorSettings) SetMinimumHosts(v int32) {
	o.MinimumHosts = &v
}

// GetRoundingRule returns the RoundingRule field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetRoundingRule() string {
	if o == nil || IsNil(o.RoundingRule) {
		var ret string
		return ret
	}
	return *o.RoundingRule
}

// GetRoundingRuleOk returns a tuple with the RoundingRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetRoundingRuleOk() (*string, bool) {
	if o == nil || IsNil(o.RoundingRule) {
		return nil, false
	}
	return o.RoundingRule, true
}

// HasRoundingRule returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasRoundingRule() bool {
	if o != nil && !IsNil(o.RoundingRule) {
		return true
	}

	return false
}

// SetRoundingRule gets a reference to the given string and assigns it to the RoundingRule field.
func (o *ModelAPIHostAllocatorSettings) SetRoundingRule(v string) {
	o.RoundingRule = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelAPIHostAllocatorSettings) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIHostAllocatorSettings) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelAPIHostAllocatorSettings) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ModelAPIHostAllocatorSettings) SetVersion(v string) {
	o.Version = &v
}

func (o ModelAPIHostAllocatorSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIHostAllocatorSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcceptableHostIdleTime) {
		toSerialize["acceptable_host_idle_time"] = o.AcceptableHostIdleTime
	}
	if !IsNil(o.FeedbackRule) {
		toSerialize["feedback_rule"] = o.FeedbackRule
	}
	if !IsNil(o.FutureHostFraction) {
		toSerialize["future_host_fraction"] = o.FutureHostFraction
	}
	if !IsNil(o.HostsOverallocatedRule) {
		toSerialize["hosts_overallocated_rule"] = o.HostsOverallocatedRule
	}
	if !IsNil(o.MaximumHosts) {
		toSerialize["maximum_hosts"] = o.MaximumHosts
	}
	if !IsNil(o.MinimumHosts) {
		toSerialize["minimum_hosts"] = o.MinimumHosts
	}
	if !IsNil(o.RoundingRule) {
		toSerialize["rounding_rule"] = o.RoundingRule
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableModelAPIHostAllocatorSettings struct {
	value *ModelAPIHostAllocatorSettings
	isSet bool
}

func (v NullableModelAPIHostAllocatorSettings) Get() *ModelAPIHostAllocatorSettings {
	return v.value
}

func (v *NullableModelAPIHostAllocatorSettings) Set(val *ModelAPIHostAllocatorSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIHostAllocatorSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIHostAllocatorSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIHostAllocatorSettings(val *ModelAPIHostAllocatorSettings) *NullableModelAPIHostAllocatorSettings {
	return &NullableModelAPIHostAllocatorSettings{value: val, isSet: true}
}

func (v NullableModelAPIHostAllocatorSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIHostAllocatorSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


