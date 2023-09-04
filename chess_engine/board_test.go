package chess_engine

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard()
	if b == nil {
		t.Errorf("Error! NewBoard function returned nil!")
	}
	if (!b.Squares[A1].Occupied) || (b.Squares[A1].CurrPiece.GetType() != ROOK) ||
		(b.Squares[A1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square A1")
	}
	if (!b.Squares[B1].Occupied) || (b.Squares[B1].CurrPiece.GetType() != KNIGHT) ||
		(b.Squares[B1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square B1")
	}
	if (!b.Squares[C1].Occupied) || (b.Squares[C1].CurrPiece.GetType() != BISHOP) ||
		(b.Squares[C1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square C1")
	}
	if (!b.Squares[D1].Occupied) || (b.Squares[D1].CurrPiece.GetType() != QUEEN) ||
		(b.Squares[D1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square D1")
	}
	if (!b.Squares[E1].Occupied) || (b.Squares[E1].CurrPiece.GetType() != KING) ||
		(b.Squares[E1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square E1")
	}
	if (!b.Squares[F1].Occupied) || (b.Squares[F1].CurrPiece.GetType() != BISHOP) ||
		(b.Squares[F1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square F1")
	}
	if (!b.Squares[G1].Occupied) || (b.Squares[G1].CurrPiece.GetType() != KNIGHT) ||
		(b.Squares[G1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square G1")
	}
	if (!b.Squares[H1].Occupied) || (b.Squares[H1].CurrPiece.GetType() != ROOK) ||
		(b.Squares[H1].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square H1")
	}
	if (!b.Squares[A2].Occupied) || (b.Squares[A2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[A2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square A2")
	}
	if (!b.Squares[B2].Occupied) || (b.Squares[B2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[B2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square B2")
	}
	if (!b.Squares[C2].Occupied) || (b.Squares[C2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[C2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square C2")
	}
	if (!b.Squares[D2].Occupied) || (b.Squares[D2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[D2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square D2")
	}
	if (!b.Squares[E2].Occupied) || (b.Squares[E2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[E2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square E2")
	}
	if (!b.Squares[F2].Occupied) || (b.Squares[F2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[F2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square F2")
	}
	if (!b.Squares[G2].Occupied) || (b.Squares[G2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[G2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square G2")
	}
	if (!b.Squares[H2].Occupied) || (b.Squares[H2].CurrPiece.GetType() != PAWN) ||
		(b.Squares[H2].CurrPiece.GetColor() != White) {
		t.Errorf("Error initialing board square H2")
	}
	if b.Squares[A3].Occupied {
		t.Errorf("Error initialing board square A3")
	}
	if b.Squares[B3].Occupied {
		t.Errorf("Error initialing board square B3")
	}
	if b.Squares[C3].Occupied {
		t.Errorf("Error initialing board square C3")
	}
	if b.Squares[D3].Occupied {
		t.Errorf("Error initialing board square D3")
	}
	if b.Squares[E3].Occupied {
		t.Errorf("Error initialing board square E3")
	}
	if b.Squares[F3].Occupied {
		t.Errorf("Error initialing board square F3")
	}
	if b.Squares[G3].Occupied {
		t.Errorf("Error initialing board square G3")
	}
	if b.Squares[H3].Occupied {
		t.Errorf("Error initialing board square H3")
	}
	if b.Squares[A4].Occupied {
		t.Errorf("Error initialing board square A4")
	}
	if b.Squares[B4].Occupied {
		t.Errorf("Error initialing board square B4")
	}
	if b.Squares[C4].Occupied {
		t.Errorf("Error initialing board square C4")
	}
	if b.Squares[D4].Occupied {
		t.Errorf("Error initialing board square D4")
	}
	if b.Squares[E4].Occupied {
		t.Errorf("Error initialing board square E4")
	}
	if b.Squares[F4].Occupied {
		t.Errorf("Error initialing board square F4")
	}
	if b.Squares[G4].Occupied {
		t.Errorf("Error initialing board square G4")
	}
	if b.Squares[H4].Occupied {
		t.Errorf("Error initialing board square H4")
	}
	if b.Squares[A5].Occupied {
		t.Errorf("Error initialing board square A5")
	}
	if b.Squares[B5].Occupied {
		t.Errorf("Error initialing board square B5")
	}
	if b.Squares[C5].Occupied {
		t.Errorf("Error initialing board square C5")
	}
	if b.Squares[D5].Occupied {
		t.Errorf("Error initialing board square D5")
	}
	if b.Squares[E5].Occupied {
		t.Errorf("Error initialing board square E5")
	}
	if b.Squares[F5].Occupied {
		t.Errorf("Error initialing board square F5")
	}
	if b.Squares[G5].Occupied {
		t.Errorf("Error initialing board square G5")
	}
	if b.Squares[H5].Occupied {
		t.Errorf("Error initialing board square H5")
	}
	if b.Squares[A6].Occupied {
		t.Errorf("Error initialing board square A6")
	}
	if b.Squares[B6].Occupied {
		t.Errorf("Error initialing board square B6")
	}
	if b.Squares[C6].Occupied {
		t.Errorf("Error initialing board square C6")
	}
	if b.Squares[D6].Occupied {
		t.Errorf("Error initialing board square D6")
	}
	if b.Squares[E6].Occupied {
		t.Errorf("Error initialing board square E6")
	}
	if b.Squares[F6].Occupied {
		t.Errorf("Error initialing board square F6")
	}
	if b.Squares[G6].Occupied {
		t.Errorf("Error initialing board square G6")
	}
	if b.Squares[H6].Occupied {
		t.Errorf("Error initialing board square H6")
	}
	if (!b.Squares[A7].Occupied) || (b.Squares[A7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[A7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square A7")
	}
	if (!b.Squares[B7].Occupied) || (b.Squares[B7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[B7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square B7")
	}
	if (!b.Squares[C7].Occupied) || (b.Squares[C7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[C7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square C7")
	}
	if (!b.Squares[D7].Occupied) || (b.Squares[D7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[D7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square D7")
	}
	if (!b.Squares[E7].Occupied) || (b.Squares[E7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[E7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square E7")
	}
	if (!b.Squares[F7].Occupied) || (b.Squares[F7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[F7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square F7")
	}
	if (!b.Squares[G7].Occupied) || (b.Squares[G7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[G7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square G7")
	}
	if (!b.Squares[H7].Occupied) || (b.Squares[H7].CurrPiece.GetType() != PAWN) ||
		(b.Squares[H7].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square H7")
	}
	if (!b.Squares[A8].Occupied) || (b.Squares[A8].CurrPiece.GetType() != ROOK) ||
		(b.Squares[A8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square A8")
	}
	if (!b.Squares[B8].Occupied) || (b.Squares[B8].CurrPiece.GetType() != KNIGHT) ||
		(b.Squares[B8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square B8")
	}
	if (!b.Squares[C8].Occupied) || (b.Squares[C8].CurrPiece.GetType() != BISHOP) ||
		(b.Squares[C8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square C8")
	}
	if (!b.Squares[D8].Occupied) || (b.Squares[D8].CurrPiece.GetType() != QUEEN) ||
		(b.Squares[D8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square D8")
	}
	if (!b.Squares[E8].Occupied) || (b.Squares[E8].CurrPiece.GetType() != KING) ||
		(b.Squares[E8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square E8")
	}
	if (!b.Squares[F8].Occupied) || (b.Squares[F8].CurrPiece.GetType() != BISHOP) ||
		(b.Squares[F8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square F8")
	}
	if (!b.Squares[G8].Occupied) || (b.Squares[G8].CurrPiece.GetType() != KNIGHT) ||
		(b.Squares[G8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square G8")
	}
	if (!b.Squares[H8].Occupied) || (b.Squares[H8].CurrPiece.GetType() != ROOK) ||
		(b.Squares[H8].CurrPiece.GetColor() != Black) {
		t.Errorf("Error initialing board square H8")
	}
}

func TestBoardScore(t *testing.T) {
	board := NewBoard()
	score := board.GetScore(White)
	if score != 39 {
		t.Errorf("Error! Board.GetScore function did not return 39 for white on initial board setup!")
	}
	score = board.GetScore(Black)
	if score != 39 {
		t.Errorf("Error! Board.GetScore function did not return 39 for black on initial board setup!")
	}
	board = &Board{}
	board.Squares[D4] = Square{true, NewPawn(D4, White)}
	board.Squares[E5] = Square{true, NewRook(E5, White)}
	board.Squares[E3] = Square{true, NewKnight(E3, White)}
	board.Squares[C5] = Square{true, NewBishop(C5, White)}
	board.Squares[C3] = Square{true, NewQueen(C3, White)}
	board.Squares[E4] = Square{true, NewKing(E4, White)}
	board.Squares[F5] = Square{true, NewPawn(F5, Black)}
	board.Squares[F6] = Square{true, NewPawn(F6, Black)}
	board.Squares[A1] = Square{true, NewRook(A1, Black)}
	board.Squares[A2] = Square{true, NewKnight(A2, Black)}
	board.Squares[A3] = Square{true, NewBishop(A3, Black)}
	board.Squares[A4] = Square{true, NewQueen(A4, Black)}
	board.Squares[A5] = Square{true, NewKing(A5, Black)}

	score = board.GetScore(White)
	if score != 21 {
		t.Errorf("Error! Board.GetScore function did not return 21 for white on random board setup!")
	}
	score = board.GetScore(Black)
	if score != 22 {
		t.Errorf("Error! Board.GetScore function did not return 22 for black on random board setup!")
	}
}

func TestBoardGetKingPos(t *testing.T) {
	board := NewBoard()
	pos := board.GetKingPosition(White)
	if pos != E1 {
		t.Errorf("Error! Board.GetKingPosition function did not return E1 for white on initial board setup!")
	}

	pos = board.GetKingPosition(Black)
	if pos != E8 {
		t.Errorf("Error! Board.GetKingPosition function did not return E8 for black on initial board setup!")
	}

	board = &Board{}
	board.Squares[D4] = Square{true, NewKing(D4, White)}
	pos = board.GetKingPosition(White)
	if pos != D4 {
		t.Errorf("Error! Board.GetKingPosition function did not return D4 for white on random board setup!")
	}

	board = &Board{}
	board.Squares[D5] = Square{true, NewKing(D5, Black)}
	pos = board.GetKingPosition(Black)
	if pos != D5 {
		t.Errorf("Error! Board.GetKingPosition function did not return D5 for black on random board setup!")
	}
}

func TestBoardGetAllMoves(t *testing.T) {
	board := NewBoard()
	moves := board.GetMoves(White)
	if len(moves) != 20 {
		t.Errorf("Error! Board.GetMoves for white function did not return 20 moves on initial board setup!")
	}

	moves = board.GetMoves(Black)
	if len(moves) != 20 {
		t.Errorf("Error! Board.GetMoves for black function did not return 20 moves on initial board setup!")
	}
}

func TestBoardKingInCheck(t *testing.T) {
	board := &Board{}
	board.Squares[D4] = Square{true, NewKing(D4, White)}
	board.Squares[D3] = Square{true, NewQueen(D3, Black)}

	if !board.KingInCheck(White) {
		t.Errorf("Error! Board.KingInCheck for white returned false when king IS in check!")
	}
}

func TestBoardKingNotInCheck(t *testing.T) {
	board := &Board{}
	board.Squares[D4] = Square{true, NewKing(D4, White)}

	if board.KingInCheck(White) {
		t.Errorf("Error! Board.KingInCheck for white returned true when king IS NOT in check!")
	}
}

func TestBoardValidMovePiece(t *testing.T) {
	board := &Board{}
	board.Squares[D4] = Square{true, NewKing(D4, White)}
	move := NewMove(D4, D5)
	res := board.MovePiece(move)
	if !res.Result {
		t.Errorf("Error! Board.MovePiece returned error for valid move!")
	}
}

func TestBoardInvalidMovePiece(t *testing.T) {
	board := &Board{}
	board.Squares[D4] = Square{true, NewKing(D4, White)}
	board.Squares[D5] = Square{true, NewPawn(D5, White)}
	move := NewMove(D4, D5)
	res := board.MovePiece(move)
	if res.Result {
		t.Errorf("Error! Board.MovePiece returned valid for invalid move!")
	}

	move = NewMove(A1, A2)
	res = board.MovePiece(move)
	if res.Result {
		t.Errorf("Error! Board.MovePiece returned valid for invalid move!")
	}
}

func TestBoardCheckMate(t *testing.T) {
	board := &Board{}
	board.Squares[E4] = Square{true, NewKing(E4, White)}
	if board.CheckMate(White) {
		t.Errorf("Error! Board.CheckMate for white returned true when king IS NOT in checkmate!")
	}

	board = &Board{}
	board.Squares[A1] = Square{true, NewKing(A1, White)}
	board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	if board.CheckMate(White) {
		t.Errorf("Error! Board.CheckMate for white returned true when king IS NOT in checkmate!")
	}

	board = &Board{}
	board.Squares[A1] = Square{true, NewKing(A1, White)}
	board.Squares[A3] = Square{true, NewQueen(A3, Black)}
	board.Squares[B3] = Square{true, NewQueen(B3, Black)}
	if !board.CheckMate(White) {
		t.Errorf("Error! Board.CheckMate for white returned false when king IS in checkmate!")
	}
}

func TestBoardKingCastle(t *testing.T) {
	board := NewBoard()
	board.Squares[F1].Occupied = false
	board.Squares[G1].Occupied = false
	move := &Move{E1, G1, false, true, false}
	res := board.MovePiece(move)
	if !res.Result {
		t.Errorf("Error! Board.MovePiece for white castling move returned failure!")
	}
	if !board.Squares[F1].Occupied || board.Squares[F1].CurrPiece.GetType() != ROOK {
		t.Errorf("Error! Board.MovePiece for white castling missing rook at F1 after move!")
	}
	if !board.Squares[G1].Occupied || board.Squares[G1].CurrPiece.GetType() != KING {
		t.Errorf("Error! Board.MovePiece for white castling missing king at G1 after move!")
	}

	board = NewBoard()
	board.Squares[B1].Occupied = false
	board.Squares[C1].Occupied = false
	board.Squares[D1].Occupied = false
	move = &Move{E1, C1, false, true, false}
	res = board.MovePiece(move)
	if !res.Result {
		t.Errorf("Error! Board.MovePiece for white castling move returned failure!")
	}
	if !board.Squares[D1].Occupied || board.Squares[D1].CurrPiece.GetType() != ROOK {
		t.Errorf("Error! Board.MovePiece for white castling missing rook at D1 after move!")
	}
	if !board.Squares[C1].Occupied || board.Squares[C1].CurrPiece.GetType() != KING {
		t.Errorf("Error! Board.MovePiece for white castling missing king at C1 after move!")
	}

	board = NewBoard()
	board.Squares[B8].Occupied = false
	board.Squares[C8].Occupied = false
	board.Squares[D8].Occupied = false
	move = &Move{E8, C8, false, true, false}
	res = board.MovePiece(move)
	if !res.Result {
		t.Errorf("Error! Board.MovePiece for black castling move returned failure!")
	}
	if !board.Squares[D8].Occupied || board.Squares[D8].CurrPiece.GetType() != ROOK {
		t.Errorf("Error! Board.MovePiece for black castling missing rook at D8 after move!")
	}
	if !board.Squares[C8].Occupied || board.Squares[C8].CurrPiece.GetType() != KING {
		t.Errorf("Error! Board.MovePiece for black castling missing king at C8 after move!")
	}

	board = NewBoard()
	board.Squares[F8].Occupied = false
	board.Squares[G8].Occupied = false
	move = &Move{E8, G8, false, true, false}
	res = board.MovePiece(move)
	if !res.Result {
		t.Errorf("Error! Board.MovePiece for black castling move returned failure!")
	}
	if !board.Squares[F8].Occupied || board.Squares[F8].CurrPiece.GetType() != ROOK {
		t.Errorf("Error! Board.MovePiece for black castling missing rook at F8 after move!")
	}
	if !board.Squares[G8].Occupied || board.Squares[G8].CurrPiece.GetType() != KING {
		t.Errorf("Error! Board.MovePiece for black castling missing king at G8 after move!")
	}
}
