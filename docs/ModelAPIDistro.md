# ModelAPIDistro

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminOnly** | Pointer to **bool** |  | [optional] 
**Aliases** | Pointer to **[]string** |  | [optional] 
**Arch** | Pointer to **string** |  | [optional] 
**AuthorizedKeysFile** | Pointer to **string** |  | [optional] 
**BootstrapSettings** | Pointer to [**ModelAPIBootstrapSettings**](ModelAPIBootstrapSettings.md) |  | [optional] 
**ContainerPool** | Pointer to **string** |  | [optional] 
**DisableShallowClone** | Pointer to **bool** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DispatcherSettings** | Pointer to [**ModelAPIDispatcherSettings**](ModelAPIDispatcherSettings.md) |  | [optional] 
**ExecUser** | Pointer to **string** |  | [optional] 
**Expansions** | Pointer to [**[]ModelAPIExpansion**](ModelAPIExpansion.md) |  | [optional] 
**FinderSettings** | Pointer to [**ModelAPIFinderSettings**](ModelAPIFinderSettings.md) |  | [optional] 
**HomeVolumeSettings** | Pointer to [**ModelAPIHomeVolumeSettings**](ModelAPIHomeVolumeSettings.md) |  | [optional] 
**HostAllocatorSettings** | Pointer to [**ModelAPIHostAllocatorSettings**](ModelAPIHostAllocatorSettings.md) |  | [optional] 
**IcecreamSettings** | Pointer to [**ModelAPIIceCreamSettings**](ModelAPIIceCreamSettings.md) |  | [optional] 
**ImageId** | Pointer to **string** |  | [optional] 
**IsCluster** | Pointer to **bool** |  | [optional] 
**IsVirtualWorkstation** | Pointer to **bool** |  | [optional] 
**Mountpoints** | Pointer to **[]string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**PlannerSettings** | Pointer to [**ModelAPIPlannerSettings**](ModelAPIPlannerSettings.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**ProviderSettings** | Pointer to **map[string]interface{}** |  | [optional] 
**Setup** | Pointer to **string** |  | [optional] 
**SetupAsSudo** | Pointer to **bool** |  | [optional] 
**SingleTaskDistro** | Pointer to **bool** |  | [optional] 
**SshOptions** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**UserSpawnAllowed** | Pointer to **bool** |  | [optional] 
**ValidProjects** | Pointer to **[]string** |  | [optional] 
**WarningNote** | Pointer to **string** |  | [optional] 
**WorkDir** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIDistro

`func NewModelAPIDistro() *ModelAPIDistro`

NewModelAPIDistro instantiates a new ModelAPIDistro object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIDistroWithDefaults

`func NewModelAPIDistroWithDefaults() *ModelAPIDistro`

NewModelAPIDistroWithDefaults instantiates a new ModelAPIDistro object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminOnly

`func (o *ModelAPIDistro) GetAdminOnly() bool`

GetAdminOnly returns the AdminOnly field if non-nil, zero value otherwise.

### GetAdminOnlyOk

`func (o *ModelAPIDistro) GetAdminOnlyOk() (*bool, bool)`

GetAdminOnlyOk returns a tuple with the AdminOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminOnly

`func (o *ModelAPIDistro) SetAdminOnly(v bool)`

SetAdminOnly sets AdminOnly field to given value.

### HasAdminOnly

`func (o *ModelAPIDistro) HasAdminOnly() bool`

HasAdminOnly returns a boolean if a field has been set.

### GetAliases

`func (o *ModelAPIDistro) GetAliases() []string`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ModelAPIDistro) GetAliasesOk() (*[]string, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ModelAPIDistro) SetAliases(v []string)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ModelAPIDistro) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetArch

