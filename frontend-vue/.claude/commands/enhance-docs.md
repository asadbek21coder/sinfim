# Enhance Documentation

You are enhancing project documentation. The user will describe a concept, rule, or guideline to add.

## Process

1. **Identify the target file.** Determine where the concept belongs:
   - `CLAUDE.md` — agent-facing instructions, workflow rules, session-level guidance
   - `.claude/skills/frontend-guidelines/SKILL.md` — coding patterns, standards, and conventions

   Most coding rules belong in the skill. Only workflow and agent orchestration guidance belongs in CLAUDE.md.

2. **Read the entire target file.** Understand the current structure, sections, and flow.

3. **Decide where the concept fits.** Don't just append to the bottom. Think about:
   - Which existing section does it belong to?
   - Does it need a new section? If so, where does it logically fit?
   - Does adding this concept make an existing section too long? If so, split it.

4. **Integrate the concept.** Write it in the same style as surrounding content:
   - Match the tone (concise, direct, no fluff)
   - Match the formatting (code blocks, tables, bullet lists — whatever the section uses)
   - If the concept relates to existing points, group them together

5. **Dual-write when adding coding rules.** If the concept is a coding standard or pattern:
   - Add it to CLAUDE.md (or the human-facing doc if it exists)
   - AND add/update the corresponding section in `.claude/skills/frontend-guidelines/SKILL.md`
   - The skill is condensed — no verbose explanations, just rules and patterns
   - The CLAUDE.md or human doc can have more context

6. **Restructure if needed.** After integration, re-read and check:
   - Does section ordering still make sense?
   - Are related concepts grouped?
   - Is anything redundant now? Remove duplication.

## Rules

- **Never lose existing concepts.** Every idea currently in the file must survive.
- **Be concise.** Add what's needed, nothing more.
- **Preserve formatting conventions.** Match existing style.
- **Show the user what changed.** After editing, briefly summarize what was added and any restructuring.
