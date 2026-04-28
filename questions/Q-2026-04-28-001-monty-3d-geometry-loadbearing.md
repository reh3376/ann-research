---
id: Q-2026-04-28-001
type: question
slug: monty-3d-geometry-loadbearing
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 2
where: [T001, T003]
tags: [monty, substrate, geometry, reference-frames]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "tbp.monty: src/tbp/monty/frameworks/models/object_model.py"
  - "tbp.monty: src/tbp/monty/frameworks/models/evidence_matching/learning_module.py"
answer: ""
---

# Do Monty's algorithms exploit properties of continuous 3D geometry that have no useful analog in symbolic structure?

## The question

Monty operates over sensorimotor input: 3D Cartesian locations, surface
normals, principal curvatures, rigid-body transforms. Its core operations
— `apply_rf_transform_to_points`, displacement matching, evidence
accumulation over `(object_id × pose)` hypothesis spaces — are
mathematically expressed in continuous 3D Euclidean geometry.

MDEMG operates over symbolic input: code observations, conversation
events, constraint records, abstract concepts. Its substrate is a graph
with vector embeddings.

The question: are Monty's algorithms **portable** to MDEMG's substrate
(by replacing 3D coordinates with structural-graph coordinates), or do
they **exploit** properties of continuous 3D geometry — smoothness,
local linearity, well-defined rotation/translation — in load-bearing
ways that don't transfer?

## Why it matters

The answer determines whether reference-frame transfer is a
**translation problem** (substitute coordinate systems and proceed) or
a **substrate problem** (the algorithms need to be redesigned for a
discrete relational substrate, and the borrowed abstractions might be
contracts without algorithmic content).

If translation: the integration path is roughly what
`DE-2026-04-28-001` proposed (after proper tracing).

If substrate: what we'd be borrowing from Monty is the *theoretical
framing* (cortical columns, voting, reference frames, hypothesis
testing) but very little code-level architecture. The work would be
substantial new design, with Monty as inspiration rather than
implementation.

The honest answer might be partial: some algorithms transfer cleanly,
others rely on geometry in ways that don't.

## What we'd need to answer it

Primarily T001 (Monty trace), with attention to:

- What does `apply_rf_transform_to_points` actually do, and what
  property of 3D rigid transforms makes it possible?
- In `_update_evidence` and `_update_evidence_with_vote`, where are 3D
  distances and angles used? Are they used as *signals* (could be
  replaced by any well-defined distance metric) or as *coordinates*
  (the geometry is the data)?
- The hypothesis space `(object_id × pose)` — is "pose" a continuous
  parameterization that gets sampled, or a discrete enumeration over
  candidate orientations? If continuous, what's the symbolic analog?
- Displacement-based prediction: when Monty predicts the next sensed
  feature given an action-displacement, is the prediction a function of
  the metric properties of the displacement vector, or just of its
  identity?

T003 may need to draw on this to answer the substrate-level question.

## Status notes

Open. No investigation yet. Tagged to T001 and T003 for the trace work
that should produce the answer.
