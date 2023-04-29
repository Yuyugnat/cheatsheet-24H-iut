/*
En terme de nom de fichier et de package en Go, il n'y a qu'une seule règle à respecter :
Quand on lance le programme avec la commande go run,
ce qui est exécuté est la fonction main() du package main.

Tous les fichiers ont une nom de package en haut du fichier.
Le nom du package est le même pour tous les fichiers au même niveau de l'arborescence du projet.

Pour lancer le programme, il faut se placer dans le dossier du projet et lancer la commande :
go run main.go ou go run *.go (pour lancer tous les fichiers du dossier)

Toutes les librairies par défaut de Go doivent être importées pour être utilisées.
Soit avec plusieurs "import math", "import fmt", "import strings", etc.
Soit avec un import groupé "import (math, fmt, strings)" avec chaque librairie sur une ligne.
*/

package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ------------------- //

// types et variables
func typeEtVariables() {
	// principaux types : int, float64, string, bool

	// on peut déclarer une variable avec var (le type est précisé après le nom de la variable)
	var age int
	age = 42 // age de dimitri
	_ = age  // cette ligne permet juste d'éviter l'erreur "unused variable age", inutile en situation réelle

	// on peut évidemment racourcir la déclaration et l'initialisation d'une variable
	var age2 int = 42
	_ = age2

	// dans ce cas là inutile de préciser le type, Go le déduit tout seul
	var age3 = 42
	_ = age3

	// on peut aussi déclarer et initialiser une variable avec :=
	age4 := 42 // équivalent à var age4 int = 42
	_ = age4

	// niveau conversion de type, à supposer que cela soit possible, il suffit d'utiliser le nom du type comme une fonction :
	var age5 float64 = 42
	age6 := int(age5)
	_ = age6

	// les constantes sont déclarées avec const
	const age7 int = 42
	_ = age7
}

// ------------------- //

// opérateurs
func operateurs() {
	// opérateurs arithmétiques : +, -, *, /, % (modulo)
	// opérateurs de comparaison : ==, !=, <, >, <=, >=
	// opérateurs logiques : &&, ||, !

	// /!\ on ne peut pas mélanger les types
	// 42 + 42.0 n'est pas possible
	// il faut convertir les types
	var age int = 42
	var age2 float64 = 42.0
	age3 := age + int(age2)
	_ = age3

	// on peut utiliser les opérateurs de comparaison sur les strings
	// mais attention, ça ne compare pas l'ordre alphabétique
	// ça compare l'ordre des caractères dans la table ASCII
	fmt.Println("a" < "b") // affiche true
	fmt.Println("a" < "B") // affiche false
	fmt.Println("coucou" < "c") // affiche false
	fmt.Println("coucou" < "coucou") // affiche false
	fmt.Println("coucou" < "coucouuuuu") // affiche true


	// la concaténation de strings se fait avec le symbole +
	fmt.Println("coucou" + " " + "toi") // affiche coucou toi
}

// ------------------- //

// affichage sortie standard
func affichage() {
	// le package fmt correspond à la librairie standard de Go pour afficher des messages
	fmt.Println("Hello, World!")        // équivalent de System.out.println("Hello, World!") en Java
	fmt.Print("Hello, World!\n")        // équivalent de System.out.print("Hello, World!\n") en Java
	fmt.Printf("Hello, %s\n", "World!") // équivalent de printf("Hello, %s\n", "World!") en C

	// on peut aussi utiliser le package os pour écrire sur la sortie standard
	fmt.Fprintln(os.Stdout, "Hello, World!") // équivalent de System.out.println("Hello, World!") en Java

	// on peut aussi utiliser le package log pour écrire sur la sortie standard
	log.Println("Hello, World!") // équivalent de System.out.println("Hello, World!") en Java

	// voici les formats de printf les plus utilisés :
	// %d : entier
	// %f : float
	// %s : string
	// %t : bool
	// %v : valeur par défaut
	// %T : type
	// %p : pointeur

	// exemple :
	test := 3.14
	fmt.Printf("La variable test vaut %f\n et son type est %T\n", test, test)

	// tips : fmt.Sprintf est identique à fmt.Printf mais renvoie une string au lieu de l'afficher
	val := fmt.Sprintf("La variable test vaut %f\n et son type est %T\n", test, test)
	fmt.Println(val)
}

// ------------------- //

// pointeurs
func pointeurs() {
	// comme en C, un pointeur vers un type T est de type *T et s'obtient avec &variable
	// un pointeur peut être nil (null en Java)
	// pour accéder à la valeur pointée par un pointeur, on utilise *variable
	text := "Coucou"
	var textPointer *string = &text
	*textPointer = "Coucou2"
	fmt.Println(text) // affiche Coucou2
}

// ------------------- //

