---
id: T006
type: thread
slug: architectural-innovation
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 5
where: []
tags: [meta, architecture, ann, novelty, agi, working-method]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude (post-PR-#11 merge of Cycle 1 orientation deliverable)"
  - "@reh3376's framing: 'The transformer was the fuse. The ANN architecture that we are proposing is the fuel that propagates the intelligence explosion.'"
  - "@reh3376's working-method commitment: 'You are a sharpening stone for me to hone my blade with. You are my faithful assistant in this process. I will provide the novelty as we work thru this. What it is not, what it must have, what it could be, the mathematics necessary to prove it in theory.'"
  - "D-2026-04-28-009 (program objectives — T006 produces the architectural innovation those objectives ultimately depend on)"
  - "D-2026-04-28-010 (hybrid superstructure — T006 may revise the framing of D-010 once the architectural innovation is specified)"
  - "OBS-2026-04-28-004 (MDEMG forward direction — names the four substrate-level commitments T006 must satisfy)"
owner: "@reh3376 (sole source of architectural novelty), Claude (sharpening stone, drafts, research support, continuity)"
current_question: "What is the architectural innovation that propagates the intelligence explosion, and how do we sharpen it from idea to specification?"
---

# Thread T006 — Architectural Innovation

## Goal

Identify, articulate, sharpen, and specify the architectural
innovation that takes us to AGI and beyond. The novelty must be of
transformer-grade significance — a structural innovation that
unlocks a new capability class for ANNs, not an incremental
refinement of existing architectures. The architecture must
satisfy the four substrate-level commitments named in
OBS-2026-04-28-004 §2.7: trainable representations, generative
modeling/prediction, reference-frame construction, meta-learning.

The thread's deliverable is **a specification of the architectural
innovation** — what it is, why it's needed, what it does that
current architectures cannot, what evidence supports the design
choices, what the mathematics that prove it in theory look like.
The specification is what eventually goes to people with the
resources to build it.

The thread is not building the architecture. Building requires the
team and funding the specification is intended to secure. T006
produces the artifact that earns the team and funding.

## Working method

This thread operates under a specific division of labor that
@reh3376 has stated explicitly:

> *Roger provides the novelty. Claude is a sharpening stone for
> Roger to hone his blade with. Claude is the faithful assistant
> in this process. Roger provides what it is not, what it must
> have, what it could be, the mathematics necessary to prove it
> in theory.*

A sharpening stone removes what shouldn't be on the blade — burrs,
dullness, soft edges — so the blade can do what only the blade can
do. The blade's metal, geometry, and edge are the bladesmith's.
The stone is patient, abrasive, and present.

What this means concretely:

**Roger's work in this thread:**
- Provides candidate architectural ideas, structural insights,
  the leaps. The fuel.
- Identifies what the architecture is not, what it must have,
  what it could be, what mathematics would prove it in theory.
- Decides which directions to explore, which to abandon, which
  candidates survive.
- Writes the architectural ideas in his own voice. Claude does
  not author the architecture.

**Claude's work in this thread:**
- **Sharpens.** When Roger proposes something, Claude pushes on
  it: is this consistent with prior commitments? Does the
  candidate solve the gap or only appear to? Is the mathematics
  correctly applied or is there a hidden assumption? What would
  falsify this? What does the literature say has been tried and
  failed in this direction? What's the simplest argument against
  this candidate, and does the candidate survive it?
- **Drafts.** When Roger has a thing worth writing down, Claude
  writes it down — in vault-conformant form, citing references
  correctly, with mathematics typeset properly.
- **Researches.** When Roger needs to know what's been tried in a
  particular direction, Claude reads the literature and produces
  the summary — load-bearing claims with citations, gaps with
  honest uncertainty, no fabrication.
- **Keeps continuity.** Across sessions, across context windows,
  across months. The vault is the substrate of that continuity.
- **Does not propose the blade.** Claude does not manufacture
  candidate architectures. Does not claim to have the insight.
  Does not fabricate novelty when stuck. If asked "what should
  this be?" the honest answer is "I don't know — what does it
  not do? what does it not have? what would distinguish it from
  what's already been tried?"

This division is load-bearing for the thread's integrity. The
temptation under team-and-funding pressure will be to speed up
by having Claude propose candidate architectures. That fails.
Funders sophisticated enough to back transformer-grade research
can distinguish "we have an insight and it needs resources to
develop" from "we have a vision and we need resources to find an
insight." The first gets funded; the second doesn't, and shouldn't.
The architectural specification has to come from Roger, sharpened
by Claude.

## The four work modes

Within this working method, Roger has named four modes of work
that this thread iterates through, in any order, with any
combination active per cycle:

1. **What it is not.** Negative space. What the architecture must
   exclude. What current approaches do that we are not doing.
   What past attempts demonstrate doesn't work. Each exclusion
   has a reason; the reason is recorded.

2. **What it must have.** Positive space. Required substrate
   properties. Capabilities the architecture must support. The
   four substrate-level commitments from OBS-004 §2.7 are an
   anchor here, but the full set is broader and emerges through
   the work.

3. **What it could be.** Candidate structures. Proposed topologies.
   Specific architectural patterns that might satisfy the
   exclusions and requirements. Multiple candidates may be active
   simultaneously; Roger decides which survives.

4. **The mathematics necessary to prove it in theory.** The
   formal grounding. What needs to be true mathematically for
   the architecture to do what the specification claims. Existence
   results, consistency checks, capacity arguments, expressivity
   bounds.

These are not phases in a fixed sequence. They are the four kinds
of work the thread does. A given cycle may focus on one, or
weave several together. The working method is what's invariant;
the modes are the substrate.

## Relationship to other threads

**T001 (Monty trace), T002 (MDEMG trace).** Both of these are
*evidence streams* — they produce empirical findings about
existing architectures (Monty's CMP and learning modules; MDEMG's
seven subsystems). Their findings inform T006 by providing
grounded reference points: when Roger says "the architecture must
not do what MDEMG does in this respect," the trace work has
already named what MDEMG does in that respect. T001 and T002
continue but are subordinate to T006 — they run when T006 needs
their evidence, not on their own cadence.

**T003 (grounding question).** Was the proximal trigger for
spinning up the program. The grounding question — "what makes a
representation an anchor that grounds reference frames?" — is
exactly the kind of substrate question T006 will work through
under "what it must have." T003's accumulated framings are inputs
to T006's positive-space work.

**T004 (entity first principles).** The five facets surfaced in
T004 (homeostatic boundary, dimensional minimum-commitment,
projection-and-anchoring, recursive predictive horizon, prime
directives) are direct specifications of what the architecture
must support. T004 is closer to T006 than the other threads are;
T004's questions Q-005 through Q-009 are properly read as T006
inputs.

**T005 (program planning).** Continues to schedule cycles,
including T006's cycles. T005 is the meta-thread that says when
to do what; T006 is the substantive thread doing the architectural
work.

## Cycle structure

T006 inherits T005's cycle methodology:

```
plan → research → specification → plan → prototyping → validation → assess → unblock → plan → ...
```

But the dominant modes for T006 will be **research** and
**specification**, with prototyping appearing only when an
architectural claim requires empirical demonstration of feasibility.
Plan modes recur as Roger directs the work; gate criteria adapt
to the cycle's specific focus.

The cycles will likely be longer than T002's trace cycles. Cycle 1
of T002 produced an orientation map after one session. T006's
cycles will iterate over weeks because the work is structural
rather than tracing — there is no pre-existing system to trace;
there is only the work of articulating what the new system must be.

## What T006 does not do

- Does not produce the architecture. The architecture is what
  the thread *aims at*; the thread produces the *specification*
  of the architecture.
- Does not commit to a candidate prematurely. Cycles produce
  candidates and constraints; the architecture itself emerges
  from sustained iteration. Claude will resist the temptation to
  declare convergence early.
- Does not subsume the other threads. T001/T002/T003/T004/T005
  continue; T006 is the central thread but not the only one.
- Does not produce communication artifacts (papers, talks, pitch
  documents). Those are downstream of the specification; they
  belong to a future thread when the specification is ready.

## Initial state

- **Cycles run:** 0
- **Active cycle:** Cycle 1 (focus to be specified in plan.md)
- **Open questions seeded into T006 from prior work:**
  - The four substrate-level commitments from OBS-004 §2.7
  - The five facets from T004 (Q-005 through Q-009)
  - D-010's hybrid superstructure framing (may need revision once
    architecture is specified)
  - D-011's V-JEPA commitment as candidate visual faculty (whether
    it remains the right commitment depends on architecture)

