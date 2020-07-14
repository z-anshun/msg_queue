package serve

type Tickets struct {
	MovieId    string `json:"movie_id" binding:"required"`
	MovieName  string `json:"movie_name" binding:"required"`
	WatchTime  string `json:"watch_time" binding:"required"`
	WatchPlace string `json:"watch_place" binding:"required"`
	Num        int    `json:"num" binding:"required"`
}
