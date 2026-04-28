package artifact

import (
	"bytes"
	_ "embed"
	"fmt"
	"text/template"
	"time"
)

//go:embed templates/observation.md
var tmplObservation string

//go:embed templates/question.md
var tmplQuestion string

//go:embed templates/hypothesis.md
var tmplHypothesis string

//go:embed templates/decision.md
var tmplDecision string

//go:embed templates/dead-end.md
var tmplDeadEnd string

//go:embed templates/thread.md
var tmplThread string

// templateFor returns the embedded template string for a given type.
func templateFor(t Type) (string, error) {
	switch t {
	case TypeObservation:
		return tmplObservation, nil
	case TypeQuestion:
		return tmplQuestion, nil
	case TypeHypothesis:
		return tmplHypothesis, nil
	case TypeDecision:
		return tmplDecision, nil
	case TypeDeadEnd:
		return tmplDeadEnd, nil
	case TypeThread:
		return tmplThread, nil
	}
	return "", fmt.Errorf("no template for type %q", t)
}

// TemplateData is the data passed to a template at render time.
type TemplateData struct {
	ID      string
	Type    Type
	Slug    string
	Title   string // human-readable, derived from slug
	Created string // ISO 8601 date
	Updated string // ISO 8601 date
}

// Render applies the template for the given type to produce the file content.
func Render(t Type, data TemplateData) ([]byte, error) {
	src, err := templateFor(t)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New(string(t)).Parse(src)
	if err != nil {
		return nil, fmt.Errorf("parse template for %s: %w", t, err)
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("execute template: %w", err)
	}
	return buf.Bytes(), nil
}

// NewTemplateData builds the data context for a new artifact.
func NewTemplateData(t Type, id, slug string, now time.Time) TemplateData {
	return TemplateData{
		ID:      id,
		Type:    t,
		Slug:    slug,
		Title:   slugToTitle(slug),
		Created: now.UTC().Format("2006-01-02"),
		Updated: now.UTC().Format("2006-01-02"),
	}
}

// slugToTitle converts "monty-vote-orchestration" to "Monty Vote Orchestration".
func slugToTitle(slug string) string {
	words := bytes.Split([]byte(slug), []byte("-"))
	for i, w := range words {
		if len(w) > 0 {
			w[0] = byte(toUpper(w[0]))
			words[i] = w
		}
	}
	return string(bytes.Join(words, []byte(" ")))
}

func toUpper(b byte) byte {
	if b >= 'a' && b <= 'z' {
		return b - 32
	}
	return b
}
