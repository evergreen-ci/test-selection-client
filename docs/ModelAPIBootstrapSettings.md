# ModelAPIBootstrapSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientDir** | Pointer to **string** |  | [optional] 
**Communication** | Pointer to **string** |  | [optional] 
**Env** | Pointer to [**[]ModelAPIEnvVar**](ModelAPIEnvVar.md) |  | [optional] 
**JasperBinaryDir** | Pointer to **string** |  | [optional] 
**JasperCredentialsPath** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**PreconditionScripts** | Pointer to [**[]ModelAPIPreconditionScript**](ModelAPIPreconditionScript.md) |  | [optional] 
**ResourceLimits** | Pointer to [**ModelAPIResourceLimits**](ModelAPIResourceLimits.md) |  | [optional] 
**RootDir** | Pointer to **string** |  | [optional] 
**ServiceUser** | Pointer to **string** |  | [optional] 
**ShellPath** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIBootstrapSettings

`func NewModelAPIBootstrapSettings() *ModelAPIBootstrapSettings`

NewModelAPIBootstrapSettings instantiates a new ModelAPIBootstrapSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIBootstrapSettingsWithDefaults

`func NewModelAPIBootstrapSettingsWithDefaults() *ModelAPIBootstrapSettings`

NewModelAPIBootstrapSettingsWithDefaults instantiates a new ModelAPIBootstrapSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientDir

`func (o *ModelAPIBootstrapSettings) GetClientDir() string`

GetClientDir returns the ClientDir field if non-nil, zero value otherwise.

### GetClientDirOk

`func (o *ModelAPIBootstrapSettings) GetClientDirOk() (*string, bool)`

GetClientDirOk returns a tuple with the ClientDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDir

`func (o *ModelAPIBootstrapSettings) SetClientDir(v string)`

SetClientDir sets ClientDir field to given value.

### HasClientDir

`func (o *ModelAPIBootstrapSettings) HasClientDir() bool`

HasClientDir returns a boolean if a field has been set.

### GetCommunication

`func (o *ModelAPIBootstrapSettings) GetCommunication() string`

GetCommunication returns the Communication field if non-nil, zero value otherwise.

### GetCommunicationOk

`func (o *ModelAPIBootstrapSettings) GetCommunicationOk() (*string, bool)`

GetCommunicationOk returns a tuple with the Communication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunication

`func (o *ModelAPIBootstrapSettings) SetCommunication(v string)`

SetCommunication sets Communication field to given value.

### HasCommunication

`func (o *ModelAPIBootstrapSettings) HasCommunication() bool`

HasCommunication returns a boolean if a field has been set.

### GetEnv

`func (o *ModelAPIBootstrapSettings) GetEnv() []ModelAPIEnvVar`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *ModelAPIBootstrapSettings) GetEnvOk() (*[]ModelAPIEnvVar, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *ModelAPIBootstrapSettings) SetEnv(v []ModelAPIEnvVar)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *ModelAPIBootstrapSettings) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetJasperBinaryDir

`func (o *ModelAPIBootstrapSettings) GetJasperBinaryDir() string`

GetJasperBinaryDir returns the JasperBinaryDir field if non-nil, zero value otherwise.

### GetJasperBinaryDirOk

`func (o *ModelAPIBootstrapSettings) GetJasperBinaryDirOk() (*string, bool)`

GetJasperBinaryDirOk returns a tuple with the JasperBinaryDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJasperBinaryDir

`func (o *ModelAPIBootstrapSettings) SetJasperBinaryDir(v string)`

SetJasperBinaryDir sets JasperBinaryDir field to given value.

### HasJasperBinaryDir

`func (o *ModelAPIBootstrapSettings) HasJasperBinaryDir() bool`

HasJasperBinaryDir returns a boolean if a field has been set.

### GetJasperCredentialsPath

`func (o *ModelAPIBootstrapSettings) GetJasperCredentialsPath() string`

GetJasperCredentialsPath returns the JasperCredentialsPath field if non-nil, zero value otherwise.

### GetJasperCredentialsPathOk

`func (o *ModelAPIBootstrapSettings) GetJasperCredentialsPathOk() (*string, bool)`

