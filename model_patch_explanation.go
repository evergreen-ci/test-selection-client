/*
Test Selection Services

Test Selection services, owner: DevProd Services & Integrations team

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PatchExplanation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchExplanation{}

// PatchExplanation Explanation of the results of a whole patch  The builds are mapped by build variant name (not display name)
type PatchExplanation struct {
	PatchId string `json:"patch_id"`
	Builds map[string]BuildExplanation `json:"builds,omitempty"`
}

type _PatchExplanation PatchExplanation

// NewPatchExplanation instantiates a new PatchExplanation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchExplanation(patchId string) *PatchExplanation {
	this := PatchExplanation{}
	this.PatchId = patchId
	return &this
}

// NewPatchExplanationWithDefaults instantiates a new PatchExplanation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchExplanationWithDefaults() *PatchExplanation {
	this := PatchExplanation{}
	return &this
}

// GetPatchId returns the PatchId field value
func (o *PatchExplanation) GetPatchId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PatchId
}

// GetPatchIdOk returns a tuple with the PatchId field value
// and a boolean to check if the value has been set.
func (o *PatchExplanation) GetPatchIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatchId, true
}

// SetPatchId sets field value
func (o *PatchExplanation) SetPatchId(v string) {
	o.PatchId = v
}

// GetBuilds returns the Builds field value if set, zero value otherwise.
func (o *PatchExplanation) GetBuilds() map[string]BuildExplanation {
	if o == nil || IsNil(o.Builds) {
		var ret map[string]BuildExplanation
		return ret
	}
	return o.Builds
}

// GetBuildsOk returns a tuple with the Builds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchExplanation) GetBuildsOk() (map[string]BuildExplanation, bool) {
	if o == nil || IsNil(o.Builds) {
		return map[string]BuildExplanation{}, false
	}
	return o.Builds, true
}

// HasBuilds returns a boolean if a field has been set.
func (o *PatchExplanation) HasBuilds() bool {
	if o != nil && !IsNil(o.Builds) {
		return true
	}

	return false
}

// SetBuilds gets a reference to the given map[string]BuildExplanation and assigns it to the Builds field.
func (o *PatchExplanation) SetBuilds(v map[string]BuildExplanation) {
	o.Builds = v
}

func (o PatchExplanation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchExplanation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["patch_id"] = o.PatchId
	if !IsNil(o.Builds) {
		toSerialize["builds"] = o.Builds
	}
	return toSerialize, nil
}

func (o *PatchExplanation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"patch_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPatchExplanation := _PatchExplanation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPatchExplanation)

	if err != nil {
		return err
	}

	*o = PatchExplanation(varPatchExplanation)

	return err
}

type NullablePatchExplanation struct {
	value *PatchExplanation
	isSet bool
}

func (v NullablePatchExplanation) Get() *PatchExplanation {
	return v.value
}

func (v *NullablePatchExplanation) Set(val *PatchExplanation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchExplanation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchExplanation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchExplanation(val *PatchExplanation) *NullablePatchExplanation {
	return &NullablePatchExplanation{value: val, isSet: true}
}

func (v NullablePatchExplanation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchExplanation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


