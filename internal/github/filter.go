package github

import (
	"sort"
	"strings"
)

// Sort field identifiers accepted by Sort.
const (
	SortByContributions = "contributions"
	SortByName          = "name"
)

// Filter removes contributors whose login matches any entry in the exclude list.
func Filter(contributors []Contributor, exclude []string) []Contributor {
	if len(exclude) == 0 {
		return contributors
	}

	excludeMap := make(map[string]bool, len(exclude))
	for _, e := range exclude {
		excludeMap[strings.ToLower(e)] = true
	}

	filtered := make([]Contributor, 0, len(contributors))
	for _, c := range contributors {
		if !excludeMap[strings.ToLower(c.Login)] {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

// FilterBots removes contributors with Type "Bot".
func FilterBots(contributors []Contributor) []Contributor {
	filtered := make([]Contributor, 0, len(contributors))
	for _, c := range contributors {
		if c.Type != "Bot" {
			filtered = append(filtered, c)
		}
	}
	return filtered
}

// Sort sorts contributors by the given field.
func Sort(contributors []Contributor, sortBy string) []Contributor {
	sorted := make([]Contributor, len(contributors))
	copy(sorted, contributors)

	switch sortBy {
	case SortByName:
		sort.Slice(sorted, func(i, j int) bool {
			return strings.ToLower(sorted[i].Login) < strings.ToLower(sorted[j].Login)
		})
	default: // SortByContributions — also the fallback for unknown values.
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].Contributions > sorted[j].Contributions
		})
	}
	return sorted
}
