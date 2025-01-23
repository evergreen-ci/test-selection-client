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

// checks if the RouteTaskRestartHandler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RouteTaskRestartHandler{}

// RouteTaskRestartHandler struct for RouteTaskRestartHandler
type RouteTaskRestartHandler struct {
	// If set for a display task, restarts only failed execution tasks. When used with a non-display task, this parameter has no effect.
	FailedOnly *bool `json:"failed_only,omitempty"`
}

// NewRouteTaskRestartHandler instantiates a new RouteTaskRestartHandler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRouteTaskRestartHandler() *RouteTaskRestartHandler {
	this := RouteTaskRestartHandler{}
	return &this
}

// NewRouteTaskRestartHandlerWithDefaults instantiates a new RouteTaskRestartHandler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRouteTaskRestartHandlerWithDefaults() *RouteTaskRestartHandler {
	this := RouteTaskRestartHandler{}
	return &this
}

// GetFailedOnly returns the FailedOnly field value if set, zero value otherwise.
func (o *RouteTaskRestartHandler) GetFailedOnly() bool {
	if o == nil || IsNil(o.FailedOnly) {
		var ret bool
		return ret
	}
	return *o.FailedOnly
}

// GetFailedOnlyOk returns a tuple with the FailedOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RouteTaskRestartHandler) GetFailedOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.FailedOnly) {
		return nil, false
	}
	return o.FailedOnly, true
}

// HasFailedOnly returns a boolean if a field has been set.
func (o *RouteTaskRestartHandler) HasFailedOnly() bool {
	if o != nil && !IsNil(o.FailedOnly) {
		return true
	}

	return false
}

// SetFailedOnly gets a reference to the given bool and assigns it to the FailedOnly field.
func (o *RouteTaskRestartHandler) SetFailedOnly(v bool) {
	o.FailedOnly = &v
}

func (o RouteTaskRestartHandler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RouteTaskRestartHandler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FailedOnly) {
		toSerialize["failed_only"] = o.FailedOnly
	}
	return toSerialize, nil
}

type NullableRouteTaskRestartHandler struct {
	value *RouteTaskRestartHandler
	isSet bool
}

func (v NullableRouteTaskRestartHandler) Get() *RouteTaskRestartHandler {
	return v.value
}

func (v *NullableRouteTaskRestartHandler) Set(val *RouteTaskRestartHandler) {
	v.value = val
	v.isSet = true
}

func (v NullableRouteTaskRestartHandler) IsSet() bool {
	return v.isSet
}

func (v *NullableRouteTaskRestartHandler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRouteTaskRestartHandler(val *RouteTaskRestartHandler) *NullableRouteTaskRestartHandler {
	return &NullableRouteTaskRestartHandler{value: val, isSet: true}
}

func (v NullableRouteTaskRestartHandler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRouteTaskRestartHandler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


