# RouteCopyVariablesOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyTo** | Pointer to **string** | Required. ProjectID to copy source_project variables to. | [optional] 
**DryRun** | Pointer to **bool** | If set to true, route returns the variables from source_project that will be copied. (If private, the values will be redacted.) If dry_run is set, then the route does not complete the copy, but returns OK if no project variables in the source project will be overwritten (this concerns all variables in the destination project, but only redacted variables in the source project). Otherwise, an error is given which includes the project variable keys that overlap. If dry_run is not set, the copy is completed, and variables could be overwritten. | [optional] 
**IncludePrivate** | Pointer to **bool** | If set to true, private variables will also be copied. | [optional] 
**Overwrite** | Pointer to **bool** | If set to true, will remove variables from the copy_to project that are not in source_project. | [optional] 

## Methods

### NewRouteCopyVariablesOptions

`func NewRouteCopyVariablesOptions() *RouteCopyVariablesOptions`

NewRouteCopyVariablesOptions instantiates a new RouteCopyVariablesOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteCopyVariablesOptionsWithDefaults

`func NewRouteCopyVariablesOptionsWithDefaults() *RouteCopyVariablesOptions`

NewRouteCopyVariablesOptionsWithDefaults instantiates a new RouteCopyVariablesOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyTo

`func (o *RouteCopyVariablesOptions) GetCopyTo() string`

GetCopyTo returns the CopyTo field if non-nil, zero value otherwise.

### GetCopyToOk

`func (o *RouteCopyVariablesOptions) GetCopyToOk() (*string, bool)`

GetCopyToOk returns a tuple with the CopyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyTo

`func (o *RouteCopyVariablesOptions) SetCopyTo(v string)`

SetCopyTo sets CopyTo field to given value.

### HasCopyTo

`func (o *RouteCopyVariablesOptions) HasCopyTo() bool`

HasCopyTo returns a boolean if a field has been set.

### GetDryRun

`func (o *RouteCopyVariablesOptions) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *RouteCopyVariablesOptions) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *RouteCopyVariablesOptions) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.

### HasDryRun

`func (o *RouteCopyVariablesOptions) HasDryRun() bool`

HasDryRun returns a boolean if a field has been set.

### GetIncludePrivate

`func (o *RouteCopyVariablesOptions) GetIncludePrivate() bool`

GetIncludePrivate returns the IncludePrivate field if non-nil, zero value otherwise.

### GetIncludePrivateOk

`func (o *RouteCopyVariablesOptions) GetIncludePrivateOk() (*bool, bool)`

GetIncludePrivateOk returns a tuple with the IncludePrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePrivate

`func (o *RouteCopyVariablesOptions) SetIncludePrivate(v bool)`

SetIncludePrivate sets IncludePrivate field to given value.

### HasIncludePrivate

`func (o *RouteCopyVariablesOptions) HasIncludePrivate() bool`

HasIncludePrivate returns a boolean if a field has been set.

### GetOverwrite

`func (o *RouteCopyVariablesOptions) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *RouteCopyVariablesOptions) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *RouteCopyVariablesOptions) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *RouteCopyVariablesOptions) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


