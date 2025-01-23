# ManifestManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | The branch of the repository. | [optional] 
**Id** | Pointer to **string** | Identifier for the version. | [optional] 
**IsBase** | Pointer to **bool** | True if the version is a mainline build. | [optional] 
**ModuleOverrides** | Pointer to **map[string]string** |  | [optional] 
**Modules** | Pointer to [**map[string]ManifestModule**](ManifestModule.md) | Map from the GitHub repository name to the module&#39;s information. | [optional] 
**Project** | Pointer to **string** | The project identifier for the version. | [optional] 
**Revision** | Pointer to **string** | The revision of the version. | [optional] 

## Methods

### NewManifestManifest

`func NewManifestManifest() *ManifestManifest`

NewManifestManifest instantiates a new ManifestManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestManifestWithDefaults

`func NewManifestManifestWithDefaults() *ManifestManifest`

NewManifestManifestWithDefaults instantiates a new ManifestManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ManifestManifest) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ManifestManifest) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ManifestManifest) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ManifestManifest) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetId

`func (o *ManifestManifest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManifestManifest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManifestManifest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ManifestManifest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsBase

`func (o *ManifestManifest) GetIsBase() bool`

GetIsBase returns the IsBase field if non-nil, zero value otherwise.

### GetIsBaseOk

`func (o *ManifestManifest) GetIsBaseOk() (*bool, bool)`

GetIsBaseOk returns a tuple with the IsBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBase

`func (o *ManifestManifest) SetIsBase(v bool)`

SetIsBase sets IsBase field to given value.

### HasIsBase

`func (o *ManifestManifest) HasIsBase() bool`

HasIsBase returns a boolean if a field has been set.

### GetModuleOverrides

`func (o *ManifestManifest) GetModuleOverrides() map[string]string`

GetModuleOverrides returns the ModuleOverrides field if non-nil, zero value otherwise.

### GetModuleOverridesOk

`func (o *ManifestManifest) GetModuleOverridesOk() (*map[string]string, bool)`

GetModuleOverridesOk returns a tuple with the ModuleOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModuleOverrides

`func (o *ManifestManifest) SetModuleOverrides(v map[string]string)`

SetModuleOverrides sets ModuleOverrides field to given value.

### HasModuleOverrides

`func (o *ManifestManifest) HasModuleOverrides() bool`

HasModuleOverrides returns a boolean if a field has been set.

### GetModules

`func (o *ManifestManifest) GetModules() map[string]ManifestModule`

GetModules returns the Modules field if non-nil, zero value otherwise.

### GetModulesOk

`func (o *ManifestManifest) GetModulesOk() (*map[string]ManifestModule, bool)`

GetModulesOk returns a tuple with the Modules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModules

`func (o *ManifestManifest) SetModules(v map[string]ManifestModule)`

SetModules sets Modules field to given value.

### HasModules

`func (o *ManifestManifest) HasModules() bool`

HasModules returns a boolean if a field has been set.

### GetProject

`func (o *ManifestManifest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ManifestManifest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ManifestManifest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *ManifestManifest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRevision

`func (o *ManifestManifest) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ManifestManifest) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ManifestManifest) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ManifestManifest) HasRevision() bool`

HasRevision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