`func (o *ModelAPIDistro) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *ModelAPIDistro) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *ModelAPIDistro) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *ModelAPIDistro) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetAuthorizedKeysFile

`func (o *ModelAPIDistro) GetAuthorizedKeysFile() string`

GetAuthorizedKeysFile returns the AuthorizedKeysFile field if non-nil, zero value otherwise.

### GetAuthorizedKeysFileOk

`func (o *ModelAPIDistro) GetAuthorizedKeysFileOk() (*string, bool)`

GetAuthorizedKeysFileOk returns a tuple with the AuthorizedKeysFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeysFile

`func (o *ModelAPIDistro) SetAuthorizedKeysFile(v string)`

SetAuthorizedKeysFile sets AuthorizedKeysFile field to given value.

### HasAuthorizedKeysFile

`func (o *ModelAPIDistro) HasAuthorizedKeysFile() bool`

HasAuthorizedKeysFile returns a boolean if a field has been set.

### GetBootstrapSettings

`func (o *ModelAPIDistro) GetBootstrapSettings() ModelAPIBootstrapSettings`

GetBootstrapSettings returns the BootstrapSettings field if non-nil, zero value otherwise.

### GetBootstrapSettingsOk

`func (o *ModelAPIDistro) GetBootstrapSettingsOk() (*ModelAPIBootstrapSettings, bool)`

GetBootstrapSettingsOk returns a tuple with the BootstrapSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootstrapSettings

`func (o *ModelAPIDistro) SetBootstrapSettings(v ModelAPIBootstrapSettings)`

SetBootstrapSettings sets BootstrapSettings field to given value.

### HasBootstrapSettings

`func (o *ModelAPIDistro) HasBootstrapSettings() bool`

HasBootstrapSettings returns a boolean if a field has been set.

### GetContainerPool

`func (o *ModelAPIDistro) GetContainerPool() string`

GetContainerPool returns the ContainerPool field if non-nil, zero value otherwise.

### GetContainerPoolOk

`func (o *ModelAPIDistro) GetContainerPoolOk() (*string, bool)`

GetContainerPoolOk returns a tuple with the ContainerPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPool

`func (o *ModelAPIDistro) SetContainerPool(v string)`

SetContainerPool sets ContainerPool field to given value.

### HasContainerPool

`func (o *ModelAPIDistro) HasContainerPool() bool`

HasContainerPool returns a boolean if a field has been set.

### GetDisableShallowClone

`func (o *ModelAPIDistro) GetDisableShallowClone() bool`

GetDisableShallowClone returns the DisableShallowClone field if non-nil, zero value otherwise.

### GetDisableShallowCloneOk

`func (o *ModelAPIDistro) GetDisableShallowCloneOk() (*bool, bool)`

GetDisableShallowCloneOk returns a tuple with the DisableShallowClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableShallowClone

`func (o *ModelAPIDistro) SetDisableShallowClone(v bool)`

SetDisableShallowClone sets DisableShallowClone field to given value.

### HasDisableShallowClone

`func (o *ModelAPIDistro) HasDisableShallowClone() bool`

HasDisableShallowClone returns a boolean if a field has been set.

### GetDisabled

`func (o *ModelAPIDistro) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ModelAPIDistro) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ModelAPIDistro) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ModelAPIDistro) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDispatcherSettings

`func (o *ModelAPIDistro) GetDispatcherSettings() ModelAPIDispatcherSettings`

GetDispatcherSettings returns the DispatcherSettings field if non-nil, zero value otherwise.

### GetDispatcherSettingsOk

`func (o *ModelAPIDistro) GetDispatcherSettingsOk() (*ModelAPIDispatcherSettings, bool)`

GetDispatcherSettingsOk returns a tuple with the DispatcherSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatcherSettings

`func (o *ModelAPIDistro) SetDispatcherSettings(v ModelAPIDispatcherSettings)`

SetDispatcherSettings sets DispatcherSettings field to given value.

### HasDispatcherSettings

`func (o *ModelAPIDistro) HasDispatcherSettings() bool`

HasDispatcherSettings returns a boolean if a field has been set.

### GetExecUser

`func (o *ModelAPIDistro) GetExecUser() string`

GetExecUser returns the ExecUser field if non-nil, zero value otherwise.

### GetExecUserOk

`func (o *ModelAPIDistro) GetExecUserOk() (*string, bool)`

GetExecUserOk returns a tuple with the ExecUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecUser

`func (o *ModelAPIDistro) SetExecUser(v string)`

SetExecUser sets ExecUser field to given value.

### HasExecUser

`func (o *ModelAPIDistro) HasExecUser() bool`

HasExecUser returns a boolean if a field has been set.

### GetExpansions

`func (o *ModelAPIDistro) GetExpansions() []ModelAPIExpansion`

GetExpansions returns the Expansions field if non-nil, zero value otherwise.

### GetExpansionsOk

`func (o *ModelAPIDistro) GetExpansionsOk() (*[]ModelAPIExpansion, bool)`