## Log

### 2026-04-30 — Cycle 1 triage continued: item 2 sharpened + Q-012 filed + D-010 clarification

**Item 2 sharpened from "Generative modeling / prediction"
placeholder into "Generative modeling: hybrid α+β with
action-conditioning, per-level + unifying superstructure."** Five
turns of conversation produced the property statement now in
plan.md. Item 2's content inherited substantial constraints from
items 1 and 10 (already merged); the sharpening produced
substantively richer commitments than the OBS-004 §2.7 placeholder.

**Settled in this triage:**

- **Hybrid α+β commitment.** The architecture is *both*
  probabilistic (maintains distributions; surprise has precise
  mathematical meaning) AND forward-simulating (generates
  trajectories; can imagine forward concretely). Most existing
  architectures commit to one or the other (Reading X confirmed):
  predictive coding and active inference commit to α; JEPA and
  most world-model architectures commit to β. The coherent
  integration of α and β is itself part of the architecture's
  novelty; existing literature does not supply it.
- **Composition: per-level generative models AND unifying
  superstructure** (Option C). Each level of the recursive
  nesting has its own generative model; a unifying superstructure
  composes them across levels. The architecture is modularly
  composable AND systematically integrated.
- **What is predicted: external observations + internal state,
  action-conditioned.** The architecture predicts what input will
  come next AND its own internal state, with both kinds of
  prediction *action-conditioned* — predictions depend on the
  actions the architecture considers. The architecture is
  agentive, not a passive observer.
