# ModelDistroInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootstrapMethod** | Pointer to **string** |  | [optional] 
**DistroId** | Pointer to **string** | Unique Identifier of this distro. Can be used to fetch more informaiton about this distro | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**IsVirtualWorkstation** | Pointer to **bool** |  | [optional] 
**IsWindows** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** | The service which provides this type of machine | [optional] 
**User** | Pointer to **string** |  | [optional] 
**WorkDir** | Pointer to **string** |  | [optional] 

## Methods

### NewModelDistroInfo

`func NewModelDistroInfo() *ModelDistroInfo`

NewModelDistroInfo instantiates a new ModelDistroInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDistroInfoWithDefaults

`func NewModelDistroInfoWithDefaults() *ModelDistroInfo`

NewModelDistroInfoWithDefaults instantiates a new ModelDistroInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootstrapMethod

`func (o *ModelDistroInfo) GetBootstrapMethod() string`

GetBootstrapMethod returns the BootstrapMethod field if non-nil, zero value otherwise.

### GetBootstrapMethodOk

`func (o *ModelDistroInfo) GetBootstrapMethodOk() (*string, bool)`

GetBootstrapMethodOk returns a tuple with the BootstrapMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapMethod

`func (o *ModelDistroInfo) SetBootstrapMethod(v string)`

SetBootstrapMethod sets BootstrapMethod field to given value.

### HasBootstrapMethod

`func (o *ModelDistroInfo) HasBootstrapMethod() bool`

HasBootstrapMethod returns a boolean if a field has been set.

### GetDistroId

`func (o *ModelDistroInfo) GetDistroId() string`

GetDistroId returns the DistroId field if non-nil, zero value otherwise.

### GetDistroIdOk

`func (o *ModelDistroInfo) GetDistroIdOk() (*string, bool)`

GetDistroIdOk returns a tuple with the DistroId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistroId

`func (o *ModelDistroInfo) SetDistroId(v string)`

SetDistroId sets DistroId field to given value.

### HasDistroId

`func (o *ModelDistroInfo) HasDistroId() bool`

HasDistroId returns a boolean if a field has been set.

### GetImageId

`func (o *ModelDistroInfo) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ModelDistroInfo) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ModelDistroInfo) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ModelDistroInfo) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetIsVirtualWorkstation

`func (o *ModelDistroInfo) GetIsVirtualWorkstation() bool`

GetIsVirtualWorkstation returns the IsVirtualWorkstation field if non-nil, zero value otherwise.

### GetIsVirtualWorkstationOk

`func (o *ModelDistroInfo) GetIsVirtualWorkstationOk() (*bool, bool)`

GetIsVirtualWorkstationOk returns a tuple with the IsVirtualWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualWorkstation

`func (o *ModelDistroInfo) SetIsVirtualWorkstation(v bool)`

SetIsVirtualWorkstation sets IsVirtualWorkstation field to given value.

### HasIsVirtualWorkstation

`func (o *ModelDistroInfo) HasIsVirtualWorkstation() bool`

HasIsVirtualWorkstation returns a boolean if a field has been set.

### GetIsWindows

`func (o *ModelDistroInfo) GetIsWindows() bool`

GetIsWindows returns the IsWindows field if non-nil, zero value otherwise.

### GetIsWindowsOk

`func (o *ModelDistroInfo) GetIsWindowsOk() (*bool, bool)`

GetIsWindowsOk returns a tuple with the IsWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWindows

`func (o *ModelDistroInfo) SetIsWindows(v bool)`

SetIsWindows sets IsWindows field to given value.

### HasIsWindows

`func (o *ModelDistroInfo) HasIsWindows() bool`

HasIsWindows returns a boolean if a field has been set.

### GetProvider

`func (o *ModelDistroInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelDistroInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelDistroInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelDistroInfo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUser

`func (o *ModelDistroInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelDistroInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelDistroInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelDistroInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetWorkDir

`func (o *ModelDistroInfo) GetWorkDir() string`

GetWorkDir returns the WorkDir field if non-nil, zero value otherwise.

### GetWorkDirOk

`func (o *ModelDistroInfo) GetWorkDirOk() (*string, bool)`

GetWorkDirOk returns a tuple with the WorkDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkDir

`func (o *ModelDistroInfo) SetWorkDir(v string)`

SetWorkDir sets WorkDir field to given value.

### HasWorkDir

`func (o *ModelDistroInfo) HasWorkDir() bool`

HasWorkDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


