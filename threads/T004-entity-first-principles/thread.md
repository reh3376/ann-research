---
id: T004
type: thread
slug: entity-first-principles
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 4
where: []
tags: [meta, first-principles, prime-directives, self-modeling, vision]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude (post-PR-#6 merge)"
  - "D-2026-04-28-009 (program objectives — this thread is the internal-facing complement)"
  - "D-2026-04-28-010 (architectural commitment — T004's entity is specifically the superstructure in the hybrid architecture)"
  - "T003 (grounding question — overlapping but distinct concerns)"
owner: "@reh3376 (PI), Claude (drafts and surfaces)"
current_question: "What is the minimal first-principles framing of what kind of entity this program is designing — from the entity's own perspective, before it has any models of self or world?"
---

# Thread T004 — Entity First Principles

## Goal

Develop a coherent first-principles framing of what an entity built on
this program's eventual architecture *is*, *needs*, and *does*, from
the entity's own perspective rather than from the engineer's. This is
the internal-facing complement to `D-2026-04-28-009`'s external-facing
objectives — D-009 says what the program is *for*; T004 develops what
the entity is *like* and *requires*.

**Per `D-2026-04-28-010` (the central architectural commitment), the
entity in question is specifically the *superstructure* in the hybrid
architecture** — the learning, predicting, world-modeling system that
maintains state, reference frames, and a homeostatic boundary, and
that orchestrates an LLM language plug-in via a teacher-pupil
dynamic. T004's questions about self/not-self, dimensionality,
projection-and-anchoring, recursive prediction, and prime directives
are about the superstructure, not about the user-facing LLM and not
about the entity-as-monolithic-agent.

The thread exists because this question must precede architectural
commitments. An entity coming online with a topology similar to a
human infant — but with no evolved priors, no body shaped by selection,
no embedded homeostatic systems — will be confused without prime
directives. The prime directives, in turn, presuppose answers to deeper
questions about boundary, dimensionality, communication, and growth.
Capturing these now means future architecture work is anchored to
considered first principles rather than to engineering convenience.

The thread does **not** propose architecture. Architectural commitments
flow downstream of T004's questions, alongside trace evidence from
T001 and T002, and with `D-2026-04-28-010` as the central
specification.

## Current question

What is the minimal first-principles framing of what kind of entity
this program is designing?

The conversation that prompted this thread surfaced five facets that
appear to be different aspects of a single coherent picture. Each is
captured as its own `Q-` artifact, with T004 holding the framing that
ties them together:

- **`Q-2026-04-28-005`** — Self/not-self as homeostatic boundary
- **`Q-2026-04-28-006`** — Minimum dimensional commitment
- **`Q-2026-04-28-007`** — Projection-and-anchoring across dimensions
- **`Q-2026-04-28-008`** — Recursive expansion of predictive horizon
- **`Q-2026-04-28-009`** — Prime directives without constraining potential

Together these constitute a working sketch:

> An entity that exists by maintaining a homeostatic boundary; that
> orients in an empirically-discovered space of unknown dimensionality;
> that communicates through projection-to-common-subspace anchored by
> sensorimotor traces; that extends its operational self outward
> through recursive prediction-improvement; whose prime directives
> orient it toward learning and self-maintenance without constraining
> what it can become.

The current question is whether this sketch is correct, complete, or
missing facets. As more facets surface, they become additional Q-
artifacts. As specific facets stabilize into committed positions, they
become D- artifacts. Until then, T004 holds the sketch in
question-form.

## Method

**Conversational surfacing followed by deliberate capture.**

1. **Conversational surfacing.** `@reh3376` raises questions and
   working framings in conversation. Claude engages substantively,
   pushing on ambiguity and surfacing structural choices. Conversation
   continues until a topic-shape stabilizes enough that further
   conversation refines rather than re-derives.

2. **Deliberate capture.** When shape stabilizes, the substance is
   captured into artifacts — typically one Q- per facet, with full
   argumentative structure preserved. Capture happens before context
   loss; "preserve to memory" is the explicit handoff from
   conversation to vault.

3. **Refinement through trace evidence.** As T001 and T002 produce
   mechanism-level observations, those observations may bear on T004's
   questions (e.g., what Monty's `pose` actually represents bears on
   `Q-006`'s dimensional question; what MDEMG's update rules actually
   compute bears on `Q-008`'s recursion question). Cross-references
   land in the Q artifacts as trace work matures.

4. **Conversation continues unconstrained.** Captured artifacts do not
   freeze the conversation. The artifacts hold what's been developed;
   the conversation can continue developing new framings, refining
   existing ones, or identifying contradictions.

## Relationship to other threads

- **`T003` (grounding question)** — partial overlap. T003 asks what
  "where am I" means for a symbolic substrate; `Q-006` and `Q-007`
  address closely related territory at a more general level. Likely
  T003 becomes a specific instance of T004 once T004 stabilizes;
  alternatively, T003 stays as the substrate-specific arm and T004
  remains general. Resolution deferred until trace evidence in T001
  and T002 makes the relationship clearer.

- **`T001` (Monty trace) and `T002` (MDEMG trace)** — feed T004 as
  evidence sources. Specifically: Monty's commitments to 3D Euclidean
  pose are evidence for `Q-006`; Monty's voting protocol is evidence
  for `Q-007`; MDEMG's substrate is evidence for both. The trace
  threads are not subordinate to T004, but T004 reads their findings
  through a particular lens.

- **`D-2026-04-28-009` and `D-2026-04-28-010`** — north-star
  objectives and central architectural commitment, respectively.
  T004 is the internal-facing complement to D-009 and develops what
  D-010's superstructure must be like. D-009 says the program targets
  a "stateful, deterministic, hybrid Bayesian prediction machine"
  with continual learning, reference frames, world models, prediction,
  attention, and surprise. D-010 commits to those capabilities living
  in a learning superstructure with an LLM as language plug-in.
  T004 develops what the superstructure must be like from its own
  point of view.

## State

**Current.**

- Framing sketch above produced through 6 turns of conversation
  on 2026-04-28.
- Five Q- artifacts capture the facets surfaced so far.
- The framing is *working* — it has internal consistency and connects
  to existing program structure — but is not committed. No D- artifacts
  yet from this thread.
- @reh3376 has explicitly noted that more questions are likely to
  surface from related territory (metacognition, prime directives that
  haven't yet been fully articulated). T004 will accumulate these as
  they arise.

## Open work

- [ ] Continue surfacing questions in conversation; capture as Q-
      artifacts when shape stabilizes.
- [ ] Watch for facets the current sketch doesn't cover. Likely
      candidates: theory of mind / other-self distinction (currently
      noted as missing in `Q-2026-04-28-009`); identity preservation
      across discontinuities; relationship to ethics and prime-
      directive content (vs. the bare structural prime directives).
- [ ] As T001 and T002 produce trace evidence, cross-reference into
      the Q- artifacts where evidence bears on the open questions.
- [ ] At some point, individual questions will stabilize into
      commitments. When they do, draft D- artifacts that the questions
      become evidence for.

## Artifacts produced by this thread

- `Q-2026-04-28-005` — Self/not-self as homeostatic boundary
- `Q-2026-04-28-006` — Minimum dimensional commitment
- `Q-2026-04-28-007` — Projection-and-anchoring across dimensions
- `Q-2026-04-28-008` — Recursive expansion of predictive horizon
- `Q-2026-04-28-009` — Prime directives without constraining potential

## Log

### 2026-04-28 — D-010 commitment refines the entity in question

`D-2026-04-28-010` committed the program's central architectural
pattern: hybrid superstructure with an LLM as language plug-in,
governed by a teacher-pupil dynamic. T004's frame updated to reflect
that the entity in question is specifically the *superstructure*, not
a monolithic agent and not the LLM. The five Q- artifacts (Q-005
through Q-009) become more tractable under this refinement: Q-005's
homeostasis is what the superstructure defends; Q-006's dimensional
commitment is the superstructure's reference frames; Q-007's
projection-and-anchoring describes how the superstructure coordinates
with its faculties (including the LLM via AI2AI); Q-008's recursive
horizon is the superstructure's learning loop; Q-009's prime
directives govern the superstructure.

The Q- artifacts themselves are not edited in this update — the
substance was correctly framed at capture time, and the architectural
context now provided by D-010 sharpens them without contradicting
them. A future PR may add cross-references from each Q- to D-010, or
the Q-artifacts may be left as-is with D-010 as the contextual frame.

### 2026-04-28 — thread opened, five facets captured

Conversation post-PR-#6-merge developed a multi-faceted first-principles
framing for the kind of entity this program is designing. Captured the
substance into T004 plus five Q- artifacts before continuing
conversation, to ensure preservation across context boundaries.

The capture itself is a deliberate act: @reh3376 explicitly said
"capture this conversation to preserve it" before continuing. Future
conversation refines what's in the artifacts; the artifacts are the
durable substrate.

Next: continue conversation. Likely territory includes metacognition
(explicitly previewed by @reh3376), other-self / theory-of-mind, and
the relationship between bare structural prime directives and
ethically-loaded content directives.
