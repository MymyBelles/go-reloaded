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
		if word == "(bin)" && i > 0 {
			fmt.Println("j'ai trouvé (bin) à la position : ", i)

			previousWord := words[i-1]
			fmt.Println("nombre binaire :", previousWord)

			number, err := strconv.ParseInt(previousWord, 2, 64)
			if err != nil {
				fmt.Println("Erreur de conversion binaire :", err)
				continue
			}
			fmt.Println("Valeur décimale :", number)

			words[i-1] = strconv.FormatInt(number, 10)
			words[i] = ""
		}
		if word == "(up)" && i > 0 {
			previousWord := words[i-1]

			words[i-1] = strings.ToUpper(previousWord)
			words[i] = ""
		}
		if word == "(low)" && i > 0 {
			previousWord := words[i-1]

			words[i-1] = strings.ToLower(previousWord)
			words[i] = ""
		}
		if word == "(cap)" && i > 0 {
			previousWord := words[i-1]

			lowerWord := strings.ToLower(previousWord)

			words[i-1] = strings.ToUpper(lowerWord[:1]) + lowerWord[1:]
			words[i] = ""
		}
		if word == "(up," && i+1 < len(words) {
			numberText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(numberText)

			if err != nil {
				fmt.Println("Erreur : nombre de mots invalide :", err)
				continue
			}

			for j := 1; j <= count; j++ {
				position := i - j

				if position >= 0 {
					words[position] = strings.ToUpper(words[position])
				}
			}

			words[i] = ""
			words[i+1] = ""
		}

		if word == "(low," && i+1 < len(words) {
			numberText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(numberText)

			if err != nil {
				fmt.Println("Erreur : nombre de mots invalide :", err)
				continue
			}

			for j := 1; j <= count; j++ {
				position := i - j

				if position >= 0 {
					words[position] = strings.ToLower(words[position])
				}
			}

			words[i] = ""
			words[i+1] = ""
		}

		if word == "(cap," && i+1 < len(words) {
			numberText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(numberText)

			if err != nil {
				fmt.Println("Erreur : nombre de mots invalide :", err)
				continue
			}

			for j := 1; j <= count; j++ {
				position := i - j

				if position >= 0 {
					words[position] = strings.Title(strings.ToLower(words[position]))
				}
			}

			words[i] = ""
			words[i+1] = ""
		}
	}
	//result := strings.Join(words, " ")

	var cleanedWords []string

	for _, word := range words {
		if word != "" {
			cleanedWords = append(cleanedWords, word)
		}
	}

	result := strings.Join(cleanedWords, " ")

	// Corriger les espaces avant les virgules
	result = strings.ReplaceAll(result, " ,", ",")

	// Corriger les espaces avant les points
	result = strings.ReplaceAll(result, " .", ".")

	// Corriger les espaces avant les points d'exclamation
	result = strings.ReplaceAll(result, " !", "!")

	// Corriger les espaces avant les points d'interrogation
	result = strings.ReplaceAll(result, " ?", "?")

	// Corriger les espaces avant les deux-points
	result = strings.ReplaceAll(result, " :", ":")

	// Corriger les espaces avant les points-virgules
	result = strings.ReplaceAll(result, " ;", ";")

	// Corriger les espaces entre les apostrophes
	parts := strings.Split(result, "'")

	for i := 1; i+1 < len(parts); i += 2 {
		parts[i] = strings.TrimSpace(parts[i])
	}
	result = strings.Join(parts, "'")

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
