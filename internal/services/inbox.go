package services

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

func (s InboxService) GetInboxByActorId(id string) ([]*Activity, error) {
	return s.repo.GetInboxByActorId(id)
}
