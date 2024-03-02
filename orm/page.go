package orm

type Page interface {
	Num() int
	Size() int
}

type PageArgs struct {
	Offset int
	Limit  int
}

func (p *PageArgs) Num() int {
	return (p.Offset / p.Limit) + 1
}

func (p *PageArgs) Size() int {
	return p.Limit
}

var _ Page = (*PageArgs)(nil)
