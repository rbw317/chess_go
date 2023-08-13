package chess_engine

import "testing"

func TestNewRook(t *testing.T) {
	rook := NewRook(B1, Black)

	if rook == nil {
		t.Errorf("Error! NewRook function returned nil!")
	}

	if rook.Type != ROOK {
		t.Errorf("Error! NewRook function returned wrong Type!")
	}

	if rook.CurrPosition != B1 {
		t.Errorf("Error! NewRook function returned wrong CurrPosition!")
	}

	if rook.Color != Black {
		t.Errorf("Error! NewRook function returned wrong Color!")
	}

	if rook.GetType() != ROOK {
		t.Errorf("Error! NewRook function returned wrong Type!")
	}

	if rook.GetCurrPosition() != B1 {
		t.Errorf("Error! NewRook function returned wrong CurrPosition!")
	}

	if rook.GetColor() != Black {
		t.Errorf("Error! NewRook function returned wrong Color!")
	}
}

func TestRookMove(t *testing.T) {
	rook := NewRook(C1, Black)
	rook.Move(C2)
	if rook.GetCurrPosition() != C2 {
		t.Errorf("Error! Rook.Move function returned wrong CurrPosition!")
	}
	if rook.MoveCount != 1 {
		t.Errorf("Error! Rook.MoveCount wrong after move!")
	}
}

func TestRookAllMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewRook(D4, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 14 {
		t.Errorf("Error! Rook.GetMoves function did not return 14 moves for rook on open center square!")
	}

	if !MovesContainMove(Move{D4, D3, false}, moves) &&
		!MovesContainMove(Move{D4, D2, false}, moves) &&
		!MovesContainMove(Move{D4, D1, false}, moves) &&
		!MovesContainMove(Move{D4, D5, false}, moves) &&
		!MovesContainMove(Move{D4, D6, false}, moves) &&
		!MovesContainMove(Move{D4, D7, false}, moves) &&
		!MovesContainMove(Move{D4, D8, false}, moves) &&
		!MovesContainMove(Move{D4, C4, false}, moves) &&
		!MovesContainMove(Move{D4, B4, false}, moves) &&
		!MovesContainMove(Move{D4, A4, false}, moves) &&
		!MovesContainMove(Move{D4, E4, false}, moves) &&
		!MovesContainMove(Move{D4, F4, false}, moves) &&
		!MovesContainMove(Move{D4, G4, false}, moves) &&
		!MovesContainMove(Move{D4, H4, false}, moves) {
		t.Errorf("Error! Rook.GetMoves for white square rook missing moves!")
	}
}

func TestRookNoMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewRook(D4, White)}
	board.Squares[D5] = Square{true, NewPawn(D5, White)}
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	board.Squares[C4] = Square{true, NewPawn(C4, White)}
	board.Squares[E4] = Square{true, NewPawn(E4, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Rook.GetMoves function did not return 0 moves for rook on surrounded center square!")
	}
}

func TestRookAttackMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewRook(D4, Black)}
	board.Squares[D5] = Square{true, NewPawn(D5, White)}
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	board.Squares[C4] = Square{true, NewPawn(C4, White)}
	board.Squares[E4] = Square{true, NewPawn(E4, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 4 {
		t.Errorf("Error! Rook.GetMoves function did not return 4 moves for rook on center square surrounded by opponents!")
	}

	if !MovesContainMove(Move{D4, D5, false}, moves) &&
		!MovesContainMove(Move{D4, C5, false}, moves) &&
		!MovesContainMove(Move{D4, C4, false}, moves) &&
		!MovesContainMove(Move{D4, C3, false}, moves) &&
		!MovesContainMove(Move{D4, D3, false}, moves) &&
		!MovesContainMove(Move{D4, E3, false}, moves) &&
		!MovesContainMove(Move{D4, E4, false}, moves) &&
		!MovesContainMove(Move{D4, E5, false}, moves) {
		t.Errorf("Error! Rook.GetMoves for white square rook surrounded by opposition missing moves!")
	}
}
