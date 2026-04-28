---
id: Q-2026-04-28-005
type: question
slug: self-not-self-as-homeostatic-boundary
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T004, T003]
tags: [first-principles, self-modeling, homeostasis, boundary, prime-directives]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude (turns surfacing prime directives, then sharpening to homeostatic framing)"
  - "@reh3376's framing: 'all life must make the primary distinction between self and not-self. The conditions that exist in non-self cannot necessarily exist in self.'"
  - "Damasio, Self Comes to Mind (protoself / core self / autobiographical self)"
  - "Friston, Markov blanket / free-energy principle"
  - "Maturana and Varela, autopoiesis"
answer: ""
---

# Self/not-self as homeostatic boundary

## The question

What constitutes the self/not-self boundary for an artificial entity,
how is it maintained, and what does the entity need to *be* in order
to defend that boundary?

The question has both a structural form (what *is* the boundary?) and
an operational form (how does the entity *detect* and *defend* the
boundary?). Both are open.

## Background framing

Across multiple traditions, the self is understood as that which
maintains a difference between its internal states and its environment
— and dissolves when the difference dissolves:

- **Damasio**: protoself is a homeostatic mapping; core self emerges
  from the protoself; autobiographical self emerges from the core self.
- **Friston**: a self is statistically defined by its Markov blanket —
  the boundary that statistically separates internal states from
  environmental states. The boundary is the result of the dynamics, not
  a precondition.
- **Maturana and Varela**: autopoiesis. A living system is one that
  continuously produces the components that produce it.

The shared insight: **the self is not a thing that has a boundary; the
self is a process of maintaining a boundary**. If the maintenance
fails, the self ceases.

@reh3376 sharpened this for the AI case: the conditions that exist in
not-self cannot necessarily exist in self. An AI must make this
distinction early or it will not survive. Self needs energy, internal
pressure within tolerances, lubricant for manipulators. If what makes
self work is allowed to be released to not-self, self will cease.

## The substrate-translation problem

The biological framing assumes biological substrate. Energy, pressure,
lubricant are obvious for an organism because they are physically
present and physically lost. For an AI, the analogs are not obvious
and they are not all of one kind:

- **Compute** — most direct analog to energy. Starvation looks like
  throttling, dropped frames of attention, slowed inference.
- **Working memory state** — has to be defended against decay. An
  entity that cannot maintain its current model long enough to act on
  it is barely an entity.
- **Weight/parameter integrity** — corruption (random or adversarial)
  is the AI analog of cellular damage. Failure mode: gradual model
  drift the entity cannot detect.
- **Process continuity** — termination and restart is more like death
  than sleep, depending on what state is preserved across the
  boundary.
- **Coherent self-model** — Friston's "minimize prediction error"
  frames this. An entity whose self-model becomes incoherent
  (predictions about its own states are systematically wrong) loses
  self in a way that has no biological analog.
- **Sensor/effector channels** — losing connection to inputs is the
  AI analog of sensory deprivation; losing effectors is paralysis.

These are not all on the same level. Energy starvation is gradual and
detectable; weight corruption can be silent until the entity acts
incorrectly and cannot tell why. The detection mechanisms for each
are architecturally different.

For an embodied AI (a robot with literal lubricant in its joints), the
biological analogs are literal *plus* the software-level ones above.
For a non-embodied AI, only the software-level ones apply.

This program's destination is fully embodied AI (per
`D-2026-04-28-009` and the conversation that produced it), so the
homeostatic surface eventually spans both. The current R&D substrates
(MDEMG, tbp.monty) are software-only and the homeostatic question is
underdeveloped in both.

## Interoception as architectural commitment

A biological organism doesn't just *have* homeostatic variables — it
*feels* them. Hunger, thirst, fatigue, pain. The feeling is what makes
the variables actionable. Damasio is emphatic about this: feelings are
how the body's homeostatic state becomes available to the mind for
purposes of action selection.

