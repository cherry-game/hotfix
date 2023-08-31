package hotfix

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/cherry-game/cherry-hotfix/model"
	"github.com/cherry-game/cherry-hotfix/symbols"
	"github.com/traefik/yaegi/interp"
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

	var (
		replaceFuncName  = "Hello"          // 替换的函数名
		replaceFuncValue = GetReplaceFunc() // 获取Hello()的替换函数
	)

	// 针对model.Foo结构的 Hello()函数进行打桩替换
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
	// 构建解析器
	interpreter := interp.New(interp.Options{})

	// 导入符号表
	interpreter.Use(symbols.Symbols)

	// 读取补丁脚本
	_, err := interpreter.EvalPath("./../_patch/foo/foo.go")
	if err != nil {
		panic(err)
	}

	// 获取替换函数
	val, err := interpreter.Eval(`foo.FixHello`)
	if err != nil {
		panic(err)
	}

	fmt.Println("patch func reflect type ->", val.Type())

	return val
}
