package model

type Foo struct {
	String string
}

func (p *Foo) Hello() string {
	return p.String
}
