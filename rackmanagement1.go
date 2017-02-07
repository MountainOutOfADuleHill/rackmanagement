/*
December 5th, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
Bonus 1:
Bonus 2:
Bonus 3:
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"math/rand"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	tiles := [7]rune{}
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < 7; i++ {
		tiles[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
	}

	fmt.Println("Your tiles are ", tiles, "\nEnter your word:")
	word, _ := reader.ReadString('\n')

	possibleWord := scrabble(tiles, word)
	fmt.Println("The statement your word works with your tiles is :", possibleWord)

}

func scrabble(tiles [7]rune, word string) bool{
	var letterWorks bool 
	succesfulWord := true

	for _, letter := range word {
		//loop through all the letters in the word entered by the user
		for x, tile := range tiles {
			//loop through all the tiles to see if there is a match with the letter entered by the user
			if string(letter) == string(tile){
				tiles[x] = '*' //so that a tile can't be used more than once
				letterWorks = true
				break //once we have found a match for the letter we don't need to keep looking at tiles
			}
		}

		if !letterWorks {
			succesfulWord = false
		}

		letterWorks = false
	}

	return succesfulWord
}