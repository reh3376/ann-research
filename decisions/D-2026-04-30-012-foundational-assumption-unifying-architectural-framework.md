---
id: D-2026-04-30-012
type: decision
slug: foundational-assumption-unifying-architectural-framework
created: 2026-04-30
updated: 2026-04-30
status: committed
confidence: 4
where: [T006]
tags: [foundational-assumption, architectural-framework, BNN, bayesian, master-algorithm, AGI, mathematics, methodology]
provenance:
  - "conversation 2026-04-30 between @reh3376 and Claude during T006 Cycle 1 foundational-assumptions exercise"
  - "@reh3376's condensed statement (verbatim): 'There exists a fundamental architectural framework that produces all knowledge. That framework is represented in the human biological neural network, but it it constrained by input types, energy, and compute power. The key to unlocking its full potential in artificial neural networks is emulating it as using Bayes and advanced mathemtics.'"
  - "Pedro Domingos, *The Master Algorithm* (2015) — origin of the master-algorithm framing the program is reframing"
  - "T006 plan items 1, 2, 3, 10 (already sharpened) — informed by this assumption and informing back"
  - "Q-2026-04-30-012 (working definition of general intelligence) — informs and is informed by this assumption"
  - "D-2026-04-28-009 (program objectives) — names AGI as destination; consistent with this decision"
  - "Hawkins / Mountcastle / grid-cell tradition — cited as evidence the BNN provides about the framework"
---

# Foundational assumption: there exists a unifying architectural framework

## Decision

The program commits to a **single foundational assumption** that
underlies all subsequent architectural work:

> *There exists a fundamental architectural framework that produces
> all knowledge sufficient for general intelligence. The framework
> is represented in the human biological neural network (BNN) as
> evidence — the BNN provides substantial structural evidence about
> the framework, but is not the framework itself; biological
> constraints (input types, energy, compute) may obscure parts of
> the framework not visible in the BNN. The methodological key for
> unlocking the framework's full potential in artificial neural
> networks is emulating the BNN's evident structure, using Bayesian
> reasoning broadly and advanced mathematics.*

This assumption is held with **strong prior confidence** (4 of 5)
based on @reh3376's program-direction reasoning and the inheritance
from items 1, 2, 3, and 10 of T006's plan. It is not a fact awaiting
proof; it is a working assumption that shapes the architectural work
and that the program's R&D cycles will test, refine, or revise as
evidence accumulates.

## Components

The assumption decomposes into three internally-coherent claims.
Decomposition is for analytical clarity; the three claims are
*components of one assumption*, not independent assumptions.

### Component 1 — Existence

**There exists a fundamental architectural framework that produces
all knowledge sufficient for general intelligence.**

The framework is *abstract* — not the BNN, not any particular ANN,
but a structural pattern that admits multiple instantiations. The
framework is what the program is working to identify, articulate,
and instantiate.

"All knowledge sufficient for general intelligence" is a *loose
universality* claim: the framework produces a sufficient range of
knowledge to support general intelligence as defined in
`Q-2026-04-30-012`, not a strict every-kind-of-knowledge-without-
gaps claim. The looseness is deliberate: strict universality would
take a position against well-known philosophical challenges to
single-source knowledge (normative knowledge, certain phenomenal
knowledge, etc.) without grounds the program is prepared to defend.

### Component 2 — Instantiation: the BNN as evidence

**The framework is represented in the human BNN as evidence — but
the BNN is not the framework itself.**

The BNN provides substantial structural evidence about the
framework. Studying the BNN tells us substantive things:

- The cortical column generalization principle (Hawkins,
  Mountcastle): the cortex's uniformity across regions, with
  region-specific representations emerging from input-driven
  growth, is evidence of the framework's input-invariant
  meta-structure (item 1)
- The hierarchical organization of cortex with bidirectional
  connections is evidence of the framework's recursive
  composability (item 10)
