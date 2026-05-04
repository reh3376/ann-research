# ann-research

*Last updated: 2026-04-30*

Working substrate of a research program toward an artificial neural
network architecture for general intelligence.

The program's claim, plainly: large language models are not the path
to AGI. They are a powerful local maximum and continuing to scale them
will not produce general intelligence. A different architecture is
required — one that learns continuously, builds and modifies its own
reference frames, predicts forward and backward in time over those
frames, and recursively composes with itself across levels of
abstraction. Identifying, articulating, and specifying that
architecture is what this program does.

This repository is where that work happens.

---

## What this repository is

A **working research vault** — the lab notebook for an active research
program — and a small Go CLI that maintains it.

It is not a finished artifact. Anyone reading at any moment will find
half-formed ideas, hypotheses awaiting evidence, and decisions that may
be revised. That is the point: the vault preserves the trajectory of
belief, not only its conclusions. The discipline of writing artifacts
that make sense to a stranger keeps the work honest, and the vault is
public for the same reason.

The vault contains:

- **Decisions** with the alternatives considered and rejected
- **Observations** made while reading code, papers, and prior work
- **Open questions**, including strong-prior research hypotheses with
  their evidence criteria and prior confidence recorded
- **Hypotheses** with explicit falsification criteria
- **Dead ends** — paths that were tried and failed, preserved so future
  work doesn't repeat them
- **Active research threads** that organize all of the above

If you are looking for the polished version of any line of thinking
here, look for an external paper, a talk, or a pitch document. If
there is no public artifact yet, the work is still in progress.

---

## What the program is working toward

This program is an active research effort, not a finished theory. The
substrate-level commitments below are working positions held with
substantial confidence; the open questions further down are exactly
that — open. We surface both because honest disclosure is what
attracts the collaborators we hope to work with.

### Sharpened substrate-level commitments

The architecture this program is building must satisfy a set of
substrate-level commitments that current architectures do not supply.
These are accumulated and sharpened in the architectural-innovation
thread (`threads/T006-architectural-innovation/`); the inventory grows
as the work proceeds. As of late April 2026, sharpened commitments
include:

- **Continuous learning via input-driven, prediction-error-driven
  topology growth** — the architecture is a framework whose
  meta-structure is general and input-invariant; data and input
  wirings drive the actual growth of nodes, edges, and reference
  frames; prediction-error and surprise drive what gets learned
- **Generative modeling: hybrid α+β with action-conditioning** — the
  architecture is *both* probabilistic (maintains distributions; surprise
  has precise meaning) AND forward-simulating (generates trajectories);
  predictions are action-conditioned, so the architecture is agentive,
  not a passive observer
- **Reference-frame construction** — reference frames are the
  foundation that learning is built upon; they emerge from co-activation
  patterns rather than explicit container machinery; they are
  domain-neutral as architectural primitives, with type acquired from
  inputs; they are identical with topologies (one architectural element,
  two perspectives — computational dynamics and representational role)
- **Meta-learning / mechanism modification** — meta-learning IS what
  the recursion produces (no separate mechanism); the parameter /
  mechanism distinction is perspective-artifact rather than substrate;
  the architecture's recursive self-improvement supersedes the
  parameter-tuning scope of conventional self-improvement frameworks
- **Recursive composability** — the architectural pattern recurs at
  every level of nested composition; topologies compose with topologies
  via inter-topology wiring; the architecture is self-similar in the
  strict sense

Together, items 3 (reference-frame construction), 4 (meta-learning),
and 10 (recursive composability) compose into a unified self-similarity
claim: the architecture has one substrate, one recursion, and multiple
perspectives at multiple levels — what we call "frames" vs "topologies",
"learning" vs "meta-learning", "parameters" vs "mechanisms" are
perspective views of the same underlying recursive pattern.

Additional substrate commitments — homeostatic boundary,
projection-and-anchoring, recursive predictive horizon, prime
directives — are in active sharpening. Each is recorded in the vault
as it stabilizes.

