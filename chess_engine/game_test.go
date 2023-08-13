package chess_engine

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g == nil {
		t.Errorf("Error! NewGame function returned nil for game!")
	}

	if g.board == nil {
		t.Errorf("Error! NewGame function returned nil for board!")
	}

	if len(g.moves) != 0 {
		t.Errorf("Error! NewGame function returned a non-empty moves array!")
	}
}
