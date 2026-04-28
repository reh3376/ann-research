---
id: {{.ID}}
type: dead-end
slug: {{.Slug}}
created: {{.Created}}
updated: {{.Updated}}
status: terminal
confidence: 4
where: []
tags: []
provenance: []
failure_mode: ""
retry_if: ""
---

# {{.Title}}

## What was attempted

<!-- Describe the approach in enough detail that future-us could reproduce the attempt. -->

## What went wrong

<!-- The specific failure mode. Be precise — "didn't work" is not useful. -->

## Why it failed (best current understanding)

<!-- Diagnosis. May be partial; flag uncertainty. -->

## Lesson

<!-- The generalized takeaway. The thing we want future-us to remember even if the
     specific attempt is forgotten. -->

## Retry conditions (if any)

<!-- Status `terminal`: do not retry. Status `blocked-on`: retry when the listed condition changes. -->

## Related artifacts

<!-- Links to decisions or hypotheses that led to this attempt. -->
