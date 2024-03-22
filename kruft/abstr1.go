package main

import (
	"fmt"
	"strings"
)

// пример #1: просто типы-псевдонимы (алиасы), для удобства

// такие псевдонимы (внимание на равно (=)) нужны исключительно для удобства,
// когда функции и методы принимают несколько параметров одинакового типа,
// чтобы не путаться.

// создадим 2 разны псевдонима для строкового типа:
type AnimalFamily = string
type AnimalSpecies = string

func PrintAnimalInfo(family AnimalFamily, species AnimalSpecies) {
	// такие типы-псевдонимы могут использоваться везде, где требуется и оригинальный тип
	// и наоборот
	// (внимание - передача переменных типа-псевдонима в методы, работающие со строками)
	fmt.Printf("животное %s относится к семейству %s\n", strings.ToLower(species), strings.ToLower(family))
}

func main() {
	var (
		family  AnimalFamily  = "кошачьи"
		species AnimalSpecies = "рысь"
	)

	// мы можем передавать в функции и методы как переменные конкретных типов-псевдонимов
	PrintAnimalInfo(family, species)

	var (
		family2  string = "псовые"
		species2 string = "рыжий волк"
	)

	PrintAnimalInfo(family2, species2)
}
