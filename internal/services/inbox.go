package services

import ap "github.com/go-ap/activitypub"

type (
	InboxRepository interface {
		Repository[Activity]
		GetInboxByActorId(id string) ([]*Activity, error)
	}
	InboxService struct {
		repo InboxRepository
	}
)

func NewInboxService(inboxRepository InboxRepository) InboxService {
	return InboxService{repo: inboxRepository}
}

func (s InboxService) GetInboxByActorId(id string) (*Collection, error) {
	activities, err := s.repo.GetInboxByActorId(id)
	if err != nil {
		return nil, err
	}
	activitiesLen := uint(len(activities))

	items := make([]ap.Item, activitiesLen)
	for _, a := range activities {
		items = append(items, *a)
	}

	c := ap.CollectionNew(ap.EmptyID)
	c.TotalItems = activitiesLen
	c.Items = items
	return c, nil
}
