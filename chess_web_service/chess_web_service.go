package chess_web_service

import (
	"chess_go/chess_engine"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type ChessWebService struct {
	Port   string
	Router *mux.Router
	Games  []*chess_engine.Game
}

type ApiResult struct {
	Result bool   `json:result`
	Error  string `json:"error"`
}

type GameInfo struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type MoveInfo struct {
	ID        int    `json:"id"`
	StartPos  string `json:"start_pos"`
	EndPos    string `json:"start_pos"`
	EnPassant bool   `json: "en_passant"`
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

var ws_instance *ChessWebService = nil

func GetWebServiceInstance(port string) *ChessWebService {
	if ws_instance == nil {
		ws_instance = &ChessWebService{}
		ws_instance.Port = port
		ws_instance.Router = mux.NewRouter()
		// Gets the list of currently running games (returns the list of ids)
		ws_instance.Router.HandleFunc("/games", GetGames).Methods("GET")
		// Used to create a new game, returns the id of the new game
		ws_instance.Router.HandleFunc("/games", CreateGame).Methods("POST")
		// Get info about a specific game
		ws_instance.Router.HandleFunc("/games/{game_id}", GetGameInfo).Methods("GET")
		// Delete the specified game
		ws_instance.Router.HandleFunc("/games/{game_id}", DeleteGame).Methods("DELETE")
		// Returns the list of moves in the specified game
		ws_instance.Router.HandleFunc("/games/{game_id}/moves", GetMoves).Methods("GET")
		// Create a new move (user or engine)
		ws_instance.Router.HandleFunc("/games/{game_id}/moves", CreateMove).Methods("POST")
		// Get the specified move from the specified game
		ws_instance.Router.HandleFunc("/games/{game_id}/moves/{move_id}", GetMove).Methods("GET")
		// Update the specified move from the specified game
		ws_instance.Router.HandleFunc("/games/{game_id}/moves/{move_id}", UpdateMove).Methods("PUT")
		// Delete the specified move from the specified game
		ws_instance.Router.HandleFunc("/games/{game_id}/moves/{move_id}", DeleteMove).Methods("DELETE")
	}

	return ws_instance
}

func (ws *ChessWebService) Run() {
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
	for id, _ := range ws.Games {
		result.Games = append(result.Games, &GameInfo{id, "game"})
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
	ws.Games = append(ws.Games, game)
	gameInfo := GameInfo{len(ws.Games), "UserMove"}
	w.Header().Set("Content-Type", "application/json")
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

}

func GetMove(w http.ResponseWriter, r *http.Request) {
	if ws_instance != nil {
		ws_instance.GetMove(w, r)
	}
}
func (ws *ChessWebService) GetMove(w http.ResponseWriter, r *http.Request) {

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
