# VariantStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariantName** | **string** |  | 
**TaskStats** | [**map[string]TaskStateInfo**](TaskStateInfo.md) |  | 

## Methods

### NewVariantStateInfo

`func NewVariantStateInfo(variantName string, taskStats map[string]TaskStateInfo, ) *VariantStateInfo`

NewVariantStateInfo instantiates a new VariantStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantStateInfoWithDefaults

`func NewVariantStateInfoWithDefaults() *VariantStateInfo`

NewVariantStateInfoWithDefaults instantiates a new VariantStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariantName

`func (o *VariantStateInfo) GetVariantName() string`

GetVariantName returns the VariantName field if non-nil, zero value otherwise.

### GetVariantNameOk

`func (o *VariantStateInfo) GetVariantNameOk() (*string, bool)`

GetVariantNameOk returns a tuple with the VariantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantName

`func (o *VariantStateInfo) SetVariantName(v string)`

SetVariantName sets VariantName field to given value.


### GetTaskStats

`func (o *VariantStateInfo) GetTaskStats() map[string]TaskStateInfo`

GetTaskStats returns the TaskStats field if non-nil, zero value otherwise.

### GetTaskStatsOk

`func (o *VariantStateInfo) GetTaskStatsOk() (*map[string]TaskStateInfo, bool)`

GetTaskStatsOk returns a tuple with the TaskStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskStats

`func (o *VariantStateInfo) SetTaskStats(v map[string]TaskStateInfo)`

SetTaskStats sets TaskStats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


