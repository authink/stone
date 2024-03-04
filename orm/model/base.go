package model

import "time"

type Created struct {
	Record
	CreatedAt *time.Time `db:"created_at"`
}

type Base struct {
	Created
	UpdatedAt *time.Time `db:"updated_at"`
}
