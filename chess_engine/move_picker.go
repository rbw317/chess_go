package chess_engine

type MovePicker interface {
	GetNextMove(game *Game) (Result, *Move)
}
