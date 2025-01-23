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

// checks if the ModelAPIPatchTriggerDefinition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIPatchTriggerDefinition{}

// ModelAPIPatchTriggerDefinition struct for ModelAPIPatchTriggerDefinition
type ModelAPIPatchTriggerDefinition struct {
	// Alias to run in the downstream project.
	Alias *string `json:"alias,omitempty"`
	// ID of the downstream project.
	ChildProjectId *string `json:"child_project_id,omitempty"`
	// Identifier of the downstream project.
	ChildProjectIdentifier *string `json:"child_project_identifier,omitempty"`
	// An optional field representing the revision at which to create the downstream patch. By default, this field is empty and the downstream patch will be based off of its most recent commit.
	DownstreamRevision *string `json:"downstream_revision,omitempty"`
	// Name of the module corresponding to the upstream project in the downstream project's YAML.
	ParentAsModule *string `json:"parent_as_module,omitempty"`
	// Status for the parent patch to conditionally kick off the child patch.
	Status *string `json:"status,omitempty"`
	// List of task specifiers.
	TaskSpecifiers []ModelAPITaskSpecifier `json:"task_specifiers,omitempty"`
	// The list of variants/tasks from the alias that will run in the downstream project.
	VariantsTasks []ModelVariantTask `json:"variants_tasks,omitempty"`
}

// NewModelAPIPatchTriggerDefinition instantiates a new ModelAPIPatchTriggerDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIPatchTriggerDefinition() *ModelAPIPatchTriggerDefinition {
	this := ModelAPIPatchTriggerDefinition{}
	return &this
}

// NewModelAPIPatchTriggerDefinitionWithDefaults instantiates a new ModelAPIPatchTriggerDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIPatchTriggerDefinitionWithDefaults() *ModelAPIPatchTriggerDefinition {
	this := ModelAPIPatchTriggerDefinition{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *ModelAPIPatchTriggerDefinition) SetAlias(v string) {
	o.Alias = &v
}

// GetChildProjectId returns the ChildProjectId field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetChildProjectId() string {
	if o == nil || IsNil(o.ChildProjectId) {
		var ret string
		return ret
	}
	return *o.ChildProjectId
}

// GetChildProjectIdOk returns a tuple with the ChildProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetChildProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChildProjectId) {
		return nil, false
	}
	return o.ChildProjectId, true
}

// HasChildProjectId returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasChildProjectId() bool {
	if o != nil && !IsNil(o.ChildProjectId) {
		return true
	}

	return false
}

// SetChildProjectId gets a reference to the given string and assigns it to the ChildProjectId field.
func (o *ModelAPIPatchTriggerDefinition) SetChildProjectId(v string) {
	o.ChildProjectId = &v
}

// GetChildProjectIdentifier returns the ChildProjectIdentifier field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetChildProjectIdentifier() string {
	if o == nil || IsNil(o.ChildProjectIdentifier) {
		var ret string
		return ret
	}
	return *o.ChildProjectIdentifier
}

// GetChildProjectIdentifierOk returns a tuple with the ChildProjectIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetChildProjectIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ChildProjectIdentifier) {
		return nil, false
	}
	return o.ChildProjectIdentifier, true
}

// HasChildProjectIdentifier returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasChildProjectIdentifier() bool {
	if o != nil && !IsNil(o.ChildProjectIdentifier) {
		return true
	}

	return false
}

// SetChildProjectIdentifier gets a reference to the given string and assigns it to the ChildProjectIdentifier field.
func (o *ModelAPIPatchTriggerDefinition) SetChildProjectIdentifier(v string) {
	o.ChildProjectIdentifier = &v
}

// GetDownstreamRevision returns the DownstreamRevision field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetDownstreamRevision() string {
	if o == nil || IsNil(o.DownstreamRevision) {
		var ret string
		return ret
	}
	return *o.DownstreamRevision
}

// GetDownstreamRevisionOk returns a tuple with the DownstreamRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetDownstreamRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.DownstreamRevision) {
		return nil, false
	}
	return o.DownstreamRevision, true
}

// HasDownstreamRevision returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasDownstreamRevision() bool {
	if o != nil && !IsNil(o.DownstreamRevision) {
		return true
	}

	return false
}

// SetDownstreamRevision gets a reference to the given string and assigns it to the DownstreamRevision field.
func (o *ModelAPIPatchTriggerDefinition) SetDownstreamRevision(v string) {
	o.DownstreamRevision = &v
}

// GetParentAsModule returns the ParentAsModule field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetParentAsModule() string {
	if o == nil || IsNil(o.ParentAsModule) {
		var ret string
		return ret
	}
	return *o.ParentAsModule
}

