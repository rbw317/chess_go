package chess_engine

import "testing"

func TestNewKnight(t *testing.T) {
	knight := NewKnight(B1, Black)

	if knight == nil {
		t.Errorf("Error! NewKnight function returned nil!")
	}

	if knight.Type != KNIGHT {
		t.Errorf("Error! NewKnight function returned wrong Type!")
	}

	if knight.CurrPosition != B1 {
		t.Errorf("Error! NewKnight function returned wrong CurrPosition!")
	}

	if knight.Color != Black {
		t.Errorf("Error! NewKnight function returned wrong Color!")
	}

	if knight.GetType() != KNIGHT {
		t.Errorf("Error! NewKnight function returned wrong Type!")
	}

	if knight.GetCurrPosition() != B1 {
		t.Errorf("Error! NewKnight function returned wrong CurrPosition!")
	}

	if knight.GetColor() != Black {
		t.Errorf("Error! NewKnight function returned wrong Color!")
	}
}

func TestKnightMove(t *testing.T) {
	knight := NewKnight(C1, Black)
	knight.Move(C2)
	if knight.GetCurrPosition() != C2 {
		t.Errorf("Error! Knight.Move function returned wrong CurrPosition!")
	}
	if knight.MoveCount != 1 {
		t.Errorf("Error! Knight.MoveCount wrong after move!")
	}
}

func TestKnightAllMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewKnight(E4, White)}

	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 8 {
		t.Errorf("Error! Knight.GetMoves function did not return 8 moves for knight on open center square!")
	}
	if !MovesContainMove(Move{E4, D6, false}, moves) &&
		!MovesContainMove(Move{E4, C5, false}, moves) &&
		!MovesContainMove(Move{E4, C3, false}, moves) &&
		!MovesContainMove(Move{E4, D2, false}, moves) &&
		!MovesContainMove(Move{E4, F2, false}, moves) &&
		!MovesContainMove(Move{E4, G3, false}, moves) &&
		!MovesContainMove(Move{E4, G5, false}, moves) &&
		!MovesContainMove(Move{E4, F6, false}, moves) {
		t.Errorf("Error! Knight.GetMoves for knight on open center square missing squares!")
	}
}

func TestKnightNoMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewKnight(E4, White)}
	board.Squares[D6] = Square{true, NewPawn(D6, White)}
	board.Squares[C5] = Square{true, NewPawn(C5, White)}
	board.Squares[C3] = Square{true, NewPawn(C3, White)}
	board.Squares[D2] = Square{true, NewPawn(D2, White)}
	board.Squares[F2] = Square{true, NewPawn(F2, White)}
	board.Squares[G3] = Square{true, NewPawn(G3, White)}
	board.Squares[G5] = Square{true, NewPawn(G5, White)}
	board.Squares[F6] = Square{true, NewPawn(F6, White)}
	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Knight.GetMoves function did not return 0 moves for king on totally blocked center square!")
	}
}

func TestKnightAttackMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[E4] = Square{true, NewKnight(E4, Black)}
	board.Squares[D6] = Square{true, NewPawn(D6, White)}
	board.Squares[C5] = Square{true, NewPawn(C5, White)}
	board.Squares[C3] = Square{true, NewPawn(C3, White)}
	board.Squares[D2] = Square{true, NewPawn(D2, White)}
	board.Squares[F2] = Square{true, NewPawn(F2, White)}
	board.Squares[G3] = Square{true, NewPawn(G3, White)}
	board.Squares[G5] = Square{true, NewPawn(G5, White)}
	board.Squares[F6] = Square{true, NewPawn(F6, White)}
	moves := board.Squares[E4].CurrPiece.GetMoves(board)
	if len(moves) != 8 {
		t.Errorf("Error! Knight.GetMoves function did not return 0 moves for king on totally blocked center square!")
	}
	if !MovesContainMove(Move{E4, D6, false}, moves) &&
		!MovesContainMove(Move{E4, C5, false}, moves) &&
		!MovesContainMove(Move{E4, C3, false}, moves) &&
		!MovesContainMove(Move{E4, D2, false}, moves) &&
		!MovesContainMove(Move{E4, F2, false}, moves) &&
		!MovesContainMove(Move{E4, G3, false}, moves) &&
		!MovesContainMove(Move{E4, G5, false}, moves) &&
		!MovesContainMove(Move{E4, F6, false}, moves) {
		t.Errorf("Error! Knight.GetMoves for knight on center square surrounded by opposing pieces missing squares!")
	}
}

func TestKnightEdgeMoves(t *testing.T) {
	board := &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[A4] = Square{true, NewKnight(A4, White)}

	moves := board.Squares[A4].CurrPiece.GetMoves(board)
	if len(moves) != 4 {
		t.Errorf("Error! Knight.GetMoves function did not return 4 moves for knight on open center edge square!")
	}
	if !MovesContainMove(Move{A4, B6, false}, moves) &&
		!MovesContainMove(Move{A4, C5, false}, moves) &&
		!MovesContainMove(Move{A4, C3, false}, moves) &&
		!MovesContainMove(Move{A4, B2, false}, moves) {
		t.Errorf("Error! Knight.GetMoves for knight on open center edge square missing squares!")
	}

	board = &Board{}
	// Test white square Bishop moving from center square to outer corners
	board.Squares[H4] = Square{true, NewKnight(H4, White)}

	moves = board.Squares[H4].CurrPiece.GetMoves(board)
	if len(moves) != 4 {
		t.Errorf("Error! Knight.GetMoves function did not return 4 moves for knight on open center edge square!")
	}
	if !MovesContainMove(Move{H4, G6, false}, moves) &&
		!MovesContainMove(Move{H4, F5, false}, moves) &&
		!MovesContainMove(Move{H4, F3, false}, moves) &&
		!MovesContainMove(Move{H4, G2, false}, moves) {
		t.Errorf("Error! Knight.GetMoves for knight on open center edge square missing squares!")
	}
}
