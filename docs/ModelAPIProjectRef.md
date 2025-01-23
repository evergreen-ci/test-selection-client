# ModelAPIProjectRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to **[]string** | Usernames of project admins. Can be null for some projects (EVG-6598). | [optional] 
**Aliases** | Pointer to [**[]ModelAPIProjectAlias**](ModelAPIProjectAlias.md) | List of aliases for the project. | [optional] 
**Banner** | Pointer to [**ModelAPIProjectBanner**](ModelAPIProjectBanner.md) | Options for banner to display for the project. | [optional] 
**BatchTime** | Pointer to **int32** | Time interval between commits for Evergreen to activate. | [optional] 
**BranchName** | Pointer to **string** | Name of tracking branch. | [optional] 
**BuildBaronSettings** | Pointer to [**ModelAPIBuildBaronSettings**](ModelAPIBuildBaronSettings.md) | Options for Build Baron. | [optional] 
**CommitQueue** | Pointer to [**ModelAPICommitQueueParams**](ModelAPICommitQueueParams.md) | Options for commit queue. | [optional] 
**ContainerSecrets** | Pointer to [**[]ModelAPIContainerSecret**](ModelAPIContainerSecret.md) | List of container secrets. | [optional] 
**ContainerSizeDefinitions** | Pointer to [**[]ModelAPIContainerResources**](ModelAPIContainerResources.md) | List of container size definitions | [optional] 
**DeactivatePrevious** | Pointer to **bool** | List of identifiers of tasks used in this patch. | [optional] 
**DeleteAdmins** | Pointer to **[]string** | Usernames of project admins to remove. | [optional] 
**DeleteContainerSecrets** | Pointer to **[]string** | Names of container secrets to be deleted. | [optional] 
**DeleteGitTagAuthorizedTeams** | Pointer to **[]string** | Names of GitHub teams authorized to submit git tag versions to remove. | [optional] 
**DeleteGitTagAuthorizedUsers** | Pointer to **[]string** | Usernames of git tag-authorized users to remove. | [optional] 
**DeleteSubscriptions** | Pointer to **[]string** | IDs of subscriptions to delete. | [optional] 
**DisabledStatsCache** | Pointer to **bool** | Disable stats caching. | [optional] 
**DispatchingDisabled** | Pointer to **bool** | Disable task dispatching. | [optional] 
**DisplayName** | Pointer to **string** | Project name displayed to users. | [optional] 
**Enabled** | Pointer to **bool** | Whether evergreen is enabled for this project. | [optional] 
**ExternalLinks** | Pointer to [**[]ModelAPIExternalLink**](ModelAPIExternalLink.md) | List of external links in the version metadata. | [optional] 
**GitTagAuthorizedTeams** | Pointer to **[]string** | Names of GitHub teams authorized to submit git tag versions. | [optional] 
**GitTagAuthorizedUsers** | Pointer to **[]string** | Usernames authorized to submit git tag versions. | [optional] 
**GitTagVersionsEnabled** | Pointer to **bool** | Enable testing when git tags are pushed. | [optional] 
**GithubChecksEnabled** | Pointer to **bool** | Enable GitHub checks. | [optional] 
**GithubDynamicTokenPermissionGroups** | Pointer to [**[]ModelAPIGitHubDynamicTokenPermissionGroup**](ModelAPIGitHubDynamicTokenPermissionGroup.md) | List of GitHub permission groups. | [optional] 
**GithubPermissionGroupByRequester** | Pointer to **map[string]string** | GitHub permission group by requester. | [optional] 
**GithubTriggerAliases** | Pointer to **[]string** | List of GitHub trigger aliases. | [optional] 
**Hidden** | Pointer to **bool** | Whether or not the project can be seen in the UI. Cannot be modified by users. | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** | Internal evergreen identifier for project. | [optional] 
**ManualPrTestingEnabled** | Pointer to **bool** | Enable GitHub manual pull request testing. | [optional] 
**NotifyOnFailure** | Pointer to **bool** | Notify original committer (or admins) when build fails. | [optional] 
**OldestAllowedMergeBase** | Pointer to **string** | Oldest allowed merge base for PR patches | [optional] 
**OwnerName** | Pointer to **string** | GitHub org name. | [optional] 
**ParsleyFilters** | Pointer to [**[]ModelAPIParsleyFilter**](ModelAPIParsleyFilter.md) | List of custom Parsley filters. | [optional] 
**PatchTriggerAliases** | Pointer to [**[]ModelAPIPatchTriggerDefinition**](ModelAPIPatchTriggerDefinition.md) | List of patch trigger aliases. | [optional] 
**PatchingDisabled** | Pointer to **bool** | Disable patching. | [optional] 
**PerfEnabled** | Pointer to **bool** | Enable the performance plugin. | [optional] 
**PeriodicBuilds** | Pointer to [**[]ModelAPIPeriodicBuildDefinition**](ModelAPIPeriodicBuildDefinition.md) | List of periodic build definitions. | [optional] 
**PrTestingEnabled** | Pointer to **bool** | Enable GitHub automated pull request testing. | [optional] 
**ProjectHealthView** | Pointer to [**ModelProjectHealthView**](ModelProjectHealthView.md) | Default project health view. | [optional] 
**RemotePath** | Pointer to **string** | Path to config file in repo. | [optional] 
**RepoName** | Pointer to **string** | GitHub repository name. | [optional] 
**RepoRefId** | Pointer to **string** | Identifier of the attached repo ref. Cannot be modified by users. | [optional] 
**RepotrackerDisabled** | Pointer to **bool** | Disable the repotracker. | [optional] 
**RepotrackerError** | Pointer to [**ModelAPIRepositoryErrorDetails**](ModelAPIRepositoryErrorDetails.md) | Error from the repotracker, if any. Cannot be modified by users. | [optional] 
**Restricted** | Pointer to **bool** | Prevent users from being able to view this project unless explicitly granted access. | [optional] 
**Revision** | Pointer to **string** | Only used when modifying projects to change the base revision and run the repotracker. | [optional] 
**SpawnHostScriptPath** | Pointer to **string** | File path to script that users can run on spawn hosts loaded with task data. | [optional] 
**StepbackBisect** | Pointer to **bool** | Use bisect stepback instead of linear. | [optional] 
**StepbackDisabled** | Pointer to **bool** | Disable stepback. | [optional] 
**Subscriptions** | Pointer to [**[]ModelAPISubscription**](ModelAPISubscription.md) | List of subscriptions for the project. | [optional] 
**TaskAnnotationSettings** | Pointer to [**ModelAPITaskAnnotationSettings**](ModelAPITaskAnnotationSettings.md) | Options for task annotations. | [optional] 
**TracksPushEvents** | Pointer to **bool** | If true, repotracker is run on github push events. If false, repotracker is run periodically every few minutes. | [optional] 
**Triggers** | Pointer to [**[]ModelAPITriggerDefinition**](ModelAPITriggerDefinition.md) | List of triggers for the project. | [optional] 
**UseRepoSettings** | Pointer to **bool** | Whether or not to default to using repo settings. | [optional] 
**Variables** | Pointer to [**ModelAPIProjectVars**](ModelAPIProjectVars.md) | Project variables information | [optional] 
**VersionControlEnabled** | Pointer to **bool** | Enable setting project aliases from version-controlled project configs. | [optional] 
**WorkstationConfig** | Pointer to [**ModelAPIWorkstationConfig**](ModelAPIWorkstationConfig.md) | Options for workstations. | [optional] 