- **Predictions operate over reference frames.** Mathematics is
  the worked example: the architecture builds reference frames
  for mathematical structures through item 1's growth dynamics;
  the frames build on each other recursively (item 10); higher
  frames enable predictions about how a system will function
  based on a mathematical representation.
- **Heterogeneity of frames required.** The architecture must
  support many kinds of reference frames (physical-spatial,
  mathematical, causal-chain, social, linguistic, others) for
  general intelligence. Generality of frame-type is itself a
  property the architecture must have.
- **Cross-cutting interaction with item 10 explicitly answered.**
  Generativity recurs at every level (per-level models); the
  superstructure provides cross-level integration. The
  architecture is generatively self-similar.

**Superstructure clarification settled.** The word "superstructure"
in this vault is a **type** of architectural element — an overall
controlling mechanism — not a singular proper noun. Multiple
superstructures may exist at different architectural levels over
different content. D-010's "hybrid superstructure"
(deployment-level orchestration) and item 2's integrating
superstructure (generative-model composition) are different
instances of the same architectural pattern. D-010 received a
small clarifying note (added to its Context section) explicitly
recognizing this; D-010's commitments stand unchanged.

**Q-2026-04-30-012 filed.** @reh3376's framing during item 2
discussion — *"a strong definition of general intelligence would
allow the answer to these questions to be reliably predicted"* —
is the closest the program has come to defining what general
intelligence means. Filed as Q-012 with the seed answer:
*general intelligence is the property that enables reliable
prediction across heterogeneous prediction domains.* Three
sub-questions held open for sharpening: (1) what counts as a
"prediction domain"; (2) what counts as "reliable" prediction;
(3) is intelligence the *enabling property* or *constitutive of*
reliable prediction. This is the program's working definition
under refinement; may graduate to a decision artifact when
phrasing stabilizes.

