# ManifestModule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Branch** | Pointer to **string** | The branch of the repository. | [optional] 
**Owner** | Pointer to **string** | The owner of the repository. | [optional] 
**Repo** | Pointer to **string** | The name of the repository. | [optional] 
**Revision** | Pointer to **string** | The revision of the head of the branch. | [optional] 
**Url** | Pointer to **string** | The url to the GitHub API call to that specific commit. | [optional] 

## Methods

### NewManifestModule

`func NewManifestModule() *ManifestModule`

NewManifestModule instantiates a new ManifestModule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManifestModuleWithDefaults

`func NewManifestModuleWithDefaults() *ManifestModule`

NewManifestModuleWithDefaults instantiates a new ManifestModule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBranch

`func (o *ManifestModule) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *ManifestModule) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *ManifestModule) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *ManifestModule) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetOwner

`func (o *ManifestModule) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ManifestModule) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ManifestModule) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ManifestModule) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetRepo

`func (o *ManifestModule) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *ManifestModule) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *ManifestModule) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *ManifestModule) HasRepo() bool`

HasRepo returns a boolean if a field has been set.

### GetRevision

`func (o *ManifestModule) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ManifestModule) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ManifestModule) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ManifestModule) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetUrl

`func (o *ManifestModule) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ManifestModule) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ManifestModule) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ManifestModule) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


