---
id: OBS-2026-04-28-002
type: observation
slug: monty-matching-step-pipeline
created: 2026-04-28
updated: 2026-04-28
status: recorded
confidence: 5
where: [T001]
tags: [monty, orchestrator, step-lifecycle, mechanism]
provenance:
  - "tbp.monty: src/tbp/monty/frameworks/models/abstract_monty_classes.py:69-90 (_matching_step)"
  - "tbp.monty: src/tbp/monty/frameworks/models/abstract_monty_classes.py:92-111 (_exploratory_step)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:163-191 (aggregate_sensory_inputs)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:233-258 (_step_learning_modules, _collect_inputs_to_lm)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:259-292 (_combine_inputs)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:305-326 (_pass_goals)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:328-340 (_step_motor_system)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:342-365 (_set_step_type_and_check_if_done)"
subject: "A Monty matching step is a strict six-phase sequence on the orchestrator. Voting happens only in matching steps, not exploratory steps. The five_lm_no_threading config has lm_to_lm_matrix=null so non-vote inter-LM signaling is disabled; vote is the only inter-LM channel."
---

# Monty matching_step pipeline

## What was observed

The `Monty._matching_step` method (defined on the abstract base class
in `abstract_monty_classes.py:69`) executes a fixed sequence of six
orchestrator phases. It is identical for all `Monty` subclasses;
`MontyBase.step` simply dispatches to `_matching_step` or
`_exploratory_step` based on `self.step_type`.

```python
def _matching_step(self, ctx, observations, proprioceptive_state):
    self.aggregate_sensory_inputs(ctx, observations, proprioceptive_state)
    self._step_learning_modules(ctx)
    self._vote()
    self._pass_goals()
    self._step_motor_system(ctx, observations, proprioceptive_state)
    self._set_step_type_and_check_if_done()
    self._post_step()
```

`_exploratory_step` is the same minus `self._vote()`. **Voting only
runs in matching steps, never in exploratory steps.** This is a
substantive property: during model-building, hypotheses are not being
adjudicated across LMs, so cross-LM voting is omitted.

### Phase-by-phase breakdown

**Phase 1 — `aggregate_sensory_inputs`** (`monty_base.py:163-190`).
For each SM in `self.sensor_modules`:
1. Pull the raw observation slice for this SM out of the environment's
   `observations` dict.
2. Update the SM with the agent's current `proprioceptive_state`
   (sensor pose).
3. Call `sm.step(ctx, raw_obs, motor_only_step)` → store result in
   `sensor_module_outputs[i]`.

Then for each LM in `self.learning_modules`:
4. Call `lm.get_output()` → store in `learning_module_outputs[i]`.

The LM outputs collected here are from the LM's *previous*
`matching_step`, since LMs haven't been stepped yet this round. This
is the temporal mechanism by which higher-level LMs receive input
from lower-level LMs in a heterarchy.

**Phase 2 — `_step_learning_modules`** (`monty_base.py:233-237`). For
each LM in order (sequential, not parallel at this level):
1. `_collect_inputs_to_lm(i)` — uses `sm_to_lm_matrix[i]` to pick which
   SM outputs go to LM_i, and `lm_to_lm_matrix[i]` (if non-null) to
   pick which LM outputs go to LM_i.
2. `_combine_inputs(...)` — filters inputs by `use_state=True`. If no
   SM input is `use_state=True`, returns `None` and the LM is
   effectively skipped this step (the "off-object / low-confidence"
   gate).
3. `getattr(lm, self.step_type)(ctx, sensory_inputs)` — calls
   `lm.matching_step(ctx, sensory_inputs)` (or `lm.exploratory_step`
   in exploratory mode). The string-dispatch via `getattr` is how the
   orchestrator runs the right LM method without an explicit branch.

**Phase 3 — `_vote`** (`monty_base.py:294-303`). Two synchronous
loops, already documented in [`OBS-2026-04-28-001`](../observations/OBS-2026-04-28-001-monty-five-lm-vote-topology.md).
Skipped entirely in `_exploratory_step`.

**Phase 4 — `_pass_goals`** (`monty_base.py:305-326`). For each LM,
call `lm.propose_goals()` (returns a `list[Goal]`). For each SM, call
`sm.propose_goals()` (returns `[]` by default, per the SM ABC). All
goals are aggregated into `self._goals`. **Reset every step** (with
the exception that `motor_only_step` does not reset, allowing prior
goals to persist into the next motor-only execution).

**Phase 5 — `_step_motor_system`** (`monty_base.py:328-340`). Calls
`self.motor_system(ctx, observations, proprioceptive_state,
self.sensor_module_outputs[0], self._goals)`. The motor system
returns `self._actions` directly. **Note that only
`sensor_module_outputs[0]` is passed** — see related question in
provenance.

**Phase 6a — `_set_step_type_and_check_if_done`**
(`monty_base.py:342-363`). Updates step counters, checks
`exceeded_min_steps`. Behaviour:
- TRAIN + matching: switch to exploratory_step
- TRAIN + exploratory: set `_is_done = True`
- EVAL + matching: set `_is_done = True`

**Phase 6b — `_post_step`** (`monty_base.py:365`). Empty hook in
the base class; subclasses can override.

## Why it might matter

Three properties surface from this pipeline that matter for the trace
and for the integration question, with confidence flagged:

1. **(confidence 5) Voting and exploration are mutually exclusive
   within a step.** If we were considering whether vote outcomes can
   drive exploratory model-building decisions, this rules out the
   intra-step pathway — they happen in different step types.

2. **(confidence 5) In the `five_lm_no_threading` config,
   `lm_to_lm_matrix` is null. This means LM-to-LM information flow
   happens *only through votes*** (see `OBS-2026-04-28-001` for the
   vote topology). Heterarchical (top-down or sibling-LM) signaling is
   architecturally available — `_collect_inputs_to_lm` reads it — but
   not exercised in this config. Worth keeping in mind: trace findings
   about inter-LM dynamics in this config are findings about
   **voting-only** dynamics, not about Monty's full inter-LM
   capability.

3. **(confidence 4) The `use_state` gate is structurally important
   beyond off-object filtering.** It's the boolean that decides
   whether an SM output participates in any LM's input, and whether an
   LM output participates in another LM's input. Monty's mechanism for
   "an LM has nothing useful to say right now" is not a low-confidence
   signal that gets weighted — it's a hard skip. This may be load-
   bearing for `Q-2026-04-28-001` (whether continuous-vs-discrete
   matters for the algorithms): one signal here is binary, not
   continuous.

## Things I cannot determine from this read alone

- **The actual content of an SM output.** `sm.step()` returns... what?
  Per `OBS-2026-04-28-001` provenance the message format is the
  Cortical Messaging Protocol `Message` class, but I haven't read
  `CameraSM.step` end-to-end yet. Next observation.
- **What `lm.matching_step` actually does internally.** The
  orchestrator calls it; the *real* mechanism (the one that bears on
  `Q-2026-04-28-001` and `Q-2026-04-28-002`) is inside that call.
  Several observations away.
- **What the motor system does with `sensor_module_outputs[0] +
  goals`.** Last in the trace.
