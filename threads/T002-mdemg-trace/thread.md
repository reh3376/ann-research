---
id: T002
type: thread
slug: mdemg-trace
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 4
where: []
tags: [mdemg, trace, mechanism]
provenance:
  - "mdemg repository: https://github.com/reh3376/mdemg"
  - "conversation 2026-04-28 between @reh3376 and Claude"
owner: "@reh3376 (primary, has end-to-end context); Claude (drafting / corroborating reads)"
current_question: "What does an MDEMG ingest+retrieve cycle actually do, mechanistically, from a code observation entering the system through retrieval results being returned?"
---

# Thread T002 — MDEMG Trace

## Goal

Develop a mechanism-level understanding of what MDEMG actually does
during a typical ingest+retrieve cycle. Mirror image of `T001`: trace
one code observation entering the system end-to-end through ingest,
node creation, hidden-layer interaction, Hebbian edge updates,
retrieval ranking, and Jiminy guidance.

The trace produces the empirical grounding that `Q-2026-04-28-003`
(eta multipliers) and any future MDEMG-side questions need.

## Current question

What does an MDEMG ingest+retrieve cycle actually do, mechanistically,
from a code observation entering the system through retrieval results
being returned? Specifically:

1. What happens at ingest: node creation, embedding, surprise scoring,
   hidden-layer assignment, edge creation.
2. What the Hebbian update does in detail: the multi-rate `eta`, the
   tanh soft-cap, the asymmetric direction edges, surprise factors.
3. What happens at retrieval: hybrid scoring (embedding + BM25 +
   graph), `query_attention`, temporal weighting, intent translation,
   cursor pagination.
4. Where Jiminy fires and what it looks at to decide.
5. What hidden-layer promotion does (the clustering step, the
   reclassifier).

## Method

This thread benefits from the asymmetry that `@reh3376` has end-to-end
operational context that the assistant lacks. The method:

1. `@reh3376` walks Claude through the ingest+retrieve flow at a
   level of detail he chooses, either in conversation or by writing
   the trace himself.
2. Claude reads the relevant files in parallel and asks clarifying
   questions, especially about:
   - Empirical grounding for tuned parameters (note `Q-2026-04-28-003`)
   - What happens at scale (clustering, graph health, scaling
     limitations called out in `02_Current_State_MDEMG.md`)
   - What components are settled and what is still in flux
3. Claude writes observations, questions, and hypotheses into the
   vault from the joint reading. `@reh3376` reviews and corrects.
4. When the trace is complete enough to answer the current question,
   move thread status to `complete`.

## State

**Initial state.** No trace work has begun in this vault. The
mdemg repository has been cloned and structurally surveyed (Go
layout, key packages identified — `internal/learning`, `internal/jiminy`,
`internal/hidden`, `internal/retrieval`, `neural/`). The structural
survey is documented in `DE-2026-04-28-001` as a known-shallow read.

The project files (`/mnt/project/`, in the prior conversation context)
provide higher-level architectural framing that the trace can
ground in mechanism.

## Open work

- [ ] Decide on the trace scope: typical ingest? a specific scenario
      (J17 protocol exchange? RSIC self-improvement loop)? minimal
      "hello world" cycle?
- [ ] Decide division of labor: who writes which parts of the trace?
- [ ] First trace pass: ingest of one code observation

## Artifacts produced by this thread

*None yet.* `Q-2026-04-28-003` (eta empirical grounding) is a question
attributed to this thread and is expected to be answered or refined
by the trace work.

## Log

### 2026-04-28

Thread created as part of bootstrap. The expected division of labor
is more `@reh3376`-led than T001 because operational context is
asymmetric in his favor.
