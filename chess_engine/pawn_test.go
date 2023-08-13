package chess_engine

import "testing"

func TestNewPawn(t *testing.T) {
	pawn := NewPawn(A2, Black)

	if pawn == nil {
		t.Errorf("Error! NewPawn function returned nil for pawn!")
	}

	if pawn.Type != PAWN {
		t.Errorf("Error! NewPawn function returned wrong Type!")
	}

	if pawn.CurrPosition != A2 {
		t.Errorf("Error! NewPawn function returned pawn with wrong CurrPosition!")
	}

	if pawn.Color != Black {
		t.Errorf("Error! NewPawn function returned pawn with wrong Color!")
	}

	if pawn.MoveCount != 0 {
		t.Errorf("Error! NewPawn function returned pawn with non-zero move count!")
	}

	if pawn.GetType() != PAWN {
		t.Errorf("Error! NewPawn function returned wrong Type!")
	}

	if pawn.GetCurrPosition() != A2 {
		t.Errorf("Error! NewPawn function returned pawn with wrong CurrPosition!")
	}

	if pawn.GetColor() != Black {
		t.Errorf("Error! NewPawn function returned pawn with wrong Color!")
	}

}

func TestPawnMove(t *testing.T) {
	pawn := NewPawn(C1, Black)
	pawn.Move(C2)
	if pawn.GetCurrPosition() != C2 {
		t.Errorf("Error! Pawn.Move function returned wrong CurrPosition!")
	}
}

func TestPawnGenNormalMoves(t *testing.T) {
	board := NewBoard()
	// Test White Pawn basic move (forward one)
	board.Squares[D3].CurrPiece = board.Squares[D2].CurrPiece // Move the pawn from D2 to D3
	board.Squares[D3].Occupied = true
	board.Squares[D3].CurrPiece.Move(D3)
	moves := board.Squares[D3].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function returned wrong number of moves for normal move test!")
	}
	if moves[0].StartPos != D3 {
		t.Errorf("Error! Pawn.GetMoves start position incorrect!")
	}
	if moves[0].EndPos != D4 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect!")
	}
}

func TestPawnBoardEdgeCases(t *testing.T) {
	// White at the end of the board
	board := NewBoard()
	// Test White Pawn basic move (forward one)
	board.Squares[D8].CurrPiece = board.Squares[D2].CurrPiece
	board.Squares[D8].Occupied = true
	board.Squares[D8].CurrPiece.Move(D8)
	moves := board.Squares[D8].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Pawn.GetMoves function returned moves for white Pawn at the end of the board!")
	}

	// Black at the end of the board
	board.InitBoard()
	// Move the black pawn on the D file to the end of the board
	board.Squares[D1].CurrPiece = board.Squares[D7].CurrPiece
	board.Squares[D1].Occupied = true
	board.Squares[D1].CurrPiece.Move(D1)
	moves = board.Squares[D1].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Pawn.GetMoves function returned moves for black Pawn at the end of the board!")
	}

	// White on the left edge
	board.InitBoard()
	// Move the white pawn on the left edge of the board one square (so it's not the first move)
	// and then get the moves
	board.Squares[A3].CurrPiece = board.Squares[A2].CurrPiece
	board.Squares[A3].Occupied = true
	board.Squares[A3].CurrPiece.Move(A3)
	moves = board.Squares[A3].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function returned invalid number of moves for white Pawn on the left edge of the board!")
	}
	if moves[0].StartPos != A3 {
		t.Errorf("Error! Pawn.GetMoves start position incorrect for white pawn on left edge!")
	}
	if moves[0].EndPos != A4 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for white pawn on left edge!")
	}

	// White on the right edge
	board.InitBoard()
	// Move the white pawn on the left edge of the board one square (so it's not the first move)
	// and then get the moves
	board.Squares[H3].CurrPiece = board.Squares[H2].CurrPiece
	board.Squares[H3].Occupied = true
	board.Squares[H3].CurrPiece.Move(H3)
	moves = board.Squares[H3].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function returned invalid number of moves for white Pawn on the right edge of the board!")
	}
	if moves[0].StartPos != H3 {
		t.Errorf("Error! Pawn.GetMoves start position incorrect for white pawn on right edge!")
	}
	if moves[0].EndPos != H4 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for white pawn on right edge!")
	}

	// Black on the left edge
	board.InitBoard()
	// Move the black pawn on the left edge of the board one square (so it's not the first move)
	// and then get the moves
	board.Squares[A6].CurrPiece = board.Squares[A7].CurrPiece
	board.Squares[A6].Occupied = true
	board.Squares[A6].CurrPiece.Move(A6)
	moves = board.Squares[A6].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function returned invalid number of moves for black Pawn on the left edge of the board!")
	}
	if moves[0].StartPos != A6 {
		t.Errorf("Error! Pawn.GetMoves start position incorrect for black pawn on left edge!")
	}
	if moves[0].EndPos != A5 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for black pawn on left edge!")
	}

	// Black on the right edge
	board.InitBoard()
	// Move the black pawn on the left edge of the board one square (so it's not the first move)
	// and then get the moves
	board.Squares[H6].CurrPiece = board.Squares[H7].CurrPiece
	board.Squares[H6].Occupied = true
	board.Squares[H6].CurrPiece.Move(H6)
	moves = board.Squares[H6].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function returned invalid number of moves for black Pawn on the right edge of the board!")
	}
	if moves[0].StartPos != H6 {
		t.Errorf("Error! Pawn.GetMoves start position incorrect for black pawn on right edge!")
	}
	if moves[0].EndPos != H5 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for black pawn on right edge!")
	}
}

