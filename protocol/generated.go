package protocol

// GENERATED CODE! DO NOT MODIFY BY HAND.

type AbsolutePathBuf = string
type AuthMode = string
type CollaborationMode = any
type ForcedLoginMethod = string
type InputModality = string
type Personality = string
type PlanType = string
type ReasoningEffort = string
type ReasoningSummary = string
type Resource = any
type ResourceTemplate = any
type ResponseItem = any
type SubAgentSource = string
type Tool = any
type Verbosity = string
type WebSearchMode = string

type AccountType = string

const (
	AccountApiKey  AccountType = "apiKey"
	AccountChatgpt AccountType = "chatgpt"
)

type Account struct {
	Type AccountType `json:"type"`
}
type ApiKeyAccount struct {
	Account
}
type ChatgptAccount struct {
	Account
	Email    string   `json:"email"`
	PlanType PlanType `json:"planType"`
}

type AccountLoginCompletedNotification struct {
	LoginId *string `json:"loginId"`
	Success bool    `json:"success"`
	Error   *string `json:"error"`
}

type AccountRateLimitsUpdatedNotification struct {
	RateLimits RateLimitSnapshot `json:"rateLimits"`
}

type AccountUpdatedNotification struct {
	AuthMode *AuthMode `json:"authMode"`
}

type AgentMessageDeltaNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
	ItemId   string `json:"itemId"`
	Delta    string `json:"delta"`
}

type AnalyticsConfig struct {
	Enabled *bool `json:"enabled"`
}

type AppBranding struct {
	Category          *string `json:"category"`
	Developer         *string `json:"developer"`
	Website           *string `json:"website"`
	PrivacyPolicy     *string `json:"privacyPolicy"`
	TermsOfService    *string `json:"termsOfService"`
	IsDiscoverableApp bool    `json:"isDiscoverableApp"`
}

type AppDisabledReason = string

const (
	AppDisabledReasonUnknown AppDisabledReason = "unknown"
	AppDisabledReasonUser    AppDisabledReason = "user"
)

type AppInfo struct {
	Id                  string         `json:"id"`
	Name                string         `json:"name"`
	Description         *string        `json:"description"`
	LogoUrl             *string        `json:"logoUrl"`
	LogoUrlDark         *string        `json:"logoUrlDark"`
	DistributionChannel *string        `json:"distributionChannel"`
	Branding            *AppBranding   `json:"branding"`
	AppMetadata         *AppMetadata   `json:"appMetadata"`
	Labels              map[string]any `json:"labels"`
	InstallUrl          *string        `json:"installUrl"`
	IsAccessible        bool           `json:"isAccessible"`
	IsEnabled           bool           `json:"isEnabled"`
}

type AppListUpdatedNotification struct {
	Data []AppInfo `json:"data"`
}

type AppMetadata struct {
	Review                     *AppReview      `json:"review"`
	Categories                 []string        `json:"categories"`
	SubCategories              []string        `json:"subCategories"`
	SeoDescription             *string         `json:"seoDescription"`
	Screenshots                []AppScreenshot `json:"screenshots"`
	Developer                  *string         `json:"developer"`
	Version                    *string         `json:"version"`
	VersionId                  *string         `json:"versionId"`
	VersionNotes               *string         `json:"versionNotes"`
	FirstPartyType             *string         `json:"firstPartyType"`
	FirstPartyRequiresInstall  *bool           `json:"firstPartyRequiresInstall"`
	ShowInComposerWhenUnlinked *bool           `json:"showInComposerWhenUnlinked"`
}

type AppReview struct {
	Status string `json:"status"`
}

type AppScreenshot struct {
	Url        *string `json:"url"`
	FileId     *string `json:"fileId"`
	UserPrompt string  `json:"userPrompt"`
}

type AppsConfig struct {
}

type AppsListParams struct {
	Cursor       *string  `json:"cursor,omitempty"`
	Limit        *float64 `json:"limit,omitempty"`
	ThreadId     *string  `json:"threadId,omitempty"`
	ForceRefetch *bool    `json:"forceRefetch,omitempty"`
}

type AppsListResponse struct {
	Data       []AppInfo `json:"data"`
	NextCursor *string   `json:"nextCursor"`
}

type AskForApproval = string

const (
	AskForApprovalUntrusted AskForApproval = "untrusted"
	AskForApprovalOnFailure AskForApproval = "on-failure"
	AskForApprovalOnRequest AskForApproval = "on-request"
	AskForApprovalNever     AskForApproval = "never"
)

type ByteRange struct {
	Start float64 `json:"start"`
	End   float64 `json:"end"`
}

type CancelLoginAccountParams struct {
	LoginId string `json:"loginId"`
}

type CancelLoginAccountResponse struct {
	Status CancelLoginAccountStatus `json:"status"`
}

type CancelLoginAccountStatus = string

const (
	CancelLoginAccountStatusCanceled CancelLoginAccountStatus = "canceled"
	CancelLoginAccountStatusNotFound CancelLoginAccountStatus = "notFound"
)

type ChatgptAuthTokensRefreshParams struct {
	Reason            ChatgptAuthTokensRefreshReason `json:"reason"`
	PreviousAccountId *string                        `json:"previousAccountId,omitempty"`
}

type ChatgptAuthTokensRefreshReason = string

type ChatgptAuthTokensRefreshResponse struct {
	AccessToken      string  `json:"accessToken"`
	ChatgptAccountId string  `json:"chatgptAccountId"`
	ChatgptPlanType  *string `json:"chatgptPlanType"`
}

type CodexErrorInfo = string

type CollabAgentState struct {
	Status  CollabAgentStatus `json:"status"`
	Message *string           `json:"message"`
}

type CollabAgentStatus = string

const (
	CollabAgentStatusPendingInit CollabAgentStatus = "pendingInit"
	CollabAgentStatusRunning     CollabAgentStatus = "running"
	CollabAgentStatusCompleted   CollabAgentStatus = "completed"
	CollabAgentStatusErrored     CollabAgentStatus = "errored"
	CollabAgentStatusShutdown    CollabAgentStatus = "shutdown"
	CollabAgentStatusNotFound    CollabAgentStatus = "notFound"
)

type CollabAgentTool = string

const (
	CollabAgentToolSpawnAgent  CollabAgentTool = "spawnAgent"
	CollabAgentToolSendInput   CollabAgentTool = "sendInput"
	CollabAgentToolResumeAgent CollabAgentTool = "resumeAgent"
	CollabAgentToolWait        CollabAgentTool = "wait"
	CollabAgentToolCloseAgent  CollabAgentTool = "closeAgent"
)

type CollabAgentToolCallStatus = string

const (
	CollabAgentToolCallStatusInProgress CollabAgentToolCallStatus = "inProgress"
	CollabAgentToolCallStatusCompleted  CollabAgentToolCallStatus = "completed"
	CollabAgentToolCallStatusFailed     CollabAgentToolCallStatus = "failed"
)

type CommandActionType = string

const (
	CommandActionRead      CommandActionType = "read"
	CommandActionListFiles CommandActionType = "listFiles"
	CommandActionSearch    CommandActionType = "search"
	CommandActionUnknown   CommandActionType = "unknown"
)

