package orm

import "time"

type Identifier interface {
	SetId(uint32)
}

type Record struct {
	Id uint32
}

// SetId implements Identifier.
func (i *Record) SetId(id uint32) {
	i.Id = id
}

var _ Identifier = (*Record)(nil)

type Created struct {
	Record
	CreatedAt *time.Time `db:"created_at"`
}

type Model struct {
	Created
	UpdatedAt *time.Time `db:"updated_at"`
}
