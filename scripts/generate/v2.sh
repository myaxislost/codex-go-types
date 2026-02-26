#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
cd "$ROOT_DIR"

if ! command -v go-jsonschema >/dev/null 2>&1; then
  echo "go-jsonschema not found in PATH" >&2
  exit 1
fi

normalize_name() {
  local base=$1
  base=$(printf '%s' "$base" | sed -E 's/([A-Z]+)([A-Z][a-z])/\1_\2/g; s/([a-z0-9])([A-Z])/\1_\2/g')
  base=$(printf '%s' "$base" | tr '[:upper:]' '[:lower:]' | sed -E 's/[^a-z0-9]+/_/g; s/^_+//; s/_+$//')
  printf '%s' "$base"
}

is_failing_schema() {
  local path=$1
  case "$path" in
    schema-json/ClientRequest.json) return 0 ;;
    schema-json/ServerNotification.json) return 0 ;;
    schema-json/v2/ConfigReadResponse.json) return 0 ;;
    schema-json/v2/ReviewStartResponse.json) return 0 ;;
    schema-json/v2/ThreadForkResponse.json) return 0 ;;
    schema-json/v2/ThreadListResponse.json) return 0 ;;
    schema-json/v2/ThreadReadResponse.json) return 0 ;;
    schema-json/v2/ThreadResumeResponse.json) return 0 ;;
    schema-json/v2/ThreadRollbackResponse.json) return 0 ;;
    schema-json/v2/ThreadStartResponse.json) return 0 ;;
    schema-json/v2/ThreadStartedNotification.json) return 0 ;;
    schema-json/v2/ThreadUnarchiveResponse.json) return 0 ;;
    schema-json/v2/TurnCompletedNotification.json) return 0 ;;
    schema-json/v2/TurnStartResponse.json) return 0 ;;
    schema-json/v2/TurnStartedNotification.json) return 0 ;;
    *) return 1 ;;
  esac
}

rm -rf protocol/root protocol/v2
mkdir -p protocol/root protocol/v2

failures=()

generate_one() {
  local schema=$1
  local family=$2

  local filename
  filename=$(basename "$schema" .json)
  local pkg
  pkg=$(normalize_name "$filename")
  local outdir="protocol/$family/$pkg"
  mkdir -p "$outdir"

  if is_failing_schema "$schema"; then
    failures+=("$schema")
    return 0
  fi

  if ! go-jsonschema \
    --only-models \
    --struct-name-from-title \
    --tags json \
    --package "$pkg" \
    --output "$outdir/types.gen.go" \
    "$schema" >/tmp/gojsonschema.log 2>&1; then
    echo "unexpected generator failure for $schema" >&2
    cat /tmp/gojsonschema.log >&2
    failures+=("$schema")
    rm -f "$outdir/types.gen.go"
  fi
}

for schema in schema-json/*.json; do
  base=$(basename "$schema")
  if [[ "$base" == "codex_app_server_protocol.schemas.json" ]]; then
    continue
  fi
  generate_one "$schema" root
done

for schema in schema-json/v2/*.json; do
  generate_one "$schema" v2
done

if ((${#failures[@]} > 0)); then
  printf '%s\n' "${failures[@]}" | sort -u > protocol/.generation_failures.txt
  echo "generation completed with known/manual schemas listed in protocol/.generation_failures.txt"
else
  rm -f protocol/.generation_failures.txt
  echo "generation completed with no failures"
fi
