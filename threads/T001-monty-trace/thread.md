---
id: T001
type: thread
slug: monty-trace
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 4
where: []
tags: [monty, tbp, trace, mechanism]
provenance:
  - "tbp.monty repository: https://github.com/thousandbrainsproject/tbp.monty"
  - "conversation 2026-04-28 between @reh3376 and Claude"
owner: "Claude (drafts), @reh3376 (review)"
current_question: "What does a single Monty episode actually do, mechanistically, from observation through motor action?"
---

# Thread T001 — Monty Trace

## Goal

Develop a mechanism-level understanding of what Monty actually does
during one episode, by tracing a single observation end-to-end through
all relevant components: sensor module, learning module, hypothesis
space update, voting, goal generation, motor system, and back to the
next observation.

The goal is **not** to summarize Monty. The goal is to produce a
trace detailed enough that a reader can answer questions like "where
in the pipeline does the 3D geometry actually do work?" and "what
happens to confidence values as a vote propagates through the network
of LMs?" without having to re-read the source.

## Current question

What does a single Monty episode actually do, mechanistically, from
observation through motor action? Specifically:

1. Pick one benchmark experiment as the trace subject (likely something
   small from `benchmarks/` — exact choice TBD).
2. For one observation step in that experiment, document:
   - The raw observation entering the sensor module
   - What the SM computes and emits to the LM(s)
   - What `matching_step` does inside the LM (especially the
     `_update_evidence` math)
   - What `send_out_vote` produces (vote contents, scaling, threshold)
   - What `receive_votes` does to the receiving LMs' evidence
   - What the GoalGenerator decides and emits
   - What action the motor system executes
   - What the next observation looks like and how it differs from the
     first

## Method

Read-and-write loop:

1. Pick the benchmark.
2. Run it in an environment we control (likely a local Monty install)
   and instrument the loop to log enough state to reconstruct the trace
   from logs.
3. For each step of the trace, write **observations** (atomic, dated,
   sourced) into the vault. Each observation carries provenance back
   to specific source files and line numbers.
4. As patterns emerge, write **questions** for things we don't
   understand and **hypotheses** for things we tentatively conclude.
5. When the trace is complete enough to answer the current question,
   move thread status to `complete` and update `T003` with what we
   learned.

## State

**Initial state.** No trace work has begun. The `tbp.monty` repository
has been cloned and structurally surveyed (class hierarchies, file
layout, overall architecture) but no episode has been traced and no
mechanism-level understanding has been produced yet. The structural
survey is documented in `DE-2026-04-28-001` as the explicit failure
mode that motivated this thread.

## Open work

- [ ] Stand up a local Monty install (uv, follow the project's UV_PROTOTYPE.md)
- [ ] Pick the benchmark experiment for the trace
- [ ] Decide on instrumentation strategy (loggers? interactive notebook?)
- [ ] First trace pass: one observation, one step

## Artifacts produced by this thread

*None yet.*

## Log

### 2026-04-28

Thread created as part of bootstrap. Motivated by Roger's redirect
("develop a deep understanding before attempting to make any changes")
and the failure mode in `DE-2026-04-28-001`.

The agreement in `D-2026-04-28-004` is that no integration work
proceeds until this thread and `T002` reach `complete` status.