### Foundational assumption

The program operates under a foundational assumption
(`D-2026-04-30-012`) that conceptually underlies all the substrate-
level commitments above. The assumption decomposes into three
internally-coherent claims, held together as one foundational
position:

**Existence.** There exists a fundamental architectural framework
that produces all knowledge sufficient for general intelligence. The
framework is *abstract* — not the human brain, not any particular
artificial neural network, but a structural pattern that admits
multiple instantiations. This is a Domingos-credited claim — Pedro
Domingos's *The Master Algorithm* (2015) is the lineage — but
reframed: the program commits to a single architectural framework
rather than a single executable algorithm.

**Instantiation.** The framework is *represented in* the human
biological neural network as **evidence** — not as the framework
itself. The BNN provides substantial structural evidence about the
framework (cortical-column uniformity, hierarchical organization
with bidirectional connections, predictive-coding dynamics,
reference-frame construction via grid cells and place cells), but
biological constraints (input types, energy budget, compute
limits) may obscure parts of the framework not visible in the BNN.
The architecture is not committed to neuromorphic implementation.

**Methodology of discovery.** The key to unlocking the framework's
full potential in artificial neural networks is emulating the BNN's
evident structure, using Bayesian reasoning broadly (probabilistic
inference including active inference, predictive coding, the
free-energy principle, and hierarchical Bayesian models — not
literal Bayes Theorem alone) and advanced mathematics. Some of the
mathematics required already exists; some of it does not, and the
program may need to develop it.

The program's destination is **AGI** (artificial general
intelligence). Whether the unconstrained framework's potential
extends beyond human-level intelligence is a question for after
AGI is demonstrated; this is not part of the program's stated
commitment at this stage.

### Open questions

These are the questions held open in the vault. Each represents
work the program does not yet know how to do, and where thoughtful
collaboration would directly advance the work:

- **`Q-2026-04-30-012` — Working definition of general
  intelligence.** Seed answer: *the property that enables reliable
  prediction across heterogeneous prediction domains*. Three
  sub-questions held open: what counts as a "prediction domain";
  what counts as "reliable" prediction; whether intelligence is
  the *enabling property* of reliable prediction or *constitutive
  of* it.
- **`Q-2026-04-28-011` — Knowledge → skills via nesting.** Strong-
  prior research hypothesis: skills are hypothesized to emerge as
  meta-frames over knowledge frames as the architecture's recursion
  deepens. Held with strong prior confidence; awaiting evidence
  from R&D cycles.
- **`Q-2026-04-30-013` — Computational complexity prerequisites
  for the architecture's first applied problem.** Three
  sub-readings (bootstrap problem; foundational primitive;
  theoretical precondition) all held open. Becomes load-bearing
  when Phase C construction begins.
- **`Q-2026-04-30-014` — 3-D Tetris as candidate first applied
  problem.** Working leading candidate based on conjunction of
  three reasons (spatial reasoning bootstrap; computational
  pressure from NP-completeness; pedagogical curriculum starting
  point). Held with explicit revision conditions — a better
  candidate would replace it.
- **Substrate-level commitments 5-9** — pending triage:
  homeostatic boundary, dimensional minimum-commitment (tentatively
  absorbed into commitment 1), projection-and-anchoring, recursive
  predictive horizon, prime directives. Each will be sharpened
  through the same workflow already applied to commitments 1-4 and
  10.

### Research hypotheses and how we hold them

Several claims in the vault are filed as **strong-prior research
hypotheses** — held with substantial conviction based on current
reasoning, but explicitly awaiting evidence from construction work
to be supported, refined, or revised. Q-2026-04-28-011 (knowledge →
skills via nesting) and Q-2026-04-30-014 (3-D Tetris first-problem
candidate) are the program's first two such hypotheses; they
establish a working pattern.

