---
id: Q-2026-04-30-013
type: question
slug: computational-complexity-prerequisites-first-problem
created: 2026-04-30
updated: 2026-04-30
status: open
confidence: 3
where: [T006]
tags: [computational-complexity, NP-completeness, first-problem, phase-1, deferred, mathematics-mode, construction-work]
provenance:
  - "conversation 2026-04-30 between @reh3376 and Claude during T006 Cycle 1 foundational-assumptions exercise"
  - "@reh3376's framing: 'I debated including an initial state that solves all N P completeness problems is the first step, but feared it would over complicate our ability to state these core assumptions clearly.'"
  - "@reh3376's clarification: 'When I say first training, I am not including the intial state assumption training for Bayes and advanced mathematical concepts. I mean the first problem we will attempt to solve once the framework and its pre-requisites are satisfied'"
  - "@reh3376's chosen handling: Option α (file as Q-artifact, defer to future cycles) over Option β (deliberate exclusion from D-NNN) or Option γ (sharpen and include in D-NNN)"
  - "D-2026-04-30-012 (foundational assumption) — does not address this question; deliberately defers to this Q-artifact"
  - "Q-2026-04-30-014 (3-D Tetris first-problem hypothesis, filed same PR) — Tetris's NP-completeness is what surfaced this question"
answer: ""
---

# Computational complexity prerequisites for the architecture's first applied problem

## The question

When the architecture transitions from **Phase 0** (substrate
initialization — Bayesian and advanced-mathematical foundations
established per D-2026-04-30-012 Component 3) to **Phase 1** (first
applied problem-solving), what does the choice of first problem
imply about the architecture's computational capabilities at that
transition moment?

Specifically: **if the architecture's first applied problem includes
NP-complete or NP-hard instances** (as 3-D Tetris would, per
`Q-2026-04-30-014`), what is required of the architecture at the
Phase 0 → Phase 1 boundary?

## Why this question is open and deferred

@reh3376 considered including a foundational claim about
NP-completeness in the foundational assumptions exercise but
deferred it: *"I debated including an initial state that solves all
N P completeness problems is the first step, but feared it would
over complicate our ability to state these core assumptions
clearly."*

