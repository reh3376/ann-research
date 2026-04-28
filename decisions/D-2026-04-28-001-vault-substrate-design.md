---
id: D-2026-04-28-001
type: decision
slug: vault-substrate-design
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 4
where: []
tags: [meta, infrastructure, substrate]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
alternatives_rejected:
  - SQLite-indexed database
  - single daily-log file
  - Notion or other hosted tool
  - per-artifact JSON
---

# Vault Substrate Design

## Decision

Use markdown files with YAML frontmatter, one artifact per file, organized
by type into top-level directories. Six artifact types:
**observation**, **question**, **hypothesis**, **decision**, **dead-end**,
**thread**.

## Context

The original problem statement: "we need to preserve our thoughts,
concepts, ideas, dead ends, failures, and successes in a manner that they
can survive context compaction or full loss of a context window. It is
important to NOT catastrophically forget important context and to be able
to access it when needed."

The substrate has to preserve **research trajectory**, not only conclusions
— what we considered, what we ruled out and why, where we changed our
minds — so that a fresh context (whether a new Claude instance or future-us
with degraded memory) can reconstitute working memory and avoid
relitigating settled questions or repeating known failures.

## Alternatives considered

### SQLite-indexed database

Rejected as premature. Files-and-grep is honest about the scale we're
operating at (low hundreds of artifacts at most for the foreseeable
future), avoids the operational overhead of schema migrations, and
survives any tool change. A SQLite index over the YAML frontmatter is a
reasonable Phase 2 if grep becomes painful, but starting with files is
the right scale.

### Single daily-log file

Rejected. A log-with-sections approach has lower write friction in the
moment but destroys the per-artifact lifecycle that git provides. With
one artifact per file, `git log -- path/to/file.md` shows the trajectory
of a single thought from creation through every revision. With a daily
log, that history is buried in noise from unrelated entries.

### Notion or other hosted tool

Rejected on three grounds: (1) portability — hosted tools die or change;
markdown survives anything that can read text; (2) public visibility —
the repo is public for the discipline that creates, and Notion adds
indirection between the work and the public surface; (3) no native
versioning of the kind git provides.

### Per-artifact JSON

Rejected. Markdown renders in any editor and on GitHub natively, supports
human-readable prose for the body, and accepts YAML frontmatter for the
structured fields. JSON would require either a separate body file or
ugly stringification of multi-paragraph prose.

## Rationale

The substrate has to satisfy four properties:

1. **Survive arbitrary tool changes.** Markdown does. Hosted formats
   don't.
2. **Be greppable.** A future Claude with no specialized infrastructure
   has to be able to find things via simple text search. Plain text in a
   filesystem satisfies this; databases or proprietary formats don't.
3. **Preserve lineage.** Belief changes over time. The substrate has to
   record what was believed, when, and what replaced it. YAML
   `supersedes` / `superseded_by` fields plus per-file git history
   satisfy this.
4. **Force the discipline of metadata.** The frontmatter schema requires
   `confidence`, `status`, `provenance`, etc. — fields that are easy to
   omit when writing free prose. Required-field validation in the CLI
   catches the omission before it persists.

Markdown + YAML in git satisfies all four with minimal infrastructure.

## Consequences

- All artifacts must have frontmatter conforming to the per-type schema.
  The CLI validator enforces this.
- The body of each artifact is free-form markdown, but each type has a
  template suggesting structure (what was observed; what would falsify
  this; alternatives rejected; etc.).
- Threads use folders (`threads/T###-slug/`) to allow accumulation of
  thread-internal scratch documents alongside the canonical `thread.md`.
- Other artifacts are flat files in their type directory.
- The `bin/` directory holds the compiled CLI; it is gitignored after
  the first build.

## Revisit conditions

- If the artifact count grows past ~500 and grep becomes too slow to be
  the primary search interface, revisit by adding a SQLite index over
  the frontmatter (additive — would not change the file substrate).
- If the per-type schema starts feeling like procrustean and we keep
  wanting to add ad-hoc fields, revisit by either generalizing the
  schema or adding a free-form `extras` map.
- If the validator's required-field enforcement becomes friction we
  route around (writing artifacts and skipping the CLI), revisit by
  either softening the requirements or making the CLI easier to use.
