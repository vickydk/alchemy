package gorm

import (
	"errors"
	"fmt"
	"time"

	domainSpaceship "alchemy/src/domain/spaceship"
	"alchemy/src/shared/constant"
	"alchemy/src/shared/database"
	"github.com/jinzhu/gorm"
)

type spaceship struct {
	db *database.Database
}

func SpaceshipSetup(db *database.Database) domainSpaceship.Repository {
	r := &spaceship{db: db}
	if r.db == nil {
		panic("please provide database client")
	}
	return r
}

func (s *spaceship) Save(entity *domainSpaceship.Entity) (err error) {
	entity.UpdatedAt = time.Now()
	err = s.db.Save(entity).Error
	if err != nil {
		return
	}

	return
}

func (s *spaceship) FindByFilter(name, class, status string) (entity []*domainSpaceship.Entity, err error) {
	query := s.db
	if len(name) > 0 {
		query.Where("name = ?", name)
	}
	if len(class) > 0 {
		query.Where("class = ?", class)
	}
	if len(status) > 0 {
		query.Where("status = ?", status)
	}

	err = query.Find(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("%s", constant.ErrorDataNotFound)
			return
		}
		err = fmt.Errorf("%s", constant.ErrorDatabase)
		return
	}

	return
}

func (s *spaceship) FindById(id int64) (entity *domainSpaceship.Entity, err error) {
	err = s.db.
		Where("id = ?", id).
		First(&entity).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = fmt.Errorf("%s", constant.ErrorDataNotFound)
			return nil, err
		}
		err = fmt.Errorf("%s", constant.ErrorDatabase)
		return nil, err
	}

	return
}
