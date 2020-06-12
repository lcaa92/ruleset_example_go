package main

import (
	"fmt"
	"testing"
)

func TestOpts(t *testing.T) {
	fmt.Println("Starting Tests...")

	rs := NewRuleSet()
	rs.AddDep("A", "B")
	rs.AddDep("B", "C")
	rs.AddConflict("A", "C")
	rs.IsCoherent()

	if rs.IsCoherent() {
		fmt.Println("RuleSet is coherent")
	} else {
		fmt.Println("RuleSet is not coherent")
	}

	fmt.Println("Done Tests ...")
}
