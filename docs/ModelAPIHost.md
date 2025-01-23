# ModelAPIHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttachedVolumeIds** | Pointer to **[]string** |  | [optional] 
**CreationTime** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Distro** | Pointer to [**ModelDistroInfo**](ModelDistroInfo.md) | Object containing information about the distro type of this host | [optional] 
**ExpirationTime** | Pointer to **string** |  | [optional] 
**HomeVolumeId** | Pointer to **string** |  | [optional] 
**HostId** | Pointer to **string** | Unique identifier of a specific host | [optional] 
**HostType** | Pointer to **string** | The instance type requested for the provider, primarily used for ec2 dynamic hosts | [optional] 
**HostUrl** | Pointer to **string** |  | [optional] 
**InstanceTags** | Pointer to [**[]HostTag**](HostTag.md) |  | [optional] 
**InstanceType** | Pointer to **string** |  | [optional] 
**LastCommunication** | Pointer to **string** |  | [optional] 
**NeedsReprovision** | Pointer to **string** |  | [optional] 
**NoExpiration** | Pointer to **bool** |  | [optional] 
**PersistentDnsName** | Pointer to **string** |  | [optional] 
**ProvisionOptions** | Pointer to [**ModelAPIProvisionOptions**](ModelAPIProvisionOptions.md) | Contains options for spawn hosts. | [optional] 
**Provisioned** | Pointer to **bool** |  | [optional] 
**RunningTask** | Pointer to [**ModelTaskInfo**](ModelTaskInfo.md) | Object containing information about the task the host is currently running | [optional] 
**StartedBy** | Pointer to **string** | Name of the process or user that started this host | [optional] 
**Status** | Pointer to **string** | The current state of the host | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**TotalIdleTime** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** | The user associated with this host. Set if this host was spawned for a specific user | [optional] 
**UserHost** | Pointer to **bool** |  | [optional] 
**Zone** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIHost

`func NewModelAPIHost() *ModelAPIHost`

NewModelAPIHost instantiates a new ModelAPIHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIHostWithDefaults

`func NewModelAPIHostWithDefaults() *ModelAPIHost`

NewModelAPIHostWithDefaults instantiates a new ModelAPIHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachedVolumeIds

`func (o *ModelAPIHost) GetAttachedVolumeIds() []string`

GetAttachedVolumeIds returns the AttachedVolumeIds field if non-nil, zero value otherwise.

### GetAttachedVolumeIdsOk

`func (o *ModelAPIHost) GetAttachedVolumeIdsOk() (*[]string, bool)`

GetAttachedVolumeIdsOk returns a tuple with the AttachedVolumeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedVolumeIds

`func (o *ModelAPIHost) SetAttachedVolumeIds(v []string)`

SetAttachedVolumeIds sets AttachedVolumeIds field to given value.

### HasAttachedVolumeIds

`func (o *ModelAPIHost) HasAttachedVolumeIds() bool`

HasAttachedVolumeIds returns a boolean if a field has been set.

### GetCreationTime

`func (o *ModelAPIHost) GetCreationTime() string`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *ModelAPIHost) GetCreationTimeOk() (*string, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *ModelAPIHost) SetCreationTime(v string)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *ModelAPIHost) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelAPIHost) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPIHost) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPIHost) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPIHost) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDistro

`func (o *ModelAPIHost) GetDistro() ModelDistroInfo`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *ModelAPIHost) GetDistroOk() (*ModelDistroInfo, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *ModelAPIHost) SetDistro(v ModelDistroInfo)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *ModelAPIHost) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetExpirationTime

`func (o *ModelAPIHost) GetExpirationTime() string`

GetExpirationTime returns the ExpirationTime field if non-nil, zero value otherwise.

### GetExpirationTimeOk

`func (o *ModelAPIHost) GetExpirationTimeOk() (*string, bool)`

GetExpirationTimeOk returns a tuple with the ExpirationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTime

`func (o *ModelAPIHost) SetExpirationTime(v string)`

SetExpirationTime sets ExpirationTime field to given value.

### HasExpirationTime

`func (o *ModelAPIHost) HasExpirationTime() bool`

HasExpirationTime returns a boolean if a field has been set.

### GetHomeVolumeId

`func (o *ModelAPIHost) GetHomeVolumeId() string`

GetHomeVolumeId returns the HomeVolumeId field if non-nil, zero value otherwise.

### GetHomeVolumeIdOk

`func (o *ModelAPIHost) GetHomeVolumeIdOk() (*string, bool)`

