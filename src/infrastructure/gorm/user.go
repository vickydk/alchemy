package gorm

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"

	domainUser "alchemy/src/domain/user"
	"alchemy/src/shared/constant"
	"alchemy/src/shared/database"
)

type user struct {
	db *database.Database
}

func UserSetup(db *database.Database) domainUser.Repository {
	r := &user{db: db}
	if r.db == nil {
		panic("please provide database client")
	}
	return r
}

func (r *user) FindByEmail(email string) (entity domainUser.Entity, err error) {
	err = r.db.
		Where("email = ?", email).
		First(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("%s", constant.ErrorUserNotFound)
			return
		}
		err = fmt.Errorf("%s", constant.ErrorDatabase)
		return
	}

	return
}
