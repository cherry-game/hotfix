package foo

import (
	"fmt"
	"github.com/cherry-game/cherry-hotfix/hotfix"
	"github.com/cherry-game/cherry-hotfix/test/model"
	"reflect"
)

func GetPatch() *hotfix.FuncPatch {
	fmt.Println("ssxxx")

	fn := func(foo *model.Foo) string {
		foo.M1Int.Int = 1
		return "func is fixed"
	}

	return &hotfix.FuncPatch{
		StructType: reflect.TypeOf(&model.Foo{}),
		FuncName:   "Hello",
		FuncValue:  reflect.ValueOf(fn),
	}
}