func TestPawnFirstMove(t *testing.T) {
	// White
	board := NewBoard()
	// Test White Pawn first move (forward one or forward two)
	moves := board.Squares[D2].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! Pawn.GetMoves function did not return 2 moves for white's first pawn move!")
	}
	if (moves[0].StartPos != D2) || (moves[1].StartPos != D2) {
		t.Errorf("Error! Pawn.GetMoves start positions incorrect for white's first move!")
	}
	if (moves[0].EndPos != D3) && (moves[1].EndPos != D3) {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for white's first move!")
	}
	if (moves[0].EndPos != D4) && (moves[1].EndPos != D4) {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for white's first move!")
	}

	// White - No moves if there is a piece right in front of the pawn
	board = NewBoard()
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	// Test White Pawn first move (forward one or forward two)
	moves = board.Squares[D2].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Pawn.GetMoves function did not return 0 moves for white's first pawn move which is blocked!")
	}

	// Black
	board = NewBoard()
	// Test Black Pawn first move (forward one or forward two)
	moves = board.Squares[D7].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! Pawn.GetMoves function did not return 2 moves for black's first pawn move!")
	}
	if (moves[0].StartPos != D7) || (moves[1].StartPos != D7) {
		t.Errorf("Error! Pawn.GetMoves start positions incorrect for black's first move!")
	}
	if (moves[0].EndPos != D6) && (moves[1].EndPos != D6) {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for black's first move!")
	}
	if (moves[0].EndPos != D5) && (moves[1].EndPos != D5) {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for black's first move!")
	}
	// Black - no moves with piece in front of the pawn
	board = NewBoard()
	board.Squares[D6] = Square{true, NewPawn(D6, Black)}
	// Test Black Pawn first move (forward one or forward two)
	moves = board.Squares[D7].CurrPiece.GetMoves(board)
	if len(moves) != 0 {
		t.Errorf("Error! Pawn.GetMoves function did not return 0 moves for black's first pawn move with piece in front!")
	}
}

