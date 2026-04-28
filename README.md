# ann-research

A research vault for novel artificial neural network topologies inspired by
the human neocortex.

This repository is the lab notebook for an ongoing research program whose
working framework is the **Adaptive Hebbian Hierarchy (AHH)** — a topological
extension of the [MDEMG](https://github.com/reh3376/mdemg) cognitive
substrate toward an architecture that learns continually, adjusts reference
frames, and behaves more like a neocortical prediction machine than a
stateless transformer.

The framework name may evolve; the broader research goal — novel ANN
topologies capable of supporting continual, reference-frame-adaptive
intelligence — is what this repository is named for.

---

## What this repository is

This is a **working research vault**, not a finished artifact. It contains:

- Observations made while reading code, papers, and our own prior work
- Open research questions
- Hypotheses with explicit falsification criteria
- Decisions, with the alternatives we considered and rejected
- Dead ends — paths we tried that did not work, preserved so future-us
  doesn't repeat them
- Active research threads that organize the above

It is public so that the discipline of writing artifacts that make sense to a
stranger keeps us honest. It is **not** a curated publication. Anyone reading
at any moment will find half-formed thoughts, hypotheses we may abandon, and
decisions we may reverse. That is the point — the substrate preserves the
trajectory of belief, not only the conclusions.

If you are looking for the polished version of any line of thinking here,
look for a public paper or talk. If there is no public artifact, the work is
still in progress.

---

## Why this exists

A central problem in working with stateful research over long time horizons
is that context — the trail of why we believe what we believe — is fragile.
For human researchers it is fragile because memory is fragile. For LLM
research assistants it is fragile because the context window is finite and
session-bound. This vault is the durable substrate that survives both.

Every artifact carries enough metadata to be understood out of context:

- `where` — which research thread it belongs to
- `created` / `updated` — when it was written and last reviewed
- `status` — its current epistemic standing (open, supported, falsified,
  superseded, etc.)
- `confidence` — how strongly we hold it (1–5)
- `provenance` — the sources behind the claim
- `supersedes` / `superseded_by` — explicit lineage when beliefs change

That is the vault's analog of "where am I / what am I" — fields that let any
fresh reader (human or AI) reconstruct enough context to act usefully.

---

## Repository structure

```
ann-research/
├── README.md            ← you are here
├── INDEX.md             ← current state of understanding (read this second)
├── BRANCH_PROTECTION.md ← workflow and protection rules
├── CONTRIBUTING.md      ← how PRs are handled
├── LICENSE              ← MIT
├── cmd/vault/           ← Go CLI source
├── internal/            ← Go CLI internals
├── sidecar/             ← Python sidecar for scripts
├── bin/                 ← compiled CLI binary (gitignored after first build)
├── observations/        ← OBS-* artifacts
├── questions/           ← Q-* artifacts
├── hypotheses/          ← H-* artifacts
├── decisions/           ← D-* artifacts
├── dead-ends/           ← DE-* artifacts
└── threads/             ← T*/thread.md per active research thread
```

### Artifact ID scheme

```
{TYPE}-{YYYY-MM-DD}-{NNN}-{slug}.md
```

For threads the date is omitted (threads have persistent identity):

```
T{NNN}-{slug}/thread.md
```

Examples:

- `OBS-2026-04-28-003-monty-vote-orchestration.md`
- `Q-2026-04-28-001-symbolic-hypothesis-space.md`
- `T001-monty-trace/thread.md`

---

## Tooling

A small Go CLI (`vault`) creates artifacts from templates, increments
sequence numbers, and validates frontmatter. Python sidecar scripts live in
`sidecar/` for tasks better suited to Python (NLP-style text manipulation,
graph visualization of the artifact dependency graph, future ML work on the
artifacts themselves).

### Build the CLI

```bash
go build -o bin/vault ./cmd/vault
```

Requires Go 1.23 or newer.

### Use the CLI

```bash
# create artifacts
./bin/vault new obs "monty vote orchestration"
./bin/vault new q  "what is the symbolic analog of pose"
./bin/vault new h  "categorical eta multipliers proxy for confidence"
./bin/vault new d  "use cmp message format for inter-column votes"
./bin/vault new de "premature mapping table"
./bin/vault new thread "evidence accumulation deep dive"

# list and validate
./bin/vault list
./bin/vault list questions
./bin/vault validate
```

The validator enforces required frontmatter fields per type — including
"every hypothesis must declare its falsification criteria" and "every
thread must state a current question." This is intentional friction: it
keeps half-thoughts from masquerading as research artifacts.

### Python sidecar

```bash
cd sidecar
uv sync
```

Currently empty by design. Scripts will appear here when we have a need
for them.

---

## Reading order for a fresh visitor

1. **`README.md`** — what this is (you just read it)
2. **`INDEX.md`** — what we currently believe and are working on
3. **`threads/T001-*/thread.md`** etc. — the current research threads
4. **`decisions/D-*-001-*.md`** through the most recent — the decisions
   that constrain everything else
5. **`dead-ends/`** — the failure modes we've already encountered

The `decisions/` and `dead-ends/` folders are the highest-value reading. The
former is what we've committed to. The latter is what stops future-us from
relitigating settled questions or repeating known failures.

---

## Acknowledgments and lineage

The work in this vault builds on:

- **MDEMG** (Multi-Dimensional Emergent Memory Graph) —
  https://github.com/reh3376/mdemg — the operational cognitive substrate
  this research extends.
- **The Thousand Brains Project / Monty** —
  https://github.com/thousandbrainsproject/tbp.monty — Numenta's open-source
  sensorimotor learning system implementing Mountcastle / Hawkins cortical
  column theory.
- The broader research program around **predictive coding** (Rao & Ballard,
  Friston), **HTM** (Hawkins, Ahmad), **Forward-Forward learning** (Hinton,
  Ororbia & Mali), **JEPA** (LeCun), **TEM** (Whittington & Behrens), and
  **active inference** (Friston, Da Costa, Buckley).

Specific citations live in artifact `provenance` fields and in the
forthcoming AHH white paper.

---

## License

MIT. See [LICENSE](LICENSE).
