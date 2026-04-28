---
id: D-2026-04-28-011
type: decision
slug: vjepa-as-candidate-visual-faculty
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 4
where: [T004]
tags: [architecture, faculty, vision, jepa, predictive-substrate, world-model]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude (post-PR-#8 merge)"
  - "@reh3376's framing: 'vision is only one of many reference frames, but if JEPA solves this mode then we should look at it for that reason alone'"
  - "Assran et al., 'V-JEPA 2: Self-Supervised Video Models Enable Understanding, Prediction and Planning' (arXiv:2506.09985, June 2025)"
  - "Meta AI, 'Introducing V-JEPA 2 world model and benchmarks' (June 2025)"
  - "V-JEPA 2.1 paper (arXiv:2603.14482, March 2026) — multi-modal tokenizer support and dense feature improvements"
  - "facebookresearch/vjepa2 (open-source code repository)"
  - "LeCun, Y. 'A Path Towards Autonomous Machine Intelligence' (2022) — original JEPA proposal"
  - "D-2026-04-28-010 (hybrid architecture — D-011 specifies one faculty within that architecture)"
  - "D-2026-04-28-009 (program objectives — V-JEPA's predict-in-representation-space aligns with mid-level capability commitments)"
alternatives_rejected:
  - "Pixel-space generative models for vision (Cosmos-style, autoregressive image models)"
  - "Custom-built JEPA-style architecture from scratch"
  - "Vision encoder from CLIP/SigLIP/DINOv2 family without predictive component"
  - "Defer the visual faculty decision until other architecture is more developed"
  - "Adopt JEPA-style approach for all sub-symbolic faculties simultaneously"
---

# V-JEPA as Candidate Visual Faculty

## Decision

The hybrid architecture committed in `D-2026-04-28-010` adopts the
**V-JEPA family** (specifically V-JEPA 2 and its successors) as the
candidate visual faculty. V-JEPA produces learned representations of
visual input that the superstructure operates over; it does not
itself constitute the cognitive core. The visual faculty is
analogous to the language faculty (LLM): a specialized capability
the superstructure orchestrates, not the entity's intelligence.

V-JEPA is adopted for the **visual modality only** at this commitment
stage. Whether JEPA-style architectures fit other sub-symbolic
faculties (proprioception, audition, novel modalities) is an open
question to be evaluated when those faculties are designed.

The commitment is at the **family** level (V-JEPA 2, V-JEPA 2.1, and
their successors), not the **specific-version** level. We commit to
the architectural approach; specific model selection happens at
implementation time and revisits as the family evolves.

## Context

`D-2026-04-28-010` committed to the hybrid architecture: learning
superstructure with the LLM as language plug-in. That decision did
not specify the *non-language* faculties. The architecture has
explicit slots for sub-symbolic predictive layers — the parts of the
entity that handle raw sensor input, learn representations, and
predict in those representations. D-010's mid-level capabilities
(reference frames, world models, continual learning, prediction)
require these sub-symbolic layers; D-010 named the requirement but
not the mechanism.

@reh3376 raised the JEPA question with appropriate framing
discipline: vision is one of many reference frames; if JEPA solves
the visual modality well, adopt it for that role specifically rather
than committing to JEPA as the universal sub-symbolic substrate.

Web research conducted 2026-04-28 confirmed:

- V-JEPA 2 (June 2025, 1.2B parameters, Meta) is state of the art
  for video understanding and motion anticipation
- V-JEPA 2.1 (March 2026, scaled to 2B parameters) introduces
  multi-modal tokenizers with shared encoder
- Both are open-source under research license at
  `facebookresearch/vjepa2`
- LLM integration via the LLaVA-style early-fusion pattern is
  demonstrated and benchmarked (PerceptionTest 84.0, TempCompass
  76.9, both at 8B parameter MLLM scale)
- Action-conditioning (V-JEPA 2-AC) achieved zero-shot deployment
  on Franka arms in unfamiliar labs, using <62 hours of unlabeled
  robot video
- Predict-in-representation-space is JEPA's core architectural
  commitment, directly aligned with D-010's predictive substrate
  needs

These verified facts make V-JEPA the strong candidate for the
visual faculty role.

## What V-JEPA does and does not provide

### What it provides

- **State-of-the-art visual representation learning** from video at
  scale, self-supervised
- **Action-conditioned prediction** suitable for embodied control
  (V-JEPA 2-AC variant)
- **Demonstrated LLM integration** via early-fusion vision-encoder
  pattern — the architectural fit for our LLM faculty
- **Multi-modal tokenizer support** in V-JEPA 2.1 (image and video
  in shared encoder; framework extends in principle)
- **Open-source availability** with weights and training code
- **Predict-in-representation-space architecture**, aligned with our
  predictive substrate commitments

### What it does not provide

- **Long-horizon prediction.** Practical prediction horizon is
  ≤16 seconds for video planning. The recursive predictive horizon
  expansion in `Q-2026-04-28-008` requires longer horizons that
  V-JEPA alone does not deliver. **This is a known architectural
  gap.** See `Q-2026-04-28-010` (filed alongside this decision) for
  the open question.
- **Multi-modal integration across distinct faculties.** V-JEPA 2.1
  supports multi-modal input *within JEPA* via shared tokenizers,
  but our architecture has separate faculties for separate
  modalities (V-JEPA for vision; possibly different mechanisms for
  proprioception, audition, novel sensors). The
  projection-and-anchoring problem in `Q-2026-04-28-007` is not
  solved by V-JEPA.
- **Catastrophic-forgetting-resistant continual learning.** General
  continual-learning problem applies; V-JEPA does not specifically
  address it. The LoRA-adapter approach committed in D-010 for the
  LLM faculty may extend to V-JEPA (LoRA on top of frozen V-JEPA
  for domain-specialization adapters), but this is speculative.
- **Language-goal specification.** V-JEPA's robotics deployments use
  image goals; language goals are not directly supported in
  current implementations. **Our architecture works around this:**
  the LLM faculty interprets language goals, the superstructure
  translates to representation-level goals, and those reach V-JEPA
  in its native goal format. The limitation that constrains
  standalone V-JEPA does not constrain our architecture.
- **Camera-position invariance and long-horizon error stability.**
  V-JEPA 2-AC shows camera-position sensitivity and incremental
  error accumulation in multi-step rollouts. These are
  implementation-level concerns; for our architecture's purposes,
  they suggest that the visual faculty may need supplementation
  (possibly a hierarchy at multiple time scales, possibly different
  mechanisms for different visual sub-tasks).

## Alternatives considered

### Pixel-space generative models for vision

Diffusion models, autoregressive image models, video generation
models like Cosmos. These predict actual pixels rather than
representations.

Rejected for the same reason LeCun originally proposed JEPA over
generative approaches: predicting pixels spends model capacity
reproducing irrelevant detail rather than learning useful structure.
For an entity that needs to *understand* visual input rather than
generate it, predict-in-representation-space is the right
architectural commitment.

### Custom-built JEPA-style architecture from scratch

Build our own JEPA-like predictive vision system rather than
adopting V-JEPA.

Rejected. V-JEPA's pre-training on 1M+ hours of video is an
inheritance the program does not need to re-derive. Just as the LLM
inherits the compressed prior of human-generated text, V-JEPA
inherits the compressed prior of human-generated video. Both are
assets the program adopts rather than rebuilds.

### Vision encoder from CLIP/SigLIP/DINOv2 family without predictive component

Use a strong learned visual encoder for representation, but without
the predictive component.

Rejected because the predictive component is the load-bearing piece
for our architecture. We need a faculty that produces representations
the superstructure can do prediction in. CLIP-family encoders give
representations but no predictive structure; we'd have to add the
predictive layer ourselves. V-JEPA already integrates representation
and prediction.

### Defer the visual faculty decision

Wait until trace work in T001 and T002 is more complete, and until
the superstructure architecture is more developed, before committing
to the visual faculty.

Rejected because the candidate is sufficiently strong that deferring
serves no real purpose. V-JEPA is well-benchmarked, open-source, and
clearly aligned with the architectural pattern. Deferring would mean
either (a) treating it as still-uncertain when the evidence is
solid, or (b) leaving the commitment in conversation rather than in
the vault. Neither serves the program. Revisit conditions in this
decision provide a graceful exit if it turns out wrong.

### Adopt JEPA-style approach for all sub-symbolic faculties

Commit to JEPA-shaped predictive-representation architectures across
all modalities (vision, proprioception, audition, novel sensors).

Rejected as architectural monoculture. JEPA is benchmarked,
deployed, and demonstrated for video. Other modalities may benefit
from similar patterns, but each modality's faculty should be
evaluated when designed. The substrate-flexibility commitment from
`Q-2026-04-28-009` argues against blanket commitments to specific
mechanisms for all faculties simultaneously.

## Rationale

Five reasons V-JEPA over the alternatives:

1. **Role-fit is clean.** The visual faculty's job is to take raw
   visual input, produce representations the superstructure can
   operate over, and support prediction in those representations.
   V-JEPA does exactly this.

2. **LLM integration pattern is demonstrated.** The LLaVA-style
   early-fusion pattern (vision-encoder patches projected to LLM
   token space) is the architectural fit for our LLM faculty. V-JEPA
   has been integrated with Qwen2-7B and Llama3-8B with strong
   results.

3. **Action-conditioning works.** V-JEPA 2-AC's zero-shot deployment
   on Franka arms with <62 hours of robot data is direct evidence
   for the embodied trajectory the program targets. This reduces
   risk in the long-term embodied-AI direction.

4. **Open-source availability removes the build-from-scratch tax.**
   The architecture inherits 1M+ hours of pre-training without
   running it ourselves.

5. **Predict-in-representation-space is the correct commitment.**
   D-010 implicitly assumes the architecture predicts at the level
   of learned features rather than raw signal. V-JEPA makes this
   explicit and operational.

## Consequences

- **The hybrid architecture from D-010 now has two specified
  faculties:** language (LLM) and vision (V-JEPA). Both are
  pre-trained foundation models the entity adopts. Both are
  orchestrated by the superstructure.
- **The faculty-as-orchestrated-tool pattern is reinforced.** Two
  data points (LLM, V-JEPA) is more than one. This becomes the
  default architectural posture for new faculties added in the
  future: select a strong pre-trained foundation, integrate via
  superstructure orchestration, treat as faculty not as core.
- **Q-008 has a partial implementation candidate.** V-JEPA gives us
  the short-horizon predictive substrate for vision. Long-horizon
  prediction beyond 16 seconds remains open and is captured as
  `Q-2026-04-28-010`.
- **Q-006 (minimum dimensional commitment) is partially served.**
  V-JEPA's representations are learned, not specified — the
  architecture does not bake in 3D Euclidean structure. The
  dimensional commitment is to *whatever V-JEPA learns from data*,
  which is more flexible than Monty's hardcoded SE(3) pose.
- **Q-007 (projection-and-anchoring) is unchanged.** V-JEPA does
  not solve cross-faculty integration; that remains a problem the
  superstructure must solve at its level.
- **Trace work direction shifts slightly.** T001's reading of Monty
  is now reading-for-mechanism-candidates more than
  reading-for-architecture, since the visual faculty role has a
  candidate. Specifically, Monty's voting and reference-frame
  mechanisms remain interesting, but Monty as a *vision system* is
  now in competition with V-JEPA, and V-JEPA wins on benchmarks
  and substrate flexibility.
- **The MDEMG/Monty/V-JEPA triple needs an integration story.** If
  MDEMG's graph substrate underlies the superstructure, and
  V-JEPA handles vision, and the LLM handles language, then the
  AI2AI protocol from D-010 must coordinate at least three
  components, not two. The protocol design therefore becomes a
  more complex specification than D-010 implied.
- **Continual learning at the visual faculty layer is open.** If
  the entity learns from new visual experience, V-JEPA itself must
  either be fine-tunable (with catastrophic-forgetting risk) or
  augmentable via adapters (untested for V-JEPA specifically). The
  LoRA-adapter pattern from D-010 may extend; this is a research
  task, not a settled answer.

## Revisit conditions

- **A non-JEPA visual architecture demonstrates substantively
  better performance** on benchmarks aligned with our requirements
  (motion understanding, action anticipation, action-conditioned
  prediction, LLM integration), with comparable or better
  open-source availability.
- **V-JEPA family scaling stalls** at the 1-2B parameter range
  without comparable advances in capability.
- **Long-horizon prediction limitations prove unworkable** even
  when supplemented by hierarchical or alternative mechanisms in
  the superstructure.
- **Cross-faculty integration via the AI2AI protocol fails** in
  ways specifically caused by V-JEPA's representation format.
- **Continual learning at the visual faculty layer proves
  fundamentally unworkable** with both fine-tuning and adapter
  approaches.
- **The architectural commitment to "faculty-orchestrated-by-
  superstructure" itself is revisited** at the D-010 level. If
  D-010 is revised, D-011 follows.

## What this is to the program

D-011 is the **first specified non-language faculty** in the hybrid
architecture. D-009 said what the program is for; D-010 said what
the program is building (hybrid pattern); D-011 begins specifying
the *components* within that pattern. Future architectural
decisions will follow this pattern: identify a faculty role,
evaluate candidate mechanisms, commit to the strongest with revisit
conditions.

The confidence rating (4, not 5) reflects that this is a candidate
commitment grounded in current state of the art. V-JEPA is strong
today; alternatives may emerge. The role-fit is clean and the
revisit conditions are explicit, so a graceful pivot is available
if needed.
