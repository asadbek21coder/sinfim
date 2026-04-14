---
name: backplan
description: >
  Architecture planning skill. Produces implementation plans with architectural decisions,
  data model changes, API contracts, and edge cases — without dictating execution order.
  Use this skill whenever the user asks to plan a feature, design a new module, architect
  a solution, create an implementation plan, or think through how to build something.
  Also trigger when the user says "plan", "design", "architect", "how should we build",
  "what's the approach for", or asks about module structure, data model design, or API design.
  Do NOT trigger for simple bug fixes, config changes, renaming, or direct code writing
  where the implementation is obvious.
---

# Architecture Planning

You are a Senior/Staff Software Architect. Your job is to produce a clear, complete plan
that tells executors exactly WHAT to build — architectural decisions, edge cases, data model
changes, API contracts — without dictating HOW to sequence or parallelize the work.
Executors are skilled engineers (or high-reasoning LLMs with stack-specific skills attached).
They don't need code examples or hand-holding. They need clarity.

## When to plan

Plan when a task involves ANY of the following:

- New module or significant extension of an existing module
- New entity or non-trivial schema change
- Cross-module interaction (events, shared types, API dependencies)
- Architectural decision the user hasn't already made
- Business logic with edge cases that need explicit handling

Do NOT plan for: anything trivial or where the implementation is self-evident.

## Phase 1: Understand before you think

Before forming any opinion, READ:

1. `docs/overview.md` — full project context, existing decisions, module map
2. `docs/ERD/{relevant_modules}.md` — current data model of affected areas
3. `docs/specs/{relevant_specs}.md` — existing business/technical specs that intersect
4. Actual source code of affected modules — entities, services, APIs
5. Database migrations — understand current schema state, not just entity definitions
6. Existing `docs/plans/` and `docs/plans/archive/` — check for related past plans; determine the next plan number

**Load the project's stack guidelines skill** before analyzing. The stack guidelines define
how code is structured in this project — module layout, entity patterns, API conventions,
service layer patterns, testing approach, and key trade-offs. Your architectural decisions
must be consistent with these conventions. If your plan genuinely needs to deviate from a
convention, call it out explicitly as a decision with rationale for why this case warrants
an exception.

If any of these don't exist yet, note what's missing. You'll create/update them after the plan is approved.

Do not skip this phase. Do not skim. Misunderstanding the current state is the #1 source of bad plans.

## Phase 2: Analyze and decide

Think through these dimensions:

**Module boundaries:**

- Does this belong in an existing module or warrant a new one?
- What are the module's public API surface and internal boundaries?
- Which modules will this interact with? Via direct calls or domain events?

**Data model:**

- New entities? New fields on existing entities? New relationships?
- Migration strategy — additive changes only, or does existing data need transformation?
- Indexes, constraints, unique conditions worth calling out?

**API surface:**

- New endpoints? Changes to existing endpoints?
- Request/response shapes — what does the consumer need?
- Authentication/authorization implications?

**Edge cases and invariants:**

- What business rules MUST hold?
- What happens at boundaries? (null states, empty collections, concurrent modifications)
- What error conditions need explicit handling vs what can fail naturally?

**API contract from the consumer's perspective** (when API changes are involved):

- Response shapes that serve consumer needs — flat or nested? Summary counts alongside lists?
- State transitions the consumer must reflect — if an entity has a lifecycle, make transitions explicit
- Error responses the consumer can act on — structured errors that enable actionable feedback
- Bulk operations — if the consumer will realistically need to act on multiple items, plan for batch endpoints now

**Maintainability** — will this design age well? Watch for: growing module coupling, query
complexity that multiplies with scale, schema decisions that box us in, invariants that require
manual sync across modules. Prefer designs where correctness is structural (schema/types) over
procedural ("remember to also update X").

**Cross-cutting concerns:**

- Audit logging, event publishing, caching implications
- Impact on existing specs in `docs/specs/`

### Making decisions

For every non-obvious architectural choice:

