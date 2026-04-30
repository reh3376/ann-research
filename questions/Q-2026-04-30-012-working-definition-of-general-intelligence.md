---
id: Q-2026-04-30-012
type: question
slug: working-definition-of-general-intelligence
created: 2026-04-30
updated: 2026-04-30
status: open
confidence: 3
where: [T006, T004]
tags: [definition, general-intelligence, agi, prediction, program-objectives, foundational]
provenance:
  - "conversation 2026-04-30 between @reh3376 and Claude during T006 Cycle 1 item 2 triage"
  - "@reh3376's seed framing: 'I think a strong definition of general intelligence would allow the answer to these questions to be reliably predicted' — referring to object trajectories through reference frames, complex mathematical formulas, and chain-reaction generalizations"
  - "D-2026-04-28-009 (program objectives) — names AGI as destination but does not define it"
  - "Item 2 (Generative modeling) — specifies what generative substrate is required to support reliable prediction across heterogeneous domains"
  - "Item 10 (Recursive composability) — pattern at every level; the recursion produces frames over which prediction operates"
answer: "Seed answer (preliminary, awaiting refinement): general intelligence is the property that enables reliable prediction across heterogeneous prediction domains — physical motion, mathematical structure, causal chains, and the indefinite further. The architecture's substrate-level commitments (items 1, 2, 10, and others) specify what is required to support such prediction; general intelligence is the capability that the substrate enables."
---

# Working definition of general intelligence

## The question

What is the program's working definition of general intelligence?

The vault has named AGI as the destination throughout (D-009
program objectives, T006 thread frame, the LinkedIn writing) but
has not yet committed to a definition of what general intelligence
*is* such that the architecture can be evaluated against it. This
question holds the definitional work open while the seed answer
informs the program.

## The seed answer (from @reh3376)

> *I think a strong definition of general intelligence would allow
> the answer to these questions to be reliably predicted.*

The "questions" referenced are heterogeneous prediction domains
@reh3376 named in the same conversation:

- **Physical-motion prediction:** the position of an object moving
  through a reference frame at a future point
- **Mathematical-structure prediction:** the answer to a complex
  mathematical formula
- **Causal-chain prediction:** what will happen if an event triggers
  a chain reaction
- **Indefinite further** — the "etc." was explicit; the list is not
  exhaustive

The seed definition is therefore: **general intelligence is the
property that enables reliable prediction across heterogeneous
prediction domains.**

## Why this is filed as a question, not a decision

The seed phrasing is preliminary, not finished. Three sub-questions
need to be addressed before the definition stabilizes:

**Sub-question 1 — What counts as a "prediction domain"?**

The seed names physics, mathematics, and causal-chain reasoning as
domains. These are extensive but unsystematic. Is a "domain" a
problem-class (physics, math, causality)? A modality (visual,
auditory, linguistic)? A reference-frame family (spatial,
conceptual, social)? Some combination? The choice matters for what
"heterogeneous" means.

**Sub-question 2 — What counts as "reliable" prediction?**

Reliable in what sense? Probabilistically calibrated (the model's
stated probabilities match observed frequencies)? Practically useful
(the prediction supports successful action)? Better than baseline
(some specified comparison)? Convergent (the model's predictions
improve with experience)? Different choices imply different
empirical tests of the architecture.

**Sub-question 3 — Is intelligence the *enabling property* of
reliable prediction, or *measured by* reliable prediction?**

This is a real ontological question, not a semantic one:

- **Enabling property reading:** intelligence is the underlying
  capability; reliable prediction is one of its visible
  consequences. The architecture has intelligence; we observe
  prediction as evidence.
- **Constitutive reading:** intelligence *is* the property of
  reliably predicting across domains. There is nothing additional
  the architecture has beyond the capacity to predict; the
  architecture's intelligence is identical to its predictive
  capacity.

The two readings have different consequences for what the program
is building. Under the enabling reading, the architecture must
have *something more* than predictive capacity; reliable prediction
is evidence of that something more. Under the constitutive reading,
the architecture's predictive capacity *is* its intelligence;
nothing additional is required.

This question may receive different answers from different members
of the AI research community. The program needs to commit to one
reading or articulate why both are compatible.

## Why this question matters now

Three places the definition bears immediately:

