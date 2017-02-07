/*
December 5th, 2016 /r/DailyProgrammer challenge
Difficulty: Easy
Bonus 1: Complete
Bonus 2: Incomplete
Bonus 3: Adjusted (built a point system into the game, not a points system into word list)
*/

package main

import (
	"fmt"
	//"bufio"
	//"os"
	"math/rand"
)

func main(){
	//reader := bufio.NewReader(os.Stdin)
	tiles := [7]rune{}
	alphabet := "abcdefghijklmnopqrstuvwxyz?"

	for i := 0; i < 7; i++ {
		tiles[i] = rune(alphabet[rand.Intn(len(alphabet)-1)])
		fmt.Println(string(tiles[i]))
	}


	//fmt.Println("Your tiles are ", tiles, "\nEnter your word:")
	//word, _ := reader.ReadString('\n')
	word := "xvl"
	totalPoints := checkPoints(word)

	possibleWord := scrabble(tiles, word)
	fmt.Println("The statement your word works with your tiles is:", possibleWord) //true or false
	if possibleWord{
		fmt.Println("Your total points are:", totalPoints)
	}
}

func scrabble(tiles [7]rune, word string) bool{
	letterWorks := false
	succesfulWord := true
	blankTileAccepted := false


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
			//check to see if there is a blank tile in the tiles array 
			for y, tile := range tiles {
				if string(tile) == "?"{
					tiles[y] = '*'
					blankTileAccepted = true 
					break
				}
			}
			if !blankTileAccepted{
				succesfulWord = false
			}
		}
		blankTileAccepted = false
		letterWorks = false
	}

	return succesfulWord
}

func checkPoints(word string) int {
	alphabet := "abcdefghijklmnopqrstuvwxyz?"
	letterPoints := []int{1, 3, 3, 2, 1, 4, 2, 4, 1, 8, 5, 1, 3, 1, 1, 3, 
						10, 1, 1, 1, 1, 4, 4, 8, 4, 10} 

	points := 0

	for _, wordElement := range word{
		for y, letter := range alphabet{
			if string(wordElement) == string(letter){
				points += letterPoints[y]
				break
			}
		}
	} 

	return points
} 