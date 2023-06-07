package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var hmArr = [7]string{
	" +------+\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"      =====\n",
	" +------+\n" +
		" 0      |\n" +
		"        |\n" +
		"        |\n" +
		"        |\n" +
		"      =====\n",
	" +------+\n" +
		" 0      |\n" +
		"/       |\n" +
		"        |\n" +
		"        |\n" +
		"      =====\n",
	" +------+\n" +
		" 0      |\n" +
		"/ \\     |\n" +
		"        |\n" +
		"        |\n" +
		"      =====\n",
	" +------+\n" +
		" 0      |\n" +
		"/|\\     |\n" +
		"        |\n" +
		"        |\n" +
		"      =====\n",
	" +------+\n" +
		" 0      |\n" +
		"/|\\     |\n" +
		"/       |\n" +
		"        |\n" +
		"      =====\n",
	" +------+\n" +
		" 0      |\n" +
		"/|\\     |\n" +
		"/ \\     |\n" +
		"        |\n" +
		"      =====\n",
}

var wordArr = [7]string{
	"BANDERA", "CHALUPA", "BARRIL", "HARINA", "PALMA", "LUNA", "SIRENA",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func generateRandomWord() {
	seedSecs := time.Now().Unix()
	rand := rand.New(rand.NewSource(seedSecs))
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
}

func areEmptyLetters() bool {
	for _, v := range correctLetters {
		if v == "" {
			return true
		}
	}
	return false
}

func main() {
	generateRandomWord()

	for {
		showDashboard()
		guess := getUserInput()
		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
			if !areEmptyLetters() {
				fmt.Printf("YOU WIN, THE SECRET WORD IS: %s\n", randWord)
				break
			}
		} else {
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)

			if len(wrongGuesses) > 5 {
				showDashboard()
				fmt.Printf("YOU LOOSE, THE SECRET WORD IS: %s\n", randWord)
				break
			}
		}
		fmt.Println()
	}

}

func showDashboard() {
	fmt.Print(hmArr[len(wrongGuesses)])

	fmt.Print("SECRET WORD: ")
	for _, l := range correctLetters {
		if l == "" {
			print("_")
		} else {
			print(l)
		}
	}
	fmt.Println()
	fmt.Print("WRONG LETTERS: ")
	for _, v := range wrongGuesses {
		fmt.Print(v + " ")
	}
	fmt.Println()
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	var guess string
	var err error

	for {
		fmt.Print("GUESS A LETTER: ")

		guess, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)

		isLetter := regexp.MustCompile(`[a-zA-Z]+$`).MatchString
		if guess == "" {
			println("PLEASE ENTER A LETTER")
			continue
		}
		if !isLetter(guess) {
			println("PLEASE ENTER A LETTER")
			continue
		}
		if len(guess) > 1 {
			println("PLEASE ENTER ONLY ONE LETTER")
			continue
		}
		if strings.Contains(guessedLetters, guess) {
			println("PLEASE ENTER A LETTER YOU HAVEN'T GUESSED")
			continue
		}
		break
	}

	return guess
}

func getAllIndexes(str, subStr string) (indices []int) {

	if len(str) == 0 || len(subStr) == 0 {
		return
	}

	offset := 0
	for {
		i := strings.Index(str[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}

}

func updateCorrectLetters(letter string) {
	indices := getAllIndexes(randWord, letter)
	for _, v := range indices {
		correctLetters[v] = letter
	}
}

// func showBoard() {
// 	fmt.Println(hmArr[len(wrongGuesses)])
// 	fmt.Println("Secret word: ")
// 	for _, v := range correctLetters {
// 		if v == "" {
// 			fmt.Print("_")
// 		} else {
// 			fmt.Print(v)
// 		}
// 	}

// 	fmt.Print("\nIncorrect guesses: ")
// 	if len(wrongGuesses) > 0 {
// 		for _, v := range wrongGuesses {
// 			fmt.Print(v + " ")
// 		}
// 	}
// 	fmt.Println()

// }

// func getUserLetter() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	var guess string

// 	for true {

// 		fmt.Print("\nGuess a letter: ")
// 		guess, err := reader.ReadString('\n')
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		guess = strings.ToUpper(guess)
// 		guess = strings.TrimSpace(guess)

// 		var isLetter = regexp.MustCompile(`[a-zA-Z]+$`).MatchString

// 		if len(guess) > 1 {
// 			fmt.Println("Please enter Only one letter")
// 		} else if !isLetter(guess) {
// 			fmt.Println("Please enter a letter")
// 		} else if strings.Contains(guessedLetters, guess) {
// 			fmt.Println("Please enter a letter you haven't guessed")
// 		} else {
// 			return guess
// 		}

// 	}
// 	return guess
// }

// func getAllIndexes(theStr, substr string) (indices []int) {

// 	if len(substr) == 0 || len(theStr) == 0 {
// 		return indices
// 	}

// 	offset := 0

// 	for {
// 		i := strings.Index(theStr[offset:], substr)
// 		if i == -1 {
// 			return indices
// 		} else {
// 			offset += i
// 			indices = append(indices, offset)
// 			offset += len(substr)
// 		}
// 	}

// }

// func updateCorrectLetters(letter string) {
// 	indexMatches := getAllIndexes(randWord, letter)

// 	for _, v := range indexMatches {
// 		correctLetters[v] = letter
// 	}
// }

// func slicesHasEmptys(theSlice []string) bool {
// 	for _, v := range theSlice {
// 		if len(v) == 0 {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// 	fmt.Println(getRandWord())

// 	for true {
// 		showBoard()
// 		guess := getUserLetter()
// 		if strings.Contains(randWord, guess) {
// 			updateCorrectLetters(guess)
// 			if slicesHasEmptys(correctLetters) {
// 				fmt.Println("More letters to guess")
// 			} else {
// 				fmt.Println("Yes the secret word is: ", randWord)
// 				break
// 			}
// 		} else {
// 			guessedLetters += guess
// 			wrongGuesses = append(wrongGuesses, guess)

// 			if len(wrongGuesses) >= 6 {
// 				fmt.Println("Sorry you are dead. The random word is: ", randWord)
// 				break
// 			}
// 		}

// 	}

// }
