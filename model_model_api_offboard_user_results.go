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

// checks if the ModelAPIOffboardUserResults type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIOffboardUserResults{}

// ModelAPIOffboardUserResults struct for ModelAPIOffboardUserResults
type ModelAPIOffboardUserResults struct {
	TerminatedHosts []string `json:"terminated_hosts,omitempty"`
	TerminatedVolumes []string `json:"terminated_volumes,omitempty"`
}

// NewModelAPIOffboardUserResults instantiates a new ModelAPIOffboardUserResults object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIOffboardUserResults() *ModelAPIOffboardUserResults {
	this := ModelAPIOffboardUserResults{}
	return &this
}

// NewModelAPIOffboardUserResultsWithDefaults instantiates a new ModelAPIOffboardUserResults object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIOffboardUserResultsWithDefaults() *ModelAPIOffboardUserResults {
	this := ModelAPIOffboardUserResults{}
	return &this
}

// GetTerminatedHosts returns the TerminatedHosts field value if set, zero value otherwise.
func (o *ModelAPIOffboardUserResults) GetTerminatedHosts() []string {
	if o == nil || IsNil(o.TerminatedHosts) {
		var ret []string
		return ret
	}
	return o.TerminatedHosts
}

// GetTerminatedHostsOk returns a tuple with the TerminatedHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIOffboardUserResults) GetTerminatedHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.TerminatedHosts) {
		return nil, false
	}
	return o.TerminatedHosts, true
}

// HasTerminatedHosts returns a boolean if a field has been set.
func (o *ModelAPIOffboardUserResults) HasTerminatedHosts() bool {
	if o != nil && !IsNil(o.TerminatedHosts) {
		return true
	}

	return false
}

// SetTerminatedHosts gets a reference to the given []string and assigns it to the TerminatedHosts field.
func (o *ModelAPIOffboardUserResults) SetTerminatedHosts(v []string) {
	o.TerminatedHosts = v
}

// GetTerminatedVolumes returns the TerminatedVolumes field value if set, zero value otherwise.
func (o *ModelAPIOffboardUserResults) GetTerminatedVolumes() []string {
	if o == nil || IsNil(o.TerminatedVolumes) {
		var ret []string
		return ret
	}
	return o.TerminatedVolumes
}

// GetTerminatedVolumesOk returns a tuple with the TerminatedVolumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIOffboardUserResults) GetTerminatedVolumesOk() ([]string, bool) {
	if o == nil || IsNil(o.TerminatedVolumes) {
		return nil, false
	}
	return o.TerminatedVolumes, true
}

// HasTerminatedVolumes returns a boolean if a field has been set.
func (o *ModelAPIOffboardUserResults) HasTerminatedVolumes() bool {
	if o != nil && !IsNil(o.TerminatedVolumes) {
		return true
	}

	return false
}

// SetTerminatedVolumes gets a reference to the given []string and assigns it to the TerminatedVolumes field.
func (o *ModelAPIOffboardUserResults) SetTerminatedVolumes(v []string) {
	o.TerminatedVolumes = v
}

func (o ModelAPIOffboardUserResults) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIOffboardUserResults) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TerminatedHosts) {
		toSerialize["terminated_hosts"] = o.TerminatedHosts
	}
	if !IsNil(o.TerminatedVolumes) {
		toSerialize["terminated_volumes"] = o.TerminatedVolumes
	}
	return toSerialize, nil
}

type NullableModelAPIOffboardUserResults struct {
	value *ModelAPIOffboardUserResults
	isSet bool
}

func (v NullableModelAPIOffboardUserResults) Get() *ModelAPIOffboardUserResults {
	return v.value
}

func (v *NullableModelAPIOffboardUserResults) Set(val *ModelAPIOffboardUserResults) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIOffboardUserResults) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIOffboardUserResults) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIOffboardUserResults(val *ModelAPIOffboardUserResults) *NullableModelAPIOffboardUserResults {
	return &NullableModelAPIOffboardUserResults{value: val, isSet: true}
}

func (v NullableModelAPIOffboardUserResults) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIOffboardUserResults) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


