package data

import (
	"log/slog"

	ap "github.com/go-ap/activitypub"
	"github.com/go-kivik/kivik/v4"
)

type (
	Activity     = ap.Activity
	ActivityRepo struct {
		Repository[Activity]
	}
)

func NewActivityRepo(log *slog.Logger, db *kivik.DB) ActivityRepo {
	return ActivityRepo{Repository[Activity]{log: log, db: db}}
}
