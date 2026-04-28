---
id: D-2026-04-28-007
type: decision
slug: pr-activity-summary-comment-required
created: 2026-04-28
updated: 2026-04-28
status: committed
confidence: 5
where: []
tags: [meta, infrastructure, workflow, governance, trajectory-preservation]
provenance:
  - "conversation 2026-04-28 between @reh3376 and Claude"
  - "GitHub PR #3 issue comment 4336494166 (first instance)"
  - "D-2026-04-28-001 (vault substrate design — same trajectory-preservation motivation)"
  - "D-2026-04-28-006 (assistant direct push — defines the workflow this convention extends)"
alternatives_rejected:
  - "include activity content in the PR description instead of a separate comment"
  - "attach as a separate file in the PR diff"
  - "rely on conversation transcripts and commit messages alone"
---

# PR Activity Summary Comment Required

## Decision

Every pull request authored by the AI research assistant gets a separate
**activity summary comment** posted to the PR after creation. The
summary captures the trajectory of the research session that produced
the PR — the working-memory record of how the deliverable was built,
distinct from the PR description which states what the deliverable
*is*.

Convention applies prospectively from PR #3 onward (where the first
activity summary was posted at issue comment 4336494166) and to all
future assistant-authored PRs. PRs #1 and #2 are not retroactively
commented; their content is sufficiently meta that the description
alone records the trajectory.

## Context

`D-2026-04-28-001` named the substrate's purpose as preserving research
*trajectory*, not just conclusions — what was considered, what was
ruled out, where reasoning shifted. The vault's artifact schema does
this for *content* decisions: every decision artifact records
`alternatives_rejected`, every dead-end records `failure_mode`, every
hypothesis declares `falsification_criteria`.

A gap remained at the PR boundary. PR descriptions cover *what changed*
and *what needs maintainer judgment*. They do not cover the things
that didn't make it into artifacts: files read in passing, hypotheses
drafted and abandoned, observations considered and rejected as too
speculative, granularity decisions that picked one path over another,
confidence calibration reasoning per claim. Roger flagged this as a
gap during the PR #3 review and asked that it be filled.

The activity summary fills that gap.

## What an activity summary contains

The shape settled on with PR #3, validated by Roger as correct:

1. **Sources consulted, in order.** Files read with line ranges.
   Tool calls in approximate sequence. Total time spent if known.
2. **Granularity decisions.** Where I picked one-artifact-vs-many or
   inline-vs-separate-hypothesis, the reasoning.
3. **Confidence calibration per claim.** A table mapping each
   non-trivial claim in the PR to its confidence rating and the
   evidence basis. Distinguishes "literally read" (5) from "inferred
   with high support" (4) from "plausible reading" (3) etc.
4. **Forking decisions.** When a noticed thing could have been a
   footnote, a question, or ignored, why I chose what I chose.
5. **Observations considered and rejected.** Things noticed during
   reading that I decided not to write up as artifacts, with the
   reason for rejection (typically: not relevant to current question,
   or insufficient evidence base).
6. **Files considered but not read.** Adjacent code I could have
   read but chose to defer, with the reason for deferral.
7. **Hypotheses drafted and abandoned.** Hypotheses I started to
   write but didn't commit to artifacts, with the reason
   (typically: would be unsupported until later trace work
   produces evidence).

The shape is not rigid. Sections may be empty if the session
genuinely had nothing for that bucket. Additional sections may
appear if the session surfaced something the seven categories don't
cover.

## Alternatives considered

### Include activity content in the PR description

Rejected. The description is for the maintainer reviewing now and is
intentionally synopsis-shaped — what changed, what needs judgment,
recommendation. Folding trajectory content in would either bloat the
description or compress the trajectory information into uselessness.
Separation matches the substrate's separation of concerns: per-artifact
content lives in artifacts; per-PR trajectory lives in the comment.

### Attach as a separate file in the PR diff

Rejected. Activity summaries describe a session, not the repository
state. Committing them as files in the diff would make them part of
`main` history, which is wrong scope — the trajectory of *one PR* is
not part of the persistent vault content. GitHub PR comments are the
right durability surface for per-PR ephemera that is more durable
than chat but less durable than artifact.

### Rely on conversation transcripts and commit messages alone

Rejected. Conversation transcripts are not durable across context
loss (the failure mode the entire vault exists to address). Commit
messages are constrained to the per-commit scope and don't naturally
carry session-level reasoning. Neither covers "I considered X and
didn't pursue it" — exactly the trajectory information most likely
to be lost.

## Rationale

Same logic that motivated `D-2026-04-28-001`: trajectory preservation
is the substrate's value proposition. Every place where reasoning
fades without being recorded is a gap. The PR boundary was the
biggest remaining gap; this closes it.

The cost is small — drafting an activity summary at PR-creation time
adds maybe 5 minutes of work to a PR that already took 30+ minutes
to produce. The benefit is large for the specific case of "future
Claude reads the merged PR history to reconstruct what the program
already considered" — without the comments, that reconstruction misses
an entire layer.

## Consequences

- **Every assistant-authored PR from #3 onward gets an activity
  summary comment** posted via the GitHub API after PR creation.
- **The comment is posted, not the description edited.** Preserves
  the original PR body as authored, keeps the trajectory comment
  distinct from the deliverable synopsis.
- **The PR template is updated** to mention the activity summary so
  the convention is visible from the GitHub UI side. Update lands
  with this PR.
- **`BRANCH_PROTECTION.md` is not edited in place** to document this
  — consistent with the precedent set in PR #2 review (`D-006` did
  not edit `D-003` or `BRANCH_PROTECTION.md`; the new artifact is the
  canonical record). Future readers needing the canonical workflow
  read the most recent `D-` artifacts in chronological order.
- **No retroactive comments on PRs #1 and #2.** Both were
  scope-setting PRs whose descriptions adequately captured the
  trajectory; a retroactive comment would mostly restate the
  description.
- **The convention is durable beyond this conversation.** A future
  Claude session reading the vault will encounter this artifact
  before any new PR drafting and apply the convention forward.

## Revisit conditions

- If activity summaries become formulaic enough that they're
  near-duplicates of the PR description, the convention isn't
  earning its cost — revisit by either tightening what counts as
  trajectory content or by folding back into the description.
- If the 7-section shape becomes procrustean (sessions where
  several sections are forced and contentless), revisit by
  loosening to "include the trajectory information that exists;
  omit sections that don't apply."
- If the activity content systematically surfaces things that
  should have been in the artifacts themselves, that's evidence
  the artifact templates are too thin — revisit the templates.
