# Claude Agent System — How It Works and Why

This document explains every piece of the `.claude/` directory used in the Go enterprise blueprint.
It is the reference for understanding the system deeply and for replicating it in other blueprints.

---

## Big Picture

Claude Code reads a project directory and picks up configuration from two places:

1. **`CLAUDE.md`** — loaded automatically at the start of every session, into the main agent's context
2. **`.claude/`** — directory with four types of artifacts:
   - `settings.json` — project-level Claude Code settings
   - `agents/` — specialized sub-agents with scoped tools and focused system prompts
   - `skills/` — reference documents auto-loaded into agents that declare them
   - `commands/` — slash commands that trigger multi-step Claude procedures

Together they form a **layered context system**: CLAUDE.md gives orientation, skills give deep rules, agents divide labor, commands automate repeated procedures.

---

## 1. `settings.json`

```json
{
  "env": {
    "CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS": "1"
  }
}
```

### What it does
Sets an environment variable that Claude Code reads at startup. The `CLAUDE_CODE_EXPERIMENTAL_AGENT_TEAMS=1` flag enables **named agent teams** — it makes the agent files in `.claude/agents/` available as `subagent_type` values when spawning sub-agents via the `Task` tool.

### Why it matters
Without this flag, the `Task` tool only spawns generic agents with no specialized context. With it, you can delegate work to `architect`, `go-coder`, `go-tester`, etc. — each pre-loaded with its own system prompt, tool restrictions, and skills. The main agent becomes an **orchestrator**, not a doer.

### Why only one setting
The `.claude/settings.json` is intentionally minimal. It sets env config that must be present before any agent runs. Everything else (rules, guidelines, tool restrictions) lives in the individual agent files where it's co-located with the context it governs.

---

## 2. `CLAUDE.md`

The root-level `CLAUDE.md` is **the first thing Claude reads** in every session. It answers three questions:

1. What is this project and how do I work in it?
2. Which specialized agents exist and when do I use them?
3. What are the verification gates before anything ships?

### Structure of the Go blueprint's CLAUDE.md

```
## Skills        → what to load before doing backend work
## Agents        → table of agents with roles and tools
## Pipeline      → architect → system-analyst → coder+tester → reviewer
## Rules         → keep docs, code, and tests in sync
## Verification  → make lint, make test, make test-system
## Workflow      → TaskCreate for planning, parallel agents for execution
```

### Design choices

**Skills section comes first.** Before touching any code, load `backend-guidelines`. This is a contract: no agent should improvise patterns. If the skill is loaded, the rules are loaded. The CLAUDE.md does not repeat the rules — it just says "load the skill."

**Pipeline is explicit.** The CLAUDE.md spells out the exact sequence: architect first (read-only, no temptation to code), system-analyst second (UC doc is the source of truth before implementation), coder and tester in parallel (both can work from the UC doc independently), reviewer last (gating step). Without this, a main agent might jump straight to coding.

**Verification is mandatory, not optional.** The three commands (`make lint`, `make test`, `make test-system`) appear in their own section with the rule "NEVER deliver code that hasn't passed all three." This makes the bar explicit — Claude cannot rationalize skipping it.

**Agents table is a lookup, not a rulebook.** The CLAUDE.md does not re-explain each agent's rules (those live in the agent files). It only says who does what and which tools each uses — enough to delegate correctly.

---

## 3. Skills — `.claude/skills/{name}/SKILL.md`

A skill is a reference document that gets **automatically injected into an agent's context** when the agent declares it in its frontmatter:

```yaml
---
name: go-coder
skills: backend-guidelines
---
```

Claude Code reads the corresponding `SKILL.md` and prepends it to the agent's system prompt.

### What `backend-guidelines` contains

The skill is a **condensed, agent-optimized version** of all 16 guideline files in `docs/guidelines/`. It covers:

- Architecture (file structure, layers, cross-module rules)
- API design (GET/POST only, response formats, no path params)
- Code style (naming, structs, comments, error handling syntax)
- Use case patterns (types, Execute method, doc-code sync)
- Controllers (forwarders only, no logic)
- PBLC (when to use, design freedom)
- Infrastructure (repogen, filter functions, filter struct conventions)
- DI containers (pattern, layer responsibilities)
- Error handling (errx, layer responsibilities, inline return)
- Validation (per layer)
- Transaction management (UOW, owned vs borrowed)
- List manipulation (pagination, sorting, filter order)
- Testing (system tests, GIVEN-WHEN-THEN, state helpers)
- Observability (logger, auto-framework handling)
- DB migrations (goose, naming, query order)

### Why a skill instead of inline in CLAUDE.md

