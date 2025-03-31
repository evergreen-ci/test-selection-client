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

// checks if the BuildExplanation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildExplanation{}

// BuildExplanation Explanation of the results of a single build variant
type BuildExplanation struct {
	// Build variant name (not display name)
	BuildVariant string `json:"build_variant"`
	Tasks map[string]TaskExplanation `json:"tasks,omitempty"`
}

type _BuildExplanation BuildExplanation

// NewBuildExplanation instantiates a new BuildExplanation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildExplanation(buildVariant string) *BuildExplanation {
	this := BuildExplanation{}
	this.BuildVariant = buildVariant
	return &this
}

// NewBuildExplanationWithDefaults instantiates a new BuildExplanation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildExplanationWithDefaults() *BuildExplanation {
	this := BuildExplanation{}
	return &this
}

// GetBuildVariant returns the BuildVariant field value
func (o *BuildExplanation) GetBuildVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildVariant
}

// GetBuildVariantOk returns a tuple with the BuildVariant field value
// and a boolean to check if the value has been set.
func (o *BuildExplanation) GetBuildVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildVariant, true
}

// SetBuildVariant sets field value
func (o *BuildExplanation) SetBuildVariant(v string) {
	o.BuildVariant = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *BuildExplanation) GetTasks() map[string]TaskExplanation {
	if o == nil || IsNil(o.Tasks) {
		var ret map[string]TaskExplanation
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildExplanation) GetTasksOk() (map[string]TaskExplanation, bool) {
	if o == nil || IsNil(o.Tasks) {
		return map[string]TaskExplanation{}, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *BuildExplanation) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given map[string]TaskExplanation and assigns it to the Tasks field.
func (o *BuildExplanation) SetTasks(v map[string]TaskExplanation) {
	o.Tasks = v
}

func (o BuildExplanation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildExplanation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["build_variant"] = o.BuildVariant
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

func (o *BuildExplanation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"build_variant",
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

	varBuildExplanation := _BuildExplanation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBuildExplanation)

	if err != nil {
		return err
	}

	*o = BuildExplanation(varBuildExplanation)

	return err
}

type NullableBuildExplanation struct {
	value *BuildExplanation
	isSet bool
}

func (v NullableBuildExplanation) Get() *BuildExplanation {
	return v.value
}

func (v *NullableBuildExplanation) Set(val *BuildExplanation) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildExplanation) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildExplanation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildExplanation(val *BuildExplanation) *NullableBuildExplanation {
	return &NullableBuildExplanation{value: val, isSet: true}
}

func (v NullableBuildExplanation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildExplanation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


