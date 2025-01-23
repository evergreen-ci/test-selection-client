# ModelAPIFeedbackSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Questions** | Pointer to [**[]ModelAPIQuestionAnswer**](ModelAPIQuestionAnswer.md) |  | [optional] 
**SubmittedAt** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewModelAPIFeedbackSubmission

`func NewModelAPIFeedbackSubmission() *ModelAPIFeedbackSubmission`

NewModelAPIFeedbackSubmission instantiates a new ModelAPIFeedbackSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIFeedbackSubmissionWithDefaults

`func NewModelAPIFeedbackSubmissionWithDefaults() *ModelAPIFeedbackSubmission`

NewModelAPIFeedbackSubmissionWithDefaults instantiates a new ModelAPIFeedbackSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestions

`func (o *ModelAPIFeedbackSubmission) GetQuestions() []ModelAPIQuestionAnswer`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *ModelAPIFeedbackSubmission) GetQuestionsOk() (*[]ModelAPIQuestionAnswer, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *ModelAPIFeedbackSubmission) SetQuestions(v []ModelAPIQuestionAnswer)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *ModelAPIFeedbackSubmission) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *ModelAPIFeedbackSubmission) GetSubmittedAt() string`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *ModelAPIFeedbackSubmission) GetSubmittedAtOk() (*string, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *ModelAPIFeedbackSubmission) SetSubmittedAt(v string)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *ModelAPIFeedbackSubmission) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetType

`func (o *ModelAPIFeedbackSubmission) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelAPIFeedbackSubmission) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelAPIFeedbackSubmission) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ModelAPIFeedbackSubmission) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUser

`func (o *ModelAPIFeedbackSubmission) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ModelAPIFeedbackSubmission) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ModelAPIFeedbackSubmission) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *ModelAPIFeedbackSubmission) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


