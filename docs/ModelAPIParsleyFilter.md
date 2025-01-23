# ModelAPIParsleyFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaseSensitive** | Pointer to **bool** | CaseSensitive indicates whether the filter is case sensitive. | [optional] 
**ExactMatch** | Pointer to **bool** | ExactMatch indicates whether the filter must be an exact match. | [optional] 
**Expression** | Pointer to **string** | Expression is a regular expression representing the filter. | [optional] 

## Methods

### NewModelAPIParsleyFilter

`func NewModelAPIParsleyFilter() *ModelAPIParsleyFilter`

NewModelAPIParsleyFilter instantiates a new ModelAPIParsleyFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIParsleyFilterWithDefaults

`func NewModelAPIParsleyFilterWithDefaults() *ModelAPIParsleyFilter`

NewModelAPIParsleyFilterWithDefaults instantiates a new ModelAPIParsleyFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaseSensitive

`func (o *ModelAPIParsleyFilter) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *ModelAPIParsleyFilter) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *ModelAPIParsleyFilter) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *ModelAPIParsleyFilter) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetExactMatch

`func (o *ModelAPIParsleyFilter) GetExactMatch() bool`

GetExactMatch returns the ExactMatch field if non-nil, zero value otherwise.

### GetExactMatchOk

`func (o *ModelAPIParsleyFilter) GetExactMatchOk() (*bool, bool)`

GetExactMatchOk returns a tuple with the ExactMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExactMatch

`func (o *ModelAPIParsleyFilter) SetExactMatch(v bool)`

SetExactMatch sets ExactMatch field to given value.

### HasExactMatch

`func (o *ModelAPIParsleyFilter) HasExactMatch() bool`

HasExactMatch returns a boolean if a field has been set.

### GetExpression

`func (o *ModelAPIParsleyFilter) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *ModelAPIParsleyFilter) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *ModelAPIParsleyFilter) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *ModelAPIParsleyFilter) HasExpression() bool`

HasExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


