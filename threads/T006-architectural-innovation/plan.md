---
id: T006-plan
type: plan
slug: t006-plan-of-record
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 4
where: [T006]
tags: [plan, cycle, architecture, what-it-must-have]
provenance:
  - "T006 thread frame (working method, four work modes, relationship to other threads)"
  - "@reh3376's selection of focus 2B in conversation 2026-04-28"
  - "OBS-2026-04-28-004 §2.7 (four substrate-level commitments) as anchor for the positive-space work"
  - "T004's five facets (Q-005 through Q-009) as additional positive-space inputs"
owner: "@reh3376 (decides what the architecture must have), Claude (sharpens, drafts, researches, surfaces consequences)"
current_question: "What does the architecture must have, and how do we sharpen each must-have from a stated requirement to a precise specification?"
---

# T006 Plan of Record

## How to read this document

This plan is a living artifact. It is revised at every cycle gate
and may be revised mid-cycle when the work demands it. The latest
version is the plan-of-record; prior versions are recoverable via
`git log` on this file.

The plan is a **shared artifact between Roger and Claude**:

- **Roger's content:** strategic phasing, focus selection, what
  goes in positive vs negative space, which candidates survive,
  which questions matter.
- **Claude's content:** plan structure, cycle bookkeeping,
  surfaced inconsistencies and gaps, drafted artifacts.

Disagreement is expected and useful. When Claude pushes back on
something Roger has written, the plan records the pushback as a
question, not as a contradiction.

---

## Strategic phasing

The program operates in three phases. Phase boundaries are
gradients, not gates — work in a later phase may begin before the
prior phase formally completes when the work is ready.

### Phase A — Architectural Innovation Development

**Goal.** Identify, articulate, sharpen, and specify the
architectural innovation. The deliverable is **a specification**
sufficient that people with the resources to build it can
recognize what it is and decide to back it.

**Lead work.** T006 cycles producing artifacts under the four
work modes (what it is not, what it must have, what it could be,
the mathematics). T004 continues surfacing facets that inform
T006. T001/T002 run when T006 needs their evidence.

**Phase A is where the program currently is.** All other phases
are downstream.

**Phase A gate criteria** (when to consider Phase A complete):
- A specification exists that names the architectural innovation
  precisely
- The specification covers all four work modes (what it is not,
  what it must have, what it could be, the mathematics)
- The specification has internal consistency (no contradictions
  between modes)
- The specification has been pressure-tested against the most
  obvious counter-arguments (does it survive the simplest
  objection?)
- Roger judges the specification ready for external consumption

The criteria are not a checklist; they are a description of when
the work is done. Roger's judgment is the binding criterion.

### Phase B — Communication and Resource Acquisition

**Goal.** Articulate the Phase A specification in forms reaching
the audiences that can fund and staff implementation. Papers,
talks, pitch documents, perhaps a working prototype that
demonstrates the core insight.

**Lead work.** Drafting communication artifacts in Roger's voice;
identifying audiences; tracking responses; iterating on what
lands and what doesn't.

**Phase B does not start until Phase A is materially complete.**
Premature Phase B work undermines Phase A's integrity (a draft
paper makes the architecture feel finished before it is) and
wastes effort on communicating a half-specification.

### Phase C — Build with Team and Funding

**Goal.** Implement the architecture, with the team and funding
Phase B has secured. Out of scope for this thread; will require
a separate program structure when reached.

---

## Active cycle

### Cycle 1 — What it must have

**Mode sequence:** plan → research → specification → plan → assess → unblock → plan

**Status:** plan-of-record drafted; awaiting research-mode start.

**Focus.** Populate the positive space of substrate properties
the architecture must support. For each must-have, sharpen from
a stated requirement to a precise specification: what exactly
must be true, what is the requirement's source/rationale, what
would falsify it, what existing architectures fail to provide
it and why.

**Why this focus first.** The four substrate-level commitments
from OBS-2026-04-28-004 §2.7 are already named at high level —
trainable representations, generative modeling/prediction,
reference-frame construction, meta-learning. T004's five facets
(Q-005 through Q-009) add another set. The positive-space work
takes these named requirements and sharpens them into
specifications. This work is high-leverage: every later cycle
of T006 (whether it is producing candidate architectures, doing
mathematics, or working negative space) will reference what was
established here.

**Why not start with negative space (2A).** Considered. Negative
space and positive space inform each other; the choice was
between A, B, or both in parallel (C). Roger selected B alone.
The reasoning (implicit but consistent with the working method):
positive space gives the architecture's *target capabilities*,
which determine what counts as relevant negative space. Without
positive space, "what it is not" can become an unbounded list
(it is not a transformer, not a CNN, not an RNN, not... — true
but not useful). With positive space established, negative space
becomes a focused set of "things that *appear* to satisfy these
must-haves but actually don't."

