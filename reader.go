package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// ReadCSV reads the data from the csv to a map struct taking into account
// the most recent timestamp for each repository
func ReadCSV(filePath string) (map[string]*Repo, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)

	// discard the header
	_, err = reader.Read()
	if err != nil {
		if err == io.EOF {
			return nil, nil
		}
		return nil, err
	}

	repoMap := make(map[string]*Repo)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		// timestamp
		timestamp, err := strconv.ParseInt(row[0], 10, 64)
		if err != nil {
			log.Println("error parsing timestamp:", err)
			return nil, err
		}

		// file changes
		fileChanges, err := strconv.Atoi(row[3])
		if err != nil {
			log.Println("error converting file changes to int:", err)
			return nil, err
		}

		// additions
		adds, err := strconv.Atoi(row[4])
		if err != nil {
			log.Println("error converting additions to int:", err)
			return nil, err
		}

		// deletions
		dels, err := strconv.Atoi(row[5])
		if err != nil {
			log.Println("error converting deletions to int:", err)
			return nil, err
		}

		repo, ok := repoMap[row[2]]
		if !ok {
			repo = &Repo{Name: row[2]}
			repoMap[row[2]] = repo
		}

		repo.TotalCommits += 1
		repo.TotalAdds += adds
		repo.TotalDel += dels
		repo.TotalFileChanges += fileChanges

		if repo.LastActivity < timestamp {
			repo.LastActivity = timestamp
		}

	}

	return repoMap, nil
}
