// vault is the CLI for the ann-research vault.
//
// Subcommands:
//
//	vault new <type> <slug>    Create a new artifact from template.
//	vault validate             Check all artifacts for schema correctness.
//	vault list <type>          List artifacts of a type.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/reh3376/ann-research/internal/artifact"
	"github.com/reh3376/ann-research/internal/validator"
	"github.com/spf13/cobra"
)

var (
	vaultRoot string
	version   = "0.1.0-dev"
)

func main() {
	root := &cobra.Command{
		Use:           "vault",
		Short:         "ann-research vault CLI",
		Long:          "Manage research artifacts (observations, questions, hypotheses, decisions, dead-ends, threads).",
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cwd, _ := os.Getwd()
	root.PersistentFlags().StringVar(&vaultRoot, "vault", cwd, "vault root directory (default: current directory)")

	root.AddCommand(newCmd())
	root.AddCommand(validateCmd())
	root.AddCommand(listCmd())

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func newCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "new <type> <slug>",
		Short: "Create a new artifact from template",
		Long: `Create a new artifact. Type may be: observation (obs), question (q),
hypothesis (h), decision (d), dead-end (de), thread (t).

Examples:
  vault new obs "monty vote orchestration"
  vault new q "what is the symbolic analog of pose"
  vault new thread "evidence accumulation deep dive"`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			t, err := artifact.ParseType(args[0])
			if err != nil {
				return err
			}
			slug := artifact.Slugify(args[1])
			if slug == "" {
				return fmt.Errorf("slug is empty after normalization")
			}

			now := time.Now().UTC()
			seq, err := artifact.NextSequenceNumber(vaultRoot, t, now)
			if err != nil {
				return fmt.Errorf("compute next sequence: %w", err)
			}

			id, filename := artifact.BuildID(t, now, seq, slug)
			data := artifact.NewTemplateData(t, id, slug, now)
			content, err := artifact.Render(t, data)
			if err != nil {
				return fmt.Errorf("render template: %w", err)
			}

			var path string
			if t == artifact.TypeThread {
				dir := filepath.Join(vaultRoot, t.Dir(), filename)
				if err := os.MkdirAll(dir, 0o755); err != nil {
					return fmt.Errorf("create thread directory: %w", err)
				}
				path = filepath.Join(dir, "thread.md")
			} else {
				dir := filepath.Join(vaultRoot, t.Dir())
				if err := os.MkdirAll(dir, 0o755); err != nil {
					return fmt.Errorf("create directory: %w", err)
				}
				path = filepath.Join(dir, filename)
			}

			if _, err := os.Stat(path); err == nil {
				return fmt.Errorf("file already exists: %s", path)
			}

			if err := os.WriteFile(path, content, 0o644); err != nil {
				return fmt.Errorf("write file: %w", err)
			}

			fmt.Println(path)
			return nil
		},
	}
}

func validateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "validate",
		Short: "Validate all artifacts in the vault",
		RunE: func(cmd *cobra.Command, args []string) error {
			issues, err := validator.ValidateVault(vaultRoot)
			if err != nil {
				return err
			}
			if len(issues) == 0 {
				fmt.Println("OK: all artifacts valid")
				return nil
			}
			for _, issue := range issues {
				fmt.Println(issue)
			}
			return fmt.Errorf("%d validation issue(s) found", len(issues))
		},
	}
}

func listCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [type]",
		Short: "List artifacts (optionally filtered by type)",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			var types []artifact.Type
			if len(args) == 1 {
				t, err := artifact.ParseType(args[0])
				if err != nil {
					return err
				}
				types = []artifact.Type{t}
			} else {
				types = artifact.AllTypes
			}

			for _, t := range types {
				paths, err := artifact.ListArtifactPaths(vaultRoot, t)
				if err != nil {
					return err
				}
				for _, p := range paths {
					f, err := artifact.ParseFile(p)
					if err != nil {
						fmt.Printf("[%s] [parse-error] %s: %v\n", t, p, err)
						continue
					}
					rel, _ := filepath.Rel(vaultRoot, p)
					fmt.Printf("[%-11s] [%-12s] [c=%d] %s — %s\n",
						f.Frontmatter.Type,
						f.Frontmatter.Status,
						f.Frontmatter.Confidence,
						f.Frontmatter.ID,
						rel,
					)
				}
			}
			return nil
		},
	}
	return cmd
}