**Working set of must-haves under triage** (item 1 sharpened in
Cycle 1's first triage pass; items 2-9 pending triage):

### Item 1 — Continuous learning via input-driven, prediction-error-driven topology growth

**Status:** sharpened (2026-04-28). Property statement, current
cycle. Logical grounding deferred to subsequent cycles. Mathematics
deferred further.

The architecture is a framework for intelligence whose
meta-structure is general, repetitive, and input-invariant —
like the cortical column, which is a general learning structure
that builds reference frames regardless of input modality.
Structural invariance enables compositional generalization
(Hawkins).

The framework is the seed. Data is the soil. Input wirings are
the mechanism of growth.

*Input wirings* are the comprehensive input-side specification of
the architecture. They include the routing of input streams to
seed-regions, the connection density and strength priors within
and between regions, and any other aspect of how input gets into
the framework. Two architectures may share seed-structure and
data while differing in input wirings; the wirings are part of
what individuates a trained instance.

Growth means: nodes and edges are created (in the graph-theoretic
sense, not committed to any particular database technology),
reference frames are constructed, predictions about current and
future state are produced. When predictions match observation,
connections strengthen. When predictions diverge from observation,
surprise triggers new learning — new nodes, new edges, modified
weights, possibly modified mechanisms.

The architecture is **deterministic given full specification**:
two instances with identical seed-structure, identical data, and
identical input wirings produce identical outputs. The
biconditional is strict — divergence in any of the three produces
divergent outputs. The architecture is therefore **path-dependent
on input wirings**: two trained instances are individuated by their
seed + data + input-wiring history. Reproducibility requires all
three.

Continuous learning is **necessary** (not sufficient) for
intelligence. Without it, systems differ only in the timescale of
their frozenness, not in kind. The full set of substrate-level
commitments together specify what the architecture *is*; this
item specifies the property that makes the architecture *capable
of being intelligence-supporting*.

The mechanism owned by this item is *prediction-error-driven
growth*. The substrate property of being a generative model is
owned by item 2; the two items co-specify the prediction-and-
surprise dynamic.

Frozen embeddings (the MDEMG ceiling), fine-tuned adapters that
update on a discrete schedule, and any architecture whose topology
is fixed at design time do not satisfy this commitment.

**Open within item 1:**

- **Q-T006-B:** Is growth append-only, or does it include pruning
  / forgetting / consolidation?
- **Q-T006-C:** Is "node and edge creation" neurogenesis-
  equivalent or capacity-recruitment-equivalent at the
  implementation level?

### Items 2-9 — Pending triage

Items 2-9 below remain placeholders awaiting triage. Each will
move through the same workflow item 1 just did: Roger sharpens;
Claude pushes back; the item gets revised, absorbed, dropped,
split, or deferred; the result is committed. Triage proceeds
one item at a time.

The four from OBS-004 §2.7:

2. **Generative modeling / prediction.** The substrate must
   produce predictions about current state and future state. A
   storage-and-retrieval substrate (the MDEMG ceiling) is not
   sufficient. *(Co-specifies prediction-and-surprise dynamic
   with item 1; substrate-property side here, mechanism in item
   1.)*
3. **Reference-frame construction.** The substrate must let the
   system carve representational space into frames that support
   compositional reasoning. Inherited frames from external
   structure (the MDEMG ceiling) are not sufficient.
4. **Meta-learning / mechanism modification.** The substrate
   must let the system modify its own learning rules over time.
   Parameter-only modification within fixed mechanisms (the
   RSIC scope) is not sufficient.

The five from T004 (Q-005 through Q-009):

5. **Homeostatic boundary** (Q-005). The system must maintain a
   distinction between self and environment that is not
   externally specified. The boundary is the substrate of
   identity.
6. **Dimensional minimum-commitment** (Q-006). The system's
   representational dimensionality is not pre-specified; it
   commits only to what the data demands. *(Tentatively absorbed
   into item 1 pending triage of item 6 — the input-driven
   growth process determines dimensionality, so this may not be
   independent.)*
7. **Projection-and-anchoring** (Q-007). Projection of current
   state forward and backward in time, anchored to invariants
   the system has identified.
8. **Recursive predictive horizon** (Q-008). The system's
   predictive horizon expands recursively — local predictions
   feed into longer-horizon predictions, which inform local
   prediction.
9. **Prime directives** (Q-009). The system's terminal goals
   (whatever they are) are themselves represented in a way that
   the system can reason about, not implicit in the
   architecture.

These are placeholders for Roger to revise, sharpen, prune,
expand, or replace. The list is the cycle's starting input,
not its conclusion.

**Cycle 1 deliverable.** An observation artifact that captures
the sharpened must-haves — for each, a precise specification of
what must be true, the rationale, what would falsify the
requirement, and (where useful) what current architectures fail
to provide it.

**Cycle 1 explicit non-goals:**
- Does not propose candidate architectures (that is "what it
  could be," a different mode)
- Does not exhaustively cover negative space (that is "what it
  is not," a different mode)
- Does not develop mathematical proofs (that is "the
  mathematics," a different mode)
- Does not produce communication artifacts (Phase B work)

**Cycle 1 gate criteria:**
- The must-have list reflects Roger's current thinking, with
  reasoning recorded
- Each must-have is sharpened beyond a one-line restatement of
  the original requirement
- Internal consistency: no must-have contradicts another
- Gaps and tensions are surfaced as questions, not papered over
- Roger judges the cycle's output as a usable foundation for
  Cycle 2

**Methodological commitment** (settled 2026-04-28 during item 1
triage): logical grounding for each must-have is the *direction*
of the work, not a Cycle 1 gate criterion. Each must-have starts
as a sharpened property statement in this cycle; logical grounding
is earned over multiple cycle iterations; mathematics mode comes
later. **Phase A is multi-cycle by construction.** Cycle 1
produces the inventory in sharpened-property-statement form;
subsequent cycles add depth.

**Likely Cycle 2 candidates** (do not commit yet — depends on
Cycle 1's outcome):
- Cycle 2a: "What it is not." Negative space, scoped by Cycle
  1's positive space.
- Cycle 2b: Continue "what it must have" if Cycle 1 surfaced
  more depth than fits one cycle.
- Cycle 2c: "What it could be." Begin candidate-architecture
  work, with Cycle 1's must-haves as filters.
- Cycle 2d: "The mathematics." Begin formal grounding for
  specific must-haves where mathematical structure is already
  visible.

The Cycle 2 choice is Roger's at unblock.

---

## Strategic phasing of remaining threads

T002 (MDEMG trace) cycles 2A-2F as previously named (J17 deep
trace, Jiminy+RSIC combined trace, MEMG deep trace, file
surfaced questions, T004 metacognition surfacing, continue T001)
are **subordinated to T006 in Phase A**. They are no longer
front-of-program; they run when T006's positive-space or
mathematical work needs their evidence.

The 10 surfaced questions from OBS-003 and OBS-004 remain queued
for filing. Filing them is light-touch work that can run between
T006 cycles when convenient, or deferred until they become
load-bearing for a specific T006 cycle. Roger directs filing
timing; not on T006's critical path.

D-010's framing as "the architecture" may need revision once
T006 produces a specification. The hybrid superstructure pattern
in D-010 may be a *deployment shape* for the architectural
innovation rather than the innovation itself. This is a question
T006's specification work may resolve. Until then, D-010 stands.

D-011's V-JEPA commitment as visual faculty depends on the
eventual architecture. If T006's specification implies a
substrate where V-JEPA is the right faculty, D-011 stands. If
not, D-011 may be revised. Until then, D-011 stands.

---

## Risks to the plan

**R1: Claude proposes candidate architectures despite the
working-method declaration.** Most likely failure mode. The
temptation is high — under pressure, drafting a "candidate"
feels productive. The mitigation is the working-method
declaration in T006's thread frame and the explicit non-goal
in Cycle 1 ("does not propose candidate architectures"). When
the temptation appears, the response is "I notice I am tempted
to propose; what is the question I should ask Roger instead?"

**R2: The must-haves become aspirational rather than
specifications.** "The architecture must support continuous
representation learning" is a sentence; "continuous representation
learning" sharpened to a specification is what the cycle
produces. The sharpening is the work. Surfacing this as a risk
because the temptation is to call the inventory list "the
specification" when it is only the input.

**R3: The cycle takes longer than expected.** Architectural
specification work is genuinely hard. Cycle 1 may need multiple
revisions. The plan accommodates this — cycle 2 may be "continue
must-have work" if cycle 1 is incomplete. The risk is Roger
losing patience with the slow pace and shifting to faster modes
(candidate architectures, communication artifacts) before the
substrate is ready. Mitigation: the working method explicitly
says fast paths fail; the thread frame holds the line.

**R4: T006's cycles drift away from the program objectives in
D-009.** Possible if must-haves are pursued for their intrinsic
interest rather than for their bearing on the long-term goal.
Mitigation: every must-have specification answers "why does
this matter for AGI/world-models/reference-frames/etc." as part
of its rationale.

---

## Log

### 2026-04-28 — Thread opened with focus 2B

T006 opened with the working-method declaration codified. Cycle 1
focus selected as 2B ("what it must have") — populating positive
space using the four §2.7 substrate commitments and the five T004
facets as initial input.

Plan-of-record drafted. Awaiting Roger's research-mode start.
