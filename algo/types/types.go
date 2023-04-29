package types

import "fmt"

// ------------------- //

// déclaration de type en Go

// un type en Go peut être :
// - un type de base (int, string, float, bool, ...)
// - un tableau (array)
// - slice
// - une fonction (func)
// - une interface (interface)
// - un pointeur (pointer)
// - une structure (struct)

// les plus importants sont les alias vers les types de base, les structures et les interfaces

// ------------------- //

// type de base

// on peut par exemple définir un alias pour un type de base
type MyInt int

// ainsi on peut utiliser MyInt comme un type de base
var i MyInt = 10

func additionMyInt(a, b MyInt) MyInt {
	return a + b
}

func testMyInt() {
	var a, b MyInt = 10, 20
	var c MyInt = additionMyInt(a, b)
	fmt.Println(c)

	// on peut utiliser un int comme un MyInt
	var d MyInt
	var e int = 10
	d = MyInt(e)
	fmt.Println(d)

	// on peut utiliser un MyInt comme un int
	var f int = int(d)
	fmt.Println(f)
}

// ------------------- //

// struct

// un type struct permet d'avoir un type contenant plusieurs champs de différents types

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func testStruct() {
	// on peut déclarer une struct en une seule ligne, mais il faut respecter l'ordre des champs
	var florimiche Person = Person{"Florimond", "Carron", 20}
	fmt.Println(florimiche)

	// on peut aussi déclarer une struct en précisant les champs, cela permet de ne pas respecter l'ordre des champs
	// mais aussi de ne pas préciser tous les champs
	var adam Person = Person{LastName: "Henchiri", FirstName: "Adam"}
	fmt.Println(adam)

	// en génaral, on déclare une struct sur plusieurs lignes
	var dimitri Person = Person{
		FirstName: "Dimitri",
		LastName:  "Copley",
		Age:       67, // bizarrerie en Go : la dernière virgule est obligatoire même si on ne met pas de champ après
	}
	fmt.Println(dimitri)
}

// ------------------- //

// interface

// une interface permet de définir un type qui contient des méthodes uniquement
// une interface ne peut pas contenir de champs

type Animal interface {
	SayHello() string
	Eat(food string) string
	IsDie() bool
}

// ------------------- //

// exemple utilisations

// Les alias vers les types de base sont notamment utilisés pour définir des enums à
// l'aide de constantes et du mot clé iota

type Color int

// ici le mot clé iota permet de définir des constantes qui vont s'incrémenter automatiquement
const (
	RED Color = iota // vaudra 0
	GREEN 		  // vaudra 1
	BLUE		  // vaudra 2
)
// ainsi on obtient un vériable enum en Go car à l'utilisation on ne se soucie pas de la valeur sous-jacente

type Fleur struct {
	Color Color
}

func testEnum() {
	var rose Fleur = Fleur{Color: RED}
	fmt.Println(rose)
}

// on peut associer des méthodes à n'importe quel type que nous avons créé

func (c Color) String() string {
	switch c {
	case RED:
		return "red"
	case GREEN:
		return "green"
	case BLUE:
		return "blue"
	default:
		return "unknown"
	}
}

func (i MyInt) IsPair() bool {
	return i%2 == 0
}

func (p Person) IsMajeur() bool {
	return p.Age >= 18
}

func test() {
	var a MyInt = 10
	fmt.Println(a.IsPair())

	var b Person = Person{Age: 20}
	fmt.Println(b.IsMajeur())
}

// pour qu'un type implémente une interface, il faut que le type définisse toutes les méthodes de l'interface
// et c'est tout, il n'y a pas de mot clé à mettre devant le type pour dire qu'il implémente l'interface

func (p Person) SayHello() string {
	return fmt.Sprintf("Hello, my name is %s %s", p.FirstName, p.LastName)
}

func (p Person) Eat(food string) string {
	return fmt.Sprintf("I eat %s", food)
}

func (p Person) IsDie() bool {
	return p.Age > 100
}

func testInterface() {
	var unMecBien Animal = Person{
		FirstName: "Charles",
		LastName:  "Darwin",
		Age:       148,
	}
	fmt.Println(unMecBien.SayHello())
	fmt.Println(unMecBien.Eat("a banana"))
	fmt.Println(unMecBien.IsDie())
}