package main

import (
	"fmt"
	"testing"
)

func TestOpts(t *testing.T) {
	fmt.Println("Starting Tests...")

	rs := RuleSets{}
	rs.AddDep()
	rs.AddConflict()
	rs.IsCoherent()
	fmt.Println("Done Tests ...")
}