type CommandAction struct {
	Type CommandActionType `json:"type"`
}
type ReadCommandAction struct {
	CommandAction
	Command string `json:"command"`
	Name    string `json:"name"`
	Path    string `json:"path"`
}
type ListFilesCommandAction struct {
	CommandAction
	Command string  `json:"command"`
	Path    *string `json:"path"`
}
type SearchCommandAction struct {
	CommandAction
	Command string  `json:"command"`
	Query   *string `json:"query"`
	Path    *string `json:"path"`
}
type UnknownCommandAction struct {
	CommandAction
	Command string `json:"command"`
}

type CommandExecParams struct {
	Command       []string       `json:"command"`
	TimeoutMs     *float64       `json:"timeoutMs,omitempty"`
	Cwd           *string        `json:"cwd,omitempty"`
	SandboxPolicy *SandboxPolicy `json:"sandboxPolicy,omitempty"`
}

type CommandExecResponse struct {
	ExitCode float64 `json:"exitCode"`
	Stdout   string  `json:"stdout"`
	Stderr   string  `json:"stderr"`
}

type CommandExecutionApprovalDecision = string

type CommandExecutionOutputDeltaNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
	ItemId   string `json:"itemId"`
	Delta    string `json:"delta"`
}

type CommandExecutionRequestApprovalParams struct {
	ThreadId                    string               `json:"threadId"`
	TurnId                      string               `json:"turnId"`
	ItemId                      string               `json:"itemId"`
	ApprovalId                  *string              `json:"approvalId,omitempty"`
	Reason                      *string              `json:"reason,omitempty"`
	Command                     *string              `json:"command,omitempty"`
	Cwd                         *string              `json:"cwd,omitempty"`
	CommandActions              []CommandAction      `json:"commandActions,omitempty"`
	ProposedExecpolicyAmendment *ExecPolicyAmendment `json:"proposedExecpolicyAmendment,omitempty"`
}

type CommandExecutionRequestApprovalResponse struct {
	Decision CommandExecutionApprovalDecision `json:"decision"`
}

type CommandExecutionStatus = string

const (
	CommandExecutionStatusInProgress CommandExecutionStatus = "inProgress"
	CommandExecutionStatusCompleted  CommandExecutionStatus = "completed"
	CommandExecutionStatusFailed     CommandExecutionStatus = "failed"
	CommandExecutionStatusDeclined   CommandExecutionStatus = "declined"
)

type Config struct {
	Model                      *string                `json:"model"`
	ReviewModel                *string                `json:"review_model"`
	ModelContextWindow         *int64                 `json:"model_context_window"`
	ModelAutoCompactTokenLimit *int64                 `json:"model_auto_compact_token_limit"`
	ModelProvider              *string                `json:"model_provider"`
	ApprovalPolicy             *AskForApproval        `json:"approval_policy"`
	SandboxMode                *SandboxMode           `json:"sandbox_mode"`
	SandboxWorkspaceWrite      *SandboxWorkspaceWrite `json:"sandbox_workspace_write"`
	ForcedChatgptWorkspaceId   *string                `json:"forced_chatgpt_workspace_id"`
	ForcedLoginMethod          *ForcedLoginMethod     `json:"forced_login_method"`
	WebSearch                  *WebSearchMode         `json:"web_search"`
	Tools                      *ToolsV2               `json:"tools"`
	Profile                    *string                `json:"profile"`
	Profiles                   map[string]ProfileV2   `json:"profiles"`
	Instructions               *string                `json:"instructions"`
	DeveloperInstructions      *string                `json:"developer_instructions"`
	CompactPrompt              *string                `json:"compact_prompt"`
	ModelReasoningEffort       *ReasoningEffort       `json:"model_reasoning_effort"`
	ModelReasoningSummary      *ReasoningSummary      `json:"model_reasoning_summary"`
	ModelVerbosity             *Verbosity             `json:"model_verbosity"`
	Analytics                  any                    `json:"analytics"`
}

type ConfigBatchWriteParams struct {
	Edits           []ConfigEdit `json:"edits"`
	FilePath        *string      `json:"filePath,omitempty"`
	ExpectedVersion *string      `json:"expectedVersion,omitempty"`
}

type ConfigEdit struct {
	KeyPath       string        `json:"keyPath"`
	Value         any           `json:"value"`
	MergeStrategy MergeStrategy `json:"mergeStrategy"`
}

type ConfigLayer struct {
	Name           ConfigLayerSource `json:"name"`
	Version        string            `json:"version"`
	Config         any               `json:"config"`
	DisabledReason *string           `json:"disabledReason"`
}

type ConfigLayerMetadata struct {
	Name    ConfigLayerSource `json:"name"`
	Version string            `json:"version"`
}

type ConfigLayerSourceType = string

const (
	ConfigLayerSourceMdm                             ConfigLayerSourceType = "mdm"
	ConfigLayerSourceSystem                          ConfigLayerSourceType = "system"
	ConfigLayerSourceUser                            ConfigLayerSourceType = "user"
	ConfigLayerSourceProject                         ConfigLayerSourceType = "project"
	ConfigLayerSourceSessionFlags                    ConfigLayerSourceType = "sessionFlags"
	ConfigLayerSourceLegacyManagedConfigTomlFromFile ConfigLayerSourceType = "legacyManagedConfigTomlFromFile"
	ConfigLayerSourceLegacyManagedConfigTomlFromMdm  ConfigLayerSourceType = "legacyManagedConfigTomlFromMdm"
)

type ConfigLayerSource struct {
	Type ConfigLayerSourceType `json:"type"`
}
type MdmConfigLayerSource struct {
	ConfigLayerSource
	Domain string `json:"domain"`
	Key    string `json:"key"`
}
type SystemConfigLayerSource struct {
	ConfigLayerSource
	File AbsolutePathBuf `json:"file"`
}
type UserConfigLayerSource struct {
	ConfigLayerSource
	File AbsolutePathBuf `json:"file"`
}
type ProjectConfigLayerSource struct {
	ConfigLayerSource
	DotCodexFolder AbsolutePathBuf `json:"dotCodexFolder"`
}
type SessionFlagsConfigLayerSource struct {
	ConfigLayerSource
}
type LegacyManagedConfigTomlFromFileConfigLayerSource struct {
	ConfigLayerSource
	File AbsolutePathBuf `json:"file"`
}
type LegacyManagedConfigTomlFromMdmConfigLayerSource struct {
	ConfigLayerSource
}

type ConfigReadParams struct {
	IncludeLayers bool    `json:"includeLayers"`
	Cwd           *string `json:"cwd,omitempty"`
}

type ConfigReadResponse struct {
	Config  Config                         `json:"config"`
	Origins map[string]ConfigLayerMetadata `json:"origins"`
	Layers  []ConfigLayer                  `json:"layers"`
}

type ConfigRequirements struct {
	AllowedApprovalPolicies []AskForApproval      `json:"allowedApprovalPolicies"`
	AllowedSandboxModes     []SandboxMode         `json:"allowedSandboxModes"`
	AllowedWebSearchModes   []WebSearchMode       `json:"allowedWebSearchModes"`
	EnforceResidency        *ResidencyRequirement `json:"enforceResidency"`
}

type ConfigRequirementsReadResponse struct {
	Requirements *ConfigRequirements `json:"requirements"`
}

