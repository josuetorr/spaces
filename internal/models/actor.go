package models

type ActorType string

const (
	Application  ActorType = "Application"
	Group        ActorType = "Group"
	Orginization ActorType = "Orginization"
	Person       ActorType = "Person"
	Service      ActorType = "Service"
)

// TODO: maybe we don't need all this info here
// since every actor has inbox, outbox, following, followers and liked are all their respective endpoints
type Actor struct {
	Id        string    `json:"id"`
	Type      ActorType `json:"type"`
	Inbox     string    `json:"inbox"`
	Outbox    string    `json:"outbox"`
	Following string    `json:"following"`
	Followers string    `json:"followers"`
	Liked     string    `json:"liked"`
}
