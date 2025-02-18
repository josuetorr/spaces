package services

import (
	"log/slog"

	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/utils"
)

type (
	Activity           = ap.Activity
	ActivityRepository interface {
		Repository[Activity]
	}
	ActivityService struct {
		log  *slog.Logger
		repo ActivityRepository
	}
)

func NewActivityService(log *slog.Logger, activityRepo ActivityRepository) ActivityService {
	return ActivityService{log: log, repo: activityRepo}
}

func (s ActivityService) ActivityCreate(data CreateActorData) error {
	a := ap.ActivityNew(ap.ID("test"), ap.FollowType, ap.ObjectNew(ap.ActorType))
	if err := s.repo.Create(a); err != nil {
		return err
	}
	return nil
}

func (s ActivityService) ActivityExists(id string) (bool, error) {
	id = utils.GetFullId("users", id)
	return s.repo.Exists(id)
}

func (s ActivityService) ActivityGetById(id string) (*Activity, error) {
	id = utils.GetFullId("users", id)
	return s.repo.GetById(id)
}
