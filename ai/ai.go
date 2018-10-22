package ai

import "TicTacToe/game"

// import "fmt"
import "math"

// Move returns the move AI should take based on game score
func Move(gameState game.State) []int {
	availableMoves := gameState.GetUnoccupiedCells()
	index := 0
	maxValue := math.MinInt8
	for i, position := range availableMoves {
		newGameState := &game.State{Cells: gameState.Cells, Player: gameState.Player,
			TimesUpdated: gameState.TimesUpdated}
		newGameState.Update(position[0], position[1])
		score := getMinScore(*newGameState, 50)
		// fmt.Printf("Ai score (%d, %d): %d\n", position[0], position[1], score)
		if score > maxValue {
			index = i
			maxValue = score
		}
	}
	return availableMoves[index]
}

func getMaxScore(gameState game.State, depth int) int {
	if game.WhoWon(&gameState) == game.PLAYER2 {
		return math.MaxInt32
	}
	if game.WhoWon(&gameState) == game.TIE {
		return math.MaxInt16
	}
	if game.WhoWon(&gameState) == game.PLAYER1 {
		return math.MinInt32
	}
	if depth == 0 {
		return 0
	}
	availableMoves := gameState.GetUnoccupiedCells()
	maxValue := math.MinInt8
	for _, position := range availableMoves {
		newGameState := &game.State{Cells: gameState.Cells, Player: gameState.Player,
			TimesUpdated: gameState.TimesUpdated}
		newGameState.Update(position[0], position[1])
		score := getMinScore(*newGameState, depth-1)
		if score > maxValue {
			maxValue = score
		}
	}
	return maxValue
}

func getMinScore(gameState game.State, depth int) int {
	if game.WhoWon(&gameState) == game.PLAYER2 {
		return math.MaxInt32
	}
	if game.WhoWon(&gameState) == game.TIE {
		return math.MaxInt16
	}
	if game.WhoWon(&gameState) == game.PLAYER1 {
		return math.MinInt32
	}
	if depth == 0 {
		return 0
	}
	availableMoves := gameState.GetUnoccupiedCells()
	minValue := math.MaxInt8
	for _, position := range availableMoves {
		newGameState := &game.State{Cells: gameState.Cells, Player: gameState.Player,
			TimesUpdated: gameState.TimesUpdated}
		newGameState.Update(position[0], position[1])
		score := getMaxScore(*newGameState, depth-1)
		if score < minValue {
			minValue = score
		}
	}
	return minValue
}