The right framing was not yet available — three sub-readings exist
(below), each with different consequences, and committing
prematurely would either over-claim (Reading R1's bootstrap
problem; Reading R2's foundational-primitive claim) or under-claim
(Reading R3's theoretical-precondition framing might be too
abstract).

Filed under **Option α** (file as Q-artifact, defer): the question
is recorded so the consideration is not lost; sub-reading
selection is held open; resolution awaits more information.

## Three sub-readings, all held open

### Reading R1 — Bootstrap problem

The architecture's first applied problem requires solving NP-hard
instances, and without efficient methods for these, the
architecture **cannot make progress on first-problem work**. The
architecture has been initialized (Phase 0 complete); but the
moment it tries to solve Tetris-class problems, it stalls — its
substrate doesn't know how to navigate the combinatorial spaces
involved.

Consequence under R1: Phase 0 must include capabilities that handle
NP-hardness in some practically-useful sense (approximation
strategies, structural exploitation, search-tree pruning, or
similar). These capabilities must be in place before Phase 1
begins. The Phase 0 / Phase 1 boundary (G-clean per Q-014) is
gated by NP-hardness readiness.

This is closest to AIXI-family concerns where the optimal agent
has computational requirements that exceed any feasible
computation; bootstrap requires either approximation or
restriction to tractable problem classes.

### Reading R2 — Foundational primitive

The architecture must contain, at the Phase 0 → Phase 1 boundary,
some computational primitive **built into the seed-structure**
(item 1's seed metaphor) capable of solving NP-complete instances
in some useful class. The primitive is not learned through
input-driven growth; it is part of what the substrate provides.
The architecture's Phase 1 work uses this primitive as a
foundation.

Consequence under R2: D-2026-04-30-012 Component 3's "advanced
mathematics" includes specific computational machinery —
combinatorial-optimization primitives, perhaps SAT-solving
analogues, perhaps approximation-algorithm scaffolding — that is
part of what gets initialized in Phase 0. Phase 0 is more
substantively-loaded than its name suggests; it doesn't just
establish Bayesian foundations, it establishes computational
foundations including NP-hardness-handling primitives.

This reading is consistent with how biological systems handle
NP-hard tasks (the brain doesn't generally compute optimal
solutions; it uses heuristics, approximation, and structural
exploitation that are evolutionarily-installed primitives) — the
BNN as evidence (per D-012 Component 2) suggests R2 is
biologically-supported.

### Reading R3 — Theoretical precondition

The question is not about the architecture's *runtime* behavior
but about the **program's mathematical foundation**. Before the
architecture can be soundly specified, the program must address
what NP-completeness implies for the framework's claims. This is
a meta-claim about Phase A specification, not about Phase C
operation.

Consequence under R3: T006's mathematics mode must include
treatment of computational complexity questions. The
architectural specification produced at the end of Phase A
includes a position on NP-hardness that the architecture's later
construction will inherit.

This reading is more abstract than R1/R2; it's a research-strategy
claim rather than an architectural claim. R3 may complement R1 or
R2 (you address the theoretical foundations and *also* implement
some primitive or bootstrap), or stand alone.

## Why all sub-readings remain open

The three sub-readings have different relationships to the
foundational assumption (D-012):

- R1 implies Phase 0 must produce an architecture that can
  *perform* NP-hardness handling
- R2 implies Phase 0 must produce an architecture that *contains*
  NP-hardness primitives as substrate
- R3 implies Phase A must address NP-hardness as part of *theoretical*
  framework specification

These are not mutually exclusive but they impose different
demands. Choosing among them prematurely would:

- Under R1: commit Phase 0 to solving NP-hardness performantly
  before we know what classes of NP-hard problems the architecture
  will actually encounter
- Under R2: commit the foundational assumption to a specific
  computational primitive without grounds for choosing one
- Under R3: commit T006's mathematics mode to a substantial
  research workstream that may not be the right shape yet

@reh3376's deferral is methodologically correct: hold the
question; resolve when more information is available.

## When this question becomes load-bearing

Three triggers expected to surface the question for decision:

**Trigger 1 — Phase C construction work begins.** When the
architecture is being built and Phase 0 substrate work is
designed, the choice of which sub-reading guides design must be
made. Q-013 must be revisited and a sub-reading committed (or
multiple sub-readings combined).

**Trigger 2 — T006 mathematics mode opens.** When the program
enters the mathematics-mode work declared in T006's working
method, computational-complexity questions are part of what gets
formalized. Reading R3's framing is most likely to land in
mathematics-mode artifacts.

**Trigger 3 — Q-014's first-problem candidate is committed.**
If 3-D Tetris (or another NP-complete first-problem candidate) is
committed beyond C-leading status, the choice of first problem
forces a sub-reading choice. The first-problem commitment and
Q-013's sub-reading are tightly coupled.

## Cross-references

- **D-2026-04-30-012** (foundational assumption) — deliberately
  does not address NP-completeness; defers here
- **Q-2026-04-30-014** (3-D Tetris first-problem hypothesis) —
  the consideration that surfaced this question; tightly coupled
- **T006 plan items 1, 2, 3, 10** — the architectural
  commitments whose computational-complexity implications are
  what this question asks about
- **OBS-2026-04-28-004 §2.7** — substrate-level commitments;
  computational complexity is not among them but may need to
  surface as items 4-9 or beyond are sharpened

## What this question is not

- **Not a P vs. NP claim.** The program is not committed to
  P=NP. R1, R2, R3 all have formulations that don't require
  P=NP (R1 via approximation, R2 via heuristic primitives, R3
  via theoretical accommodation rather than resolution).
- **Not a question about all NP-complete problems.** The
  question is scoped to **the first applied problem the
  architecture encounters** — possibly 3-D Tetris, possibly
  another candidate. What's required is sufficient capability
  for that specific problem class, not arbitrary NP-completeness.
- **Not a Phase A specification gate.** Phase A can complete
  with this question deferred. Construction work (Phase C) is
  where the resolution becomes necessary; Phase A produces an
  architectural specification that Phase C inherits including
  this open question.

## Status updates expected

Revisited when:

- Phase C construction work begins (most likely trigger)
- T006 mathematics mode opens
- Q-014's first-problem candidate is committed beyond C-leading
- Items 4-9 of T006's plan are sharpened and surface
  computational-complexity implications
- Broader AI research community produces results bearing on the
  question (e.g., new approximation algorithms, new structural
  results on the problem classes the architecture will encounter)
