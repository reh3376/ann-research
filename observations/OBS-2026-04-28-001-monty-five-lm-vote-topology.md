---
id: OBS-2026-04-28-001
type: observation
slug: monty-five-lm-vote-topology
created: 2026-04-28
updated: 2026-04-28
status: recorded
confidence: 5
where: [T001]
tags: [monty, voting, topology, connectivity]
provenance:
  - "tbp.monty: src/tbp/monty/conf/monty/connectivity/5lm_5sm.yaml"
  - "tbp.monty: src/tbp/monty/frameworks/models/monty_base.py:294-303 (the _vote method)"
subject: "Vote topology in the 5-LM connectivity config is fully connected (each LM votes to all 4 others), and the orchestrator routes votes via a hand-specified matrix rather than a learned one."
---

# Monty 5-LM Vote Topology

## What was observed

`src/tbp/monty/conf/monty/connectivity/5lm_5sm.yaml` defines the
connectivity for the 5-LM test configurations. Three matrices are
specified:

- `sm_to_lm_matrix` — maps sensor module index to learning module
  index. In this config, SM_i feeds LM_i (1-to-1) for i in 0..4. The
  view_finder is a sixth sensor module that does not feed an LM.
- `lm_to_lm_matrix: null` — no hierarchical (non-vote) inter-LM
  connections in this config.
- `lm_to_lm_vote_matrix` — the voting topology. **Fully connected.**
  Each LM votes to every other LM:

```yaml
lm_to_lm_vote_matrix:
  - [1, 2, 3, 4]   # LM_0 sends votes to LMs 1, 2, 3, 4
  - [0, 2, 3, 4]   # LM_1 sends votes to LMs 0, 2, 3, 4
  - [0, 1, 3, 4]   # LM_2 ...
  - [0, 1, 2, 4]   # LM_3 ...
  - [0, 1, 2, 3]   # LM_4 ...
```

The orchestrator at `monty_base.py:294-303` reads this matrix in `_vote()`:

```python
def _vote(self):
    if self.lm_to_lm_vote_matrix is not None:
        votes_per_lm = []
        for i in range(len(self.learning_modules)):
            votes_per_lm.append(self.learning_modules[i].send_out_vote())
        for i in range(len(self.learning_modules)):
            voting_data = [votes_per_lm[j] for j in self.lm_to_lm_vote_matrix[i]]
            self.learning_modules[i].receive_votes(voting_data)
```

Each step: every LM produces a vote (`send_out_vote()`), then every LM
receives the votes from its in-neighbors per the matrix
(`receive_votes(voting_data)`).

## Why it might matter

Three things to flag for the trace work, none load-bearing yet:

1. **Topology is hand-specified, not learned.** The matrix is a config
   parameter. Whether topology should adapt over training is not raised
   here; the test configs all use static fully-connected. This is a
   property of the *current* implementation, not necessarily a property
   of the architecture.

2. **`lm_to_lm_matrix` (the non-vote channel) is `null` in this
   config.** Monty supports another type of inter-LM connection beyond
   voting (likely hierarchical / heterarchy connections per the
   `connecting-lms-into-a-heterarchy` doc). For the first trace this
   means voting is the only inter-LM signal we'll see.

3. **The `_vote` orchestration is two synchronous loops,** not async
   message passing. Every LM completes `send_out_vote` before any LM
   begins `receive_votes`. This is a property worth remembering when
   the trace reaches voting, because it constrains what state any LM
   can have when it processes a vote (the sender's state at time of
   send, frozen).

This observation does not bear on `Q-2026-04-28-001` (3D geometry) or
`Q-2026-04-28-002` (hypothesis space). It is structural setup for
reading the trace.