func TestPawnAttack(t *testing.T) {
	// White
	board := NewBoard()
	// Test White Pawn first move (forward one or forward two)
	board.Squares[D3] = Square{true, NewPawn(D3, White)}
	board.Squares[C3] = Square{true, NewPawn(C3, Black)}

	moves := board.Squares[D2].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function did not return 1 move for white attack test!")
	}
	if moves[0].StartPos != D2 {
		t.Errorf("Error! Pawn.GetMoves start positions incorrect for white's first move!")
	}
	if moves[0].EndPos != C3 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for white's first move!")
	}

	// Black
	board = NewBoard()
	// Test Black Pawn first move (forward one or forward two)
	board.Squares[D6] = Square{true, NewPawn(D6, Black)}
	board.Squares[C6] = Square{true, NewPawn(C6, White)}

	moves = board.Squares[D7].CurrPiece.GetMoves(board)
	if len(moves) != 1 {
		t.Errorf("Error! Pawn.GetMoves function did not return 1 move for black attack test")
	}
	if moves[0].StartPos != D7 {
		t.Errorf("Error! Pawn.GetMoves start positions incorrect for black's attack move!")
	}
	if moves[0].EndPos != C6 {
		t.Errorf("Error! Pawn.GetMoves end position incorrect for black's attack move!")
	}
}

func TestPawnWhiteEnPassantAttack(t *testing.T) {
	// White
	board := NewBoard()
	// Test White Pawn first move (forward one or forward two)
	whitePawn := NewPawn(D5, White)
	blackPawn := NewPawn(C5, Black)
	blackPawn.PrevPosition = C7
	blackPawn.MoveCount = 1
	blackPawn2 := NewPawn(E5, Black)
	blackPawn2.PrevPosition = E7
	blackPawn2.MoveCount = 1
	board.Squares[D5] = Square{true, whitePawn}
	board.Squares[D6] = Square{true, whitePawn}
	board.Squares[C5] = Square{true, blackPawn}
	board.Squares[E5] = Square{true, blackPawn2}
	moves := board.Squares[D5].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! Pawn.GetMoves function did not return 1 move for white attack test!")
	} else {
		if moves[0].StartPos != D5 {
			t.Errorf("Error! Pawn.GetMoves start positions incorrect for white's enPassant!")
		}
		if moves[0].EndPos != C6 {
			t.Errorf("Error! Pawn.GetMoves end position incorrect for white's enPassant!")
		}
		if !moves[0].EnPassant {
			t.Errorf("Error! Pawn.GetMoves first move is not en passant!")
		}
		if moves[1].StartPos != D5 {
			t.Errorf("Error! Pawn.GetMoves start positions incorrect for white's enPassant!")
		}
		if moves[1].EndPos != E6 {
			t.Errorf("Error! Pawn.GetMoves end position incorrect for white's enPassant!")
		}
		if !moves[1].EnPassant {
			t.Errorf("Error! Pawn.GetMoves first move is not enPassant!")
		}
	}
}

func TestPawnBlackEnPassantAttack(t *testing.T) {
	// White
	board := NewBoard()
	// Test White Pawn first move (forward one or forward two)
	blackPawn := NewPawn(D4, Black)
	whitePawn := NewPawn(C4, White)
	whitePawn.PrevPosition = C2
	whitePawn.MoveCount = 1
	whitePawn2 := NewPawn(E4, White)
	whitePawn2.PrevPosition = E2
	whitePawn2.MoveCount = 1
	board.Squares[D4] = Square{true, blackPawn}
	board.Squares[D3] = Square{true, blackPawn}
	board.Squares[C4] = Square{true, whitePawn}
	board.Squares[E4] = Square{true, whitePawn2}
	moves := board.Squares[D4].CurrPiece.GetMoves(board)
	if len(moves) != 2 {
		t.Errorf("Error! Pawn.GetMoves function did not return 1 move for black enpassant attack test!")
	} else {
		if moves[0].StartPos != D4 {
			t.Errorf("Error! Pawn.GetMoves start positions incorrect for black's enPassant!")
		}
		if moves[0].EndPos != C3 {
			t.Errorf("Error! Pawn.GetMoves end position incorrect for black's enPassant!")
		}
		if !moves[0].EnPassant {
			t.Errorf("Error! Pawn.GetMoves first move is not en passant!")
		}
		if moves[1].StartPos != D4 {
			t.Errorf("Error! Pawn.GetMoves start positions incorrect for black's enPassant!")
		}
		if moves[1].EndPos != E3 {
			t.Errorf("Error! Pawn.GetMoves end position incorrect for black's enPassant!")
		}
		if !moves[1].EnPassant {
			t.Errorf("Error! Pawn.GetMoves first move is not enPassant!")
		}
	}
}
