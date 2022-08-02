package main

import (
	"fmt"
	"github.com/jckli/valorant.go"
)

func main() {
	a := valorant.Authentication("laurabyohashi", "Oofergang1")
	fmt.Println(a)
	b, _ := valorant.Contracts_fetch(a)
	fmt.Println(b.ProcessedMatches[0].Id)
}