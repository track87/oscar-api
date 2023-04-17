// Package http declare something
// MarsDong 2023/4/17
package http

import (
	"reflect"

	"github.com/track87/oscar-api/utils"
)

var all = utils.NewSet()

func AllRequests() map[string]reflect.Type {
	return all.Clone()
}

func ExistsRequest(action string) bool {
	return all.Exist(action)
}

func GetRequestObject(action string) reflect.Type {
	return all.Get(action)
}
