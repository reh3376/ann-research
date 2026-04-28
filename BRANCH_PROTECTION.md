# Branch Protection and Workflow

This document describes the branching, review, and merge rules for this
repository, and the one-time setup needed to enforce them on GitHub.

The asymmetric protection here matches the actual relationship: `@reh3376`
is the principal investigator on this research program; an AI research
assistant (Claude) drafts artifacts via PR for review.

---

## Workflow summary

| Actor | Path to `main` |
|---|---|
| `@reh3376` | Direct push permitted (admin bypass) |
| Anyone else, including AI assistant | Pull request → `@reh3376` approval → merge |

The AI assistant works on branches named `claude/{YYYY-MM-DD}-{slug}` and
opens PRs into `main`. Each PR uses the template in
`.github/pull_request_template.md` and requires:

- Approval from `@reh3376` (enforced via CODEOWNERS)
- Conversation resolution before merging
- Linear history (squash merge)

---

## One-time GitHub setup

Run through this checklist after the initial commit lands on `main`. These
settings cannot be applied before there is a `main` branch to protect.

### 1. Enable branch protection on `main`

Go to **Settings → Branches → Branch protection rules → Add rule**.

Branch name pattern: `main`

Enable:

- [x] **Require a pull request before merging**
  - [x] Require approvals: **1**
  - [x] Dismiss stale pull request approvals when new commits are pushed
  - [x] **Require review from Code Owners**
- [x] **Require conversation resolution before merging**
- [x] **Require linear history**
- [x] **Do not allow bypassing the above settings**
- [x] **Restrict who can push to matching branches** — leave empty (admins
      can still push by default; we override below for `@reh3376`)
- [ ] *Do not* enable "require signed commits" — adds friction without
      meaningful security benefit at this scale.
- [ ] *Do not* enable "require status checks" — no CI yet.

### 2. Configure admin bypass for `@reh3376`

In the same branch protection rule:

- Enable **"Allow specified actors to bypass required pull requests"**
- Add `@reh3376`

This lets the maintainer push directly to `main` while still routing
assistant contributions through review.

### 3. Verify CODEOWNERS

`.github/CODEOWNERS` should contain:

```
* @reh3376
```

This makes `@reh3376` a required reviewer on every PR.

### 4. Repository settings

Under **Settings → General → Pull Requests**:

- [x] Allow squash merging — *default*
- [ ] Allow merge commits — *disable*
- [ ] Allow rebase merging — *disable*
- [x] Always suggest updating pull request branches
- [x] Automatically delete head branches after merge

Under **Settings → General → Features**:

- [x] Issues — enabled
- [x] Discussions — enabled (for substantive engagement)
- [ ] Wikis — disabled (the vault itself is the wiki)
- [ ] Projects — disabled for now

---

## PR conventions for the AI assistant

The assistant follows these conventions when drafting PRs:

1. **Branch name:** `claude/{YYYY-MM-DD}-{short-slug}`. Example:
   `claude/2026-04-28-monty-vote-trace-step-3`.
2. **Commit per artifact:** one observation, question, hypothesis,
   decision, dead-end, or thread update per commit. The temptation to
   batch is real and is the wrong move; per-artifact commits make
   `git log -- path/to/file.md` useful as a per-thought history.
3. **Commit message form:** `add: <id> <slug>`, `update: <id> <change>`,
   `supersede: <old-id> by <new-id>`, `dead-end: <id> <slug>`,
   `decide: <id> <slug>`. Mirrors artifact types so `git log --oneline`
   reads as a research timeline.
4. **PR description:** uses the template at
   `.github/pull_request_template.md`. Lists artifacts added/changed,
   threads affected, what specifically needs maintainer judgment, and the
   assistant's recommendation.
5. **Squash merge:** each merged PR appears as one commit on `main` with
   a clean message.

---

## When this document changes

Updates to branch protection or workflow themselves go through PR (or
direct push by `@reh3376`) and are recorded as `decisions/` artifacts so
the change history is preserved in the vault as well as in `git log`.
