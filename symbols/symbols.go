package symbols

import (
	"reflect"

	"github.com/traefik/yaegi/stdlib"
)

var Symbols = map[string]map[string]reflect.Value{}

func init() {
	for k, v := range stdlib.Symbols {
		Symbols[k] = v
	}
}

// 点击生成符号表
//go:generate yaegi extract github.com/cherry-game/cherry-hotfix/model model
