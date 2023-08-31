package patch

import "github.com/cherry-game/cherry-hotfix/model"

func FixFooHello(foo *model.Foo) string {
	return "func is fixed"
}
