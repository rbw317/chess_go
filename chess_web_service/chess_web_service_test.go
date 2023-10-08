package chess_web_service

import (
	"chess_go/chess_engine"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetInstance(t *testing.T) {
	ws := GetWebServiceInstance(":8080")

	if ws == nil {
		t.Errorf("Error! GetWebServiceInstance function returned nil for ws instance!")
	}

	ws2 := GetWebServiceInstance(":8080")

	if ws != ws2 {
		t.Errorf("Error! GetWebServiceInstance function returned a different pointer for second call!")
	}
}

func TestGetGames(t *testing.T) {
	var result GetGamesResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	ws.Games[2] = chess_engine.NewGame(chess_engine.Black)
	req := httptest.NewRequest(http.MethodGet, "/api/games", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("GetGames return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	json.Unmarshal(data, &result)
	if !result.Result.Result {
		t.Fatalf("GetGames returned failue!")
	}

	if len(result.Games) != 2 {
		t.Fatalf("Expected GetGames to return 2 games, actually returned %d!", len(result.Games))
	}

	if result.Games[0].ID != 1 {
		t.Fatalf("Expected GetGames first id to be 0, actually returned %d!", result.Games[0].ID)
	}

	if result.Games[1].ID != 2 {
		t.Fatalf("Expected GetGames first id to be 1, actually returned %d!", result.Games[1].ID)
	}

}

func TestCreateGame(t *testing.T) {
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	req := httptest.NewRequest(http.MethodPost, "/api/games?color=black", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	gameLen := len(ws.Games)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("CreateGame return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	var m map[string]string
	json.Unmarshal(data, &m)
	if m["error"] != "" {
		t.Errorf("Received error %s from Game POST call", m["error"])
	}

	if gameLen != 1 {
		t.Fatalf("After CreateGame expected 1 game in the web service list but have %d", len(ws.Games))
	}

	if ws.Games[1].UserColor != chess_engine.Black {
		t.Fatalf("After CreateGame expected user color to be white but is black.")
	}
}

func TestGetGameInfo(t *testing.T) {
	var result GetGameInfoResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	ws.Games[2] = chess_engine.NewGame(chess_engine.Black)
	req := httptest.NewRequest(http.MethodGet, "/api/games/2", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("GetGameInfo return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	json.Unmarshal(data, &result)
	if !result.Result.Result {
		t.Fatalf("GetGameInfo returned failue!")
	}

	if result.Game == nil {
		t.Fatalf("GetGameInfo did not return game info")
	}

	if result.Game.ID != 2 {
		t.Fatalf("Expected GetGameInfo id to be 1, actually returned %d!", result.Game.ID)
	}

	if result.Game.Status != "Engine Move" {
		t.Fatalf("Expected GetGameInfo status to be Engine Move, actually returned %s!", result.Game.Status)
	}
}

func TestDeleteGame(t *testing.T) {
	var result GetGameInfoResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	ws.Games[2] = chess_engine.NewGame(chess_engine.Black)
	req := httptest.NewRequest(http.MethodDelete, "/api/games/2", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("DeleteGame return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	json.Unmarshal(data, &result)
	if !result.Result.Result {
		t.Fatalf("DeleteGame returned failue!")
	}

	if len(ws.Games) != 1 {
		t.Fatalf("DeleteGame did not delete the specified game from the web service.")
	}
}

func TestUserMove(t *testing.T) {
	var result CreateMoveResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	req := httptest.NewRequest(http.MethodPost, "/api/games/1/moves?start=a2&end=a3", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("CreateMove return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Error parsing response from CreateMove call: %s!", err.Error())
	}

	if !result.Result.Result {
		t.Fatalf("CreateMove returned failue!")
	}

	if chess_engine.GetPositionFromString(result.Move.StartPos) != chess_engine.A2 ||
		chess_engine.GetPositionFromString(result.Move.EndPos) != chess_engine.A3 {
		t.Fatalf("CreateMove returned wrong positions: %s, %s", result.Move.StartPos, result.Move.EndPos)
	}

	if result.Move.Castle {
		t.Fatalf("CreateMove indicates Castle for basic pawn move.")
	}

	if result.Move.EnPassant {
		t.Fatalf("CreateMove indicates EnPassant for basic pawn move.")
	}

	if result.Move.Promotion {
		t.Fatalf("CreateMove indicates Promotion for basic pawn move.")
	}
}

func TestInvalidUserMove(t *testing.T) {
	var result CreateMoveResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	req := httptest.NewRequest(http.MethodPost, "/api/games/1/moves?start=a2&end=d3", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusBadRequest != res.StatusCode {
		t.Fatalf("CreateMove return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Error parsing response from CreateMove call: %s!", err.Error())
	}

	if result.Result.Result {
		t.Fatalf("CreateMove returned success for an invalid user move!")
	}
}

func TestEngineMove(t *testing.T) {
	var result CreateMoveResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	req := httptest.NewRequest(http.MethodPost, "/api/games/1/moves?start=a2&end=a3", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("CreateMove return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		t.Fatalf("Error parsing response from CreateMove call: %s!", err.Error())
	}

	if !result.Result.Result {
		t.Fatalf("CreateMove returned failue!")
	}

	if result.Move.Mover != "user" {
		t.Fatalf("CreateMove returned %s for the mover when user was expected.", result.Move.Mover)
	}

	if chess_engine.GetPositionFromString(result.Move.StartPos) != chess_engine.A2 ||
		chess_engine.GetPositionFromString(result.Move.EndPos) != chess_engine.A3 {
		t.Fatalf("CreateMove returned wrong positions: %s, %s", result.Move.StartPos, result.Move.EndPos)
	}

	if result.Move.Castle {
		t.Fatalf("CreateMove indicates Castle for basic pawn move.")
	}

	if result.Move.EnPassant {
		t.Fatalf("CreateMove indicates EnPassant for basic pawn move.")
	}

	if result.Move.Promotion {
		t.Fatalf("CreateMove indicates Promotion for basic pawn move.")
	}
}

func TestGetMoveInfo(t *testing.T) {
	var result GetMoveInfoResult
	ws := GetWebServiceInstance(":8080") // Have to create the web service
	ws.Games = make(map[int]*chess_engine.Game)
	ws.Games[1] = chess_engine.NewGame(chess_engine.White)
	// Make a valid move
	req := httptest.NewRequest(http.MethodPost, "/api/games/1/moves?start=a2&end=a3", nil)
	w := httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)

	req = httptest.NewRequest(http.MethodGet, "/api/games/1/moves/0", nil)
	w = httptest.NewRecorder()
	ws.Router.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if http.StatusOK != res.StatusCode {
		t.Fatalf("GetMoveInfo return status expected a StatusOK, instead got: %d", res.StatusCode)
	}

	json.Unmarshal(data, &result)
	if !result.Result.Result {
		t.Fatalf("GetMoveInfo returned failue!")
	}

	if result.Move == nil {
		t.Fatalf("GetMoveInfo did not return game info")
	}

	if result.Move.ID != 1 {
		t.Fatalf("Expected GetMoveInfo id to be 1, actually returned %d!", result.Move.ID)
	}

	if result.Move.StartPos != "A2" {
		t.Fatalf("Expected GetMoveInfo.StartPos to be A2, actually returned %s!", result.Move.StartPos)
	}

	if result.Move.EndPos != "A3" {
		t.Fatalf("Expected GetMoveInfo.EndPos to be A3, actually returned %s!", result.Move.EndPos)
	}
}
