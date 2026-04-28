---
id: Q-2026-04-28-010
type: question
slug: long-horizon-prediction-beyond-jepa-window
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 4
where: [T004]
tags: [prediction, horizon, jepa, world-model, hierarchy, time-scale]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "V-JEPA 2 paper (arXiv:2506.09985) — practical prediction horizon ≤16 seconds for video planning"
  - "V-JEPA family limitations summary (emergentmind.com survey, December 2025)"
  - "D-2026-04-28-011 (V-JEPA as candidate visual faculty) — surfaces this gap"
  - "Q-2026-04-28-008 (recursive expansion of predictive horizon) — this is the specific case of that general commitment"
answer: ""
---

# Long-horizon prediction beyond JEPA's 16-second window

## The question

V-JEPA's practical prediction horizon is ≤16 seconds for video
planning. The recursive expansion of predictive horizon committed in
`Q-2026-04-28-008` requires the entity to extend prediction outward —
to seconds, then minutes, then hours, then longer time scales as
operational reach demands. V-JEPA alone does not deliver this.

What mechanism handles long-horizon prediction in the architecture,
above and beyond what V-JEPA provides at the visual-faculty layer?

## Why this matters

The recursive expansion of the predictive horizon is the entity's
*operating loop*, not a decorative feature. As described in
`Q-2026-04-28-008`:

> The entity uses its current model to predict; from prediction
> errors, it improves the model; the improved model produces better
> predictions, which let it act further out; acting further out
> produces new prediction errors, which feed back into improvement.

If the architecture's primary visual predictor caps at 16 seconds,
the entity can only act out 16 seconds visually. That's enough for
short-horizon manipulation tasks (the V-JEPA 2-AC robot demos) but
insufficient for:

- **Multi-step planning** that requires reasoning about what the
  world will look like in minutes, hours, or longer
- **Anticipating outcomes of decisions** whose consequences unfold
  over time scales beyond 16 seconds
- **Maintaining temporal coherence of identity** — knowing what
  "yesterday" or "next week" will mean for the entity's predictions
- **Long-horizon goal pursuit** — sustaining behavior toward goals
  that take more than 16 seconds to achieve

The 16-second limit is therefore an architectural constraint at the
*visual* faculty level, but the architecture must operate at longer
time scales overall. Something must close the gap.

## Candidate mechanisms

Three architectural candidates worth tracking, in order from least
to most ambitious:

### Candidate 1 — Hierarchical JEPAs at multiple time scales

Stack JEPAs operating at progressively coarser temporal resolutions.
Level 0: V-JEPA at 16-second horizon. Level 1: JEPA-shaped predictor
operating on Level 0's representations at minute-scale. Level 2:
hour-scale. Level 3: day-scale. Each level predicts in the
representations of the level below.

This is closest to how cortical hierarchy is hypothesized to operate
(slower time scales in higher-order cortex). It preserves the
predict-in-representation-space commitment uniformly. It is also
unproven at the temporal scales we'd need; nobody has demonstrated
JEPA-shaped architectures producing useful predictions at hour or
day scales.

### Candidate 2 — Symbolic / graph-based long-horizon model

Long-horizon predictions live in a different substrate from
short-horizon visual predictions. The superstructure (possibly
MDEMG-shaped) maintains a longer-horizon world model that operates
on summaries of V-JEPA's outputs and predicts at coarse temporal
granularity using non-JEPA mechanisms.

This is closer to how mainstream agent architectures handle
long-horizon planning. It introduces a substrate split, which has
costs (different formalism for each level) and benefits (each level
uses the right tool). It is also more aligned with MDEMG's existing
graph-based pattern.

### Candidate 3 — Language-mediated long-horizon prediction

The LLM faculty handles long-horizon prediction by reasoning in
language about future states. The superstructure invokes the LLM
for "what happens in the next hour given current state?" and treats
the LLM's output as a predicted future. The LLM's training data
includes vast amounts of human reasoning about long-horizon
outcomes, so the LLM has implicit long-horizon knowledge.

This is the lowest-engineering-cost option and has the most existing
machinery. It is also the most subject to the load-bearing-creep
failure mode flagged in D-010 — if long-horizon prediction is
offloaded entirely to the LLM, the superstructure has not actually
solved long-horizon prediction, it has just outsourced it.

## What the answer probably looks like

The architecture will likely combine elements of all three:

- **V-JEPA at the short horizon for visual prediction**, as
  committed in D-011
- **Some hierarchy or longer-time-scale predictor** at the
  representation level (Candidate 1 or some variant) for
  intermediate horizons
- **A symbolic/graph long-horizon model** at the superstructure
  level (Candidate 2) for the longest horizons, where the relevant
  predictions are about events, not pixels or representation patches
- **LLM consultation** as a faculty the superstructure can deploy
  for language-mediated reasoning (Candidate 3), but not as the
  load-bearing long-horizon mechanism

The exact split of which mechanism handles which time scale is open
and bears on T002 (MDEMG's existing patterns may map to longer
horizons), T004 (the entity's relationship to time as a first-
principles question), and on architecture-design work that follows
trace work.

## What it would take to answer

Resolution requires:

1. **Empirical evidence** on whether JEPA-style architectures scale
   to longer horizons via hierarchy. As of 2026-04-28 the literature
   does not contain such demonstrations at the time scales we need;
   the question may need to be answered via our own experiments.
2. **A trace-based picture of MDEMG's existing long-horizon
   capabilities.** Whether MDEMG already does anything resembling
   long-horizon prediction, and if so, at what time scales and via
   what mechanism. T002 trace work bears on this.
3. **An architectural specification of the time-scale split** —
   which faculty / mechanism handles which horizon, with explicit
   handoff between them. This is downstream of (1) and (2).
4. **A position on whether the LLM is allowed to be load-bearing
   for any horizon.** If yes, under what conditions and with what
   safeguards. If no, the architecture must produce non-LLM
   mechanisms for all horizons.

## Status

Open. Surfaced 2026-04-28 alongside D-011 (V-JEPA adoption). This
is the specific architectural-gap form of the more general
`Q-2026-04-28-008` (recursive expansion of predictive horizon) — Q-008
asks the question in principle, Q-010 asks it in the specific
context of V-JEPA's known limitation. Resolution of Q-010 is part
of resolving Q-008.
