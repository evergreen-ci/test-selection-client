# RouteRenameUserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The old email of the user | 
**NewEmail** | **string** | The new email of the user | 

## Methods

### NewRouteRenameUserInfo

`func NewRouteRenameUserInfo(email string, newEmail string, ) *RouteRenameUserInfo`

NewRouteRenameUserInfo instantiates a new RouteRenameUserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteRenameUserInfoWithDefaults

`func NewRouteRenameUserInfoWithDefaults() *RouteRenameUserInfo`

NewRouteRenameUserInfoWithDefaults instantiates a new RouteRenameUserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *RouteRenameUserInfo) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *RouteRenameUserInfo) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *RouteRenameUserInfo) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetNewEmail

`func (o *RouteRenameUserInfo) GetNewEmail() string`

GetNewEmail returns the NewEmail field if non-nil, zero value otherwise.

### GetNewEmailOk

`func (o *RouteRenameUserInfo) GetNewEmailOk() (*string, bool)`

GetNewEmailOk returns a tuple with the NewEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewEmail

`func (o *RouteRenameUserInfo) SetNewEmail(v string)`

SetNewEmail sets NewEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


