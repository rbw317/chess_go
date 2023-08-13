package chess_engine

type PieceType uint8

const (
	NoPiece PieceType = iota
	PAWN
	ROOK
	KNIGHT
	BISHOP
	QUEEN
	KING
)

type PieceColor uint8

const (
	White = 0
	Black = 1
)

type BoardPosition uint8

const (
	A1 BoardPosition = iota
	A2
	A3
	A4
	A5
	A6
	A7
	A8
	B1
	B2
	B3
	B4
	B5
	B6
	B7
	B8
	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	D1
	D2
	D3
	D4
	D5
	D6
	D7
	D8
	E1
	E2
	E3
	E4
	E5
	E6
	E7
	E8
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	G1
	G2
	G3
	G4
	G5
	G6
	G7
	G8
	H1
	H2
	H3
	H4
	H5
	H6
	H7
	H8
	INVALID_BOARD_POSITION
)

type Square struct {
	Occupied  bool
	CurrPiece ChessPiece
}

type Board struct {
	Squares [64]Square
}

func NewBoard() *Board {
	b := Board{}
	b.InitBoard()
	return &b
}

func (board *Board) InitBoard() {
	board.Squares[A1] = Square{Occupied: true, CurrPiece: NewRook(A1, White)}
	board.Squares[B1] = Square{Occupied: true, CurrPiece: NewKnight(B1, White)}
	board.Squares[C1] = Square{Occupied: true, CurrPiece: NewBishop(C1, White)}
	board.Squares[D1] = Square{Occupied: true, CurrPiece: NewQueen(D1, White)}
	board.Squares[E1] = Square{Occupied: true, CurrPiece: NewKing(E1, White)}
	board.Squares[F1] = Square{Occupied: true, CurrPiece: NewBishop(F1, White)}
	board.Squares[G1] = Square{Occupied: true, CurrPiece: NewKnight(G1, White)}
	board.Squares[H1] = Square{Occupied: true, CurrPiece: NewRook(H1, White)}
	board.Squares[A2] = Square{Occupied: true, CurrPiece: NewPawn(A2, White)}
	board.Squares[B2] = Square{Occupied: true, CurrPiece: NewPawn(B2, White)}
	board.Squares[C2] = Square{Occupied: true, CurrPiece: NewPawn(C2, White)}
	board.Squares[D2] = Square{Occupied: true, CurrPiece: NewPawn(D2, White)}
	board.Squares[E2] = Square{Occupied: true, CurrPiece: NewPawn(E2, White)}
	board.Squares[F2] = Square{Occupied: true, CurrPiece: NewPawn(F2, White)}
	board.Squares[G2] = Square{Occupied: true, CurrPiece: NewPawn(G2, White)}
	board.Squares[H2] = Square{Occupied: true, CurrPiece: NewPawn(H2, White)}
	board.Squares[A8] = Square{Occupied: true, CurrPiece: NewRook(A8, Black)}
	board.Squares[B8] = Square{Occupied: true, CurrPiece: NewKnight(B8, Black)}
	board.Squares[C8] = Square{Occupied: true, CurrPiece: NewBishop(C8, Black)}
	board.Squares[D8] = Square{Occupied: true, CurrPiece: NewQueen(D8, Black)}
	board.Squares[E8] = Square{Occupied: true, CurrPiece: NewKing(E8, Black)}
	board.Squares[F8] = Square{Occupied: true, CurrPiece: NewBishop(F8, Black)}
	board.Squares[G8] = Square{Occupied: true, CurrPiece: NewKnight(G8, Black)}
	board.Squares[H8] = Square{Occupied: true, CurrPiece: NewRook(H8, Black)}
	board.Squares[A7] = Square{Occupied: true, CurrPiece: NewPawn(A7, Black)}
	board.Squares[B7] = Square{Occupied: true, CurrPiece: NewPawn(B7, Black)}
	board.Squares[C7] = Square{Occupied: true, CurrPiece: NewPawn(C7, Black)}
	board.Squares[D7] = Square{Occupied: true, CurrPiece: NewPawn(D7, Black)}
	board.Squares[E7] = Square{Occupied: true, CurrPiece: NewPawn(E7, Black)}
	board.Squares[F7] = Square{Occupied: true, CurrPiece: NewPawn(F7, Black)}
	board.Squares[G7] = Square{Occupied: true, CurrPiece: NewPawn(G7, Black)}
	board.Squares[H7] = Square{Occupied: true, CurrPiece: NewPawn(H7, Black)}
}
