# AI Blueprint

Universal AI agent configuration for any tech stack. This is the **source of truth** — every project blueprint instances from it.

---

## What this is

A complete, stack-agnostic AI agent system that gives Claude Code (and Codex) structured, specialized agents that plan, implement, review, and orchestrate work — without making up conventions or getting confused about roles.

It provides universal skills that work for any stack, plus templates for the stack-specific pieces that each project fills in.

It also includes a universal handoff system so Claude and Codex can continue each other's work through repository files when a session ends or token budget runs low.

---

## Structure

```
ai-blueprint/
├── README.md                        # This file
├── AGENTS.md.template               # Dispatcher rules — copy to each project, fill in agents
├── .claude/
│   ├── settings.json                # Enables agent teams
│   ├── rules/
│   │   └── RULES.md                 # Points to AGENTS.md (same pattern for all projects)
│   ├── agents/
│   │   ├── developer.md.template    # Developer agent — fill in stack-specific content
│   │   └── reviewer.md.template     # Reviewer agent — fill in stack-specific content
│   └── skills/
│       ├── orchestrate/
│       │   └── SKILL.md             # Universal — task orchestration, git workflow, agent dispatch
│       ├── backplan/
│       │   └── SKILL.md             # Universal — architecture planning protocol (stack-agnostic)
│       ├── red-green-tdd/
│       │   └── SKILL.md             # Universal — TDD protocol with hard gates
│       └── handoff-protocol/
│           └── SKILL.md             # Universal — cross-agent session handoff workflow
├── docs/
│   └── ai-context/
│       ├── SESSION.md               # Current task and status snapshot
│       ├── WORKLOG.md               # Append-only progress log
│       └── HANDOFF.md               # Cross-agent transfer note
└── .codex/
    ├── config.toml                  # Codex shared agent settings
    ├── agents/
    │   ├── developer.toml.template  # Codex developer agent — fill in stack-specific content
    │   └── reviewer.toml.template   # Codex reviewer agent — fill in stack-specific content
    └── skills/
        ├── orchestrate/
        │   └── SKILL.md             # Universal Codex skill mirror
        ├── backplan/
        │   └── SKILL.md             # Universal Codex skill mirror
        ├── red-green-tdd/
        │   └── SKILL.md             # Universal Codex skill mirror
        └── handoff-protocol/
            └── SKILL.md             # Universal Codex skill mirror
```

---

## Universal skills

These universal skills are **identical across all stacks** — copy them to any blueprint without modification:

| Skill | What it does |
|---|---|
| `orchestrate` | Decomposes plans into tasks, dispatches agents, manages git workflow. The orchestrator never writes code. |
| `backplan` | Produces implementation plans (what to build, architectural decisions, data model, API contracts, edge cases) before any code is written. |
| `red-green-tdd` | Strict RED-GREEN-REFACTOR protocol with hard gates. Auto-triggers for any logic-bearing code. |
| `handoff-protocol` | Keeps active work in `docs/ai-context/` so Claude and Codex can continue each other's sessions reliably. |

## Cross-Agent Handoff

Use `docs/ai-context/` as the shared working memory between agents.

| File | Purpose |
|---|---|
| `docs/ai-context/SESSION.md` | Snapshot of the active task, decisions, blockers, and exact next step |
| `docs/ai-context/WORKLOG.md` | Append-only log of meaningful progress, command results, and discoveries |
| `docs/ai-context/HANDOFF.md` | Clean handoff note for the next Claude or Codex session |

Recommended workflow:

1. At task start, read `HANDOFF.md` and `SESSION.md` if they exist
2. During work, keep `SESSION.md` current and append concise `WORKLOG.md` entries
3. Before stopping or switching agents, update `HANDOFF.md` with one exact next step

---

## Stack-specific pieces (fill in per project)

| File | What to fill in |
|---|---|
| `AGENTS.md` | Agent names for `[impl]` and `[review]` tasks |
| `.claude/agents/developer.md` | Stack guidelines skill, verification command, specific rules |
| `.claude/agents/reviewer.md` | Stack guidelines skill, review checklist |
| `.codex/agents/developer.toml` | Mirror of the Claude developer agent in TOML |
| `.codex/agents/reviewer.toml` | Mirror of the Claude reviewer agent in TOML |
| `.claude/skills/{stack-guidelines}` | Stack-specific Claude skill |
| `.codex/skills/{stack-guidelines}` | Stack-specific Codex skill mirror |

---

## How to use in a new project blueprint

1. Copy the four universal skills from `.claude/skills/` to your blueprint's `.claude/skills/`
2. Copy the same four universal skills from `.codex/skills/` to your blueprint's `.codex/skills/`
3. Copy `.codex/config.toml` and `.codex/agents/` templates
4. Copy `.claude/settings.json` and `.claude/rules/RULES.md`
5. Copy `docs/ai-context/` into the blueprint
6. Copy `AGENTS.md.template` to `AGENTS.md` and fill in your agent names
7. Fill in developer and reviewer agents with your stack-specific content
8. Create your stack's `{stack}-guidelines` skill in both `.claude/skills/` and `.codex/skills/`

---

## Instances

| Blueprint | Stack | Developer agent | Reviewer agent | Stack guidelines |
|---|---|---|---|---|
| `kotlin-spring-blueprint` | Kotlin + Spring Boot | `kotlin-spring-developer` | `kotlin-spring-reviewer` | `kotlin-spring-guidelines` |
| `go-enterprise-blueprint` | Go + Fiber | `architect`, `go-coder`, `go-tester` | `reviewer` | `backend-guidelines` |
| `vue-blueprint-web` | Vue 3 + TypeScript | `vue-coder` | `vue-reviewer` | `frontend-guidelines` |
| `go-vue-monorepo-blueprint` | Go + Vue | `go-coder`, `vue-coder` | `reviewer`, `vue-reviewer` | both |
