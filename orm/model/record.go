package model

type Record struct {
	Id uint32
}

// SetId implements Identifier.
func (i *Record) SetId(id uint32) {
	i.Id = id
}

var _ Identifier = (*Record)(nil)