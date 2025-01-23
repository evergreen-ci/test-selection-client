# ModelAPIDBUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BetaFeatures** | Pointer to [**ModelAPIBetaFeatures**](ModelAPIBetaFeatures.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EmailAddress** | Pointer to **string** |  | [optional] 
**OnlyApi** | Pointer to **bool** | will be set to true if the user represents a service user | [optional] 
**ParsleyFilters** | Pointer to [**[]ModelAPIParsleyFilter**](ModelAPIParsleyFilter.md) |  | [optional] 
**ParsleySettings** | Pointer to [**ModelAPIParsleySettings**](ModelAPIParsleySettings.md) |  | [optional] 
**Roles** | Pointer to **[]string** |  | [optional] 
**Settings** | Pointer to [**ModelAPIUserSettings**](ModelAPIUserSettings.md) |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIDBUser

`func NewModelAPIDBUser() *ModelAPIDBUser`

NewModelAPIDBUser instantiates a new ModelAPIDBUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIDBUserWithDefaults

`func NewModelAPIDBUserWithDefaults() *ModelAPIDBUser`

NewModelAPIDBUserWithDefaults instantiates a new ModelAPIDBUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBetaFeatures

`func (o *ModelAPIDBUser) GetBetaFeatures() ModelAPIBetaFeatures`

GetBetaFeatures returns the BetaFeatures field if non-nil, zero value otherwise.

### GetBetaFeaturesOk

`func (o *ModelAPIDBUser) GetBetaFeaturesOk() (*ModelAPIBetaFeatures, bool)`

GetBetaFeaturesOk returns a tuple with the BetaFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaFeatures

`func (o *ModelAPIDBUser) SetBetaFeatures(v ModelAPIBetaFeatures)`

SetBetaFeatures sets BetaFeatures field to given value.

### HasBetaFeatures

`func (o *ModelAPIDBUser) HasBetaFeatures() bool`

HasBetaFeatures returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelAPIDBUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPIDBUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPIDBUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPIDBUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *ModelAPIDBUser) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ModelAPIDBUser) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ModelAPIDBUser) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ModelAPIDBUser) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetOnlyApi

`func (o *ModelAPIDBUser) GetOnlyApi() bool`

GetOnlyApi returns the OnlyApi field if non-nil, zero value otherwise.

### GetOnlyApiOk

`func (o *ModelAPIDBUser) GetOnlyApiOk() (*bool, bool)`

GetOnlyApiOk returns a tuple with the OnlyApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyApi

`func (o *ModelAPIDBUser) SetOnlyApi(v bool)`

SetOnlyApi sets OnlyApi field to given value.

### HasOnlyApi

`func (o *ModelAPIDBUser) HasOnlyApi() bool`

HasOnlyApi returns a boolean if a field has been set.

### GetParsleyFilters

`func (o *ModelAPIDBUser) GetParsleyFilters() []ModelAPIParsleyFilter`

GetParsleyFilters returns the ParsleyFilters field if non-nil, zero value otherwise.

### GetParsleyFiltersOk

`func (o *ModelAPIDBUser) GetParsleyFiltersOk() (*[]ModelAPIParsleyFilter, bool)`

GetParsleyFiltersOk returns a tuple with the ParsleyFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsleyFilters

`func (o *ModelAPIDBUser) SetParsleyFilters(v []ModelAPIParsleyFilter)`

SetParsleyFilters sets ParsleyFilters field to given value.

### HasParsleyFilters

`func (o *ModelAPIDBUser) HasParsleyFilters() bool`

HasParsleyFilters returns a boolean if a field has been set.

### GetParsleySettings

`func (o *ModelAPIDBUser) GetParsleySettings() ModelAPIParsleySettings`

GetParsleySettings returns the ParsleySettings field if non-nil, zero value otherwise.

### GetParsleySettingsOk

`func (o *ModelAPIDBUser) GetParsleySettingsOk() (*ModelAPIParsleySettings, bool)`

GetParsleySettingsOk returns a tuple with the ParsleySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsleySettings

`func (o *ModelAPIDBUser) SetParsleySettings(v ModelAPIParsleySettings)`

SetParsleySettings sets ParsleySettings field to given value.

### HasParsleySettings

`func (o *ModelAPIDBUser) HasParsleySettings() bool`

HasParsleySettings returns a boolean if a field has been set.

### GetRoles

`func (o *ModelAPIDBUser) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ModelAPIDBUser) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ModelAPIDBUser) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ModelAPIDBUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetSettings

`func (o *ModelAPIDBUser) GetSettings() ModelAPIUserSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ModelAPIDBUser) GetSettingsOk() (*ModelAPIUserSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ModelAPIDBUser) SetSettings(v ModelAPIUserSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *ModelAPIDBUser) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetUserId

`func (o *ModelAPIDBUser) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ModelAPIDBUser) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ModelAPIDBUser) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ModelAPIDBUser) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