## Methods

### NewModelAPIProjectRef

`func NewModelAPIProjectRef() *ModelAPIProjectRef`

NewModelAPIProjectRef instantiates a new ModelAPIProjectRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelAPIProjectRefWithDefaults

`func NewModelAPIProjectRefWithDefaults() *ModelAPIProjectRef`

NewModelAPIProjectRefWithDefaults instantiates a new ModelAPIProjectRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *ModelAPIProjectRef) GetAdmins() []string`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *ModelAPIProjectRef) GetAdminsOk() (*[]string, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *ModelAPIProjectRef) SetAdmins(v []string)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *ModelAPIProjectRef) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetAliases

`func (o *ModelAPIProjectRef) GetAliases() []ModelAPIProjectAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *ModelAPIProjectRef) GetAliasesOk() (*[]ModelAPIProjectAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *ModelAPIProjectRef) SetAliases(v []ModelAPIProjectAlias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *ModelAPIProjectRef) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetBanner

`func (o *ModelAPIProjectRef) GetBanner() ModelAPIProjectBanner`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *ModelAPIProjectRef) GetBannerOk() (*ModelAPIProjectBanner, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *ModelAPIProjectRef) SetBanner(v ModelAPIProjectBanner)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *ModelAPIProjectRef) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### GetBatchTime

`func (o *ModelAPIProjectRef) GetBatchTime() int32`

GetBatchTime returns the BatchTime field if non-nil, zero value otherwise.

### GetBatchTimeOk

`func (o *ModelAPIProjectRef) GetBatchTimeOk() (*int32, bool)`

GetBatchTimeOk returns a tuple with the BatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchTime

`func (o *ModelAPIProjectRef) SetBatchTime(v int32)`

SetBatchTime sets BatchTime field to given value.

### HasBatchTime

`func (o *ModelAPIProjectRef) HasBatchTime() bool`

HasBatchTime returns a boolean if a field has been set.

### GetBranchName

`func (o *ModelAPIProjectRef) GetBranchName() string`

GetBranchName returns the BranchName field if non-nil, zero value otherwise.

### GetBranchNameOk

`func (o *ModelAPIProjectRef) GetBranchNameOk() (*string, bool)`

GetBranchNameOk returns a tuple with the BranchName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranchName

`func (o *ModelAPIProjectRef) SetBranchName(v string)`

SetBranchName sets BranchName field to given value.

### HasBranchName

`func (o *ModelAPIProjectRef) HasBranchName() bool`

HasBranchName returns a boolean if a field has been set.

### GetBuildBaronSettings

`func (o *ModelAPIProjectRef) GetBuildBaronSettings() ModelAPIBuildBaronSettings`

GetBuildBaronSettings returns the BuildBaronSettings field if non-nil, zero value otherwise.

### GetBuildBaronSettingsOk

`func (o *ModelAPIProjectRef) GetBuildBaronSettingsOk() (*ModelAPIBuildBaronSettings, bool)`

GetBuildBaronSettingsOk returns a tuple with the BuildBaronSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildBaronSettings

`func (o *ModelAPIProjectRef) SetBuildBaronSettings(v ModelAPIBuildBaronSettings)`

SetBuildBaronSettings sets BuildBaronSettings field to given value.

### HasBuildBaronSettings

`func (o *ModelAPIProjectRef) HasBuildBaronSettings() bool`

HasBuildBaronSettings returns a boolean if a field has been set.

### GetCommitQueue

`func (o *ModelAPIProjectRef) GetCommitQueue() ModelAPICommitQueueParams`

GetCommitQueue returns the CommitQueue field if non-nil, zero value otherwise.

### GetCommitQueueOk

`func (o *ModelAPIProjectRef) GetCommitQueueOk() (*ModelAPICommitQueueParams, bool)`

GetCommitQueueOk returns a tuple with the CommitQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommitQueue

`func (o *ModelAPIProjectRef) SetCommitQueue(v ModelAPICommitQueueParams)`

SetCommitQueue sets CommitQueue field to given value.

### HasCommitQueue

`func (o *ModelAPIProjectRef) HasCommitQueue() bool`

HasCommitQueue returns a boolean if a field has been set.

### GetContainerSecrets

`func (o *ModelAPIProjectRef) GetContainerSecrets() []ModelAPIContainerSecret`

GetContainerSecrets returns the ContainerSecrets field if non-nil, zero value otherwise.

### GetContainerSecretsOk

`func (o *ModelAPIProjectRef) GetContainerSecretsOk() (*[]ModelAPIContainerSecret, bool)`

