package github

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchContributors(t *testing.T) {
	mockContributors := []Contributor{
		{Login: "alice", ID: 1, Contributions: 100, Type: "User", AvatarURL: "https://example.com/alice.png", HTMLURL: "https://github.com/alice"},
		{Login: "bob", ID: 2, Contributions: 50, Type: "User", AvatarURL: "https://example.com/bob.png", HTMLURL: "https://github.com/bob"},
		{Login: "ci-bot", ID: 3, Contributions: 30, Type: "Bot", AvatarURL: "https://example.com/bot.png", HTMLURL: "https://github.com/ci-bot"},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			t.Errorf("unexpected auth header: %s", r.Header.Get("Authorization"))
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockContributors)
	}))
	defer server.Close()

	client := NewClientWithBaseURL("test-token", server.URL)

	contributors, err := client.FetchContributors("owner", "repo", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(contributors) != 2 {
		t.Fatalf("expected 2 contributors (excluding bots), got %d", len(contributors))
	}
	if contributors[0].Login != "alice" {
		t.Errorf("expected first contributor alice, got %s", contributors[0].Login)
	}
}

func TestFetchContributorsIncludeBots(t *testing.T) {
	mockContributors := []Contributor{
		{Login: "alice", Type: "User", Contributions: 100},
		{Login: "ci-bot", Type: "Bot", Contributions: 30},
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockContributors)
	}))
	defer server.Close()

	client := NewClientWithBaseURL("", server.URL)

	contributors, err := client.FetchContributors("owner", "repo", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(contributors) != 2 {
		t.Fatalf("expected 2 contributors (including bots), got %d", len(contributors))
	}
}

func TestFetchContributorsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "Not Found"}`))
	}))
	defer server.Close()

	client := NewClientWithBaseURL("token", server.URL)

	_, err := client.FetchContributors("owner", "repo", false)
	if err == nil {
		t.Fatal("expected error for 404 response")
	}
}

func TestNewClient(t *testing.T) {
	client := NewClient("my-token")
	if client.token != "my-token" {
		t.Errorf("expected token my-token, got %s", client.token)
	}
	if client.baseURL != apiBaseURL {
		t.Errorf("expected baseURL %s, got %s", apiBaseURL, client.baseURL)
	}
	if client.httpClient == nil {
		t.Error("expected httpClient to be non-nil")
	}
}

func TestFetchContributorsEmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
	}))
	defer server.Close()

	client := NewClientWithBaseURL("", server.URL)

	contributors, err := client.FetchContributors("owner", "repo", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(contributors) != 0 {
		t.Errorf("expected 0 contributors, got %d", len(contributors))
	}
}

func TestFetchContributorsInvalidJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{invalid json"))
	}))
	defer server.Close()

	client := NewClientWithBaseURL("", server.URL)

	_, err := client.FetchContributors("owner", "repo", false)
	if err == nil {
		t.Fatal("expected error for invalid JSON")
	}
}

func TestFetchContributorsNoToken(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			t.Error("expected no Authorization header when token is empty")
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Contributor{{Login: "alice", Type: "User"}})
	}))
	defer server.Close()

	client := NewClientWithBaseURL("", server.URL)

	contributors, err := client.FetchContributors("owner", "repo", true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(contributors) != 1 {
		t.Errorf("expected 1 contributor, got %d", len(contributors))
	}
}

func TestFetchContributorsPagination(t *testing.T) {
	callCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		if callCount == 1 {
			contributors := make([]Contributor, 100)
			for i := range contributors {
				contributors[i] = Contributor{Login: "user", Type: "User", Contributions: 1}
			}
			json.NewEncoder(w).Encode(contributors)
		} else {
			contributors := []Contributor{{Login: "last", Type: "User", Contributions: 1}}
			json.NewEncoder(w).Encode(contributors)
		}
	}))
	defer server.Close()

	client := NewClientWithBaseURL("", server.URL)

	contributors, err := client.FetchContributors("owner", "repo", false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(contributors) != 101 {
		t.Errorf("expected 101 contributors, got %d", len(contributors))
	}
	if callCount != 2 {
		t.Errorf("expected 2 API calls, got %d", callCount)
	}
}
