// Package chaos declare something
// MarsDong 2023/4/11
package chaos

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHostMemStress_Object(t *testing.T) {
	chaos := HostMemStress{}
	v, _ := json.MarshalIndent(chaos, "", "  ")
	fmt.Println(string(v))
}
