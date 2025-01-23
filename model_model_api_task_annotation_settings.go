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

// checks if the ModelAPITaskAnnotationSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPITaskAnnotationSettings{}

// ModelAPITaskAnnotationSettings struct for ModelAPITaskAnnotationSettings
type ModelAPITaskAnnotationSettings struct {
	// Options for webhooks.
	WebHook *ModelAPIWebHook `json:"web_hook,omitempty"`
}

// NewModelAPITaskAnnotationSettings instantiates a new ModelAPITaskAnnotationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPITaskAnnotationSettings() *ModelAPITaskAnnotationSettings {
	this := ModelAPITaskAnnotationSettings{}
	return &this
}

// NewModelAPITaskAnnotationSettingsWithDefaults instantiates a new ModelAPITaskAnnotationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPITaskAnnotationSettingsWithDefaults() *ModelAPITaskAnnotationSettings {
	this := ModelAPITaskAnnotationSettings{}
	return &this
}

// GetWebHook returns the WebHook field value if set, zero value otherwise.
func (o *ModelAPITaskAnnotationSettings) GetWebHook() ModelAPIWebHook {
	if o == nil || IsNil(o.WebHook) {
		var ret ModelAPIWebHook
		return ret
	}
	return *o.WebHook
}

// GetWebHookOk returns a tuple with the WebHook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPITaskAnnotationSettings) GetWebHookOk() (*ModelAPIWebHook, bool) {
	if o == nil || IsNil(o.WebHook) {
		return nil, false
	}
	return o.WebHook, true
}

// HasWebHook returns a boolean if a field has been set.
func (o *ModelAPITaskAnnotationSettings) HasWebHook() bool {
	if o != nil && !IsNil(o.WebHook) {
		return true
	}

	return false
}

// SetWebHook gets a reference to the given ModelAPIWebHook and assigns it to the WebHook field.
func (o *ModelAPITaskAnnotationSettings) SetWebHook(v ModelAPIWebHook) {
	o.WebHook = &v
}

func (o ModelAPITaskAnnotationSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPITaskAnnotationSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebHook) {
		toSerialize["web_hook"] = o.WebHook
	}
	return toSerialize, nil
}

type NullableModelAPITaskAnnotationSettings struct {
	value *ModelAPITaskAnnotationSettings
	isSet bool
}

func (v NullableModelAPITaskAnnotationSettings) Get() *ModelAPITaskAnnotationSettings {
	return v.value
}

func (v *NullableModelAPITaskAnnotationSettings) Set(val *ModelAPITaskAnnotationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPITaskAnnotationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPITaskAnnotationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPITaskAnnotationSettings(val *ModelAPITaskAnnotationSettings) *NullableModelAPITaskAnnotationSettings {
	return &NullableModelAPITaskAnnotationSettings{value: val, isSet: true}
}

func (v NullableModelAPITaskAnnotationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPITaskAnnotationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


