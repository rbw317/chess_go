package chess_engine

type Move struct {
	StartPos  BoardPosition
	EndPos    BoardPosition
	EnPassant bool
}

func NewMove(start BoardPosition, end BoardPosition) *Move {
	m := new(Move)
	m.StartPos = start
	m.EndPos = end
	m.EnPassant = false
	return m
}

func NewEnPassantMove(start BoardPosition, end BoardPosition) *Move {
	m := new(Move)
	m.StartPos = start
	m.EndPos = end
	m.EnPassant = true
	return m
}
