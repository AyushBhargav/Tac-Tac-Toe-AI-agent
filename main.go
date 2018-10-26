package main

import (
	"TicTacToe/ai"
	"TicTacToe/game"
	"fmt"
)

func main() {
	gameState := &game.State{}
	gameState.New()
	for game.WhoWon(gameState) == game.VOID {
		gameState.Render()
		if gameState.Player == game.PLAYER1 { //
			fmt.Printf("Your turn row col->")
			var row, col int
			_, err := fmt.Scanf("%d %d\n", &row, &col)
			if err != nil {
				panic(err)
			}
			gameState.Update(row, col)
		} else { // AI
			moveAi := ai.Move(*gameState)
			fmt.Printf("Ai wants to take (%d, %d).\n", moveAi[0], moveAi[1])
			gameState.Update(moveAi[0], moveAi[1])
		}
	}
	winnerPlayer := game.WhoWon(gameState)
	switch winnerPlayer {
	case game.PLAYER1:
		fmt.Printf("you won :(.")
	case game.PLAYER2:
		fmt.Printf("AI won. You are not that smart after all :D.")
	case game.TIE:
		fmt.Printf("Match Tied.")
	}
	fmt.Scanln()
}
