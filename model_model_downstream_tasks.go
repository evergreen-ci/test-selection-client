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

// checks if the ModelDownstreamTasks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDownstreamTasks{}

// ModelDownstreamTasks struct for ModelDownstreamTasks
type ModelDownstreamTasks struct {
	Project *string `json:"project,omitempty"`
	Tasks []string `json:"tasks,omitempty"`
	VariantTasks []ModelVariantTask `json:"variant_tasks,omitempty"`
}

// NewModelDownstreamTasks instantiates a new ModelDownstreamTasks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDownstreamTasks() *ModelDownstreamTasks {
	this := ModelDownstreamTasks{}
	return &this
}

// NewModelDownstreamTasksWithDefaults instantiates a new ModelDownstreamTasks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDownstreamTasksWithDefaults() *ModelDownstreamTasks {
	this := ModelDownstreamTasks{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ModelDownstreamTasks) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDownstreamTasks) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ModelDownstreamTasks) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *ModelDownstreamTasks) SetProject(v string) {
	o.Project = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *ModelDownstreamTasks) GetTasks() []string {
	if o == nil || IsNil(o.Tasks) {
		var ret []string
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDownstreamTasks) GetTasksOk() ([]string, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *ModelDownstreamTasks) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []string and assigns it to the Tasks field.
func (o *ModelDownstreamTasks) SetTasks(v []string) {
	o.Tasks = v
}

// GetVariantTasks returns the VariantTasks field value if set, zero value otherwise.
func (o *ModelDownstreamTasks) GetVariantTasks() []ModelVariantTask {
	if o == nil || IsNil(o.VariantTasks) {
		var ret []ModelVariantTask
		return ret
	}
	return o.VariantTasks
}

// GetVariantTasksOk returns a tuple with the VariantTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelDownstreamTasks) GetVariantTasksOk() ([]ModelVariantTask, bool) {
	if o == nil || IsNil(o.VariantTasks) {
		return nil, false
	}
	return o.VariantTasks, true
}

// HasVariantTasks returns a boolean if a field has been set.
func (o *ModelDownstreamTasks) HasVariantTasks() bool {
	if o != nil && !IsNil(o.VariantTasks) {
		return true
	}

	return false
}

// SetVariantTasks gets a reference to the given []ModelVariantTask and assigns it to the VariantTasks field.
func (o *ModelDownstreamTasks) SetVariantTasks(v []ModelVariantTask) {
	o.VariantTasks = v
}

func (o ModelDownstreamTasks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDownstreamTasks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	if !IsNil(o.VariantTasks) {
		toSerialize["variant_tasks"] = o.VariantTasks
	}
	return toSerialize, nil
}

type NullableModelDownstreamTasks struct {
	value *ModelDownstreamTasks
	isSet bool
}

func (v NullableModelDownstreamTasks) Get() *ModelDownstreamTasks {
	return v.value
}

func (v *NullableModelDownstreamTasks) Set(val *ModelDownstreamTasks) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDownstreamTasks) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDownstreamTasks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDownstreamTasks(val *ModelDownstreamTasks) *NullableModelDownstreamTasks {
	return &NullableModelDownstreamTasks{value: val, isSet: true}
}

func (v NullableModelDownstreamTasks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDownstreamTasks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