**Token efficiency.** Every session starts with a limited context window. If all 16 guideline files were embedded in CLAUDE.md, they'd consume ~4,000+ tokens before the user says a word. With skills, guidelines only load for agents that need them — the main orchestrator, the reviewer, and the coder each get them; the architect does NOT (it only plans structure, not code style).

**Separation of concerns.** The CLAUDE.md tells Claude *how to work*. The skill tells Claude *how to write code*. They're different levels of concern. Keeping them separate makes each easier to maintain.

**Condensed format.** The skill is NOT a copy of the docs. It's written for an agent that already understands Go — no lengthy explanations, just rules and patterns. Human-readable docs in `docs/guidelines/` have more context; the skill has the rules in the most compact form possible.

**Sync is enforced by `/enhance-docs`.** The `/enhance-docs` command requires updating the skill whenever the underlying docs change. This prevents drift between what developers read and what agents follow.

### Which agents load the skill

| Agent | Loads backend-guidelines? | Reason |
|---|---|---|
| architect | No | Plans structure, not code details |
| system-analyst | No | Writes docs, not code |
| go-coder | Yes | Needs all code patterns |
| go-tester | Yes | Needs testing patterns + code conventions |
| reviewer | Yes | Needs all rules to check against |

---

## 4. Agents — `.claude/agents/*.md`

Each agent file is a markdown file with **YAML frontmatter** followed by a system prompt:

```yaml
---
name: agent-name
description: one sentence — when to use this agent
tools: Read, Glob, Grep
model: opus
skills: backend-guidelines
disallowedTools: Write, Edit
---

System prompt goes here...
```

### Frontmatter fields

| Field | Purpose |
|---|---|
| `name` | The `subagent_type` value used when spawning via Task tool |
| `description` | Claude Code displays this in agent selection; also helps the main agent know when to delegate |
| `tools` | Whitelist of tools this agent can use |
| `model` | `opus` for complex reasoning, `sonnet` for implementation |
| `skills` | Skill names to auto-load (space-separated) |
| `disallowedTools` | Explicit blocklist — stronger than omitting from `tools` |

### Agent deep-dives

---

#### `architect` — Plan, never code

```yaml
tools: Read, Glob, Grep
model: opus
```

**Purpose:** Takes a requirement and produces a precise implementation plan that other agents follow.

**Why opus?** Architectural planning requires understanding the full project, reasoning about cross-layer dependencies, identifying UOW boundaries, and spotting cross-module concerns (Portal/Embassy). Opus handles complex multi-step reasoning better. The cost is worth it because architect runs once per feature, not continuously.

**Why read-only tools?** This is not just a convention — it is a *design constraint*. By giving the architect only `Read, Glob, Grep`, it is physically impossible for it to write code even if it wanted to. This enforces the plan-first principle. If the architect could write files, it might skip planning and start scaffolding — which undermines the whole pipeline.

**What a good architect plan includes:**
- Exact file paths to create or modify
- Layer-by-layer breakdown in implementation order (bottom-up)
- Dependencies and portals required
- New domain entities and interfaces
- Migration needs (new tables, indexes)
- UOW boundaries (where to Start and Apply)
- New error codes
- What can be parallelized

---

#### `system-analyst` — Docs first, always

```yaml
tools: Read, Write, Edit, Glob, Grep
model: sonnet
```

**Purpose:** Writes and maintains use case documents, ERDs, flow docs, and module overviews. The documents it creates are the source of truth — implementation follows from them, never the reverse.

**Why no Bash?** Documentation work never needs shell commands. Excluding Bash removes an entire class of accidental side effects (running make, touching infra, triggering tests). It's a precision tool restriction.

**Why separate from architect?** The architect plans *how* to implement (which layers, which files, what order). The system-analyst writes *what* to implement (the formal spec that frontend devs and testers also read). These are different audiences and different output types. The architect's output is consumed internally by the coding pipeline; the analyst's output is a durable project artifact.

**Key doc-code sync principle:** Execute method comments in code must mirror the steps in the UC document exactly. This means the UC doc is not just documentation — it becomes embedded in the code itself. If the doc changes, the code comments must change. If implementation reveals an edge case, the doc is updated first. The system-analyst owns this canonical record.

---

#### `go-coder` — Implement only, follow the plan

```yaml
tools: Read, Write, Edit, Bash, Glob, Grep
model: sonnet
skills: backend-guidelines
```

**Purpose:** Takes a UC document and the architect's plan, implements the code layer-by-layer (bottom-up), following every pattern in `backend-guidelines` exactly.

**Why full tools including Bash?** Code implementation requires: writing files (Write, Edit), reading existing patterns (Read, Glob, Grep), and running verification commands like `make fmt`, `make lint` to check work. Bash is essential.

