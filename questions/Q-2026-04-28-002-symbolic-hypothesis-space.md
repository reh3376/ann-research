---
id: Q-2026-04-28-002
type: question
slug: symbolic-hypothesis-space
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 2
where: [T001, T002, T003]
tags: [monty, mdemg, hypothesis-space, reference-frames, pose]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "tbp.monty: src/tbp/monty/frameworks/models/evidence_matching/learning_module.py — EvidenceGraphLM"
answer: ""
---

# What is the symbolic-substrate analog of Monty's hypothesis space (object_id × pose)?

## The question

Monty's `EvidenceGraphLM` accumulates evidence over a hypothesis space
indexed by `(object_id, pose)`, where:

- `object_id` enumerates known object models
- `pose` parameterizes a 3D rotation/translation

For a code symbol, a constraint, or a process unit in MDEMG, what
plays the role of `pose`?

Candidates that came up in initial reading:

1. **Context fingerprint** (from MDEMG note 05) — a sparse bit vector
   indicating which contextual factors are active. Plausibly the
   "what context is this entity in" analog of "what pose is this object
   in," but the parameterization is categorical rather than dimensional.

2. **Structural coordinate** — position in the call graph, AST,
   dependency graph, or git history. Has dimensionality but the
   dimensions aren't continuous in the same way 3D space is.

3. **Role assignment** — what role does this symbol play in the current
   task / agent / session? Discrete. Probably too coarse.

4. **No analog at all** — pose may be a load-bearing geometric concept
   that doesn't exist in the symbolic substrate, in which case the
   "object_id × pose" framing has to be replaced rather than translated.

## Why it matters

If we adopt Monty's `LearningModule` / `LMMemory` abstractions in MDEMG,
the hypothesis space is the central data structure each LM operates on.
A wrong choice here propagates everywhere.

This question is downstream of `Q-2026-04-28-001` (whether 3D geometry
is load-bearing in Monty). If the answer there is "yes, geometry is
load-bearing," then this question may be vacuous — there isn't a useful
analog and we need a different substrate.

## What we'd need to answer it

- Trace what Monty actually *does* with pose in `_update_evidence` and
  `receive_votes`. Specifically: is pose used as a continuous variable
  (compared with metric tolerance, transformed by rotation matrices,
  etc.) or as a hypothesis identifier (compared by index, voted on)?
- For each candidate analog above: trace what an MDEMG operation
  analogous to Monty's `_update_evidence` would look like with that
  candidate substituted for pose. Some candidates will fail mechanically.
- Check whether the hypothesis-resampling work in Monty RFCs 0009 / 0013
  changes the answer — those may be reframing the hypothesis space
  itself.

## Status notes

Open. Bootstrap question; no investigation yet.