GetExpansionsOk returns a tuple with the Expansions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpansions

`func (o *ModelAPIDistro) SetExpansions(v []ModelAPIExpansion)`

SetExpansions sets Expansions field to given value.

### HasExpansions

`func (o *ModelAPIDistro) HasExpansions() bool`

HasExpansions returns a boolean if a field has been set.

### GetFinderSettings

`func (o *ModelAPIDistro) GetFinderSettings() ModelAPIFinderSettings`

GetFinderSettings returns the FinderSettings field if non-nil, zero value otherwise.

### GetFinderSettingsOk

`func (o *ModelAPIDistro) GetFinderSettingsOk() (*ModelAPIFinderSettings, bool)`

GetFinderSettingsOk returns a tuple with the FinderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinderSettings

`func (o *ModelAPIDistro) SetFinderSettings(v ModelAPIFinderSettings)`

SetFinderSettings sets FinderSettings field to given value.

### HasFinderSettings

`func (o *ModelAPIDistro) HasFinderSettings() bool`

HasFinderSettings returns a boolean if a field has been set.

### GetHomeVolumeSettings

`func (o *ModelAPIDistro) GetHomeVolumeSettings() ModelAPIHomeVolumeSettings`

GetHomeVolumeSettings returns the HomeVolumeSettings field if non-nil, zero value otherwise.

### GetHomeVolumeSettingsOk

`func (o *ModelAPIDistro) GetHomeVolumeSettingsOk() (*ModelAPIHomeVolumeSettings, bool)`

GetHomeVolumeSettingsOk returns a tuple with the HomeVolumeSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeVolumeSettings

`func (o *ModelAPIDistro) SetHomeVolumeSettings(v ModelAPIHomeVolumeSettings)`

SetHomeVolumeSettings sets HomeVolumeSettings field to given value.

### HasHomeVolumeSettings

`func (o *ModelAPIDistro) HasHomeVolumeSettings() bool`

HasHomeVolumeSettings returns a boolean if a field has been set.

### GetHostAllocatorSettings

`func (o *ModelAPIDistro) GetHostAllocatorSettings() ModelAPIHostAllocatorSettings`

GetHostAllocatorSettings returns the HostAllocatorSettings field if non-nil, zero value otherwise.

### GetHostAllocatorSettingsOk

`func (o *ModelAPIDistro) GetHostAllocatorSettingsOk() (*ModelAPIHostAllocatorSettings, bool)`

GetHostAllocatorSettingsOk returns a tuple with the HostAllocatorSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostAllocatorSettings

`func (o *ModelAPIDistro) SetHostAllocatorSettings(v ModelAPIHostAllocatorSettings)`

SetHostAllocatorSettings sets HostAllocatorSettings field to given value.

### HasHostAllocatorSettings

`func (o *ModelAPIDistro) HasHostAllocatorSettings() bool`

HasHostAllocatorSettings returns a boolean if a field has been set.

### GetIcecreamSettings

`func (o *ModelAPIDistro) GetIcecreamSettings() ModelAPIIceCreamSettings`

GetIcecreamSettings returns the IcecreamSettings field if non-nil, zero value otherwise.

### GetIcecreamSettingsOk

`func (o *ModelAPIDistro) GetIcecreamSettingsOk() (*ModelAPIIceCreamSettings, bool)`

GetIcecreamSettingsOk returns a tuple with the IcecreamSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcecreamSettings

`func (o *ModelAPIDistro) SetIcecreamSettings(v ModelAPIIceCreamSettings)`

SetIcecreamSettings sets IcecreamSettings field to given value.

### HasIcecreamSettings

`func (o *ModelAPIDistro) HasIcecreamSettings() bool`

HasIcecreamSettings returns a boolean if a field has been set.

### GetImageId

