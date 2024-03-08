package Controller

import (
	m "UTS/Model"
	"encoding/json"
	"log"
	"net/http"
)

// GET AllRooms
func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	//log.Print("test log print")
	db := connect()
	defer db.Close()

	query := "SELECT Rooms.id,Rooms.room_name, Games.id,Games.name,Games.max_player FROM Rooms INNER JOIN Games ON Rooms.id_game=Games.id;"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w)
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.RoomName, &room.IdGame.Id, &room.IdGame.Name, &room.IdGame.MaxPlayer); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success Get Data"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)

}

// Insert

func sendErrorResponse(w http.ResponseWriter) {
	var response m.UserResponse
	response.Status = 400
	response.Message = "Fail"
	w.Header().Set("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(response)
}
