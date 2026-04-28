---
id: Q-2026-04-28-007
type: question
slug: projection-and-anchoring-across-dimensions
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T004, T003]
tags: [first-principles, communication, reference-frames, multi-modal, anchoring]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "@reh3376's framing: 'with a much more methodical approach to anchoring we should be able to work through the maths that provide commonality based on the dimensional space of the entity with the least dimensionality'"
  - "Q-2026-04-28-006 (minimum dimensional commitment — this question follows from it)"
answer: ""
---

# Projection-and-anchoring across dimensions

## The question

How do entities (or sub-modules within a single entity) operating in
representational spaces of different dimensionality, topology, or
modality communicate meaningfully?

The question has two parts: a **mathematical** part (what operation
maps representations between such spaces?) and an **anchoring** part
(what makes the operation meaningful rather than arbitrary?).

## The principle: projection to the common subspace

When entities of different dimensionalities need to communicate,
communication occurs in the subspace that both can represent — which
is the subspace defined by the lower-dimensional entity. Higher-
dimensional entities project their representations down to that
subspace; lower-dimensional entities pass theirs through unchanged.

Concretely: if Entity A operates in 3D and Entity B operates in 11D,
communication happens in 3D. B projects its 11-D state down to the
3-D subspace A can represent; A's representation passes through. B
knows the projection has lost information; A doesn't (and in many
cases doesn't need to).

@reh3376 noted that this is how humans communicate across cognitive-
space dimensionality differences — between specialists and generalists,
between domains of expertise. It is "a source of much frustration and
misunderstanding on occasion, but it is possible." Misunderstanding
happens precisely when one party doesn't realize a projection has
occurred — the generalist treats the projected-down version as the
full picture; the specialist forgets the generalist isn't operating
with the higher-resolution version.

## The math is partially solved

The general operation has names in different traditions:

- **Linear case.** If both entities have linear representation spaces,
  projection is matrix multiplication — find the basis of the lower-
  dimensional space within the higher-dimensional space and project.
  Well understood. The interesting question is *which* basis: the one
  that preserves the most variance? The one that preserves task-
  relevant structure? The one the lower entity actually has perceptual
  access to? Different choices, different costs.

- **Nonlinear case.** If representations are on manifolds (which they
  are for any realistic entity), the projection is to a submanifold.
  Math exists (manifold learning, diffusion maps, autoencoder
  bottlenecks) but is not unified.

- **Information-theoretic case.** If we treat the entities as
  information channels, the right framing is "what's the maximum
  mutual information achievable through a channel of capacity equal
  to the lower entity's representational bandwidth?" That's a rate-
  distortion problem. There's a literature.

- **Category-theoretic case.** If we want to be principled about
  *what* gets preserved through projection (which structure, which
  relations), the answer is functorial mappings between categories of
  representations. Heavier machinery, but the right kind of heaviness
  for getting "what survives projection" rigorous.

These are not competing — they are different levels at which the same
operation can be analyzed. For our purposes, the practical question is
*which level does the architecture need to operate at?*

Initial guess: the manifold one, because that's where the empirical
evidence about how learned representations actually look (curved,
locally-Euclidean, often low-effective-dimensional) lives. But the
information-theoretic framing is what gives us bounds, and the
category-theoretic framing tells us when we are losing structure that
matters. A complete architecture will likely need elements from all
three.

## The anchoring problem

The load-bearing word in @reh3376's formulation is "anchoring."
Projection without anchoring is meaningless: two entities with
different dimensionalities can each invent a 3-dimensional projection
of their representations, but if their projections are not anchored
to the *same underlying referents*, they are not communicating —
they are producing locally-coherent gibberish.

In human cognitive communication, anchoring happens through:

- **Shared physical referents** — pointing at the same object,
  sharing a sensorimotor experience, having both walked around the
  same building.
- **Shared symbolic referents** — agreed-upon words for things, math
  notation, formal definitions.
- **Bootstrapped shared experience** — the conversation itself
  becomes shared experience that grounds future projections.

For artificial entities, candidate anchoring strategies:

