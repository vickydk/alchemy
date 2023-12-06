package spaceship

type CreateSpaceshipRequest struct {
	Name     string     `json:"name" validate:"required"`
	Class    string     `json:"class" validate:"required"`
	Crew     int64      `json:"crew"`
	Image    string     `json:"image"`
	Value    float64    `json:"value"`
	Status   string     `json:"status" validate:"required"`
	Armament []Armament `json:"armament" validate:"required"`
}

type Armament struct {
	Title string `json:"title"`
	Qty   string `json:"qty"`
}

type FilterSpaceship struct {
	Name   string `json:"name" query:"name"`
	Class  string `json:"class" query:"name"`
	Status string `json:"status" query:"name"`
}

type FindSpaceshipResponse struct {
	Data []*DataSpaceship `json:"data"`
}

type DataSpaceship struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

type SpaceshipResponse struct {
	Id       int64      `json:"id"`
	Name     string     `json:"name"`
	Class    string     `json:"class"`
	Crew     int64      `json:"crew"`
	Image    string     `json:"image"`
	Value    float64    `json:"value"`
	Status   string     `json:"status"`
	Armament []Armament `json:"armament"`
}
