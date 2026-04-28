---
id: Q-2026-04-28-004
type: question
slug: motor-system-receives-only-first-sm
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T001]
tags: [monty, motor-system, sensor-module, orchestrator]
provenance:
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:328-340 (_step_motor_system)"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:163-190 (aggregate_sensory_inputs builds sensor_module_outputs in SM order)"
answer: ""
---

# Why does the motor system receive only sensor_module_outputs[0]?

## The question

In `MontyBase._step_motor_system`:

```python
self._actions = self.motor_system(
    ctx,
    observations,
    proprioceptive_state,
    self.sensor_module_outputs[0],  # only the first SM
    self._goals,
)
```

The motor system receives the orchestrator's full `observations` dict
and the full `proprioceptive_state`, but only **the first SM's
processed output** (`sensor_module_outputs[0]`) — not the full list of
SM outputs.

In the `five_lm_no_threading` config (the trace target per
`D-2026-04-28-005`), there are six SMs: five `patch_*` SMs feeding
five LMs, plus a `view_finder` SM that doesn't feed an LM at all. The
order in `self.sensor_modules` determines which one is index 0.

The question has three parts:

1. **Which SM ends up at index 0** in our trace config? Need to
   confirm by reading the SM-list construction. Likely `patch_0`, but
   `view_finder` is also a candidate if the convention places it first.

2. **Why only one SM's output, not all of them?** Possibilities:
   - The motor system is designed for a single sensor agent and
     multi-SM is a config we evolved into without rewiring the motor
     system fully (technical debt).
   - The motor system uses one SM's output to determine "where am I
     looking" and the rest are redundant for that purpose.
   - There's a config-aware motor policy that handles multi-SM
     differently and this single-SM passing is just a default.

3. **Does the answer change between motor policies?** Different
   `motor_policy_selectors` may make different use of the SM input. If
   the goal-driven policy ignores the SM output anyway and uses
   `self._goals` exclusively, then which SM is passed in doesn't
   matter for that branch.

## Why it matters

Surfaces because of `Q-2026-04-28-001` (whether 3D geometry is load-
bearing). If the motor system uses 3D pose information from one
specific SM to compute the next action, then "where the agent moves"
is grounded in one specific 3D coordinate frame. That's a small but
real piece of evidence that the substrate is geometric, not relational.

Also because of `T003` (the grounding question). If multi-SM
operation depends on a single SM being privileged for motor control,
the geometric "where am I" signal in Monty is *singular* — there's one
canonical agent pose, not a distributed/voted one. That's substantive
information about how Monty handles "where am I."

## What we'd need to answer it

- Read the SM list construction in the config build path to determine
  which SM is at index 0 in our trace target.
- Read the `MotorSystem.__call__` and the active motor policy
  (`policy_selector` config) to see what it does with the SM output.
- Check whether the motor policy uses the goals (`self._goals`) as the
  primary signal, with SM output as a fallback or context, or whether
  it weights them differently.

This is a sub-question of the matching_step trace and will likely
get answered when we read `motor_system.py` and `motor_policies.py`.
Capturing it now so the question is on record at the moment it
emerged.

## Status notes

Open. Surfaced 2026-04-28 while reading `MontyBase._step_motor_system`
for `OBS-2026-04-28-002`.
