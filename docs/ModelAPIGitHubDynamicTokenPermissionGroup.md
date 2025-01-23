# ModelAPIGitHubDynamicTokenPermissionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllPermissions** | Pointer to **bool** | AllPermissions is a flag that indicates that the group has all permissions. If this is set to true, the Permissions field is ignored. If this is set to false, the Permissions field is used (and may be nil, representing no permissions). | [optional] 
**Name** | Pointer to **string** | Name of the GitHub permission group. | [optional] 
**Permissions** | Pointer to **map[string]string** | Permissions for the GitHub permission group. | [optional] 

## Methods

### NewModelAPIGitHubDynamicTokenPermissionGroup

`func NewModelAPIGitHubDynamicTokenPermissionGroup() *ModelAPIGitHubDynamicTokenPermissionGroup`

NewModelAPIGitHubDynamicTokenPermissionGroup instantiates a new ModelAPIGitHubDynamicTokenPermissionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIGitHubDynamicTokenPermissionGroupWithDefaults

`func NewModelAPIGitHubDynamicTokenPermissionGroupWithDefaults() *ModelAPIGitHubDynamicTokenPermissionGroup`

NewModelAPIGitHubDynamicTokenPermissionGroupWithDefaults instantiates a new ModelAPIGitHubDynamicTokenPermissionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllPermissions

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) GetAllPermissions() bool`

GetAllPermissions returns the AllPermissions field if non-nil, zero value otherwise.

### GetAllPermissionsOk

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) GetAllPermissionsOk() (*bool, bool)`

GetAllPermissionsOk returns a tuple with the AllPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPermissions

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) SetAllPermissions(v bool)`

SetAllPermissions sets AllPermissions field to given value.

### HasAllPermissions

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) HasAllPermissions() bool`

HasAllPermissions returns a boolean if a field has been set.

### GetName

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPermissions

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) GetPermissions() map[string]string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) GetPermissionsOk() (*map[string]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) SetPermissions(v map[string]string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ModelAPIGitHubDynamicTokenPermissionGroup) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