GetContainerSecretsOk returns a tuple with the ContainerSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSecrets

`func (o *ModelAPIProjectRef) SetContainerSecrets(v []ModelAPIContainerSecret)`

SetContainerSecrets sets ContainerSecrets field to given value.

### HasContainerSecrets

`func (o *ModelAPIProjectRef) HasContainerSecrets() bool`

HasContainerSecrets returns a boolean if a field has been set.

### GetContainerSizeDefinitions

`func (o *ModelAPIProjectRef) GetContainerSizeDefinitions() []ModelAPIContainerResources`

GetContainerSizeDefinitions returns the ContainerSizeDefinitions field if non-nil, zero value otherwise.

### GetContainerSizeDefinitionsOk

`func (o *ModelAPIProjectRef) GetContainerSizeDefinitionsOk() (*[]ModelAPIContainerResources, bool)`

GetContainerSizeDefinitionsOk returns a tuple with the ContainerSizeDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerSizeDefinitions

`func (o *ModelAPIProjectRef) SetContainerSizeDefinitions(v []ModelAPIContainerResources)`

SetContainerSizeDefinitions sets ContainerSizeDefinitions field to given value.

### HasContainerSizeDefinitions

`func (o *ModelAPIProjectRef) HasContainerSizeDefinitions() bool`

HasContainerSizeDefinitions returns a boolean if a field has been set.

### GetDeactivatePrevious

`func (o *ModelAPIProjectRef) GetDeactivatePrevious() bool`

GetDeactivatePrevious returns the DeactivatePrevious field if non-nil, zero value otherwise.

### GetDeactivatePreviousOk

`func (o *ModelAPIProjectRef) GetDeactivatePreviousOk() (*bool, bool)`

GetDeactivatePreviousOk returns a tuple with the DeactivatePrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeactivatePrevious

`func (o *ModelAPIProjectRef) SetDeactivatePrevious(v bool)`

SetDeactivatePrevious sets DeactivatePrevious field to given value.

### HasDeactivatePrevious

`func (o *ModelAPIProjectRef) HasDeactivatePrevious() bool`

HasDeactivatePrevious returns a boolean if a field has been set.

### GetDeleteAdmins

`func (o *ModelAPIProjectRef) GetDeleteAdmins() []string`

GetDeleteAdmins returns the DeleteAdmins field if non-nil, zero value otherwise.

### GetDeleteAdminsOk

`func (o *ModelAPIProjectRef) GetDeleteAdminsOk() (*[]string, bool)`

GetDeleteAdminsOk returns a tuple with the DeleteAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteAdmins

`func (o *ModelAPIProjectRef) SetDeleteAdmins(v []string)`

SetDeleteAdmins sets DeleteAdmins field to given value.

### HasDeleteAdmins

`func (o *ModelAPIProjectRef) HasDeleteAdmins() bool`

HasDeleteAdmins returns a boolean if a field has been set.

### GetDeleteContainerSecrets

`func (o *ModelAPIProjectRef) GetDeleteContainerSecrets() []string`

GetDeleteContainerSecrets returns the DeleteContainerSecrets field if non-nil, zero value otherwise.

### GetDeleteContainerSecretsOk

`func (o *ModelAPIProjectRef) GetDeleteContainerSecretsOk() (*[]string, bool)`

GetDeleteContainerSecretsOk returns a tuple with the DeleteContainerSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteContainerSecrets

`func (o *ModelAPIProjectRef) SetDeleteContainerSecrets(v []string)`

SetDeleteContainerSecrets sets DeleteContainerSecrets field to given value.

### HasDeleteContainerSecrets

`func (o *ModelAPIProjectRef) HasDeleteContainerSecrets() bool`

HasDeleteContainerSecrets returns a boolean if a field has been set.

### GetDeleteGitTagAuthorizedTeams

`func (o *ModelAPIProjectRef) GetDeleteGitTagAuthorizedTeams() []string`

GetDeleteGitTagAuthorizedTeams returns the DeleteGitTagAuthorizedTeams field if non-nil, zero value otherwise.

### GetDeleteGitTagAuthorizedTeamsOk

`func (o *ModelAPIProjectRef) GetDeleteGitTagAuthorizedTeamsOk() (*[]string, bool)`

GetDeleteGitTagAuthorizedTeamsOk returns a tuple with the DeleteGitTagAuthorizedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteGitTagAuthorizedTeams

`func (o *ModelAPIProjectRef) SetDeleteGitTagAuthorizedTeams(v []string)`

SetDeleteGitTagAuthorizedTeams sets DeleteGitTagAuthorizedTeams field to given value.

### HasDeleteGitTagAuthorizedTeams

`func (o *ModelAPIProjectRef) HasDeleteGitTagAuthorizedTeams() bool`

HasDeleteGitTagAuthorizedTeams returns a boolean if a field has been set.

### GetDeleteGitTagAuthorizedUsers

`func (o *ModelAPIProjectRef) GetDeleteGitTagAuthorizedUsers() []string`

GetDeleteGitTagAuthorizedUsers returns the DeleteGitTagAuthorizedUsers field if non-nil, zero value otherwise.

### GetDeleteGitTagAuthorizedUsersOk

`func (o *ModelAPIProjectRef) GetDeleteGitTagAuthorizedUsersOk() (*[]string, bool)`

GetDeleteGitTagAuthorizedUsersOk returns a tuple with the DeleteGitTagAuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteGitTagAuthorizedUsers

`func (o *ModelAPIProjectRef) SetDeleteGitTagAuthorizedUsers(v []string)`

SetDeleteGitTagAuthorizedUsers sets DeleteGitTagAuthorizedUsers field to given value.

### HasDeleteGitTagAuthorizedUsers

`func (o *ModelAPIProjectRef) HasDeleteGitTagAuthorizedUsers() bool`

HasDeleteGitTagAuthorizedUsers returns a boolean if a field has been set.

### GetDeleteSubscriptions

`func (o *ModelAPIProjectRef) GetDeleteSubscriptions() []string`

GetDeleteSubscriptions returns the DeleteSubscriptions field if non-nil, zero value otherwise.

