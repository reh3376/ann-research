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
