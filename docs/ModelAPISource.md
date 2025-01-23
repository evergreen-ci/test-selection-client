# ModelAPISource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** | The author of the edit | [optional] 
**Requester** | Pointer to **string** | The source of the request (api or ui) | [optional] 
**Time** | Pointer to **string** | The time of the edit | [optional] 

## Methods

### NewModelAPISource

`func NewModelAPISource() *ModelAPISource`

NewModelAPISource instantiates a new ModelAPISource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPISourceWithDefaults

`func NewModelAPISourceWithDefaults() *ModelAPISource`

NewModelAPISourceWithDefaults instantiates a new ModelAPISource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *ModelAPISource) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelAPISource) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelAPISource) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelAPISource) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetRequester

`func (o *ModelAPISource) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ModelAPISource) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ModelAPISource) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ModelAPISource) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetTime

`func (o *ModelAPISource) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ModelAPISource) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ModelAPISource) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *ModelAPISource) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


