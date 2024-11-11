# RepoRanker

Small project that reads a csv file with multiple repository commits and ranks them based on certain conditions.

## How it works

### Reader
The commits are extracted from the csv file and organized by the repository name in a `map[string]*Repo`, calculating the total number of commits, file changes, additions, deletions and the most recent timestamp for each repository.

``` Go
type Repo struct {
	Name             string
	TotalCommits     int
	TotalFileChanges int
	TotalAdds        int
	TotalDel         int
	LastActivity     int64
	Score            float64
}
```

### Calculate Score
The score is calculated based on the total commits, files changes, additions and deletions with their respective weight of importance:

```Go
const (
	commitsWeight   = 0.4
	changesWeight   = 0.3
	filesWeight     = 0.2
	startTimeDecay  = 15 // days
)
```
So for each repository the totals are multiplied by their corresponding weight and added to get a total score:
```Go
score := float64(r.TotalAdds+r.TotalDel)*changesWeight+
		float64(r.TotalCommits)*commitsWeight +
		float64(r.TotalFileChanges)*filesWeight
```
Additionally there is a time weight, that starts at 1 and gradually decays for inactivity beyond 15 days.  
The time score is applied to the initial score:

```Go
timeWeight := 1.0

currentTime := time.Now().UTC().Truncate(24 * time.Hour)

days := time.Now().UTC().Sub(time.Unix(r.LastActivity, 0)).Hours() / 24
if days > startTimeDecay {
    timeWeight = startTimeDecay / days
}

r.Score = math.Round((score * timeWeight))
```

### Result
The results are printed to the console, but also written to a new `results.csv` file in the project root.

## CLI Usage

```bash
$ go run . [input_file_path] [top_rank]
```
### Arguments

- **`input_file_path`**  
  The path to the CSV file containing the repository data.
  
  **Example**:  
  ```bash
  $ go run . /path/to/commits.csv
  ```

- **`top_rank`**  
  Optional integer specifying how many top repositories to display.
  The default value is `10`.
  
  **Example**:  
  ```bash
  $ go run . /path/to/commits.csv 5
  ```

## Top 10 Most Active Repositories:

1. **repo476** - Score: 12234
2. **repo260** - Score: 3831
3. **repo920** - Score: 2211
4. **repo795** - Score: 1900
5. **repo161** - Score: 1391
6. **repo1143** - Score: 1306
7. **repo518** - Score: 1187
8. **repo1185** - Score: 1013
9. **repo1243** - Score: 940
10. **repo250** - Score: 811


**Note**: the scores can change slightly based on the current time, as it is used in the calculation.  
The current results where calculated on **2024-11-10**.