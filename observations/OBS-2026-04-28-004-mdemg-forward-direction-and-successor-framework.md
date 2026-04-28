---
id: OBS-2026-04-28-004
type: observation
slug: mdemg-forward-direction-and-successor-framework
created: 2026-04-28
updated: 2026-04-28
status: recorded
confidence: 4
where: [T002, T004, T005]
tags: [mdemg, sprint-plans, successor-framework, predictive-coding, fep, htm, ff, fork-timing, alignment-assessment]
provenance:
  - "@reh3376 supplied the docs/research/mdemg_sprint_ideas/ directory contents on 2026-04-28 (uploaded as 13 markdown files; the directory exists locally but is not present on the public mdemg repo at the time of read)"
  - "mdemg-specification.md (155K, 1328 lines) — MDEMG-SPEC-RD-001 v1.0 dated 2026-04-27; 'R&D Phase Stewardship and Successor Framework Preparation'; written by 'technical reviewer working with reh3376'; calibration explicitly 'loyal opposition'"
  - "Nine sprint plans dated 2026-04-21 covering: 01-pc-reframe, 02-precision-weighted-eta, 03-top-down-predictions, 04-column-voting, 05-context-fingerprints, 06-sparse-retrieval, 07-ff-shallow-heads, 08-htm-sequence-memory, 09-active-inference-unification"
  - "risk-opp-04232026-01.md (24K) — MDEMG-SPR-PLAN-001 v1.0 dated 2026-04-24; risk and opportunity remediation sprint planning input"
  - "MDEMG_FT_LORA_PACKAGING_SPEC.md (17K) — MDEMG-FT-PKG-001 v1.0 dated 2026-04-24; LoRA packaging and Homebrew distribution"
  - "lora-training-glossary.md (54K) dated 2026-04-27 — LoRA fine-tuning vocabulary reference"
  - "OBS-2026-04-28-003 (companion observation: MDEMG current-state subsystem orientation)"
  - "All decisions D-009 through D-011 and questions Q-005 through Q-010 (this artifact bears on most of them substantively)"
subject: "MDEMG's forward-looking architectural research corpus reveals that MDEMG itself is explicitly framed as 'R&D vehicle, not production target' with a 'successor framework on the horizon' that is a separate artifact, not a future MDEMG version. This artifact captures (a) the strategic frame (MDEMG-as-R&D-vehicle, fork-timing constraints, alignment assessment against long-term goal), (b) the nine-sprint research corpus (PC/FEP track + independent tracks + active-inference capstone), (c) the seven carry-forward patterns identified for the successor, and (d) substantial implications for our hybrid-superstructure architectural commitments."
---

# MDEMG Forward Direction and Successor Framework

## What this artifact is

OBS-2026-04-28-003 captured **what MDEMG is today** — seven major
subsystems, their relationships, current operational state. This
artifact captures **what MDEMG plans to become next**, organized
around the explicit framing the source documents establish: MDEMG
as a finite-duration R&D vehicle producing architectural
intuitions for a successor framework that is a separate artifact.

The two observations are companions. OBS-003 is the orientation of
the current substrate; OBS-004 is the orientation of the forward
direction. Together they form Cycle 1's complete orientation
deliverable.

## The strategic frame

`mdemg-specification.md` (1328 lines, dated 2026-04-27) is the
master document. Five claims from §0.2 are load-bearing for
everything that follows:

1. **MDEMG is the R&D vehicle, not the production target.** The
   user has been explicit: MDEMG has served as a year-long R&D
   substrate, has produced architectural intuitions that are now
   coherent, and is the *template* from which a more abstracted
   successor framework will be built.

2. **The successor framework is a separate artifact, not a future
   MDEMG version.** This is a structural commitment, not a naming
   choice. The successor is built on a different foundation; MDEMG
   is the template for patterns, not the structural starting point.

3. **The long-term goal is not "improve MDEMG."** The stated
   long-term goal is *continuous-learning artificial neural
   networks with biologically-inspired topologies — systems that
   establish reference frames, build world models, continuously
   predict current state, project forward and backward in time,
   and improve themselves recursively such that they become
   smarter and wiser the way a biological neural network does*.
   The BNN analogy is "vocabulary stand-in"; the actual artifact
   may not resemble a BNN, but no better vocabulary currently
   exists.

4. **Fork-timing is constrained.** The right moment to fork is not
   before FT-LORA is fully complete — all adapters trained,
   multi-LoRA serving validated end-to-end at the 16 MDEMG call
   sites. Forking earlier produces a worse template; forking later
   risks design-choice calcification.

5. **Calibration is loyal opposition.** Honest critique. Placation
   explicitly disallowed. The document's role is to surface
   observations sharply so decisions are well-informed.

