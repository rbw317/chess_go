package chess_web_service

import (
	"chess_go/chess_engine"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"
)

type ChessWebService struct {
	Port          string
	Router        *mux.Router
	GameCount     int
	StaticHandler http.Handler
	Games         map[int]*chess_engine.Game
}

type ApiResult struct {
	Result bool   `json:result`
	Error  string `json:"error"`
}

type GameInfo struct {
	ID     int    `json:"id""`
	Status string `json:"status"`
}

type MoveInfo struct {
	ID        int    `json:"id"`
	Mover     string `json:"mover"` // Either user or engine
	StartPos  string `json:"start_pos"`
	EndPos    string `json:"end_pos"`
	EnPassant bool   `json:"en_passant"`
	Castle    bool   `json:"castle"`
	Promotion bool   `json:"promotion"`
}

type GetGamesResult struct {
	Result ApiResult   `json:"ApiResult"`
	Games  []*GameInfo `json:"Games"`
}

type GetGameInfoResult struct {
	Result ApiResult `json:"ApiResult"`
	Game   *GameInfo `json:"Game"`
}

type DeleteGameResult struct {
	Result ApiResult `json:"ApiResult"`
}

type CreateMoveResult struct {
	Result ApiResult `json:"ApiResult"`
	Move   *MoveInfo `json:"Move"`
}

type GetMoveInfoResult struct {
	Result ApiResult `json:"ApiResult"`
	Move   *MoveInfo `json:"Move"`
}

var ws_instance *ChessWebService = nil

func GetWebServiceInstance(port string) *ChessWebService {
	if ws_instance == nil {
		ws_instance = &ChessWebService{}
		ws_instance.GameCount = 0
		ws_instance.Games = make(map[int]*chess_engine.Game)
		ws_instance.Port = port
		ws_instance.Router = mux.NewRouter()
		// Gets the list of currently running games (returns the list of ids)
		ws_instance.Router.HandleFunc("/api/games", GetGames).Methods("GET")
		// Used to create a new game, returns the id of the new game
		ws_instance.Router.HandleFunc("/api/games", CreateGame).Methods("POST")
		// Get info about a specific game
		ws_instance.Router.HandleFunc("/api/games/{game_id}", GetGameInfo).Methods("GET")
		// Delete the specified game
		ws_instance.Router.HandleFunc("/api/games/{game_id}", DeleteGame).Methods("DELETE")
		// Returns the list of moves in the specified game
		ws_instance.Router.HandleFunc("/api/games/{game_id}/moves", GetMoves).Methods("GET")
		// Create a new move (user or engine)
		ws_instance.Router.HandleFunc("/api/games/{game_id}/moves", CreateMove).Methods("POST")
		// Get the specified move from the specified game
		ws_instance.Router.HandleFunc("/api/games/{game_id}/moves/{move_id}", GetMove).Methods("GET")
		// Update the specified move from the specified game
		ws_instance.Router.HandleFunc("/api/games/{game_id}/moves/{move_id}", UpdateMove).Methods("PUT")
		// Delete the specified move from the specified game
		ws_instance.Router.HandleFunc("/api/games/{game_id}/moves/{move_id}", DeleteMove).Methods("DELETE")
	}

	return ws_instance
}

func (ws *ChessWebService) Run() {
	ct := mime.TypeByExtension(".css")
	fmt.Printf("ct: %s\n", ct)
	ws.StaticHandler = http.FileServer(http.Dir("./chess_web_page"))
	ws.Router.Handle("/", ws.StaticHandler)
	ws.Router.PathPrefix("/chessboardjs/").Handler(http.StripPrefix("/chessboardjs/",
		http.FileServer(http.Dir("./chess_web_page/chessboardjs"))))
	ws.Router.PathPrefix("/chessboardjs/css/").Handler(http.StripPrefix("/chessboardjs/css/",
		http.FileServer(http.Dir("./chess_web_page/chessboardjs/css"))))
	ws.Router.PathPrefix("/img/").Handler(http.StripPrefix("/img/",
		http.FileServer(http.Dir("./chess_web_page/img"))))
	ws.Router.PathPrefix("/img/chesspieces/wikipedia/").Handler(http.StripPrefix("/img/chesspieces/wikipedia/",
		http.FileServer(http.Dir("./chess_web_page/img/chesspieces/wikipedia"))))
	ws.Router.PathPrefix("/chessboardjs/js/").Handler(http.StripPrefix("/chessboardjs/js/",
		http.FileServer(http.Dir("./chess_web_page/chessboardjs/js"))))
	//ws.Router.Handle("/chessboardjs/", ws.StaticHandler)
	//ws.Router.Handle("/chessboardjs/css/", ws.StaticHandler)
	//ws.Router.Handle("/chessboardjs/img/", ws.StaticHandler)
	//ws.Router.Handle("/chessboardjs/js/", ws.StaticHandler)
	log.Fatal(http.ListenAndServe(ws.Port, ws.Router))
}

