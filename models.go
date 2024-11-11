package main

import (
	"math"
	"time"
)

// ranking weights
const (
	commitsWeight  = 0.4
	changesWeight  = 0.3
	filesWeight    = 0.2
	startTimeDecay = 15 // days
)

// Repo represents a repository information
type Repo struct {
	Name             string
	TotalCommits     int
	TotalFileChanges int
	TotalAdds        int
	TotalDel         int
	LastActivity     int64
	Score            float64
}

// CalculateScore calculate the repository activity score
func (r *Repo) CalculateScore() {
	score := float64(r.TotalAdds+r.TotalDel)*changesWeight +
		float64(r.TotalCommits)*commitsWeight +
		float64(r.TotalFileChanges)*filesWeight

	timeWeight := 1.0

	currentTime := time.Now().UTC().Truncate(24 * time.Hour)

	days := currentTime.Sub(time.Unix(r.LastActivity, 0)).Hours() / 24
	if days > startTimeDecay {
		timeWeight = startTimeDecay / days
	}

	r.Score = math.Round((score * timeWeight))
}
