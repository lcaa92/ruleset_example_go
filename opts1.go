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

func getArrDep(element string, deps [][2]string) (arrDeps []string) {
	for i := 0; i < len(deps); i++ {
		if deps[i][0] != element {
			continue
		}
		arrDeps = append(arrDeps, deps[i][1])
	}
	return
}

// IsCoherent validate deps and conclicts RuleSet and return true or false
func (rs *RuleSet) IsCoherent() (isCoherent bool) {
	isCoherent = true
	var itemVerify []string

	for i := 0; i < len(rs.Deps); i++ {
		// Get all item dependencies
		var dependencies []string
		var elementsChecked []string
		var elementsPending []string

		itemDep := rs.Deps[i][0]

		// Not repeat check dep
		if inArray(itemDep, itemVerify) {
			continue
		}
		itemVerify = append(itemVerify, itemDep)
		elementsPending = append(elementsPending, itemDep)

		for {
			if len(elementsPending) == 0 {
				break
			}

			// Alter element to check and remove from pending elements
			checkDeps := elementsPending[0]
			elementsPending = append(elementsPending[:0], elementsPending[0+1:]...)

			if (inArray(checkDeps, elementsChecked)) == false {
				elementsChecked = append(elementsChecked, checkDeps)
			}

			getDeps := getArrDep(checkDeps, rs.Deps)

			for _, el := range getDeps {
				// Add el in array dependencies
				if inArray(el, dependencies) == false {
					dependencies = append(dependencies, el)
				}

				// Verify if el in elements already checked
				if inArray(el, elementsChecked) {
					continue
				}
				elementsChecked = append(elementsChecked, el)

				// Validate if needs to add element in elementsPending
				if inArray(el, elementsPending) == false {
					elementsPending = append(elementsPending, el)
				}
			}
		}

		// Check conflicts
		for _, dep := range dependencies {
			for k := 0; k < len(rs.Conflicts); k++ {
				if rs.Conflicts[k][0] != dep {
					continue
				}
				conflict := rs.Conflicts[k][1]
				if inArray(conflict, dependencies) {
					return false
				}
			}
		}
	}

	fmt.Println("")
	return
}

func main() {
	fmt.Println("Starting ...")

	// Example RuleSet not Coherent
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

	// Example RuleSet Coherent
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

	fmt.Println("")

	fmt.Println("Done ...")
}
