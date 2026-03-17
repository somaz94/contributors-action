package formatter

import (
	"strings"
	"testing"

	"github.com/somaz94/contributors-action/internal/github"
)

var testContributors = []github.Contributor{
	{Login: "alice", AvatarURL: "https://example.com/alice.png", HTMLURL: "https://github.com/alice", Contributions: 100},
	{Login: "bob", AvatarURL: "https://example.com/bob.png", HTMLURL: "https://github.com/bob", Contributions: 50},
	{Login: "charlie", AvatarURL: "https://example.com/charlie.png", HTMLURL: "https://github.com/charlie", Contributions: 20},
}

func TestFormatTable(t *testing.T) {
	result := Format(testContributors, "table", 2, 100)

	if !strings.Contains(result, "<table>") {
		t.Error("expected table tag")
	}
	if !strings.Contains(result, "alice") {
		t.Error("expected alice in output")
	}
	if !strings.Contains(result, "width=\"100\"") {
		t.Error("expected avatar size 100")
	}
	if strings.Count(result, "<tr>") != 2 {
		t.Errorf("expected 2 rows for 3 contributors with 2 columns, got %d", strings.Count(result, "<tr>"))
	}
}

func TestFormatTableEmpty(t *testing.T) {
	result := Format(nil, "table", 6, 100)
	if result != "" {
		t.Errorf("expected empty string for nil contributors, got %q", result)
	}
}

func TestFormatList(t *testing.T) {
	result := Format(testContributors, "list", 6, 50)

	if !strings.Contains(result, "- [") {
		t.Error("expected list format")
	}
	if !strings.Contains(result, "100 contributions") {
		t.Error("expected contribution count")
	}
	if !strings.Contains(result, "width=\"50\"") {
		t.Error("expected avatar size 50")
	}
}

func TestFormatImage(t *testing.T) {
	result := Format(testContributors, "image", 6, 80)

	if !strings.Contains(result, "title=\"alice\"") {
		t.Error("expected title attribute")
	}
	if !strings.Contains(result, "width=\"80\"") {
		t.Error("expected avatar size 80")
	}
}

func TestFormatDefaultIsTable(t *testing.T) {
	result := Format(testContributors, "unknown", 6, 100)
	if !strings.Contains(result, "<table>") {
		t.Error("expected table format as default")
	}
}
