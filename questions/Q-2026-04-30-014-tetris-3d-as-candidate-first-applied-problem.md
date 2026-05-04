---
id: Q-2026-04-30-014
type: question
slug: tetris-3d-as-candidate-first-applied-problem
created: 2026-04-30
updated: 2026-04-30
status: open
confidence: 3
where: [T006]
tags: [first-problem, phase-1, training-task, hypothesis, strong-prior, 3d-tetris, NP-completeness, deferred-to-phase-c]
provenance:
  - "conversation 2026-04-30 between @reh3376 and Claude during T006 Cycle 1 foundational-assumptions exercise"
  - "@reh3376's framing: 'I have been considering the use of a game like 3-D tetris as an initial training mechanism. This is where the original NP-completeness issue was birthed in my thought process.'"
  - "@reh3376's clarification on Phase 0 vs Phase 1: 'When I say first training, I am not including the intial state assumption training for Bayes and advanced mathematical concepts. I mean the first problem we will attempt to solve once the framework and its pre-requisites are satisfied'"
  - "@reh3376's hypothesis-not-fact framing: 'I lean towards M-1 as well. It is a hypothesis not a fact.'"
  - "@reh3376's commitment-strength framing (C-leading): 'I am currently saying Tetris is the leading hypothesis but I likely am not familiar with all possible candidates and a better candidate may exist. I commit to 3D tetris based on my base knowledge but it may change if someone or something presents a better candidate.'"
  - "D-2026-04-30-012 (foundational assumption) — establishes the framework whose first applied problem this question concerns"
  - "Q-2026-04-30-013 (computational complexity prerequisites) — Tetris's NP-completeness is the link"
  - "Q-2026-04-28-011 (knowledge → skills via nesting) — precedent for strong-prior research hypothesis filed as Q-artifact"
answer: "Seed answer (working hypothesis, C-leading commitment): 3-D Tetris is the leading candidate for the architecture's first applied problem (Phase 1, post-substrate-initialization). The candidate is selected on the conjunctive basis of three reasons that co-justify it: spatial reasoning bootstrap, computational pressure (Tetris's NP-completeness), and pedagogical curriculum starting point. The hypothesis is held with strong prior confidence based on @reh3376's current knowledge; the commitment is revisable upon presentation of a better candidate."
---

# 3-D Tetris as candidate first applied problem (working hypothesis)

## The question

When the architecture transitions from Phase 0 (substrate
initialization with Bayesian and advanced-mathematical
foundations) to Phase 1 (first applied problem-solving), what
problem is the architecture's first applied target?

@reh3376 has been considering **3-D Tetris** as the candidate.
This Q-artifact captures the consideration as a working hypothesis,
records the reasoning behind it, and holds open the question of
whether better candidates exist.

## Hypothesis (Sense 1: whole-claim hypothesis)

> *3-D Tetris is the right first applied problem for the architecture
> at the Phase 0 → Phase 1 transition, on the conjunctive basis of
> three co-justifying reasons (spatial reasoning, computational
> pressure, curriculum starting point).*

This is **Sense 1** from the foundational-assumptions sharpening:
the whole claim is the hypothesis. If Tetris turns out to be the
right choice for reasons different from these three, the hypothesis
gets revised or replaced. The hypothesis's structure is more
falsifiable than a Sense-2 nested-hypotheses framing would be;
@reh3376 chose Sense 1 deliberately for cleaner falsifiability.

## Commitment strength: C-leading

@reh3376's framing: *"I am currently saying Tetris is the leading
hypothesis but I likely am not familiar with all possible candidates
and a better candidate may exist. I commit to 3D tetris based on my
base knowledge but it may change if someone or something presents a
better candidate."*

This is **C-leading** — a fourth commitment-strength category
distinct from C-strong (likely commitment), C-weak (one of several),
and C-deferred (too early):

- **Working leading candidate:** Tetris is what the program proceeds
  with absent better information
