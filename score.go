package main

import (
	"cmp"
	"slices"
)

// CalculateScores loops the repository map and calculates the score for each one
// returning a slice
func CalculateScores(repos map[string]*Repo) []*Repo {
	arr := make([]*Repo, len(repos))
	i := 0
	for _, v := range repos {
		v.CalculateScore()
		arr[i] = v
		i++
	}

	return arr
}

// Rank sorts the repository slice by their score and returns the first X top positions
func Rank(repos []*Repo, top int) []*Repo {
	slices.SortFunc(repos,
		func(a, b *Repo) int {
			return cmp.Compare(b.Score, a.Score)
		})

	if len(repos) < top {
		return repos
	}

	return repos[:top]
}
