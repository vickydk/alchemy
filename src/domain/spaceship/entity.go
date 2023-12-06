package spaceship

import (
	"encoding/json"
	"time"
)

type Entity struct {
	Id        int64           `json:"id"`
	Name      string          `json:"name"`
	Class     string          `json:"class"`
	Crew      int64           `json:"crew"`
	Image     string          `json:"image"`
	Value     float64         `json:"value"`
	Status    string          `json:"status"`
	Armament  json.RawMessage `json:"armament"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *Entity) TableName() string {
	return "spaceship"
}
