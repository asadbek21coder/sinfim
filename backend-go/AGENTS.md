# AGENT INSTRUCTIONS

## Project Context

Read these to gain project context:

    - {PROJECT_ROOT}/README.md
    - {PROJECT_ROOT}/docs/specs/modules/{module}/overview.md
    - {PROJECT_ROOT}/docs/specs/modules/{module}/usecases/{domain}/{operation}.md
    - {PROJECT_ROOT}/docs/ai-context/HANDOFF.md
    - {PROJECT_ROOT}/docs/ai-context/SESSION.md

## Commands

Read {PROJECT_ROOT}/Makefile for available commands.

## Guidelines

ALWAYS follow the project guidelines:

    - stack guidelines: Skills "backend-guidelines"
    - workflow: Skills "red-green-tdd", "orchestrate", "backplan", "handoff-protocol"
    - project-specific: {PROJECT_ROOT}/docs/

## Pull Request Readiness

- Every commit must compile, pass lint, and pass all tests. No "fix later" commits.
- If `make lint`, `make test`, or `make test-system` does not pass, the code is not done.

## Subagent Selection

| Task type  | Agent         |
| ---------- | ------------- |
| `[impl]`   | `go-coder`    |
| `[review]` | `reviewer`    |
| `[fix]`    | same as `[impl]` |

## Parallel Dispatch

Backend tasks are generally safe to parallelize when they touch different modules or use cases with no shared state.
