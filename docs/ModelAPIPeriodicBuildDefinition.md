# ModelAPIPeriodicBuildDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias to run for the periodic build. | [optional] 
**ConfigFile** | Pointer to **string** | Project config file to use for the periodic build. | [optional] 
**Cron** | Pointer to **string** | Cron specification for when to run periodic builds. | [optional] 
**Id** | Pointer to **string** | Identifier for the periodic build. | [optional] 
**IntervalHours** | Pointer to **int32** | Interval (in hours) between periodic build runs. | [optional] 
**Message** | Pointer to **string** | Message to display in the version metadata. | [optional] 
**NextRunTime** | Pointer to **string** | Next time that the periodic build will run. | [optional] 

## Methods

### NewModelAPIPeriodicBuildDefinition

`func NewModelAPIPeriodicBuildDefinition() *ModelAPIPeriodicBuildDefinition`

NewModelAPIPeriodicBuildDefinition instantiates a new ModelAPIPeriodicBuildDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIPeriodicBuildDefinitionWithDefaults

`func NewModelAPIPeriodicBuildDefinitionWithDefaults() *ModelAPIPeriodicBuildDefinition`

NewModelAPIPeriodicBuildDefinitionWithDefaults instantiates a new ModelAPIPeriodicBuildDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *ModelAPIPeriodicBuildDefinition) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *ModelAPIPeriodicBuildDefinition) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *ModelAPIPeriodicBuildDefinition) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *ModelAPIPeriodicBuildDefinition) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetConfigFile

`func (o *ModelAPIPeriodicBuildDefinition) GetConfigFile() string`

GetConfigFile returns the ConfigFile field if non-nil, zero value otherwise.

### GetConfigFileOk

`func (o *ModelAPIPeriodicBuildDefinition) GetConfigFileOk() (*string, bool)`

GetConfigFileOk returns a tuple with the ConfigFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFile

`func (o *ModelAPIPeriodicBuildDefinition) SetConfigFile(v string)`

SetConfigFile sets ConfigFile field to given value.

### HasConfigFile

`func (o *ModelAPIPeriodicBuildDefinition) HasConfigFile() bool`

HasConfigFile returns a boolean if a field has been set.

### GetCron

`func (o *ModelAPIPeriodicBuildDefinition) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ModelAPIPeriodicBuildDefinition) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ModelAPIPeriodicBuildDefinition) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ModelAPIPeriodicBuildDefinition) HasCron() bool`

HasCron returns a boolean if a field has been set.

### GetId

`func (o *ModelAPIPeriodicBuildDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPIPeriodicBuildDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPIPeriodicBuildDefinition) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPIPeriodicBuildDefinition) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntervalHours

`func (o *ModelAPIPeriodicBuildDefinition) GetIntervalHours() int32`

GetIntervalHours returns the IntervalHours field if non-nil, zero value otherwise.

### GetIntervalHoursOk

`func (o *ModelAPIPeriodicBuildDefinition) GetIntervalHoursOk() (*int32, bool)`

GetIntervalHoursOk returns a tuple with the IntervalHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalHours

`func (o *ModelAPIPeriodicBuildDefinition) SetIntervalHours(v int32)`

SetIntervalHours sets IntervalHours field to given value.

### HasIntervalHours

`func (o *ModelAPIPeriodicBuildDefinition) HasIntervalHours() bool`

HasIntervalHours returns a boolean if a field has been set.

### GetMessage

`func (o *ModelAPIPeriodicBuildDefinition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelAPIPeriodicBuildDefinition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelAPIPeriodicBuildDefinition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelAPIPeriodicBuildDefinition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNextRunTime

`func (o *ModelAPIPeriodicBuildDefinition) GetNextRunTime() string`

GetNextRunTime returns the NextRunTime field if non-nil, zero value otherwise.

### GetNextRunTimeOk

`func (o *ModelAPIPeriodicBuildDefinition) GetNextRunTimeOk() (*string, bool)`

GetNextRunTimeOk returns a tuple with the NextRunTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunTime

`func (o *ModelAPIPeriodicBuildDefinition) SetNextRunTime(v string)`

SetNextRunTime sets NextRunTime field to given value.

### HasNextRunTime

`func (o *ModelAPIPeriodicBuildDefinition) HasNextRunTime() bool`

HasNextRunTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