type ConfigValueWriteParams struct {
	KeyPath         string        `json:"keyPath"`
	Value           any           `json:"value"`
	MergeStrategy   MergeStrategy `json:"mergeStrategy"`
	FilePath        *string       `json:"filePath,omitempty"`
	ExpectedVersion *string       `json:"expectedVersion,omitempty"`
}

type ConfigWarningNotification struct {
	Summary string     `json:"summary"`
	Details *string    `json:"details"`
	Path    *string    `json:"path,omitempty"`
	Range   *TextRange `json:"range,omitempty"`
}

type ConfigWriteResponse struct {
	Status             WriteStatus         `json:"status"`
	Version            string              `json:"version"`
	FilePath           AbsolutePathBuf     `json:"filePath"`
	OverriddenMetadata *OverriddenMetadata `json:"overriddenMetadata"`
}

type ContextCompactedNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
}

type CreditsSnapshot struct {
	HasCredits bool    `json:"hasCredits"`
	Unlimited  bool    `json:"unlimited"`
	Balance    *string `json:"balance"`
}

type DeprecationNoticeNotification struct {
	Summary string  `json:"summary"`
	Details *string `json:"details"`
}

type DynamicToolCallOutputContentItemType = string

const (
	DynamicToolCallOutputContentItemInputText  DynamicToolCallOutputContentItemType = "inputText"
	DynamicToolCallOutputContentItemInputImage DynamicToolCallOutputContentItemType = "inputImage"
)

type DynamicToolCallOutputContentItem struct {
	Type DynamicToolCallOutputContentItemType `json:"type"`
}
type InputTextDynamicToolCallOutputContentItem struct {
	DynamicToolCallOutputContentItem
	Text string `json:"text"`
}
type InputImageDynamicToolCallOutputContentItem struct {
	DynamicToolCallOutputContentItem
	ImageUrl string `json:"imageUrl"`
}

type DynamicToolCallParams struct {
	ThreadId  string `json:"threadId"`
	TurnId    string `json:"turnId"`
	CallId    string `json:"callId"`
	Tool      string `json:"tool"`
	Arguments any    `json:"arguments"`
}

type DynamicToolCallResponse struct {
	ContentItems []DynamicToolCallOutputContentItem `json:"contentItems"`
	Success      bool                               `json:"success"`
}

type DynamicToolSpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	InputSchema any    `json:"inputSchema"`
}

type ErrorNotification struct {
	Error     TurnError `json:"error"`
	WillRetry bool      `json:"willRetry"`
	ThreadId  string    `json:"threadId"`
	TurnId    string    `json:"turnId"`
}

type ExecPolicyAmendment = []string

type ExperimentalFeature struct {
	Name           string                   `json:"name"`
	Stage          ExperimentalFeatureStage `json:"stage"`
	DisplayName    *string                  `json:"displayName"`
	Description    *string                  `json:"description"`
	Announcement   *string                  `json:"announcement"`
	Enabled        bool                     `json:"enabled"`
	DefaultEnabled bool                     `json:"defaultEnabled"`
}

type ExperimentalFeatureListParams struct {
	Cursor *string  `json:"cursor,omitempty"`
	Limit  *float64 `json:"limit,omitempty"`
}

type ExperimentalFeatureListResponse struct {
	Data       []ExperimentalFeature `json:"data"`
	NextCursor *string               `json:"nextCursor"`
}

type ExperimentalFeatureStage = string

const (
	ExperimentalFeatureStageBeta             ExperimentalFeatureStage = "beta"
	ExperimentalFeatureStageUnderDevelopment ExperimentalFeatureStage = "underDevelopment"
	ExperimentalFeatureStageStable           ExperimentalFeatureStage = "stable"
	ExperimentalFeatureStageDeprecated       ExperimentalFeatureStage = "deprecated"
	ExperimentalFeatureStageRemoved          ExperimentalFeatureStage = "removed"
)

type FeedbackUploadParams struct {
	Classification string  `json:"classification"`
	Reason         *string `json:"reason,omitempty"`
	ThreadId       *string `json:"threadId,omitempty"`
	IncludeLogs    bool    `json:"includeLogs"`
}

type FeedbackUploadResponse struct {
	ThreadId string `json:"threadId"`
}

type FileChangeApprovalDecision = string

const (
	FileChangeApprovalDecisionAccept           FileChangeApprovalDecision = "accept"
	FileChangeApprovalDecisionAcceptForSession FileChangeApprovalDecision = "acceptForSession"
	FileChangeApprovalDecisionDecline          FileChangeApprovalDecision = "decline"
	FileChangeApprovalDecisionCancel           FileChangeApprovalDecision = "cancel"
)

type FileChangeOutputDeltaNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
	ItemId   string `json:"itemId"`
	Delta    string `json:"delta"`
}

type FileChangeRequestApprovalParams struct {
	ThreadId  string  `json:"threadId"`
	TurnId    string  `json:"turnId"`
	ItemId    string  `json:"itemId"`
	Reason    *string `json:"reason,omitempty"`
	GrantRoot *string `json:"grantRoot,omitempty"`
}

type FileChangeRequestApprovalResponse struct {
	Decision FileChangeApprovalDecision `json:"decision"`
}

type FileUpdateChange struct {
	Path string          `json:"path"`
	Kind PatchChangeKind `json:"kind"`
	Diff string          `json:"diff"`
}

type GetAccountParams struct {
	RefreshToken bool `json:"refreshToken"`
}

type GetAccountRateLimitsResponse struct {
	RateLimits          RateLimitSnapshot `json:"rateLimits"`
	RateLimitsByLimitId map[string]any    `json:"rateLimitsByLimitId"`
}

type GetAccountResponse struct {
	Account            *Account `json:"account"`
	RequiresOpenaiAuth bool     `json:"requiresOpenaiAuth"`
}

type GitInfo struct {
	Sha       *string `json:"sha"`
	Branch    *string `json:"branch"`
	OriginUrl *string `json:"originUrl"`
}

type HazelnutScope = string

const (
	HazelnutScopeExample         HazelnutScope = "example"
	HazelnutScopeWorkspaceShared HazelnutScope = "workspace-shared"
	HazelnutScopeAllShared       HazelnutScope = "all-shared"
	HazelnutScopePersonal        HazelnutScope = "personal"
)

type ItemCompletedNotification struct {
	Item     ThreadItem `json:"item"`
	ThreadId string     `json:"threadId"`
	TurnId   string     `json:"turnId"`
}

type ItemStartedNotification struct {
	Item     ThreadItem `json:"item"`
	ThreadId string     `json:"threadId"`
	TurnId   string     `json:"turnId"`
}

type ListMcpServerStatusParams struct {
	Cursor *string  `json:"cursor,omitempty"`
	Limit  *float64 `json:"limit,omitempty"`
}

type ListMcpServerStatusResponse struct {
	Data       []McpServerStatus `json:"data"`
	NextCursor *string           `json:"nextCursor"`
}

type LoginAccountParamsType = string

const (
	LoginAccountParamsApiKey            LoginAccountParamsType = "apiKey"
	LoginAccountParamsChatgpt           LoginAccountParamsType = "chatgpt"
	LoginAccountParamsChatgptAuthTokens LoginAccountParamsType = "chatgptAuthTokens"
)

