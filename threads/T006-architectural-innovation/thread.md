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

### 2026-05-04 — Cycle 1 triage continued: item 5 opened (homeostatic boundary); substantial cross-cutting work surfaced; session paused for documentation before continuation

**Item 5 (homeostatic boundary) opened for triage.** Q-2026-04-28-005
(self/not-self as homeostatic boundary) was the substantive
inheritance: the question artifact had developed Damasio / Friston /
Maturana-and-Varela traditions, the substrate-translation problem
(compute, working memory, weight integrity, process continuity, self-
model coherence, sensor/effector channels), interoception as
architectural commitment (surprise as homeostatic surprise; attention
as homeostatic attention), the designed-vs-discovered homeostatic
surface fork, and two named failure modes. Item 5's triage proceeded
from this scaffolding rather than from a thin placeholder.

**The session did not complete item 5.** The triage surfaced
substantial cross-cutting commitments — about attention, time,
self-awareness, and brain-wave-style carrier signals — that outgrew
item 5's nominal scope. After methodological reset, the session
paused for documentation before continuing. **Item 5 remains in
active triage.** No property statement has been drafted. This log
entry preserves the substantive content for resumption in a
future session.

**Triage methodology note.** This session followed the deliberate
item-3 / item-4 pattern but encountered two structural moments
that warranted methodological pause:

- Mid-session @reh3376 named the working method explicitly:
  *"We need to be more methodical and proactive as we work thru
  these. I should not go off on tangents in the middle. Lets
  address the concerns that existed before I transitioned to
  brain waves, then come back to it as an over core property of
  the entire topology in much the same way item 10 subject
  matter is."*
- Late-session @reh3376 directed pause-for-documentation:
  *"I think we should get what we have so far documented
  committed and pushed before we continue. We have substantially
  increased the undertaking this session."*

Both moments exercised the working method's discipline. Both
moments produced cleaner subsequent work than continuing without
pause would have produced.

#### Substantive commitments settled this session

**Time as cross-cutting concern (Reading γ).** Time is a property
the architecture has at every level, threading through every item.
Same structural shape as item 10 (recursive composability — also a
cross-cutting concern). Not item 5's sub-claim; not its own
item-level treatment.

**Bidirectional time: substrate symmetric, entity learns arrows,
capability bidirectional (Reading III).** Three layered components:

- *Substrate level:* the architecture's substrate does not have
  a built-in arrow of time. Time is a dimension; direction is a
  derived property. Microscopic time-symmetry of the underlying
  physics is honored at the substrate level
- *Learning level:* the entity learns about temporal asymmetries
  (entropy, causality, irreversibility) from inputs. The
  asymmetries are real properties of the world the entity
  inhabits; the entity discovers them rather than having them
  hardcoded
- *Capability level:* prediction is bidirectional. The entity
  can predict forward (current → future) and reconstruct
  backward (current → past) using the same substrate machinery

