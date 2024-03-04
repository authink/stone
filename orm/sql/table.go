package sql

type Table interface {
	TbName() string
}

const (
	Id        = "id"
	CreatedAt = "created_at"
	UpdatedAt = "updated_at"
)
