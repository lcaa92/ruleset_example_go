package main

import "fmt"

// Item struct
type Item struct{}

// RuleSet struct
type RuleSet struct{}

// NewRuleSet return empty RuleSet
func NewRuleSet() RuleSet {
	return RuleSet{}
}

// AddDep return nil
func (rs RuleSet) AddDep(A Item, B Item) {
	fmt.Println("AddDep func")
}

// AddConflict return nil
func (rs RuleSet) AddConflict(A Item, B Item) {
	fmt.Println("AddConflict func")
}

// IsCoherent return nil
func (rs RuleSet) IsCoherent(A Item, B Item) {
	fmt.Println("IsCoherent func")
}

func main() {
	fmt.Println("Starting ...")
	rs := NewRuleSet()
	itemA := Item{}
	itemB := Item{}
	rs.AddDep(itemA, itemB)
	rs.AddConflict(itemA, itemB)
	rs.IsCoherent(itemA, itemB)
	fmt.Println("Done ...")
}
