# Plan of Record

> **Living document.** Revised at every cycle gate. Git history
> preserves evolution; the current state is always the working plan.
>
> **Last revised:** 2026-04-28 (Cycle 1 revised mid-flight; orientation produced)
> **Current state:** Cycle 1 in `assess` mode pending @reh3376 review of OBS-2026-04-28-003
> **Next gate:** Cycle 1 `assess → unblock` decision

---

## Where we are

**Architectural commitments are firm but may be revised.** D-009
(program objectives), D-010 (hybrid superstructure: LLM language
plug-in, teacher-pupil dynamic, AI2AI protocol, LoRA-managed
continuous learning), and D-011 (V-JEPA family as candidate visual
faculty) form the current architectural foundation. **D-010 and
D-011 may be revised once T006 produces an architectural
specification** — the hybrid superstructure pattern in D-010 may
be a deployment shape over the architecture rather than the
architecture itself. Until T006 has a specification, D-009/D-010/
D-011 stand.

**First-principles framing is captured.** T004 holds five facets
(homeostatic boundary, dimensional minimum-commitment,
projection-and-anchoring, recursive predictive horizon, prime
directives) as open questions. T006 reads these as positive-space
inputs to its "what it must have" mode.

**Trace work — orientation phase complete; subordinated to T006.**
T001 (Monty) has produced two mechanism observations. T002 (MDEMG)
has produced OBS-003 (current state) and OBS-004 (forward
direction). Trace cycles 2A-2F (J17 deep, Jiminy+RSIC combined,
MEMG deep, file questions, T004 surfacing, continue Monty) are
subordinated to T006: they run when T006's work needs their
evidence, not on their own cadence.

**T006 (Architectural Innovation) is the central thread.** Opened
2026-04-28 with explicit working-method declaration: Roger provides
the novelty; Claude is a sharpening stone. T006's plan.md owns
the program's strategic phasing (Phase A: Architectural Innovation
Development; Phase B: Communication and Resource Acquisition;
Phase C: Build with Team and Funding). Cycle 1 of T006 is in plan
mode with focus 2B ("what it must have").

**Cycle structure is committed.** T005 codifies how planning and
work proceed across all threads.

---

## Strategic priorities (the long view)

**Strategic phasing now lives in T006's plan.md.** T006 was opened
on 2026-04-28 as the central thread for the program's architectural
work, with explicit working-method declaration (Roger provides the
novelty; Claude is a sharpening stone). T006's plan.md owns the
three-phase structure (Phase A: Architectural Innovation
Development; Phase B: Communication and Resource Acquisition; Phase
C: Build with Team and Funding).

T005's role under this revised structure is:

- **Schedule cycles across the program**, including T006's cycles.
- **Track cycle bookkeeping** (which cycle is active, what mode,
  what gate criteria, when to assess and unblock).
- **Surface methodology questions** that bear on multiple threads.
- **Maintain the cycle methodology itself** (the
  plan/research/specification/prototyping/validation/assess/unblock
  loop).

T005 no longer owns the strategic phasing. T006 does. This is a
deliberate split: T005 is the meta-thread (how we work); T006 is
the substantive thread for the program's central architectural
work (what we work on).

### Trace cycles 2A-2F: re-sequenced under T006

The trace cycles previously named in T005's cycle queue (J17 deep
trace, Jiminy+RSIC combined trace, MEMG deep trace, file surfaced
questions, T004 metacognition surfacing, continue T001) are now
**subordinated to T006 in Phase A**. They run when T006's
positive-space, negative-space, candidate-architecture, or
mathematical work needs their evidence — not on their own
cadence.