**Why not write tests?** "Do not write tests — that's go-tester's job." Separating these responsibilities means: (a) coder and tester can run in parallel after the UC doc exists, (b) tester derives tests from the *spec*, not from reading the implementation — this avoids test bias where tests just mirror code paths instead of verifying behavior.

**Bottom-up implementation order enforced:** The agent is explicitly told to implement in order: Migrations → Domain → Infra → PBLC → Use Case → Controller → DI Containers. This is not just a suggestion — it's in the system prompt. The reason is dependency direction: each layer depends on the one below. Building top-down means you can't compile or test incrementally.

**Doc-code sync via Execute comments:** Every step in the UC document becomes a comment in the Execute method. The code under each comment implements that step. This is enforced so that anyone reading the code can trace it back to the spec without looking up the UC doc.

---

#### `go-tester` — Test documented behavior, not implementation

```yaml
tools: Read, Write, Edit, Bash, Glob, Grep
model: sonnet
skills: backend-guidelines
```

**Purpose:** Writes system tests for every use case and unit tests for `pkg/` components. Tests are derived from UC documents — not from reading the implementation.

**Why same tools as go-coder?** Testers need to write test files (Write, Edit), read UC docs and existing test patterns (Read, Glob, Grep), and run test suites (Bash: `make test`, `make test-system`).

**GIVEN-WHEN-THEN structure:** Every system test follows this pattern. GIVEN sets up state using state helpers (no hardcoded IDs, no cross-test dependencies). WHEN triggers the UC (via typed trigger functions for each UC type). THEN asserts response and DB state. This structure makes tests readable and maps directly to the UC document's success/failure scenarios.

**Why state helpers?** Direct DB inserts in tests would couple tests to the schema. State helpers (`GivenAdmins(t, ...)`, `GetAdminByUsername(t, ...)`) abstract the setup and validation. If the schema changes, only the helpers change — not every test. They also use `map[string]any` for flexible field overrides, which keeps them general enough to cover all scenarios.

**Why `database.Empty(t)` on every test?** Test isolation. If one test creates data and the next test queries it, the second test's outcome depends on the first. `database.Empty(t)` truncates all tables before the test, ensuring the test runs in a clean state regardless of what ran before it.

**Deriving tests from docs:**

| Doc element | Test type |
|---|---|
| Success execute steps | One comprehensive test (response + DB state + side effects) |
| Simple failures (validation) | Table-driven test |
| Complex failures (inactive user, limit exceeded) | Dedicated test per scenario |
| Start/Apply UOW present | Atomicity test (partial failure = no state change) |

---

#### `reviewer` — No mercy, no modifications

```yaml
tools: Read, Glob, Grep, Bash
disallowedTools: Write, Edit
model: sonnet
skills: backend-guidelines
```

**Purpose:** Reviews all code against every guideline. Runs lint and tests. Reports findings in a structured format. Cannot modify files.

**Why `disallowedTools: Write, Edit` explicitly?** This is a double lock. The `tools` field could be set to only `Read, Glob, Grep, Bash` — but `disallowedTools` adds an explicit prohibition even if the tools field were ever expanded. It makes the read-only contract unmistakable in the file itself.

**Why does the reviewer run lint/tests?** A code review that doesn't run verification is incomplete. The reviewer runs `make lint`, `make test`, `make test-system` as part of its review — not as an optional step. This means the reviewer confirms the code actually works, not just that it looks correct.

**Report format: CRITICAL > WARNING > SUGGESTION.** Structured severity prevents the reviewer from presenting a wall of findings with no prioritization. CRITICAL items must be fixed before delivery. WARNING items should be fixed. SUGGESTION items are nice-to-have. This maps directly to how a dev would triage the feedback.

**Why not let go-coder self-review?** Self-review misses violations because the author already has a mental model of "this is correct." A separate agent has no attachment to the code and no prior decisions to rationalize. It applies the checklist mechanically and comprehensively.

---

## 5. Commands — `.claude/commands/*.md`

Commands are markdown files that Claude executes when you type `/command-name` in a session. The file content becomes the prompt. `$ARGUMENTS` is replaced with whatever the user typed after the command name.

### `/generate-module $ARGUMENTS`

A 7-step procedure that creates a complete new module from scratch:

1. Validate module name input
2. Ask user for module purpose (to fill overview.md — doc-first, even in scaffolding)
3. Create the full directory structure with all layers
4. Create the portal interface
5. Create documentation (overview.md, ERD.md, usecases/ dir)
6. Create test directories
7. Wire the module into app.go, run.go, shutdown.go, portal/container.go

