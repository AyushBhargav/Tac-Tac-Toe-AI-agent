package game

import "fmt"

const (
	_ = iota
	// VOID is empty cell
	VOID
	// PLAYER1 is human Player
	PLAYER1
	// PLAYER2 is AI
	PLAYER2
	// TIE denotes that every field on board has been filled but there is no winner
	TIE
)

// State holds game info.
type State struct {
	Cells        [3][3]int
	Player       int
	TimesUpdated int
}

// New function returns initialized game state.
func (state *State) New() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			state.Cells[i][j] = VOID
		}
	}
	state.Player = PLAYER1
	state.TimesUpdated = 0
}

// Update updates the current game state.
func (state *State) Update(row int, col int) *State {
	if state.Cells[row][col] != VOID {
		return nil
	}
	state.Cells[row][col] = state.Player
	if state.Player == PLAYER1 {
		state.Player = PLAYER2
	} else {
		state.Player = PLAYER1
	}
	state.TimesUpdated++
	return state
}

// GetUnoccupiedCells returns coordinates of unoccupied Cells.
func (state State) GetUnoccupiedCells() [][]int {
	var unoccupiedCells [][]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if state.Cells[i][j] == VOID {
				unoccupiedCells = append(unoccupiedCells, []int{i, j})
			}
		}
	}
	// fmt.Printf("Unoccupied Cells %T", unoccupiedCells)
	return unoccupiedCells
}

func (state *State) Render() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", state.Cells[i][j])
		}
		fmt.Println()
	}
}