// tableaux, slices et maps
func tableauxSlicesMaps() {
	// les tableaux sont déclarés avec [taille]type
	var tableau [5]int
	// on accède à un élément du tableau (lecture ou écriture) avec tableau[index]
	tableau[0] = 1
	tableau[1] = 2
	// on peut rapidement initialiser un tableau avec des valeurs
	tableau2 := [5]int{1, 2, 3, 4, 5}
	// on peut aussi ne pas préciser la taille du tableau et laisser Go la déduire du nombre d'éléments
	tableau3 := [...]int{1, 2, 3, 4, 5}
	_ = tableau2
	_ = tableau3

	// les slices sont des tableaux dynamiques
	// ils sont déclarés avec []type
	var slice []int
	// on peut ajouter des éléments à un slice avec append(slice, element)
	// il faut noter que append n'ajoute pas l'élément à la fin du slice mais renvoie un nouveau slice
	slice = append(slice, 1)
	slice = append(slice, 2)
	// on peut aussi initialiser un slice avec des valeurs
	slice2 := []int{1, 2, 3, 4, 5}
	_ = slice2

	// les maps sont des tableaux associatifs
	// elles sont déclarées avec map[keyType]valueType et initialisées avec make(map[keyType]valueType)
	map1 := make(map[string]int) // il faut utiliser make pour initialiser une map
	// on peut ajouter des éléments à une map avec map[key] = value
	map1["age"] = 42
	map1["taille"] = 180
	// on peut aussi initialiser une map avec des valeurs
	map2 := map[string]int{"age": 42, "taille": 180}
	_ = map2
}

// ------------------- //

// conditions
func conditions() {
	// /!\ pas d'opérateur ternaire en Go

	// les conditions en Go sont les mêmes qu'en C ou Java (sauf qu'il n'y a pas de parenthèses)
	age := 42
	if age > 18 {
		fmt.Println("Majeur")
	} else if age == 18 {
		fmt.Println("Tout juste majeur")
	} else {
		fmt.Println("Mineur")
	}
	// on peut également déclarer une variable dans la condition
	if ageTest := 42; ageTest > 18 {
		fmt.Println("Majeur")
	} else if ageTest == 18 {
		fmt.Println("Tout juste majeur")
	} else {
		fmt.Println("Mineur")
	}

	// on peut utiliser le switch case classique de C ou Java
	// il n'y a pas de break en Go, le break est implicite
	mot := "caca"
	switch mot {
	case "caca":
		fmt.Println("ça sent mauvais")
	case "fleur":
		fmt.Println("ça sent bon")
	default:
		fmt.Println("jsp ce que ça sent")
	}
	// on peut combiner plusieurs cas
	lettre := "a"
	switch lettre {
	case "a", "e", "i", "o", "u", "y":
		fmt.Println("c'est une voyelle")
	default:
		fmt.Println("c'est une consonne")
	}
}

// ------------------- //

// les boucles
func boucles() {
	// TOUTES les boucles en Go s'écrivent toute avec le mot clé for

	// il y a la boucle for classique de C ou Java
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// il y a la boucle while de C ou Java
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// il y a la boucle infinie de C ou Java
	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 10 {
			break // break permet de sortir de la boucle sinon c'est juste une boucle infinie
		}
	}

	// il y a la boucle for each de Java
	// elle s'écrit for index, value := range tableau
	greetings := []string{"salut", "hello", "hola", "ni hao"}
	for index, value := range greetings {
		fmt.Printf("La valeur à l'index %d est %s\n", index, value)
	}

	// il y a la boucle for each de Java pour les maps
	// elle s'écrit de la même façon que pour les tableaux
	ages := map[string]int{"Bob": 42, "Alice": 18, "John": 25}
	for name, age := range ages {
		fmt.Printf("%s a %d ans\n", name, age)
	}

	// si on ne veut pas utiliser l'index ou la valeur, on peut les remplacer par _ afin de les ignorer
	// sinon, Go renvoie une erreur car on ne peut pas déclarer une variable qu'on n'utilise pas
	for _, value := range greetings {
		fmt.Println(value)
	}

	// on peut aussi utiliser continue pour passer à l'itération suivante (valable pour tous les types de boucles)
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// ------------------- //

// fonctions
/*
	les fonctions en Go s'écrivent avec le mot clé func
	le type de retour est précisé après les paramètres
	si la fonction ne retourne rien, on met rien
	normalement elles sont écrites en dehors d'une autre fonction
	mais on peut aussi les écrire dans une autre fonction
*/
// fonction prenant deux entiers en paramètres et retournant un entier
func addition(a int, b int) int {
	return a + b
}

// lorsque le type de plusieurs paramètres à la suite est le même, on peut ne pas répéter le type
func addition2(a, b int) int {
	return a + b
}

func isMajeur(age int) bool {
	if age > 18 {
		return true
	} else {
		return false
	}
}

// en Go, une foncton peut retourner plusieurs valeurs
func division(a, b int) (int, int) {
	return a / b, a % b
}

func testDivision() {
	quotient, reste := division(42, 5)
	fmt.Println(quotient, reste)
}

// fonction dans une variable
func fonctionDansVariable() {
	maFonction := func() {
		fmt.Println("Coucou")
	}
	maFonction()
}

