---
id: OBS-2026-04-28-003
type: observation
slug: mdemg-subsystem-orientation-map
created: 2026-04-28
updated: 2026-04-28
status: recorded
confidence: 4
where: [T002, T005]
tags: [mdemg, orientation, architecture, subsystems, jiminy, j17, cms, rsic, uxts]
provenance:
  - "MDEMG repository at github.com/reh3376/mdemg, main branch HEAD 3322c9f (PR #361 merged)"
  - "VISION.md (634 lines, top-level)"
  - "AGENT_HANDOFF.md (1025 lines, top-level — current development state, phase registry, open work items)"
  - "README.md (29K, top-level — benchmark methodology, key features, project structure)"
  - "docs/architecture/00_README.md (documentation index)"
  - "docs/architecture/01_Architecture.md (services, flows, integration modes)"
  - "CMS.md (top-level, 24K)"
  - "docs/features/jiminy-inner-voice.md (29K)"
  - "docs/features/j17-ai2ai-protocol.md (57K — section 1-3.1 fully read; sections 4-8 read in extracted summary form covering RSIC cycles, comprehension metrics, continuous learning architecture)"
  - "docs/features/j17-prompt-compression.md (4.2K, fully read)"
  - "docs/features/neural-training-pipeline.md (19K — three training workstreams, generative LoRA pipeline)"
  - "docs/features/rsic-feedback-loop.md (12K — feedback loop, health weighting formula DH-005)"
  - "docs/uxts01/README.md and docs/specs/UXTS_PORTABLE_AGENT_SPEC.md (head only)"
  - "T005 Cycle 1 plan (orientation-only mode sequence)"
subject: "MDEMG is a substantially developed cognitive substrate (~227K Go LOC, 7 phases complete, 611 files, 2,381 tests, 150 API endpoints) organized as a Multi-Layer Emergent Memory Graph plus six functional subsystems plus a Modular Intelligence plugin layer. This artifact captures the structural skeleton without going to mechanism depth; subsequent cycles will go deep on individual subsystems."
---

# MDEMG Subsystem Orientation Map

## What was observed

This is the orientation deliverable for Cycle 1 of T005 program planning.
It does not go to mechanism depth on any single subsystem. It maps the
subsystems and their relationships at the level of "what does each one
do, where does it live, and how does it relate to the others," so that
subsequent cycles can deep-trace specific subsystems with confidence
about scope and boundaries.

## High-level framing

MDEMG is described in `VISION.md` as a *cognitive substrate* for AI
coding agents — a learning system whose higher-level concepts emerge
through Hebbian principles applied to a multi-layer graph. The
substrate has evolved through seven complete development phases into a
self-improving system with proactive guidance (Jiminy), AI-to-AI
communication (J17 Protocol), reflexive quality improvement (RSIC),
conversation-memory persistence (CMS), and a specification framework
(UxTS) for governing test contracts.

The system runs as a **Docker Compose stack of 5 services**:

| Service | Role |
|---|---|
| `mdemg` (Go) | API, retrieval, consolidation, RSIC orchestration |
| `neo4j` | Knowledge graph (5-layer memory hierarchy) |
| `timescaledb` | Time-series metrics + LLM interaction recording for training |
| `neural-sidecar` (Python) | ML services (re-ranking, NLI, tier prediction) |
| `grafana` | Observability dashboards |

LLM calls inside MDEMG follow a **no-tool-calling architectural policy**
— all 16 LLM call sites are single-shot structured-output or reasoning,
with nine banned patterns grep-audited each sprint. The fine-tuning
target is **Qwen3.6-35B-A3B MoE** (Apache 2.0, 35B total / 3B active,
256 experts, 262K context), with a two-tier MoE-Sieve LoRA strategy
(tier 1 universal: attention + shared expert; tier 2 per-family:
top-25% routed experts). This work is current as of sprint FT-LORA
(2026-04-21).

### Tech stack (from AGENT_HANDOFF.md §2)

| Component | Technology |
|---|---|
| Graph DB | Neo4j 5.x |
| Backend | Go 1.24 |
| gRPC | Protocol Buffers (`api/proto/*.proto`) |
| Embeddings | OpenAI `text-embedding-3-large` (3072d); vector index at 3072d |
| Plugins | Binary sidecar via gRPC Unix sockets (`plugins/*/`) |
| Neural Sidecar | Python FastAPI (`neural/`) — re-ranking + NLI |

### Internal package layout (`internal/`, from AGENT_HANDOFF.md §2)

22 packages, with the load-bearing ones for our orientation:

