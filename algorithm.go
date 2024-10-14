package main

import (
	"math"
	"math/rand"
)

type Record struct {
	challenge            int
	difference           int
	difficulty           float64
	randomValue          float64
	won                  bool
	estimatedLevelBefore float64
	estimatedLevelAfter  float64
	temporaryLevels      []int
}

// RunDynamicDifficultyAdjustmentAlgorithm
// The algorithm looks like this:
// - we have 11 challenges with levels from 0 to 10
// - the player's level is also expressed on the same scale (a player of
// (level 5 is considered to have 50% chance to lose a challenge of level 5)
// - player level: the mean of the 5 last temporary levels
// - player temporary level:
//   - first, we compute a temporary level for the player
//   - player's temporary level changes if
//   - he wins a hard challenge (player of level n loses a challenge of
//     level m >= n, player gets temporary level m+1)
//   - loses an easy one (player of level n loses a challenge of level
//     m <= n, player gets temporary level m-1)
func RunDynamicDifficultyAdjustmentAlgorithm(p *Player, nSessions int) (records []Record) {
	for i := 0; i < nSessions; i++ {
		r := Record{}
		r.estimatedLevelBefore = p.EstimatedLevel()
		r.challenge = NextChallengeConstantDifficultySlope(r.estimatedLevelBefore)
		r.difference, r.difficulty, r.randomValue, r.won = p.PlayChallenge(r.challenge)
		p.RegisterResult(r.challenge, r.won)
		r.estimatedLevelAfter = p.EstimatedLevel()
		r.temporaryLevels = p.TemporaryLevels
		records = append(records, r)
	}
	return
}

// NextChallengeConstantDifficultySlope - attention, NOT difficulty CURVE
// Choose the challenge according to the player's estimated level.
// Our target is to reach the player's actual level. So if the estimated
// level is 5, we give him a challenge of 5.
// But what if it's 4.5?
// If we always give him a challenge of 5 and he loses, his score
// doesn't change, because he's expected to lose.
// If we always give him a challenge of 4 and he wins, his score
// doesn't change, because he's expected to win.
// So we must alternate.
func NextChallengeConstantDifficultySlope(estimatedPlayerLevel float64) int {
	smallChallenge := int(math.Floor(estimatedPlayerLevel))
	bigChallenge := int(math.Ceil(estimatedPlayerLevel))
	challenge := smallChallenge
	if rand.Int()%2 == 0 {
		challenge = bigChallenge
	}
	return challenge
}
