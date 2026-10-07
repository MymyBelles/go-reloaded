package main

import (
	"fmt"
	"os"
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
	err = os.WriteFile(outputFile, data, 0644)

	if err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier de sortie:", err)
		return
	}

	fmt.Println("Fichier d'entrée:", inputFile)
	fmt.Println("Fichier de sortie:", outputFile)
	fmt.Println("Contenu:", string(data))
}