- Predictive-coding evidence in cortex is evidence of the
  framework's prediction-and-surprise dynamics (items 1 + 2)
- The BNN's reference-frame construction (grid cells, place
  cells, broader cortical reference-frame literature) is
  evidence of the framework's reference-frame substrate (item 3)

But the BNN is **not** the framework. Biological constraints —
input types limited to biological sensors, energy budget bounded
by metabolic cost, compute bounded by neuron count and connection
density — may obscure parts of the framework not visible in the
BNN. The framework's full structure may exceed what the BNN
reveals; an ANN instantiation freed from biological constraints
may instantiate the framework more completely than the BNN does.

This commitment ("BNN provides evidence, not realization") is
**Reading R-evidence** from the foundational-assumptions
sharpening conversation. It commits the program methodologically
to:

- Treating neuroscience as a substantial source of architectural
  insight (the BNN is evidence the program studies)
- NOT treating the BNN as a literal implementation target (the
  architecture is not committed to neuromorphic implementation;
  it does not need to look like brain tissue at the substrate
  level)
- Leaving open the possibility that the framework's full
  structure exceeds what the BNN reveals (the program's R&D
  may surface framework content not present in any biological
  organism)

### Component 3 — Methodology of discovery

**The key to unlocking the framework's full potential in ANNs is
emulating the BNN's evident structure, using Bayesian reasoning
broadly and advanced mathematics.**

Three sub-commitments:

**(a) Emulation of the BNN's evident structure.** The architecture
takes structural cues from what the BNN reveals: cortical-column
uniformity at every level, hierarchical recursive organization,
prediction-error-driven learning, reference-frame construction.
These structural cues constrain the architecture's design.
Per Component 2, the BNN's structure is evidence, not
prescription — emulation extends to the framework-evident
features, not to biological details that are constraint-artifacts.

**(b) Bayesian reasoning broadly.** Probabilistic reasoning with
prior + likelihood + evidence updating, including derived
frameworks: active inference, predictive coding, the free-energy
principle, hierarchical Bayesian models, and related. **Not**
literal Bayes Theorem alone (P(A|B) = P(B|A)P(A)/P(B) and direct
extensions); the broader Bayesian paradigm is what's committed.
This is consistent with item 2's hybrid α+β commitment — the α
component (probabilistic) inherits from this commitment.

**(c) Advanced mathematics.** The framework's full specification
will require mathematical machinery beyond current ANN practice.
Specific examples already surfaced by program work:

- Mathematics of N-level recursion with determinism-given-full-
  specification (item 10 / item 1 interaction; flagged by
  @reh3376 during item 10 triage as requiring new mathematics)
