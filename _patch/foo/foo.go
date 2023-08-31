package foo

import (
	"github.com/cherry-game/cherry-hotfix/model"
)

// FixHello Hello()修复函数
func FixHello(foo *model.Foo) string {
	//fmt.Println("输出M1Int值: foo.M1Int.Int = ", foo.M1Int.Int)
	//fmt.Println("输出M1Int类型:", reflect.TypeOf(foo.M1Int))

	foo.M1Int.Int = 1

	return "func is fixed"
}
