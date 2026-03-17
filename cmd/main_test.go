package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/somaz94/contributors-action/internal/config"
	"github.com/somaz94/contributors-action/internal/github"
)

func setupMockServer(t *testing.T) *httptest.Server {
	t.Helper()
	contributors := []github.Contributor{
		{Login: "alice", ID: 1, Contributions: 100, Type: "User", AvatarURL: "https://example.com/alice.png", HTMLURL: "https://github.com/alice"},
		{Login: "bob", ID: 2, Contributions: 50, Type: "User", AvatarURL: "https://example.com/bob.png", HTMLURL: "https://github.com/bob"},
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(contributors)
	}))
}

func TestSetOutputWithFile(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "github-output-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	os.Setenv("GITHUB_OUTPUT", tmpFile.Name())
	defer os.Unsetenv("GITHUB_OUTPUT")

	setOutput("test_key", "test_value")

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	if !strings.Contains(string(data), "test_key=test_value") {
		t.Errorf("expected output to contain test_key=test_value, got %q", string(data))
	}
}

func TestSetOutputWithoutFile(t *testing.T) {
	os.Unsetenv("GITHUB_OUTPUT")
	// Should not panic, just prints to stdout
	setOutput("key", "value")
}

func TestSetOutputInvalidPath(t *testing.T) {
	os.Setenv("GITHUB_OUTPUT", "/nonexistent/dir/output")
	defer os.Unsetenv("GITHUB_OUTPUT")

	// Should not panic, logs a warning
	setOutput("key", "value")
}

func TestExecuteDryRun(t *testing.T) {
	server := setupMockServer(t)
	defer server.Close()

	client := github.NewClientWithBaseURL("test-token", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: "CONTRIBUTORS.md",
		Format:     "table",
		Columns:    6,
		AvatarSize: 100,
		SortBy:     "contributions",
		DryRun:     true,
	}

	os.Unsetenv("GITHUB_OUTPUT")

	err := execute(cfg, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExecuteWriteFile(t *testing.T) {
	server := setupMockServer(t)
	defer server.Close()

	tmpFile, err := os.CreateTemp("", "contributors-*.md")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	client := github.NewClientWithBaseURL("test-token", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: tmpFile.Name(),
		Format:     "table",
		Columns:    6,
		AvatarSize: 100,
		SortBy:     "contributions",
		DryRun:     false,
	}

	os.Unsetenv("GITHUB_OUTPUT")

	err = execute(cfg, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "<table>") {
		t.Error("expected table in output file")
	}
	if !strings.Contains(string(data), "alice") {
		t.Error("expected alice in output file")
	}
}

func TestExecuteWithMaxContributors(t *testing.T) {
	server := setupMockServer(t)
	defer server.Close()

	client := github.NewClientWithBaseURL("", server.URL)

	cfg := &config.Config{
		Owner:           "somaz94",
		Repo:            "contributors-action",
		OutputFile:      "CONTRIBUTORS.md",
		Format:          "list",
		Columns:         6,
		AvatarSize:      100,
		SortBy:          "contributions",
		MaxContributors: 1,
		DryRun:          true,
	}

	os.Unsetenv("GITHUB_OUTPUT")

	err := execute(cfg, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExecuteWithExclude(t *testing.T) {
	server := setupMockServer(t)
	defer server.Close()

	client := github.NewClientWithBaseURL("", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: "CONTRIBUTORS.md",
		Format:     "image",
		Columns:    6,
		AvatarSize: 80,
		SortBy:     "name",
		Exclude:    []string{"alice"},
		DryRun:     true,
	}

	os.Unsetenv("GITHUB_OUTPUT")

	err := execute(cfg, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExecuteWithGitHubOutput(t *testing.T) {
	server := setupMockServer(t)
	defer server.Close()

	tmpOutput, err := os.CreateTemp("", "github-output-*")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpOutput.Name())
	tmpOutput.Close()

	os.Setenv("GITHUB_OUTPUT", tmpOutput.Name())
	defer os.Unsetenv("GITHUB_OUTPUT")

	client := github.NewClientWithBaseURL("", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: "CONTRIBUTORS.md",
		Format:     "table",
		Columns:    6,
		AvatarSize: 100,
		SortBy:     "contributions",
		DryRun:     true,
	}

	err = execute(cfg, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	data, err := os.ReadFile(tmpOutput.Name())
	if err != nil {
		t.Fatal(err)
	}
	output := string(data)
	if !strings.Contains(output, "contributors_count=2") {
		t.Errorf("expected contributors_count=2 in output, got %q", output)
	}
	if !strings.Contains(output, "top_contributor=alice") {
		t.Errorf("expected top_contributor=alice in output, got %q", output)
	}
}

func TestExecuteFetchError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("server error"))
	}))
	defer server.Close()

	client := github.NewClientWithBaseURL("", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: "CONTRIBUTORS.md",
		Format:     "table",
		Columns:    6,
		AvatarSize: 100,
		SortBy:     "contributions",
		DryRun:     true,
	}

	err := execute(cfg, client)
	if err == nil {
		t.Fatal("expected error for API failure")
	}
	if !strings.Contains(err.Error(), "failed to fetch contributors") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestExecuteWriteError(t *testing.T) {
	server := setupMockServer(t)
	defer server.Close()

	client := github.NewClientWithBaseURL("", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: "/nonexistent/dir/output.md",
		Format:     "table",
		Columns:    6,
		AvatarSize: 100,
		SortBy:     "contributions",
		DryRun:     false,
	}

	os.Unsetenv("GITHUB_OUTPUT")

	err := execute(cfg, client)
	if err == nil {
		t.Fatal("expected error for write failure")
	}
	if !strings.Contains(err.Error(), "failed to write output") {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestExecuteEmptyContributors(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	}))
	defer server.Close()

	client := github.NewClientWithBaseURL("", server.URL)

	cfg := &config.Config{
		Owner:      "somaz94",
		Repo:       "contributors-action",
		OutputFile: "CONTRIBUTORS.md",
		Format:     "table",
		Columns:    6,
		AvatarSize: 100,
		SortBy:     "contributions",
		DryRun:     true,
	}

	os.Unsetenv("GITHUB_OUTPUT")

	err := execute(cfg, client)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunConfigError(t *testing.T) {
	os.Unsetenv("GITHUB_REPOSITORY")
	os.Unsetenv("INPUT_OWNER")
	os.Unsetenv("INPUT_REPO")

	err := run()
	if err == nil {
		t.Fatal("expected error for missing config")
	}
	if !strings.Contains(err.Error(), "failed to load config") {
		t.Errorf("unexpected error: %v", err)
	}
}
