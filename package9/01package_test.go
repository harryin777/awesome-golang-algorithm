package package9

import (
	"fmt"
	"testing"
)

func Test_01package(t *testing.T) {
	fmt.Println(package01(6, 4, []int{4, 2, 3, 6}, []int{3, 4, 1, 2}))
	fmt.Println(package01V2(6, 4, []int{4, 2, 3, 6}, []int{3, 4, 1, 2}))
}
