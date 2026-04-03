package config

import (
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Setenv("GITHUB_REPOSITORY", "somaz94/contributors-action")
	t.Setenv("INPUT_TOKEN", "test-token")
	t.Setenv("INPUT_OUTPUT_FILE", "CONTRIBUTORS.md")
	t.Setenv("INPUT_FORMAT", "table")
	t.Setenv("INPUT_COLUMNS", "6")
	t.Setenv("INPUT_MAX_CONTRIBUTORS", "10")
	t.Setenv("INPUT_EXCLUDE", "bot1,bot2")
	t.Setenv("INPUT_INCLUDE_BOTS", "false")
	t.Setenv("INPUT_AVATAR_SIZE", "100")
	t.Setenv("INPUT_SORT_BY", "contributions")
	t.Setenv("INPUT_UPDATE_SECTION", "false")
	t.Setenv("INPUT_DRY_RUN", "true")

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
	t.Setenv("GITHUB_REPOSITORY", "")
	t.Setenv("INPUT_OWNER", "")
	t.Setenv("INPUT_REPO", "")

	_, err := Load()
	if err == nil {
		t.Fatal("expected error when repository info is missing")
	}
}

func TestLoadParsing(t *testing.T) {
	t.Run("InvalidColumns", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_COLUMNS", "abc")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for invalid columns")
		}
	})

	t.Run("InvalidMaxContributors", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_MAX_CONTRIBUTORS", "xyz")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for invalid max_contributors")
		}
	})

	t.Run("InvalidAvatarSize", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_AVATAR_SIZE", "big")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for invalid avatar_size")
		}
	})
}

func TestLoadWithExplicitOwnerRepo(t *testing.T) {
	t.Setenv("GITHUB_REPOSITORY", "")
	t.Setenv("INPUT_OWNER", "custom-owner")
	t.Setenv("INPUT_REPO", "custom-repo")

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
	t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	t.Setenv("INPUT_EXCLUDE", " bot1 , bot2 , , bot3 ")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cfg.Exclude) != 3 {
		t.Errorf("expected 3 excludes, got %d: %v", len(cfg.Exclude), cfg.Exclude)
	}
}

func TestLoadBooleanFields(t *testing.T) {
	t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
	t.Setenv("INPUT_INCLUDE_BOTS", "true")
	t.Setenv("INPUT_UPDATE_SECTION", "true")
	t.Setenv("INPUT_DRY_RUN", "false")

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

func TestLoadValidation(t *testing.T) {
	t.Run("InvalidFormat", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_FORMAT", "yaml")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for invalid format")
		}
		if !strings.Contains(err.Error(), "invalid format") {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("InvalidSortBy", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_SORT_BY", "date")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for invalid sort_by")
		}
		if !strings.Contains(err.Error(), "invalid sort_by") {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("ColumnsZero", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_COLUMNS", "0")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for columns = 0")
		}
		if !strings.Contains(err.Error(), "columns must be >= 1") {
			t.Errorf("unexpected error: %v", err)
		}
	})

	t.Run("AvatarSizeZero", func(t *testing.T) {
		t.Setenv("GITHUB_REPOSITORY", "somaz94/test")
		t.Setenv("INPUT_AVATAR_SIZE", "0")

		_, err := Load()
		if err == nil {
			t.Fatal("expected error for avatar_size = 0")
		}
		if !strings.Contains(err.Error(), "avatar_size must be >= 1") {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestParseRepository(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		wantOwner string
		wantRepo  string
	}{
		{"ValidRepo", "somaz94/contributors-action", "somaz94", "contributors-action"},
		{"Empty", "", "", ""},
		{"Invalid", "invalid", "", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("GITHUB_REPOSITORY", tt.input)
			owner, repo := parseRepository()
			if owner != tt.wantOwner || repo != tt.wantRepo {
				t.Errorf("parseRepository(%q) = (%q, %q), want (%q, %q)", tt.input, owner, repo, tt.wantOwner, tt.wantRepo)
			}
		})
	}
}