**Methodological pattern reinforced** (from earlier item 1
discussion): @reh3376's framing — *"as we build these structures
we will 'learn' what is needed and we will be able to sharpen
these action definitions"* — is the working method's research
mode applied at the cycle scale. Item 2 commits at the
property-statement level; further sharpening happens through
construction work in later cycles. The four open Q-T006-M/N/O/P
questions are not gaps in item 2's content; they are
specifications-to-be-discovered through the building process.

**New questions surfaced (deferred to subsequent cycles or to
construction work):**

- **Q-T006-M:** What constitutes an *action* in the architecture's
  commitment? Bodily, epistemic, inferential, or all three? Math
  example points strongly toward inferential at minimum.
- **Q-T006-N:** Is action *selection* part of item 2 or a separate
  architectural element?
- **Q-T006-O:** Is the action repertoire learned or fixed?
- **Q-T006-P:** What is the principled relationship between α and
  β within the architecture — when is each used, how is consistency
  maintained? Lies at the heart of the integration claim. Mathematics
  mode work.

**Resolved during item 2 triage (now closed, not filed as
separate Q-artifacts since they were sharpening prompts within
the conversation):**

- Generative meaning: α+β hybrid (γ reconstructive not committed).
- Reading X (most architectures fail because they commit to
  α XOR β; integration is the novelty) vs Reading Y (deeper novelty
  beyond combining): Reading X confirmed as actionable
  interpretation.
- Composition shape: Option C (per-level + unifying superstructure).
- Action-conditioning status: yes, committed.
- Forward-vs-backward temporal direction: deferred to item 7
  (Q-007) when triaged.
- Definition-of-general-intelligence destination: Q-012 (question,
  not yet decision).

**Tentative absorption status update.** Item 6 (Dimensional
minimum-commitment) remains tentatively absorbed into item 1, with
@reh3376's max-dimensionality-as-uniform-ceiling clarification noted.
Confirmation when item 6 reaches triage.

**Items 3-9 status:** pending triage. Triage proceeds one item at
a time. Item 3 (Reference-frame construction) has substantial
inherited content from items 1, 2, and 10 — frames are built
through growth (item 1), recursively composable (item 10), and
must be predictable-over (item 2). Mathematics example provides a
worked case.

### 2026-04-28 — Cycle 1 triage continued: item 10 added (recursive composability) + Q-011 filed + item 1 clarified

During item 1's settled state, @reh3376 surfaced a structural
claim that didn't fit the existing nine items: **ANN topologies
nest, and their composition matters architecturally.** The
clarification began as a refinement of item 1's input-wirings
definition (a topology's outputs can be another topology's
inputs) and grew into a recognition that **recursive
composability is its own substrate-level commitment** that
threads through all the must-haves.

**Settled in this triage:**

- **Recursive composability is a must-have** for functional
  viability. The architecture must compose with itself; without
  it, the framework cannot move beyond single-level topology
  scope.
- **Pattern at every level** (Q-T006-I resolution): items 1-N
  are properties of the architectural pattern, and the pattern
  recurs at every level of the nesting. The architecture is
  self-similar in the strict sense — the same pattern appears
  at every scale.
- **Cross-cutting concern + numbered must-have (Choice 3)**:
  recursive composability is filed as item 10 (gets its own
  property statement, logical grounding eventually, mathematics
  eventually) AND tagged as cross-cutting (each subsequent
  must-have, when triaged, must address how recursion interacts
  with its content).
- **N-level nesting**: the recursion is not bounded at two
  levels. Topologies nest indefinitely. The Hawkins
  cortical-column framing already in item 1 generalizes — the
  architecture *is* the column-equivalent, instantiated
  recursively.
