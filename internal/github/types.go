package github

// Contributor represents a GitHub repository contributor.
type Contributor struct {
	Login         string `json:"login"`
	ID            int    `json:"id"`
	AvatarURL     string `json:"avatar_url"`
	HTMLURL       string `json:"html_url"`
	Contributions int    `json:"contributions"`
	Type          string `json:"type"`
}
