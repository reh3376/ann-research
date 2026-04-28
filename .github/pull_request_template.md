<!--
PR template for ann-research.
Brief is fine. The point is to give the maintainer enough context to approve
or request changes without having to dig through the diff.
-->

## What changed

<!-- One-line summary, then a list of artifacts added or modified. -->

## Threads affected

<!-- Which T### threads does this PR touch? -->

## What needs maintainer judgment

<!-- Specific points where you want @reh3376's eyes. Examples:
     - "Confidence rating on H-2026-04-28-001 — I have it at 2; you may know
        domain-specific evidence that should bump it up."
     - "Whether DE-... should be terminal or blocked-on."
     - "None — straightforward observation, just recording for the trail." -->

## Assistant's recommendation

<!-- One of: approve and merge / approve with edits requested below /
     this is provisional and may need rework. -->

## Checklist

- [ ] All new artifacts pass `vault validate`
- [ ] Frontmatter `confidence` reflects actual epistemic state, not aspiration
- [ ] `provenance` cites specific files / line numbers / URLs / conversation refs
- [ ] If anything supersedes prior work, `supersedes` / `superseded_by` are linked
- [ ] No secrets, internal paths, or non-public material in the diff
