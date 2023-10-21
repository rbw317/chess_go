package chess_engine

import (
	"fmt"
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
	bestScore := -100

	// Check for engine in checkmate
	if !game.IsOver() {
		moves := game.Board.GetMoves(game.EngineColor)

		for _, m := range moves {
			tempBoard := game.Board.CopyBoard()
			tempBoard.MovePiece(m)
			if !tempBoard.KingInCheck(game.EngineColor) {
				validMoves = append(validMoves, m)
			}
		}

		movesMap := make(map[string][]*Move)
		for _, m := range validMoves {
			tempBoard := game.Board.CopyBoard()
			tempBoard.MovePiece(m)
			// Look for CheckMate first
			if tempBoard.CheckMate(game.UserColor) {
				retMove = m
				break
			}
			// Get the user moves from this position
			userMoves := tempBoard.GetMoves(game.UserColor)
			bestUserScore := -100
			var bestUserBoard *Board = nil
			// Get the move which results in the highest user score
			for _, userMove := range userMoves {
				tempBoard2 := tempBoard.CopyBoard()
				tempBoard2.MovePiece(userMove)
				userScore := tempBoard2.GetScore(game.UserColor)
				engineScore := tempBoard2.GetScore(game.EngineColor)
				currScore := userScore - engineScore
				if currScore > bestUserScore {
					bestUserScore = currScore
					bestUserBoard = tempBoard2
				}
			}
			// Get the diff of the user's best move with this move.  See if this engine/user move combo ends up with the
			// engine having the biggest score advantage. If so then that's the return move.
			engineScore := bestUserBoard.GetScore(game.EngineColor)
			userScore := bestUserBoard.GetScore(game.UserColor)
			currScore := engineScore - userScore
			if currScore >= bestScore {
				bestScore = currScore
				currScoreStr := fmt.Sprintf("%d", currScore)
				if movesMap[currScoreStr] == nil {
					movesMap[currScoreStr] = make([]*Move, 1, 1)
					movesMap[currScoreStr][0] = m
				} else {
					movesMap[currScoreStr] = append(movesMap[currScoreStr], m)
				}
			}

		}

		bestScoreStr := fmt.Sprintf("%d", bestScore)
		bestLen := len(movesMap[bestScoreStr])
		if bestLen > 0 {
			idx := rand.Int() % bestLen
			retMove = movesMap[bestScoreStr][idx]
		}
	}

	if retMove != nil {
		retVal = Result{true, NO_ERROR, ""}
	}

	return retVal, retMove
}
