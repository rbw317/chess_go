package chess_engine

import (
	"math/rand"
)

type BasicMovePicker struct {
	EngineColor PieceColor
}

func NewBasicMovePicker(engineColor PieceColor) *BasicMovePicker {
	movePicker := &BasicMovePicker{engineColor}
	return movePicker
}

func (picker *BasicMovePicker) GetNextMove(game *Game) (Result, *Move) {
	retVal := Result{false, INVALID_MOVE, ""}
	validMoves := []*Move{}
	var retMove *Move = nil

	// Check for engine in checkmate
	if game.Status != GAME_STATUS_ENGINE_CHECKMATE && game.Status != GAME_STATUS_USER_CHECKMATE {
		moves := game.Board.GetMoves(game.EngineColor)
		attackMoves := []*Move{}
		noAttackMoves := []*Move{}
		for _, m := range moves {
			tempBoard := game.Board.CopyBoard()
			tempBoard.MovePiece(m)
			if !tempBoard.KingInCheck(game.EngineColor) {
				validMoves = append(validMoves, m)
			}
		}

		// Look for CheckMate first
		for _, m := range validMoves {
			tempBoard := game.Board.CopyBoard()
			tempBoard.MovePiece(m)
			if tempBoard.CheckMate(game.UserColor) {
				retMove = m
				break
			}
			userMoves := tempBoard.GetMoves(game.UserColor)
			if MovesContainMove(*m, userMoves) {
				attackMoves = append(attackMoves, m)
			} else {
				noAttackMoves = append(noAttackMoves, m)
			}
		}

		// Pick the move with the highest score from the no attack moves
		noAttackBestScore := 100
		noAttackMovesMap := make(map[int][]*Move)
		if retMove == nil {
			for _, m := range noAttackMoves {
				tempBoard := game.Board.CopyBoard()
				tempBoard.MovePiece(m)
				currScore := tempBoard.GetScore(game.UserColor)
				if currScore <= noAttackBestScore {
					retMove = m
					if noAttackMovesMap[currScore] == nil {
						noAttackMovesMap[currScore] = make([]*Move, 1, 1)
						noAttackMovesMap[currScore][0] = m
					} else {
						noAttackMovesMap[currScore] = append(noAttackMovesMap[currScore], m)
					}
					noAttackBestScore = currScore
				}
			}

			bestLen := len(noAttackMovesMap[noAttackBestScore])
			if bestLen > 0 {
				idx := rand.Int() % bestLen
				retMove = noAttackMovesMap[noAttackBestScore][idx]
			}
		}

		// If still no move pick the move with the best score from the attack moves
		attackBestScore := 100
		attackMovesMap := make(map[int][]*Move)
		if retMove == nil {
			for _, m := range attackMoves {
				tempBoard := game.Board.CopyBoard()
				tempBoard.MovePiece(m)
				currScore := tempBoard.GetScore(game.UserColor)
				if currScore <= attackBestScore {
					retMove = m
					if attackMovesMap[currScore] == nil {
						attackMovesMap[currScore] = make([]*Move, 1, 1)
						attackMovesMap[currScore][0] = m
					} else {
						attackMovesMap[currScore] = append(attackMovesMap[currScore], m)
					}
					attackBestScore = currScore
				}
			}

			bestLen := len(attackMovesMap[attackBestScore])
			if bestLen > 0 {
				idx := rand.Int() % bestLen
				retMove = attackMovesMap[attackBestScore][idx]
			}
		}
	}

	if retMove != nil {
		retVal = Result{true, NO_ERROR, ""}
	}
	return retVal, retMove
}