- **Confidence basis:** @reh3376's current base knowledge
- **Revision condition:** a better candidate presented (by someone
  or something — including the program's own subsequent thinking)
- **Default behavior:** absent revision, Tetris is the choice

This is honest about epistemic state — committing where @reh3376
is while acknowledging information is incomplete.

## Phase boundary: G-clean

The Phase 0 → Phase 1 transition is treated as **discrete**: Phase
0 completes (substrate initialization with Bayesian and advanced-
mathematical foundations established and verified), then Phase 1
begins with the first applied problem. The phases do not overlap
in this hypothesis's framing.

What "verifiably ready" means at the Phase 0 → Phase 1 boundary
is itself an open question (related to but not identical to
Q-2026-04-30-013); deferred to construction work.

## The three co-justifying reasons (M-1 conjunction)

Per @reh3376's settlement, all three reasons co-justify the choice;
no single reason is sufficient on its own; the combination is what
makes Tetris the leading candidate.

### Reading X — Spatial reasoning bootstrap

3-D Tetris exercises spatial reasoning, geometric prediction, and
reference-frame construction over a simplified physical domain. The
architecture builds its first reference frames on a domain where:

- The rules are clear (well-defined game state, deterministic
  physics)
- The inputs are bounded (finite playing field, finite piece
  inventory, finite action space)
- The relationships between objects are spatially grounded (3-D
  geometric relations between pieces and the playing volume)

Under this reading, Tetris is selected because it is **a natural
bootstrap domain for an architecture committed to reference-frame
construction (item 3) and Hawkins-style cortical-column generalization
(item 1)**. The architecture's first reference frames form most
easily where the substrate's spatial-reasoning machinery is
exercised against a structured, bounded environment.

### Reading Y — Computational pressure

3-D Tetris is **NP-complete** in its general form — deciding optimal
piece placement to clear lines or maximize survival is provably
hard. The hardness is not incidental; it is part of the value:

- Encountering an NP-complete first problem **forces the architecture
  to develop sophisticated strategies** rather than relying on
  brute-force enumeration
- After Phase 0 has established Bayesian and advanced-mathematical
  foundations, Tetris **tests what the architecture can do with its
  foundations** — not whether it has them
- The architecture's response to Tetris (heuristic search, structural
  exploitation, learned approximation) is itself substrate-revealing

Under this reading, Tetris is selected because it is **harder than
an arbitrary tractable problem would be** — the difficulty is the
training value.

This reading directly raises Q-2026-04-30-013's question (what
computational-complexity prerequisites does the architecture need at
Phase 1 entry?). The two Q-artifacts are tightly coupled.

### Reading Z — Pedagogical curriculum starting point

3-D Tetris fits as the first item in a curriculum that progressively
introduces more complex domains. The choice of starting point is
partly pedagogical:

- Clear ground truth available (well-defined optimal play; large
  pre-existing literature on Tetris-playing algorithms; benchmark
  AI Tetris players exist)
- Bounded inputs simplify what's expected of the architecture's
  first sensory wiring
- Performance is measurable on multiple axes (clearing rate,
  survival time, planning depth, generalization to different piece
  sequences)

Under this reading, Tetris is selected because it provides the
**right shape for training and evaluation at the first stage**.
A more complex first problem (e.g., natural language understanding,
open-world physical reasoning) would not have the same
ground-truth and benchmarking properties.

## Why M-1 conjunction matters

Under M-1, the three readings co-justify; if any one of them turns
out to be wrong about Tetris, the case for Tetris is weakened.
Specifically:

- If Reading X is wrong (Tetris turns out not to exercise the
  relevant spatial-reasoning capabilities the architecture needs),
  the bootstrap argument fails and Tetris's value as first problem
  reduces
