# ModelAPIWorkstationConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitClone** | Pointer to **bool** | Git clone the project in the workstation. | [optional] 
**SetupCommands** | Pointer to [**[]ModelAPIWorkstationSetupCommand**](ModelAPIWorkstationSetupCommand.md) | List of setup commands to run. | [optional] 

## Methods

### NewModelAPIWorkstationConfig

`func NewModelAPIWorkstationConfig() *ModelAPIWorkstationConfig`

NewModelAPIWorkstationConfig instantiates a new ModelAPIWorkstationConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIWorkstationConfigWithDefaults

`func NewModelAPIWorkstationConfigWithDefaults() *ModelAPIWorkstationConfig`

NewModelAPIWorkstationConfigWithDefaults instantiates a new ModelAPIWorkstationConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitClone

`func (o *ModelAPIWorkstationConfig) GetGitClone() bool`

GetGitClone returns the GitClone field if non-nil, zero value otherwise.

### GetGitCloneOk

`func (o *ModelAPIWorkstationConfig) GetGitCloneOk() (*bool, bool)`

GetGitCloneOk returns a tuple with the GitClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitClone

`func (o *ModelAPIWorkstationConfig) SetGitClone(v bool)`

SetGitClone sets GitClone field to given value.

### HasGitClone

`func (o *ModelAPIWorkstationConfig) HasGitClone() bool`

HasGitClone returns a boolean if a field has been set.

### GetSetupCommands

`func (o *ModelAPIWorkstationConfig) GetSetupCommands() []ModelAPIWorkstationSetupCommand`

GetSetupCommands returns the SetupCommands field if non-nil, zero value otherwise.

### GetSetupCommandsOk

`func (o *ModelAPIWorkstationConfig) GetSetupCommandsOk() (*[]ModelAPIWorkstationSetupCommand, bool)`

GetSetupCommandsOk returns a tuple with the SetupCommands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupCommands

`func (o *ModelAPIWorkstationConfig) SetSetupCommands(v []ModelAPIWorkstationSetupCommand)`

SetSetupCommands sets SetupCommands field to given value.

### HasSetupCommands

`func (o *ModelAPIWorkstationConfig) HasSetupCommands() bool`

HasSetupCommands returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


