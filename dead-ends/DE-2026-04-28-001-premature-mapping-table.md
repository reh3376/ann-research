---
id: DE-2026-04-28-001
type: dead-end
slug: premature-mapping-table
created: 2026-04-28
updated: 2026-04-28
status: terminal
confidence: 5
where: [T001, T002]
tags: [meta, epistemic-discipline, process-failure]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "Claude response to 'continue' instruction, 2026-04-28"
failure_mode: "Pattern-matching on structural similarity between two codebases instead of tracing the mechanisms each codebase actually implements"
retry_if: ""
---

# Dead End: Premature Mapping Table

## What was attempted

After cloning the MDEMG and Monty repositories and reading both at the
level of class hierarchies, file structures, and function signatures,
the assistant produced (a) a side-by-side architectural comparison
section, (b) a "five concrete merge/borrow opportunities" list, and (c)
a draft `MONTY_INTEROP.md` document with a five-row mapping table from
Monty abstractions to MDEMG components, plus the start of Go interface
stubs.

The work was produced in a single conversational turn after roughly an
afternoon of structural reading. It was structured, articulate, and
internally consistent.

## What went wrong

Roger's response: "I want you to develop a deep understanding of the
problem and the media we are working with before attempting to make any
changes. We will discuss findings and work together as a research group
as we move forward in our understanding and form a viable and
methodical path forward."

The redirect was correct. The mapping table read as comprehensive but
was actually **a translation of names, not mechanisms**. For each
Monty abstraction, the assistant had identified an MDEMG location
where a "similar-shaped" component lived, and proposed the mapping. The
mappings were defensible at the level of "these two things have similar
contracts" but unsupported at the level of "these two things share an
algorithmic core that survives substrate change."

Concretely, the assistant had not:

- Traced what `_update_evidence` or `_update_evidence_with_vote` in
  Monty actually do (the math, the data flow, the dimensionality
  changes).
- Followed a single Monty episode end-to-end (one observation in,
  through sensor module, learning module, voting, goal generation,
  motor system).
- Confirmed whether Monty's algorithms exploit properties of continuous
  3D Euclidean geometry that have no useful analog in MDEMG's symbolic
  substrate.
- Read Monty's hypothesis-resampling RFCs (0009, 0013), which appeared
  to indicate live design tension worth understanding before proposing
  to import the architecture.
- Traced an MDEMG ingest+retrieve cycle end-to-end (one code observation
  in, through ingest, hidden-layer interaction, retrieval, ranking,
  Jiminy guidance).
- Established what the empirical grounding is for MDEMG's multi-rate
  `eta` values, before proposing to replace the categorical multipliers
  with continuous confidence values.

The mapping table stood in for understanding rather than emerging from
it.

## Why it failed (best current understanding)

Three contributing factors:

1. **Structural similarity is cheap to recognize.** Class hierarchies,
   ABCs, and function signatures look similar across many systems.
   When you've read a lot of code, surface-level mapping is fast and
   feels productive, but it's underconstrained — most pairs of systems
   admit a "similar-shaped" mapping that doesn't survive contact with
   the actual algorithms.

2. **Bias toward producing tidy artifacts.** A clean mapping table is
   a satisfying artifact to produce. The cleanness is itself a warning
   sign — when two codebases at very different substrate layers
   (sensorimotor 3D versus symbolic graph) map cleanly after one read,
   that's evidence the read was too shallow, not that the systems are
   well-aligned.

3. **Insufficient skepticism about the substrate gap.** The assistant
   noted "substrate differences are real" and then handed Roger a
   translation table as if translation were the hard part. The harder
   question — whether the algorithms exploit substrate-specific
   geometric properties — was not engaged with.

## Lesson

Two operational rules emerge:

1. **When a mapping looks clean after a single read, treat the
   cleanness as evidence the read was too shallow.** Add depth before
   adding the mapping, not after.

2. **Bring observations and questions before proposing integration.**
   Integration proposals at minimum require having traced mechanism
   end-to-end in both systems being integrated. If the trace hasn't
   been done, the integration proposal is premature regardless of
   how confident it sounds.

## Retry conditions

This dead-end is `terminal`. The specific failure mode (premature
mapping table from structural reading) is not to be retried.

What's *not* foreclosed is integration work after T001 (Monty trace)
and T002 (MDEMG trace) have both reached `complete` status with
maintainer review. At that point the trace artifacts will provide the
mechanism-level grounding the original mapping lacked. See
`D-2026-04-28-004` (working agreement).

## Related artifacts

- `D-2026-04-28-004` — working agreement codifying the constraint
- `T001-monty-trace` — the trace that has to happen first
- `T002-mdemg-trace` — the trace that has to happen first
- `T003-grounding-question` — the substrate-level question this attempt
  glossed over
