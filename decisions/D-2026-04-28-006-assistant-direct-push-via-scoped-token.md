---
id: D-2026-04-28-006
type: decision
slug: assistant-direct-push-via-scoped-token
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 4
where: []
tags: [meta, infrastructure, workflow, security, governance]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "D-2026-04-28-003 (branch protection and workflow)"
  - "GitHub PR #1: https://github.com/reh3376/ann-research/pull/1"
alternatives_rejected:
  - "patch-and-apply via git format-patch (the original mechanism in D-003)"
  - "no token; assistant produces files for manual drag-and-drop"
  - "broader-scoped token with admin or merge capabilities"
---

# Assistant Direct Push via Scoped Token

## Decision

The AI research assistant authenticates to GitHub with a fine-grained
personal access token, scoped to `reh3376/ann-research` only, with
permissions `Contents: read/write`, `Pull requests: read/write`, and
`Metadata: read`. No `Administration` permission. The assistant uses
this token to push to feature branches (`claude/{YYYY-MM-DD}-{slug}`)
and open PRs via the GitHub API, but **cannot approve or merge PRs**
— the protection rules from `D-2026-04-28-003` (CODEOWNERS = `@reh3376`,
required Code Owner review, linear history) remain the gate, and only
`@reh3376` satisfies them.

The token lives in `~/ann-research/.env` on Roger's machine, gitignored,
and is uploaded to active conversations as needed. Token expiration is
30–60 days; rotation is manual.

## Context

`D-2026-04-28-003` codified the asymmetric workflow (Roger pushes
direct, assistant via PR with CODEOWNERS approval) but did not
specify the *mechanism* by which the assistant gets work onto a feature
branch. The implicit assumption was that the assistant would produce
patches (via `git format-patch`) and Roger would apply them locally
with `git am`, then push the branch himself.

That mechanism turned out to be unnecessarily heavy for a vault
consisting almost entirely of markdown files. Roger flagged this:
*"You should be able to handle all this without my help."* The
patch-and-apply ritual added ~20 minutes of friction per PR for
handoff that didn't proportionally serve the discipline of review.

The actual review checkpoint — Roger reading the work before it lands
on `main` — happens at the GitHub PR review UI regardless of how the
branch got there. The mechanism for getting the branch there is
operational plumbing, not a substantive control.

## Alternatives considered

### Patch-and-apply via `git format-patch` (the original D-003 mechanism)

Rejected. Adds operational ceremony without proportional benefit. The
review checkpoint is at PR approval, not at branch creation. Roger's
direct objection: the friction cost outweighed the protection benefit
at this scale. Documented as superseded mechanism here.

### No token; manual file drag-and-drop into the repo

Considered. Assistant produces files in the conversation, Roger drags
them into his working tree, commits, pushes. Lower credential surface
(no token at all). Five-minutes-a-week friction cost.

Rejected because Roger explicitly preferred automation: *"You should
be able to handle all this without my help."* For a vault with
expected commit volume of multiple PRs per active research week, the
manual handoff would create a steady drag.

### Broader-scoped token (with Administration or merge capabilities)

Rejected on principle. The asymmetric protection in
`D-2026-04-28-003` is the substantive control that keeps Roger as the
sole approver and merger. A token with `Administration` permission
could bypass branch protection; a token that could trigger merges
could collapse the asymmetry. The minimum scope sufficient for "push
and PR" is exactly what we want: anything more is unnecessary blast
radius.

## Rationale

The token mechanism *implements* the workflow that `D-2026-04-28-003`
*decided*. It does not change what the workflow is. After this
change:

- Assistant work still goes through PR
- PRs still require approval from CODEOWNERS (only `@reh3376`)
- Linear history is still enforced
- Roger still has admin bypass for direct pushes to `main`

What changes:

- Assistant pushes feature branches and opens PRs via API rather than
  producing patches for Roger to apply
- Token lives in a gitignored `.env` file, uploaded to conversations
  as needed
- `.env` is now part of the project's gitignore baseline (added in
  commit `41ff6f1` on `main`)

## Consequences

- **Cleanup hygiene.** After a PR merges, the feature branch should
  be deleted. The assistant has the token permission to delete the
  branch via `DELETE /repos/.../git/refs/heads/claude/...` and will do
  so unless `Settings → General → Pull Requests → Automatically delete
  head branches` is enabled (recommended; one-time setting on the
  GitHub repo).
- **Token rotation.** Every 30–60 days, Roger generates a new
  token, replaces the value in `~/ann-research/.env`, and uploads
  the new file to the active conversation. The old token gets
  revoked at GitHub.
- **Periodic protection check.** The protection now consists of
  three layers: token scope (no admin, no merge), branch protection
  (PR required, code-owner review required, linear history), and
  CODEOWNERS (only `@reh3376` approves). All three must be in place;
  if any fails, the asymmetry is broken. A monthly sanity check is
  appropriate: try to merge a test PR without approval and confirm
  GitHub blocks it.
- **Documentation drift.** `BRANCH_PROTECTION.md` and the prose
  description in `D-2026-04-28-003` both reference the
  patch-and-apply mechanism. They are now partially inaccurate. This
  artifact serves as the canonical record of the operational change;
  the older documents are not edited (preserves the audit trail) but
  should be read in conjunction with this one.

## Revisit conditions

- **Token compromise.** If the token leaks (committed to a public
  repo, posted in a public chat, etc.), revoke immediately, rotate,
  audit recent commits and PRs for unauthorized activity.
- **Expanded scope.** If a third contributor joins, revisit whether
  the per-actor token model still works or whether a deploy-key /
  GitHub App approach is more appropriate.
- **Protection layer failure.** If a routine sanity check shows that
  branch protection has been weakened (admin bypass enabled too
  broadly, code-owner review accidentally disabled, etc.), this
  decision is suspended pending re-establishment of the protection
  layers.
- **High-volume workflow.** If the PR cadence increases dramatically
  (e.g., dozens of PRs per day), revisit whether per-PR review
  becomes a bottleneck. At current research-cadence scale, it is not.
