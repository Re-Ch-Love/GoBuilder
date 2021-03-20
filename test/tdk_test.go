package test

import (
	"fmt"
	tdk "gitee.com/KongchengPro/GoBuilder/pkg/tdk/commands"
	"reflect"
	"testing"
)

func TestStringToCommand(t *testing.T) {
	cmd, args := tdk.ParseCommand(`a "aa bb" "" aabb`)
	fmt.Println(cmd, args)
	if cmd != "a" || !reflect.DeepEqual(args, []string{"aa bb", "", "aabb"}) {
		t.FailNow()
	}
}
