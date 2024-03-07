package orm

import (
	"database/sql"
	"errors"

	"github.com/authink/inkstone/orm/model"
)

func afterInsert(result sql.Result, m model.Identifier) (err error) {
	lastId, err := result.LastInsertId()
	if err != nil {
		return
	}

	m.SetId(uint32(lastId))
	return
}

func afterSave(result sql.Result, m model.Identifier) (err error) {
	if err = afterInsert(result, m); err != nil {
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return
	} else if rowsAffected == 0 {
		err = errors.New("duplicate key")
	}
	return
}

func afterUpdate(sql.Result, model.Identifier) error {
	return nil
}

func afterDelete(sql.Result, model.Identifier) error {
	return nil
}