### GetDeleteSubscriptionsOk

`func (o *ModelAPIProjectRef) GetDeleteSubscriptionsOk() (*[]string, bool)`

GetDeleteSubscriptionsOk returns a tuple with the DeleteSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteSubscriptions

`func (o *ModelAPIProjectRef) SetDeleteSubscriptions(v []string)`

SetDeleteSubscriptions sets DeleteSubscriptions field to given value.

### HasDeleteSubscriptions

`func (o *ModelAPIProjectRef) HasDeleteSubscriptions() bool`

HasDeleteSubscriptions returns a boolean if a field has been set.

### GetDisabledStatsCache

`func (o *ModelAPIProjectRef) GetDisabledStatsCache() bool`

GetDisabledStatsCache returns the DisabledStatsCache field if non-nil, zero value otherwise.

### GetDisabledStatsCacheOk

`func (o *ModelAPIProjectRef) GetDisabledStatsCacheOk() (*bool, bool)`

GetDisabledStatsCacheOk returns a tuple with the DisabledStatsCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabledStatsCache

`func (o *ModelAPIProjectRef) SetDisabledStatsCache(v bool)`

SetDisabledStatsCache sets DisabledStatsCache field to given value.

### HasDisabledStatsCache

`func (o *ModelAPIProjectRef) HasDisabledStatsCache() bool`

HasDisabledStatsCache returns a boolean if a field has been set.

### GetDispatchingDisabled

`func (o *ModelAPIProjectRef) GetDispatchingDisabled() bool`

GetDispatchingDisabled returns the DispatchingDisabled field if non-nil, zero value otherwise.

### GetDispatchingDisabledOk

`func (o *ModelAPIProjectRef) GetDispatchingDisabledOk() (*bool, bool)`

GetDispatchingDisabledOk returns a tuple with the DispatchingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchingDisabled

`func (o *ModelAPIProjectRef) SetDispatchingDisabled(v bool)`

SetDispatchingDisabled sets DispatchingDisabled field to given value.

### HasDispatchingDisabled

`func (o *ModelAPIProjectRef) HasDispatchingDisabled() bool`

HasDispatchingDisabled returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelAPIProjectRef) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelAPIProjectRef) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelAPIProjectRef) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelAPIProjectRef) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *ModelAPIProjectRef) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelAPIProjectRef) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelAPIProjectRef) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelAPIProjectRef) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalLinks

`func (o *ModelAPIProjectRef) GetExternalLinks() []ModelAPIExternalLink`

GetExternalLinks returns the ExternalLinks field if non-nil, zero value otherwise.

### GetExternalLinksOk

`func (o *ModelAPIProjectRef) GetExternalLinksOk() (*[]ModelAPIExternalLink, bool)`

GetExternalLinksOk returns a tuple with the ExternalLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalLinks

`func (o *ModelAPIProjectRef) SetExternalLinks(v []ModelAPIExternalLink)`

SetExternalLinks sets ExternalLinks field to given value.

### HasExternalLinks

`func (o *ModelAPIProjectRef) HasExternalLinks() bool`

HasExternalLinks returns a boolean if a field has been set.

### GetGitTagAuthorizedTeams

`func (o *ModelAPIProjectRef) GetGitTagAuthorizedTeams() []string`

GetGitTagAuthorizedTeams returns the GitTagAuthorizedTeams field if non-nil, zero value otherwise.

### GetGitTagAuthorizedTeamsOk

`func (o *ModelAPIProjectRef) GetGitTagAuthorizedTeamsOk() (*[]string, bool)`

GetGitTagAuthorizedTeamsOk returns a tuple with the GitTagAuthorizedTeams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagAuthorizedTeams

`func (o *ModelAPIProjectRef) SetGitTagAuthorizedTeams(v []string)`

SetGitTagAuthorizedTeams sets GitTagAuthorizedTeams field to given value.

### HasGitTagAuthorizedTeams

`func (o *ModelAPIProjectRef) HasGitTagAuthorizedTeams() bool`

HasGitTagAuthorizedTeams returns a boolean if a field has been set.

### GetGitTagAuthorizedUsers

`func (o *ModelAPIProjectRef) GetGitTagAuthorizedUsers() []string`

GetGitTagAuthorizedUsers returns the GitTagAuthorizedUsers field if non-nil, zero value otherwise.

### GetGitTagAuthorizedUsersOk

`func (o *ModelAPIProjectRef) GetGitTagAuthorizedUsersOk() (*[]string, bool)`

GetGitTagAuthorizedUsersOk returns a tuple with the GitTagAuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagAuthorizedUsers

`func (o *ModelAPIProjectRef) SetGitTagAuthorizedUsers(v []string)`

SetGitTagAuthorizedUsers sets GitTagAuthorizedUsers field to given value.

### HasGitTagAuthorizedUsers

`func (o *ModelAPIProjectRef) HasGitTagAuthorizedUsers() bool`

HasGitTagAuthorizedUsers returns a boolean if a field has been set.

### GetGitTagVersionsEnabled

`func (o *ModelAPIProjectRef) GetGitTagVersionsEnabled() bool`

GetGitTagVersionsEnabled returns the GitTagVersionsEnabled field if non-nil, zero value otherwise.

### GetGitTagVersionsEnabledOk

`func (o *ModelAPIProjectRef) GetGitTagVersionsEnabledOk() (*bool, bool)`

GetGitTagVersionsEnabledOk returns a tuple with the GitTagVersionsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitTagVersionsEnabled

`func (o *ModelAPIProjectRef) SetGitTagVersionsEnabled(v bool)`

SetGitTagVersionsEnabled sets GitTagVersionsEnabled field to given value.

### HasGitTagVersionsEnabled

`func (o *ModelAPIProjectRef) HasGitTagVersionsEnabled() bool`

HasGitTagVersionsEnabled returns a boolean if a field has been set.

### GetGithubChecksEnabled

`func (o *ModelAPIProjectRef) GetGithubChecksEnabled() bool`

