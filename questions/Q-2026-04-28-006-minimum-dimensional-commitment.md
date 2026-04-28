---
id: Q-2026-04-28-006
type: question
slug: minimum-dimensional-commitment
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T004, T003, T001]
tags: [first-principles, dimensionality, reference-frames, substrate, vision]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "@reh3376's framing: 'I want to be careful about committing too strongly to 3 dimensions. As humans we assume space has 3 dimensions but is this merely a limit imposed by our sensors?'"
  - "Q-2026-04-28-001 (whether 3D geometry is load-bearing in Monty's algorithms) — sharpened version"
  - "T003 (grounding question — overlapping concerns)"
answer: ""
---

# Minimum dimensional commitment

## The question

What dimensional, topological, and geometric commitments should the
entity's architecture make a priori, and what should the entity
discover empirically from its inputs?

The question is sharper than the original `Q-2026-04-28-001`, which
asked whether Monty's 3D geometry is load-bearing for Monty's
algorithms. The sharpening is: which of Monty's algorithms (and which
of any candidate architecture's algorithms) would *also* fail in any
non-3D-Euclidean setting — not just symbolic, but, say, 11-dimensional,
or non-Euclidean, or graph-structured?

The answer gates how much of any system's mechanisms transfer to the
program's eventual target, and how much is scaffolding tied to the
particular substrate it was developed in.

## The principle

The dimensionality of perceived space is empirical, not metaphysical.
Humans perceive 3D because human sensors and human scale produce
reliable 3D regularities — vision triangulates, vestibular system
detects 6 DoF acceleration, proprioception reports body configuration
in 3D. That is a fact about *the human sensorium / environment
relationship*, not about space itself. The actual world has whatever
dimensionality the regularities require, and the regularities depend
on the sensor.

A bee detects polarization patterns. A shark senses electric fields.
A pit viper sees infrared. Each of these gives access to regularities
humans don't see. An entity equipped with all of them simultaneously
would build reference frames over a richer feature space than humans
can. Whether you call that "more dimensions" or "more features" is
partly a vocabulary choice, but the implication for architecture is
the same.

And this is just current biology. An entity with a SQUID magnetometer,
a neutrino detector, gravitational-wave interferometry, or sensors for
things we haven't named yet would find structure invisible to any
biological organism. None of that is hypothetical — just expensive.

The principle: **never hardcode the dimensionality of the entity's
world**, because the entity's world is whatever its current sensor
suite can detect regularities in, and the sensor suite is contingent
and replaceable.

## Three spaces, only one of which is necessarily 3D

Distinguish three spaces because they have different implications for
architecture:

**Effective space.** The dimensionality of regularities the entity
can learn from. Depends entirely on sensors. Could be high. Could be
variable (different tasks find different effective dimensionality
from the same sensors).

**Body space.** The dimensionality the entity's effectors operate in.
For a wheeled robot, partly 2D. For a humanoid, 3D-but-constrained-by-
DoF. For a software agent with no body, none of the conventional
sense. **This is where 3D shows up most reliably**, because physical
effectors live in physical space — but the entity's cognitive
architecture should not conflate "my body is in 3D" with "my reference
frames must be 3D."

**Cognitive space.** The dimensionality of the abstract spaces the
entity reasons over. In principle unbounded. Concept spaces, social
spaces, planning spaces — all routinely high-dimensional and often
non-metric. Hardcoding 3D here would be absurd even for humans; we
do not reason about concepts in 3D.

The mistake to avoid: **letting body-space dimensionality leak into
cognitive-space architecture**. Monty does this in a particular way
— its `pose` is SE(3), 6-DoF — appropriate for body space, but baked
into the *cognitive* representation of objects (the hypothesis is
"object × pose"). For an embodied AI in a 3D body this might be fine.
For an embodied AI in a more general body, or for the abstract
reasoning the cognitive layer needs to do *about non-bodily things*,
it is a constraint that doesn't carry.

## The principle of minimum commitment

The architecture's right posture is to make the *minimum* dimensional
and topological commitment at every stage and let the entity discover
the rest. Roughly:

