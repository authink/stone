package inkstone

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type Model struct {
	Id        uint32
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ORM[T any] interface {
	Save(*T) error
	SaveWithTx(*T, *sqlx.Tx) error
	Get(int) (*T, error)
	Find() ([]T, error)
	Delete(int) error
}

type TxFunc func(tx *sqlx.Tx) error

func Transaction(app *AppContext, txFunc TxFunc) (err error) {
	tx := app.DB.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

type SQL interface {
	Insert() string
	Delete() string
	Update() string
	Get() string
	Find() string
}
