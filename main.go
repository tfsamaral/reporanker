package main

import (
	"log"
	"os"
	"strconv"
)

const (
	defaultTopRank = 10
)

func main() {
	if len(os.Args) == 1 {
		log.Fatalln("please provide a file path to read")
	}

	filePath := os.Args[1]

	topRank := defaultTopRank
	if len(os.Args) > 2 {
		topRankArg, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Println("invalid top rank:", err)
			log.Println("using default", defaultTopRank)
			topRankArg = defaultTopRank
		}

		topRank = topRankArg
	}

	// parse csv file
	repoMap, err := ReadCSV(filePath)
	if err != nil {
		log.Fatalln("error getting repository information from file:", err)
	}

	if len(repoMap) == 0 {
		log.Fatalln("no repositories, file might be empty")
	}

	// calculate repository scores
	repos := CalculateScores(repoMap)

	// rank the repositories based on top rank option
	rankedRepos := Rank(repos, topRank)

	// print to console
	for i, v := range rankedRepos {
		log.Println("Rank:", i+1, "Name:", v.Name, "Score:", v.Score)
	}

	// write the result to new csv file
	if err := WriteCSV(rankedRepos); err != nil {
		log.Fatalln("error writing resulting csv:", err)
	}

}
