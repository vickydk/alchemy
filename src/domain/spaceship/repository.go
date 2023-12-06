package spaceship

type Repository interface {
	Save(entity *Entity) (err error)
	FindByFilter(name, class, status string) (entity []*Entity, err error)
	FindById(id int64) (entity *Entity, err error)
}
