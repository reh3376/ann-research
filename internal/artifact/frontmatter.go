package artifact

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// frontmatterDelimiter separates the YAML frontmatter from the markdown body.
const frontmatterDelimiter = "---"

// File is the parsed representation of an artifact file.
type File struct {
	Frontmatter Frontmatter
	Body        string // markdown body after frontmatter
	Path        string // absolute path on disk
}

// ParseFile reads and parses an artifact file from disk.
func ParseFile(path string) (*File, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	fm, body, err := splitFrontmatter(raw)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	var parsed Frontmatter
	if err := yaml.Unmarshal(fm, &parsed); err != nil {
		return nil, fmt.Errorf("yaml in %s: %w", path, err)
	}

	return &File{
		Frontmatter: parsed,
		Body:        string(body),
		Path:        path,
	}, nil
}

// splitFrontmatter separates `---\nYAML\n---\nbody` into the YAML bytes and body bytes.
// Returns an error if the file does not start with frontmatter.
func splitFrontmatter(raw []byte) (frontmatter, body []byte, err error) {
	// Normalize line endings — git on Windows occasionally produces CRLF.
	normalized := bytes.ReplaceAll(raw, []byte("\r\n"), []byte("\n"))

	if !bytes.HasPrefix(normalized, []byte(frontmatterDelimiter+"\n")) {
		return nil, nil, fmt.Errorf("file does not start with %q frontmatter delimiter", frontmatterDelimiter)
	}

	rest := normalized[len(frontmatterDelimiter)+1:] // skip "---\n"
	closeIdx := bytes.Index(rest, []byte("\n"+frontmatterDelimiter+"\n"))
	if closeIdx < 0 {
		// Allow trailing delimiter without newline (file ends with `---`).
		if bytes.HasSuffix(rest, []byte("\n"+frontmatterDelimiter)) {
			closeIdx = len(rest) - len(frontmatterDelimiter) - 1
		} else {
			return nil, nil, fmt.Errorf("frontmatter never closes (missing trailing %q)", frontmatterDelimiter)
		}
	}

	frontmatter = rest[:closeIdx]
	bodyStart := closeIdx + len("\n"+frontmatterDelimiter+"\n")
	if bodyStart > len(rest) {
		bodyStart = len(rest)
	}
	body = rest[bodyStart:]
	return frontmatter, body, nil
}

// Render serializes a File back to disk format (frontmatter + body).
func (f *File) Render() ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteString(frontmatterDelimiter + "\n")

	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(f.Frontmatter); err != nil {
		return nil, fmt.Errorf("encode frontmatter: %w", err)
	}
	if err := enc.Close(); err != nil {
		return nil, fmt.Errorf("close yaml encoder: %w", err)
	}

	buf.WriteString(frontmatterDelimiter + "\n")

	body := f.Body
	if !strings.HasPrefix(body, "\n") {
		buf.WriteString("\n")
	}
	buf.WriteString(body)
	if !strings.HasSuffix(body, "\n") {
		buf.WriteString("\n")
	}

	return buf.Bytes(), nil
}

// Write saves the artifact to disk at f.Path.
func (f *File) Write() error {
	data, err := f.Render()
	if err != nil {
		return err
	}
	return os.WriteFile(f.Path, data, 0o644)
}
