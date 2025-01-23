# ModelAPISpawnHostModify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**AddHours** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **string** |  | [optional] 
**HostId** | Pointer to **string** |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**NewName** | Pointer to **string** |  | [optional] 
**RdpPwd** | Pointer to **string** |  | [optional] 
**TagsToAdd** | Pointer to **[]string** |  | [optional] 
**TagsToDelete** | Pointer to **[]string** |  | [optional] 
**VolumeId** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPISpawnHostModify

`func NewModelAPISpawnHostModify() *ModelAPISpawnHostModify`

NewModelAPISpawnHostModify instantiates a new ModelAPISpawnHostModify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPISpawnHostModifyWithDefaults

`func NewModelAPISpawnHostModifyWithDefaults() *ModelAPISpawnHostModify`

NewModelAPISpawnHostModifyWithDefaults instantiates a new ModelAPISpawnHostModify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ModelAPISpawnHostModify) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ModelAPISpawnHostModify) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ModelAPISpawnHostModify) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ModelAPISpawnHostModify) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetAddHours

`func (o *ModelAPISpawnHostModify) GetAddHours() string`

GetAddHours returns the AddHours field if non-nil, zero value otherwise.

### GetAddHoursOk

`func (o *ModelAPISpawnHostModify) GetAddHoursOk() (*string, bool)`

GetAddHoursOk returns a tuple with the AddHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddHours

`func (o *ModelAPISpawnHostModify) SetAddHours(v string)`

SetAddHours sets AddHours field to given value.

### HasAddHours

`func (o *ModelAPISpawnHostModify) HasAddHours() bool`

HasAddHours returns a boolean if a field has been set.

### GetExpiration

`func (o *ModelAPISpawnHostModify) GetExpiration() string`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *ModelAPISpawnHostModify) GetExpirationOk() (*string, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *ModelAPISpawnHostModify) SetExpiration(v string)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *ModelAPISpawnHostModify) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetHostId

`func (o *ModelAPISpawnHostModify) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ModelAPISpawnHostModify) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ModelAPISpawnHostModify) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ModelAPISpawnHostModify) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetInstanceType

`func (o *ModelAPISpawnHostModify) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ModelAPISpawnHostModify) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ModelAPISpawnHostModify) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ModelAPISpawnHostModify) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetNewName

`func (o *ModelAPISpawnHostModify) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *ModelAPISpawnHostModify) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *ModelAPISpawnHostModify) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *ModelAPISpawnHostModify) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRdpPwd

`func (o *ModelAPISpawnHostModify) GetRdpPwd() string`

GetRdpPwd returns the RdpPwd field if non-nil, zero value otherwise.

### GetRdpPwdOk

`func (o *ModelAPISpawnHostModify) GetRdpPwdOk() (*string, bool)`

GetRdpPwdOk returns a tuple with the RdpPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpPwd

`func (o *ModelAPISpawnHostModify) SetRdpPwd(v string)`

SetRdpPwd sets RdpPwd field to given value.

### HasRdpPwd

`func (o *ModelAPISpawnHostModify) HasRdpPwd() bool`

HasRdpPwd returns a boolean if a field has been set.

### GetTagsToAdd

`func (o *ModelAPISpawnHostModify) GetTagsToAdd() []string`

GetTagsToAdd returns the TagsToAdd field if non-nil, zero value otherwise.

### GetTagsToAddOk

`func (o *ModelAPISpawnHostModify) GetTagsToAddOk() (*[]string, bool)`

GetTagsToAddOk returns a tuple with the TagsToAdd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsToAdd

`func (o *ModelAPISpawnHostModify) SetTagsToAdd(v []string)`

SetTagsToAdd sets TagsToAdd field to given value.

### HasTagsToAdd

`func (o *ModelAPISpawnHostModify) HasTagsToAdd() bool`

HasTagsToAdd returns a boolean if a field has been set.

### GetTagsToDelete

`func (o *ModelAPISpawnHostModify) GetTagsToDelete() []string`

GetTagsToDelete returns the TagsToDelete field if non-nil, zero value otherwise.

### GetTagsToDeleteOk

`func (o *ModelAPISpawnHostModify) GetTagsToDeleteOk() (*[]string, bool)`

GetTagsToDeleteOk returns a tuple with the TagsToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagsToDelete

`func (o *ModelAPISpawnHostModify) SetTagsToDelete(v []string)`

SetTagsToDelete sets TagsToDelete field to given value.

### HasTagsToDelete

`func (o *ModelAPISpawnHostModify) HasTagsToDelete() bool`

HasTagsToDelete returns a boolean if a field has been set.

### GetVolumeId

`func (o *ModelAPISpawnHostModify) GetVolumeId() string`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *ModelAPISpawnHostModify) GetVolumeIdOk() (*string, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *ModelAPISpawnHostModify) SetVolumeId(v string)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *ModelAPISpawnHostModify) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


