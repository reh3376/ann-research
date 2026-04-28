---
id: T005
type: thread
slug: program-planning
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 5
where: []
tags: [meta, planning, cycle, methodology, governance]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude (post-PR-#9 merge)"
  - "@reh3376's cycle formulation: plan / research / specification / plan / prototyping / validation / assess / unblock / plan"
  - "D-2026-04-28-009 (program objectives — T005 produces the operational plan that advances those objectives)"
  - "D-2026-04-28-010 (hybrid architecture — what the plan plans toward)"
  - "D-2026-04-28-011 (V-JEPA visual faculty — first faculty-level architectural commitment T005 will plan execution for)"
owner: "@reh3376 (PI), Claude (drafts and surfaces)"
current_question: "What is the next cycle's focus, what mode are we in, and what does it take to reach the next unblock gate?"
---

# Thread T005 — Program Planning

## Goal

Plan the work of the program in iterative cycles, capturing the
plan as a living document that revises at every cycle gate. T005 is
the planning thread; it does not produce research findings,
architectural decisions, or implementation. It produces *the
sequence of work that produces* those things.

The thread exists because long-horizon research programs that
attempt linear sequencing fail in characteristic ways: the early
plan is wrong because the program hasn't yet learned what's
load-bearing; mid-plan revisions get lost in conversation rather
than recorded; the relationship between completed work and remaining
work becomes unclear over time. T005 prevents these failures by
making planning explicit, recurring, and durable.

## The cycle

The program runs through cycles with the following structure:

```
plan → research → specification → plan → prototyping → validation → assess → unblock → plan → ...
```

Each cycle has a **focus** — a specific question, hypothesis,
capability, or commitment that the cycle advances. The cycle's
modes (research, specification, prototyping, validation) operate
on that focus.

### The modes

**plan.** Set or revise the plan. Identify the next cycle's focus.
Identify the modes that will be active in this cycle. Identify the
gate criteria for moving from one mode to the next. Identify the
unblock gate criteria for closing the cycle. Plan is a recurring
mode, not a one-time start; it appears at cycle start, between
specification and prototyping, and after the unblock gate.

**research.** Produce mechanism-level understanding relevant to
the cycle's focus. Trace work in MDEMG, `tbp.monty`, or external
literature; surface observations, questions, hypotheses. Output
lands as `OBS-`, `Q-`, or `H-` artifacts in the vault.

**specification.** Refine architectural commitments into
buildable specifications for the cycle's focus. Component
definitions, interfaces, behavioral contracts, schemas. Output
lands as `D-` artifacts (committed specifications) or detailed
appendices to existing decisions.

**prototyping.** Build working code that exercises the
specification. Prototypes are not the final system; they are
learning instruments whose job is to surface gaps between
specification and implementation reality. Output lands as code
in implementation repositories (likely MDEMG itself for some
prototypes, separate repositories for others), with
cross-references in the vault.

**validation.** Measure what the prototype does. Test against the
expected behavior the specification described. Validation answers
*"did the cycle do what we expected?"* — execution success or
failure. Output lands as observation artifacts (`OBS-`) capturing
what was measured.

**assess.** Interpret the validation. Was what we expected the
right thing to expect? Assessment answers *"was the direction
correct?"* — design success or failure. A cycle can validate
successfully (it did what we expected) but assess as wrong (what
we expected was the wrong thing). Output lands as decision
artifacts (`D-` for "the direction was right; continue" or `DE-`
for "the direction was wrong; this is a dead-end") and as new
questions for the next cycle.

**unblock.** A gate that opens new cycles. Given the cycle's
output (what was learned from research, what was specified, what
was built, what was validated, what was assessed), what cycles
does this unblock? Unblock is generative: each completed cycle
produces zero, one, or several candidate next cycles. Unblock is
the decision about which to pursue. The unblock decision is
itself part of the next plan.

### Gate criteria

Each transition between modes has a gate. The gate criteria are
defined in the cycle's plan at the start of the cycle. Common
gates:

- **plan → research:** the focus is named; the mechanism question
  is articulable; the success criterion for research is named.
- **research → specification:** evidence is sufficient that
  specification can be drafted without re-deriving evidence
  midway through.
- **specification → plan:** the specification is committed; what
  remains is implementation work.
- **plan → prototyping:** implementation scope is bounded; the
  success criterion for the prototype is named.
- **prototyping → validation:** the prototype is built and runs.
- **validation → assess:** measurements are recorded.
- **assess → unblock:** the assessment is captured as decision
  or dead-end.
- **unblock → plan:** next cycle's focus is identified.

A mode does not advance until its gate criterion is met. If a
gate criterion cannot be met, that itself is a finding — usually
that the cycle's plan was wrong and needs revision (back to
plan).

