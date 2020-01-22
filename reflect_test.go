package utils

import (
	"fmt"
	"testing"
)

func TestStructSetFieldValue(t *testing.T) {

	type E struct{ N string }
	e1 := E{N:"222"}
	StructSetFieldValue(&e1, "N", "fff")
	fmt.Println(e1)

}