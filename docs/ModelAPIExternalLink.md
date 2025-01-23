# ModelAPIExternalLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name for the URL. | [optional] 
**Requesters** | Pointer to **[]string** | Requester filter for when to display the link. | [optional] 
**UrlTemplate** | Pointer to **string** | URL format to add to the version metadata panel. | [optional] 

## Methods

### NewModelAPIExternalLink

`func NewModelAPIExternalLink() *ModelAPIExternalLink`

NewModelAPIExternalLink instantiates a new ModelAPIExternalLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIExternalLinkWithDefaults

`func NewModelAPIExternalLinkWithDefaults() *ModelAPIExternalLink`

NewModelAPIExternalLinkWithDefaults instantiates a new ModelAPIExternalLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *ModelAPIExternalLink) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPIExternalLink) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPIExternalLink) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPIExternalLink) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetRequesters

`func (o *ModelAPIExternalLink) GetRequesters() []string`

GetRequesters returns the Requesters field if non-nil, zero value otherwise.

### GetRequestersOk

`func (o *ModelAPIExternalLink) GetRequestersOk() (*[]string, bool)`

GetRequestersOk returns a tuple with the Requesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesters

`func (o *ModelAPIExternalLink) SetRequesters(v []string)`

SetRequesters sets Requesters field to given value.

### HasRequesters

`func (o *ModelAPIExternalLink) HasRequesters() bool`

HasRequesters returns a boolean if a field has been set.

### GetUrlTemplate

`func (o *ModelAPIExternalLink) GetUrlTemplate() string`

GetUrlTemplate returns the UrlTemplate field if non-nil, zero value otherwise.

### GetUrlTemplateOk

`func (o *ModelAPIExternalLink) GetUrlTemplateOk() (*string, bool)`

GetUrlTemplateOk returns a tuple with the UrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTemplate

`func (o *ModelAPIExternalLink) SetUrlTemplate(v string)`

SetUrlTemplate sets UrlTemplate field to given value.

### HasUrlTemplate

`func (o *ModelAPIExternalLink) HasUrlTemplate() bool`

HasUrlTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