GetJasperCredentialsPathOk returns a tuple with the JasperCredentialsPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJasperCredentialsPath

`func (o *ModelAPIBootstrapSettings) SetJasperCredentialsPath(v string)`

SetJasperCredentialsPath sets JasperCredentialsPath field to given value.

### HasJasperCredentialsPath

`func (o *ModelAPIBootstrapSettings) HasJasperCredentialsPath() bool`

HasJasperCredentialsPath returns a boolean if a field has been set.

### GetMethod

`func (o *ModelAPIBootstrapSettings) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ModelAPIBootstrapSettings) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ModelAPIBootstrapSettings) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *ModelAPIBootstrapSettings) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetPreconditionScripts

`func (o *ModelAPIBootstrapSettings) GetPreconditionScripts() []ModelAPIPreconditionScript`

GetPreconditionScripts returns the PreconditionScripts field if non-nil, zero value otherwise.

### GetPreconditionScriptsOk

`func (o *ModelAPIBootstrapSettings) GetPreconditionScriptsOk() (*[]ModelAPIPreconditionScript, bool)`

GetPreconditionScriptsOk returns a tuple with the PreconditionScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreconditionScripts

`func (o *ModelAPIBootstrapSettings) SetPreconditionScripts(v []ModelAPIPreconditionScript)`

SetPreconditionScripts sets PreconditionScripts field to given value.

### HasPreconditionScripts

`func (o *ModelAPIBootstrapSettings) HasPreconditionScripts() bool`

HasPreconditionScripts returns a boolean if a field has been set.

### GetResourceLimits

`func (o *ModelAPIBootstrapSettings) GetResourceLimits() ModelAPIResourceLimits`

GetResourceLimits returns the ResourceLimits field if non-nil, zero value otherwise.

### GetResourceLimitsOk

`func (o *ModelAPIBootstrapSettings) GetResourceLimitsOk() (*ModelAPIResourceLimits, bool)`

GetResourceLimitsOk returns a tuple with the ResourceLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceLimits

`func (o *ModelAPIBootstrapSettings) SetResourceLimits(v ModelAPIResourceLimits)`

SetResourceLimits sets ResourceLimits field to given value.

### HasResourceLimits

`func (o *ModelAPIBootstrapSettings) HasResourceLimits() bool`

HasResourceLimits returns a boolean if a field has been set.

### GetRootDir

`func (o *ModelAPIBootstrapSettings) GetRootDir() string`

GetRootDir returns the RootDir field if non-nil, zero value otherwise.

### GetRootDirOk

`func (o *ModelAPIBootstrapSettings) GetRootDirOk() (*string, bool)`

GetRootDirOk returns a tuple with the RootDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDir

`func (o *ModelAPIBootstrapSettings) SetRootDir(v string)`

SetRootDir sets RootDir field to given value.

### HasRootDir

`func (o *ModelAPIBootstrapSettings) HasRootDir() bool`

HasRootDir returns a boolean if a field has been set.

### GetServiceUser

`func (o *ModelAPIBootstrapSettings) GetServiceUser() string`

GetServiceUser returns the ServiceUser field if non-nil, zero value otherwise.

### GetServiceUserOk

`func (o *ModelAPIBootstrapSettings) GetServiceUserOk() (*string, bool)`

GetServiceUserOk returns a tuple with the ServiceUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUser

`func (o *ModelAPIBootstrapSettings) SetServiceUser(v string)`

SetServiceUser sets ServiceUser field to given value.

### HasServiceUser

`func (o *ModelAPIBootstrapSettings) HasServiceUser() bool`

HasServiceUser returns a boolean if a field has been set.

### GetShellPath

`func (o *ModelAPIBootstrapSettings) GetShellPath() string`

GetShellPath returns the ShellPath field if non-nil, zero value otherwise.

### GetShellPathOk

`func (o *ModelAPIBootstrapSettings) GetShellPathOk() (*string, bool)`

GetShellPathOk returns a tuple with the ShellPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShellPath

`func (o *ModelAPIBootstrapSettings) SetShellPath(v string)`

SetShellPath sets ShellPath field to given value.

### HasShellPath

`func (o *ModelAPIBootstrapSettings) HasShellPath() bool`

HasShellPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


