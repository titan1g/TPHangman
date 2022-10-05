package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() { /*

		FirstWord()
		LastWord()
		AfterWord()
	*/
	file, _ := os.Open("Hangman.txt")     // Permet d'accéder au fichier Hangman.txt
	fileScanner := bufio.NewScanner(file) // Permet de scanner le fichier Hangman.txt et donc de lire l'intégrité du fichier
	var array []string                    // Création d'un tableau de string afin d'accéder plus simplement aux mots du fichier Hangman.txt

	for fileScanner.Scan() { // Boucle qui permet de scanner le fichier Hangman.txt

		array = append(array, fileScanner.Text()) // Ajoute les mots du fichier Hangman.txt dans le tableau "array"

	}
	file.Close()
	//fmt.Println(array)    // Affiche le tableau "array" qui contient tous les mots du fichier Hangman.txt
	fmt.Println(array[0]) // Affiche le premier mot du fichier Hangman.txt
	fmt.Println(array[len(array)-1])

	var a string // Affiche le troisième mot du fichier Hangman.txt
	for i := 0; i < len(array); i++ {

		if array[i] == "Before" {
			a = array[i+1]
		}
	}

	aInt, _ := strconv.Atoi(a)
	fmt.Println(array[aInt])

}

// Vu avec maxime, le troisième mot ne fonctionne pas, je ne comprends pas pourquoi.
