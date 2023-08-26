package chess_engine

type UserInterface interface {
	EngineNextMove() *Move
	UserNextMove(move *Move) Result
}
