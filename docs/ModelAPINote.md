# ModelAPINote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | Comment about the task failure | [optional] 
**Source** | Pointer to [**ModelAPISource**](ModelAPISource.md) | The source of the note | [optional] 

## Methods

### NewModelAPINote

`func NewModelAPINote() *ModelAPINote`

NewModelAPINote instantiates a new ModelAPINote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPINoteWithDefaults

`func NewModelAPINoteWithDefaults() *ModelAPINote`

NewModelAPINoteWithDefaults instantiates a new ModelAPINote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ModelAPINote) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ModelAPINote) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ModelAPINote) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ModelAPINote) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSource

`func (o *ModelAPINote) GetSource() ModelAPISource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ModelAPINote) GetSourceOk() (*ModelAPISource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ModelAPINote) SetSource(v ModelAPISource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ModelAPINote) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


