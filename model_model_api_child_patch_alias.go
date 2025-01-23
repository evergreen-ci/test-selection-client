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

// checks if the ModelAPIChildPatchAlias type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIChildPatchAlias{}

// ModelAPIChildPatchAlias struct for ModelAPIChildPatchAlias
type ModelAPIChildPatchAlias struct {
	Alias *string `json:"alias,omitempty"`
	PatchId *string `json:"patch_id,omitempty"`
}

// NewModelAPIChildPatchAlias instantiates a new ModelAPIChildPatchAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIChildPatchAlias() *ModelAPIChildPatchAlias {
	this := ModelAPIChildPatchAlias{}
	return &this
}

// NewModelAPIChildPatchAliasWithDefaults instantiates a new ModelAPIChildPatchAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIChildPatchAliasWithDefaults() *ModelAPIChildPatchAlias {
	this := ModelAPIChildPatchAlias{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModelAPIChildPatchAlias) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIChildPatchAlias) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelAPIChildPatchAlias) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModelAPIChildPatchAlias) SetAlias(v string) {
	o.Alias = &v
}

// GetPatchId returns the PatchId field value if set, zero value otherwise.
func (o *ModelAPIChildPatchAlias) GetPatchId() string {
	if o == nil || IsNil(o.PatchId) {
		var ret string
		return ret
	}
	return *o.PatchId
}

// GetPatchIdOk returns a tuple with the PatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIChildPatchAlias) GetPatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.PatchId) {
		return nil, false
	}
	return o.PatchId, true
}

// HasPatchId returns a boolean if a field has been set.
func (o *ModelAPIChildPatchAlias) HasPatchId() bool {
	if o != nil && !IsNil(o.PatchId) {
		return true
	}

	return false
}

// SetPatchId gets a reference to the given string and assigns it to the PatchId field.
func (o *ModelAPIChildPatchAlias) SetPatchId(v string) {
	o.PatchId = &v
}

func (o ModelAPIChildPatchAlias) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIChildPatchAlias) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.PatchId) {
		toSerialize["patch_id"] = o.PatchId
	}
	return toSerialize, nil
}

type NullableModelAPIChildPatchAlias struct {
	value *ModelAPIChildPatchAlias
	isSet bool
}

func (v NullableModelAPIChildPatchAlias) Get() *ModelAPIChildPatchAlias {
	return v.value
}

func (v *NullableModelAPIChildPatchAlias) Set(val *ModelAPIChildPatchAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIChildPatchAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIChildPatchAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIChildPatchAlias(val *ModelAPIChildPatchAlias) *NullableModelAPIChildPatchAlias {
	return &NullableModelAPIChildPatchAlias{value: val, isSet: true}
}

func (v NullableModelAPIChildPatchAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIChildPatchAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