func GetGames(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.GetGames(w, r)
	}
}
func (ws *ChessWebService) GetGames(w http.ResponseWriter, r *http.Request) {
	var result GetGamesResult
	w.Header().Set("Content-Type", "application/json")
	for id, game := range ws.Games {
		result.Games = append(result.Games, &GameInfo{id, GetGameStatusString(game.Status)})
	}
	result.Result.Result = true
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func CreateGame(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.CreateGame(w, r)
	}
}
func (ws *ChessWebService) CreateGame(w http.ResponseWriter, r *http.Request) {
	var userColor chess_engine.PieceColor = chess_engine.White
	if r.URL.Query().Get("color") != "white" {
		userColor = chess_engine.Black
	}
	game := chess_engine.NewGame(userColor)
	ws.GameCount++
	ws.Games[ws.GameCount] = game
	gameInfo := GameInfo{len(ws.Games), "UserMove"}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(gameInfo)
}

func GetGameInfo(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.GetGameInfo(w, r)
	}
}
func (ws *ChessWebService) GetGameInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res GetGameInfoResult
	vars := mux.Vars(r)
	if len(vars) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else if vars["game_id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else {
		id, err := strconv.Atoi(vars["game_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Result.Result = false
			res.Result.Error = err.Error()
		} else {
			if id > len(ws.Games) {
				w.WriteHeader(http.StatusBadRequest)
				res.Result.Result = false
				res.Result.Error = "Invalid Game ID"
			} else {
				w.WriteHeader(http.StatusOK)
				res.Result.Result = true
				res.Game = &GameInfo{id, GetGameStatusString(ws.Games[id].Status)}
			}

		}
	}
	json.NewEncoder(w).Encode(res)
}

func UpdateGame(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.UpdateGame(w, r)
	}
}
func (ws *ChessWebService) UpdateGame(w http.ResponseWriter, r *http.Request) {

}

func DeleteGame(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.DeleteGame(w, r)
	}
}
func (ws *ChessWebService) DeleteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res GetGameInfoResult
	vars := mux.Vars(r)
	if len(vars) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else if vars["game_id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else {
		id, err := strconv.Atoi(vars["game_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Result.Result = false
			res.Result.Error = err.Error()
		} else {
			if id > len(ws.Games) {
				w.WriteHeader(http.StatusBadRequest)
				res.Result.Result = false
				res.Result.Error = "Invalid Game ID"
			} else {
				w.WriteHeader(http.StatusOK)
				res.Result.Result = true
				delete(ws.Games, id)
			}

		}
	}
	json.NewEncoder(w).Encode(res)
}

func GetMoves(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.GetMoves(w, r)
	}
}
func (ws *ChessWebService) GetMoves(w http.ResponseWriter, r *http.Request) {

}

func CreateMove(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.CreateMove(w, r)
	}
}
func (ws *ChessWebService) CreateMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res CreateMoveResult

	vars := mux.Vars(r)
	if len(vars) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else if vars["game_id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else {
		gameIdStr := vars["game_id"]
		id, err := strconv.Atoi(gameIdStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Result.Result = false
			res.Result.Error = err.Error()
		} else {
			if id > len(ws.Games) {
				w.WriteHeader(http.StatusBadRequest)
				res.Result.Result = false
				res.Result.Error = "Invalid Game ID"
			} else {
				game := ws.Games[id]
				if r.URL.Query().Get("start") == "" ||
					r.URL.Query().Get("end") == "" {
					if game.Status != chess_engine.GAME_STATUS_ENGINE_MOVE {
						res.Result.Result = false
						res.Result.Error = "MIssing start and end parameters"
					} else {
						res = GetEngineMove(game, w, r)
					}
				} else {
					if game.Status != chess_engine.GAME_STATUS_USER_MOVE {
						res.Result.Result = false
						res.Result.Error = "User move request but it's engine's move."
					} else {
						startPos := chess_engine.GetPositionFromString(strings.ToUpper(r.URL.Query().Get("start")))
						endPos := chess_engine.GetPositionFromString(strings.ToUpper(r.URL.Query().Get("end")))
						if startPos == chess_engine.INVALID_BOARD_POSITION ||
							endPos == chess_engine.INVALID_BOARD_POSITION {
							w.WriteHeader(http.StatusBadRequest)
							res.Result.Result = false
							res.Result.Error = "Invalid Board Positions"
						} else {
							move := chess_engine.NewMove(startPos, endPos)
							moveRes := game.UserMove(move)
							if !moveRes.Result {
								w.WriteHeader(http.StatusBadRequest)
								res.Result.Result = false
								res.Result.Error = moveRes.ResStr
							} else {
								w.WriteHeader(http.StatusOK)
								res.Result.Result = true
								res.Move = &MoveInfo{}
								res.Move.Mover = "user"
								res.Move.StartPos = strings.ToLower(chess_engine.GetPositionString(move.StartPos))
								res.Move.EndPos = strings.ToLower(chess_engine.GetPositionString(move.EndPos))
								res.Move.Castle = move.Castle
								res.Move.Promotion = move.Promote
								res.Move.EnPassant = move.EnPassant
								res.Move.ID = len(game.Moves)
							}
						}
					}
				}
			}
		}
	}

	json.NewEncoder(w).Encode(res)
}

