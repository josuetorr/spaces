package models

type ActivityType string

const (
	Accept          ActivityType = "Accept"
	Add             ActivityType = "Add"
	Announce        ActivityType = "Announce"
	Arrive          ActivityType = "Arrive"
	Block           ActivityType = "Block"
	Create          ActivityType = "Create"
	Delete          ActivityType = "Delete"
	Dislike         ActivityType = "Dislike"
	Flag            ActivityType = "Flag"
	Follow          ActivityType = "Follow"
	Ignore          ActivityType = "Ignore"
	Invite          ActivityType = "Invite"
	Join            ActivityType = "Join"
	Leave           ActivityType = "Leave"
	Like            ActivityType = "Like"
	Listen          ActivityType = "Listen"
	Move            ActivityType = "Move"
	Offer           ActivityType = "Offer"
	Question        ActivityType = "Question"
	Reject          ActivityType = "Reject"
	Read            ActivityType = "Read"
	Remove          ActivityType = "Remove"
	TentativeReject ActivityType = "TentativeReject"
	TentativeAccept ActivityType = "TentativeAccept"
	Travel          ActivityType = "Travel"
	Undo            ActivityType = "Undo"
	Update          ActivityType = "Update"
	View            ActivityType = "View"
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
