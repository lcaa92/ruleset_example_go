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
	rs.Deps = append(rs.Deps, [2]string{A, B})
}

// AddConflict add conflict in RuleSet
func (rs *RuleSet) AddConflict(A string, B string) {
	rs.Conflicts = append(rs.Conflicts, [2]string{A, B})
}

func inArray(element string, arr []string) bool {
	for _, el := range arr {
		if el == element {
			return true
		}
	}
	return false
}

func getDep(element string, deps [][2]string) (string, bool) {
	for i := 0; i < len(deps); i++ {
		if deps[i][0] != element {
			continue
		}
		return deps[i][1], true
	}
	return "", false
}

// IsCoherent validate deps and conclicts RuleSet and return true or false
func (rs *RuleSet) IsCoherent() (isCoherent bool) {
	isCoherent = true

	for i := 0; i < len(rs.Deps); i++ {
		// Get all item dependencies
		var dependecy []string
		next := false
		itemDep := rs.Deps[i][0]

		dependecy = append(dependecy, rs.Deps[i][1])
		checkDeps := rs.Deps[i][1]
		for {
			newDeps, checkNext := getDep(checkDeps, rs.Deps)

			if checkNext == true {
				checkDeps = newDeps
				dependecy = append(dependecy, newDeps)
			} else {
				next = checkNext
			}

			if next == false {
				break
			}
		}

		// Check conflicts
		for k := 0; k < len(rs.Conflicts); k++ {
			if rs.Conflicts[k][0] != itemDep {
				continue
			}
			conflict := rs.Conflicts[k][1]
			if inArray(conflict, dependecy) {
				return false
			}
		}
	}

	fmt.Println("")
	return
}

func main() {
	fmt.Println("Starting ...")
	rs := NewRuleSet()

	rs.AddDep("A", "B")
	rs.AddDep("B", "C")
	rs.AddConflict("A", "C")

	if rs.IsCoherent() {
		fmt.Println("RuleSet is coherent")
	} else {
		fmt.Println("RuleSet is not coherent")
	}

	fmt.Println("")

	fmt.Println("Done ...")
}
