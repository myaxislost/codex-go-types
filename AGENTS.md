# Purpose
This repo purpose is generating go types and structs for app-server rpc protocol

# Repo structure
- ./schema-json contains type definitions generates with app-server `codex app-server generate-json-schema --out ./schema-json`. it's in .gitignore and not to be commited.
- ./protocol - put all the generated go types code here. 
- ./internal - if you ever need to write golang code to facilitate coversion/generation it should go here. Use subfolders for better orgranization 
- ./scripts - if you ever need to write scripts/automations they go in this folder. Use subfolders for better orgranization if necessary

# IMPORTANT
0. SKIP ./schema-json/v1
1. always use `go test ./protocol` to check for errors
2. you need to generate initial code using `go-jsonschema`. then fix any errors and organize
3. with oneOf/anyOf shapes prefer explicit structs. Follow appropriate naming convention  
4. you're not to commit anything unless explicitely told so by the user
5. this repo is for type definitions only. DO NOT CODE any logic.
