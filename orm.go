package inkstone

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type ModelOf interface {
	Of() *Model
}

type Model struct {
	Id        uint32
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}

func (m *Model) Of() *Model {
	return m
}

var _ ModelOf = (*Model)(nil)

type ORM[T any] interface {
	Insert(*T) error
	InsertWithTx(*T, *sqlx.Tx) error
	Save(*T) error
	SaveWithTx(*T, *sqlx.Tx) error
	Get(int) (*T, error)
	Find() ([]T, error)
	Delete(int) error
}

type TxFunc func(tx *sqlx.Tx) error

func Transaction(app *AppContext, txFunc TxFunc) (err error) {
	tx := app.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

type SQL interface {
	Insert() string
	Save() string
	Delete() string
	Update() string
	Get() string
	Find() string
}
