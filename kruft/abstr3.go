package main

import (
	"fmt"
	"os"
)

// пример #3: полноценные типы на основе существующих типов

// такие псевдонимы (внимание на __отсутствие__ равно (=)) используются,
// когда требуется к существующим (как встроенным, так и пользовательским) типам
// добавить некую функциональность (т.е. методы)

// создадим 4 разных типа на основе встроенного типа строка
type Psovie string
type Koshachiyi string
type Polosatikovie string
type Loshadinie string

// интерфейсный тип: требуется, чтобы нечто имело метод PrintSpeciesName(),
// тогда это "нечто" будет удовлетворять интерфейсу Species
// и это "нечто" можно будет передать везде, где требуется переменная типа Species
type Species interface {
	GetSpeciesInfo() string
}

// добавляем метод PrintSpeciesName к типу Psovie
func (p Psovie) GetSpeciesInfo() string {
	return fmt.Sprintf("%s - это представитель семейства псовых", p)
}

// добавляем метод PrintSpeciesName к типу Koshachiyi
func (c Koshachiyi) GetSpeciesInfo() string {
	return fmt.Sprintf("%s - это представитель семейства кошачьих", c)
}

// NB: типу Polosatikovie НЕ добавляем метод PrintSpeciesName

// NB: типу Loshadinie добавляем метод PrintSpeciesName с другой сигнатурой
func (p Loshadinie) GetSpeciesInfo() error {
	fmt.Printf("%s - это представитель семейства лашадиных", p)
	return nil
}

// создаем функцию, которая принмиает интерфейсный типа Species
// т.е. ей можно передать все, что угодно, что удовлетворяет интерфейсу Species
func printInfoAboutSpecies(species Species) error {
	// NB: т.к мы принимаем интерфейсный тип, то мы знаем о переданной переменной только то,
	// что определено в описании интерфейсного типа - если в реальности тип переданной переменной
	// имеет еще какие-то типы, мы внутри этой функции этого не знаем и вызывать их не может

	// вызываем метод PrintSpeciesName у переданной переменной
	info := species.GetSpeciesInfo()
	if info == "" {
		return fmt.Errorf("species info is empty. this is unexpected")
	}

	fmt.Println(info)

	return nil
}

func main() {
	var (
		d = Psovie("собака")
	)

	// создем слайс из переменных типа Species
	// NB:  добавить в него переменные Polosatikovie и/или Loshadinie мы не можем,
	// т.к. они не удовлетворяют сигнатуре интерфейсного типа Species
	listOfSpecies := []Species{
		d,
		Koshachiyi("рысь"),
	}

	// запускаем цикл по слайсу
	// NB: внутри цикла мы также, как и внутри функции printInfoAboutSpecies выше
	// мы можем только вызывать описанные в сигнатуре интефейса методы, больше ничего
	for _, species := range listOfSpecies {
		info := species.GetSpeciesInfo()
		if info == "" {
			fmt.Println("[ERROR] cannot print information about species")
			os.Exit(1)
		}

		fmt.Println(info)
	}
}
