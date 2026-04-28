---
id: D-2026-04-28-008
type: decision
slug: admin-bypass-merge-pattern
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 5
where: []
tags: [meta, infrastructure, workflow, governance, github]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "GitHub: 'Pull requests cannot be approved by their own author' — platform-level rule, not configurable"
  - "D-2026-04-28-003 (branch protection and workflow)"
  - "D-2026-04-28-006 (assistant direct push via scoped token)"
alternatives_rejected:
  - "create separate GitHub account for the assistant (bot identity)"
  - "GitHub App with installation token"
  - "remove the approval requirement entirely from branch protection"
---

# Admin-Bypass Merge Pattern

## Decision

Pull requests authored by the AI research assistant are merged by
`@reh3376` via **admin bypass**, not via the approval-then-merge
workflow. The "Code Owner approval required" rule remains configured
on the branch protection ruleset, but is bypassed at merge time
because the PR author and the only Code Owner are the same GitHub
identity.

The substantive control — *Roger reads the PR before it lands on
`main`* — is preserved. What is not preserved is a formal "approved
by" record on the PR; merges show as direct merges by the repo owner.

## Context

`D-2026-04-28-006` set up the assistant to push branches and open PRs
using a fine-grained token scoped to `reh3376/ann-research`. That
token is issued under Roger's GitHub account, so PRs the assistant
opens are authored by `@reh3376`.

The CODEOWNERS file (`* @reh3376`) and the branch protection rule
"require review from Code Owners" make `@reh3376` the only valid
approver. GitHub then enforces a platform-level rule that cannot be
configured away: **a pull request cannot be approved by its own
author.** Roger is therefore unable to approve any PR the assistant
opens, because GitHub views Roger as the author.

This is an identity collision, not a permissions misconfiguration.
The branch protection rules from `D-2026-04-28-003` are correctly
configured; the constraint is that the assistant operates under
Roger's identity rather than a separate one.

## Alternatives considered

### Create a separate GitHub account for the assistant (bot identity)

Considered. Standard pattern for automation: a dedicated user account
(e.g., `reh3376-claude` or `ahh-research-bot`) added as a collaborator
with `Write` access. PRs authored by the bot account; Roger approves
as Code Owner; standard workflow restored.

Rejected for now on cost grounds: another account to maintain,
another token rotation cadence, an additional identity-management
surface for a two-actor research repo. The protection benefit is real
but small — the substantive control (Roger reads before merge) is
preserved by admin bypass, so what's lost is mostly the formal
approval audit record on each PR.

Revisit if the identity question becomes load-bearing for some other
reason (e.g., a third contributor joins; the repo needs to integrate
with CI that distinguishes bot commits; the approval audit trail
becomes needed for external review).

### GitHub App with installation token

Considered. More principled than a bot user — apps have their own
identity, short-lived installation tokens, and clean revocation at
the app level rather than at the user level. Better audit story for
a more mature program.

Rejected on the same grounds as the bot account but more so: more
setup, more conceptual machinery (app permissions, installation,
token refresh), and the same protection benefit. Right answer for a
mature multi-contributor repo; wrong cost-benefit for current scale.

Same revisit conditions as the bot account.

### Remove the approval requirement from branch protection

Rejected. The "Code Owner approval required" rule is the structural
record of *who is authorized to approve PRs*, even if it can't be
exercised currently. Removing it would silently weaken the
configuration; the present state ("rule is on, bypassed at merge by
admin") is more honest about what's happening than ("rule is off,
Roger remembers to read PRs before merging") would be.

If the rule is left in place, future-us reading the protection
configuration sees: "approval is required, the only valid approver is
Roger, and the way merges happen in practice is via admin bypass."
That's the truth and it's discoverable from settings.

## Rationale

The substantive purpose of the protection workflow — that Roger reads
each PR before it lands on `main` — does not depend on the approval
button being clicked. It depends on Roger reading the PR. As long as
Roger continues to read each PR before clicking "Merge" (whether via
the approval-then-merge flow or via admin bypass), the protection
holds.

Documenting this explicitly is the substantive cost of this decision.
Without the record, a future reader of the repo settings sees a
configured-but-bypassed rule and wonders why; with this record, the
"why" is in the vault and the configuration is comprehensible.

The honest version of the protection is therefore:

```text
PR is opened by the assistant (under @reh3376's identity via API token)
↓
@reh3376 reads the PR (this is the substantive check)
↓
@reh3376 merges via admin bypass
```

The "Code Owner approval" step is a configuration artifact that
remains in place for future identity-correctness, but is not exercised
in the current setup.

## Consequences

- **Merges show as direct merges by `@reh3376`**, not as
  "approved-then-merged-after-CI." The audit trail is the merge
  commit author plus this decision artifact.
- **Branch protection rules stay configured as-is.** The "Require
  review from Code Owners" rule remains on; it is not removed.
- **The substantive review still happens.** Roger continues to read
  PR descriptions, activity summary comments, and diffs before
  merging. This is a behavioral commitment, not a technical
  enforcement.
- **The `D-007` activity-summary-comment convention becomes
  load-bearing.** Without a formal approval record, the activity
  summary comment is the most-detailed durable record of what
  Roger had access to at merge time. Future audits ("what did Roger
  see when he merged this?") rely on the comment being present and
  thorough.
- **No changes to the token, the branch naming, or any artifact
  schemas.** This decision changes none of that.
- **`D-2026-04-28-003` is partially superseded** in its mechanism
  description (the approval-then-merge pathway is no longer the
  operational path), but its substantive decision (asymmetric
  protection with Roger as sole gatekeeper) is unchanged. Per
  precedent set in PR #2 review, `D-003` is left as-is; this artifact
  is the canonical correction.

## Revisit conditions

- **A second human contributor joins.** Identity collisions go away
  the moment there are two human approvers; the bot-account or App
  options become natural at that point.
- **External audit requirement.** If we ever need formal "approved
  by" records for compliance, regulatory, or publication reasons,
  the bot-account or App options become necessary.
- **Drift in the substantive control.** If Roger ever finds himself
  merging assistant PRs without reading them in detail (the failure
  mode this whole workflow exists to prevent), revisit immediately
  — at that point the technical enforcement is the only thing
  standing between the program and the failure mode in
  `DE-2026-04-28-001`.
- **Token compromise.** Same as `D-006`: revoke, rotate, audit.
- **GitHub changes the platform rule.** If GitHub ever makes the
  author-can't-approve rule configurable, this decision becomes
  obsolete and the approval-then-merge workflow can be restored
  directly.
