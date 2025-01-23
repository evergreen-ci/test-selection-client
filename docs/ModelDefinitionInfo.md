# ModelDefinitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Batchtime** | Pointer to **int32** | The batchtime defined for the variant, if provided, as defined in the project settings | [optional] 
**Cron** | Pointer to **string** | The cron defined for the variant, if provided, as defined in the project settings | [optional] 

## Methods

### NewModelDefinitionInfo

`func NewModelDefinitionInfo() *ModelDefinitionInfo`

NewModelDefinitionInfo instantiates a new ModelDefinitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelDefinitionInfoWithDefaults

`func NewModelDefinitionInfoWithDefaults() *ModelDefinitionInfo`

NewModelDefinitionInfoWithDefaults instantiates a new ModelDefinitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBatchtime

`func (o *ModelDefinitionInfo) GetBatchtime() int32`

GetBatchtime returns the Batchtime field if non-nil, zero value otherwise.

### GetBatchtimeOk

`func (o *ModelDefinitionInfo) GetBatchtimeOk() (*int32, bool)`

GetBatchtimeOk returns a tuple with the Batchtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchtime

`func (o *ModelDefinitionInfo) SetBatchtime(v int32)`

SetBatchtime sets Batchtime field to given value.

### HasBatchtime

`func (o *ModelDefinitionInfo) HasBatchtime() bool`

HasBatchtime returns a boolean if a field has been set.

### GetCron

`func (o *ModelDefinitionInfo) GetCron() string`

GetCron returns the Cron field if non-nil, zero value otherwise.

### GetCronOk

`func (o *ModelDefinitionInfo) GetCronOk() (*string, bool)`

GetCronOk returns a tuple with the Cron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCron

`func (o *ModelDefinitionInfo) SetCron(v string)`

SetCron sets Cron field to given value.

### HasCron

`func (o *ModelDefinitionInfo) HasCron() bool`

HasCron returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


