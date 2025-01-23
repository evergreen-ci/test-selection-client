# ModelAPIUserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateFormat** | Pointer to **string** |  | [optional] 
**GithubUser** | Pointer to [**ModelAPIGithubUser**](ModelAPIGithubUser.md) |  | [optional] 
**Notifications** | Pointer to [**ModelAPINotificationPreferences**](ModelAPINotificationPreferences.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SlackMemberId** | Pointer to **string** |  | [optional] 
**SlackUsername** | Pointer to **string** |  | [optional] 
**SpruceFeedback** | Pointer to [**ModelAPIFeedbackSubmission**](ModelAPIFeedbackSubmission.md) |  | [optional] 
**TimeFormat** | Pointer to **string** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**UseSpruceOptions** | Pointer to [**ModelAPIUseSpruceOptions**](ModelAPIUseSpruceOptions.md) |  | [optional] 

## Methods

### NewModelAPIUserSettings

`func NewModelAPIUserSettings() *ModelAPIUserSettings`

NewModelAPIUserSettings instantiates a new ModelAPIUserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIUserSettingsWithDefaults

`func NewModelAPIUserSettingsWithDefaults() *ModelAPIUserSettings`

NewModelAPIUserSettingsWithDefaults instantiates a new ModelAPIUserSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateFormat

`func (o *ModelAPIUserSettings) GetDateFormat() string`

GetDateFormat returns the DateFormat field if non-nil, zero value otherwise.

### GetDateFormatOk

`func (o *ModelAPIUserSettings) GetDateFormatOk() (*string, bool)`

GetDateFormatOk returns a tuple with the DateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateFormat

`func (o *ModelAPIUserSettings) SetDateFormat(v string)`

SetDateFormat sets DateFormat field to given value.

### HasDateFormat

`func (o *ModelAPIUserSettings) HasDateFormat() bool`

HasDateFormat returns a boolean if a field has been set.

### GetGithubUser

`func (o *ModelAPIUserSettings) GetGithubUser() ModelAPIGithubUser`

GetGithubUser returns the GithubUser field if non-nil, zero value otherwise.

### GetGithubUserOk

`func (o *ModelAPIUserSettings) GetGithubUserOk() (*ModelAPIGithubUser, bool)`

GetGithubUserOk returns a tuple with the GithubUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubUser

`func (o *ModelAPIUserSettings) SetGithubUser(v ModelAPIGithubUser)`

SetGithubUser sets GithubUser field to given value.

### HasGithubUser

`func (o *ModelAPIUserSettings) HasGithubUser() bool`

HasGithubUser returns a boolean if a field has been set.

### GetNotifications

`func (o *ModelAPIUserSettings) GetNotifications() ModelAPINotificationPreferences`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *ModelAPIUserSettings) GetNotificationsOk() (*ModelAPINotificationPreferences, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *ModelAPIUserSettings) SetNotifications(v ModelAPINotificationPreferences)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *ModelAPIUserSettings) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetRegion

`func (o *ModelAPIUserSettings) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ModelAPIUserSettings) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ModelAPIUserSettings) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ModelAPIUserSettings) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSlackMemberId

`func (o *ModelAPIUserSettings) GetSlackMemberId() string`

GetSlackMemberId returns the SlackMemberId field if non-nil, zero value otherwise.

### GetSlackMemberIdOk

`func (o *ModelAPIUserSettings) GetSlackMemberIdOk() (*string, bool)`

GetSlackMemberIdOk returns a tuple with the SlackMemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackMemberId

`func (o *ModelAPIUserSettings) SetSlackMemberId(v string)`

SetSlackMemberId sets SlackMemberId field to given value.

### HasSlackMemberId

`func (o *ModelAPIUserSettings) HasSlackMemberId() bool`

HasSlackMemberId returns a boolean if a field has been set.

### GetSlackUsername

`func (o *ModelAPIUserSettings) GetSlackUsername() string`

GetSlackUsername returns the SlackUsername field if non-nil, zero value otherwise.

### GetSlackUsernameOk

`func (o *ModelAPIUserSettings) GetSlackUsernameOk() (*string, bool)`

GetSlackUsernameOk returns a tuple with the SlackUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackUsername

`func (o *ModelAPIUserSettings) SetSlackUsername(v string)`

SetSlackUsername sets SlackUsername field to given value.

### HasSlackUsername

`func (o *ModelAPIUserSettings) HasSlackUsername() bool`

HasSlackUsername returns a boolean if a field has been set.

### GetSpruceFeedback

`func (o *ModelAPIUserSettings) GetSpruceFeedback() ModelAPIFeedbackSubmission`

GetSpruceFeedback returns the SpruceFeedback field if non-nil, zero value otherwise.

### GetSpruceFeedbackOk

`func (o *ModelAPIUserSettings) GetSpruceFeedbackOk() (*ModelAPIFeedbackSubmission, bool)`

GetSpruceFeedbackOk returns a tuple with the SpruceFeedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpruceFeedback

`func (o *ModelAPIUserSettings) SetSpruceFeedback(v ModelAPIFeedbackSubmission)`

SetSpruceFeedback sets SpruceFeedback field to given value.

### HasSpruceFeedback

`func (o *ModelAPIUserSettings) HasSpruceFeedback() bool`

HasSpruceFeedback returns a boolean if a field has been set.

### GetTimeFormat

`func (o *ModelAPIUserSettings) GetTimeFormat() string`

GetTimeFormat returns the TimeFormat field if non-nil, zero value otherwise.

### GetTimeFormatOk

`func (o *ModelAPIUserSettings) GetTimeFormatOk() (*string, bool)`

GetTimeFormatOk returns a tuple with the TimeFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeFormat

`func (o *ModelAPIUserSettings) SetTimeFormat(v string)`

SetTimeFormat sets TimeFormat field to given value.

### HasTimeFormat

`func (o *ModelAPIUserSettings) HasTimeFormat() bool`

HasTimeFormat returns a boolean if a field has been set.

### GetTimezone

`func (o *ModelAPIUserSettings) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *ModelAPIUserSettings) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *ModelAPIUserSettings) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *ModelAPIUserSettings) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUseSpruceOptions

`func (o *ModelAPIUserSettings) GetUseSpruceOptions() ModelAPIUseSpruceOptions`

GetUseSpruceOptions returns the UseSpruceOptions field if non-nil, zero value otherwise.

### GetUseSpruceOptionsOk

`func (o *ModelAPIUserSettings) GetUseSpruceOptionsOk() (*ModelAPIUseSpruceOptions, bool)`

GetUseSpruceOptionsOk returns a tuple with the UseSpruceOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpruceOptions

`func (o *ModelAPIUserSettings) SetUseSpruceOptions(v ModelAPIUseSpruceOptions)`

SetUseSpruceOptions sets UseSpruceOptions field to given value.

### HasUseSpruceOptions

`func (o *ModelAPIUserSettings) HasUseSpruceOptions() bool`

HasUseSpruceOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


