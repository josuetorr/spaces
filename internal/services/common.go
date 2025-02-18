package services

import (
	ap "github.com/go-ap/activitypub"
)

type WithID interface {
	GetID() ap.ID
}

type Storable interface {
	WithID
}

type Repository[T Storable] interface {
	Create(*T) error
	Update(string, T) error
	Patch(string, T) error

	Exists(string) (bool, error)
	GetById(string) (*T, error)
	GetAll() ([]T, error)

	Delete(string, bool) error
}