`func (o *ModelAPIDistro) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *ModelAPIDistro) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *ModelAPIDistro) SetImageId(v string)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *ModelAPIDistro) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetIsCluster

`func (o *ModelAPIDistro) GetIsCluster() bool`

GetIsCluster returns the IsCluster field if non-nil, zero value otherwise.

### GetIsClusterOk

`func (o *ModelAPIDistro) GetIsClusterOk() (*bool, bool)`

GetIsClusterOk returns a tuple with the IsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCluster

`func (o *ModelAPIDistro) SetIsCluster(v bool)`

SetIsCluster sets IsCluster field to given value.

### HasIsCluster

`func (o *ModelAPIDistro) HasIsCluster() bool`

HasIsCluster returns a boolean if a field has been set.

### GetIsVirtualWorkstation

`func (o *ModelAPIDistro) GetIsVirtualWorkstation() bool`

GetIsVirtualWorkstation returns the IsVirtualWorkstation field if non-nil, zero value otherwise.

### GetIsVirtualWorkstationOk

`func (o *ModelAPIDistro) GetIsVirtualWorkstationOk() (*bool, bool)`

GetIsVirtualWorkstationOk returns a tuple with the IsVirtualWorkstation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVirtualWorkstation

`func (o *ModelAPIDistro) SetIsVirtualWorkstation(v bool)`

SetIsVirtualWorkstation sets IsVirtualWorkstation field to given value.

### HasIsVirtualWorkstation

`func (o *ModelAPIDistro) HasIsVirtualWorkstation() bool`

HasIsVirtualWorkstation returns a boolean if a field has been set.

### GetMountpoints

`func (o *ModelAPIDistro) GetMountpoints() []string`

GetMountpoints returns the Mountpoints field if non-nil, zero value otherwise.

### GetMountpointsOk

`func (o *ModelAPIDistro) GetMountpointsOk() (*[]string, bool)`

GetMountpointsOk returns a tuple with the Mountpoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountpoints

`func (o *ModelAPIDistro) SetMountpoints(v []string)`

SetMountpoints sets Mountpoints field to given value.

### HasMountpoints

`func (o *ModelAPIDistro) HasMountpoints() bool`

HasMountpoints returns a boolean if a field has been set.

### GetName

`func (o *ModelAPIDistro) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelAPIDistro) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelAPIDistro) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelAPIDistro) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNote

`func (o *ModelAPIDistro) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ModelAPIDistro) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ModelAPIDistro) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ModelAPIDistro) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetPlannerSettings

`func (o *ModelAPIDistro) GetPlannerSettings() ModelAPIPlannerSettings`

GetPlannerSettings returns the PlannerSettings field if non-nil, zero value otherwise.

### GetPlannerSettingsOk

`func (o *ModelAPIDistro) GetPlannerSettingsOk() (*ModelAPIPlannerSettings, bool)`

GetPlannerSettingsOk returns a tuple with the PlannerSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlannerSettings

`func (o *ModelAPIDistro) SetPlannerSettings(v ModelAPIPlannerSettings)`

SetPlannerSettings sets PlannerSettings field to given value.

### HasPlannerSettings

`func (o *ModelAPIDistro) HasPlannerSettings() bool`

HasPlannerSettings returns a boolean if a field has been set.

### GetProvider

`func (o *ModelAPIDistro) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ModelAPIDistro) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ModelAPIDistro) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ModelAPIDistro) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetProviderSettings

`func (o *ModelAPIDistro) GetProviderSettings() map[string]interface{}`

GetProviderSettings returns the ProviderSettings field if non-nil, zero value otherwise.

### GetProviderSettingsOk

`func (o *ModelAPIDistro) GetProviderSettingsOk() (*map[string]interface{}, bool)`

GetProviderSettingsOk returns a tuple with the ProviderSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSettings

`func (o *ModelAPIDistro) SetProviderSettings(v map[string]interface{})`

SetProviderSettings sets ProviderSettings field to given value.

### HasProviderSettings

`func (o *ModelAPIDistro) HasProviderSettings() bool`

HasProviderSettings returns a boolean if a field has been set.

### GetSetup

`func (o *ModelAPIDistro) GetSetup() string`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *ModelAPIDistro) GetSetupOk() (*string, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *ModelAPIDistro) SetSetup(v string)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *ModelAPIDistro) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSetupAsSudo

`func (o *ModelAPIDistro) GetSetupAsSudo() bool`

GetSetupAsSudo returns the SetupAsSudo field if non-nil, zero value otherwise.

### GetSetupAsSudoOk

`func (o *ModelAPIDistro) GetSetupAsSudoOk() (*bool, bool)`

GetSetupAsSudoOk returns a tuple with the SetupAsSudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupAsSudo

`func (o *ModelAPIDistro) SetSetupAsSudo(v bool)`

SetSetupAsSudo sets SetupAsSudo field to given value.

