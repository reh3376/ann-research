package artifact

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

// slugRegex matches non-alphanumeric characters that should be replaced with hyphens.
var slugRegex = regexp.MustCompile(`[^a-z0-9]+`)

// Slugify converts free-form text to a filesystem-friendly slug.
//
//	"Monty Vote Orchestration"     -> "monty-vote-orchestration"
//	"   leading/trailing  spaces " -> "leading-trailing-spaces"
//	"emoji 🚀 stripped"             -> "emoji-stripped"
func Slugify(s string) string {
	lower := strings.ToLower(strings.TrimSpace(s))
	hyphenated := slugRegex.ReplaceAllString(lower, "-")
	return strings.Trim(hyphenated, "-")
}

// NextSequenceNumber finds the highest existing sequence number for the given
// type and date in the vault, and returns the next available number.
//
// For thread artifacts, the date is ignored — threads use a global sequence
// since their identity is persistent.
func NextSequenceNumber(vaultRoot string, t Type, date time.Time) (int, error) {
	dir := filepath.Join(vaultRoot, t.Dir())
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, fmt.Errorf("read %s: %w", dir, err)
	}

	var prefix string
	if t == TypeThread {
		// Threads: T{NNN}-{slug}/
		prefix = t.Prefix() + "-"
	} else {
		// Other artifacts: {PREFIX}-{YYYY-MM-DD}-{NNN}-{slug}.md
		prefix = fmt.Sprintf("%s-%s-", t.Prefix(), date.Format("2006-01-02"))
	}

	maxSeq := 0
	for _, entry := range entries {
		name := entry.Name()
		if !strings.HasPrefix(name, prefix) {
			continue
		}
		rest := strings.TrimPrefix(name, prefix)
		// rest looks like "001-some-slug.md" or "001-some-slug" (for threads which are dirs)
		parts := strings.SplitN(rest, "-", 2)
		if len(parts) == 0 {
			continue
		}
		seq, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}
		if seq > maxSeq {
			maxSeq = seq
		}
	}

	return maxSeq + 1, nil
}

// BuildID constructs the artifact ID and filename for a new artifact.
//
// For non-thread types: returns ("OBS-2026-04-28-003", "OBS-2026-04-28-003-monty-vote-orchestration.md").
// For threads:          returns ("T004",                "T004-evidence-deep-dive").
//
// Note: thread "filename" is actually the folder name; the file inside is always thread.md.
func BuildID(t Type, date time.Time, seq int, slug string) (id, filename string) {
	if t == TypeThread {
		id = fmt.Sprintf("%s%03d", t.Prefix(), seq)
		filename = fmt.Sprintf("%s-%s", id, slug)
		return id, filename
	}
	id = fmt.Sprintf("%s-%s-%03d", t.Prefix(), date.Format("2006-01-02"), seq)
	filename = fmt.Sprintf("%s-%s.md", id, slug)
	return id, filename
}

// ListArtifactPaths returns paths to all artifact files of a given type in the vault.
// For threads, returns paths to each thread.md file.
func ListArtifactPaths(vaultRoot string, t Type) ([]string, error) {
	dir := filepath.Join(vaultRoot, t.Dir())
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	var paths []string
	for _, entry := range entries {
		if t == TypeThread {
			if !entry.IsDir() {
				continue
			}
			threadFile := filepath.Join(dir, entry.Name(), "thread.md")
			if _, err := os.Stat(threadFile); err == nil {
				paths = append(paths, threadFile)
			}
			continue
		}
		if entry.IsDir() {
			continue
		}
		name := entry.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}
		if !strings.HasPrefix(name, t.Prefix()+"-") {
			continue
		}
		paths = append(paths, filepath.Join(dir, name))
	}

	sort.Strings(paths)
	return paths, nil
}