1. There's structure in the input.
2. Some inputs are *near* others — the entity discovers a neighborhood
   relation. (Topological structure.)
3. Distance might be meaningful — the entity discovers whether a
   metric exists. (Metric structure.)
4. The metric might be Euclidean — the entity discovers whether it is.
   (Geometric structure.)
5. The Euclidean space has *N* dimensions — the entity discovers *N*.
   (Dimensional structure.)

Each step is an empirical finding, not a designed-in assumption.
Steps 4 and 5 might not even be appropriate for some inputs (concept
spaces aren't Euclidean; graph structures aren't necessarily metric).
The entity that needs Euclidean 3D for its body and non-Euclidean N-D
for its concepts would arrive at both *as discoveries about its
inputs*, not as prior architectural commitments.

This is harder than it sounds. The math for "learn the dimensionality
and topology of your input space online from streaming data, while
also using that structure to guide further learning" is not a solved
problem at the scale needed. Manifold learning gives one-shot
dimensionality estimates from batches; topological data analysis gives
structural summaries; representation learning discovers latent
dimensionality but only relative to the architectural prior. An
online, continuous, sensor-modality-agnostic version is genuine open
research.

This is one of the places where the **mathematics** layer of the
framework/topology/math triple bites — the math is not yet there. See
`D-2026-04-28-009` on the framework/topology/math distinction; this
question lives in the math layer.

## How current substrates behave

**Monty** bakes in 3D more pervasively than was initially
characterized. Not just `pose` — the assumption shows up in:
- Sensor module outputs that include 3D location and 3D surface
  normal as primary features
- Vote messages that carry pose in SE(3)
- Object hypotheses indexed by `(object_id, pose)` with pose drawn
  from SE(3)
- Motor system operating on agent pose in 3D space

How much of this is incidental scaffolding vs. load-bearing for the
algorithms requires careful tracing. The original Q-001 will be
sharpened into specific sub-questions as T001 reads into the LM
internals.

**MDEMG** is potentially closer to dimension-agnostic. Its substrate
is graph-based, and a graph is a topological object first, becoming
metric/dimensional only if you give it those structures. MDEMG does
not bake in dimensionality the way Monty does. It is also possible
that MDEMG bakes in *other* commitments worth relaxing (assumptions
about discreteness, about node-edge structure, about update locality).
T002 will produce evidence on this when it begins.

## Architectural implications, if minimum commitment is adopted

Several:

1. **Reference frames must carry their dimensionality and topology
   as metadata, not as assumption.** A reference frame must be able
   to declare what kind of space it organizes, so consumers of that
   frame know how to interpret it.
2. **Algorithms must be specifiable in dimension- and topology-agnostic
   form**, with substrate-specific specializations as opt-in
   refinements rather than baked-in assumptions.
3. **The discovery process itself must be part of the architecture.**
   The entity discovers structure online; the discovery loop is not
   an offline preprocessing step.
4. **Some current literature must be re-read with this critique
   active.** Hawkins's columns, place cells, head direction cells,
   grid cells in ML form — most of those proposals carry 2D or 3D
   commitments that need to be examined for whether they are
   load-bearing or incidental.

## What it would take to answer

Resolution requires:

1. Trace evidence from T001 establishing which Monty algorithms are
   intrinsically 3D-Euclidean and which are dimension-agnostic at
   their core.
2. Trace evidence from T002 establishing what dimensional/topological
   commitments MDEMG implicitly makes.
3. A position on whether the architecture should pursue strict
   minimum commitment (with the math gap as research debt) or pragmatic
   commitment (with explicit dimensionality assumed at some level,
   noted as a limitation).
4. If strict minimum commitment is pursued: a research line on the
   online structure-discovery math. This is potentially years of
   work and overlaps with manifold learning, topological data
   analysis, and active inference literatures.

## Status

Open. Surfaced 2026-04-28. Sharpens `Q-2026-04-28-001` (originally
"is 3D geometry load-bearing in Monty's algorithms?") to a more
general question about dimensional commitment across the program.
