# Monorepo vs Separate Repos — Honest Comparison

This document compares the **go-vue-monorepo-blueprint** (one repo, both stacks) against the
**go-enterprise-blueprint + vue-blueprint-web** pair (two separate repos). No sugarcoating.

---

## The Core Trade-off

**Monorepo:** Everything in one place. One clone, one history, one CLAUDE.md. Features ship as atomic commits that include backend and frontend. The full-stack AI agent sees everything.

**Separate repos:** Each stack lives and breathes independently. Backend team and frontend team never step on each other. Each deploys on its own schedule. Scaling an org is easier.

The monorepo wins for **speed at small scale**. Separate repos win for **autonomy at larger scale**.

The question is not which is better. It's which fits your current situation.

---

## Honest Pros of Monorepo

### 1. One commit = one feature

In separate repos, shipping a billing feature means:
- Commit in `go-backend`: "feat: add invoice creation use case"
- Commit in `vue-frontend`: "feat: add invoice page"
- If you roll back, you roll back twice, separately, and hope the timing matches

In the monorepo:
- One commit: "feat: billing — invoice creation (backend + frontend)"
- Roll back is `git revert` on one commit. Done.

This is not a small thing for a solo developer shipping fast.

### 2. The AI architect sees the full picture

In separate repos, you ask the Go architect to plan invoice creation — it plans the use case, endpoints, and UC doc. Then you ask the Vue architect to plan the invoice page — it reads the Go UC doc if you copy it over, or guesses at the API contract.

In the monorepo, one architect agent reads both the Go module structure AND the existing Vue views and plans:
- Which Go use cases to add
- The exact API contract (request/response shapes)
- Which Vue types/api/views to create
- Where the frontend calls the backend

One planning pass. No copy-pasting. No contract drift.

### 3. No branch coordination

In separate repos, feature branches must be coordinated:
- `feature/billing` in backend must match `feature/billing` in frontend
- If backend is reviewed first, frontend dev merges against a contract not yet in main
- CI for frontend runs against main backend, which doesn't have the new endpoints yet

In the monorepo, `feature/billing` has both. CI runs against both. There's nothing to coordinate.

### 4. Shared documentation

`docs/specs/` lives at the root. Backend use case docs and frontend feature specs are siblings.
No question of "which repo has the system flow doc" or "did the frontend team get the updated UC?"
One place, one source of truth.

### 5. Simpler setup for new contributors

One `git clone`. One `docker-compose up`. One `.env` file. One `README.md` to read.

With separate repos: clone two repos, set up two environments, configure the proxy in the frontend to point at the running backend, make sure both are on compatible branches, etc.

### 6. API contract violations caught immediately

The frontend dev and backend dev are in the same repo. If a Go UC changes its response shape, the TypeScript types in `frontend/src/types/` break on `npm run build`. The breakage is visible in the same PR.

In separate repos, the backend can break the contract and the frontend doesn't know until it pulls and builds.

---

## Honest Cons of Monorepo

### 1. CI is harder to split

If you push a Go change, you don't want to run `npm run build`. If you push a Vue change, you don't want to run `make test-system` (which spins up Postgres + Kafka).

You have to configure path-based CI triggers:
```yaml
# GitHub Actions
on:
  push:
    paths:
      - 'backend/**'   → run Go CI
      - 'frontend/**'  → run Vue CI
```

It's doable but adds CI complexity. Separate repos don't have this problem — each repo's CI is self-contained.

### 2. `node_modules` and Go modules coexist

Running `npm install` in a Go workspace feels wrong. It's not wrong — it just creates a cognitive boundary violation that bothers some people. Some tools (IDEs, linters, global git hooks) try to handle the whole repo as one thing and get confused.

Practically:
- GoLand doesn't natively understand Vue/TypeScript
- VS Code + extensions handles it reasonably well
- `git diff` stats mix `.go` and `.vue` files
- Some grep/glob tools need extra ignore patterns

This is friction, not a blocker.

### 3. Scales poorly past ~5 people

If you hire a backend team and a frontend team, they'll conflict on:
- `main` branch ownership — who reviews PRs touching both?
- Release management — backend wants to deploy Friday, frontend says not yet
- `CLAUDE.md` ownership — whose context docs take priority?
- Permission control — you can't give a contractor read access to only the Vue code

None of these are unsolvable. But they require process overhead that separate repos avoid by default.

### 4. Repository size grows from both stacks

`go.sum`, Go build artifacts, `node_modules` (in lockfiles), Vite build output — if you commit things you shouldn't, the repo grows from two directions. Not a real problem if you write a proper `.gitignore`, but it requires attention.

### 5. Harder to reuse the backend with another frontend

If you later build a mobile app, or a second web interface for a different role, you want the Go backend as an independent service. In the monorepo, it's physically inside a `backend/` folder. You'd need to either:
- Extract it to its own repo (painful)
- Have the mobile app depend on a subdirectory (messy)

Separate repos don't have this coupling.

### 6. Local development requires running two processes

`make infra-up` starts Postgres/Kafka. Then you run `go run ./cmd run` (backend on 8080) AND `npm run dev` (frontend on 5173). Both must be running at the same time. Vite proxies `/api` to 8080.

This is the same as separate repos, but with separate repos you can work on the frontend by mocking the API without running the backend at all. In a monorepo people sometimes forget this is still an option.

---

## When to Use Monorepo

- Solo developer or team of 1-3
- MVP / startup — you're shipping features, not managing teams
- Every meaningful feature touches both backend and frontend
- You want the AI agent to have full context without context-switching repos
- Speed is the highest priority right now
- You don't plan on having multiple frontend clients (no mobile app)

## When to Use Separate Repos

- Team of 5+, or distinct backend/frontend teams
- Multiple frontend clients (web + mobile + admin panel)
- Different deploy cadences (backend needs to stay stable; frontend ships constantly)
- Backend is a shared service used by multiple products
- Contractors need access to only one stack

---

## Migration Path

Starting monorepo and want to split later:
1. `git subtree split --prefix=backend -b backend-only`
2. `git subtree split --prefix=frontend -b frontend-only`
3. Push each to a new repo

Git history is preserved. It's not trivial but it's not destructive either.

Starting separate and want to merge later:
1. `git subtree add --prefix=backend <backend-repo-url> main`
2. `git subtree add --prefix=frontend <frontend-repo-url> main`

Also works, also preserves history.

**Bottom line: this decision is reversible. Don't over-engineer it.**

---

## Summary Table

| Factor | Monorepo | Separate Repos |
|---|---|---|
| Initial setup | Simpler (1 clone) | Two clones, two envs |
| Feature commits | Atomic (1 commit) | Coordinated (2 commits) |
| AI agent context | Full picture, 1 CLAUDE.md | Two separate contexts |
| API contract safety | Immediate (same PR) | Delayed (separate PRs) |
| CI configuration | Path-based filtering needed | Self-contained per repo |
| Team scaling | Gets messy >5 people | Cleaner separation |
| Multiple frontends | Awkward | Natural |
| Deploy independence | Harder | Natural |
| IDE experience | Minor friction | Each stack in its IDE |
| Reversibility | Yes (git subtree) | Yes (git subtree) |

**Best fit:** Solo dev building an MVP SPA + Go API with no immediate plans to scale the team or add more frontend clients.
