package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"reflect"
	"testing"
)

func TestInt_AddTemporaryLevel(t *testing.T) {
	p := Player{}
	for i := 1; i <= 5; i++ {
		p.AddTemporaryLevel(i)
		assert.True(t, len(p.TemporaryLevels) == i)
	}
	for i := 1; i < 100; i++ {
		p.AddTemporaryLevel(i)
		assert.True(t, len(p.TemporaryLevels) == 5)
	}
}

func TestInt_EstimatedLevel(t *testing.T) {
	p := Player{}
	assert.Panics(t, func() { p.EstimatedLevel() })
	p.AddTemporaryLevel(1)
	assert.Equal(t, p.EstimatedLevel(), 1.0)
	p.AddTemporaryLevel(2)
	assert.Equal(t, p.EstimatedLevel(), 1.5)
	p.AddTemporaryLevel(3)
	assert.Equal(t, p.EstimatedLevel(), 2.0)
}

func TestInt_RegisterResult(t *testing.T) {
	{
		p := Player{}
		p.AddTemporaryLevel(3)
		p.ActualLevel = 3
		p.RegisterResult(3, true)
		assert.True(t, reflect.DeepEqual(p.TemporaryLevels, []int{3, 4}))
	}

	{
		p := Player{}
		p.AddTemporaryLevel(3)
		p.ActualLevel = 3
		p.RegisterResult(3, false)
		assert.True(t, reflect.DeepEqual(p.TemporaryLevels, []int{3, 2}))
	}

	{
		p := Player{}
		p.AddTemporaryLevel(3)
		p.ActualLevel = 3
		p.RegisterResult(2, true)
		assert.True(t, reflect.DeepEqual(p.TemporaryLevels, []int{3}))
	}

	{
		p := Player{}
		p.AddTemporaryLevel(3)
		p.ActualLevel = 3
		p.RegisterResult(4, false)
		assert.True(t, reflect.DeepEqual(p.TemporaryLevels, []int{3}))
	}
}

func TestInt_PlayChallenge(t *testing.T) {
	{
		p := Player{}
		p.AddTemporaryLevel(3)
		p.ActualLevel = 3
		nLosses := 0
		estimatedDifficulty := 0.0
		for i := 1; i <= 1000; i++ {
			difference, difficulty, randomValue, won := p.PlayChallenge(3)
			estimatedDifficulty = difficulty
			assert.Equal(t, difference, 0)
			assert.True(t, math.Abs(difficulty-0.5) < 0.001)
			assert.True(t, won == (randomValue > difficulty))
			if !won {
				nLosses++
			}
		}
		assert.True(t, math.Abs(float64(nLosses)/1000.0-estimatedDifficulty) < 0.1)
	}

	{
		p := Player{}
		p.AddTemporaryLevel(3)
		p.ActualLevel = 3
		nLosses := 0
		estimatedDifficulty := 0.0
		for i := 1; i <= 1000; i++ {
			difference, difficulty, randomValue, won := p.PlayChallenge(5)
			estimatedDifficulty = difficulty
			assert.Equal(t, difference, -2)
			assert.True(t, math.Abs(difficulty-0.8) < 0.05)
			assert.True(t, won == (randomValue > difficulty))
			if !won {
				nLosses++
			}
		}
		assert.True(t, math.Abs(float64(nLosses)/1000.0-estimatedDifficulty) < 0.1)
	}
}
