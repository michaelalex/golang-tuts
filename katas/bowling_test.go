package katas

import "testing"

func Test_CalculateScore(t *testing.T) {
	var expectedScore = 10
	game := Game{bowls: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	var actualScore = CalculateScore(&game)

	if actualScore != expectedScore {
		t.Errorf("Score is not correct expected %d actual %d", expectedScore, actualScore)
	}
}

func Test_CalculateScoreWithStrike(t *testing.T) {
	var expectedScore = 21
	game := Game{bowls: []int{10, 1, 1, 1, 1, 1, 1, 1, 1, 1}}
	var actualScore = CalculateScore(&game)

	if actualScore != expectedScore {
		t.Errorf("Score is not correct expected %d actual %d", expectedScore, actualScore)
	}
}

func Test_CalculateScoreWithSpare(t *testing.T) {
	var expectedScore = 19
	game := Game{bowls: []int{6, 4, 1, 1, 1, 1, 1, 1, 1, 1}}
	var actualScore = CalculateScore(&game)

	if actualScore != expectedScore {
		t.Errorf("Score is not correct expected %d actual %d", expectedScore, actualScore)
	}
}
