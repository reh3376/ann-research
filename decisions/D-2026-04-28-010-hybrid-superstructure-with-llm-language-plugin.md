---
id: D-2026-04-28-010
type: decision
slug: hybrid-superstructure-with-llm-language-plugin
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 5
where: [T004, T001, T002, T003]
tags: [architecture, central-commitment, hybrid, llm, superstructure, teacher-pupil, lora-adapters, ai2ai-protocol]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude (post-PR-#7 merge)"
  - "@reh3376's framing: hybrid system, LLM as language plug-in, superstructure as the 'real power'"
  - "@reh3376's elaboration: teacher-pupil dynamic, default non-intervention, override with generation"
  - "@reh3376's elaboration: superstructure has its own generative LLM for internal reasoning; AI2AI protocol for inter-LLM communication; LoRA adapters as the continuous-learning surface"
  - "MDEMG: github.com/reh3376/mdemg — Jiminy / coding-agents pattern as small-scale precedent for the architecture"
  - "MDEMG: AI2AI communications protocol (specifics to be captured in T002 trace work)"
  - "D-2026-04-28-009 (program objectives — D-010 is the central architectural commitment that operationalizes D-009's three-layer objectives)"
  - "T004 (entity first principles — D-010 specifies what kind of entity is being designed)"
alternatives_rejected:
  - "LLM-centric architectures with augmentations (current mainstream: RAG, tools, agents)"
  - "Pure non-LLM architectures (would require re-deriving language capabilities the LLM has compressed from training)"
  - "End-to-end neural architectures that don't separate language from cognition"
  - "Two-LLM mode of single model rather than hard architectural split"
  - "Full base-model fine-tuning as the LLM-layer learning surface (rejected in favor of LoRA adapters)"
  - "Natural-language inter-LLM communication (rejected in favor of token-efficient AI2AI protocol)"
---

# Hybrid superstructure with LLM as language plug-in

## Decision

The architecture this program is designing is a **hybrid system** with
the following structure:

**The superstructure is the entity.** A learning, predicting,
world-modeling system that maintains state across time, develops and
revises reference frames, defends a homeostatic boundary, and extends
its operational self outward through recursive prediction-improvement.
The superstructure is the cognitive core, the seat of self, and the
locus of continual learning. Everything in `D-2026-04-28-009`'s mid-
level capabilities (reference frames, world models, continual
learning, prediction, attention, surprise) lives in the superstructure.

**The LLM is a language plug-in.** A pre-trained base model that
provides language capabilities to the entity. The LLM is *not* the
cognitive core. The LLM has no state across calls, no self-model, no
reference frames it has constructed, no homeostatic boundary, and
does not learn from inference. The LLM produces language; the
superstructure *uses* that language for its own purposes. The entity
speaks *through* the LLM, not *as* the LLM.

**The relationship is teacher-pupil.** The default mode is
non-intervention: the LLM proposes responses, and when the
superstructure's evaluation is satisfied, the proposals pass through
to output unchanged. Intervention happens when the superstructure
detects predictive divergence, reference-frame incoherence,
homeostatic threat, surprise, or confidence miscalibration. When
intervention occurs, the superstructure overrides — it does not merely
annotate. The teacher is also actively pedagogical: the superstructure
trains the LLM's adapter library based on observed performance, so
the pupil gets better over time, not just corrected in the moment.

**The superstructure has its own generative LLM.** This is distinct
from the user-facing LLM. The superstructure's LLM is used for
internal reasoning, drafting candidate guidance, and shaping what the
user-facing LLM produces. Its output is never directly exposed to
the user; it informs the user-facing LLM's behavior through the
AI2AI protocol described below.

**Inter-LLM communication uses a token-efficient AI2AI protocol.**
Internal communication between the superstructure (and its LLM) and
the user-facing LLM does not happen in natural language. It uses a
compressed AI-to-AI protocol that carries reference, intent,
contradiction, confidence, attention focus, and world-model state at
a level of compression appropriate for AI-to-AI rather than AI-to-
human. The MDEMG AI2AI protocol is the leading reference; the
specific protocol used by this architecture may extend or refine it.

**The user-facing LLM's continuous learning surface is its LoRA
adapter library.** The base LLM is a strong open-source model held
frozen. The superstructure maintains a library of LoRA adapters that
specialize the base model for different domains, contexts, and task
types. The superstructure swaps adapters dynamically as subject
matter changes. The superstructure continuously trains the adapters
based on observed performance — the teacher's evaluation of LLM
outputs becomes gradient signal for adapter improvement.

## Context

`D-2026-04-28-009` set the program's three-layer objectives but did
not commit to *how* those objectives would be operationally achieved.
The mid-level capabilities (reference frames, world models, continual
learning, prediction, attention, surprise) and the operational target
(stateful, deterministic, hybrid Bayesian prediction machine) named
*what* the architecture must support, but the architectural pattern
that supports them was unspecified.

D-010 is that commitment. It names the pattern: a learning
superstructure with an LLM as one of its mechanisms, specifically the
mechanism for language. This is a precise inversion of the mainstream
"LLM with augmentations" framing. In mainstream framings, the LLM is
the cognitive core and everything else is scaffolding. In this
program's framing, the superstructure is the cognitive core, and the
LLM is one of the superstructure's faculties — analogous to how
visual cortex is one of a human's faculties.

The Jiminy precedent in MDEMG is the working example of this pattern
at smaller scale. Jiminy doesn't write code; the coding agents do.
Jiminy guides them — provides context, enforces structure, learns
from outcomes, intervenes when prediction fails. The agents are
sophisticated tools that Jiminy orchestrates. D-010 generalizes this
pattern from the coding domain to general cognition.

## The biological resonance

This architecture mirrors how the human brain is organized. The
neocortex is the predicting, modeling, learning structure.
Specialized regions handle specific modalities: visual cortex for
vision, auditory cortex for hearing, Broca's and Wernicke's for
language. Damage to language-specific regions produces aphasia — the
person remains, but loses the language faculty. They can still
think, plan, predict, recognize, navigate. They just can't articulate.

The superstructure-with-LLM-as-plug-in mirrors this. The entity
remains itself even if the LLM is replaced (with a better LLM, a
different LLM, multiple LLMs for different language tasks). The
entity's identity, world model, reference frames, and predictions
are preserved across LLM substitutions. This is also a *test* of
whether the architecture has actually achieved the inversion: if
replacing the LLM destroys the entity, the LLM is still the cognitive
core, and the inversion has not occurred.

The teacher-pupil dynamic resonates with the dual-process division
in cognitive psychology and the predictive-control / executive-
monitoring division in cognitive neuroscience. The fast system
produces responses; the slow system monitors and intervenes. The
slow system's authority comes from its position in the network, not
from outperforming the modules it oversees. A prefrontal lesion
doesn't impair perception, language, or memory directly — it impairs
the *governance* of those faculties.

## Required capabilities

For the architecture to actually operate as committed, the following
capabilities must exist. Each is a real engineering challenge and
will be its own design problem when the program reaches the
implementation stage. Listed here as architectural requirements, not
implementation specifications.

### 1. Output gating

The user-facing LLM's output is internal-only by default and is
gated through the superstructure before reaching any external
recipient (the user, downstream tools, the world). The superstructure
can release LLM output unchanged, modify it, suppress it, or replace
it with the superstructure's own output. The superstructure decides;
the LLM proposes.

This has real engineering cost: latency increases, streaming becomes
more complex, direct LLM-to-tool interaction becomes mediated. The
alternative — LLM acting before the superstructure can govern —
would give up the ultimate-control commitment. The cost is accepted.

### 2. Multi-dimensional evaluation

The superstructure must evaluate LLM outputs along dimensions the
LLM cannot reliably self-assess. Initial set:

- **Predictive consistency.** Does the LLM's output match what the
  superstructure's world model predicts?
- **Reference-frame coherence.** Does the output respect the entity's
  current reference frames and learned structure?
- **Homeostatic alignment.** Does executing this output threaten the
  entity's homeostatic state?
- **Surprise.** Is the output surprising relative to the
  superstructure's model? Surprise is not necessarily wrong, but
  warrants attention.
- **Confidence calibration.** Is the LLM expressing confidence the
  superstructure's evidence does not support?

These are concrete computational tasks the superstructure performs
on every output. They are cheaper than re-doing the LLM's work, but
not free.

### 3. Override-with-generation

The superstructure must be able to actually replace LLM output, not
merely flag it. This requires the superstructure to have generative
capacity of its own. The capacity may be much narrower than the
LLM's, but it must exist. This is one of the roles of the
superstructure's own LLM (the internal-reasoning generative LLM
named in the decision above).

