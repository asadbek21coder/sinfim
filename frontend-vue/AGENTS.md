# AGENT INSTRUCTIONS

## Project Context

Read these to gain project context:

    - {PROJECT_ROOT}/README.md
    - {PROJECT_ROOT}/src/views/
    - {PROJECT_ROOT}/src/types/
    - {PROJECT_ROOT}/../docs/ai-context/HANDOFF.md
    - {PROJECT_ROOT}/../docs/ai-context/SESSION.md

## Commands

Read {PROJECT_ROOT}/package.json for available scripts.

## Guidelines

ALWAYS follow the project guidelines:

    - stack guidelines: Skills "frontend-guidelines"
    - workflow: Skills "red-green-tdd", "orchestrate", "backplan", "handoff-protocol"
    - project-specific: {PROJECT_ROOT}/README.md

## Pull Request Readiness

- Every change must build cleanly. No "fix later" commits.
- If `npm run build` does not pass, the code is not done.

## Subagent Selection

| Task type  | Agent          |
| ---------- | -------------- |
| `[impl]`   | `vue-coder`    |
| `[review]` | `vue-reviewer` |
| `[fix]`    | same as `[impl]` |

## Parallel Dispatch

Frontend tasks are generally safe to parallelize when they touch different domains or views with no shared files.
