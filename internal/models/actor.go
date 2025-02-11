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
	Id                string    `json:"id"`
	Type              ActorType `json:"type"`
	PreferredUsername string    `json:"preferredUsername"`
}
