package user

type Repository interface {
	FindByEmail(email string) (entity Entity, err error)
}
