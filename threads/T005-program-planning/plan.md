# Plan of Record

> **Living document.** Revised at every cycle gate. Git history
> preserves evolution; the current state is always the working plan.
>
> **Last revised:** 2026-04-28 (Cycle 1 revised mid-flight; orientation produced)
> **Current state:** Cycle 1 in `assess` mode pending @reh3376 review of OBS-2026-04-28-003
> **Next gate:** Cycle 1 `assess → unblock` decision

---

## Where we are

**Architectural commitments are firm.** D-009 (program objectives),
D-010 (hybrid architecture: superstructure + LLM language plug-in,
teacher-pupil dynamic, AI2AI protocol, LoRA-managed continuous
learning), and D-011 (V-JEPA family as candidate visual faculty)
form the architectural foundation. The pattern is: superstructure
orchestrates faculty-as-foundation-model.

**First-principles framing is captured.** T004 holds five facets
(homeostatic boundary, dimensional minimum-commitment,
projection-and-anchoring, recursive predictive horizon, prime
directives) as open questions. None has stabilized into commitment;
each will be revisited as evidence accumulates.

**Trace work — orientation phase complete.** T001 (Monty) has
produced two mechanism observations. T002 (MDEMG) has produced
OBS-2026-04-28-003, the structural orientation map of MDEMG's
seven major subsystems. T002 has shifted from "not started" to
"orientation complete; deep traces queued."

**Cycle structure is committed.** T005 codifies how planning and
work proceed.

---

## Strategic priorities (the long view)

The program advances through three phases, in this order:

### Phase A — Evidence base for architecture

Bring trace work in T001 and T002 to mechanism-level depth. Surface
the gaps between current substrates (MDEMG, Monty) and what the
hybrid architecture (D-010, D-011) requires. Produce a clear-eyed
account of what each substrate provides, what each does not, and
what the architecture inherits versus what it must build.

This phase is dominated by **research mode**. Specification work
happens opportunistically — for components whose mechanism is
already understood (LLM faculty interfaces, V-JEPA integration
patterns from published work). Prototyping is sparse in this phase;
mostly tiny prototypes that surface specific specification gaps.

Phase A ends when: T001 has read into LM internals (where
geometry/dimensionality questions become concrete), T002 has
mechanism observations covering at least the ingest-retrieve cycle
and the AI2AI protocol, and the substrate-flexibility implications
(Q-006, Q-007) are answerable from evidence rather than principle.

### Phase B — Architectural specification

Refine D-010 and D-011 into buildable component specifications.
Define the superstructure's subsystems (state management, world
model, reference frames, predictive horizon, evaluation, governance).
Define the AI2AI protocol's message schema. Define the LoRA adapter
management interface. Define how the visual faculty (V-JEPA)
integrates with the superstructure.

This phase is dominated by **specification mode**. Research
continues to surface gaps. Prototyping increases — small components
get built to validate specifications before they're finalized.

Phase B ends when: the architecture has buildable specifications
for the subsystems whose first prototypes are next. Not all
subsystems need full specification at once; some can stay at
high-level commitment while others descend to detail.

### Phase C — First end-to-end prototype

Build the smallest meaningful slice that exercises the architecture
end-to-end: input through faculty, faculty output through
superstructure governance, superstructure decision back through
faculty as guidance, output to user. This is the milestone proving
the pattern works.

This phase is dominated by **prototyping and validation modes**.
The prototype is not the final system; it's the first end-to-end
check that the architecture is buildable.

Phase C ends when: a functioning end-to-end prototype demonstrates
the teacher-pupil dynamic on a constrained task, with measurements
captured. Assessment of that demonstration determines what cycles
follow.

After Phase C, the cycle structure continues — there is no Phase D.
Each subsequent cycle extends the prototype, refines the
specifications, or addresses gaps surfaced by validation. The
phase distinction stops mattering once the program is iterating
on a working artifact.

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
