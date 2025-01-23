# ModelAPIContainerSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | External identifier for the container secret. Cannot be modified by users. | [optional] 
**ExternalName** | Pointer to **string** | External name of the container secrets. Cannot be modified by users. | [optional] 
**Name** | Pointer to **string** | Name of the container secret. | [optional] 
**RepoCreds** | Pointer to [**ModelAPIRepositoryCredentials**](ModelAPIRepositoryCredentials.md) | RepoCreds, if set, are the new repository credentials to store. This only applies to repository credentials. | [optional] 
**ShouldRotate** | Pointer to **bool** | ShouldRotate indicates that the user requested the pod secret to be rotated to a new value. This only applies to the project&#39;s pod secret. | [optional] 
**Type** | Pointer to **string** | Type of container secret. | [optional] 
**Value** | Pointer to **string** | Container secret value to set. | [optional] 

## Methods

### NewModelAPIContainerSecret

`func NewModelAPIContainerSecret() *ModelAPIContainerSecret`

NewModelAPIContainerSecret instantiates a new ModelAPIContainerSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIContainerSecretWithDefaults

`func NewModelAPIContainerSecretWithDefaults() *ModelAPIContainerSecret`

NewModelAPIContainerSecretWithDefaults instantiates a new ModelAPIContainerSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ModelAPIContainerSecret) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ModelAPIContainerSecret) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ModelAPIContainerSecret) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ModelAPIContainerSecret) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetExternalName

`func (o *ModelAPIContainerSecret) GetExternalName() string`

GetExternalName returns the ExternalName field if non-nil, zero value otherwise.

### GetExternalNameOk

`func (o *ModelAPIContainerSecret) GetExternalNameOk() (*string, bool)`

GetExternalNameOk returns a tuple with the ExternalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalName

`func (o *ModelAPIContainerSecret) SetExternalName(v string)`

SetExternalName sets ExternalName field to given value.

### HasExternalName

`func (o *ModelAPIContainerSecret) HasExternalName() bool`

HasExternalName returns a boolean if a field has been set.

### GetName

`func (o *ModelAPIContainerSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelAPIContainerSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelAPIContainerSecret) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelAPIContainerSecret) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRepoCreds

`func (o *ModelAPIContainerSecret) GetRepoCreds() ModelAPIRepositoryCredentials`

GetRepoCreds returns the RepoCreds field if non-nil, zero value otherwise.

### GetRepoCredsOk

`func (o *ModelAPIContainerSecret) GetRepoCredsOk() (*ModelAPIRepositoryCredentials, bool)`

GetRepoCredsOk returns a tuple with the RepoCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoCreds

`func (o *ModelAPIContainerSecret) SetRepoCreds(v ModelAPIRepositoryCredentials)`

SetRepoCreds sets RepoCreds field to given value.

### HasRepoCreds

`func (o *ModelAPIContainerSecret) HasRepoCreds() bool`

HasRepoCreds returns a boolean if a field has been set.

### GetShouldRotate

`func (o *ModelAPIContainerSecret) GetShouldRotate() bool`

GetShouldRotate returns the ShouldRotate field if non-nil, zero value otherwise.

### GetShouldRotateOk

`func (o *ModelAPIContainerSecret) GetShouldRotateOk() (*bool, bool)`

GetShouldRotateOk returns a tuple with the ShouldRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldRotate

`func (o *ModelAPIContainerSecret) SetShouldRotate(v bool)`

SetShouldRotate sets ShouldRotate field to given value.

### HasShouldRotate

`func (o *ModelAPIContainerSecret) HasShouldRotate() bool`

HasShouldRotate returns a boolean if a field has been set.

### GetType

`func (o *ModelAPIContainerSecret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelAPIContainerSecret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelAPIContainerSecret) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelAPIContainerSecret) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ModelAPIContainerSecret) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ModelAPIContainerSecret) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ModelAPIContainerSecret) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ModelAPIContainerSecret) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


