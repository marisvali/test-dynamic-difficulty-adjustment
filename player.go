package main

import "math/rand"

type Player struct {
	ActualLevel     int   // This specifies the player's actual level.
	TemporaryLevels []int // Slice to hold the last 5 temporary levels.
}

// AddTemporaryLevel adds a new temporary level and maintains only the last 5
// temporary levels.
func (p *Player) AddTemporaryLevel(level int) {
	p.TemporaryLevels = append(p.TemporaryLevels, level)
	if len(p.TemporaryLevels) > 5 {
		p.TemporaryLevels = p.TemporaryLevels[1:] // Keep only the last 5 levels
	}
}

// EstimatedLevel calculates the mean of the last 5 temporary levels.
func (p *Player) EstimatedLevel() float64 {
	if len(p.TemporaryLevels) == 0 {
		panic("we must have some temporary levels to compute the level")
	}

	sum := 0
	for _, level := range p.TemporaryLevels {
		sum += level
	}
	return float64(sum) / float64(len(p.TemporaryLevels))
}

// RegisterResult adjusts the player's level according to his latest result.
func (p *Player) RegisterResult(challengeLevel int, won bool) {
	if won {
		if float64(challengeLevel) >= p.EstimatedLevel() {
			p.AddTemporaryLevel(challengeLevel + 1) // Hard challenge won
		}
	} else { // Player lost
		if float64(challengeLevel) <= p.EstimatedLevel() {
			p.AddTemporaryLevel(challengeLevel - 1) // Easy challenge lost
		}
	}
}

// PlayChallenge gives the result of this player playing a challenge of a
// certain level.
func (p *Player) PlayChallenge(challengeLevel int) (difference int,
	difficulty float64, randomValue float64, won bool) {
	difference = p.ActualLevel - challengeLevel
	// We will use the difficulty curve that the authors presumed will hold true
	// before they ran the experiment. This curve assigns a chance for a player
	// of level X to lose a challenge of level Y. This implicitly assumes that
	// a player has SOME chance to win/lose any challenge.
	// EstimatedLevel difference 	Theoretic difficulty
	// < 2 					0.95
	// -2 					0.8
	// -1 					0.6
	// 0					0.5
	// 1					0.4
	// 2					0.2
	// > 2					0.05
	// Polynomial that approximates this rule:
	// difficulty(difference) = -1.464 * difference + 0.5
	difficulty = -0.1464*float64(difference) + 0.5
	// difficulty represents the chance to lose
	// You win if you roll a die that gives you a random value between 0 and 1
	// and your value is larger than difficulty. The larger difficulty is, the
	// fewer your chances are to win.
	randomValue = rand.Float64()
	won = randomValue > difficulty
	return
}
