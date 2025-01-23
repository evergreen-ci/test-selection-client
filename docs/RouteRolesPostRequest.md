# RouteRolesPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateUser** | Pointer to **bool** | if true, will also create a shell user document for the user. By default, specifying a user that does not exist will error | [optional] 
**RemoveRoles** | Pointer to **[]string** | The list of roles to remove for the user | [optional] 
**Roles** | Pointer to **[]string** | the list of roles to add for the user | [optional] 

## Methods

### NewRouteRolesPostRequest

`func NewRouteRolesPostRequest() *RouteRolesPostRequest`

NewRouteRolesPostRequest instantiates a new RouteRolesPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteRolesPostRequestWithDefaults

`func NewRouteRolesPostRequestWithDefaults() *RouteRolesPostRequest`

NewRouteRolesPostRequestWithDefaults instantiates a new RouteRolesPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateUser

`func (o *RouteRolesPostRequest) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *RouteRolesPostRequest) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *RouteRolesPostRequest) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *RouteRolesPostRequest) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetRemoveRoles

`func (o *RouteRolesPostRequest) GetRemoveRoles() []string`

GetRemoveRoles returns the RemoveRoles field if non-nil, zero value otherwise.

### GetRemoveRolesOk

`func (o *RouteRolesPostRequest) GetRemoveRolesOk() (*[]string, bool)`

GetRemoveRolesOk returns a tuple with the RemoveRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveRoles

`func (o *RouteRolesPostRequest) SetRemoveRoles(v []string)`

SetRemoveRoles sets RemoveRoles field to given value.

### HasRemoveRoles

`func (o *RouteRolesPostRequest) HasRemoveRoles() bool`

HasRemoveRoles returns a boolean if a field has been set.

### GetRoles

`func (o *RouteRolesPostRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *RouteRolesPostRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *RouteRolesPostRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *RouteRolesPostRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