GetGithubChecksEnabled returns the GithubChecksEnabled field if non-nil, zero value otherwise.

### GetGithubChecksEnabledOk

`func (o *ModelAPIProjectRef) GetGithubChecksEnabledOk() (*bool, bool)`

GetGithubChecksEnabledOk returns a tuple with the GithubChecksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubChecksEnabled

`func (o *ModelAPIProjectRef) SetGithubChecksEnabled(v bool)`

SetGithubChecksEnabled sets GithubChecksEnabled field to given value.

### HasGithubChecksEnabled

`func (o *ModelAPIProjectRef) HasGithubChecksEnabled() bool`

HasGithubChecksEnabled returns a boolean if a field has been set.

### GetGithubDynamicTokenPermissionGroups

`func (o *ModelAPIProjectRef) GetGithubDynamicTokenPermissionGroups() []ModelAPIGitHubDynamicTokenPermissionGroup`

GetGithubDynamicTokenPermissionGroups returns the GithubDynamicTokenPermissionGroups field if non-nil, zero value otherwise.

### GetGithubDynamicTokenPermissionGroupsOk

`func (o *ModelAPIProjectRef) GetGithubDynamicTokenPermissionGroupsOk() (*[]ModelAPIGitHubDynamicTokenPermissionGroup, bool)`

GetGithubDynamicTokenPermissionGroupsOk returns a tuple with the GithubDynamicTokenPermissionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubDynamicTokenPermissionGroups

`func (o *ModelAPIProjectRef) SetGithubDynamicTokenPermissionGroups(v []ModelAPIGitHubDynamicTokenPermissionGroup)`

SetGithubDynamicTokenPermissionGroups sets GithubDynamicTokenPermissionGroups field to given value.

### HasGithubDynamicTokenPermissionGroups

`func (o *ModelAPIProjectRef) HasGithubDynamicTokenPermissionGroups() bool`

HasGithubDynamicTokenPermissionGroups returns a boolean if a field has been set.

### GetGithubPermissionGroupByRequester

`func (o *ModelAPIProjectRef) GetGithubPermissionGroupByRequester() map[string]string`

GetGithubPermissionGroupByRequester returns the GithubPermissionGroupByRequester field if non-nil, zero value otherwise.

### GetGithubPermissionGroupByRequesterOk

`func (o *ModelAPIProjectRef) GetGithubPermissionGroupByRequesterOk() (*map[string]string, bool)`

GetGithubPermissionGroupByRequesterOk returns a tuple with the GithubPermissionGroupByRequester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubPermissionGroupByRequester

`func (o *ModelAPIProjectRef) SetGithubPermissionGroupByRequester(v map[string]string)`

SetGithubPermissionGroupByRequester sets GithubPermissionGroupByRequester field to given value.

### HasGithubPermissionGroupByRequester

`func (o *ModelAPIProjectRef) HasGithubPermissionGroupByRequester() bool`

HasGithubPermissionGroupByRequester returns a boolean if a field has been set.

### GetGithubTriggerAliases

`func (o *ModelAPIProjectRef) GetGithubTriggerAliases() []string`

GetGithubTriggerAliases returns the GithubTriggerAliases field if non-nil, zero value otherwise.

### GetGithubTriggerAliasesOk

`func (o *ModelAPIProjectRef) GetGithubTriggerAliasesOk() (*[]string, bool)`

GetGithubTriggerAliasesOk returns a tuple with the GithubTriggerAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGithubTriggerAliases

`func (o *ModelAPIProjectRef) SetGithubTriggerAliases(v []string)`

SetGithubTriggerAliases sets GithubTriggerAliases field to given value.

### HasGithubTriggerAliases

`func (o *ModelAPIProjectRef) HasGithubTriggerAliases() bool`

HasGithubTriggerAliases returns a boolean if a field has been set.

### GetHidden

`func (o *ModelAPIProjectRef) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ModelAPIProjectRef) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ModelAPIProjectRef) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ModelAPIProjectRef) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetId

`func (o *ModelAPIProjectRef) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelAPIProjectRef) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelAPIProjectRef) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelAPIProjectRef) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentifier

`func (o *ModelAPIProjectRef) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ModelAPIProjectRef) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ModelAPIProjectRef) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *ModelAPIProjectRef) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetManualPrTestingEnabled

`func (o *ModelAPIProjectRef) GetManualPrTestingEnabled() bool`

GetManualPrTestingEnabled returns the ManualPrTestingEnabled field if non-nil, zero value otherwise.

### GetManualPrTestingEnabledOk

`func (o *ModelAPIProjectRef) GetManualPrTestingEnabledOk() (*bool, bool)`

GetManualPrTestingEnabledOk returns a tuple with the ManualPrTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManualPrTestingEnabled

`func (o *ModelAPIProjectRef) SetManualPrTestingEnabled(v bool)`

SetManualPrTestingEnabled sets ManualPrTestingEnabled field to given value.

### HasManualPrTestingEnabled

`func (o *ModelAPIProjectRef) HasManualPrTestingEnabled() bool`

HasManualPrTestingEnabled returns a boolean if a field has been set.

### GetNotifyOnFailure

`func (o *ModelAPIProjectRef) GetNotifyOnFailure() bool`

GetNotifyOnFailure returns the NotifyOnFailure field if non-nil, zero value otherwise.

### GetNotifyOnFailureOk

`func (o *ModelAPIProjectRef) GetNotifyOnFailureOk() (*bool, bool)`

GetNotifyOnFailureOk returns a tuple with the NotifyOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyOnFailure

`func (o *ModelAPIProjectRef) SetNotifyOnFailure(v bool)`

SetNotifyOnFailure sets NotifyOnFailure field to given value.

### HasNotifyOnFailure

`func (o *ModelAPIProjectRef) HasNotifyOnFailure() bool`

HasNotifyOnFailure returns a boolean if a field has been set.

### GetOldestAllowedMergeBase

`func (o *ModelAPIProjectRef) GetOldestAllowedMergeBase() string`

GetOldestAllowedMergeBase returns the OldestAllowedMergeBase field if non-nil, zero value otherwise.

### GetOldestAllowedMergeBaseOk

`func (o *ModelAPIProjectRef) GetOldestAllowedMergeBaseOk() (*string, bool)`

GetOldestAllowedMergeBaseOk returns a tuple with the OldestAllowedMergeBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestAllowedMergeBase

`func (o *ModelAPIProjectRef) SetOldestAllowedMergeBase(v string)`

SetOldestAllowedMergeBase sets OldestAllowedMergeBase field to given value.

### HasOldestAllowedMergeBase

`func (o *ModelAPIProjectRef) HasOldestAllowedMergeBase() bool`

HasOldestAllowedMergeBase returns a boolean if a field has been set.

### GetOwnerName

`func (o *ModelAPIProjectRef) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *ModelAPIProjectRef) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *ModelAPIProjectRef) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *ModelAPIProjectRef) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetParsleyFilters

`func (o *ModelAPIProjectRef) GetParsleyFilters() []ModelAPIParsleyFilter`

GetParsleyFilters returns the ParsleyFilters field if non-nil, zero value otherwise.

### GetParsleyFiltersOk

`func (o *ModelAPIProjectRef) GetParsleyFiltersOk() (*[]ModelAPIParsleyFilter, bool)`

GetParsleyFiltersOk returns a tuple with the ParsleyFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsleyFilters

`func (o *ModelAPIProjectRef) SetParsleyFilters(v []ModelAPIParsleyFilter)`

SetParsleyFilters sets ParsleyFilters field to given value.

### HasParsleyFilters

`func (o *ModelAPIProjectRef) HasParsleyFilters() bool`

HasParsleyFilters returns a boolean if a field has been set.

### GetPatchTriggerAliases

`func (o *ModelAPIProjectRef) GetPatchTriggerAliases() []ModelAPIPatchTriggerDefinition`

GetPatchTriggerAliases returns the PatchTriggerAliases field if non-nil, zero value otherwise.

### GetPatchTriggerAliasesOk

`func (o *ModelAPIProjectRef) GetPatchTriggerAliasesOk() (*[]ModelAPIPatchTriggerDefinition, bool)`

GetPatchTriggerAliasesOk returns a tuple with the PatchTriggerAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchTriggerAliases

`func (o *ModelAPIProjectRef) SetPatchTriggerAliases(v []ModelAPIPatchTriggerDefinition)`

SetPatchTriggerAliases sets PatchTriggerAliases field to given value.

### HasPatchTriggerAliases

`func (o *ModelAPIProjectRef) HasPatchTriggerAliases() bool`

HasPatchTriggerAliases returns a boolean if a field has been set.

### GetPatchingDisabled

`func (o *ModelAPIProjectRef) GetPatchingDisabled() bool`

GetPatchingDisabled returns the PatchingDisabled field if non-nil, zero value otherwise.

### GetPatchingDisabledOk

`func (o *ModelAPIProjectRef) GetPatchingDisabledOk() (*bool, bool)`

GetPatchingDisabledOk returns a tuple with the PatchingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatchingDisabled

`func (o *ModelAPIProjectRef) SetPatchingDisabled(v bool)`

SetPatchingDisabled sets PatchingDisabled field to given value.

### HasPatchingDisabled

`func (o *ModelAPIProjectRef) HasPatchingDisabled() bool`

HasPatchingDisabled returns a boolean if a field has been set.

### GetPerfEnabled

`func (o *ModelAPIProjectRef) GetPerfEnabled() bool`

GetPerfEnabled returns the PerfEnabled field if non-nil, zero value otherwise.

### GetPerfEnabledOk

`func (o *ModelAPIProjectRef) GetPerfEnabledOk() (*bool, bool)`

GetPerfEnabledOk returns a tuple with the PerfEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfEnabled

`func (o *ModelAPIProjectRef) SetPerfEnabled(v bool)`

SetPerfEnabled sets PerfEnabled field to given value.

### HasPerfEnabled

`func (o *ModelAPIProjectRef) HasPerfEnabled() bool`

HasPerfEnabled returns a boolean if a field has been set.

### GetPeriodicBuilds

`func (o *ModelAPIProjectRef) GetPeriodicBuilds() []ModelAPIPeriodicBuildDefinition`

GetPeriodicBuilds returns the PeriodicBuilds field if non-nil, zero value otherwise.

### GetPeriodicBuildsOk

`func (o *ModelAPIProjectRef) GetPeriodicBuildsOk() (*[]ModelAPIPeriodicBuildDefinition, bool)`

GetPeriodicBuildsOk returns a tuple with the PeriodicBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodicBuilds

`func (o *ModelAPIProjectRef) SetPeriodicBuilds(v []ModelAPIPeriodicBuildDefinition)`

SetPeriodicBuilds sets PeriodicBuilds field to given value.

### HasPeriodicBuilds

`func (o *ModelAPIProjectRef) HasPeriodicBuilds() bool`

HasPeriodicBuilds returns a boolean if a field has been set.

### GetPrTestingEnabled

`func (o *ModelAPIProjectRef) GetPrTestingEnabled() bool`

GetPrTestingEnabled returns the PrTestingEnabled field if non-nil, zero value otherwise.

### GetPrTestingEnabledOk

`func (o *ModelAPIProjectRef) GetPrTestingEnabledOk() (*bool, bool)`

GetPrTestingEnabledOk returns a tuple with the PrTestingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrTestingEnabled

`func (o *ModelAPIProjectRef) SetPrTestingEnabled(v bool)`

SetPrTestingEnabled sets PrTestingEnabled field to given value.

### HasPrTestingEnabled

`func (o *ModelAPIProjectRef) HasPrTestingEnabled() bool`

HasPrTestingEnabled returns a boolean if a field has been set.

### GetProjectHealthView

`func (o *ModelAPIProjectRef) GetProjectHealthView() ModelProjectHealthView`

GetProjectHealthView returns the ProjectHealthView field if non-nil, zero value otherwise.

### GetProjectHealthViewOk

