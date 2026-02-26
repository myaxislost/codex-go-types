package protocol

type Method = string

const (
	// ---- Thread lifecycle (client -> server requests) ----
	MethodInitialize                     Method = "initialize"
	MethodInitialized                    Method = "initialized"
	MethodThreadStart                    Method = "thread/start"
	MethodThreadStarted                  Method = "thread/started"
	MethodThreadResume                   Method = "thread/resume"
	MethodThreadFork                     Method = "thread/fork"
	MethodThreadArchive                  Method = "thread/archive"
	MethodThreadUnsubscribe              Method = "thread/unsubscribe"
	MethodThreadNameSet                  Method = "thread/name/set"
	MethodThreadUnarchive                Method = "thread/unarchive"
	MethodThreadCompactStart             Method = "thread/compact/start"
	MethodThreadBackgroundTerminalsClean Method = "thread/backgroundTerminals/clean"
	MethodThreadRollback                 Method = "thread/rollback"
	MethodThreadList                     Method = "thread/list"
	MethodThreadLoadedList               Method = "thread/loaded/list"
	MethodThreadRead                     Method = "thread/read"

	// ---- Thread lifecycle (server -> client notifications) ----
	MethodThreadStatusChanged     Method = "thread/status/changed"
	MethodThreadArchived          Method = "thread/archived"
	MethodThreadUnarchived        Method = "thread/unarchived"
	MethodThreadClosed            Method = "thread/closed"
	MethodThreadNameUpdated       Method = "thread/name/updated"
	MethodThreadTokenUsageUpdated Method = "thread/tokenUsage/updated"
	// Deprecated (still emitted in some builds)
	MethodThreadCompacted Method = "thread/compacted"

	// ---- Turn lifecycle (client -> server requests) ----
	MethodTurnStart     Method = "turn/start"
	MethodTurnSteer     Method = "turn/steer"
	MethodTurnInterrupt Method = "turn/interrupt"

	// ---- Turn lifecycle (server -> client notifications) ----
	MethodTurnStarted     Method = "turn/started"
	MethodTurnCompleted   Method = "turn/completed"
	MethodTurnDiffUpdated Method = "turn/diff/updated"
	MethodTurnPlanUpdated Method = "turn/plan/updated"

	// ---- Item lifecycle + streaming (server -> client notifications) ----
	MethodItemStarted                             Method = "item/started"
	MethodItemCompleted                           Method = "item/completed"
	MethodItemAgentMessageDelta                   Method = "item/agentMessage/delta"
	MethodItemPlanDelta                           Method = "item/plan/delta"
	MethodItemCommandExecutionOutputDelta         Method = "item/commandExecution/outputDelta"
	MethodItemCommandExecutionTerminalInteraction Method = "item/commandExecution/terminalInteraction"
	MethodItemFileChangeOutputDelta               Method = "item/fileChange/outputDelta"
	MethodItemMcpToolCallProgress                 Method = "item/mcpToolCall/progress"
	MethodItemReasoningSummaryTextDelta           Method = "item/reasoning/summaryTextDelta"
	MethodItemReasoningSummaryPartAdded           Method = "item/reasoning/summaryPartAdded"
	MethodItemReasoningTextDelta                  Method = "item/reasoning/textDelta"

	// ---- Server -> client requests (approval + tool interactions) ----
	MethodItemCommandExecutionRequestApproval Method = "item/commandExecution/requestApproval"
	MethodItemFileChangeRequestApproval       Method = "item/fileChange/requestApproval"
	MethodItemToolRequestUserInput            Method = "item/tool/requestUserInput"
	MethodItemToolCall                        Method = "item/tool/call"

	// ---- MCP / config (client -> server requests) ----
	MethodMcpServerOauthLogin   Method = "mcpServer/oauth/login"
	MethodConfigMcpServerReload Method = "config/mcpServer/reload"
	MethodMcpServerStatusList   Method = "mcpServerStatus/list"

	// ---- MCP / auth (server -> client notifications) ----
	MethodMcpServerOauthLoginCompleted Method = "mcpServer/oauthLogin/completed"

	// ---- Account/auth (client -> server requests) ----
	MethodAccountRead           Method = "account/read"
	MethodAccountLoginStart     Method = "account/login/start"
	MethodAccountLoginCancel    Method = "account/login/cancel"
	MethodAccountLogout         Method = "account/logout"
	MethodAccountRateLimitsRead Method = "account/rateLimits/read"

	// ---- Account/auth (server -> client requests) ----
	MethodAccountChatgptAuthTokensRefresh Method = "account/chatgptAuthTokens/refresh"

	// ---- Account/auth (server -> client notifications) ----
	MethodAccountUpdated           Method = "account/updated"
	MethodAccountRateLimitsUpdated Method = "account/rateLimits/updated"
	MethodAccountLoginCompleted    Method = "account/login/completed"

	// ---- Skills / apps (client -> server requests) ----
	MethodSkillsList         Method = "skills/list"
	MethodSkillsRemoteList   Method = "skills/remote/list"
	MethodSkillsRemoteExport Method = "skills/remote/export"
	MethodSkillsConfigWrite  Method = "skills/config/write"
	MethodAppList            Method = "app/list"

	// ---- Apps (server -> client notifications) ----
	MethodAppListUpdated Method = "app/list/updated"

	// ---- Review (client -> server request) ----
	MethodReviewStart Method = "review/start"

	// ---- Models + features ----
	MethodModelList               Method = "model/list"
	MethodExperimentalFeatureList Method = "experimentalFeature/list"
	MethodCollaborationModeList   Method = "collaborationMode/list"

	// ---- Model routing (server -> client notifications) ----
	MethodModelRerouted Method = "model/rerouted"

	// ---- Feedback + command exec + config ----
	MethodFeedbackUpload            Method = "feedback/upload"
	MethodCommandExec               Method = "command/exec"
	MethodConfigRead                Method = "config/read"
	MethodConfigValueWrite          Method = "config/value/write"
	MethodConfigBatchWrite          Method = "config/batchWrite"
	MethodConfigRequirementsRead    Method = "configRequirements/read"
	MethodExternalAgentConfigDetect Method = "externalAgentConfig/detect"
	MethodExternalAgentConfigImport Method = "externalAgentConfig/import"

	// ---- Windows sandbox ----
	MethodWindowsSandboxSetupStart     Method = "windowsSandbox/setupStart"
	MethodWindowsSandboxSetupCompleted Method = "windowsSandbox/setupCompleted"
	MethodWindowsWorldWritableWarning  Method = "windows/worldWritableWarning"

	// ---- Thread realtime (experimental) ----
	MethodThreadRealtimeStart       Method = "thread/realtime/start"
	MethodThreadRealtimeAppendAudio Method = "thread/realtime/appendAudio"
	MethodThreadRealtimeAppendText  Method = "thread/realtime/appendText"
	MethodThreadRealtimeStop        Method = "thread/realtime/stop"

	MethodThreadRealtimeStarted          Method = "thread/realtime/started"
	MethodThreadRealtimeItemAdded        Method = "thread/realtime/itemAdded"
	MethodThreadRealtimeOutputAudioDelta Method = "thread/realtime/outputAudio/delta"
	MethodThreadRealtimeError            Method = "thread/realtime/error"
	MethodThreadRealtimeClosed           Method = "thread/realtime/closed"

	// ---- Fuzzy file search ----
	MethodFuzzyFileSearch              Method = "fuzzyFileSearch"
	MethodFuzzyFileSearchSessionStart  Method = "fuzzyFileSearch/sessionStart"
	MethodFuzzyFileSearchSessionUpdate Method = "fuzzyFileSearch/sessionUpdate"
	MethodFuzzyFileSearchSessionStop   Method = "fuzzyFileSearch/sessionStop"

	MethodFuzzyFileSearchSessionUpdated   Method = "fuzzyFileSearch/sessionUpdated"
	MethodFuzzyFileSearchSessionCompleted Method = "fuzzyFileSearch/sessionCompleted"

	// ---- Generic server notifications ----
	MethodError             Method = "error"
	MethodDeprecationNotice Method = "deprecationNotice"
	MethodConfigWarning     Method = "configWarning"

	// ---- Internal-only notification (Codex Cloud) ----
	MethodRawResponseItemCompleted Method = "rawResponseItem/completed"

	// ---- Test-only method ----
	MethodMockExperimentalMethod Method = "mock/experimentalMethod"
)
