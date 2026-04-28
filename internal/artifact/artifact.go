// Package artifact defines the research vault artifact types and their schemas.
//
// The vault holds six artifact types:
//
//   - observation: an atomic noticed thing while reading code or docs.
//   - question:    an open, unresolved research question.
//   - hypothesis:  a claim entertained but not yet tested. Requires falsification criteria.
//   - decision:    a directional choice. Records alternatives rejected and why.
//   - dead-end:    a path tried that failed. Records the failure mode for future reference.
//   - thread:      an active research line. Persistent identity; folder-based.
//
// Each artifact is a single markdown file with YAML frontmatter. The frontmatter
// schema is defined per type below. The body is free-form markdown for human prose.
package artifact

import (
	"fmt"
	"time"
)

// Type is the kind of artifact. Determines schema and storage location.
type Type string

const (
	TypeObservation Type = "observation"
	TypeQuestion    Type = "question"
	TypeHypothesis  Type = "hypothesis"
	TypeDecision    Type = "decision"
	TypeDeadEnd     Type = "dead-end"
	TypeThread      Type = "thread"
)

// AllTypes is the canonical list, used for validation and CLI help.
var AllTypes = []Type{
	TypeObservation,
	TypeQuestion,
	TypeHypothesis,
	TypeDecision,
	TypeDeadEnd,
	TypeThread,
}

// Prefix returns the filename prefix for this type (e.g. "OBS", "Q", "T").
func (t Type) Prefix() string {
	switch t {
	case TypeObservation:
		return "OBS"
	case TypeQuestion:
		return "Q"
	case TypeHypothesis:
		return "H"
	case TypeDecision:
		return "D"
	case TypeDeadEnd:
		return "DE"
	case TypeThread:
		return "T"
	}
	return ""
}

// Dir returns the vault subdirectory for this type.
func (t Type) Dir() string {
	switch t {
	case TypeObservation:
		return "observations"
	case TypeQuestion:
		return "questions"
	case TypeHypothesis:
		return "hypotheses"
	case TypeDecision:
		return "decisions"
	case TypeDeadEnd:
		return "dead-ends"
	case TypeThread:
		return "threads"
	}
	return ""
}

// TemplateName returns the template filename for this type.
func (t Type) TemplateName() string {
	return string(t) + ".md"
}

// Aliases maps short forms to canonical types. Used by the CLI.
var Aliases = map[string]Type{
	"obs":         TypeObservation,
	"observation": TypeObservation,
	"q":           TypeQuestion,
	"question":    TypeQuestion,
	"h":           TypeHypothesis,
	"hyp":         TypeHypothesis,
	"hypothesis":  TypeHypothesis,
	"d":           TypeDecision,
	"dec":         TypeDecision,
	"decision":    TypeDecision,
	"de":          TypeDeadEnd,
	"dead-end":    TypeDeadEnd,
	"deadend":     TypeDeadEnd,
	"t":           TypeThread,
	"thr":         TypeThread,
	"thread":      TypeThread,
}

// ParseType normalizes a user-provided type string to a canonical Type.
func ParseType(s string) (Type, error) {
	if t, ok := Aliases[s]; ok {
		return t, nil
	}
	return "", fmt.Errorf("unknown artifact type %q (valid: observation, question, hypothesis, decision, dead-end, thread)", s)
}

// Frontmatter is the union of all YAML frontmatter fields across artifact types.
// Fields are populated based on type; unused fields are omitted from output.
type Frontmatter struct {
	// Identity
	ID   string `yaml:"id"`
	Type Type   `yaml:"type"`
	Slug string `yaml:"slug"`

	// Trajectory
	Created      time.Time `yaml:"created"`
	Updated      time.Time `yaml:"updated"`
	Supersedes   []string  `yaml:"supersedes,omitempty"`
	SupersededBy string    `yaml:"superseded_by,omitempty"`

	// Epistemic state
	Status     string `yaml:"status"`
	Confidence int    `yaml:"confidence"` // 1-5

	// Lineage
	Where      []string `yaml:"where,omitempty"`      // thread IDs (e.g. "T001")
	Tags       []string `yaml:"tags,omitempty"`
	Provenance []string `yaml:"provenance,omitempty"` // file paths, URLs, conversation refs

	// Type-specific fields (only populated where relevant)
	Subject              string   `yaml:"subject,omitempty"`               // observation
	Answer               string   `yaml:"answer,omitempty"`                // question
	FalsificationCriteria string  `yaml:"falsification_criteria,omitempty"` // hypothesis
	AlternativesRejected []string `yaml:"alternatives_rejected,omitempty"` // decision
	FailureMode          string   `yaml:"failure_mode,omitempty"`          // dead-end
	RetryIf              string   `yaml:"retry_if,omitempty"`              // dead-end (blocked-on)
	Owner                string   `yaml:"owner,omitempty"`                 // thread
	CurrentQuestion      string   `yaml:"current_question,omitempty"`      // thread
}

// ValidStatusesByType returns the allowed status values for a given type.
func ValidStatusesByType(t Type) []string {
	switch t {
	case TypeObservation:
		return []string{"recorded", "stale"}
	case TypeQuestion:
		return []string{"open", "investigating", "answered", "abandoned"}
	case TypeHypothesis:
		return []string{"untested", "supported", "falsified", "abandoned"}
	case TypeDecision:
		return []string{"provisional", "committed", "superseded"}
	case TypeDeadEnd:
		return []string{"terminal", "blocked-on"}
	case TypeThread:
		return []string{"active", "paused", "complete", "abandoned"}
	}
	return nil
}
