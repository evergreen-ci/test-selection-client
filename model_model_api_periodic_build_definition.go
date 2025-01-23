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

// checks if the ModelAPIPeriodicBuildDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIPeriodicBuildDefinition{}

// ModelAPIPeriodicBuildDefinition struct for ModelAPIPeriodicBuildDefinition
type ModelAPIPeriodicBuildDefinition struct {
	// Alias to run for the periodic build.
	Alias *string `json:"alias,omitempty"`
	// Project config file to use for the periodic build.
	ConfigFile *string `json:"config_file,omitempty"`
	// Cron specification for when to run periodic builds.
	Cron *string `json:"cron,omitempty"`
	// Identifier for the periodic build.
	Id *string `json:"id,omitempty"`
	// Interval (in hours) between periodic build runs.
	IntervalHours *int32 `json:"interval_hours,omitempty"`
	// Message to display in the version metadata.
	Message *string `json:"message,omitempty"`
	// Next time that the periodic build will run.
	NextRunTime *string `json:"next_run_time,omitempty"`
}

// NewModelAPIPeriodicBuildDefinition instantiates a new ModelAPIPeriodicBuildDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIPeriodicBuildDefinition() *ModelAPIPeriodicBuildDefinition {
	this := ModelAPIPeriodicBuildDefinition{}
	return &this
}

// NewModelAPIPeriodicBuildDefinitionWithDefaults instantiates a new ModelAPIPeriodicBuildDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIPeriodicBuildDefinitionWithDefaults() *ModelAPIPeriodicBuildDefinition {
	this := ModelAPIPeriodicBuildDefinition{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModelAPIPeriodicBuildDefinition) SetAlias(v string) {
	o.Alias = &v
}

// GetConfigFile returns the ConfigFile field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetConfigFile() string {
	if o == nil || IsNil(o.ConfigFile) {
		var ret string
		return ret
	}
	return *o.ConfigFile
}

// GetConfigFileOk returns a tuple with the ConfigFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetConfigFileOk() (*string, bool) {
	if o == nil || IsNil(o.ConfigFile) {
		return nil, false
	}
	return o.ConfigFile, true
}

// HasConfigFile returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasConfigFile() bool {
	if o != nil && !IsNil(o.ConfigFile) {
		return true
	}

	return false
}

// SetConfigFile gets a reference to the given string and assigns it to the ConfigFile field.
func (o *ModelAPIPeriodicBuildDefinition) SetConfigFile(v string) {
	o.ConfigFile = &v
}

// GetCron returns the Cron field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetCron() string {
	if o == nil || IsNil(o.Cron) {
		var ret string
		return ret
	}
	return *o.Cron
}

// GetCronOk returns a tuple with the Cron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetCronOk() (*string, bool) {
	if o == nil || IsNil(o.Cron) {
		return nil, false
	}
	return o.Cron, true
}

// HasCron returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasCron() bool {
	if o != nil && !IsNil(o.Cron) {
		return true
	}

	return false
}

// SetCron gets a reference to the given string and assigns it to the Cron field.
func (o *ModelAPIPeriodicBuildDefinition) SetCron(v string) {
	o.Cron = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelAPIPeriodicBuildDefinition) SetId(v string) {
	o.Id = &v
}

// GetIntervalHours returns the IntervalHours field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetIntervalHours() int32 {
	if o == nil || IsNil(o.IntervalHours) {
		var ret int32
		return ret
	}
	return *o.IntervalHours
}

// GetIntervalHoursOk returns a tuple with the IntervalHours field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetIntervalHoursOk() (*int32, bool) {
	if o == nil || IsNil(o.IntervalHours) {
		return nil, false
	}
	return o.IntervalHours, true
}

// HasIntervalHours returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasIntervalHours() bool {
	if o != nil && !IsNil(o.IntervalHours) {
		return true
	}

	return false
}

// SetIntervalHours gets a reference to the given int32 and assigns it to the IntervalHours field.
func (o *ModelAPIPeriodicBuildDefinition) SetIntervalHours(v int32) {
	o.IntervalHours = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ModelAPIPeriodicBuildDefinition) SetMessage(v string) {
	o.Message = &v
}

// GetNextRunTime returns the NextRunTime field value if set, zero value otherwise.
func (o *ModelAPIPeriodicBuildDefinition) GetNextRunTime() string {
	if o == nil || IsNil(o.NextRunTime) {
		var ret string
		return ret
	}
	return *o.NextRunTime
}

// GetNextRunTimeOk returns a tuple with the NextRunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPeriodicBuildDefinition) GetNextRunTimeOk() (*string, bool) {
	if o == nil || IsNil(o.NextRunTime) {
		return nil, false
	}
	return o.NextRunTime, true
}

// HasNextRunTime returns a boolean if a field has been set.
func (o *ModelAPIPeriodicBuildDefinition) HasNextRunTime() bool {
	if o != nil && !IsNil(o.NextRunTime) {
		return true
	}

	return false
}

// SetNextRunTime gets a reference to the given string and assigns it to the NextRunTime field.
func (o *ModelAPIPeriodicBuildDefinition) SetNextRunTime(v string) {
	o.NextRunTime = &v
}

func (o ModelAPIPeriodicBuildDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIPeriodicBuildDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.ConfigFile) {
		toSerialize["config_file"] = o.ConfigFile
	}
	if !IsNil(o.Cron) {
		toSerialize["cron"] = o.Cron
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntervalHours) {
		toSerialize["interval_hours"] = o.IntervalHours
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.NextRunTime) {
		toSerialize["next_run_time"] = o.NextRunTime
	}
	return toSerialize, nil
}

type NullableModelAPIPeriodicBuildDefinition struct {
	value *ModelAPIPeriodicBuildDefinition
	isSet bool
}

func (v NullableModelAPIPeriodicBuildDefinition) Get() *ModelAPIPeriodicBuildDefinition {
	return v.value
}

func (v *NullableModelAPIPeriodicBuildDefinition) Set(val *ModelAPIPeriodicBuildDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIPeriodicBuildDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIPeriodicBuildDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIPeriodicBuildDefinition(val *ModelAPIPeriodicBuildDefinition) *NullableModelAPIPeriodicBuildDefinition {
	return &NullableModelAPIPeriodicBuildDefinition{value: val, isSet: true}
}

func (v NullableModelAPIPeriodicBuildDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIPeriodicBuildDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