### Cycle scoping

Cycles are not all the same size. A cycle's scope is set at
plan time and should be the **smallest scope that produces a
genuine learning** — small enough to complete, large enough to
matter. Scoping principles:

- **A cycle should fit inside a coherent stretch of working
  attention.** A cycle that spans months of calendar time loses
  context between modes. If a cycle is that big, decompose it.
- **A cycle should have one focus.** Multi-focus cycles produce
  blurred learnings. If multiple things are being investigated
  simultaneously, run them as parallel cycles, not as one
  multi-focus cycle.
- **Not every cycle hits every mode.** A pure-research cycle
  may stop at "assess" without specification, prototyping, or
  validation. A specification-refinement cycle may not need
  prototyping. The cycle structure is the *available* modes,
  not a checklist that must be exhausted.

## Method

This thread operates as follows:

1. **Plan-of-record lives in `plan.md`** alongside this thread
   file. The plan is updated at every plan-mode transition. Git
   history preserves the plan's evolution; the file always shows
   the *current* plan.

2. **Each cycle gets a log entry.** As cycles complete, the
   thread log captures the cycle's focus, what was done, what was
   learned, and what was unblocked. The log is the cycle history.

3. **Cycles produce artifacts in their own right.** A research
   cycle produces observations and questions; a specification
   cycle produces decisions; a prototyping cycle produces code
   (with cross-references); a validation cycle produces
   measurements; an assess cycle produces direction decisions or
   dead-ends. Those artifacts live in their normal vault locations,
   with the thread log noting which cycle produced them.

4. **Plan revisions are decisions when substantial.** A small
   plan adjustment (sequence change, scope tweak) lives in the
   `plan.md` revision history. A substantial plan change
   (priority reordering, mode reconfiguration, cycle abandonment)
   gets its own decision artifact recording the change explicitly.

5. **The plan is not a contract.** It is the current best
   understanding of what to do next. New evidence overrides the
   plan; the plan revises freely. The plan's job is not to bind
   future work but to make current work coherent and current
   priorities explicit.

## Relationship to other threads

**T001 (Monty trace), T002 (MDEMG trace), T003 (grounding question),
T004 (entity first principles).** These threads contain ongoing
work that T005's cycles operate on. A research cycle whose focus
is "what is in `sensor_module_outputs[i]` after `sm.step()`?"
advances T001. A research cycle whose focus is "what does an
MDEMG ingest cycle actually do?" starts T002. A specification
cycle whose focus is "what is the AI2AI protocol's message
schema?" advances both T002 and the architectural specification
agenda.

T005 does not subsume these threads. They remain the homes for
their respective subject matter. T005 schedules the work that
flows through them.

**`D-2026-04-28-009` (program objectives) and `D-2026-04-28-010`
(architectural commitment).** T005's cycles advance toward what
these decisions name as the program's purpose and architecture.
Every cycle's focus should be readable as contributing to
something named in D-009 or D-010.

## State