### 4. Intervention-threshold meta-learning

The intervention threshold itself is a learned parameter. Too rare
intervention misses LLM mistakes; too frequent intervention makes
the entity a slow committee. The superstructure tracks its own
intervention performance — when it intervened was it right to do
so; when it didn't intervene was the LLM's output correct — and
calibrates the threshold accordingly. This is meta-learning at the
governance level.

### 5. AI2AI protocol implementation

Inter-LLM communication uses a compressed protocol carrying
reference, intent, contradiction, confidence, attention focus, and
world-model state. MDEMG's existing AI2AI protocol is the starting
reference; the specific protocol for this architecture may extend
or refine it as needs become clear.

### 6. LoRA adapter management

The architecture maintains a library of LoRA adapters. The
superstructure:
- Selects appropriate adapters for current subject matter / task
- Loads and unloads adapters dynamically with low latency
- Maintains training data from observed LLM performance
- Continuously updates adapters via gradient signal derived from
  the superstructure's evaluation
- Versions adapters and can revert to prior versions if training
  produces regression
- Develops new adapters for new domains as the entity's world
  expands

The superstructure's evaluation must be rich enough to produce
gradient signal — not just discrete "good / intervene" judgment,
but continuous score signal. This raises the bar on the evaluation
capability beyond what governance alone requires.