- If Reading Y is wrong (the architecture handles NP-completeness
  in some way that doesn't require first-problem training pressure),
  the computational-pressure argument fails and a tractable first
  problem may serve as well
- If Reading Z is wrong (the curriculum-starting-point reasoning
  is replaced by a different pedagogical structure), the
  curriculum-start argument fails and other choices become
  competitive

@reh3376's framing — *"It will be hard to commit to just one of
them"* — acknowledges that the case for Tetris is intrinsically
multi-faceted. Reading H-design from the sharpening conversation
(the multi-facetedness is substantive, not just epistemic) is what
this Q-artifact records.

## Sub-claims worth tracking separately

The hypothesis (Sense 1) is the whole claim. Three sub-claims are
implicit in how the hypothesis would be tested:

**Sub-claim (a) — Tetris exercises the right capabilities for
first-problem training.** Spatial reasoning, prediction over
geometric trajectories, plan-ahead-multiple-pieces, all in a way
that develops the architecture's reference-frame and prediction
machinery.

**Sub-claim (b) — The architecture's response to Tetris is
substrate-revealing.** How the architecture solves Tetris (or fails
to) tells us things about the substrate that other tasks would
not reveal.

**Sub-claim (c) — Performance on Tetris generalizes to subsequent
curriculum stages.** Capabilities developed during Tetris training
transfer to the next problems in the curriculum, justifying the
"first problem" framing.

These sub-claims are not separately committed; they are the testing
structure for the M-1 conjunctive hypothesis.

## Inviting alternative candidates

Per C-leading commitment strength, this Q-artifact explicitly
invites alternative first-problem candidates:

> *If a problem is identified that:
> 1. Better exercises spatial reasoning bootstrap (Reading X), OR
> 2. Provides equivalent or stronger computational pressure
>    (Reading Y) AT a level the architecture can productively
>    engage with, OR
> 3. Better fits as a pedagogical starting point (Reading Z),
> this hypothesis should be revised or replaced.*

Discoverability mechanisms:

- This Q-artifact is tagged with broad terms (`training-task`,
  `first-problem`, `phase-1`) that future searches will hit
- Cross-references from D-2026-04-30-012 and Q-2026-04-30-013
  make this artifact visible from the foundational-assumption
  work
- The C-leading status is named explicitly so future readers know
  the commitment is revisable

## Cross-references

- **D-2026-04-30-012** (foundational assumption) — establishes the
  framework whose first applied problem this question concerns
- **Q-2026-04-30-013** (computational complexity prerequisites) —
  Tetris's NP-completeness is the link; the two Q-artifacts are
  tightly coupled
- **Q-2026-04-28-011** (knowledge → skills via nesting) — precedent
  for strong-prior research hypothesis filed as Q-artifact (Sense 1
  framing inherited from this precedent)
- **T006 plan items 1, 2, 3, 10** — the architectural commitments
  Tetris would exercise during Phase 1
- **OBS-2026-04-28-004** — MDEMG observational work; MDEMG's
  experience of operating on real human-use cases provides context
  for what "first problem" should look like

## What this question is not

- **Not a Phase A commitment.** This is Phase C work (training-task
  selection for the constructed architecture). Filed in Phase A so
  the consideration is not lost; resolved in Phase C.
- **Not committing 3-D Tetris specifically as a final choice.**
  C-leading status means revisable upon better-candidate presentation.
- **Not committing any specific Tetris variant or rule set.** "3-D
  Tetris" is a class of related games; the specific variant is itself
  open (standard 3-D variants vary in piece set, board geometry,
  scoring rules).
- **Not the only training task.** The "first" problem framing
  implies a curriculum; subsequent problems are not addressed by
  this question.

## Status updates expected

Revisited when:

- Phase C construction work begins and training-task selection
  becomes immediate
- A better candidate is presented (per C-leading revision condition)
- Q-2026-04-30-013's sub-readings are addressed (Tetris's
  computational-complexity demands are part of what's at stake)
- The curriculum structure is articulated (the post-Tetris stages
  may bear on whether Tetris is the right *first* stage)

## Confidence

**Question-confidence:** 3. The question is well-formed; sub-claims
identified; relationships to other vault content explicit. Not 5
because what counts as a "better candidate" is not specified
operationally; the M-1 conjunctive structure means evaluation
involves multiple weighing decisions.

**Prior confidence in the seed answer (Tetris as right choice):**
Moderate-high. @reh3376's C-leading framing reflects working-level
commitment based on current knowledge; the prior is strong enough to
proceed but explicit about its revisability.
