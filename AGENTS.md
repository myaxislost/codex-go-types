# Purpose
This repo purpose is generating go types and structs for app-server rpc protocol

# Repo structure
- ./schema-ts contains type definitions generates with app-server `codex app-server generate-ts --out ./schema-ts`. it's in .gitignore and not to be commited.
- ./protocol - put all the generated go types code here
- ./internal - if you ever need to write golang code to facilitate coversion/generation it should go here. Use subfolders for better orgranization 
- ./scripts - if you ever need to write scripts/automations they go in this folder. Use subfolders for better orgranization if necessary

# IMPORTANT
0. Only Use schema-ts/v2 !!!
1. always use `go test ./protocol` to check for errors
2. you're not to commit anything unless explicitely told so by the user
3. this repo is for type definitions only. DO NOT CODE any logic.
4. focus on generatinc clear types. don't focus on regenerating the results. the final product are the generated types not your scripts.


# Criteria
1. avoid primitive type aliases, for example TurnInterruptParams.ThreadId should be `string` rather than `type ThreadId = string` .
2. avoid unnecessary suffixes, for example `ThreadStartParamsConfig_ThreadStartParams` should be `ThreadStartParamsConfig`
3. deduplicate repeating types.
4. use strong typing avoid using `interface{}` or `any` when possible

# Type Variants
for stucts like ThreadItem you should generate the base class
```go
type ThreadItem struct {
	ID     string     `json:"id"`
	Type   ItemType   `json:"type"`
}
```

then use it in other variants like
```go
type AgentMessageThreadItem struct {
	ThreadItem
	Text string `json:"text"`
}  
```

# String literal union type
Handle string literal union types
```ts
type Personality = "none" | "friendly" | "pragmatic";
```

Like this:
```go
type Personality = string
const (
	PersonalityNone = "none"
	PersonalityFirendly = "friendly"
	PersonalityPragmatic = "pragmatic"
)
```

# JsonValue type
```ts
export type JsonValue = number | string | boolean | Array<JsonValue> | { [key in string]?: JsonValue } | null;
...

{ structuredContent?: JsonValue }
```
When you see something typed as JsonValue you supposed to use `any`

```go
struct {
	structuredContent any
}
```