type LoginAccountParams struct {
	Type LoginAccountParamsType `json:"type"`
}
type ApiKeyLoginAccountParams struct {
	LoginAccountParams
	ApiKey string `json:"apiKey"`
}
type ChatgptLoginAccountParams struct {
	LoginAccountParams
}
type ChatgptAuthTokensLoginAccountParams struct {
	LoginAccountParams
	AccessToken      string  `json:"accessToken"`
	ChatgptAccountId string  `json:"chatgptAccountId"`
	ChatgptPlanType  *string `json:"chatgptPlanType,omitempty"`
}

type LoginAccountResponseType = string

const (
	LoginAccountResponseApiKey            LoginAccountResponseType = "apiKey"
	LoginAccountResponseChatgpt           LoginAccountResponseType = "chatgpt"
	LoginAccountResponseChatgptAuthTokens LoginAccountResponseType = "chatgptAuthTokens"
)

type LoginAccountResponse struct {
	Type LoginAccountResponseType `json:"type"`
}
type ApiKeyLoginAccountResponse struct {
	LoginAccountResponse
}
type ChatgptLoginAccountResponse struct {
	LoginAccountResponse
	LoginId string `json:"loginId"`
	AuthUrl string `json:"authUrl"`
}
type ChatgptAuthTokensLoginAccountResponse struct {
	LoginAccountResponse
}

type LogoutAccountResponse struct{}

type McpAuthStatus = string

const (
	McpAuthStatusUnsupported McpAuthStatus = "unsupported"
	McpAuthStatusNotLoggedIn McpAuthStatus = "notLoggedIn"
	McpAuthStatusBearerToken McpAuthStatus = "bearerToken"
	McpAuthStatusOAuth       McpAuthStatus = "oAuth"
)

type McpServerOauthLoginCompletedNotification struct {
	Name    string  `json:"name"`
	Success bool    `json:"success"`
	Error   *string `json:"error,omitempty"`
}

type McpServerOauthLoginParams struct {
	Name        string   `json:"name"`
	Scopes      []string `json:"scopes,omitempty"`
	TimeoutSecs *int64   `json:"timeoutSecs,omitempty"`
}

type McpServerOauthLoginResponse struct {
	AuthorizationUrl string `json:"authorizationUrl"`
}

type McpServerRefreshResponse struct{}

type McpServerStatus struct {
	Name              string             `json:"name"`
	Tools             map[string]Tool    `json:"tools"`
	Resources         []Resource         `json:"resources"`
	ResourceTemplates []ResourceTemplate `json:"resourceTemplates"`
	AuthStatus        McpAuthStatus      `json:"authStatus"`
}

type McpToolCallError struct {
	Message string `json:"message"`
}

type McpToolCallProgressNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
	ItemId   string `json:"itemId"`
	Message  string `json:"message"`
}

type McpToolCallResult struct {
	Content           any `json:"content"`
	StructuredContent any `json:"structuredContent"`
}

type McpToolCallStatus = string

const (
	McpToolCallStatusInProgress McpToolCallStatus = "inProgress"
	McpToolCallStatusCompleted  McpToolCallStatus = "completed"
	McpToolCallStatusFailed     McpToolCallStatus = "failed"
)

type MergeStrategy = string

const (
	MergeStrategyReplace MergeStrategy = "replace"
	MergeStrategyUpsert  MergeStrategy = "upsert"
)

type Model struct {
	Id                        string                  `json:"id"`
	Model                     string                  `json:"model"`
	Upgrade                   *string                 `json:"upgrade"`
	DisplayName               string                  `json:"displayName"`
	Description               string                  `json:"description"`
	Hidden                    bool                    `json:"hidden"`
	SupportedReasoningEfforts []ReasoningEffortOption `json:"supportedReasoningEfforts"`
	DefaultReasoningEffort    ReasoningEffort         `json:"defaultReasoningEffort"`
	InputModalities           []InputModality         `json:"inputModalities"`
	SupportsPersonality       bool                    `json:"supportsPersonality"`
	IsDefault                 bool                    `json:"isDefault"`
}

type ModelListParams struct {
	Cursor        *string  `json:"cursor,omitempty"`
	Limit         *float64 `json:"limit,omitempty"`
	IncludeHidden *bool    `json:"includeHidden,omitempty"`
}

type ModelListResponse struct {
	Data       []Model `json:"data"`
	NextCursor *string `json:"nextCursor"`
}

type ModelRerouteReason = string

type ModelReroutedNotification struct {
	ThreadId  string             `json:"threadId"`
	TurnId    string             `json:"turnId"`
	FromModel string             `json:"fromModel"`
	ToModel   string             `json:"toModel"`
	Reason    ModelRerouteReason `json:"reason"`
}

type NetworkAccess = string

const (
	NetworkAccessRestricted NetworkAccess = "restricted"
	NetworkAccessEnabled    NetworkAccess = "enabled"
)

type NetworkRequirements struct {
	Enabled                          *bool    `json:"enabled"`
	HttpPort                         *float64 `json:"httpPort"`
	SocksPort                        *float64 `json:"socksPort"`
	AllowUpstreamProxy               *bool    `json:"allowUpstreamProxy"`
	DangerouslyAllowNonLoopbackProxy *bool    `json:"dangerouslyAllowNonLoopbackProxy"`
	DangerouslyAllowNonLoopbackAdmin *bool    `json:"dangerouslyAllowNonLoopbackAdmin"`
	AllowedDomains                   []string `json:"allowedDomains"`
	DeniedDomains                    []string `json:"deniedDomains"`
	AllowUnixSockets                 []string `json:"allowUnixSockets"`
	AllowLocalBinding                *bool    `json:"allowLocalBinding"`
}

type OverriddenMetadata struct {
	Message         string              `json:"message"`
	OverridingLayer ConfigLayerMetadata `json:"overridingLayer"`
	EffectiveValue  any                 `json:"effectiveValue"`
}

type PatchApplyStatus = string

const (
	PatchApplyStatusInProgress PatchApplyStatus = "inProgress"
	PatchApplyStatusCompleted  PatchApplyStatus = "completed"
	PatchApplyStatusFailed     PatchApplyStatus = "failed"
	PatchApplyStatusDeclined   PatchApplyStatus = "declined"
)

type PatchChangeKindType = string

const (
	PatchChangeKindAdd    PatchChangeKindType = "add"
	PatchChangeKindDelete PatchChangeKindType = "delete"
	PatchChangeKindUpdate PatchChangeKindType = "update"
)

type PatchChangeKind struct {
	Type PatchChangeKindType `json:"type"`
}
type AddPatchChangeKind struct {
	PatchChangeKind
}
type DeletePatchChangeKind struct {
	PatchChangeKind
}
type UpdatePatchChangeKind struct {
	PatchChangeKind
	MovePath *string `json:"move_path"`
}

type PlanDeltaNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
	ItemId   string `json:"itemId"`
	Delta    string `json:"delta"`
}

type ProductSurface = string

const (
	ProductSurfaceChatgpt ProductSurface = "chatgpt"
	ProductSurfaceCodex   ProductSurface = "codex"
	ProductSurfaceApi     ProductSurface = "api"
	ProductSurfaceAtlas   ProductSurface = "atlas"
)

