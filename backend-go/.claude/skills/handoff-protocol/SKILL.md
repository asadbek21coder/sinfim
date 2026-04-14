---
name: handoff-protocol
description: >
  Universal cross-agent handoff workflow for Claude and Codex. Use whenever work spans
  multiple sessions, context is getting large, token budget is low, the current agent is
  about to stop, or the user explicitly asks to continue work in another agent. Keeps the
  real task context in repository files instead of chat history.
---

# Handoff Protocol

The repository is the source of truth for active work. Do not rely on chat history alone.

## Files

- `docs/ai-context/SESSION.md`
- `docs/ai-context/WORKLOG.md`
- `docs/ai-context/HANDOFF.md`

## At Task Start

If `docs/ai-context/HANDOFF.md` exists, read it first.

Then read `docs/ai-context/SESSION.md` if the task looks ongoing or multi-step.

Use these files to recover:

- current goal
- work already completed
- decisions already made
- exact next step

If the files are stale, update them once you understand the current state.

## During Work

Keep `SESSION.md` current for the active task.

Append concise entries to `WORKLOG.md` when any of these happen:

- a meaningful code change is completed
- an important command is run
- a failure or blocker is discovered
- the implementation approach changes

Prefer short, high-signal notes. Do not turn the worklog into a transcript.

## Before Stopping or Switching Agents

You MUST update `HANDOFF.md` when:

- token budget is getting low
- context is large enough that a fresh agent would struggle
- the user wants another agent to continue
- you are ending the session with unfinished work

`HANDOFF.md` must include:

- task summary
- current status
- files touched
- important decisions
- commands run and key results
- open issues
- one exact next step

The next step must be concrete and immediately actionable.

Good:

- `Run npm run build and fix the type error in src/views/UsersView.vue`
- `Implement the repository method in backend/internal/user/repo.go and then run make backend-test`

Bad:

- `Continue working`
- `Finish the feature`

## Scope Rules

- Keep handoff notes stack-agnostic and factual.
- Record what was actually done, not what you intended to do.
- Do not duplicate large diffs or paste long command output.
- Reference exact file paths and command names.

## Reviewer Behavior

Reviewers do not own long-running task state, but if they discover important blockers or follow-up work, they should add a short `WORKLOG.md` entry and update `HANDOFF.md` when the review is the last step in the current session.
