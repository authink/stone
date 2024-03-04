package sql

import (
	"database/sql"
	"fmt"
	"strings"
)

func Col(tableAlias, col string) string {
	return fmt.Sprintf("%s.%s", tableAlias, col)
}

func Named(name string) sql.NamedArg {
	return sql.Named(name, nil)
}

func ReplaceAtWithColon(statement string) string {
	return strings.ReplaceAll(statement, "@", ":")
}

func EQ(op1, op2 string) string {
	return fmt.Sprintf("%s = %s", op1, op2)
}

func LimitAndOffset(statement string) string {
	return fmt.Sprintf("%s LIMIT :limit OFFSET :offset", statement)
}