This frame **substantively re-orients our program's relationship to
MDEMG**. Earlier vault artifacts (D-009, D-010, OBS-003) treated
MDEMG as "an R&D substrate" without naming the fork. The fork is
real, planned, and gated. The successor framework — what
@reh3376's documents call it — is approximately what D-010 calls
the *hybrid superstructure*. This is the deepest cross-reference in
the program to date.

## The architectural ceiling (§2 alignment assessment)

The document's §2 is "the most consequential" by its own framing.
It evaluates how MDEMG's architecture aligns with the long-term goal,
calibrated to loyal opposition. The findings are sharp:

### Three boundaries of MDEMG's learning (§2.1)

MDEMG learns edge weights over a graph whose embeddings come from
external models. The boundaries:

- **No representation learning.** The embedding manifold is fixed
  by the external embedding model. MDEMG's learning operates over
  coordinates the system did not choose.
- **No structural learning.** The graph schema (node types, edge
  types, layer structure) is fixed by migrations.
- **No mechanism learning.** The Hebbian rule is fixed. RSIC
  tunes the rule's *parameters* (eta, decay, thresholds); it
  does not modify the rule itself.

The +58.4% retrieval improvement is real, but it is what
"retrieval-over-frozen-representations can produce when applied
with discipline." It is not evidence the substrate can support
continuous-learning ANNs.

### Vocabulary drift risks (§2.2-§2.4)

The document explicitly flags three places where MDEMG vocabulary
overlaps with the long-term-goal vocabulary in ways that risk
conceptual confusion:

- **"Emergence."** MDEMG's emergence is DBSCAN clustering plus
  LLM-driven naming within a fixed embedding space. This is not
  emergence in the manifesto's sense (where new representational
  units form and the system carves new conceptual axes that did
  not exist before).
- **"Reference frames."** MDEMG's reference frames are inherited
  from external structure (file paths, package hierarchies, symbol
  trees). They are static external scaffolding. Hawkins' Thousand
  Brains framing — distributed reference frames the system
  *constructs* to support compositional reasoning, learned from
  experience, modifiable as understanding evolves — is materially
  different.
- **"Self-improvement."** RSIC's "Recursive Self-Improvement Cycle"
  modifies parameters within fixed mechanisms. The 19 reflection
  patterns, the action types, the safety gates — all pre-defined.
  RSIC is "calibrated parameter search inside a designed envelope."
  The long-term goal's recursive self-improvement is more
  ambitious — modify learning rules, restructure architecture,
  refine optimization targets — and is not what RSIC attempts.

The document is explicit: "The fact that RSIC is called 'Recursive
Self-Improvement Cycle' risks the same vocabulary drift as
'reference frames' — useful conceptual framing for the work that
*is* done, misleading if read as a claim about the work that is
*not* done."

**This bears directly on our vault's vocabulary.** D-009 names
"recursive self-improvement" as a mid-level capability target.
Q-008 frames "recursive expansion of predictive horizon." We
should be careful that those framings remain claims about the
*entity we're designing*, not descriptions of MDEMG's RSIC.

### What is *not* in the code (§1.10)

Direct code search confirms absence of:

- No `predictive_coding`, `prediction_error`, or `PredictiveCoding`
  symbols anywhere in Go or Python
- No `world_model` or `WorldModel` symbols
- No `reference_frame` or `ReferenceFrame` symbols
- No `active_inference` or `ActiveInference` symbols
- No `free_energy` symbols
- No object-centric slot structure (no slot-attention, no
  DETR-style query slots, no equivalent)

The "predict" usages that *do* exist in the Python sidecar
(`tier_model.py`, `app.py`) are application-level prediction (NLI
tier classification, response generation), not architectural
prediction. They predict what an agent will comprehend, not what
the world will do next.

### Four substrate-level commitments demanded by the long-term goal that MDEMG cannot supply by design (§2.7)

This is the central architectural finding:

1. **Trainable representations.** Substrate must support
   representation learning. MDEMG's embedding layer is external
   and frozen.
2. **Generative modeling / prediction.** Substrate must produce
   predictions about current state and future state. MDEMG's
   graph stores observations; nothing produces predictions over
   them.
3. **Reference-frame construction.** Substrate must let the system
   carve representational space into frames that support
   compositional reasoning. MDEMG inherits frames from external
   structure.
4. **Meta-learning / mechanism modification.** Substrate must let
   the system modify its own learning rules over time. MDEMG has
   fixed mechanisms with parameter-only modification.

All four must be substrate-level commitments in the successor
framework. The MDEMG R&D phase should produce *empirical input* on
how each commitment should be implemented, "drawn from the
architectural intuitions generated by working without them."

**This is the sharpest possible statement of the gap our program
has been working in.** D-009's mid-level capability commitments
(reference frames, world models, continual learning, prediction)
are exactly these four. D-010's hybrid superstructure architecture
is the program's commitment to a substrate that supplies them.

