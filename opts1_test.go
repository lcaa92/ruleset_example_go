package main

import (
	"fmt"
	"testing"
)

func TestOpts(t *testing.T) {
	fmt.Println("Starting Tests...")

	rs := NewRuleSet()
	itemA := Item{}
	itemB := Item{}
	rs.AddDep(itemA, itemB)
	rs.AddConflict(itemA, itemB)
	rs.IsCoherent(itemA, itemB)
	fmt.Println("Done Tests ...")
}
