package main

import (
	"fmt"
	"os"
)

// пример #2: полноценные типы на основе существующих типов

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
	// создаем переменные-экземпляры каждого типа.
	// т.к. все типы - это "псевдонимы" на основе строкового типа,
	// то для создания экземпляров требуется передать строки
	var (
		d = Psovie("собака")
		c = Koshachiyi("рысь")
		w = Polosatikovie("синий кит")
		h = Loshadinie("домашняя лошадь")
	)

	// тип Psovie имеет нужный (описанный в сигнатуре интерфейса) метод GetSpeciesInfo,
	// метод имеет верную сигнатуру, мы можем передать переменную типа Psovie
	// в качестве переменной типа Species
	if err := printInfoAboutSpecies(d); err != nil {
		fmt.Println("[ERROR] cannot print information about species")
		os.Exit(1)
	}

	// тип Koshachiyi имеет нужный (описанный в сигнатуре интерфейса) метод GetSpeciesInfo,
	// метод имеет верную сигнатуру, мы можем передать переменную типа Koshachiyi
	// в качестве переменной типа Species
	if err := printInfoAboutSpecies(c); err != nil {
		fmt.Println("[ERROR] cannot print information about species")
		os.Exit(1)
	}

	// тип Polosatikovie не имеет нужного метода, поэтому мы не можем передать экземпляр
	// данного типа в качестве Species (компиляция упадет на этом месте)
	if err := printInfoAboutSpecies(w); err != nil {
		fmt.Println("[ERROR] cannot print information about species")
		os.Exit(1)
	}

	// тип Loshadinie имеет метод с нужным именем, но его сигнатура отличается  не имеет нужного метода, поэтому мы не можем передать экземпляр
	// от того, что описано в сигнатуре интерфейсного типа Species, поэтому
	// мы не можем передать переменную типа Loshadinie в Species (компиляция упадет на этом месте)
	if err := printInfoAboutSpecies(h); err != nil {
		fmt.Println("[ERROR] cannot print information about species")
		os.Exit(1)
	}

	// ожидаемое поведение:
	// вначале компиляция падает на последних двух вызоваз (ошибки компилятор понятные)
	// далее, если закомментировать/удалить код вызова функции в 3 и 4 раз (по очереди для понимания),
	// и соответствующие переменные, чтобы не оставалось неиспользуемых переменных в коде,
	// то компиляция и запуск начинают работать
}