type ProfileV2 struct {
	Model                 *string           `json:"model"`
	ModelProvider         *string           `json:"model_provider"`
	ApprovalPolicy        *AskForApproval   `json:"approval_policy"`
	ModelReasoningEffort  *ReasoningEffort  `json:"model_reasoning_effort"`
	ModelReasoningSummary *ReasoningSummary `json:"model_reasoning_summary"`
	ModelVerbosity        *Verbosity        `json:"model_verbosity"`
	WebSearch             *WebSearchMode    `json:"web_search"`
	ChatgptBaseUrl        *string           `json:"chatgpt_base_url"`
}

type RateLimitSnapshot struct {
	LimitId   *string          `json:"limitId"`
	LimitName *string          `json:"limitName"`
	Primary   *RateLimitWindow `json:"primary"`
	Secondary *RateLimitWindow `json:"secondary"`
	Credits   *CreditsSnapshot `json:"credits"`
	PlanType  *PlanType        `json:"planType"`
}

type RateLimitWindow struct {
	UsedPercent        float64  `json:"usedPercent"`
	WindowDurationMins *float64 `json:"windowDurationMins"`
	ResetsAt           *float64 `json:"resetsAt"`
}

type RawResponseItemCompletedNotification struct {
	ThreadId string       `json:"threadId"`
	TurnId   string       `json:"turnId"`
	Item     ResponseItem `json:"item"`
}

type ReadOnlyAccessType = string

const (
	ReadOnlyAccessRestricted ReadOnlyAccessType = "restricted"
	ReadOnlyAccessFullAccess ReadOnlyAccessType = "fullAccess"
)

type ReadOnlyAccess struct {
	Type ReadOnlyAccessType `json:"type"`
}
type RestrictedReadOnlyAccess struct {
	ReadOnlyAccess
	IncludePlatformDefaults bool              `json:"includePlatformDefaults"`
	ReadableRoots           []AbsolutePathBuf `json:"readableRoots"`
}
type FullAccessReadOnlyAccess struct {
	ReadOnlyAccess
}

type ReasoningEffortOption struct {
	ReasoningEffort ReasoningEffort `json:"reasoningEffort"`
	Description     string          `json:"description"`
}

type ReasoningSummaryPartAddedNotification struct {
	ThreadId     string  `json:"threadId"`
	TurnId       string  `json:"turnId"`
	ItemId       string  `json:"itemId"`
	SummaryIndex float64 `json:"summaryIndex"`
}

type ReasoningSummaryTextDeltaNotification struct {
	ThreadId     string  `json:"threadId"`
	TurnId       string  `json:"turnId"`
	ItemId       string  `json:"itemId"`
	Delta        string  `json:"delta"`
	SummaryIndex float64 `json:"summaryIndex"`
}

type ReasoningTextDeltaNotification struct {
	ThreadId     string  `json:"threadId"`
	TurnId       string  `json:"turnId"`
	ItemId       string  `json:"itemId"`
	Delta        string  `json:"delta"`
	ContentIndex float64 `json:"contentIndex"`
}

type RemoteSkillSummary struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ResidencyRequirement = string

type ReviewDelivery = string

const (
	ReviewDeliveryInline   ReviewDelivery = "inline"
	ReviewDeliveryDetached ReviewDelivery = "detached"
)

type ReviewStartParams struct {
	ThreadId string          `json:"threadId"`
	Target   ReviewTarget    `json:"target"`
	Delivery *ReviewDelivery `json:"delivery,omitempty"`
}

type ReviewStartResponse struct {
	Turn           Turn   `json:"turn"`
	ReviewThreadId string `json:"reviewThreadId"`
}

type ReviewTargetType = string

const (
	ReviewTargetUncommittedChanges ReviewTargetType = "uncommittedChanges"
	ReviewTargetBaseBranch         ReviewTargetType = "baseBranch"
	ReviewTargetCommit             ReviewTargetType = "commit"
	ReviewTargetCustom             ReviewTargetType = "custom"
)

type ReviewTarget struct {
	Type ReviewTargetType `json:"type"`
}
type UncommittedChangesReviewTarget struct {
	ReviewTarget
}
type BaseBranchReviewTarget struct {
	ReviewTarget
	Branch string `json:"branch"`
}
type CommitReviewTarget struct {
	ReviewTarget
	Sha   string  `json:"sha"`
	Title *string `json:"title"`
}
type CustomReviewTarget struct {
	ReviewTarget
	Instructions string `json:"instructions"`
}

type SandboxMode = string

const (
	SandboxModeReadOnly         SandboxMode = "read-only"
	SandboxModeWorkspaceWrite   SandboxMode = "workspace-write"
	SandboxModeDangerFullAccess SandboxMode = "danger-full-access"
)

type SandboxPolicyType = string

const (
	SandboxPolicyDangerFullAccess SandboxPolicyType = "dangerFullAccess"
	SandboxPolicyReadOnly         SandboxPolicyType = "readOnly"
	SandboxPolicyExternalSandbox  SandboxPolicyType = "externalSandbox"
	SandboxPolicyWorkspaceWrite   SandboxPolicyType = "workspaceWrite"
)

type SandboxPolicy struct {
	Type SandboxPolicyType `json:"type"`
}
type DangerFullAccessSandboxPolicy struct {
	SandboxPolicy
}
type ReadOnlySandboxPolicy struct {
	SandboxPolicy
	Access ReadOnlyAccess `json:"access"`
}
type ExternalSandboxSandboxPolicy struct {
	SandboxPolicy
	NetworkAccess NetworkAccess `json:"networkAccess"`
}
type WorkspaceWriteSandboxPolicy struct {
	SandboxPolicy
	WritableRoots       []AbsolutePathBuf `json:"writableRoots"`
	ReadOnlyAccess      ReadOnlyAccess    `json:"readOnlyAccess"`
	NetworkAccess       bool              `json:"networkAccess"`
	ExcludeTmpdirEnvVar bool              `json:"excludeTmpdirEnvVar"`
	ExcludeSlashTmp     bool              `json:"excludeSlashTmp"`
}

type SandboxWorkspaceWrite struct {
	WritableRoots       []string `json:"writable_roots"`
	NetworkAccess       bool     `json:"network_access"`
	ExcludeTmpdirEnvVar bool     `json:"exclude_tmpdir_env_var"`
	ExcludeSlashTmp     bool     `json:"exclude_slash_tmp"`
}

type SessionSource = string

type SkillDependencies struct {
	Tools []SkillToolDependency `json:"tools"`
}

type SkillErrorInfo struct {
	Path    string `json:"path"`
	Message string `json:"message"`
}

type SkillInterface struct {
	DisplayName      *string `json:"displayName,omitempty"`
	ShortDescription *string `json:"shortDescription,omitempty"`
	IconSmall        *string `json:"iconSmall,omitempty"`
	IconLarge        *string `json:"iconLarge,omitempty"`
	BrandColor       *string `json:"brandColor,omitempty"`
	DefaultPrompt    *string `json:"defaultPrompt,omitempty"`
}

