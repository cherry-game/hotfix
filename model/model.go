package model

import "github.com/cherry-game/cherry-hotfix/model/m1"

type Foo struct {
	String string
	M1Int  m1.M1
}

func (p *Foo) Hello() string {
	return p.String
}
