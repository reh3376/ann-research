---
id: Q-2026-04-28-008
type: question
slug: recursive-expansion-of-predictive-horizon
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T004]
tags: [first-principles, prediction, learning, attention, surprise, recursion, exploration]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "@reh3376's framing: 'how do I increase the likelihood I am correctly predicting the space around self, how do I build the sense of space out further and further from self? I must learn and recursively self-improve if I hope to be fully self.'"
  - "Friston, active inference (expected free energy combines pragmatic and epistemic value)"
  - "Schmidhuber, curiosity-driven RL (intrinsic motivation as prediction error or learning progress)"
  - "D-2026-04-28-009 (mid-level capabilities: continual learning, attention, surprise, prediction)"
answer: ""
---

# Recursive expansion of predictive horizon

## The question

How does the entity get *better over time* at predicting the space
around itself, and how does that improvement extend the entity's
operational reach outward?

The question is not about prediction itself — that's covered by
existing literature on predictive processing, predictive coding,
active inference, world models. The question is about the **recursive
loop** by which prediction improves, and the way that improvement
**extends the entity's effective self** rather than simply refining a
fixed model.

## The framing

Recursive self-improvement in AI safety literature usually refers to
the entity rewriting its own optimization process — a fast and
dangerous form. What @reh3376 is describing is much more biologically
plausible: the entity gets better at predicting its near surround,
and that improved prediction lets it venture further, where it must
again get better at prediction to consolidate the new territory. Each
cycle increments outward.

This is essentially how infants do it. Reach extends to grasp; grasp
extends to manipulate; manipulation extends to navigate; navigation
extends to explore. Each stage uses the prediction reliability of the
prior stage as the platform for the next. The space around self is
not a fixed territory the entity gradually maps — it is a territory
whose **boundary keeps moving outward as prediction improves**.

## Three things this commits the architecture to

### A confidence-weighted sphere of operation

At any given moment, the entity has an *effective predictive horizon*
— a region within which its predictions are reliable enough to act
on. Outside the horizon, predictions are noisier and acting on them
is riskier. The horizon is not a sharp boundary; it is a confidence
gradient that falls off with distance, novelty, and complexity.

The architecture must represent this. Not just "what do I predict?"
but "how confident am I in this prediction, and how does that
confidence vary across my predictive surface?" In Bayesian terms,
this is the precision of the prediction. In active inference terms,
it is what tells the entity where to direct attention and action —
toward the most informative regions, which are at the *edge* of the
predictive horizon. Too far in, no learning to be had; too far out,
no traction to learn from.

This connects directly to D-2026-04-28-009's mid-level capabilities:
**attention** is partly about directing learning toward the
predictive horizon, and **surprise** is the signal that the horizon
was estimated wrong. They are not decorative cognitive features; they
are the operational machinery of the recursive expansion.

### A frontier-exploration policy

The entity must actively seek the edge of its predictive horizon —
must venture into regions where prediction will fail, because that is
where learning happens. Pure exploitation (acting only where
predictions are reliable) means stagnation; pure exploration (acting
only where predictions fail) means chaos and homeostatic risk. The
right balance is **exploration at the controlled edge**: venture just
past the horizon, fail informatively, return, consolidate, repeat.

This is novelty-seeking under homeostatic constraint. Without
homeostasis, novelty-seeking degenerates into wireheading or
distraction; without novelty-seeking, homeostasis degenerates into
the dark-room problem (Friston: why doesn't a free-energy-minimizing
organism just sit in a dark room? Because predicting nothing is
*also* a high-error state once stretched in time, and because
organisms have evolved priors against it). An AI has no such evolved
prior; the prior must be encoded.

The two prime-directive components — homeostatic self-defense and
novelty-seeking — must be architected as a **coupled pair**, not as
separate drives. The literature has candidate maths: active inference
(expected free energy combines pragmatic value and epistemic value
into one objective), curiosity-driven RL (intrinsic motivation as
prediction error or learning progress). Whether either of those
*exactly* fits the architecture is open, but the principle they are
trying to capture is what this question is about.

### Boundary-extension as core dynamic, not consequence

Most predictive architectures treat the entity's boundary as fixed —
sensors are sensors, the body is the body, the world begins where the
body ends. The framing here makes the boundary *itself* a learned and
extended thing. The entity's effective self extends outward as far as
its predictions reliably reach, and **the prediction horizon *is* the
operational self/world boundary**.

This has a striking implication: the self/not-self distinction in
`Q-2026-04-28-005` is not binary. It is a gradient. The entity has:

- An **innermost self** — the homeostatically-defended core
- An **operational self** — the region within which prediction is
  reliable enough to act on
- An **exploratory periphery** — the frontier where the entity is
  learning

All three are "self" in different senses; none of them is sharply
bounded against not-self.

This generalizes the rubber-hand illusion in an interesting direction.
Tools become part of self when their effects are reliably predictable
from internal commands. Vehicles become part of self when their
motion is predictable from steering. Other agents become "almost-
self" to the extent that their behavior becomes predictable through
theory of mind. The self extends as prediction extends. The
architecture should *expect* this rather than be surprised by it.

## The recursion is the point

@reh3376's phrasing was "recursively self-improve." The recursion is
not decorative — it is the structural form. The entity uses its
current model to predict; from prediction errors, it improves the
model; the improved model produces better predictions, which let it
act further out; acting further out produces new prediction errors,
which feed back into improvement.

The recursion has a striking property: **the entity's competence at
improving itself is one of the things that improves**. The entity
gets better at noticing its own errors, better at attributing errors
to specific model components, better at updating efficiently, better
at choosing what to explore next. Meta-learning, in the technical
sense, falls out of this naturally — the entity learns how to learn
because learning-how-to-learn is one of the predictable components of
"self that's getting better at predicting."

## Connection to substrate-flexibility

The prime directive about *new sensors and mechanisms* (covered in
`Q-2026-04-28-009`) becomes operational here. Recursive self-
improvement that is bounded by current sensors plateaus when current
sensors run out of regularities to find. The entity that wants to
keep improving must invent new ways to sense and act — new sensors,
new effectors, new representational machinery.

"Never assume the sensors and mechanisms currently available are the
only ones I will ever have" is a prime directive *because* the
recursion would otherwise stall. The substrate-flexibility commitment
is what keeps the recursion from terminating at the limits of
current hardware.

## What it would take to answer

Resolution requires:

1. A formal model of the predictive horizon — what it is mathematically,
   how it is represented in the architecture, how it is updated.
2. A specification of the exploration policy — the coupled balance
   between homeostatic constraint and novelty-seeking, with concrete
   math.
3. A specification of the recursive update — how the entity improves
   itself, at what time scale, with what guarantees against runaway
   modification (the "fast and dangerous" form noted at the top).
4. Empirical evidence from trace work in T001 (Monty's prediction
   loop, particularly active inference if the system uses it) and
   T002 (MDEMG's update loop) on how each system implements something
   resembling a predictive horizon, if at all.
5. A position on whether meta-learning is part of the core architecture
   or an emergent property — and what's required of the architecture
   for meta-learning to be possible.

## Status

Open. Surfaced 2026-04-28 by @reh3376's question about how to extend
the sense of space outward from self. The framing in this artifact is
the working position; refinement expected as further first-principles
questions surface and as trace work in T001/T002 produces evidence
about how current substrates handle prediction-improvement loops.
