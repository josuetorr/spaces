package data

import (
	"log/slog"

	ap "github.com/go-ap/activitypub"
	"github.com/go-kivik/kivik/v4"
)

type Activity = ap.Activity

type ActivityRepository struct {
	Repository[Activity]
}

func NewActivityRepository(log *slog.Logger, db *kivik.DB) ActivityRepository {
	return ActivityRepository{Repository[Activity]{log: log, db: db}}
}
