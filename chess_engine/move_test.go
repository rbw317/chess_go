package chess_engine

import "testing"

func TestNewMove(t *testing.T) {
	m := NewMove(A2, A3)
	if m == nil {
		t.Errorf("Error! NewMove function returned nil for move!")
	}

	if m.StartPos != A2 {
		t.Errorf("Error! NewMove function returned move with incorrect Start Position!")
	}

	if m.EndPos != A3 {
		t.Errorf("Error! NewMove function returned move with incorrect End Position!")
	}
}
