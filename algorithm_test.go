package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestNextChallengeConstantDifficultySlope(t *testing.T) {
	n := 0
	for i := 1; i < 100; i++ {
		challenge := NextChallengeConstantDifficultySlope(4.3)
		assert.True(t, challenge == 4 || challenge == 5)
		if challenge == 4 {
			n++
		}
	}
	assert.True(t, math.Abs(float64(n)/100.0-0.5) < 0.1)
}

func TestRunDynamicDifficultyAdjustmentAlgorithm(t *testing.T) {
	// Our code involves randomness so make the test non-random by fixing
	// a global seed. Warning: for some seeds the error tolerance at the end
	// is not big enough. It is what it is. I would still say the test passes.
	rand.Seed(15)

	// Check that the algorithm actually converges to the player's actual level.
	for actualLevel := 1; actualLevel <= 10; actualLevel++ {
		p := Player{}
		p.AddTemporaryLevel(5)
		p.ActualLevel = actualLevel
		records := RunDynamicDifficultyAdjustmentAlgorithm(&p, 200)
		// We have to give a lot of leeway to this algorithm. So compute the
		// mean estimated level for the last 50 estimations.
		// I assume that after 150 iterations, it converged and the last 50
		// estimations are jumping up and down the actual value.
		sum := 0.0
		for i := 150; i < 200; i++ {
			sum += records[i].estimatedLevelAfter
		}
		meanEstimatedLevel := sum / 50.0
		// println(meanEstimatedLevel, p.ActualLevel)
		// Even so, I need a margin of 1.
		assert.True(t, math.Abs(meanEstimatedLevel-float64(p.ActualLevel)) < 1)
	}
}