`func (o *ModelAPIProjectRef) GetProjectHealthViewOk() (*ModelProjectHealthView, bool)`

GetProjectHealthViewOk returns a tuple with the ProjectHealthView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectHealthView

`func (o *ModelAPIProjectRef) SetProjectHealthView(v ModelProjectHealthView)`

SetProjectHealthView sets ProjectHealthView field to given value.

### HasProjectHealthView

`func (o *ModelAPIProjectRef) HasProjectHealthView() bool`

HasProjectHealthView returns a boolean if a field has been set.

### GetRemotePath

`func (o *ModelAPIProjectRef) GetRemotePath() string`

GetRemotePath returns the RemotePath field if non-nil, zero value otherwise.

### GetRemotePathOk

`func (o *ModelAPIProjectRef) GetRemotePathOk() (*string, bool)`

GetRemotePathOk returns a tuple with the RemotePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemotePath

`func (o *ModelAPIProjectRef) SetRemotePath(v string)`

SetRemotePath sets RemotePath field to given value.

### HasRemotePath

`func (o *ModelAPIProjectRef) HasRemotePath() bool`

HasRemotePath returns a boolean if a field has been set.

### GetRepoName

`func (o *ModelAPIProjectRef) GetRepoName() string`

GetRepoName returns the RepoName field if non-nil, zero value otherwise.

### GetRepoNameOk

`func (o *ModelAPIProjectRef) GetRepoNameOk() (*string, bool)`

GetRepoNameOk returns a tuple with the RepoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoName

`func (o *ModelAPIProjectRef) SetRepoName(v string)`

SetRepoName sets RepoName field to given value.

### HasRepoName

`func (o *ModelAPIProjectRef) HasRepoName() bool`

HasRepoName returns a boolean if a field has been set.

### GetRepoRefId

`func (o *ModelAPIProjectRef) GetRepoRefId() string`

GetRepoRefId returns the RepoRefId field if non-nil, zero value otherwise.

### GetRepoRefIdOk

`func (o *ModelAPIProjectRef) GetRepoRefIdOk() (*string, bool)`

GetRepoRefIdOk returns a tuple with the RepoRefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoRefId

`func (o *ModelAPIProjectRef) SetRepoRefId(v string)`

SetRepoRefId sets RepoRefId field to given value.

### HasRepoRefId

`func (o *ModelAPIProjectRef) HasRepoRefId() bool`

HasRepoRefId returns a boolean if a field has been set.

### GetRepotrackerDisabled

`func (o *ModelAPIProjectRef) GetRepotrackerDisabled() bool`

GetRepotrackerDisabled returns the RepotrackerDisabled field if non-nil, zero value otherwise.

### GetRepotrackerDisabledOk

`func (o *ModelAPIProjectRef) GetRepotrackerDisabledOk() (*bool, bool)`

GetRepotrackerDisabledOk returns a tuple with the RepotrackerDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepotrackerDisabled

`func (o *ModelAPIProjectRef) SetRepotrackerDisabled(v bool)`

SetRepotrackerDisabled sets RepotrackerDisabled field to given value.

### HasRepotrackerDisabled

`func (o *ModelAPIProjectRef) HasRepotrackerDisabled() bool`

HasRepotrackerDisabled returns a boolean if a field has been set.

### GetRepotrackerError

`func (o *ModelAPIProjectRef) GetRepotrackerError() ModelAPIRepositoryErrorDetails`

GetRepotrackerError returns the RepotrackerError field if non-nil, zero value otherwise.

### GetRepotrackerErrorOk

`func (o *ModelAPIProjectRef) GetRepotrackerErrorOk() (*ModelAPIRepositoryErrorDetails, bool)`

GetRepotrackerErrorOk returns a tuple with the RepotrackerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepotrackerError

`func (o *ModelAPIProjectRef) SetRepotrackerError(v ModelAPIRepositoryErrorDetails)`

SetRepotrackerError sets RepotrackerError field to given value.

### HasRepotrackerError

`func (o *ModelAPIProjectRef) HasRepotrackerError() bool`

HasRepotrackerError returns a boolean if a field has been set.

### GetRestricted

`func (o *ModelAPIProjectRef) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *ModelAPIProjectRef) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *ModelAPIProjectRef) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *ModelAPIProjectRef) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetRevision

`func (o *ModelAPIProjectRef) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ModelAPIProjectRef) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ModelAPIProjectRef) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ModelAPIProjectRef) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetSpawnHostScriptPath

`func (o *ModelAPIProjectRef) GetSpawnHostScriptPath() string`

GetSpawnHostScriptPath returns the SpawnHostScriptPath field if non-nil, zero value otherwise.

### GetSpawnHostScriptPathOk

`func (o *ModelAPIProjectRef) GetSpawnHostScriptPathOk() (*string, bool)`

GetSpawnHostScriptPathOk returns a tuple with the SpawnHostScriptPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpawnHostScriptPath

`func (o *ModelAPIProjectRef) SetSpawnHostScriptPath(v string)`

SetSpawnHostScriptPath sets SpawnHostScriptPath field to given value.

### HasSpawnHostScriptPath

`func (o *ModelAPIProjectRef) HasSpawnHostScriptPath() bool`

HasSpawnHostScriptPath returns a boolean if a field has been set.

### GetStepbackBisect

`func (o *ModelAPIProjectRef) GetStepbackBisect() bool`

GetStepbackBisect returns the StepbackBisect field if non-nil, zero value otherwise.

### GetStepbackBisectOk

`func (o *ModelAPIProjectRef) GetStepbackBisectOk() (*bool, bool)`

GetStepbackBisectOk returns a tuple with the StepbackBisect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepbackBisect

`func (o *ModelAPIProjectRef) SetStepbackBisect(v bool)`

SetStepbackBisect sets StepbackBisect field to given value.

### HasStepbackBisect

`func (o *ModelAPIProjectRef) HasStepbackBisect() bool`

HasStepbackBisect returns a boolean if a field has been set.

### GetStepbackDisabled

`func (o *ModelAPIProjectRef) GetStepbackDisabled() bool`

