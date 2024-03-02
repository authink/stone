package model

type Pager interface {
	Num() int
	Size() int
}

type Page struct {
	Offset int
	Limit  int
}

func (p *Page) Num() int {
	return (p.Offset / p.Limit) + 1
}

func (p *Page) Size() int {
	return p.Limit
}

var _ Pager = (*Page)(nil)