GetHomeVolumeIdOk returns a tuple with the HomeVolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeVolumeId

`func (o *ModelAPIHost) SetHomeVolumeId(v string)`

SetHomeVolumeId sets HomeVolumeId field to given value.

### HasHomeVolumeId

`func (o *ModelAPIHost) HasHomeVolumeId() bool`

HasHomeVolumeId returns a boolean if a field has been set.

### GetHostId

`func (o *ModelAPIHost) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *ModelAPIHost) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *ModelAPIHost) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *ModelAPIHost) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetHostType

`func (o *ModelAPIHost) GetHostType() string`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *ModelAPIHost) GetHostTypeOk() (*string, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *ModelAPIHost) SetHostType(v string)`

SetHostType sets HostType field to given value.

### HasHostType

`func (o *ModelAPIHost) HasHostType() bool`

HasHostType returns a boolean if a field has been set.

### GetHostUrl

`func (o *ModelAPIHost) GetHostUrl() string`

GetHostUrl returns the HostUrl field if non-nil, zero value otherwise.

### GetHostUrlOk

`func (o *ModelAPIHost) GetHostUrlOk() (*string, bool)`

GetHostUrlOk returns a tuple with the HostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUrl

`func (o *ModelAPIHost) SetHostUrl(v string)`

SetHostUrl sets HostUrl field to given value.

### HasHostUrl

`func (o *ModelAPIHost) HasHostUrl() bool`

HasHostUrl returns a boolean if a field has been set.

### GetInstanceTags

`func (o *ModelAPIHost) GetInstanceTags() []HostTag`

GetInstanceTags returns the InstanceTags field if non-nil, zero value otherwise.

### GetInstanceTagsOk

`func (o *ModelAPIHost) GetInstanceTagsOk() (*[]HostTag, bool)`

GetInstanceTagsOk returns a tuple with the InstanceTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTags

`func (o *ModelAPIHost) SetInstanceTags(v []HostTag)`

SetInstanceTags sets InstanceTags field to given value.

### HasInstanceTags

`func (o *ModelAPIHost) HasInstanceTags() bool`

HasInstanceTags returns a boolean if a field has been set.

### GetInstanceType

`func (o *ModelAPIHost) GetInstanceType() string`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ModelAPIHost) GetInstanceTypeOk() (*string, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ModelAPIHost) SetInstanceType(v string)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ModelAPIHost) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetLastCommunication

`func (o *ModelAPIHost) GetLastCommunication() string`

GetLastCommunication returns the LastCommunication field if non-nil, zero value otherwise.

### GetLastCommunicationOk

`func (o *ModelAPIHost) GetLastCommunicationOk() (*string, bool)`

GetLastCommunicationOk returns a tuple with the LastCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunication

`func (o *ModelAPIHost) SetLastCommunication(v string)`

SetLastCommunication sets LastCommunication field to given value.

### HasLastCommunication

`func (o *ModelAPIHost) HasLastCommunication() bool`

HasLastCommunication returns a boolean if a field has been set.

### GetNeedsReprovision

`func (o *ModelAPIHost) GetNeedsReprovision() string`

GetNeedsReprovision returns the NeedsReprovision field if non-nil, zero value otherwise.

### GetNeedsReprovisionOk

`func (o *ModelAPIHost) GetNeedsReprovisionOk() (*string, bool)`

GetNeedsReprovisionOk returns a tuple with the NeedsReprovision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsReprovision

`func (o *ModelAPIHost) SetNeedsReprovision(v string)`

SetNeedsReprovision sets NeedsReprovision field to given value.

### HasNeedsReprovision

`func (o *ModelAPIHost) HasNeedsReprovision() bool`

HasNeedsReprovision returns a boolean if a field has been set.

### GetNoExpiration

`func (o *ModelAPIHost) GetNoExpiration() bool`

GetNoExpiration returns the NoExpiration field if non-nil, zero value otherwise.

### GetNoExpirationOk

`func (o *ModelAPIHost) GetNoExpirationOk() (*bool, bool)`

GetNoExpirationOk returns a tuple with the NoExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoExpiration

`func (o *ModelAPIHost) SetNoExpiration(v bool)`

SetNoExpiration sets NoExpiration field to given value.

### HasNoExpiration

`func (o *ModelAPIHost) HasNoExpiration() bool`

HasNoExpiration returns a boolean if a field has been set.

### GetPersistentDnsName

`func (o *ModelAPIHost) GetPersistentDnsName() string`

