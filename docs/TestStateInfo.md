# TestStateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**StateMachineEnum**](StateMachineEnum.md) |  | 
**OverrideState** | Pointer to [**NullableStateMachineEnum**](StateMachineEnum.md) |  | [optional] 

## Methods

### NewTestStateInfo

`func NewTestStateInfo(state StateMachineEnum, ) *TestStateInfo`

NewTestStateInfo instantiates a new TestStateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTestStateInfoWithDefaults

`func NewTestStateInfoWithDefaults() *TestStateInfo`

NewTestStateInfoWithDefaults instantiates a new TestStateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *TestStateInfo) GetState() StateMachineEnum`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TestStateInfo) GetStateOk() (*StateMachineEnum, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TestStateInfo) SetState(v StateMachineEnum)`

SetState sets State field to given value.


### GetOverrideState

`func (o *TestStateInfo) GetOverrideState() StateMachineEnum`

GetOverrideState returns the OverrideState field if non-nil, zero value otherwise.

### GetOverrideStateOk

`func (o *TestStateInfo) GetOverrideStateOk() (*StateMachineEnum, bool)`

GetOverrideStateOk returns a tuple with the OverrideState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideState

`func (o *TestStateInfo) SetOverrideState(v StateMachineEnum)`

SetOverrideState sets OverrideState field to given value.

### HasOverrideState

`func (o *TestStateInfo) HasOverrideState() bool`

HasOverrideState returns a boolean if a field has been set.

### SetOverrideStateNil

`func (o *TestStateInfo) SetOverrideStateNil(b bool)`

 SetOverrideStateNil sets the value for OverrideState to be an explicit nil

### UnsetOverrideState
`func (o *TestStateInfo) UnsetOverrideState()`

UnsetOverrideState ensures that no value is present for OverrideState, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


