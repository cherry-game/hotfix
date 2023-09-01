package hotfix

import "reflect"

type FuncPatch struct {
	StructType reflect.Type
	FuncName   string
	FuncValue  reflect.Value
}
