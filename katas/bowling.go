package katas

type Game struct {
	bowls []int
}

func CalculateScore(game *Game) int {
	var score = 0
	for index := 0; index < len(game.bowls); index++ {
		var bowl = game.bowls[index]
		if isStrike(bowl) {
			score += bowl + game.bowls[index+1] + game.bowls[index+2]
		} else if index > 0 && isSpare(game.bowls[index-1], bowl) {
			score += bowl + game.bowls[index+1]
		} else {
			score += bowl
		}
	}
	return score
}

func isStrike(pins int) bool {
	return pins == 10
}

func isSpare(pinsA int, pinsB int) bool {
	return pinsA+pinsB == 10
}