The pattern matters because it lets the program be honest about
what's commitment vs. what's hypothesis. Substrate-level commitments
go into property statements with the discipline of "no claim
without grounding." Strong-prior hypotheses go into Q-artifacts
with explicit prior-confidence framing and explicit revision
conditions — a falsifying piece of evidence, an alternative
candidate, a better articulation. Both inform the program's work;
neither is treated as the other.

This is unusual epistemic infrastructure. We use it because the
distinction matters — for our own thinking, for honest external
communication, and for collaborators who deserve to know exactly
what we're sure about and what we're betting on.

### Inheritance from prior work

The architecture inherits from work already in the field:

- **MDEMG** (Multi-Dimensional Emergent Memory Graph,
  https://github.com/reh3376/mdemg) — the operational cognitive
  substrate this program builds on. MDEMG is the R&D vehicle whose
  18 months of operational experience surfaced what frozen-
  representation systems cannot do.
- **JEPA** (Yann LeCun and collaborators) — the commitment to
  embedding-space prediction is the right substrate; the embedding
  space alone is under-structured for the continuous-learning problem
  this program targets.
- **Numenta's Thousand Brains Project** (Hawkins, Ahmad, and
  collaborators, https://github.com/thousandbrainsproject/tbp.monty) —
  the cortical-column generalization principle is load-bearing for the
  architecture's claim that the same pattern recurs at every level.

The program does not claim to have the architecture finished or even
fully specified. It claims the problem is well-posed enough to work
on, and that the pieces are visible enough across existing research
that the synthesis is buildable.

---

## How the work is organized

The program runs in three phases. Phase boundaries are gradients, not
gates — work in a later phase can begin before the prior phase
formally completes when the work is ready.

### Phase A — Architectural Innovation Development (current phase)

The work to identify, articulate, sharpen, and specify the
architectural innovation. The deliverable is a specification
sufficient that people with the resources to build it can recognize
what it is and decide to back it.

This is where the program currently is. The work happens in cycles
that follow a documented methodology
(`threads/T005-program-planning/`); each cycle produces a
substantive artifact — a sharpened must-have, a research hypothesis,
a settled architectural decision — and closes when the cycle's
question has been answered or honestly identified as needing more
preliminary work.

### Phase B — Communication and Resource Acquisition

The work to articulate the Phase A specification in forms reaching the
audiences who can fund and staff implementation: papers, talks, pitch
documents, possibly working prototype demonstrations. Phase B does not
start until Phase A is materially complete; premature communication
artifacts make the architecture feel finished before it is.

### Phase C — Build with Team and Funding

Implementation, with the team and resources Phase B has secured. Out
of scope for the current planning structure; will require its own
program organization when reached.

---

## How the work is done

The architectural-innovation thread (T006) operates under an explicit
working method:

> *Roger provides the novelty. Claude (the AI research assistant) is a
> sharpening stone for Roger to hone his blade with. Roger provides
> what it is not, what it must have, what it could be, the mathematics
> necessary to prove it in theory.*

A sharpening stone removes burrs and dullness from a blade so it can
do what only the blade can do. The blade's metal, geometry, and edge
are the bladesmith's. The stone is patient, abrasive, and present.

In practice this means:

- Roger authors all architectural claims and chooses which directions
  to explore
- Claude pushes back, surfaces what's been tried, identifies gaps in
  candidate proposals, and drafts artifacts in vault-conformant form
- Claude does not propose candidate architectures or fabricate novelty
  when stuck — the working method is explicit that this is the failure
  mode the discipline protects against
- Architectural commitments live in the vault under
  `decisions/` and `threads/T006-architectural-innovation/plan.md`;
  research questions and hypotheses live in `questions/`; what has
  been explored and abandoned lives in `dead-ends/`

The full working-method declaration is in the T006 thread frame
(`threads/T006-architectural-innovation/thread.md`).

---

## Repository structure

```
ann-research/
├── README.md            ← you are here
├── INDEX.md             ← current state of understanding (read this second)
├── BRANCH_PROTECTION.md ← workflow and protection rules
├── CONTRIBUTING.md      ← how PRs are handled
├── LICENSE              ← MIT
├── cmd/vault/           ← Go CLI source
├── internal/            ← Go CLI internals
├── sidecar/             ← Python sidecar for scripts
├── observations/        ← OBS-* artifacts
├── questions/           ← Q-* artifacts
├── hypotheses/          ← H-* artifacts
├── decisions/           ← D-* artifacts
├── dead-ends/           ← DE-* artifacts
└── threads/             ← T*/thread.md per active research thread
```

### Artifact ID scheme

```
{TYPE}-{YYYY-MM-DD}-{NNN}-{slug}.md
```

For threads the date is omitted (threads have persistent identity):

```
T{NNN}-{slug}/thread.md
```

Examples:

- `OBS-2026-04-28-003-mdemg-subsystem-orientation-map.md`
- `D-2026-04-30-012-foundational-assumption-unifying-architectural-framework.md`
- `Q-2026-04-30-012-working-definition-of-general-intelligence.md`
- `T006-architectural-innovation/thread.md`

### Artifact metadata

Every artifact carries enough frontmatter to be understood out of
context:

- `where` — which research thread it belongs to
- `created` / `updated` — when it was written and last reviewed
- `status` — current epistemic standing (open, supported, falsified,
  superseded, etc.)
- `confidence` — how strongly the claim is held (1–5)
- `provenance` — sources behind the claim
- `supersedes` / `superseded_by` — explicit lineage when beliefs change

The validator enforces required fields per type. This is intentional
friction: it keeps half-thoughts from masquerading as research
artifacts.

---

## Tooling

A small Go CLI (`vault`) creates artifacts from templates, increments
sequence numbers, and validates frontmatter. Python sidecar scripts
live in `sidecar/` for tasks better suited to Python (text
manipulation, graph visualization of artifact dependencies, future ML
work on the artifacts themselves).

### Build the CLI

```bash
go build -o bin/vault ./cmd/vault
```

Requires Go 1.23 or newer.

### Use the CLI

```bash
# create artifacts
./bin/vault new obs    "monty vote orchestration"
./bin/vault new q      "what is the symbolic analog of pose"
./bin/vault new h      "categorical eta multipliers proxy for confidence"
./bin/vault new d      "use cmp message format for inter-column votes"
./bin/vault new de     "premature mapping table"
./bin/vault new thread "evidence accumulation deep dive"

# list and validate
./bin/vault list
./bin/vault list questions
./bin/vault validate
```

### Python sidecar

```bash
cd sidecar
uv sync
```

Currently empty by design. Scripts will appear here when there is a
specific need for them.

---

## Reading order for a fresh visitor

1. **`README.md`** — what this is (you just read it)
2. **`INDEX.md`** — what is currently believed and being worked on
3. **`threads/T006-architectural-innovation/thread.md`** — the central
   thread, with the working-method declaration in full
4. **`threads/T006-architectural-innovation/plan.md`** — the
   substrate-level commitments accumulated so far, with their
   sharpening status
5. **`decisions/`** — the architectural commitments that constrain the
   rest of the work
6. **`dead-ends/`** — failure modes already encountered

The `decisions/` and `dead-ends/` directories are the highest-value
reading. The former records what the program has committed to. The
latter prevents future work from relitigating settled questions or
repeating known failures.

---

## A note on the work in progress

This program is at the beginning of long-horizon work. The vault has
been active for less than a month. Phase A has produced a small number
of sharpened architectural commitments, a working definition of
general intelligence held as an open question, and a research method
applied with discipline. Substantial work remains.

If you are working on similar questions — continuous learning, world
models, reference frames, recursive self-improvement at the substrate
level — I would welcome the conversation. Reach out via the
contact information on https://github.com/reh3376.

---

## License

MIT. See [LICENSE](LICENSE).
