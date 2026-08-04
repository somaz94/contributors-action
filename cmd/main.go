package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/somaz94/contributors-action/internal/config"
	"github.com/somaz94/contributors-action/internal/formatter"
	"github.com/somaz94/contributors-action/internal/github"
	"github.com/somaz94/contributors-action/internal/writer"
)

// Output key names — must match action.yml outputs.
const (
	outputContributorsCount = "contributors_count"
	outputOutputFile        = "output_file"
	outputTopContributor    = "top_contributor"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

// run contains the main logic, extracted for testability.
func run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	client := github.NewClient(cfg.Token)
	return execute(ctx, cfg, client)
}

// execute runs the core logic with the given config and client.
func execute(ctx context.Context, cfg *config.Config, client *github.Client) error {
	contributors, err := client.FetchContributors(ctx, cfg.Owner, cfg.Repo, cfg.IncludeBots)
	if err != nil {
		return fmt.Errorf("failed to fetch contributors: %w", err)
	}

	contributors = applyPipeline(contributors, cfg)
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

	setOutput(outputContributorsCount, fmt.Sprintf("%d", len(contributors)))
	setOutput(outputOutputFile, cfg.OutputFile)
	setOutput(outputTopContributor, topContributor)

	fmt.Printf("Successfully processed %d contributors\n", len(contributors))
	return nil
}

// applyPipeline runs the post-fetch transforms: exclude filter, sort, then
// max-contributors truncation.
func applyPipeline(contributors []github.Contributor, cfg *config.Config) []github.Contributor {
	contributors = github.Filter(contributors, cfg.Exclude)
	contributors = github.Sort(contributors, cfg.SortBy)
	if cfg.MaxContributors > 0 && len(contributors) > cfg.MaxContributors {
		contributors = contributors[:cfg.MaxContributors]
	}
	return contributors
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
	// os.File writes are unbuffered, so Close carries no pending data.
	defer func() { _ = f.Close() }()
	fmt.Fprintf(f, "%s=%s\n", name, value)
}
