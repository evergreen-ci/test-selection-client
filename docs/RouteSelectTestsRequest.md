# RouteSelectTestsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildVariant** | Pointer to **string** | BuildVariant is the Evergreen build variant. | [optional] 
**Project** | Pointer to **string** | Project is the project identifier. | [optional] 
**Requester** | Pointer to **string** | Requester is the Evergreen requester type. | [optional] 
**TaskId** | Pointer to **string** | TaskID is the Evergreen task ID. | [optional] 
**TaskName** | Pointer to **string** | TaskName is the Evergreen task name. | [optional] 
**Tests** | Pointer to **[]string** | Tests is a list of test names. | [optional] 

## Methods

### NewRouteSelectTestsRequest

`func NewRouteSelectTestsRequest() *RouteSelectTestsRequest`

NewRouteSelectTestsRequest instantiates a new RouteSelectTestsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouteSelectTestsRequestWithDefaults

`func NewRouteSelectTestsRequestWithDefaults() *RouteSelectTestsRequest`

NewRouteSelectTestsRequestWithDefaults instantiates a new RouteSelectTestsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildVariant

`func (o *RouteSelectTestsRequest) GetBuildVariant() string`

GetBuildVariant returns the BuildVariant field if non-nil, zero value otherwise.

### GetBuildVariantOk

`func (o *RouteSelectTestsRequest) GetBuildVariantOk() (*string, bool)`

GetBuildVariantOk returns a tuple with the BuildVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVariant

`func (o *RouteSelectTestsRequest) SetBuildVariant(v string)`

SetBuildVariant sets BuildVariant field to given value.

### HasBuildVariant

`func (o *RouteSelectTestsRequest) HasBuildVariant() bool`

HasBuildVariant returns a boolean if a field has been set.

### GetProject

`func (o *RouteSelectTestsRequest) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RouteSelectTestsRequest) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RouteSelectTestsRequest) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *RouteSelectTestsRequest) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetRequester

`func (o *RouteSelectTestsRequest) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *RouteSelectTestsRequest) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *RouteSelectTestsRequest) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *RouteSelectTestsRequest) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetTaskId

`func (o *RouteSelectTestsRequest) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *RouteSelectTestsRequest) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *RouteSelectTestsRequest) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *RouteSelectTestsRequest) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskName

`func (o *RouteSelectTestsRequest) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *RouteSelectTestsRequest) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *RouteSelectTestsRequest) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *RouteSelectTestsRequest) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetTests

`func (o *RouteSelectTestsRequest) GetTests() []string`

GetTests returns the Tests field if non-nil, zero value otherwise.

### GetTestsOk

`func (o *RouteSelectTestsRequest) GetTestsOk() (*[]string, bool)`

GetTestsOk returns a tuple with the Tests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTests

`func (o *RouteSelectTestsRequest) SetTests(v []string)`

SetTests sets Tests field to given value.

### HasTests

`func (o *RouteSelectTestsRequest) HasTests() bool`

HasTests returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


