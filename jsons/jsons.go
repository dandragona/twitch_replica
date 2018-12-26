package jsons

//LittleGameData is the Json format for the data in GameData.
type LittleGameData struct {
	ID          string
	Name        string
	Box_Art_URL string
}

//GameData is the Json format for a game api request.
type GameData struct {
	Data []LittleGameData
}

//TopGames contains list of games and Pagination information.
type TopGames struct {
	Total int64
	Top   []Games
}

//Pagination contains cursor information.
type Games struct {
	Channels int64
	Viewers  int64
	Game     Game
}

type Game struct {
	_ID          int64
	Box          Box
	Giantbomb_id int64
	Logo         Box
	Name         string
	Popularity   int64
}

type Box struct {
	Large    string
	Medium   string
	Small    string
	Template string
}

type Streams struct {
	Data       []Stream
	Pagination Pagination
}

type Stream struct {
	Id            string
	User_ID       string
	User_Name     string
	Game_Id       string
	Community_Ids []string
	Type          string
	Title         string
	Viewer_Count  int64
	Started_At    string
	Language      string
	Thumbnail_URL string
}

type Pagination struct {
	Cursor string
}
