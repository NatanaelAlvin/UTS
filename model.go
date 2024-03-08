package Model

type Accounts struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}
type Games struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	MaxPlayer int    `json:"maxPlayer"`
}
type GetRooms struct {
	Id       int    `json:"id"`
	RoomName string `json:"roomn_name"`
}
type GetRoomsResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    []GetRooms `json:"data"`
}

type Rooms struct {
	Id       int    `json:"id"`
	RoomName string `json:"roomn_name"`
	IdGame   Games  `json:"id_game"`
}
type Participant struct {
	Id        int        `json:"id"`
	IdRoom    Rooms      `json:"room_name"`
	IdAccount []Accounts `json:"id_accounts"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}

type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
