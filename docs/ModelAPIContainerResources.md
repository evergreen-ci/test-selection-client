# ModelAPIContainerResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** | CPU (1024 CPU &#x3D; 1 vCPU) for the container. | [optional] 
**MemoryMb** | Pointer to **int32** | Memory (in MB) for the container. | [optional] 
**Name** | Pointer to **string** | Name for container resource definition. | [optional] 

## Methods

### NewModelAPIContainerResources

`func NewModelAPIContainerResources() *ModelAPIContainerResources`

NewModelAPIContainerResources instantiates a new ModelAPIContainerResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIContainerResourcesWithDefaults

`func NewModelAPIContainerResourcesWithDefaults() *ModelAPIContainerResources`

NewModelAPIContainerResourcesWithDefaults instantiates a new ModelAPIContainerResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *ModelAPIContainerResources) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ModelAPIContainerResources) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ModelAPIContainerResources) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ModelAPIContainerResources) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemoryMb

`func (o *ModelAPIContainerResources) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *ModelAPIContainerResources) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *ModelAPIContainerResources) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *ModelAPIContainerResources) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetName

`func (o *ModelAPIContainerResources) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelAPIContainerResources) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelAPIContainerResources) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelAPIContainerResources) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


