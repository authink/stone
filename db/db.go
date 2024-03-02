package db

import (
	"github.com/jmoiron/sqlx"
)

type DB struct {
	*sqlx.DB
}

type TxFunc func(tx *sqlx.Tx) error

func (db *DB) Transaction(txFunc TxFunc) (err error) {
	tx := db.MustBegin()

	if err = txFunc(tx); err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}