**Sensorimotor traces.** If two entities have both acted on the same
external world (even if they sense it differently), the action-
consequence pairs they each recorded are anchors. "1m forward" is the
same for both if both have a "1m forward" action that produces
commensurate sensory changes. Strongest form of anchoring because it
is grounded in physical regularity neither entity can hallucinate.

**Predictive consistency.** If entity A sends a representation to
entity B, and B can use it to make predictions that come true, the
representation is anchored — it is referring to something B can
verify. Wrong predictions reveal wrong anchoring. Weaker than
sensorimotor anchoring but doesn't require shared action.

**Iterated negotiation.** Two entities exchange representations,
predict each other's responses, update when wrong. Over time, they
develop a shared protocol that *is* their anchoring — emergent from
the back-and-forth rather than designed in. This is how humans
actually do most of it; Lewis-style signaling games converge this
way.

**Principled construction.** A common substrate (the lowest-
dimensional entity's space) plus axioms about what's preserved under
projection. Heaviest machinery. Not how biological systems do it, but
maybe how artificial ones can.

For an AI architecture aiming at multi-modal, multi-sensor, eventually-
multi-agent operation, the working position is that the answer involves
**sensorimotor traces as the bottom-of-stack anchor**, with predictive
consistency as the operating mechanism, with iterated negotiation as
the runtime protocol, with principled construction as the *spec* the
runtime is approximating. All four, layered.

## Architectural implications

Several:

1. **Reference frames must carry their dimensionality and structure
   as metadata.** When LM_A votes for LM_B, the vote message must
   include "this is a vote in a 3D Euclidean space" or "this is a
   vote in a 7D space with these topological properties." Otherwise
   B cannot know how to project A's vote into B's space, and the vote
   is meaningless. Monty's current vote messages do not carry this —
   pose is hardcoded as SE(3). This is one of the places the
   architecture needs to relax beyond what Monty currently provides.

2. **The projection operator itself must be learnable.** Entities
   cannot be born knowing how to project from N-D to M-D; they have
   to learn the projection from joint experience. This is closer to
   cross-modal representation learning than to fixed projection —
   which means an entity's interaction with another entity is itself
   a learning task.

3. **The "lowest-dimensional entity sets the protocol" rule has
   consequences for power dynamics.** An entity that can detect 11
   dimensions of structure but is communicating with a 3-D entity
   must operate in 3-D for that exchange. This is fine for cooperative
   cases; it is a vulnerability in adversarial cases (a high-
   dimensional entity that *could* notice deception in dimensions the
   low-dimensional entity cannot perceive must give up that advantage
   to communicate). Not architectural for now, but worth flagging —
   the prime-directive layer eventually has to address this.

4. **Intra-entity multi-modal fusion is the same problem.** If an
   entity has multiple sensor modalities of different intrinsic
   dimensionalities, its internal architecture is doing the same
   projection work across its own modules. The sub-architectures that
   handle vision (high-dim) and audition (lower-dim) and proprioception
   (different again) need to communicate within the entity, and they
   need the same anchoring story we are proposing for inter-entity
   communication. **The intra-entity case is just the inter-entity
   case with shorter cables.**

The fourth implication is significant: solving the dimensional-
translation question once solves both inter-entity multi-agent
communication and intra-entity multi-modal fusion. That is the kind of
architectural insight worth holding onto.

## What it would take to answer

Resolution requires:

1. A formal specification of the projection operation appropriate
   for the architecture (which of the four mathematical framings, at
   what level of detail).
2. A specification of the anchoring protocol — which of the four
   anchoring strategies, in what combination, with what fallback
   behavior when an anchor fails.
3. A demonstration in either MDEMG or Monty (or both) of cross-
   modality or cross-dimensionality communication with explicit
   anchoring. This may be premature given the trace-first strategy.
4. Re-examination of Monty's voting protocol with the projection-
   and-anchoring framing active — does the existing protocol implicitly
   have an anchoring story (e.g., shared object models acting as
   anchors), and if so, how does it generalize?

## Status

Open. Surfaced 2026-04-28. Closely tied to `Q-2026-04-28-006`
(minimum dimensional commitment) — projection-and-anchoring is what
makes minimum dimensional commitment workable in a multi-module or
multi-agent setting. T003 (the grounding question) addresses related
territory specifically for symbolic substrates; this question is the
more general form.
