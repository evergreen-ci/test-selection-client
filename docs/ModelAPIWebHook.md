# ModelAPIWebHook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | Pointer to **string** | Webhook endpoint | [optional] 
**Secret** | Pointer to **string** | Webhook secret | [optional] 

## Methods

### NewModelAPIWebHook

`func NewModelAPIWebHook() *ModelAPIWebHook`

NewModelAPIWebHook instantiates a new ModelAPIWebHook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIWebHookWithDefaults

`func NewModelAPIWebHookWithDefaults() *ModelAPIWebHook`

NewModelAPIWebHookWithDefaults instantiates a new ModelAPIWebHook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *ModelAPIWebHook) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ModelAPIWebHook) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ModelAPIWebHook) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ModelAPIWebHook) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetSecret

`func (o *ModelAPIWebHook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ModelAPIWebHook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ModelAPIWebHook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ModelAPIWebHook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


