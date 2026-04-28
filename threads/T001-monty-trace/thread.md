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
- One structural observation recorded (`OBS-2026-04-28-001`): the
  vote topology and how `_vote()` orchestrates send/receive.
- No mechanism-level trace work has begun yet. Next step is to read
  `MontyBase.step()` and write the first mechanism observation.

**Open assumption** worth flagging (will become a question or
observation as the trace progresses): the chosen config presumes that
the test configs in `tbp.monty/main` are runnable without missing
dependencies or environment-specific shims. If running becomes
necessary and the test config doesn't run cleanly, the read-only path
becomes more load-bearing.

## Open work

- [ ] Read `MontyBase.step()` and `_matching_step()`. Write
      `OBS-2026-04-28-NNN` covering what the orchestrator does on a
      single step.
- [ ] Read sensor module pipeline (`sensor_modules.py:CameraSM`,
      `ObservationProcessor`). Write observation covering what comes
      out of an SM in this config.
- [ ] Read `EvidenceGraphLM.matching_step` and trace `_update_evidence`.
      This is where 3D geometry questions (`Q-2026-04-28-001`) will
      surface.
- [ ] After the first three observations, decide whether reading is
      sufficient or whether we need to run-and-instrument.

## Artifacts produced by this thread

- `D-2026-04-28-005` — trace target choice
- `OBS-2026-04-28-001` — 5-LM vote topology (structural setup)

## Log

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
