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
  - "D-2026-04-28-005 (trace target choice)"
owner: "Claude (drafts), @reh3376 (review)"
current_question: "What does each LM compute and emit during one matching_step in the five_lm_no_threading config, and what does each LM do with received votes on the next step?"
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

What does each LM compute and emit during one `matching_step` in the
`five_lm_no_threading` config (see `D-2026-04-28-005`), and what does
each LM do with received votes on the next step?

The first version of this question was broader ("what does an episode
do, end to end"). Narrowing to a single step gives the trace a
tractable target. Once one step is traced cleanly, expand to the next
step, then to multi-step patterns.

## Method

**Read-until-ambiguous, then run-and-instrument selectively.**

1. **Read pass.** Follow the source code path that processes a single
   observation through one `matching_step` in the chosen config:
   - `MontyBase.step()` (entry)
   - `MontyBase._matching_step()` and `aggregate_sensory_inputs()`
   - `EvidenceGraphLM.matching_step()` and `_update_evidence`
   - `EvidenceGraphLM.send_out_vote()`
   - `MontyBase._vote()` → `EvidenceGraphLM.receive_votes()` →
     `_update_evidence_with_vote`
   - `EvidenceGoalGenerator.step()` and `_generate_goal()`
   - `MotorSystem.step()` (or whatever the chosen motor policy uses)

2. **Document each step as observations.** One observation per
   mechanically distinct step. Each carries provenance to specific
   source file + line range. Confidence rated honestly: a 5 means I
   read the code and it does the thing I said it does; a 3 means I
   inferred from context and could be wrong.

3. **When ambiguous, run and instrument.** "Ambiguous" means: source
   reading cannot determine what a piece of code does without seeing
   actual values, OR the documented behavior conflicts with what the
   code appears to do, OR there is a specific runtime behavior (like
   evidence value scaling) that depends on data and can't be predicted
   from code alone. The fallback is to install Monty locally, run the
   chosen config, and capture the actual state. This is authorized by
   `D-2026-04-28-005` and does not need a separate decision.

4. **Write hypotheses for things the trace tentatively concludes.**
   Trace findings that aren't certain become hypotheses with explicit
   falsification criteria.

5. **Write questions for things the trace surfaces but can't answer.**
   Particularly: questions whose answers would require running the code
   or that point at the substrate-level concerns of `T003`.

## State

**Current.**

- Trace target chosen (`D-2026-04-28-005`): the
  `five_lm_no_threading.yaml` config, 5 LMs, fully connected vote
  matrix, deterministic, bounded.
- Structural observation recorded (`OBS-2026-04-28-001`): the vote
  topology and how `_vote()` orchestrates send/receive.
- **First mechanism observation recorded (`OBS-2026-04-28-002`):
  the matching_step pipeline as a six-phase orchestrator sequence,
  with the substantive properties (voting only in matching steps;
  `lm_to_lm_matrix=null` in our config means votes are the only
  inter-LM channel; the binary `use_state` gate).** One question
  surfaced (`Q-2026-04-28-004`) about why the motor system receives
  only `sensor_module_outputs[0]` rather than all SM outputs.

The trace is now inside the orchestrator at the level of "what does
each phase delegate to." The next layer down is what the SMs and LMs
actually do when invoked. Confidence in current mechanism
understanding: high for orchestrator (literally read), unknown for
SM/LM internals (not yet read).

**Open assumption** worth flagging (will become a question or
observation as the trace progresses): the chosen config presumes that
the test configs in `tbp.monty/main` are runnable without missing
dependencies or environment-specific shims. If running becomes
necessary and the test config doesn't run cleanly, the read-only path
becomes more load-bearing.

## Open work

- [x] Read `MontyBase.step()` and `_matching_step()`. Wrote
      `OBS-2026-04-28-002` (matching_step pipeline). Surfaced
      `Q-2026-04-28-004` (motor system / single-SM question).
- [ ] Read sensor module pipeline (`sensor_modules.py:CameraSM`,
      `ObservationProcessor`). Write observation covering what comes
      out of an SM in this config — i.e., what's actually inside a
      `sensor_module_outputs[i]` element after `sm.step()`.
- [ ] Read `EvidenceGraphLM.matching_step` and trace `_update_evidence`.
      This is where 3D geometry questions (`Q-2026-04-28-001`) will
      surface.
- [ ] Read motor system + active motor policy. Answers
      `Q-2026-04-28-004`.
- [ ] After the first three mechanism observations, decide whether
      reading is sufficient or whether we need to run-and-instrument.

## Artifacts produced by this thread

- `D-2026-04-28-005` — trace target choice
- `OBS-2026-04-28-001` — 5-LM vote topology (structural setup)
- `OBS-2026-04-28-002` — matching_step pipeline (first mechanism)
- `Q-2026-04-28-004` — motor system / single-SM question

## Log

### 2026-04-28 — first mechanism observation

Read `MontyBase.step`, `Monty._matching_step`, `aggregate_sensory_inputs`,
`_step_learning_modules`, `_collect_inputs_to_lm`, `_combine_inputs`,
`_pass_goals`, `_step_motor_system`, `_set_step_type_and_check_if_done`.

Wrote `OBS-2026-04-28-002` documenting the six-phase matching_step
pipeline with confidence-flagged properties: voting only in matching
steps; `lm_to_lm_matrix=null` in our config makes votes the sole
inter-LM channel; the `use_state` gate is binary, not weighted.

Surfaced `Q-2026-04-28-004` while reading `_step_motor_system`: the
motor system receives only `sensor_module_outputs[0]`, not all SM
outputs. Worth understanding before drawing conclusions about how
"where am I" is grounded.

Next step: read sensor_modules.py to understand what's inside
`sensor_module_outputs[i]`. After that, drop down into
`EvidenceGraphLM.matching_step` — that's where the geometry questions
(`Q-2026-04-28-001`) will start to bear on the trace.

### 2026-04-28 — trace target chosen

Read the experiment configs in `src/tbp/monty/conf/experiment/` and
the connectivity configs in `src/tbp/monty/conf/monty/connectivity/`.
Selected `five_lm_no_threading.yaml` as the first trace target;
documented in `D-2026-04-28-005`. Recorded one structural observation
about the vote topology (`OBS-2026-04-28-001`).

Next step: read `MontyBase.step()` and produce the first
mechanism-level observation. This will probably also be the first PR
that produces "real trace content" rather than scope-setting.

### 2026-04-28 — thread created

Thread created as part of bootstrap. Motivated by Roger's redirect
("develop a deep understanding before attempting to make any changes")
and the failure mode in `DE-2026-04-28-001`.

The agreement in `D-2026-04-28-004` is that no integration work
proceeds until this thread and `T002` reach `complete` status.
