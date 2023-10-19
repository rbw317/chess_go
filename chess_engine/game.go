package chess_engine

type GameStatus uint8

const (
	GAME_STATUS_USER_MOVE        = 0
	GAME_STATUS_ENGINE_MOVE      = 1
	GAME_STATUS_USER_CHECKMATE   = 2
	GAME_STATUS_ENGINE_CHECKMATE = 3
	GAME_STATUS_STALEMATE        = 4
	GAME_STATUS_DRAW             = 5
)

type Game struct {
	Board       *Board
	Moves       []*Move
	UserColor   PieceColor
	EngineColor PieceColor
	Status      GameStatus
	Picker      MovePicker
}

func NewGame(userColor PieceColor) *Game {
	g := &Game{}
	g.Board = NewBoard()
	g.UserColor = userColor
	g.EngineColor = OppositeColor(userColor)
	if userColor == White {
		g.Status = GAME_STATUS_USER_MOVE
	} else {
		g.Status = GAME_STATUS_ENGINE_MOVE
	}
	g.Picker = NewBasicMovePicker(g.EngineColor)
	return g
}

func (game *Game) CheckMate() bool {
	return false
}

func (game *Game) UserMove(move *Move) Result {
	retVal := Result{false, INVALID_MOVE, ""}

	if game.Status == GAME_STATUS_USER_CHECKMATE {
		retVal = Result{false, USER_CHECK_MATE, "User is in checkmate."}
	} else if game.Status == GAME_STATUS_ENGINE_CHECKMATE {
		retVal = Result{false, ENGINE_CHECK_MATE, "Engine is in checkmate."}
	} else if game.Status != GAME_STATUS_USER_MOVE {
		retVal = Result{false, NOT_USER_MOVE, "Not user move"}
	} else if game.Board.Squares[move.StartPos].Occupied &&
		game.Board.Squares[move.StartPos].CurrPiece.GetColor() == game.UserColor {
		userMoves := game.Board.GetMoves(game.UserColor)
		if MovesContainMove(*move, userMoves) {
			move.Castle = game.UserMoveIsCastle(move)
			tempBoard := game.Board.CopyBoard()
			tempRes := tempBoard.MovePiece(move)
			if tempRes.Result {
				if !tempBoard.KingInCheck(game.UserColor) {
					retVal = game.Board.MovePiece(move)
					if retVal.Result {
						game.Moves = append(game.Moves, move)
						game.UpdateGameStatus()
					} else {
						retVal.Result = false
					}

				} else {
					retVal = Result{false, INVALID_MOVE, "Invalid move.  Would put King in check."}
				}
			} else {
				retVal = tempRes
			}
		} else {
			retVal = Result{false, INVALID_MOVE, "Invalid move"}
		}
	}
	return retVal
}

func (game *Game) EngineMove() (Result, *Move) {
	retVal := Result{false, INVALID_MOVE, ""}
	var retMove *Move = nil
	retVal, retMove = game.Picker.GetNextMove(game)
	if retVal.Result {
		retVal = game.Board.MovePiece(retMove)
		if retVal.Result {
			game.UpdateGameStatus()
			retVal = Result{true, NO_ERROR, ""}
		}
	}

	return retVal, retMove
}

func (game *Game) UpdateGameStatus() {
	if game.Board.CheckMate(game.UserColor) {
		game.Status = GAME_STATUS_USER_CHECKMATE
	} else if game.Board.CheckMate(game.EngineColor) {
		game.Status = GAME_STATUS_ENGINE_CHECKMATE
	} else if game.Board.IsStalemate() {
		game.Status = GAME_STATUS_STALEMATE
	} else if game.Board.IsDraw() {
		game.Status = GAME_STATUS_DRAW
	} else if game.Status == GAME_STATUS_USER_MOVE {
		game.Status = GAME_STATUS_ENGINE_MOVE
	} else {
		game.Status = GAME_STATUS_USER_MOVE
	}
	game.Board.KingInCheck(White)
	game.Board.KingInCheck(Black)
}

func (game *Game) UserMoveIsCastle(move *Move) bool {
	retVal := false
	if (move.StartPos == E1 && (move.EndPos == G1 || move.EndPos == C1)) ||
		(move.StartPos == E8 && (move.EndPos == G8 || move.EndPos == C8)) {
		retVal = true
	}
	return retVal
}

func (game *Game) IsOver() bool {
	retVal := true
	if game.Status != GAME_STATUS_ENGINE_CHECKMATE && game.Status != GAME_STATUS_USER_CHECKMATE &&
		game.Status != GAME_STATUS_STALEMATE && game.Status != GAME_STATUS_DRAW {
		retVal = false
	}

	return retVal
}
