package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	apiBaseURL         = "https://api.github.com"
	defaultHTTPTimeout = 30 * time.Second
	perPage            = 100
)

// Client is a GitHub API client.
type Client struct {
	token      string
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new GitHub API client.
func NewClient(token string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{Timeout: defaultHTTPTimeout},
		baseURL:    apiBaseURL,
	}
}

// NewClientWithBaseURL creates a new client with a custom base URL (for testing).
func NewClientWithBaseURL(token, baseURL string) *Client {
	return &Client{
		token:      token,
		httpClient: &http.Client{Timeout: defaultHTTPTimeout},
		baseURL:    baseURL,
	}
}

// FetchContributors fetches all contributors for the given repository.
// The context bounds the entire pagination loop.
func (c *Client) FetchContributors(ctx context.Context, owner, repo string, includeBots bool) ([]Contributor, error) {
	var allContributors []Contributor
	page := 1

	for {
		url := fmt.Sprintf("%s/repos/%s/%s/contributors?per_page=%d&page=%d", c.baseURL, owner, repo, perPage, page)
		contributors, err := c.fetchPage(ctx, url)
		if err != nil {
			return nil, err
		}

		if len(contributors) == 0 {
			break
		}

		allContributors = append(allContributors, contributors...)

		if len(contributors) < perPage {
			break
		}
		page++
	}

	if !includeBots {
		allContributors = FilterBots(allContributors)
	}

	return allContributors, nil
}

// fetchPage fetches a single page. Body is drained and closed before return
// to keep connection reuse healthy and avoid per-loop defer accumulation.
func (c *Client) fetchPage(ctx context.Context, url string) ([]Contributor, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fetching contributors: %w", err)
	}
	defer func() {
		// Drain and close so the connection can be reused; neither failure is
		// actionable once the response has been read.
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("GitHub API returned %d: %s", resp.StatusCode, string(body))
	}

	var contributors []Contributor
	if err := json.NewDecoder(resp.Body).Decode(&contributors); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}

	return contributors, nil
}
