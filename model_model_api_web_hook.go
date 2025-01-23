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

// checks if the ModelAPIWebHook type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIWebHook{}

// ModelAPIWebHook struct for ModelAPIWebHook
type ModelAPIWebHook struct {
	// Webhook endpoint
	Endpoint *string `json:"endpoint,omitempty"`
	// Webhook secret
	Secret *string `json:"secret,omitempty"`
}

// NewModelAPIWebHook instantiates a new ModelAPIWebHook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIWebHook() *ModelAPIWebHook {
	this := ModelAPIWebHook{}
	return &this
}

// NewModelAPIWebHookWithDefaults instantiates a new ModelAPIWebHook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIWebHookWithDefaults() *ModelAPIWebHook {
	this := ModelAPIWebHook{}
	return &this
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *ModelAPIWebHook) GetEndpoint() string {
	if o == nil || IsNil(o.Endpoint) {
		var ret string
		return ret
	}
	return *o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIWebHook) GetEndpointOk() (*string, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *ModelAPIWebHook) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given string and assigns it to the Endpoint field.
func (o *ModelAPIWebHook) SetEndpoint(v string) {
	o.Endpoint = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ModelAPIWebHook) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIWebHook) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ModelAPIWebHook) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ModelAPIWebHook) SetSecret(v string) {
	o.Secret = &v
}

func (o ModelAPIWebHook) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIWebHook) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableModelAPIWebHook struct {
	value *ModelAPIWebHook
	isSet bool
}

func (v NullableModelAPIWebHook) Get() *ModelAPIWebHook {
	return v.value
}

func (v *NullableModelAPIWebHook) Set(val *ModelAPIWebHook) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIWebHook) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIWebHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIWebHook(val *ModelAPIWebHook) *NullableModelAPIWebHook {
	return &NullableModelAPIWebHook{value: val, isSet: true}
}

func (v NullableModelAPIWebHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIWebHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


