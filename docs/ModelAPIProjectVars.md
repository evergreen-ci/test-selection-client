# ModelAPIProjectVars

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminOnlyVars** | Pointer to **map[string]bool** | Admin-only variable names. | [optional] 
**PrivateVars** | Pointer to **map[string]bool** | Private variable names. | [optional] 
**Vars** | Pointer to **map[string]string** | Regular project variable names and their values. | [optional] 
**VarsToDelete** | Pointer to **[]string** | Names of project variables to delete. | [optional] 

## Methods

### NewModelAPIProjectVars

`func NewModelAPIProjectVars() *ModelAPIProjectVars`

NewModelAPIProjectVars instantiates a new ModelAPIProjectVars object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIProjectVarsWithDefaults

`func NewModelAPIProjectVarsWithDefaults() *ModelAPIProjectVars`

NewModelAPIProjectVarsWithDefaults instantiates a new ModelAPIProjectVars object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminOnlyVars

`func (o *ModelAPIProjectVars) GetAdminOnlyVars() map[string]bool`

GetAdminOnlyVars returns the AdminOnlyVars field if non-nil, zero value otherwise.

### GetAdminOnlyVarsOk

`func (o *ModelAPIProjectVars) GetAdminOnlyVarsOk() (*map[string]bool, bool)`

GetAdminOnlyVarsOk returns a tuple with the AdminOnlyVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOnlyVars

`func (o *ModelAPIProjectVars) SetAdminOnlyVars(v map[string]bool)`

SetAdminOnlyVars sets AdminOnlyVars field to given value.

### HasAdminOnlyVars

`func (o *ModelAPIProjectVars) HasAdminOnlyVars() bool`

HasAdminOnlyVars returns a boolean if a field has been set.

### GetPrivateVars

`func (o *ModelAPIProjectVars) GetPrivateVars() map[string]bool`

GetPrivateVars returns the PrivateVars field if non-nil, zero value otherwise.

### GetPrivateVarsOk

`func (o *ModelAPIProjectVars) GetPrivateVarsOk() (*map[string]bool, bool)`

GetPrivateVarsOk returns a tuple with the PrivateVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateVars

`func (o *ModelAPIProjectVars) SetPrivateVars(v map[string]bool)`

SetPrivateVars sets PrivateVars field to given value.

### HasPrivateVars

`func (o *ModelAPIProjectVars) HasPrivateVars() bool`

HasPrivateVars returns a boolean if a field has been set.

### GetVars

`func (o *ModelAPIProjectVars) GetVars() map[string]string`

GetVars returns the Vars field if non-nil, zero value otherwise.

### GetVarsOk

`func (o *ModelAPIProjectVars) GetVarsOk() (*map[string]string, bool)`

GetVarsOk returns a tuple with the Vars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVars

`func (o *ModelAPIProjectVars) SetVars(v map[string]string)`

SetVars sets Vars field to given value.

### HasVars

`func (o *ModelAPIProjectVars) HasVars() bool`

HasVars returns a boolean if a field has been set.

### GetVarsToDelete

`func (o *ModelAPIProjectVars) GetVarsToDelete() []string`

GetVarsToDelete returns the VarsToDelete field if non-nil, zero value otherwise.

### GetVarsToDeleteOk

`func (o *ModelAPIProjectVars) GetVarsToDeleteOk() (*[]string, bool)`

GetVarsToDeleteOk returns a tuple with the VarsToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarsToDelete

`func (o *ModelAPIProjectVars) SetVarsToDelete(v []string)`

SetVarsToDelete sets VarsToDelete field to given value.

### HasVarsToDelete

`func (o *ModelAPIProjectVars) HasVarsToDelete() bool`

HasVarsToDelete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


