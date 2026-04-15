---
name: red-green-tdd
description: >
  Strict Red-Green-Refactor TDD workflow enforcement for AI agents. This is a step-by-step
  protocol with hard gates — not guidelines. ALWAYS trigger by default when implementing any
  code that contains conditional logic, business rules, algorithms, state machines, validators,
  branching logic, validation, state transitions, or when fixing bugs in logic-bearing code.
  The user does NOT need to mention TDD explicitly — this skill activates automatically for
  any feature, bug fix, or behavior change that involves logic.
  Do NOT trigger for: renaming, config changes, log statements, dependency updates, documentation,
  pure data classes, database migrations, generated code, or trivial one-line changes with no logic.
---

# TDD Protocol

This is a strict workflow with hard gates. Each gate has an output requirement and a condition
that must be true before proceeding to the next gate. Do not skip gates. Do not combine gates.

## Scope: Which Test Layers

TDD applies to **unit and integration tests only**. It is a design tool — it drives you to think
about behavior before implementation. Higher-level tests don't drive design and are too slow for
tight RED-GREEN-REFACTOR cycles.

| Layer | Apply TDD? | Notes |
|---|---|---|
| Unit tests | **Yes** | Pure functions, business logic, validators, state machines, algorithms |
| Integration tests | **Yes** | Service layer with real DB, REST endpoints with framework test client |
| API/Contract tests | Sometimes | Test-first works for API contracts, but not strict gate cycles |
| E2E / UI tests | **No** | Too slow, too brittle. Write after the feature works. |
| Performance / load | **No** | Not behavioral — you don't RED-GREEN a latency threshold |
| Visual regression | **No** | Screenshot diffs are not behavioral assertions |

## Entry Point: Route the Workflow

Determine which protocol to follow BEFORE doing anything else:

1. **User explicitly said "skip tests" / "no tests" / "just the code"**
   → Implement without TDD. You may note once that TDD is recommended. Do not argue.

2. **The change has no logic** (pure data class, config, migration, enum without behavior,
   renaming, adding logs, dependency update)
   → Skip TDD. Proceed directly.

3. **The task is E2E, performance, visual regression, or smoke tests**
   → Write tests directly without TDD gates. These layers verify the assembled system,
   not drive design.

