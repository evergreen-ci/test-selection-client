# ModelAPITaskAnnotationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebHook** | Pointer to [**ModelAPIWebHook**](ModelAPIWebHook.md) | Options for webhooks. | [optional] 

## Methods

### NewModelAPITaskAnnotationSettings

`func NewModelAPITaskAnnotationSettings() *ModelAPITaskAnnotationSettings`

NewModelAPITaskAnnotationSettings instantiates a new ModelAPITaskAnnotationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPITaskAnnotationSettingsWithDefaults

`func NewModelAPITaskAnnotationSettingsWithDefaults() *ModelAPITaskAnnotationSettings`

NewModelAPITaskAnnotationSettingsWithDefaults instantiates a new ModelAPITaskAnnotationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebHook

`func (o *ModelAPITaskAnnotationSettings) GetWebHook() ModelAPIWebHook`

GetWebHook returns the WebHook field if non-nil, zero value otherwise.

### GetWebHookOk

`func (o *ModelAPITaskAnnotationSettings) GetWebHookOk() (*ModelAPIWebHook, bool)`

GetWebHookOk returns a tuple with the WebHook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebHook

`func (o *ModelAPITaskAnnotationSettings) SetWebHook(v ModelAPIWebHook)`

SetWebHook sets WebHook field to given value.

### HasWebHook

`func (o *ModelAPITaskAnnotationSettings) HasWebHook() bool`

HasWebHook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


