package data

import (
	"context"
	"log/slog"

	"github.com/go-kivik/kivik/v4"
	"gitlab.com/josuetorr/spaces/internal/services"
)

// The "Repository" type can store any data type that implement "Storable".
// It stores them as a document for couchDB
type Repository[T services.Storable] struct {
	log *slog.Logger
	db  *kivik.DB
}

func NewRepository[T services.Storable](log *slog.Logger, db *kivik.DB) Repository[T] {
	return Repository[T]{log: log, db: db}
}

func (r Repository[T]) Create(data *T) (string, error) {
	docId, _, err := r.db.CreateDoc(context.TODO(), data)
	if err != nil {
		r.log.Error(err.Error())
		return docId, err
	}

	return docId, nil
}

func (r Repository[T]) Update(id string, data T) error {
	panic("implement Update")
}

func (r Repository[T]) Patch(id string, data T) error {
	panic("implement Patch")
}

func (r Repository[T]) Exists(id string) (bool, error) {
	a, err := r.GetById(id)

	if err != nil && err.Error() == "Not Found: missing" {
		return false, nil
	}

	return a != nil, err
}

func (r Repository[T]) GetById(id string) (*T, error) {
	var a *T
	query := map[string]any{
		"selector": map[string]string{
			"id": id,
		},
		"limit": 1,
	}
	rows := r.db.Find(context.TODO(), query)
	defer rows.Close()

	for rows.Next() {
		if err := rows.ScanDoc(&a); err != nil {
			return nil, err
		}
	}
	return a, nil
}

func (r Repository[T]) GetByDocId(id string) (*T, error) {
	var a T
	if err := r.db.Get(context.TODO(), id).ScanDoc(&a); err != nil {
		return nil, err
	}
	return &a, nil
}

func (r Repository[T]) GetAll() ([]T, error) {
	panic("implement GetAll")
}

func (r Repository[T]) Delete(id string, hard bool) error {
	panic("implement Delete")
}
