---
id: D-2026-04-28-002
type: decision
slug: repo-structure-tooling
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 4
where: []
tags: [meta, infrastructure, tooling]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
alternatives_rejected:
  - "Python-only CLI"
  - "private repository"
  - "name: ahh-research (framework-named)"
  - "name: mdemg-ahh-vault (bridge-named)"
---

# Repo Structure and Tooling

## Decision

- **Public GitHub repository** at `github.com/reh3376/ann-research`.
- **MIT licensed.**
- **Go-based CLI** (`vault`) for artifact creation, validation, listing.
- **Python sidecar** (`sidecar/`, Python 3.12+, uv-managed, ruff-linted) for
  scripts where Python's library ecosystem makes the work easier than Go.
- **Repository named `ann-research`**, naming the broader research program
  rather than the current proposed framework.

## Context

Three coupled questions: where does the vault live, what tooling builds it,
what is it called.

## Alternatives considered

### Python-only CLI

Rejected. The MDEMG architecture is Go-base + Python-sidecar; using the
same pattern for this research vault has two benefits: (1) stack
consistency means skills transfer between repos, (2) we eat our own
dogfood on the architecture pattern we're advocating for in the broader
research program. The argument for Python-only was speed-of-iteration;
that's still recoverable through the sidecar without forcing the whole
CLI into Python.

### Private repository

Rejected. Public repository forces the discipline of writing artifacts
that make sense to a stranger, which is the same discipline as writing
artifacts that make sense to a future-us with no context. The two
failure modes converge. Nothing in the vault should be sensitive (see
CONTRIBUTING.md and the `.gitignore` for what's excluded).

### Name: `ahh-research` or `mdemg-ahh-vault`

Rejected because both name a *current* framework or bridge. The framework
name (AHH — Adaptive Hebbian Hierarchy) may evolve as the research
matures. The repository should outlast any specific framework name. The
broader goal — research on novel ANN topologies inspired by neocortical
computation — is what the repository is *for*, and `ann-research` names
that.

## Rationale

Go + Python is the operational pattern in MDEMG; using it here has zero
incremental cost and avoids the inconsistency of two parallel substrates
having different toolchains. Public repository is honest about the work
being in-progress and forces clearer writing. Naming for the long-term
goal rather than the current framework hedges against the framework name
changing.

## Consequences

- Go 1.23 or newer required to build the CLI.
- Python 3.12 or newer for the sidecar.
- `uv` + `ruff` are the standard sidecar toolchain; pinned in
  `sidecar/pyproject.toml`.
- The repository is open to public visibility; all new artifacts must
  pass the public-readability bar.
- Future tooling additions (CI, pre-commit hooks, etc.) are deferred
  until concrete need arises, to avoid premature infrastructure.

## Revisit conditions

- If a meaningful collaboration partner needs a private subset of the
  work, consider a separate private repo for that subset rather than
  flipping this one to private.
- If the framework name solidifies into something that's genuinely the
  same thing as the long-term goal, the repository name could be
  revisited; until then, leave it general.
- If the sidecar accumulates substantial code, revisit whether it
  should be a separate repository.
