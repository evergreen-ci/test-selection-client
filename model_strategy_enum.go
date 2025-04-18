/*
Test Selection Services

Test Selection services, owner: DevProd Services & Integrations team

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// StrategyEnum the model 'StrategyEnum'
type StrategyEnum string

// List of StrategyEnum
const (
	OPTIMISTIC StrategyEnum = "Optimistic"
	STARTS_WITH_T StrategyEnum = "StartsWithT"
)

// All allowed values of StrategyEnum enum
var AllowedStrategyEnumEnumValues = []StrategyEnum{
	"Optimistic",
	"StartsWithT",
}

func (v *StrategyEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StrategyEnum(value)
	for _, existing := range AllowedStrategyEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StrategyEnum", value)
}

// NewStrategyEnumFromValue returns a pointer to a valid StrategyEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStrategyEnumFromValue(v string) (*StrategyEnum, error) {
	ev := StrategyEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StrategyEnum: valid values are %v", v, AllowedStrategyEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StrategyEnum) IsValid() bool {
	for _, existing := range AllowedStrategyEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StrategyEnum value
func (v StrategyEnum) Ptr() *StrategyEnum {
	return &v
}

type NullableStrategyEnum struct {
	value *StrategyEnum
	isSet bool
}

func (v NullableStrategyEnum) Get() *StrategyEnum {
	return v.value
}

func (v *NullableStrategyEnum) Set(val *StrategyEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableStrategyEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableStrategyEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStrategyEnum(val *StrategyEnum) *NullableStrategyEnum {
	return &NullableStrategyEnum{value: val, isSet: true}
}

func (v NullableStrategyEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStrategyEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