- `api/` — HTTP handlers + middleware
- `ape/` — Active Participant Engine + RSIC (covered in subsystem #3)
- `consulting/` — Agent consulting service (consult, suggest, constraints)
- `conversation/` — CMS (covered in subsystem #2)
- `db/` — Neo4j driver + schema validation + migrations
- `guardrail/` — MCP guardrail validation pipeline
- `hidden/` — Hidden layer abstraction/consolidation pipeline
- `jiminy/` — Jiminy inner voice guidance (covered in subsystem #4)
- `learning/` — Hebbian learning (`CO_ACTIVATED_WITH`)
- `llmclient/` — Unified LLM client (OpenAI/Ollama)
- `metalearn/` — Global meta-learning (cross-space promotion)
- `metrics/` — Prometheus metrics + determinism scoring
- `plugins/` — Plugin manager + scaffold
- `retrieval/` — Core retrieval pipeline (vector + activation + scoring + cache)
- `sanitize/` — Prompt injection sanitization + control char stripping
- `summarize/` — LLM summary service
- `symbols/` — Symbol extraction (tree-sitter)
- `transfer/` — Space transfer (export/import)

Plus `pkg/mdemg/` (public API types for external Go consumers) and
`neural/training/` (Python LoRA pipeline).

### Retrieval pipeline scoring formula (from AGENT_HANDOFF.md §2)

The actual scoring weights:

```
score = vector(0.55) + activation(0.30) + recency(0.10) + confidence(0.05)
      - hub_penalty(0.08) - redundancy(0.12)
```

This is the canonical retrieval scoring; the score is computed
per candidate after vector recall, symbol search, bounded expansion
(max depth=3), and spreading activation. Output goes through the
TTL-LRU cache before returning.

### Phase registry (from AGENT_HANDOFF.md §4)

The README's "all 7 phases complete" claim is confirmed and refined:

- **105 core phases** (numbered 31..105 with sub-phases) — all ✅
- **16 sidecar phases** (S0..S16) — all ✅
- **5 cognitive gap phases** (101..105) — all ✅ (gaps closed)

The sub-phase structure is detailed in AGENT_HANDOFF.md §4. Notable
patterns: phase 45 (Modular Intelligence) split into 45.1-45.5 with
45.3 cancelled (parser RPC researched, deemed bad idea); phase 47
(Incremental Updates) wired APE INGEST via StartScheduledSync; phase
46-PR (Dynamic Pipeline Registry) is the mechanism for automatic
layer promotion triggers; phase 80 is CMS Meta-Cognition.

### Fine-tuning state (from AGENT_HANDOFF.md §5, current 2026-04-21)

The FT-LORA sprint state, since this is operationally active and
relevant to D-010's LoRA-managed continuous-learning commitment:

**Local LoRA pipeline** (PRs #246-250, all complete):
- Phase 1: LLM Interaction Logger (TSDB writer + scrubber + quality
  pipeline + data CLI; 16 consumers labeled via `WithContext()`)
- Phase 2: Think Mode + Response Sanitization (`StripThinkBlock +
  StripCodeFence`, 11 call sites)
- Phase 3: vllm-mlx Integration
- Phase 4+: SFT/GRPO/DPO local-LoRA training

**Hosted OpenAI fine-tuning track** (FT-OAI):
- **FT-OAI-001 ✅ complete (2026-04-21)** — first production fine-tune
  on `gpt-4.1-mini`. Result: +0.032 mean cosine, 7.8:1 W/L, parse-pass
  preserved at 0.973. **Verdict: MARGINAL.** Cross-base bench against
  `gpt-5.4-mini` showed Δ=−0.034. Strategic read: closed ~48% of the
  stock-4.1-mini → stock-5.4-mini gap (0.832 → 0.864 → 0.898; 0.0319
  of 0.0658). Deploy on hold pending per-token cost ratio analysis.
- **FT-OAI-002 ✅ complete (2026-04-21)** — pure tooling + telemetry +
  investigation hardening. E1-E9 epics; E3 (cap-symmetric baseline
  re-eval) + E8 (10-record integration smoke) staged pending
  live-spend authorization. No re-train; FT-OAI-001's 0.864 stands.
- **FT-OAI-003 📋 planned** — north-star sprint to close the remaining
  ~52% of the gap, target `FT(cheap-base) ≈ prod(gpt-5.4-mini)` at
  the cheap base's inference cost. 10 epics + rollback. Cost cap
  ≤$250. Quality floor ≥0.8322.

This is significant for our architectural work because it shows the
**LoRA-managed continuous-learning surface (D-010 commitment)
operating in production**, with explicit quality gates, cost caps,
and rollback plans. The pattern is more disciplined than D-010
implied; specifications can adopt FT-OAI's gating discipline as a
reference.

### Architecture maps for Jiminy context injection

A pattern worth flagging: `docs/architecture/maps/` contains 10
compact architecture maps that are **auto-generated** by
`scripts/generate_arch_maps.py` (with `--checksum`, `--dry-run`,
`--force` modes). These maps are then consumed by Jiminy as
context injection — meaning the system documents itself for its
own use. UITS optimization protection (`metadata.optimized: true`)
prevents the generator from overwriting maps that have been
manually converged.

This is directly relevant to D-010's superstructure: the entity
having structural self-knowledge that's machine-readable and
machine-maintained. Worth cross-referencing in future architectural
specification work.

---

## The seven major components

The structural decomposition. Each is a candidate for its own deep-trace
cycle.

### 1. The Multi-Layer Emergent Memory Graph (MEMG itself)

**What it is.** A 5-layer (L0..L5) Neo4j graph with vector indexes,
adaptive DBSCAN clustering parameters that loosen at higher layers,
and Hebbian learning rules over typed edges.

**Layer constraints.**
- Minimum: 1 layer (raw observations only)
- Maximum: unconstrained (hardware-limited)
- Growth: dynamic — layers emerge as data density warrants
- Adaptive parameters: `eps = base_eps * (1 + 0.4 * layer)`,
  `minSamples = max(2, base_min - layer)`. Higher layers cluster
  more freely.

**Key labels.** `:TapRoot` (singleton per `space_id`), `:MemoryNode`
(main nodes, has `embedding`), `:Observation` (append-only events),
`:SchemaMeta` (schema version tracking).

**Key relationships.** `ASSOCIATED_WITH` (semantic similarity),
`CO_ACTIVATED_WITH` (Hebbian learning signal), `CAUSES`, `ENABLES`,
`ABSTRACTS_TO` / `INSTANTIATES` (layer hierarchy), `CONTRADICTS`
(used by Jiminy).

**Hebbian learning extensions.** Tanh soft-capping on edge weights,
squared activation amplification, multi-rate learning rates by
edge type, time-based LR schedule, cautious decay of
frequently-activated edges, local-first activation spreading,
value residual bypass for high-confidence nodes, L0 skip
connections, explicit negative-result tracking, frontier
detection. (Source: VISION.md "Hebbian Learning Optimizations".)

**Technical invariants** (named as inviolable):
- Vector index = recall (fast candidate generation)
- Graph = reasoning (typed edges with evidence)
- Runtime = activation physics (computed in-memory, not persisted)
- DB writes = learning deltas only (bounded; no per-request
  activation writes)

**Where to look in the codebase.** `docs/architecture/02_Graph_Schema.md`,
`docs/architecture/04_Activation_and_Learning.md`,
`internal/retrieval/`, `internal/ingest/` (paths inferred from
description; need verification on next read).

**Deep-trace candidate.** Yes — the substrate everything else stands
on. Comparable in importance to Monty's matching_step pipeline.

### 2. CMS — Conversation Memory System

**What it is.** Persistent memory for LLM coding agents across context
window boundaries. Captures significant conversational events as
structured observations, stores them in Neo4j with semantic
embeddings, surprise scores, and quality assessments, and restores
the most relevant context on session resume.

**Core problems addressed.** Context loss on compaction (~every 20-30
minutes), poor context selection on resume, signal-vs-noise in
observations, multi-agent isolation, cross-session continuity,
memory degradation over time.

**Memory lifecycle.**
1. Observe — significant events captured during session
2. Store — embedding + scores + persist in Neo4j
3. Resume — retrieve most relevant on session start
4. Reinforce — co-activation strengthens edges (Hebbian)
5. Self-improve — RSIC assesses health between sessions

**Observation types** (4 priority tiers): HIGHEST (correction); HIGH
(error, blocker, decision); MEDIUM (preference, learning, task,
insight); LOW (progress, context, technical_note).

**Volatile vs permanent.** New observations start volatile (stability
~0.1). Co-activation reinforces (+0.15). Stability > 0.8 graduates
to permanent. Stability < 0.05 tombstones (removed). Mimics
biological consolidation.

**Resume relevance scoring.** 40% importance (type+surprise+stability)
+ 30% recency (24h half-life exponential decay) + 30% task
relevance (embedding similarity).

**Surprise detection.** Correction patterns ("No...", "Actually..."),
term novelty, embedding distance, contradiction with stored
knowledge.

**Where in the codebase.** `internal/conversation/` package —
`service.go`, `cooler.go` (volatile→permanent graduation),
`quality.go`, `dedup.go`, `relevance.go`, `truncation.go`,
`templates.go`, `snapshot.go`, `org_review.go`,
`session_tracker.go`, `types.go`.

**Deep-trace candidate.** Yes — closely related to and entangled with
RSIC. Worth tracing together.

### 3. RSIC — Recursive Self-Improvement Cycle

**What it is.** The mechanism by which MDEMG evaluates and improves
the quality of its own knowledge. A 5-stage cycle
(Assess→Reflect→Plan→Execute→Validate) operating at three temporal
scales.

**Time scales.**
- Micro (minutes): immediate quality checks after learning events
- Meso (hours): aggregate effectiveness analysis, tier drift
- Macro (days): structural health, cross-space consistency,
  calibration review

**Stage outputs.**
- Assess: 7-dimension health score (retrieval, memory, edge, task,
  guidance, protocol, synergy)
- Reflect: pattern detection across 19 reflection patterns (low
  comprehension, tier drift, calibration drift, volatile backlog,
  synergy health, etc.)
- Plan: prioritized improvement tasks with estimated impact
- Execute: automated actions (tier threshold adjustment, code
  retirement, constraint archival) with safety gates
- Validate: before/after comparison ensuring changes improved
  rather than degraded quality

**Safety mechanisms.** Dry-run mode, rollback snapshots, confidence
thresholds, cooldown policies, calibration tracking per action.

**Where in the codebase.** `internal/ape/` package (Active Participant
Engine) — `types_rsic.go`, `self_assess.go`, `self_reflect.go`,
`improvement_plan.go`, `task_spec.go`, `task_dispatch.go`,
`task_monitor.go`, `calibration.go`, `watchdog.go` (4-level
escalation), `cycle.go` (orchestrator), `orchestration_policy.go`,
`safety_validator.go`, `action_snapshot.go`, `rsic_store.go`.
HTTP handlers in `internal/api/handlers_self_improve.go`.

**Phase specs.** `docs/specs/phase87-rsic-orchestration-activation.md`,
`phase88-rsic-safety-policy-enforcement.md`,
`phase89-rsic-persistence-multi-space.md`,
`phase90-rsic-conformance-ci-gating.md`,
`phase91-rsic-observability-operations.md`.

**Closed-loop feedback (AR-1).** RSIC was originally open-loop —
actions ran but their effectiveness wasn't measured. Phase AR-1
(2026-03-30) added closed-loop semantics: post-cycle re-assessment
captures `MetricsAfter`; per-task `SuccessCriteria` are evaluated
(comparing `MetricsAfter` to `MetricsBefore` against thresholds);
auto-rollback reverses **reversible actions** when criteria fail.

**Reversible vs irreversible actions.**

| Reversible | Irreversible |
|---|---|
| `tombstone_stale` | `prune_decayed_edges` |
| `graduate_volatile` | `prune_excess_edges` |
|  | `trigger_consolidation` |
|  | `refresh_stale_edges` |

For irreversible actions, failure is logged as `completed_no_benefit`.
Calibration confidence only increases when both task-completed AND
criteria-met. Rollback events emit `mdemg_rsic_rollbacks_total`
counter with `{action, reason}` labels.

**DH-005 health-weighting formula** (substantial finding — load-bearing).
Overall health is a *confidence-weighted* sum across 7 dimensions:

```
overall_health = Σ(w_i · c_i · s_i) / Σ(w_i · c_i)
```

where `w_i` = base prior weight, `c_i` = data-sufficiency confidence
∈ [0,1], `s_i` = dimension score ∈ [0,1]. **Dimensions with `w_i ≤ 0`
or `c_i ≤ 0` are automatically excluded, not penalized.** This is
architecturally significant: the metric system is *honest about what
it doesn't know* rather than fabricating low scores from missing
data.

Default weights reflect a hybrid of *reliability* (how trustworthy
is the score?) and *user-impact* (how visible in operation?):

| Dimension | Default weight | Confidence threshold |
|---|---:|---|
| ProtocolHealth (J17 composite) | **0.20** | 30 J17 events |
| TaskPerformance | **0.20** | 50 observations |
| GuidanceHealth (Jiminy) | **0.17** | 30 guidance events |
| MemoryHealth | **0.15** | 100 nodes |
| EdgeHealth | **0.15** | 50 edges |
| RetrievalQuality | **0.08** | LearningPhase ≥ warm |
| SynergyHealth | **0.05** | Jiminy healthy + files |

Operator-tunable via `RSIC_HEALTH_WEIGHT_*` env vars without
rebuild. The formula structure means dashboard thresholds (red <0.4,
green ≥0.7) remain meaningful regardless of which dimensions are
present — a powerful pattern for evolving systems where capabilities
come online incrementally.

**This pattern is directly applicable to the entity's homeostatic
boundary** (Q-005). The entity's self-model can use the same formula:
each homeostatic dimension contributes proportional to weight × confidence
× score, missing dimensions are excluded not penalized. As the entity
matures and more aspects of self become measurable, the formula
gracefully extends.

**Deep-trace candidate.** Yes — directly relevant to D-009's
"recursive self-improvement" framing and Q-008's predictive horizon
expansion. Bears on T004 framing significantly.

### 4. Jiminy Inner-Voice Service

**What it is.** Pre-prompt guidance injection. Every user prompt is
intercepted via Claude Code hooks (`prompt-context.sh`); Jiminy
runs a 5-source parallel fan-out (6-second timeout), merges
results, and produces an injectable `═══ JIMINY GUIDANCE ═══`
block that arrives before the agent sees the prompt.

**The five sources** (parallel goroutines):
- A: Consulting Suggestions (`consulting.Suggest()`) — constraints,
  conflicts, patterns, concepts
- B: Correction Vector Search — past corrections via semantic
  similarity
- C: Contradiction Edge Lookup — `CONTRADICTS` edges between
  relevant nodes
- D: Frontier Surfacing — `is_frontier=true` nodes (thin knowledge
  areas)
- E: Full Retrieval Pipeline (J7) — all 8 obs_types, L0-L5 concepts,
  Hebbian edges, 14-component scoring

**Processing pipeline.** Generate embedding → fan out → lock-protected
merge → filter by min confidence → deduplicate by content → sort
by priority then confidence → truncate to max items → escalation
effects → LLM synthesis (J8, optional) → format prompt
augmentation.

**Graceful degradation.** If embeddings fail, vector-based sources
(B, C, D) skipped; A and E still work. Late arrivals past timeout
dropped silently. LLM synthesis falls back to static formatting on
error.

**Effectiveness tracking.** Every guidance item can receive follow-up
feedback (followed / contradicted / ignored). This data feeds RSIC
self-calibration.

**Operational character.** "Conscience" framing — not a tool the agent
chooses to use; a persistent voice that speaks before every
action. This is the architectural inversion of retrieval-on-demand
(by the time the agent asks, it has often already made the
mistake).

**Where in the codebase.** `internal/jiminy/`. Phase Jiminy:
J1-J16 implementing source A-E and J7-J16 features (J7 fan-out,
J8 LLM synthesis, J9 evaluation, J11-J12 escalation).

**Deep-trace candidate.** Yes — this is D-010's teacher-pupil pattern
in operational form. The architectural pattern in D-010 is named
as inherited from this; understanding Jiminy is understanding the
precedent.

### 5. J17 AI-to-AI Protocol

**What it is.** The compact, self-improving AI-to-AI protocol that
replaces verbose natural-language guidance with a three-tier
encoding. Designed for the Jiminy guidance use case but
generalizable. Achieves 5.2x token compression while maintaining
100% comprehension accuracy (10.0/10 T1 mean score) on baseline
benchmarks.

**Three tiers.**
- T1 (Coded, ~15-25 tokens, 80% of traffic): `TYPE:SEV|code|[ann]|esc:N|src:NODE_ID`
  - Type prefixes: C=constraint, X=correction, F=frontier,
    D=decision, P=pattern, L=learning
  - Severity: !=must/must_not, ?=should/should_not, ~=info
  - Annotations (most-actionable-first): alt:, neg:, scope:, ctx:
- T2 (Telegraphic, ~50-100 tokens, 15% of traffic): abbreviated
  natural language
- T3 (Full NL, ~80+ tokens, 5% of traffic): complete explanation
  with rationale

(Documentation also references T0 and T4 in some places — five-tier
naming. T0/T4 not yet investigated; question filed.)

**Bootstrap header.** Sent once per session. Includes encoding
specification + dynamically-built `DICT:` glossary of active
constraint codes (~15 words per code). The DICT eliminates the
primary T1 failure mode (codes interpreted literally rather than
as mnemonic references).

**State persistence — Session Tickets.** HMAC-SHA256 signed payloads
that survive context resets. Carry: version, space_id, session_id,
last_seq (replay detection), trust_score, escalation_snapshot
(per-constraint level + ignore count), active_constraint_ids,
conversation_phase, issued_at, ttl (default 168h / 7-day).

**Renewal triggers** (event-driven, not periodic): pre-compaction
hook (~20-30 min), session start (context reset / new session),
escalation state transition (immediate), constraint set mutation
(next hook cycle).

**Graceful fallback.** Invalid or expired ticket triggers full
re-guidance, not an error. Agent still receives guidance, just
without state continuity.

**RSIC-driven evolution.** Tier thresholds adjust based on measured
comprehension; ineffective codes retired; new codes generated from
frequently-repeated T2 constraints. This is the protocol *learning
to communicate better* through measured outcomes.

**Comprehension feedback.** NLI-based scoring via the neural-sidecar.
Per-tier effectiveness grading, calibration tracking, RSIC drift
detection.

**Design axioms.**
- No size constraints (compactness by design, not by limit)
- Human readability not a design factor (auditability via logging)
- Encryption: signing only (HMAC-SHA256); transparency is a goal
- Living protocol (RSIC-driven evolution)

**Self-improvement journey (sections 4-5 of the protocol doc).** The
three-tier protocol's current 91% perfect comprehension and 5.2×
compression are *not* initial values — they are the result of four
documented RSIC self-improvement cycles:

- **Baseline:** T1=8.3/10, compression 1.9×, perfect 14/23 (61%).
  Five identified failure modes: model-name mangling, scope
  softening, trade-off context loss, negation-only codes,
  implementation detail loss.
- **Cycle 1 (3 iterations, code-level):** Better mnemonics via LLM
  regeneration. T1 reached 9.1, then plateaued.
- **Cycle 2 (3 iterations, code-level + server feedback):** Same
  loop with server-side metric ingestion. T1=8.9. Still plateaued.
- **Cycle 3 (1 iteration, protocol-level):** Three protocol
  enhancements — DICT glossary in bootstrap header, inline
  annotations, action-focused codes. **Zero code changes needed
  for immediate convergence to T1=9.9, perfect 21/23 (91%).**
- **Cycle 4 (4 experiments, T1 compression):** Three compression
  levers. T1=10.0 at 5.2× compression (19% of original).

The lesson stated in the doc: "when individual codes fail, the
protocol itself must evolve — better codes within a weak protocol
hit a ceiling, but a stronger protocol makes even simple codes
work." This is exactly the recursive expansion of predictive horizon
in `Q-2026-04-28-008` — when local optimization plateaus, the
optimization moves up a level.

**ProtocolHealth dimension formula.** RSIC's protocol dimension is
itself a four-component composite:

```
ProtocolHealth = 0.40 × comprehension     (codes understood?)
               + 0.25 × compression       (saving tokens?)
               + 0.20 × coverage          (constraints have T1 codes?)
               + 0.15 × stability         (tickets restoring, replays infrequent?)
```

**Five reflection patterns + four mutation actions** (RSIC-driven
protocol evolution, all with auto-rollback):

| Pattern | Trigger | Mutation |
|---|---|---|
| `j17_codification_opportunity` | T2 sent >30 times | `codify_constraint` |
| `j17_low_comprehension` | code <70% | `retire_code` |
| `j17_high_replay` | replay >5/hr | `adjust_replay_buffer` |
| `j17_compression_regression` | ratio <2.0× | `adjust_tier_threshold` |
| `j17_low_code_coverage` | <80% coded | `codify_constraint` |

**JSONL training data accumulation** for ML-powered tier selection
(Phase J17-5, planned). Every protocol event records constraint
code, tier used, token count, outcome, comprehension score, trust
score, session id. The trained model will replace rule-based tier
selection with learned per-constraint optimum.

**J17-PC: Internal LLM Caller Prompt Compression** (separate but
related feature). J17's compression utilities (`TelegraphicCompress`,
`CodedEncoder`) applied not only to outgoing AI2AI guidance but to
the **inputs** of the 5 highest-value internal LLM callers. Per-caller
opt-in via `*_COMPRESS` env vars. Prose-only compression with code
diffs, JSON schemas, and enum values verbatim. **Aggregate 25-35%
token reduction with zero quality loss.**

| Caller | Savings | Strategy |
|---|---|---|
| RSIC LLM Reflector | 40-50% | CompactJSON, criteria truncation |
| LLM Reranking | 30-40% | Summary truncation 300ch, pipe-separated |
| SME Synthesis | 25-35% | CompressSection(400), concepts capped at 10 |
| Guardrail Evaluation | 20-30% | Compact prompts (17→6 lines) |
| Outcome Classification | 20-30% | Compact prompts (14→3 lines) |

**Where in the codebase.** Implementation locations not yet confirmed
by reading the source. Documentation: `docs/features/j17-ai2ai-protocol.md`,
`docs/features/j17-feedback-loop-closure.md`,
`docs/features/j17-prompt-compression.md`,
`docs/research/ai2ai/` (6 research docs: protocol survey,
encoding schemes, emergent communication, state persistence,
synthesis, recommendations).

**Deep-trace candidate.** Yes — directly fills D-010's AI2AI protocol
commitment. The hybrid architecture's inter-faculty communication
specification depends on understanding J17 thoroughly. This is the
priority-1 deep-trace target.

### 6. UxTS — Universal-x Test Specification framework

**What it is.** A specification methodology, not a runtime subsystem.
A pattern for organizing programmatic verification into
domain-specific frameworks (UATS for APIs, UPTS for parsers, UBTS
for benchmarks, UNTS for hash verification, etc.) that share a
common four-layer architecture: Schema → Specs → Runner → CI Gate.

**Core concepts.**
- Spec: declarative JSON file defining a verifiable contract (data,
  not code; diffable, reviewable, generatable)
- Schema: JSON Schema defining the valid structure of specs in a
  framework (canonical authority)
- Runner: executable that reads specs, executes verification,
  reports pass/fail
- Fixture: static input file referenced by a spec, with integrity
  hashing
- Framework: the combination (schema + specs + runner + CI wiring)
  for one concern domain

**Frameworks.** 12 frameworks documented. Naming convention: `Ux?TS`
where `x` is the wildcard domain identifier and `?` is the
specific letter. UATS (API tests), UPTS (parser tests), UBTS
(benchmark tests), UNTS (hash verification, "Universal-N" for
"non-cryptographic"), UAITS (AI tests), ULTS (linter tests, etc.).

**Why it matters to MDEMG.** UxTS governs test contracts across the
codebase. Frameworks generate spec files that runners verify; CI
enforces gates. This is how MDEMG keeps its 2,381 tests organized
across heterogeneous domains.

**Why it matters to our architectural work.** UxTS is essentially a
**specification language** for the entity's constraints. The
hybrid architecture's superstructure may benefit from a similar
formalism — declarative constraint specifications validated by
runners, with the superstructure responsible for evaluating
constraint satisfaction. This is a meta-architectural connection,
not a direct adoption.

**Where in the codebase.** `docs/uxts01/` (spec package),
`docs/specs/UXTS_PORTABLE_AGENT_SPEC.md` (the canonical spec),
`docs/guides/UXTS_DEVELOPER_GUIDE.md` (developer guide),
`docs/research/UXTS_FRAMEWORK_DEEP_DIVE_LOG.md`. Implementation
distributed across multiple packages depending on which framework.

**Deep-trace candidate.** Possibly. Lower priority than Jiminy/J17/RSIC
for our architectural work because UxTS is governance-of-tests,
not core cognitive substrate. Worth understanding, may be
deferred.

### 7. The LLM Call Layer (and the Modular Intelligence Plugin Architecture)

**What it is.** Two related concerns combined for orientation:
(a) the way MDEMG actually invokes LLMs internally, and
(b) the broader plugin architecture that allows LLM-based and
non-LLM modules to extend MDEMG's capabilities.

**LLM call layer (a).**
- 16 LLM call sites total, all single-shot structured-output or
  reasoning
- No tool-calling — explicitly banned across the stack
- 9 banned patterns grep-audited each sprint (`tool_use`,
  `tool_call`, `function_call`, `--tool-call-parser`,
  `enable-auto-tool-choice`, `preserve_thinking`, etc.)
- Fine-tuning target: Qwen3.6-35B-A3B MoE with two-tier MoE-Sieve
  LoRA (current sprint FT-LORA from 2026-04-21)
- Training data: TimescaleDB recording of LLM interactions during
  normal operation; LoRA SFT and RLHF DPO on the recorded data
- Pipeline scripts in `neural/training/`

**Three training workstreams** (from `neural-training-pipeline.md`,
substantially more developed than the AGENT_HANDOFF excerpt
suggested):

| Workstream | Technique | Model | Data Source | Status |
|---|---|---|---|---|
| **Cross-encoder reranker (NR-4)** | MSE regression | ms-marco-MiniLM-L-6-v2 | `rerank_collector` JSONL | Built |
| **Generative LoRA (Phases 2-12)** | SFT + GRPO | Qwen3.6-35B-A3B | `llm_interactions` (16 tasks) | Pipeline complete |
| **Embedding fine-tuning (Phase D)** | Contrastive | 3072-dim domain-tuned | `embedding_events` + `retrieval_events` | Collecting data |

The cross-encoder reranker has a 100× speedup story: ~5ms CPU vs
~500ms LLM API call. For typical recall with 20 candidates, that's
10s of LLM re-ranking replaced by 100ms of local inference. The
small local model progressively learns to approximate the LLM's
judgment via MSE regression on collected `(query, candidate, score)`
tuples. Spearman rank correlation tracks ranking-quality preservation.

**Generative LoRA pipeline** (9 explicit steps, all complete):

1. **Collect** — `mdemg data export-auto` (LaunchAgent, daily TSDB export with retention)
2. **Filter** — `quality_filter.py` with 8 quality gates: privacy, empty, error, duplicate, latency, model, prompt hash, ULTS schema
3. **Convert** — `format_converter.py` (JSONL → HuggingFace MLX chat with RAFT context, think-mode wrapping)
4. **Version** — `dataset_versioner.py` (temporal train/test/val splits, dedup, exogenous ratio checks, manifest)
5. **Train** — `train_ft.py` (LoRA via mlx-lm-lora, manifest validation, anti-collapse gate)
6. **Evaluate** — `evaluate_ft.py` (per-task vs held-out test, ULTS quality_metrics contract)
7. **Gate** — `regression_gate.py` (deployment decision: no task regresses >5%, ≥2 improve, ≥95% JSON validity)
8. **Deploy** — `quantize_deploy.py` (fuse adapter, quantize 4-bit, verify)
9. **Serve** — `vllm-mlx` (OpenAI-compatible inference with prefix caching)

Plus supplementary tools: `teacher_distill.py` (synthetic data for
under-represented tasks), `reward_functions.py` (21 GRPO reward
functions), `test_vllm_mlx.py` (smoke test all 16 ULTS tasks).

**Multi-paradigm curation** (UAITS framework, added 2026-04-10):
`paradigm_router.py` for spec-driven curation across SFT/DPO/RAFT/curriculum
paradigms. `dpo_builder.py` constructs DPO preference pairs from
`constraint_outcomes` + `llm_interactions`. CLI surfaces:
`mdemg data curate` and `mdemg data validate`.

**Modular Intelligence Plugin Architecture (b).** Plugin types:
- **INGESTION**: parse external sources into observations
  (`linear-module`, `obsidian-module`)
- **REASONING**: process or re-rank retrieval results
  (`keyword-booster`, `llm-reranker`)
- **APE** (Active Participant Engine): background tasks and
  autonomous actions (`reflection-module`, `consistency-checker`)

**Communication.** gRPC over Unix sockets (low latency).

**Plugin Manager.** `internal/plugins/` — discovery (scans `/plugins`
directory), lifecycle (spawn, handshake, health check, shutdown),
routing (match requests to capable modules).

**Retrieval pipeline integration points** for REASONING modules:
1. Vector recall → candidates
2. BM25 + RRF fusion → hybrid search (if enabled)
3. Graph expansion → 1-D hop traversal
4. Spreading activation → transient scores
5. Initial scoring → ScoreAndRank
6. **REASONING MODULES** → plugin-based re-ranking
7. Built-in LLM rerank → optional LLM scoring
8. Jiminy explanations → explainable retrieval + Jiminy Guide

**APE event triggers.** `session_end`, `consolidate`, `ingest`,
`schedule` (cron-based).

**Deep-trace candidate.** The plugin architecture itself, yes — it's
the mechanism by which MDEMG is extensible. The LLM call sites
specifically, lower priority unless the trace surfaces something
unexpected. The Qwen fine-tuning pipeline is operationally
significant but not architecturally load-bearing for our
hybrid-architecture work yet.

## Subsystem relationship graph

The subsystems are densely interconnected. Key relationships:

- **MEMG (graph) is the substrate** that all other subsystems read
  from and write to.
- **CMS produces the most observations** (via the conversation
  capture path) and uses the graph as its persistence layer.
- **Jiminy reads from the graph** (5 parallel sources) and produces
  guidance. Jiminy is consumed by external agents; J17 is how that
  guidance is encoded.
- **J17 carries Jiminy's output** in compact form; J17's session
  tickets persist Jiminy's state across context resets.
- **RSIC monitors and improves** all of the above. RSIC's 7
  dimensions cover retrieval, memory, edge, task, guidance,
  protocol (J17), and synergy.
- **UxTS governs tests** that verify all of the above. Including
  J17 conformance tests.
- **The plugin architecture extends** any of the above. Reasoning
  modules plug into retrieval; APE modules run background tasks
  including RSIC orchestration.
- **The LLM call layer is invoked** by Jiminy synthesis (J8), CMS
  observation classification, RSIC reflection (LLM-driven pattern
  detection), and several other points.

The seven subsystems are not orthogonal layers; they are entangled
capabilities that share the graph as substrate and coordinate through
the plugin and APE machinery.

## Things this orientation does NOT include

This is a structural skeleton, not mechanism-level understanding. It
deliberately omits:

- The actual code paths and call graphs for any subsystem
- The exact data structures and message formats
- The specific learning rules and their parameter values
- How subsystems interact at runtime — only their high-level
  relationships
- Concrete failure modes and edge cases
- The 14-component scoring function in retrieval (mentioned in
  passing in Jiminy docs but not investigated)
- Phase-by-phase development history beyond what VISION.md states
- The TimescaleDB metric schema and the LLM training pipeline
  details

These belong in deep-trace cycles, one subsystem at a time.

## Confidence calibration

| Claim | Rating | Why |
|---|---|---|
| The seven subsystems and their general roles | 5 | Direct restatement from VISION.md and the docs index. |
| The graph's layer structure and Hebbian extensions | 5 | Direct from VISION.md and 01_Architecture.md. |
| CMS observation types, lifecycle, scoring | 5 | Direct from CMS.md. |
| RSIC's 5-stage cycle and 3 time scales | 5 | Direct from VISION.md. |
| Jiminy's 5-source fan-out architecture | 5 | Direct from jiminy-inner-voice.md. |
| J17's three-tier encoding and session tickets | 5 | Direct from j17-ai2ai-protocol.md (sections 1-3.1 read in full). |
| J17's 5-tier framing (T0-T4 mentioned in summary) | 3 | Documentation references both 3-tier and 5-tier; question filed for resolution. |
| UxTS's 4-layer architecture and 12-framework count | 4 | UxTS portable agent spec read at head only; the 12-framework count is from the developer guide reference. |
| LLM call layer's no-tool-calling policy and 16 call sites | 5 | Direct from VISION.md operational architecture section. |
| The plugin architecture's three types and gRPC-over-Unix-socket transport | 5 | Direct from 01_Architecture.md. |
| Tech stack specifics (Neo4j 5.x, Go 1.24, embedding model + 3072d) | 5 | Direct from AGENT_HANDOFF.md §2. |
| `internal/` 22-package layout | 5 | Direct from AGENT_HANDOFF.md §2. |
| Retrieval pipeline scoring formula (specific weights) | 5 | Direct from AGENT_HANDOFF.md §2. |
| 105 core phases / 16 sidecar phases / 5 cognitive gap phases all complete | 5 | Direct from AGENT_HANDOFF.md §1 and §4. |
| Fine-tuning state (FT-OAI-001/002 complete, FT-OAI-003 planned, with specific deltas) | 5 | Direct from AGENT_HANDOFF.md §5 (current as of 2026-04-21). |
| Architecture-maps-as-self-documentation pattern | 4 | Direct from AGENT_HANDOFF.md §5; the architectural significance for our work is my inference. |
| J17's 4-cycle self-improvement journey (8.3 → 9.9, 1.9× → 5.2×) | 5 | Direct from j17-ai2ai-protocol.md §5 (sections 5.1-5.7). |
| ProtocolHealth 4-component formula with weights | 5 | Direct from j17-ai2ai-protocol.md §8.1. |
| J17 5 reflection patterns + 4 mutation actions | 5 | Direct from j17-ai2ai-protocol.md §8.1. |
| J17-PC's 25-35% aggregate token reduction across 5 callers | 5 | Direct from j17-prompt-compression.md (full read). |
| RSIC closed-loop (AR-1) post-cycle re-assessment with auto-rollback | 5 | Direct from rsic-feedback-loop.md (full read). |
| DH-005 health-weighting formula and 7-dimension defaults | 5 | Direct from rsic-feedback-loop.md §"Health Weighting & Confidence". |
| Three training workstreams (cross-encoder/generative/embedding) | 5 | Direct from neural-training-pipeline.md §"Three Training Workstreams". |
| Generative LoRA 9-step pipeline with quality gates | 5 | Direct from neural-training-pipeline.md. |
| Cross-encoder 100× speedup (5ms vs 500ms) | 5 | Direct from neural-training-pipeline.md §"Why This Matters". |
| Benchmark methodology and 0.898 / +5.2% / 100% results | 5 | Direct from README.md §"Benchmark Receipts". |
| Cross-subsystem relationships in the relationship graph | 4 | Inferred from the documents' cross-references; a direct verification would require reading code paths. |

## Surfaced questions (filed separately as Q-artifacts in subsequent PR)

1. **What is the relationship between J17's 3-tier and 5-tier framings?**
   The summary references "5 encoding tiers (T0-T4)" but the body
   describes only T1-T3. T0 and T4 are referenced but not defined
   in the read sections.

2. **What does the 14-component scoring function in Jiminy J7 actually
   compute?** Mentioned in passing as "14-component scoring." Bears
   on understanding what Jiminy considers when ranking guidance
   candidates.

3. **How does RSIC's "Synergy" dimension (7th of 7) actually work?**
   Synergy is described as Claude Code ↔ MDEMG token reduction
   with overflow interceptor, but the metric and the action set
   aren't clear from the orientation reads.

4. **What is the difference between "MEMG" and "MDEMG"?** The graph
   itself is called the MEMG (Multi-Layer Emergent Memory Graph,
   the substrate); MDEMG is the system as a whole. The naming
   distinction is consistent in some places and conflated in
   others. Worth verifying the canonical usage.

5. **How do session tickets resolve trust score conflicts when an
   agent presents a stale ticket with high trust?** Renewal
   semantics described; conflict semantics not.

These will be filed as Q- artifacts in a follow-up PR alongside
this one or in the next cycle, depending on Cycle 1 unblock
decision.

## Next-cycle deep-trace candidates, ranked

In rough order of architectural-work load-bearing:

1. **J17 deep trace** — fills D-010's AI2AI protocol commitment.
   Priority 1 because the protocol design for our hybrid
   architecture cannot proceed without this evidence.
2. **Jiminy + RSIC combined trace** — D-010's teacher-pupil pattern
   in operational form, plus the recursive-self-improvement
   precedent for Q-008. Treat as one cycle because they're tightly
   coupled.
3. **The graph (MEMG) deep trace** — the substrate; bears on
   Q-006 (dimensional commitment) and Q-007 (projection-and-
   anchoring) more than initially expected because the graph is
   already topology-first, metric-optional.
4. **CMS deep trace** — important for the entity's persistent
   memory; less urgent than the above three because the
   architectural commitment in D-010 doesn't yet specify CMS-shape.
5. **The plugin architecture trace** — important for substrate
   flexibility (Q-009) but lower priority because we're not yet
   designing extensions.
6. **UxTS trace** — useful for understanding the testing/governance
   surface, lower priority than core cognitive subsystems.
7. **LLM call layer + Qwen training pipeline trace** — operational
   pattern is interesting (especially the no-tool-calling policy
   and the LoRA strategy), but lower architectural priority for
   the program's hybrid-architecture work.

This ranking is the recommended Cycle 2 / Cycle 3 / Cycle 4 sequencing.

## Implications for our architectural work

**MDEMG is more than R&D substrate.** Earlier vault artifacts
described MDEMG as "an R&D substrate" without further qualification.
This orientation reveals MDEMG to be a **substantially developed
operational system** (~227K Go LOC, 7 phases complete, deployed via
Docker Compose, with an active fine-tuning pipeline). The framing in
prior decisions (D-010's "MDEMG is the seed of the eventual
architecture") is sharpened: the seed is *much closer* to the
eventual architecture than the R&D-substrate framing suggested.

**The teacher-pupil pattern in D-010 has direct precedent.** Jiminy's
"conscience" framing is exactly the teacher-pupil dynamic. Jiminy
runs before every prompt; intercepts via hooks; injects guidance;
tracks effectiveness through feedback; calibrates via RSIC. The
pattern is operational, not aspirational.

**The AI2AI protocol commitment in D-010 has concrete reference.**
J17 is exactly the kind of protocol D-010 named as needed. The
specific commitments (reference, intent, contradiction, confidence,
attention focus, world-model state) map to J17's encoding fields
(type, severity, annotations, escalation, source). Some commitments
go further than J17's current scope and may need protocol
extension; this is design work, not invention from scratch.

**The recursive-self-improvement commitment in Q-008 has direct
precedent.** RSIC is exactly that — a 5-stage cycle operating at
three time scales, with safety mechanisms (dry-run, rollback,
calibration). The architectural pattern is proven; what we need to
add for the eventual entity is dimensions RSIC doesn't currently
cover (predictive horizon, world model coherence, sensorimotor
ground-truth checking).

**Substrate flexibility (Q-009) is partially served.** The plugin
architecture is the substrate-flexibility commitment in operational
form: new modules can extend the system without modifying the core.
Whether this generalizes to the substrate-flexibility we want for
sensors and effectors (in the eventual embodied system) is open,
but the precedent is encouraging.

**The architecture inversion in D-010 is correct but the LLM is
already in the right place.** D-010 inverted "LLM with augmentations"
to "superstructure with LLM as faculty." MDEMG already operates with
the LLM as a called-out faculty (16 invocation sites, no tool
calling, single-shot calls, structured outputs). The Qwen fine-tuning
pipeline is the LoRA-managed continuous-learning surface D-010
committed to. The inversion is not aspirational; it's operational.

**The LoRA-managed continuous learning surface (D-010 commitment)
is more disciplined than D-010 implied.** AGENT_HANDOFF.md §5 reveals
the FT-OAI-001/002/003 sprint structure: explicit quality gates
(MARGINAL verdict on 4.1-mini fine-tune), cost caps (≤$250 for
FT-OAI-003), quality floors (≥0.8322 cross-base), strategic
gap-closure framing (the FT(cheap-base) ≈ prod(prod-base) north
star). Future architectural specification of the LoRA adapter
management interface should adopt this gating discipline as a
reference pattern. The eventual entity's continuous learning will
need similar quality/cost gates.

**MDEMG documents itself for its own use.** The
`docs/architecture/maps/` pattern — 10 compact architecture maps
auto-generated by `scripts/generate_arch_maps.py` and consumed by
Jiminy as context injection — is a precedent for the eventual
entity having structural self-knowledge that's machine-readable
and machine-maintained. UITS optimization protection prevents the
generator from overwriting manually-converged maps. This connects
to T004's homeostatic-boundary framing (Q-005): the entity needs
to know what it is in order to maintain itself; auto-generated
self-description is one mechanism for that.

**Data-sufficiency-aware aggregation is load-bearing for the
entity's homeostatic boundary** (Q-005). RSIC's DH-005 formula
`overall_health = Σ(w_i · c_i · s_i) / Σ(w_i · c_i)` excludes
dimensions with zero confidence rather than penalizing them. The
formula is honest about what it doesn't know. This pattern is
directly applicable to the entity's self-model: each homeostatic
dimension contributes proportional to weight × confidence × score,
missing dimensions are excluded not penalized. As the entity matures
and more aspects of self become measurable, the formula gracefully
extends. **This is a reference pattern for D-010's superstructure
state representation.**

**The J17 self-improvement journey is the entity's recursive
predictive horizon expansion in operational form** (Q-008). The
documented finding — "when individual codes fail, the protocol
itself must evolve; better codes within a weak protocol hit a
ceiling, but a stronger protocol makes even simple codes work" —
is exactly the recursive expansion principle: when local
optimization plateaus, optimization moves up a level. T005's
cycle methodology and J17's RSIC integration are conceptually
parallel. The entity's predictive horizon expansion may reuse the
same iteration structure: empirical loop → plateau → meta-loop
intervention → new plateau at higher capacity.

**The RSIC closed-loop pattern (AR-1) is reference architecture
for the entity's self-corrective cycles.** Post-action re-assess,
evaluate against pre-declared success criteria, auto-rollback for
reversible actions, log irreversible failures as
`completed_no_benefit`. The pattern is more disciplined than naive
"try and see" — every action carries its own success criterion;
calibration confidence only increases when both completed AND
successful. The eventual entity's continuous learning will benefit
from this gating discipline.

**The training pipeline is more disciplined than D-010 implied.**
Beyond AGENT_HANDOFF.md's FT-OAI quality gates, the generative LoRA
pipeline has 8 quality gates at filter time, deployment gates (no
task regresses >5%, ≥2 tasks improve, ≥95% JSON validity), and
multi-paradigm curation across SFT/DPO/RAFT/curriculum. Reference
architecture for the eventual entity's continuous learning specification.

## What this artifact does for Cycle 1

The orientation map satisfies the cycle's research-mode gate
criterion: it produces a structural map of MDEMG's subsystems
sufficient to scope subsequent cycles. It does not produce
mechanism-level understanding of any single subsystem; that's
correctly the work of Cycles 2..N.

The cycle should now move to **assess** mode: was this orientation the
right thing to do, did it produce the kind of evidence that unblocks
next-cycle planning, and what cycles does this unblock?

That assessment is the next conversational turn with @reh3376, not
something to land in this artifact.
