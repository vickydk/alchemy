package spaceship

import (
	"encoding/json"

	domainSpaceship "alchemy/src/domain/spaceship"
	ctxSess "alchemy/src/shared/utils/context"
)

type service struct {
	spaceshipRepo domainSpaceship.Repository
}

func NewService(spaceshipRepo domainSpaceship.Repository) Service {
	s := &service{
		spaceshipRepo: spaceshipRepo,
	}
	if s.spaceshipRepo == nil {
		panic("please provide spaceship repo")
	}
	return s
}

func (s *service) CreateSpaceship(ctxSess *ctxSess.Context, req *CreateSpaceshipRequest) (err error) {
	entity := &domainSpaceship.Entity{
		Name:   req.Name,
		Class:  req.Class,
		Crew:   req.Crew,
		Image:  req.Image,
		Value:  req.Value,
		Status: req.Status,
	}
	if len(req.Armament) > 0 {
		entity.Armament, _ = json.Marshal(req.Armament)
	}
	err = s.spaceshipRepo.Save(entity)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}
	return
}

func (s *service) UpdateSpaceship(ctxSess *ctxSess.Context, id int64, req *CreateSpaceshipRequest) (err error) {
	entity, err := s.spaceshipRepo.FindById(id)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}
	entity.Name = req.Name
	entity.Class = req.Class
	entity.Crew = req.Crew
	entity.Image = req.Image
	entity.Value = req.Value
	entity.Status = req.Status
	if len(req.Armament) > 0 {
		entity.Armament, _ = json.Marshal(req.Armament)
	}
	err = s.spaceshipRepo.Save(entity)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return
	}
	return
}

func (s *service) FindSpaceship(ctxSess *ctxSess.Context, filter *FilterSpaceship) (resp *FindSpaceshipResponse, err error) {
	entities, err := s.spaceshipRepo.FindByFilter(filter.Name, filter.Class, filter.Status)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return nil, err
	}

	resp = &FindSpaceshipResponse{}
	for _, v := range entities {
		resp.Data = append(resp.Data, &DataSpaceship{
			ID:     v.Id,
			Name:   v.Name,
			Status: v.Status,
		})
	}

	return
}

func (s *service) FindByID(ctxSess *ctxSess.Context, id int64) (resp interface{}, err error) {
	resp, err = s.spaceshipRepo.FindById(id)
	if err != nil {
		ctxSess.ErrorMessage = err.Error()
		return nil, err
	}

	return
}
