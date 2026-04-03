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

	cfg := &Config{
		Token:           getEnv("INPUT_TOKEN", ""),
		Owner:           getEnv("INPUT_OWNER", owner),
		Repo:            getEnv("INPUT_REPO", repo),
		OutputFile:      getEnv("INPUT_OUTPUT_FILE", "CONTRIBUTORS.md"),
		Format:          getEnv("INPUT_FORMAT", "table"),
		Columns:         columns,
		MaxContributors: maxContributors,
		Exclude:         parseCSV(getEnv("INPUT_EXCLUDE", "")),
		IncludeBots:     getEnvBool("INPUT_INCLUDE_BOTS"),
		AvatarSize:      avatarSize,
		SortBy:          getEnv("INPUT_SORT_BY", "contributions"),
		UpdateSection:   getEnvBool("INPUT_UPDATE_SECTION"),
		SectionStart:    getEnv("INPUT_SECTION_START", "<!-- CONTRIBUTORS-START -->"),
		SectionEnd:      getEnv("INPUT_SECTION_END", "<!-- CONTRIBUTORS-END -->"),
		DryRun:          getEnvBool("INPUT_DRY_RUN"),
	}

	if cfg.Owner == "" || cfg.Repo == "" {
		return nil, fmt.Errorf("owner and repo must be set via inputs or GITHUB_REPOSITORY")
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	switch c.Format {
	case "table", "list", "image":
	default:
		return fmt.Errorf("invalid format %q: must be table, list, or image", c.Format)
	}

	switch c.SortBy {
	case "contributions", "name":
	default:
		return fmt.Errorf("invalid sort_by %q: must be contributions or name", c.SortBy)
	}

	if c.Columns < 1 {
		return fmt.Errorf("columns must be >= 1, got %d", c.Columns)
	}

	if c.AvatarSize < 1 {
		return fmt.Errorf("avatar_size must be >= 1, got %d", c.AvatarSize)
	}

	return nil
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvBool(key string) bool {
	return getEnv(key, "false") == "true"
}

func parseCSV(s string) []string {
	if s == "" {
		return nil
	}
	var result []string
	for _, e := range strings.Split(s, ",") {
		trimmed := strings.TrimSpace(e)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
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
