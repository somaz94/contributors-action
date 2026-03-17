package github

import (
	"testing"
)

func TestFilter(t *testing.T) {
	contributors := []Contributor{
		{Login: "alice", Contributions: 100},
		{Login: "bob", Contributions: 50},
		{Login: "dependabot[bot]", Contributions: 30},
		{Login: "charlie", Contributions: 20},
	}

	result := Filter(contributors, []string{"dependabot[bot]", "Bob"})
	if len(result) != 2 {
		t.Fatalf("expected 2 contributors, got %d", len(result))
	}
	if result[0].Login != "alice" || result[1].Login != "charlie" {
		t.Errorf("unexpected result: %v", result)
	}
}

func TestFilterEmpty(t *testing.T) {
	contributors := []Contributor{
		{Login: "alice", Contributions: 100},
	}
	result := Filter(contributors, nil)
	if len(result) != 1 {
		t.Fatalf("expected 1 contributor, got %d", len(result))
	}
}

func TestSortByContributions(t *testing.T) {
	contributors := []Contributor{
		{Login: "bob", Contributions: 50},
		{Login: "alice", Contributions: 100},
		{Login: "charlie", Contributions: 20},
	}

	result := Sort(contributors, "contributions")
	if result[0].Login != "alice" || result[1].Login != "bob" || result[2].Login != "charlie" {
		t.Errorf("unexpected sort order: %v", result)
	}
}

func TestSortByName(t *testing.T) {
	contributors := []Contributor{
		{Login: "charlie", Contributions: 20},
		{Login: "alice", Contributions: 100},
		{Login: "Bob", Contributions: 50},
	}

	result := Sort(contributors, "name")
	if result[0].Login != "alice" || result[1].Login != "Bob" || result[2].Login != "charlie" {
		t.Errorf("unexpected sort order: %v", result)
	}
}

func TestSortDoesNotMutateOriginal(t *testing.T) {
	contributors := []Contributor{
		{Login: "bob", Contributions: 50},
		{Login: "alice", Contributions: 100},
	}

	Sort(contributors, "name")
	if contributors[0].Login != "bob" {
		t.Error("original slice was mutated")
	}
}
