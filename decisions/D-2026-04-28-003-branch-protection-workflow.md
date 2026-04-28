---
id: D-2026-04-28-003
type: decision
slug: branch-protection-workflow
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 4
where: []
tags: [meta, infrastructure, workflow, governance]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "BRANCH_PROTECTION.md (this repository)"
alternatives_rejected:
  - "symmetric protection (Roger also PRs his own work)"
  - "no protection"
  - "two-reviewer requirement"
---

# Branch Protection and Workflow

## Decision

`main` is protected. PRs are required for all contributions from the AI
research assistant; `@reh3376` pushes directly to `main` as repository
admin (admin bypass enabled). `CODEOWNERS` requires `@reh3376` review on
all paths; squash-merge is the only merge strategy; linear history is
required.

## Context

This is a two-actor research program: `@reh3376` as principal
investigator, an AI research assistant (Claude) as collaborator. The
question was how to balance review discipline with the friction-of-
two-people scale.

## Alternatives considered

### Symmetric protection (Roger also PRs his own work)

Rejected for friction. The PR mechanism has value as a checkpoint on
the assistant's contributions, where the assistant might pattern-match
or hallucinate; it has less value as a checkpoint on Roger's
contributions, where the bottleneck would be Roger reviewing his own
work or pulling in a non-existent second reviewer.

### No protection

Rejected. Without PR-required review on the assistant's work, there's no
structural moment for Roger to read what the assistant has done before
it lands on `main`. That moment is the entire point of the workflow:
not to catch malicious behavior, but to ensure the assistant's
artifacts are something Roger has read and stands behind.

### Two-reviewer requirement

Rejected. There are only two of us. Two-reviewer requirements deadlock
at this scale.

## Rationale

The asymmetric protection matches the actual relationship. Roger is the
PI; the assistant drafts artifacts that Roger reviews. Required PR review
on assistant contributions creates the mandatory read-before-merge
checkpoint. Direct push for Roger preserves agility on his own work.

## Consequences

- The assistant works on `claude/{YYYY-MM-DD}-{slug}` branches and opens
  PRs into `main`, using the template in
  `.github/pull_request_template.md`.
- Each PR commit is per-artifact, not per-session.
- Squash-merge produces a clean `main` history.
- Linear history is required (no merge commits).
- Roger configures branch protection on GitHub per the checklist in
  `BRANCH_PROTECTION.md`.

## Revisit conditions

- If a third collaborator is added, revisit whether protection should be
  symmetric.
- If admin-bypass becomes a vector for accidental direct-pushes that
  should have been PRs, revisit the bypass.
- If PR friction starts producing commit-batching (multiple artifacts per
  commit), revisit either the workflow or the convention.
