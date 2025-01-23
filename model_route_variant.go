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

// checks if the RouteVariant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteVariant{}

// RouteVariant struct for RouteVariant
type RouteVariant struct {
	Id *string `json:"id,omitempty"`
	Tasks []string `json:"tasks,omitempty"`
}

// NewRouteVariant instantiates a new RouteVariant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteVariant() *RouteVariant {
	this := RouteVariant{}
	return &this
}

// NewRouteVariantWithDefaults instantiates a new RouteVariant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteVariantWithDefaults() *RouteVariant {
	this := RouteVariant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RouteVariant) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteVariant) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RouteVariant) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RouteVariant) SetId(v string) {
	o.Id = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *RouteVariant) GetTasks() []string {
	if o == nil || IsNil(o.Tasks) {
		var ret []string
		return ret
	}
	return o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteVariant) GetTasksOk() ([]string, bool) {
	if o == nil || IsNil(o.Tasks) {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *RouteVariant) HasTasks() bool {
	if o != nil && !IsNil(o.Tasks) {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []string and assigns it to the Tasks field.
func (o *RouteVariant) SetTasks(v []string) {
	o.Tasks = v
}

func (o RouteVariant) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteVariant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Tasks) {
		toSerialize["tasks"] = o.Tasks
	}
	return toSerialize, nil
}

type NullableRouteVariant struct {
	value *RouteVariant
	isSet bool
}

func (v NullableRouteVariant) Get() *RouteVariant {
	return v.value
}

func (v *NullableRouteVariant) Set(val *RouteVariant) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteVariant) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteVariant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteVariant(val *RouteVariant) *NullableRouteVariant {
	return &NullableRouteVariant{value: val, isSet: true}
}

func (v NullableRouteVariant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteVariant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


