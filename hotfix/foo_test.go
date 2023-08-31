package hotfix

import (
	"fmt"
	"github.com/agiledragon/gomonkey/v2"
	"github.com/cherry-game/cherry-hotfix/model"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"reflect"
	"testing"
)

func TestFixFooHelloFunc(t *testing.T) {
	// 构建foo1
	foo1 := model.Foo{
		String: "foo1",
	}

	// 打印foo1的信息
	t.Logf("替换前: foo1 -> %+v, Helo() -> %s", &foo1, foo1.Hello())

	// 获取Hello()的替换函数
	replaceFuncValueOf := GetReplaceFunc()

	// 打桩替换
	patches := monkeyApplyMethod(reflect.TypeOf(&model.Foo{}), "Hello", replaceFuncValueOf)

	// 打印foo1的信息，Hello()函数已被替换
	t.Logf("替换后: foo1 -> %+v, Helo() -> %s", &foo1, foo1.Hello())

	// 重置还原Hello()
	patches.Reset()

	// 打印foo1的信息，已还原
	t.Logf("还原后: foo1 -> %+v, Helo() -> %s", &foo1, foo1.Hello())
}

func monkeyApplyMethod(target reflect.Type, methodName string, dest reflect.Value) *gomonkey.Patches {
	patches := gomonkey.NewPatches()
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic("retrieve method by name failed")
	}
	return patches.ApplyCore(m.Func, dest)
}

func GetReplaceFunc() reflect.Value {
	// 演示用，先写死
	patchScript := `
package patch

import "github.com/cherry-game/cherry-hotfix/model"

func FixFooHello() func(foo *model.Foo) string {
	return func(foo *model.Foo) string {
		return "foo.Hello() is fixed"
	}
}
`
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	// 演示用，先写死,可以通过yaegi生成
	i.Use(map[string]map[string]reflect.Value{
		"github.com/cherry-game/cherry-hotfix/model/model": {
			"Foo": reflect.ValueOf((*model.Foo)(nil)),
		},
	})

	if _, err := i.Eval(patchScript); err != nil {
		fmt.Println(err)
	}

	// 获取函数对象
	val, err := i.Eval(`patch.FixFooHello`)
	if err != nil {
		fmt.Println(err)
	}

	values := val.Call(nil)
	//fmt.Println("patch func reflect type ->", values[0].Type())

	return values[0]
}
