package orm

import (
	"github.com/authink/inkstone/orm/model"
	"github.com/jmoiron/sqlx"
)

type Inserter[T any] interface {
	Insert(*T) error
	InsertTx(*sqlx.Tx, *T) error
}

type Saver[T any] interface {
	Save(*T) error
	SaveTx(*sqlx.Tx, *T) error
}

type Updater[T any] interface {
	Update(*T) error
	UpdateTx(*sqlx.Tx, *T) error
}

type Geter[T any] interface {
	Get(int) (*T, error)
	GetTx(*sqlx.Tx, int) (*T, error)
}

type Deleter[T any] interface {
	Delete(int) error
	DeleteTx(*sqlx.Tx, int) error
}

type Finder[T any] interface {
	Find(...any) ([]T, error)
}

type Counter interface {
	Count(...any) (int, error)
	CountTx(*sqlx.Tx, ...any) (int, error)
}

type Pager[T any] interface {
	PaginationTx(*sqlx.Tx, model.Pager) ([]T, error)
}
