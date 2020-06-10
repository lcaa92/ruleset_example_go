package main

import "fmt"

// Item struct
type Item struct {
	name string
}

// RuleSets struct
type RuleSets struct {
	itemA Item
	itemB Item
}

// AddDep return nil
func (rs RuleSets) AddDep() {
	fmt.Println("AddDep func")
}

// AddConflict return nil
func (rs RuleSets) AddConflict() {
	fmt.Println("AddConflict func")
}

// IsCoherent return nil
func (rs RuleSets) IsCoherent() {
	fmt.Println("IsCoherent func")
}

func main() {
	fmt.Println("Starting ...")
	rs := RuleSets{}
	rs.AddDep()
	rs.AddConflict()
	rs.IsCoherent()
	fmt.Println("Done ...")
}