### HasSetupAsSudo

`func (o *ModelAPIDistro) HasSetupAsSudo() bool`

HasSetupAsSudo returns a boolean if a field has been set.

### GetSingleTaskDistro

`func (o *ModelAPIDistro) GetSingleTaskDistro() bool`

GetSingleTaskDistro returns the SingleTaskDistro field if non-nil, zero value otherwise.

### GetSingleTaskDistroOk

`func (o *ModelAPIDistro) GetSingleTaskDistroOk() (*bool, bool)`

GetSingleTaskDistroOk returns a tuple with the SingleTaskDistro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleTaskDistro

`func (o *ModelAPIDistro) SetSingleTaskDistro(v bool)`

SetSingleTaskDistro sets SingleTaskDistro field to given value.

### HasSingleTaskDistro

`func (o *ModelAPIDistro) HasSingleTaskDistro() bool`

HasSingleTaskDistro returns a boolean if a field has been set.

### GetSshOptions

`func (o *ModelAPIDistro) GetSshOptions() []string`

GetSshOptions returns the SshOptions field if non-nil, zero value otherwise.

### GetSshOptionsOk

`func (o *ModelAPIDistro) GetSshOptionsOk() (*[]string, bool)`

GetSshOptionsOk returns a tuple with the SshOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshOptions

`func (o *ModelAPIDistro) SetSshOptions(v []string)`

SetSshOptions sets SshOptions field to given value.

### HasSshOptions

`func (o *ModelAPIDistro) HasSshOptions() bool`

HasSshOptions returns a boolean if a field has been set.

### GetUser

`func (o *ModelAPIDistro) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelAPIDistro) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelAPIDistro) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelAPIDistro) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserSpawnAllowed

`func (o *ModelAPIDistro) GetUserSpawnAllowed() bool`

GetUserSpawnAllowed returns the UserSpawnAllowed field if non-nil, zero value otherwise.

### GetUserSpawnAllowedOk

`func (o *ModelAPIDistro) GetUserSpawnAllowedOk() (*bool, bool)`

GetUserSpawnAllowedOk returns a tuple with the UserSpawnAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSpawnAllowed

`func (o *ModelAPIDistro) SetUserSpawnAllowed(v bool)`

SetUserSpawnAllowed sets UserSpawnAllowed field to given value.

### HasUserSpawnAllowed

`func (o *ModelAPIDistro) HasUserSpawnAllowed() bool`

HasUserSpawnAllowed returns a boolean if a field has been set.

### GetValidProjects

`func (o *ModelAPIDistro) GetValidProjects() []string`

GetValidProjects returns the ValidProjects field if non-nil, zero value otherwise.

### GetValidProjectsOk

`func (o *ModelAPIDistro) GetValidProjectsOk() (*[]string, bool)`

GetValidProjectsOk returns a tuple with the ValidProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidProjects

`func (o *ModelAPIDistro) SetValidProjects(v []string)`

SetValidProjects sets ValidProjects field to given value.

### HasValidProjects

`func (o *ModelAPIDistro) HasValidProjects() bool`

HasValidProjects returns a boolean if a field has been set.

### GetWarningNote

`func (o *ModelAPIDistro) GetWarningNote() string`

GetWarningNote returns the WarningNote field if non-nil, zero value otherwise.

### GetWarningNoteOk

`func (o *ModelAPIDistro) GetWarningNoteOk() (*string, bool)`

GetWarningNoteOk returns a tuple with the WarningNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningNote

`func (o *ModelAPIDistro) SetWarningNote(v string)`

SetWarningNote sets WarningNote field to given value.

### HasWarningNote

`func (o *ModelAPIDistro) HasWarningNote() bool`

HasWarningNote returns a boolean if a field has been set.

### GetWorkDir

`func (o *ModelAPIDistro) GetWorkDir() string`

GetWorkDir returns the WorkDir field if non-nil, zero value otherwise.

### GetWorkDirOk

`func (o *ModelAPIDistro) GetWorkDirOk() (*string, bool)`

GetWorkDirOk returns a tuple with the WorkDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkDir

`func (o *ModelAPIDistro) SetWorkDir(v string)`

SetWorkDir sets WorkDir field to given value.

### HasWorkDir

`func (o *ModelAPIDistro) HasWorkDir() bool`

HasWorkDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


