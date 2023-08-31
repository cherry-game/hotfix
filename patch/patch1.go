package patch

import "github.com/cherry-game/cherry-hotfix/model"

func FixFooHello() func(foo *model.Foo) string {
	return func(foo *model.Foo) string {
		return "foo.Hello() is fixed"
	}
}