type SkillMetadata struct {
	Name             string             `json:"name"`
	Description      string             `json:"description"`
	ShortDescription *string            `json:"shortDescription,omitempty"`
	Interface        *SkillInterface    `json:"interface,omitempty"`
	Dependencies     *SkillDependencies `json:"dependencies,omitempty"`
	Path             string             `json:"path"`
	Scope            SkillScope         `json:"scope"`
	Enabled          bool               `json:"enabled"`
}

type SkillScope = string

const (
	SkillScopeUser   SkillScope = "user"
	SkillScopeRepo   SkillScope = "repo"
	SkillScopeSystem SkillScope = "system"
	SkillScopeAdmin  SkillScope = "admin"
)

type SkillToolDependency struct {
	Type        string  `json:"type"`
	Value       string  `json:"value"`
	Description *string `json:"description,omitempty"`
	Transport   *string `json:"transport,omitempty"`
	Command     *string `json:"command,omitempty"`
	Url         *string `json:"url,omitempty"`
}

type SkillsConfigWriteParams struct {
	Path    string `json:"path"`
	Enabled bool   `json:"enabled"`
}

type SkillsConfigWriteResponse struct {
	EffectiveEnabled bool `json:"effectiveEnabled"`
}

type SkillsListEntry struct {
	Cwd    string           `json:"cwd"`
	Skills []SkillMetadata  `json:"skills"`
	Errors []SkillErrorInfo `json:"errors"`
}

type SkillsListExtraRootsForCwd struct {
	Cwd            string   `json:"cwd"`
	ExtraUserRoots []string `json:"extraUserRoots"`
}

type SkillsListParams struct {
	Cwds                 []string                     `json:"cwds,omitempty"`
	ForceReload          *bool                        `json:"forceReload,omitempty"`
	PerCwdExtraUserRoots []SkillsListExtraRootsForCwd `json:"perCwdExtraUserRoots,omitempty"`
}

type SkillsListResponse struct {
	Data []SkillsListEntry `json:"data"`
}

type SkillsRemoteReadParams struct {
	HazelnutScope  HazelnutScope  `json:"hazelnutScope"`
	ProductSurface ProductSurface `json:"productSurface"`
	Enabled        bool           `json:"enabled"`
}

type SkillsRemoteReadResponse struct {
	Data []RemoteSkillSummary `json:"data"`
}

type SkillsRemoteWriteParams struct {
	HazelnutId string `json:"hazelnutId"`
}

type SkillsRemoteWriteResponse struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

type TerminalInteractionNotification struct {
	ThreadId  string `json:"threadId"`
	TurnId    string `json:"turnId"`
	ItemId    string `json:"itemId"`
	ProcessId string `json:"processId"`
	Stdin     string `json:"stdin"`
}

type TextElement struct {
	ByteRange   ByteRange `json:"byteRange"`
	Placeholder *string   `json:"placeholder"`
}

type TextPosition struct {
	Line   float64 `json:"line"`
	Column float64 `json:"column"`
}

type TextRange struct {
	Start TextPosition `json:"start"`
	End   TextPosition `json:"end"`
}

type Thread struct {
	Id            string        `json:"id"`
	Preview       string        `json:"preview"`
	ModelProvider string        `json:"modelProvider"`
	CreatedAt     float64       `json:"createdAt"`
	UpdatedAt     float64       `json:"updatedAt"`
	Path          *string       `json:"path"`
	Cwd           string        `json:"cwd"`
	CliVersion    string        `json:"cliVersion"`
	Source        SessionSource `json:"source"`
	GitInfo       *GitInfo      `json:"gitInfo"`
	Turns         []Turn        `json:"turns"`
}

type ThreadArchiveParams struct {
	ThreadId string `json:"threadId"`
}

type ThreadArchiveResponse struct{}

type ThreadArchivedNotification struct {
	ThreadId string `json:"threadId"`
}

type ThreadCompactStartParams struct {
	ThreadId string `json:"threadId"`
}

type ThreadCompactStartResponse struct{}

type ThreadForkParams struct {
	ThreadId               string          `json:"threadId"`
	Path                   *string         `json:"path,omitempty"`
	Model                  *string         `json:"model,omitempty"`
	ModelProvider          *string         `json:"modelProvider,omitempty"`
	Cwd                    *string         `json:"cwd,omitempty"`
	ApprovalPolicy         *AskForApproval `json:"approvalPolicy,omitempty"`
	Sandbox                *SandboxMode    `json:"sandbox,omitempty"`
	Config                 any             `json:"config,omitempty"`
	BaseInstructions       *string         `json:"baseInstructions,omitempty"`
	DeveloperInstructions  *string         `json:"developerInstructions,omitempty"`
	PersistExtendedHistory bool            `json:"persistExtendedHistory"`
}

type ThreadForkResponse struct {
	Thread          Thread           `json:"thread"`
	Model           string           `json:"model"`
	ModelProvider   string           `json:"modelProvider"`
	Cwd             string           `json:"cwd"`
	ApprovalPolicy  AskForApproval   `json:"approvalPolicy"`
	Sandbox         SandboxPolicy    `json:"sandbox"`
	ReasoningEffort *ReasoningEffort `json:"reasoningEffort"`
}

type ThreadItemType = string

const (
	ThreadItemUserMessage         ThreadItemType = "userMessage"
	ThreadItemAgentMessage        ThreadItemType = "agentMessage"
	ThreadItemPlan                ThreadItemType = "plan"
	ThreadItemReasoning           ThreadItemType = "reasoning"
	ThreadItemCommandExecution    ThreadItemType = "commandExecution"
	ThreadItemFileChange          ThreadItemType = "fileChange"
	ThreadItemMcpToolCall         ThreadItemType = "mcpToolCall"
	ThreadItemCollabAgentToolCall ThreadItemType = "collabAgentToolCall"
	ThreadItemWebSearch           ThreadItemType = "webSearch"
	ThreadItemImageView           ThreadItemType = "imageView"
	ThreadItemEnteredReviewMode   ThreadItemType = "enteredReviewMode"
	ThreadItemExitedReviewMode    ThreadItemType = "exitedReviewMode"
	ThreadItemContextCompaction   ThreadItemType = "contextCompaction"
)

