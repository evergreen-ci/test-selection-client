# KnownIssue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**KnownIssueSource**](KnownIssueSource.md) |  | 
**DiscoveredAt** | **time.Time** |  | 

## Methods

### NewKnownIssue

`func NewKnownIssue(source KnownIssueSource, discoveredAt time.Time, ) *KnownIssue`

NewKnownIssue instantiates a new KnownIssue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKnownIssueWithDefaults

`func NewKnownIssueWithDefaults() *KnownIssue`

NewKnownIssueWithDefaults instantiates a new KnownIssue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *KnownIssue) GetSource() KnownIssueSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *KnownIssue) GetSourceOk() (*KnownIssueSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *KnownIssue) SetSource(v KnownIssueSource)`

SetSource sets Source field to given value.


### GetDiscoveredAt

`func (o *KnownIssue) GetDiscoveredAt() time.Time`

GetDiscoveredAt returns the DiscoveredAt field if non-nil, zero value otherwise.

### GetDiscoveredAtOk

`func (o *KnownIssue) GetDiscoveredAtOk() (*time.Time, bool)`

GetDiscoveredAtOk returns a tuple with the DiscoveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredAt

`func (o *KnownIssue) SetDiscoveredAt(v time.Time)`

SetDiscoveredAt sets DiscoveredAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