// ------------------- //

// quelques types particuliers
/*
	les bytes sont des entiers non signés sur 8 bits, ils sont souvent utilisés pour représenter des caractères
	les runes sont des entiers non signés sur 32 bits, ils sont souvent utilisés pour représenter des caractères
*/

func exempleByteEtRune() {
	var b byte = 'a'
	fmt.Println(b) // affiche 97
	var r rune = 'a'
	fmt.Println(r) // affiche 97

	// on peut convertir un byte en rune et inversement
	var b2 byte = 'a'
	var r2 rune = rune(b2)
	fmt.Println(r2) // affiche 97
	var r3 rune = 'a'
	var b3 byte = byte(r3)
	fmt.Println(b3) // affiche 97

	// on peut aussi convertir un byte ou un rune en string
	var b4 byte = 'a'
	var r4 rune = 'a'
	var s string = string(b4)
	fmt.Println(s) // affiche a
	var s2 string = string(r4)
	fmt.Println(s2) // affiche a

	// lorsque l'on accède à un caractère d'une chaîne de caractères, on obtient un byte
	var s3 string = "Hello"
	fmt.Println(s3[0])        // affiche 72
	fmt.Printf("%T\n", s3[0]) // affiche uint8
	// pour obtenir un string, on peut utiliser une conversion
	fmt.Println(string(s3[0])) // affiche H
}

// ------------------- //

// maintenant quelques packages utiles
/*
	Le package math contient les fonctions mathématiques de base
	Le package math/rand contient des fonctions pour générer des nombres aléatoires
	Le package time contient des fonctions pour manipuler le temps
	Le package strings contient des fonctions pour manipuler des chaînes de caractères
	Le package strconv contient des fonctions pour convertir des valeurs en chaînes de caractères
	Le package os contient des fonctions pour manipuler le système d'exploitation
*/

// math
func packageMathEtMathRand() {
	// math
	fmt.Println(math.Abs(-42))  // affiche 42
	fmt.Println(math.Max(1, 2)) // affiche 2
	fmt.Println(math.Pow(2, 8)) // affiche 256
	// math/rand
	fmt.Println(rand.Intn(100)) // affiche un nombre aléatoire entre 0 et 99
	fmt.Println(rand.Float64()) // affiche un nombre aléatoire entre 0 et 1
}

// time
func packageTime() {
	now := time.Now()                      // renvoie la date et l'heure actuelle
	time.Sleep(1 * time.Second)            // fait une pause de 1 seconde
	time.Sleep(1 * time.Millisecond)       // fait une pause de 1 milliseconde
	fmt.Println(time.Since(now).Seconds()) // affiche le temps écoulé depuis now en secondes
}

// strings et strconv
func packageStringsEtStrconv() {
	// strings
	fmt.Println(strings.ToUpper("coucou"))                      // affiche COUCOU
	fmt.Println(strings.ToLower("COUCOU"))                      // affiche coucou
	fmt.Println(strings.Contains("coucou", "ou"))               // affiche true
	fmt.Println(strings.Index("coucou", "ou"))                  // affiche 1
	fmt.Println(strings.Replace("coucou", "ou", "a", -1))       // affiche caucau, le paramètre number permet de préciser le nombre de remplacements à faire, -1 signifie qu'on remplace tout
	fmt.Println(strings.Split("coucou", "o"))                   // affiche [c c oucou], un slice de chaînes de caractères
	fmt.Println(strings.Join([]string{"c", "c", "oucou"}, "-")) // affiche c-c-oucou
	// strconv
	fmt.Println(strconv.Itoa(42))   // affiche 42
	fmt.Println(strconv.Atoi("42")) // affiche 42
}

// os
func packageOs() {
	fmt.Println(os.Getenv("PATH")) // affiche la variable d'environnement PATH
	fmt.Println(os.Getpid())       // affiche l'identifiant du processus
	fmt.Println(os.Getppid())      // affiche l'identifiant du processus parent
	fmt.Println(os.Getwd())        // affiche le répertoire de travail
	// il y a aussi des fonctions pour manipuler les fichiers
	// la fonction OpenFile permet d'ouvrir un fichier, elle renvoie deux valeurs :
	// un pointeur vers un objet de type File et une erreur, si l'ouverture s'est bien passée, l'erreur vaut nil
	file, err := os.OpenFile("fichier.txt", os.O_RDONLY, 0666) // ouvre le fichier en lecture seule
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // le mot clé defer permet d'exécuter une fonction à la fin de la fonction courante

	// la fonction Read permet de lire des données dans un fichier
	// elle prend en paramètre un slice de bytes et renvoie le nombre de bytes lus et une erreur
	// si la lecture s'est bien passée, l'erreur vaut nil
	data := make([]byte, 100) // on crée un slice de 100 bytes pour stocker les données lues
	nb, err := file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(nb)         // affiche le nombre de bytes lus
	fmt.Println(string(data)) // affiche les données lues
}

// ------------------- //

func main() {
	operateurs()
}