type ThreadItem struct {
	Type ThreadItemType `json:"type"`
}
type UserMessageThreadItem struct {
	ThreadItem
	Id      string      `json:"id"`
	Content []UserInput `json:"content"`
}
type AgentMessageThreadItem struct {
	ThreadItem
	Id   string `json:"id"`
	Text string `json:"text"`
}
type PlanThreadItem struct {
	ThreadItem
	Id   string `json:"id"`
	Text string `json:"text"`
}
type ReasoningThreadItem struct {
	ThreadItem
	Id      string   `json:"id"`
	Summary []string `json:"summary"`
	Content []string `json:"content"`
}
type CommandExecutionThreadItem struct {
	ThreadItem
	Id               string                 `json:"id"`
	Command          string                 `json:"command"`
	Cwd              string                 `json:"cwd"`
	ProcessId        *string                `json:"processId"`
	Status           CommandExecutionStatus `json:"status"`
	CommandActions   []CommandAction        `json:"commandActions"`
	AggregatedOutput *string                `json:"aggregatedOutput"`
	ExitCode         *float64               `json:"exitCode"`
	DurationMs       *float64               `json:"durationMs"`
}
type FileChangeThreadItem struct {
	ThreadItem
	Id      string             `json:"id"`
	Changes []FileUpdateChange `json:"changes"`
	Status  PatchApplyStatus   `json:"status"`
}
type McpToolCallThreadItem struct {
	ThreadItem
	Id         string             `json:"id"`
	Server     string             `json:"server"`
	Tool       string             `json:"tool"`
	Status     McpToolCallStatus  `json:"status"`
	Arguments  any                `json:"arguments"`
	Result     *McpToolCallResult `json:"result"`
	Error      *McpToolCallError  `json:"error"`
	DurationMs *float64           `json:"durationMs"`
}
type CollabAgentToolCallThreadItem struct {
	ThreadItem
	Id                string                      `json:"id"`
	Tool              CollabAgentTool             `json:"tool"`
	Status            CollabAgentToolCallStatus   `json:"status"`
	SenderThreadId    string                      `json:"senderThreadId"`
	ReceiverThreadIds []string                    `json:"receiverThreadIds"`
	Prompt            *string                     `json:"prompt"`
	AgentsStates      map[string]CollabAgentState `json:"agentsStates"`
}
type WebSearchThreadItem struct {
	ThreadItem
	Id     string           `json:"id"`
	Query  string           `json:"query"`
	Action *WebSearchAction `json:"action"`
}
type ImageViewThreadItem struct {
	ThreadItem
	Id   string `json:"id"`
	Path string `json:"path"`
}
type EnteredReviewModeThreadItem struct {
	ThreadItem
	Id     string `json:"id"`
	Review string `json:"review"`
}
type ExitedReviewModeThreadItem struct {
	ThreadItem
	Id     string `json:"id"`
	Review string `json:"review"`
}
type ContextCompactionThreadItem struct {
	ThreadItem
	Id string `json:"id"`
}

type ThreadListParams struct {
	Cursor         *string            `json:"cursor,omitempty"`
	Limit          *float64           `json:"limit,omitempty"`
	SortKey        *ThreadSortKey     `json:"sortKey,omitempty"`
	ModelProviders []string           `json:"modelProviders,omitempty"`
	SourceKinds    []ThreadSourceKind `json:"sourceKinds,omitempty"`
	Archived       *bool              `json:"archived,omitempty"`
	Cwd            *string            `json:"cwd,omitempty"`
}

type ThreadListResponse struct {
	Data       []Thread `json:"data"`
	NextCursor *string  `json:"nextCursor"`
}

type ThreadLoadedListParams struct {
	Cursor *string  `json:"cursor,omitempty"`
	Limit  *float64 `json:"limit,omitempty"`
}

type ThreadLoadedListResponse struct {
	Data       []string `json:"data"`
	NextCursor *string  `json:"nextCursor"`
}

type ThreadNameUpdatedNotification struct {
	ThreadId   string  `json:"threadId"`
	ThreadName *string `json:"threadName,omitempty"`
}

type ThreadReadParams struct {
	ThreadId     string `json:"threadId"`
	IncludeTurns bool   `json:"includeTurns"`
}

type ThreadReadResponse struct {
	Thread Thread `json:"thread"`
}

type ThreadResumeParams struct {
	ThreadId               string          `json:"threadId"`
	History                []ResponseItem  `json:"history,omitempty"`
	Path                   *string         `json:"path,omitempty"`
	Model                  *string         `json:"model,omitempty"`
	ModelProvider          *string         `json:"modelProvider,omitempty"`
	Cwd                    *string         `json:"cwd,omitempty"`
	ApprovalPolicy         *AskForApproval `json:"approvalPolicy,omitempty"`
	Sandbox                *SandboxMode    `json:"sandbox,omitempty"`
	Config                 any             `json:"config,omitempty"`
	BaseInstructions       *string         `json:"baseInstructions,omitempty"`
	DeveloperInstructions  *string         `json:"developerInstructions,omitempty"`
	Personality            *Personality    `json:"personality,omitempty"`
	PersistExtendedHistory bool            `json:"persistExtendedHistory"`
}

type ThreadResumeResponse struct {
	Thread          Thread           `json:"thread"`
	Model           string           `json:"model"`
	ModelProvider   string           `json:"modelProvider"`
	Cwd             string           `json:"cwd"`
	ApprovalPolicy  AskForApproval   `json:"approvalPolicy"`
	Sandbox         SandboxPolicy    `json:"sandbox"`
	ReasoningEffort *ReasoningEffort `json:"reasoningEffort"`
}

type ThreadRollbackParams struct {
	ThreadId string  `json:"threadId"`
	NumTurns float64 `json:"numTurns"`
}

type ThreadRollbackResponse struct {
	Thread Thread `json:"thread"`
}

type ThreadSetNameParams struct {
	ThreadId string `json:"threadId"`
	Name     string `json:"name"`
}

type ThreadSetNameResponse struct{}

type ThreadSortKey = string

const (
	ThreadSortKeyCreatedAt ThreadSortKey = "created_at"
	ThreadSortKeyUpdatedAt ThreadSortKey = "updated_at"
)

type ThreadSourceKind = string

const (
	ThreadSourceKindCli                 ThreadSourceKind = "cli"
	ThreadSourceKindVscode              ThreadSourceKind = "vscode"
	ThreadSourceKindExec                ThreadSourceKind = "exec"
	ThreadSourceKindAppServer           ThreadSourceKind = "appServer"
	ThreadSourceKindSubAgent            ThreadSourceKind = "subAgent"
	ThreadSourceKindSubAgentReview      ThreadSourceKind = "subAgentReview"
	ThreadSourceKindSubAgentCompact     ThreadSourceKind = "subAgentCompact"
	ThreadSourceKindSubAgentThreadSpawn ThreadSourceKind = "subAgentThreadSpawn"
	ThreadSourceKindSubAgentOther       ThreadSourceKind = "subAgentOther"
	ThreadSourceKindUnknown             ThreadSourceKind = "unknown"
)

type ThreadStartParams struct {
	Model                  *string         `json:"model,omitempty"`
	ModelProvider          *string         `json:"modelProvider,omitempty"`
	Cwd                    *string         `json:"cwd,omitempty"`
	ApprovalPolicy         *AskForApproval `json:"approvalPolicy,omitempty"`
	Sandbox                *SandboxMode    `json:"sandbox,omitempty"`
	Config                 any             `json:"config,omitempty"`
	BaseInstructions       *string         `json:"baseInstructions,omitempty"`
	DeveloperInstructions  *string         `json:"developerInstructions,omitempty"`
	Personality            *Personality    `json:"personality,omitempty"`
	Ephemeral              *bool           `json:"ephemeral,omitempty"`
	ExperimentalRawEvents  bool            `json:"experimentalRawEvents"`
	PersistExtendedHistory bool            `json:"persistExtendedHistory"`
}

type ThreadStartResponse struct {
	Thread          Thread           `json:"thread"`
	Model           string           `json:"model"`
	ModelProvider   string           `json:"modelProvider"`
	Cwd             string           `json:"cwd"`
	ApprovalPolicy  AskForApproval   `json:"approvalPolicy"`
	Sandbox         SandboxPolicy    `json:"sandbox"`
	ReasoningEffort *ReasoningEffort `json:"reasoningEffort"`
}

type ThreadStartedNotification struct {
	Thread Thread `json:"thread"`
}