### Three options for reaching the long-term goal (§2.6)

The architectural ceiling is not a failure of effort; it is a
property of the architectural shape. Reaching the long-term goal
from MDEMG's shape would require one of:

1. **Replace the frozen-representation foundation** with a
   learned-representation foundation (the embedding layer becomes
   trainable, the manifold becomes continuously updated), and
   rebuild everything that keys on it.
2. **Add new substrate alongside the existing one** (a
   generative-model layer, a representation-learning layer, a
   meta-learning layer) and re-architect how the substrates compose.
3. **Build the successor framework on a different foundation
   entirely**, with MDEMG as a template for the patterns that work
   but not as the structural starting point.

@reh3376's stated framing is **option 3**. The document operates
within that framing. Options 1 and 2 are not zero-cost; option 1
("rebuild everything that keys on the embedding layer") would
require touching every subsystem in the codebase.

**Our D-010 hybrid superstructure aligns most closely with option
3.** The superstructure is built on a different foundation (its
own predictive substrate, its own world model, its own homeostatic
boundary); MDEMG patterns are inherited as orchestration and
governance machinery, not as the cognitive substrate itself.

## Seven carry-forward patterns for the successor (§2.8)

The document names seven MDEMG patterns that should carry forward
to the successor as substrate-agnostic commitments. Each is named
explicitly:

1. **The interceptor pattern.** Five-component closed feedback
   loop around action: Guide → Validate → Evaluate → Track →
   Learn. Substrate-agnostic. MDEMG's Jiminy is a reference
   implementation of "Guide"; NLI is a reference implementation of
   "Track."

2. **The multi-temporal RSIC scaffolding.** Five-stage cycle
   (Assess → Reflect → Plan → Execute → Validate) at three
   temporal scales (Micro / Meso / Macro). Self-improvement at
   multiple temporal scales is generalizable; the specific 19
   reflection patterns are MDEMG-specific.

3. **The UxTS framework family pattern.** Schema + spec + runner
   + CI as a uniform pattern for governing testable surfaces.
   Possibly MDEMG's strongest carry-forward. The pattern produces
   "empirically reliable governance across heterogeneous testable
   surfaces."

4. **The constraint-and-evidence schema.** Typed constraints with
   confidence, contradictions, frontiers, evidence counts. The
   conceptual layout is substrate-agnostic; the specific Cypher
   schema is MDEMG-specific.

5. **The multi-LoRA serving infrastructure.** One base model, many
   adapters, programmatic per-request adapter selection. The
   pattern is the production-ready realization of multi-task
   specialization; the specific adapter format and base model are
   MDEMG-specific.

6. **The fail-open default for LLM-backed capabilities.** Every
   LLM-dependent capability in MDEMG degrades gracefully when the
   provider is unavailable. A generalizable resilience commitment.

7. **The dual-gate regression discipline.** 5a aggregate gate +
   5b per-task gate. Catches the failure mode where aggregate
   metrics mask catastrophic per-task regressions. Generalizable
   testing pattern for ML system updates.

**This is the most direct and complete map of MDEMG → successor
yet produced.** Our D-010 hybrid superstructure can adopt all
seven explicitly. The *interceptor pattern* is the Jiminy precedent
in D-010 named at substrate-agnostic level. The *multi-temporal
RSIC scaffolding* is the Q-008 recursive horizon expansion in
operational form. The *constraint-and-evidence schema* is the
underpinning of D-010's AI2AI protocol. The *multi-LoRA serving
infrastructure* is D-010's LoRA-managed continuous-learning
surface. Five of the seven map directly to existing program
commitments.

## Fork-timing specification (§3)

The document operationalizes the fork gate as a five-indicator
binary checklist:

