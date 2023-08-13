package chess_engine

type ChessPiece interface {
	GetMoves(board *Board) []*Move
	GetCurrPosition() BoardPosition
	GetPrevPosition() BoardPosition
	GetColor() PieceColor
	GetType() PieceType
	GetMoveCount() int
	Move(position BoardPosition)
}
