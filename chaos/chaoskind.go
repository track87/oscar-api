// Package chaos declare something
// MarsDong 2023/3/30
package chaos

import (
	"reflect"

	"github.com/track87/oscar-api/utils"
)

var all = utils.NewSet()

func AllChaosKinds() map[string]reflect.Type {
	return all.Clone()
}

func ExistsChaosKind(kind string) bool {
	return all.Exist(kind)
}