// GetParentAsModuleOk returns a tuple with the ParentAsModule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetParentAsModuleOk() (*string, bool) {
	if o == nil || IsNil(o.ParentAsModule) {
		return nil, false
	}
	return o.ParentAsModule, true
}

// HasParentAsModule returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasParentAsModule() bool {
	if o != nil && !IsNil(o.ParentAsModule) {
		return true
	}

	return false
}

// SetParentAsModule gets a reference to the given string and assigns it to the ParentAsModule field.
func (o *ModelAPIPatchTriggerDefinition) SetParentAsModule(v string) {
	o.ParentAsModule = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelAPIPatchTriggerDefinition) SetStatus(v string) {
	o.Status = &v
}

// GetTaskSpecifiers returns the TaskSpecifiers field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetTaskSpecifiers() []ModelAPITaskSpecifier {
	if o == nil || IsNil(o.TaskSpecifiers) {
		var ret []ModelAPITaskSpecifier
		return ret
	}
	return o.TaskSpecifiers
}

// GetTaskSpecifiersOk returns a tuple with the TaskSpecifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetTaskSpecifiersOk() ([]ModelAPITaskSpecifier, bool) {
	if o == nil || IsNil(o.TaskSpecifiers) {
		return nil, false
	}
	return o.TaskSpecifiers, true
}

// HasTaskSpecifiers returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasTaskSpecifiers() bool {
	if o != nil && !IsNil(o.TaskSpecifiers) {
		return true
	}

	return false
}

// SetTaskSpecifiers gets a reference to the given []ModelAPITaskSpecifier and assigns it to the TaskSpecifiers field.
func (o *ModelAPIPatchTriggerDefinition) SetTaskSpecifiers(v []ModelAPITaskSpecifier) {
	o.TaskSpecifiers = v
}

// GetVariantsTasks returns the VariantsTasks field value if set, zero value otherwise.
func (o *ModelAPIPatchTriggerDefinition) GetVariantsTasks() []ModelVariantTask {
	if o == nil || IsNil(o.VariantsTasks) {
		var ret []ModelVariantTask
		return ret
	}
	return o.VariantsTasks
}

// GetVariantsTasksOk returns a tuple with the VariantsTasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelAPIPatchTriggerDefinition) GetVariantsTasksOk() ([]ModelVariantTask, bool) {
	if o == nil || IsNil(o.VariantsTasks) {
		return nil, false
	}
	return o.VariantsTasks, true
}

// HasVariantsTasks returns a boolean if a field has been set.
func (o *ModelAPIPatchTriggerDefinition) HasVariantsTasks() bool {
	if o != nil && !IsNil(o.VariantsTasks) {
		return true
	}

	return false
}

// SetVariantsTasks gets a reference to the given []ModelVariantTask and assigns it to the VariantsTasks field.
func (o *ModelAPIPatchTriggerDefinition) SetVariantsTasks(v []ModelVariantTask) {
	o.VariantsTasks = v
}

func (o ModelAPIPatchTriggerDefinition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIPatchTriggerDefinition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.ChildProjectId) {
		toSerialize["child_project_id"] = o.ChildProjectId
	}
	if !IsNil(o.ChildProjectIdentifier) {
		toSerialize["child_project_identifier"] = o.ChildProjectIdentifier
	}
	if !IsNil(o.DownstreamRevision) {
		toSerialize["downstream_revision"] = o.DownstreamRevision
	}
	if !IsNil(o.ParentAsModule) {
		toSerialize["parent_as_module"] = o.ParentAsModule
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TaskSpecifiers) {
		toSerialize["task_specifiers"] = o.TaskSpecifiers
	}
	if !IsNil(o.VariantsTasks) {
		toSerialize["variants_tasks"] = o.VariantsTasks
	}
	return toSerialize, nil
}

type NullableModelAPIPatchTriggerDefinition struct {
	value *ModelAPIPatchTriggerDefinition
	isSet bool
}

func (v NullableModelAPIPatchTriggerDefinition) Get() *ModelAPIPatchTriggerDefinition {
	return v.value
}

func (v *NullableModelAPIPatchTriggerDefinition) Set(val *ModelAPIPatchTriggerDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIPatchTriggerDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIPatchTriggerDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIPatchTriggerDefinition(val *ModelAPIPatchTriggerDefinition) *NullableModelAPIPatchTriggerDefinition {
	return &NullableModelAPIPatchTriggerDefinition{value: val, isSet: true}
}

func (v NullableModelAPIPatchTriggerDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIPatchTriggerDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


