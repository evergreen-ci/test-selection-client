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

// checks if the ModelAPIBuildBaronSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIBuildBaronSettings{}

// ModelAPIBuildBaronSettings struct for ModelAPIBuildBaronSettings
type ModelAPIBuildBaronSettings struct {
	BfSuggestionFeaturesUrl *string `json:"bf_suggestion_features_url,omitempty"`
	BfSuggestionPassword *string `json:"bf_suggestion_password,omitempty"`
	BfSuggestionServer *string `json:"bf_suggestion_server,omitempty"`
	BfSuggestionTimeoutSecs *int32 `json:"bf_suggestion_timeout_secs,omitempty"`
	BfSuggestionUsername *string `json:"bf_suggestion_username,omitempty"`
	// Type of ticket to create.
	TicketCreateIssueType *string `json:"ticket_create_issue_type,omitempty"`
	// Jira project where tickets should be created.
	TicketCreateProject *string `json:"ticket_create_project,omitempty"`
	// Jira project to search for tickets.
	TicketSearchProjects []string `json:"ticket_search_projects,omitempty"`
}

// NewModelAPIBuildBaronSettings instantiates a new ModelAPIBuildBaronSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIBuildBaronSettings() *ModelAPIBuildBaronSettings {
	this := ModelAPIBuildBaronSettings{}
	return &this
}

// NewModelAPIBuildBaronSettingsWithDefaults instantiates a new ModelAPIBuildBaronSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIBuildBaronSettingsWithDefaults() *ModelAPIBuildBaronSettings {
	this := ModelAPIBuildBaronSettings{}
	return &this
}

// GetBfSuggestionFeaturesUrl returns the BfSuggestionFeaturesUrl field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionFeaturesUrl() string {
	if o == nil || IsNil(o.BfSuggestionFeaturesUrl) {
		var ret string
		return ret
	}
	return *o.BfSuggestionFeaturesUrl
}

// GetBfSuggestionFeaturesUrlOk returns a tuple with the BfSuggestionFeaturesUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionFeaturesUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BfSuggestionFeaturesUrl) {
		return nil, false
	}
	return o.BfSuggestionFeaturesUrl, true
}

// HasBfSuggestionFeaturesUrl returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasBfSuggestionFeaturesUrl() bool {
	if o != nil && !IsNil(o.BfSuggestionFeaturesUrl) {
		return true
	}

	return false
}

// SetBfSuggestionFeaturesUrl gets a reference to the given string and assigns it to the BfSuggestionFeaturesUrl field.
func (o *ModelAPIBuildBaronSettings) SetBfSuggestionFeaturesUrl(v string) {
	o.BfSuggestionFeaturesUrl = &v
}

// GetBfSuggestionPassword returns the BfSuggestionPassword field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionPassword() string {
	if o == nil || IsNil(o.BfSuggestionPassword) {
		var ret string
		return ret
	}
	return *o.BfSuggestionPassword
}

// GetBfSuggestionPasswordOk returns a tuple with the BfSuggestionPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.BfSuggestionPassword) {
		return nil, false
	}
	return o.BfSuggestionPassword, true
}

// HasBfSuggestionPassword returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasBfSuggestionPassword() bool {
	if o != nil && !IsNil(o.BfSuggestionPassword) {
		return true
	}

	return false
}

// SetBfSuggestionPassword gets a reference to the given string and assigns it to the BfSuggestionPassword field.
func (o *ModelAPIBuildBaronSettings) SetBfSuggestionPassword(v string) {
	o.BfSuggestionPassword = &v
}

// GetBfSuggestionServer returns the BfSuggestionServer field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionServer() string {
	if o == nil || IsNil(o.BfSuggestionServer) {
		var ret string
		return ret
	}
	return *o.BfSuggestionServer
}

// GetBfSuggestionServerOk returns a tuple with the BfSuggestionServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionServerOk() (*string, bool) {
	if o == nil || IsNil(o.BfSuggestionServer) {
		return nil, false
	}
	return o.BfSuggestionServer, true
}

// HasBfSuggestionServer returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasBfSuggestionServer() bool {
	if o != nil && !IsNil(o.BfSuggestionServer) {
		return true
	}

	return false
}

// SetBfSuggestionServer gets a reference to the given string and assigns it to the BfSuggestionServer field.
func (o *ModelAPIBuildBaronSettings) SetBfSuggestionServer(v string) {
	o.BfSuggestionServer = &v
}

// GetBfSuggestionTimeoutSecs returns the BfSuggestionTimeoutSecs field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionTimeoutSecs() int32 {
	if o == nil || IsNil(o.BfSuggestionTimeoutSecs) {
		var ret int32
		return ret
	}
	return *o.BfSuggestionTimeoutSecs
}

// GetBfSuggestionTimeoutSecsOk returns a tuple with the BfSuggestionTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionTimeoutSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.BfSuggestionTimeoutSecs) {
		return nil, false
	}
	return o.BfSuggestionTimeoutSecs, true
}

// HasBfSuggestionTimeoutSecs returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasBfSuggestionTimeoutSecs() bool {
	if o != nil && !IsNil(o.BfSuggestionTimeoutSecs) {
		return true
	}

	return false
}

// SetBfSuggestionTimeoutSecs gets a reference to the given int32 and assigns it to the BfSuggestionTimeoutSecs field.
func (o *ModelAPIBuildBaronSettings) SetBfSuggestionTimeoutSecs(v int32) {
	o.BfSuggestionTimeoutSecs = &v
}

