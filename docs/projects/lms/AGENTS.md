# AGENT INSTRUCTIONS

## Project Context

Read these to gain project context:

    - {PROJECT_ROOT}/docs/startup-idea.md
    - {PROJECT_ROOT}/docs/auth-access-notes.md
    - {PROJECT_ROOT}/docs/ai-context/HANDOFF.md
    - {PROJECT_ROOT}/docs/ai-context/SESSION.md

## Commands

This is currently a planning workspace, not the implementation repository.
When the actual LMS repo is created, read its Makefile, package.json, justfile, or equivalent task runner before running commands.

## Guidelines

ALWAYS follow the project guidelines:

    - stack guidelines: use the selected backend/frontend blueprint guidelines once the implementation repo exists
    - workflow: Skills "red-green-tdd", "orchestrate", "backplan", "handoff-protocol"
    - project-specific: {PROJECT_ROOT}/docs/

## Pull Request Readiness

- Every commit must compile, pass lint, and pass all tests. No "fix later" commits.
- Verification command is not defined yet because the implementation repo has not been created.
- For the planned Go + Vue stack, expect backend `make fmt && make lint && make test && make test-system` and frontend `npm run type-check && npm run lint` once the repo exists.

## Subagent Selection

| Task type  | Agent                   |
| ---------- | ----------------------- |
| `[impl]`   | implementation agent for the chosen stack |
| `[review]` | reviewer agent for the chosen stack |
| `[fix]`    | same as `[impl]`        |

## Parallel Dispatch

Tasks are generally safe to parallelize when they touch different modules or entities with no shared state.
