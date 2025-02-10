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
	// MUST
	Context   string    `json:"@context"`
	Id        string    `json:"id"`
	Type      ActorType `json:"type"`
	Inbox     string    `json:"inbox"`
	Outbox    string    `json:"outbox"`
	Following string    `json:"following"`
	Followers string    `json:"followers"`
	Liked     string    `json:"liked"`

	// MAY
	Stream                     string `json:"stream"`
	PreferredUsername          string `json:"preferred_username"`
	Endpoints                  string `json:"endpoints"`
	ProxyUrl                   string `json:"proxy_url"`
	OAuthAuthorizationEndpoint string `json:"oAuthAuthorizationEndpoint"`
	OAuthTokenEndpoint         string `json:"oAuthTokenEndpoint"`
	SignClientKey              string `json:"signClientKey"`
	SharedInbox                string `json:"sharedInbox"`

	// TODO: add public key
}
