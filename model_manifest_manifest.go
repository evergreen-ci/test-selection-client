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

// checks if the ManifestManifest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManifestManifest{}

// ManifestManifest struct for ManifestManifest
type ManifestManifest struct {
	// The branch of the repository.
	Branch *string `json:"branch,omitempty"`
	// Identifier for the version.
	Id *string `json:"id,omitempty"`
	// True if the version is a mainline build.
	IsBase *bool `json:"is_base,omitempty"`
	ModuleOverrides *map[string]string `json:"module_overrides,omitempty"`
	// Map from the GitHub repository name to the module's information.
	Modules *map[string]ManifestModule `json:"modules,omitempty"`
	// The project identifier for the version.
	Project *string `json:"project,omitempty"`
	// The revision of the version.
	Revision *string `json:"revision,omitempty"`
}

// NewManifestManifest instantiates a new ManifestManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManifestManifest() *ManifestManifest {
	this := ManifestManifest{}
	return &this
}

// NewManifestManifestWithDefaults instantiates a new ManifestManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManifestManifestWithDefaults() *ManifestManifest {
	this := ManifestManifest{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *ManifestManifest) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *ManifestManifest) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *ManifestManifest) SetBranch(v string) {
	o.Branch = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ManifestManifest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ManifestManifest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ManifestManifest) SetId(v string) {
	o.Id = &v
}

// GetIsBase returns the IsBase field value if set, zero value otherwise.
func (o *ManifestManifest) GetIsBase() bool {
	if o == nil || IsNil(o.IsBase) {
		var ret bool
		return ret
	}
	return *o.IsBase
}

// GetIsBaseOk returns a tuple with the IsBase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetIsBaseOk() (*bool, bool) {
	if o == nil || IsNil(o.IsBase) {
		return nil, false
	}
	return o.IsBase, true
}

// HasIsBase returns a boolean if a field has been set.
func (o *ManifestManifest) HasIsBase() bool {
	if o != nil && !IsNil(o.IsBase) {
		return true
	}

	return false
}

// SetIsBase gets a reference to the given bool and assigns it to the IsBase field.
func (o *ManifestManifest) SetIsBase(v bool) {
	o.IsBase = &v
}

// GetModuleOverrides returns the ModuleOverrides field value if set, zero value otherwise.
func (o *ManifestManifest) GetModuleOverrides() map[string]string {
	if o == nil || IsNil(o.ModuleOverrides) {
		var ret map[string]string
		return ret
	}
	return *o.ModuleOverrides
}

// GetModuleOverridesOk returns a tuple with the ModuleOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetModuleOverridesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.ModuleOverrides) {
		return nil, false
	}
	return o.ModuleOverrides, true
}

// HasModuleOverrides returns a boolean if a field has been set.
func (o *ManifestManifest) HasModuleOverrides() bool {
	if o != nil && !IsNil(o.ModuleOverrides) {
		return true
	}

	return false
}

// SetModuleOverrides gets a reference to the given map[string]string and assigns it to the ModuleOverrides field.
func (o *ManifestManifest) SetModuleOverrides(v map[string]string) {
	o.ModuleOverrides = &v
}

// GetModules returns the Modules field value if set, zero value otherwise.
func (o *ManifestManifest) GetModules() map[string]ManifestModule {
	if o == nil || IsNil(o.Modules) {
		var ret map[string]ManifestModule
		return ret
	}
	return *o.Modules
}

// GetModulesOk returns a tuple with the Modules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetModulesOk() (*map[string]ManifestModule, bool) {
	if o == nil || IsNil(o.Modules) {
		return nil, false
	}
	return o.Modules, true
}

// HasModules returns a boolean if a field has been set.
func (o *ManifestManifest) HasModules() bool {
	if o != nil && !IsNil(o.Modules) {
		return true
	}

	return false
}

// SetModules gets a reference to the given map[string]ManifestModule and assigns it to the Modules field.
func (o *ManifestManifest) SetModules(v map[string]ManifestModule) {
	o.Modules = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ManifestManifest) GetProject() string {
	if o == nil || IsNil(o.Project) {
		var ret string
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetProjectOk() (*string, bool) {
	if o == nil || IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ManifestManifest) HasProject() bool {
	if o != nil && !IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given string and assigns it to the Project field.
func (o *ManifestManifest) SetProject(v string) {
	o.Project = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *ManifestManifest) GetRevision() string {
	if o == nil || IsNil(o.Revision) {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestManifest) GetRevisionOk() (*string, bool) {
	if o == nil || IsNil(o.Revision) {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *ManifestManifest) HasRevision() bool {
	if o != nil && !IsNil(o.Revision) {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *ManifestManifest) SetRevision(v string) {
	o.Revision = &v
}

func (o ManifestManifest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManifestManifest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsBase) {
		toSerialize["is_base"] = o.IsBase
	}
	if !IsNil(o.ModuleOverrides) {
		toSerialize["module_overrides"] = o.ModuleOverrides
	}
	if !IsNil(o.Modules) {
		toSerialize["modules"] = o.Modules
	}
	if !IsNil(o.Project) {
		toSerialize["project"] = o.Project
	}
	if !IsNil(o.Revision) {
		toSerialize["revision"] = o.Revision
	}
	return toSerialize, nil
}

type NullableManifestManifest struct {
	value *ManifestManifest
	isSet bool
}

func (v NullableManifestManifest) Get() *ManifestManifest {
	return v.value
}

func (v *NullableManifestManifest) Set(val *ManifestManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableManifestManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableManifestManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManifestManifest(val *ManifestManifest) *NullableManifestManifest {
	return &NullableManifestManifest{value: val, isSet: true}
}

func (v NullableManifestManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManifestManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