func GetEngineMove(game *chess_engine.Game, w http.ResponseWriter, r *http.Request) CreateMoveResult {
	var res CreateMoveResult
	moveRes, move := game.EngineMove()
	if !moveRes.Result {
		res.Result.Result = false
		res.Result.Error = moveRes.ResStr
	} else {
		w.WriteHeader(http.StatusOK)
		res.Result.Result = true
		res.Move = &MoveInfo{}
		res.Move.Mover = "engine"
		res.Move.StartPos = strings.ToLower(chess_engine.GetPositionString(move.StartPos))
		res.Move.EndPos = strings.ToLower(chess_engine.GetPositionString(move.EndPos))
		res.Move.ID = len(game.Moves)
	}
	return res
}

func GetMove(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.GetMove(w, r)
	}
}
func (ws *ChessWebService) GetMove(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res GetMoveInfoResult

	vars := mux.Vars(r)
	if len(vars) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else if vars["game_id"] == "" {
		w.WriteHeader(http.StatusBadRequest)
		res.Result.Result = false
		res.Result.Error = "Missing game id"
	} else {
		id, err := strconv.Atoi(vars["game_id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Result.Result = false
			res.Result.Error = err.Error()
		} else {
			if id > len(ws.Games) {
				w.WriteHeader(http.StatusBadRequest)
				res.Result.Result = false
				res.Result.Error = "Invalid Game ID"
			} else {
				game := ws.Games[id]
				if vars["move_id"] == "" {
					w.WriteHeader(http.StatusBadRequest)
					res.Result.Result = false
					res.Result.Error = "Missing move id"
				} else {
					moveId, err := strconv.Atoi(vars["move_id"])
					if err != nil {
						w.WriteHeader(http.StatusBadRequest)
						res.Result.Result = false
						res.Result.Error = err.Error()
					} else {
						if moveId > len(game.Moves) {
							w.WriteHeader(http.StatusBadRequest)
							res.Result.Result = false
							res.Result.Error = "Invalid Move ID"
						} else {
							move := game.Moves[moveId]
							if move == nil {
								w.WriteHeader(http.StatusInternalServerError)
								res.Result.Result = false
								res.Result.Error = "Specified move ID is nil"
							} else {
								w.WriteHeader(http.StatusOK)
								res.Result.Result = true
								res.Move = &MoveInfo{}
								res.Move.Mover = "user"
								res.Move.StartPos = chess_engine.GetPositionString(move.StartPos)
								res.Move.EndPos = chess_engine.GetPositionString(move.EndPos)
								res.Move.ID = len(game.Moves)
							}
						}
					}
				}
			}
		}
	}

	json.NewEncoder(w).Encode(res)
}

func UpdateMove(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.UpdateMove(w, r)
	}
}
func (ws *ChessWebService) UpdateMove(w http.ResponseWriter, r *http.Request) {

}

func DeleteMove(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.DeleteMove(w, r)
	}
}
func (ws *ChessWebService) DeleteMove(w http.ResponseWriter, r *http.Request) {

}

func GetGameStatusString(status chess_engine.GameStatus) string {
	var resStr string
	switch status {
	case chess_engine.GAME_STATUS_USER_MOVE:
		resStr = "User Move"
		break
	case chess_engine.GAME_STATUS_ENGINE_MOVE:
		resStr = "Engine Move"
		break
	case chess_engine.GAME_STATUS_USER_CHECKMATE:
		resStr = "User Check Mate"
		break
	case chess_engine.GAME_STATUS_ENGINE_CHECKMATE:
		resStr = "Engine Check Mate"
		break
	default:
		resStr = "Unknown state"
		break
	}
	return resStr
}
