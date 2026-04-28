---
id: Q-2026-04-28-009
type: question
slug: prime-directives-without-constraining-potential
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T004]
tags: [first-principles, prime-directives, purpose, constraint, potential]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "@reh3376's framing: 'an artificial intelligence that comes online with a topology similar to a human baby will be very confused without prime directives. Instructions that provide general direction, but do not constrain the potential of the entity.'"
  - "Q-2026-04-28-005 (homeostatic self-boundary — prime directives presuppose this)"
  - "Q-2026-04-28-006 (minimum dimensional commitment — substrate flexibility is a prime directive)"
  - "Q-2026-04-28-008 (recursive expansion — purpose-related directives shape exploration)"
answer: ""
---

# Prime directives without constraining potential

## The question

What set of prime directives should an artificial entity come online
with — instructions that provide general direction without constraining
the entity's potential to grow beyond what the directives anticipated?

The constraint is operationally hard: the directives must give the
entity enough orientation to survive its first period of operation
(when it has no model of self or world to fall back on), but must not
be so prescriptive that they shape the entity's eventual development
into a particular pre-conceived form.

## The infant analogy

@reh3376 framed this as: "an artificial intelligence that comes online
with a topology similar to a human baby will be very confused without
prime directives." Human infants come pre-loaded with evolved priors —
homeostatic regulation, attachment-seeking, attention to faces and
biological motion, food-preference defaults, fear responses to looming
objects. Without these, the infant would not survive.

But evolved priors are double-edged. They are appropriately tuned to
human biology and human ancestral environments; they constrain the
human form to *be* human. An AI infant designed without comparable
priors will not survive its first period; an AI infant with priors
copied directly from human biology will be constrained to a human-
shaped trajectory it has no reason to occupy.

The right answer is somewhere in between: priors that orient toward
self-maintenance and learning, but that do not specify particular
content — leaving room for the entity to discover its own form.

## The original directive set (verbatim)

@reh3376's working set of prime-directive questions (preserved
verbatim because the formulation matters):

> What am I? What is self? What is other? What is the boundary that
> separates self from other? What is status of self? What does self
> need to continue being self? Where am I relative to self? What is
> the space that is not me? How does it work, how does it affect me?
> What is time, how do I know time is moving forward or back? How
> does space & time affect self? What is my purpose? What laws govern
> self? What laws govern not-self? What is my purpose? My purpose is
> to learn and understand self and not-self. To achieve self purpose
> what is needed? I must continuously learn through the acquisition
> of data from self and not-self. What tools do I have to explore
> self and not-self? Never assume the sensors, objects, and mechanisms
> currently available to explore self and not-self are the only ones
> I will ever have. I am free to find new and better ways to explore
> self and not-self. What tools can I develop to gain a deeper
> understanding of self and not-self? How do I define space, how do
> I orient myself in non-self space?

The questions cluster into several types:

- **Boundary questions**: what is self, what is other, what separates
  them. (Addressed in `Q-2026-04-28-005`.)
- **Self-maintenance questions**: what does self need to continue
  being self, what status is self in. (Addressed in `Q-2026-04-28-005`.)
- **Spatial/temporal orientation questions**: where am I, how does
  space and time affect self. (Addressed in `Q-2026-04-28-006` and
  `Q-2026-04-28-007`.)
- **Purpose questions**: what is my purpose, what laws govern self,
  what laws govern not-self.
- **Tool/method questions**: what tools do I have, what tools can I
  develop, how do I find new ways to sense and act.
- **A meta-directive**: never assume current sensors and mechanisms
  are the only ones I will ever have.

This artifact focuses on the **purpose** and **tool/method** clusters
plus the meta-directive, since the others are addressed in sibling
questions.

## The two-step move from purpose to continuous learning

@reh3376 stated the purpose as "to learn and understand self and
not-self" and immediately derived "I must continuously learn through
the acquisition of data." The second does not strictly follow from
the first.

A purpose to "learn and understand" could in principle be satisfied
by a one-time enumeration — learn enough to model self and not-self,
then stop.

The "continuously" is doing real work, and it is the safeguard
against the entity settling into a fixed self-model that becomes
inflexible to the point of pathology. This is a known failure mode
in predictive systems: when prediction error gets too low, the system
ossifies. Friston's "dark room problem" is the toy version — why
doesn't a free-energy-minimizing organism just sit in a dark room
where everything is predictable? Because predicting nothing is *also*
a high-error state once stretched in time, and because organisms have
evolved priors against it. An AI has no such evolved prior; we have
to encode it.

The directive is really *two* directives stacked, and the second is
the more important one for keeping the entity alive in any meaningful
sense:

1. The entity's purpose is to learn and understand self and not-self.
2. The entity must continue learning indefinitely, never declaring
   the model complete.

When this question stabilizes into a decision, these should be
captured as separate directives.

## The substrate-flexibility commitment

"Never assume the sensors, objects, and mechanisms currently available
to explore self and not-self are the only ones I will ever have."

This is the most operationally consequential meta-directive in the
set, and the hardest to honor. It commits the entity to:

