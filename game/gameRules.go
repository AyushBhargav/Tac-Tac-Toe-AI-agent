package game

// WhoWon returns the player who won the game.In case game is not finished, VOID will be return.
func WhoWon(state *State) int {
	// Check row-wise
	for i := 0; i < 3; i++ {
		if state.Cells[i][0] == VOID {
			continue
		}
		flagWon := true
		player := state.Cells[i][0]
		for j := 1; j < 3; j++ {
			if player != state.Cells[i][j] {
				flagWon = false
				break
			}
		}
		if flagWon {
			return player
		}
	}
	// Check col-wise
	for i := 0; i < 3; i++ {
		if state.Cells[0][i] == VOID {
			continue
		}
		flagWon := true
		player := state.Cells[0][i]
		for j := 1; j < 3; j++ {
			if player != state.Cells[j][i] {
				flagWon = false
				break
			}
		}
		if flagWon {
			return player
		}
	}
	// Check diagonal
	playerLeft := state.Cells[0][0]
	playerRight := state.Cells[2][0]
	if playerLeft == state.Cells[1][1] && playerLeft == state.Cells[2][2] {
		return playerLeft
	}
	if playerRight == state.Cells[1][1] && playerRight == state.Cells[0][2] {
		return playerRight
	}
	if state.TimesUpdated == 9 {
		return TIE
	}
	return VOID
}
