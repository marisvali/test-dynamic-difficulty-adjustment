package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

// The article "Difficulty in Video Games - An Experimental Validation of
// a Formal Definition" uses an algorithm for finding the level of a player
// based on his performance in the previous levels. They call it "The
// Dynamic Difficulty Adjustment (DDA) algorithm". The purpose of this repo
// is to simulate that algorithm in order to answer some questions like:
// - How long would it take to converge to the player's actual level?
// - What does the road to convergence look like for a player? How many
// wins/losses do you deal with before you end up converging?
func main() {
	for rSeed := 1; rSeed <= 5; rSeed++ {
		rand.Seed(int64(rSeed))

		p := Player{}
		// The article doesn't specify the starting temporary level, so let's assume
		// it's 5.
		p.AddTemporaryLevel(5)
		// Let's assume the player's actual level is 2.
		p.ActualLevel = 2

		// Let's see what happens now if we simulate the difficulty curve that tries
		// to match the level of the challenge given to the player with the player's
		// actual level.
		records := RunDynamicDifficultyAdjustmentAlgorithm(&p, 100)
		OutputRecords(records, fmt.Sprintf("output_%d.csv", rSeed))
	}
}

func OutputRecords(records []Record, filename string) {
	// Create a new CSV file
	file, err := os.Create(filename)
	Check(err)
	// Ensure the file is closed when done
	defer func(file *os.File) { Check(file.Close()) }(file)
	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush() // Ensure all data is written to the file
	err = writer.Write([]string{
		"challenge_level",
		"difference",
		"difficulty",
		"random_value",
		"won",
		"old_estimated_level",
		"new_estimated_level",
		"temporary_levels"})
	Check(err)
	for _, r := range records {
		err = writer.Write([]string{
			fmt.Sprintf("%d", r.challenge),
			fmt.Sprintf("%d", r.difference),
			fmt.Sprintf("%.2f", r.difficulty),
			fmt.Sprintf("%.2f", r.randomValue),
			fmt.Sprintf("%t", r.won),
			fmt.Sprintf("%.2f", r.estimatedLevelBefore),
			fmt.Sprintf("%.2f", r.estimatedLevelAfter),
			fmt.Sprintf("%v", r.temporaryLevels)})
		Check(err)
	}
}
