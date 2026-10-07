package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)

	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier d'entrée:", err)
		return
	}
	text := string(data)
	words := strings.Fields(text)

	fmt.Println("Mots :", words)
	for i, word := range words {
		if word == "(hex)" && i > 0 {
			fmt.Println("J'ai trouvé le mot hexadécimal à la position ", i)

			previousword := words[i-1]
			fmt.Println("Mot précédent (hex):", previousword)

			number, err := strconv.ParseInt(previousword, 16, 64)

			if err != nil {
				fmt.Println("Erreur lors de la conversion du mot précédent en entier:", err)
				continue
			}
			fmt.Println("valeur décimale:", number)

			words[i-1] = strconv.FormatInt(number, 10)
			words[i] = " "
		}

	}
	result := strings.Join(words, " ")
	fmt.Println("Résultat :", result)

	err = os.WriteFile(outputFile, []byte(result), 0644)

	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier de sortie:", err)
		return
	}

	fmt.Println("Fichier d'entrée:", inputFile)
	fmt.Println("Fichier de sortie:", outputFile)
	fmt.Println("Contenu:", string(data))

}
