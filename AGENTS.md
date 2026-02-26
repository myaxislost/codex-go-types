# Purpose
This repo purpose is generating go types and structs for app-server rpc protocol

# Repo structure
- ./schema-ts contains type definitions generates with app-server `codex app-server generate-ts --out ./schema-ts`. it's in .gitignore and not to be commited.
- ./protocol - put all the generated go types code here

# IMPORTANT
0. Only Use schema-ts/v2 !!!
1. always use `go test ./protocol` to check for errors
2. you're not to commit anything unless explicitely told so by the user
3. this repo is for type definitions only. DO NOT CODE any logic.
4. focus on generatinc clear types. don't focus on regenerating the results. the final product are the generated types not your scripts.
5. final output MUST be a single file: `./protocol/generated.go`
6. do NOT split generated types across multiple files/folders unless user explicitly asks for that
7. keep `generated.go` as the source of truth for protocol types (overwrite/update this file directly)

# Output contract (must follow)
1. package name in generated file must be `protocol`
2. all exported protocol types from `schema-ts/v2` must exist in `generated.go`
3. include required shared types referenced from `schema-ts/v2` imports (for example `../Personality`, `../Tool`, etc.) so `go test ./protocol` passes
4. use discriminated union base + variant structs for `type` tagged unions (example: `ThreadItem`, `UserInput`, `WebSearchAction`)
5. keep names clean and deduplicated; avoid generated noise suffixes
6. preserve JSON field names via struct tags

# Regeneration workflow (every time)
1. read only from `schema-ts/v2` as the authoritative root
2. generate/update `./protocol/generated.go`
3. run `go test ./protocol`
4. if tests fail, fix the generated types in `generated.go` and rerun tests until passing
5. do not leave temporary generated split files in `./protocol`


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

# Nullability and optional fields
Use pointers to represent nullable/optional non-collection fields.

Examples:
```ts
{ loginId: string | null }
{ path?: string | null }
```

```go
LoginID *string `json:"loginId"`
Path    *string `json:"path,omitempty"`
```
