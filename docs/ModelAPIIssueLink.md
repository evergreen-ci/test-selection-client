# ModelAPIIssueLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfidenceScore** | Pointer to **float32** | The confidence score of the issue | [optional] 
**IssueKey** | Pointer to **string** | Text to be displayed | [optional] 
**Source** | Pointer to [**ModelAPISource**](ModelAPISource.md) | The source of the edit | [optional] 
**Url** | Pointer to **string** | The url of the ticket | [optional] 

## Methods

### NewModelAPIIssueLink

`func NewModelAPIIssueLink() *ModelAPIIssueLink`

NewModelAPIIssueLink instantiates a new ModelAPIIssueLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIIssueLinkWithDefaults

`func NewModelAPIIssueLinkWithDefaults() *ModelAPIIssueLink`

NewModelAPIIssueLinkWithDefaults instantiates a new ModelAPIIssueLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfidenceScore

`func (o *ModelAPIIssueLink) GetConfidenceScore() float32`

GetConfidenceScore returns the ConfidenceScore field if non-nil, zero value otherwise.

### GetConfidenceScoreOk

`func (o *ModelAPIIssueLink) GetConfidenceScoreOk() (*float32, bool)`

GetConfidenceScoreOk returns a tuple with the ConfidenceScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceScore

`func (o *ModelAPIIssueLink) SetConfidenceScore(v float32)`

SetConfidenceScore sets ConfidenceScore field to given value.

### HasConfidenceScore

`func (o *ModelAPIIssueLink) HasConfidenceScore() bool`

HasConfidenceScore returns a boolean if a field has been set.

### GetIssueKey

`func (o *ModelAPIIssueLink) GetIssueKey() string`

GetIssueKey returns the IssueKey field if non-nil, zero value otherwise.

### GetIssueKeyOk

`func (o *ModelAPIIssueLink) GetIssueKeyOk() (*string, bool)`

GetIssueKeyOk returns a tuple with the IssueKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueKey

`func (o *ModelAPIIssueLink) SetIssueKey(v string)`

SetIssueKey sets IssueKey field to given value.

### HasIssueKey

`func (o *ModelAPIIssueLink) HasIssueKey() bool`

HasIssueKey returns a boolean if a field has been set.

### GetSource

`func (o *ModelAPIIssueLink) GetSource() ModelAPISource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ModelAPIIssueLink) GetSourceOk() (*ModelAPISource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ModelAPIIssueLink) SetSource(v ModelAPISource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ModelAPIIssueLink) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetUrl

`func (o *ModelAPIIssueLink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ModelAPIIssueLink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ModelAPIIssueLink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ModelAPIIssueLink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


