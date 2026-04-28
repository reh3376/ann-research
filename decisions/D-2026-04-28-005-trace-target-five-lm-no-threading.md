---
id: D-2026-04-28-005
type: decision
slug: trace-target-five-lm-no-threading
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 4
where: [T001]
tags: [monty, trace, scope]
provenance:
  - "tbp.monty: src/tbp/monty/conf/experiment/test/evidence_lm/five_lm_no_threading.yaml"
  - "tbp.monty: src/tbp/monty/conf/monty/connectivity/5lm_5sm.yaml"
  - "tbp.monty: src/tbp/monty/conf/monty/learning_module/test_evidence_5lm_multithreading0.yaml (referenced by the experiment config)"
alternatives_rejected:
  - "test/evidence_lm/evidence.yaml (1-LM, no voting)"
  - "test/evidence_lm/five_lm.yaml (threaded; non-deterministic execution)"
  - "test/evidence_lm/five_lm_basic_logging.yaml"
  - "production benchmarks: ycb_10objs / ycb_77objs / randrot_noise variants"
---

# Trace Target: five_lm_no_threading

## Decision

The first Monty trace, opening work for `T001`, will use the experiment
configuration at:

```
src/tbp/monty/conf/experiment/test/evidence_lm/five_lm_no_threading.yaml
```

This configures a 5-LM evidence-based learning module setup with a fully
connected vote matrix, single-threaded execution, deterministic seed, and
bounded step counts (30 train / 30 eval / 60 total).

## Context

`T001`'s current question is "what does a single Monty episode actually do,
mechanistically, from observation through motor action?" Answering that
requires picking a specific experiment to trace. The Monty repository
contains roughly 40 experiment configs across `src/tbp/monty/conf/experiment/`,
ranging from 1-LM tutorials to multi-agent multi-modal benchmarks. The
trace target needs three properties:

- **Voting must be active.** The previous attempt (`DE-2026-04-28-001`)
  pattern-matched `lm_to_lm_vote_matrix` orchestration to MDEMG retrieval
  columns without understanding what voting actually does. The first trace
  has to exercise voting end-to-end.
- **Execution must be traceable.** Threading introduces non-determinism in
  observation → vote ordering. For the first trace, sequential execution
  makes the data flow inspectable.
- **Scope must be bounded.** A trace pass over 30 steps with 2-3 objects is
  tractable. A trace pass over 1000 steps with 77 objects is not.

## Alternatives considered

### `test/evidence_lm/evidence.yaml` (1-LM)

Rejected. Single learning module means no vote matrix, no
`receive_votes` / `send_out_vote` exercise. The voting abstraction is
load-bearing for the integration question and skipping it here would
leave the most-misunderstood mechanism untraced.

### `test/evidence_lm/five_lm.yaml` (threaded)

Rejected. Identical to the chosen target except `use_multithreading=True`
inside the LMs. The `_update_evidence` logic in
`evidence_matching/learning_module.py` (around line 770) launches one
thread per known object when computing vote updates. Thread interleaving
makes "what did LM_2 see when it computed evidence at step 17" a
non-deterministic question.

### `test/evidence_lm/five_lm_basic_logging.yaml`

Considered. Differs in the logging subsystem rather than in core
mechanics. The `detailed_debug_monty_runs` logger configured in
`five_lm_no_threading.yaml` already provides the diagnostic output needed
for tracing; the "basic_logging" variant would log less, not more.

### Production benchmarks (`ycb_10objs`, `ycb_77objs`,
`randrot_noise_*`)

Rejected for first trace. These configs are tuned for performance
measurement, not for inspectability. They use larger object sets, often
include randomized rotations and sensor noise, and have step budgets in
the thousands. Appropriate for later trace passes once the mechanism is
understood; the wrong scope for first contact.

## Rationale

The chosen config is deliberately a **test config**, not a production
benchmark. It exists in the Monty codebase precisely for the kind of
inspection-with-bounded-scope the first trace needs:

- `min_lms_match: 5` requires all LMs to converge before terminal,
  forcing voting to actually do work
- `seed: 42` makes the run reproducible, so a trace written today should
  match a trace written next week
- `max_total_steps: 60` upper-bounds episode length
- `multithreading0` in the LM config means evidence updates run in the
  main thread
- The connectivity matrix in `5lm_5sm.yaml` is a complete graph
  (`lm_to_lm_vote_matrix` has every LM voting to every other), so the
  topology itself doesn't introduce edge cases that obscure the
  voting mechanism

## Consequences

- T001 work proceeds against this config specifically. References to
  "the trace" in subsequent artifacts mean this config unless explicitly
  noted otherwise.
- The trace approach starts as **read-only** (read source + config to
  predict mechanism), with the fallback option to **run-and-instrument**
  if reading leaves ambiguity that source inspection can't resolve. See
  T001 thread.md for the methodological note.
- Other configs (1-LM, threaded, randrot/noise) become candidates for
  follow-up traces if findings from this one raise questions specific to
  those variants.

## Revisit conditions

- If `five_lm_no_threading.yaml` turns out to be broken or unrunnable in
  the current Monty main branch, fall back to the closest working config
  with the same properties (5 LMs, no threading, bounded steps).
- If reading the source cannot determine what some step does without
  actually executing the code, the trace expands to include
  run-and-instrument for that specific step. This does not require a new
  decision — it's already authorized by this one.
- If the trace reveals that single-threaded behavior diverges materially
  from threaded behavior, threading itself becomes a research question
  worth tracing separately.
