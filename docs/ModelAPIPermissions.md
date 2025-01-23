# ModelAPIPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistroPermissions** | Pointer to [**[]ModelAPIPermission**](ModelAPIPermission.md) |  | [optional] 
**ProjectPermissions** | Pointer to [**[]ModelAPIPermission**](ModelAPIPermission.md) |  | [optional] 

## Methods

### NewModelAPIPermissions

`func NewModelAPIPermissions() *ModelAPIPermissions`

NewModelAPIPermissions instantiates a new ModelAPIPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIPermissionsWithDefaults

`func NewModelAPIPermissionsWithDefaults() *ModelAPIPermissions`

NewModelAPIPermissionsWithDefaults instantiates a new ModelAPIPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistroPermissions

`func (o *ModelAPIPermissions) GetDistroPermissions() []ModelAPIPermission`

GetDistroPermissions returns the DistroPermissions field if non-nil, zero value otherwise.

### GetDistroPermissionsOk

`func (o *ModelAPIPermissions) GetDistroPermissionsOk() (*[]ModelAPIPermission, bool)`

GetDistroPermissionsOk returns a tuple with the DistroPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroPermissions

`func (o *ModelAPIPermissions) SetDistroPermissions(v []ModelAPIPermission)`

SetDistroPermissions sets DistroPermissions field to given value.

### HasDistroPermissions

`func (o *ModelAPIPermissions) HasDistroPermissions() bool`

HasDistroPermissions returns a boolean if a field has been set.

### GetProjectPermissions

`func (o *ModelAPIPermissions) GetProjectPermissions() []ModelAPIPermission`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *ModelAPIPermissions) GetProjectPermissionsOk() (*[]ModelAPIPermission, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *ModelAPIPermissions) SetProjectPermissions(v []ModelAPIPermission)`

SetProjectPermissions sets ProjectPermissions field to given value.

### HasProjectPermissions

`func (o *ModelAPIPermissions) HasProjectPermissions() bool`

HasProjectPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