// GetBfSuggestionUsername returns the BfSuggestionUsername field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionUsername() string {
	if o == nil || IsNil(o.BfSuggestionUsername) {
		var ret string
		return ret
	}
	return *o.BfSuggestionUsername
}

// GetBfSuggestionUsernameOk returns a tuple with the BfSuggestionUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetBfSuggestionUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.BfSuggestionUsername) {
		return nil, false
	}
	return o.BfSuggestionUsername, true
}

// HasBfSuggestionUsername returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasBfSuggestionUsername() bool {
	if o != nil && !IsNil(o.BfSuggestionUsername) {
		return true
	}

	return false
}

// SetBfSuggestionUsername gets a reference to the given string and assigns it to the BfSuggestionUsername field.
func (o *ModelAPIBuildBaronSettings) SetBfSuggestionUsername(v string) {
	o.BfSuggestionUsername = &v
}

// GetTicketCreateIssueType returns the TicketCreateIssueType field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetTicketCreateIssueType() string {
	if o == nil || IsNil(o.TicketCreateIssueType) {
		var ret string
		return ret
	}
	return *o.TicketCreateIssueType
}

// GetTicketCreateIssueTypeOk returns a tuple with the TicketCreateIssueType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetTicketCreateIssueTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TicketCreateIssueType) {
		return nil, false
	}
	return o.TicketCreateIssueType, true
}

// HasTicketCreateIssueType returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasTicketCreateIssueType() bool {
	if o != nil && !IsNil(o.TicketCreateIssueType) {
		return true
	}

	return false
}

// SetTicketCreateIssueType gets a reference to the given string and assigns it to the TicketCreateIssueType field.
func (o *ModelAPIBuildBaronSettings) SetTicketCreateIssueType(v string) {
	o.TicketCreateIssueType = &v
}

// GetTicketCreateProject returns the TicketCreateProject field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetTicketCreateProject() string {
	if o == nil || IsNil(o.TicketCreateProject) {
		var ret string
		return ret
	}
	return *o.TicketCreateProject
}

// GetTicketCreateProjectOk returns a tuple with the TicketCreateProject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetTicketCreateProjectOk() (*string, bool) {
	if o == nil || IsNil(o.TicketCreateProject) {
		return nil, false
	}
	return o.TicketCreateProject, true
}

// HasTicketCreateProject returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasTicketCreateProject() bool {
	if o != nil && !IsNil(o.TicketCreateProject) {
		return true
	}

	return false
}

// SetTicketCreateProject gets a reference to the given string and assigns it to the TicketCreateProject field.
func (o *ModelAPIBuildBaronSettings) SetTicketCreateProject(v string) {
	o.TicketCreateProject = &v
}

// GetTicketSearchProjects returns the TicketSearchProjects field value if set, zero value otherwise.
func (o *ModelAPIBuildBaronSettings) GetTicketSearchProjects() []string {
	if o == nil || IsNil(o.TicketSearchProjects) {
		var ret []string
		return ret
	}
	return o.TicketSearchProjects
}

// GetTicketSearchProjectsOk returns a tuple with the TicketSearchProjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIBuildBaronSettings) GetTicketSearchProjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.TicketSearchProjects) {
		return nil, false
	}
	return o.TicketSearchProjects, true
}

// HasTicketSearchProjects returns a boolean if a field has been set.
func (o *ModelAPIBuildBaronSettings) HasTicketSearchProjects() bool {
	if o != nil && !IsNil(o.TicketSearchProjects) {
		return true
	}

	return false
}

// SetTicketSearchProjects gets a reference to the given []string and assigns it to the TicketSearchProjects field.
func (o *ModelAPIBuildBaronSettings) SetTicketSearchProjects(v []string) {
	o.TicketSearchProjects = v
}

func (o ModelAPIBuildBaronSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIBuildBaronSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BfSuggestionFeaturesUrl) {
		toSerialize["bf_suggestion_features_url"] = o.BfSuggestionFeaturesUrl
	}
	if !IsNil(o.BfSuggestionPassword) {
		toSerialize["bf_suggestion_password"] = o.BfSuggestionPassword
	}
	if !IsNil(o.BfSuggestionServer) {
		toSerialize["bf_suggestion_server"] = o.BfSuggestionServer
	}
	if !IsNil(o.BfSuggestionTimeoutSecs) {
		toSerialize["bf_suggestion_timeout_secs"] = o.BfSuggestionTimeoutSecs
	}
	if !IsNil(o.BfSuggestionUsername) {
		toSerialize["bf_suggestion_username"] = o.BfSuggestionUsername
	}
	if !IsNil(o.TicketCreateIssueType) {
		toSerialize["ticket_create_issue_type"] = o.TicketCreateIssueType
	}
	if !IsNil(o.TicketCreateProject) {
		toSerialize["ticket_create_project"] = o.TicketCreateProject
	}
	if !IsNil(o.TicketSearchProjects) {
		toSerialize["ticket_search_projects"] = o.TicketSearchProjects
	}
	return toSerialize, nil
}

type NullableModelAPIBuildBaronSettings struct {
	value *ModelAPIBuildBaronSettings
	isSet bool
}

func (v NullableModelAPIBuildBaronSettings) Get() *ModelAPIBuildBaronSettings {
	return v.value
}

func (v *NullableModelAPIBuildBaronSettings) Set(val *ModelAPIBuildBaronSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIBuildBaronSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIBuildBaronSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIBuildBaronSettings(val *ModelAPIBuildBaronSettings) *NullableModelAPIBuildBaronSettings {
	return &NullableModelAPIBuildBaronSettings{value: val, isSet: true}
}

func (v NullableModelAPIBuildBaronSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIBuildBaronSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


