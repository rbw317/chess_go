package chess_engine

type Move struct {
	StartPos  BoardPosition
	EndPos    BoardPosition
	EnPassant bool
	Castle    bool
	Promote   bool
}

func NewMove(start BoardPosition, end BoardPosition) *Move {
	m := new(Move)
	m.StartPos = start
	m.EndPos = end
	m.EnPassant = false
	m.Castle = false
	m.Promote = false
	return m
}

func NewEnPassantMove(start BoardPosition, end BoardPosition) *Move {
	m := new(Move)
	m.StartPos = start
	m.EndPos = end
	m.EnPassant = true
	return m
}

func NewCastleMove(start BoardPosition, end BoardPosition) *Move {
	m := new(Move)
	m.StartPos = start
	m.EndPos = end
	m.Castle = true
	return m
}

func NewPromoteMove(start BoardPosition, end BoardPosition) *Move {
	m := new(Move)
	m.StartPos = start
	m.EndPos = end
	m.Promote = true
	return m
}