This is a substantive re-sequencing. Earlier in the program the
trace work was front-of-program; under the new framing, the
architectural innovation is front-of-program and trace work is
back-of-program (it produces evidence that informs the
architecture's specification, but it is not the central work).

### D-010 and D-011 may need revision once T006 has a specification

D-010 (hybrid superstructure) was committed before T006 existed.
Its framing as "the architecture" may be revised to "a deployment
shape over the architecture" once T006's specification is
produced. D-011 (V-JEPA visual faculty) similarly depends on the
eventual architecture. Until T006's specification exists, both
stand. After, both may be revised.

---

## Active cycle: Cycle 1 (in `assess` mode)

### Original focus and revision

**Original focus** (drafted in initial plan): Bring T002 (MDEMG trace)
to first-observation status — produce a mechanism-level observation
comparable to OBS-2026-04-28-002 (Monty's matching_step pipeline)
covering MDEMG's primary cognitive loop end-to-end.

**Revision** (mid-cycle, 2026-04-28): The original scope was wrong.
@reh3376 flagged that MDEMG is over 1M LOC with multiple major
subsystems — too much surface for one observation to cover at
mechanism depth. The cycle scope was narrowed to **architectural
orientation**: produce a structural map of MDEMG's subsystems
sufficient to scope subsequent deep-trace cycles, without going to
mechanism depth on any single subsystem.

**Cycle 1 method (revised).** Doc-first reading. The MDEMG codebase
is exceptionally well documented (643 .md files); reading the
architectural and feature-level documents produces high-bandwidth
orientation faster than walking @reh3376 through the codebase
verbally. Method: read VISION.md and the architecture index,
followed by per-subsystem feature docs (CMS.md, jiminy-inner-voice.md,
j17-ai2ai-protocol.md, UxTS spec head). Capture as a single
orientation observation. @reh3376 reviews; cycle assess + unblock
follows.

### Cycle 1 deliverable: OBS-2026-04-28-003

The orientation observation has been produced. Identified seven
major subsystems:

1. The Multi-Layer Emergent Memory Graph (MEMG) — the substrate
2. CMS — Conversation Memory System
3. RSIC — Recursive Self-Improvement Cycle
4. Jiminy Inner-Voice Service
5. J17 AI-to-AI Protocol
6. UxTS — Universal-x Test Specification framework
7. The LLM Call Layer + Modular Intelligence Plugin Architecture

The artifact captures each subsystem's role, codebase location,
relationships to other subsystems, and deep-trace candidacy ranking.
Five questions surfaced. Confidence ratings flagged per claim.

### Cycle 1 mode sequence (revised)

```
plan → plan-revision → research (doc-reading) → assess → unblock → plan
```

The original mode sequence was `plan → research → assess → unblock
→ plan`. The plan-revision sub-step is what @reh3376's "MDEMG is
1M LOC" feedback triggered.

### Gate criteria (revised)

**research → assess.** Cycle 1's research-mode gate criterion was:

- One artifact captures MDEMG's structural skeleton at the level
  of "what does each subsystem do, where does it live, and how do
  they relate" — sufficient to scope subsequent cycles. ✓ (OBS-003)
- The seven major subsystems are named and characterized. ✓
- Cross-subsystem relationships are mapped at high level. ✓
- A deep-trace candidacy ranking is proposed for Cycles 2..N. ✓
- Surfaced questions are filed (or noted for filing). 5 noted.

**assess → unblock.** Assessment for Cycle 1 asks (next conversation
turn):

- Did the orientation produce evidence at the level we expected?
  (validation question — execution)
- Was orientation the right mode, or should we have gone deeper
  on one subsystem instead of broad on all? (assess question —
  design)
- Where did the orientation surface unexpected things — gaps the
  architecture must close, or capabilities the architecture can
  inherit?

Initial findings worth flagging for assess discussion (from
OBS-003's "Implications" section):

- MDEMG is more developed than prior vault artifacts characterized.
  "R&D substrate" undersells it; "operational system serving as
  R&D vehicle" is closer.
- The teacher-pupil pattern in D-010 has direct precedent in Jiminy.
  The architectural pattern is operational, not aspirational.
- The AI2AI protocol commitment in D-010 has concrete reference in
  J17. Specifications can build on J17 rather than invent from
  scratch.
- The recursive self-improvement commitment in Q-008 has direct
  precedent in RSIC.
- The plugin architecture is the substrate-flexibility commitment in
  operational form.

### Unblock candidates (for next plan revision after assess)

Likely Cycle 2 focuses, ranked:

- **Cycle 2A: J17 deep trace.** Priority 1 candidate — fills
  D-010's AI2AI protocol commitment.
- **Cycle 2B: Jiminy + RSIC combined trace.** Priority 2 — D-010's
  teacher-pupil pattern in operational form, plus Q-008 precedent.
- **Cycle 2C: MEMG (graph) deep trace.** Priority 3 — substrate;
  bears on Q-006 and Q-007.
- **Cycle 2D: T004 metacognition surfacing.** Priority X (orthogonal)
  — @reh3376 has previewed this and will surface when ready;
  conversational, not a research cycle in the trace sense.
- **Cycle 2E: Continue T001 with Monty's LM internals.** Priority Y
  (parallel evidence) — would advance T003 synthesis.

The unblock decision happens at the close of Cycle 1, informed by
@reh3376's assess judgment.

---

## Cycle queue (provisional)

These cycles are projected from current understanding. The actual
sequence will be set by each cycle's unblock step, not from this
document.

**Cycle 2** — TBD (one of the candidates above, decided by Cycle 1's unblock)

**Cycle 3** — TBD

**Cycle 4** — TBD

**Long-projected cycles** (mentioned for orientation, will be planned
when their predecessors complete):

- Specification cycle for the AI2AI protocol message schema
- Specification cycle for the LoRA adapter management interface
- Specification cycle for the LLM faculty / superstructure boundary
- Research cycle for active inference / predictive coding literature
  bearing on Q-008 (recursive predictive horizon)
- Research cycle for hierarchical-temporal-prediction approaches
  bearing on Q-010 (long-horizon beyond JEPA)
- First end-to-end prototype cycle (Phase C) — far enough out that
  scoping it now is premature

---

## Open priorities not currently in cycle

These are tracked but not yet scheduled. They surface for inclusion
in the queue at unblock gates.

**Continue T004 first-principles surfacing.** @reh3376 has previewed
metacognition material as forthcoming. This is conversational
surfacing followed by deliberate capture, not a cycle in the
research-spec-prototype sense — but the artifacts it produces feed
all other cycles. Likely cadence: when @reh3376 raises the topic,
capture promptly; otherwise let lie.

**Web research on adjacent literatures.** Active inference
(Friston/Parr/Pezzulo), reference-frame and grid-cell literature,
JEPA family beyond what we've covered, biologically-plausible
learning rules. Each is a candidate research cycle. Sequencing
depends on which open questions become load-bearing for the next
specification cycle.

**Trace work continues for T001.** Reading Monty's LM internals
specifically (`EvidenceGraphLM.matching_step`, `_update_evidence`,
the motor system) bears directly on Q-001 (3D geometry load-bearing
in algorithms) and Q-006 (minimum dimensional commitment). Likely
becomes a cycle when its evidence becomes needed for specification.

---

## Things deliberately not in the plan

- Detailed sequencing for Phase B and Phase C work. The detail
  comes when the cycles approach.
- Resource estimates (time, compute, dollars). Premature.
- A "completion criterion" for the program as a whole. The program
  is open-ended research; "complete" is the wrong frame.
- A commitment to specific implementation languages, frameworks, or
  infrastructure for the eventual prototype. These choices belong
  in Phase C cycles, not in this plan.
