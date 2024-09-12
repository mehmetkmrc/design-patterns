package main

import (
	"fmt"

	factorymethod "github.com/mehmetkmrc/design-patterns/factory_method"
)

func main() {
	ak47, _ := factorymethod.getGun("Ak47")
	musket, _ := factorymethod.getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g factorymethod.IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}