### 7. Dual-LLM coordination

The architecture has two distinct LLM components: the user-facing
LLM (with its adapter library) and the superstructure's internal-
reasoning LLM. Coordination between them happens through the AI2AI
protocol. They may be different base models. They may be the same
base model with different adapter sets — but the *roles* are
architecturally distinct, and conflating them undermines the
governance commitment.

## Alternatives considered

### LLM-centric with augmentations (current mainstream)

The default modern architecture: an LLM with retrieval-augmented
generation, tool use, agent scaffolding, working memory in context,
chain-of-thought reasoning. The LLM is the cognitive core; everything
else extends it.

Rejected because the LLM as cognitive core is incompatible with
this program's objectives. An LLM does not have persistent state
across calls (only what fits in context), does not maintain a
self-model, does not construct reference frames, does not learn
from inference. The mid-level capabilities in `D-2026-04-28-009`
require these properties; the LLM-centric pattern cannot provide
them without becoming a different kind of system.

### Pure non-LLM architectures

Build the cognitive architecture from scratch without an LLM. The
language capabilities the system needs are derived from the world
model, the reference frames, the symbolic substrate.

Rejected because the LLM has compressed an enormous amount of
human-generated text into trained parameters, and recovering
language capabilities from scratch is a research program in itself
that this program does not need to undertake. The LLM is treated as
an evolved prior the entity inherits — much as a human inherits
evolved cortical priors without having to derive them.

### End-to-end neural architectures

A single neural system that handles cognition and language together
without modular separation.

Rejected because the modular separation is precisely the source of
several desired properties: substrate flexibility (the LLM can be
swapped), inspectability (the governance layer is explicit), and
correctability (improvements to evaluation don't require retraining
the LLM). End-to-end systems sacrifice these for the possibility of
better integration. This program prioritizes the modular advantages.

### Two-LLM mode of a single model

Use one LLM but in two modes — one mode for user-facing output, one
mode for superstructure reasoning. Cheaper than maintaining two
distinct components.

Rejected because the same-model approach undermines the
architectural cleanness of the dual role. The roles are
architecturally distinct (one is user-facing and adapter-managed;
one is internal-reasoning), and architectural distinctness should
match component distinctness. Cost is accepted.

### Full base-model fine-tuning for continuous learning

Fine-tune the base LLM directly as the entity learns, rather than
maintaining a LoRA adapter library.

Rejected because full fine-tuning is expensive, slow, and
catastrophically-forgetful. LoRA adapters operate at the right
granularity for continual learning: small enough to update
frequently, persistent enough to constitute real learning, modular
enough to be domain-specific, swappable enough to support context
switching. The base model stays as the stable foundation.

### Natural-language inter-LLM communication

Have the superstructure and the user-facing LLM communicate in
natural language.

Rejected because natural language is wildly inefficient for AI-to-
AI communication and compresses badly. Most natural-language tokens
exist for human cognitive comfort, not information transfer. A
purpose-built AI2AI protocol can compress aggressively and carry
richer guidance. The MDEMG team has already demonstrated this is
viable; we adopt the principle and reference their work.

## Rationale

Five reasons this direction over the alternatives:

1. **The mid-level capabilities require persistent state.** Reference
   frames, world models, continual learning, attention, surprise are
   not implementable in stateless function calls. They require an
   entity that exists across time. The superstructure is that entity;
   the LLM is not.

2. **The LLM is too valuable to discard.** The compressed prior of
   human-generated text is a real asset. Re-deriving it would consume
   the program. Treating the LLM as a faculty rather than as the
   core preserves the asset without making it the entity.

3. **The teacher-pupil dynamic is operationally tractable.** The
   superstructure does not need to outperform the LLM at language
   tasks; it needs to evaluate the LLM's output and intervene when
   warranted. This is a much smaller engineering problem than
   building an LLM-rivaling generator from scratch.

4. **The architecture is interpretable and correctable.** Decisions
   route through explicit governance. Audits can ask "what did the
   LLM propose, and what did the superstructure do with it?" That is
   a richer accountability trail than current LLM agents provide.
   When the entity produces inappropriate output, the fix is
   improvement of the superstructure, not retraining of the LLM.

