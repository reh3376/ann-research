---
id: Q-2026-04-28-003
type: question
slug: mdemg-eta-empirical-grounding
created: 2026-04-28
updated: 2026-04-28
status: open
confidence: 3
where: [T002]
tags: [mdemg, hebbian-learning, parameters, empirical]
provenance:
  - "mdemg: internal/learning/service.go lines 350-460 (Hebbian update Cypher)"
  - "mdemg: internal/config/* (where these multipliers are likely defined)"
answer: ""
---

# What is the empirical grounding for MDEMG's etaConvMult, etaConfigMult, etaSameDirMult values?

## The question

MDEMG's Hebbian learning service applies a categorical multiplier to the
base learning rate `η` based on edge context:

```
etaMult = etaConvMult     if either node is conversation_observation
       * etaSameDirMult   if pathSim > 0.5
       * etaConfigMult    if obs_types differ AND one is config
```

Each multiplier is a numeric value loaded from configuration. Three
sub-questions:

1. **Where do these values come from?** Were they tuned against a
   specific benchmark? Chosen heuristically from intuition about
   relative importance? Inherited from an earlier version and never
   revisited?

2. **What work are they doing?** The categorical CASE structure
   suggests the design intent was to apply different learning rates
   to different *kinds* of co-activation. Is the categorical structure
   load-bearing, or is it an approximation of something more
   continuous?

3. **What would be lost by replacing them?** AHH note 02
   (precision-weighted Hebbian) proposes replacing categorical
   multipliers with continuous confidence values: `η_eff = η · c_i ·
   c_j`. To know whether that's an improvement or a regression, we
   need to know what the current categorical structure achieves that
   the continuous form would lose.

## Why it matters

AHH note 02 is one of the early-stage extensions in the roadmap. If
the categorical multipliers are doing important work that the
proposed continuous form wouldn't capture, note 02 needs revision. If
they're approximating a continuous form, note 02 is well-motivated.
We can't tell without empirical grounding.

This is also an instance of a more general question: **MDEMG has many
tuned parameters; for each, do we know what it's tuned against?**
This question is one entry; future tracing in T002 may surface others
that deserve their own artifacts.

## What we'd need to answer it

- Read the configuration system to find where these values are set and
  what their defaults are.
- Search git history (`git log -p` on the relevant config files) for
  the commits that introduced or tuned these values; commit messages
  may explain the empirical basis.
- Search MDEMG benchmarks (`internal/anomaly`, `neural/benchmarks`,
  any PR descriptions) for ablations that varied these multipliers.
- If no documented empirical basis exists: design and run an ablation
  ourselves, with the multipliers individually and jointly toggled.

## Status notes

Open. T002 trace work should produce at least the first three steps.