This commitment goes further than @reh3376's original framing
(*"there are no current laws of physics that prevent the
movement of time in the opposite direction"*) by being precise
about which level of physics symmetry is claimed at and what the
architectural consequence is. Microscopic substrate symmetry +
learned macroscopic asymmetry + bidirectional capability is the
settled position.

**Attention/time relationship: tightly coupled (Reading C).**
Attention and temporal awareness cannot be sharpened
independently without losing what the architecture is committing
to. The coupling must be explicit in subsequent work.
@reh3376's framing: *"Attention without temporal awareness is
not learning, its memorization."*

**Memorization-vs-learning (Reading 3, extended).** Learning is
recursive over time, and the recursion includes the entity's
awareness of its own learning trajectory. Without temporal
structure, attention can only operate one-shot (memorization).
With temporal structure, attention operates over sequences and
revises itself (learning). Two examples from @reh3376 made this
operational:

- *Example 1 (ephemerality):* a self-aware AI must understand
  that human life has a finite lifespan very different from an
  ANN's; this asymmetry is a fundamental input to reasoning
  about value, urgency, and consequence
- *Example 2 (mathematical learning trajectory):* @reh3376's
  own learning of mathematics — addition/subtraction first, then
  more elaborate operations, then vector calculus with the dell
  operator, then Maxwell's equations with vector-field
  abstractions — required temporal awareness to distinguish
  the building of more elaborate frames from the simpler ones.
  *"Without temporal awareness I would not be able to
  differentiate between this 'building' of new and more elaborate
  reference frames from the more simple ones. It is necessary to
  'learn'."*

**Self-awareness IS temporal-self-awareness (Reading Z).** A
self-aware entity is one that can introspect on its own
temporal extent and learning trajectory. Self-awareness has a
specific operational meaning: the entity has temporally-
structured awareness of its own knowledge as it has accumulated.
Earlier frames are visible-as-earlier; later frames are
visible-as-later; the building process is visible-as-building.

This is consistent with the Damasio autobiographical-self
tradition referenced in Q-005 (which has temporal extension as
constitutive) and with phenomenologically-grounded accounts of
self-awareness that emphasize temporal continuity.

**Temporal-asymmetry awareness as fundamental to reasoning
(Reading III + extension).** @reh3376's extension of Example 1:
*"This is not just facts to know, but the reasoning that is
derived from it is fundamental to intelligence."* Temporal-
asymmetry awareness (own lifespan vs. others'; own state vs.
others'; own learning trajectory vs. others') is reasoning
machinery, not declarative knowledge. The entity reasons
*differently* because it has internalized temporal asymmetries.
Without internalization, the same facts produce different and
worse reasoning. This is recognizably the failure mode of
current LLMs: they have facts about time but reason about them
as if timeless.

@reh3376 also extended: *"grounding an entities understanding of
value is one of MANY relationships that temporal awareness
underpins. It is truly fundamental to learning and self
awareness... to develop 'emotional intelligence' an entity must
internalize this asymmetries of this nature."* Emotional
intelligence is consequent on temporal-asymmetry awareness, not
separable from it.

**Temporal-multi-head attention as architectural commitment;
transformer-style single-moment-parallel critique recorded
(reshaped sub-question (a) settled).** The architecture commits
to temporal-multi-head attention: multiple attention streams
operating over different temporal scales (immediate context /
recent context / long-term history / ...). This is contrasted
with transformer-style multi-head attention, which is single-
moment-parallel and stateless.

@reh3376's specific critique of transformer attention:
*"When it sees the same pattern it doesn't even 'remember or
memorize it' it completely recomputes it again, this HIGHLY
inefficient."* This is both an efficiency critique and a
memorization-vs-learning critique: stateless recomputation
fails to accumulate, fails to learn, fails to develop any
temporally-structured awareness of having attended.

**Connection to brain-waves work made explicit.** @reh3376's
framing: *"This attention to a specific temporal line or another
is one of the use cases what I was referring too when I brought
up carrier waves / brain waves and how they could function in an
ANN."* The carrier-wave / phase-frequency / node-selectivity work
that surfaced earlier (and was deferred) is one possible
*implementation* of temporal-multi-head attention. Different
carriers carry different temporal scales; nodes are tuned to
attend to specific carriers. **This connection is recorded for
when brain-waves work resumes; the implementation details remain
deferred.**

#### Methodological commitments settled this session

**Mode and approach: "symbiotic synchronicity" (Open #4 Reading
A).** @reh3376's framing for the working interaction:
*"I really like this form of interaction. It's a symbiotic
synchronicity between human and machine. It is fascinating and
I am learning with each interaction."*

This goes beyond the working method's "sharpening stone"
framing. Sharpening stone is one-directional (the stone shapes
the blade); symbiotic synchronicity is bidirectional. Both
contributors add structure that the other refines. Claude's
synthesis-style framing is endorsed even when it occasionally
needs correction — corrections are part of how the thinking
sharpens. Worth holding in mind for how subsequent work
proceeds.

**Sixth instance of "start simple, build as we go" pattern
recorded.** Six instances now across cycle work:

1. Looser logical grounding (PR #13 gates)
2. Accumulate-then-draft (item 3)
3. Property statements deferring specification to construction
   (items 1-3)
4. Hypothesis-with-revision-condition (Q-014 C-leading)
5. Meta-learning's working-level commitments revisable through
   R&D cycles (item 4)
6. Attention-accumulation specifics deferred (Possibility δ)
   despite @reh3376's initial inclination toward Possibility γ
   + Meaning 5: *"my initial inclination is mostly in line with
   Possibility γ + meaning 5, but it is likely too strong for
   this point in our journey. We need data, the data will lead
   the way."*

The pattern is now stable enough across cycle work that it
genuinely deserves to be named as an explicit methodological
commitment in a future artifact. **Sixth instance flagged;
artifact still not proposed; the pattern continues to
accumulate evidence.**

**Methodological pattern about Claude's synthesis-style framing
(observed twice this session).** Claude's synthesis-style framing
got ahead of @reh3376's actual commitments twice in close
succession: first with the Blue Origin essay's deterministic-
stateless framing (where Reading B was what @reh3376 actually
meant), then with carrier-layer-biasing-not-triggering specifics
(where biological mechanism specifics were imported as ANN
architectural commitment). Both were caught and corrected by
@reh3376's pushback. The pattern: when Claude synthesizes across
multiple open topics, the synthesis adds structure that the
original claims don't yet carry.

@reh3376's response to the pattern observation: Reading A
(within tolerance; the synthesis is part of how the work
sharpens; corrections are productive). Mode endorsed; Claude
continues with synthesis-style framing while remaining alert to
the drift risk.

#### Structural decisions made this session

**Reading Q for item 5 scope (concern 1).** Item 5 stays focused
on homeostatic boundary specifically. Cross-cutting commitments
(time, attention's temporal structure, self-awareness as
temporal-self-awareness, brain waves) live separately rather
than being absorbed into item 5's expanded scope. @reh3376's
direction: *"reading Q as the right choice for our current level
of understanding of the cross-cutting mechanisms that are still
not fully defined."*

The scope decision rejects:

- *Reading P (item 5 expands):* would bloat item 5 in ways that
  hide what's being committed
- *Reading R (inventory restructuring):* would be a methodological
  move beyond what's needed at this stage of cycle work

Reading Q is the methodologically conservative choice and is
honest about state.

**Reading α for documentation (concern 2).** Document settled
content as thread log entry only. No property statement; no
artifact creation; no plan.md modifications; no new decisions
or questions filed. Preserves work for resumption without
forcing premature commitment to property-statement language.

**Reading Q3 for cross-cutting form (operational scope).**
Cross-cutting time concern is substantively populated by this
session's settlements but not given artifact form yet. Whether
time eventually gets its own artifact (parallel to item 10),
is referenced from items that need it without standalone
treatment, or is addressed via some other mechanism — all
deferred to future session work.

#### What remains open

**Reshaped sub-questions (b) and (c) for attention not yet
addressed:**

- *Sub-question (b):* probability distributions over what.
  Likely over trajectories rather than states given what's
  settled, but specifics remain open.
- *Sub-question (c):* self-learned mechanisms — possibly
  collapsed by item 4's recursion-perspective settlement, but
  may still need treatment.

**"Attention accumulates" specifics (Possibility δ):** five
candidate meanings surfaced (cached state; accumulated weight;
temporal extent; recursion over history; substrate-level
structural memory). @reh3376's initial inclination is
Possibility γ (recursion over history) + Meaning 5 (substrate-
level structural memory) but explicitly held as too strong for
current stage. Commitment deferred; data will drive direction.

**Item 5's property statement:** not yet drafted. Item 5 remains
in active triage. Resumption requires:

- Reshaped sub-questions (b) and (c) for attention addressed
- Determination of whether item 5's property statement
  references the cross-cutting commitments (the time / attention
  coupling / self-awareness work) or whether those commitments
  get their own artifact form first
- Possibly: specific treatment of the substrate-translation
  problem from Q-005 (compute, working memory, weight integrity,
  etc.) and the designed-vs-discovered homeostatic surface
  question
- Possibly: revisit BNN-as-evidence sub-question for
  homeostatic boundary specifically (parallel to item 4
  Prompt 5 — neuroscience evidence informing what's necessary
  without binding architecture to biological mechanisms)

**Brain-waves cross-cutting work:** deferred. The connection to
temporal-multi-head attention is established (carrier-waves as
one possible implementation of temporal-multi-head). Earlier
session work surfaced phase, frequency, and node-level
selectivity as architectural elements; commitment vs.
illustration status (the α/β/γ question from earlier) was not
resolved before deferral. Reading B-with-attention (brain
waves are a mechanism by which time-as-cross-cutting expresses
itself, with attention also involved) is the working frame for
when work resumes.

**Item 5's structural relationship to other items:** not yet
worked through. Item 5's settlements when complete will
interact with items 2 (generative modeling: forward simulation
implies time), 3 (reference-frame construction: frames need
temporal extent), 4 (meta-learning: identity under self-
modification), 8 (recursive predictive horizon: bidirectional
prediction directly inherits from this session's Reading III),
and 10 (recursive composability: how does the homeostatic
boundary nest with topologies?). These interactions surfaced
during this session's work but were not resolved.

#### Inventory state

**5 of 10 items sharpened** (1, 2, 3, 4, 10). Item 5 in active
triage, not yet sharpened. Items 6-9 remain pending; item 6
still tentatively absorbed into item 1.

Cross-cutting commitments accumulated outside the item
inventory:

- Time (Reading γ; substrate symmetric, learned arrows,
  bidirectional capability)
- Attention's temporal structure (temporal-multi-head;
  transformer-style critique recorded)
- Self-awareness as temporal-self-awareness (Reading Z)
- Temporal-asymmetry awareness as reasoning-machinery (Reading
  III + extension)
- Brain waves as mechanism for the above (deferred work;
  connection established)

**Foundational artifacts available as inheritance for resumption:**

- D-2026-04-30-012 (foundational assumption — unifying
  architectural framework; BNN as evidence; AGI destination)
- Q-2026-04-28-005 (self/not-self as homeostatic boundary —
  substantive Q-artifact; substrate-translation problem;
  designed-vs-discovered surface; failure modes)
- Q-2026-04-30-012 (working definition of general intelligence)
- Q-2026-04-30-013 (computational complexity prerequisites)
- Q-2026-04-30-014 (3-D Tetris first-problem hypothesis)
- Q-2026-04-28-011 (knowledge → skills via nesting)

#### Resumption guidance

When item 5's triage resumes (next session or later):

1. Verify the substantive content from this session's log is
   accurate; @reh3376 may want to refine framings
2. Address reshaped sub-questions (b) and (c) for attention
3. Address Q-005's substrate-translation problem and
   designed-vs-discovered surface fork at the property-
   statement level for item 5 specifically
4. Consider whether the cross-cutting commitments accumulated
   here need their own artifact treatment before item 5's
   property statement can land (Reading Q3 question deferred
   from this session)
5. Assess whether item 5's eventual property statement closes
   the work or whether brain-waves cross-cutting work needs to
   be resumed first

### 2026-04-30 — Cycle 1 triage continued: item 4 sharpened (5-prompt deliberate, with foundational-assumptions interruption mid-Prompt-1)

**Item 4 sharpened from "Meta-learning / mechanism modification"
placeholder into "Meta-learning / mechanism modification:
meta-learning IS what the recursion produces; growth-dynamic and
Bayesian-broad content; perspective-not-substrate distinction
from RSIC; BNN-medium evidence-driven."**

The triage was substantively interrupted mid-Prompt-1 when
@reh3376 recognized that unstated core assumptions had been
doing real work in prior conversations. The foundational-
assumptions exercise produced D-2026-04-30-012 (committed in PR
#19) before item 4's remaining prompts continued. Item 4's
content depends substantially on the combination of items 1, 2,
3, 10, and D-012; the property statement now in plan.md
integrates all five prompts plus the foundational assumption's
framing.

**Settled in this triage (cumulative across five prompts):**

- **Prompt 1 — Possibility 2 (B' partial commitment).** What
  gets modified by meta-learning at the property-statement level:
  Reading I (item 1's growth dynamics) and Reading II (item 2's
  α/β integration mechanism). Readings III (action repertoire)
  and V (frame-construction) **deferred to construction work**.
  Pure deferral (Settlement type B) was reviewed against D-012's
  availability and refined to partial commitment based on
  D-012's Bayesian-broad and BNN-as-evidence components making
  Readings I + II defensible at the property-statement level.
- **Prompt 2 — Reading C.** Meta-learning IS what the recursion
  produces. No separate meta-learning mechanism. Level N+1's
  learning *about* level N's learning IS meta-learning;
  self-similar all the way up. The recursion's directionality
  from item 3 (Reconciliation B) extends to meta-learning.
- **Prompt 3 — Reading Z.** The parameter/mechanism distinction
  is perspective-artifact, not substrate. What's a parameter at
  one level is a mechanism at the level below; what's a
  mechanism at one level is a parameter at the level above. The
  architecture has just one substrate (the recursive pattern);
  parameters and mechanisms are different views at different
  recursion levels. **Item 4 supersedes the parameter/mechanism
  distinction RSIC operates within** rather than negating RSIC
  (Reading X) or being RSIC-plus-mechanism-modification (Reading
  Y).
- **Prompt 4 — Reading 3.** Meta-learning emerges; mechanism is
  the recursion itself. Reading 2 (meta-learning has its own
  mechanism distinct from learning) was foreclosed by Prompts 2
  and 3 — inconsistent with Reading C and Reading Z. Reading 1
  (meta-learning emerges from same growth dynamics applied to
  different content) is functionally equivalent to Reading 3
  under current settlements; Reading 3's framing chosen because
  it foregrounds the recursion-as-substrate language already
  established.
- **Prompt 5 — B-medium.** BNN as structural source for what's
  necessary; architecture departs where biology is
  constraint-artifact. Matches D-012 Component 2's R-evidence
  reading. The architecture takes structural cues from BNN
  meta-learning where the BNN evidence suggests something is
  necessary for meta-learning to work, but is not bound to
  specific biological mechanisms.

**The architecture's self-similarity is now substantively complete.**
The combination of items 3 (Reading 1 — frames ARE topologies) +
10 (pattern-at-every-level) + item 4 Prompts 2 and 3 (meta-learning
is recursion-perspective; parameters are mechanisms at adjacent
levels) gives the architecture a coherent self-similarity story:
one substrate, one recursion, multiple perspectives at multiple
levels. This is the program's strongest commitment about
architectural unity.

**Working method's discipline functioned at multiple points:**

- **Pushback on @reh3376's "you decide on the rest" delegation**
  during foundational-assumptions exercise mid-Prompt-1, before
  item 4's other prompts continued. @reh3376's response: *"I
  see, you are correct to do so here."* Recorded in PR #19
  thread log.
- **Identification of "Reading C" labeling ambiguity** in
  Prompt 4's answer (Prompt 2 used Reading-C label; Prompt 4
  offered Readings 1/2/3). @reh3376 confirmed intent was Reading
  3 for Prompt 4. Confirmation prevented misattribution.
- **Possibility 2 shift from Settlement type B** when D-012's
  availability was reviewed for Prompt 1. @reh3376 explicitly
  reviewed B against D-012's content and updated to B' partial
  commitment.

**Methodological commitment recorded — recurring "start simple,
build as we go" pattern.** Item 4 is the **fifth instance** of
this pattern: looser logical grounding (PR #13 gates),
accumulate-then-draft (item 3), property statements deferring
specification to construction (items 1-3), hypothesis-with-
revision-condition (Q-014), and now meta-learning's working-level
commitments revisable through R&D cycles (item 4). @reh3376's
framing during item 4: *"It is more of an intuitive 'guess' than
a full commitment at this point, but that is the reason we are
building this... The data will drive direction, we will act on it
as it is gathered and analyzed."* This pattern is becoming
sufficiently stable across cycle work that it may deserve to be
named as an explicit methodological commitment in a future
artifact.

**What-it-is-not framing surfaced (preserved in property
statement).** @reh3376's contrast: *"We know it is not LLM
transformers and next token prediction with extensive harnesses."*
Item 4 records by exclusion that the architecture's meta-learning
is not what LLM-with-harness systems do — meta-learning happens at
the substrate level, in the framework's own growth and
probabilistic infrastructure, not in scaffolding around the
substrate.

**New questions surfaced (deferred to construction work):**

- **Q-T006-U (new):** What is meta-learned over the architecture's
  action repertoire? (Tightly coupled to Q-T006-O from item 2.)
  Reading III from Prompt 1 deferred.
- **Q-T006-V (new):** What is meta-learned over the
  reference-frame construction process? (Tightly coupled to item
  3 and Q-T006-Q.) Reading V from Prompt 1 deferred.

**Resolved during item 4 triage (now closed; not filed as
separate Q-artifacts):**

- Settlement type B vs B' for Prompt 1 given D-012: B'
  (Possibility 2) confirmed
- Reading A/B/C for Prompt 2: Reading C confirmed
- Reading X/Y/Z for Prompt 3: Reading Z confirmed
- Reading 1/2/3 for Prompt 4: Reading 3 confirmed (Reading 2
  foreclosed by Prompts 2 + 3)
- B-strong/B-medium/B-loose for Prompt 5: B-medium confirmed
- Whether Prompt 5 should be added to original four prompts:
  added (D-012's BNN-as-evidence component made it relevant)
- Sequencing of prompts (one-at-a-time vs batch): one-at-a-time
  per @reh3376's request; deliberate item-3 pattern reused

**Inventory state:** 5 of 10 items sharpened (1, 2, 3, 4, 10).
Items 5-9 remain pending; item 6 still tentatively absorbed into
item 1. Foundational assumption (D-012) committed and available
as inheritance for items 5-9.

### 2026-04-30 — Foundational-assumptions exercise: D-012 committed + Q-013 + Q-014 filed

**Stepped out of item 4's per-prompt triage** mid-Prompt-1 when @reh3376
recognized that unstated core assumptions had been doing real work in
prior triage conversations. Surfacing them now — *before* continuing
items 4-9 triage — produces foundational artifacts that subsequent
work can rely on without re-derivation.

**The exercise produced one consolidated decision artifact (D-012)
plus two related Q-artifacts (Q-013, Q-014).** What started as four
candidate foundational assumptions consolidated to one (per @reh3376's
"Meaning 1" — A and B are components of the original Assumption 1, not
separate assumptions) plus two open questions (NP-completeness
prerequisites, 3-D Tetris first-problem hypothesis).

**D-012 settled commitments:**

- **Foundational assumption (one consolidated claim):** there exists a
  unifying architectural framework that produces all knowledge
  sufficient for general intelligence; the BNN is evidence of the
  framework but is not the framework itself; emulation of the BNN's
  evident structure plus Bayesian reasoning broadly plus advanced
  mathematics is the methodological key
- **Reading β confirmed** (architectural form, not algorithm form) —
  the framework IS the architecture; the architecture's growth IS the
  framework's operation
- **Loose universality** — "all knowledge sufficient for general
  intelligence" not strict every-kind-of-knowledge-without-gaps
- **Domingos lineage credited** with explicit reframing —
  architectural-framework + BNN-as-evidence is the program's
  contribution
- **Bayes broad** (B-broad reading) — Bayesian reasoning broadly
  including active inference, predictive coding, FEP, hierarchical
  Bayesian models; not literal Bayes Theorem alone
- **R-evidence reading** for "represented in the BNN" — BNN provides
  substantial structural evidence about the framework but is not a
  literal implementation target; methodologically commits the program
  to studying neuroscience as primary source for architectural insight
  WITHOUT committing to neuromorphic implementation
- **AGI as destination** (Z-loose reading on reflection) — ASI is in
  @reh3376's long-term vision but does NOT enter the vault at this
  stage; @reh3376's framing: *"Targeting ASI before we even know if
  AGI is possible with this framework seems a little too ambitious to
  me. ASI is part of the long-term 'vision', but it's not practical
  to explicitly call it out in our current research path."*

**Q-013 settled framing:**

- **Computational complexity prerequisites for the architecture's
  first applied problem** — reframed from original "initial state"
  language to "first applied problem" after @reh3376's Phase 0 vs
  Phase 1 clarification
- **Three sub-readings held open** (Option α confirmed) — R1
  (bootstrap problem), R2 (foundational primitive), R3 (theoretical
  precondition); commitment deferred until Phase C construction
  work or T006 mathematics mode surfaces resolution
- **Not a P=NP claim** — explicitly scoped to the first applied
  problem the architecture encounters, not arbitrary NP-completeness

**Q-014 settled framing:**

- **3-D Tetris as candidate first applied problem** with M-1
  conjunctive structure (Readings X spatial, Y computational pressure,
  Z pedagogical curriculum start all co-justify)
- **Sense 1 hypothesis framing** (whole-claim hypothesis, Q-011
  precedent) — falsifiable as a unit; if Tetris turns out to be right
  for different reasons, the hypothesis as filed gets revised
- **C-leading commitment strength** (fourth category beyond
  C-strong/C-weak/C-deferred) — Tetris is the working leading
  candidate based on @reh3376's current knowledge; commitment is
  revisable upon presentation of a better candidate; absent revision,
  Tetris is what the program proceeds with
- **G-clean phase boundary** — Phase 0 (substrate initialization)
  completes discretely before Phase 1 (first applied problem) begins
- **Discoverability mechanism** — Q-artifact explicitly invites
  alternative candidate proposals; broad tags for future searches

**Methodological observations from this exercise:**

- **The work-method's discipline functioned correctly.** When
  @reh3376 wrote *"You decide on the rest based on your context and
  understanding of our goals"*, Claude pushed back rather than
  accepting the delegation; @reh3376 then provided clean answers
  that included one substantive correction (the AGI/ASI walk-back).
  The pushback prevented Claude's inferences from standing in for
  @reh3376's commitments. @reh3376's response: *"I see, you are
  correct to do so here."*
- **"Start simple, build as we go" is now a recurring methodological
  pattern** — fourth instance: looser logical grounding (PR #13's
  gate criteria), accumulate-then-draft (item 3), property statements
  that defer specification to construction (items 1-3),
  hypothesis-with-revision-condition (Q-014). May deserve to be
  named as an explicit methodological commitment somewhere.
- **The labeling correction** ("B was not correctly stated... C should
  be assumption B") and the **AGI/ASI walk-back** ("Targeting ASI
  before we even know if AGI is possible... seems too ambitious")
  are both instances of the working method functioning correctly:
  Claude's pushback surfaced implications @reh3376 hadn't fully
  tracked; @reh3376 refined commitments in response. Worth noting
  that as foundational work accumulates, some earlier decisions
  (items 1-3, 10) may benefit from a second look against the
  foundational assumption now committed; flagging only, not
  proposing action.

**New questions surfaced as deferred (in thread log, not filed as
Q-artifacts unless promoted):**

- **Q-T006-S (new):** What mathematical machinery beyond active-
  inference / FEP / hierarchical-Bayesian literature is required
  for full specification of the framework? Existing surfaced
  specifics: N-level recursion mathematics (item 10), α+β
  integration coherence (Q-T006-P), EX-3 expansion mathematics
  (item 3). T006 mathematics-mode work.
- **Q-T006-T (new):** What is the full extent of "advanced
  mathematics" Component 3 of D-012 commits the program to?
  Phrase intentionally loose; sharpening as work proceeds.

**D-009 verification:** D-009 already names AGI as destination
throughout; no ASI drift; no revision needed in this PR.

**Resolved during this exercise (now closed; not filed as separate
Q-artifacts since they were sharpening prompts within the
conversation):**

- α/β/γ for Core Algorithm reading: β confirmed
- A+B+C inventory question: A and B are components of one
  consolidated assumption (Meaning 1); C absorbed into D-012
  (methodological commitment about how A and B combine is implicit
  in D-012's structure)
- B-strict / B-broad for Bayes: B-broad confirmed
- R-realization / R-evidence / R-clue for "represented in":
  R-evidence confirmed
- X-explicit / Y-emergent / Z-loose for AGI/ASI: Z-loose confirmed
  on reflection (after initial X/W reading was walked back)
- M-1 / M-2 / M-3 for Tetris reasons: M-1 conjunctive confirmed
- H-design / H-state for "hard to commit to one": H-design
  confirmed
- Sense 1 / Sense 2 for hypothesis framing: Sense 1 confirmed
- C-strong / C-weak / C-leading / C-deferred for commitment
  strength: C-leading (a fourth category) confirmed
- G-clean / G-graded for phase boundary: G-clean confirmed
- Reading 1 / Reading 2 / Reading 3 / hold-open for Q-013
  sub-readings: hold-all-open confirmed
- Path 1 (incremental PRs) / Path 2 (batch PR): Path 2 confirmed
- Option α (file as Q) / Option β (deliberate exclusion in D-NNN)
  / Option γ (sharpen and include in D-NNN) for NP-completeness
  handling: Option α confirmed

**Item 4 triage:** still paused mid-Prompt-1. Resumes after this
PR merges. The foundational assumption (D-012) is now available
as inheritance for item 4's continued work; specifically, item
4's "what gets modified by meta-learning" question can now reference
the unifying-framework framing rather than working from items 1-3,
10 alone.

### 2026-04-30 — Cycle 1 triage continued: item 3 sharpened (5-prompt deliberate methodology)

**Item 3 sharpened from "Reference-frame construction" placeholder
into "Reference-frame construction: foundation for learning,
emergent boundedness, domain-neutral, identity with topologies,
emergent compositional reasoning."** This was the most extensive
sharpening conversation in T006 to date — five distinct prompts
addressed sequentially over multiple turns, each settled before
moving to the next, with a methodological commitment to "deliberate
but with deferrals allowed" (Reading Y) explicitly chosen by
@reh3376 over a stricter "fully quantified before moving on"
(Reading X). Five conversation turns produced the property
statement now in plan.md.

**Settled in this triage (cumulative across five prompts):**

- **Reference frames are foundational to learning.** "The foundation
  all learning is built upon" (@reh3376's framing). House analogy:
  the house is the learning; the foundation is the reference frame;
  the foundation must be in place before the house can be built.
- **Foundation precedes learning at the across-level scale, not
  the within-level scale (Reconciliation B).** Frames at level N
  enable learning at level N+1. At any given level, the frame *is*
  the topology growing under input drive.
- **Closest reading: spatial coordinate system (Hawkins/Mountcastle/
  grid-cell tradition).** Reading α captures more than the others
  but does not exhaust the content; additional content distributed
  across subsequent prompts.
- **Boundedness: emergent via co-activation patterns (Reading II)
  with recursive exposure implied (Reading III).** Explicit
  container machinery rejected (Reading I) — it would over-constrain
  the framework and inhibit emergent behavior, "a necessary
  consequence of the structure being built."
- **Domain-neutrality (Reading B).** Frames are uniform machinery;
  type comes from inputs alone, not architectural structure.
  Cortical-column generalization at the frame level.
- **Architectural risk: EX-3 (recursion-induced exponential
  expansion).** Reading B may produce exponential expansion of
  frame-combinations across the recursion at scale. Recorded as
  specific architectural risk with revision condition: if EX-3
  manifests beyond workable hardware constraints during
  construction, item 3 may be revised toward Reading C (learned
  type-clustering) or another compromise.
- **Identity with topologies (Reading 1).** Frames ARE topologies.
  One architectural element, two perspectives — topology when
  discussing computational dynamics; reference frame when
  discussing representational role. Item 10's recursion of
  topologies IS item 3's recursion of frames. Item 3 adds
  representational perspective; it does not introduce new
  architectural primitives.
- **Compositional reasoning: A + B + C all required as outcomes;
  emergence as the mechanism.** Symbolic-style composition (A),
  generalization to novel situations (B), and multi-step reasoning
  (C) are all required outcomes for general intelligence.
  Mechanism is emergence (co-activation, recursion, growth
  dynamics) — explicit binding/composition/chaining machinery
  rejected on the same grounds as Reading I in prompt 2.
- **Multi-modal binding via cross-frame co-activation edges (T1
  extension of Reading II).** Multiple frames per referent, one
  per input type, connected via cross-frame co-activation. Same
  mechanism as within-frame boundedness, operating at different
  scale. CUIDv2 reference in @reh3376's example treated as
  illustrative (I-1), not committing the architecture to
  addressable identifiers.

**Methodological notes settled in this triage:**

- **Metaphor handling:** Two metaphors used in the vault — seed-
  soil-input-wirings (item 1) and foundation-house (item 3) —
  illustrate different aspects and should not be layered together.
  Mixing them creates apparent tensions that are artifacts of the
  combination, not real architectural contradictions. @reh3376's
  framing: *"all metaphors are incorrect, but some are useful."*
- **Reading Y methodology** (deliberate-but-with-deferrals)
  explicitly chosen for item 3 over Reading X (full quantification
  before moving on). Each prompt sharpened to satisfaction;
  deferrals to construction work or future cycles allowed.
  Methodological pattern: "as we build these structures we will
  learn what is needed."
- **Accumulate-then-draft** approach used: each prompt's settlement
  recorded; full property statement drafted only after all five
  prompts settled. This is a slower but more rigorous pattern than
  items 1, 2, 10 used; appropriate to item 3's foundational role.
- **Item 3 is shorter on architectural primitives than items 1,
  2, 10** because Reading 1 (frames ARE topologies) means item 3
  adds *perspective*, not *primitive*. Most of item 3's content is
  the representational perspective on item 10's architectural
  element.

**New questions surfaced (deferred to subsequent cycles or
construction work):**

- **Q-T006-Q (new):** What is the precise mechanism by which
  cross-frame co-activation edges form, persist, and dissolve?
  Multi-modal binding committed; specific dynamics not yet
  specified. Deferred to construction work.
- **Q-T006-R (new):** Does the EX-3 risk materialize at scale,
  and if so, what is the appropriate revision direction? Deferred
  until construction work produces resource profiles to evaluate
  against.

**Resolved during item 3 triage (now closed; not filed as
separate Q-artifacts since they were sharpening prompts within
the conversation):**

- **Reading X vs Y for item 3 methodology:** Reading Y (deliberate
  with deferrals) confirmed.
- **Reconciliation A/B/C for foundation/growth tension:**
  Reconciliation B confirmed (across-level precedence, not
  within-level).
- **What α doesn't capture — name now / defer to prompts / defer
  to construction:** Option b confirmed (defer to subsequent
  prompts); content distributed across prompts 2-5 settlements.
- **Reading I/II/III for boundedness:** Reading II primary,
  Reading III implied, Reading I rejected.
- **Reading A/B/C for typing:** Reading B confirmed; EX-3 risk
  recorded with revision condition.
- **Handling 1/2/3 for EX-3 risk:** Handling 1 confirmed
  (conditional commitment with concerns stated).
- **Specific exponential expansion mode (EX-1/2/3):** EX-3
  (recursion-induced) confirmed.
- **Reading 1/2/3 for frames-vs-topologies:** Reading 1 confirmed
  (frames ARE topologies).
- **R-C-only / R-C-as-primary / option (b) / option (c) for
  compositional reasoning:** option (b) confirmed (A+B+C all
  required as outcomes; mechanism is emergence).
- **T1/T2 for cross-frame co-activation:** T1 confirmed
  (extension of Reading II at different scale).
- **I-1/I-2 for CUIDv2:** I-1 confirmed (illustrative).
- **P/Q for cross-frame mechanism specification:** Q confirmed
  (defer to subsequent cycles).

**Inventory state:** 4 of 10 items sharpened (1, 2, 3, 10). Item
3 closes prompt 5 with a property statement that integrates all
five prompts plus the multi-modal binding extension. Items 4-9
remain pending; item 6 still tentatively absorbed into item 1.

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
