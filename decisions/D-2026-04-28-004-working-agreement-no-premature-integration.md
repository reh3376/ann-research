---
id: D-2026-04-28-004
type: decision
slug: working-agreement-no-premature-integration
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 5
where: [T001, T002, T003]
tags: [meta, governance, epistemic-discipline]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "DE-2026-04-28-001 (premature mapping table)"
alternatives_rejected:
  - "no agreement / proceed normally"
  - "stricter form (no merge proposals ever without explicit invitation)"
  - "looser form (caveat assertions instead of forbidding them)"
---

# Working Agreement: No Premature Integration

## Decision

Until research threads **T001 (Monty trace)** and **T002 (MDEMG trace)**
are complete and reviewed, the AI research assistant will not:

- Propose merges or integrations between MDEMG and Monty
- Draft interop specifications, mapping documents, or interface bridges
- Produce side-by-side comparison tables of MDEMG and Monty components

The assistant will instead bring **observations** and **questions**, and
the team decides together which warrant deeper investigation. When
uncertainty exists in the assistant's understanding of either codebase,
the assistant flags it explicitly rather than smoothing over it.

## Context

The first substantive response in this collaboration (`continue` request,
2026-04-28) produced a clean-looking interop document with a five-row
side-by-side mapping table of Monty abstractions to MDEMG components,
after one afternoon of skimming both repos. Roger correctly identified
this as pattern-matching on shape rather than tracing mechanism, and
redirected the work toward deep understanding before any change
proposal. See `DE-2026-04-28-001`.

## Alternatives considered

### No agreement / proceed normally

Rejected. Without an explicit precommitment, the failure mode is likely
to recur — the bias toward producing tidy integration documents is
strong, and "proceed carefully" is too vague to act as a constraint.

### Stricter form (no merge proposals ever without explicit invitation)

Considered. Would be safer but would also block legitimate proposals
that emerge naturally during thread work. The current form ties the
unblock to evidence (thread completion) rather than to permission.

### Looser form (caveat assertions instead of forbidding them)

Rejected. Caveats slide into ignored disclaimers. A behavioral
constraint ("don't produce X until Y") is enforceable; a presentational
constraint ("produce X with caveat Z") is not.

## Rationale

The substrate-level question — whether reference-frame transfer between
sensorimotor and symbolic substrates is a translation problem or a
substrate problem — is open and load-bearing. Producing integration
proposals before that question has been investigated risks committing
the program to architectural commitments the evidence doesn't yet
support.

The trace work in T001 and T002 is the cheapest path to that evidence:
following actual mechanism end-to-end in both codebases will surface
either (a) genuine algorithmic alignment that supports integration, or
(b) substrate-specific load-bearing properties that don't transfer.
Either finding is more valuable than the integration proposals it
preempts.

## Consequences

- The MONTY_INTEROP.md draft and Go interface stubs from the prior
  attempt are not committed to this repository.
- The assistant's PR drafts during T001 and T002 will be observations,
  questions, hypotheses, dead-ends, and thread updates — not
  integration documents.
- The agreement explicitly does not block: questions *about* whether
  integration would work, hypotheses *about* what integration would
  require, or dead-ends *recording* failed attempts at integration.
  These are research artifacts, not integration proposals.

## Revisit conditions

- T001 and T002 each reach `complete` status with maintainer review.
- The team explicitly invites integration work earlier (e.g., a
  collaborator joins and asks for a draft).
- New evidence emerges that the substrate-level question is settled by
  existing prior art we hadn't read.
