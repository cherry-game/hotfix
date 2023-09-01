package symbols

import (
	"reflect"

	"github.com/cherry-game/cherry-hotfix/hotfix"
	"github.com/traefik/yaegi/stdlib"
)

var Symbols = map[string]map[string]reflect.Value{}

func init() {
	for k, v := range stdlib.Symbols {
		Symbols[k] = v
	}

	Symbols["github.com/cherry-game/cherry-hotfix/hotfix/hotfix"] = map[string]reflect.Value{
		// type definitions
		"FuncPatch": reflect.ValueOf((*hotfix.FuncPatch)(nil)),
	}
}

// 点击生成符号表
//go:generate yaegi extract github.com/cherry-game/cherry-hotfix/test/model model