type ThreadTokenUsage struct {
	Total              TokenUsageBreakdown `json:"total"`
	Last               TokenUsageBreakdown `json:"last"`
	ModelContextWindow *float64            `json:"modelContextWindow"`
}

type ThreadTokenUsageUpdatedNotification struct {
	ThreadId   string           `json:"threadId"`
	TurnId     string           `json:"turnId"`
	TokenUsage ThreadTokenUsage `json:"tokenUsage"`
}

type ThreadUnarchiveParams struct {
	ThreadId string `json:"threadId"`
}

type ThreadUnarchiveResponse struct {
	Thread Thread `json:"thread"`
}

type ThreadUnarchivedNotification struct {
	ThreadId string `json:"threadId"`
}

type TokenUsageBreakdown struct {
	TotalTokens           float64 `json:"totalTokens"`
	InputTokens           float64 `json:"inputTokens"`
	CachedInputTokens     float64 `json:"cachedInputTokens"`
	OutputTokens          float64 `json:"outputTokens"`
	ReasoningOutputTokens float64 `json:"reasoningOutputTokens"`
}

type ToolRequestUserInputAnswer struct {
	Answers []string `json:"answers"`
}

type ToolRequestUserInputOption struct {
	Label       string `json:"label"`
	Description string `json:"description"`
}

type ToolRequestUserInputParams struct {
	ThreadId  string                         `json:"threadId"`
	TurnId    string                         `json:"turnId"`
	ItemId    string                         `json:"itemId"`
	Questions []ToolRequestUserInputQuestion `json:"questions"`
}

type ToolRequestUserInputQuestion struct {
	Id       string                       `json:"id"`
	Header   string                       `json:"header"`
	Question string                       `json:"question"`
	IsOther  bool                         `json:"isOther"`
	IsSecret bool                         `json:"isSecret"`
	Options  []ToolRequestUserInputOption `json:"options"`
}

type ToolRequestUserInputResponse struct {
	Answers map[string]ToolRequestUserInputAnswer `json:"answers"`
}

type ToolsV2 struct {
	WebSearch *bool `json:"web_search"`
	ViewImage *bool `json:"view_image"`
}

type Turn struct {
	Id     string       `json:"id"`
	Items  []ThreadItem `json:"items"`
	Status TurnStatus   `json:"status"`
	Error  *TurnError   `json:"error"`
}

type TurnCompletedNotification struct {
	ThreadId string `json:"threadId"`
	Turn     Turn   `json:"turn"`
}

type TurnDiffUpdatedNotification struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
	Diff     string `json:"diff"`
}

type TurnError struct {
	Message           string          `json:"message"`
	CodexErrorInfo    *CodexErrorInfo `json:"codexErrorInfo"`
	AdditionalDetails *string         `json:"additionalDetails"`
}

type TurnInterruptParams struct {
	ThreadId string `json:"threadId"`
	TurnId   string `json:"turnId"`
}

type TurnInterruptResponse struct{}

type TurnPlanStep struct {
	Step   string             `json:"step"`
	Status TurnPlanStepStatus `json:"status"`
}

type TurnPlanStepStatus = string

const (
	TurnPlanStepStatusPending    TurnPlanStepStatus = "pending"
	TurnPlanStepStatusInProgress TurnPlanStepStatus = "inProgress"
	TurnPlanStepStatusCompleted  TurnPlanStepStatus = "completed"
)

type TurnPlanUpdatedNotification struct {
	ThreadId    string         `json:"threadId"`
	TurnId      string         `json:"turnId"`
	Explanation *string        `json:"explanation"`
	Plan        []TurnPlanStep `json:"plan"`
}

type TurnStartParams struct {
	ThreadId          string             `json:"threadId"`
	Input             []UserInput        `json:"input"`
	Cwd               *string            `json:"cwd,omitempty"`
	ApprovalPolicy    *AskForApproval    `json:"approvalPolicy,omitempty"`
	SandboxPolicy     *SandboxPolicy     `json:"sandboxPolicy,omitempty"`
	Model             *string            `json:"model,omitempty"`
	Effort            *ReasoningEffort   `json:"effort,omitempty"`
	Summary           *ReasoningSummary  `json:"summary,omitempty"`
	Personality       *Personality       `json:"personality,omitempty"`
	OutputSchema      any                `json:"outputSchema,omitempty"`
	CollaborationMode *CollaborationMode `json:"collaborationMode,omitempty"`
}

type TurnStartResponse struct {
	Turn Turn `json:"turn"`
}

type TurnStartedNotification struct {
	ThreadId string `json:"threadId"`
	Turn     Turn   `json:"turn"`
}

type TurnStatus = string

const (
	TurnStatusCompleted   TurnStatus = "completed"
	TurnStatusInterrupted TurnStatus = "interrupted"
	TurnStatusFailed      TurnStatus = "failed"
	TurnStatusInProgress  TurnStatus = "inProgress"
)

type TurnSteerParams struct {
	ThreadId       string      `json:"threadId"`
	Input          []UserInput `json:"input"`
	ExpectedTurnId string      `json:"expectedTurnId"`
}

type TurnSteerResponse struct {
	TurnId string `json:"turnId"`
}

type UserInputType = string

const (
	UserInputText       UserInputType = "text"
	UserInputImage      UserInputType = "image"
	UserInputLocalImage UserInputType = "localImage"
	UserInputSkill      UserInputType = "skill"
	UserInputMention    UserInputType = "mention"
)

type UserInput struct {
	Type UserInputType `json:"type"`
}
type TextUserInput struct {
	UserInput
	Text         string        `json:"text"`
	TextElements []TextElement `json:"text_elements"`
}
type ImageUserInput struct {
	UserInput
	Url string `json:"url"`
}
type LocalImageUserInput struct {
	UserInput
	Path string `json:"path"`
}
type SkillUserInput struct {
	UserInput
	Name string `json:"name"`
	Path string `json:"path"`
}
type MentionUserInput struct {
	UserInput
	Name string `json:"name"`
	Path string `json:"path"`
}

type WebSearchActionType = string

const (
	WebSearchActionSearch     WebSearchActionType = "search"
	WebSearchActionOpenPage   WebSearchActionType = "openPage"
	WebSearchActionFindInPage WebSearchActionType = "findInPage"
	WebSearchActionOther      WebSearchActionType = "other"
)

type WebSearchAction struct {
	Type WebSearchActionType `json:"type"`
}
type SearchWebSearchAction struct {
	WebSearchAction
	Query   *string  `json:"query"`
	Queries []string `json:"queries"`
}
type OpenPageWebSearchAction struct {
	WebSearchAction
	Url *string `json:"url"`
}
type FindInPageWebSearchAction struct {
	WebSearchAction
	Url     *string `json:"url"`
	Pattern *string `json:"pattern"`
}
type OtherWebSearchAction struct {
	WebSearchAction
}

type WindowsWorldWritableWarningNotification struct {
	SamplePaths []string `json:"samplePaths"`
	ExtraCount  float64  `json:"extraCount"`
	FailedScan  bool     `json:"failedScan"`
}

type WriteStatus = string

const (
	WriteStatusOk           WriteStatus = "ok"
	WriteStatusOkOverridden WriteStatus = "okOverridden"
)
