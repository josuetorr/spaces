package models

type ActivityType string

const (
	Accept          ActorType = "Accept"
	Add             ActorType = "Add"
	Announce        ActorType = "Announce"
	Arrive          ActorType = "Arrive"
	Block           ActorType = "Block"
	Create          ActorType = "Create"
	Delete          ActorType = "Delete"
	Dislike         ActorType = "Dislike"
	Flag            ActorType = "Flag"
	Follow          ActorType = "Follow"
	Ignore          ActorType = "Ignore"
	Invite          ActorType = "Invite"
	Join            ActorType = "Join"
	Leave           ActorType = "Leave"
	Like            ActorType = "Like"
	Listen          ActorType = "Listen"
	Move            ActorType = "Move"
	Offer           ActorType = "Offer"
	Question        ActorType = "Question"
	Reject          ActorType = "Reject"
	Read            ActorType = "Read"
	Remove          ActorType = "Remove"
	TentativeReject ActorType = "TentativeReject"
	TentativeAccept ActorType = "TentativeAccept"
	Travel          ActorType = "Travel"
	Undo            ActorType = "Undo"
	Update          ActorType = "Update"
	View            ActorType = "View"
)

type Activity struct {
	// TODO: add properties extended from Object: "https://www.w3.org/TR/activitystreams-vocabulary/#dfn-object"
	Id     string       `json:"id"`
	Type   ActivityType `json:"type"`
	Actor  string       `json:"actor"`
	Object any          `json:"object"`
	Target string       `json:"target"`

	Result     string `json:"result"`
	Origin     string `json:"origin"`
	Instrument string `json:"instrument"`
}
