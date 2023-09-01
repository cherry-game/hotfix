package hotfix_test

import (
	"fmt"
	"github.com/cherry-game/cherry-hotfix/hotfix"
	"github.com/cherry-game/cherry-hotfix/test/model"
	"testing"
	"time"

	"github.com/cherry-game/cherry-hotfix/symbols"
)

func TestFixFooHelloFunc(t *testing.T) {
	// 构建foo1
	foo1 := &model.Foo{
		String: "foo1",
	}

	// 打印foo1的信息
	fmt.Println("初始: ", foo1, "-> Hello() ->", foo1.Hello())

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
		filePath = "./_patch_files/foo/foo.go"
		evalText = "foo.GetPatch()"
	)

	// 修复Hello()函数
	patches, err := hotfix.ApplyFunc(filePath, evalText, symbols.Symbols)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 打印foo1的信息，Hello()函数已被替换!
	fmt.Println("替换: ", foo1, "-> Hello() ->", foo1.Hello())

	// 重置Hello()函数
	patches.Reset()

	// 打印foo1的信息，已还原
	fmt.Println("还原: ", foo1, "-> Hello() ->", foo1.Hello())
}