**Why include file templates?** Without templates, Claude might invent field names, import paths, or patterns that don't match existing modules. The command embeds exact templates for every file (module.go, domain/container.go, embassy.go, etc.). This prevents pattern drift and ensures the scaffolded module is immediately compilable.

**Why does it ask for purpose before creating files?** Because the first output is `overview.md` — the module's documentation. Creating the doc from the user's description before scaffolding code reinforces the document-first principle even for new modules.

**Why wire the application layer?** A module that isn't wired into app.go, run.go, and portal/container.go is unusable. Many scaffolding tools stop at directory creation. This command does the full integration so the module can start immediately.

### `/enhance-docs`

Integrates a new concept into the correct documentation file AND syncs the change to the corresponding section of `backend-guidelines` SKILL.md.

**Why the dual-write requirement?** The guideline docs in `docs/guidelines/` are the human-readable source. The SKILL.md is the agent-optimized condensed version. If you update the docs without updating the skill, agents will follow stale rules. The command enforces the sync as part of the procedure.

**Why does it decide WHERE to place content?** Not just appending. The command reads the target file, finds where the new concept logically belongs (which section, which position), and integrates it into the existing structure. Then checks for redundancy and restructures if needed. This keeps docs clean over time rather than growing a "misc" tail section.

### `/review-full` and `/review-module`

Thin commands that invoke the reviewer agent on the full project or a specific module. Their content is minimal — just a one-sentence instruction. Their value is: (a) one keystroke to trigger a comprehensive review, (b) they establish a habit of reviewing before delivery.

---

## 6. The Pipeline in Practice

When a new feature request arrives, the full pipeline looks like this:

```
User: "Add a billing module with invoice creation"

Main agent (orchestrator):
  1. Spawn architect → "Plan the billing module implementation"
     architect reads: docs/specs/modules/, existing modules, portal/
     architect returns: file list, layer breakdown, migration plan, UOW boundaries

  2. Spawn system-analyst → "Write UC doc for invoice creation"
     system-analyst reads: architect's plan, UC templates
     system-analyst writes: docs/specs/modules/billing/usecases/invoice/create.md

  3. In parallel:
     Spawn go-coder → "Implement create-invoice use case per the UC doc and architect's plan"
     Spawn go-tester → "Write system tests for create-invoice per the UC doc"

  4. Spawn reviewer → "Review billing/invoice/create implementation"
     reviewer reads: code, tests, UC doc
     reviewer runs: make lint, make test, make test-system
     reviewer reports: CRITICAL/WARNING/SUGGESTION

  5. go-coder fixes CRITICAL items
  6. reviewer re-reviews if needed
  7. Deliver
```

**Why this sequence and not a simpler one?**

- **Architect before coder**: Prevents the coder from discovering architectural problems mid-implementation (wrong layer, missing portal, bad UOW boundary). Fixing architecture mid-code is expensive.
- **Docs before code**: UC docs are the specification. Without them, the coder makes assumptions that may differ from what the frontend dev or product owner expects.
- **Parallel coder+tester**: After the UC doc exists, both agents have a complete specification to work from independently. Tests derived from specs, not from implementation — prevents test bias.
- **Reviewer last**: Nothing is delivered without verification. The reviewer is a gate, not a formality.

---

## 7. Key Principles to Carry Forward

These are the underlying ideas that make the system work. Any blueprint you build should respect them:

**1. Context is currency.** Every token in an agent's context window is either working for you or wasting space. Skills load only for agents that need them. CLAUDE.md is concise. Agents don't repeat guidelines — they reference skills.

**2. Tools define agent character.** Restricting tools is not just safety — it's role definition. An architect with write tools is just a coder. A reviewer with edit tools is a reviewer who might rationalize fixes instead of reporting them. Tool restrictions enforce the agent's purpose.

**3. Specialization beats generalism.** A generic agent asked to "plan and implement and test and review" makes compromises at every step. Specialized agents make no compromises within their domain. The pipeline has overhead but produces better output than a single all-purpose agent.

**4. Docs are not optional.** The doc-code sync requirement (Execute comments mirror UC steps) is not documentation for humans — it's a contract between agents. The system-analyst writes the spec; the coder implements it; the reviewer checks they match. This chain only works if the docs are real and maintained.

**5. Verification is mandatory, not aspirational.** Lint, unit tests, system tests — all three must pass before delivery. This is stated in CLAUDE.md, enforced by the reviewer, and baked into each agent's instructions. "It works" is not a verification strategy.

**6. Commands reduce cognitive load.** `/generate-module` is not just convenience — it encodes all the architectural knowledge about what a complete module looks like. Without it, each module created would require the developer (or main agent) to remember 7 steps and all the templates. The command is expertise made reusable.
