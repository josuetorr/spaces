package services

import (
	"log/slog"

	ap "github.com/go-ap/activitypub"
	"gitlab.com/josuetorr/spaces/internal/utils"
)

type (
	Activity           = ap.Activity
	CreateActivityData struct {
		Type   string
		Actor  string
		Object string
	}
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

func (s ActivityService) ActivityCreate(data CreateActivityData) (string, error) {
	a := ap.ActivityNew(ap.EmptyID, ap.ActivityVocabularyType(data.Type), ap.ObjectNew(ap.ActorType))
	docId, err := s.repo.Create(a)
	if err != nil {
		return docId, err
	}
	return docId, nil
}

func (s ActivityService) ActivityExists(id string) (bool, error) {
	id = utils.GetFullId("users", id)
	return s.repo.Exists(id)
}

func (s ActivityService) ActivityGetById(id string) (*Activity, error) {
	id = utils.GetFullId("users", id)
	return s.repo.GetById(id)
}
