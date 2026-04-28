# Python sidecar

Scripts and utilities for the `ann-research` vault that are easier to write
in Python than in Go.

This directory is intentionally near-empty at scaffold time. Scripts will be
added here when we have a need for them — likely candidates:

- Generating a graph visualization of artifact dependencies
  (observation → hypothesis → decision lineage)
- Exporting subsets of the vault to other formats (Obsidian, HTML, PDF)
- NLP-style analysis of artifact bodies (clustering, topic modeling)
- Bridging code to interact with `tbp.monty` or other Python-native research
  systems

## Setup

Requires Python 3.12 or newer and [uv](https://docs.astral.sh/uv/).

```bash
cd sidecar
uv sync
```

To run scripts:

```bash
uv run python -m ann_research.<script_name>
```

## Why this exists separately from the Go CLI

The Go CLI (`vault`) handles the latency-sensitive, structurally simple
operations: artifact creation, frontmatter validation, listing. The Python
sidecar is for tasks where Python's library ecosystem (numpy, networkx,
matplotlib, pyyaml, the Monty research code) makes the work substantially
easier than writing it in Go.

The boundary is pragmatic, not principled. If a Python script becomes
load-bearing or performance-critical, it can be rewritten in Go. If a Go
operation needs scientific Python libraries, it can be exposed via a
sidecar wrapper.
