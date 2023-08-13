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
