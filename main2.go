package main

import "fmt"

// Creature представляет существо с именем и приветствием.
type Creature struct {
	Name     string
	Greeting string
}

// Greet выводит приветственное сообщение и возвращает Creature.
func (c Creature) Greet() Creature {
	fmt.Printf("%s говорит: %s!\n", c.Name, c.Greeting)
	return c
}

// SayGoodbye выводит прощальное сообщение, используя имя существа.
func (c Creature) SayGoodbye(name string) {
	fmt.Println("Прощай", name, "!")
}

func main() {
	// Создание экземпляра Creature с именем sammy
	sammy := Creature{
		Name:     "Сэмми",
		Greeting: "Привет!",
	}

	// Вызов методов Greet и SayGoodbye для sammy
	sammy.Greet().SayGoodbye("гоферы")
}
