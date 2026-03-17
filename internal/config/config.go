package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Token           string
	Owner           string
	Repo            string
	OutputFile      string
	Format          string
	Columns         int
	MaxContributors int
	Exclude         []string
	IncludeBots     bool
	AvatarSize      int
	SortBy          string
	UpdateSection   bool
	SectionStart    string
	SectionEnd      string
	DryRun          bool
}

func Load() (*Config, error) {
	owner, repo := parseRepository()

	columns, err := strconv.Atoi(getEnv("INPUT_COLUMNS", "6"))
	if err != nil {
		return nil, fmt.Errorf("invalid columns value: %w", err)
	}

	maxContributors, err := strconv.Atoi(getEnv("INPUT_MAX_CONTRIBUTORS", "0"))
	if err != nil {
		return nil, fmt.Errorf("invalid max_contributors value: %w", err)
	}

	avatarSize, err := strconv.Atoi(getEnv("INPUT_AVATAR_SIZE", "100"))
	if err != nil {
		return nil, fmt.Errorf("invalid avatar_size value: %w", err)
	}

	excludeStr := getEnv("INPUT_EXCLUDE", "")
	var exclude []string
	if excludeStr != "" {
		for _, e := range strings.Split(excludeStr, ",") {
			trimmed := strings.TrimSpace(e)
			if trimmed != "" {
				exclude = append(exclude, trimmed)
			}
		}
	}

	cfg := &Config{
		Token:           getEnv("INPUT_TOKEN", ""),
		Owner:           getEnvOr("INPUT_OWNER", owner),
		Repo:            getEnvOr("INPUT_REPO", repo),
		OutputFile:      getEnv("INPUT_OUTPUT_FILE", "CONTRIBUTORS.md"),
		Format:          getEnv("INPUT_FORMAT", "table"),
		Columns:         columns,
		MaxContributors: maxContributors,
		Exclude:         exclude,
		IncludeBots:     getEnv("INPUT_INCLUDE_BOTS", "false") == "true",
		AvatarSize:      avatarSize,
		SortBy:          getEnv("INPUT_SORT_BY", "contributions"),
		UpdateSection:   getEnv("INPUT_UPDATE_SECTION", "false") == "true",
		SectionStart:    getEnv("INPUT_SECTION_START", "<!-- CONTRIBUTORS-START -->"),
		SectionEnd:      getEnv("INPUT_SECTION_END", "<!-- CONTRIBUTORS-END -->"),
		DryRun:          getEnv("INPUT_DRY_RUN", "false") == "true",
	}

	if cfg.Owner == "" || cfg.Repo == "" {
		return nil, fmt.Errorf("owner and repo must be set via inputs or GITHUB_REPOSITORY")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parseRepository() (string, string) {
	repo := os.Getenv("GITHUB_REPOSITORY")
	if repo == "" {
		return "", ""
	}
	parts := strings.SplitN(repo, "/", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}
