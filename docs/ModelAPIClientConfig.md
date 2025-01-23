# ModelAPIClientConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientBinaries** | Pointer to [**[]ModelAPIClientBinary**](ModelAPIClientBinary.md) |  | [optional] 
**LatestRevision** | Pointer to **string** |  | [optional] 
**S3ClientBinaries** | Pointer to [**[]ModelAPIClientBinary**](ModelAPIClientBinary.md) |  | [optional] 

## Methods

### NewModelAPIClientConfig

`func NewModelAPIClientConfig() *ModelAPIClientConfig`

NewModelAPIClientConfig instantiates a new ModelAPIClientConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIClientConfigWithDefaults

`func NewModelAPIClientConfigWithDefaults() *ModelAPIClientConfig`

NewModelAPIClientConfigWithDefaults instantiates a new ModelAPIClientConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientBinaries

`func (o *ModelAPIClientConfig) GetClientBinaries() []ModelAPIClientBinary`

GetClientBinaries returns the ClientBinaries field if non-nil, zero value otherwise.

### GetClientBinariesOk

`func (o *ModelAPIClientConfig) GetClientBinariesOk() (*[]ModelAPIClientBinary, bool)`

GetClientBinariesOk returns a tuple with the ClientBinaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBinaries

`func (o *ModelAPIClientConfig) SetClientBinaries(v []ModelAPIClientBinary)`

SetClientBinaries sets ClientBinaries field to given value.

### HasClientBinaries

`func (o *ModelAPIClientConfig) HasClientBinaries() bool`

HasClientBinaries returns a boolean if a field has been set.

### GetLatestRevision

`func (o *ModelAPIClientConfig) GetLatestRevision() string`

GetLatestRevision returns the LatestRevision field if non-nil, zero value otherwise.

### GetLatestRevisionOk

`func (o *ModelAPIClientConfig) GetLatestRevisionOk() (*string, bool)`

GetLatestRevisionOk returns a tuple with the LatestRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevision

`func (o *ModelAPIClientConfig) SetLatestRevision(v string)`

SetLatestRevision sets LatestRevision field to given value.

### HasLatestRevision

`func (o *ModelAPIClientConfig) HasLatestRevision() bool`

HasLatestRevision returns a boolean if a field has been set.

### GetS3ClientBinaries

`func (o *ModelAPIClientConfig) GetS3ClientBinaries() []ModelAPIClientBinary`

GetS3ClientBinaries returns the S3ClientBinaries field if non-nil, zero value otherwise.

### GetS3ClientBinariesOk

`func (o *ModelAPIClientConfig) GetS3ClientBinariesOk() (*[]ModelAPIClientBinary, bool)`

GetS3ClientBinariesOk returns a tuple with the S3ClientBinaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3ClientBinaries

`func (o *ModelAPIClientConfig) SetS3ClientBinaries(v []ModelAPIClientBinary)`

SetS3ClientBinaries sets S3ClientBinaries field to given value.

### HasS3ClientBinaries

`func (o *ModelAPIClientConfig) HasS3ClientBinaries() bool`

HasS3ClientBinaries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


