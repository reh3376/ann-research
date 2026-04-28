---
id: T003
type: thread
slug: grounding-question
created: 2026-04-28
updated: 2026-04-28
status: active
confidence: 4
where: []
tags: [substrate, reference-frames, grounding, theory]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "Hawkins & Ahmad, 'The Thousand Brains Theory'"
  - "mdemg-collaboration-brief.md (project files)"
owner: "@reh3376 + Claude jointly"
current_question: "What does it mean for a symbolic substrate to have a 'where am I' signal of the kind proprioception provides Monty?"
---

# Thread T003 — The Grounding Question

## Goal

Resolve, as far as we can, the substrate-level question: is reference-
frame transfer between Monty's sensorimotor substrate and MDEMG's
symbolic substrate a **translation problem** (the algorithms are
substrate-agnostic; substituting coordinate systems suffices) or a
**substrate problem** (the algorithms exploit properties of continuous
3D geometry that don't transfer, and the borrowed abstractions would
be contracts without algorithmic content)?

This thread is the slowest of the three. It depends on `T001` and
`T002` for evidence and on theoretical reading and discussion for
synthesis. It's also the most consequential — its answer determines
the shape of any integration work this program eventually pursues.

## Current question

What does it mean for a symbolic substrate to have a "where am I"
signal of the kind proprioception provides Monty?

In Monty: proprioception (joint angles, sensor pose, motor state)
tells the system where the sensor is in space. That's what makes
sensed features meaningful: a "feature at location X" claim is only
useful if you know where X is.

In MDEMG: a code symbol or constraint exists in some structural
context. "Where am I" candidates include AST traversal state, git
position, call-graph location, agent/session identity, current
task. None of these feels equivalent to proprioception in the
load-bearing sense — they all feel more like *labels on the
observation* than like *self-knowledge of the observer*.

Is that asymmetry a translation detail or the actual story?

## Method

This thread doesn't have a clean "trace one thing" method. It will
proceed through:

1. **Evidence from T001 and T002.** As those threads produce
   observations about where geometry is and isn't load-bearing in
   Monty, and what the structural-graph analogs do or don't preserve
   in MDEMG, this thread accumulates the synthesis.

2. **Theoretical reading.** Specifically:
   - Hawkins / Ahmad / Mountcastle on cortical columns and reference
     frames (what *is* a reference frame, formally?)
   - Whittington & Behrens (TEM) — structure/content factorization,
     which is the most direct prior art for separating "where" from
     "what"
   - Friston / Da Costa on discrete-state active inference — does the
     FEP framework care about substrate continuity?

3. **Discussion.** This thread benefits from joint reasoning between
   `@reh3376` and Claude more than the trace threads do. The output
   is likely to take the form of decisions ("we hold X about
   substrate transfer") and hypotheses ("if Y, then Z would follow"),
   not just observations.

## State

**Initial state.** The question is articulated, no answer yet. The
asymmetry between proprioception and structural labeling is the
strongest prima facie case for "substrate problem" but is not
conclusive — it might be a feature of how we've described the
candidates rather than a genuine asymmetry.

## Open work

- [ ] Wait for early observations from `T001` and `T002` (no
      productive work in this thread until then).
- [ ] Reading: Whittington & Behrens 2020 (TEM), Hawkins &
      Mountcastle on reference frames, Da Costa et al. 2020 on
      discrete active inference.
- [ ] Reframe the "where am I" question more formally: what does a
      reference frame *minimally require*? Once stated formally, the
      question of substrate analog becomes more tractable.

## Artifacts produced by this thread

`Q-2026-04-28-001` and `Q-2026-04-28-002` both bear on this thread,
though they're primarily attached to `T001`. `T003` will produce its
own artifacts as the work progresses.

## Log

### 2026-04-28

Thread created as part of bootstrap. Roger flagged this as the
question he most wants joint reasoning on, distinct from the trace
threads which can largely proceed via reading and reporting.

Claude's prima facie suspicion (low confidence): the asymmetry is
substantive, not cosmetic — proprioception in Monty is about the
*observer's* state, while every "where" candidate in MDEMG is about
the *observation's* context. If that distinction holds up, it
suggests either (a) the symbolic substrate needs something playing
the observer-state role that we haven't yet identified, or (b) the
symbolic-substrate analog of a "reference frame" is structurally
different from Monty's.

This is a hunch, not a finding. It belongs in the log, not in a
hypothesis artifact, until there's something to falsify it against.
