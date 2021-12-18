package main

import (
	"fmt"
	"github.com/KartikeyaMalimath/go-mini-games/tictactoe"
)

func main() {
	var input int
	fmt.Println("===================================")
	fmt.Println("Welcome to KM's Go-lang Mini Games")
	fmt.Printf("===================================\n\n")
	fmt.Printf("Select the game you wanna Play!\n 1.) Tic Tac Toe\n2.) Exit\n------------------------\nEnter your choice: ")
	fmt.Scan(&input)
	switch(input){
	case 1: TicTacToe()
	case 2: break
	}

}
