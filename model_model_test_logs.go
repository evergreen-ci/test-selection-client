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

// checks if the ModelTestLogs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelTestLogs{}

// ModelTestLogs struct for ModelTestLogs
type ModelTestLogs struct {
	// Line number in the log file corresponding to information about this test
	LineNum *int32 `json:"line_num,omitempty"`
	RenderingType *string `json:"rendering_type,omitempty"`
	// URL where the log can be fetched
	Url *string `json:"url,omitempty"`
	UrlParsley *string `json:"url_parsley,omitempty"`
	// URL of the unprocessed version of the logs file for this test
	UrlRaw *string `json:"url_raw,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewModelTestLogs instantiates a new ModelTestLogs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelTestLogs() *ModelTestLogs {
	this := ModelTestLogs{}
	return &this
}

// NewModelTestLogsWithDefaults instantiates a new ModelTestLogs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelTestLogsWithDefaults() *ModelTestLogs {
	this := ModelTestLogs{}
	return &this
}

// GetLineNum returns the LineNum field value if set, zero value otherwise.
func (o *ModelTestLogs) GetLineNum() int32 {
	if o == nil || IsNil(o.LineNum) {
		var ret int32
		return ret
	}
	return *o.LineNum
}

// GetLineNumOk returns a tuple with the LineNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTestLogs) GetLineNumOk() (*int32, bool) {
	if o == nil || IsNil(o.LineNum) {
		return nil, false
	}
	return o.LineNum, true
}

// HasLineNum returns a boolean if a field has been set.
func (o *ModelTestLogs) HasLineNum() bool {
	if o != nil && !IsNil(o.LineNum) {
		return true
	}

	return false
}

// SetLineNum gets a reference to the given int32 and assigns it to the LineNum field.
func (o *ModelTestLogs) SetLineNum(v int32) {
	o.LineNum = &v
}

// GetRenderingType returns the RenderingType field value if set, zero value otherwise.
func (o *ModelTestLogs) GetRenderingType() string {
	if o == nil || IsNil(o.RenderingType) {
		var ret string
		return ret
	}
	return *o.RenderingType
}

// GetRenderingTypeOk returns a tuple with the RenderingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTestLogs) GetRenderingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RenderingType) {
		return nil, false
	}
	return o.RenderingType, true
}

// HasRenderingType returns a boolean if a field has been set.
func (o *ModelTestLogs) HasRenderingType() bool {
	if o != nil && !IsNil(o.RenderingType) {
		return true
	}

	return false
}

// SetRenderingType gets a reference to the given string and assigns it to the RenderingType field.
func (o *ModelTestLogs) SetRenderingType(v string) {
	o.RenderingType = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ModelTestLogs) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTestLogs) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ModelTestLogs) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ModelTestLogs) SetUrl(v string) {
	o.Url = &v
}

// GetUrlParsley returns the UrlParsley field value if set, zero value otherwise.
func (o *ModelTestLogs) GetUrlParsley() string {
	if o == nil || IsNil(o.UrlParsley) {
		var ret string
		return ret
	}
	return *o.UrlParsley
}

// GetUrlParsleyOk returns a tuple with the UrlParsley field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTestLogs) GetUrlParsleyOk() (*string, bool) {
	if o == nil || IsNil(o.UrlParsley) {
		return nil, false
	}
	return o.UrlParsley, true
}

// HasUrlParsley returns a boolean if a field has been set.
func (o *ModelTestLogs) HasUrlParsley() bool {
	if o != nil && !IsNil(o.UrlParsley) {
		return true
	}

	return false
}

// SetUrlParsley gets a reference to the given string and assigns it to the UrlParsley field.
func (o *ModelTestLogs) SetUrlParsley(v string) {
	o.UrlParsley = &v
}

// GetUrlRaw returns the UrlRaw field value if set, zero value otherwise.
func (o *ModelTestLogs) GetUrlRaw() string {
	if o == nil || IsNil(o.UrlRaw) {
		var ret string
		return ret
	}
	return *o.UrlRaw
}

// GetUrlRawOk returns a tuple with the UrlRaw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTestLogs) GetUrlRawOk() (*string, bool) {
	if o == nil || IsNil(o.UrlRaw) {
		return nil, false
	}
	return o.UrlRaw, true
}

// HasUrlRaw returns a boolean if a field has been set.
func (o *ModelTestLogs) HasUrlRaw() bool {
	if o != nil && !IsNil(o.UrlRaw) {
		return true
	}

	return false
}

// SetUrlRaw gets a reference to the given string and assigns it to the UrlRaw field.
func (o *ModelTestLogs) SetUrlRaw(v string) {
	o.UrlRaw = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelTestLogs) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelTestLogs) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelTestLogs) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *ModelTestLogs) SetVersion(v int32) {
	o.Version = &v
}

func (o ModelTestLogs) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelTestLogs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LineNum) {
		toSerialize["line_num"] = o.LineNum
	}
	if !IsNil(o.RenderingType) {
		toSerialize["rendering_type"] = o.RenderingType
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.UrlParsley) {
		toSerialize["url_parsley"] = o.UrlParsley
	}
	if !IsNil(o.UrlRaw) {
		toSerialize["url_raw"] = o.UrlRaw
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableModelTestLogs struct {
	value *ModelTestLogs
	isSet bool
}

func (v NullableModelTestLogs) Get() *ModelTestLogs {
	return v.value
}

func (v *NullableModelTestLogs) Set(val *ModelTestLogs) {
	v.value = val
	v.isSet = true
}

func (v NullableModelTestLogs) IsSet() bool {
	return v.isSet
}

func (v *NullableModelTestLogs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelTestLogs(val *ModelTestLogs) *NullableModelTestLogs {
	return &NullableModelTestLogs{value: val, isSet: true}
}

func (v NullableModelTestLogs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelTestLogs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