**Current.** Thread opened 2026-04-28 with the cycle structure
codified per @reh3376's formulation. Initial plan-of-record
populated in `plan.md` with the first cycle's focus, mode, and
gate criteria.

The first cycle is a planning-and-research cycle to bring T002
(MDEMG trace) up to a level of evidence comparable to T001
(Monty trace), since T002 has not yet produced any observations
and is currently the binding constraint on architectural
specification work. See `plan.md` for the cycle definition.

## Open work

- [ ] First cycle execution: research-mode work to bring T002 to
      first-observation status.
- [ ] As cycles complete, accumulate cycle log entries here.
- [ ] As the plan evolves substantially, capture significant
      plan changes as decision artifacts.

## Artifacts produced by this thread

- `plan.md` (alongside this thread file) — living plan-of-record
- Cycle log entries (in this thread file's Log section)
- Plan-change decisions when revisions are substantial

## Log

### 2026-04-28 — Cycle 1 produces second observation (OBS-004) on MDEMG forward direction

After OBS-003 (current state) reached substantial depth, @reh3376
clarified that the previously-missing `mdemg_sprint_ideas`
directory is at `docs/research/mdemg_sprint_ideas/` (existing
locally but not pushed to the public mdemg repo) and uploaded
all 13 markdown files directly. The corpus is substantively
different from current-state docs — it is forward-looking
architectural research.

**Cycle hygiene call:** Per the cycle methodology, this is a
distinct subject matter from OBS-003. Rather than continuing to
extend OBS-003 (already 994 lines, growing unwieldy), produced
**OBS-2026-04-28-004 as a companion observation**. OBS-003 = what
MDEMG is today; OBS-004 = what MDEMG plans to become next. Both
are part of Cycle 1's orientation deliverable.

**Files read:**

Strategic:
- `mdemg-specification.md` (155K, 1328 lines) — the master document
  framing MDEMG as R&D vehicle with successor framework on the
  horizon. §0.2 strategic frame, §2 alignment assessment, §3 fork-
  timing, §7 open questions, §8 successor framework foundational
  document.
- `risk-opp-04232026-01.md` (24K) — sprint planning input
  (Executive Summary read; 9 risks + 9 opportunities surveyed).

Sprints (forward-looking architecture):
- `01-pc-reframe-and-surprise-routing.md` (33K) — Theoretical
  Foundation read; identifies DH-005 as precision-weighted Bayesian
  / PC framing.
- `02-precision-weighted-hebbian-eta.md` (17K) — framing + key
  changes read.
- `03-top-down-predictions-and-promotion.md` (18K) — framing +
  schema read; identified as highest-risk PC/FEP sprint.
- `04-column-voting-retrieval.md` (15K) — full read (uploaded).
- `05-context-specific-node-activations.md` (15K) — full read
  (uploaded).
- `06-sparse-retrieval-activation.md` (11K) — full read (uploaded).
- `07-ff-shallow-heads.md` (14K) — full read (uploaded).
- `08-htm-sequence-memory.md` (12K) — full read (uploaded).
- `09-active-inference-unification.md` (13K) — full read (uploaded).

Reference:
- `MDEMG_FT_LORA_PACKAGING_SPEC.md` (17K) — head read; identifies
  Qwen3-14B FP4 dense as the base model.
- `lora-training-glossary.md` (54K) — head read; vocabulary
  reference, not architectural.

**Key findings captured in OBS-004:**

1. **MDEMG is explicitly framed as R&D vehicle, not production
   target.** The successor framework is a separate artifact, not
   a future MDEMG version. This re-orients our program's
   relationship to MDEMG.

2. **D-010's hybrid superstructure aligns directly with the
   successor framework.** Five of seven carry-forward patterns map
   to D-010 components; the four substrate-level commitments
   (§2.7) map to our D-009 mid-level capability targets exactly.

3. **DH-005 is precision-weighted Bayesian / hierarchical PC, not
   just data-sufficiency-aware.** Sprint 01 sharpens OBS-003's
   identification. Mathematical framing is more load-bearing.

4. **Vocabulary drift risks named explicitly.** Emergence,
   reference frames, recursive self-improvement — MDEMG terms
   that overlap with long-term-goal terms in misleading ways. Our
   vault must be careful.

5. **Fork-timing creates operational sequencing for our work.** At
   3-of-5 fork-indicators green, the specification recommends
   drafting the successor's foundational document in parallel.
   Currently estimated weeks to 3-of-5.

**Five new questions surfaced** (Q-011 through Q-015), to be filed
as Q-artifacts in a follow-up PR or in Cycle 2 unblock.

**One candidate new cycle surfaced:** read the
`mdemg-white-paper-review.md` referenced by all nine sprint plans
but not in the corpus uploaded.

### 2026-04-28 — Cycle 1 orientation extended again with five additional reads

@reh3376 added six files to the read list. Five existed and were
read; one (`docs/mdemg_sprint_ideas/`) does not exist in the repo
(verified: no directory by that name; `docs/development/ft-lora/`
and `docs/development/ft-oai/` contain sprint-plan files but the
named path is not present). Question to @reh3376 surfaced for
clarification.

**Files read and integrated into OBS-003:**

- `README.md` (29K) — benchmark methodology (whk-wms 507K LOC TS,
  120 questions, +5.2% vs baseline, -51% variance), key-features
  list including "Capability gap detection"
- `docs/features/j17-ai2ai-protocol.md` sections 4-8 (previously
  only 1-3.1) — full self-improvement journey, ProtocolHealth
  formula, reflection patterns, mutation actions, NLI scoring,
  JSONL training data
- `docs/features/j17-prompt-compression.md` (4.2K, full) — J17-PC
  applied to inputs of 5 LLM callers, 25-35% aggregate reduction
- `docs/features/neural-training-pipeline.md` (19K) — three
  training workstreams (cross-encoder/generative/embedding), 9-step
  generative LoRA pipeline with quality gates, 100× cross-encoder
  speedup
- `docs/features/rsic-feedback-loop.md` (12K) — closed-loop AR-1,
  reversible vs irreversible actions, **DH-005 health-weighting
  formula** (substantial finding)

**Key new findings extending OBS-003:**

- J17's current state (5.2× compression, 91% comprehension) is the
  result of four documented RSIC cycles, not initial design. The
  documented lesson — "better codes within a weak protocol hit a
  ceiling; the protocol itself must evolve" — is the recursive
  predictive horizon expansion principle in operational form.
- DH-005's `overall_health = Σ(w_i · c_i · s_i) / Σ(w_i · c_i)`
  excludes dimensions with zero confidence rather than penalizing
  them. The formula is *honest about what it doesn't know*. This
  is reference architecture for the entity's homeostatic boundary
  (Q-005) — the self-model can use the same data-sufficiency-aware
  aggregation.
- J17-PC compresses LLM *inputs*, not just outgoing AI2AI guidance.
  Per-caller opt-in. Prose-only with code/JSON/enums verbatim.
- Three training workstreams more developed than the AGENT_HANDOFF
  excerpt suggested. Generative LoRA has 9 explicit steps with
  multiple quality gates. Reference architecture for the entity's
  continuous learning.

### 2026-04-28 — Cycle 1 orientation extended with AGENT_HANDOFF.md

After the orientation deliverable was produced and PR #11 opened,
@reh3376 added `AGENT_HANDOFF.md` (1025 lines, 132K, top-level) to
the read list. Read in full and integrated into OBS-2026-04-28-003.

**Cycle hygiene call:** This extends Cycle 1 rather than opens
Cycle 2. PR #11 had not yet merged; the orientation was still in
`assess` mode. AGENT_HANDOFF.md is the same kind of source as the
others (markdown documentation about MDEMG structure). Reading it
now extends the orientation rather than starting a new cycle. The
artifact gets revised in the same PR before merge, so the
orientation is complete when assess closes.

**New material integrated:**
- Tech stack subsection (Neo4j 5.x, Go 1.24, OpenAI
  text-embedding-3-large 3072d, gRPC plugins, Python FastAPI sidecar)
- `internal/` 22-package layout
- Retrieval pipeline scoring formula with specific weights
  (vector 0.55 + activation 0.30 + recency 0.10 + confidence 0.05
  − hub_penalty 0.08 − redundancy 0.12)
- Phase registry confirmation (105 core ✅, 16 sidecar ✅, 5 cognitive
  gap ✅)
- Fine-tuning current state — FT-OAI-001 ✅ MARGINAL verdict,
  FT-OAI-002 ✅ tooling/telemetry, FT-OAI-003 📋 planned;
  local LoRA pipeline (PRs #246-250) all complete
- Architecture-maps-as-self-documentation pattern
  (`docs/architecture/maps/` auto-generated for Jiminy context)

**Implications added** to OBS-003:
- LoRA-managed continuous learning is more disciplined than D-010
  implied (explicit quality gates, cost caps, quality floors,
  gap-closure framing)
- MDEMG documents itself for its own use (precedent for entity
  structural self-knowledge; connects to Q-005 homeostatic boundary)

### 2026-04-28 — Cycle 1 plan revised mid-flight; orientation produced

**Original Cycle 1 plan:** "Bring T002 (MDEMG trace) to first-
observation status — produce a mechanism-level observation comparable
to OBS-2026-04-28-002 covering MDEMG's primary cognitive loop."

**Revision trigger:** @reh3376 noted MDEMG is over 1M LOC with
multiple major subsystems (graph base, UxTS, Jiminy, J17, CMS, RSIC,
plus the LLM call layer). The original mechanism-depth-on-primary-
loop scope was wrong for this size — would have produced a shallow
tour of subsystems without depth on any.

**Revised Cycle 1 scope:** Architectural orientation only. Produce a
structural map of MDEMG's subsystems, sufficient to scope subsequent
deep-trace cycles. Do not go to mechanism depth on any subsystem.

**Method change:** Doc-first reading, not @reh3376-led walk-through.
The MDEMG codebase is exceptionally well-documented (643 .md files,
covering all subsystems at multiple levels). Reading the
architectural and feature-level documents produced high-bandwidth
orientation faster than verbal walk-through would have. @reh3376
flagged this directly: "the mdemg codebase is very well documented,
it will probably be much quicker for you to find and review the .md
files relevant to each of those main topics."

**Deliverable produced:** OBS-2026-04-28-003 (MDEMG subsystem
orientation map). Identifies seven major subsystems, characterizes
each, maps their relationships, surfaces five questions, ranks
deep-trace candidacy. Cycle 1 now in `assess` mode pending @reh3376
review.

**Methodology validation noted:** This is the first plan revision
recorded under the cycle methodology. The plan-revision-mid-flight
worked as designed: scope-was-wrong was caught and corrected without
abandoning the cycle, the revised scope was achievable in the same
working session, and the deliverable now exists. The cycle
methodology survives its first contact with reality.

### 2026-04-28 — thread opened, cycle structure codified

Conversation post-PR-#9-merge converged on a planning methodology:
iterative cycles with explicit plan-research-specification-plan-
prototyping-validation-assess-unblock-plan structure. The
formulation came from @reh3376; my role was to verify the cycle
structure and capture it durably.

Key distinctions clarified in conversation:
- Validation answers "did the cycle do what we expected?" (execution)
- Assess answers "was what we expected the right thing to expect?" (design)
- Unblock is a gate that opens new cycles, not a debugging step
- Plan recurs at multiple points in the cycle, not just at start

Initial `plan.md` populated with first cycle's focus (T002 trace
work to first-observation status). Cycle execution begins next
turn.
