# Index — Current State of Understanding

> A fresh visitor (human or AI) should read this **second** (after `README.md`).
> It is the working-memory entry point and is rewritten periodically as
> understanding evolves.
>
> Last reviewed: **2026-04-28** (initial bootstrap)

---

## What we are working on

The research program is to develop a deep, mechanism-level understanding
of two systems — **Monty** (the Thousand Brains Project's sensorimotor
learning system, implementing cortical column theory) and **MDEMG**
(an operational graph-based cognitive substrate for AI coding agents) —
and from that understanding, determine whether and how their architectural
commitments can inform the design of novel ANN topologies for continual,
reference-frame-adaptive intelligence.

The current proposed framework name is **AHH** (Adaptive Hebbian
Hierarchy). The framework name may evolve.

## What we have decided

Four bootstrap decisions, all made on 2026-04-28:

- **`D-2026-04-28-001`** — Vault substrate is markdown + YAML frontmatter,
  one artifact per file, six artifact types, in git.
- **`D-2026-04-28-002`** — Repository at `github.com/reh3376/ann-research`,
  public, MIT licensed, Go CLI base + Python sidecar.
- **`D-2026-04-28-003`** — Branch protection: PRs required for assistant
  contributions; `@reh3376` admin bypass; CODEOWNERS = `@reh3376`.
- **`D-2026-04-28-004`** — Working agreement: no integration / interop /
  mapping work until `T001` and `T002` are complete.

## What we have ruled out

- **`DE-2026-04-28-001`** — Producing mapping tables or interop documents
  from structural reading alone. Specific failure mode: pattern-matching
  on shape rather than tracing mechanism. Terminal.

## What we are investigating

Three active research threads:

| ID | Status | Current question |
|---|---|---|
| **`T001-monty-trace`** | active | What does a single Monty episode actually do, mechanistically, from observation through motor action? |
| **`T002-mdemg-trace`** | active | What does an MDEMG ingest+retrieve cycle actually do, mechanistically? |
| **`T003-grounding-question`** | active | What does it mean for a symbolic substrate to have a "where am I" signal of the kind proprioception provides Monty? |

`T001` and `T002` are mostly trace work that proceeds through reading and
reporting. `T003` is synthesis work that depends on the trace threads
and on theoretical reading.

## Open questions

- **`Q-2026-04-28-001`** — Do Monty's algorithms exploit properties of
  continuous 3D geometry that have no useful analog in symbolic
  structure? *(load-bearing for whether integration is translation or
  substrate work)*
- **`Q-2026-04-28-002`** — What is the symbolic-substrate analog of
  Monty's `(object_id × pose)` hypothesis space?
- **`Q-2026-04-28-003`** — What is the empirical grounding for MDEMG's
  `etaConvMult` / `etaConfigMult` / `etaSameDirMult` learning-rate
  multipliers?

## Hypotheses under test

*None yet.* The current state is "we have questions, not hypotheses."
Hypotheses become appropriate once the trace work has produced enough
mechanism-level understanding to make falsifiable claims.

## What's load-bearing right now

The single most load-bearing artifact in the vault is
**`D-2026-04-28-004`** (the working agreement). It governs what kinds
of artifacts the assistant produces during the trace phase. If it's
violated, the program reverts to the failure mode in `DE-2026-04-28-001`.

The single most load-bearing thread is **`T003`**, because its answer
shapes what integration work — if any — eventually makes sense. But
`T003` can't progress meaningfully until `T001` and `T002` produce
evidence, so the practical near-term load-bearing is on those two.

## How this index is maintained

Rewrite this file when one of the following changes:

- A thread reaches `complete` or `abandoned`
- A decision is committed or superseded
- A hypothesis becomes supported or falsified
- A new thread is opened
- Roughly every two weeks regardless, to mark "still current as of"

The full per-artifact history lives in the vault folders and in
`git log`. This file is the projection of "what's relevant *now*."
