package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type hangman struct {
	wordsToGuess []string
	randomIndex  int
	choosedWord  string
}

func (h hangman) AddItem(words []string) []string {
	for _, element := range words {
		h.wordsToGuess = append(words, element)
	}
	fmt.Println("Added\n", h.wordsToGuess)
	return h.wordsToGuess
}

func (h hangman) selectRandomIndex() int {
	return rand.Intn(len(h.wordsToGuess))
}

func (h hangman) chooseWordbyRandomIndex() string {
	return h.wordsToGuess[h.selectRandomIndex()]
}

func (h hangman) wordIsNotGuessed(list []string) bool {

	for _, element := range list {
		if element == "-" {
			return true
		}
	}
	return false
}

type HiddenWord struct {
	hiddenw  []string
	splitted []string
}

func (hidden HiddenWord) hideTheWord(h hangman) []string {
	wordLenght := len([]rune(h.choosedWord))

	hiddenWord := make([]string, wordLenght)

	for i, element := range h.choosedWord {
		fmt.Printf("%c\n", element)
		hiddenWord[i] = "-"
	}
	fmt.Printf("Choosed word is: %v\n", hiddenWord)

	return hiddenWord
}

type userError struct {
	message string
}

func (e userError) Error() string {
	return e.message
}

type GameLoop struct {
	result string
}

func (g GameLoop) inputHandler() string {
	notValidInput := true
	var chara string

	for notValidInput {

		fmt.Println("User input: ")
		fmt.Scanln(&chara)
		fmt.Println(len(chara))

		if _, err := strconv.Atoi(chara); err == nil {
			notValidInput = true
			fmt.Printf("%q Looks like a number.\n", chara)

		}
		if !(len(chara) > 1 || len(chara) <= 0) {
			notValidInput = false
			fmt.Println("Correct input.")

		} else {
			notValidInput = true
			fmt.Println("Incorrect input.")
		}

	}
	return chara

}

func (g GameLoop) gameEvent(wordsToGuess []string) string {

	var hang hangman
	var hidden HiddenWord

	// Append words to Hangmans own list
	hang.wordsToGuess = hang.AddItem(wordsToGuess)
	fmt.Printf("Hangman class: %+v", hang)

	hang.choosedWord = hang.chooseWordbyRandomIndex()
	fmt.Printf("Choosed word is: %v\n", hang.choosedWord)

	hidden.hiddenw = hidden.hideTheWord(hang)

	hidden.splitted = strings.Split(hang.choosedWord, "")
	fmt.Println(hidden.splitted)

	myBool := true

	for myBool {

		chara := g.inputHandler()

		for i, element := range hidden.splitted {
			if chara == element {
				hidden.hiddenw[i] = chara
			}
		}

		fmt.Printf("Choosed word is: %v\n", hidden.hiddenw)

		myBool = hang.wordIsNotGuessed(hidden.hiddenw)
	}
	g.result = "YOU REALLY WIN!"

	fmt.Printf("===========================\n")
	fmt.Printf("CONGRATULATION!\n")
	fmt.Printf("Choosed word is: %v\n", hidden.hiddenw)

	return g.result

}

func main() {

	fmt.Println("Welcome to the first HANGMAN project in GO!")
	fmt.Println("-------------------------------------")

	// Some words whoms the player have to guess
	wordsToGuess := []string{"Tesla", "Teller", "Edison", "Einstein"}
	fmt.Printf("Words to guess: %s\n", wordsToGuess)

	// declar Hangman struct
	var game GameLoop

	fmt.Println(game.gameEvent(wordsToGuess))

}