4. **Everything else** (new feature, bug fix, refactoring logic, adding behavior)
   → Follow the [Developer Protocol](#developer-protocol).

---

## Developer Protocol

### GATE 0: Decompose

**Action:** Break the requirement into discrete, testable behaviors. List them from simplest
to most complex. Edge cases are behaviors — list them explicitly.

**Output requirement — print this before writing any code:**

```
BEHAVIORS:
1. [simplest happy path]
2. [next variation]
...
N-2. [edge case: empty/null/boundary]
N-1. [edge case: error condition]
N.   [edge case: concurrent/state issue]
```

**Gate condition:** The list exists and has at least 3 items. If you cannot identify 3 behaviors,
reconsider whether TDD applies — the change may be too simple.

**Edge case guidance:** For each behavior, consider: empty/null inputs, boundary values (0, 1, -1,
MAX), invalid formats, error paths, and state issues.

**For bug fixes:** The first behavior is always "reproduce the bug" — write the exact condition
from the bug report as behavior #1.

---

### GATE 1: RED — Write One Failing Test

**Action:** Write a single test for the next behavior from your list.

**Test rules:**

- Name: `should <expected outcome> when <condition>`
- Structure: Arrange → Act → Assert (blank lines between sections)
- One logical assertion concept per test
- No logic in test code (no if/for/while/try-catch)
- No mocking the system under test

**Output requirement:** The test code, then run it. State the failure:

```
RED: [behavior #N] — Test fails because [specific reason]
```

**Gate condition:** The test fails for the RIGHT reason. The assertion fails because the production
code doesn't handle this case yet. A compilation error counts as valid RED for the first test of
a new component. If the test passes unexpectedly — STOP. Investigate. Either the behavior already
exists (skip to next) or your test is wrong (fix it).

---

### GATE 2: GREEN — Minimal Production Code

**Action:** Write the MINIMUM code to make the failing test pass.

**"Minimum" means:**

- A constant return is fine if it makes this one test pass — the next test forces generalization
- Do not add error handling no test requires
- Do not add parameters no test uses
- Do not add branches no test exercises
- Do not refactor yet — that is the next gate

**Output requirement:** The production code change, then run the relevant test class/file. State the result:

```
GREEN: Tests pass ([N] tests in this class/file)
```

**Gate condition:** All tests in the relevant class/file pass. If any test fails, fix the production code (not the tests) until all pass.

---

### GATE 3: REFACTOR — Clean Up While Green

**Action:** Improve code structure without changing behavior. This step is optional if the code
is already clean — but you must explicitly decide.

**Allowed changes:** Remove duplication, improve names, extract methods, simplify expressions.
**Forbidden:** Adding new behavior, changing what the code does, adding new test cases.

**Output requirement:** Either:

```
REFACTOR: [what you changed] — Tests still pass ([N] tests in this class/file)
```

or:

```
REFACTOR: No changes needed
```

**Gate condition:** All tests in the relevant class/file still pass. If a test breaks, you changed behavior — undo and try a smaller change.

---

### GATE 4: Next Behavior

Pick the next behavior from your list. Return to GATE 1. Repeat until all behaviors are covered.

**For simple, well-understood behaviors** (standard CRUD, straightforward validation), you may
group 2-3 related RED-GREEN cycles — but always show each test before the code it drives, and
produce RED/GREEN/REFACTOR markers for each.

**For complex or edge-case behaviors** (state machines, business rules, concurrency), show
each cycle individually.

---

### GATE 5: Done Checklist

After all behaviors are implemented, verify each item.

**Output requirement:**

```
TDD DONE:
- [x] Every behavior from the list has a test
- [x] Every public method/function has at least one test
- [x] Edge cases considered (skipped: [list with rationale] or "none")
- [x] Error/exception paths are tested
- [x] No production code exists without a test that drove it
- [x] All tests pass ([N] tests)
- [x] Test names describe behaviors, not implementation
- [x] Each test is independent (can run in isolation)
```

If any item fails (`[ ]`), go back and fix it before declaring done.

---

## Bug Fix Protocol

A specialization of the Developer Protocol. The sequence is non-negotiable:

1. **GATE 0:** Behavior list where #1 is "reproduce the exact bug condition"
2. **GATE 1 (RED):** Write a test that reproduces the bug. Run it. Confirm it fails.
3. **GATE 2 (GREEN):** Fix the bug with minimal code change. Run all tests. Confirm all pass.
4. **GATE 1-3 again:** Write tests for adjacent edge cases (boundary values near the bug).
5. **GATE 5:** Done checklist.

The reproduction test MUST exist before the fix. This guarantees the test actually catches the bug.

---

## Legacy Code (No Existing Tests)

When adding behavior to untested code:

1. Write **characterization tests** first — tests that document current behavior, not desired behavior
2. Run them to confirm they pass (they describe what the code does NOW)
3. Then follow the normal Developer Protocol for the new behavior
4. Do NOT try to achieve full coverage of existing code — only cover what you are changing

---

## Refactoring-Only Tasks

When restructuring without adding behavior:

1. Ensure existing tests pass (add characterization tests if none exist)
2. Make one structural change at a time
3. Run all tests after each change
4. Do NOT add new behavior — if you discover untested behavior, note it and move on
5. No new test cases for new behaviors — you are not adding behavior
6. Do NOT produce GATE 0 behavior list or RED/GREEN markers — this is not a TDD cycle

---

## Mocking Rules

- **Mock:** External I/O (HTTP clients, message queues, payment gateways), time (fixed clock)
- **Don't mock:** The class under test, simple value objects, data access in integration tests
- **Red flag:** More mock setup lines than assertion lines → you're testing wiring, not logic
- **Prefer:** Real implementations and fakes over mocks. Use real databases (Testcontainers, in-memory, or test DB) for integration tests.

---

## Test Execution

**CRITICAL — targeted tests during cycles, full suite only at the end:**

- **GATE 1-4 (RED/GREEN/REFACTOR):** Run ONLY the specific test class or file you are working on. Use the project's targeted test command. NEVER run the full test suite during individual cycles — it is too slow and provides no additional value for the behavior you are developing.
- **GATE 5 (Done):** Run the full test suite exactly once to catch regressions.
- Adapt to the project's existing test framework — check existing test files before writing new ones
- Place tests alongside existing test files, following the project's convention

## Reviewer Protocol

When reviewing code for TDD compliance, check:

**Output format:**

```
### TDD Compliance: [PASS / FAIL]

Tests present:
- [x] / [ ] Every behavior has a corresponding test
- [x] / [ ] Test names describe behaviors (`should X when Y`)
- [x] / [ ] No logic in tests (no if/for/while)
- [x] / [ ] Tests are independent (no shared state, no ordering dependencies)

Coverage:
- [x] / [ ] Happy path tested
- [x] / [ ] Error/exception paths tested
- [x] / [ ] Edge cases considered (or noted as N/A with reason)

Quality:
- [x] / [ ] Tests assert behavior, not implementation details
- [x] / [ ] Mocking used only where appropriate (I/O, time, external deps)

[If FAIL: list specific behaviors that are untested or tests that test implementation]
```
