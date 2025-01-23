# ModelAPIPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** |  | [optional] 
**Levels** | Pointer to [**[]EvergreenPermissionLevel**](EvergreenPermissionLevel.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIPermission

`func NewModelAPIPermission() *ModelAPIPermission`

NewModelAPIPermission instantiates a new ModelAPIPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIPermissionWithDefaults

`func NewModelAPIPermissionWithDefaults() *ModelAPIPermission`

NewModelAPIPermissionWithDefaults instantiates a new ModelAPIPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ModelAPIPermission) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelAPIPermission) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelAPIPermission) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelAPIPermission) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLevels

`func (o *ModelAPIPermission) GetLevels() []EvergreenPermissionLevel`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *ModelAPIPermission) GetLevelsOk() (*[]EvergreenPermissionLevel, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *ModelAPIPermission) SetLevels(v []EvergreenPermissionLevel)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *ModelAPIPermission) HasLevels() bool`

HasLevels returns a boolean if a field has been set.

### GetName

`func (o *ModelAPIPermission) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelAPIPermission) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelAPIPermission) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelAPIPermission) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