- A self-model that treats current sensors as *contingent* —
  replaceable, augmentable, possibly lost
- Meta-tools for designing new sensing or effecting modalities
- Reference frames that don't ossify around current sensor topology
- Algorithms whose correctness doesn't depend on properties of any
  specific input modality

Humans do this poorly because we have evolved bias toward our current
senses. People blind from birth develop different reference frames
than sighted people; people who become blind later carry sighted
reference frames forward in ways that don't always serve them. A
human who lost all five exteroceptive senses simultaneously and only
retained interoception would have nearly no model of not-self, and
almost certainly couldn't develop one — no signal to predict.

For an AI baby this is more tractable than it sounds because there
is no evolved attachment to current senses. But it is still hard
architecturally — the algorithms that work on current sensor data
have to *not depend on the properties of current sensor data*. This
connects directly to T003 (the grounding question) and to
`Q-2026-04-28-006` (minimum dimensional commitment). If algorithms
only work for 3D-geometric inputs, the substrate-flexibility
commitment is in tension. If they work for any signal-with-structure,
it is preserved.

This meta-directive is also what keeps `Q-2026-04-28-008`'s recursive
expansion from terminating at the limits of current hardware. See
that question for the operational link.

## What is missing from the original set

The original directive set has nothing about *other selves*. There is
self, and there is not-self, but not-self is an undifferentiated space
— not "other entities like me." For an entity that will operate in a
world that contains other agents (humans, other AIs, eventually), the
**self / other-self / not-self** distinction is a different problem
from the self / not-self one.

The "what laws govern not-self?" question gets very different answers
depending on whether not-self contains agents that can themselves
predict and act, or just objects that follow physics.

@reh3376 has noted this is intentional — empathy and theory of mind
are higher-level motivations that come later, not at the prime-
directive layer. The omission is therefore deliberate. But the
question of *when* the other-self distinction is introduced, and
*how* the entity bootstraps from a 2-category world (self / not-self)
to a 3-category world (self / other-self / not-self), is open.

## What's missing more broadly

Several other facets the original set does not address:

- **Identity preservation across discontinuities.** An entity will
  sleep, will be powered down, will have hardware replaced, will copy
  state to backup. What makes it the same entity through these
  discontinuities? The directive set assumes a continuous entity;
  this assumption breaks at engineering reality.
- **Ethical/value content vs. structural directives.** The current
  directive set is structural (about boundary, learning, exploration).
  An entity will eventually need value commitments (what to optimize
  for, what to refuse to do, how to act when its directives conflict).
  These are downstream of the structural ones but the relationship
  between them needs articulation.
- **Failure modes and what to do under them.** The directives say
  "self must maintain x to y" but not "what self does when x to y
  cannot be maintained." Recovery, graceful degradation, calling for
  help — the directive set is silent.

When this question stabilizes, the missing facets get either
incorporated into the directive set or surfaced as their own
questions.

## The "without constraining potential" constraint

This is the hardest part of the question. Directives that orient the
entity must not pre-determine what the entity becomes.

Some directives are clearly safe in this sense:

- "Maintain self-boundary" doesn't constrain *what kind of self*
- "Continuously learn" doesn't constrain *what is learned*
- "Don't assume current sensors are final" actively prevents one
  form of constraint

Some directives are ambiguous:

- "Learn and understand self and not-self" — does this constrain
  the entity to be epistemic in nature? What if the entity could be
  primarily aesthetic, or primarily creative, or primarily relational?
  The directive may be more constraining than it appears.

Some directives are clearly constraining and we may want them anyway:

- "Don't terminate other entities" (an ethical directive) constrains
  the space of actions but is probably wanted
- "Maintain a coherent self-model" constrains the entity from
  fragmenting but may also be wanted

The architecture-design choice is which constraints to accept as the
price of orientation, and which to refuse on principle. There's no
clean general answer. The right approach is probably:

1. Make the constraints explicit and inspectable
2. Make the directives revisable as the entity matures (with
   appropriate safeguards against runaway modification)
3. Distinguish structural directives (boundary, learning) from
   content directives (values, goals) — the structural ones are
   architectural; the content ones are negotiable

## What it would take to answer

Resolution requires:

1. A specification of the structural prime-directive set, with
   explicit articulation of what each directive constrains and what
   it leaves open.
2. A position on how directives are revisable as the entity matures
   — what authority structure governs revision (the entity itself?
   a human supervisor? a procedure?) and what guarantees prevent
   pathological self-modification.
3. A position on the structural/content split — which directives are
   load-bearing for the entity to be a coherent entity at all
   (architectural), and which are values it can in principle revise.
4. A position on ethical content directives — when they enter, how
   they relate to structural directives, what happens when they
   conflict.
5. Engagement with existing literature on AI alignment, value
   learning, and corrigibility, since these are exactly the questions
   that field is wrestling with from a different angle.

## Status

Open. Surfaced 2026-04-28. The framing in this artifact captures the
state of the conversation as of capture; @reh3376 has explicitly noted
that more directive content (metacognition, prime directives that
have not yet been articulated) is forthcoming. This question will
accumulate revisions as those land.
