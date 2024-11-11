package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// WriteCSV writes the given repositories information to a new csv file
func WriteCSV(repos []*Repo) error {
	f, err := os.Create("results.csv")
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	if err := w.Write(
		[]string{
			"name",
			"commits",
			"file_changes",
			"additions",
			"deletions",
			"last_activity",
			"score",
		}); err != nil {
		return err
	}

	for _, v := range repos {
		if err := w.Write(
			[]string{
				v.Name,
				strconv.Itoa(v.TotalCommits),
				strconv.Itoa(v.TotalFileChanges),
				strconv.Itoa(v.TotalAdds),
				strconv.Itoa(v.TotalDel),
				strconv.FormatInt(v.LastActivity, 10),
				strconv.Itoa(int(v.Score)),
			},
		); err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}
