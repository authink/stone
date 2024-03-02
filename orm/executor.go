package orm

import (
	"database/sql"

	"github.com/authink/inkstone/model"
	"github.com/jmoiron/sqlx"
)

type Executor interface {
	Exec(string, ...any) (sql.Result, error)
	NamedExec(string, any) (sql.Result, error)
	PrepareNamed(string) (*sqlx.NamedStmt, error)
	Get(any, string, ...any) error
	Select(any, string, ...any) error
}

func NamedInsert(executor Executor, statement string, m model.Identifier) error {
	return namedExec(executor, statement, m, afterInsert)
}

func NamedUpdate(executor Executor, statement string, m model.Identifier) error {
	return namedExec(executor, statement, m, afterUpdate)
}

func NamedSave(executor Executor, statement string, m model.Identifier) error {
	return namedExec(executor, statement, m, afterSave)
}

func Get(executor Executor, dest any, statement string, args ...any) error {
	return executor.Get(
		dest,
		statement,
		args...,
	)
}

func Select(executor Executor, list any, statement string, args ...any) error {
	return executor.Select(
		list,
		statement,
		args...,
	)
}

func Delete(executor Executor, statement string, args ...any) (err error) {
	result, err := executor.Exec(
		statement,
		args...,
	)
	if err != nil {
		return
	}
	afterDelete(result)
	return
}

func Count(executor Executor, statement string, dest, arg any) (err error) {
	stmt, err := executor.PrepareNamed(statement)
	if err != nil {
		return
	}
	err = stmt.Get(dest, arg)
	return
}

func Pagination(executor Executor, statement string, dest, arg any) (err error) {
	stmt, err := executor.PrepareNamed(statement)
	if err != nil {
		return
	}
	err = stmt.Select(dest, arg)
	return
}

func namedExec(executor Executor, statement string, m model.Identifier, afterExec func(sql.Result, model.Identifier) error) (err error) {
	result, err := executor.NamedExec(
		statement,
		m,
	)
	if err != nil {
		return
	}

	err = afterExec(result, m)
	return
}
