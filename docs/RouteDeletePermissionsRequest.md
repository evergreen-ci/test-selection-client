# RouteDeletePermissionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | resource_id - the resource ID for which to delete permissions.   Required unless deleting all permissions. | [optional] 
**ResourceType** | Pointer to **string** | resource_type - the type of resources for which to delete permissions. Must   be one of \&quot;project\&quot;, \&quot;distro\&quot;, \&quot;superuser\&quot;, or \&quot;all\&quot;. \&quot;all\&quot; will revoke all   permissions for the user. | [optional] 

## Methods

### NewRouteDeletePermissionsRequest

`func NewRouteDeletePermissionsRequest() *RouteDeletePermissionsRequest`

NewRouteDeletePermissionsRequest instantiates a new RouteDeletePermissionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteDeletePermissionsRequestWithDefaults

`func NewRouteDeletePermissionsRequestWithDefaults() *RouteDeletePermissionsRequest`

NewRouteDeletePermissionsRequestWithDefaults instantiates a new RouteDeletePermissionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *RouteDeletePermissionsRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *RouteDeletePermissionsRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *RouteDeletePermissionsRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *RouteDeletePermissionsRequest) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetResourceType

`func (o *RouteDeletePermissionsRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *RouteDeletePermissionsRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *RouteDeletePermissionsRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *RouteDeletePermissionsRequest) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