GetPersistentDnsName returns the PersistentDnsName field if non-nil, zero value otherwise.

### GetPersistentDnsNameOk

`func (o *ModelAPIHost) GetPersistentDnsNameOk() (*string, bool)`

GetPersistentDnsNameOk returns a tuple with the PersistentDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentDnsName

`func (o *ModelAPIHost) SetPersistentDnsName(v string)`

SetPersistentDnsName sets PersistentDnsName field to given value.

### HasPersistentDnsName

`func (o *ModelAPIHost) HasPersistentDnsName() bool`

HasPersistentDnsName returns a boolean if a field has been set.

### GetProvisionOptions

`func (o *ModelAPIHost) GetProvisionOptions() ModelAPIProvisionOptions`

GetProvisionOptions returns the ProvisionOptions field if non-nil, zero value otherwise.

### GetProvisionOptionsOk

`func (o *ModelAPIHost) GetProvisionOptionsOk() (*ModelAPIProvisionOptions, bool)`

GetProvisionOptionsOk returns a tuple with the ProvisionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionOptions

`func (o *ModelAPIHost) SetProvisionOptions(v ModelAPIProvisionOptions)`

SetProvisionOptions sets ProvisionOptions field to given value.

### HasProvisionOptions

`func (o *ModelAPIHost) HasProvisionOptions() bool`

HasProvisionOptions returns a boolean if a field has been set.

### GetProvisioned

`func (o *ModelAPIHost) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *ModelAPIHost) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *ModelAPIHost) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *ModelAPIHost) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetRunningTask

`func (o *ModelAPIHost) GetRunningTask() ModelTaskInfo`

GetRunningTask returns the RunningTask field if non-nil, zero value otherwise.

### GetRunningTaskOk

`func (o *ModelAPIHost) GetRunningTaskOk() (*ModelTaskInfo, bool)`

GetRunningTaskOk returns a tuple with the RunningTask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningTask

`func (o *ModelAPIHost) SetRunningTask(v ModelTaskInfo)`

SetRunningTask sets RunningTask field to given value.

### HasRunningTask

`func (o *ModelAPIHost) HasRunningTask() bool`

HasRunningTask returns a boolean if a field has been set.

### GetStartedBy

`func (o *ModelAPIHost) GetStartedBy() string`

GetStartedBy returns the StartedBy field if non-nil, zero value otherwise.

### GetStartedByOk

`func (o *ModelAPIHost) GetStartedByOk() (*string, bool)`

GetStartedByOk returns a tuple with the StartedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedBy

`func (o *ModelAPIHost) SetStartedBy(v string)`

SetStartedBy sets StartedBy field to given value.

### HasStartedBy

`func (o *ModelAPIHost) HasStartedBy() bool`

HasStartedBy returns a boolean if a field has been set.

### GetStatus

`func (o *ModelAPIHost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ModelAPIHost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ModelAPIHost) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ModelAPIHost) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTag

`func (o *ModelAPIHost) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ModelAPIHost) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ModelAPIHost) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ModelAPIHost) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetTotalIdleTime

`func (o *ModelAPIHost) GetTotalIdleTime() int32`

GetTotalIdleTime returns the TotalIdleTime field if non-nil, zero value otherwise.

### GetTotalIdleTimeOk

`func (o *ModelAPIHost) GetTotalIdleTimeOk() (*int32, bool)`

GetTotalIdleTimeOk returns a tuple with the TotalIdleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIdleTime

`func (o *ModelAPIHost) SetTotalIdleTime(v int32)`

SetTotalIdleTime sets TotalIdleTime field to given value.

### HasTotalIdleTime

`func (o *ModelAPIHost) HasTotalIdleTime() bool`

HasTotalIdleTime returns a boolean if a field has been set.

### GetUser

`func (o *ModelAPIHost) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelAPIHost) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelAPIHost) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelAPIHost) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserHost

`func (o *ModelAPIHost) GetUserHost() bool`

GetUserHost returns the UserHost field if non-nil, zero value otherwise.

### GetUserHostOk

`func (o *ModelAPIHost) GetUserHostOk() (*bool, bool)`

GetUserHostOk returns a tuple with the UserHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserHost

`func (o *ModelAPIHost) SetUserHost(v bool)`

SetUserHost sets UserHost field to given value.

### HasUserHost

`func (o *ModelAPIHost) HasUserHost() bool`

HasUserHost returns a boolean if a field has been set.

### GetZone

`func (o *ModelAPIHost) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ModelAPIHost) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ModelAPIHost) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ModelAPIHost) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


