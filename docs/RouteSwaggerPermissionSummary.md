# RouteSwaggerPermissionSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **map[string]map[string]int32** | permissions - an object whose keys are the resources for which the user has   permissions. Note that these objects will often have many keys, since   logged-in users have basic permissions to every project and distro. The   values in the keys are objects representing the permissions that the user   has for that resource, identical to the format of the permissions field in   the POST /users/\\&lt;user_id\\&gt;/permissions API. | [optional] 
**Type** | Pointer to **string** | type - the type of resources for which the listed permissions apply.   Will be \&quot;project\&quot;, \&quot;distro\&quot;, or \&quot;superuser\&quot; | [optional] 

## Methods

### NewRouteSwaggerPermissionSummary

`func NewRouteSwaggerPermissionSummary() *RouteSwaggerPermissionSummary`

NewRouteSwaggerPermissionSummary instantiates a new RouteSwaggerPermissionSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteSwaggerPermissionSummaryWithDefaults

`func NewRouteSwaggerPermissionSummaryWithDefaults() *RouteSwaggerPermissionSummary`

NewRouteSwaggerPermissionSummaryWithDefaults instantiates a new RouteSwaggerPermissionSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *RouteSwaggerPermissionSummary) GetPermissions() map[string]map[string]int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *RouteSwaggerPermissionSummary) GetPermissionsOk() (*map[string]map[string]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *RouteSwaggerPermissionSummary) SetPermissions(v map[string]map[string]int32)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *RouteSwaggerPermissionSummary) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetType

`func (o *RouteSwaggerPermissionSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RouteSwaggerPermissionSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RouteSwaggerPermissionSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RouteSwaggerPermissionSummary) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


