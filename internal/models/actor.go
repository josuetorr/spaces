package models

type ActorType string

const (
	Application  ActorType = "Application"
	Group        ActorType = "Group"
	Orginization ActorType = "Orginization"
	Person       ActorType = "Person"
	Service      ActorType = "Service"
)

type Actor struct {
	Context   string    `json:"@context"`
	Id        string    `json:"id"`
	Type      ActorType `json:"type"`
	Inbox     string    `json:"inbox"`
	Outbox    string    `json:"outbox"`
	Following string    `json:"following"`
	Followers string    `json:"followers"`
	Liked     string    `json:"liked"`
}
