# ModelAPIFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** |  | [optional] 
**IgnoreForFetch** | Pointer to **bool** | When true, these artifacts are excluded from reproduction | [optional] 
**Name** | Pointer to **string** | Human-readable name of the file | [optional] 
**Url** | Pointer to **string** | Link to the file | [optional] 
**UrlParsley** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** | Determines who can see the file in the UI | [optional] 

## Methods

### NewModelAPIFile

`func NewModelAPIFile() *ModelAPIFile`

NewModelAPIFile instantiates a new ModelAPIFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIFileWithDefaults

`func NewModelAPIFileWithDefaults() *ModelAPIFile`

NewModelAPIFileWithDefaults instantiates a new ModelAPIFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ModelAPIFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ModelAPIFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ModelAPIFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ModelAPIFile) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetIgnoreForFetch

`func (o *ModelAPIFile) GetIgnoreForFetch() bool`

GetIgnoreForFetch returns the IgnoreForFetch field if non-nil, zero value otherwise.

### GetIgnoreForFetchOk

`func (o *ModelAPIFile) GetIgnoreForFetchOk() (*bool, bool)`

GetIgnoreForFetchOk returns a tuple with the IgnoreForFetch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreForFetch

`func (o *ModelAPIFile) SetIgnoreForFetch(v bool)`

SetIgnoreForFetch sets IgnoreForFetch field to given value.

### HasIgnoreForFetch

`func (o *ModelAPIFile) HasIgnoreForFetch() bool`

HasIgnoreForFetch returns a boolean if a field has been set.

### GetName

`func (o *ModelAPIFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelAPIFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelAPIFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelAPIFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *ModelAPIFile) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelAPIFile) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelAPIFile) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelAPIFile) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUrlParsley

`func (o *ModelAPIFile) GetUrlParsley() string`

GetUrlParsley returns the UrlParsley field if non-nil, zero value otherwise.

### GetUrlParsleyOk

`func (o *ModelAPIFile) GetUrlParsleyOk() (*string, bool)`

GetUrlParsleyOk returns a tuple with the UrlParsley field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlParsley

`func (o *ModelAPIFile) SetUrlParsley(v string)`

SetUrlParsley sets UrlParsley field to given value.

### HasUrlParsley

`func (o *ModelAPIFile) HasUrlParsley() bool`

HasUrlParsley returns a boolean if a field has been set.

### GetVisibility

`func (o *ModelAPIFile) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ModelAPIFile) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ModelAPIFile) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ModelAPIFile) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