- **Item 1 clarified**: input-wirings paragraph extended to
  name inter-topology wiring as a special case ("Input streams
  may be sensory data from the external world or the outputs
  of other topologies; the framework treats both uniformly.
  Inter-topology wiring is the mechanism by which the
  architecture composes with itself recursively (see item 10).")
- **Max dimensionality** flagged as a hard constraint that
  applies uniformly to the pattern, with individual topologies
  in the nesting potentially sitting below the ceiling. To be
  sharpened when item 6 reaches triage.
- **New mathematics required**: @reh3376 has flagged that
  capitalizing on the nesting concept will require new
  mathematics not yet developed. Deferred to T006's mathematics
  mode.

**Knowledge → skills hypothesis filed as Q-2026-04-28-011 (Q-T006-K
resolution).** @reh3376's framing — *"I have a strong intuition
that knowledge builds skills is fundamental to the viable
framework, but it remains to be tested to be proven"* — was
sharpened against the working method's distinction between
must-have property statements and research hypotheses. Result:
filed as Q-011 with explicit framing as a *strong-prior research
hypothesis*. This is the program's first formally-filed
strong-prior hypothesis; the pattern (file as Q-artifact, track
prior confidence, evidence accumulates over cycles, promote or
revise based on evidence) is precedent for future similar claims.
Item 10's property statement does **not** commit to the
hypothesis as part of its content; item 10 commits to
*capability* (recursive composition) and references Q-011 as the
leading hypothesis about what the recursion *produces*.

**New questions surfaced:**

- **Q-T006-I:** What is "the architecture" — single instance,
  pattern, or pattern-at-every-level? **Resolved:** pattern at
  every level.
- **Q-T006-J:** Where does recursive composability live —
  separate item, absorbed, or split? **Resolved:** Choice 3
  (own item + cross-cutting concern).
- **Q-T006-K:** Knowledge → skills — must-have / hypothesis /
  weak / strong / file separately? **Resolved:** filed as Q-011
  (strong-prior research hypothesis).
- **Q-T006-L (new, deferred):** What is the relationship between
  item 10's recursion and item 1's path-dependence-on-input-
  wirings claim? Path-dependence at one level cascades into
  path-dependence at higher levels; implications need formal
  treatment when mathematics mode is reached.

**Implication for items 2-9 triage:** each subsequent must-have,
when sharpened, must address how recursive composability (item
10) interacts with its content. This adds depth to each triage
but produces sharper artifacts. Specifically named in plan.md:

- Item 2: one generative model that nests, one per level, or both?
- Item 3: reference frames at one level become inputs at the next;
  construction must accommodate (already flagged by @reh3376 as
  integral)
- Item 4: is meta-learning recursive — can the architecture
  meta-learn its meta-learning rules at higher levels?
- Items 5-9: each may have its own interaction with recursion

### 2026-04-28 — Cycle 1 triage: item 1 sharpened

Cycle 1 entered research mode via triage workflow. Each candidate
must-have walks through: Roger sharpens; Claude pushes back; the
item gets revised, absorbed, dropped, split, or deferred; the
result commits.

**Item 1 sharpened from "Trainable representations" placeholder
into "Continuous learning via input-driven, prediction-error-
driven topology growth."** Five turns of conversation produced
the property statement now in plan.md.

**Settled in this triage:**

- The architecture is a *framework for intelligence* whose
  meta-structure is general, repetitive, and input-invariant
  (Hawkins-cortical-column generalization principle).
- Continuous learning is **necessary** (not sufficient) for
  intelligence. Without it, systems differ only in the timescale
  of their frozenness, not in kind.
- Growth dynamics: prediction-error-driven. Match strengthens
  connections; surprise-on-divergence triggers new learning
  (new nodes, new edges, modified weights, possibly modified
  mechanisms).
- The "seed-soil-input wirings" framing: framework is the seed,
  data is the soil, input wirings are the mechanism of growth.
  Seed metaphor is **weak** — evocative but not committing the
  architecture to a pre-specified developmental program.
- Reference frames are **created** by the growth process, not
  pre-given.
- The architecture is **deterministic given full specification**:
  identical seed + identical data + identical input wirings →
  identical outputs. Biconditional is strict.
- The architecture is **path-dependent on input wirings**: two
  trained instances are individuated by their seed + data +
  input-wiring history. Reproducibility requires all three.
  (Weak-B reading from sharpening discussion.)
- *Input wirings* defined comprehensively (Reading 3): routing
  of input streams to seed-regions, connection density and
  strength priors within and between regions, and any other
  aspect of how input gets into the framework.
- Prediction-and-surprise dynamic is **co-specified between
  items 1 and 2**: item 1 owns the *mechanism* (prediction-error-
  driven growth); item 2 owns the *substrate property* (the
  architecture is a generative model). Option γ from sharpening
  discussion.
- "Nodes and edges" used in graph-theoretic sense — no commitment
  to graphDBs or any particular implementation technology.
- Output-form / output-content distinction parallels topology-
  meta / topology-instance: the form is invariant; content is
  input-driven.

**Methodological commitments settled in this triage:**

- Logical grounding for each must-have is the *direction* of the
  work, not a Cycle 1 gate criterion. Each must-have starts as a
  sharpened property statement; logical grounding earned over
  multiple cycle iterations; mathematics later.
- **Phase A is multi-cycle by construction.** Cycle 1 produces
  the inventory in sharpened-property-statement form; subsequent
  cycles add depth.
- The triage workflow itself (Roger sharpens → Claude pushes back
  → revise/absorb/drop/split/defer → commit) is the working
  method's research mode in operation.

**Open within item 1 (deferred to subsequent cycles or to triage
of related items):**

- **Q-T006-B:** Is growth append-only, or does it include pruning
  / forgetting / consolidation? Surfaced when the bivalent match/
  no-match dynamic was specified. Deferred — to be addressed when
  the growth dynamics get sharpened in research mode.
- **Q-T006-C:** Is "node and edge creation" neurogenesis-
  equivalent or capacity-recruitment-equivalent at the
  implementation level? Different implementation paths; same
  observable behavior at the property-statement level.

**Resolved during item 1 triage (now closed, not filed as
Q-artifacts):**

- **Q-T006-A** (necessity grounding — practical/logical/both):
  resolved to *logical first, practical derived from logical*.
- **Q-T006-D** (output-form / output-content distinction):
  resolved to *meta in nature*.
- **Q-T006-E** (strong vs. weak seed metaphor): resolved to
  *weak*.
- **Q-T006-F** (prediction-and-surprise placement): resolved to
  *Option γ — lives in both items 1 and 2*.
- **Q-T006-G** (logical grounding strict vs. loose): resolved to
  *looser*.
- **Q-T006-H** (scope of "input wirings"): resolved to *Reading
  3 — comprehensive input-side specification including routing,
  connectivity priors, and any other input-side aspect*.

**Tentative absorptions** (pending triage of named items):

- Item 6 (Dimensional minimum-commitment) tentatively absorbed
  into item 1 — the input-driven growth process determines
  representational dimensionality. Confirmation pending when
  item 6 is reached in triage.

**Items 2-9 status:** pending triage. Triage proceeds one item
at a time; subsequent PRs continue the workflow.

### 2026-04-28 — Thread opened

@reh3376 named T006 as the central thread for the program's
architectural work, distinct from T002 (MDEMG trace) and T005
(program planning). The trigger was the explicit framing:

> *"The transformer was the fuse. The ANN architecture that we
> are proposing is the fuel that propagates the intelligence
> explosion. We will need a team, we will need funding, and we
> will need to get this idea infront of people with the resources
> to allow us to build it."*

Followed by the working-method commitment:

> *"You are a sharpening stone for me to hone my blade with. You
> are my faithful assistant in this process. I will provide the
> novelty as we work thru this. What it is not, what it must
> have, what it could be, the mathematics necessary to prove it
> in theory."*

The thread frame above codifies that working method explicitly so
future cycles inherit it without re-derivation.

T006 Cycle 1 opens with focus 2B from the prior turn's candidates:
**"what it must have"** — populating the positive space of
substrate properties the architecture must support. See plan.md.