**1. Item 2 (Generative modeling) is partly defined by it.** Item 2
commits to hybrid α+β with action-conditioning, per-level
generative models composed by a unifying superstructure. The
*reason* these specific commitments are made — rather than weaker
alternatives — is that they are required to support the kind of
prediction the seed definition specifies. If the definition shifts,
item 2's commitments may shift.

**2. Phase B (Communication) needs a defensible definition.** When
the program articulates publicly what it is working toward, the
working definition is what audiences will press on. "Reliable
prediction across heterogeneous domains" is a strong defensible
position; it is also one of several positions in the field. The
program should commit to its definition with awareness of where it
sits in the broader debate.

**3. The relationship to existing AGI definitions matters.** The
field has accumulated multiple definitions:

- *Behavioral / task-performance:* AGI is a system that can perform
  any cognitive task a human can perform, at human level or better
  (the loose Turing-test descendants)
- *Generality / compositional:* AGI is a system whose competence
  generalizes outside its training distribution by composing
  primitive capabilities (Chollet's ARC framing, others)
- *Meta-learning / adaptation:* AGI is a system that learns how to
  learn, adapting rapidly to new domains (recent meta-learning
  literature)
- *Predictive / world-modeling:* AGI is a system that maintains an
  accurate generative model of its world and uses it for action
  selection (LeCun's recent framings, Hawkins' more limited claim)

The program's seed definition is closest to the predictive /
world-modeling family but with explicit emphasis on heterogeneity
across domains. Articulating the program's definition clearly
positions it within the field's debate.

## Sub-claims in the seed definition

Three sub-claims the seed answer implicitly contains:

**Sub-claim (a) — Prediction is the central operation.** General
intelligence reduces to (or fundamentally involves) prediction.
Other cognitive capacities (memory, reasoning, planning, learning)
are secondary or are themselves species of prediction.

**Sub-claim (b) — Heterogeneity is required.** A system that
predicts well in one domain (or a few) is not generally
intelligent. Intelligence is constitutively about cross-domain
generality of the predictive capacity.

**Sub-claim (c) — Reliability is the success criterion.** A system
whose predictions are unreliable is not generally intelligent,
even if it predicts widely. Reliability is the test, not just
breadth.

Each sub-claim is defensible but not uncontroversial. Each may be
revised independently as the question is sharpened.

## What would constitute progress on this question

The question stays open for multi-cycle work. Progress looks like:

- Sub-questions 1-3 receiving sharpened answers
- The relationship to existing AGI definitions being articulated
  (positioning move)
- The relationship between this definition and the architectural
  commitments (items 1-N) being made explicit
- Possibly graduation to a decision artifact when the definition
  stabilizes, with the seed answer refined into committed phrasing

The question may also be revised based on:

- Empirical findings from the eventual prototype (Phase C)
- Engagement with the broader AI research community's framings
- New cognitive capacities surfacing that don't fit the
  predictive framing cleanly

## What this question is not

- **Not a placeholder.** The seed answer is substantive; the
  question holds it open for refinement, but the seed itself
  informs the program now.
- **Not a definition of intelligence in general.** The question is
  scoped to *general* intelligence as the program's destination.
  Narrow intelligence, animal intelligence, and human intelligence
  are not in scope.
- **Not a commitment to a specific empirical test.** Defining
  intelligence is upstream of testing for it; the question
  addresses the definition; tests follow once the definition
  stabilizes.

## Status updates expected

This question should be revisited:

- Periodically as item 2's content matures (the architecture's
  commitments inform what the definition needs to support)
- When Phase B begins (the definition becomes externally
  consequential)
- When new cognitive capacities surface in the program's work
  that may extend or complicate the predictive framing
- If the broader AI research community shifts in ways that make
  the program's positioning clearer or more contested

## Confidence

**Question-confidence:** 3. The question is well-formed; its
relationship to other vault content is clear; the sub-questions
that need addressing are identified. Not 5 because the relationship
to existing AGI definitions and the answers to the three
sub-questions remain to be developed.

**Prior confidence in the seed answer's direction:** Moderate. The
prediction-centric framing is consistent with predictive-coding,
active-inference, and world-model traditions; it is one of several
defensible framings; @reh3376 has held it consistently in
conversation but the program has not formally pressure-tested it
against the alternatives.