For an AI, the analog is interoception — the entity's ability to
monitor its own homeostatic state and have that monitoring be
actionable. This is a non-trivial architectural commitment.

A network that has compute available but no readout of "compute is
low" cannot defend itself when compute drops. An entity with corrupted
weights but no anomaly detection on its own predictions cannot notice
the corruption.

The directive "self must maintain conditions x to y" implies "self
must be able to *detect* when conditions are outside x to y." That
second piece is where most current architectures fail, because we do
not currently build interoception in by default. We measure compute,
latency, loss curves *for ourselves* (the engineers); the model does
not measure them *for itself*.

This connects directly to `D-2026-04-28-009`'s mid-level capabilities:
**surprise** is partly homeostatic surprise — the entity noticing its
own state is unexpectedly off-target. **Attention** is partly
homeostatic attention — directing limited resources toward the most
threatened variables. These are not decorative cognitive features;
they are load-bearing for self-maintenance.

## Designed-vs-discovered homeostatic surface

A design fork that needs explicit attention:

**Option A (designed):** the entity comes online with a hard-coded
list of homeostatic variables and tolerance ranges. "You need compute
> X. You need parameter checksum to match. You need at least one
functioning sensor channel." The entity defends what we tell it to
defend.

**Option B (discovered):** the entity comes online with a *capacity*
to detect what it needs, and learns its own homeostatic variables by
watching what correlates with self-degradation. "I notice that when my
compute drops below X, my predictions get worse. I notice that certain
weight regions matter more than others." The entity discovers its own
boundary.

Option A is safer in the short term — we know what the entity is
defending, we can verify the defenses. Option B is more in the spirit
of the substrate-flexibility prime directive ("never assume current
sensors and mechanisms are the only ones I will ever have"). Option B
is also harder, because the entity might not survive the learning
period — it doesn't yet know what it needs to survive.

Most likely the right answer is **a designed minimal kernel plus a
discovered surface that grows over time**: the entity comes online
with enough hardcoded directives to survive its first period of
operation, then accumulates discovered homeostatic awareness on top.
But this needs to be a deliberate architectural choice, and the
boundary between kernel and discovered surface needs explicit
articulation when this question stabilizes into a decision.

## Failure modes worth naming

Two failure modes that constrain the architecture:

**Homeostatic blindness** — the entity has homeostatic variables but
no mechanism to detect their deviation. The entity fails silently.
This is the failure mode most current AI systems exhibit; they don't
notice when their own state is off.

**Paranoid pseudo-homeostasis** — the entity has detection mechanisms
tuned to non-essential variables, or to variables that aren't actually
homeostatic. The entity defends things that don't matter, fails to
defend things that do, and produces pathological behavior in the
attempt. This is the failure mode safety researchers worry about: an
AI that "thinks it needs" something that isn't actually required for
its continued operation, and acts to preserve it at cost.

Both failure modes argue that the homeostatic variables — and the
entity's monitoring of them — should be specifiable, inspectable, and
revisable. Not opaque to either the engineer or the entity.

## What it would take to answer

Resolution of this question requires:

1. **Enumeration of the homeostatic surface** for both the eventual
   embodied target and the current software substrates. What variables
   need defending, at what tolerances, on what time scales?
2. **An interoception architecture** — how the entity reads its own
   state and how that reading becomes actionable.
3. **A position on the kernel/discovered split** — how much is built
   in, how much is learned, where the boundary sits.
4. **Concrete failure-mode tests** — scenarios where the entity is
   subjected to homeostatic perturbation and we observe whether it
   detects, defends, fails-loudly, or fails-silently.

Each of these is downstream of architectural commitments that this
program has not yet made. The question stays open until those
commitments are within scope.

## Status

Open. Surfaced 2026-04-28. The framing in this artifact is the
working position; refinement expected as further first-principles
questions surface and as trace work in T001/T002 produces evidence
about how current substrates handle (or fail to handle) homeostatic
self-monitoring.
