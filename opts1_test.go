package main

import (
	"fmt"
	"testing"
)

func TestOpts(t *testing.T) {
	fmt.Println("Starting Tests...")

	// Example RuleSet Coherent
	rs := NewRuleSet()

	rs.AddDep("A", "B")
	rs.AddDep("B", "C")
	rs.AddDep("C", "D")
	rs.AddDep("D", "E")
	rs.AddConflict("A", "E")

	if rs.IsCoherent() {
		fmt.Println("RuleSet1 is coherent")
	} else {
		fmt.Println("RuleSet1 is not coherent")
	}

	// Example RuleSet not Coherent
	rs2 := NewRuleSet()
	rs2.AddDep("A", "B")
	rs2.AddDep("B", "T")
	rs2.AddDep("A", "F")
	rs2.AddDep("B", "C")

	rs2.AddConflict("B", "F")

	if rs2.IsCoherent() {
		fmt.Println("RuleSet2 is coherent")
	} else {
		fmt.Println("RuleSet2 is not coherent")
	}

	fmt.Println("Done Tests ...")
}
