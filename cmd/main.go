package main

import (
	"fmt"
	"log"
	"os"

	"github.com/somaz94/contributors-action/internal/config"
	"github.com/somaz94/contributors-action/internal/formatter"
	"github.com/somaz94/contributors-action/internal/github"
	"github.com/somaz94/contributors-action/internal/writer"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

// run contains the main logic, extracted for testability.
func run() error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := github.NewClient(cfg.Token)
	return execute(cfg, client)
}

// execute runs the core logic with the given config and client.
func execute(cfg *config.Config, client *github.Client) error {
	contributors, err := client.FetchContributors(cfg.Owner, cfg.Repo, cfg.IncludeBots)
	if err != nil {
		return fmt.Errorf("failed to fetch contributors: %w", err)
	}

	contributors = github.Filter(contributors, cfg.Exclude)
	contributors = github.Sort(contributors, cfg.SortBy)

	if cfg.MaxContributors > 0 && len(contributors) > cfg.MaxContributors {
		contributors = contributors[:cfg.MaxContributors]
	}

	content := formatter.Format(contributors, cfg.Format, cfg.Columns, cfg.AvatarSize)

	if cfg.DryRun {
		fmt.Println("--- DRY RUN ---")
		fmt.Println(content)
	} else {
		if err := writer.Write(cfg, content); err != nil {
			return fmt.Errorf("failed to write output: %w", err)
		}
	}

	topContributor := ""
	if len(contributors) > 0 {
		topContributor = contributors[0].Login
	}

	setOutput("contributors_count", fmt.Sprintf("%d", len(contributors)))
	setOutput("output_file", cfg.OutputFile)
	setOutput("top_contributor", topContributor)

	fmt.Printf("Successfully processed %d contributors\n", len(contributors))
	return nil
}

func setOutput(name, value string) {
	outputFile := os.Getenv("GITHUB_OUTPUT")
	if outputFile == "" {
		fmt.Printf("%s=%s\n", name, value)
		return
	}
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o644)
	if err != nil {
		log.Printf("warning: could not write to GITHUB_OUTPUT: %v", err)
		return
	}
	defer f.Close()
	fmt.Fprintf(f, "%s=%s\n", name, value)
}
