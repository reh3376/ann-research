---
id: Q-2026-04-28-011
type: question
slug: knowledge-builds-skills-via-nesting
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 3
where: [T006]
tags: [hypothesis, recursive-composability, skills, nesting, framework-viability, research-target]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude during T006 Cycle 1 item 1 / item 10 triage"
  - "@reh3376's framing: 'It seems likely to me that this is how knowledge is turned into skills that translate into means of interacting with the world.'"
  - "@reh3376's settlement under sharpening: 'I have a strong intuition that knowledge builds skills is fundamental to the viable framework, but it remains to be tested to be proven.'"
  - "Item 10 (Recursive composability) — the architectural property this question hypothesizes the function of"
  - "Hawkins, Thousand Brains framing — distributed reference frames composing across cortical regions"
  - "Item 1 (Continuous learning via input-driven, prediction-error-driven topology growth) — the substrate this hypothesis sits on"
answer: ""
---

# Knowledge builds skills via nesting (research hypothesis)

## The question

When the architecture's recursive composability (item 10) is
exercised — when topologies receive other topologies' outputs as
inputs and grow reference frames over them — what does the
recursion produce? Specifically: **does nesting transform knowledge
into skills?**

A first-level topology grows reference frames over sensory inputs
— learning what objects are, what relations hold, what situations
mean. This is *knowledge* in the colloquial sense.

A higher-level topology takes those reference frames as its inputs
and grows reference frames over *them*. The hypothesis is that the
result is *skills*: ways of interacting with the world, structured
as meta-frames over the knowledge frames they operate on. Skills
are what knowledge becomes when it is composed at higher levels of
the recursion.

If true, this answers a load-bearing question for the framework's
viability: **how does the architecture move from passive
representation to active capability?** Without an answer, the
architecture risks being a sophisticated knowledge representation
with no story for how it acts on what it knows.

## Why this is filed as a question, not a must-have

@reh3376 articulated this carefully under sharpening: *"I have a
strong intuition that knowledge builds skills is fundamental to
the viable framework, but it remains to be tested to be proven."*

The working method T006 operates under requires distinguishing:

- **Must-have property statements** = commitments the architecture
  must satisfy. Defensible today. Falsifying a must-have means
  rejecting the architecture.
- **Hypotheses** = claims the program is investigating, with
  recorded prior confidence. Falsifying a hypothesis means
  updating the claim, not rejecting the architecture.

The knowledge-builds-skills mapping has a strong prior (intuition
about framework viability) but does not yet have empirical or
formal grounding. Putting it in a must-have property statement
would commit the architecture publicly to a claim whose evidence
is forthcoming, not present. Filing as a research hypothesis with
explicit prior-confidence framing protects the must-have list's
integrity while recording the claim with full force.

This Q-artifact is the program's first formally-filed
**strong-prior research hypothesis** — a claim whose prior
confidence is substantial but whose evidence is to be produced by
future R&D cycles. The pattern is: file as Q-artifact with
explicit prior confidence; track evidence for and against as
cycles produce it; promote to a decision artifact or must-have
when evidence accumulates sufficiently; or revise / retract if
evidence accumulates against.

## The hypothesis as @reh3376 stated it

> *It seems likely to me that this is how knowledge is turned
> into 'skills' that translate into means of interacting with
> the world.*

Three sub-claims are implicit and may need separate tracking as
the hypothesis matures:

**Sub-claim (a) — Skills emerge from nesting.** When the
recursion is exercised, what emerges at higher levels has
skill-like structure (procedural, action-oriented, capable of
being applied to multiple knowledge configurations).

**Sub-claim (b) — The level of nesting at which skills emerge is
structurally identifiable.** It is possible to look at the
architecture and locate where knowledge ends and skills begin —
the levels are distinguishable, not blurred.

**Sub-claim (c) — Higher-level nesting produces meta-skills.**
Skills that compose other skills emerge as the recursion deepens.
Skill-of-skills, meta-strategies, generalized interaction
patterns.

@reh3376 has stated (a) directly. (b) and (c) are inferences from
(a) that may or may not be part of his commitment; flagged here
for tracking but not yet pinned down.

## What would constitute evidence

**Deferred** at the time of filing. To be specified as the program
matures and as Phase A produces architectural specifications
sharp enough to design empirical tests against.

A first-pass criterion that may evolve: in a working prototype of
the architecture, exposed to sensory data and then to
interaction-task contexts, can a structurally-distinct level of
nesting be identified at which the topology's outputs are
*action-shaped* (recommending what to do given a knowledge
configuration) rather than *representation-shaped* (describing
what the knowledge configuration is)? If yes, with the
action-shaped outputs demonstrably grounded in the knowledge
representations beneath them, this would be substantial evidence
for sub-claim (a).

This criterion is preliminary. Refinement of evidence criteria
is itself part of what the R&D cycles will produce.

## Where this question lives operationally

**T006 thread.** The hypothesis is internal to the
architectural-innovation work. Its resolution shapes how item 10
(recursive composability) is understood — as a property whose
function is partly known (composes topologies) and partly
hypothesized (composes them in a way that produces skills).

**Phase A → Phase C transition.** When the program transitions
to building a prototype, this question becomes one of the
empirical claims the prototype tests. Knowing now what a key
testing target is may shape how the prototype gets specified.

**Cross-references:**
- **Item 10 (Recursive composability)** — the architectural
  property this question hypothesizes the function of. Item 10's
  property statement does not commit to this hypothesis; it
  references this Q-artifact as the leading hypothesis about
  what the recursion produces.
- **Item 1 (Continuous learning via input-driven, prediction-
  error-driven topology growth)** — the substrate that supports
  the recursion within which skills are hypothesized to emerge.

## Confidence

**Question-confidence:** 3. The question is well-formed; what is
being asked is clear; the conditions under which it could be
answered are partially specified (with refinement deferred). Not
5 because the evidence criteria are still preliminary.

**Prior confidence in the answer being yes:** High but not
formally calibrated. @reh3376's framing — "fundamental to the
viable framework" — is a strong prior. Pending either (a)
empirical results from R&D cycles, or (b) a logical argument
constructed during T006's mathematics mode, the prior is the
ground for tracking this as a strong-prior hypothesis rather
than treating it as casual speculation.

## What this question is not

- Not a commitment that the architecture will produce skills via
  this exact mechanism. The architecture commits (via item 10) to
  recursive composability; what the recursion produces is what
  this question asks.
- Not a falsifying test of the architecture. If the hypothesis is
  disproved, the architecture's recursive composability still
  stands; what changes is the program's understanding of what
  the recursion is good for.
- Not exhaustive. There may be other things the recursion
  produces beyond skills-from-knowledge. This question tracks the
  specific knowledge → skills hypothesis; other hypotheses may
  emerge and be filed separately.

## Status updates expected

This Q-artifact should be revisited:

- When item 10's property statement is sharpened in a future
  cycle and the relationship to this hypothesis becomes clearer
- When Phase A transitions to specification work and empirical
  test design begins
- When Phase B begins and the question of what to communicate
  publicly about the framework's expected capabilities arises
- Periodically during T006's mathematics mode, in case logical
  grounding for the hypothesis can be constructed

Each revisit may refine the sub-claims, sharpen the evidence
criteria, or update the prior confidence based on accumulated
reasoning.
