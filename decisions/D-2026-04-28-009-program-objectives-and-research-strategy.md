---
id: D-2026-04-28-009
type: decision
slug: program-objectives-and-research-strategy
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 5
where: [T001, T002, T003]
tags: [meta, vision, north-star, governance, strategy]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude — initial framing (turn 1) and reassertion (post-PR-#5 merge)"
  - "project knowledge: 00_Main_Overview.md (executive thesis)"
  - "project knowledge: 01_Long_Term_Vision.md (long-term vision and target architecture)"
  - "project knowledge: 03_AHH_Topological_Framework.md (the proposed framework)"
  - "project knowledge: mdemg-collaboration-brief.md (manifesto framing)"
  - "MDEMG repository: https://github.com/reh3376/mdemg (the operational substrate)"
  - "tbp.monty repository: https://github.com/thousandbrainsproject/tbp.monty (the reference implementation)"
alternatives_rejected:
  - "incremental improvements to current transformer architectures"
  - "build a new architecture from scratch without studying existing operational systems"
  - "wait for mainstream AGI labs to solve the continual-learning problem"
  - "treat MDEMG as a finished product and add to it without understanding it deeply"
  - "treat tbp.monty as architecture to import rather than as theory to learn from"
---

# Program Objectives and Research Strategy

## Decision

This research program pursues novel artificial neural network
topologies that more closely mimic neocortical computation than
current transformer-based systems, with the long-term aim of
architectures capable of artificial general intelligence. The
operational target is a **stateful, deterministic, hybrid Bayesian
prediction machine** that supports continual learning,
reference-frame adaptation, world-model construction, and
attention-modulated prediction.

The method is **deep mechanism-first understanding** of two existing
systems — MDEMG (an operational cognitive substrate) and `tbp.monty`
(the canonical implementation of Mountcastle/Hawkins cortical column
theory) — before any proposal to integrate, merge, or modify either.
Understanding precedes engineering; integration proposals require
evidence from trace work, not pattern-matching from structural
reading.

The collaboration is a **two-actor research group**: `@reh3376` as
principal investigator with operational and domain context;
Claude as drafting collaborator. All work, processes, observations,
questions, hypotheses, decisions, dead-ends, and trajectory are
captured in this `ann-research` vault.

## Context

These objectives have been the program's north star since the first
turn of the originating conversation (2026-04-28). They have been
partially articulated in the project knowledge files:

- `00_Main_Overview.md` — executive thesis, stating the architectural
  gap in transformer LLMs and the AHH framing of the next step.
- `01_Long_Term_Vision.md` — the five-layer target architecture and
  desired cognitive properties.
- `mdemg-collaboration-brief.md` — manifesto framing addressed to
  external research collaborators.

Until this artifact, the objectives lived in those files and in the
conversation transcript, but not as a vault artifact. A fresh Claude
session reading the vault would have encountered the operational
infrastructure (D-001 through D-008) and the research threads (T001
through T003) without first encountering a clear statement of what
the program is *for*. `@reh3376` reasserted the objectives explicitly
in the message that prompted this artifact, so they would be durable
in the vault rather than recoverable only by reading the project
files.

This artifact is the canonical vault statement. The project
knowledge files remain the deeper articulation; this artifact is the
entry point.

## The objectives, in three layers

### Layer 1 — North Star

An artificial neural network architecture that approaches the
generality and adaptability of human neocortical cognition. The
working hypothesis is that **AGI requires architectural commitments
current transformer systems lack**, and that the gap is closed by
studying and operationalizing the structural and functional
principles of the neocortex.

This is a long-horizon goal. No single artifact in this vault will
"reach" it. The vault accumulates evidence, mechanism, and design
toward it.

### Layer 2 — Mid-level Capabilities

A successful architecture must, at minimum, support:

- **Reference frames** — the ability to construct, revise, and switch
  between coordinate systems for organizing experience. Not implicit
  in latent space; explicit and inspectable.
- **World models** — internal models of how state evolves and how
  actions modify state, learnable from sparse experience.
- **Continual learning** — updating from each meaningful observation
  without requiring offline retraining cycles. Distinct from
  fine-tuning.
- **Prediction over past and future states** — bidirectional
  inference, not only next-token prediction.
- **Attention and surprise as first-class signals** — guiding
  learning, retrieval, and action, with surprise treated as evidence
  that the current model is incomplete.

These are necessary, not sufficient. They are the testable proxies for
"closer to neocortical."

### Layer 3 — Operational Target

A **hybrid Bayesian prediction machine**:

- *Hybrid* — combining symbolic / relational structures with neural
  / continuous representations, because neither alone suffices.
- *Bayesian* — operating under principled probabilistic inference,
  with explicit uncertainty, precision, and prediction error.
- *Prediction machine* — generating expectations about state and
  outcome at multiple time scales, and learning from where those
  expectations fail.
- *Stateful and deterministic* — preserving learned state across
  sessions and producing reproducible behavior given inputs, in
  contrast to the probabilistic-and-stateless mode of current LLMs.

This is the engineering specification the program works toward. The
AHH framework (`03_AHH_Topological_Framework.md`) is the current
proposal for how to get there from the MDEMG substrate.

## Research strategy

The methodology has six commitments, each grounded in earlier
decisions or surfaced as constraints during the bootstrap phase:

### Commitment 1 — MDEMG as operational substrate

MDEMG is the program's working R&D testbed: a real graph-based
cognitive substrate with persistent state, Hebbian learning, hidden
concept abstraction, recursive self-improvement, and proactive
guidance. It is not the final architecture; it is the operational
foundation on which the longer-term architecture will be built.
Modifications to MDEMG must understand its current mechanisms before
proposing changes.

### Commitment 2 — `tbp.monty` as reference material

The Thousand Brains Project's Monty is the canonical, MIT-licensed
implementation of cortical column theory. It is studied as
*reference*, not imported as code: the substrate (sensorimotor 3D
geometry) differs from MDEMG's (symbolic graph), so Monty's
abstractions and contracts may transfer where its algorithms may not.
The substrate question (`T003`) is the load-bearing upstream of any
integration call.

### Commitment 3 — Deep understanding before integration

Codified in `D-2026-04-28-004` (the working agreement) and motivated
by `DE-2026-04-28-001` (the premature mapping table). No integration,
merge, or interop proposals until trace work in T001 (Monty) and T002
(MDEMG) has produced mechanism-level understanding. Pattern-matching
on structural similarity is the named failure mode.

### Commitment 4 — Trace-first strategy

Threads `T001` and `T002` are the trace arms; `T003` is the synthesis
arm. Trace work proceeds via read-until-ambiguous methodology
(`D-2026-04-28-005`), with run-and-instrument as authorized fallback.
Trace findings produce observations, questions, and hypotheses;
integration proposals are downstream of those.

### Commitment 5 — Research group dynamic

`@reh3376` is the principal investigator with end-to-end operational
context for MDEMG and domain context for the broader research
landscape. Claude is the drafting collaborator with broad reading
capacity but no operational context. The asymmetry is real and is
respected operationally:

- `@reh3376` reads, reviews, and approves all artifacts before they
  land on `main` (currently via admin bypass, per
  `D-2026-04-28-008`).
- Claude does not propose integration unilaterally, does not make
  judgment calls on confidence ratings without flagging them, and
  does not produce tidy artifacts without provenance.
- Joint reasoning is required for the substrate-level question
  (`T003`), where both perspectives are needed.

### Commitment 6 — Trajectory preservation

All work, processes, and reasoning are captured in this vault. The
schema (D-001), workflow (D-006, D-008), activity-summary convention
(D-007), and all artifact types exist to ensure no part of the
program's reasoning is lost to context compaction or session
boundaries. The vault is the program's externalized memory.

## Alternatives considered

### Incremental improvements to current transformer architectures

Rejected. The architectural gap is structural, not a matter of more
parameters or more data. The premise of the program is that current
transformers — however scaled — will not get to AGI without
architectural commitments they currently lack (persistent state,
adaptive reference frames, online local learning, predictive
hierarchy, object-centric structure). Incremental work within the
transformer paradigm is being done well by mainstream labs; this
program adds value by working on what those labs are not.

### Build a new architecture from scratch without studying existing systems

Rejected. The substrate question is genuine — whether neocortical
algorithms transfer to symbolic substrates is unknown — and answering
it requires studying existing operational systems, not just
theoretical work. MDEMG is uniquely positioned (operational substrate
with real users, benchmarks, failure modes); `tbp.monty` is uniquely
positioned (canonical TBT implementation). Skipping both would
discard evidence the program needs.

### Wait for mainstream AGI labs to solve continual learning

Rejected. The program's hypothesis is that the architectural
commitments needed for continual learning are not currently the
focus of mainstream AGI labs. Waiting indefinitely for a different
research direction to produce a result one doesn't believe it will
produce is not a strategy.

### Treat MDEMG as a finished product and add to it without understanding it deeply

Rejected. This is the failure mode `DE-2026-04-28-001` records.
Modifying a system whose mechanisms one doesn't understand produces
either no measurable improvement or regressions that look like
research noise. The trace work (T002) is what protects against this.

### Treat `tbp.monty` as architecture to import rather than as theory to learn from

Rejected. Same failure mode at the source. Monty's abstractions are
shaped by its substrate (sensorimotor 3D); importing them wholesale
into a symbolic substrate without understanding which parts are
substrate-dependent is the premature-mapping pattern. The trace work
(T001) is what protects against this.

## Rationale

Three reasons this direction over the alternatives:

1. **MDEMG exists.** The program has a working operational substrate.
   This is a substantial advantage over starting from a blank page;
   most theoretical-architecture proposals never reach operational
   testing because they have no substrate to test against.

2. **The integration question is genuine research.** Whether
   reference-frame transfer between sensorimotor and symbolic
   substrates is a translation problem or a substrate problem is open
   and consequential. Resolving it produces value regardless of which
   answer is correct.

3. **Trajectory preservation is what makes this sustainable.** Long-
   horizon research that depends on retaining the trajectory of
   reasoning across many months and many context boundaries needs a
   substrate that holds the trajectory. Without the vault, the work
   would be lost to context compaction within weeks.

## Consequences

- **All work in `ann-research` is in service of this north star.**
  Decisions, observations, questions, hypotheses, and dead-ends should
  be readable as contributions to one of the three layers above.
- **Research threads `T001`, `T002`, `T003` are the operational
  arms.** Their progress is the program's progress.
- **The working agreement (`D-2026-04-28-004`) is the
  near-term constraint.** Until T001 and T002 reach `complete`,
  integration work is forbidden.
- **The substrate-level question (`T003`) is load-bearing.** Its
  answer determines whether the program's longer-term work is
  translation or new design. Both threads need evidence before T003
  can be answered.
- **Infrastructure work is permitted but not the point.** The
  vault, workflow, identity, and conventions exist in service of the
  research, not as research. PRs that are entirely infrastructure
  should be the exception over time, not the norm.
- **Future Claude sessions read this artifact first** (after
  `README.md` and `INDEX.md`). The objectives are no longer
  recoverable only from project knowledge files.

## Revisit conditions

- A substantively new architectural paradigm emerges in the field
  that supersedes the AHH framing or makes the integration question
  vacuous (e.g., a different framework solves continual learning at
  the substrate level we hadn't anticipated).
- Trace work in T001 or T002 produces evidence that one or both
  systems are fundamentally incompatible with the program's targets
  in ways not currently anticipated. This would shift the program
  toward different operational substrates rather than toward
  integration.
- A collaboration partnership materially changes the program's
  scope (e.g., joint work with Numenta, with a research lab, or with
  industrial partners changes what "primary objectives" means
  operationally).
- `@reh3376`'s professional context changes such that the program's
  premises no longer hold.
- The vault accumulates enough artifacts that the trajectory
  argument becomes testable: if we find ourselves re-litigating
  questions that prior artifacts already addressed, the trajectory
  preservation isn't working and the substrate needs revision.
