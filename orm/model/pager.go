package model

type Pager interface {
	Arg
	Num() int
	Size() int
}

type Page struct {
	Argument
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
