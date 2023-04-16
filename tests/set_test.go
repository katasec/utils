package tests

import (
	"fmt"
	"testing"

	set "github.com/katasec/utils/set"
)

func TestIntSet(t *testing.T) {
	a := set.NewSet[int]()
	a.Add(5)
	a.Add(6)
	a.Add(7)
	a.Del(7)
	for i := range a.Items {
		fmt.Printf("Key=%d\n", i)
	}
}

func TestStringSet(t *testing.T) {
	a := set.NewSet[string]()
	a.Add("Ameer")
	a.Add("Is")
	a.Add("Testing")
	a.Del("Ameer")

	for i := range a.Items {
		fmt.Printf("Key=%s\n", i)
	}
}
