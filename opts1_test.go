package main

import (
	"fmt"
	"testing"
)

func TestOpts(t *testing.T) {
	fmt.Println("Starting Tests...")

	rs := NewRuleSet()
	rs.AddDep("A", "B")
	rs.AddDep("C", "D")
	rs.AddConflict("D", "F")
	rs.IsCoherent()

	if rs.IsCoherent() {
		fmt.Println("RuleSet is coherent")
	} else {
		fmt.Println("RuleSet is not coherent")
	}

	fmt.Println("Done Tests ...")
}
