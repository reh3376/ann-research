# Plan of Record

> **Living document.** Revised at every cycle gate. Git history
> preserves evolution; the current state is always the working plan.
>
> **Last revised:** 2026-04-28 (initial draft)
> **Next gate:** completion of Cycle 1 → unblock → plan revision

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

**Trace work is asymmetric.** T001 (Monty) has produced two
mechanism observations (vote topology, matching_step pipeline). T002
(MDEMG) has not started. This asymmetry is the binding constraint
on architectural specification work because D-010's Jiminy precedent
and the AI2AI protocol both depend on T002 evidence we don't yet
have.

**Cycle structure is committed.** T005 codifies how planning and
work proceed: plan / research / specification / plan / prototyping /
validation / assess / unblock / plan, iterative, with explicit gates
between modes. This document is the plan-of-record under that
structure.

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

## Active cycle: Cycle 1

### Focus

**Bring T002 (MDEMG trace) to first-observation status.** Specifically: produce
a mechanism-level observation comparable to OBS-2026-04-28-002
(Monty's matching_step pipeline) covering MDEMG's primary cognitive
loop — what an ingest cycle does mechanistically, end to end,
when the system is processing input.

### Why this cycle

T002 has been on the open-work list since project bootstrap and
hasn't started. This is the binding constraint on architectural
work for two reasons:

1. **D-010's Jiminy precedent depends on T002 evidence.** The
   teacher-pupil pattern is named as inherited from MDEMG's
   existing Jiminy/coding-agents structure. We've committed to
   that pattern as architectural foundation but have no
   trace-level evidence of how it actually operates. The
   specification work in Phase B cannot proceed responsibly
   until we have that evidence.
2. **The AI2AI protocol is committed as architectural component
   without specification.** D-010 names MDEMG's existing AI2AI
   protocol as the leading reference. Until T002 produces a clear
   picture of what AI2AI carries, how it's structured, and what
   its compression strategy is, the protocol design for the
   hybrid architecture cannot begin.

### Mode sequence for Cycle 1

This is a **research cycle**. The mode sequence is:

```
plan → research → assess → unblock → plan
```

Specification, prototyping, and validation are not part of this
cycle. The cycle's job is to produce mechanism-level observations
that subsequent cycles can specify and prototype against.

### Cycle 1 method

This cycle requires asymmetric collaboration that previous T001
work did not. Specifically:

**@reh3376 leads.** MDEMG is your operational substrate. You have
end-to-end context I lack. The cycle proceeds by you walking
through MDEMG's primary loop — the actual code paths, the actual
data flows, what's load-bearing vs. what's incidental — at a level
of mechanism detail.

**Claude drafts and surfaces.** As you describe, I capture into
artifacts. I surface questions where your description leaves
ambiguity. I produce observation drafts you review. The
trace-via-conversation mode is heavier collaboration than the
trace-via-reading mode T001 has used.

**Reading happens in support, not in lead.** I can read MDEMG
source code on the working clone (`/home/claude/mdemg`) to
ground or verify what you describe. Reading does not lead the
cycle; your operational walk-through does.

### Gate criteria

**research → assess.** Cycle 1's research mode is complete when:

- One observation artifact captures the MDEMG primary loop at
  mechanism-level depth, comparable to OBS-2026-04-28-002 in
  detail and confidence calibration
- One observation artifact captures the AI2AI protocol's
  current state (what it carries, how, what it does not carry
  well)
- At least 3 questions are surfaced about MDEMG mechanisms
  whose understanding requires further work
- Cross-references to D-010, D-011, and T004 questions are
  in place where evidence bears

**assess → unblock.** Assessment for Cycle 1 asks:

- Did the trace work produce evidence at the level we expected?
  (validation question — execution)
- Was the mechanism-level abstraction the right depth, or do we
  need to go shallower or deeper for what comes next? (assess
  question — design)
- Where did the trace surface unexpected things — gaps the
  architecture must close, or capabilities the architecture can
  inherit?

**unblock.** Unblocking Cycle 1 produces candidate Cycle 2
focuses. Likely candidates (we'll know after Cycle 1 actually runs):

- Cycle 2A: continue T002 — read deeper into Hebbian learning
  rules, hidden concept abstraction, Jiminy intervention logic
- Cycle 2B: pivot to specification — the AI2AI protocol message
  schema, given the trace evidence
- Cycle 2C: pivot to T001 — finish reading Monty's LM internals
  (the geometry/dimensionality concrete questions) since
  parallel evidence streams help T003 synthesis
- Cycle 2D: pivot to T004 first-principles continuation — the
  metacognition material @reh3376 has previewed

The unblock decision happens at the close of Cycle 1, informed by
what Cycle 1 actually surfaces. The candidates above are the
*current* projection of likely options, not commitments.

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
