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

// checks if the ModelProjectTaskExecutionResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelProjectTaskExecutionResp{}

// ModelProjectTaskExecutionResp struct for ModelProjectTaskExecutionResp
type ModelProjectTaskExecutionResp struct {
	NumCompleted *int32 `json:"num_completed,omitempty"`
}

// NewModelProjectTaskExecutionResp instantiates a new ModelProjectTaskExecutionResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelProjectTaskExecutionResp() *ModelProjectTaskExecutionResp {
	this := ModelProjectTaskExecutionResp{}
	return &this
}

// NewModelProjectTaskExecutionRespWithDefaults instantiates a new ModelProjectTaskExecutionResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelProjectTaskExecutionRespWithDefaults() *ModelProjectTaskExecutionResp {
	this := ModelProjectTaskExecutionResp{}
	return &this
}

// GetNumCompleted returns the NumCompleted field value if set, zero value otherwise.
func (o *ModelProjectTaskExecutionResp) GetNumCompleted() int32 {
	if o == nil || IsNil(o.NumCompleted) {
		var ret int32
		return ret
	}
	return *o.NumCompleted
}

// GetNumCompletedOk returns a tuple with the NumCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelProjectTaskExecutionResp) GetNumCompletedOk() (*int32, bool) {
	if o == nil || IsNil(o.NumCompleted) {
		return nil, false
	}
	return o.NumCompleted, true
}

// HasNumCompleted returns a boolean if a field has been set.
func (o *ModelProjectTaskExecutionResp) HasNumCompleted() bool {
	if o != nil && !IsNil(o.NumCompleted) {
		return true
	}

	return false
}

// SetNumCompleted gets a reference to the given int32 and assigns it to the NumCompleted field.
func (o *ModelProjectTaskExecutionResp) SetNumCompleted(v int32) {
	o.NumCompleted = &v
}

func (o ModelProjectTaskExecutionResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelProjectTaskExecutionResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NumCompleted) {
		toSerialize["num_completed"] = o.NumCompleted
	}
	return toSerialize, nil
}

type NullableModelProjectTaskExecutionResp struct {
	value *ModelProjectTaskExecutionResp
	isSet bool
}

func (v NullableModelProjectTaskExecutionResp) Get() *ModelProjectTaskExecutionResp {
	return v.value
}

func (v *NullableModelProjectTaskExecutionResp) Set(val *ModelProjectTaskExecutionResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelProjectTaskExecutionResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelProjectTaskExecutionResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelProjectTaskExecutionResp(val *ModelProjectTaskExecutionResp) *NullableModelProjectTaskExecutionResp {
	return &NullableModelProjectTaskExecutionResp{value: val, isSet: true}
}

func (v NullableModelProjectTaskExecutionResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelProjectTaskExecutionResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


