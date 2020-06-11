package main

import "fmt"

// RuleSet struct
type RuleSet struct {
	Deps      [][2]string
	Conflicts [][2]string
}

// NewRuleSet return empty RuleSet
func NewRuleSet() RuleSet {
	return RuleSet{}
}

// AddDep add dep in RuleSet
func (rs *RuleSet) AddDep(A string, B string) {
	fmt.Println("AddDep func")
	rs.Deps = append(rs.Deps, [2]string{A, B})
}

// AddConflict add conflict in RuleSet
func (rs *RuleSet) AddConflict(A string, B string) {
	fmt.Println("AddConflict func")
	rs.Conflicts = append(rs.Conflicts, [2]string{A, B})
}

// IsCoherent validate deps and conclicts RuleSet and return true or false
func (rs *RuleSet) IsCoherent() (isCoherent bool) {
	fmt.Println("IsCoherent func")
	return
}

func main() {
	fmt.Println("Starting ...")
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

	fmt.Printf("%s", rs)
	fmt.Println("")

	fmt.Println("Done ...")
}
