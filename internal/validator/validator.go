// Package validator walks the vault and checks artifact frontmatter for required fields and valid values.
package validator

import (
	"fmt"
	"slices"

	"github.com/reh3376/ann-research/internal/artifact"
)

// Issue describes one validation problem found in an artifact.
type Issue struct {
	Path    string
	Message string
}

func (i Issue) String() string {
	return fmt.Sprintf("%s: %s", i.Path, i.Message)
}

// ValidateFile checks a single artifact for schema correctness.
// Returns nil if valid, or a list of issues otherwise.
func ValidateFile(f *artifact.File) []Issue {
	var issues []Issue
	fm := f.Frontmatter

	// Required for all types
	if fm.ID == "" {
		issues = append(issues, Issue{f.Path, "missing required field: id"})
	}
	if fm.Type == "" {
		issues = append(issues, Issue{f.Path, "missing required field: type"})
	}
	if fm.Slug == "" {
		issues = append(issues, Issue{f.Path, "missing required field: slug"})
	}
	if fm.Created.IsZero() {
		issues = append(issues, Issue{f.Path, "missing required field: created"})
	}
	if fm.Updated.IsZero() {
		issues = append(issues, Issue{f.Path, "missing required field: updated"})
	}
	if fm.Status == "" {
		issues = append(issues, Issue{f.Path, "missing required field: status"})
	}
	if fm.Confidence < 1 || fm.Confidence > 5 {
		issues = append(issues, Issue{f.Path, fmt.Sprintf("confidence must be 1-5, got %d", fm.Confidence)})
	}

	// Status must be valid for the type
	if fm.Type != "" && fm.Status != "" {
		valid := artifact.ValidStatusesByType(fm.Type)
		if valid != nil && !slices.Contains(valid, fm.Status) {
			issues = append(issues, Issue{
				f.Path,
				fmt.Sprintf("status %q invalid for type %s (valid: %v)", fm.Status, fm.Type, valid),
			})
		}
	}

	// Type-specific required fields
	switch fm.Type {
	case artifact.TypeHypothesis:
		if fm.FalsificationCriteria == "" {
			issues = append(issues, Issue{
				f.Path,
				"hypothesis requires falsification_criteria (what would falsify this?)",
			})
		}
	case artifact.TypeDeadEnd:
		if fm.FailureMode == "" {
			issues = append(issues, Issue{f.Path, "dead-end requires failure_mode"})
		}
		if fm.Status == "blocked-on" && fm.RetryIf == "" {
			issues = append(issues, Issue{
				f.Path,
				"dead-end with status=blocked-on requires retry_if",
			})
		}
	case artifact.TypeThread:
		if fm.CurrentQuestion == "" {
			issues = append(issues, Issue{f.Path, "thread requires current_question"})
		}
	}

	return issues
}

// ValidateVault walks all artifacts in vaultRoot and returns the cumulative issues.
// Parsing failures are returned as issues (one per file).
func ValidateVault(vaultRoot string) ([]Issue, error) {
	var allIssues []Issue

	for _, t := range artifact.AllTypes {
		paths, err := artifact.ListArtifactPaths(vaultRoot, t)
		if err != nil {
			return nil, fmt.Errorf("list %s: %w", t, err)
		}
		for _, p := range paths {
			f, err := artifact.ParseFile(p)
			if err != nil {
				allIssues = append(allIssues, Issue{p, "parse error: " + err.Error()})
				continue
			}
			allIssues = append(allIssues, ValidateFile(f)...)
		}
	}

	return allIssues, nil
}