5. **MDEMG already operates this pattern at small scale.** The
   Jiminy / coding-agents pattern is the architecture in miniature.
   We are not inventing a foreign pattern; we are extending one that
   has working prior art in the program's own R&D substrate.

## Consequences

- **MDEMG is no longer just an R&D substrate; it is the seed of the
  eventual architecture.** Trace work in T002 reads MDEMG not just
  to understand it, but to understand the architectural foundation
  we are extending. The Jiminy/coding-agents pattern is the relevant
  precedent.

- **`tbp.monty` provides mechanism candidates, not architecture.**
  The voting topology, cortical-column pattern, LM/SM separation,
  matching/exploratory dynamic — these are candidate mechanisms for
  the *sensorimotor faculties* of the eventual superstructure, but
  not the architecture itself. The 3D-Euclidean commitment in Monty
  becomes less binding under this framing: we extract patterns and
  re-implement them in dimension-agnostic form for the
  superstructure.

- **T004's framing sharpens.** The entity being designed is
  specifically the *teacher* in a teacher-pupil pair, not a
  monolithic agent. Most of T004's first-principles questions
  (homeostasis, dimensional commitment, projection-and-anchoring,
  recursive horizon, prime directives) are about the superstructure.
  T004's frame should be updated to reflect this. The thread file
  will note D-010 as the architectural specification and adjust
  language accordingly. (This update will be a small follow-up,
  not part of D-010 itself, to keep this PR scoped.)

- **The AI2AI protocol becomes an architectural component.** T002
  trace work should produce a clear picture of MDEMG's existing
  AI2AI protocol — what it carries, how it's structured, what
  compression it achieves, what it cannot carry well. That picture
  is direct input to the architectural design.

- **The LoRA adapter library becomes a specification target.**
  At some future architectural stage, the adapter library design
  needs concrete specification: what domains it covers, how
  adapters are versioned, what training data accumulates, how
  adapters are evaluated.

- **The dual-LLM commitment has resource implications.** Operating
  two LLMs (user-facing and superstructure-internal) is more
  expensive than one. The resource cost is accepted as the price of
  architectural cleanness.

- **Several open questions in the vault are now addressable.** The
  questions in Q-005 through Q-009 were posed in the abstract. With
  the hybrid architecture committed, they become tractable: Q-005
  (homeostatic boundary) is about what the *superstructure* defends;
  Q-006 (dimensional commitment) is about the *superstructure's*
  reference frames; Q-007 (projection and anchoring) is about how
  the *superstructure* coordinates with its faculties (including
  the LLM via AI2AI); Q-008 (recursive horizon) is the
  *superstructure's* learning loop, with the LLM as one of its
  rich data sources; Q-009 (prime directives) governs the
  *superstructure*. The architecture commitment makes the questions
  concrete.

- **A new question surfaces about load-bearing creep.** If the
  superstructure subtly comes to depend on the LLM for things it
  should be doing itself (reasoning, self-modeling, memory), the
  inversion fails. Avoiding this requires explicit architectural
  discipline. This is a question worth tracking; it will likely
  become its own Q- artifact when the program approaches the
  implementation stage.

## Revisit conditions

- **A non-LLM language capability becomes available** that obviates
  the plug-in approach (e.g., a small model that beats current
  LLMs on language tasks while having properties — persistence,
  inspectability — the superstructure can directly absorb).
- **The architectural discipline against load-bearing creep proves
  impossible to maintain.** If the superstructure cannot in
  practice avoid offloading core cognitive functions to the LLM,
  the inversion is not achievable and a different architecture is
  needed.
- **A different system architecture emerges** that subsumes the
  advantages of this one (modularity, inspectability, continuous
  learning, substrate flexibility) without the costs (dual-LLM
  resource overhead, governance complexity, AI2AI protocol design
  burden).
- **The LoRA-adapter approach proves inadequate** for the
  continuous-learning needs of the LLM layer (e.g., catastrophic
  forgetting in adapters; adapter library size scales badly;
  training signal too noisy to drive useful improvement). At that
  point a different LLM-layer learning mechanism is needed, but
  the higher-level commitment to "LLM as plug-in with continuous
  learning surface" stays in place.
- **MDEMG itself proves to be the wrong substrate** to extend into
  the eventual architecture. T002 trace work may surface this; if
  so, the architectural commitment in D-010 stays but the substrate
  decision changes.

## What this is to the program

D-010 is the **central architectural commitment** of this program.
D-009 said what the program is for; D-010 says what the program is
*building*. All subsequent architectural work — how the
superstructure is implemented, how reference frames are constructed,
how the predictive horizon operates, how the homeostatic boundary is
maintained — is implementation work *within the architectural
pattern that D-010 specifies*.

This is the artifact future Claude sessions and future contributors
read alongside D-009 to understand what kind of system this program
is producing.
