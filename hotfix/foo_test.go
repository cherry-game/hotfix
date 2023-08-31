package hotfix

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/cherry-game/cherry-hotfix/model"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func TestFixFooHelloFunc(t *testing.T) {
	// 构建foo1
	foo1 := &model.Foo{
		String: "foo1",
	}

	// 打印foo1的信息
	fmt.Println("替换前: ", foo1, "-> Hello() ->", foo1.Hello())

	// 模拟Hello()被调用
	for i := 0; i < 1000; i++ {
		go func(foo *model.Foo) {
			for {
				foo.Hello()
				time.Sleep(1 * time.Millisecond)
			}
		}(foo1)
	}

	// 获取Hello()的替换函数
	var (
		replaceFuncValue = GetReplaceFunc()
		replaceFuncName  = "Hello"
	)

	// 针对model.Foo的Hello()函数进行打桩替换
	patches := monkeyApplyMethod(reflect.TypeOf(&model.Foo{}), replaceFuncName, replaceFuncValue)

	// 打印foo1的信息，Hello()函数已被替换!
	fmt.Println("替换后: ", foo1, "-> Hello() ->", foo1.Hello())

	// 重置Hello()函数
	patches.Reset()

	// 打印foo1的信息，已还原
	fmt.Println("还原后: ", foo1, "-> Hello() ->", foo1.Hello())
}

func monkeyApplyMethod(target reflect.Type, methodName string, dest reflect.Value) *gomonkey.Patches {
	m, ok := target.MethodByName(methodName)
	if !ok {
		panic("retrieve method by name failed")
	}
	patches := gomonkey.NewPatches()
	return patches.ApplyCore(m.Func, dest)
}

func GetReplaceFunc() reflect.Value {
	// 演示用，先写死
	patchScript := `
package patch

import "github.com/cherry-game/cherry-hotfix/model"

func FixFooHello(foo *model.Foo) string {
	return "func is fixed"
}

`
	interpreter := interp.New(interp.Options{})
	interpreter.Use(stdlib.Symbols)

	// 演示用，先写死,可以通过yaegi生成
	interpreter.Use(map[string]map[string]reflect.Value{
		"github.com/cherry-game/cherry-hotfix/model/model": {
			"Foo": reflect.ValueOf((*model.Foo)(nil)),
		},
	})

	if _, err := interpreter.Eval(patchScript); err != nil {
		panic(err)
	}

	// 获取函数对象
	val, err := interpreter.Eval(`patch.FixFooHello`)
	if err != nil {
		panic(err)
	}

	fmt.Println("patch func reflect type ->", val.Type())

	return val
}
