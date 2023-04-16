package tests

import (
	"fmt"
	"testing"

	set "github.com/katasec/utils/set"
)

func TestIntSet(t *testing.T) {
	myset := set.NewSet[int]()
	myset.Add(5)
	myset.Add(6)
	myset.Add(7)
	myset.Del(7)
	for i := range myset.Items {
		fmt.Printf("item=%d\n", i)
	}
}

func TestStringSet(t *testing.T) {
	a := set.NewSet[string]()
	a.Add("Ameer")
	a.Add("Is")
	a.Add("Testing")
	a.Del("Ameer")

	for i := range a.Items {
		fmt.Printf("item=%s\n", i)
	}
}