GetStepbackDisabled returns the StepbackDisabled field if non-nil, zero value otherwise.

### GetStepbackDisabledOk

`func (o *ModelAPIProjectRef) GetStepbackDisabledOk() (*bool, bool)`

GetStepbackDisabledOk returns a tuple with the StepbackDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepbackDisabled

`func (o *ModelAPIProjectRef) SetStepbackDisabled(v bool)`

SetStepbackDisabled sets StepbackDisabled field to given value.

### HasStepbackDisabled

`func (o *ModelAPIProjectRef) HasStepbackDisabled() bool`

HasStepbackDisabled returns a boolean if a field has been set.

### GetSubscriptions

`func (o *ModelAPIProjectRef) GetSubscriptions() []ModelAPISubscription`

GetSubscriptions returns the Subscriptions field if non-nil, zero value otherwise.

### GetSubscriptionsOk

`func (o *ModelAPIProjectRef) GetSubscriptionsOk() (*[]ModelAPISubscription, bool)`

GetSubscriptionsOk returns a tuple with the Subscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptions

`func (o *ModelAPIProjectRef) SetSubscriptions(v []ModelAPISubscription)`

SetSubscriptions sets Subscriptions field to given value.

### HasSubscriptions

`func (o *ModelAPIProjectRef) HasSubscriptions() bool`

HasSubscriptions returns a boolean if a field has been set.

### GetTaskAnnotationSettings

`func (o *ModelAPIProjectRef) GetTaskAnnotationSettings() ModelAPITaskAnnotationSettings`

GetTaskAnnotationSettings returns the TaskAnnotationSettings field if non-nil, zero value otherwise.

### GetTaskAnnotationSettingsOk

`func (o *ModelAPIProjectRef) GetTaskAnnotationSettingsOk() (*ModelAPITaskAnnotationSettings, bool)`

GetTaskAnnotationSettingsOk returns a tuple with the TaskAnnotationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskAnnotationSettings

`func (o *ModelAPIProjectRef) SetTaskAnnotationSettings(v ModelAPITaskAnnotationSettings)`

SetTaskAnnotationSettings sets TaskAnnotationSettings field to given value.

### HasTaskAnnotationSettings

`func (o *ModelAPIProjectRef) HasTaskAnnotationSettings() bool`

HasTaskAnnotationSettings returns a boolean if a field has been set.

### GetTracksPushEvents

`func (o *ModelAPIProjectRef) GetTracksPushEvents() bool`

GetTracksPushEvents returns the TracksPushEvents field if non-nil, zero value otherwise.

### GetTracksPushEventsOk

`func (o *ModelAPIProjectRef) GetTracksPushEventsOk() (*bool, bool)`

GetTracksPushEventsOk returns a tuple with the TracksPushEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracksPushEvents

`func (o *ModelAPIProjectRef) SetTracksPushEvents(v bool)`

SetTracksPushEvents sets TracksPushEvents field to given value.

### HasTracksPushEvents

`func (o *ModelAPIProjectRef) HasTracksPushEvents() bool`

HasTracksPushEvents returns a boolean if a field has been set.

### GetTriggers

`func (o *ModelAPIProjectRef) GetTriggers() []ModelAPITriggerDefinition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ModelAPIProjectRef) GetTriggersOk() (*[]ModelAPITriggerDefinition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ModelAPIProjectRef) SetTriggers(v []ModelAPITriggerDefinition)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ModelAPIProjectRef) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.

### GetUseRepoSettings

`func (o *ModelAPIProjectRef) GetUseRepoSettings() bool`

GetUseRepoSettings returns the UseRepoSettings field if non-nil, zero value otherwise.

### GetUseRepoSettingsOk

`func (o *ModelAPIProjectRef) GetUseRepoSettingsOk() (*bool, bool)`

GetUseRepoSettingsOk returns a tuple with the UseRepoSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRepoSettings

`func (o *ModelAPIProjectRef) SetUseRepoSettings(v bool)`

SetUseRepoSettings sets UseRepoSettings field to given value.

### HasUseRepoSettings

`func (o *ModelAPIProjectRef) HasUseRepoSettings() bool`

HasUseRepoSettings returns a boolean if a field has been set.

### GetVariables

`func (o *ModelAPIProjectRef) GetVariables() ModelAPIProjectVars`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ModelAPIProjectRef) GetVariablesOk() (*ModelAPIProjectVars, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ModelAPIProjectRef) SetVariables(v ModelAPIProjectVars)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ModelAPIProjectRef) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVersionControlEnabled

`func (o *ModelAPIProjectRef) GetVersionControlEnabled() bool`

GetVersionControlEnabled returns the VersionControlEnabled field if non-nil, zero value otherwise.

### GetVersionControlEnabledOk

`func (o *ModelAPIProjectRef) GetVersionControlEnabledOk() (*bool, bool)`

GetVersionControlEnabledOk returns a tuple with the VersionControlEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionControlEnabled

`func (o *ModelAPIProjectRef) SetVersionControlEnabled(v bool)`

SetVersionControlEnabled sets VersionControlEnabled field to given value.

### HasVersionControlEnabled

`func (o *ModelAPIProjectRef) HasVersionControlEnabled() bool`

HasVersionControlEnabled returns a boolean if a field has been set.

### GetWorkstationConfig

`func (o *ModelAPIProjectRef) GetWorkstationConfig() ModelAPIWorkstationConfig`

GetWorkstationConfig returns the WorkstationConfig field if non-nil, zero value otherwise.

### GetWorkstationConfigOk

`func (o *ModelAPIProjectRef) GetWorkstationConfigOk() (*ModelAPIWorkstationConfig, bool)`

GetWorkstationConfigOk returns a tuple with the WorkstationConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkstationConfig

`func (o *ModelAPIProjectRef) SetWorkstationConfig(v ModelAPIWorkstationConfig)`

SetWorkstationConfig sets WorkstationConfig field to given value.

### HasWorkstationConfig

`func (o *ModelAPIProjectRef) HasWorkstationConfig() bool`

HasWorkstationConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


