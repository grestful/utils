package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestNextId(t *testing.T) {
	i := 0
	for {
		if i == 100 {
			break
		}
		fmt.Println(NextId())
		c := time.After(100*time.Second)
		<- c
		i++
	}
}
