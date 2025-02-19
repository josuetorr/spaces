package data

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/go-kivik/kivik/v4"
)

type InboxRepository struct {
	Repository[Activity]
}

func NewInboxRepository(log *slog.Logger, db *kivik.DB) InboxRepository {
	return InboxRepository{Repository[Activity]{log: log, db: db}}
}

func (r InboxRepository) GetInboxByActorId(id string) ([]*Activity, error) {
	opts := kivik.Params(map[string]any{"key": id})
	rows := r.db.Query(context.TODO(), "_design/inbox", "_view/actor-inbox", opts)
	defer rows.Close()

	for rows.Next() {
		var a *Activity
		if err := rows.ScanDoc(a); err != nil {
			return nil, err
		}

		fmt.Printf("%+v\n", a)
	}
	return nil, nil
}
