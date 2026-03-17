package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	os.Setenv("GITHUB_REPOSITORY", "somaz94/contributors-action")
	os.Setenv("INPUT_TOKEN", "test-token")
	os.Setenv("INPUT_OUTPUT_FILE", "CONTRIBUTORS.md")
	os.Setenv("INPUT_FORMAT", "table")
	os.Setenv("INPUT_COLUMNS", "6")
	os.Setenv("INPUT_MAX_CONTRIBUTORS", "10")
	os.Setenv("INPUT_EXCLUDE", "bot1,bot2")
	os.Setenv("INPUT_INCLUDE_BOTS", "false")
	os.Setenv("INPUT_AVATAR_SIZE", "100")
	os.Setenv("INPUT_SORT_BY", "contributions")
	os.Setenv("INPUT_UPDATE_SECTION", "false")
	os.Setenv("INPUT_DRY_RUN", "true")
	defer func() {
		for _, key := range []string{
			"GITHUB_REPOSITORY", "INPUT_TOKEN", "INPUT_OUTPUT_FILE", "INPUT_FORMAT",
			"INPUT_COLUMNS", "INPUT_MAX_CONTRIBUTORS", "INPUT_EXCLUDE", "INPUT_INCLUDE_BOTS",
			"INPUT_AVATAR_SIZE", "INPUT_SORT_BY", "INPUT_UPDATE_SECTION", "INPUT_DRY_RUN",
		} {
			os.Unsetenv(key)
		}
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.Owner != "somaz94" {
		t.Errorf("expected owner somaz94, got %s", cfg.Owner)
	}
	if cfg.Repo != "contributors-action" {
		t.Errorf("expected repo contributors-action, got %s", cfg.Repo)
	}
	if cfg.MaxContributors != 10 {
		t.Errorf("expected max_contributors 10, got %d", cfg.MaxContributors)
	}
	if len(cfg.Exclude) != 2 {
		t.Errorf("expected 2 excludes, got %d", len(cfg.Exclude))
	}
	if !cfg.DryRun {
		t.Error("expected dry_run to be true")
	}
}

func TestLoadMissingRepository(t *testing.T) {
	os.Unsetenv("GITHUB_REPOSITORY")
	os.Unsetenv("INPUT_OWNER")
	os.Unsetenv("INPUT_REPO")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error when repository info is missing")
	}
}

func TestLoadInvalidColumns(t *testing.T) {
	os.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	os.Setenv("INPUT_COLUMNS", "abc")
	defer func() {
		os.Unsetenv("GITHUB_REPOSITORY")
		os.Unsetenv("INPUT_COLUMNS")
	}()

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid columns")
	}
}

func TestLoadInvalidMaxContributors(t *testing.T) {
	os.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	os.Setenv("INPUT_MAX_CONTRIBUTORS", "xyz")
	defer func() {
		os.Unsetenv("GITHUB_REPOSITORY")
		os.Unsetenv("INPUT_MAX_CONTRIBUTORS")
	}()

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid max_contributors")
	}
}

func TestLoadInvalidAvatarSize(t *testing.T) {
	os.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	os.Setenv("INPUT_AVATAR_SIZE", "big")
	defer func() {
		os.Unsetenv("GITHUB_REPOSITORY")
		os.Unsetenv("INPUT_AVATAR_SIZE")
	}()

	_, err := Load()
	if err == nil {
		t.Fatal("expected error for invalid avatar_size")
	}
}

func TestLoadWithExplicitOwnerRepo(t *testing.T) {
	os.Unsetenv("GITHUB_REPOSITORY")
	os.Setenv("INPUT_OWNER", "custom-owner")
	os.Setenv("INPUT_REPO", "custom-repo")
	defer func() {
		os.Unsetenv("INPUT_OWNER")
		os.Unsetenv("INPUT_REPO")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Owner != "custom-owner" {
		t.Errorf("expected owner custom-owner, got %s", cfg.Owner)
	}
	if cfg.Repo != "custom-repo" {
		t.Errorf("expected repo custom-repo, got %s", cfg.Repo)
	}
}

func TestLoadExcludeWithSpaces(t *testing.T) {
	os.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	os.Setenv("INPUT_EXCLUDE", " bot1 , bot2 , , bot3 ")
	defer func() {
		os.Unsetenv("GITHUB_REPOSITORY")
		os.Unsetenv("INPUT_EXCLUDE")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Exclude) != 3 {
		t.Errorf("expected 3 excludes, got %d: %v", len(cfg.Exclude), cfg.Exclude)
	}
}

func TestLoadBooleanFields(t *testing.T) {
	os.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	os.Setenv("INPUT_INCLUDE_BOTS", "true")
	os.Setenv("INPUT_UPDATE_SECTION", "true")
	os.Setenv("INPUT_DRY_RUN", "false")
	defer func() {
		os.Unsetenv("GITHUB_REPOSITORY")
		os.Unsetenv("INPUT_INCLUDE_BOTS")
		os.Unsetenv("INPUT_UPDATE_SECTION")
		os.Unsetenv("INPUT_DRY_RUN")
	}()

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !cfg.IncludeBots {
		t.Error("expected include_bots to be true")
	}
	if !cfg.UpdateSection {
		t.Error("expected update_section to be true")
	}
	if cfg.DryRun {
		t.Error("expected dry_run to be false")
	}
}

func TestParseRepository(t *testing.T) {
	tests := []struct {
		input     string
		wantOwner string
		wantRepo  string
	}{
		{"somaz94/contributors-action", "somaz94", "contributors-action"},
		{"", "", ""},
		{"invalid", "", ""},
	}

	for _, tt := range tests {
		os.Setenv("GITHUB_REPOSITORY", tt.input)
		owner, repo := parseRepository()
		if owner != tt.wantOwner || repo != tt.wantRepo {
			t.Errorf("parseRepository(%q) = (%q, %q), want (%q, %q)", tt.input, owner, repo, tt.wantOwner, tt.wantRepo)
		}
	}
	os.Unsetenv("GITHUB_REPOSITORY")
}
