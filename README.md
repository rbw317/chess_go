# chess_go
Chess engine written in Go

# Rest Web Service
Provides an API for creating and playing chess games with the Go chess engine.

API JSON Objects

 should just contain an Error object.  Otherwise it will contain the JSON objects expected based on the call that was made.
- Error
  - Fields
    - error_code (int)
    - error_string (string)
- GameInfo
  - Fields
    - id (int)
    - status (string) - Either UserMove, EngineMove, UserCheckMate or EngineCheckMate
- MoveInfo
  - Fields
    - id (int)
    - start_pos (string)
    - end_pos (string)
    - en_passant (bool)
    - castle (bool)
    - promotion (bool)
- MoveRequest
  - Fields
    - user_move (bool)
    - start_pos (string)
    - end_pos (string)
    
API Calls
- /games GET - Gets the list of games.  Returns a JSON array of GameInfo objects
    i.e. [{'id':'1','status':'UserMove'}]

- /games POST - Creates a new chess game.  Returns a GameInfo object for the new game
    i.e. {'id':'1','status':'UserMove'}
- /games/{game_id} GET - Returns the GameInfo object for the specified game id or an Error object.
- /games/{game_id} DELETE - Deletes the specified game
- /games/{game_id}/moves GET - Returns a JSON array of MoveInfo objects which are the moves that have been made 
in the specified game or an Error object.
- /games/{game_id}/moves POST - Makes a new move in the game.  Passed in a MoveRequest object.  
Returns a MoveInfo with an updated GameInfo object or an Error object.
- /games/{game_id}/moves/{move_id} GET - Returns a MoveInfo object for the specified move ID or an Error.
- /games/{game_id}/moves/{move_id} PUT - Updates the move for the specified move ID with the options in the passed in
MoveRequest object.  Returns a MoveInfo object or Error.
- /games/{game_id}/moves/{move_id} DELETE - Deletes the specified move from the game.  Returns Ok or an Error.