- Mathematics of α + β integration coherence (Q-T006-P from
  item 2's triage)
- Mathematics of recursion-induced exponential expansion
  (EX-3 from item 3's triage)

The mathematics required is not yet fully specified. Some of it
exists in the literature (Friston's free-energy mathematics;
hierarchical Bayesian modeling; cortical-column mathematical
models); some of it does not yet exist and the program may need
to develop it during T006's mathematics mode.

## Domingos lineage

The program credits Pedro Domingos's *The Master Algorithm* (2015)
for the general framing. Domingos argued the five major ML
paradigms (symbolists, connectionists, evolutionaries, Bayesians,
analogizers) might converge on a single unifying algorithm capable
of learning anything from data.

This program reframes Domingos's claim:

- **Where Domingos posited a single executable algorithm**, this
  program posits a **single architectural framework** (Reading β
  from sharpening — architectural form, not algorithm form). The
  framework IS the architecture; the architecture's growth IS the
  framework's operation. The algorithmic and architectural
  perspectives converge.
- **Where Domingos's claim was about ML paradigm convergence**,
  this program's claim is about **structural convergence in
  the BNN as evidence** — the BNN is not predominantly any of
  the five paradigms; it is the framework, instantiated.
- **Where Domingos's claim required theoretical proof**, this
  program's claim requires **structured testing and R&D to
  prove out**. The program will follow the data; the assumption
  is held with strong prior confidence and is revisable based
  on what construction reveals.

Crediting Domingos preserves the lineage. Reframing into an
architectural-framework + BNN-as-evidence claim is the program's
contribution.

## Program destination: AGI

The program's destination is **artificial general intelligence
(AGI)**. The framework, when freed from BNN's biological
constraints, is intended to produce general intelligence as
defined in `Q-2026-04-30-012`.

Whether the unconstrained framework's potential extends beyond
human-level intelligence is a question for after AGI is
demonstrated; this question is **not addressed by this decision**
and is not part of the program's stated commitment. Premature
commitment to a destination beyond AGI would over-commit on
grounds the program does not yet have.

## What this decision commits the program to

**Architectural commitments inherited or implied:**

- The framework is a single architectural pattern (consistent
  with item 10's pattern-at-every-level)
- The framework instantiates via input-driven growth (item 1)
- The framework supports prediction-and-surprise as substrate
  property (items 1 + 2)
- The framework supports reference-frame construction (item 3,
  via item 10's pattern element)
- The framework supports recursive composability (item 10)

This decision does **not** introduce new architectural primitives.
It articulates the foundational assumption underneath items 1, 2,
3, 10 — the assumption that motivates the work and constrains
what counts as success.

**Methodological commitments:**

- Neuroscience is studied as a primary source of architectural
  insight (BNN as evidence)
- Bayesian reasoning broadly — not literal Bayes Theorem alone —
  is the probabilistic foundation
- Advanced mathematics, including new mathematics where
  necessary, is required for full specification
- The destination is AGI; ASI is not in scope at this stage

**Research commitments:**

- The assumption is held with strong prior confidence; it is
  not a fact
- Construction work in Phase C will test whether the assumption
  bears out; the assumption may be revised based on what
  construction reveals
- Where the data leads, the program follows

## Cross-references

This decision is referenced from and informs:

- **Items 1, 2, 3, 10** (T006 plan) — the architectural
  commitments this assumption underlies
- **Q-2026-04-30-012** — working definition of general
  intelligence; this decision's destination commitment depends
  on Q-012's framing
- **D-2026-04-28-009** — program objectives; AGI destination
  consistent with both
- **Q-2026-04-30-013** (filed in same PR as this decision) —
  computational complexity prerequisites for the architecture's
  first applied problem; raises questions this decision does not
  address
- **Q-2026-04-30-014** (filed in same PR as this decision) —
  3-D Tetris as candidate first applied problem; relates to this
  decision via the Phase 0 / Phase 1 distinction in construction
  work

## Open within this decision

The following questions are deferred to future cycles or to
construction work; they do not block the assumption's commitment:

- **Q-T006-S (new):** What mathematical machinery — beyond what
  already exists in the active-inference / FEP / hierarchical-
  Bayesian literature — is required for full specification of
  the framework? Existing surfaced specifics: mathematics of
  N-level recursion (item 10), α+β integration coherence
  (Q-T006-P), EX-3 expansion (item 3). To be sharpened in T006's
  mathematics mode.
- **Q-T006-T (new):** What is the full extent of "advanced
  mathematics" Component 3 commits the program to? The phrase is
  intentionally loose; sharpening as the work proceeds.
- The relationship between the framework's full potential and
  AGI specifically — whether the framework's unconstrained
  potential is *equal to* AGI or *exceeds* AGI in some way that
  is not yet known — is not addressed by this decision and is
  open.

## Status updates expected

This decision should be revisited:

- When items 4-9 of T006's plan are sharpened, to verify
  consistency with this assumption
- When T006's mathematics mode begins, to refine Component 3's
  mathematical commitments
- When Phase C construction work surfaces evidence about the
  framework's structure, to refine or revise components
- Periodically as the broader AI research community produces
  results bearing on the master-algorithm question
