package chess_engine

import "testing"

func TestNewQueen(t *testing.T) {
	queen := NewQueen(B1, Black)

	if queen == nil {
		t.Errorf("Error! NewQueen function returned nil!")
	}

	if queen.Type != QUEEN {
		t.Errorf("Error! NewQueen function returned wrong Type!")
	}

	if queen.CurrPosition != B1 {
		t.Errorf("Error! NewQueen function returned wrong CurrPosition!")
	}

	if queen.Color != Black {
		t.Errorf("Error! NewQueen function returned wrong Color!")
	}

	if queen.GetType() != QUEEN {
		t.Errorf("Error! NewQueen function returned wrong Type!")
	}

	if queen.GetCurrPosition() != B1 {
		t.Errorf("Error! NewQueen function returned wrong CurrPosition!")
	}

	if queen.GetColor() != Black {
		t.Errorf("Error! NewQueen function returned wrong Color!")
	}
}

func TestQueenMove(t *testing.T) {
	queen := NewQueen(C1, Black)
	queen.Move(C2)
	if queen.GetCurrPosition() != C2 {
		t.Errorf("Error! Queen.Move function returned wrong CurrPosition!")
	}
	if queen.MoveCount != 1 {
		t.Errorf("Error! Queen.MoveCount wrong after move!")
	}
}

func TestQueenAllMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewQueen(D4, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 27 {
		t.Errorf("Error! Queen.GetMoves function did not return 27 moves for queen on open center square!")
	}

	if !MovesContainMove(Move{D4, C3, false}, moves) &&
		!MovesContainMove(Move{D4, B2, false}, moves) &&
		!MovesContainMove(Move{D4, A1, false}, moves) &&
		!MovesContainMove(Move{D4, E5, false}, moves) &&
		!MovesContainMove(Move{D4, F6, false}, moves) &&
		!MovesContainMove(Move{D4, G7, false}, moves) &&
		!MovesContainMove(Move{D4, H8, false}, moves) &&
		!MovesContainMove(Move{D4, E3, false}, moves) &&
		!MovesContainMove(Move{D4, F2, false}, moves) &&
		!MovesContainMove(Move{D4, G1, false}, moves) &&
		!MovesContainMove(Move{D4, C5, false}, moves) &&
		!MovesContainMove(Move{D4, B6, false}, moves) &&
		!MovesContainMove(Move{D4, D3, false}, moves) &&
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
		!MovesContainMove(Move{D4, H4, false}, moves) &&
		!MovesContainMove(Move{D4, A7, false}, moves) {
		t.Errorf("Error! Queen.GetMoves for white square queen missing moves!")
	}
}

func TestQueenNoMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewQueen(D4, White)}
	board.Squares[E5] = Square{true, NewPawn(E5, White)}
	board.Squares[E3] = Square{true, NewPawn(E3, White)}
	board.Squares[C5] = Square{true, NewPawn(C5, White)}
	board.Squares[C3] = Square{true, NewPawn(C3, White)}
	board.Squares[D5] = Square{true, NewPawn(D5, White)}
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	board.Squares[C4] = Square{true, NewPawn(C4, White)}
	board.Squares[E4] = Square{true, NewPawn(E4, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Queen.GetMoves function did not return 0 moves for queen on surrounded center square!")
	}
}

func TestQueenAttackMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[D4] = Square{true, NewQueen(D4, Black)}
	board.Squares[E5] = Square{true, NewPawn(E5, White)}
	board.Squares[E3] = Square{true, NewPawn(E3, White)}
	board.Squares[C5] = Square{true, NewPawn(C5, White)}
	board.Squares[C3] = Square{true, NewPawn(C3, White)}
	board.Squares[D5] = Square{true, NewPawn(D5, White)}
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	board.Squares[C4] = Square{true, NewPawn(C4, White)}
	board.Squares[E4] = Square{true, NewPawn(E4, White)}

	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 8 {
		t.Errorf("Error! Queen.GetMoves function did not return 8 moves for queen on open center square!")
	}

	if !MovesContainMove(Move{D4, D5, false}, moves) &&
		!MovesContainMove(Move{D4, C5, false}, moves) &&
		!MovesContainMove(Move{D4, C4, false}, moves) &&
		!MovesContainMove(Move{D4, C3, false}, moves) &&
		!MovesContainMove(Move{D4, D3, false}, moves) &&
		!MovesContainMove(Move{D4, E3, false}, moves) &&
		!MovesContainMove(Move{D4, E4, false}, moves) &&
		!MovesContainMove(Move{D4, E5, false}, moves) {
		t.Errorf("Error! Queen.GetMoves for white square queen missing moves!")
	}
}
