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
	Get(*T) error
	GetTx(*sqlx.Tx, *T) error
}

type Deleter[T any] interface {
	Delete(*T) error
	DeleteTx(*sqlx.Tx, *T) error
}

type Finder[T any] interface {
	Find(...model.Arg) ([]T, error)
}

type Counter interface {
	Count(...model.Arg) (int, error)
	CountTx(*sqlx.Tx, ...model.Arg) (int, error)
}

type Pager[T any] interface {
	PaginationTx(*sqlx.Tx, model.Pager) ([]T, error)
}