- **FT-1:** Phase 11 GRPO retry produces a definitive signal
  (currently in flight at PR #358; ETA was 2026-04-27 ~17:00 UTC)
- **FT-2:** Phase 12 HITL DPO either runs or is consciously
  skipped with documented rationale
- **FT-3:** Multi-adapter serving operational at all 16 call sites
  with 1+ week of production validation
- **FG-1:** Documentation freshness (AGENT_HANDOFF.md, CHANGELOG.md
  within 2 days of HEAD; mandatory doc phase compliance verified
  for last 2 weeks of PRs)
- **FG-2:** Architectural intuitions documented (each §2.8
  carry-forward template specified as substrate-agnostic pattern)

When all five indicators are green, the fork is unblocked. The
document recommends: at 3-of-5 green, **begin drafting the
successor framework's foundational document in parallel**.

**This is operationally significant for our program.** D-010 is
already a candidate for a piece of the successor's foundational
document. The fork-readiness state determines when our
specification work transitions from "informed by MDEMG" to
"governing post-fork work." Once 3-of-5 indicators are green, we
should be drafting in parallel.

## The successor framework foundational document (§8)

§8 specifies what the successor's foundational document must
contain. Two key claims:

**It is written before implementation begins**, not after 105
phases of work — the successor benefits from MDEMG R&D's
architectural intuitions. **It commits at the substrate level**,
not the application level — MDEMG's VISION.md describes an
application; the successor's foundational document describes a
substrate.

The document specifies **five substrate commitments and five
governance commitments**:

**Substrate commitments (long-term goal made architecturally
concrete):**

1. Trainable representations
2. Reference-frame mechanism
3. World-model component
4. Recursive self-improvement mechanism (architecture-and-rule
   level, distinguished from RSIC's parameter-tuning scope)
5. Object-centric slot structure (persistent entity identity
   surviving context changes)

**Governance commitments (carry-forward patterns from MDEMG):**

6. The interceptor pattern
7. Multi-temporal RSIC scaffolding
8. The UxTS framework family pattern
9. Constraint-and-evidence schema
10. Multi-LoRA serving infrastructure

(Substrate items 1-5 align with §2.7's four demanded commitments
plus the additional commitment for object-centric slot structure
referenced in the Risk Analysis. Governance items 6-10 are five of
the seven carry-forward patterns from §2.8; the fail-open default
and dual-gate regression discipline are not promoted to commitment
level.)

**The foundational document is the operational contract** between
the long-term goal (research-grade and high-level) and the
successor's first implementation sprint (concrete and bounded).

## The nine-sprint research corpus

With the strategic frame established, the nine sprint plans
(01-09) become readable as **MDEMG-internal prototypes that
inform the successor framework**. Per §8.3, each sprint produces
one of: architectural intuitions (FG-2 fork-gate work), prototype
empirical signal (O-PROTO opportunities), or specification work
(O-LT specifications).

The sprint corpus organizes into **three threads**:

### Thread A — Predictive Coding / Free Energy Principle (sprints 01 → 02 → 03 → ... → 09)

Strict sequential dependency. Each builds on the prior. Capstone
is sprint 09 (active-inference unification), explicitly research-
stage, with a stated 9-12 month calendar duration.

**Sprint 01 — PC Reframe & Surprise Routing (PC-REFRAME-A).** The
sprint identifies that MDEMG already implements precision-weighted
evidence integration in DH-005:

```
overall = Σ(w_i · c_i · s_i) / Σ(w_i · c_i)
```

The document's §1 framing: "this is mathematically identical to a
Bayesian precision-weighted posterior mean and, structurally, to
predictive coding's hierarchical weighted-error integration." The
formula is *also* identical to "the foundational quantity in PC,
Kalman filtering, and Bayesian evidence accumulation." Surprise
detection (`internal/conversation/surprise.go`) is identified as a
"4-factor variational surprise proxy." The Jiminy interceptor loop
(detect → correct → iterate → learn) is identified as active
inference (perception → action → re-observation → learning).

The sprint's claim: "the formal PC framing has mostly arrived
without anyone naming it." Sprint 01 names it (no math change) and
adds the missing routing piece (fast-track / normal / quarantine
based on surprise × source-trust).

**This is a substantial framing finding for our program.** OBS-003
identified DH-005 as "data-sufficiency-aware aggregation" —
correct but undersold. Sprint 01 identifies it as
**precision-weighted Bayesian evidence integration / hierarchical
predictive coding**. The math is the same; the framing is more
load-bearing. **D-010's hybrid superstructure should commit to the
PC/FEP framing explicitly** rather than merely inheriting the
formula. Filing as Q-011.

**Sprint 02 — Precision-Weighted Hebbian η.** Extends DH-005's
precision weighting from the health formula to the Hebbian edge
update rule itself. Per-node `activation_confidence c_n ∈ [0,1]`;
per-edge effective learning rate `η_eff = η · c_a · c_b`. High-
confidence updates have stronger effects than low-confidence ones.

**Sprint 03 — Top-Down Predictions & Prediction-Error Promotion.**
The "highest-risk architectural sprint" in the PC/FEP thread.
Adds a reciprocal prediction stream: each higher-layer node
maintains a predicted activation pattern over its lower-layer
children; divergence drives learning. Replaces evidence-counting
promotion with prediction-error-driven promotion. Without this,
"MDEMG is a one-way hierarchy (bottom-up promotion only), not a
true PC hierarchy."

**Sprint 09 — Active Inference Unification (capstone).** Unify
Jiminy, RSIC, and Consulting under a single variational free
energy objective. Each currently has its own tunable parameters;
FEP would unify them — epistemic value (information gain) and
pragmatic value (outcome quality) emerging as the two natural
components of free-energy minimization. Explicitly research-stage,
explicitly uncertain ("a placeholder; its real purpose is to hold
a spot in the sprint sequence so that we remember FEP unification
is the theoretical endpoint of this thread"). Calendar estimate
9-12 months.

**Important:** Sprint 09 includes an "Alternative: Don't Do This"
section. Decision criterion: after Sprints 01-07 ship and soak 3
months, if "conflicting-guidance" incidents (Jiminy / RSIC /
Consulting disagree) are <1% of sessions, heuristics compose fine
and FEP unification is not worth the risk. Only if frequency is
>5% does unification become natural.

### Thread B — Independent tracks (sprints 04, 05, 06, 07, 08)

Can run in parallel with Thread A. Each independently scoped.

**Sprint 04 — Column-Voting Retrieval.** Replace weighted-linear-
combination retrieval with an ensemble of independent retrieval
"columns" (embedding, BM25, graph-proximity, structural, temporal,
role-scoped) combined by Reciprocal Rank Fusion. Adds **consensus
strength** as an uncertainty metric. Hawkins-inspired (Paper 1).

**Sprint 05 — Context-Specific Node Activations.** Shift from
"one embedding per node" to "one embedding per (node, context)
pair." Add lightweight `context_fingerprint` (256-bit sparse
representation, ~2% active) to observations. HTM distal-dendrite
analog. Solves the polysemy problem: `ErrorHandler` in `auth.go`
during authentication review is a different thing from
`ErrorHandler` in `payments.go` during incident triage.

**Sprint 06 — Sparse Retrieval Activation.** HTM SDR pattern.
Only top 2% of nodes "fire" for a query (gated by activation
threshold). Enables set algebra across queries. Cleaner signal,
cheaper downstream, principled uncertainty.

**Sprint 07 — Forward-Forward Shallow Heads.** FF training (no
backprop) for three shallow decision heads: promotion classifier
("should this L_n node be promoted to L_{n+1}?"), trust-bin
classifier (replacing the rule from sprint 01), context-
appropriateness classifier. Bounded scope: FF only for shallow
heads over frozen LLM embeddings (the literature-identified sweet
spot per Ororbia & Mali paper). Cross-encoder re-ranker stays on
backprop.

**Sprint 08 — HTM Sequence Memory for Predictive Jiminy.**
Variable-order Markov model with HTM-style representation: each
(state, context) pair has its own "cell"; cell records what came
next. Stored as a Neo4j subgraph. Fast learning (1-3 observations
makes a pattern predictive). Transforms Jiminy from reactive to
genuinely predictive — anticipates next action.

### Carry-forward implications across the corpus

The nine sprints are not just MDEMG improvements; they're
**operational tests of architectural patterns the successor will
need at substrate level**. The PC/FEP framing in sprint 01, the
prediction-error mechanism in sprint 03, the column voting in
sprint 04, the context-fingerprints in sprint 05, the sparse
activation in sprint 06, the FF training in sprint 07, the
sequence memory in sprint 08, the FEP unification in sprint 09 —
each is a *probe* of a substrate commitment.

Sprints 01-03 probe predictive-coding substrate.
Sprint 04 probes ensemble/voting reasoning (relates to D-010's
multi-faculty orchestration).
Sprint 05 probes compositional reference-frame structure (the
manifesto's reference-frame commitment).
Sprint 06 probes sparse representations (HTM-aligned).
Sprint 07 probes biologically-plausible learning (no-backprop, the
NGRAD-style local-credit pattern from the white paper review).
Sprint 08 probes predictive sequence learning.
Sprint 09 probes free-energy-principle unification.

**This is exactly the empirical input MDEMG R&D is supposed to
produce per §8.3.** The sprints are not feature work; they are
the mechanism by which MDEMG produces architectural intuitions
for the successor.

## The supporting documents

### `risk-opp-04232026-01.md` — Sprint Planning Input

24K, dated 2026-04-24. Translates the April 23 MDEMG Risk &
Opportunity Analysis into sprint-ready entries. Identifies **nine
risks** and **nine opportunities**. Five most consequential themes:

- **Community & contribution structure** — zero external
  contributors; vision requires community collaboration
- **Model & deployment portability** — FT-LoRA implicitly Mac/MLX-
  first, vision targets cross-platform Docker-first
- **Scale ceilings** — DBSCAN O(n²) and two parallel decay systems
  threaten emergence at target scale
- **North-star FT objective** — FT-OAI-003 blocked on
  authorization; "cheap FT ≈ prod" thesis unverified
- **Multi-agent completeness** — Phase 36 (Observation Forwarding)
  unimplemented

This document operates within the same risk/opportunity framing as
mdemg-specification.md but at a finer-grained planning level. The
two documents complement: spec sets strategy; risk-opp produces
sprint-ready operational items.

### `MDEMG_FT_LORA_PACKAGING_SPEC.md` — Packaging spec

17K, dated 2026-04-24. Specifies how to package, distribute, and
install the MDEMG LoRA fine-tuned model (built on Qwen3-14B FP4
dense base) as an **optional component** of the `brew install
mdemg` flow. Architectural decision, Homebrew integration pattern,
distribution backend, deployment targets, sprint-ready task
breakdown.

The base model is **Qwen3-14B FP4 dense** — concretizing
@reh3376's earlier note about running LoRA SFT + RLHF DPO on
Qwen-3.6-14b-dense. This is the actual operational target.

### `lora-training-glossary.md` — LoRA vocabulary reference

54K, dated 2026-04-27. Substantial glossary of LoRA fine-tuning
terms. The opening paragraphs articulate the key architectural
property: "a single base model can host many adapters
simultaneously, with the active adapter selected programmatically
per request." This is the multi-LoRA serving pattern — one of the
seven carry-forward patterns identified in §2.8 of the
specification.

The glossary is reference material for sprint planners; it is not
itself architectural. Worth noting it exists.

## How this changes our program

The corpus provides multiple substantial findings, ranked by
architectural significance:

### Finding 1 — D-010's hybrid superstructure aligns with the successor framework

**Confidence: 5.** D-010's commitment to "learning superstructure +
LLM as language plug-in + teacher-pupil dynamic + AI2AI protocol +
LoRA-managed continuous learning" is approximately what the source
documents call the "successor framework." Both:

- Treat MDEMG-current as substrate for inheriting patterns, not as
  the target architecture
- Commit to substrate-level capabilities MDEMG cannot supply by
  design (representation learning, prediction, reference-frame
  construction, meta-learning)
- Use multi-LoRA serving as the continuous-learning surface
- Adopt the interceptor pattern as governance machinery
- Use multi-temporal cycles for self-improvement

**The mapping is so direct that our D-010 should be reframed as a
candidate contribution to the successor framework's foundational
document (§8).** The vault's program objectives (D-009) and
hybrid architecture (D-010) are *exactly* the kind of input §8.3
calls for.

### Finding 2 — DH-005 is precision-weighted Bayesian / PC framing, not just data-sufficiency-aware

**Confidence: 5.** OBS-003 identified DH-005 as
"data-sufficiency-aware aggregation" and flagged it as reference
architecture for the entity's homeostatic boundary. Sprint 01
sharpens this: the formula is mathematically identical to **the
foundational quantity in PC, Kalman filtering, and Bayesian
evidence accumulation**. The pattern is not just useful; it is the
right substrate-level mathematical commitment.

**Implication: D-010 should explicitly commit to the PC/FEP
mathematical framing for the entity's state representation.**
This is filed as Q-011 for capture as a decision when the
implications stabilize.

### Finding 3 — The four substrate-level commitments name our open architectural questions exactly

**Confidence: 5.** §2.7 names the four substrate-level commitments
the long-term goal demands that MDEMG cannot supply:

| Commitment | Our vault artifact |
|---|---|
| Trainable representations | D-009 mid-level capability; partly Q-006 |
| Generative modeling / prediction | Q-008 (recursive horizon); Q-010 (long-horizon beyond JEPA) |
| Reference-frame construction | D-009 mid-level capability; Q-007 (projection-and-anchoring) |
| Meta-learning / mechanism modification | Q-008; partial in D-010's LoRA continuous learning |

The vocabulary alignment is striking. The four commitments are
exactly the architectural questions our program has been working.
**This validates that our program objectives target the right
gap.**

### Finding 4 — The seven carry-forward patterns map directly to D-010 components

**Confidence: 4.** Five of the seven carry-forward patterns map
directly to commitments already in D-010 or to D-010-adjacent
work:

| Carry-forward pattern | D-010 component |
|---|---|
| Interceptor pattern | Teacher-pupil dynamic (Jiminy precedent) |
| Multi-temporal RSIC scaffolding | Recursive predictive horizon (Q-008) |
| Multi-LoRA serving | LoRA-managed continuous-learning surface |
| Constraint-and-evidence schema | AI2AI protocol foundation |
| Fail-open default | Implicit in robust-systems thinking |
| UxTS pattern | Not yet in D-010; candidate |
| Dual-gate regression discipline | Not yet in D-010; candidate |

**The two unmapped patterns (UxTS, dual-gate regression) are
candidates for explicit incorporation into D-010 or its
successors.** Filing as Q-012 (UxTS as constraint specification
language for the superstructure) and Q-013 (dual-gate regression
discipline for the entity's self-evaluation).

### Finding 5 — The fork-timing concept is operationally relevant to our program

**Confidence: 4.** Until reading this corpus, our program treated
MDEMG as a stable backdrop. The fork-timing framing changes that
— MDEMG itself has a finite-duration R&D phase with a planned
transition. Our program's specification work (D-009, D-010,
upcoming D-NNN) is candidate input to the successor framework's
foundational document.

**The 3-of-5 fork-readiness threshold means our specification work
should be ready to "begin drafting the successor framework's
foundational document in parallel" within a reasonable timeframe.**
Currently FT-1 (PR #358 retry) and possibly FT-2 are at or near
green; FT-3 (multi-adapter at 16 call sites) is in progress;
FG-1 and FG-2 are likely partial. Estimated 3-of-5 may be reached
within weeks.

### Finding 6 — Vocabulary care is non-trivially important

**Confidence: 5.** The §2 alignment assessment explicitly flags
three vocabulary drifts: emergence, reference frames, self-
improvement. Each is an MDEMG term that overlaps with long-term-
goal vocabulary in misleading ways.

**Our vault must be careful about the same drifts.** Specifically:

- "Recursive self-improvement" in D-009 and Q-008 must be
  understood as a *target capability of the entity we're
  designing*, not as a description of MDEMG's RSIC. RSIC is
  parameter-tuning within fixed mechanisms; the entity's recursive
  self-improvement extends to mechanism modification, architecture
  modification, and goal modification — explicitly *not* what
  RSIC does.
- "Reference frames" in D-009 must mean Hawkins-style learned
  distributed reference frames, not MDEMG-style inherited frames.
- "Emergence" if it appears in our artifacts should mean
  manifesto-level emergence (new representational units, new
  conceptual axes), not DBSCAN clustering.

This is a vocabulary discipline issue that bears on every future
artifact we produce.

### Finding 7 — The PC/FEP unification path provides a research arc

**Confidence: 3.** Sprint 09's framing — that all three of MDEMG's
decision systems (Jiminy, RSIC, Consulting) could be unified under
variational free energy — suggests our hybrid superstructure may
have a similar unification path. **The superstructure's
orchestration of multiple faculties (LLM, V-JEPA, eventual
proprioception/audition/etc.) may admit a free-energy-principle
unification at the architectural level.** This is highly
speculative and far in the future, but worth flagging as a
candidate organizing principle for the superstructure's eventual
mathematical foundation.

## Confidence calibration

| Claim | Rating | Why |
|---|---|---|
| MDEMG framed as R&D vehicle, successor as separate artifact | 5 | Direct from mdemg-specification.md §0.2, repeatedly stated. |
| Four substrate-level commitments demanded by long-term goal | 5 | Direct from §2.7. |
| Seven carry-forward patterns | 5 | Direct from §2.8. |
| Five-indicator fork-timing checklist | 5 | Direct from §3.3. |
| Successor foundational document specification | 5 | Direct from §8. |
| §2 alignment assessment characterization (three vocabulary drifts) | 5 | Direct from §§2.2, 2.3, 2.4. |
| DH-005 = precision-weighted Bayesian = PC framing | 5 | Direct from sprint 01 Theoretical Foundation. |
| Nine sprint plans organized into PC/FEP track + independent + capstone | 5 | Direct from reading all nine. |
| Sprint dependencies and sequencing | 5 | Direct from each sprint's "Dependencies" section. |
| Each sprint's risk analysis and gate criteria | 4 | Read closely for sprints 04-09 in full; sprints 01-03 read at framing + key sections. |
| MDEMG-as-substrate vs successor-as-different-foundation | 5 | Direct from §2.6 option 3 framing. |
| §1.10 absences (no PC, no world model, no reference frames, etc.) | 5 | Direct from §1.10; document's own grep-verification claim. |
| Sprint 09's "9-12 month calendar" estimate | 4 | Direct from sprint 09 size estimate. |
| Mapping of carry-forward patterns to D-010 components | 3 | My inference based on pattern descriptions; load-bearing for our work. |
| D-010 ↔ successor framework alignment claim | 4 | Strong textual evidence in both directions; the mapping is striking but my inference about the equivalence is interpretive. |
| Vocabulary-drift implications for our vault | 4 | Direct from §2 plus my inference about which of our artifacts are at risk. |
| PC/FEP unification as candidate organizing principle for superstructure | 3 | Speculative; sprint 09 suggests the path; our application is interpretive. |

## Surfaced questions (file as Q-artifacts in subsequent PR)

1. **Q-011 — Should D-010 explicitly commit to the PC/FEP
   mathematical framing for the superstructure's state
   representation?** Sprint 01 identifies DH-005 as precision-
   weighted Bayesian / hierarchical PC. Adopting the framing
   explicitly would tighten our mathematical foundation and align
   with the successor framework's architectural intuitions.

2. **Q-012 — Should the entity's constraint-specification language
   adopt the UxTS pattern (schema + spec + runner + CI)?** UxTS is
   one of the seven carry-forward patterns, named by the
   specification document as "possibly MDEMG's strongest." Our
   vault has not yet specified how the entity's constraints are
   expressed; UxTS is a candidate.

3. **Q-013 — Should the entity's self-evaluation adopt the dual-
   gate regression discipline (aggregate + per-task gates)?** This
   is the seventh carry-forward pattern, not yet in D-010. The
   pattern catches catastrophic per-task regressions masked by
   aggregate metrics — directly applicable to the entity's
   recursive self-improvement (Q-008).

4. **Q-014 — How does our program relate operationally to MDEMG's
   fork-timing checklist?** When MDEMG reaches 3-of-5 indicators
   green, the specification recommends parallel drafting of the
   successor's foundational document. Our specifications work
   should sequence against this.

5. **Q-015 — Is the active-inference unification path (sprint 09)
   a candidate organizing principle for the superstructure's
   mathematical foundation?** Speculative; the architectural
   parallel is suggestive but not derived. Worth tracking as the
   superstructure's specifications mature.

These will be filed as Q-artifacts in a follow-up PR alongside
this one or in the next cycle.

## Implications for next cycles

**Cycle 2A (J17 deep trace) priority is unchanged but its scope
sharpens.** J17 should be deep-traced not just for its current
operational structure but for what it tells us about the
constraint-and-evidence schema as a substrate-agnostic carry-
forward pattern (governance commitment 9 in §8.2).

**Cycle 2B (Jiminy + RSIC combined trace) gains a third axis.**
Beyond their operational mechanism (D-010 teacher-pupil precedent)
and recursive self-improvement (Q-008), the trace should also
identify:

- The interceptor pattern at substrate-agnostic level (governance
  commitment 6)
- The multi-temporal scaffolding at substrate-agnostic level
  (governance commitment 7)
- The vocabulary drift between RSIC's operational scope and the
  long-term goal's recursive self-improvement scope

**A new candidate cycle emerges: Cycle 2X — read the white-paper
review.** All nine sprint plans cite a `mdemg-white-paper-review.md`
that I have not seen. If it is in `docs/research/` of the public
repo or available locally, reading it would substantially deepen
the PC/FEP/HTM framing. **Surfaced for @reh3376 to consider.**

## What this artifact does NOT include

- Detailed read of risk-opp-04232026-01.md beyond Executive Summary
  (24K of risk and opportunity entries; valuable for sprint
  planning but not strategy)
- Detailed read of MDEMG_FT_LORA_PACKAGING_SPEC.md beyond
  identification of Qwen3-14B FP4 dense as base
- Detailed read of lora-training-glossary.md (vocabulary reference;
  not architectural)
- Sprint-by-sprint deep trace of internal phases (each sprint has
  6-8 phases with detailed gap entries; full trace would balloon
  this artifact)
- Code-level verification that the sprint plans' identifications
  (e.g., "DH-005 is precision-weighted Bayesian") match the actual
  code at the level of mechanism (the sprint plans claim it; code
  verification is its own cycle)

These belong in deep-trace cycles, when the relevant subsystem
becomes load-bearing for specification work.

## Summary

**MDEMG is not the cognitive substrate the program is building. It
is the R&D vehicle that produces architectural intuitions for the
substrate the program is building.** That substrate is what
@reh3376's documents call the *successor framework* and what our
vault calls the *hybrid superstructure*. The two terms describe
the same artifact at different levels of architectural maturity:
the successor framework specification is the operational form;
D-010 is the program-level form.

Five substrate commitments and five governance commitments (§8.2)
specify what the successor framework must contain. The four
substrate-level commitments demanded by the long-term goal map
exactly to D-009's mid-level capability targets and the open
architectural questions Q-005 through Q-010. The seven carry-
forward patterns map five-of-seven to D-010 components with two
candidates (UxTS, dual-gate regression discipline) for explicit
incorporation.

The fork-timing specification creates an operational sequence: as
MDEMG's R&D phase completes (currently 2-3 of 5 fork-indicators
green, estimated weeks to 3-of-5), the program transitions from
"informed by MDEMG" to "drafting the successor's foundational
document in parallel." Our D-009 / D-010 / D-011 work is candidate
input to that document.

The nine sprint plans are not just MDEMG improvements; they are
operational tests of architectural patterns the successor needs at
substrate level. The PC/FEP framing in sprint 01 sharpens DH-005
from "data-sufficiency-aware aggregation" to "precision-weighted
Bayesian / hierarchical PC" — a substantial mathematical commitment
candidate for our superstructure.

Vocabulary discipline matters: emergence, reference frames, and
self-improvement are MDEMG terms that overlap with long-term-goal
terms in misleading ways. Our vault must be careful that these
terms in our artifacts (D-009, Q-008, etc.) refer to the entity
we're designing, not to MDEMG's operational implementations.

This is the most substantively reframing observation produced in
the program to date.