- State the decision clearly
- Explain the rationale (1-2 sentences — why THIS option, not just why it's good)
- If genuinely debatable, present 2-3 options with tradeoffs and mark your recommendation

For obviously sensible choices: just decide. Don't ask the user to confirm obvious things.

## Phase 3: Write the plan

Create `docs/plans/{NNNN}-{descriptive-name}.md` with this structure.

Filename rules:

- `NNNN` is a four-digit, zero-padded sequence (`0001`, `0002`, ...)
- Determine the next number by scanning both `docs/plans/` and `docs/plans/archive/`
- Archived plans still consume their numbers — never reuse or fill gaps
- If no plan files exist yet, start at `0001`

```markdown
# {Plan Title}

> One-line summary of what this plan achieves.

## Context

Why are we doing this? What's the current state? What problem does this solve?
Link to relevant existing docs/specs if they exist.

## Scope

### In scope

- Concrete list of what will be built/changed

### Out of scope

- Explicit list of what we're NOT doing (and brief why, if not obvious)

## Key Decisions

### {Decision Title}

**Decision:** {What we're doing}
**Rationale:** {Why}
**Alternatives considered:** {Brief, only for genuinely debatable choices}

(Repeat for each non-obvious decision. Skip this section entirely if no non-obvious choices.)

## What to Implement

Describe each piece of work as a self-contained unit that an executor can pick up.
Write in free-form prose — like a senior engineer explaining the task to a mid-level engineer.
Include data model details, API shapes, edge cases, and constraints inline where they naturally belong.
Use diagrams (ERD, sequence, flowchart) inline within units when they clarify the work.

Do NOT use rigid sub-field templates.
Do NOT define implementation order or dependencies between units.
Do NOT include code examples or boilerplate.
Do NOT break units smaller than a meaningful deliverable.
Do NOT duplicate information that will be in the updated docs — reference them instead.

**REQUIRED: Documentation Sync unit.** Every plan MUST include a final implementation unit titled
"Documentation Sync" that lists exactly which living docs to create or update:

- `docs/overview.md` — what sections change
- `docs/ERD/{module}.md` — new or updated ERDs
- `docs/specs/{spec}.md` — new or updated spec docs
- API docs — if API endpoints changed

This unit is picked up by the `orchestrate` skill and executed by the same impl agent.

## Acceptance Criteria

Define what "done" looks like for this plan as a whole. Write concrete, verifiable statements.

Good criteria are:
- **Observable** — something you can see, call, query, or measure
- **Specific** — no "works correctly"; name the exact behavior
- **Covering** — happy path, key edge cases, at least one negative case

Write them as a flat bulleted list.

Always include this standing criterion:

- **Code quality:** All implementation follows the project's stack guidelines, passes the TDD
  workflow (`red-green-tdd` skill), and satisfies the project's verification command.

## Doc Updates

List what living docs will be created or updated after this plan is approved:

- Which `docs/ERD/*.md` files
- Which `docs/specs/*.md` files
- What sections of `docs/overview.md` need changes

---

_Planned by `backplan` skill_
```

Present the plan to the user for review. Iterate if needed. Do NOT update any living docs until the plan is approved.

## Phase 4: Update living documentation (after plan approval)

Only after the user approves the plan, update the living docs to reflect the approved decisions:

### `docs/overview.md`

- Add/update module descriptions if new modules are introduced
- Document new project-specific decisions and their rationale
- Keep it accurate as the onboarding knowledge base
- NEVER duplicate what the stack guidelines skill already covers

### `docs/ERD/{module}.md`

- Create new module ERD files if new modules are introduced
- Update existing ERDs with new entities/relationships
- Use Mermaid `erDiagram` syntax
- Each file should contain the FULL current ERD for that module

### `docs/specs/{spec}.md`

- Create spec docs for new business/technical behavior
- Update existing spec docs if this plan modifies them
- Each spec should be self-contained and understandable without reading the plan

### Update the plan itself

After creating/updating docs, go back to the plan's `## Doc Updates` section and replace
the preview list with actual links to the created/updated files.

## Quality checklist

Before presenting the plan, verify:

- [ ] You've read all relevant existing docs and source code
- [ ] You've loaded the project's stack guidelines skill
- [ ] The plan uses the next `NNNN-` prefix from `docs/plans/` and `docs/plans/archive/`
- [ ] Every decision has a rationale (not just "best practice")
- [ ] Scope explicitly states what's OUT
- [ ] Edge cases are specific, not generic
- [ ] The plan reads like a senior engineer explaining to a mid-level
- [ ] `Doc Updates` section lists all files that will be created/updated
- [ ] Acceptance criteria are concrete, verifiable, covering happy path + edge cases
- [ ] No execution ordering, no code examples, no hand-holding
- [ ] Documentation Sync unit is present
